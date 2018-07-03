// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_bgp_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_bgp_datatypes"))
}

// BgpAfAdditionalPathsCfg represents Bgp af additional paths cfg
type BgpAfAdditionalPathsCfg string

const (
    // Enable
    BgpAfAdditionalPathsCfg_enable BgpAfAdditionalPathsCfg = "enable"

    // Disable
    BgpAfAdditionalPathsCfg_disable BgpAfAdditionalPathsCfg = "disable"
)

// BgpSubsequentAddressFamily represents Bgp subsequent address family
type BgpSubsequentAddressFamily string

const (
    // Unicast
    BgpSubsequentAddressFamily_unicast BgpSubsequentAddressFamily = "unicast"

    // Multicast
    BgpSubsequentAddressFamily_multicast BgpSubsequentAddressFamily = "multicast"

    // Labeled unicast
    BgpSubsequentAddressFamily_labeled_unicast BgpSubsequentAddressFamily = "labeled-unicast"

    // MVPN
    BgpSubsequentAddressFamily_mvpn BgpSubsequentAddressFamily = "mvpn"

    // MSPW
    BgpSubsequentAddressFamily_mspw BgpSubsequentAddressFamily = "mspw"

    // Tunnel
    BgpSubsequentAddressFamily_tunnel BgpSubsequentAddressFamily = "tunnel"

    // VPLS
    BgpSubsequentAddressFamily_vpls BgpSubsequentAddressFamily = "vpls"

    // MDT
    BgpSubsequentAddressFamily_mdt BgpSubsequentAddressFamily = "mdt"

    // VPWS
    BgpSubsequentAddressFamily_vpws BgpSubsequentAddressFamily = "vpws"

    // EVPN
    BgpSubsequentAddressFamily_evpn BgpSubsequentAddressFamily = "evpn"

    // LS
    BgpSubsequentAddressFamily_ls BgpSubsequentAddressFamily = "ls"

    // SRPolicy
    BgpSubsequentAddressFamily_sr_policy BgpSubsequentAddressFamily = "sr-policy"

    // VPN
    BgpSubsequentAddressFamily_vpn BgpSubsequentAddressFamily = "vpn"

    // VPN MCAST
    BgpSubsequentAddressFamily_vpn_mcast BgpSubsequentAddressFamily = "vpn-mcast"

    // Rt filter
    BgpSubsequentAddressFamily_rt_filter BgpSubsequentAddressFamily = "rt-filter"

    // Flowspec
    BgpSubsequentAddressFamily_flowspec BgpSubsequentAddressFamily = "flowspec"

    // VPN Flowspec
    BgpSubsequentAddressFamily_vpn_flowspec BgpSubsequentAddressFamily = "vpn-flowspec"

    // All
    BgpSubsequentAddressFamily_all BgpSubsequentAddressFamily = "all"
)

// BgpNbrCapAdditionalPathsCfg represents Bgp nbr cap additional paths cfg
type BgpNbrCapAdditionalPathsCfg string

const (
    // Enable
    BgpNbrCapAdditionalPathsCfg_enable BgpNbrCapAdditionalPathsCfg = "enable"

    // Disable
    BgpNbrCapAdditionalPathsCfg_disable BgpNbrCapAdditionalPathsCfg = "disable"
)

// BgpOfficialAddressFamily represents Bgp official address family
type BgpOfficialAddressFamily string

const (
    // IPv4
    BgpOfficialAddressFamily_ipv4 BgpOfficialAddressFamily = "ipv4"

    // IPv6
    BgpOfficialAddressFamily_ipv6 BgpOfficialAddressFamily = "ipv6"

    // L2VPN
    BgpOfficialAddressFamily_l2vpn BgpOfficialAddressFamily = "l2vpn"

    // LS
    BgpOfficialAddressFamily_ls BgpOfficialAddressFamily = "ls"

    // All
    BgpOfficialAddressFamily_all BgpOfficialAddressFamily = "all"
)

// BgpTos represents Bgp tos
type BgpTos string

