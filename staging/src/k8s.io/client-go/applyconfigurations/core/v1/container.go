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
	corev1 "k8s.io/api/core/v1"
)

// ContainerApplyConfiguration represents an declarative configuration of the Container type for use
// with apply.
type ContainerApplyConfiguration struct {
	Name                     *string                                 `json:"name,omitempty"`
	Image                    *string                                 `json:"image,omitempty"`
	Command                  *[]string                               `json:"command,omitempty"`
	Args                     *[]string                               `json:"args,omitempty"`
	WorkingDir               *string                                 `json:"workingDir,omitempty"`
	Ports                    *ContainerPortList                      `json:"ports,omitempty"`
	EnvFrom                  *EnvFromSourceList                      `json:"envFrom,omitempty"`
	Env                      *EnvVarList                             `json:"env,omitempty"`
	Resources                *ResourceRequirementsApplyConfiguration `json:"resources,omitempty"`
	VolumeMounts             *VolumeMountList                        `json:"volumeMounts,omitempty"`
	VolumeDevices            *VolumeDeviceList                       `json:"volumeDevices,omitempty"`
	LivenessProbe            *ProbeApplyConfiguration                `json:"livenessProbe,omitempty"`
	ReadinessProbe           *ProbeApplyConfiguration                `json:"readinessProbe,omitempty"`
	StartupProbe             *ProbeApplyConfiguration                `json:"startupProbe,omitempty"`
	Lifecycle                *LifecycleApplyConfiguration            `json:"lifecycle,omitempty"`
	TerminationMessagePath   *string                                 `json:"terminationMessagePath,omitempty"`
	TerminationMessagePolicy *corev1.TerminationMessagePolicy        `json:"terminationMessagePolicy,omitempty"`
	ImagePullPolicy          *corev1.PullPolicy                      `json:"imagePullPolicy,omitempty"`
	SecurityContext          *SecurityContextApplyConfiguration      `json:"securityContext,omitempty"`
	Stdin                    *bool                                   `json:"stdin,omitempty"`
	StdinOnce                *bool                                   `json:"stdinOnce,omitempty"`
	TTY                      *bool                                   `json:"tty,omitempty"`
}

// ContainerList represents a listAlias of ContainerApplyConfiguration.
type ContainerList []*ContainerApplyConfiguration

// ContainerList represents a map of ContainerApplyConfiguration.
type ContainerMap map[string]ContainerApplyConfiguration
