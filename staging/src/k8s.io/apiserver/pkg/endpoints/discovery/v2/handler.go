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

package v2

import (
	"net/http"
	"sync"

	"github.com/emicklei/go-restful/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/handlers/negotiation"
	"k8s.io/klog/v2"
)

const DiscoveryEndpointRoot = "/discovery"

// This handler serves the /discovery/<version> endpoint for a given list of
// api resources indexed by their group version.
type ResourceManager interface {
	// Adds knowledge of the given groupversion to the discovery document
	// If it was already being tracked, updates the stored APIVersionDiscovery
	// Thread-safe
	AddGroupVersion(groupName string, value metav1.APIVersionDiscovery)

	// Removes all group versions for a given group
	// Thread-safe
	RemoveGroup(groupName string)

	// Removes a specific groupversion. If all versions of a group have been
	// removed, then the entire group is unlisted.
	// Thread-safe
	RemoveGroupVersion(gv metav1.GroupVersion)

	// Resets the manager's known list of group-versions and replaces them
	// with the given groups
	// Thread-Safe
	SetGroups([]metav1.APIGroupDiscovery)

	// Returns a restful webservice which responds to discovery requests
	// Thread-safe
	WebService() *restful.WebService

	http.Handler
}

type noopResourceManager struct{}

func (noopResourceManager) AddGroupVersion(groupName string, value metav1.APIVersionDiscovery) {}
func (noopResourceManager) RemoveGroup(groupName string)                                       {}
func (noopResourceManager) RemoveGroupVersion(gv metav1.GroupVersion)                          {}
func (noopResourceManager) SetGroups([]metav1.APIGroupDiscovery)                               {}
func (noopResourceManager) WebService() *restful.WebService                                    { return nil }
func (noopResourceManager) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusNotImplemented)
}

var _ ResourceManager = noopResourceManager{}

func NewNoopResourceManager() ResourceManager {
	return noopResourceManager{}
}

type resourceDiscoveryManager struct {
	// Protects writes to all fields in struct
	lock sync.RWMutex

	// Writes protected by the lock. List if all apigroups & resources indexed
	// by the resource manager
	apiGroups     map[string]*metav1.APIGroupDiscovery
	apiGroupNames []string // apiGroupNames preserves insertion order

	serializer runtime.NegotiatedSerializer
	cachedGroupList
}

type cachedGroupList struct {
	// Most up to date version of all discovery api groups
	cachedResponse *metav1.APIGroupDiscoveryList
	// Hash of the cachedResponse used for cache-busting
	cachedResponseETag string
}

func NewResourceManager(serializer runtime.NegotiatedSerializer) ResourceManager {
	result := &resourceDiscoveryManager{serializer: serializer}
	return result
}

func (rdm *resourceDiscoveryManager) SetGroups(groups []metav1.APIGroupDiscovery) {
	rdm.lock.Lock()
	defer rdm.lock.Unlock()

	rdm.apiGroups = nil
	rdm.apiGroupNames = nil
	rdm.cachedGroupList = cachedGroupList{}

	for _, group := range groups {
		for _, version := range group.Versions {
			rdm.addGroupVersionLocked(group.Name, version)
		}
	}
}

func (rdm *resourceDiscoveryManager) AddGroups(groups ...metav1.APIGroupDiscovery) {
	rdm.lock.Lock()
	defer rdm.lock.Unlock()

	for _, group := range groups {
		for _, version := range group.Versions {
			rdm.addGroupVersionLocked(group.Name, version)
		}
	}
}

func (rdm *resourceDiscoveryManager) AddGroupVersion(groupName string, value metav1.APIVersionDiscovery) {
	rdm.lock.Lock()
	defer rdm.lock.Unlock()

	rdm.addGroupVersionLocked(groupName, value)
}