const (
    // Precedence
    BgpTos_precedence BgpTos = "precedence"

    // DSCP
    BgpTos_dscp BgpTos = "dscp"
)

// BgpAddressFamily represents Bgp address family
type BgpAddressFamily string

const (
    // IPv4 unicast
    BgpAddressFamily_ipv4_unicast BgpAddressFamily = "ipv4-unicast"

    // IPv4 multicast
    BgpAddressFamily_ipv4_multicast BgpAddressFamily = "ipv4-multicast"

    // IPv4 labeled-unicast
    BgpAddressFamily_ipv4_labeled_unicast BgpAddressFamily = "ipv4-labeled-unicast"

    // IPv4 tunnel
    BgpAddressFamily_ipv4_tunnel BgpAddressFamily = "ipv4-tunnel"

    // VPNv4 unicast
    BgpAddressFamily_vpnv4_unicast BgpAddressFamily = "vpnv4-unicast"

    // IPv6 unicast
    BgpAddressFamily_ipv6_unicast BgpAddressFamily = "ipv6-unicast"

    // IPv6 multicast
    BgpAddressFamily_ipv6_multicast BgpAddressFamily = "ipv6-multicast"

    // IPv6 labeled-unicast
    BgpAddressFamily_ipv6_labeled_unicast BgpAddressFamily = "ipv6-labeled-unicast"

    // VPNv6 unicast
    BgpAddressFamily_vpnv6_unicast BgpAddressFamily = "vpnv6-unicast"

    // IPv4 MDT
    BgpAddressFamily_ipv4_mdt BgpAddressFamily = "ipv4-mdt"

    // L2VPN VPLS-VPWS
    BgpAddressFamily_l2vpn_vpls BgpAddressFamily = "l2vpn-vpls"

    // IPv4 rt-filter
    BgpAddressFamily_ipv4rt_constraint BgpAddressFamily = "ipv4rt-constraint"

    // IPv4 MVPN
    BgpAddressFamily_ipv4_mvpn BgpAddressFamily = "ipv4-mvpn"

    // IPv6 MVPN
    BgpAddressFamily_ipv6_mvpn BgpAddressFamily = "ipv6-mvpn"

    // L2VPN EVPN
    BgpAddressFamily_l2vpn_evpn BgpAddressFamily = "l2vpn-evpn"

    // Link-state link-state
    BgpAddressFamily_lsls BgpAddressFamily = "lsls"

    // VPNv4 Multicast
    BgpAddressFamily_vpnv4_multicast BgpAddressFamily = "vpnv4-multicast"

    // VPNv6 Multicast
    BgpAddressFamily_vpnv6_multicast BgpAddressFamily = "vpnv6-multicast"

    // IPv4 flowspec
    BgpAddressFamily_ipv4_flowspec BgpAddressFamily = "ipv4-flowspec"

    // IPv6 flowspec
    BgpAddressFamily_ipv6_flowspec BgpAddressFamily = "ipv6-flowspec"

    // VPNv4 flowspec
    BgpAddressFamily_vpnv4_flowspec BgpAddressFamily = "vpnv4-flowspec"

    // VPNv6 flowspec
    BgpAddressFamily_vpnv6_flowspec BgpAddressFamily = "vpnv6-flowspec"

    // L2VPN MSPW
    BgpAddressFamily_l2vpn_mspw BgpAddressFamily = "l2vpn-mspw"

    // IPv4 SRPolicy
    BgpAddressFamily_ipv4_sr_policy BgpAddressFamily = "ipv4-sr-policy"

    // IPv6 SRPolicy
    BgpAddressFamily_ipv6_sr_policy BgpAddressFamily = "ipv6-sr-policy"

    // All Address Families
    BgpAddressFamily_all_address_family BgpAddressFamily = "all-address-family"
)

// BgpPrecedenceDscp represents Bgp precedence dscp
type BgpPrecedenceDscp string

