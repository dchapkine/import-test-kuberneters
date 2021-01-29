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

package v2beta2

import (
	v2beta2 "k8s.io/api/autoscaling/v2beta2"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// MetricTargetApplyConfiguration represents an declarative configuration of the MetricTarget type for use
// with apply.
type MetricTargetApplyConfiguration struct {
	Type               *v2beta2.MetricTargetType `json:"type,omitempty"`
	Value              *resource.Quantity        `json:"value,omitempty"`
	AverageValue       *resource.Quantity        `json:"averageValue,omitempty"`
	AverageUtilization *int32                    `json:"averageUtilization,omitempty"`
}

// MetricTargetList represents a listAlias of MetricTargetApplyConfiguration.
type MetricTargetList []*MetricTargetApplyConfiguration

// MetricTargetList represents a map of MetricTargetApplyConfiguration.
type MetricTargetMap map[string]MetricTargetApplyConfiguration
