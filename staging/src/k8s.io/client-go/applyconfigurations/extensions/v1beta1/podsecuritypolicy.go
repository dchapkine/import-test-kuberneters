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
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// PodSecurityPolicyApplyConfiguration represents an declarative configuration of the PodSecurityPolicy type for use
// with apply.
type PodSecurityPolicyApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration         `json:"metadata,omitempty"`
	Spec                          *PodSecurityPolicySpecApplyConfiguration `json:"spec,omitempty"`
}

// PodSecurityPolicyList represents a listAlias of PodSecurityPolicyApplyConfiguration.
type PodSecurityPolicyList []*PodSecurityPolicyApplyConfiguration

// PodSecurityPolicyList represents a map of PodSecurityPolicyApplyConfiguration.
type PodSecurityPolicyMap map[string]PodSecurityPolicyApplyConfiguration
