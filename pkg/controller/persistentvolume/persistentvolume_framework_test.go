/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package persistentvolume

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"testing"

	"github.com/golang/glog"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/api/testapi"
	"k8s.io/kubernetes/pkg/client/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/fake"
	"k8s.io/kubernetes/pkg/client/testing/core"
	"k8s.io/kubernetes/pkg/controller/framework"
	"k8s.io/kubernetes/pkg/conversion"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/types"
	"k8s.io/kubernetes/pkg/util/diff"
)

// This is a unit test framework for persistent volume controller.
// It fills the controller with test claims/volumes and can simulate these
// scenarios:
// 1) Call syncClaim/syncVolume once.
// 2) Call syncClaim/syncVolume several times (both simulating "claim/volume
//    modified" events and periodic sync), until the controller settles down and
//    does not modify anything.
// 3) Simulate almost real API server/etcd and call add/update/delete
//    volume/claim.
// In all these scenarios, when the test finishes, the framework can compare
// resulting claims/volumes with list of expected claims/volumes and report
// differences.

// controllerTest contains a single controller test input.
// Each test has initial set of volumes and claims that are filled into the
// controller before the test starts. The test then contains a reference to
// function to call as the actual test. Available functions are:
//   - testSyncClaim - calls syncClaim on the first claim in initialClaims.
//   - testSyncClaimError - calls syncClaim on the first claim in initialClaims
//                          and expects an error to be returned.
//   - testSyncVolume - calls syncVolume on the first volume in initialVolumes.
//   - any custom function for specialized tests.
// The test then contains list of volumes/claims that are expected at the end
// of the test.
type controllerTest struct {
	// Name of the test, for logging
	name string
	// Initial content of controller volume cache.
	initialVolumes []*api.PersistentVolume
	// Expected content of controller volume cache at the end of the test.
	expectedVolumes []*api.PersistentVolume
	// Initial content of controller claim cache.
	initialClaims []*api.PersistentVolumeClaim
	// Expected content of controller claim cache at the end of the test.
	expectedClaims []*api.PersistentVolumeClaim
	// Function to call as the test.
	test testCall
}

type testCall func(ctrl *PersistentVolumeController, reactor *volumeReactor, test controllerTest) error

const testNamespace = "default"
const versionConflictError = "VersionError"

var noVolumes []*api.PersistentVolume
var noClaims []*api.PersistentVolumeClaim

// volumeReactor is a core.Reactor that simulates etcd and API server. It
// stores:
// - Latest version of claims volumes saved by the controller.
// - Queue of all saves (to simulate "volume/claim updated" events). This queue
//   contains all intermediate state of an object - e.g. a claim.VolumeName
//   is updated first and claim.Phase second. This queue will then contain both
//   updates as separate entries.
// - Number of changes since the last call to volumeReactor.syncAll().
// - Optionally, volume and claim event sources. When set, all changed
//   volumes/claims are sent as Modify event to these sources. These sources can
//   be linked back to the controller watcher as "volume/claim updated" events.
type volumeReactor struct {
	volumes              map[string]*api.PersistentVolume
	claims               map[string]*api.PersistentVolumeClaim
	changedObjects       []interface{}
	changedSinceLastSync int
	ctrl                 *PersistentVolumeController
	volumeSource         *framework.FakeControllerSource
	claimSource          *framework.FakeControllerSource
	lock                 sync.Mutex
}

