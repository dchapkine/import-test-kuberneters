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

package v1beta2

import (
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// DeploymentSpecApplyConfiguration represents an declarative configuration of the DeploymentSpec type for use
// with apply.
type DeploymentSpecApplyConfiguration struct {
	Replicas                *int32                                    `json:"replicas,omitempty"`
	Selector                *v1.LabelSelectorApplyConfiguration       `json:"selector,omitempty"`
	Template                *corev1.PodTemplateSpecApplyConfiguration `json:"template,omitempty"`
	Strategy                *DeploymentStrategyApplyConfiguration     `json:"strategy,omitempty"`
	MinReadySeconds         *int32                                    `json:"minReadySeconds,omitempty"`
	RevisionHistoryLimit    *int32                                    `json:"revisionHistoryLimit,omitempty"`
	Paused                  *bool                                     `json:"paused,omitempty"`
	ProgressDeadlineSeconds *int32                                    `json:"progressDeadlineSeconds,omitempty"`
}

// DeploymentSpecList represents a listAlias of DeploymentSpecApplyConfiguration.
type DeploymentSpecList []*DeploymentSpecApplyConfiguration

// DeploymentSpecList represents a map of DeploymentSpecApplyConfiguration.
type DeploymentSpecMap map[string]DeploymentSpecApplyConfiguration
