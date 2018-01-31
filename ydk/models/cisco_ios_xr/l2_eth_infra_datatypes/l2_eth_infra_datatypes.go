// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package l2_eth_infra_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package l2_eth_infra_datatypes"))
}

// VlanTagOrCvp represents Vlan tag or cvp
type VlanTagOrCvp string

const (
    // This is the Native VLAN and C-VLAN
    // preservation is enabled
    VlanTagOrCvp_native_with_cvlan_preservation VlanTagOrCvp = "native-with-cvlan-preservation"
)

// Rewrite represents Rewrite
type Rewrite string

const (
    // Pop 1 tag
    Rewrite_pop1 Rewrite = "pop1"

    // Pop 2 tags
    Rewrite_pop2 Rewrite = "pop2"

    // Push 1 tag
    Rewrite_push1 Rewrite = "push1"

    // Push 2 tags
    Rewrite_push2 Rewrite = "push2"

    // Translate 1-to-1
    Rewrite_translate1to1 Rewrite = "translate1to1"

    // Translate 1-to-2
    Rewrite_translate1to2 Rewrite = "translate1to2"

    // Translate 2-to-1
    Rewrite_translate2to1 Rewrite = "translate2to1"

    // Translate 2-to-2
    Rewrite_translate2to2 Rewrite = "translate2to2"
)

// Vlan represents Vlan
type Vlan string

const (
    // An 802.1ad VLAN
    Vlan_vlan_type_dot1ad Vlan = "vlan-type-dot1ad"

    // An 802.1q VLAN
    Vlan_vlan_type_dot1q Vlan = "vlan-type-dot1q"
)

// EthertypeMatch represents Ethertype match
type EthertypeMatch string

const (
    // PPP over Ethernet
    EthertypeMatch_ppp_over_ethernet EthertypeMatch = "ppp-over-ethernet"
)

// VlanTagOrNull represents Vlan tag or null
type VlanTagOrNull string

const (
    // Match any inner VLAN tag value
    VlanTagOrNull_any VlanTagOrNull = "any"
)

// VlanTagOrAny represents Vlan tag or any
type VlanTagOrAny string

const (
    // Match any VLAN tag value
    VlanTagOrAny_any VlanTagOrAny = "any"
)

// VlanTagOrNative represents Vlan tag or native
type VlanTagOrNative string

const (
    // This is the Native VLAN
    VlanTagOrNative_native VlanTagOrNative = "native"

    // This is the Native VLAN and C-VLAN
    // preservation is enabled
    VlanTagOrNative_native_with_cvlan_preservation VlanTagOrNative = "native-with-cvlan-preservation"
)

// Match represents Match
type Match string

const (
    // All otherwise unmatched packets
    Match_match_default Match = "match-default"

    // Untagged packets
    Match_match_untagged Match = "match-untagged"

    // Match Dot1Q tags
    Match_match_dot1q Match = "match-dot1q"

    // Match Dot1ad tags
    Match_match_dot1ad Match = "match-dot1ad"

    // Match Dot1Q priority-tagged packets
    Match_match_dot1q_priority Match = "match-dot1q-priority"

    // Match Dot1ad priority-tagged packets
    Match_match_dot1ad_priority Match = "match-dot1ad-priority"
)

