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
)

// ResourceQuotaSpecApplyConfiguration represents an declarative configuration of the ResourceQuotaSpec type for use
// with apply.
type ResourceQuotaSpecApplyConfiguration struct {
	Hard          *v1.ResourceList                 `json:"hard,omitempty"`
	Scopes        *[]v1.ResourceQuotaScope         `json:"scopes,omitempty"`
	ScopeSelector *ScopeSelectorApplyConfiguration `json:"scopeSelector,omitempty"`
}

// ResourceQuotaSpecList represents a listAlias of ResourceQuotaSpecApplyConfiguration.
type ResourceQuotaSpecList []*ResourceQuotaSpecApplyConfiguration

// ResourceQuotaSpecList represents a map of ResourceQuotaSpecApplyConfiguration.
type ResourceQuotaSpecMap map[string]ResourceQuotaSpecApplyConfiguration
