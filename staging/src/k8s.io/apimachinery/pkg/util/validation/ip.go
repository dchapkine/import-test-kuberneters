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

package validation

import (
	"fmt"

	utilip "k8s.io/apimachinery/pkg/util/ip"
	"k8s.io/apimachinery/pkg/util/validation/field"
	netutils "k8s.io/utils/net"
)

// "Good" validation functions, for new API objects and fields

// ValidateIP tests that value is a valid IP address (either IPv4 or IPv6). Note that
// this rejects some values that were considered valid in older Kubernetes APIs; those
// fields must use ValidateIPForLegacyAPI instead.
func ValidateIP(fldPath *field.Path, value string) field.ErrorList {
	var allErrors field.ErrorList
	if _, err := utilip.ParseIP(value); err != nil {
		allErrors = append(allErrors, field.Invalid(fldPath, value, fmt.Sprintf("must be a valid IP address: %v", err)))
	}
	return allErrors
}

// ValidateIPv4Address tests that the argument is a valid IPv4 address. Note that
// this rejects some values that were considered valid in older Kubernetes APIs; those
// fields must use ValidateIP4AddressForLegacyAPI instead.
func ValidateIPv4Address(fldPath *field.Path, value string) field.ErrorList {
	var allErrors field.ErrorList
	ip, err := utilip.ParseIP(value)
	if err != nil {
		allErrors = append(allErrors, field.Invalid(fldPath, value, fmt.Sprintf("must be a valid IPv4 address: %v", err)))
	} else if !(ip.Is4() || ip.Is4In6()) /* FIXME !utilip.IsIPv4(ip) */ {
		allErrors = append(allErrors, field.Invalid(fldPath, value, "must be a valid IPv4 address"))
	}
	return allErrors
}

// ValidateIPv6Address tests that the argument is a valid IPv6 address. Note that
// this rejects some values that were considered valid in older Kubernetes APIs; those
// fields must use ValidateIP6AddressForLegacyAPI instead.
func ValidateIPv6Address(fldPath *field.Path, value string) field.ErrorList {
	var allErrors field.ErrorList
	ip, err := utilip.ParseIP(value)
	if err != nil {
		allErrors = append(allErrors, field.Invalid(fldPath, value, fmt.Sprintf("must be a valid IPv6 address: %v", err)))
	} else if !(ip.Is6() && !ip.Is4In6()) /* FIXME !utilip.IsIPv6(ip) */ {
		allErrors = append(allErrors, field.Invalid(fldPath, value, "must be a valid IPv6 address"))
	}
	return allErrors
}

// ValidateNonIP tests that the argument is NOT an IP address (valid or legacy). This is
// intended for contexts where a hostname is expected and an IP address is not allowed,
// but it does not validate that value is a hostname since different fields have different
// requirements there (e.g., allowing or not allowing wildcards).
func ValidateNonIP(fldPath *field.Path, value string) field.ErrorList {
	var allErrors field.ErrorList
	if netutils.ParseIPSloppy(value) != nil {
		allErrors = append(allErrors, field.Invalid(fldPath, value, fmt.Sprintf("must be a DNS name, not an IP address")))
	}
	return allErrors
}

// ValidateCIDR tests that value is a valid CIDR address (either IPv4 or IPv6). Note that
// this rejects some values that were considered valid in older Kubernetes APIs; those
// fields must use ValidateCIDRForLegacyAPI instead.
func ValidateCIDR(fldPath *field.Path, value string) field.ErrorList {
	var allErrors field.ErrorList
	if _, err := utilip.ParseCIDR(value); err != nil {
		allErrors = append(allErrors, field.Invalid(fldPath, value, fmt.Sprintf("must be a valid CIDR address: %v", err)))
	}
	return allErrors
}

// Legacy validation functions.

// ValidateIPForLegacyAPI tests that value was considered an IP according to legacy
// Kubernetes IP validation rules. This must be used for validating API fields that
// historically used these rules, but MUST NOT be used for new APIs.
func ValidateIPForLegacyAPI(fldPath *field.Path, value string, context utilip.LegacyIPStringContext) field.ErrorList {
	var allErrors field.ErrorList
	if _, _, err := utilip.ParseLegacyIP(value, context); err != nil {
		allErrors = append(allErrors, field.Invalid(fldPath, value, fmt.Sprintf("must be a valid IP address: %v", err)))
	}
	return allErrors
}

// ValidateIPv4AddressForLegacyAPI tests that the argument was considered a valid IPv4
// address according to legacy Kubernetes IP validation rules. This must be used for
// validating API fields that historically used these rules, but MUST NOT be used for new
// APIs.
//
// In addition to accepting IPv4 addresses with leading "0"s, this also considers
// IPv6-wrapped IPv4 addresses (e.g., "::ffff:1.2.3.4") to be valid IPv4 addresses.
func ValidateIPv4AddressForLegacyAPI(fldPath *field.Path, value string, context utilip.LegacyIPStringContext) field.ErrorList {
	var allErrors field.ErrorList
	_, ip, err := utilip.ParseLegacyIP(value, context)
	if err != nil {
		allErrors = append(allErrors, field.Invalid(fldPath, value, fmt.Sprintf("must be a valid IPv4 address: %v", err)))
	} else if !(ip.Is4() || ip.Is4In6()) /* FIXME !utilip.IsIPv4(ip) */ {
		allErrors = append(allErrors, field.Invalid(fldPath, value, "must be a valid IPv4 address"))
	}
	return allErrors
}

// ValidateIPv6AddressForLegacyAPI tests that the argument was considered a valid IPv6
// address according to legacy Kubernetes IP validation rules. This must be used for
// validating API fields that historically used these rules, but MUST NOT be used for new
// APIs.
func ValidateIPv6AddressForLegacyAPI(fldPath *field.Path, value string, context utilip.LegacyIPStringContext) field.ErrorList {
	var allErrors field.ErrorList
	_, ip, err := utilip.ParseLegacyIP(value, context)
	if err != nil {
		allErrors = append(allErrors, field.Invalid(fldPath, value, fmt.Sprintf("must be a valid IPv6 address: %v", err)))
	} else if !(ip.Is6() && !ip.Is4In6()) /* FIXME !utilip.IsIPv6(ip) */ {
		allErrors = append(allErrors, field.Invalid(fldPath, value, "must be a valid IPv6 address"))
	}
	return allErrors
}

// ValidateCIDRForLegacyAPI tests that value was considered a CIDR address according to
// legacy Kubernetes IP validation rules. This must be used for validating API fields that
// historically used these rules, but MUST NOT be used for new APIs.
func ValidateCIDRForLegacyAPI(fldPath *field.Path, value string, context utilip.LegacyIPStringContext) field.ErrorList {
	var allErrors field.ErrorList
	if _, _, err := utilip.ParseLegacyCIDR(value, context); err != nil {
		allErrors = append(allErrors, field.Invalid(fldPath, value, fmt.Sprintf("must be a valid CIDR address: %v", err)))
	}
	return allErrors
}
