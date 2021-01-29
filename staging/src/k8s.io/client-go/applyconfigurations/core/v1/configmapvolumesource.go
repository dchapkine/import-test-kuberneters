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

// ConfigMapVolumeSourceApplyConfiguration represents an declarative configuration of the ConfigMapVolumeSource type for use
// with apply.
type ConfigMapVolumeSourceApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Items                                  *KeyToPathList `json:"items,omitempty"`
	DefaultMode                            *int32         `json:"defaultMode,omitempty"`
	Optional                               *bool          `json:"optional,omitempty"`
}

// ConfigMapVolumeSourceList represents a listAlias of ConfigMapVolumeSourceApplyConfiguration.
type ConfigMapVolumeSourceList []*ConfigMapVolumeSourceApplyConfiguration

// ConfigMapVolumeSourceList represents a map of ConfigMapVolumeSourceApplyConfiguration.
type ConfigMapVolumeSourceMap map[string]ConfigMapVolumeSourceApplyConfiguration