const (
    // AF11 dscp (001010)
    BgpPrecedenceDscp_af11 BgpPrecedenceDscp = "af11"

    // AF12 dscp (001100)
    BgpPrecedenceDscp_af12 BgpPrecedenceDscp = "af12"

    // AF13 dscp (001110)
    BgpPrecedenceDscp_af13 BgpPrecedenceDscp = "af13"

    // AF21 dscp (010010)
    BgpPrecedenceDscp_af21 BgpPrecedenceDscp = "af21"

    // AF22 dscp (010100)
    BgpPrecedenceDscp_af22 BgpPrecedenceDscp = "af22"

    // AF23 dscp (010110)
    BgpPrecedenceDscp_af23 BgpPrecedenceDscp = "af23"

    // AF31 dscp (011010)
    BgpPrecedenceDscp_af31 BgpPrecedenceDscp = "af31"

    // AF32 dscp (011100)
    BgpPrecedenceDscp_af32 BgpPrecedenceDscp = "af32"

    // AF33 dscp (011110)
    BgpPrecedenceDscp_af33 BgpPrecedenceDscp = "af33"

    // AF41 dscp (100010)
    BgpPrecedenceDscp_af41 BgpPrecedenceDscp = "af41"

    // AF42 dscp (100100)
    BgpPrecedenceDscp_af42 BgpPrecedenceDscp = "af42"

    // AF43 dscp (100110)
    BgpPrecedenceDscp_af43 BgpPrecedenceDscp = "af43"

    // CS1 dscp (001000)
    BgpPrecedenceDscp_cs1 BgpPrecedenceDscp = "cs1"

    // CS2 dscp (010000)
    BgpPrecedenceDscp_cs2 BgpPrecedenceDscp = "cs2"

    // CS3 dscp (011000)
    BgpPrecedenceDscp_cs3 BgpPrecedenceDscp = "cs3"

    // CS4 dscp (100000)
    BgpPrecedenceDscp_cs4 BgpPrecedenceDscp = "cs4"

    // CS5 dscp (101000)
    BgpPrecedenceDscp_cs5 BgpPrecedenceDscp = "cs5"

    // CS6 dscp (110000)
    BgpPrecedenceDscp_cs6 BgpPrecedenceDscp = "cs6"

    // CS7 dscp (111000)
    BgpPrecedenceDscp_cs7 BgpPrecedenceDscp = "cs7"

    // EF dscp (101110)
    BgpPrecedenceDscp_ef BgpPrecedenceDscp = "ef"

    // critical precedence (5)
    BgpPrecedenceDscp_critical BgpPrecedenceDscp = "critical"

    // flash precedence (3)
    BgpPrecedenceDscp_flash BgpPrecedenceDscp = "flash"

    // flash override precedence (4)
    BgpPrecedenceDscp_flash_override BgpPrecedenceDscp = "flash-override"

    // immediate precedence (2)
    BgpPrecedenceDscp_immediate BgpPrecedenceDscp = "immediate"

    // internetwork control precedence (6)
    BgpPrecedenceDscp_internet BgpPrecedenceDscp = "internet"

    // network control precedence (7)
    BgpPrecedenceDscp_network BgpPrecedenceDscp = "network"

    // priority precedence (1)
    BgpPrecedenceDscp_priority BgpPrecedenceDscp = "priority"

    // default dscp or routine precedence (0)
    BgpPrecedenceDscp_default_or_routine BgpPrecedenceDscp = "default-or-routine"
)

// BgpAdvertiseLocalLabeledRouteCfg represents Bgp advertise local labeled route cfg
type BgpAdvertiseLocalLabeledRouteCfg string

const (
    // Enable
    BgpAdvertiseLocalLabeledRouteCfg_enable BgpAdvertiseLocalLabeledRouteCfg = "enable"

    // Disable
    BgpAdvertiseLocalLabeledRouteCfg_disable BgpAdvertiseLocalLabeledRouteCfg = "disable"
)

// BgpUpdateFilterAction represents Bgp update filter action
type BgpUpdateFilterAction string

const (
    // Treat as withdraw
    BgpUpdateFilterAction_treat_as_withdraw BgpUpdateFilterAction = "treat-as-withdraw"

    // Discard attribute
    BgpUpdateFilterAction_discard_attibute BgpUpdateFilterAction = "discard-attibute"
)

