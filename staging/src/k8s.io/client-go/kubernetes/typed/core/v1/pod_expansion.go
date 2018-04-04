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

package v1

import (
	"k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
)

// The PodExpansion interface allows manually adding extra methods to the PodInterface.
type PodExpansion interface {
	Bind(binding *v1.Binding) error
	Evict(eviction *policy.Eviction) error
	GetLogs(name string, opts *v1.PodLogOptions) *restclient.Request
	PatchStatus(name string, data []byte) (*v1.Pod, error)
}

// Bind applies the provided binding to the named pod in the current namespace (binding.Namespace is ignored).
func (c *pods) Bind(binding *v1.Binding) error {
	return c.client.Post().Namespace(c.ns).Resource("pods").Name(binding.Name).SubResource("binding").Body(binding).Do().Error()
}

func (c *pods) Evict(eviction *policy.Eviction) error {
	return c.client.Post().Namespace(c.ns).Resource("pods").Name(eviction.Name).SubResource("eviction").Body(eviction).Do().Error()
}

// Get constructs a request for getting the logs for a pod
func (c *pods) GetLogs(name string, opts *v1.PodLogOptions) *restclient.Request {
	return c.client.Get().Namespace(c.ns).Name(name).Resource("pods").SubResource("log").VersionedParams(opts, scheme.ParameterCodec)
}

// PatchStatus modifies the status of an existing pod. It returns the copy of
// the pod that the server returns, or an error.
func (c *pods) PatchStatus(name string, data []byte) (*v1.Pod, error) {
	result := &v1.Pod{}
	err := c.client.Patch(types.StrategicMergePatchType).
		Namespace(c.ns).
		Resource("pods").
		Name(name).
		SubResource("status").
		Body(data).
		Do().
		Into(result)
	return result, err
}
