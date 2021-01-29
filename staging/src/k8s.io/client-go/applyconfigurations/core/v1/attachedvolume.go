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

import (
	v1 "k8s.io/api/core/v1"
)

// AttachedVolumeApplyConfiguration represents an declarative configuration of the AttachedVolume type for use
// with apply.
type AttachedVolumeApplyConfiguration struct {
	Name       *v1.UniqueVolumeName `json:"name,omitempty"`
	DevicePath *string              `json:"devicePath,omitempty"`
}

// AttachedVolumeList represents a listAlias of AttachedVolumeApplyConfiguration.
type AttachedVolumeList []*AttachedVolumeApplyConfiguration

// AttachedVolumeList represents a map of AttachedVolumeApplyConfiguration.
type AttachedVolumeMap map[string]AttachedVolumeApplyConfiguration
