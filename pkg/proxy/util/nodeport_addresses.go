/*
Copyright 2022 The Kubernetes Authors.

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
	"fmt"
	"net"

	v1 "k8s.io/api/core/v1"
	utilip "k8s.io/apimachinery/pkg/util/ip"
	netutils "k8s.io/utils/net"
)

// NodePortAddresses is used to handle the --nodeport-addresses flag
type NodePortAddresses struct {
	cidrStrings []string

	cidrs                []*net.IPNet
	containsIPv4Loopback bool
	matchAll             bool
}

// RFC 5735 127.0.0.0/8 - This block is assigned for use as the Internet host loopback address
var ipv4LoopbackStart = net.IPv4(127, 0, 0, 0)

// NewNodePortAddresses takes an IP family and the `--nodeport-addresses` value (which is
// assumed to contain only valid CIDRs, potentially of both IP families) and the primary IP
// (which will be used as node port address when `--nodeport-addresses` is empty).
// It will return a NodePortAddresses object for the given family. If there are no CIDRs of
// the given family then the CIDR "0.0.0.0/0" or "::/0" will be added (even if there are
// CIDRs of the other family).
func NewNodePortAddresses(family v1.IPFamily, cidrStrings []string, primaryIP net.IP) *NodePortAddresses {
	npa := &NodePortAddresses{}

	// Filter CIDRs to correct family
	for _, str := range cidrStrings {
		if utilip.IPFamilyOfCIDR(str) == family {
			npa.cidrStrings = append(npa.cidrStrings, str)
		}
	}
	if len(npa.cidrStrings) == 0 {
		if primaryIP == nil {
			if family == v1.IPv4Protocol {
				npa.cidrStrings = []string{IPv4ZeroCIDR}
			} else {
				npa.cidrStrings = []string{IPv6ZeroCIDR}
			}
		} else {
			if family == v1.IPv4Protocol {
				npa.cidrStrings = []string{fmt.Sprintf("%s/32", primaryIP.String())}
			} else {
				npa.cidrStrings = []string{fmt.Sprintf("%s/128", primaryIP.String())}
			}
		}
	}

	// Now parse
	for _, str := range npa.cidrStrings {
		_, cidr, _ := netutils.ParseCIDRSloppy(str)

		if utilip.IsIPv4CIDR(cidr) {
			if cidr.IP.IsLoopback() || cidr.Contains(ipv4LoopbackStart) {
				npa.containsIPv4Loopback = true
			}
		}

		if IsZeroCIDR(str) {
			// Ignore everything else
			npa.cidrs = []*net.IPNet{cidr}
			npa.matchAll = true
			break
		}

		npa.cidrs = append(npa.cidrs, cidr)
	}

	return npa
}

func (npa *NodePortAddresses) String() string {
	return fmt.Sprintf("%v", npa.cidrStrings)
}

// MatchAll returns true if npa matches all node IPs (of npa's given family)
func (npa *NodePortAddresses) MatchAll() bool {
	return npa.matchAll
}

// GetNodeIPs return all matched node IP addresses for npa's CIDRs. If no matching
// IPs are found, it returns an empty list.
// NetworkInterfacer is injected for test purpose.
func (npa *NodePortAddresses) GetNodeIPs(nw NetworkInterfacer) ([]net.IP, error) {
	addrs, err := nw.InterfaceAddrs()
	if err != nil {
		return nil, fmt.Errorf("error listing all interfaceAddrs from host, error: %v", err)
	}

	// Use a map to dedup matches
	addresses := make(map[string]net.IP)
	for _, cidr := range npa.cidrs {
		for _, addr := range addrs {
			ip := utilip.IPFromInterfaceAddr(addr)
			if ip != nil && cidr.Contains(ip) {
				addresses[ip.String()] = ip
			}
		}
	}

	ips := make([]net.IP, 0, len(addresses))
	for _, ip := range addresses {
		ips = append(ips, ip)
	}

	return ips, nil
}

// ContainsIPv4Loopback returns true if npa's CIDRs contain an IPv4 loopback address.
func (npa *NodePortAddresses) ContainsIPv4Loopback() bool {
	return npa.containsIPv4Loopback
}
