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
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ValidatingWebhookApplyConfiguration represents an declarative configuration of the ValidatingWebhook type for use
// with apply.
type ValidatingWebhookApplyConfiguration struct {
	Name                    *string                                         `json:"name,omitempty"`
	ClientConfig            *WebhookClientConfigApplyConfiguration          `json:"clientConfig,omitempty"`
	Rules                   *RuleWithOperationsList                         `json:"rules,omitempty"`
	FailurePolicy           *admissionregistrationv1beta1.FailurePolicyType `json:"failurePolicy,omitempty"`
	MatchPolicy             *admissionregistrationv1beta1.MatchPolicyType   `json:"matchPolicy,omitempty"`
	NamespaceSelector       *v1.LabelSelectorApplyConfiguration             `json:"namespaceSelector,omitempty"`
	ObjectSelector          *v1.LabelSelectorApplyConfiguration             `json:"objectSelector,omitempty"`
	SideEffects             *admissionregistrationv1beta1.SideEffectClass   `json:"sideEffects,omitempty"`
	TimeoutSeconds          *int32                                          `json:"timeoutSeconds,omitempty"`
	AdmissionReviewVersions *[]string                                       `json:"admissionReviewVersions,omitempty"`
}

// ValidatingWebhookList represents a listAlias of ValidatingWebhookApplyConfiguration.
type ValidatingWebhookList []*ValidatingWebhookApplyConfiguration

// ValidatingWebhookList represents a map of ValidatingWebhookApplyConfiguration.
type ValidatingWebhookMap map[string]ValidatingWebhookApplyConfiguration
