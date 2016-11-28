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

package install

import (
	"k8s.io/kubernetes/apifederator/pkg/apis/apifederation"
	"k8s.io/kubernetes/apifederator/pkg/apis/apifederation/v1alpha1"
	"k8s.io/kubernetes/pkg/apimachinery/announced"
	"k8s.io/kubernetes/pkg/util/sets"
)

func init() {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  apifederation.GroupName,
			RootScopedKinds:            sets.NewString("APIServer"),
			VersionPreferenceOrder:     []string{v1alpha1.SchemeGroupVersion.Version},
			ImportPrefix:               "k8s.io/kubernetes/apifederator/pkg/apis/apifederation",
			AddInternalObjectsToScheme: apifederation.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			v1alpha1.SchemeGroupVersion.Version: v1alpha1.AddToScheme,
		},
	).Announce().RegisterAndEnable(); err != nil {
		panic(err)
	}
}
