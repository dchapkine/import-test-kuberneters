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

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	coordination "k8s.io/kubernetes/pkg/apis/coordination"
	scheme "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/scheme"
)

// LeasesGetter has a method to return a LeaseInterface.
// A group's client should implement this interface.
type LeasesGetter interface {
	Leases(namespace string) LeaseInterface
}

// LeaseInterface has methods to work with Lease resources.
type LeaseInterface interface {
	Create(*coordination.Lease) (*coordination.Lease, error)
	Update(*coordination.Lease) (*coordination.Lease, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*coordination.Lease, error)
	List(opts v1.ListOptions) (*coordination.LeaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *coordination.Lease, err error)
	LeaseExpansion
}

// leases implements LeaseInterface
type leases struct {
	client rest.Interface
	ns     string
}

// newLeases returns a Leases
func newLeases(c *CoordinationClient, namespace string) *leases {
	return &leases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lease, and returns the corresponding lease object, and an error if there is any.
func (c *leases) Get(name string, options v1.GetOptions) (result *coordination.Lease, err error) {
	result = &coordination.Lease{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("leases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Leases that match those selectors.
func (c *leases) List(opts v1.ListOptions) (result *coordination.LeaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds)
	}
	result = &coordination.LeaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("leases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested leases.
func (c *leases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds)
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("leases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lease and creates it.  Returns the server's representation of the lease, and an error, if there is any.
func (c *leases) Create(lease *coordination.Lease) (result *coordination.Lease, err error) {
	result = &coordination.Lease{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("leases").
		Body(lease).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lease and updates it. Returns the server's representation of the lease, and an error, if there is any.
func (c *leases) Update(lease *coordination.Lease) (result *coordination.Lease, err error) {
	result = &coordination.Lease{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("leases").
		Name(lease.Name).
		Body(lease).
		Do().
		Into(result)
	return
}

// Delete takes name of the lease and deletes it. Returns an error if one occurs.
func (c *leases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("leases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *leases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds)
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("leases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lease.
func (c *leases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *coordination.Lease, err error) {
	result = &coordination.Lease{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("leases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
