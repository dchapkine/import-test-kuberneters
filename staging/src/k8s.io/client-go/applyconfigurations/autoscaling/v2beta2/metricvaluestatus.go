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
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// MetricValueStatusApplyConfiguration represents an declarative configuration of the MetricValueStatus type for use
// with apply.
type MetricValueStatusApplyConfiguration struct {
	Value              *resource.Quantity `json:"value,omitempty"`
	AverageValue       *resource.Quantity `json:"averageValue,omitempty"`
	AverageUtilization *int32             `json:"averageUtilization,omitempty"`
}

// MetricValueStatusList represents a listAlias of MetricValueStatusApplyConfiguration.
type MetricValueStatusList []*MetricValueStatusApplyConfiguration

// MetricValueStatusList represents a map of MetricValueStatusApplyConfiguration.
type MetricValueStatusMap map[string]MetricValueStatusApplyConfiguration
