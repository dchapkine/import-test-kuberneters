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

package v1beta2

import (
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// DaemonSetApplyConfiguration represents an declarative configuration of the DaemonSet type for use
// with apply.
type DaemonSetApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration   `json:"metadata,omitempty"`
	Spec                          *DaemonSetSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                        *DaemonSetStatusApplyConfiguration `json:"status,omitempty"`
}

// DaemonSetList represents a listAlias of DaemonSetApplyConfiguration.
type DaemonSetList []*DaemonSetApplyConfiguration

// DaemonSetList represents a map of DaemonSetApplyConfiguration.
type DaemonSetMap map[string]DaemonSetApplyConfiguration
