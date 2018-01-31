// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-nd-subscriber package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_nd_subscriber_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_nd_subscriber_cfg"))
}

// Ipv6NdRouterPrefTemplate represents Ipv6 nd router pref template
type Ipv6NdRouterPrefTemplate string

const (
    // High preference
    Ipv6NdRouterPrefTemplate_high Ipv6NdRouterPrefTemplate = "high"

    // Medium preference
    Ipv6NdRouterPrefTemplate_medium Ipv6NdRouterPrefTemplate = "medium"

    // Low preference
    Ipv6NdRouterPrefTemplate_low Ipv6NdRouterPrefTemplate = "low"
)