// React is a callback called by fake kubeClient from the controller.
// In other words, every claim/volume change performed by the controller ends
// here.
// This callback checks versions of the updated objects and refuse those that
// are too old (simulating real etcd).
// All updated objects are stored locally to keep track of object versions and
// to evaluate test results.
// All updated objects are also inserted into changedObjects queue and
// optionally sent back to the controller via its watchers.
func (r *volumeReactor) React(action core.Action) (handled bool, ret runtime.Object, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	glog.V(4).Infof("reactor got operation %q on %q", action.GetVerb(), action.GetResource())

	switch {
	case action.Matches("update", "persistentvolumes"):
		obj := action.(core.UpdateAction).GetObject()
		volume := obj.(*api.PersistentVolume)

		// Check and bump object version
		storedVolume, found := r.volumes[volume.Name]
		if found {
			storedVer, _ := strconv.Atoi(storedVolume.ResourceVersion)
			requestedVer, _ := strconv.Atoi(volume.ResourceVersion)
			if storedVer > requestedVer {
				return true, obj, fmt.Errorf(versionConflictError)
			}
			volume.ResourceVersion = strconv.Itoa(storedVer + 1)
		}

		// Store the updated object to appropriate places.
		if r.volumeSource != nil {
			r.volumeSource.Modify(volume)
		}
		r.volumes[volume.Name] = volume
		r.changedObjects = append(r.changedObjects, volume)
		r.changedSinceLastSync++
		glog.V(4).Infof("saved updated volume %s", volume.Name)
		return true, volume, nil

	case action.Matches("update", "persistentvolumeclaims"):
		obj := action.(core.UpdateAction).GetObject()
		claim := obj.(*api.PersistentVolumeClaim)

		// Check and bump object version
		storedClaim, found := r.claims[claim.Name]
		if found {
			storedVer, _ := strconv.Atoi(storedClaim.ResourceVersion)
			requestedVer, _ := strconv.Atoi(claim.ResourceVersion)
			if storedVer > requestedVer {
				return true, obj, fmt.Errorf("VersionError")
			}
			claim.ResourceVersion = strconv.Itoa(storedVer + 1)
		}

		// Store the updated object to appropriate places.
		r.claims[claim.Name] = claim
		if r.claimSource != nil {
			r.claimSource.Modify(claim)
		}
		r.changedObjects = append(r.changedObjects, claim)
		r.changedSinceLastSync++
		glog.V(4).Infof("saved updated claim %s", claim.Name)
		return true, claim, nil
	}
	return false, nil, nil
}

// checkVolumes compares all expectedVolumes with set of volumes at the end of
// the test and reports differences.
func (r *volumeReactor) checkVolumes(t *testing.T, expectedVolumes []*api.PersistentVolume) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	expectedMap := make(map[string]*api.PersistentVolume)
	gotMap := make(map[string]*api.PersistentVolume)
	// Clear any ResourceVersion from both sets
	for _, v := range expectedVolumes {
		v.ResourceVersion = ""
		expectedMap[v.Name] = v
	}
	for _, v := range r.volumes {
		// We must clone the volume because of golang race check - it was
		// written by the controller without any locks on it.
		clone, _ := conversion.NewCloner().DeepCopy(v)
		v = clone.(*api.PersistentVolume)
		v.ResourceVersion = ""
		if v.Spec.ClaimRef != nil {
			v.Spec.ClaimRef.ResourceVersion = ""
		}
		gotMap[v.Name] = v
	}
	if !reflect.DeepEqual(expectedMap, gotMap) {
		// Print ugly but useful diff of expected and received objects for
		// easier debugging.
		return fmt.Errorf("Volume check failed [A-expected, B-got]: %s", diff.ObjectDiff(expectedMap, gotMap))
	}
	return nil
}

// checkClaims compares all expectedClaims with set of claims at the end of the
// test and reports differences.
func (r *volumeReactor) checkClaims(t *testing.T, expectedClaims []*api.PersistentVolumeClaim) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	expectedMap := make(map[string]*api.PersistentVolumeClaim)
	gotMap := make(map[string]*api.PersistentVolumeClaim)
	for _, c := range expectedClaims {
		c.ResourceVersion = ""
		expectedMap[c.Name] = c
	}
	for _, c := range r.claims {
		// We must clone the claim because of golang race check - it was
		// written by the controller without any locks on it.
		clone, _ := conversion.NewCloner().DeepCopy(c)
		c = clone.(*api.PersistentVolumeClaim)
		c.ResourceVersion = ""
		gotMap[c.Name] = c
	}
	if !reflect.DeepEqual(expectedMap, gotMap) {
		// Print ugly but useful diff of expected and received objects for
		// easier debugging.
		return fmt.Errorf("Claim check failed [A-expected, B-got result]: %s", diff.ObjectDiff(expectedMap, gotMap))
	}
	return nil
}

