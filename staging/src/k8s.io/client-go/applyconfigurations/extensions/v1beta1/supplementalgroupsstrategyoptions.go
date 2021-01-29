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
	v1beta1 "k8s.io/api/extensions/v1beta1"
)

// SupplementalGroupsStrategyOptionsApplyConfiguration represents an declarative configuration of the SupplementalGroupsStrategyOptions type for use
// with apply.
type SupplementalGroupsStrategyOptionsApplyConfiguration struct {
	Rule   *v1beta1.SupplementalGroupsStrategyType `json:"rule,omitempty"`
	Ranges *IDRangeList                            `json:"ranges,omitempty"`
}

// SupplementalGroupsStrategyOptionsList represents a listAlias of SupplementalGroupsStrategyOptionsApplyConfiguration.
type SupplementalGroupsStrategyOptionsList []*SupplementalGroupsStrategyOptionsApplyConfiguration

// SupplementalGroupsStrategyOptionsList represents a map of SupplementalGroupsStrategyOptionsApplyConfiguration.
type SupplementalGroupsStrategyOptionsMap map[string]SupplementalGroupsStrategyOptionsApplyConfiguration
