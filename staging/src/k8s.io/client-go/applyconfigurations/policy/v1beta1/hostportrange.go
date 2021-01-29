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

package v1beta1

// HostPortRangeApplyConfiguration represents an declarative configuration of the HostPortRange type for use
// with apply.
type HostPortRangeApplyConfiguration struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// HostPortRangeList represents a listAlias of HostPortRangeApplyConfiguration.
type HostPortRangeList []*HostPortRangeApplyConfiguration

// HostPortRangeList represents a map of HostPortRangeApplyConfiguration.
type HostPortRangeMap map[string]HostPortRangeApplyConfiguration