// popChange returns one recorded updated object, either *api.PersistentVolume
// or *api.PersistentVolumeClaim. Returns nil when there are no changes.
func (r *volumeReactor) popChange() interface{} {
	r.lock.Lock()
	defer r.lock.Unlock()

	if len(r.changedObjects) == 0 {
		return nil
	}

	// For debugging purposes, print the queue
	for _, obj := range r.changedObjects {
		switch obj.(type) {
		case *api.PersistentVolume:
			vol, _ := obj.(*api.PersistentVolume)
			glog.V(4).Infof("reactor queue: %s", vol.Name)
		case *api.PersistentVolumeClaim:
			claim, _ := obj.(*api.PersistentVolumeClaim)
			glog.V(4).Infof("reactor queue: %s", claim.Name)
		}
	}

	// Pop the first item from the queue and return it
	obj := r.changedObjects[0]
	r.changedObjects = r.changedObjects[1:]
	return obj
}

// syncAll simulates the controller periodic sync of volumes and claim. It
// simply adds all these objects to the internal queue of updates. This method
// should be used when the test manually calls syncClaim/syncVolume. Test that
// use real controller loop (ctrl.Run()) will get periodic sync automatically.
func (r *volumeReactor) syncAll() {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, c := range r.claims {
		r.changedObjects = append(r.changedObjects, c)
	}
	for _, v := range r.volumes {
		r.changedObjects = append(r.changedObjects, v)
	}
	r.changedSinceLastSync = 0
}

func (r *volumeReactor) getChangeCount() int {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.changedSinceLastSync
}

func newVolumeReactor(client *fake.Clientset, ctrl *PersistentVolumeController, volumeSource, claimSource *framework.FakeControllerSource) *volumeReactor {
	reactor := &volumeReactor{
		volumes:      make(map[string]*api.PersistentVolume),
		claims:       make(map[string]*api.PersistentVolumeClaim),
		ctrl:         ctrl,
		volumeSource: volumeSource,
		claimSource:  claimSource,
	}
	client.AddReactor("*", "*", reactor.React)
	return reactor
}

func newPersistentVolumeController(kubeClient clientset.Interface) *PersistentVolumeController {
	ctrl := &PersistentVolumeController{
		volumes:    newPersistentVolumeOrderedIndex(),
		claims:     cache.NewStore(cache.MetaNamespaceKeyFunc),
		kubeClient: kubeClient,
	}
	return ctrl
}

// newVolume returns a new volume with given attributes
func newVolume(name, capacity, boundToClaimUID, boundToClaimName string, phase api.PersistentVolumePhase, annotations ...string) *api.PersistentVolume {
	volume := api.PersistentVolume{
		ObjectMeta: api.ObjectMeta{
			Name:            name,
			ResourceVersion: "1",
		},
		Spec: api.PersistentVolumeSpec{
			Capacity: api.ResourceList{
				api.ResourceName(api.ResourceStorage): resource.MustParse(capacity),
			},
			PersistentVolumeSource: api.PersistentVolumeSource{
				GCEPersistentDisk: &api.GCEPersistentDiskVolumeSource{},
			},
			AccessModes: []api.PersistentVolumeAccessMode{api.ReadWriteOnce, api.ReadOnlyMany},
		},
		Status: api.PersistentVolumeStatus{
			Phase: phase,
		},
	}

	if boundToClaimName != "" {
		volume.Spec.ClaimRef = &api.ObjectReference{
			Kind:       "PersistentVolumeClaim",
			APIVersion: "v1",
			UID:        types.UID(boundToClaimUID),
			Namespace:  testNamespace,
			Name:       boundToClaimName,
		}
	}

	if len(annotations) > 0 {
		volume.Annotations = make(map[string]string)
		for _, a := range annotations {
			volume.Annotations[a] = "yes"
		}
	}

	return &volume
}

// newVolumeArray returns array with a single volume that would be returned by
// newVolume() with the same parameters.
func newVolumeArray(name, capacity, boundToClaimUID, boundToClaimName string, phase api.PersistentVolumePhase, annotations ...string) []*api.PersistentVolume {
	return []*api.PersistentVolume{
		newVolume(name, capacity, boundToClaimUID, boundToClaimName, phase, annotations...),
	}
}

