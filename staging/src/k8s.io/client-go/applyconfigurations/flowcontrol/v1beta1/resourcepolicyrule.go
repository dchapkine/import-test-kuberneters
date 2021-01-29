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

// ResourcePolicyRuleApplyConfiguration represents an declarative configuration of the ResourcePolicyRule type for use
// with apply.
type ResourcePolicyRuleApplyConfiguration struct {
	Verbs        *[]string `json:"verbs,omitempty"`
	APIGroups    *[]string `json:"apiGroups,omitempty"`
	Resources    *[]string `json:"resources,omitempty"`
	ClusterScope *bool     `json:"clusterScope,omitempty"`
	Namespaces   *[]string `json:"namespaces,omitempty"`
}

// ResourcePolicyRuleList represents a listAlias of ResourcePolicyRuleApplyConfiguration.
type ResourcePolicyRuleList []*ResourcePolicyRuleApplyConfiguration

// ResourcePolicyRuleList represents a map of ResourcePolicyRuleApplyConfiguration.
type ResourcePolicyRuleMap map[string]ResourcePolicyRuleApplyConfiguration
