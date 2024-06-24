/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	appsv1beta2 "k8s.io/client-go/applyconfigurations/apps/v1beta2"
	testing "k8s.io/client-go/testing"
)

// FakeStatefulSets implements StatefulSetInterface
type FakeStatefulSets struct {
	Fake *FakeAppsV1beta2
	ns   string
}

var statefulsetsResource = v1beta2.SchemeGroupVersion.WithResource("statefulsets")

var statefulsetsKind = v1beta2.SchemeGroupVersion.WithKind("StatefulSet")

// Get takes name of the statefulSet, and returns the corresponding statefulSet object, and an error if there is any.
func (c *FakeStatefulSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.StatefulSet, err error) {
	emptyResult := &v1beta2.StatefulSet{}
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(statefulsetsResource, c.ns, name), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.StatefulSet), err
}

// List takes label and field selectors, and returns the list of StatefulSets that match those selectors.
func (c *FakeStatefulSets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.StatefulSetList, err error) {
	emptyResult := &v1beta2.StatefulSetList{}
	obj, err := c.Fake.
		Invokes(testing.NewListAction(statefulsetsResource, statefulsetsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	list := &v1beta2.StatefulSetList{ListMeta: obj.(*v1beta2.StatefulSetList).ListMeta}
	for _, item := range obj.(*v1beta2.StatefulSetList).Items {
		if testing.ExtractFromListOptions(opts).Labels.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested statefulSets.
func (c *FakeStatefulSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(statefulsetsResource, c.ns, opts))

}

// Create takes the representation of a statefulSet and creates it.  Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *FakeStatefulSets) Create(ctx context.Context, statefulSet *v1beta2.StatefulSet, opts v1.CreateOptions) (result *v1beta2.StatefulSet, err error) {
	emptyResult := &v1beta2.StatefulSet{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(statefulsetsResource, c.ns, statefulSet), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.StatefulSet), err
}

// Update takes the representation of a statefulSet and updates it. Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *FakeStatefulSets) Update(ctx context.Context, statefulSet *v1beta2.StatefulSet, opts v1.UpdateOptions) (result *v1beta2.StatefulSet, err error) {
	emptyResult := &v1beta2.StatefulSet{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(statefulsetsResource, c.ns, statefulSet), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.StatefulSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStatefulSets) UpdateStatus(ctx context.Context, statefulSet *v1beta2.StatefulSet, opts v1.UpdateOptions) (result *v1beta2.StatefulSet, err error) {
	emptyResult := &v1beta2.StatefulSet{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(statefulsetsResource, "status", c.ns, statefulSet), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.StatefulSet), err
}

// Delete takes name of the statefulSet and deletes it. Returns an error if one occurs.
func (c *FakeStatefulSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(statefulsetsResource, c.ns, name, opts), &v1beta2.StatefulSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStatefulSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(statefulsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.StatefulSetList{})
	return err
}

// Patch applies the patch and returns the patched statefulSet.
func (c *FakeStatefulSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.StatefulSet, err error) {
	emptyResult := &v1beta2.StatefulSet{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(statefulsetsResource, c.ns, name, pt, data, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.StatefulSet), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied statefulSet.
func (c *FakeStatefulSets) Apply(ctx context.Context, statefulSet *appsv1beta2.StatefulSetApplyConfiguration, opts v1.ApplyOptions) (result *v1beta2.StatefulSet, err error) {
	if statefulSet == nil {
		return nil, fmt.Errorf("statefulSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(statefulSet)
	if err != nil {
		return nil, err
	}
	name := statefulSet.Name
	if name == nil {
		return nil, fmt.Errorf("statefulSet.Name must be provided to Apply")
	}
	emptyResult := &v1beta2.StatefulSet{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(statefulsetsResource, c.ns, *name, types.ApplyPatchType, data), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.StatefulSet), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeStatefulSets) ApplyStatus(ctx context.Context, statefulSet *appsv1beta2.StatefulSetApplyConfiguration, opts v1.ApplyOptions) (result *v1beta2.StatefulSet, err error) {
	if statefulSet == nil {
		return nil, fmt.Errorf("statefulSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(statefulSet)
	if err != nil {
		return nil, err
	}
	name := statefulSet.Name
	if name == nil {
		return nil, fmt.Errorf("statefulSet.Name must be provided to Apply")
	}
	emptyResult := &v1beta2.StatefulSet{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(statefulsetsResource, c.ns, *name, types.ApplyPatchType, data, "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.StatefulSet), err
}

// GetScale takes name of the statefulSet, and returns the corresponding scale object, and an error if there is any.
func (c *FakeStatefulSets) GetScale(ctx context.Context, statefulSetName string, options v1.GetOptions) (result *v1beta2.Scale, err error) {
	emptyResult := &v1beta2.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(statefulsetsResource, c.ns, "scale", statefulSetName), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeStatefulSets) UpdateScale(ctx context.Context, statefulSetName string, scale *v1beta2.Scale, opts v1.UpdateOptions) (result *v1beta2.Scale, err error) {
	emptyResult := &v1beta2.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(statefulsetsResource, "scale", c.ns, scale), &v1beta2.Scale{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.Scale), err
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *FakeStatefulSets) ApplyScale(ctx context.Context, statefulSetName string, scale *appsv1beta2.ScaleApplyConfiguration, opts v1.ApplyOptions) (result *v1beta2.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	data, err := json.Marshal(scale)
	if err != nil {
		return nil, err
	}
	emptyResult := &v1beta2.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(statefulsetsResource, c.ns, statefulSetName, types.ApplyPatchType, data, "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta2.Scale), err
}
