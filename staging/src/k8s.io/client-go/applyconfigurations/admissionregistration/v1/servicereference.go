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

// ServiceReferenceApplyConfiguration represents an declarative configuration of the ServiceReference type for use
// with apply.
type ServiceReferenceApplyConfiguration struct {
	Namespace *string `json:"namespace,omitempty"`
	Name      *string `json:"name,omitempty"`
	Path      *string `json:"path,omitempty"`
	Port      *int32  `json:"port,omitempty"`
}

// ServiceReferenceList represents a listAlias of ServiceReferenceApplyConfiguration.
type ServiceReferenceList []*ServiceReferenceApplyConfiguration

// ServiceReferenceList represents a map of ServiceReferenceApplyConfiguration.
type ServiceReferenceMap map[string]ServiceReferenceApplyConfiguration
