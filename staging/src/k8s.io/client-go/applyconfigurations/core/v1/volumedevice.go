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

// VolumeDeviceApplyConfiguration represents an declarative configuration of the VolumeDevice type for use
// with apply.
type VolumeDeviceApplyConfiguration struct {
	Name       *string `json:"name,omitempty"`
	DevicePath *string `json:"devicePath,omitempty"`
}

// VolumeDeviceList represents a listAlias of VolumeDeviceApplyConfiguration.
type VolumeDeviceList []*VolumeDeviceApplyConfiguration

// VolumeDeviceList represents a map of VolumeDeviceApplyConfiguration.
type VolumeDeviceMap map[string]VolumeDeviceApplyConfiguration
