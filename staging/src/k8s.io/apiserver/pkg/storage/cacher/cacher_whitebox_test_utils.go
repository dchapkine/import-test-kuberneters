/*
Copyright 2016 The Kubernetes Authors.

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

package cacher

import (
	"context"
	"fmt"
	"sync"

	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/apis/example"
	examplev1 "k8s.io/apiserver/pkg/apis/example/v1"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/utils/clock"
)

var (
	scheme = runtime.NewScheme()
	codecs = serializer.NewCodecFactory(scheme)
)

func InitTestSchema() {
	metav1.AddToGroupVersion(scheme, metav1.SchemeGroupVersion)
	utilruntime.Must(example.AddToScheme(scheme))
	utilruntime.Must(examplev1.AddToScheme(scheme))
}

func newTestCacher(s storage.Interface) (*Cacher, storage.Versioner, error) {
	prefix := "pods"
	config := Config{
		Storage:        s,
		Versioner:      storage.APIObjectVersioner{},
		GroupResource:  schema.GroupResource{Resource: "pods"},
		ResourcePrefix: prefix,
		KeyFunc:        func(obj runtime.Object) (string, error) { return storage.NamespaceKeyFunc(prefix, obj) },
		GetAttrsFunc: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			pod, ok := obj.(*example.Pod)
			if !ok {
				return storage.DefaultNamespaceScopedAttr(obj)
			}
			labelsSet, fieldsSet, err := storage.DefaultNamespaceScopedAttr(obj)
			if err != nil {
				return nil, nil, err
			}
			fieldsSet["spec.nodeName"] = pod.Spec.NodeName
			return labelsSet, fieldsSet, nil
		},
		NewFunc:     func() runtime.Object { return &example.Pod{} },
		NewListFunc: func() runtime.Object { return &example.PodList{} },
		Codec:       codecs.LegacyCodec(examplev1.SchemeGroupVersion),
		Clock:       clock.RealClock{},
	}
	cacher, err := NewCacherFromConfig(config)
	return cacher, storage.APIObjectVersioner{}, err
}

type dummyStorage struct {
	sync.RWMutex
	err       error
	getListFn func(_ context.Context, _ string, _ storage.ListOptions, listObj runtime.Object) error
	watchFn   func(_ context.Context, _ string, _ storage.ListOptions) (watch.Interface, error)

	// use getRequestWatchProgressCounter when reading
	// the value of the counter
	requestWatchProgressCounter int
}

func (d *dummyStorage) RequestWatchProgress(ctx context.Context) error {
	d.Lock()
	defer d.Unlock()
	d.requestWatchProgressCounter++
	return nil
}

func (d *dummyStorage) getRequestWatchProgressCounter() int {
	d.RLock()
	defer d.RUnlock()
	return d.requestWatchProgressCounter
}

type dummyWatch struct {
	ch chan watch.Event
}

func (w *dummyWatch) ResultChan() <-chan watch.Event {
	return w.ch
}

func (w *dummyWatch) Stop() {
	close(w.ch)
}

func newDummyWatch() watch.Interface {
	return &dummyWatch{
		ch: make(chan watch.Event),
	}
}

func (d *dummyStorage) Versioner() storage.Versioner { return nil }
func (d *dummyStorage) Create(_ context.Context, _ string, _, _ runtime.Object, _ uint64) error {
	return fmt.Errorf("unimplemented")
}
func (d *dummyStorage) Delete(_ context.Context, _ string, _ runtime.Object, _ *storage.Preconditions, _ storage.ValidateObjectFunc, _ runtime.Object) error {
	return fmt.Errorf("unimplemented")
}
func (d *dummyStorage) Watch(ctx context.Context, key string, opts storage.ListOptions) (watch.Interface, error) {
	if d.watchFn != nil {
		return d.watchFn(ctx, key, opts)
	}
	d.RLock()
	defer d.RUnlock()

	return newDummyWatch(), d.err
}
func (d *dummyStorage) Get(_ context.Context, _ string, _ storage.GetOptions, _ runtime.Object) error {
	d.RLock()
	defer d.RUnlock()

	return d.err
}
func (d *dummyStorage) GetList(ctx context.Context, resPrefix string, opts storage.ListOptions, listObj runtime.Object) error {
	if d.getListFn != nil {
		return d.getListFn(ctx, resPrefix, opts, listObj)
	}
	d.RLock()
	defer d.RUnlock()
	podList := listObj.(*example.PodList)
	podList.ListMeta = metav1.ListMeta{ResourceVersion: "100"}
	return d.err
}
func (d *dummyStorage) GuaranteedUpdate(_ context.Context, _ string, _ runtime.Object, _ bool, _ *storage.Preconditions, _ storage.UpdateFunc, _ runtime.Object) error {
	return fmt.Errorf("unimplemented")
}
func (d *dummyStorage) Count(_ string) (int64, error) {
	return 0, fmt.Errorf("unimplemented")
}
func (d *dummyStorage) injectError(err error) {
	d.Lock()
	defer d.Unlock()

	d.err = err
}
