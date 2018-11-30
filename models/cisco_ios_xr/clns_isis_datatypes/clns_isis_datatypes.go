// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package clns_isis_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package clns_isis_datatypes"))
}

// IsisInternalLevel represents Isis internal level
type IsisInternalLevel string

const (
    // Level not set
    IsisInternalLevel_not_set IsisInternalLevel = "not-set"

    // Level1
    IsisInternalLevel_level1 IsisInternalLevel = "level1"

    // Level2
    IsisInternalLevel_level2 IsisInternalLevel = "level2"
)

// IsisAddressFamily represents Isis address family
type IsisAddressFamily string

const (
    // IPv4
    IsisAddressFamily_ipv4 IsisAddressFamily = "ipv4"

    // IPv6
    IsisAddressFamily_ipv6 IsisAddressFamily = "ipv6"
)

// IsisSubAddressFamily represents Isis sub address family
type IsisSubAddressFamily string

const (
    // Unicast
    IsisSubAddressFamily_unicast IsisSubAddressFamily = "unicast"

    // Multicast
    IsisSubAddressFamily_multicast IsisSubAddressFamily = "multicast"
)

