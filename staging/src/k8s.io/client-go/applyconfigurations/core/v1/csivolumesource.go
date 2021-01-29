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

// CSIVolumeSourceApplyConfiguration represents an declarative configuration of the CSIVolumeSource type for use
// with apply.
type CSIVolumeSourceApplyConfiguration struct {
	Driver               *string                                 `json:"driver,omitempty"`
	ReadOnly             *bool                                   `json:"readOnly,omitempty"`
	FSType               *string                                 `json:"fsType,omitempty"`
	VolumeAttributes     *map[string]string                      `json:"volumeAttributes,omitempty"`
	NodePublishSecretRef *LocalObjectReferenceApplyConfiguration `json:"nodePublishSecretRef,omitempty"`
}

// CSIVolumeSourceList represents a listAlias of CSIVolumeSourceApplyConfiguration.
type CSIVolumeSourceList []*CSIVolumeSourceApplyConfiguration

// CSIVolumeSourceList represents a map of CSIVolumeSourceApplyConfiguration.
type CSIVolumeSourceMap map[string]CSIVolumeSourceApplyConfiguration
