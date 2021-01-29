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

// StatefulSetStatusApplyConfiguration represents an declarative configuration of the StatefulSetStatus type for use
// with apply.
type StatefulSetStatusApplyConfiguration struct {
	ObservedGeneration *int64                    `json:"observedGeneration,omitempty"`
	Replicas           *int32                    `json:"replicas,omitempty"`
	ReadyReplicas      *int32                    `json:"readyReplicas,omitempty"`
	CurrentReplicas    *int32                    `json:"currentReplicas,omitempty"`
	UpdatedReplicas    *int32                    `json:"updatedReplicas,omitempty"`
	CurrentRevision    *string                   `json:"currentRevision,omitempty"`
	UpdateRevision     *string                   `json:"updateRevision,omitempty"`
	CollisionCount     *int32                    `json:"collisionCount,omitempty"`
	Conditions         *StatefulSetConditionList `json:"conditions,omitempty"`
}

// StatefulSetStatusList represents a listAlias of StatefulSetStatusApplyConfiguration.
type StatefulSetStatusList []*StatefulSetStatusApplyConfiguration

// StatefulSetStatusList represents a map of StatefulSetStatusApplyConfiguration.
type StatefulSetStatusMap map[string]StatefulSetStatusApplyConfiguration
