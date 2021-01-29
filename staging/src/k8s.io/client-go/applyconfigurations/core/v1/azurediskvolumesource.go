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

// AzureDiskVolumeSourceApplyConfiguration represents an declarative configuration of the AzureDiskVolumeSource type for use
// with apply.
type AzureDiskVolumeSourceApplyConfiguration struct {
	DiskName    *string                      `json:"diskName,omitempty"`
	DataDiskURI *string                      `json:"diskURI,omitempty"`
	CachingMode *v1.AzureDataDiskCachingMode `json:"cachingMode,omitempty"`
	FSType      *string                      `json:"fsType,omitempty"`
	ReadOnly    *bool                        `json:"readOnly,omitempty"`
	Kind        *v1.AzureDataDiskKind        `json:"kind,omitempty"`
}

// AzureDiskVolumeSourceList represents a listAlias of AzureDiskVolumeSourceApplyConfiguration.
type AzureDiskVolumeSourceList []*AzureDiskVolumeSourceApplyConfiguration

// AzureDiskVolumeSourceList represents a map of AzureDiskVolumeSourceApplyConfiguration.
type AzureDiskVolumeSourceMap map[string]AzureDiskVolumeSourceApplyConfiguration
