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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// EventApplyConfiguration represents an declarative configuration of the Event type for use
// with apply.
type EventApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration `json:",inline"`
	ObjectMeta                    *v1.ObjectMetaApplyConfiguration          `json:"metadata,omitempty"`
	EventTime                     *metav1.MicroTime                         `json:"eventTime,omitempty"`
	Series                        *EventSeriesApplyConfiguration            `json:"series,omitempty"`
	ReportingController           *string                                   `json:"reportingController,omitempty"`
	ReportingInstance             *string                                   `json:"reportingInstance,omitempty"`
	Action                        *string                                   `json:"action,omitempty"`
	Reason                        *string                                   `json:"reason,omitempty"`
	Regarding                     *corev1.ObjectReferenceApplyConfiguration `json:"regarding,omitempty"`
	Related                       *corev1.ObjectReferenceApplyConfiguration `json:"related,omitempty"`
	Note                          *string                                   `json:"note,omitempty"`
	Type                          *string                                   `json:"type,omitempty"`
	DeprecatedSource              *corev1.EventSourceApplyConfiguration     `json:"deprecatedSource,omitempty"`
	DeprecatedFirstTimestamp      *metav1.Time                              `json:"deprecatedFirstTimestamp,omitempty"`
	DeprecatedLastTimestamp       *metav1.Time                              `json:"deprecatedLastTimestamp,omitempty"`
	DeprecatedCount               *int32                                    `json:"deprecatedCount,omitempty"`
}

// EventList represents a listAlias of EventApplyConfiguration.
type EventList []*EventApplyConfiguration

// EventList represents a map of EventApplyConfiguration.
type EventMap map[string]EventApplyConfiguration
