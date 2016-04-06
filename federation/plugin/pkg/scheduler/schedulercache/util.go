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
	"k8s.io/kubernetes/federation/apis/federation"
)

// CreateClusterNameToInfoMap obtains a list of pods and pivots that list into a map where the keys are cluster names
// and the values are the aggregated information for that cluster.
func CreateClusterNameToInfoMap(replicaSets []*extensions.ReplicaSet) map[string]*ClusterInfo {
	clusterNameToInfo := make(map[string]*ClusterInfo)
	for _, replicasSet := range replicaSets {
		clusterName := replicasSet.Annotations[federation.TargetClusterKey]
		clusterInfo, ok := clusterNameToInfo[clusterName]
		if !ok {
			clusterInfo = NewClusterInfo()
			clusterNameToInfo[clusterName] = clusterInfo
		}
		clusterInfo.addReplicaSet(replicasSet)
	}
	return clusterNameToInfo
}
