// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-io package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_io_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_io_cfg"))
}

// Ipv4DefaultPing represents Ipv4 default ping
type Ipv4DefaultPing string

const (
    // Default route is not allowed to match when
    // checking source address
    Ipv4DefaultPing_disabled Ipv4DefaultPing = "disabled"

    // Allow default route to match when checking
    // source address
    Ipv4DefaultPing_enabled Ipv4DefaultPing = "enabled"
)

// Ipv4SelfPing represents Ipv4 self ping
type Ipv4SelfPing string

const (
    // Doesn't allow router to ping itself
    Ipv4SelfPing_disabled Ipv4SelfPing = "disabled"

    // Allow router to ping itself
    Ipv4SelfPing_enabled Ipv4SelfPing = "enabled"
)

// Ipv4Reachable represents Ipv4 reachable
type Ipv4Reachable string

const (
    // Source is reachable via any interface
    Ipv4Reachable_any Ipv4Reachable = "any"

    // Source is reachable via interface on which
    // packet was received
    Ipv4Reachable_received Ipv4Reachable = "received"
)

// Ipv4InterfaceQppb represents Ipv4 interface qppb
type Ipv4InterfaceQppb string

const (
    // Enable IP precedence based QPPB
    Ipv4InterfaceQppb_ip_precedence Ipv4InterfaceQppb = "ip-precedence"

    // Enable QoS-group based QPPB
    Ipv4InterfaceQppb_qos_group Ipv4InterfaceQppb = "qos-group"

    // Enable both IP precedence and QoS-group based
    // QPPB
    Ipv4InterfaceQppb_both Ipv4InterfaceQppb = "both"
)

