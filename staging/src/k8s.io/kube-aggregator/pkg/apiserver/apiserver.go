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

package apiserver

import (
	"time"

	"k8s.io/apimachinery/pkg/util/sets"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	kubeinformers "k8s.io/client-go/informers"
	kubeclientset "k8s.io/client-go/kubernetes"
	v1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/version"

	"k8s.io/kube-aggregator/pkg/apis/apiregistration"
	"k8s.io/kube-aggregator/pkg/apis/apiregistration/v1alpha1"
	"k8s.io/kube-aggregator/pkg/client/clientset_generated/internalclientset"
	informers "k8s.io/kube-aggregator/pkg/client/informers/internalversion"
	listers "k8s.io/kube-aggregator/pkg/client/listers/apiregistration/internalversion"
	apiservicestorage "k8s.io/kube-aggregator/pkg/registry/apiservice/etcd"

	_ "k8s.io/client-go/pkg/api/install"
)

// legacyAPIServiceName is the fixed name of the only non-groupified API version
const legacyAPIServiceName = "v1."

type Config struct {
	GenericConfig       *genericapiserver.Config
	CoreAPIServerClient kubeclientset.Interface

	// ProxyClientCert/Key are the client cert used to identify this proxy. Backing APIServices use
	// this to confirm the proxy's identity
	ProxyClientCert []byte
	ProxyClientKey  []byte
}

// APIAggregator contains state for a Kubernetes cluster master/api server.
type APIAggregator struct {
	GenericAPIServer *genericapiserver.GenericAPIServer

	contextMapper genericapirequest.RequestContextMapper

	// proxyClientCert/Key are the client cert used to identify this proxy. Backing APIServices use
	// this to confirm the proxy's identity
	proxyClientCert []byte
	proxyClientKey  []byte

	// proxyHandlers are the proxy handlers that are currently registered, keyed by apiservice.name
	proxyHandlers map[string]*proxyHandler
	// handledGroups are the groups that already have routes
	handledGroups sets.String

	// lister is used to add group handling for /apis/<group> aggregator lookups based on
	// controller state
	lister listers.APIServiceLister

	// serviceLister is used by the aggregator handler to determine whether or not to try to expose the group
	serviceLister v1listers.ServiceLister
	// endpointsLister is used by the aggregator handler to determine whether or not to try to expose the group
	endpointsLister v1listers.EndpointsLister
}

type completedConfig struct {
	*Config
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *Config) Complete() completedConfig {
	// generic discovery must be disabled so that we can register our own discovery handler
	c.GenericConfig.EnableDiscovery = false
	c.GenericConfig.Complete()

	version := version.Get()
	c.GenericConfig.Version = &version

	return completedConfig{c}
}

// SkipComplete provides a way to construct a server instance without config completion.
func (c *Config) SkipComplete() completedConfig {
	return completedConfig{c}
}

