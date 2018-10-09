// This module contains a collection of YANG definitions
// for Cisco IOS-XR platform-pifib package operational data.
// 
// This YANG module augments the
//   Cisco-IOS-XR-lpts-pre-ifib-oper
// module with state data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package platform_pifib_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package platform_pifib_oper"))
}

// UsageAddressFamily represents Usage address family
type UsageAddressFamily string

const (
    // Ipv4 af
    UsageAddressFamily_ipv4 UsageAddressFamily = "ipv4"

    // Ipv6 af
    UsageAddressFamily_ipv6 UsageAddressFamily = "ipv6"
)

