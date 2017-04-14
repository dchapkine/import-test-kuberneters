/*
Copyright 2016 The Kubernetes Authors.

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

package stubs

import (
	"google.golang.org/api/googleapi"
	"k8s.io/kubernetes/federation/pkg/dnsprovider/providers/google/clouddns/internal/interfaces"
	"strings"
)

// Compile time check for interface adherence
var _ interfaces.ResourceRecordSetsListCall = &ResourceRecordSetsListCall{}

type ResourceRecordSetsListCall struct {
	Response_   *ResourceRecordSetsListResponse
	Err_        error
	Name_       string
	Type_       string
	MaxResults_ int64
}

func (call *ResourceRecordSetsListCall) Do(opts ...googleapi.CallOption) (interfaces.ResourceRecordSetsListResponse, error) {
	if len(call.Name_) > 0 {
		records := make([]interfaces.ResourceRecordSet, 0)
		for _, record := range call.Response_.impl {
			if strings.TrimSuffix(call.Name_, ".") == strings.TrimSuffix(record.Name(), ".") {
				records = append(records, record)
			}
		}
		call.Response_.impl = records
	}
	return call.Response_, call.Err_
}

func (call *ResourceRecordSetsListCall) Name(name string) interfaces.ResourceRecordSetsListCall {
	call.Name_ = name
	return call
}

func (call *ResourceRecordSetsListCall) MaxResults(maxResults int64) interfaces.ResourceRecordSetsListCall {
	call.MaxResults_ = maxResults
	return call
}

func (call *ResourceRecordSetsListCall) Type(type_ string) interfaces.ResourceRecordSetsListCall {
	call.Type_ = type_
	return call
}
