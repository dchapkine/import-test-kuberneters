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
	types "k8s.io/apimachinery/pkg/types"
)

// ObjectReferenceApplyConfiguration represents an declarative configuration of the ObjectReference type for use
// with apply.
type ObjectReferenceApplyConfiguration struct {
	Kind            *string    `json:"kind,omitempty"`
	Namespace       *string    `json:"namespace,omitempty"`
	Name            *string    `json:"name,omitempty"`
	UID             *types.UID `json:"uid,omitempty"`
	APIVersion      *string    `json:"apiVersion,omitempty"`
	ResourceVersion *string    `json:"resourceVersion,omitempty"`
	FieldPath       *string    `json:"fieldPath,omitempty"`
}

// ObjectReferenceList represents a listAlias of ObjectReferenceApplyConfiguration.
type ObjectReferenceList []*ObjectReferenceApplyConfiguration

// ObjectReferenceList represents a map of ObjectReferenceApplyConfiguration.
type ObjectReferenceMap map[string]ObjectReferenceApplyConfiguration
