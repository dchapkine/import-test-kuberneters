/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/autoscaling/v2alpha1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1beta1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/rbac/v1beta1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	api "k8s.io/kubernetes/pkg/api"
	v1 "k8s.io/kubernetes/pkg/api/v1"
)

// ReplicationControllerLister helps list ReplicationControllers.
type ReplicationControllerLister interface {
	// List lists all ReplicationControllers in the indexer.
	List(selector labels.Selector) (ret []*v1.ReplicationController, err error)
	// ReplicationControllers returns an object that can list and get ReplicationControllers.
	ReplicationControllers(namespace string) ReplicationControllerNamespaceLister
	ReplicationControllerListerExpansion
}

// replicationControllerLister implements the ReplicationControllerLister interface.
type replicationControllerLister struct {
	indexer cache.Indexer
}

// NewReplicationControllerLister returns a new ReplicationControllerLister.
func NewReplicationControllerLister(indexer cache.Indexer) ReplicationControllerLister {
	return &replicationControllerLister{indexer: indexer}
}

// List lists all ReplicationControllers in the indexer.
func (s *replicationControllerLister) List(selector labels.Selector) (ret []*v1.ReplicationController, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ReplicationController))
	})
	return ret, err
}

// ReplicationControllers returns an object that can list and get ReplicationControllers.
func (s *replicationControllerLister) ReplicationControllers(namespace string) ReplicationControllerNamespaceLister {
	return replicationControllerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReplicationControllerNamespaceLister helps list and get ReplicationControllers.
type ReplicationControllerNamespaceLister interface {
	// List lists all ReplicationControllers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ReplicationController, err error)
	// Get retrieves the ReplicationController from the indexer for a given namespace and name.
	Get(name string) (*v1.ReplicationController, error)
	ReplicationControllerNamespaceListerExpansion
}

// replicationControllerNamespaceLister implements the ReplicationControllerNamespaceLister
// interface.
type replicationControllerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReplicationControllers in the indexer for a given namespace.
func (s replicationControllerNamespaceLister) List(selector labels.Selector) (ret []*v1.ReplicationController, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ReplicationController))
	})
	return ret, err
}

// Get retrieves the ReplicationController from the indexer for a given namespace and name.
func (s replicationControllerNamespaceLister) Get(name string) (*v1.ReplicationController, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("replicationcontroller"), name)
	}
	return obj.(*v1.ReplicationController), nil
}
