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

package internalversion

import (
	"k8s.io/kubernetes/pkg/api"
	policy "k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/client/restclient"
)

// The PodExpansion interface allows manually adding extra methods to the PodInterface.
type PodExpansion interface {
	Bind(binding *api.Binding) error
	Evict(eviction *policy.Eviction) error
	GetLogs(name string, opts *api.PodLogOptions) *restclient.Request
}

// Bind applies the provided binding to the named pod in the current namespace (binding.Namespace is ignored).
func (c *pods) Bind(binding *api.Binding) error {
	return c.client.Post().Namespace(c.ns).Resource("pods").Name(binding.Name).SubResource("binding").Body(binding).Do().Error()
}

// Evict trys to voluntarily evict pods based on its corresponding PodDisruptionBudget policy
func (c *pods) Evict(eviction *policy.Eviction) error {
	// TODO: We hack which codec func Body() is using to build the request body. To be fixed, when we found a good way
	// to make eviction works cross Core and Policy group.
	return c.client.Post().Namespace(c.ns).Resource("pods").Name(eviction.Name).SubResource("eviction").Body(eviction).Do().Error()
}

// Get constructs a request for getting the logs for a pod
func (c *pods) GetLogs(name string, opts *api.PodLogOptions) *restclient.Request {
	return c.client.Get().Namespace(c.ns).Name(name).Resource("pods").SubResource("log").VersionedParams(opts, api.ParameterCodec)
}
