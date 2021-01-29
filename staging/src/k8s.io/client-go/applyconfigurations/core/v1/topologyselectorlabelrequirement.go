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

// TopologySelectorLabelRequirementApplyConfiguration represents an declarative configuration of the TopologySelectorLabelRequirement type for use
// with apply.
type TopologySelectorLabelRequirementApplyConfiguration struct {
	Key    *string   `json:"key,omitempty"`
	Values *[]string `json:"values,omitempty"`
}

// TopologySelectorLabelRequirementList represents a listAlias of TopologySelectorLabelRequirementApplyConfiguration.
type TopologySelectorLabelRequirementList []*TopologySelectorLabelRequirementApplyConfiguration

// TopologySelectorLabelRequirementList represents a map of TopologySelectorLabelRequirementApplyConfiguration.
type TopologySelectorLabelRequirementMap map[string]TopologySelectorLabelRequirementApplyConfiguration
