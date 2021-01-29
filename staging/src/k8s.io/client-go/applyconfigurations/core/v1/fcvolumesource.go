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

// FCVolumeSourceApplyConfiguration represents an declarative configuration of the FCVolumeSource type for use
// with apply.
type FCVolumeSourceApplyConfiguration struct {
	TargetWWNs *[]string `json:"targetWWNs,omitempty"`
	Lun        *int32    `json:"lun,omitempty"`
	FSType     *string   `json:"fsType,omitempty"`
	ReadOnly   *bool     `json:"readOnly,omitempty"`
	WWIDs      *[]string `json:"wwids,omitempty"`
}

// FCVolumeSourceList represents a listAlias of FCVolumeSourceApplyConfiguration.
type FCVolumeSourceList []*FCVolumeSourceApplyConfiguration

// FCVolumeSourceList represents a map of FCVolumeSourceApplyConfiguration.
type FCVolumeSourceMap map[string]FCVolumeSourceApplyConfiguration
