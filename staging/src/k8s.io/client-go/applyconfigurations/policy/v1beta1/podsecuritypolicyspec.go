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

import (
	v1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/policy/v1beta1"
)

// PodSecurityPolicySpecApplyConfiguration represents an declarative configuration of the PodSecurityPolicySpec type for use
// with apply.
type PodSecurityPolicySpecApplyConfiguration struct {
	Privileged                      *bool                                                `json:"privileged,omitempty"`
	DefaultAddCapabilities          *[]v1.Capability                                     `json:"defaultAddCapabilities,omitempty"`
	RequiredDropCapabilities        *[]v1.Capability                                     `json:"requiredDropCapabilities,omitempty"`
	AllowedCapabilities             *[]v1.Capability                                     `json:"allowedCapabilities,omitempty"`
	Volumes                         *[]v1beta1.FSType                                    `json:"volumes,omitempty"`
	HostNetwork                     *bool                                                `json:"hostNetwork,omitempty"`
	HostPorts                       *HostPortRangeList                                   `json:"hostPorts,omitempty"`
	HostPID                         *bool                                                `json:"hostPID,omitempty"`
	HostIPC                         *bool                                                `json:"hostIPC,omitempty"`
	SELinux                         *SELinuxStrategyOptionsApplyConfiguration            `json:"seLinux,omitempty"`
	RunAsUser                       *RunAsUserStrategyOptionsApplyConfiguration          `json:"runAsUser,omitempty"`
	RunAsGroup                      *RunAsGroupStrategyOptionsApplyConfiguration         `json:"runAsGroup,omitempty"`
	SupplementalGroups              *SupplementalGroupsStrategyOptionsApplyConfiguration `json:"supplementalGroups,omitempty"`
	FSGroup                         *FSGroupStrategyOptionsApplyConfiguration            `json:"fsGroup,omitempty"`
	ReadOnlyRootFilesystem          *bool                                                `json:"readOnlyRootFilesystem,omitempty"`
	DefaultAllowPrivilegeEscalation *bool                                                `json:"defaultAllowPrivilegeEscalation,omitempty"`
	AllowPrivilegeEscalation        *bool                                                `json:"allowPrivilegeEscalation,omitempty"`
	AllowedHostPaths                *AllowedHostPathList                                 `json:"allowedHostPaths,omitempty"`
	AllowedFlexVolumes              *AllowedFlexVolumeList                               `json:"allowedFlexVolumes,omitempty"`
	AllowedCSIDrivers               *AllowedCSIDriverList                                `json:"allowedCSIDrivers,omitempty"`
	AllowedUnsafeSysctls            *[]string                                            `json:"allowedUnsafeSysctls,omitempty"`
	ForbiddenSysctls                *[]string                                            `json:"forbiddenSysctls,omitempty"`
	AllowedProcMountTypes           *[]v1.ProcMountType                                  `json:"allowedProcMountTypes,omitempty"`
	RuntimeClass                    *RuntimeClassStrategyOptionsApplyConfiguration       `json:"runtimeClass,omitempty"`
}

// PodSecurityPolicySpecList represents a listAlias of PodSecurityPolicySpecApplyConfiguration.
type PodSecurityPolicySpecList []*PodSecurityPolicySpecApplyConfiguration

// PodSecurityPolicySpecList represents a map of PodSecurityPolicySpecApplyConfiguration.
type PodSecurityPolicySpecMap map[string]PodSecurityPolicySpecApplyConfiguration
