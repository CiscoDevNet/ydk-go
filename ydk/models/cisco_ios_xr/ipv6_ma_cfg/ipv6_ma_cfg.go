// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-ma package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_ma_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_ma_cfg"))
}

// Ipv6SelfPing represents Ipv6 self ping
type Ipv6SelfPing string

const (
    // Doesn't allow router to ping itself
    Ipv6SelfPing_disabled Ipv6SelfPing = "disabled"

    // Allow router to ping itself
    Ipv6SelfPing_enabled Ipv6SelfPing = "enabled"
)

// Ipv6Reachable represents Ipv6 reachable
type Ipv6Reachable string

const (
    // Source is reachable via any interface
    Ipv6Reachable_any Ipv6Reachable = "any"

    // Source is reachable via interface on which
    // packet was received
    Ipv6Reachable_received Ipv6Reachable = "received"
)

// Ipv6DefaultPing represents Ipv6 default ping
type Ipv6DefaultPing string

const (
    // Default route is not allowed to match when
    // checking source address
    Ipv6DefaultPing_disabled Ipv6DefaultPing = "disabled"

    // Allow default route to match when checking
    // source address
    Ipv6DefaultPing_enabled Ipv6DefaultPing = "enabled"
)

// Ipv6Qppb represents Ipv6 qppb
type Ipv6Qppb string

const (
    // No QPPB configuration
    Ipv6Qppb_none Ipv6Qppb = "none"

    // Enable ip-precedence based QPPB
    Ipv6Qppb_ip_precedence Ipv6Qppb = "ip-precedence"

    // Enable qos-group based QPPB
    Ipv6Qppb_qos_group Ipv6Qppb = "qos-group"

    // Enable both ip-precedence and qos-group based
    // QPPB
    Ipv6Qppb_both Ipv6Qppb = "both"
)