func (rdm *resourceDiscoveryManager) addGroupVersionLocked(groupName string, value metav1.APIVersionDiscovery) {
	klog.Infof("Adding GroupVersion %s %s to ResourceManager", groupName, value.Version)

	if rdm.apiGroups == nil {
		rdm.apiGroups = make(map[string]*metav1.APIGroupDiscovery)
	}

	if existing, groupExists := rdm.apiGroups[groupName]; groupExists {
		// If this version already exists, replace it
		versionExists := false

		// Not very efficient, but in practice there are generally not many versions
		for i := range existing.Versions {
			if existing.Versions[i].Version == value.Version {
				existing.Versions[i] = value
				versionExists = true
				break
			}
		}

		if !versionExists {
			existing.Versions = append(existing.Versions, value)
		}

	} else {
		rdm.apiGroups[groupName] = &metav1.APIGroupDiscovery{
			TypeMeta: metav1.TypeMeta{
				Kind:       "APIGroupDiscovery",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: groupName,
			},
			Versions: []metav1.APIVersionDiscovery{value},
		}
		rdm.apiGroupNames = append(rdm.apiGroupNames, groupName)
	}

	// Reset response document so it is recreated lazily
	rdm.cachedGroupList = cachedGroupList{}
}

func (rdm *resourceDiscoveryManager) RemoveGroupVersion(apiGroup metav1.GroupVersion) {
	rdm.lock.Lock()
	defer rdm.lock.Unlock()

	group, exists := rdm.apiGroups[apiGroup.Group]
	if !exists {
		return
	}

	for i := range group.Versions {
		if group.Versions[i].Version == apiGroup.Version {
			group.Versions = append(group.Versions[:i], group.Versions[i+1:]...)
			break
		}
	}

	if len(group.Versions) == 0 {
		delete(rdm.apiGroups, group.Name)
		for i := range rdm.apiGroupNames {
			if rdm.apiGroupNames[i] == group.Name {
				rdm.apiGroupNames = append(rdm.apiGroupNames[:i], rdm.apiGroupNames[i+1:]...)
				break
			}
		}
	}

	// Reset response document so it is recreated lazily
	rdm.cachedGroupList = cachedGroupList{}
}

func (rdm *resourceDiscoveryManager) RemoveGroup(groupName string) {
	rdm.lock.Lock()
	defer rdm.lock.Unlock()

	delete(rdm.apiGroups, groupName)
	for i := range rdm.apiGroupNames {
		if rdm.apiGroupNames[i] == groupName {
			rdm.apiGroupNames = append(rdm.apiGroupNames[:i], rdm.apiGroupNames[i+1:]...)
			break
		}
	}

	// Reset response document so it is recreated lazily
	rdm.cachedGroupList = cachedGroupList{}
}

func (rdm *resourceDiscoveryManager) WebService() *restful.WebService {
	mediaTypes, _ := negotiation.MediaTypesForSerializer(rdm.serializer)
	ws := new(restful.WebService)
	ws.Path(DiscoveryEndpointRoot)
	ws.Doc("get available API groupversions and resources")

	ws.Route(ws.GET("/v2").To(func(req *restful.Request, resp *restful.Response) {
		rdm.ServeHTTP(resp.ResponseWriter, req.Request)
	}).
		Doc("get available API groupversions and their resources").
		Operation("getDiscoveryResources").
		Produces(mediaTypes...).
		Consumes(mediaTypes...).
		Writes(metav1.APIGroupDiscoveryList{}))
	return ws
}

func (rdm *resourceDiscoveryManager) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	rdm.lock.RLock()
	response := rdm.cachedResponse
	etag := rdm.cachedResponseETag

	// The cachedResponse is wiped out every time there might be a change to its
	// contents.
	if response != nil {
		defer rdm.lock.RUnlock()
	} else {
		rdm.lock.RUnlock()

		// Document does not exist, recreate it
		rdm.lock.Lock()
		defer rdm.lock.Unlock()

		// Now that we have taken exclusive lock, check to see if another thread
		// recreated the document while we were waiting for the lock
		response, etag = rdm.cachedResponse, rdm.cachedResponseETag
		if response == nil {
			// Re-order the apiGroups by their insertion order
			orderedGroups := []metav1.APIGroupDiscovery{}
			for _, groupName := range rdm.apiGroupNames {
				orderedGroups = append(orderedGroups, *rdm.apiGroups[groupName])
			}

			var err error
			response = &metav1.APIGroupDiscoveryList{
				TypeMeta: metav1.TypeMeta{
					Kind:       "APIGroupDiscoveryList",
					APIVersion: "v1",
				},
				Groups: orderedGroups,
			}
			etag, err = CalculateETag(response)

			if err != nil {
				klog.Errorf("failed to calculate etag for discovery document: %s", etag)
			}

			rdm.cachedResponse = response
			rdm.cachedResponseETag = etag
		}

	}

	ServeHTTPWithETag(
		response,
		etag,
		rdm.serializer,
		resp,
		req,
	)
}
