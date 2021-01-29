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
	v1beta1 "k8s.io/api/flowcontrol/v1beta1"
)

// SubjectApplyConfiguration represents an declarative configuration of the Subject type for use
// with apply.
type SubjectApplyConfiguration struct {
	Kind           *v1beta1.SubjectKind                     `json:"kind,omitempty"`
	User           *UserSubjectApplyConfiguration           `json:"user,omitempty"`
	Group          *GroupSubjectApplyConfiguration          `json:"group,omitempty"`
	ServiceAccount *ServiceAccountSubjectApplyConfiguration `json:"serviceAccount,omitempty"`
}

// SubjectList represents a listAlias of SubjectApplyConfiguration.
type SubjectList []*SubjectApplyConfiguration

// SubjectList represents a map of SubjectApplyConfiguration.
type SubjectMap map[string]SubjectApplyConfiguration