// New returns a new instance of APIAggregator from the given config.
func (c completedConfig) New(stopCh <-chan struct{}) (*APIAggregator, error) {
	informerFactory := informers.NewSharedInformerFactory(
		internalclientset.NewForConfigOrDie(c.Config.GenericConfig.LoopbackClientConfig),
		5*time.Minute, // this is effectively used as a refresh interval right now.  Might want to do something nicer later on.
	)
	kubeInformers := kubeinformers.NewSharedInformerFactory(c.CoreAPIServerClient, 5*time.Minute)

	genericServer, err := c.Config.GenericConfig.SkipComplete().New() // completion is done in Complete, no need for a second time
	if err != nil {
		return nil, err
	}

	s := &APIAggregator{
		GenericAPIServer: genericServer,
		contextMapper:    c.GenericConfig.RequestContextMapper,
		proxyClientCert:  c.ProxyClientCert,
		proxyClientKey:   c.ProxyClientKey,
		proxyHandlers:    map[string]*proxyHandler{},
		handledGroups:    sets.String{},
		lister:           informerFactory.Apiregistration().InternalVersion().APIServices().Lister(),
		serviceLister:    kubeInformers.Core().V1().Services().Lister(),
		endpointsLister:  kubeInformers.Core().V1().Endpoints().Lister(),
	}

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(apiregistration.GroupName, api.Registry, api.Scheme, api.ParameterCodec, api.Codecs)
	apiGroupInfo.GroupMeta.GroupVersion = v1alpha1.SchemeGroupVersion
	v1alpha1storage := map[string]rest.Storage{}
	v1alpha1storage["apiservices"] = apiservicestorage.NewREST(c.GenericConfig.RESTOptionsGetter)
	apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	apisHandler := &apisHandler{
		lister:          s.lister,
		delegate:        s.GenericAPIServer.FallThroughHandler,
		serviceLister:   s.serviceLister,
		endpointsLister: s.endpointsLister,
	}
	s.GenericAPIServer.HandlerContainer.Handle("/apis", apisHandler)

	apiserviceRegistrationController := NewAPIServiceRegistrationController(informerFactory.Apiregistration().InternalVersion().APIServices(), s)

	s.GenericAPIServer.AddPostStartHook("start-informers", func(context genericapiserver.PostStartHookContext) error {
		informerFactory.Start(stopCh)
		kubeInformers.Start(stopCh)
		return nil
	})
	s.GenericAPIServer.AddPostStartHook("apiservice-registration-controller", func(context genericapiserver.PostStartHookContext) error {
		apiserviceRegistrationController.Run(stopCh)
		return nil
	})

	return s, nil
}

// AddAPIService adds an API service.  It is not thread-safe, so only call it on one thread at a time please.
// It's a slow moving API, so its ok to run the controller on a single thread
func (s *APIAggregator) AddAPIService(apiService *apiregistration.APIService) {
	// if the proxyHandler already exists, it needs to be updated. The aggregation bits do not
	// since they are wired against listers because they require multiple resources to respond
	if proxyHandler, exists := s.proxyHandlers[apiService.Name]; exists {
		proxyHandler.updateAPIService(apiService)
		return
	}

	proxyPath := "/apis/" + apiService.Spec.Group + "/" + apiService.Spec.Version
	// v1. is a special case for the legacy API.  It proxies to a wider set of endpoints.
	if apiService.Name == "v1." {
		proxyPath = "/api"
	}

	// register the proxy handler
	proxyHandler := &proxyHandler{
		contextMapper:   s.contextMapper,
		localDelegate:   s.GenericAPIServer.FallThroughHandler,
		proxyClientCert: s.proxyClientCert,
		proxyClientKey:  s.proxyClientKey,
	}
	proxyHandler.updateAPIService(apiService)
	s.proxyHandlers[apiService.Name] = proxyHandler
	s.GenericAPIServer.FallThroughHandler.Handle(proxyPath, proxyHandler)
	s.GenericAPIServer.FallThroughHandler.Handle(proxyPath+"/", proxyHandler)

	// if we're dealing with the legacy group, we're done here
	if apiService.Name == legacyAPIServiceName {
		return
	}

	// if we've already registered the path with the handler, we don't want to do it again.
	if s.handledGroups.Has(apiService.Spec.Group) {
		return
	}

	// it's time to register the group aggregation endpoint
	groupPath := "/apis/" + apiService.Spec.Group
	groupDiscoveryHandler := &apiGroupHandler{
		groupName:       apiService.Spec.Group,
		lister:          s.lister,
		serviceLister:   s.serviceLister,
		endpointsLister: s.endpointsLister,
	}
	// aggregation is protected
	s.GenericAPIServer.FallThroughHandler.UnlistedHandle(groupPath, groupDiscoveryHandler)
	s.GenericAPIServer.FallThroughHandler.UnlistedHandle(groupPath+"/", groupDiscoveryHandler)
	s.handledGroups.Insert(apiService.Spec.Group)
}

// RemoveAPIService removes the APIService from being handled.  Later on it will disable the proxy endpoint.
// Right now it does nothing because our handler has to properly 404 itself since muxes don't unregister
func (s *APIAggregator) RemoveAPIService(apiServiceName string) {
	proxyHandler, exists := s.proxyHandlers[apiServiceName]
	if !exists {
		return
	}
	proxyHandler.removeAPIService()
}
