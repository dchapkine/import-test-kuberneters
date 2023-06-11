/*
Copyright 2023 The Kubernetes Authors.

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

package ip

// LegacyIPStringContext is a context in which a legacy IP exists that must be parsed with
// ParseLegacyIP or ParseLegacyCIDR.
//
// Note that while it is possible to cast an arbitrary string to LegacyIPStringContext, it
// is preferable to register new values here, so we are explicitly keeping track of all
// APIs that contain legacy IP strings.
type LegacyIPStringContext string

const (
	// Legacy API fields
	EndpointAddressIPContext            LegacyIPStringContext = "v1.EndpointAddress.IP"
	HostAliasIPContext                  LegacyIPStringContext = "v1.HostAlias.IP"
	NodeSpecPodCIDRsContext             LegacyIPStringContext = "v1.NodeSpec.PodCIDRs"
	PodDNSConfigNameserversContext      LegacyIPStringContext = "v1.PodDNSConfig.Nameservers"
	PodStatusPodIPsContext              LegacyIPStringContext = "v1.PodStatus.PodIPs"
	ServiceSpecClusterIPsContext        LegacyIPStringContext = "v1.ServiceSpec.ClusterIPs"
	ServiceSpecExternalIPsContext       LegacyIPStringContext = "v1.ServiceSpec.ExternalIPs"
	ServiceLoadBalancerIngressIPContext LegacyIPStringContext = "v1.LoadBalancerIngress.IP"

	EndpointSliceAddressesContext LegacyIPStringContext = "discoveryv1.Endpoint.Addresses"

	ClusterCIDRSpecContext            LegacyIPStringContext = "networkingv1.ClusterCIDRSpec"
	IngressLoadBalancerIngressContext LegacyIPStringContext = "networkingv1.IngressLoadBalancer.Ingress"
	NetworkPolicyIPBlockContext       LegacyIPStringContext = "networkingv1.IPBlock"

	// Legacy command-line arguments
	KubeProxyBindAddress LegacyIPStringContext = "kube-proxy --bind-address"
	// ...

	// Legacy annotations
	NodeCloudNodeIPAnnotation LegacyIPStringContext = "alpha.kubernetes.io/provided-node-ip"
	// ...
)
