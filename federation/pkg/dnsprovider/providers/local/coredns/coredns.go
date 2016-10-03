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

// Package coredns is the implementation of pkg/dnsprovider interface for CoreDNS
package coredns

import (
	"fmt"
	etcd "github.com/coreos/etcd/client"
	"github.com/golang/glog"
	"gopkg.in/gcfg.v1"
	"io"
	"k8s.io/kubernetes/federation/pkg/dnsprovider"
	"strings"
)

// "local-coredns" should be used to use this DNS provider
const (
	ProviderName = "local-coredns"
)

// Config to override defaults
type Config struct {
	Global struct {
		EtcdEndpoints string `gcfg:"etcd-endpoints"`
		Domain        string `gcfg:"domain"`
	}
}

func init() {
	dnsprovider.RegisterDnsProvider(ProviderName, func(config io.Reader, zoneName string) (dnsprovider.Interface, error) {
		return newCoreDNSProviderInterface(config, zoneName)
	})
}

// newCoreDnsProviderInterface creates a new instance of an CoreDNS DNS Interface.
func newCoreDNSProviderInterface(config io.Reader, zoneName string) (*Interface, error) {
	etcdEndpoints := "http://federation-dns-server-etcd:2379"
	etcdPathPrefix := "skydns"
	domain := zoneName

	// Possibly override defaults with config below
	if config != nil {
		var cfg Config
		if err := gcfg.ReadInto(&cfg, config); err != nil {
			glog.Errorf("Couldn't read config: %v", err)
			return nil, err
		}
		etcdEndpoints = cfg.Global.EtcdEndpoints
		domain = cfg.Global.Domain
	}
	glog.Infof("Using local CoreDNS DNS provider")

	etcdCfg := etcd.Config{
		Endpoints: strings.Split(etcdEndpoints, ","),
		Transport: etcd.DefaultTransport,
	}

	c, err := etcd.New(etcdCfg)
	if err != nil {
		return nil, fmt.Errorf("Create etcd client from the config failed")
	}
	etcdKeysAPI := etcd.NewKeysAPI(c)

	i := newInterfaceWithStub(etcdKeysAPI)
	i.domain = domain
	i.etcdPathPrefix = etcdPathPrefix

	return i, nil
}
