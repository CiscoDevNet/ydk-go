// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package lmp_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lmp_datatypes"))
}

// OlmSwitchingCap represents Olm switching cap
type OlmSwitchingCap string

const (
    // Lambda switch capable
    OlmSwitchingCap_lsc OlmSwitchingCap = "lsc"

    // Fiber switch capable
    OlmSwitchingCap_fsc OlmSwitchingCap = "fsc"
)

// OlmAddr represents Olm addr
type OlmAddr string

const (
    // IPv4 address
    OlmAddr_ipv4 OlmAddr = "ipv4"

    // IPv6 address
    OlmAddr_ipv6 OlmAddr = "ipv6"

    // Unnumbered address
    OlmAddr_unnumbered OlmAddr = "unnumbered"

    // NSAP address
    OlmAddr_nsap OlmAddr = "nsap"
)

