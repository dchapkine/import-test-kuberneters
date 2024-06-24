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

	v1beta1 "k8s.io/api/certificates/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	certificatesv1beta1 "k8s.io/client-go/applyconfigurations/certificates/v1beta1"
	testing "k8s.io/client-go/testing"
)

// FakeCertificateSigningRequests implements CertificateSigningRequestInterface
type FakeCertificateSigningRequests struct {
	Fake *FakeCertificatesV1beta1
}

var certificatesigningrequestsResource = v1beta1.SchemeGroupVersion.WithResource("certificatesigningrequests")

var certificatesigningrequestsKind = v1beta1.SchemeGroupVersion.WithKind("CertificateSigningRequest")

// Get takes name of the certificateSigningRequest, and returns the corresponding certificateSigningRequest object, and an error if there is any.
func (c *FakeCertificateSigningRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CertificateSigningRequest, err error) {
	emptyResult := &v1beta1.CertificateSigningRequest{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(certificatesigningrequestsResource, name), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CertificateSigningRequest), err
}

// List takes label and field selectors, and returns the list of CertificateSigningRequests that match those selectors.
func (c *FakeCertificateSigningRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CertificateSigningRequestList, err error) {
	emptyResult := &v1beta1.CertificateSigningRequestList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(certificatesigningrequestsResource, certificatesigningrequestsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	list := &v1beta1.CertificateSigningRequestList{ListMeta: obj.(*v1beta1.CertificateSigningRequestList).ListMeta}
	for _, item := range obj.(*v1beta1.CertificateSigningRequestList).Items {
		if testing.ExtractFromListOptions(opts).Labels.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificateSigningRequests.
func (c *FakeCertificateSigningRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(certificatesigningrequestsResource, opts))
}

// Create takes the representation of a certificateSigningRequest and creates it.  Returns the server's representation of the certificateSigningRequest, and an error, if there is any.
func (c *FakeCertificateSigningRequests) Create(ctx context.Context, certificateSigningRequest *v1beta1.CertificateSigningRequest, opts v1.CreateOptions) (result *v1beta1.CertificateSigningRequest, err error) {
	emptyResult := &v1beta1.CertificateSigningRequest{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(certificatesigningrequestsResource, certificateSigningRequest), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CertificateSigningRequest), err
}

// Update takes the representation of a certificateSigningRequest and updates it. Returns the server's representation of the certificateSigningRequest, and an error, if there is any.
func (c *FakeCertificateSigningRequests) Update(ctx context.Context, certificateSigningRequest *v1beta1.CertificateSigningRequest, opts v1.UpdateOptions) (result *v1beta1.CertificateSigningRequest, err error) {
	emptyResult := &v1beta1.CertificateSigningRequest{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(certificatesigningrequestsResource, certificateSigningRequest), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CertificateSigningRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificateSigningRequests) UpdateStatus(ctx context.Context, certificateSigningRequest *v1beta1.CertificateSigningRequest, opts v1.UpdateOptions) (result *v1beta1.CertificateSigningRequest, err error) {
	emptyResult := &v1beta1.CertificateSigningRequest{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(certificatesigningrequestsResource, "status", certificateSigningRequest), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CertificateSigningRequest), err
}

// Delete takes name of the certificateSigningRequest and deletes it. Returns an error if one occurs.
func (c *FakeCertificateSigningRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(certificatesigningrequestsResource, name, opts), &v1beta1.CertificateSigningRequest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificateSigningRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(certificatesigningrequestsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CertificateSigningRequestList{})
	return err
}

// Patch applies the patch and returns the patched certificateSigningRequest.
func (c *FakeCertificateSigningRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CertificateSigningRequest, err error) {
	emptyResult := &v1beta1.CertificateSigningRequest{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(certificatesigningrequestsResource, name, pt, data, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CertificateSigningRequest), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied certificateSigningRequest.
func (c *FakeCertificateSigningRequests) Apply(ctx context.Context, certificateSigningRequest *certificatesv1beta1.CertificateSigningRequestApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.CertificateSigningRequest, err error) {
	if certificateSigningRequest == nil {
		return nil, fmt.Errorf("certificateSigningRequest provided to Apply must not be nil")
	}
	data, err := json.Marshal(certificateSigningRequest)
	if err != nil {
		return nil, err
	}
	name := certificateSigningRequest.Name
	if name == nil {
		return nil, fmt.Errorf("certificateSigningRequest.Name must be provided to Apply")
	}
	emptyResult := &v1beta1.CertificateSigningRequest{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(certificatesigningrequestsResource, *name, types.ApplyPatchType, data), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CertificateSigningRequest), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeCertificateSigningRequests) ApplyStatus(ctx context.Context, certificateSigningRequest *certificatesv1beta1.CertificateSigningRequestApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.CertificateSigningRequest, err error) {
	if certificateSigningRequest == nil {
		return nil, fmt.Errorf("certificateSigningRequest provided to Apply must not be nil")
	}
	data, err := json.Marshal(certificateSigningRequest)
	if err != nil {
		return nil, err
	}
	name := certificateSigningRequest.Name
	if name == nil {
		return nil, fmt.Errorf("certificateSigningRequest.Name must be provided to Apply")
	}
	emptyResult := &v1beta1.CertificateSigningRequest{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(certificatesigningrequestsResource, *name, types.ApplyPatchType, data, "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.CertificateSigningRequest), err
}
