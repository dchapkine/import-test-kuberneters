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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PersistentVolumeClaimConditionApplyConfiguration represents an declarative configuration of the PersistentVolumeClaimCondition type for use
// with apply.
type PersistentVolumeClaimConditionApplyConfiguration struct {
	Type               *v1.PersistentVolumeClaimConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus                    `json:"status,omitempty"`
	LastProbeTime      *metav1.Time                           `json:"lastProbeTime,omitempty"`
	LastTransitionTime *metav1.Time                           `json:"lastTransitionTime,omitempty"`
	Reason             *string                                `json:"reason,omitempty"`
	Message            *string                                `json:"message,omitempty"`
}

// PersistentVolumeClaimConditionList represents a listAlias of PersistentVolumeClaimConditionApplyConfiguration.
type PersistentVolumeClaimConditionList []*PersistentVolumeClaimConditionApplyConfiguration

// PersistentVolumeClaimConditionList represents a map of PersistentVolumeClaimConditionApplyConfiguration.
type PersistentVolumeClaimConditionMap map[string]PersistentVolumeClaimConditionApplyConfiguration
