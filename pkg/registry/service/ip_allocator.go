/*
Copyright 2014 Google Inc. All rights reserved.

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

package service

import (
	"fmt"
	"net"
	"sync"

	"github.com/golang/glog"
)

type ipAllocator struct {
	subnet *net.IPNet
	// TODO: This could be smarter, but for now a bitmap will suffice.
	lock sync.Mutex // protects 'used'
	used []byte     // a bitmap of allocated IPs
}

// newIPAllocator creates and intializes a new ipAllocator object.
func newIPAllocator(subnet *net.IPNet) *ipAllocator {
	if subnet == nil || subnet.IP == nil || subnet.Mask == nil {
		return nil
	}

	ones, bits := subnet.Mask.Size()
	numIps := 1 << uint(bits-ones)
	ipa := &ipAllocator{
		subnet: subnet,
		used:   make([]byte, numIps/8),
	}
	ipa.used[0] = 0x01            // block the network addr
	ipa.used[(numIps/8)-1] = 0x80 // block the broadcast addr
	return ipa
}

// Allocate allocates a specific IP.  This is useful when recovering saved state.
func (ipa *ipAllocator) Allocate(ip net.IP) error {
	ipa.lock.Lock()
	defer ipa.lock.Unlock()

	if !ipa.subnet.Contains(ip) {
		return fmt.Errorf("IP %s does not fall within subnet %s", ip, ipa.subnet)
	}
	offset := ipSub(ip, ipa.subnet.IP)
	i := offset / 8
	m := byte(1 << byte(offset%8))
	if ipa.used[i]&m != 0 {
		return fmt.Errorf("IP %s is already allocated", ip)
	}
	ipa.used[i] |= m
	return nil
}

// AllocateNext allocates and returns a new IP.
func (ipa *ipAllocator) AllocateNext() (net.IP, error) {
	ipa.lock.Lock()
	defer ipa.lock.Unlock()

	for i := range ipa.used {
		if ipa.used[i] != 0xff {
			freeMask := ^ipa.used[i]
			nextBit, err := ffs(freeMask)
			if err != nil {
				// If this happens, something really weird is going on.
				glog.Errorf("ffs(%#x) had an unexpected error: %s", freeMask, err)
				return nil, err
			}
			ipa.used[i] |= 1 << nextBit
			offset := (i * 8) + int(nextBit)
			ip := ipAdd(copyIP(ipa.subnet.IP), offset)
			return ip, nil
		}
	}
	return nil, fmt.Errorf("can't find a free IP in %s", ipa.subnet)
}

// This is a really dumb implementation of find-first-set-bit.
func ffs(val byte) (uint, error) {
	if val == 0 {
		return 0, fmt.Errorf("Can't find-first-set on 0")
	}
	i := uint(0)
	for ; i < 8 && (val&(1<<i) == 0); i++ {
	}
	return i, nil
}

// Add an offset to an IP address - used for joining network addr and host addr parts.
func ipAdd(ip net.IP, offset int) net.IP {
	for i := 0; offset > 0; i++ {
		b := &ip[len(ip)-1-i]
		offset += int(*b)
		*b = byte(offset % 256)
		offset >>= 8
	}
	return ip
}

// Subtract two IPs, returning the difference as an offset - used or splitting an IP into
// network addr and host addr parts.
func ipSub(lhs, rhs net.IP) int {
	// If they are not the same length, normalize them.  Make copies because net.IP is
	// a slice underneath. Sneaky sneaky.
	if len(lhs) != len(rhs) {
		lhs = copyIP(lhs).To16()
		rhs = copyIP(rhs).To16()
	}
	offset := 0
	for i := range lhs {
		offset *= 256
		offset += int(lhs[i] - rhs[i])
	}
	return offset
}

// Make a copy of a net.IP.  It appears to be a value type, but it is actually defined as a
// slice, so value assignment is shallow.  Why does a poor dumb user like me need to know
// this sort of implementation detail?
func copyIP(in net.IP) net.IP {
	out := make(net.IP, len(in))
	copy(out, in)
	return out
}

// Release de-allocates an IP.
func (ipa *ipAllocator) Release(ip net.IP) error {
	ipa.lock.Lock()
	defer ipa.lock.Unlock()

	if !ipa.subnet.Contains(ip) {
		return fmt.Errorf("IP %s does not fall within subnet %s", ip, ipa.subnet)
	}
	offset := ipSub(ip, ipa.subnet.IP)
	i := offset / 8
	m := byte(1 << byte(offset%8))
	ipa.used[i] &^= m
	return nil
}
