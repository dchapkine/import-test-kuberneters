/*
Copyright 2015 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package namespace

import (
	"fmt"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client/cache"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/watch"
	"github.com/golang/glog"
)

// NamespaceManager is responsible for performing actions dependent upon a namespace phase
type NamespaceManager struct {
	kubeClient client.Interface
	store      cache.Store
	syncTime   <-chan time.Time

	// To allow injection for testing.
	syncHandler func(namespace api.Namespace) error
}

// NewNamespaceManager creates a new NamespaceManager
func NewNamespaceManager(kubeClient client.Interface) *NamespaceManager {
	store := cache.NewStore(cache.MetaNamespaceKeyFunc)
	//fieldSelector, _ := labels.ParseSelector("status.phase=Terminating")
	fieldSelector := labels.Everything()
	reflector := cache.NewReflector(
		&cache.ListWatch{
			ListFunc: func() (runtime.Object, error) {
				return kubeClient.Namespaces().List(labels.Everything(), fieldSelector)
			},
			WatchFunc: func(resourceVersion string) (watch.Interface, error) {
				return kubeClient.Namespaces().Watch(labels.Everything(), fieldSelector, resourceVersion)
			},
		},
		&api.Namespace{},
		store,
		0,
	)
	reflector.Run()
	nm := &NamespaceManager{
		kubeClient: kubeClient,
		store:      store,
	}
	// set the synchronization handler
	nm.syncHandler = nm.terminateNamespace
	return nm
}

// Run begins syncing at the specified period interval
func (nm *NamespaceManager) Run(period time.Duration) {
	nm.syncTime = time.Tick(period)
	go util.Forever(func() { nm.synchronize() }, period)
}

// Iterate over the each namespace that is in terminating phase and perform necessary clean-up
func (nm *NamespaceManager) synchronize() {
	namespaceObjs := nm.store.List()
	wg := sync.WaitGroup{}
	wg.Add(len(namespaceObjs))
	for ix := range namespaceObjs {
		go func(ix int) {
			defer wg.Done()
			namespace := namespaceObjs[ix].(*api.Namespace)
			glog.V(4).Infof("periodic sync of namespace: %v", namespace.Name)
			err := nm.syncHandler(*namespace)
			if err != nil {
				glog.Errorf("Error synchronizing: %v", err)
			}
		}(ix)
	}
	wg.Wait()
}

// finalized returns true if the spec.finalizers is equal to the status.finalizers
func finalized(namespace api.Namespace) bool {
	specSet := util.NewStringSet()
	for i := range namespace.Spec.Finalizers {
		specSet.Insert(string(namespace.Spec.Finalizers[i]))
	}
	statusSet := util.NewStringSet()
	for i := range namespace.Status.Finalizers {
		statusSet.Insert(string(namespace.Status.Finalizers[i]))
	}
	return statusSet.HasAll(specSet.List()...)
}

// finalize will finalize the namespace for kubernetes
func finalize(kubeClient client.Interface, namespace api.Namespace) (*api.Namespace, error) {
	namespaceFinalize := api.Namespace{
		ObjectMeta: api.ObjectMeta{
			Name:            namespace.Name,
			ResourceVersion: namespace.ResourceVersion,
		},
		Status: api.NamespaceStatus{},
	}
	namespaceFinalize.Status.Finalizers = make([]api.FinalizerName, len(namespace.Status.Finalizers), len(namespace.Status.Finalizers)+1)
	copy(namespaceFinalize.Status.Finalizers, namespace.Status.Finalizers)
	namespaceFinalize.Status.Finalizers = append(namespaceFinalize.Status.Finalizers, api.FinalizerKubernetes)

	return kubeClient.Namespaces().Finalize(&namespaceFinalize)
}

// deleteAllContent will delete all content known to the system in a namespace
func deleteAllContent(kubeClient client.Interface, namespace string) (err error) {
	err = deleteServices(kubeClient, namespace)
	if err != nil {
		return err
	}
	err = deleteReplicationControllers(kubeClient, namespace)
	if err != nil {
		return err
	}
	err = deletePods(kubeClient, namespace)
	if err != nil {
		return err
	}
	err = deleteSecrets(kubeClient, namespace)
	if err != nil {
		return err
	}
	err = deleteLimitRanges(kubeClient, namespace)
	if err != nil {
		return err
	}
	err = deleteResourceQuotas(kubeClient, namespace)
	if err != nil {
		return err
	}
	err = deleteEvents(kubeClient, namespace)
	if err != nil {
		return err
	}
	return nil
}

// terminateNamespace attempts to do graceful termination of a namespace
func (nm *NamespaceManager) terminateNamespace(namespace api.Namespace) (err error) {

	// remove this when https://github.com/GoogleCloudPlatform/kubernetes/pull/5389 merges since our
	// watch will already be filtered
	if namespace.Status.Phase != api.NamespaceTerminating {
		return fmt.Errorf("Namespace %v is not in the terminating phase.", namespace.Name)
	}

	// if the namespace is already finalized, delete it
	if finalized(namespace) {
		err = nm.kubeClient.Namespaces().Delete(namespace.Name)
		return err
	}

	// there may still be content for us to remove
	err = deleteAllContent(nm.kubeClient, namespace.Name)
	if err != nil {
		return err
	}

	// we have removed content, so mark it finalized by us
	result, err := finalize(nm.kubeClient, namespace)
	if err != nil {
		return err
	}

	// now check if all finalizers have reported that we delete now
	if finalized(*result) {
		err = nm.kubeClient.Namespaces().Delete(namespace.Name)
		return err
	}

	return nil
}

func deleteLimitRanges(kubeClient client.Interface, ns string) error {
	items, err := kubeClient.LimitRanges(ns).List(labels.Everything())
	if err != nil {
		return err
	}
	for i := range items.Items {
		err := kubeClient.LimitRanges(ns).Delete(items.Items[i].Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteResourceQuotas(kubeClient client.Interface, ns string) error {
	resourceQuotas, err := kubeClient.ResourceQuotas(ns).List(labels.Everything())
	if err != nil {
		return err
	}
	for i := range resourceQuotas.Items {
		err := kubeClient.ResourceQuotas(ns).Delete(resourceQuotas.Items[i].Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteServices(kubeClient client.Interface, ns string) error {
	items, err := kubeClient.Services(ns).List(labels.Everything())
	if err != nil {
		return err
	}
	for i := range items.Items {
		err := kubeClient.Services(ns).Delete(items.Items[i].Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteReplicationControllers(kubeClient client.Interface, ns string) error {
	items, err := kubeClient.ReplicationControllers(ns).List(labels.Everything())
	if err != nil {
		return err
	}
	for i := range items.Items {
		err := kubeClient.ReplicationControllers(ns).Delete(items.Items[i].Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func deletePods(kubeClient client.Interface, ns string) error {
	items, err := kubeClient.Pods(ns).List(labels.Everything())
	if err != nil {
		return err
	}
	for i := range items.Items {
		err := kubeClient.Pods(ns).Delete(items.Items[i].Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteEvents(kubeClient client.Interface, ns string) error {
	items, err := kubeClient.Events(ns).List(labels.Everything(), labels.Everything())
	if err != nil {
		return err
	}
	for i := range items.Items {
		err := kubeClient.Events(ns).Delete(items.Items[i].Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteSecrets(kubeClient client.Interface, ns string) error {
	items, err := kubeClient.Secrets(ns).List(labels.Everything(), labels.Everything())
	if err != nil {
		return err
	}
	for i := range items.Items {
		err := kubeClient.Secrets(ns).Delete(items.Items[i].Name)
		if err != nil {
			return err
		}
	}
	return nil
}
