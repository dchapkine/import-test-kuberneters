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
	authorizationv1beta1 "k8s.io/api/authorization/v1beta1"
)

// SubjectAccessReviewSpecApplyConfiguration represents an declarative configuration of the SubjectAccessReviewSpec type for use
// with apply.
type SubjectAccessReviewSpecApplyConfiguration struct {
	ResourceAttributes    *ResourceAttributesApplyConfiguration       `json:"resourceAttributes,omitempty"`
	NonResourceAttributes *NonResourceAttributesApplyConfiguration    `json:"nonResourceAttributes,omitempty"`
	User                  *string                                     `json:"user,omitempty"`
	Groups                *[]string                                   `json:"group,omitempty"`
	Extra                 *map[string]authorizationv1beta1.ExtraValue `json:"extra,omitempty"`
	UID                   *string                                     `json:"uid,omitempty"`
}

// SubjectAccessReviewSpecList represents a listAlias of SubjectAccessReviewSpecApplyConfiguration.
type SubjectAccessReviewSpecList []*SubjectAccessReviewSpecApplyConfiguration

// SubjectAccessReviewSpecList represents a map of SubjectAccessReviewSpecApplyConfiguration.
type SubjectAccessReviewSpecMap map[string]SubjectAccessReviewSpecApplyConfiguration
