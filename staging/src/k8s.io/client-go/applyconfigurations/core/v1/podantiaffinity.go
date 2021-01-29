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

// PodAntiAffinityApplyConfiguration represents an declarative configuration of the PodAntiAffinity type for use
// with apply.
type PodAntiAffinityApplyConfiguration struct {
	RequiredDuringSchedulingIgnoredDuringExecution  *PodAffinityTermList         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution *WeightedPodAffinityTermList `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// PodAntiAffinityList represents a listAlias of PodAntiAffinityApplyConfiguration.
type PodAntiAffinityList []*PodAntiAffinityApplyConfiguration

// PodAntiAffinityList represents a map of PodAntiAffinityApplyConfiguration.
type PodAntiAffinityMap map[string]PodAntiAffinityApplyConfiguration
