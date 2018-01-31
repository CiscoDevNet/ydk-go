// This module contains a collection of YANG definitions
// for Cisco IOS-XR ethernet-cfm package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-l2-eth-infra-cfg,
//   Cisco-IOS-XR-snmp-agent-cfg,
//   Cisco-IOS-XR-infra-sla-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_cfm_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_cfm_cfg"))
}

// CfmLmCountersCfg represents Cfm lm counters cfg
type CfmLmCountersCfg string

const (
    // Aggregate Counters
    CfmLmCountersCfg_aggregate CfmLmCountersCfg = "aggregate"

    // List of per-CoS counters
    CfmLmCountersCfg_list CfmLmCountersCfg = "list"

    // Range of per-CoS counters
    CfmLmCountersCfg_range_ CfmLmCountersCfg = "range"
)

// CfmMdidFormat represents Cfm mdid format
type CfmMdidFormat string

const (
    // Null MDID
    CfmMdidFormat_null CfmMdidFormat = "null"

    // DNS-like MDID
    CfmMdidFormat_dns_like CfmMdidFormat = "dns-like"

    // MDID Comprising MAC Address and 16-bit integer
    CfmMdidFormat_mac_address CfmMdidFormat = "mac-address"

    // String MDID
    CfmMdidFormat_string CfmMdidFormat = "string"
)

// CfmShortMaNameFormat represents Cfm short ma name format
type CfmShortMaNameFormat string

const (
    // VLAN ID
    CfmShortMaNameFormat_vlan_id CfmShortMaNameFormat = "vlan-id"

    // String Short MA Name
    CfmShortMaNameFormat_string CfmShortMaNameFormat = "string"

    // Numeric Short MA Name
    CfmShortMaNameFormat_number CfmShortMaNameFormat = "number"

    // RFC 2685 VPN ID
    CfmShortMaNameFormat_vpn_id CfmShortMaNameFormat = "vpn-id"

    // ICC-based format
    CfmShortMaNameFormat_icc_based CfmShortMaNameFormat = "icc-based"
)

// CfmService represents Cfm service
type CfmService string

const (
    // Use a Bridge Domain - all MEPs will be Up MEPs
    // and MIPs are permitted
    CfmService_bridge_domain CfmService = "bridge-domain"

    // Use a P2P Cross Connect - all MEPs will be Up
    // MEPs and MIPs are permitted
    CfmService_p2p_cross_connect CfmService = "p2p-cross-connect"

    // Use a MP2MP Cross Connect - all MEPs will be Up
    // MEPs and MIPs are permitted
    CfmService_mp2mp_cross_connect CfmService = "mp2mp-cross-connect"

    // Use a VLAN-aware Flexible Cross Connect - all
    // MEPs will be Up MEPs and MIPs are permitted
    CfmService_vlan_aware_flexible_cross_connect CfmService = "vlan-aware-flexible-cross-connect"

    // Use a VLAN-unaware Flexible Cross Connect - all
    // MEPs will be Up MEPs and MIPs are permitted
    CfmService_vlan_unaware_flexible_cross_connect CfmService = "vlan-unaware-flexible-cross-connect"

    // Down MEPs - no MIPs permitted
    CfmService_down_meps CfmService = "down-meps"
)

// CfmMipPolicy represents Cfm mip policy
type CfmMipPolicy string

const (
    // Create MIPs on all ports in the Bridge Domain
    // or Cross-connect
    CfmMipPolicy_all CfmMipPolicy = "all"

    // Create MIPs on ports which have a MEP at a
    // lower level
    CfmMipPolicy_lower_mep_only CfmMipPolicy = "lower-mep-only"
)

