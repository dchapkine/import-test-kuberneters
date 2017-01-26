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

package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	fedclient "k8s.io/kubernetes/federation/client/clientset_generated/federation_clientset"
	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	kubectlcmd "k8s.io/kubernetes/pkg/kubectl/cmd"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	// KubeconfigSecretDataKey is the key name used in the secret to
	// stores a cluster's credentials.
	KubeconfigSecretDataKey = "kubeconfig"

	// DefaultFederationSystemNamespace is the namespace in which
	// federation system components are hosted.
	DefaultFederationSystemNamespace = "federation-system"

	lbAddrRetryInterval = 5 * time.Second
	podWaitInterval     = 2 * time.Second
)

// AdminConfig provides a filesystem based kubeconfig (via
// `PathOptions()`) and a mechanism to talk to the federation
// host cluster and the federation control plane api server.
type AdminConfig interface {
	// PathOptions provides filesystem based kubeconfig access.
	PathOptions() *clientcmd.PathOptions
	// FedClientSet provides a federation API compliant clientset
	// to communicate with the federation control plane api server
	FederationClientset(context, kubeconfigPath string) (*fedclient.Clientset, error)
	// HostFactory provides a mechanism to communicate with the
	// cluster where federation control plane is hosted.
	HostFactory(hostcontext, kubeconfigPath string) cmdutil.Factory
}

// adminConfig implements the AdminConfig interface.
type adminConfig struct {
	pathOptions *clientcmd.PathOptions
}

// NewAdminConfig creates an admin config for `kubefed` commands.
func NewAdminConfig(pathOptions *clientcmd.PathOptions) AdminConfig {
	return &adminConfig{
		pathOptions: pathOptions,
	}
}

func (a *adminConfig) PathOptions() *clientcmd.PathOptions {
	return a.pathOptions
}

func (a *adminConfig) FederationClientset(context, kubeconfigPath string) (*fedclient.Clientset, error) {
	fedConfig := a.getClientConfig(context, kubeconfigPath)
	fedClientConfig, err := fedConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	return fedclient.NewForConfigOrDie(fedClientConfig), nil
}

func (a *adminConfig) HostFactory(hostcontext, kubeconfigPath string) cmdutil.Factory {
	hostClientConfig := a.getClientConfig(hostcontext, kubeconfigPath)
	return cmdutil.NewFactory(hostClientConfig)
}

func (a *adminConfig) getClientConfig(context, kubeconfigPath string) clientcmd.ClientConfig {
	loadingRules := *a.pathOptions.LoadingRules
	loadingRules.Precedence = a.pathOptions.GetLoadingPrecedence()
	loadingRules.ExplicitPath = kubeconfigPath
	overrides := &clientcmd.ConfigOverrides{
		CurrentContext: context,
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(&loadingRules, overrides)
}

// SubcommandFlags holds the flags required by the subcommands of
// `kubefed`.
type SubcommandFlags struct {
	Name                      string
	Host                      string
	FederationSystemNamespace string
	Kubeconfig                string
}

// AddSubcommandFlags adds the definition for `kubefed` subcommand
// flags.
func AddSubcommandFlags(cmd *cobra.Command) {
	cmd.Flags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	cmd.Flags().String("host-cluster-context", "", "Host cluster context")
	cmd.Flags().String("federation-system-namespace", DefaultFederationSystemNamespace, "Namespace in the host cluster where the federation system components are installed")
}

// GetSubcommandFlags retrieves the command line flag values for the
// `kubefed` subcommands.
func GetSubcommandFlags(cmd *cobra.Command, args []string) (*SubcommandFlags, error) {
	name, err := kubectlcmd.NameFromCommandArgs(cmd, args)
	if err != nil {
		return nil, err
	}
	return &SubcommandFlags{
		Name: name,
		Host: cmdutil.GetFlagString(cmd, "host-cluster-context"),
		FederationSystemNamespace: cmdutil.GetFlagString(cmd, "federation-system-namespace"),
		Kubeconfig:                cmdutil.GetFlagString(cmd, "kubeconfig"),
	}, nil
}

func CreateKubeconfigSecret(clientset *client.Clientset, kubeconfig *clientcmdapi.Config, namespace, name string, dryRun bool) (*api.Secret, error) {
	configBytes, err := clientcmd.Write(*kubeconfig)
	if err != nil {
		return nil, err
	}

	// Build the secret object with the minified and flattened
	// kubeconfig content.
	secret := &api.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: map[string][]byte{
			KubeconfigSecretDataKey: configBytes,
		},
	}

	if !dryRun {
		return clientset.Core().Secrets(namespace).Create(secret)
	}
	return secret, nil
}

func WaitForLoadBalancerAddress(clientset *client.Clientset, svc *api.Service) ([]string, []string, error) {
	ips := []string{}
	hostnames := []string{}

	err := wait.PollImmediateInfinite(lbAddrRetryInterval, func() (bool, error) {
		pollSvc, err := clientset.Core().Services(svc.Namespace).Get(svc.Name, metav1.GetOptions{})
		if err != nil {
			return false, nil
		}
		if ings := pollSvc.Status.LoadBalancer.Ingress; len(ings) > 0 {
			for _, ing := range ings {
				if len(ing.IP) > 0 {
					ips = append(ips, ing.IP)
				}
				if len(ing.Hostname) > 0 {
					hostnames = append(hostnames, ing.Hostname)
				}
			}
			if len(ips) > 0 || len(hostnames) > 0 {
				return true, nil
			}
		}
		return false, nil
	})
	if err != nil {
		return nil, nil, err
	}

	return ips, hostnames, nil
}

func WaitForPods(clientset *client.Clientset, pods []string, namespace string) error {
	err := wait.PollInfinite(podWaitInterval, func() (bool, error) {
		podCheck := len(pods)
		podList, err := clientset.Core().Pods(namespace).List(metav1.ListOptions{})
		if err != nil {
			return false, nil
		}
		for _, pod := range podList.Items {
			for _, fedPod := range pods {
				if strings.HasPrefix(pod.Name, fedPod) && pod.Status.Phase == "Running" {
					podCheck -= 1
				}
			}
			//ensure that all pods are in running state or keep waiting
			if podCheck == 0 {
				return true, nil
			}
		}
		return false, nil
	})
	return err
}

func GetFirstNodeIP(clientset *client.Clientset) ([]string, error) {
	address := ""
	nodeList, err := clientset.Core().Nodes().List(metav1.ListOptions{})
	if err == nil {
		if len(nodeList.Items) > 0 {
			// Try to get IP for node from list of node addresses
			// prefer NodeExternalIP over NodeInternalIP over other types (may be Legacy IP's)
			for _, addr := range nodeList.Items[0].Status.Addresses {
				if addr.Type == api.NodeExternalIP {
					address = addr.Address
					break
				} else if addr.Type == api.NodeInternalIP {
					address = addr.Address
					continue
				}
				address = addr.Address
			}
		}
	}

	return []string{address}, nil
}
