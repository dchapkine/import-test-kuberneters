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

// IngressRuleValueApplyConfiguration represents an declarative configuration of the IngressRuleValue type for use
// with apply.
type IngressRuleValueApplyConfiguration struct {
	HTTP *HTTPIngressRuleValueApplyConfiguration `json:"http,omitempty"`
}

// IngressRuleValueList represents a listAlias of IngressRuleValueApplyConfiguration.
type IngressRuleValueList []*IngressRuleValueApplyConfiguration

// IngressRuleValueList represents a map of IngressRuleValueApplyConfiguration.
type IngressRuleValueMap map[string]IngressRuleValueApplyConfiguration
