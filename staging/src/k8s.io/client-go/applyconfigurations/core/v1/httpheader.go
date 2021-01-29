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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// HTTPHeaderApplyConfiguration represents an declarative configuration of the HTTPHeader type for use
// with apply.
type HTTPHeaderApplyConfiguration struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// HTTPHeaderList represents a listAlias of HTTPHeaderApplyConfiguration.
type HTTPHeaderList []*HTTPHeaderApplyConfiguration

// HTTPHeaderList represents a map of HTTPHeaderApplyConfiguration.
type HTTPHeaderMap map[string]HTTPHeaderApplyConfiguration