// newClaim returns a new claim with given attributes
func newClaim(name, claimUID, capacity, boundToVolume string, phase api.PersistentVolumeClaimPhase, annotations ...string) *api.PersistentVolumeClaim {
	claim := api.PersistentVolumeClaim{
		ObjectMeta: api.ObjectMeta{
			Name:            name,
			Namespace:       testNamespace,
			UID:             types.UID(claimUID),
			ResourceVersion: "1",
		},
		Spec: api.PersistentVolumeClaimSpec{
			AccessModes: []api.PersistentVolumeAccessMode{api.ReadWriteOnce, api.ReadOnlyMany},
			Resources: api.ResourceRequirements{
				Requests: api.ResourceList{
					api.ResourceName(api.ResourceStorage): resource.MustParse(capacity),
				},
			},
			VolumeName: boundToVolume,
		},
		Status: api.PersistentVolumeClaimStatus{
			Phase: phase,
		},
	}
	// Make sure api.GetReference(claim) works
	claim.ObjectMeta.SelfLink = testapi.Default.SelfLink("pvc", name)

	if len(annotations) > 0 {
		claim.Annotations = make(map[string]string)
		for _, a := range annotations {
			claim.Annotations[a] = "yes"
		}
	}
	return &claim
}

// newClaimArray returns array with a single claim that would be returned by
// newClaim() with the same parameters.
func newClaimArray(name, claimUID, capacity, boundToVolume string, phase api.PersistentVolumeClaimPhase, annotations ...string) []*api.PersistentVolumeClaim {
	return []*api.PersistentVolumeClaim{
		newClaim(name, claimUID, capacity, boundToVolume, phase, annotations...),
	}
}

func testSyncClaim(ctrl *PersistentVolumeController, reactor *volumeReactor, test controllerTest) error {
	return ctrl.syncClaim(test.initialClaims[0])
}

func testSyncClaimError(ctrl *PersistentVolumeController, reactor *volumeReactor, test controllerTest) error {
	err := ctrl.syncClaim(test.initialClaims[0])

	if err != nil {
		return nil
	}
	return fmt.Errorf("syncClaim succeeded when failure was expected")
}

func testSyncVolume(ctrl *PersistentVolumeController, reactor *volumeReactor, test controllerTest) error {
	return ctrl.syncVolume(test.initialVolumes[0])
}

func evaluateTestResults(ctrl *PersistentVolumeController, reactor *volumeReactor, test controllerTest, t *testing.T) {
	// Evaluate results
	if err := reactor.checkClaims(t, test.expectedClaims); err != nil {
		t.Errorf("Test %q: %v", test.name, err)

	}
	if err := reactor.checkVolumes(t, test.expectedVolumes); err != nil {
		t.Errorf("Test %q: %v", test.name, err)
	}
}

// Test single call to syncClaim and syncVolume methods.
// For all tests:
// 1. Fill in the controller with initial data
// 2. Call the tested function (syncClaim/syncVolume) via
//    controllerTest.testCall *once*.
// 3. Compare resulting volumes and claims with expected volumes and claims.
func runSyncTests(t *testing.T, tests []controllerTest) {
	for _, test := range tests {
		glog.V(4).Infof("starting test %q", test.name)

		// Initialize the controller
		client := &fake.Clientset{}
		ctrl := newPersistentVolumeController(client)
		reactor := newVolumeReactor(client, ctrl, nil, nil)
		for _, claim := range test.initialClaims {
			ctrl.claims.Add(claim)
			reactor.claims[claim.Name] = claim
		}
		for _, volume := range test.initialVolumes {
			ctrl.volumes.store.Add(volume)
			reactor.volumes[volume.Name] = volume
		}

		// Run the tested functions
		err := test.test(ctrl, reactor, test)
		if err != nil {
			t.Errorf("Test %q failed: %v", test.name, err)
		}

		evaluateTestResults(ctrl, reactor, test, t)
	}
}
