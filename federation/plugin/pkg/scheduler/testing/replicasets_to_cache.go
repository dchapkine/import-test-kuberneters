/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package schedulercache

import (
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/federation/plugin/pkg/scheduler/schedulercache"
)

// ReplicaSetsToCache is used for testing
type ReplicaSetsToCache []*extensions.ReplicaSet

func (p ReplicaSetsToCache) AssumeReplicaSetIfBindSucceed(pod *extensions.ReplicaSet, bind func() bool) error {
	if !bind() {
		return nil
	}
	return nil
}

func (p ReplicaSetsToCache) AddReplicaSet(pod *extensions.ReplicaSet) error { return nil }

func (p ReplicaSetsToCache) UpdateReplicaSet(oldReplicaSet, newReplicaSet *extensions.ReplicaSet) error { return nil }

func (p ReplicaSetsToCache) RemoveReplicaSet(pod *extensions.ReplicaSet) error { return nil }

func (p ReplicaSetsToCache) GetClusterNameToInfoMap() (map[string]*schedulercache.ClusterInfo, error) {
	return schedulercache.CreateClusterNameToInfoMap(p), nil
}

func (p ReplicaSetsToCache) List(s labels.Selector) (selected []*extensions.ReplicaSet, err error) {
	for _, pod := range p {
		if s.Matches(labels.Set(pod.Labels)) {
			selected = append(selected, pod)
		}
	}
	return selected, nil
}
