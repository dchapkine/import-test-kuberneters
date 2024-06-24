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

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	testing "k8s.io/client-go/testing"
)

// FakePersistentVolumeClaims implements PersistentVolumeClaimInterface
type FakePersistentVolumeClaims struct {
	Fake *FakeCoreV1
	ns   string
}

var persistentvolumeclaimsResource = v1.SchemeGroupVersion.WithResource("persistentvolumeclaims")

var persistentvolumeclaimsKind = v1.SchemeGroupVersion.WithKind("PersistentVolumeClaim")

// Get takes name of the persistentVolumeClaim, and returns the corresponding persistentVolumeClaim object, and an error if there is any.
func (c *FakePersistentVolumeClaims) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PersistentVolumeClaim, err error) {
	emptyResult := &v1.PersistentVolumeClaim{}
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(persistentvolumeclaimsResource, c.ns, name), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolumeClaim), err
}

// List takes label and field selectors, and returns the list of PersistentVolumeClaims that match those selectors.
func (c *FakePersistentVolumeClaims) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PersistentVolumeClaimList, err error) {
	emptyResult := &v1.PersistentVolumeClaimList{}
	obj, err := c.Fake.
		Invokes(testing.NewListAction(persistentvolumeclaimsResource, persistentvolumeclaimsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	list := &v1.PersistentVolumeClaimList{ListMeta: obj.(*v1.PersistentVolumeClaimList).ListMeta}
	for _, item := range obj.(*v1.PersistentVolumeClaimList).Items {
		if testing.ExtractFromListOptions(opts).Labels.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested persistentVolumeClaims.
func (c *FakePersistentVolumeClaims) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(persistentvolumeclaimsResource, c.ns, opts))

}

// Create takes the representation of a persistentVolumeClaim and creates it.  Returns the server's representation of the persistentVolumeClaim, and an error, if there is any.
func (c *FakePersistentVolumeClaims) Create(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaim, opts metav1.CreateOptions) (result *v1.PersistentVolumeClaim, err error) {
	emptyResult := &v1.PersistentVolumeClaim{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(persistentvolumeclaimsResource, c.ns, persistentVolumeClaim), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolumeClaim), err
}

// Update takes the representation of a persistentVolumeClaim and updates it. Returns the server's representation of the persistentVolumeClaim, and an error, if there is any.
func (c *FakePersistentVolumeClaims) Update(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaim, opts metav1.UpdateOptions) (result *v1.PersistentVolumeClaim, err error) {
	emptyResult := &v1.PersistentVolumeClaim{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(persistentvolumeclaimsResource, c.ns, persistentVolumeClaim), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolumeClaim), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePersistentVolumeClaims) UpdateStatus(ctx context.Context, persistentVolumeClaim *v1.PersistentVolumeClaim, opts metav1.UpdateOptions) (result *v1.PersistentVolumeClaim, err error) {
	emptyResult := &v1.PersistentVolumeClaim{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(persistentvolumeclaimsResource, "status", c.ns, persistentVolumeClaim), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolumeClaim), err
}

// Delete takes name of the persistentVolumeClaim and deletes it. Returns an error if one occurs.
func (c *FakePersistentVolumeClaims) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(persistentvolumeclaimsResource, c.ns, name, opts), &v1.PersistentVolumeClaim{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePersistentVolumeClaims) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(persistentvolumeclaimsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.PersistentVolumeClaimList{})
	return err
}

// Patch applies the patch and returns the patched persistentVolumeClaim.
func (c *FakePersistentVolumeClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PersistentVolumeClaim, err error) {
	emptyResult := &v1.PersistentVolumeClaim{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(persistentvolumeclaimsResource, c.ns, name, pt, data, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolumeClaim), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied persistentVolumeClaim.
func (c *FakePersistentVolumeClaims) Apply(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolumeClaim, err error) {
	if persistentVolumeClaim == nil {
		return nil, fmt.Errorf("persistentVolumeClaim provided to Apply must not be nil")
	}
	data, err := json.Marshal(persistentVolumeClaim)
	if err != nil {
		return nil, err
	}
	name := persistentVolumeClaim.Name
	if name == nil {
		return nil, fmt.Errorf("persistentVolumeClaim.Name must be provided to Apply")
	}
	emptyResult := &v1.PersistentVolumeClaim{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(persistentvolumeclaimsResource, c.ns, *name, types.ApplyPatchType, data), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolumeClaim), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePersistentVolumeClaims) ApplyStatus(ctx context.Context, persistentVolumeClaim *corev1.PersistentVolumeClaimApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolumeClaim, err error) {
	if persistentVolumeClaim == nil {
		return nil, fmt.Errorf("persistentVolumeClaim provided to Apply must not be nil")
	}
	data, err := json.Marshal(persistentVolumeClaim)
	if err != nil {
		return nil, err
	}
	name := persistentVolumeClaim.Name
	if name == nil {
		return nil, fmt.Errorf("persistentVolumeClaim.Name must be provided to Apply")
	}
	emptyResult := &v1.PersistentVolumeClaim{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(persistentvolumeclaimsResource, c.ns, *name, types.ApplyPatchType, data, "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolumeClaim), err
}
