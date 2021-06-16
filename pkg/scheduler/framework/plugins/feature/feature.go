/*
Copyright 2021 The Kubernetes Authors.

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

package feature

// Features carries feature gate values used by various plugins.
// This struct allows us to break the dependency of the plugins on
// the internal k8s features pkg.
type Features struct {
	EnableCSIMigration                  bool // Default: true
	EnableCSIMigrationAWS               bool
	EnableCSIMigrationGCE               bool
	EnableCSIMigrationAzureDisk         bool
	EnableCSIMigrationOpenStack         bool // Default: true
	EnableCSIStorageCapacity            bool // Default: true
	EnableDefaultPodTopologySpread      bool // Default: true
	EnableGenericEphemeralVolume        bool // Default: true
	EnableLocalStorageCapacityIsolation bool // Default: true
	EnablePodAffinityNamespaceSelector  bool // Default: true
	EnablePodDisruptionBudget           bool // Default: true
	EnablePodOverhead                   bool // Default: true
	EnableReadWriteOncePod              bool
	EnableStorageObjectInUseProtection  bool // Default: true
	EnableVolumeCapacityPriority        bool
}
