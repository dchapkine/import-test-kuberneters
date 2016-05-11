// +build integration,!no-etcd

/*
Copyright 2015 The Kubernetes Authors.

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

package kubectl

import (
	"testing"

	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	clientcmdapi "k8s.io/kubernetes/pkg/client/unversioned/clientcmd/api"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/test/integration/framework"
)

func kubectlValidation(t *testing.T, cluster *clientcmdapi.Cluster) {
	testCases := []struct {
		data string
		err  bool
	}{
		{`{"apiVersion": "v1", "kind": "thisObjectShouldNotExistInAnyGroup"}`, true},
		{`{"apiVersion": "invalidVersion", "kind": "Pod"}`, true},
		{`{"apiVersion": "v1", "kind": "Pod"}`, false},

		// The following test the experimental api.
		// TODO: Replace with something more robust. These may move.
		{`{"apiVersion": "extensions/v1beta1", "kind": "Ingress"}`, false},
		{`{"apiVersion": "extensions/v1beta1", "kind": "Job"}`, false},
		{`{"apiVersion": "vNotAVersion", "kind": "Job"}`, true},
	}

	ctx := clientcmdapi.NewContext()
	cfg := clientcmdapi.NewConfig()

	cfg.Contexts = map[string]*clientcmdapi.Context{"test": ctx}
	cfg.CurrentContext = "test"
	overrides := clientcmd.ConfigOverrides{
		ClusterInfo: *cluster,
	}
	cmdConfig := clientcmd.NewNonInteractiveClientConfig(*cfg, "test", &overrides, nil)
	factory := util.NewFactory(cmdConfig)
	schema, err := factory.Validator(true, "")
	if err != nil {
		t.Errorf("failed to get validator: %v", err)
		return
	}
	for i, test := range testCases {
		err := schema.ValidateBytes([]byte(test.data))
		if err == nil {
			if test.err {
				t.Errorf("case %d: expected error", i)
			}
		} else {
			if !test.err {
				t.Errorf("case %d: unexpected error: %v", i, err)
			}
		}
	}
}

func TestKubectlValidation(t *testing.T) {
	components := framework.NewMasterComponents(&framework.Config{})
	defer components.Stop(true, true)
	// Enable swagger api on master.
	components.KubeMaster.InstallSwaggerAPI()
	cluster := clientcmdapi.NewCluster()
	cluster.Server = components.ApiServer.URL
	cluster.InsecureSkipTLSVerify = true
	kubectlValidation(t, cluster)
}

func TestKubectlValidationWithFailover(t *testing.T) {
	components := framework.NewMasterComponents(&framework.Config{})
	defer components.Stop(true, true)
	// Enable swagger api on master.
	components.KubeMaster.InstallSwaggerAPI()
	cluster := clientcmdapi.NewCluster()
	cluster.Server = "http://localhost:7777"
	cluster.Servers = []string{components.ApiServer.URL}
	cluster.InsecureSkipTLSVerify = true
	kubectlValidation(t, cluster)
}
