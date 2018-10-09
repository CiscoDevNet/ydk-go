// This module contains type definitions common to all Cisco IOS XE native models
// 
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package common_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package common_types"))
}

// AddrType represents IP address type
type AddrType string

const (
    AddrType_address_none AddrType = "address-none"

    AddrType_ipv4_address AddrType = "ipv4-address"

    AddrType_ipv6_address AddrType = "ipv6-address"

    AddrType_ipv4_address_mcast AddrType = "ipv4-address-mcast"

    AddrType_ipv6_address_mcast AddrType = "ipv6-address-mcast"
)

