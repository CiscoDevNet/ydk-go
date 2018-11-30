// Cisco XE Native Multiprotocol Label Switching (MPLS) Yang model.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls"))
}

// LdpDiscoveryAddressType
type LdpDiscoveryAddressType string

const (
    LdpDiscoveryAddressType_interface_ LdpDiscoveryAddressType = "interface"
)

// MplsTeTiebreakerType
type MplsTeTiebreakerType string

const (
    // Use max-fill tiebreaker
    MplsTeTiebreakerType_max_fill MplsTeTiebreakerType = "max-fill"

    // Use min-fill tiebreaker
    MplsTeTiebreakerType_min_fill MplsTeTiebreakerType = "min-fill"

    // Use random tiebreaker
    MplsTeTiebreakerType_random MplsTeTiebreakerType = "random"
)

