// This module contains a collection of YANG definitions
// for Cisco IOS-XR evpn package operational data.
// 
// This module contains definitions
// for the following management objects:
//   evpn: EVPN Operational Table
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package evpn_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package evpn_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-evpn-oper evpn}", reflect.TypeOf(Evpn{}))
    ydk.RegisterEntity("Cisco-IOS-XR-evpn-oper:evpn", reflect.TypeOf(Evpn{}))
}

// L2vpnRgRole represents L2vpn rg role
type L2vpnRgRole string

const (
    // l2vpn rg role not defined
    L2vpnRgRole_l2vpn_rg_role_not_defined L2vpnRgRole = "l2vpn-rg-role-not-defined"

    // l2vpn rg role active
    L2vpnRgRole_l2vpn_rg_role_active L2vpnRgRole = "l2vpn-rg-role-active"

    // l2vpn rg role standby
    L2vpnRgRole_l2vpn_rg_role_standby L2vpnRgRole = "l2vpn-rg-role-standby"

    // l2vpn rg role max
    L2vpnRgRole_l2vpn_rg_role_max L2vpnRgRole = "l2vpn-rg-role-max"
)

// L2vpnEvpnScMode represents EVPN Ethernet-Segment service carving mode
type L2vpnEvpnScMode string

const (
    // Invalid service carving mode
    L2vpnEvpnScMode_invalid L2vpnEvpnScMode = "invalid"

    // Auto service carving mode
    L2vpnEvpnScMode_auto L2vpnEvpnScMode = "auto"

    // Manual service carving
    L2vpnEvpnScMode_manual L2vpnEvpnScMode = "manual"

    // Manual List service carving
    L2vpnEvpnScMode_manual_list L2vpnEvpnScMode = "manual-list"

    // HRW service carving
    L2vpnEvpnScMode_hrw L2vpnEvpnScMode = "hrw"

    // Preferential service carving
    L2vpnEvpnScMode_pref L2vpnEvpnScMode = "pref"
)

// L2vpnEvpnMfMode represents L2VPN EVPN MAC flushing mode
type L2vpnEvpnMfMode string

const (
    // Invalid MAC Flushing mode
    L2vpnEvpnMfMode_invalid L2vpnEvpnMfMode = "invalid"

    // TCN STP MAC Flushing mode
    L2vpnEvpnMfMode_tcn_stp L2vpnEvpnMfMode = "tcn-stp"

    // MVRP MAC Flushing mode
    L2vpnEvpnMfMode_mvrp L2vpnEvpnMfMode = "mvrp"
)

// L2vpnEvpnLbMode represents L2VPN EVPN load balancing mode
type L2vpnEvpnLbMode string

const (
    // Invalid load balancing
    L2vpnEvpnLbMode_invalid_load_balancing L2vpnEvpnLbMode = "invalid-load-balancing"

    // Single-homed site or network
    L2vpnEvpnLbMode_single_homed L2vpnEvpnLbMode = "single-homed"

    // Multi-homed access network active/active per
    // flow
    L2vpnEvpnLbMode_multi_homed_aa_per_flow L2vpnEvpnLbMode = "multi-homed-aa-per-flow"

    // Multi-homed access network active/active per
    // service
    L2vpnEvpnLbMode_multi_homed_aa_per_service L2vpnEvpnLbMode = "multi-homed-aa-per-service"
)

// L2vpnEvpnEsi represents EVPN ESI types
type L2vpnEvpnEsi string

const (
    // ESI type zero
    L2vpnEvpnEsi_esi_type0 L2vpnEvpnEsi = "esi-type0"

    // ESI type one
    L2vpnEvpnEsi_esi_type1 L2vpnEvpnEsi = "esi-type1"

    // ESI type two
    L2vpnEvpnEsi_esi_type2 L2vpnEvpnEsi = "esi-type2"

    // ESI type three
    L2vpnEvpnEsi_esi_type3 L2vpnEvpnEsi = "esi-type3"

    // ESI type four
    L2vpnEvpnEsi_esi_type4 L2vpnEvpnEsi = "esi-type4"

    // ESI type five
    L2vpnEvpnEsi_esi_type5 L2vpnEvpnEsi = "esi-type5"

    // ESI type legacy
    L2vpnEvpnEsi_l2vpn_evpn_esi_type_legacy L2vpnEvpnEsi = "l2vpn-evpn-esi-type-legacy"

    // ESI type override (10-octet value)
    L2vpnEvpnEsi_l2vpn_evpn_esi_type_override L2vpnEvpnEsi = "l2vpn-evpn-esi-type-override"

    // ESI type invalid
    L2vpnEvpnEsi_esi_type_invalid L2vpnEvpnEsi = "esi-type-invalid"
)

// L2vpnTdmRtpOption represents L2VPN TDM RTP option
type L2vpnTdmRtpOption string

const (
    // Unknown RTP option
    L2vpnTdmRtpOption_unknown L2vpnTdmRtpOption = "unknown"

    // RTP option present
    L2vpnTdmRtpOption_present L2vpnTdmRtpOption = "present"

    // RTP option absent
    L2vpnTdmRtpOption_absent L2vpnTdmRtpOption = "absent"
)

// BgpRouteTargetRole represents Bgp route target role
type BgpRouteTargetRole string

const (
    // Both Import and export roles
    BgpRouteTargetRole_both BgpRouteTargetRole = "both"

    // Import role
    BgpRouteTargetRole_import_ BgpRouteTargetRole = "import"

    // Export role
    BgpRouteTargetRole_export BgpRouteTargetRole = "export"
)

// L2vpnEvpn represents L2vpn evpn
type L2vpnEvpn string

const (
    // Unspecify type for that EVI entry
    L2vpnEvpn_evpn_type_invalid L2vpnEvpn = "evpn-type-invalid"

    // EVPN service type
    L2vpnEvpn_evpn_type_evpn L2vpnEvpn = "evpn-type-evpn"

    // PBB EVPN service type
    L2vpnEvpn_evpn_type_pbb_evpn L2vpnEvpn = "evpn-type-pbb-evpn"

    // EVPN VPWS vlan-unaware service type
    L2vpnEvpn_evpn_type_evpn_vpws_vlan_unaware L2vpnEvpn = "evpn-type-evpn-vpws-vlan-unaware"

    // EVPN VPWS vlan-aware service type
    L2vpnEvpn_evpn_type_evpn_vpws_vlan_aware L2vpnEvpn = "evpn-type-evpn-vpws-vlan-aware"

    // Max EVPN type
    L2vpnEvpn_evpn_type_max L2vpnEvpn = "evpn-type-max"
)

// L2vpnFrMode represents L2vpn fr mode
type L2vpnFrMode string

const (
    // Frame Relay port mode
    L2vpnFrMode_l2vpn_fr_port_mode L2vpnFrMode = "l2vpn-fr-port-mode"

    // Frame Relay DLCI mode
    L2vpnFrMode_l2vpn_fr_dlci_mode L2vpnFrMode = "l2vpn-fr-dlci-mode"
)

// L2vpnEvpnRtOrigin represents L2vpn evpn rt origin
type L2vpnEvpnRtOrigin string

const (
    // Incomplete Configuration
    L2vpnEvpnRtOrigin_invalid L2vpnEvpnRtOrigin = "invalid"

    // From ESI
    L2vpnEvpnRtOrigin_extracted L2vpnEvpnRtOrigin = "extracted"

    // Locally configured
    L2vpnEvpnRtOrigin_configured L2vpnEvpnRtOrigin = "configured"
)

// L2vpnInterface represents L2vpn interface
type L2vpnInterface string

const (
    // Unknown
    L2vpnInterface_l2vpn_intf_type_unknown L2vpnInterface = "l2vpn-intf-type-unknown"

    // Ethernet
    L2vpnInterface_l2vpn_intf_type_ethernet L2vpnInterface = "l2vpn-intf-type-ethernet"

    // Ethernet Vlan
    L2vpnInterface_l2vpn_intf_type_vlan L2vpnInterface = "l2vpn-intf-type-vlan"

    // ATM
    L2vpnInterface_l2vpn_intf_type_atm L2vpnInterface = "l2vpn-intf-type-atm"

    // Frame Relay
    L2vpnInterface_l2vpn_intf_type_frame_relay L2vpnInterface = "l2vpn-intf-type-frame-relay"

    // HDLC
    L2vpnInterface_l2vpn_intf_type_hdlc L2vpnInterface = "l2vpn-intf-type-hdlc"

    // PPP
    L2vpnInterface_l2vpn_intf_type_ppp L2vpnInterface = "l2vpn-intf-type-ppp"

    // SPAN
    L2vpnInterface_l2vpn_intf_type_span L2vpnInterface = "l2vpn-intf-type-span"

    // BVI
    L2vpnInterface_l2vpn_intf_type_bvi L2vpnInterface = "l2vpn-intf-type-bvi"

    // CEM
    L2vpnInterface_l2vpn_intf_type_cem L2vpnInterface = "l2vpn-intf-type-cem"

    // PsuedowireEther
    L2vpnInterface_l2vpn_intf_type_pw_ether L2vpnInterface = "l2vpn-intf-type-pw-ether"

    // PsuedowireIW
    L2vpnInterface_l2vpn_intf_type_pw_iw L2vpnInterface = "l2vpn-intf-type-pw-iw"

    // VXLAN
    L2vpnInterface_l2vpn_intf_type_vni L2vpnInterface = "l2vpn-intf-type-vni"
)

// EvpnIgmpSource represents Evpn igmp source
type EvpnIgmpSource string

const (
    // Local
    EvpnIgmpSource_local EvpnIgmpSource = "local"

    // Remote
    EvpnIgmpSource_remote EvpnIgmpSource = "remote"
)

// L2vpnAtmMode represents L2vpn atm mode
type L2vpnAtmMode string

const (
    // ATM port mode
    L2vpnAtmMode_l2vpn_atm_port_mode L2vpnAtmMode = "l2vpn-atm-port-mode"

    // ATM Virtual Path mode
    L2vpnAtmMode_l2vpn_atm_vp_mode L2vpnAtmMode = "l2vpn-atm-vp-mode"

    // ATM Virtual Channel mode
    L2vpnAtmMode_l2vpn_atm_vc_mode L2vpnAtmMode = "l2vpn-atm-vc-mode"
)

// L2vpnAdRtRole represents L2vpn ad rt role
type L2vpnAdRtRole string

const (
    // Both
    L2vpnAdRtRole_both L2vpnAdRtRole = "both"

    // Import
    L2vpnAdRtRole_import_ L2vpnAdRtRole = "import"

    // Export
    L2vpnAdRtRole_export L2vpnAdRtRole = "export"
)

// ImStateEnum represents Im state enum
type ImStateEnum string

const (
    // im state not ready
    ImStateEnum_im_state_not_ready ImStateEnum = "im-state-not-ready"

    // im state admin down
    ImStateEnum_im_state_admin_down ImStateEnum = "im-state-admin-down"

    // im state down
    ImStateEnum_im_state_down ImStateEnum = "im-state-down"

    // im state up
    ImStateEnum_im_state_up ImStateEnum = "im-state-up"

    // im state shutdown
    ImStateEnum_im_state_shutdown ImStateEnum = "im-state-shutdown"

    // im state err disable
    ImStateEnum_im_state_err_disable ImStateEnum = "im-state-err-disable"

    // im state down immediate
    ImStateEnum_im_state_down_immediate ImStateEnum = "im-state-down-immediate"

    // im state down immediate admin
    ImStateEnum_im_state_down_immediate_admin ImStateEnum = "im-state-down-immediate-admin"

    // im state down graceful
    ImStateEnum_im_state_down_graceful ImStateEnum = "im-state-down-graceful"

    // im state begin shutdown
    ImStateEnum_im_state_begin_shutdown ImStateEnum = "im-state-begin-shutdown"

    // im state end shutdown
    ImStateEnum_im_state_end_shutdown ImStateEnum = "im-state-end-shutdown"

    // im state begin error disable
    ImStateEnum_im_state_begin_error_disable ImStateEnum = "im-state-begin-error-disable"

    // im state end error disable
    ImStateEnum_im_state_end_error_disable ImStateEnum = "im-state-end-error-disable"

    // im state begin down graceful
    ImStateEnum_im_state_begin_down_graceful ImStateEnum = "im-state-begin-down-graceful"

    // im state reset
    ImStateEnum_im_state_reset ImStateEnum = "im-state-reset"

    // im state operational
    ImStateEnum_im_state_operational ImStateEnum = "im-state-operational"

    // im state not operational
    ImStateEnum_im_state_not_operational ImStateEnum = "im-state-not-operational"

    // im state unknown
    ImStateEnum_im_state_unknown ImStateEnum = "im-state-unknown"

    // im state last
    ImStateEnum_im_state_last ImStateEnum = "im-state-last"
)

// L2vpnTdmMode represents L2VPN TDM modes
type L2vpnTdmMode string

const (
    // Unknown mode
    L2vpnTdmMode_unknown L2vpnTdmMode = "unknown"

    // CESoPSN mode
    L2vpnTdmMode_ce_so_psn L2vpnTdmMode = "ce-so-psn"

    // SAToP E1 mode
    L2vpnTdmMode_sa_to_p_e1 L2vpnTdmMode = "sa-to-p-e1"

    // SAToP T1 mode
    L2vpnTdmMode_sa_to_p_t1 L2vpnTdmMode = "sa-to-p-t1"

    // SAToP E3 mode
    L2vpnTdmMode_sa_to_p_e3 L2vpnTdmMode = "sa-to-p-e3"

    // SAToP T3 mode
    L2vpnTdmMode_sa_to_p_t3 L2vpnTdmMode = "sa-to-p-t3"
)

// IflistRepStatus represents Interface list replication status
type IflistRepStatus string

const (
    // Invalid
    IflistRepStatus_invalid IflistRepStatus = "invalid"

    // Pending
    IflistRepStatus_pending IflistRepStatus = "pending"

    // Done
    IflistRepStatus_done IflistRepStatus = "done"

    // Not supported
    IflistRepStatus_not_supported IflistRepStatus = "not-supported"

    // Failed
    IflistRepStatus_failed IflistRepStatus = "failed"
)

// EvpnIgmpGrp represents Evpn igmp grp
type EvpnIgmpGrp string

const (
    // Include
    EvpnIgmpGrp_include EvpnIgmpGrp = "include"

    // Exclude
    EvpnIgmpGrp_exclude EvpnIgmpGrp = "exclude"
)

// EvpnIgmpMsg represents Evpn igmp msg
type EvpnIgmpMsg string

const (
    // Join
    EvpnIgmpMsg_join EvpnIgmpMsg = "join"

    // Leave
    EvpnIgmpMsg_leave EvpnIgmpMsg = "leave"
)

// BgpRouteTargetFormat represents Bgp route target format
type BgpRouteTargetFormat string

const (
    // No route target
    BgpRouteTargetFormat_none BgpRouteTargetFormat = "none"

    // 2 Byte AS:nn format
    BgpRouteTargetFormat_two_byte_as BgpRouteTargetFormat = "two-byte-as"

    // 4 byte AS:nn format
    BgpRouteTargetFormat_four_byte_as BgpRouteTargetFormat = "four-byte-as"

    // IP:nn format
    BgpRouteTargetFormat_ipv4_address BgpRouteTargetFormat = "ipv4-address"

    // a.a.i format
    BgpRouteTargetFormat_es_import BgpRouteTargetFormat = "es-import"
)

// EvpnGrp represents EVPN Group State
type EvpnGrp string

const (
    // Deisolating
    EvpnGrp_deisolating EvpnGrp = "deisolating"

    // Isolated
    EvpnGrp_isolated EvpnGrp = "isolated"

    // Ready
    EvpnGrp_ready EvpnGrp = "ready"

    // Incomplete
    EvpnGrp_incomplete EvpnGrp = "incomplete"
)

// L2vpnEvpnSmacSrc represents L2vpn evpn smac src
type L2vpnEvpnSmacSrc string

const (
    // Incomplete Configuration
    L2vpnEvpnSmacSrc_invalid L2vpnEvpnSmacSrc = "invalid"

    // Source MAC Not Applicable (EVPN)
    L2vpnEvpnSmacSrc_not_applicable L2vpnEvpnSmacSrc = "not-applicable"

    // Local
    L2vpnEvpnSmacSrc_local L2vpnEvpnSmacSrc = "local"

    // PBB BSA
    L2vpnEvpnSmacSrc_pbb_bsa L2vpnEvpnSmacSrc = "pbb-bsa"

    // From ESI
    L2vpnEvpnSmacSrc_esi L2vpnEvpnSmacSrc = "esi"

    // From ESI, Error
    L2vpnEvpnSmacSrc_esi_invalid L2vpnEvpnSmacSrc = "esi-invalid"

    // PBB BSA, no ESI
    L2vpnEvpnSmacSrc_pbb_bsa_overrride L2vpnEvpnSmacSrc = "pbb-bsa-overrride"
)

// L2vpnAdRt represents L2vpn ad rt
type L2vpnAdRt string

const (
    // Route target not set
    L2vpnAdRt_l2vpn_ad_rt_none L2vpnAdRt = "l2vpn-ad-rt-none"

    // Route Target with 2 Byte AS number
    L2vpnAdRt_l2vpn_ad_rt_as L2vpnAdRt = "l2vpn-ad-rt-as"

    // Route Target with 4 Byte AS number
    L2vpnAdRt_l2vpn_ad_rt_4byte_as L2vpnAdRt = "l2vpn-ad-rt-4byte-as"

    // Route Target with IPv4 Address
    L2vpnAdRt_l2vpn_ad_rt_v4_addr L2vpnAdRt = "l2vpn-ad-rt-v4-addr"

    // Ethernet Segment Route Target from BGP
    L2vpnAdRt_es_import L2vpnAdRt = "es-import"
)

// EvpnIgmpVersion represents Evpn igmp version
type EvpnIgmpVersion string

const (
    // Version 1
    EvpnIgmpVersion_version1 EvpnIgmpVersion = "version1"

    // Version 2
    EvpnIgmpVersion_version2 EvpnIgmpVersion = "version2"

    // Version 3
    EvpnIgmpVersion_version3 EvpnIgmpVersion = "version3"
)

// L2vpnTimeStampMode represents L2VPN TDM Time stamp modes
type L2vpnTimeStampMode string

const (
    // Unknown time stamp mode
    L2vpnTimeStampMode_unknown L2vpnTimeStampMode = "unknown"

    // Differential time stamp mode
    L2vpnTimeStampMode_differential L2vpnTimeStampMode = "differential"

    // Absolute Time Stamp mode
    L2vpnTimeStampMode_absolute L2vpnTimeStampMode = "absolute"

    // time stamp mode none
    L2vpnTimeStampMode_none L2vpnTimeStampMode = "none"
)

// L2vpnAdRd represents L2vpn ad rd
type L2vpnAdRd string

const (
    // Route Distinguisher not set
    L2vpnAdRd_l2vpn_ad_rd_none L2vpnAdRd = "l2vpn-ad-rd-none"

    // Route Distinguisher auto-generated
    L2vpnAdRd_l2vpn_ad_rd_auto L2vpnAdRd = "l2vpn-ad-rd-auto"

    // Route Distinguisher with 2 Byte AS number
    L2vpnAdRd_l2vpn_ad_rd_as L2vpnAdRd = "l2vpn-ad-rd-as"

    // Route Distinguisher with 4 Byte AS number
    L2vpnAdRd_l2vpn_ad_rd_4byte_as L2vpnAdRd = "l2vpn-ad-rd-4byte-as"

    // Route Distinguisher with IPv4 Address
    L2vpnAdRd_l2vpn_ad_rd_v4_addr L2vpnAdRd = "l2vpn-ad-rd-v4-addr"
)

// Evpn
// EVPN Operational Table
type Evpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of EVPN operational data for a particular node.
    Nodes Evpn_Nodes

    // Active EVPN operational data.
    Active Evpn_Active

    // Standby EVPN operational data.
    Standby Evpn_Standby
}

func (evpn *Evpn) GetEntityData() *types.CommonEntityData {
    evpn.EntityData.YFilter = evpn.YFilter
    evpn.EntityData.YangName = "evpn"
    evpn.EntityData.BundleName = "cisco_ios_xr"
    evpn.EntityData.ParentYangName = "Cisco-IOS-XR-evpn-oper"
    evpn.EntityData.SegmentPath = "Cisco-IOS-XR-evpn-oper:evpn"
    evpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpn.EntityData.Children = types.NewOrderedMap()
    evpn.EntityData.Children.Append("nodes", types.YChild{"Nodes", &evpn.Nodes})
    evpn.EntityData.Children.Append("active", types.YChild{"Active", &evpn.Active})
    evpn.EntityData.Children.Append("standby", types.YChild{"Standby", &evpn.Standby})
    evpn.EntityData.Leafs = types.NewOrderedMap()

    evpn.EntityData.YListKeys = []string {}

    return &(evpn.EntityData)
}

// Evpn_Nodes
// Table of EVPN operational data for a particular
// node
type Evpn_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN operational data for a particular node. The type is slice of
    // Evpn_Nodes_Node.
    Node []*Evpn_Nodes_Node
}

func (nodes *Evpn_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "evpn"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Evpn_Nodes_Node
// EVPN operational data for a particular node
type Evpn_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // EVPN Group Table.
    EvpnGroups Evpn_Nodes_Node_EvpnGroups

    // EVPN Remote SHG table.
    RemoteShgs Evpn_Nodes_Node_RemoteShgs

    // L2VPN EVPN Client.
    Client Evpn_Nodes_Node_Client

    // EVPN IGMP table.
    Igmps Evpn_Nodes_Node_Igmps

    // L2VPN EVPN EVI Table.
    Evis Evpn_Nodes_Node_Evis

    // L2VPN EVPN Summary.
    Summary Evpn_Nodes_Node_Summary

    // L2VPN EVI Detail Table.
    EviDetail Evpn_Nodes_Node_EviDetail

    // L2VPN EVPN TEP Table.
    Teps Evpn_Nodes_Node_Teps

    // EVPN Internal Label Table.
    InternalLabels Evpn_Nodes_Node_InternalLabels

    // EVPN Ethernet-Segment Table.
    EthernetSegments Evpn_Nodes_Node_EthernetSegments

    // EVPN AC ID table.
    AcIds Evpn_Nodes_Node_AcIds
}

func (node *Evpn_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("evpn-groups", types.YChild{"EvpnGroups", &node.EvpnGroups})
    node.EntityData.Children.Append("remote-shgs", types.YChild{"RemoteShgs", &node.RemoteShgs})
    node.EntityData.Children.Append("client", types.YChild{"Client", &node.Client})
    node.EntityData.Children.Append("igmps", types.YChild{"Igmps", &node.Igmps})
    node.EntityData.Children.Append("evis", types.YChild{"Evis", &node.Evis})
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Children.Append("evi-detail", types.YChild{"EviDetail", &node.EviDetail})
    node.EntityData.Children.Append("teps", types.YChild{"Teps", &node.Teps})
    node.EntityData.Children.Append("internal-labels", types.YChild{"InternalLabels", &node.InternalLabels})
    node.EntityData.Children.Append("ethernet-segments", types.YChild{"EthernetSegments", &node.EthernetSegments})
    node.EntityData.Children.Append("ac-ids", types.YChild{"AcIds", &node.AcIds})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})

    node.EntityData.YListKeys = []string {"NodeId"}

    return &(node.EntityData)
}

// Evpn_Nodes_Node_EvpnGroups
// EVPN Group Table
type Evpn_Nodes_Node_EvpnGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Group information. The type is slice of
    // Evpn_Nodes_Node_EvpnGroups_EvpnGroup.
    EvpnGroup []*Evpn_Nodes_Node_EvpnGroups_EvpnGroup
}

func (evpnGroups *Evpn_Nodes_Node_EvpnGroups) GetEntityData() *types.CommonEntityData {
    evpnGroups.EntityData.YFilter = evpnGroups.YFilter
    evpnGroups.EntityData.YangName = "evpn-groups"
    evpnGroups.EntityData.BundleName = "cisco_ios_xr"
    evpnGroups.EntityData.ParentYangName = "node"
    evpnGroups.EntityData.SegmentPath = "evpn-groups"
    evpnGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroups.EntityData.Children = types.NewOrderedMap()
    evpnGroups.EntityData.Children.Append("evpn-group", types.YChild{"EvpnGroup", nil})
    for i := range evpnGroups.EvpnGroup {
        evpnGroups.EntityData.Children.Append(types.GetSegmentPath(evpnGroups.EvpnGroup[i]), types.YChild{"EvpnGroup", evpnGroups.EvpnGroup[i]})
    }
    evpnGroups.EntityData.Leafs = types.NewOrderedMap()

    evpnGroups.EntityData.YListKeys = []string {}

    return &(evpnGroups.EntityData)
}

// Evpn_Nodes_Node_EvpnGroups_EvpnGroup
// EVPN Group information
type Evpn_Nodes_Node_EvpnGroups_EvpnGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN group number. The type is interface{} with
    // range: 1..4294967295.
    GroupNumber interface{}

    // EVPN Group ID. The type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // EVPN Group State. The type is EvpnGrp.
    State interface{}

    // EVPN Group Core Interface table. The type is slice of
    // Evpn_Nodes_Node_EvpnGroups_EvpnGroup_CoreInterface.
    CoreInterface []*Evpn_Nodes_Node_EvpnGroups_EvpnGroup_CoreInterface

    // EVPN Access Core Interface table. The type is slice of
    // Evpn_Nodes_Node_EvpnGroups_EvpnGroup_AccessInterface.
    AccessInterface []*Evpn_Nodes_Node_EvpnGroups_EvpnGroup_AccessInterface
}

func (evpnGroup *Evpn_Nodes_Node_EvpnGroups_EvpnGroup) GetEntityData() *types.CommonEntityData {
    evpnGroup.EntityData.YFilter = evpnGroup.YFilter
    evpnGroup.EntityData.YangName = "evpn-group"
    evpnGroup.EntityData.BundleName = "cisco_ios_xr"
    evpnGroup.EntityData.ParentYangName = "evpn-groups"
    evpnGroup.EntityData.SegmentPath = "evpn-group" + types.AddKeyToken(evpnGroup.GroupNumber, "group-number")
    evpnGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroup.EntityData.Children = types.NewOrderedMap()
    evpnGroup.EntityData.Children.Append("core-interface", types.YChild{"CoreInterface", nil})
    for i := range evpnGroup.CoreInterface {
        evpnGroup.EntityData.Children.Append(types.GetSegmentPath(evpnGroup.CoreInterface[i]), types.YChild{"CoreInterface", evpnGroup.CoreInterface[i]})
    }
    evpnGroup.EntityData.Children.Append("access-interface", types.YChild{"AccessInterface", nil})
    for i := range evpnGroup.AccessInterface {
        evpnGroup.EntityData.Children.Append(types.GetSegmentPath(evpnGroup.AccessInterface[i]), types.YChild{"AccessInterface", evpnGroup.AccessInterface[i]})
    }
    evpnGroup.EntityData.Leafs = types.NewOrderedMap()
    evpnGroup.EntityData.Leafs.Append("group-number", types.YLeaf{"GroupNumber", evpnGroup.GroupNumber})
    evpnGroup.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", evpnGroup.GroupId})
    evpnGroup.EntityData.Leafs.Append("state", types.YLeaf{"State", evpnGroup.State})

    evpnGroup.EntityData.YListKeys = []string {"GroupNumber"}

    return &(evpnGroup.EntityData)
}

// Evpn_Nodes_Node_EvpnGroups_EvpnGroup_CoreInterface
// EVPN Group Core Interface table
type Evpn_Nodes_Node_EvpnGroups_EvpnGroup_CoreInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface State. The type is ImStateEnum.
    State interface{}
}

func (coreInterface *Evpn_Nodes_Node_EvpnGroups_EvpnGroup_CoreInterface) GetEntityData() *types.CommonEntityData {
    coreInterface.EntityData.YFilter = coreInterface.YFilter
    coreInterface.EntityData.YangName = "core-interface"
    coreInterface.EntityData.BundleName = "cisco_ios_xr"
    coreInterface.EntityData.ParentYangName = "evpn-group"
    coreInterface.EntityData.SegmentPath = "core-interface"
    coreInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coreInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coreInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coreInterface.EntityData.Children = types.NewOrderedMap()
    coreInterface.EntityData.Leafs = types.NewOrderedMap()
    coreInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", coreInterface.InterfaceName})
    coreInterface.EntityData.Leafs.Append("state", types.YLeaf{"State", coreInterface.State})

    coreInterface.EntityData.YListKeys = []string {}

    return &(coreInterface.EntityData)
}

// Evpn_Nodes_Node_EvpnGroups_EvpnGroup_AccessInterface
// EVPN Access Core Interface table
type Evpn_Nodes_Node_EvpnGroups_EvpnGroup_AccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface State. The type is ImStateEnum.
    State interface{}
}

func (accessInterface *Evpn_Nodes_Node_EvpnGroups_EvpnGroup_AccessInterface) GetEntityData() *types.CommonEntityData {
    accessInterface.EntityData.YFilter = accessInterface.YFilter
    accessInterface.EntityData.YangName = "access-interface"
    accessInterface.EntityData.BundleName = "cisco_ios_xr"
    accessInterface.EntityData.ParentYangName = "evpn-group"
    accessInterface.EntityData.SegmentPath = "access-interface"
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = types.NewOrderedMap()
    accessInterface.EntityData.Leafs = types.NewOrderedMap()
    accessInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterface.InterfaceName})
    accessInterface.EntityData.Leafs.Append("state", types.YLeaf{"State", accessInterface.State})

    accessInterface.EntityData.YListKeys = []string {}

    return &(accessInterface.EntityData)
}

// Evpn_Nodes_Node_RemoteShgs
// EVPN Remote SHG table
type Evpn_Nodes_Node_RemoteShgs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Remote SHG. The type is slice of Evpn_Nodes_Node_RemoteShgs_RemoteShg.
    RemoteShg []*Evpn_Nodes_Node_RemoteShgs_RemoteShg
}

func (remoteShgs *Evpn_Nodes_Node_RemoteShgs) GetEntityData() *types.CommonEntityData {
    remoteShgs.EntityData.YFilter = remoteShgs.YFilter
    remoteShgs.EntityData.YangName = "remote-shgs"
    remoteShgs.EntityData.BundleName = "cisco_ios_xr"
    remoteShgs.EntityData.ParentYangName = "node"
    remoteShgs.EntityData.SegmentPath = "remote-shgs"
    remoteShgs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteShgs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteShgs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteShgs.EntityData.Children = types.NewOrderedMap()
    remoteShgs.EntityData.Children.Append("remote-shg", types.YChild{"RemoteShg", nil})
    for i := range remoteShgs.RemoteShg {
        remoteShgs.EntityData.Children.Append(types.GetSegmentPath(remoteShgs.RemoteShg[i]), types.YChild{"RemoteShg", remoteShgs.RemoteShg[i]})
    }
    remoteShgs.EntityData.Leafs = types.NewOrderedMap()

    remoteShgs.EntityData.YListKeys = []string {}

    return &(remoteShgs.EntityData)
}

// Evpn_Nodes_Node_RemoteShgs_RemoteShg
// EVPN Remote SHG
type Evpn_Nodes_Node_RemoteShgs_RemoteShg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Segment id. The type is slice of
    // Evpn_Nodes_Node_RemoteShgs_RemoteShg_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Nodes_Node_RemoteShgs_RemoteShg_EthernetSegmentIdentifier

    // Remote split horizon group labels. The type is slice of
    // Evpn_Nodes_Node_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel.
    RemoteSplitHorizonGroupLabel []*Evpn_Nodes_Node_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel
}

func (remoteShg *Evpn_Nodes_Node_RemoteShgs_RemoteShg) GetEntityData() *types.CommonEntityData {
    remoteShg.EntityData.YFilter = remoteShg.YFilter
    remoteShg.EntityData.YangName = "remote-shg"
    remoteShg.EntityData.BundleName = "cisco_ios_xr"
    remoteShg.EntityData.ParentYangName = "remote-shgs"
    remoteShg.EntityData.SegmentPath = "remote-shg"
    remoteShg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteShg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteShg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteShg.EntityData.Children = types.NewOrderedMap()
    remoteShg.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range remoteShg.EthernetSegmentIdentifier {
        remoteShg.EntityData.Children.Append(types.GetSegmentPath(remoteShg.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", remoteShg.EthernetSegmentIdentifier[i]})
    }
    remoteShg.EntityData.Children.Append("remote-split-horizon-group-label", types.YChild{"RemoteSplitHorizonGroupLabel", nil})
    for i := range remoteShg.RemoteSplitHorizonGroupLabel {
        remoteShg.EntityData.Children.Append(types.GetSegmentPath(remoteShg.RemoteSplitHorizonGroupLabel[i]), types.YChild{"RemoteSplitHorizonGroupLabel", remoteShg.RemoteSplitHorizonGroupLabel[i]})
    }
    remoteShg.EntityData.Leafs = types.NewOrderedMap()
    remoteShg.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", remoteShg.Esi1})
    remoteShg.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", remoteShg.Esi2})
    remoteShg.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", remoteShg.Esi3})
    remoteShg.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", remoteShg.Esi4})
    remoteShg.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", remoteShg.Esi5})

    remoteShg.EntityData.YListKeys = []string {}

    return &(remoteShg.EntityData)
}

// Evpn_Nodes_Node_RemoteShgs_RemoteShg_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Nodes_Node_RemoteShgs_RemoteShg_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_RemoteShgs_RemoteShg_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "remote-shg"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Nodes_Node_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Nodes_Node_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Split horizon label associated with next-hop address. The type is
    // interface{} with range: 0..4294967295.
    Label interface{}
}

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel) GetEntityData() *types.CommonEntityData {
    remoteSplitHorizonGroupLabel.EntityData.YFilter = remoteSplitHorizonGroupLabel.YFilter
    remoteSplitHorizonGroupLabel.EntityData.YangName = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.BundleName = "cisco_ios_xr"
    remoteSplitHorizonGroupLabel.EntityData.ParentYangName = "remote-shg"
    remoteSplitHorizonGroupLabel.EntityData.SegmentPath = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteSplitHorizonGroupLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteSplitHorizonGroupLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteSplitHorizonGroupLabel.EntityData.Children = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", remoteSplitHorizonGroupLabel.NextHop})
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("label", types.YLeaf{"Label", remoteSplitHorizonGroupLabel.Label})

    remoteSplitHorizonGroupLabel.EntityData.YListKeys = []string {}

    return &(remoteSplitHorizonGroupLabel.EntityData)
}

// Evpn_Nodes_Node_Client
// L2VPN EVPN Client
type Evpn_Nodes_Node_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (client *Evpn_Nodes_Node_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "node"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// Evpn_Nodes_Node_Igmps
// EVPN IGMP table
type Evpn_Nodes_Node_Igmps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Remote. The type is slice of Evpn_Nodes_Node_Igmps_Igmp.
    Igmp []*Evpn_Nodes_Node_Igmps_Igmp
}

func (igmps *Evpn_Nodes_Node_Igmps) GetEntityData() *types.CommonEntityData {
    igmps.EntityData.YFilter = igmps.YFilter
    igmps.EntityData.YangName = "igmps"
    igmps.EntityData.BundleName = "cisco_ios_xr"
    igmps.EntityData.ParentYangName = "node"
    igmps.EntityData.SegmentPath = "igmps"
    igmps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igmps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igmps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igmps.EntityData.Children = types.NewOrderedMap()
    igmps.EntityData.Children.Append("igmp", types.YChild{"Igmp", nil})
    for i := range igmps.Igmp {
        igmps.EntityData.Children.Append(types.GetSegmentPath(igmps.Igmp[i]), types.YChild{"Igmp", igmps.Igmp[i]})
    }
    igmps.EntityData.Leafs = types.NewOrderedMap()

    igmps.EntityData.YListKeys = []string {}

    return &(igmps.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp
// EVPN Remote
type Evpn_Nodes_Node_Igmps_Igmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Join=0, Leave=1. The type is interface{} with range: 0..4294967295.
    IsLeave interface{}

    // BP xcid. The type is interface{} with range: 0..4294967295.
    Bpxcid interface{}

    // EVI/BD. The type is interface{} with range: 0..4294967295.
    Evibd interface{}

    // Source IP Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SrcIp interface{}

    // Group IP Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GrpIp interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Segment Name. The type is string.
    EthernetSegmentName interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // BD id. The type is interface{} with range: 0..4294967295.
    BdId interface{}

    // Route Type. The type is EvpnIgmpMsg.
    RouteType interface{}

    // Source IP Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddr interface{}

    // Group IP Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddr interface{}

    // Ethernet Tag id. The type is interface{} with range: 0..4294967295.
    EthernetTagId interface{}

    // IGMP Version. The type is EvpnIgmpVersion.
    IgmpVersion interface{}

    // IGMP Group Type. The type is EvpnIgmpGrp.
    IgmpGroupType interface{}

    // Max Response Time. The type is interface{} with range: 0..255.
    MaXResponseTime interface{}

    // Resolved. The type is bool.
    Resolved interface{}

    // Source Info.
    SourceInfo Evpn_Nodes_Node_Igmps_Igmp_SourceInfo

    // Ethernet Segment id. The type is slice of
    // Evpn_Nodes_Node_Igmps_Igmp_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Nodes_Node_Igmps_Igmp_EthernetSegmentIdentifier

    // List of nexthop IPv6 addresses. The type is slice of
    // Evpn_Nodes_Node_Igmps_Igmp_NextHop.
    NextHop []*Evpn_Nodes_Node_Igmps_Igmp_NextHop
}

func (igmp *Evpn_Nodes_Node_Igmps_Igmp) GetEntityData() *types.CommonEntityData {
    igmp.EntityData.YFilter = igmp.YFilter
    igmp.EntityData.YangName = "igmp"
    igmp.EntityData.BundleName = "cisco_ios_xr"
    igmp.EntityData.ParentYangName = "igmps"
    igmp.EntityData.SegmentPath = "igmp"
    igmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igmp.EntityData.Children = types.NewOrderedMap()
    igmp.EntityData.Children.Append("source-info", types.YChild{"SourceInfo", &igmp.SourceInfo})
    igmp.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range igmp.EthernetSegmentIdentifier {
        igmp.EntityData.Children.Append(types.GetSegmentPath(igmp.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", igmp.EthernetSegmentIdentifier[i]})
    }
    igmp.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range igmp.NextHop {
        igmp.EntityData.Children.Append(types.GetSegmentPath(igmp.NextHop[i]), types.YChild{"NextHop", igmp.NextHop[i]})
    }
    igmp.EntityData.Leafs = types.NewOrderedMap()
    igmp.EntityData.Leafs.Append("is-leave", types.YLeaf{"IsLeave", igmp.IsLeave})
    igmp.EntityData.Leafs.Append("bpxcid", types.YLeaf{"Bpxcid", igmp.Bpxcid})
    igmp.EntityData.Leafs.Append("evibd", types.YLeaf{"Evibd", igmp.Evibd})
    igmp.EntityData.Leafs.Append("src-ip", types.YLeaf{"SrcIp", igmp.SrcIp})
    igmp.EntityData.Leafs.Append("grp-ip", types.YLeaf{"GrpIp", igmp.GrpIp})
    igmp.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", igmp.Esi1})
    igmp.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", igmp.Esi2})
    igmp.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", igmp.Esi3})
    igmp.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", igmp.Esi4})
    igmp.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", igmp.Esi5})
    igmp.EntityData.Leafs.Append("ethernet-segment-name", types.YLeaf{"EthernetSegmentName", igmp.EthernetSegmentName})
    igmp.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", igmp.Evi})
    igmp.EntityData.Leafs.Append("bd-id", types.YLeaf{"BdId", igmp.BdId})
    igmp.EntityData.Leafs.Append("route-type", types.YLeaf{"RouteType", igmp.RouteType})
    igmp.EntityData.Leafs.Append("source-addr", types.YLeaf{"SourceAddr", igmp.SourceAddr})
    igmp.EntityData.Leafs.Append("group-addr", types.YLeaf{"GroupAddr", igmp.GroupAddr})
    igmp.EntityData.Leafs.Append("ethernet-tag-id", types.YLeaf{"EthernetTagId", igmp.EthernetTagId})
    igmp.EntityData.Leafs.Append("igmp-version", types.YLeaf{"IgmpVersion", igmp.IgmpVersion})
    igmp.EntityData.Leafs.Append("igmp-group-type", types.YLeaf{"IgmpGroupType", igmp.IgmpGroupType})
    igmp.EntityData.Leafs.Append("ma-x-response-time", types.YLeaf{"MaXResponseTime", igmp.MaXResponseTime})
    igmp.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", igmp.Resolved})

    igmp.EntityData.YListKeys = []string {}

    return &(igmp.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo
// Source Info
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is EvpnIgmpSource.
    Type interface{}

    // remote info. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteInfo interface{}

    // local info.
    LocalInfo Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo
}

func (sourceInfo *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo) GetEntityData() *types.CommonEntityData {
    sourceInfo.EntityData.YFilter = sourceInfo.YFilter
    sourceInfo.EntityData.YangName = "source-info"
    sourceInfo.EntityData.BundleName = "cisco_ios_xr"
    sourceInfo.EntityData.ParentYangName = "igmp"
    sourceInfo.EntityData.SegmentPath = "source-info"
    sourceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceInfo.EntityData.Children = types.NewOrderedMap()
    sourceInfo.EntityData.Children.Append("local-info", types.YChild{"LocalInfo", &sourceInfo.LocalInfo})
    sourceInfo.EntityData.Leafs = types.NewOrderedMap()
    sourceInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", sourceInfo.Type})
    sourceInfo.EntityData.Leafs.Append("remote-info", types.YLeaf{"RemoteInfo", sourceInfo.RemoteInfo})

    sourceInfo.EntityData.YListKeys = []string {}

    return &(sourceInfo.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo
// local info
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..81.
    Name interface{}

    // Interface MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Payload bytes. The type is interface{} with range: 0..65535. Units are
    // byte.
    PayloadBytes interface{}

    // Interface parameters.
    Parameters Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters
}

func (localInfo *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo) GetEntityData() *types.CommonEntityData {
    localInfo.EntityData.YFilter = localInfo.YFilter
    localInfo.EntityData.YangName = "local-info"
    localInfo.EntityData.BundleName = "cisco_ios_xr"
    localInfo.EntityData.ParentYangName = "source-info"
    localInfo.EntityData.SegmentPath = "local-info"
    localInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localInfo.EntityData.Children = types.NewOrderedMap()
    localInfo.EntityData.Children.Append("parameters", types.YChild{"Parameters", &localInfo.Parameters})
    localInfo.EntityData.Leafs = types.NewOrderedMap()
    localInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", localInfo.Name})
    localInfo.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", localInfo.Mtu})
    localInfo.EntityData.Leafs.Append("payload-bytes", types.YLeaf{"PayloadBytes", localInfo.PayloadBytes})

    localInfo.EntityData.YListKeys = []string {}

    return &(localInfo.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters
// Interface parameters
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is L2vpnInterface.
    Type interface{}

    // Ethernet.
    Ethernet Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet

    // VLAN.
    Vlan Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan

    // TDM.
    Tdm Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm

    // ATM.
    Atm Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm

    // Frame Relay.
    Fr Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr

    // PW Ether.
    PseudowireEther Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther

    // PW IW.
    PseudowireIw Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw
}

func (parameters *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters) GetEntityData() *types.CommonEntityData {
    parameters.EntityData.YFilter = parameters.YFilter
    parameters.EntityData.YangName = "parameters"
    parameters.EntityData.BundleName = "cisco_ios_xr"
    parameters.EntityData.ParentYangName = "local-info"
    parameters.EntityData.SegmentPath = "parameters"
    parameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parameters.EntityData.Children = types.NewOrderedMap()
    parameters.EntityData.Children.Append("ethernet", types.YChild{"Ethernet", &parameters.Ethernet})
    parameters.EntityData.Children.Append("vlan", types.YChild{"Vlan", &parameters.Vlan})
    parameters.EntityData.Children.Append("tdm", types.YChild{"Tdm", &parameters.Tdm})
    parameters.EntityData.Children.Append("atm", types.YChild{"Atm", &parameters.Atm})
    parameters.EntityData.Children.Append("fr", types.YChild{"Fr", &parameters.Fr})
    parameters.EntityData.Children.Append("pseudowire-ether", types.YChild{"PseudowireEther", &parameters.PseudowireEther})
    parameters.EntityData.Children.Append("pseudowire-iw", types.YChild{"PseudowireIw", &parameters.PseudowireIw})
    parameters.EntityData.Leafs = types.NewOrderedMap()
    parameters.EntityData.Leafs.Append("type", types.YLeaf{"Type", parameters.Type})

    parameters.EntityData.YListKeys = []string {}

    return &(parameters.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet
// Ethernet
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XConnect tags. The type is interface{} with range: 0..255.
    XconnectTags interface{}
}

func (ethernet *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "cisco_ios_xr"
    ethernet.EntityData.ParentYangName = "parameters"
    ethernet.EntityData.SegmentPath = "ethernet"
    ethernet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernet.EntityData.Children = types.NewOrderedMap()
    ethernet.EntityData.Leafs = types.NewOrderedMap()
    ethernet.EntityData.Leafs.Append("xconnect-tags", types.YLeaf{"XconnectTags", ethernet.XconnectTags})

    ethernet.EntityData.YListKeys = []string {}

    return &(ethernet.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan
// VLAN
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XConnect tags. The type is interface{} with range: 0..255.
    XconnectTags interface{}

    // VLAN rewrite tag. The type is interface{} with range: 0..65535.
    VlanRewriteTag interface{}

    // Simple EFP. The type is interface{} with range: 0..255.
    SimpleEfp interface{}

    // Encapsulation Type. The type is interface{} with range: 0..255.
    EncapsulationType interface{}

    // Outer Tag. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Rewrite Tags. The type is slice of
    // Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag.
    RewriteTag []*Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag

    // vlan range. The type is slice of
    // Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange.
    VlanRange []*Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange
}

func (vlan *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "cisco_ios_xr"
    vlan.EntityData.ParentYangName = "parameters"
    vlan.EntityData.SegmentPath = "vlan"
    vlan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlan.EntityData.Children = types.NewOrderedMap()
    vlan.EntityData.Children.Append("rewrite-tag", types.YChild{"RewriteTag", nil})
    for i := range vlan.RewriteTag {
        vlan.EntityData.Children.Append(types.GetSegmentPath(vlan.RewriteTag[i]), types.YChild{"RewriteTag", vlan.RewriteTag[i]})
    }
    vlan.EntityData.Children.Append("vlan-range", types.YChild{"VlanRange", nil})
    for i := range vlan.VlanRange {
        vlan.EntityData.Children.Append(types.GetSegmentPath(vlan.VlanRange[i]), types.YChild{"VlanRange", vlan.VlanRange[i]})
    }
    vlan.EntityData.Leafs = types.NewOrderedMap()
    vlan.EntityData.Leafs.Append("xconnect-tags", types.YLeaf{"XconnectTags", vlan.XconnectTags})
    vlan.EntityData.Leafs.Append("vlan-rewrite-tag", types.YLeaf{"VlanRewriteTag", vlan.VlanRewriteTag})
    vlan.EntityData.Leafs.Append("simple-efp", types.YLeaf{"SimpleEfp", vlan.SimpleEfp})
    vlan.EntityData.Leafs.Append("encapsulation-type", types.YLeaf{"EncapsulationType", vlan.EncapsulationType})
    vlan.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", vlan.OuterTag})

    vlan.EntityData.YListKeys = []string {}

    return &(vlan.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag
// Rewrite Tags
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535.
    Entry interface{}
}

func (rewriteTag *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag) GetEntityData() *types.CommonEntityData {
    rewriteTag.EntityData.YFilter = rewriteTag.YFilter
    rewriteTag.EntityData.YangName = "rewrite-tag"
    rewriteTag.EntityData.BundleName = "cisco_ios_xr"
    rewriteTag.EntityData.ParentYangName = "vlan"
    rewriteTag.EntityData.SegmentPath = "rewrite-tag"
    rewriteTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rewriteTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rewriteTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rewriteTag.EntityData.Children = types.NewOrderedMap()
    rewriteTag.EntityData.Leafs = types.NewOrderedMap()
    rewriteTag.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", rewriteTag.Entry})

    rewriteTag.EntityData.YListKeys = []string {}

    return &(rewriteTag.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange
// vlan range
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lower. The type is interface{} with range: 0..65535.
    Lower interface{}

    // Upper. The type is interface{} with range: 0..65535.
    Upper interface{}
}

func (vlanRange *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange) GetEntityData() *types.CommonEntityData {
    vlanRange.EntityData.YFilter = vlanRange.YFilter
    vlanRange.EntityData.YangName = "vlan-range"
    vlanRange.EntityData.BundleName = "cisco_ios_xr"
    vlanRange.EntityData.ParentYangName = "vlan"
    vlanRange.EntityData.SegmentPath = "vlan-range"
    vlanRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanRange.EntityData.Children = types.NewOrderedMap()
    vlanRange.EntityData.Leafs = types.NewOrderedMap()
    vlanRange.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", vlanRange.Lower})
    vlanRange.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", vlanRange.Upper})

    vlanRange.EntityData.YListKeys = []string {}

    return &(vlanRange.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm
// TDM
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeslots separated by , or - from 1 to 31. : indicates individual timeslot
    // and - represents a range.E.g. 1-3,5 represents timeslots 1, 2, 3, and 5.
    // The type is string.
    TimeslotGroup interface{}

    // Timeslot rate in units of Kbps. The type is interface{} with range: 0..255.
    // Units are kbit/s.
    TimeslotRate interface{}

    // TDM mode. The type is L2vpnTdmMode.
    TdmMode interface{}

    // TDM options.
    TdmOptions Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions
}

func (tdm *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm) GetEntityData() *types.CommonEntityData {
    tdm.EntityData.YFilter = tdm.YFilter
    tdm.EntityData.YangName = "tdm"
    tdm.EntityData.BundleName = "cisco_ios_xr"
    tdm.EntityData.ParentYangName = "parameters"
    tdm.EntityData.SegmentPath = "tdm"
    tdm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tdm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tdm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tdm.EntityData.Children = types.NewOrderedMap()
    tdm.EntityData.Children.Append("tdm-options", types.YChild{"TdmOptions", &tdm.TdmOptions})
    tdm.EntityData.Leafs = types.NewOrderedMap()
    tdm.EntityData.Leafs.Append("timeslot-group", types.YLeaf{"TimeslotGroup", tdm.TimeslotGroup})
    tdm.EntityData.Leafs.Append("timeslot-rate", types.YLeaf{"TimeslotRate", tdm.TimeslotRate})
    tdm.EntityData.Leafs.Append("tdm-mode", types.YLeaf{"TdmMode", tdm.TdmMode})

    tdm.EntityData.YListKeys = []string {}

    return &(tdm.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions
// TDM options
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TDM payload bytes. The type is interface{} with range: 0..65535. Units are
    // byte.
    PayloadBytes interface{}

    // TDM bit rate in units of Kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    BitRate interface{}

    // RTP header. The type is L2vpnTdmRtpOption.
    Rtp interface{}

    // TDM Timestamping mode. The type is L2vpnTimeStampMode.
    TimestampMode interface{}

    // Signalling packets. The type is interface{} with range: 0..255.
    SignallingPackets interface{}

    // CAS. The type is interface{} with range: 0..255.
    Cas interface{}

    // RTP header payload type. The type is interface{} with range: 0..255.
    RtpHeaderPayloadType interface{}

    // Timestamping clock frequency in units of 8Khz. The type is interface{} with
    // range: 0..65535.
    TimestampClockFreq interface{}

    // Synchronization Source identifier. The type is interface{} with range:
    // 0..4294967295.
    Ssrc interface{}
}

func (tdmOptions *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions) GetEntityData() *types.CommonEntityData {
    tdmOptions.EntityData.YFilter = tdmOptions.YFilter
    tdmOptions.EntityData.YangName = "tdm-options"
    tdmOptions.EntityData.BundleName = "cisco_ios_xr"
    tdmOptions.EntityData.ParentYangName = "tdm"
    tdmOptions.EntityData.SegmentPath = "tdm-options"
    tdmOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tdmOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tdmOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tdmOptions.EntityData.Children = types.NewOrderedMap()
    tdmOptions.EntityData.Leafs = types.NewOrderedMap()
    tdmOptions.EntityData.Leafs.Append("payload-bytes", types.YLeaf{"PayloadBytes", tdmOptions.PayloadBytes})
    tdmOptions.EntityData.Leafs.Append("bit-rate", types.YLeaf{"BitRate", tdmOptions.BitRate})
    tdmOptions.EntityData.Leafs.Append("rtp", types.YLeaf{"Rtp", tdmOptions.Rtp})
    tdmOptions.EntityData.Leafs.Append("timestamp-mode", types.YLeaf{"TimestampMode", tdmOptions.TimestampMode})
    tdmOptions.EntityData.Leafs.Append("signalling-packets", types.YLeaf{"SignallingPackets", tdmOptions.SignallingPackets})
    tdmOptions.EntityData.Leafs.Append("cas", types.YLeaf{"Cas", tdmOptions.Cas})
    tdmOptions.EntityData.Leafs.Append("rtp-header-payload-type", types.YLeaf{"RtpHeaderPayloadType", tdmOptions.RtpHeaderPayloadType})
    tdmOptions.EntityData.Leafs.Append("timestamp-clock-freq", types.YLeaf{"TimestampClockFreq", tdmOptions.TimestampClockFreq})
    tdmOptions.EntityData.Leafs.Append("ssrc", types.YLeaf{"Ssrc", tdmOptions.Ssrc})

    tdmOptions.EntityData.YListKeys = []string {}

    return &(tdmOptions.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm
// ATM
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max number of cells packed. The type is interface{} with range: 0..65535.
    MaximumNumberCellsPacked interface{}

    // Max number of cells unpacked. The type is interface{} with range: 0..65535.
    MaximumNumberCellsUnPacked interface{}

    // ATM mode. The type is L2vpnAtmMode.
    AtmMode interface{}

    // Virtual path identifier. The type is interface{} with range: 0..65535.
    Vpi interface{}

    // Virtual channel identifier. The type is interface{} with range: 0..65535.
    Vci interface{}
}

func (atm *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm) GetEntityData() *types.CommonEntityData {
    atm.EntityData.YFilter = atm.YFilter
    atm.EntityData.YangName = "atm"
    atm.EntityData.BundleName = "cisco_ios_xr"
    atm.EntityData.ParentYangName = "parameters"
    atm.EntityData.SegmentPath = "atm"
    atm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    atm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    atm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    atm.EntityData.Children = types.NewOrderedMap()
    atm.EntityData.Leafs = types.NewOrderedMap()
    atm.EntityData.Leafs.Append("maximum-number-cells-packed", types.YLeaf{"MaximumNumberCellsPacked", atm.MaximumNumberCellsPacked})
    atm.EntityData.Leafs.Append("maximum-number-cells-un-packed", types.YLeaf{"MaximumNumberCellsUnPacked", atm.MaximumNumberCellsUnPacked})
    atm.EntityData.Leafs.Append("atm-mode", types.YLeaf{"AtmMode", atm.AtmMode})
    atm.EntityData.Leafs.Append("vpi", types.YLeaf{"Vpi", atm.Vpi})
    atm.EntityData.Leafs.Append("vci", types.YLeaf{"Vci", atm.Vci})

    atm.EntityData.YListKeys = []string {}

    return &(atm.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr
// Frame Relay
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Frame Relay mode. The type is L2vpnFrMode.
    FrMode interface{}

    // Data-link connection identifier. The type is interface{} with range:
    // 0..4294967295.
    Dlci interface{}
}

func (fr *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr) GetEntityData() *types.CommonEntityData {
    fr.EntityData.YFilter = fr.YFilter
    fr.EntityData.YangName = "fr"
    fr.EntityData.BundleName = "cisco_ios_xr"
    fr.EntityData.ParentYangName = "parameters"
    fr.EntityData.SegmentPath = "fr"
    fr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fr.EntityData.Children = types.NewOrderedMap()
    fr.EntityData.Leafs = types.NewOrderedMap()
    fr.EntityData.Leafs.Append("fr-mode", types.YLeaf{"FrMode", fr.FrMode})
    fr.EntityData.Leafs.Append("dlci", types.YLeaf{"Dlci", fr.Dlci})

    fr.EntityData.YListKeys = []string {}

    return &(fr.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther
// PW Ether
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is this Interface list valid. The type is bool.
    IsValid interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Interface list data.
    InterfaceList Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList
}

func (pseudowireEther *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther) GetEntityData() *types.CommonEntityData {
    pseudowireEther.EntityData.YFilter = pseudowireEther.YFilter
    pseudowireEther.EntityData.YangName = "pseudowire-ether"
    pseudowireEther.EntityData.BundleName = "cisco_ios_xr"
    pseudowireEther.EntityData.ParentYangName = "parameters"
    pseudowireEther.EntityData.SegmentPath = "pseudowire-ether"
    pseudowireEther.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireEther.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireEther.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireEther.EntityData.Children = types.NewOrderedMap()
    pseudowireEther.EntityData.Children.Append("interface-list", types.YChild{"InterfaceList", &pseudowireEther.InterfaceList})
    pseudowireEther.EntityData.Leafs = types.NewOrderedMap()
    pseudowireEther.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", pseudowireEther.IsValid})
    pseudowireEther.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", pseudowireEther.InternalLabel})

    pseudowireEther.EntityData.YListKeys = []string {}

    return &(pseudowireEther.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList
// Interface list data
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface-list name. The type is string with length: 0..33.
    InterfaceListName interface{}

    // Interface internal ID. The type is interface{} with range: 0..4294967295.
    InterfaceListId interface{}

    // Interfaces. The type is slice of
    // Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface.
    Interface []*Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface
}

func (interfaceList *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList) GetEntityData() *types.CommonEntityData {
    interfaceList.EntityData.YFilter = interfaceList.YFilter
    interfaceList.EntityData.YangName = "interface-list"
    interfaceList.EntityData.BundleName = "cisco_ios_xr"
    interfaceList.EntityData.ParentYangName = "pseudowire-ether"
    interfaceList.EntityData.SegmentPath = "interface-list"
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = types.NewOrderedMap()
    interfaceList.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceList.Interface {
        interfaceList.EntityData.Children.Append(types.GetSegmentPath(interfaceList.Interface[i]), types.YChild{"Interface", interfaceList.Interface[i]})
    }
    interfaceList.EntityData.Leafs = types.NewOrderedMap()
    interfaceList.EntityData.Leafs.Append("interface-list-name", types.YLeaf{"InterfaceListName", interfaceList.InterfaceListName})
    interfaceList.EntityData.Leafs.Append("interface-list-id", types.YLeaf{"InterfaceListId", interfaceList.InterfaceListId})

    interfaceList.EntityData.YListKeys = []string {}

    return &(interfaceList.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface
// Interfaces
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Replicate status. The type is IflistRepStatus.
    ReplicateStatus interface{}
}

func (self *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interface-list"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("replicate-status", types.YLeaf{"ReplicateStatus", self.ReplicateStatus})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw
// PW IW
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is this Interface list valid. The type is bool.
    IsValid interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Interface list data.
    InterfaceList Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList
}

func (pseudowireIw *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw) GetEntityData() *types.CommonEntityData {
    pseudowireIw.EntityData.YFilter = pseudowireIw.YFilter
    pseudowireIw.EntityData.YangName = "pseudowire-iw"
    pseudowireIw.EntityData.BundleName = "cisco_ios_xr"
    pseudowireIw.EntityData.ParentYangName = "parameters"
    pseudowireIw.EntityData.SegmentPath = "pseudowire-iw"
    pseudowireIw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireIw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireIw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireIw.EntityData.Children = types.NewOrderedMap()
    pseudowireIw.EntityData.Children.Append("interface-list", types.YChild{"InterfaceList", &pseudowireIw.InterfaceList})
    pseudowireIw.EntityData.Leafs = types.NewOrderedMap()
    pseudowireIw.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", pseudowireIw.IsValid})
    pseudowireIw.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", pseudowireIw.InternalLabel})

    pseudowireIw.EntityData.YListKeys = []string {}

    return &(pseudowireIw.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList
// Interface list data
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface-list name. The type is string with length: 0..33.
    InterfaceListName interface{}

    // Interface internal ID. The type is interface{} with range: 0..4294967295.
    InterfaceListId interface{}

    // Interfaces. The type is slice of
    // Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface.
    Interface []*Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface
}

func (interfaceList *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList) GetEntityData() *types.CommonEntityData {
    interfaceList.EntityData.YFilter = interfaceList.YFilter
    interfaceList.EntityData.YangName = "interface-list"
    interfaceList.EntityData.BundleName = "cisco_ios_xr"
    interfaceList.EntityData.ParentYangName = "pseudowire-iw"
    interfaceList.EntityData.SegmentPath = "interface-list"
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = types.NewOrderedMap()
    interfaceList.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceList.Interface {
        interfaceList.EntityData.Children.Append(types.GetSegmentPath(interfaceList.Interface[i]), types.YChild{"Interface", interfaceList.Interface[i]})
    }
    interfaceList.EntityData.Leafs = types.NewOrderedMap()
    interfaceList.EntityData.Leafs.Append("interface-list-name", types.YLeaf{"InterfaceListName", interfaceList.InterfaceListName})
    interfaceList.EntityData.Leafs.Append("interface-list-id", types.YLeaf{"InterfaceListId", interfaceList.InterfaceListId})

    interfaceList.EntityData.YListKeys = []string {}

    return &(interfaceList.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface
// Interfaces
type Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Replicate status. The type is IflistRepStatus.
    ReplicateStatus interface{}
}

func (self *Evpn_Nodes_Node_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interface-list"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("replicate-status", types.YLeaf{"ReplicateStatus", self.ReplicateStatus})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Nodes_Node_Igmps_Igmp_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_Igmps_Igmp_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "igmp"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Nodes_Node_Igmps_Igmp_NextHop
// List of nexthop IPv6 addresses
type Evpn_Nodes_Node_Igmps_Igmp_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // DF Dont Premption. The type is bool.
    DfDontPrempt interface{}

    // DF Election Mode Configured. The type is interface{} with range: 0..255.
    DfType interface{}

    // DF Election Preference Set. The type is interface{} with range: 0..65535.
    DfPref interface{}
}

func (nextHop *Evpn_Nodes_Node_Igmps_Igmp_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "igmp"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", nextHop.NextHop})
    nextHop.EntityData.Leafs.Append("df-dont-prempt", types.YLeaf{"DfDontPrempt", nextHop.DfDontPrempt})
    nextHop.EntityData.Leafs.Append("df-type", types.YLeaf{"DfType", nextHop.DfType})
    nextHop.EntityData.Leafs.Append("df-pref", types.YLeaf{"DfPref", nextHop.DfPref})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// Evpn_Nodes_Node_Evis
// L2VPN EVPN EVI Table
type Evpn_Nodes_Node_Evis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Entry. The type is slice of Evpn_Nodes_Node_Evis_Evi.
    Evi []*Evpn_Nodes_Node_Evis_Evi
}

func (evis *Evpn_Nodes_Node_Evis) GetEntityData() *types.CommonEntityData {
    evis.EntityData.YFilter = evis.YFilter
    evis.EntityData.YangName = "evis"
    evis.EntityData.BundleName = "cisco_ios_xr"
    evis.EntityData.ParentYangName = "node"
    evis.EntityData.SegmentPath = "evis"
    evis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evis.EntityData.Children = types.NewOrderedMap()
    evis.EntityData.Children.Append("evi", types.YChild{"Evi", nil})
    for i := range evis.Evi {
        evis.EntityData.Children.Append(types.GetSegmentPath(evis.Evi[i]), types.YChild{"Evi", evis.Evi[i]})
    }
    evis.EntityData.Leafs = types.NewOrderedMap()

    evis.EntityData.YListKeys = []string {}

    return &(evis.EntityData)
}

// Evpn_Nodes_Node_Evis_Evi
// L2VPN EVPN EVI Entry
type Evpn_Nodes_Node_Evis_Evi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evi *Evpn_Nodes_Node_Evis_Evi) GetEntityData() *types.CommonEntityData {
    evi.EntityData.YFilter = evi.YFilter
    evi.EntityData.YangName = "evi"
    evi.EntityData.BundleName = "cisco_ios_xr"
    evi.EntityData.ParentYangName = "evis"
    evi.EntityData.SegmentPath = "evi"
    evi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evi.EntityData.Children = types.NewOrderedMap()
    evi.EntityData.Leafs = types.NewOrderedMap()
    evi.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", evi.Evi})
    evi.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", evi.Encapsulation})
    evi.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evi.EthernetVpnId})
    evi.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evi.EncapsulationXr})
    evi.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evi.BdName})
    evi.EntityData.Leafs.Append("type", types.YLeaf{"Type", evi.Type})

    evi.EntityData.YListKeys = []string {}

    return &(evi.EntityData)
}

// Evpn_Nodes_Node_Summary
// L2VPN EVPN Summary
type Evpn_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Router ID. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // BGP AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // Number of EVI DB Entries. The type is interface{} with range:
    // 0..4294967295.
    EvIs interface{}

    // Number of Tunnel Endpoint DB Entries. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpoints interface{}

    // Number of Local MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalMacRoutes interface{}

    // Number of Local IPv4 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalIpv4MacRoutes interface{}

    // Number of Local IPv6 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalIpv6MacRoutes interface{}

    // Number of ES:Global MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    EsGlobalMacRoutes interface{}

    // Number of Remote MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteMacRoutes interface{}

    // Number of Remote Soo MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteSooMacRoutes interface{}

    // Number of Remote IPv4 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteIpv4MacRoutes interface{}

    // Number of Remote IPv6 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteIpv6MacRoutes interface{}

    // Number of Local IMCAST Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalImcastRoutes interface{}

    // Number of Remote IMCAST Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteImcastRoutes interface{}

    // Number of Internal Labels. The type is interface{} with range:
    // 0..4294967295.
    Labels interface{}

    // Number of ES Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    EsEntries interface{}

    // Number of neighbor Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    NeighborEntries interface{}

    // Number of Local EAD Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    LocalEadRoutes interface{}

    // Number of Remote EAD Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    RemoteEadRoutes interface{}

    // Global Source MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    GlobalSourceMac interface{}

    // EVPN ES Peering Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    PeeringTime interface{}

    // EVPN ES Recovery Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    RecoveryTime interface{}

    // EVPN ES Carving Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    CarvingTime interface{}

    // Number of moves within the move interval before locking the MAC. The type
    // is interface{} with range: 0..4294967295.
    MacSecureMoveCount interface{}

    // Interval to watch for subsequent mac moves before locking the MAC. The type
    // is interface{} with range: 0..4294967295.
    MacSecureMoveInterval interface{}

    // Length of time to lock the mac after a MAC security violation. The type is
    // interface{} with range: 0..4294967295.
    MacSecureFreezeTime interface{}

    // Number of times to retry after a MAC un-freezes. The type is interface{}
    // with range: 0..4294967295.
    MacSecureRetryCount interface{}

    // EVPN Node Cost-out. The type is bool.
    CostOut interface{}

    // EVPN Node startup cost-in Time (minutes). The type is interface{} with
    // range: 0..4294967295. Units are minute.
    StartupCostInTime interface{}

    // Send to L2RIB Throttled. The type is bool.
    L2ribThrottle interface{}

    // Logging EVPN Designated Forwarder changes enabled. The type is bool.
    LoggingDfElectionEnabled interface{}
}

func (summary *Evpn_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", summary.RouterId})
    summary.EntityData.Leafs.Append("as", types.YLeaf{"As", summary.As})
    summary.EntityData.Leafs.Append("ev-is", types.YLeaf{"EvIs", summary.EvIs})
    summary.EntityData.Leafs.Append("tunnel-endpoints", types.YLeaf{"TunnelEndpoints", summary.TunnelEndpoints})
    summary.EntityData.Leafs.Append("local-mac-routes", types.YLeaf{"LocalMacRoutes", summary.LocalMacRoutes})
    summary.EntityData.Leafs.Append("local-ipv4-mac-routes", types.YLeaf{"LocalIpv4MacRoutes", summary.LocalIpv4MacRoutes})
    summary.EntityData.Leafs.Append("local-ipv6-mac-routes", types.YLeaf{"LocalIpv6MacRoutes", summary.LocalIpv6MacRoutes})
    summary.EntityData.Leafs.Append("es-global-mac-routes", types.YLeaf{"EsGlobalMacRoutes", summary.EsGlobalMacRoutes})
    summary.EntityData.Leafs.Append("remote-mac-routes", types.YLeaf{"RemoteMacRoutes", summary.RemoteMacRoutes})
    summary.EntityData.Leafs.Append("remote-soo-mac-routes", types.YLeaf{"RemoteSooMacRoutes", summary.RemoteSooMacRoutes})
    summary.EntityData.Leafs.Append("remote-ipv4-mac-routes", types.YLeaf{"RemoteIpv4MacRoutes", summary.RemoteIpv4MacRoutes})
    summary.EntityData.Leafs.Append("remote-ipv6-mac-routes", types.YLeaf{"RemoteIpv6MacRoutes", summary.RemoteIpv6MacRoutes})
    summary.EntityData.Leafs.Append("local-imcast-routes", types.YLeaf{"LocalImcastRoutes", summary.LocalImcastRoutes})
    summary.EntityData.Leafs.Append("remote-imcast-routes", types.YLeaf{"RemoteImcastRoutes", summary.RemoteImcastRoutes})
    summary.EntityData.Leafs.Append("labels", types.YLeaf{"Labels", summary.Labels})
    summary.EntityData.Leafs.Append("es-entries", types.YLeaf{"EsEntries", summary.EsEntries})
    summary.EntityData.Leafs.Append("neighbor-entries", types.YLeaf{"NeighborEntries", summary.NeighborEntries})
    summary.EntityData.Leafs.Append("local-ead-routes", types.YLeaf{"LocalEadRoutes", summary.LocalEadRoutes})
    summary.EntityData.Leafs.Append("remote-ead-routes", types.YLeaf{"RemoteEadRoutes", summary.RemoteEadRoutes})
    summary.EntityData.Leafs.Append("global-source-mac", types.YLeaf{"GlobalSourceMac", summary.GlobalSourceMac})
    summary.EntityData.Leafs.Append("peering-time", types.YLeaf{"PeeringTime", summary.PeeringTime})
    summary.EntityData.Leafs.Append("recovery-time", types.YLeaf{"RecoveryTime", summary.RecoveryTime})
    summary.EntityData.Leafs.Append("carving-time", types.YLeaf{"CarvingTime", summary.CarvingTime})
    summary.EntityData.Leafs.Append("mac-secure-move-count", types.YLeaf{"MacSecureMoveCount", summary.MacSecureMoveCount})
    summary.EntityData.Leafs.Append("mac-secure-move-interval", types.YLeaf{"MacSecureMoveInterval", summary.MacSecureMoveInterval})
    summary.EntityData.Leafs.Append("mac-secure-freeze-time", types.YLeaf{"MacSecureFreezeTime", summary.MacSecureFreezeTime})
    summary.EntityData.Leafs.Append("mac-secure-retry-count", types.YLeaf{"MacSecureRetryCount", summary.MacSecureRetryCount})
    summary.EntityData.Leafs.Append("cost-out", types.YLeaf{"CostOut", summary.CostOut})
    summary.EntityData.Leafs.Append("startup-cost-in-time", types.YLeaf{"StartupCostInTime", summary.StartupCostInTime})
    summary.EntityData.Leafs.Append("l2rib-throttle", types.YLeaf{"L2ribThrottle", summary.L2ribThrottle})
    summary.EntityData.Leafs.Append("logging-df-election-enabled", types.YLeaf{"LoggingDfElectionEnabled", summary.LoggingDfElectionEnabled})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Evpn_Nodes_Node_EviDetail
// L2VPN EVI Detail Table
type Evpn_Nodes_Node_EviDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info Elements.
    Elements Evpn_Nodes_Node_EviDetail_Elements

    // Container for all EVI detail info.
    EviChildren Evpn_Nodes_Node_EviDetail_EviChildren
}

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetEntityData() *types.CommonEntityData {
    eviDetail.EntityData.YFilter = eviDetail.YFilter
    eviDetail.EntityData.YangName = "evi-detail"
    eviDetail.EntityData.BundleName = "cisco_ios_xr"
    eviDetail.EntityData.ParentYangName = "node"
    eviDetail.EntityData.SegmentPath = "evi-detail"
    eviDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviDetail.EntityData.Children = types.NewOrderedMap()
    eviDetail.EntityData.Children.Append("elements", types.YChild{"Elements", &eviDetail.Elements})
    eviDetail.EntityData.Children.Append("evi-children", types.YChild{"EviChildren", &eviDetail.EviChildren})
    eviDetail.EntityData.Leafs = types.NewOrderedMap()

    eviDetail.EntityData.YListKeys = []string {}

    return &(eviDetail.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements
// EVI BGP RT Detail Info Elements
type Evpn_Nodes_Node_EviDetail_Elements struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info. The type is slice of
    // Evpn_Nodes_Node_EviDetail_Elements_Element.
    Element []*Evpn_Nodes_Node_EviDetail_Elements_Element
}

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetEntityData() *types.CommonEntityData {
    elements.EntityData.YFilter = elements.YFilter
    elements.EntityData.YangName = "elements"
    elements.EntityData.BundleName = "cisco_ios_xr"
    elements.EntityData.ParentYangName = "evi-detail"
    elements.EntityData.SegmentPath = "elements"
    elements.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    elements.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    elements.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    elements.EntityData.Children = types.NewOrderedMap()
    elements.EntityData.Children.Append("element", types.YChild{"Element", nil})
    for i := range elements.Element {
        elements.EntityData.Children.Append(types.GetSegmentPath(elements.Element[i]), types.YChild{"Element", elements.Element[i]})
    }
    elements.EntityData.Leafs = types.NewOrderedMap()

    elements.EntityData.YListKeys = []string {}

    return &(elements.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element
// EVI BGP RT Detail Info
type Evpn_Nodes_Node_EviDetail_Elements_Element struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // EVPN Instance encapsulation. The type is interface{} with range: 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // EVI description. The type is string.
    Description interface{}

    // Unicast Label. The type is interface{} with range: 0..4294967295.
    UnicastLabel interface{}

    // Multicast Label. The type is interface{} with range: 0..4294967295.
    MulticastLabel interface{}

    // Control-Word Disable. The type is bool.
    CwDisable interface{}

    // Table-policy Name. The type is string.
    TablePolicyName interface{}

    // Forward Class attribute. The type is interface{} with range: 0..255.
    ForwardClass interface{}

    // Is Import RT None set. The type is bool.
    RtImportBlockSet interface{}

    // Is Export RT None set. The type is bool.
    RtExportBlockSet interface{}

    // Advertise MAC-only routes on this EVI. The type is bool.
    AdvertiseMac interface{}

    // Advertise BVI MACs routes on this EVI. The type is bool.
    AdvertiseBviMac interface{}

    // Route Aliasing is disabled. The type is bool.
    AliasingDisabled interface{}

    // Unknown-unicast flooding is disabled. The type is bool.
    UnknownUnicastFloodingDisabled interface{}

    // Route Re-origination is disabled. The type is bool.
    ReoriginateDisabled interface{}

    // EVPN Instance is Regular/Stitching side. The type is bool.
    Stitching interface{}

    // EVI is connected to multicast source. The type is bool.
    MulticastSourceConnected interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Nodes_Node_EviDetail_Elements_Element_EvpnInstance

    // Flow Label Information.
    FlowLabel Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel

    // Automatic Route Distingtuisher.
    RdAuto Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto

    // Configured Route Distinguisher.
    RdConfigured Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured

    // Automatic Route Target.
    RtAuto Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto
}

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetEntityData() *types.CommonEntityData {
    element.EntityData.YFilter = element.YFilter
    element.EntityData.YangName = "element"
    element.EntityData.BundleName = "cisco_ios_xr"
    element.EntityData.ParentYangName = "elements"
    element.EntityData.SegmentPath = "element"
    element.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    element.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    element.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    element.EntityData.Children = types.NewOrderedMap()
    element.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &element.EvpnInstance})
    element.EntityData.Children.Append("flow-label", types.YChild{"FlowLabel", &element.FlowLabel})
    element.EntityData.Children.Append("rd-auto", types.YChild{"RdAuto", &element.RdAuto})
    element.EntityData.Children.Append("rd-configured", types.YChild{"RdConfigured", &element.RdConfigured})
    element.EntityData.Children.Append("rt-auto", types.YChild{"RtAuto", &element.RtAuto})
    element.EntityData.Leafs = types.NewOrderedMap()
    element.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", element.Evi})
    element.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", element.Encapsulation})
    element.EntityData.Leafs.Append("evi-xr", types.YLeaf{"EviXr", element.EviXr})
    element.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", element.EncapsulationXr})
    element.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", element.BdName})
    element.EntityData.Leafs.Append("type", types.YLeaf{"Type", element.Type})
    element.EntityData.Leafs.Append("description", types.YLeaf{"Description", element.Description})
    element.EntityData.Leafs.Append("unicast-label", types.YLeaf{"UnicastLabel", element.UnicastLabel})
    element.EntityData.Leafs.Append("multicast-label", types.YLeaf{"MulticastLabel", element.MulticastLabel})
    element.EntityData.Leafs.Append("cw-disable", types.YLeaf{"CwDisable", element.CwDisable})
    element.EntityData.Leafs.Append("table-policy-name", types.YLeaf{"TablePolicyName", element.TablePolicyName})
    element.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", element.ForwardClass})
    element.EntityData.Leafs.Append("rt-import-block-set", types.YLeaf{"RtImportBlockSet", element.RtImportBlockSet})
    element.EntityData.Leafs.Append("rt-export-block-set", types.YLeaf{"RtExportBlockSet", element.RtExportBlockSet})
    element.EntityData.Leafs.Append("advertise-mac", types.YLeaf{"AdvertiseMac", element.AdvertiseMac})
    element.EntityData.Leafs.Append("advertise-bvi-mac", types.YLeaf{"AdvertiseBviMac", element.AdvertiseBviMac})
    element.EntityData.Leafs.Append("aliasing-disabled", types.YLeaf{"AliasingDisabled", element.AliasingDisabled})
    element.EntityData.Leafs.Append("unknown-unicast-flooding-disabled", types.YLeaf{"UnknownUnicastFloodingDisabled", element.UnknownUnicastFloodingDisabled})
    element.EntityData.Leafs.Append("reoriginate-disabled", types.YLeaf{"ReoriginateDisabled", element.ReoriginateDisabled})
    element.EntityData.Leafs.Append("stitching", types.YLeaf{"Stitching", element.Stitching})
    element.EntityData.Leafs.Append("multicast-source-connected", types.YLeaf{"MulticastSourceConnected", element.MulticastSourceConnected})

    element.EntityData.YListKeys = []string {}

    return &(element.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_EvpnInstance
// EVPN Instance summary information
type Evpn_Nodes_Node_EviDetail_Elements_Element_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Nodes_Node_EviDetail_Elements_Element_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "element"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel
// Flow Label Information
type Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static flow label. The type is bool.
    StaticFlowLabel interface{}

    // Globally configured flow label. The type is bool.
    GlobalFlowLabel interface{}
}

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetEntityData() *types.CommonEntityData {
    flowLabel.EntityData.YFilter = flowLabel.YFilter
    flowLabel.EntityData.YangName = "flow-label"
    flowLabel.EntityData.BundleName = "cisco_ios_xr"
    flowLabel.EntityData.ParentYangName = "element"
    flowLabel.EntityData.SegmentPath = "flow-label"
    flowLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabel.EntityData.Children = types.NewOrderedMap()
    flowLabel.EntityData.Leafs = types.NewOrderedMap()
    flowLabel.EntityData.Leafs.Append("static-flow-label", types.YLeaf{"StaticFlowLabel", flowLabel.StaticFlowLabel})
    flowLabel.EntityData.Leafs.Append("global-flow-label", types.YLeaf{"GlobalFlowLabel", flowLabel.GlobalFlowLabel})

    flowLabel.EntityData.YListKeys = []string {}

    return &(flowLabel.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto
// Automatic Route Distingtuisher
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RD. The type is L2vpnAdRd.
    Rd interface{}

    // auto.
    Auto Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto

    // two byte as.
    TwoByteAs Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs

    // v4 addr.
    V4Addr Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr
}

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetEntityData() *types.CommonEntityData {
    rdAuto.EntityData.YFilter = rdAuto.YFilter
    rdAuto.EntityData.YangName = "rd-auto"
    rdAuto.EntityData.BundleName = "cisco_ios_xr"
    rdAuto.EntityData.ParentYangName = "element"
    rdAuto.EntityData.SegmentPath = "rd-auto"
    rdAuto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdAuto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdAuto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdAuto.EntityData.Children = types.NewOrderedMap()
    rdAuto.EntityData.Children.Append("auto", types.YChild{"Auto", &rdAuto.Auto})
    rdAuto.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &rdAuto.TwoByteAs})
    rdAuto.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &rdAuto.FourByteAs})
    rdAuto.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &rdAuto.V4Addr})
    rdAuto.EntityData.Leafs = types.NewOrderedMap()
    rdAuto.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", rdAuto.Rd})

    rdAuto.EntityData.YListKeys = []string {}

    return &(rdAuto.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto
// auto
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetEntityData() *types.CommonEntityData {
    auto.EntityData.YFilter = auto.YFilter
    auto.EntityData.YangName = "auto"
    auto.EntityData.BundleName = "cisco_ios_xr"
    auto.EntityData.ParentYangName = "rd-auto"
    auto.EntityData.SegmentPath = "auto"
    auto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auto.EntityData.Children = types.NewOrderedMap()
    auto.EntityData.Leafs = types.NewOrderedMap()
    auto.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", auto.RouterId})
    auto.EntityData.Leafs.Append("auto-index", types.YLeaf{"AutoIndex", auto.AutoIndex})

    auto.EntityData.YListKeys = []string {}

    return &(auto.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rd-auto"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rd-auto"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rd-auto"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured
// Configured Route Distinguisher
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RD. The type is L2vpnAdRd.
    Rd interface{}

    // auto.
    Auto Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto

    // two byte as.
    TwoByteAs Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs

    // v4 addr.
    V4Addr Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr
}

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetEntityData() *types.CommonEntityData {
    rdConfigured.EntityData.YFilter = rdConfigured.YFilter
    rdConfigured.EntityData.YangName = "rd-configured"
    rdConfigured.EntityData.BundleName = "cisco_ios_xr"
    rdConfigured.EntityData.ParentYangName = "element"
    rdConfigured.EntityData.SegmentPath = "rd-configured"
    rdConfigured.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdConfigured.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdConfigured.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdConfigured.EntityData.Children = types.NewOrderedMap()
    rdConfigured.EntityData.Children.Append("auto", types.YChild{"Auto", &rdConfigured.Auto})
    rdConfigured.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &rdConfigured.TwoByteAs})
    rdConfigured.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &rdConfigured.FourByteAs})
    rdConfigured.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &rdConfigured.V4Addr})
    rdConfigured.EntityData.Leafs = types.NewOrderedMap()
    rdConfigured.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", rdConfigured.Rd})

    rdConfigured.EntityData.YListKeys = []string {}

    return &(rdConfigured.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto
// auto
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetEntityData() *types.CommonEntityData {
    auto.EntityData.YFilter = auto.YFilter
    auto.EntityData.YangName = "auto"
    auto.EntityData.BundleName = "cisco_ios_xr"
    auto.EntityData.ParentYangName = "rd-configured"
    auto.EntityData.SegmentPath = "auto"
    auto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auto.EntityData.Children = types.NewOrderedMap()
    auto.EntityData.Leafs = types.NewOrderedMap()
    auto.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", auto.RouterId})
    auto.EntityData.Leafs.Append("auto-index", types.YLeaf{"AutoIndex", auto.AutoIndex})

    auto.EntityData.YListKeys = []string {}

    return &(auto.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rd-configured"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rd-configured"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rd-configured"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto
// Automatic Route Target
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs

    // v4 addr.
    V4Addr Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr

    // es import.
    EsImport Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport
}

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetEntityData() *types.CommonEntityData {
    rtAuto.EntityData.YFilter = rtAuto.YFilter
    rtAuto.EntityData.YangName = "rt-auto"
    rtAuto.EntityData.BundleName = "cisco_ios_xr"
    rtAuto.EntityData.ParentYangName = "element"
    rtAuto.EntityData.SegmentPath = "rt-auto"
    rtAuto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rtAuto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rtAuto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rtAuto.EntityData.Children = types.NewOrderedMap()
    rtAuto.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &rtAuto.TwoByteAs})
    rtAuto.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &rtAuto.FourByteAs})
    rtAuto.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &rtAuto.V4Addr})
    rtAuto.EntityData.Children.Append("es-import", types.YChild{"EsImport", &rtAuto.EsImport})
    rtAuto.EntityData.Leafs = types.NewOrderedMap()
    rtAuto.EntityData.Leafs.Append("rt", types.YLeaf{"Rt", rtAuto.Rt})

    rtAuto.EntityData.YListKeys = []string {}

    return &(rtAuto.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rt-auto"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rt-auto"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rt-auto"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport
// es import
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "rt-auto"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = types.NewOrderedMap()
    esImport.EntityData.Leafs = types.NewOrderedMap()
    esImport.EntityData.Leafs.Append("high-bytes", types.YLeaf{"HighBytes", esImport.HighBytes})
    esImport.EntityData.Leafs.Append("low-bytes", types.YLeaf{"LowBytes", esImport.LowBytes})

    esImport.EntityData.YListKeys = []string {}

    return &(esImport.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren
// Container for all EVI detail info
type Evpn_Nodes_Node_EviDetail_EviChildren struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Neighbor table.
    Neighbors Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors

    // EVPN Ethernet Auto-Discovery table.
    EthernetAutoDiscoveries Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries

    // L2VPN EVPN IMCAST table.
    InclusiveMulticasts Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts

    // L2VPN EVPN EVI RT Child Table.
    RouteTargets Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets

    // L2VPN EVPN EVI MAC table.
    Macs Evpn_Nodes_Node_EviDetail_EviChildren_Macs
}

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetEntityData() *types.CommonEntityData {
    eviChildren.EntityData.YFilter = eviChildren.YFilter
    eviChildren.EntityData.YangName = "evi-children"
    eviChildren.EntityData.BundleName = "cisco_ios_xr"
    eviChildren.EntityData.ParentYangName = "evi-detail"
    eviChildren.EntityData.SegmentPath = "evi-children"
    eviChildren.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviChildren.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviChildren.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviChildren.EntityData.Children = types.NewOrderedMap()
    eviChildren.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &eviChildren.Neighbors})
    eviChildren.EntityData.Children.Append("ethernet-auto-discoveries", types.YChild{"EthernetAutoDiscoveries", &eviChildren.EthernetAutoDiscoveries})
    eviChildren.EntityData.Children.Append("inclusive-multicasts", types.YChild{"InclusiveMulticasts", &eviChildren.InclusiveMulticasts})
    eviChildren.EntityData.Children.Append("route-targets", types.YChild{"RouteTargets", &eviChildren.RouteTargets})
    eviChildren.EntityData.Children.Append("macs", types.YChild{"Macs", &eviChildren.Macs})
    eviChildren.EntityData.Leafs = types.NewOrderedMap()

    eviChildren.EntityData.YListKeys = []string {}

    return &(eviChildren.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors
// EVPN Neighbor table
type Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Neighbor table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor.
    Neighbor []*Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor
}

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "evi-children"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor
// EVPN Neighbor table
type Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Neighbor IP. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborIp interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance
}

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &neighbor.EvpnInstance})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", neighbor.Evi})
    neighbor.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", neighbor.Encapsulation})
    neighbor.EntityData.Leafs.Append("neighbor-ip", types.YLeaf{"NeighborIp", neighbor.NeighborIp})
    neighbor.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", neighbor.Neighbor})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance
// EVPN Instance summary information
type Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "neighbor"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries
// EVPN Ethernet Auto-Discovery table
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet Auto-Discovery Entry. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery.
    EthernetAutoDiscovery []*Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
}

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetEntityData() *types.CommonEntityData {
    ethernetAutoDiscoveries.EntityData.YFilter = ethernetAutoDiscoveries.YFilter
    ethernetAutoDiscoveries.EntityData.YangName = "ethernet-auto-discoveries"
    ethernetAutoDiscoveries.EntityData.BundleName = "cisco_ios_xr"
    ethernetAutoDiscoveries.EntityData.ParentYangName = "evi-children"
    ethernetAutoDiscoveries.EntityData.SegmentPath = "ethernet-auto-discoveries"
    ethernetAutoDiscoveries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetAutoDiscoveries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetAutoDiscoveries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetAutoDiscoveries.EntityData.Children = types.NewOrderedMap()
    ethernetAutoDiscoveries.EntityData.Children.Append("ethernet-auto-discovery", types.YChild{"EthernetAutoDiscovery", nil})
    for i := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
        ethernetAutoDiscoveries.EntityData.Children.Append(types.GetSegmentPath(ethernetAutoDiscoveries.EthernetAutoDiscovery[i]), types.YChild{"EthernetAutoDiscovery", ethernetAutoDiscoveries.EthernetAutoDiscovery[i]})
    }
    ethernetAutoDiscoveries.EntityData.Leafs = types.NewOrderedMap()

    ethernetAutoDiscoveries.EntityData.YListKeys = []string {}

    return &(ethernetAutoDiscoveries.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
// EVPN Ethernet Auto-Discovery Entry
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Local nexthop IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalNextHop interface{}

    // Associated local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Indication of EthernetAutoDiscovery Route is local. The type is bool.
    IsLocalEad interface{}

    // Single-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleActive interface{}

    // Single-flow-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleFlowActive interface{}

    // Number of items in path list buffer. The type is interface{} with range:
    // 0..4294967295.
    NumPaths interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance

    // Ethernet Segment id. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer.
    PathBuffer []*Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
}

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetEntityData() *types.CommonEntityData {
    ethernetAutoDiscovery.EntityData.YFilter = ethernetAutoDiscovery.YFilter
    ethernetAutoDiscovery.EntityData.YangName = "ethernet-auto-discovery"
    ethernetAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    ethernetAutoDiscovery.EntityData.ParentYangName = "ethernet-auto-discoveries"
    ethernetAutoDiscovery.EntityData.SegmentPath = "ethernet-auto-discovery"
    ethernetAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetAutoDiscovery.EntityData.Children = types.NewOrderedMap()
    ethernetAutoDiscovery.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &ethernetAutoDiscovery.EvpnInstance})
    ethernetAutoDiscovery.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
        ethernetAutoDiscovery.EntityData.Children.Append(types.GetSegmentPath(ethernetAutoDiscovery.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", ethernetAutoDiscovery.EthernetSegmentIdentifier[i]})
    }
    ethernetAutoDiscovery.EntityData.Children.Append("path-buffer", types.YChild{"PathBuffer", nil})
    for i := range ethernetAutoDiscovery.PathBuffer {
        ethernetAutoDiscovery.EntityData.Children.Append(types.GetSegmentPath(ethernetAutoDiscovery.PathBuffer[i]), types.YChild{"PathBuffer", ethernetAutoDiscovery.PathBuffer[i]})
    }
    ethernetAutoDiscovery.EntityData.Leafs = types.NewOrderedMap()
    ethernetAutoDiscovery.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", ethernetAutoDiscovery.Evi})
    ethernetAutoDiscovery.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", ethernetAutoDiscovery.Encapsulation})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", ethernetAutoDiscovery.Esi1})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", ethernetAutoDiscovery.Esi2})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", ethernetAutoDiscovery.Esi3})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", ethernetAutoDiscovery.Esi4})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", ethernetAutoDiscovery.Esi5})
    ethernetAutoDiscovery.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", ethernetAutoDiscovery.EthernetTag})
    ethernetAutoDiscovery.EntityData.Leafs.Append("ethernet-tag-xr", types.YLeaf{"EthernetTagXr", ethernetAutoDiscovery.EthernetTagXr})
    ethernetAutoDiscovery.EntityData.Leafs.Append("local-next-hop", types.YLeaf{"LocalNextHop", ethernetAutoDiscovery.LocalNextHop})
    ethernetAutoDiscovery.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", ethernetAutoDiscovery.LocalLabel})
    ethernetAutoDiscovery.EntityData.Leafs.Append("is-local-ead", types.YLeaf{"IsLocalEad", ethernetAutoDiscovery.IsLocalEad})
    ethernetAutoDiscovery.EntityData.Leafs.Append("redundancy-single-active", types.YLeaf{"RedundancySingleActive", ethernetAutoDiscovery.RedundancySingleActive})
    ethernetAutoDiscovery.EntityData.Leafs.Append("redundancy-single-flow-active", types.YLeaf{"RedundancySingleFlowActive", ethernetAutoDiscovery.RedundancySingleFlowActive})
    ethernetAutoDiscovery.EntityData.Leafs.Append("num-paths", types.YLeaf{"NumPaths", ethernetAutoDiscovery.NumPaths})

    ethernetAutoDiscovery.EntityData.YListKeys = []string {}

    return &(ethernetAutoDiscovery.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance
// EVPN Instance summary information
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "ethernet-auto-discovery"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "ethernet-auto-discovery"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
// Path List Buffer
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetEntityData() *types.CommonEntityData {
    pathBuffer.EntityData.YFilter = pathBuffer.YFilter
    pathBuffer.EntityData.YangName = "path-buffer"
    pathBuffer.EntityData.BundleName = "cisco_ios_xr"
    pathBuffer.EntityData.ParentYangName = "ethernet-auto-discovery"
    pathBuffer.EntityData.SegmentPath = "path-buffer"
    pathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathBuffer.EntityData.Children = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", pathBuffer.TunnelEndpointId})
    pathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", pathBuffer.NextHop})
    pathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", pathBuffer.OutputLabel})
    pathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel})

    pathBuffer.EntityData.YListKeys = []string {}

    return &(pathBuffer.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts
// L2VPN EVPN IMCAST table
type Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN IMCAST table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast.
    InclusiveMulticast []*Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
}

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetEntityData() *types.CommonEntityData {
    inclusiveMulticasts.EntityData.YFilter = inclusiveMulticasts.YFilter
    inclusiveMulticasts.EntityData.YangName = "inclusive-multicasts"
    inclusiveMulticasts.EntityData.BundleName = "cisco_ios_xr"
    inclusiveMulticasts.EntityData.ParentYangName = "evi-children"
    inclusiveMulticasts.EntityData.SegmentPath = "inclusive-multicasts"
    inclusiveMulticasts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inclusiveMulticasts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inclusiveMulticasts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inclusiveMulticasts.EntityData.Children = types.NewOrderedMap()
    inclusiveMulticasts.EntityData.Children.Append("inclusive-multicast", types.YChild{"InclusiveMulticast", nil})
    for i := range inclusiveMulticasts.InclusiveMulticast {
        inclusiveMulticasts.EntityData.Children.Append(types.GetSegmentPath(inclusiveMulticasts.InclusiveMulticast[i]), types.YChild{"InclusiveMulticast", inclusiveMulticasts.InclusiveMulticast[i]})
    }
    inclusiveMulticasts.EntityData.Leafs = types.NewOrderedMap()

    inclusiveMulticasts.EntityData.YListKeys = []string {}

    return &(inclusiveMulticasts.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
// L2VPN EVPN IMCAST table
type Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // Originating IP. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIp interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Originating IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIpXr interface{}

    // Tunnel Endpoint ID. The type is interface{} with range: 0..4294967295.
    TunnelEndpointId interface{}

    // IP of nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Local entry. The type is bool.
    IsLocalEntry interface{}

    // Proxy entry. The type is bool.
    IsProxyEntry interface{}

    // SR-TE Policy. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SrtePolicy interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance
}

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetEntityData() *types.CommonEntityData {
    inclusiveMulticast.EntityData.YFilter = inclusiveMulticast.YFilter
    inclusiveMulticast.EntityData.YangName = "inclusive-multicast"
    inclusiveMulticast.EntityData.BundleName = "cisco_ios_xr"
    inclusiveMulticast.EntityData.ParentYangName = "inclusive-multicasts"
    inclusiveMulticast.EntityData.SegmentPath = "inclusive-multicast"
    inclusiveMulticast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inclusiveMulticast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inclusiveMulticast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inclusiveMulticast.EntityData.Children = types.NewOrderedMap()
    inclusiveMulticast.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &inclusiveMulticast.EvpnInstance})
    inclusiveMulticast.EntityData.Leafs = types.NewOrderedMap()
    inclusiveMulticast.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", inclusiveMulticast.Evi})
    inclusiveMulticast.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", inclusiveMulticast.Encapsulation})
    inclusiveMulticast.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", inclusiveMulticast.EthernetTag})
    inclusiveMulticast.EntityData.Leafs.Append("originating-ip", types.YLeaf{"OriginatingIp", inclusiveMulticast.OriginatingIp})
    inclusiveMulticast.EntityData.Leafs.Append("ethernet-tag-xr", types.YLeaf{"EthernetTagXr", inclusiveMulticast.EthernetTagXr})
    inclusiveMulticast.EntityData.Leafs.Append("originating-ip-xr", types.YLeaf{"OriginatingIpXr", inclusiveMulticast.OriginatingIpXr})
    inclusiveMulticast.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", inclusiveMulticast.TunnelEndpointId})
    inclusiveMulticast.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", inclusiveMulticast.NextHop})
    inclusiveMulticast.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", inclusiveMulticast.OutputLabel})
    inclusiveMulticast.EntityData.Leafs.Append("is-local-entry", types.YLeaf{"IsLocalEntry", inclusiveMulticast.IsLocalEntry})
    inclusiveMulticast.EntityData.Leafs.Append("is-proxy-entry", types.YLeaf{"IsProxyEntry", inclusiveMulticast.IsProxyEntry})
    inclusiveMulticast.EntityData.Leafs.Append("srte-policy", types.YLeaf{"SrtePolicy", inclusiveMulticast.SrtePolicy})

    inclusiveMulticast.EntityData.YListKeys = []string {}

    return &(inclusiveMulticast.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance
// EVPN Instance summary information
type Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "inclusive-multicast"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets
// L2VPN EVPN EVI RT Child Table
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI RT Table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget.
    RouteTarget []*Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget
}

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "evi-children"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = types.NewOrderedMap()
    routeTargets.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children.Append(types.GetSegmentPath(routeTargets.RouteTarget[i]), types.YChild{"RouteTarget", routeTargets.RouteTarget[i]})
    }
    routeTargets.EntityData.Leafs = types.NewOrderedMap()

    routeTargets.EntityData.YListKeys = []string {}

    return &(routeTargets.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget
// L2VPN EVPN EVI RT Table
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Role of the route target. The type is BgpRouteTargetRole.
    Role interface{}

    // Format of the route target. The type is BgpRouteTargetFormat.
    Format interface{}

    // Two or Four byte AS Number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // RT AS Index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // RT IP Index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}

    // RT IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // RT Role. The type is L2vpnAdRtRole.
    RouteTargetRole interface{}

    // RT Stitching. The type is bool.
    RouteTargetStitching interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance

    // Route Target.
    RouteTarget Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target"
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &routeTarget.EvpnInstance})
    routeTarget.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", &routeTarget.RouteTarget})
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", routeTarget.Evi})
    routeTarget.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", routeTarget.Encapsulation})
    routeTarget.EntityData.Leafs.Append("role", types.YLeaf{"Role", routeTarget.Role})
    routeTarget.EntityData.Leafs.Append("format", types.YLeaf{"Format", routeTarget.Format})
    routeTarget.EntityData.Leafs.Append("as", types.YLeaf{"As", routeTarget.As})
    routeTarget.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", routeTarget.AsIndex})
    routeTarget.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", routeTarget.AddrIndex})
    routeTarget.EntityData.Leafs.Append("address", types.YLeaf{"Address", routeTarget.Address})
    routeTarget.EntityData.Leafs.Append("route-target-role", types.YLeaf{"RouteTargetRole", routeTarget.RouteTargetRole})
    routeTarget.EntityData.Leafs.Append("route-target-stitching", types.YLeaf{"RouteTargetStitching", routeTarget.RouteTargetStitching})

    routeTarget.EntityData.YListKeys = []string {}

    return &(routeTarget.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance
// EVPN Instance summary information
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "route-target"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
// Route Target
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs

    // v4 addr.
    V4Addr Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr

    // es import.
    EsImport Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-target"
    routeTarget.EntityData.SegmentPath = "route-target"
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &routeTarget.TwoByteAs})
    routeTarget.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &routeTarget.FourByteAs})
    routeTarget.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &routeTarget.V4Addr})
    routeTarget.EntityData.Children.Append("es-import", types.YChild{"EsImport", &routeTarget.EsImport})
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("rt", types.YLeaf{"Rt", routeTarget.Rt})

    routeTarget.EntityData.YListKeys = []string {}

    return &(routeTarget.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "route-target"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "route-target"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "route-target"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport
// es import
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "route-target"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = types.NewOrderedMap()
    esImport.EntityData.Leafs = types.NewOrderedMap()
    esImport.EntityData.Leafs.Append("high-bytes", types.YLeaf{"HighBytes", esImport.HighBytes})
    esImport.EntityData.Leafs.Append("low-bytes", types.YLeaf{"LowBytes", esImport.LowBytes})

    esImport.EntityData.YListKeys = []string {}

    return &(esImport.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs
// L2VPN EVPN EVI MAC table
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN MAC table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac.
    Mac []*Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac
}

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetEntityData() *types.CommonEntityData {
    macs.EntityData.YFilter = macs.YFilter
    macs.EntityData.YangName = "macs"
    macs.EntityData.BundleName = "cisco_ios_xr"
    macs.EntityData.ParentYangName = "evi-children"
    macs.EntityData.SegmentPath = "macs"
    macs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macs.EntityData.Children = types.NewOrderedMap()
    macs.EntityData.Children.Append("mac", types.YChild{"Mac", nil})
    for i := range macs.Mac {
        macs.EntityData.Children.Append(types.GetSegmentPath(macs.Mac[i]), types.YChild{"Mac", macs.Mac[i]})
    }
    macs.EntityData.Leafs = types.NewOrderedMap()

    macs.EntityData.YListKeys = []string {}

    return &(macs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac
// L2VPN EVPN MAC table
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // IP Address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddressXr interface{}

    // IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddressXr interface{}

    // Associated local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Number of items in path list buffer. The type is interface{} with range:
    // 0..4294967295.
    NumPaths interface{}

    // Indication of MAC being locally generated. The type is bool.
    IsLocalMac interface{}

    // Proxy entry. The type is bool.
    IsProxyEntry interface{}

    // Indication of MAC being remotely generated. The type is bool.
    IsRemoteMac interface{}

    // SOO nexthop (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SooNexthop interface{}

    // IP nexthop address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpnhAddress interface{}

    // ESI port key. The type is interface{} with range: 0..65535.
    EsiPortKey interface{}

    // Encap type of local MAC. The type is interface{} with range: 0..255.
    LocalEncapType interface{}

    // Encap type of remote MAC. The type is interface{} with range: 0..255.
    RemoteEncapType interface{}

    // Port the MAC was learned on. The type is string.
    LearnedBridgePortName interface{}

    // local seq id. The type is interface{} with range: 0..4294967295.
    LocalSeqId interface{}

    // remote seq id. The type is interface{} with range: 0..4294967295.
    RemoteSeqId interface{}

    // local l3 label. The type is interface{} with range: 0..4294967295.
    LocalL3Label interface{}

    // Router MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    RouterMacAddress interface{}

    // Number of flushes requested . The type is interface{} with range: 0..65535.
    MacFlushRequested interface{}

    // Number of flushes received . The type is interface{} with range: 0..65535.
    MacFlushReceived interface{}

    // MPLS Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Internal Label has resolved per-ES EAD and per-EVI EAD or MAC routes. The
    // type is bool.
    Resolved interface{}

    // Indication if Local MAC is statically configured. The type is bool.
    LocalIsStatic interface{}

    // Indication if Remote MAC is statically configured. The type is bool.
    RemoteIsStatic interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_EvpnInstance

    // Local Ethernet Segment id. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier.
    LocalEthernetSegmentIdentifier []*Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier

    // Remote Ethernet Segment id. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier.
    RemoteEthernetSegmentIdentifier []*Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer.
    PathBuffer []*Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer
}

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "macs"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &mac.EvpnInstance})
    mac.EntityData.Children.Append("local-ethernet-segment-identifier", types.YChild{"LocalEthernetSegmentIdentifier", nil})
    for i := range mac.LocalEthernetSegmentIdentifier {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.LocalEthernetSegmentIdentifier[i]), types.YChild{"LocalEthernetSegmentIdentifier", mac.LocalEthernetSegmentIdentifier[i]})
    }
    mac.EntityData.Children.Append("remote-ethernet-segment-identifier", types.YChild{"RemoteEthernetSegmentIdentifier", nil})
    for i := range mac.RemoteEthernetSegmentIdentifier {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.RemoteEthernetSegmentIdentifier[i]), types.YChild{"RemoteEthernetSegmentIdentifier", mac.RemoteEthernetSegmentIdentifier[i]})
    }
    mac.EntityData.Children.Append("path-buffer", types.YChild{"PathBuffer", nil})
    for i := range mac.PathBuffer {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.PathBuffer[i]), types.YChild{"PathBuffer", mac.PathBuffer[i]})
    }
    mac.EntityData.Leafs = types.NewOrderedMap()
    mac.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", mac.Evi})
    mac.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", mac.Encapsulation})
    mac.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", mac.EthernetTag})
    mac.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", mac.MacAddress})
    mac.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", mac.IpAddress})
    mac.EntityData.Leafs.Append("ethernet-tag-xr", types.YLeaf{"EthernetTagXr", mac.EthernetTagXr})
    mac.EntityData.Leafs.Append("mac-address-xr", types.YLeaf{"MacAddressXr", mac.MacAddressXr})
    mac.EntityData.Leafs.Append("ip-address-xr", types.YLeaf{"IpAddressXr", mac.IpAddressXr})
    mac.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", mac.LocalLabel})
    mac.EntityData.Leafs.Append("num-paths", types.YLeaf{"NumPaths", mac.NumPaths})
    mac.EntityData.Leafs.Append("is-local-mac", types.YLeaf{"IsLocalMac", mac.IsLocalMac})
    mac.EntityData.Leafs.Append("is-proxy-entry", types.YLeaf{"IsProxyEntry", mac.IsProxyEntry})
    mac.EntityData.Leafs.Append("is-remote-mac", types.YLeaf{"IsRemoteMac", mac.IsRemoteMac})
    mac.EntityData.Leafs.Append("soo-nexthop", types.YLeaf{"SooNexthop", mac.SooNexthop})
    mac.EntityData.Leafs.Append("ipnh-address", types.YLeaf{"IpnhAddress", mac.IpnhAddress})
    mac.EntityData.Leafs.Append("esi-port-key", types.YLeaf{"EsiPortKey", mac.EsiPortKey})
    mac.EntityData.Leafs.Append("local-encap-type", types.YLeaf{"LocalEncapType", mac.LocalEncapType})
    mac.EntityData.Leafs.Append("remote-encap-type", types.YLeaf{"RemoteEncapType", mac.RemoteEncapType})
    mac.EntityData.Leafs.Append("learned-bridge-port-name", types.YLeaf{"LearnedBridgePortName", mac.LearnedBridgePortName})
    mac.EntityData.Leafs.Append("local-seq-id", types.YLeaf{"LocalSeqId", mac.LocalSeqId})
    mac.EntityData.Leafs.Append("remote-seq-id", types.YLeaf{"RemoteSeqId", mac.RemoteSeqId})
    mac.EntityData.Leafs.Append("local-l3-label", types.YLeaf{"LocalL3Label", mac.LocalL3Label})
    mac.EntityData.Leafs.Append("router-mac-address", types.YLeaf{"RouterMacAddress", mac.RouterMacAddress})
    mac.EntityData.Leafs.Append("mac-flush-requested", types.YLeaf{"MacFlushRequested", mac.MacFlushRequested})
    mac.EntityData.Leafs.Append("mac-flush-received", types.YLeaf{"MacFlushReceived", mac.MacFlushReceived})
    mac.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", mac.InternalLabel})
    mac.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", mac.Resolved})
    mac.EntityData.Leafs.Append("local-is-static", types.YLeaf{"LocalIsStatic", mac.LocalIsStatic})
    mac.EntityData.Leafs.Append("remote-is-static", types.YLeaf{"RemoteIsStatic", mac.RemoteIsStatic})

    mac.EntityData.YListKeys = []string {}

    return &(mac.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_EvpnInstance
// EVPN Instance summary information
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "mac"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier
// Local Ethernet Segment id
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    localEthernetSegmentIdentifier.EntityData.YFilter = localEthernetSegmentIdentifier.YFilter
    localEthernetSegmentIdentifier.EntityData.YangName = "local-ethernet-segment-identifier"
    localEthernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    localEthernetSegmentIdentifier.EntityData.ParentYangName = "mac"
    localEthernetSegmentIdentifier.EntityData.SegmentPath = "local-ethernet-segment-identifier"
    localEthernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localEthernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localEthernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localEthernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    localEthernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    localEthernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", localEthernetSegmentIdentifier.Entry})

    localEthernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(localEthernetSegmentIdentifier.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier
// Remote Ethernet Segment id
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    remoteEthernetSegmentIdentifier.EntityData.YFilter = remoteEthernetSegmentIdentifier.YFilter
    remoteEthernetSegmentIdentifier.EntityData.YangName = "remote-ethernet-segment-identifier"
    remoteEthernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    remoteEthernetSegmentIdentifier.EntityData.ParentYangName = "mac"
    remoteEthernetSegmentIdentifier.EntityData.SegmentPath = "remote-ethernet-segment-identifier"
    remoteEthernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteEthernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteEthernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteEthernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    remoteEthernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    remoteEthernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", remoteEthernetSegmentIdentifier.Entry})

    remoteEthernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(remoteEthernetSegmentIdentifier.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer
// Path List Buffer
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetEntityData() *types.CommonEntityData {
    pathBuffer.EntityData.YFilter = pathBuffer.YFilter
    pathBuffer.EntityData.YangName = "path-buffer"
    pathBuffer.EntityData.BundleName = "cisco_ios_xr"
    pathBuffer.EntityData.ParentYangName = "mac"
    pathBuffer.EntityData.SegmentPath = "path-buffer"
    pathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathBuffer.EntityData.Children = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", pathBuffer.TunnelEndpointId})
    pathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", pathBuffer.NextHop})
    pathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", pathBuffer.OutputLabel})
    pathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel})

    pathBuffer.EntityData.YListKeys = []string {}

    return &(pathBuffer.EntityData)
}

// Evpn_Nodes_Node_Teps
// L2VPN EVPN TEP Table
type Evpn_Nodes_Node_Teps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN TEP Entry. The type is slice of Evpn_Nodes_Node_Teps_Tep.
    Tep []*Evpn_Nodes_Node_Teps_Tep
}

func (teps *Evpn_Nodes_Node_Teps) GetEntityData() *types.CommonEntityData {
    teps.EntityData.YFilter = teps.YFilter
    teps.EntityData.YangName = "teps"
    teps.EntityData.BundleName = "cisco_ios_xr"
    teps.EntityData.ParentYangName = "node"
    teps.EntityData.SegmentPath = "teps"
    teps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teps.EntityData.Children = types.NewOrderedMap()
    teps.EntityData.Children.Append("tep", types.YChild{"Tep", nil})
    for i := range teps.Tep {
        teps.EntityData.Children.Append(types.GetSegmentPath(teps.Tep[i]), types.YChild{"Tep", teps.Tep[i]})
    }
    teps.EntityData.Leafs = types.NewOrderedMap()

    teps.EntityData.YListKeys = []string {}

    return &(teps.EntityData)
}

// Evpn_Nodes_Node_Teps_Tep
// L2VPN EVPN TEP Entry
type Evpn_Nodes_Node_Teps_Tep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TEP id. The type is interface{} with range:
    // 0..4294967295.
    TepId interface{}

    // Tunnel Endpoint id. The type is interface{} with range: 0..4294967295.
    TunnelEndpointId interface{}

    // EVPN Tunnel Endpoint Type. The type is interface{} with range: 0..255.
    Type interface{}

    // in-use counter. The type is interface{} with range: 0..4294967295.
    UseCount interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Table Id in RIB. The type is interface{} with range: 0..4294967295.
    VrfTableId interface{}

    // UDP port. The type is interface{} with range: 0..65535.
    UdpPort interface{}

    // Local TEP Information.
    LocalInfo Evpn_Nodes_Node_Teps_Tep_LocalInfo

    // Remote TEP Information.
    RemoteInfo Evpn_Nodes_Node_Teps_Tep_RemoteInfo
}

func (tep *Evpn_Nodes_Node_Teps_Tep) GetEntityData() *types.CommonEntityData {
    tep.EntityData.YFilter = tep.YFilter
    tep.EntityData.YangName = "tep"
    tep.EntityData.BundleName = "cisco_ios_xr"
    tep.EntityData.ParentYangName = "teps"
    tep.EntityData.SegmentPath = "tep" + types.AddKeyToken(tep.TepId, "tep-id")
    tep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tep.EntityData.Children = types.NewOrderedMap()
    tep.EntityData.Children.Append("local-info", types.YChild{"LocalInfo", &tep.LocalInfo})
    tep.EntityData.Children.Append("remote-info", types.YChild{"RemoteInfo", &tep.RemoteInfo})
    tep.EntityData.Leafs = types.NewOrderedMap()
    tep.EntityData.Leafs.Append("tep-id", types.YLeaf{"TepId", tep.TepId})
    tep.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", tep.TunnelEndpointId})
    tep.EntityData.Leafs.Append("type", types.YLeaf{"Type", tep.Type})
    tep.EntityData.Leafs.Append("use-count", types.YLeaf{"UseCount", tep.UseCount})
    tep.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", tep.VrfName})
    tep.EntityData.Leafs.Append("vrf-table-id", types.YLeaf{"VrfTableId", tep.VrfTableId})
    tep.EntityData.Leafs.Append("udp-port", types.YLeaf{"UdpPort", tep.UdpPort})

    tep.EntityData.YListKeys = []string {"TepId"}

    return &(tep.EntityData)
}

// Evpn_Nodes_Node_Teps_Tep_LocalInfo
// Local TEP Information
type Evpn_Nodes_Node_Teps_Tep_LocalInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Tunnel encapsulation. The type is interface{} with range: 0..255.
    Encapsulation interface{}

    // IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}
}

func (localInfo *Evpn_Nodes_Node_Teps_Tep_LocalInfo) GetEntityData() *types.CommonEntityData {
    localInfo.EntityData.YFilter = localInfo.YFilter
    localInfo.EntityData.YangName = "local-info"
    localInfo.EntityData.BundleName = "cisco_ios_xr"
    localInfo.EntityData.ParentYangName = "tep"
    localInfo.EntityData.SegmentPath = "local-info"
    localInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localInfo.EntityData.Children = types.NewOrderedMap()
    localInfo.EntityData.Leafs = types.NewOrderedMap()
    localInfo.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", localInfo.EthernetVpnId})
    localInfo.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", localInfo.Encapsulation})
    localInfo.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", localInfo.Ip})

    localInfo.EntityData.YListKeys = []string {}

    return &(localInfo.EntityData)
}

// Evpn_Nodes_Node_Teps_Tep_RemoteInfo
// Remote TEP Information
type Evpn_Nodes_Node_Teps_Tep_RemoteInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Tunnel encapsulation. The type is interface{} with range: 0..255.
    Encapsulation interface{}

    // IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}
}

func (remoteInfo *Evpn_Nodes_Node_Teps_Tep_RemoteInfo) GetEntityData() *types.CommonEntityData {
    remoteInfo.EntityData.YFilter = remoteInfo.YFilter
    remoteInfo.EntityData.YangName = "remote-info"
    remoteInfo.EntityData.BundleName = "cisco_ios_xr"
    remoteInfo.EntityData.ParentYangName = "tep"
    remoteInfo.EntityData.SegmentPath = "remote-info"
    remoteInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteInfo.EntityData.Children = types.NewOrderedMap()
    remoteInfo.EntityData.Leafs = types.NewOrderedMap()
    remoteInfo.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", remoteInfo.EthernetVpnId})
    remoteInfo.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", remoteInfo.Encapsulation})
    remoteInfo.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", remoteInfo.Ip})

    remoteInfo.EntityData.YListKeys = []string {}

    return &(remoteInfo.EntityData)
}

// Evpn_Nodes_Node_InternalLabels
// EVPN Internal Label Table
type Evpn_Nodes_Node_InternalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN Internal Label. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel.
    InternalLabel []*Evpn_Nodes_Node_InternalLabels_InternalLabel
}

func (internalLabels *Evpn_Nodes_Node_InternalLabels) GetEntityData() *types.CommonEntityData {
    internalLabels.EntityData.YFilter = internalLabels.YFilter
    internalLabels.EntityData.YangName = "internal-labels"
    internalLabels.EntityData.BundleName = "cisco_ios_xr"
    internalLabels.EntityData.ParentYangName = "node"
    internalLabels.EntityData.SegmentPath = "internal-labels"
    internalLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalLabels.EntityData.Children = types.NewOrderedMap()
    internalLabels.EntityData.Children.Append("internal-label", types.YChild{"InternalLabel", nil})
    for i := range internalLabels.InternalLabel {
        internalLabels.EntityData.Children.Append(types.GetSegmentPath(internalLabels.InternalLabel[i]), types.YChild{"InternalLabel", internalLabels.InternalLabel[i]})
    }
    internalLabels.EntityData.Leafs = types.NewOrderedMap()

    internalLabels.EntityData.YListKeys = []string {}

    return &(internalLabels.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel
// L2VPN EVPN Internal Label
type Evpn_Nodes_Node_InternalLabels_InternalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // Ethernet Segment id. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Esi interface{}

    // Label Tag. The type is interface{} with range: 0..4294967295.
    Tag interface{}

    // MPLS Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Number of items in the MAC path buffer. The type is interface{} with range:
    // 0..4294967295.
    MacNumPaths interface{}

    // Number of items in the ead path buffer. The type is interface{} with range:
    // 0..4294967295.
    EadNumPaths interface{}

    // Number of items in the evi path buffer. The type is interface{} with range:
    // 0..4294967295.
    EviNumPaths interface{}

    // Number of items in the sum path buffer. The type is interface{} with range:
    // 0..4294967295.
    SumNumPaths interface{}

    // Number of items in the sum path buffer that are Active Paths. The type is
    // interface{} with range: 0..4294967295.
    SumNumActivePaths interface{}

    // Internal Label has resolved per-ES EAD and per-EVI EAD or MAC routes. The
    // type is bool.
    Resolved interface{}

    // ECMP Disable Per EVI Resolution. The type is bool.
    EcmpDisable interface{}

    // Single-active redundancy configured at remote ES. The type is bool.
    RedundancySingleActive interface{}

    // Single-flow-active redundancy at remote ES (MST-AG). The type is bool.
    RedundancySingleFlowActive interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Nodes_Node_InternalLabels_InternalLabel_EvpnInstance

    // MAC Path list buffer. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel_MacPathBuffer.
    MacPathBuffer []*Evpn_Nodes_Node_InternalLabels_InternalLabel_MacPathBuffer

    // EAD/ES Path list buffer. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel_EadPathBuffer.
    EadPathBuffer []*Evpn_Nodes_Node_InternalLabels_InternalLabel_EadPathBuffer

    // EAD/EVI Path list buffer. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel_EviPathBuffer.
    EviPathBuffer []*Evpn_Nodes_Node_InternalLabels_InternalLabel_EviPathBuffer

    // Summary Path list buffer. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel_SummaryPathBuffer.
    SummaryPathBuffer []*Evpn_Nodes_Node_InternalLabels_InternalLabel_SummaryPathBuffer
}

func (internalLabel *Evpn_Nodes_Node_InternalLabels_InternalLabel) GetEntityData() *types.CommonEntityData {
    internalLabel.EntityData.YFilter = internalLabel.YFilter
    internalLabel.EntityData.YangName = "internal-label"
    internalLabel.EntityData.BundleName = "cisco_ios_xr"
    internalLabel.EntityData.ParentYangName = "internal-labels"
    internalLabel.EntityData.SegmentPath = "internal-label"
    internalLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalLabel.EntityData.Children = types.NewOrderedMap()
    internalLabel.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &internalLabel.EvpnInstance})
    internalLabel.EntityData.Children.Append("mac-path-buffer", types.YChild{"MacPathBuffer", nil})
    for i := range internalLabel.MacPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.MacPathBuffer[i]), types.YChild{"MacPathBuffer", internalLabel.MacPathBuffer[i]})
    }
    internalLabel.EntityData.Children.Append("ead-path-buffer", types.YChild{"EadPathBuffer", nil})
    for i := range internalLabel.EadPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.EadPathBuffer[i]), types.YChild{"EadPathBuffer", internalLabel.EadPathBuffer[i]})
    }
    internalLabel.EntityData.Children.Append("evi-path-buffer", types.YChild{"EviPathBuffer", nil})
    for i := range internalLabel.EviPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.EviPathBuffer[i]), types.YChild{"EviPathBuffer", internalLabel.EviPathBuffer[i]})
    }
    internalLabel.EntityData.Children.Append("summary-path-buffer", types.YChild{"SummaryPathBuffer", nil})
    for i := range internalLabel.SummaryPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.SummaryPathBuffer[i]), types.YChild{"SummaryPathBuffer", internalLabel.SummaryPathBuffer[i]})
    }
    internalLabel.EntityData.Leafs = types.NewOrderedMap()
    internalLabel.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", internalLabel.Evi})
    internalLabel.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", internalLabel.Encapsulation})
    internalLabel.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", internalLabel.Esi1})
    internalLabel.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", internalLabel.Esi2})
    internalLabel.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", internalLabel.Esi3})
    internalLabel.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", internalLabel.Esi4})
    internalLabel.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", internalLabel.Esi5})
    internalLabel.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", internalLabel.EthernetTag})
    internalLabel.EntityData.Leafs.Append("esi", types.YLeaf{"Esi", internalLabel.Esi})
    internalLabel.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", internalLabel.Tag})
    internalLabel.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", internalLabel.InternalLabel})
    internalLabel.EntityData.Leafs.Append("mac-num-paths", types.YLeaf{"MacNumPaths", internalLabel.MacNumPaths})
    internalLabel.EntityData.Leafs.Append("ead-num-paths", types.YLeaf{"EadNumPaths", internalLabel.EadNumPaths})
    internalLabel.EntityData.Leafs.Append("evi-num-paths", types.YLeaf{"EviNumPaths", internalLabel.EviNumPaths})
    internalLabel.EntityData.Leafs.Append("sum-num-paths", types.YLeaf{"SumNumPaths", internalLabel.SumNumPaths})
    internalLabel.EntityData.Leafs.Append("sum-num-active-paths", types.YLeaf{"SumNumActivePaths", internalLabel.SumNumActivePaths})
    internalLabel.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", internalLabel.Resolved})
    internalLabel.EntityData.Leafs.Append("ecmp-disable", types.YLeaf{"EcmpDisable", internalLabel.EcmpDisable})
    internalLabel.EntityData.Leafs.Append("redundancy-single-active", types.YLeaf{"RedundancySingleActive", internalLabel.RedundancySingleActive})
    internalLabel.EntityData.Leafs.Append("redundancy-single-flow-active", types.YLeaf{"RedundancySingleFlowActive", internalLabel.RedundancySingleFlowActive})

    internalLabel.EntityData.YListKeys = []string {}

    return &(internalLabel.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel_EvpnInstance
// EVPN Instance summary information
type Evpn_Nodes_Node_InternalLabels_InternalLabel_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Nodes_Node_InternalLabels_InternalLabel_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "internal-label"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel_MacPathBuffer
// MAC Path list buffer
type Evpn_Nodes_Node_InternalLabels_InternalLabel_MacPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (macPathBuffer *Evpn_Nodes_Node_InternalLabels_InternalLabel_MacPathBuffer) GetEntityData() *types.CommonEntityData {
    macPathBuffer.EntityData.YFilter = macPathBuffer.YFilter
    macPathBuffer.EntityData.YangName = "mac-path-buffer"
    macPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    macPathBuffer.EntityData.ParentYangName = "internal-label"
    macPathBuffer.EntityData.SegmentPath = "mac-path-buffer"
    macPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macPathBuffer.EntityData.Children = types.NewOrderedMap()
    macPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    macPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", macPathBuffer.TunnelEndpointId})
    macPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", macPathBuffer.NextHop})
    macPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", macPathBuffer.OutputLabel})
    macPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", macPathBuffer.SrteTunnel})

    macPathBuffer.EntityData.YListKeys = []string {}

    return &(macPathBuffer.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel_EadPathBuffer
// EAD/ES Path list buffer
type Evpn_Nodes_Node_InternalLabels_InternalLabel_EadPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (eadPathBuffer *Evpn_Nodes_Node_InternalLabels_InternalLabel_EadPathBuffer) GetEntityData() *types.CommonEntityData {
    eadPathBuffer.EntityData.YFilter = eadPathBuffer.YFilter
    eadPathBuffer.EntityData.YangName = "ead-path-buffer"
    eadPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    eadPathBuffer.EntityData.ParentYangName = "internal-label"
    eadPathBuffer.EntityData.SegmentPath = "ead-path-buffer"
    eadPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eadPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eadPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eadPathBuffer.EntityData.Children = types.NewOrderedMap()
    eadPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    eadPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", eadPathBuffer.TunnelEndpointId})
    eadPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", eadPathBuffer.NextHop})
    eadPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", eadPathBuffer.OutputLabel})
    eadPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", eadPathBuffer.SrteTunnel})

    eadPathBuffer.EntityData.YListKeys = []string {}

    return &(eadPathBuffer.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel_EviPathBuffer
// EAD/EVI Path list buffer
type Evpn_Nodes_Node_InternalLabels_InternalLabel_EviPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (eviPathBuffer *Evpn_Nodes_Node_InternalLabels_InternalLabel_EviPathBuffer) GetEntityData() *types.CommonEntityData {
    eviPathBuffer.EntityData.YFilter = eviPathBuffer.YFilter
    eviPathBuffer.EntityData.YangName = "evi-path-buffer"
    eviPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    eviPathBuffer.EntityData.ParentYangName = "internal-label"
    eviPathBuffer.EntityData.SegmentPath = "evi-path-buffer"
    eviPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviPathBuffer.EntityData.Children = types.NewOrderedMap()
    eviPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    eviPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", eviPathBuffer.TunnelEndpointId})
    eviPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", eviPathBuffer.NextHop})
    eviPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", eviPathBuffer.OutputLabel})
    eviPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", eviPathBuffer.SrteTunnel})

    eviPathBuffer.EntityData.YListKeys = []string {}

    return &(eviPathBuffer.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel_SummaryPathBuffer
// Summary Path list buffer
type Evpn_Nodes_Node_InternalLabels_InternalLabel_SummaryPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (summaryPathBuffer *Evpn_Nodes_Node_InternalLabels_InternalLabel_SummaryPathBuffer) GetEntityData() *types.CommonEntityData {
    summaryPathBuffer.EntityData.YFilter = summaryPathBuffer.YFilter
    summaryPathBuffer.EntityData.YangName = "summary-path-buffer"
    summaryPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    summaryPathBuffer.EntityData.ParentYangName = "internal-label"
    summaryPathBuffer.EntityData.SegmentPath = "summary-path-buffer"
    summaryPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPathBuffer.EntityData.Children = types.NewOrderedMap()
    summaryPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    summaryPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", summaryPathBuffer.TunnelEndpointId})
    summaryPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", summaryPathBuffer.NextHop})
    summaryPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", summaryPathBuffer.OutputLabel})
    summaryPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", summaryPathBuffer.SrteTunnel})

    summaryPathBuffer.EntityData.YListKeys = []string {}

    return &(summaryPathBuffer.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments
// EVPN Ethernet-Segment Table
type Evpn_Nodes_Node_EthernetSegments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet-Segment Entry. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment.
    EthernetSegment []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment
}

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetEntityData() *types.CommonEntityData {
    ethernetSegments.EntityData.YFilter = ethernetSegments.YFilter
    ethernetSegments.EntityData.YangName = "ethernet-segments"
    ethernetSegments.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegments.EntityData.ParentYangName = "node"
    ethernetSegments.EntityData.SegmentPath = "ethernet-segments"
    ethernetSegments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegments.EntityData.Children = types.NewOrderedMap()
    ethernetSegments.EntityData.Children.Append("ethernet-segment", types.YChild{"EthernetSegment", nil})
    for i := range ethernetSegments.EthernetSegment {
        ethernetSegments.EntityData.Children.Append(types.GetSegmentPath(ethernetSegments.EthernetSegment[i]), types.YChild{"EthernetSegment", ethernetSegments.EthernetSegment[i]})
    }
    ethernetSegments.EntityData.Leafs = types.NewOrderedMap()

    ethernetSegments.EntityData.YListKeys = []string {}

    return &(ethernetSegments.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment
// EVPN Ethernet-Segment Entry
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // ESI Type. The type is L2vpnEvpnEsi.
    EsiType interface{}

    // ESI System Identifier. The type is string.
    EsiSystemIdentifier interface{}

    // ESI Port Key. The type is interface{} with range: 0..4294967295.
    EsiPortKey interface{}

    // ESI System Priority. The type is interface{} with range: 0..4294967295.
    EsiSystemPriority interface{}

    // Ethernet Segment Name. The type is string.
    EthernetSegmentName interface{}

    // State of the ethernet segment. The type is interface{} with range:
    // 0..4294967295.
    EthernetSegmentState interface{}

    // Main port ifhandle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    IfHandle interface{}

    // Main port redundancy group role. The type is L2vpnRgRole.
    MainPortRole interface{}

    // Main Port MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MainPortMac interface{}

    // Number of PWs in Up state. The type is interface{} with range:
    // 0..4294967295.
    NumUpPWs interface{}

    // ES-Import Route Target. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    RouteTarget interface{}

    // Origin of operational ES-Import RT. The type is L2vpnEvpnRtOrigin.
    RtOrigin interface{}

    // ES BGP Gates. The type is string.
    EsBgpGates interface{}

    // ES L2FIB Gates. The type is string.
    EsL2fibGates interface{}

    // Configured MAC Flushing mode. The type is L2vpnEvpnMfMode.
    MacFlushingModeConfig interface{}

    // Configured load balancing mode. The type is L2vpnEvpnLbMode.
    LoadBalanceModeConfig interface{}

    // Load balancing mode is default. The type is bool.
    LoadBalanceModeIsDefault interface{}

    // Operational load balancing mode. The type is L2vpnEvpnLbMode.
    LoadBalanceModeOper interface{}

    // Ethernet-Segment forced to single home. The type is bool.
    ForceSingleHome interface{}

    // Operational Source MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacOper interface{}

    // Origin of operational source MAC address. The type is L2vpnEvpnSmacSrc.
    SourceMacOrigin interface{}

    // Configured timer for triggering DF election (seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    PeeringTimer interface{}

    // Milliseconds left on DF election timer. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    PeeringTimerLeft interface{}

    // Configured timer for (STP) recovery (seconds). The type is interface{} with
    // range: 0..4294967295. Units are second.
    RecoveryTimer interface{}

    // Milliseconds left on (STP) recovery timer. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RecoveryTimerLeft interface{}

    // Configured timer for delaying DF election (seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    CarvingTimer interface{}

    // Milliseconds left on carving timer. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    CarvingTimerLeft interface{}

    // Service carving mode. The type is L2vpnEvpnScMode.
    ServiceCarvingMode interface{}

    // Input string of Primary services ESI/I-SIDs. The type is string.
    PrimaryServicesInput interface{}

    // Input string of Secondary services ESI/I-SIDs. The type is string.
    SecondaryServicesInput interface{}

    // Count of Forwarders (AC, AC PW, VFI PW). The type is interface{} with
    // range: 0..4294967295.
    ForwarderPorts interface{}

    // Count of Forwarders with permanent service. The type is interface{} with
    // range: 0..4294967295.
    PermanentForwarderPorts interface{}

    // Count of Forwarders with elected service. The type is interface{} with
    // range: 0..4294967295.
    ElectedForwarderPorts interface{}

    // Count of Forwarders with not elected service. The type is interface{} with
    // range: 0..4294967295.
    NotElectedForwarderPorts interface{}

    // Count of forwarders with missing config detected. The type is interface{}
    // with range: 0..4294967295.
    NotConfigForwarderPorts interface{}

    // MP is protected and not under EVPN control. The type is bool.
    MpProtected interface{}

    // Anycast VTEP mode on NVE main-interface. The type is bool.
    NveAnycastVtep interface{}

    // Ingress-Replication is configured on NVE main-interface. The type is bool.
    NveIngressReplication interface{}

    // Local split horizon group label is valid. The type is bool.
    LocalSplitHorizonGroupLabelValid interface{}

    // Local split horizon group label. The type is interface{} with range:
    // 0..4294967295.
    LocalSplitHorizonGroupLabel interface{}

    // Ethernet Segment id. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier

    // List of Primary services ESI/I-SIDs. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService.
    PrimaryService []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService

    // List of Secondary services ESI/I-SIDs. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService.
    SecondaryService []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService

    // Elected ISID service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult.
    ServiceCarvingISidelectedResult []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult

    // Not elected ISID service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult.
    ServiceCarvingIsidNotElectedResult []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult

    // Elected EVI service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult.
    ServiceCarvingEviElectedResult []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult

    // Not elected EVI service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult.
    ServiceCarvingEviNotElectedResult []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult

    // Elected VNI service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult.
    ServiceCarvingVniElectedResult []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult

    // Not elected VNI service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult.
    ServiceCarvingVniNotElectedResult []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult

    // List of nexthop IPv6 addresses. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop.
    NextHop []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop

    // Permanent EVPN VPWS service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult.
    ServiceCarvingVpwsPermanentResult []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult

    // Remote split horizon group labels. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel.
    RemoteSplitHorizonGroupLabel []*Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
}

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetEntityData() *types.CommonEntityData {
    ethernetSegment.EntityData.YFilter = ethernetSegment.YFilter
    ethernetSegment.EntityData.YangName = "ethernet-segment"
    ethernetSegment.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegment.EntityData.ParentYangName = "ethernet-segments"
    ethernetSegment.EntityData.SegmentPath = "ethernet-segment"
    ethernetSegment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegment.EntityData.Children = types.NewOrderedMap()
    ethernetSegment.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range ethernetSegment.EthernetSegmentIdentifier {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", ethernetSegment.EthernetSegmentIdentifier[i]})
    }
    ethernetSegment.EntityData.Children.Append("primary-service", types.YChild{"PrimaryService", nil})
    for i := range ethernetSegment.PrimaryService {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.PrimaryService[i]), types.YChild{"PrimaryService", ethernetSegment.PrimaryService[i]})
    }
    ethernetSegment.EntityData.Children.Append("secondary-service", types.YChild{"SecondaryService", nil})
    for i := range ethernetSegment.SecondaryService {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.SecondaryService[i]), types.YChild{"SecondaryService", ethernetSegment.SecondaryService[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-i-sidelected-result", types.YChild{"ServiceCarvingISidelectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingISidelectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingISidelectedResult[i]), types.YChild{"ServiceCarvingISidelectedResult", ethernetSegment.ServiceCarvingISidelectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-isid-not-elected-result", types.YChild{"ServiceCarvingIsidNotElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingIsidNotElectedResult[i]), types.YChild{"ServiceCarvingIsidNotElectedResult", ethernetSegment.ServiceCarvingIsidNotElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-evi-elected-result", types.YChild{"ServiceCarvingEviElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingEviElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingEviElectedResult[i]), types.YChild{"ServiceCarvingEviElectedResult", ethernetSegment.ServiceCarvingEviElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-evi-not-elected-result", types.YChild{"ServiceCarvingEviNotElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingEviNotElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingEviNotElectedResult[i]), types.YChild{"ServiceCarvingEviNotElectedResult", ethernetSegment.ServiceCarvingEviNotElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-vni-elected-result", types.YChild{"ServiceCarvingVniElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingVniElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingVniElectedResult[i]), types.YChild{"ServiceCarvingVniElectedResult", ethernetSegment.ServiceCarvingVniElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-vni-not-elected-result", types.YChild{"ServiceCarvingVniNotElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingVniNotElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingVniNotElectedResult[i]), types.YChild{"ServiceCarvingVniNotElectedResult", ethernetSegment.ServiceCarvingVniNotElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range ethernetSegment.NextHop {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.NextHop[i]), types.YChild{"NextHop", ethernetSegment.NextHop[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-vpws-permanent-result", types.YChild{"ServiceCarvingVpwsPermanentResult", nil})
    for i := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingVpwsPermanentResult[i]), types.YChild{"ServiceCarvingVpwsPermanentResult", ethernetSegment.ServiceCarvingVpwsPermanentResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("remote-split-horizon-group-label", types.YChild{"RemoteSplitHorizonGroupLabel", nil})
    for i := range ethernetSegment.RemoteSplitHorizonGroupLabel {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.RemoteSplitHorizonGroupLabel[i]), types.YChild{"RemoteSplitHorizonGroupLabel", ethernetSegment.RemoteSplitHorizonGroupLabel[i]})
    }
    ethernetSegment.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegment.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ethernetSegment.InterfaceName})
    ethernetSegment.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", ethernetSegment.Esi1})
    ethernetSegment.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", ethernetSegment.Esi2})
    ethernetSegment.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", ethernetSegment.Esi3})
    ethernetSegment.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", ethernetSegment.Esi4})
    ethernetSegment.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", ethernetSegment.Esi5})
    ethernetSegment.EntityData.Leafs.Append("esi-type", types.YLeaf{"EsiType", ethernetSegment.EsiType})
    ethernetSegment.EntityData.Leafs.Append("esi-system-identifier", types.YLeaf{"EsiSystemIdentifier", ethernetSegment.EsiSystemIdentifier})
    ethernetSegment.EntityData.Leafs.Append("esi-port-key", types.YLeaf{"EsiPortKey", ethernetSegment.EsiPortKey})
    ethernetSegment.EntityData.Leafs.Append("esi-system-priority", types.YLeaf{"EsiSystemPriority", ethernetSegment.EsiSystemPriority})
    ethernetSegment.EntityData.Leafs.Append("ethernet-segment-name", types.YLeaf{"EthernetSegmentName", ethernetSegment.EthernetSegmentName})
    ethernetSegment.EntityData.Leafs.Append("ethernet-segment-state", types.YLeaf{"EthernetSegmentState", ethernetSegment.EthernetSegmentState})
    ethernetSegment.EntityData.Leafs.Append("if-handle", types.YLeaf{"IfHandle", ethernetSegment.IfHandle})
    ethernetSegment.EntityData.Leafs.Append("main-port-role", types.YLeaf{"MainPortRole", ethernetSegment.MainPortRole})
    ethernetSegment.EntityData.Leafs.Append("main-port-mac", types.YLeaf{"MainPortMac", ethernetSegment.MainPortMac})
    ethernetSegment.EntityData.Leafs.Append("num-up-p-ws", types.YLeaf{"NumUpPWs", ethernetSegment.NumUpPWs})
    ethernetSegment.EntityData.Leafs.Append("route-target", types.YLeaf{"RouteTarget", ethernetSegment.RouteTarget})
    ethernetSegment.EntityData.Leafs.Append("rt-origin", types.YLeaf{"RtOrigin", ethernetSegment.RtOrigin})
    ethernetSegment.EntityData.Leafs.Append("es-bgp-gates", types.YLeaf{"EsBgpGates", ethernetSegment.EsBgpGates})
    ethernetSegment.EntityData.Leafs.Append("es-l2fib-gates", types.YLeaf{"EsL2fibGates", ethernetSegment.EsL2fibGates})
    ethernetSegment.EntityData.Leafs.Append("mac-flushing-mode-config", types.YLeaf{"MacFlushingModeConfig", ethernetSegment.MacFlushingModeConfig})
    ethernetSegment.EntityData.Leafs.Append("load-balance-mode-config", types.YLeaf{"LoadBalanceModeConfig", ethernetSegment.LoadBalanceModeConfig})
    ethernetSegment.EntityData.Leafs.Append("load-balance-mode-is-default", types.YLeaf{"LoadBalanceModeIsDefault", ethernetSegment.LoadBalanceModeIsDefault})
    ethernetSegment.EntityData.Leafs.Append("load-balance-mode-oper", types.YLeaf{"LoadBalanceModeOper", ethernetSegment.LoadBalanceModeOper})
    ethernetSegment.EntityData.Leafs.Append("force-single-home", types.YLeaf{"ForceSingleHome", ethernetSegment.ForceSingleHome})
    ethernetSegment.EntityData.Leafs.Append("source-mac-oper", types.YLeaf{"SourceMacOper", ethernetSegment.SourceMacOper})
    ethernetSegment.EntityData.Leafs.Append("source-mac-origin", types.YLeaf{"SourceMacOrigin", ethernetSegment.SourceMacOrigin})
    ethernetSegment.EntityData.Leafs.Append("peering-timer", types.YLeaf{"PeeringTimer", ethernetSegment.PeeringTimer})
    ethernetSegment.EntityData.Leafs.Append("peering-timer-left", types.YLeaf{"PeeringTimerLeft", ethernetSegment.PeeringTimerLeft})
    ethernetSegment.EntityData.Leafs.Append("recovery-timer", types.YLeaf{"RecoveryTimer", ethernetSegment.RecoveryTimer})
    ethernetSegment.EntityData.Leafs.Append("recovery-timer-left", types.YLeaf{"RecoveryTimerLeft", ethernetSegment.RecoveryTimerLeft})
    ethernetSegment.EntityData.Leafs.Append("carving-timer", types.YLeaf{"CarvingTimer", ethernetSegment.CarvingTimer})
    ethernetSegment.EntityData.Leafs.Append("carving-timer-left", types.YLeaf{"CarvingTimerLeft", ethernetSegment.CarvingTimerLeft})
    ethernetSegment.EntityData.Leafs.Append("service-carving-mode", types.YLeaf{"ServiceCarvingMode", ethernetSegment.ServiceCarvingMode})
    ethernetSegment.EntityData.Leafs.Append("primary-services-input", types.YLeaf{"PrimaryServicesInput", ethernetSegment.PrimaryServicesInput})
    ethernetSegment.EntityData.Leafs.Append("secondary-services-input", types.YLeaf{"SecondaryServicesInput", ethernetSegment.SecondaryServicesInput})
    ethernetSegment.EntityData.Leafs.Append("forwarder-ports", types.YLeaf{"ForwarderPorts", ethernetSegment.ForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("permanent-forwarder-ports", types.YLeaf{"PermanentForwarderPorts", ethernetSegment.PermanentForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("elected-forwarder-ports", types.YLeaf{"ElectedForwarderPorts", ethernetSegment.ElectedForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("not-elected-forwarder-ports", types.YLeaf{"NotElectedForwarderPorts", ethernetSegment.NotElectedForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("not-config-forwarder-ports", types.YLeaf{"NotConfigForwarderPorts", ethernetSegment.NotConfigForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("mp-protected", types.YLeaf{"MpProtected", ethernetSegment.MpProtected})
    ethernetSegment.EntityData.Leafs.Append("nve-anycast-vtep", types.YLeaf{"NveAnycastVtep", ethernetSegment.NveAnycastVtep})
    ethernetSegment.EntityData.Leafs.Append("nve-ingress-replication", types.YLeaf{"NveIngressReplication", ethernetSegment.NveIngressReplication})
    ethernetSegment.EntityData.Leafs.Append("local-split-horizon-group-label-valid", types.YLeaf{"LocalSplitHorizonGroupLabelValid", ethernetSegment.LocalSplitHorizonGroupLabelValid})
    ethernetSegment.EntityData.Leafs.Append("local-split-horizon-group-label", types.YLeaf{"LocalSplitHorizonGroupLabel", ethernetSegment.LocalSplitHorizonGroupLabel})

    ethernetSegment.EntityData.YListKeys = []string {}

    return &(ethernetSegment.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "ethernet-segment"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService
// List of Primary services ESI/I-SIDs
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetEntityData() *types.CommonEntityData {
    primaryService.EntityData.YFilter = primaryService.YFilter
    primaryService.EntityData.YangName = "primary-service"
    primaryService.EntityData.BundleName = "cisco_ios_xr"
    primaryService.EntityData.ParentYangName = "ethernet-segment"
    primaryService.EntityData.SegmentPath = "primary-service"
    primaryService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryService.EntityData.Children = types.NewOrderedMap()
    primaryService.EntityData.Leafs = types.NewOrderedMap()
    primaryService.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", primaryService.Entry})

    primaryService.EntityData.YListKeys = []string {}

    return &(primaryService.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService
// List of Secondary services ESI/I-SIDs
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetEntityData() *types.CommonEntityData {
    secondaryService.EntityData.YFilter = secondaryService.YFilter
    secondaryService.EntityData.YangName = "secondary-service"
    secondaryService.EntityData.BundleName = "cisco_ios_xr"
    secondaryService.EntityData.ParentYangName = "ethernet-segment"
    secondaryService.EntityData.SegmentPath = "secondary-service"
    secondaryService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryService.EntityData.Children = types.NewOrderedMap()
    secondaryService.EntityData.Leafs = types.NewOrderedMap()
    secondaryService.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", secondaryService.Entry})

    secondaryService.EntityData.YListKeys = []string {}

    return &(secondaryService.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult
// Elected ISID service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingISidelectedResult.EntityData.YFilter = serviceCarvingISidelectedResult.YFilter
    serviceCarvingISidelectedResult.EntityData.YangName = "service-carving-i-sidelected-result"
    serviceCarvingISidelectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingISidelectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingISidelectedResult.EntityData.SegmentPath = "service-carving-i-sidelected-result"
    serviceCarvingISidelectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingISidelectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingISidelectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingISidelectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingISidelectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingISidelectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingISidelectedResult.Entry})

    serviceCarvingISidelectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingISidelectedResult.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult
// Not elected ISID service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingIsidNotElectedResult.EntityData.YFilter = serviceCarvingIsidNotElectedResult.YFilter
    serviceCarvingIsidNotElectedResult.EntityData.YangName = "service-carving-isid-not-elected-result"
    serviceCarvingIsidNotElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingIsidNotElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingIsidNotElectedResult.EntityData.SegmentPath = "service-carving-isid-not-elected-result"
    serviceCarvingIsidNotElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingIsidNotElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingIsidNotElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingIsidNotElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingIsidNotElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingIsidNotElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingIsidNotElectedResult.Entry})

    serviceCarvingIsidNotElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingIsidNotElectedResult.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult
// Elected EVI service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingEviElectedResult.EntityData.YFilter = serviceCarvingEviElectedResult.YFilter
    serviceCarvingEviElectedResult.EntityData.YangName = "service-carving-evi-elected-result"
    serviceCarvingEviElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingEviElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingEviElectedResult.EntityData.SegmentPath = "service-carving-evi-elected-result"
    serviceCarvingEviElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingEviElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingEviElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingEviElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingEviElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingEviElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingEviElectedResult.Entry})

    serviceCarvingEviElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingEviElectedResult.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult
// Not elected EVI service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingEviNotElectedResult.EntityData.YFilter = serviceCarvingEviNotElectedResult.YFilter
    serviceCarvingEviNotElectedResult.EntityData.YangName = "service-carving-evi-not-elected-result"
    serviceCarvingEviNotElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingEviNotElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingEviNotElectedResult.EntityData.SegmentPath = "service-carving-evi-not-elected-result"
    serviceCarvingEviNotElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingEviNotElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingEviNotElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingEviNotElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingEviNotElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingEviNotElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingEviNotElectedResult.Entry})

    serviceCarvingEviNotElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingEviNotElectedResult.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult
// Elected VNI service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingVniElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingVniElectedResult.EntityData.YFilter = serviceCarvingVniElectedResult.YFilter
    serviceCarvingVniElectedResult.EntityData.YangName = "service-carving-vni-elected-result"
    serviceCarvingVniElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingVniElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingVniElectedResult.EntityData.SegmentPath = "service-carving-vni-elected-result"
    serviceCarvingVniElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingVniElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingVniElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingVniElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingVniElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingVniElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingVniElectedResult.Entry})

    serviceCarvingVniElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingVniElectedResult.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult
// Not elected VNI service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingVniNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingVniNotElectedResult.EntityData.YFilter = serviceCarvingVniNotElectedResult.YFilter
    serviceCarvingVniNotElectedResult.EntityData.YangName = "service-carving-vni-not-elected-result"
    serviceCarvingVniNotElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingVniNotElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingVniNotElectedResult.EntityData.SegmentPath = "service-carving-vni-not-elected-result"
    serviceCarvingVniNotElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingVniNotElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingVniNotElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingVniNotElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingVniNotElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingVniNotElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingVniNotElectedResult.Entry})

    serviceCarvingVniNotElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingVniNotElectedResult.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop
// List of nexthop IPv6 addresses
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // DF Dont Premption. The type is bool.
    DfDontPrempt interface{}

    // DF Election Mode Configured. The type is interface{} with range: 0..255.
    DfType interface{}

    // DF Election Preference Set. The type is interface{} with range: 0..65535.
    DfPref interface{}
}

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "ethernet-segment"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", nextHop.NextHop})
    nextHop.EntityData.Leafs.Append("df-dont-prempt", types.YLeaf{"DfDontPrempt", nextHop.DfDontPrempt})
    nextHop.EntityData.Leafs.Append("df-type", types.YLeaf{"DfType", nextHop.DfType})
    nextHop.EntityData.Leafs.Append("df-pref", types.YLeaf{"DfPref", nextHop.DfPref})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult
// Permanent EVPN VPWS service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPN ID. The type is interface{} with range: 0..4294967295.
    VpnId interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}
}

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingVpwsPermanentResult.EntityData.YFilter = serviceCarvingVpwsPermanentResult.YFilter
    serviceCarvingVpwsPermanentResult.EntityData.YangName = "service-carving-vpws-permanent-result"
    serviceCarvingVpwsPermanentResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingVpwsPermanentResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingVpwsPermanentResult.EntityData.SegmentPath = "service-carving-vpws-permanent-result"
    serviceCarvingVpwsPermanentResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingVpwsPermanentResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingVpwsPermanentResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingVpwsPermanentResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingVpwsPermanentResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingVpwsPermanentResult.EntityData.Leafs.Append("vpn-id", types.YLeaf{"VpnId", serviceCarvingVpwsPermanentResult.VpnId})
    serviceCarvingVpwsPermanentResult.EntityData.Leafs.Append("type", types.YLeaf{"Type", serviceCarvingVpwsPermanentResult.Type})
    serviceCarvingVpwsPermanentResult.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", serviceCarvingVpwsPermanentResult.EthernetTag})

    serviceCarvingVpwsPermanentResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingVpwsPermanentResult.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Split horizon label associated with next-hop address. The type is
    // interface{} with range: 0..4294967295.
    Label interface{}
}

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetEntityData() *types.CommonEntityData {
    remoteSplitHorizonGroupLabel.EntityData.YFilter = remoteSplitHorizonGroupLabel.YFilter
    remoteSplitHorizonGroupLabel.EntityData.YangName = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.BundleName = "cisco_ios_xr"
    remoteSplitHorizonGroupLabel.EntityData.ParentYangName = "ethernet-segment"
    remoteSplitHorizonGroupLabel.EntityData.SegmentPath = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteSplitHorizonGroupLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteSplitHorizonGroupLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteSplitHorizonGroupLabel.EntityData.Children = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", remoteSplitHorizonGroupLabel.NextHop})
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("label", types.YLeaf{"Label", remoteSplitHorizonGroupLabel.Label})

    remoteSplitHorizonGroupLabel.EntityData.YListKeys = []string {}

    return &(remoteSplitHorizonGroupLabel.EntityData)
}

// Evpn_Nodes_Node_AcIds
// EVPN AC ID table
type Evpn_Nodes_Node_AcIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN AC ID table. The type is slice of Evpn_Nodes_Node_AcIds_AcId.
    AcId []*Evpn_Nodes_Node_AcIds_AcId
}

func (acIds *Evpn_Nodes_Node_AcIds) GetEntityData() *types.CommonEntityData {
    acIds.EntityData.YFilter = acIds.YFilter
    acIds.EntityData.YangName = "ac-ids"
    acIds.EntityData.BundleName = "cisco_ios_xr"
    acIds.EntityData.ParentYangName = "node"
    acIds.EntityData.SegmentPath = "ac-ids"
    acIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acIds.EntityData.Children = types.NewOrderedMap()
    acIds.EntityData.Children.Append("ac-id", types.YChild{"AcId", nil})
    for i := range acIds.AcId {
        acIds.EntityData.Children.Append(types.GetSegmentPath(acIds.AcId[i]), types.YChild{"AcId", acIds.AcId[i]})
    }
    acIds.EntityData.Leafs = types.NewOrderedMap()

    acIds.EntityData.YListKeys = []string {}

    return &(acIds.EntityData)
}

// Evpn_Nodes_Node_AcIds_AcId
// EVPN AC ID table
type Evpn_Nodes_Node_AcIds_AcId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // AC ID. The type is interface{} with range: 0..4294967295.
    AcId interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Nodes_Node_AcIds_AcId_EvpnInstance
}

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetEntityData() *types.CommonEntityData {
    acId.EntityData.YFilter = acId.YFilter
    acId.EntityData.YangName = "ac-id"
    acId.EntityData.BundleName = "cisco_ios_xr"
    acId.EntityData.ParentYangName = "ac-ids"
    acId.EntityData.SegmentPath = "ac-id"
    acId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acId.EntityData.Children = types.NewOrderedMap()
    acId.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &acId.EvpnInstance})
    acId.EntityData.Leafs = types.NewOrderedMap()
    acId.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", acId.Evi})
    acId.EntityData.Leafs.Append("ac-id", types.YLeaf{"AcId", acId.AcId})
    acId.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", acId.Neighbor})

    acId.EntityData.YListKeys = []string {}

    return &(acId.EntityData)
}

// Evpn_Nodes_Node_AcIds_AcId_EvpnInstance
// EVPN Instance summary information
type Evpn_Nodes_Node_AcIds_AcId_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Nodes_Node_AcIds_AcId_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "ac-id"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Active
// Active EVPN operational data
type Evpn_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Group Table.
    EvpnGroups Evpn_Active_EvpnGroups

    // EVPN Remote SHG table.
    RemoteShgs Evpn_Active_RemoteShgs

    // L2VPN EVPN Client.
    Client Evpn_Active_Client

    // EVPN IGMP table.
    Igmps Evpn_Active_Igmps

    // L2VPN EVPN EVI Table.
    Evis Evpn_Active_Evis

    // L2VPN EVPN Summary.
    Summary Evpn_Active_Summary

    // L2VPN EVI Detail Table.
    EviDetail Evpn_Active_EviDetail

    // L2VPN EVPN TEP Table.
    Teps Evpn_Active_Teps

    // EVPN Internal Label Table.
    InternalLabels Evpn_Active_InternalLabels

    // EVPN Ethernet-Segment Table.
    EthernetSegments Evpn_Active_EthernetSegments

    // EVPN AC ID table.
    AcIds Evpn_Active_AcIds
}

func (active *Evpn_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "evpn"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("evpn-groups", types.YChild{"EvpnGroups", &active.EvpnGroups})
    active.EntityData.Children.Append("remote-shgs", types.YChild{"RemoteShgs", &active.RemoteShgs})
    active.EntityData.Children.Append("client", types.YChild{"Client", &active.Client})
    active.EntityData.Children.Append("igmps", types.YChild{"Igmps", &active.Igmps})
    active.EntityData.Children.Append("evis", types.YChild{"Evis", &active.Evis})
    active.EntityData.Children.Append("summary", types.YChild{"Summary", &active.Summary})
    active.EntityData.Children.Append("evi-detail", types.YChild{"EviDetail", &active.EviDetail})
    active.EntityData.Children.Append("teps", types.YChild{"Teps", &active.Teps})
    active.EntityData.Children.Append("internal-labels", types.YChild{"InternalLabels", &active.InternalLabels})
    active.EntityData.Children.Append("ethernet-segments", types.YChild{"EthernetSegments", &active.EthernetSegments})
    active.EntityData.Children.Append("ac-ids", types.YChild{"AcIds", &active.AcIds})
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// Evpn_Active_EvpnGroups
// EVPN Group Table
type Evpn_Active_EvpnGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Group information. The type is slice of
    // Evpn_Active_EvpnGroups_EvpnGroup.
    EvpnGroup []*Evpn_Active_EvpnGroups_EvpnGroup
}

func (evpnGroups *Evpn_Active_EvpnGroups) GetEntityData() *types.CommonEntityData {
    evpnGroups.EntityData.YFilter = evpnGroups.YFilter
    evpnGroups.EntityData.YangName = "evpn-groups"
    evpnGroups.EntityData.BundleName = "cisco_ios_xr"
    evpnGroups.EntityData.ParentYangName = "active"
    evpnGroups.EntityData.SegmentPath = "evpn-groups"
    evpnGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroups.EntityData.Children = types.NewOrderedMap()
    evpnGroups.EntityData.Children.Append("evpn-group", types.YChild{"EvpnGroup", nil})
    for i := range evpnGroups.EvpnGroup {
        evpnGroups.EntityData.Children.Append(types.GetSegmentPath(evpnGroups.EvpnGroup[i]), types.YChild{"EvpnGroup", evpnGroups.EvpnGroup[i]})
    }
    evpnGroups.EntityData.Leafs = types.NewOrderedMap()

    evpnGroups.EntityData.YListKeys = []string {}

    return &(evpnGroups.EntityData)
}

// Evpn_Active_EvpnGroups_EvpnGroup
// EVPN Group information
type Evpn_Active_EvpnGroups_EvpnGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN group number. The type is interface{} with
    // range: 1..4294967295.
    GroupNumber interface{}

    // EVPN Group ID. The type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // EVPN Group State. The type is EvpnGrp.
    State interface{}

    // EVPN Group Core Interface table. The type is slice of
    // Evpn_Active_EvpnGroups_EvpnGroup_CoreInterface.
    CoreInterface []*Evpn_Active_EvpnGroups_EvpnGroup_CoreInterface

    // EVPN Access Core Interface table. The type is slice of
    // Evpn_Active_EvpnGroups_EvpnGroup_AccessInterface.
    AccessInterface []*Evpn_Active_EvpnGroups_EvpnGroup_AccessInterface
}

func (evpnGroup *Evpn_Active_EvpnGroups_EvpnGroup) GetEntityData() *types.CommonEntityData {
    evpnGroup.EntityData.YFilter = evpnGroup.YFilter
    evpnGroup.EntityData.YangName = "evpn-group"
    evpnGroup.EntityData.BundleName = "cisco_ios_xr"
    evpnGroup.EntityData.ParentYangName = "evpn-groups"
    evpnGroup.EntityData.SegmentPath = "evpn-group" + types.AddKeyToken(evpnGroup.GroupNumber, "group-number")
    evpnGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroup.EntityData.Children = types.NewOrderedMap()
    evpnGroup.EntityData.Children.Append("core-interface", types.YChild{"CoreInterface", nil})
    for i := range evpnGroup.CoreInterface {
        evpnGroup.EntityData.Children.Append(types.GetSegmentPath(evpnGroup.CoreInterface[i]), types.YChild{"CoreInterface", evpnGroup.CoreInterface[i]})
    }
    evpnGroup.EntityData.Children.Append("access-interface", types.YChild{"AccessInterface", nil})
    for i := range evpnGroup.AccessInterface {
        evpnGroup.EntityData.Children.Append(types.GetSegmentPath(evpnGroup.AccessInterface[i]), types.YChild{"AccessInterface", evpnGroup.AccessInterface[i]})
    }
    evpnGroup.EntityData.Leafs = types.NewOrderedMap()
    evpnGroup.EntityData.Leafs.Append("group-number", types.YLeaf{"GroupNumber", evpnGroup.GroupNumber})
    evpnGroup.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", evpnGroup.GroupId})
    evpnGroup.EntityData.Leafs.Append("state", types.YLeaf{"State", evpnGroup.State})

    evpnGroup.EntityData.YListKeys = []string {"GroupNumber"}

    return &(evpnGroup.EntityData)
}

// Evpn_Active_EvpnGroups_EvpnGroup_CoreInterface
// EVPN Group Core Interface table
type Evpn_Active_EvpnGroups_EvpnGroup_CoreInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface State. The type is ImStateEnum.
    State interface{}
}

func (coreInterface *Evpn_Active_EvpnGroups_EvpnGroup_CoreInterface) GetEntityData() *types.CommonEntityData {
    coreInterface.EntityData.YFilter = coreInterface.YFilter
    coreInterface.EntityData.YangName = "core-interface"
    coreInterface.EntityData.BundleName = "cisco_ios_xr"
    coreInterface.EntityData.ParentYangName = "evpn-group"
    coreInterface.EntityData.SegmentPath = "core-interface"
    coreInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coreInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coreInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coreInterface.EntityData.Children = types.NewOrderedMap()
    coreInterface.EntityData.Leafs = types.NewOrderedMap()
    coreInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", coreInterface.InterfaceName})
    coreInterface.EntityData.Leafs.Append("state", types.YLeaf{"State", coreInterface.State})

    coreInterface.EntityData.YListKeys = []string {}

    return &(coreInterface.EntityData)
}

// Evpn_Active_EvpnGroups_EvpnGroup_AccessInterface
// EVPN Access Core Interface table
type Evpn_Active_EvpnGroups_EvpnGroup_AccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface State. The type is ImStateEnum.
    State interface{}
}

func (accessInterface *Evpn_Active_EvpnGroups_EvpnGroup_AccessInterface) GetEntityData() *types.CommonEntityData {
    accessInterface.EntityData.YFilter = accessInterface.YFilter
    accessInterface.EntityData.YangName = "access-interface"
    accessInterface.EntityData.BundleName = "cisco_ios_xr"
    accessInterface.EntityData.ParentYangName = "evpn-group"
    accessInterface.EntityData.SegmentPath = "access-interface"
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = types.NewOrderedMap()
    accessInterface.EntityData.Leafs = types.NewOrderedMap()
    accessInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterface.InterfaceName})
    accessInterface.EntityData.Leafs.Append("state", types.YLeaf{"State", accessInterface.State})

    accessInterface.EntityData.YListKeys = []string {}

    return &(accessInterface.EntityData)
}

// Evpn_Active_RemoteShgs
// EVPN Remote SHG table
type Evpn_Active_RemoteShgs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Remote SHG. The type is slice of Evpn_Active_RemoteShgs_RemoteShg.
    RemoteShg []*Evpn_Active_RemoteShgs_RemoteShg
}

func (remoteShgs *Evpn_Active_RemoteShgs) GetEntityData() *types.CommonEntityData {
    remoteShgs.EntityData.YFilter = remoteShgs.YFilter
    remoteShgs.EntityData.YangName = "remote-shgs"
    remoteShgs.EntityData.BundleName = "cisco_ios_xr"
    remoteShgs.EntityData.ParentYangName = "active"
    remoteShgs.EntityData.SegmentPath = "remote-shgs"
    remoteShgs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteShgs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteShgs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteShgs.EntityData.Children = types.NewOrderedMap()
    remoteShgs.EntityData.Children.Append("remote-shg", types.YChild{"RemoteShg", nil})
    for i := range remoteShgs.RemoteShg {
        remoteShgs.EntityData.Children.Append(types.GetSegmentPath(remoteShgs.RemoteShg[i]), types.YChild{"RemoteShg", remoteShgs.RemoteShg[i]})
    }
    remoteShgs.EntityData.Leafs = types.NewOrderedMap()

    remoteShgs.EntityData.YListKeys = []string {}

    return &(remoteShgs.EntityData)
}

// Evpn_Active_RemoteShgs_RemoteShg
// EVPN Remote SHG
type Evpn_Active_RemoteShgs_RemoteShg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Segment id. The type is slice of
    // Evpn_Active_RemoteShgs_RemoteShg_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Active_RemoteShgs_RemoteShg_EthernetSegmentIdentifier

    // Remote split horizon group labels. The type is slice of
    // Evpn_Active_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel.
    RemoteSplitHorizonGroupLabel []*Evpn_Active_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel
}

func (remoteShg *Evpn_Active_RemoteShgs_RemoteShg) GetEntityData() *types.CommonEntityData {
    remoteShg.EntityData.YFilter = remoteShg.YFilter
    remoteShg.EntityData.YangName = "remote-shg"
    remoteShg.EntityData.BundleName = "cisco_ios_xr"
    remoteShg.EntityData.ParentYangName = "remote-shgs"
    remoteShg.EntityData.SegmentPath = "remote-shg"
    remoteShg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteShg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteShg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteShg.EntityData.Children = types.NewOrderedMap()
    remoteShg.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range remoteShg.EthernetSegmentIdentifier {
        remoteShg.EntityData.Children.Append(types.GetSegmentPath(remoteShg.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", remoteShg.EthernetSegmentIdentifier[i]})
    }
    remoteShg.EntityData.Children.Append("remote-split-horizon-group-label", types.YChild{"RemoteSplitHorizonGroupLabel", nil})
    for i := range remoteShg.RemoteSplitHorizonGroupLabel {
        remoteShg.EntityData.Children.Append(types.GetSegmentPath(remoteShg.RemoteSplitHorizonGroupLabel[i]), types.YChild{"RemoteSplitHorizonGroupLabel", remoteShg.RemoteSplitHorizonGroupLabel[i]})
    }
    remoteShg.EntityData.Leafs = types.NewOrderedMap()
    remoteShg.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", remoteShg.Esi1})
    remoteShg.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", remoteShg.Esi2})
    remoteShg.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", remoteShg.Esi3})
    remoteShg.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", remoteShg.Esi4})
    remoteShg.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", remoteShg.Esi5})

    remoteShg.EntityData.YListKeys = []string {}

    return &(remoteShg.EntityData)
}

// Evpn_Active_RemoteShgs_RemoteShg_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Active_RemoteShgs_RemoteShg_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Active_RemoteShgs_RemoteShg_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "remote-shg"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Active_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Active_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Split horizon label associated with next-hop address. The type is
    // interface{} with range: 0..4294967295.
    Label interface{}
}

func (remoteSplitHorizonGroupLabel *Evpn_Active_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel) GetEntityData() *types.CommonEntityData {
    remoteSplitHorizonGroupLabel.EntityData.YFilter = remoteSplitHorizonGroupLabel.YFilter
    remoteSplitHorizonGroupLabel.EntityData.YangName = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.BundleName = "cisco_ios_xr"
    remoteSplitHorizonGroupLabel.EntityData.ParentYangName = "remote-shg"
    remoteSplitHorizonGroupLabel.EntityData.SegmentPath = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteSplitHorizonGroupLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteSplitHorizonGroupLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteSplitHorizonGroupLabel.EntityData.Children = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", remoteSplitHorizonGroupLabel.NextHop})
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("label", types.YLeaf{"Label", remoteSplitHorizonGroupLabel.Label})

    remoteSplitHorizonGroupLabel.EntityData.YListKeys = []string {}

    return &(remoteSplitHorizonGroupLabel.EntityData)
}

// Evpn_Active_Client
// L2VPN EVPN Client
type Evpn_Active_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (client *Evpn_Active_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "active"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// Evpn_Active_Igmps
// EVPN IGMP table
type Evpn_Active_Igmps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Remote. The type is slice of Evpn_Active_Igmps_Igmp.
    Igmp []*Evpn_Active_Igmps_Igmp
}

func (igmps *Evpn_Active_Igmps) GetEntityData() *types.CommonEntityData {
    igmps.EntityData.YFilter = igmps.YFilter
    igmps.EntityData.YangName = "igmps"
    igmps.EntityData.BundleName = "cisco_ios_xr"
    igmps.EntityData.ParentYangName = "active"
    igmps.EntityData.SegmentPath = "igmps"
    igmps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igmps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igmps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igmps.EntityData.Children = types.NewOrderedMap()
    igmps.EntityData.Children.Append("igmp", types.YChild{"Igmp", nil})
    for i := range igmps.Igmp {
        igmps.EntityData.Children.Append(types.GetSegmentPath(igmps.Igmp[i]), types.YChild{"Igmp", igmps.Igmp[i]})
    }
    igmps.EntityData.Leafs = types.NewOrderedMap()

    igmps.EntityData.YListKeys = []string {}

    return &(igmps.EntityData)
}

// Evpn_Active_Igmps_Igmp
// EVPN Remote
type Evpn_Active_Igmps_Igmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Join=0, Leave=1. The type is interface{} with range: 0..4294967295.
    IsLeave interface{}

    // BP xcid. The type is interface{} with range: 0..4294967295.
    Bpxcid interface{}

    // EVI/BD. The type is interface{} with range: 0..4294967295.
    Evibd interface{}

    // Source IP Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SrcIp interface{}

    // Group IP Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GrpIp interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Segment Name. The type is string.
    EthernetSegmentName interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // BD id. The type is interface{} with range: 0..4294967295.
    BdId interface{}

    // Route Type. The type is EvpnIgmpMsg.
    RouteType interface{}

    // Source IP Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddr interface{}

    // Group IP Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddr interface{}

    // Ethernet Tag id. The type is interface{} with range: 0..4294967295.
    EthernetTagId interface{}

    // IGMP Version. The type is EvpnIgmpVersion.
    IgmpVersion interface{}

    // IGMP Group Type. The type is EvpnIgmpGrp.
    IgmpGroupType interface{}

    // Max Response Time. The type is interface{} with range: 0..255.
    MaXResponseTime interface{}

    // Resolved. The type is bool.
    Resolved interface{}

    // Source Info.
    SourceInfo Evpn_Active_Igmps_Igmp_SourceInfo

    // Ethernet Segment id. The type is slice of
    // Evpn_Active_Igmps_Igmp_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Active_Igmps_Igmp_EthernetSegmentIdentifier

    // List of nexthop IPv6 addresses. The type is slice of
    // Evpn_Active_Igmps_Igmp_NextHop.
    NextHop []*Evpn_Active_Igmps_Igmp_NextHop
}

func (igmp *Evpn_Active_Igmps_Igmp) GetEntityData() *types.CommonEntityData {
    igmp.EntityData.YFilter = igmp.YFilter
    igmp.EntityData.YangName = "igmp"
    igmp.EntityData.BundleName = "cisco_ios_xr"
    igmp.EntityData.ParentYangName = "igmps"
    igmp.EntityData.SegmentPath = "igmp"
    igmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igmp.EntityData.Children = types.NewOrderedMap()
    igmp.EntityData.Children.Append("source-info", types.YChild{"SourceInfo", &igmp.SourceInfo})
    igmp.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range igmp.EthernetSegmentIdentifier {
        igmp.EntityData.Children.Append(types.GetSegmentPath(igmp.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", igmp.EthernetSegmentIdentifier[i]})
    }
    igmp.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range igmp.NextHop {
        igmp.EntityData.Children.Append(types.GetSegmentPath(igmp.NextHop[i]), types.YChild{"NextHop", igmp.NextHop[i]})
    }
    igmp.EntityData.Leafs = types.NewOrderedMap()
    igmp.EntityData.Leafs.Append("is-leave", types.YLeaf{"IsLeave", igmp.IsLeave})
    igmp.EntityData.Leafs.Append("bpxcid", types.YLeaf{"Bpxcid", igmp.Bpxcid})
    igmp.EntityData.Leafs.Append("evibd", types.YLeaf{"Evibd", igmp.Evibd})
    igmp.EntityData.Leafs.Append("src-ip", types.YLeaf{"SrcIp", igmp.SrcIp})
    igmp.EntityData.Leafs.Append("grp-ip", types.YLeaf{"GrpIp", igmp.GrpIp})
    igmp.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", igmp.Esi1})
    igmp.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", igmp.Esi2})
    igmp.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", igmp.Esi3})
    igmp.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", igmp.Esi4})
    igmp.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", igmp.Esi5})
    igmp.EntityData.Leafs.Append("ethernet-segment-name", types.YLeaf{"EthernetSegmentName", igmp.EthernetSegmentName})
    igmp.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", igmp.Evi})
    igmp.EntityData.Leafs.Append("bd-id", types.YLeaf{"BdId", igmp.BdId})
    igmp.EntityData.Leafs.Append("route-type", types.YLeaf{"RouteType", igmp.RouteType})
    igmp.EntityData.Leafs.Append("source-addr", types.YLeaf{"SourceAddr", igmp.SourceAddr})
    igmp.EntityData.Leafs.Append("group-addr", types.YLeaf{"GroupAddr", igmp.GroupAddr})
    igmp.EntityData.Leafs.Append("ethernet-tag-id", types.YLeaf{"EthernetTagId", igmp.EthernetTagId})
    igmp.EntityData.Leafs.Append("igmp-version", types.YLeaf{"IgmpVersion", igmp.IgmpVersion})
    igmp.EntityData.Leafs.Append("igmp-group-type", types.YLeaf{"IgmpGroupType", igmp.IgmpGroupType})
    igmp.EntityData.Leafs.Append("ma-x-response-time", types.YLeaf{"MaXResponseTime", igmp.MaXResponseTime})
    igmp.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", igmp.Resolved})

    igmp.EntityData.YListKeys = []string {}

    return &(igmp.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo
// Source Info
type Evpn_Active_Igmps_Igmp_SourceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is EvpnIgmpSource.
    Type interface{}

    // remote info. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteInfo interface{}

    // local info.
    LocalInfo Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo
}

func (sourceInfo *Evpn_Active_Igmps_Igmp_SourceInfo) GetEntityData() *types.CommonEntityData {
    sourceInfo.EntityData.YFilter = sourceInfo.YFilter
    sourceInfo.EntityData.YangName = "source-info"
    sourceInfo.EntityData.BundleName = "cisco_ios_xr"
    sourceInfo.EntityData.ParentYangName = "igmp"
    sourceInfo.EntityData.SegmentPath = "source-info"
    sourceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceInfo.EntityData.Children = types.NewOrderedMap()
    sourceInfo.EntityData.Children.Append("local-info", types.YChild{"LocalInfo", &sourceInfo.LocalInfo})
    sourceInfo.EntityData.Leafs = types.NewOrderedMap()
    sourceInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", sourceInfo.Type})
    sourceInfo.EntityData.Leafs.Append("remote-info", types.YLeaf{"RemoteInfo", sourceInfo.RemoteInfo})

    sourceInfo.EntityData.YListKeys = []string {}

    return &(sourceInfo.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo
// local info
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..81.
    Name interface{}

    // Interface MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Payload bytes. The type is interface{} with range: 0..65535. Units are
    // byte.
    PayloadBytes interface{}

    // Interface parameters.
    Parameters Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters
}

func (localInfo *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo) GetEntityData() *types.CommonEntityData {
    localInfo.EntityData.YFilter = localInfo.YFilter
    localInfo.EntityData.YangName = "local-info"
    localInfo.EntityData.BundleName = "cisco_ios_xr"
    localInfo.EntityData.ParentYangName = "source-info"
    localInfo.EntityData.SegmentPath = "local-info"
    localInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localInfo.EntityData.Children = types.NewOrderedMap()
    localInfo.EntityData.Children.Append("parameters", types.YChild{"Parameters", &localInfo.Parameters})
    localInfo.EntityData.Leafs = types.NewOrderedMap()
    localInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", localInfo.Name})
    localInfo.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", localInfo.Mtu})
    localInfo.EntityData.Leafs.Append("payload-bytes", types.YLeaf{"PayloadBytes", localInfo.PayloadBytes})

    localInfo.EntityData.YListKeys = []string {}

    return &(localInfo.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters
// Interface parameters
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is L2vpnInterface.
    Type interface{}

    // Ethernet.
    Ethernet Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet

    // VLAN.
    Vlan Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan

    // TDM.
    Tdm Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm

    // ATM.
    Atm Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm

    // Frame Relay.
    Fr Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr

    // PW Ether.
    PseudowireEther Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther

    // PW IW.
    PseudowireIw Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw
}

func (parameters *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters) GetEntityData() *types.CommonEntityData {
    parameters.EntityData.YFilter = parameters.YFilter
    parameters.EntityData.YangName = "parameters"
    parameters.EntityData.BundleName = "cisco_ios_xr"
    parameters.EntityData.ParentYangName = "local-info"
    parameters.EntityData.SegmentPath = "parameters"
    parameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parameters.EntityData.Children = types.NewOrderedMap()
    parameters.EntityData.Children.Append("ethernet", types.YChild{"Ethernet", &parameters.Ethernet})
    parameters.EntityData.Children.Append("vlan", types.YChild{"Vlan", &parameters.Vlan})
    parameters.EntityData.Children.Append("tdm", types.YChild{"Tdm", &parameters.Tdm})
    parameters.EntityData.Children.Append("atm", types.YChild{"Atm", &parameters.Atm})
    parameters.EntityData.Children.Append("fr", types.YChild{"Fr", &parameters.Fr})
    parameters.EntityData.Children.Append("pseudowire-ether", types.YChild{"PseudowireEther", &parameters.PseudowireEther})
    parameters.EntityData.Children.Append("pseudowire-iw", types.YChild{"PseudowireIw", &parameters.PseudowireIw})
    parameters.EntityData.Leafs = types.NewOrderedMap()
    parameters.EntityData.Leafs.Append("type", types.YLeaf{"Type", parameters.Type})

    parameters.EntityData.YListKeys = []string {}

    return &(parameters.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet
// Ethernet
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XConnect tags. The type is interface{} with range: 0..255.
    XconnectTags interface{}
}

func (ethernet *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "cisco_ios_xr"
    ethernet.EntityData.ParentYangName = "parameters"
    ethernet.EntityData.SegmentPath = "ethernet"
    ethernet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernet.EntityData.Children = types.NewOrderedMap()
    ethernet.EntityData.Leafs = types.NewOrderedMap()
    ethernet.EntityData.Leafs.Append("xconnect-tags", types.YLeaf{"XconnectTags", ethernet.XconnectTags})

    ethernet.EntityData.YListKeys = []string {}

    return &(ethernet.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan
// VLAN
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XConnect tags. The type is interface{} with range: 0..255.
    XconnectTags interface{}

    // VLAN rewrite tag. The type is interface{} with range: 0..65535.
    VlanRewriteTag interface{}

    // Simple EFP. The type is interface{} with range: 0..255.
    SimpleEfp interface{}

    // Encapsulation Type. The type is interface{} with range: 0..255.
    EncapsulationType interface{}

    // Outer Tag. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Rewrite Tags. The type is slice of
    // Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag.
    RewriteTag []*Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag

    // vlan range. The type is slice of
    // Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange.
    VlanRange []*Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange
}

func (vlan *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "cisco_ios_xr"
    vlan.EntityData.ParentYangName = "parameters"
    vlan.EntityData.SegmentPath = "vlan"
    vlan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlan.EntityData.Children = types.NewOrderedMap()
    vlan.EntityData.Children.Append("rewrite-tag", types.YChild{"RewriteTag", nil})
    for i := range vlan.RewriteTag {
        vlan.EntityData.Children.Append(types.GetSegmentPath(vlan.RewriteTag[i]), types.YChild{"RewriteTag", vlan.RewriteTag[i]})
    }
    vlan.EntityData.Children.Append("vlan-range", types.YChild{"VlanRange", nil})
    for i := range vlan.VlanRange {
        vlan.EntityData.Children.Append(types.GetSegmentPath(vlan.VlanRange[i]), types.YChild{"VlanRange", vlan.VlanRange[i]})
    }
    vlan.EntityData.Leafs = types.NewOrderedMap()
    vlan.EntityData.Leafs.Append("xconnect-tags", types.YLeaf{"XconnectTags", vlan.XconnectTags})
    vlan.EntityData.Leafs.Append("vlan-rewrite-tag", types.YLeaf{"VlanRewriteTag", vlan.VlanRewriteTag})
    vlan.EntityData.Leafs.Append("simple-efp", types.YLeaf{"SimpleEfp", vlan.SimpleEfp})
    vlan.EntityData.Leafs.Append("encapsulation-type", types.YLeaf{"EncapsulationType", vlan.EncapsulationType})
    vlan.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", vlan.OuterTag})

    vlan.EntityData.YListKeys = []string {}

    return &(vlan.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag
// Rewrite Tags
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535.
    Entry interface{}
}

func (rewriteTag *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag) GetEntityData() *types.CommonEntityData {
    rewriteTag.EntityData.YFilter = rewriteTag.YFilter
    rewriteTag.EntityData.YangName = "rewrite-tag"
    rewriteTag.EntityData.BundleName = "cisco_ios_xr"
    rewriteTag.EntityData.ParentYangName = "vlan"
    rewriteTag.EntityData.SegmentPath = "rewrite-tag"
    rewriteTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rewriteTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rewriteTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rewriteTag.EntityData.Children = types.NewOrderedMap()
    rewriteTag.EntityData.Leafs = types.NewOrderedMap()
    rewriteTag.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", rewriteTag.Entry})

    rewriteTag.EntityData.YListKeys = []string {}

    return &(rewriteTag.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange
// vlan range
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lower. The type is interface{} with range: 0..65535.
    Lower interface{}

    // Upper. The type is interface{} with range: 0..65535.
    Upper interface{}
}

func (vlanRange *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange) GetEntityData() *types.CommonEntityData {
    vlanRange.EntityData.YFilter = vlanRange.YFilter
    vlanRange.EntityData.YangName = "vlan-range"
    vlanRange.EntityData.BundleName = "cisco_ios_xr"
    vlanRange.EntityData.ParentYangName = "vlan"
    vlanRange.EntityData.SegmentPath = "vlan-range"
    vlanRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanRange.EntityData.Children = types.NewOrderedMap()
    vlanRange.EntityData.Leafs = types.NewOrderedMap()
    vlanRange.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", vlanRange.Lower})
    vlanRange.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", vlanRange.Upper})

    vlanRange.EntityData.YListKeys = []string {}

    return &(vlanRange.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm
// TDM
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeslots separated by , or - from 1 to 31. : indicates individual timeslot
    // and - represents a range.E.g. 1-3,5 represents timeslots 1, 2, 3, and 5.
    // The type is string.
    TimeslotGroup interface{}

    // Timeslot rate in units of Kbps. The type is interface{} with range: 0..255.
    // Units are kbit/s.
    TimeslotRate interface{}

    // TDM mode. The type is L2vpnTdmMode.
    TdmMode interface{}

    // TDM options.
    TdmOptions Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions
}

func (tdm *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm) GetEntityData() *types.CommonEntityData {
    tdm.EntityData.YFilter = tdm.YFilter
    tdm.EntityData.YangName = "tdm"
    tdm.EntityData.BundleName = "cisco_ios_xr"
    tdm.EntityData.ParentYangName = "parameters"
    tdm.EntityData.SegmentPath = "tdm"
    tdm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tdm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tdm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tdm.EntityData.Children = types.NewOrderedMap()
    tdm.EntityData.Children.Append("tdm-options", types.YChild{"TdmOptions", &tdm.TdmOptions})
    tdm.EntityData.Leafs = types.NewOrderedMap()
    tdm.EntityData.Leafs.Append("timeslot-group", types.YLeaf{"TimeslotGroup", tdm.TimeslotGroup})
    tdm.EntityData.Leafs.Append("timeslot-rate", types.YLeaf{"TimeslotRate", tdm.TimeslotRate})
    tdm.EntityData.Leafs.Append("tdm-mode", types.YLeaf{"TdmMode", tdm.TdmMode})

    tdm.EntityData.YListKeys = []string {}

    return &(tdm.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions
// TDM options
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TDM payload bytes. The type is interface{} with range: 0..65535. Units are
    // byte.
    PayloadBytes interface{}

    // TDM bit rate in units of Kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    BitRate interface{}

    // RTP header. The type is L2vpnTdmRtpOption.
    Rtp interface{}

    // TDM Timestamping mode. The type is L2vpnTimeStampMode.
    TimestampMode interface{}

    // Signalling packets. The type is interface{} with range: 0..255.
    SignallingPackets interface{}

    // CAS. The type is interface{} with range: 0..255.
    Cas interface{}

    // RTP header payload type. The type is interface{} with range: 0..255.
    RtpHeaderPayloadType interface{}

    // Timestamping clock frequency in units of 8Khz. The type is interface{} with
    // range: 0..65535.
    TimestampClockFreq interface{}

    // Synchronization Source identifier. The type is interface{} with range:
    // 0..4294967295.
    Ssrc interface{}
}

func (tdmOptions *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions) GetEntityData() *types.CommonEntityData {
    tdmOptions.EntityData.YFilter = tdmOptions.YFilter
    tdmOptions.EntityData.YangName = "tdm-options"
    tdmOptions.EntityData.BundleName = "cisco_ios_xr"
    tdmOptions.EntityData.ParentYangName = "tdm"
    tdmOptions.EntityData.SegmentPath = "tdm-options"
    tdmOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tdmOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tdmOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tdmOptions.EntityData.Children = types.NewOrderedMap()
    tdmOptions.EntityData.Leafs = types.NewOrderedMap()
    tdmOptions.EntityData.Leafs.Append("payload-bytes", types.YLeaf{"PayloadBytes", tdmOptions.PayloadBytes})
    tdmOptions.EntityData.Leafs.Append("bit-rate", types.YLeaf{"BitRate", tdmOptions.BitRate})
    tdmOptions.EntityData.Leafs.Append("rtp", types.YLeaf{"Rtp", tdmOptions.Rtp})
    tdmOptions.EntityData.Leafs.Append("timestamp-mode", types.YLeaf{"TimestampMode", tdmOptions.TimestampMode})
    tdmOptions.EntityData.Leafs.Append("signalling-packets", types.YLeaf{"SignallingPackets", tdmOptions.SignallingPackets})
    tdmOptions.EntityData.Leafs.Append("cas", types.YLeaf{"Cas", tdmOptions.Cas})
    tdmOptions.EntityData.Leafs.Append("rtp-header-payload-type", types.YLeaf{"RtpHeaderPayloadType", tdmOptions.RtpHeaderPayloadType})
    tdmOptions.EntityData.Leafs.Append("timestamp-clock-freq", types.YLeaf{"TimestampClockFreq", tdmOptions.TimestampClockFreq})
    tdmOptions.EntityData.Leafs.Append("ssrc", types.YLeaf{"Ssrc", tdmOptions.Ssrc})

    tdmOptions.EntityData.YListKeys = []string {}

    return &(tdmOptions.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm
// ATM
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max number of cells packed. The type is interface{} with range: 0..65535.
    MaximumNumberCellsPacked interface{}

    // Max number of cells unpacked. The type is interface{} with range: 0..65535.
    MaximumNumberCellsUnPacked interface{}

    // ATM mode. The type is L2vpnAtmMode.
    AtmMode interface{}

    // Virtual path identifier. The type is interface{} with range: 0..65535.
    Vpi interface{}

    // Virtual channel identifier. The type is interface{} with range: 0..65535.
    Vci interface{}
}

func (atm *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm) GetEntityData() *types.CommonEntityData {
    atm.EntityData.YFilter = atm.YFilter
    atm.EntityData.YangName = "atm"
    atm.EntityData.BundleName = "cisco_ios_xr"
    atm.EntityData.ParentYangName = "parameters"
    atm.EntityData.SegmentPath = "atm"
    atm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    atm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    atm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    atm.EntityData.Children = types.NewOrderedMap()
    atm.EntityData.Leafs = types.NewOrderedMap()
    atm.EntityData.Leafs.Append("maximum-number-cells-packed", types.YLeaf{"MaximumNumberCellsPacked", atm.MaximumNumberCellsPacked})
    atm.EntityData.Leafs.Append("maximum-number-cells-un-packed", types.YLeaf{"MaximumNumberCellsUnPacked", atm.MaximumNumberCellsUnPacked})
    atm.EntityData.Leafs.Append("atm-mode", types.YLeaf{"AtmMode", atm.AtmMode})
    atm.EntityData.Leafs.Append("vpi", types.YLeaf{"Vpi", atm.Vpi})
    atm.EntityData.Leafs.Append("vci", types.YLeaf{"Vci", atm.Vci})

    atm.EntityData.YListKeys = []string {}

    return &(atm.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr
// Frame Relay
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Frame Relay mode. The type is L2vpnFrMode.
    FrMode interface{}

    // Data-link connection identifier. The type is interface{} with range:
    // 0..4294967295.
    Dlci interface{}
}

func (fr *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr) GetEntityData() *types.CommonEntityData {
    fr.EntityData.YFilter = fr.YFilter
    fr.EntityData.YangName = "fr"
    fr.EntityData.BundleName = "cisco_ios_xr"
    fr.EntityData.ParentYangName = "parameters"
    fr.EntityData.SegmentPath = "fr"
    fr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fr.EntityData.Children = types.NewOrderedMap()
    fr.EntityData.Leafs = types.NewOrderedMap()
    fr.EntityData.Leafs.Append("fr-mode", types.YLeaf{"FrMode", fr.FrMode})
    fr.EntityData.Leafs.Append("dlci", types.YLeaf{"Dlci", fr.Dlci})

    fr.EntityData.YListKeys = []string {}

    return &(fr.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther
// PW Ether
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is this Interface list valid. The type is bool.
    IsValid interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Interface list data.
    InterfaceList Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList
}

func (pseudowireEther *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther) GetEntityData() *types.CommonEntityData {
    pseudowireEther.EntityData.YFilter = pseudowireEther.YFilter
    pseudowireEther.EntityData.YangName = "pseudowire-ether"
    pseudowireEther.EntityData.BundleName = "cisco_ios_xr"
    pseudowireEther.EntityData.ParentYangName = "parameters"
    pseudowireEther.EntityData.SegmentPath = "pseudowire-ether"
    pseudowireEther.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireEther.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireEther.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireEther.EntityData.Children = types.NewOrderedMap()
    pseudowireEther.EntityData.Children.Append("interface-list", types.YChild{"InterfaceList", &pseudowireEther.InterfaceList})
    pseudowireEther.EntityData.Leafs = types.NewOrderedMap()
    pseudowireEther.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", pseudowireEther.IsValid})
    pseudowireEther.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", pseudowireEther.InternalLabel})

    pseudowireEther.EntityData.YListKeys = []string {}

    return &(pseudowireEther.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList
// Interface list data
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface-list name. The type is string with length: 0..33.
    InterfaceListName interface{}

    // Interface internal ID. The type is interface{} with range: 0..4294967295.
    InterfaceListId interface{}

    // Interfaces. The type is slice of
    // Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface.
    Interface []*Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface
}

func (interfaceList *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList) GetEntityData() *types.CommonEntityData {
    interfaceList.EntityData.YFilter = interfaceList.YFilter
    interfaceList.EntityData.YangName = "interface-list"
    interfaceList.EntityData.BundleName = "cisco_ios_xr"
    interfaceList.EntityData.ParentYangName = "pseudowire-ether"
    interfaceList.EntityData.SegmentPath = "interface-list"
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = types.NewOrderedMap()
    interfaceList.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceList.Interface {
        interfaceList.EntityData.Children.Append(types.GetSegmentPath(interfaceList.Interface[i]), types.YChild{"Interface", interfaceList.Interface[i]})
    }
    interfaceList.EntityData.Leafs = types.NewOrderedMap()
    interfaceList.EntityData.Leafs.Append("interface-list-name", types.YLeaf{"InterfaceListName", interfaceList.InterfaceListName})
    interfaceList.EntityData.Leafs.Append("interface-list-id", types.YLeaf{"InterfaceListId", interfaceList.InterfaceListId})

    interfaceList.EntityData.YListKeys = []string {}

    return &(interfaceList.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface
// Interfaces
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Replicate status. The type is IflistRepStatus.
    ReplicateStatus interface{}
}

func (self *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interface-list"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("replicate-status", types.YLeaf{"ReplicateStatus", self.ReplicateStatus})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw
// PW IW
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is this Interface list valid. The type is bool.
    IsValid interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Interface list data.
    InterfaceList Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList
}

func (pseudowireIw *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw) GetEntityData() *types.CommonEntityData {
    pseudowireIw.EntityData.YFilter = pseudowireIw.YFilter
    pseudowireIw.EntityData.YangName = "pseudowire-iw"
    pseudowireIw.EntityData.BundleName = "cisco_ios_xr"
    pseudowireIw.EntityData.ParentYangName = "parameters"
    pseudowireIw.EntityData.SegmentPath = "pseudowire-iw"
    pseudowireIw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireIw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireIw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireIw.EntityData.Children = types.NewOrderedMap()
    pseudowireIw.EntityData.Children.Append("interface-list", types.YChild{"InterfaceList", &pseudowireIw.InterfaceList})
    pseudowireIw.EntityData.Leafs = types.NewOrderedMap()
    pseudowireIw.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", pseudowireIw.IsValid})
    pseudowireIw.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", pseudowireIw.InternalLabel})

    pseudowireIw.EntityData.YListKeys = []string {}

    return &(pseudowireIw.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList
// Interface list data
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface-list name. The type is string with length: 0..33.
    InterfaceListName interface{}

    // Interface internal ID. The type is interface{} with range: 0..4294967295.
    InterfaceListId interface{}

    // Interfaces. The type is slice of
    // Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface.
    Interface []*Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface
}

func (interfaceList *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList) GetEntityData() *types.CommonEntityData {
    interfaceList.EntityData.YFilter = interfaceList.YFilter
    interfaceList.EntityData.YangName = "interface-list"
    interfaceList.EntityData.BundleName = "cisco_ios_xr"
    interfaceList.EntityData.ParentYangName = "pseudowire-iw"
    interfaceList.EntityData.SegmentPath = "interface-list"
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = types.NewOrderedMap()
    interfaceList.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceList.Interface {
        interfaceList.EntityData.Children.Append(types.GetSegmentPath(interfaceList.Interface[i]), types.YChild{"Interface", interfaceList.Interface[i]})
    }
    interfaceList.EntityData.Leafs = types.NewOrderedMap()
    interfaceList.EntityData.Leafs.Append("interface-list-name", types.YLeaf{"InterfaceListName", interfaceList.InterfaceListName})
    interfaceList.EntityData.Leafs.Append("interface-list-id", types.YLeaf{"InterfaceListId", interfaceList.InterfaceListId})

    interfaceList.EntityData.YListKeys = []string {}

    return &(interfaceList.EntityData)
}

// Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface
// Interfaces
type Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Replicate status. The type is IflistRepStatus.
    ReplicateStatus interface{}
}

func (self *Evpn_Active_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interface-list"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("replicate-status", types.YLeaf{"ReplicateStatus", self.ReplicateStatus})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Evpn_Active_Igmps_Igmp_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Active_Igmps_Igmp_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Active_Igmps_Igmp_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "igmp"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Active_Igmps_Igmp_NextHop
// List of nexthop IPv6 addresses
type Evpn_Active_Igmps_Igmp_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // DF Dont Premption. The type is bool.
    DfDontPrempt interface{}

    // DF Election Mode Configured. The type is interface{} with range: 0..255.
    DfType interface{}

    // DF Election Preference Set. The type is interface{} with range: 0..65535.
    DfPref interface{}
}

func (nextHop *Evpn_Active_Igmps_Igmp_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "igmp"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", nextHop.NextHop})
    nextHop.EntityData.Leafs.Append("df-dont-prempt", types.YLeaf{"DfDontPrempt", nextHop.DfDontPrempt})
    nextHop.EntityData.Leafs.Append("df-type", types.YLeaf{"DfType", nextHop.DfType})
    nextHop.EntityData.Leafs.Append("df-pref", types.YLeaf{"DfPref", nextHop.DfPref})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// Evpn_Active_Evis
// L2VPN EVPN EVI Table
type Evpn_Active_Evis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Entry. The type is slice of Evpn_Active_Evis_Evi.
    Evi []*Evpn_Active_Evis_Evi
}

func (evis *Evpn_Active_Evis) GetEntityData() *types.CommonEntityData {
    evis.EntityData.YFilter = evis.YFilter
    evis.EntityData.YangName = "evis"
    evis.EntityData.BundleName = "cisco_ios_xr"
    evis.EntityData.ParentYangName = "active"
    evis.EntityData.SegmentPath = "evis"
    evis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evis.EntityData.Children = types.NewOrderedMap()
    evis.EntityData.Children.Append("evi", types.YChild{"Evi", nil})
    for i := range evis.Evi {
        evis.EntityData.Children.Append(types.GetSegmentPath(evis.Evi[i]), types.YChild{"Evi", evis.Evi[i]})
    }
    evis.EntityData.Leafs = types.NewOrderedMap()

    evis.EntityData.YListKeys = []string {}

    return &(evis.EntityData)
}

// Evpn_Active_Evis_Evi
// L2VPN EVPN EVI Entry
type Evpn_Active_Evis_Evi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evi *Evpn_Active_Evis_Evi) GetEntityData() *types.CommonEntityData {
    evi.EntityData.YFilter = evi.YFilter
    evi.EntityData.YangName = "evi"
    evi.EntityData.BundleName = "cisco_ios_xr"
    evi.EntityData.ParentYangName = "evis"
    evi.EntityData.SegmentPath = "evi"
    evi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evi.EntityData.Children = types.NewOrderedMap()
    evi.EntityData.Leafs = types.NewOrderedMap()
    evi.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", evi.Evi})
    evi.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", evi.Encapsulation})
    evi.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evi.EthernetVpnId})
    evi.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evi.EncapsulationXr})
    evi.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evi.BdName})
    evi.EntityData.Leafs.Append("type", types.YLeaf{"Type", evi.Type})

    evi.EntityData.YListKeys = []string {}

    return &(evi.EntityData)
}

// Evpn_Active_Summary
// L2VPN EVPN Summary
type Evpn_Active_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Router ID. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // BGP AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // Number of EVI DB Entries. The type is interface{} with range:
    // 0..4294967295.
    EvIs interface{}

    // Number of Tunnel Endpoint DB Entries. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpoints interface{}

    // Number of Local MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalMacRoutes interface{}

    // Number of Local IPv4 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalIpv4MacRoutes interface{}

    // Number of Local IPv6 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalIpv6MacRoutes interface{}

    // Number of ES:Global MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    EsGlobalMacRoutes interface{}

    // Number of Remote MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteMacRoutes interface{}

    // Number of Remote Soo MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteSooMacRoutes interface{}

    // Number of Remote IPv4 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteIpv4MacRoutes interface{}

    // Number of Remote IPv6 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteIpv6MacRoutes interface{}

    // Number of Local IMCAST Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalImcastRoutes interface{}

    // Number of Remote IMCAST Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteImcastRoutes interface{}

    // Number of Internal Labels. The type is interface{} with range:
    // 0..4294967295.
    Labels interface{}

    // Number of ES Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    EsEntries interface{}

    // Number of neighbor Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    NeighborEntries interface{}

    // Number of Local EAD Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    LocalEadRoutes interface{}

    // Number of Remote EAD Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    RemoteEadRoutes interface{}

    // Global Source MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    GlobalSourceMac interface{}

    // EVPN ES Peering Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    PeeringTime interface{}

    // EVPN ES Recovery Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    RecoveryTime interface{}

    // EVPN ES Carving Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    CarvingTime interface{}

    // Number of moves within the move interval before locking the MAC. The type
    // is interface{} with range: 0..4294967295.
    MacSecureMoveCount interface{}

    // Interval to watch for subsequent mac moves before locking the MAC. The type
    // is interface{} with range: 0..4294967295.
    MacSecureMoveInterval interface{}

    // Length of time to lock the mac after a MAC security violation. The type is
    // interface{} with range: 0..4294967295.
    MacSecureFreezeTime interface{}

    // Number of times to retry after a MAC un-freezes. The type is interface{}
    // with range: 0..4294967295.
    MacSecureRetryCount interface{}

    // EVPN Node Cost-out. The type is bool.
    CostOut interface{}

    // EVPN Node startup cost-in Time (minutes). The type is interface{} with
    // range: 0..4294967295. Units are minute.
    StartupCostInTime interface{}

    // Send to L2RIB Throttled. The type is bool.
    L2ribThrottle interface{}

    // Logging EVPN Designated Forwarder changes enabled. The type is bool.
    LoggingDfElectionEnabled interface{}
}

func (summary *Evpn_Active_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "active"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", summary.RouterId})
    summary.EntityData.Leafs.Append("as", types.YLeaf{"As", summary.As})
    summary.EntityData.Leafs.Append("ev-is", types.YLeaf{"EvIs", summary.EvIs})
    summary.EntityData.Leafs.Append("tunnel-endpoints", types.YLeaf{"TunnelEndpoints", summary.TunnelEndpoints})
    summary.EntityData.Leafs.Append("local-mac-routes", types.YLeaf{"LocalMacRoutes", summary.LocalMacRoutes})
    summary.EntityData.Leafs.Append("local-ipv4-mac-routes", types.YLeaf{"LocalIpv4MacRoutes", summary.LocalIpv4MacRoutes})
    summary.EntityData.Leafs.Append("local-ipv6-mac-routes", types.YLeaf{"LocalIpv6MacRoutes", summary.LocalIpv6MacRoutes})
    summary.EntityData.Leafs.Append("es-global-mac-routes", types.YLeaf{"EsGlobalMacRoutes", summary.EsGlobalMacRoutes})
    summary.EntityData.Leafs.Append("remote-mac-routes", types.YLeaf{"RemoteMacRoutes", summary.RemoteMacRoutes})
    summary.EntityData.Leafs.Append("remote-soo-mac-routes", types.YLeaf{"RemoteSooMacRoutes", summary.RemoteSooMacRoutes})
    summary.EntityData.Leafs.Append("remote-ipv4-mac-routes", types.YLeaf{"RemoteIpv4MacRoutes", summary.RemoteIpv4MacRoutes})
    summary.EntityData.Leafs.Append("remote-ipv6-mac-routes", types.YLeaf{"RemoteIpv6MacRoutes", summary.RemoteIpv6MacRoutes})
    summary.EntityData.Leafs.Append("local-imcast-routes", types.YLeaf{"LocalImcastRoutes", summary.LocalImcastRoutes})
    summary.EntityData.Leafs.Append("remote-imcast-routes", types.YLeaf{"RemoteImcastRoutes", summary.RemoteImcastRoutes})
    summary.EntityData.Leafs.Append("labels", types.YLeaf{"Labels", summary.Labels})
    summary.EntityData.Leafs.Append("es-entries", types.YLeaf{"EsEntries", summary.EsEntries})
    summary.EntityData.Leafs.Append("neighbor-entries", types.YLeaf{"NeighborEntries", summary.NeighborEntries})
    summary.EntityData.Leafs.Append("local-ead-routes", types.YLeaf{"LocalEadRoutes", summary.LocalEadRoutes})
    summary.EntityData.Leafs.Append("remote-ead-routes", types.YLeaf{"RemoteEadRoutes", summary.RemoteEadRoutes})
    summary.EntityData.Leafs.Append("global-source-mac", types.YLeaf{"GlobalSourceMac", summary.GlobalSourceMac})
    summary.EntityData.Leafs.Append("peering-time", types.YLeaf{"PeeringTime", summary.PeeringTime})
    summary.EntityData.Leafs.Append("recovery-time", types.YLeaf{"RecoveryTime", summary.RecoveryTime})
    summary.EntityData.Leafs.Append("carving-time", types.YLeaf{"CarvingTime", summary.CarvingTime})
    summary.EntityData.Leafs.Append("mac-secure-move-count", types.YLeaf{"MacSecureMoveCount", summary.MacSecureMoveCount})
    summary.EntityData.Leafs.Append("mac-secure-move-interval", types.YLeaf{"MacSecureMoveInterval", summary.MacSecureMoveInterval})
    summary.EntityData.Leafs.Append("mac-secure-freeze-time", types.YLeaf{"MacSecureFreezeTime", summary.MacSecureFreezeTime})
    summary.EntityData.Leafs.Append("mac-secure-retry-count", types.YLeaf{"MacSecureRetryCount", summary.MacSecureRetryCount})
    summary.EntityData.Leafs.Append("cost-out", types.YLeaf{"CostOut", summary.CostOut})
    summary.EntityData.Leafs.Append("startup-cost-in-time", types.YLeaf{"StartupCostInTime", summary.StartupCostInTime})
    summary.EntityData.Leafs.Append("l2rib-throttle", types.YLeaf{"L2ribThrottle", summary.L2ribThrottle})
    summary.EntityData.Leafs.Append("logging-df-election-enabled", types.YLeaf{"LoggingDfElectionEnabled", summary.LoggingDfElectionEnabled})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Evpn_Active_EviDetail
// L2VPN EVI Detail Table
type Evpn_Active_EviDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info Elements.
    Elements Evpn_Active_EviDetail_Elements

    // Container for all EVI detail info.
    EviChildren Evpn_Active_EviDetail_EviChildren
}

func (eviDetail *Evpn_Active_EviDetail) GetEntityData() *types.CommonEntityData {
    eviDetail.EntityData.YFilter = eviDetail.YFilter
    eviDetail.EntityData.YangName = "evi-detail"
    eviDetail.EntityData.BundleName = "cisco_ios_xr"
    eviDetail.EntityData.ParentYangName = "active"
    eviDetail.EntityData.SegmentPath = "evi-detail"
    eviDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviDetail.EntityData.Children = types.NewOrderedMap()
    eviDetail.EntityData.Children.Append("elements", types.YChild{"Elements", &eviDetail.Elements})
    eviDetail.EntityData.Children.Append("evi-children", types.YChild{"EviChildren", &eviDetail.EviChildren})
    eviDetail.EntityData.Leafs = types.NewOrderedMap()

    eviDetail.EntityData.YListKeys = []string {}

    return &(eviDetail.EntityData)
}

// Evpn_Active_EviDetail_Elements
// EVI BGP RT Detail Info Elements
type Evpn_Active_EviDetail_Elements struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info. The type is slice of
    // Evpn_Active_EviDetail_Elements_Element.
    Element []*Evpn_Active_EviDetail_Elements_Element
}

func (elements *Evpn_Active_EviDetail_Elements) GetEntityData() *types.CommonEntityData {
    elements.EntityData.YFilter = elements.YFilter
    elements.EntityData.YangName = "elements"
    elements.EntityData.BundleName = "cisco_ios_xr"
    elements.EntityData.ParentYangName = "evi-detail"
    elements.EntityData.SegmentPath = "elements"
    elements.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    elements.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    elements.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    elements.EntityData.Children = types.NewOrderedMap()
    elements.EntityData.Children.Append("element", types.YChild{"Element", nil})
    for i := range elements.Element {
        elements.EntityData.Children.Append(types.GetSegmentPath(elements.Element[i]), types.YChild{"Element", elements.Element[i]})
    }
    elements.EntityData.Leafs = types.NewOrderedMap()

    elements.EntityData.YListKeys = []string {}

    return &(elements.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element
// EVI BGP RT Detail Info
type Evpn_Active_EviDetail_Elements_Element struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // EVPN Instance encapsulation. The type is interface{} with range: 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // EVI description. The type is string.
    Description interface{}

    // Unicast Label. The type is interface{} with range: 0..4294967295.
    UnicastLabel interface{}

    // Multicast Label. The type is interface{} with range: 0..4294967295.
    MulticastLabel interface{}

    // Control-Word Disable. The type is bool.
    CwDisable interface{}

    // Table-policy Name. The type is string.
    TablePolicyName interface{}

    // Forward Class attribute. The type is interface{} with range: 0..255.
    ForwardClass interface{}

    // Is Import RT None set. The type is bool.
    RtImportBlockSet interface{}

    // Is Export RT None set. The type is bool.
    RtExportBlockSet interface{}

    // Advertise MAC-only routes on this EVI. The type is bool.
    AdvertiseMac interface{}

    // Advertise BVI MACs routes on this EVI. The type is bool.
    AdvertiseBviMac interface{}

    // Route Aliasing is disabled. The type is bool.
    AliasingDisabled interface{}

    // Unknown-unicast flooding is disabled. The type is bool.
    UnknownUnicastFloodingDisabled interface{}

    // Route Re-origination is disabled. The type is bool.
    ReoriginateDisabled interface{}

    // EVPN Instance is Regular/Stitching side. The type is bool.
    Stitching interface{}

    // EVI is connected to multicast source. The type is bool.
    MulticastSourceConnected interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Active_EviDetail_Elements_Element_EvpnInstance

    // Flow Label Information.
    FlowLabel Evpn_Active_EviDetail_Elements_Element_FlowLabel

    // Automatic Route Distingtuisher.
    RdAuto Evpn_Active_EviDetail_Elements_Element_RdAuto

    // Configured Route Distinguisher.
    RdConfigured Evpn_Active_EviDetail_Elements_Element_RdConfigured

    // Automatic Route Target.
    RtAuto Evpn_Active_EviDetail_Elements_Element_RtAuto
}

func (element *Evpn_Active_EviDetail_Elements_Element) GetEntityData() *types.CommonEntityData {
    element.EntityData.YFilter = element.YFilter
    element.EntityData.YangName = "element"
    element.EntityData.BundleName = "cisco_ios_xr"
    element.EntityData.ParentYangName = "elements"
    element.EntityData.SegmentPath = "element"
    element.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    element.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    element.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    element.EntityData.Children = types.NewOrderedMap()
    element.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &element.EvpnInstance})
    element.EntityData.Children.Append("flow-label", types.YChild{"FlowLabel", &element.FlowLabel})
    element.EntityData.Children.Append("rd-auto", types.YChild{"RdAuto", &element.RdAuto})
    element.EntityData.Children.Append("rd-configured", types.YChild{"RdConfigured", &element.RdConfigured})
    element.EntityData.Children.Append("rt-auto", types.YChild{"RtAuto", &element.RtAuto})
    element.EntityData.Leafs = types.NewOrderedMap()
    element.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", element.Evi})
    element.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", element.Encapsulation})
    element.EntityData.Leafs.Append("evi-xr", types.YLeaf{"EviXr", element.EviXr})
    element.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", element.EncapsulationXr})
    element.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", element.BdName})
    element.EntityData.Leafs.Append("type", types.YLeaf{"Type", element.Type})
    element.EntityData.Leafs.Append("description", types.YLeaf{"Description", element.Description})
    element.EntityData.Leafs.Append("unicast-label", types.YLeaf{"UnicastLabel", element.UnicastLabel})
    element.EntityData.Leafs.Append("multicast-label", types.YLeaf{"MulticastLabel", element.MulticastLabel})
    element.EntityData.Leafs.Append("cw-disable", types.YLeaf{"CwDisable", element.CwDisable})
    element.EntityData.Leafs.Append("table-policy-name", types.YLeaf{"TablePolicyName", element.TablePolicyName})
    element.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", element.ForwardClass})
    element.EntityData.Leafs.Append("rt-import-block-set", types.YLeaf{"RtImportBlockSet", element.RtImportBlockSet})
    element.EntityData.Leafs.Append("rt-export-block-set", types.YLeaf{"RtExportBlockSet", element.RtExportBlockSet})
    element.EntityData.Leafs.Append("advertise-mac", types.YLeaf{"AdvertiseMac", element.AdvertiseMac})
    element.EntityData.Leafs.Append("advertise-bvi-mac", types.YLeaf{"AdvertiseBviMac", element.AdvertiseBviMac})
    element.EntityData.Leafs.Append("aliasing-disabled", types.YLeaf{"AliasingDisabled", element.AliasingDisabled})
    element.EntityData.Leafs.Append("unknown-unicast-flooding-disabled", types.YLeaf{"UnknownUnicastFloodingDisabled", element.UnknownUnicastFloodingDisabled})
    element.EntityData.Leafs.Append("reoriginate-disabled", types.YLeaf{"ReoriginateDisabled", element.ReoriginateDisabled})
    element.EntityData.Leafs.Append("stitching", types.YLeaf{"Stitching", element.Stitching})
    element.EntityData.Leafs.Append("multicast-source-connected", types.YLeaf{"MulticastSourceConnected", element.MulticastSourceConnected})

    element.EntityData.YListKeys = []string {}

    return &(element.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_EvpnInstance
// EVPN Instance summary information
type Evpn_Active_EviDetail_Elements_Element_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Active_EviDetail_Elements_Element_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "element"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_FlowLabel
// Flow Label Information
type Evpn_Active_EviDetail_Elements_Element_FlowLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static flow label. The type is bool.
    StaticFlowLabel interface{}

    // Globally configured flow label. The type is bool.
    GlobalFlowLabel interface{}
}

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetEntityData() *types.CommonEntityData {
    flowLabel.EntityData.YFilter = flowLabel.YFilter
    flowLabel.EntityData.YangName = "flow-label"
    flowLabel.EntityData.BundleName = "cisco_ios_xr"
    flowLabel.EntityData.ParentYangName = "element"
    flowLabel.EntityData.SegmentPath = "flow-label"
    flowLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabel.EntityData.Children = types.NewOrderedMap()
    flowLabel.EntityData.Leafs = types.NewOrderedMap()
    flowLabel.EntityData.Leafs.Append("static-flow-label", types.YLeaf{"StaticFlowLabel", flowLabel.StaticFlowLabel})
    flowLabel.EntityData.Leafs.Append("global-flow-label", types.YLeaf{"GlobalFlowLabel", flowLabel.GlobalFlowLabel})

    flowLabel.EntityData.YListKeys = []string {}

    return &(flowLabel.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdAuto
// Automatic Route Distingtuisher
type Evpn_Active_EviDetail_Elements_Element_RdAuto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RD. The type is L2vpnAdRd.
    Rd interface{}

    // auto.
    Auto Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto

    // two byte as.
    TwoByteAs Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs

    // v4 addr.
    V4Addr Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr
}

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetEntityData() *types.CommonEntityData {
    rdAuto.EntityData.YFilter = rdAuto.YFilter
    rdAuto.EntityData.YangName = "rd-auto"
    rdAuto.EntityData.BundleName = "cisco_ios_xr"
    rdAuto.EntityData.ParentYangName = "element"
    rdAuto.EntityData.SegmentPath = "rd-auto"
    rdAuto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdAuto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdAuto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdAuto.EntityData.Children = types.NewOrderedMap()
    rdAuto.EntityData.Children.Append("auto", types.YChild{"Auto", &rdAuto.Auto})
    rdAuto.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &rdAuto.TwoByteAs})
    rdAuto.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &rdAuto.FourByteAs})
    rdAuto.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &rdAuto.V4Addr})
    rdAuto.EntityData.Leafs = types.NewOrderedMap()
    rdAuto.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", rdAuto.Rd})

    rdAuto.EntityData.YListKeys = []string {}

    return &(rdAuto.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto
// auto
type Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetEntityData() *types.CommonEntityData {
    auto.EntityData.YFilter = auto.YFilter
    auto.EntityData.YangName = "auto"
    auto.EntityData.BundleName = "cisco_ios_xr"
    auto.EntityData.ParentYangName = "rd-auto"
    auto.EntityData.SegmentPath = "auto"
    auto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auto.EntityData.Children = types.NewOrderedMap()
    auto.EntityData.Leafs = types.NewOrderedMap()
    auto.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", auto.RouterId})
    auto.EntityData.Leafs.Append("auto-index", types.YLeaf{"AutoIndex", auto.AutoIndex})

    auto.EntityData.YListKeys = []string {}

    return &(auto.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rd-auto"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs
// four byte as
type Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rd-auto"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rd-auto"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdConfigured
// Configured Route Distinguisher
type Evpn_Active_EviDetail_Elements_Element_RdConfigured struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RD. The type is L2vpnAdRd.
    Rd interface{}

    // auto.
    Auto Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto

    // two byte as.
    TwoByteAs Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs

    // v4 addr.
    V4Addr Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr
}

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetEntityData() *types.CommonEntityData {
    rdConfigured.EntityData.YFilter = rdConfigured.YFilter
    rdConfigured.EntityData.YangName = "rd-configured"
    rdConfigured.EntityData.BundleName = "cisco_ios_xr"
    rdConfigured.EntityData.ParentYangName = "element"
    rdConfigured.EntityData.SegmentPath = "rd-configured"
    rdConfigured.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdConfigured.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdConfigured.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdConfigured.EntityData.Children = types.NewOrderedMap()
    rdConfigured.EntityData.Children.Append("auto", types.YChild{"Auto", &rdConfigured.Auto})
    rdConfigured.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &rdConfigured.TwoByteAs})
    rdConfigured.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &rdConfigured.FourByteAs})
    rdConfigured.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &rdConfigured.V4Addr})
    rdConfigured.EntityData.Leafs = types.NewOrderedMap()
    rdConfigured.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", rdConfigured.Rd})

    rdConfigured.EntityData.YListKeys = []string {}

    return &(rdConfigured.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto
// auto
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetEntityData() *types.CommonEntityData {
    auto.EntityData.YFilter = auto.YFilter
    auto.EntityData.YangName = "auto"
    auto.EntityData.BundleName = "cisco_ios_xr"
    auto.EntityData.ParentYangName = "rd-configured"
    auto.EntityData.SegmentPath = "auto"
    auto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auto.EntityData.Children = types.NewOrderedMap()
    auto.EntityData.Leafs = types.NewOrderedMap()
    auto.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", auto.RouterId})
    auto.EntityData.Leafs.Append("auto-index", types.YLeaf{"AutoIndex", auto.AutoIndex})

    auto.EntityData.YListKeys = []string {}

    return &(auto.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rd-configured"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs
// four byte as
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rd-configured"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rd-configured"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAuto
// Automatic Route Target
type Evpn_Active_EviDetail_Elements_Element_RtAuto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs

    // v4 addr.
    V4Addr Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr

    // es import.
    EsImport Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport
}

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetEntityData() *types.CommonEntityData {
    rtAuto.EntityData.YFilter = rtAuto.YFilter
    rtAuto.EntityData.YangName = "rt-auto"
    rtAuto.EntityData.BundleName = "cisco_ios_xr"
    rtAuto.EntityData.ParentYangName = "element"
    rtAuto.EntityData.SegmentPath = "rt-auto"
    rtAuto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rtAuto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rtAuto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rtAuto.EntityData.Children = types.NewOrderedMap()
    rtAuto.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &rtAuto.TwoByteAs})
    rtAuto.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &rtAuto.FourByteAs})
    rtAuto.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &rtAuto.V4Addr})
    rtAuto.EntityData.Children.Append("es-import", types.YChild{"EsImport", &rtAuto.EsImport})
    rtAuto.EntityData.Leafs = types.NewOrderedMap()
    rtAuto.EntityData.Leafs.Append("rt", types.YLeaf{"Rt", rtAuto.Rt})

    rtAuto.EntityData.YListKeys = []string {}

    return &(rtAuto.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rt-auto"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs
// four byte as
type Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rt-auto"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rt-auto"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport
// es import
type Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "rt-auto"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = types.NewOrderedMap()
    esImport.EntityData.Leafs = types.NewOrderedMap()
    esImport.EntityData.Leafs.Append("high-bytes", types.YLeaf{"HighBytes", esImport.HighBytes})
    esImport.EntityData.Leafs.Append("low-bytes", types.YLeaf{"LowBytes", esImport.LowBytes})

    esImport.EntityData.YListKeys = []string {}

    return &(esImport.EntityData)
}

// Evpn_Active_EviDetail_EviChildren
// Container for all EVI detail info
type Evpn_Active_EviDetail_EviChildren struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Neighbor table.
    Neighbors Evpn_Active_EviDetail_EviChildren_Neighbors

    // EVPN Ethernet Auto-Discovery table.
    EthernetAutoDiscoveries Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries

    // L2VPN EVPN IMCAST table.
    InclusiveMulticasts Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts

    // L2VPN EVPN EVI RT Child Table.
    RouteTargets Evpn_Active_EviDetail_EviChildren_RouteTargets

    // L2VPN EVPN EVI MAC table.
    Macs Evpn_Active_EviDetail_EviChildren_Macs
}

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetEntityData() *types.CommonEntityData {
    eviChildren.EntityData.YFilter = eviChildren.YFilter
    eviChildren.EntityData.YangName = "evi-children"
    eviChildren.EntityData.BundleName = "cisco_ios_xr"
    eviChildren.EntityData.ParentYangName = "evi-detail"
    eviChildren.EntityData.SegmentPath = "evi-children"
    eviChildren.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviChildren.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviChildren.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviChildren.EntityData.Children = types.NewOrderedMap()
    eviChildren.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &eviChildren.Neighbors})
    eviChildren.EntityData.Children.Append("ethernet-auto-discoveries", types.YChild{"EthernetAutoDiscoveries", &eviChildren.EthernetAutoDiscoveries})
    eviChildren.EntityData.Children.Append("inclusive-multicasts", types.YChild{"InclusiveMulticasts", &eviChildren.InclusiveMulticasts})
    eviChildren.EntityData.Children.Append("route-targets", types.YChild{"RouteTargets", &eviChildren.RouteTargets})
    eviChildren.EntityData.Children.Append("macs", types.YChild{"Macs", &eviChildren.Macs})
    eviChildren.EntityData.Leafs = types.NewOrderedMap()

    eviChildren.EntityData.YListKeys = []string {}

    return &(eviChildren.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Neighbors
// EVPN Neighbor table
type Evpn_Active_EviDetail_EviChildren_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Neighbor table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor.
    Neighbor []*Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor
}

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "evi-children"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor
// EVPN Neighbor table
type Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Neighbor IP. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborIp interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance
}

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &neighbor.EvpnInstance})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", neighbor.Evi})
    neighbor.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", neighbor.Encapsulation})
    neighbor.EntityData.Leafs.Append("neighbor-ip", types.YLeaf{"NeighborIp", neighbor.NeighborIp})
    neighbor.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", neighbor.Neighbor})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance
// EVPN Instance summary information
type Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "neighbor"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries
// EVPN Ethernet Auto-Discovery table
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet Auto-Discovery Entry. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery.
    EthernetAutoDiscovery []*Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
}

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetEntityData() *types.CommonEntityData {
    ethernetAutoDiscoveries.EntityData.YFilter = ethernetAutoDiscoveries.YFilter
    ethernetAutoDiscoveries.EntityData.YangName = "ethernet-auto-discoveries"
    ethernetAutoDiscoveries.EntityData.BundleName = "cisco_ios_xr"
    ethernetAutoDiscoveries.EntityData.ParentYangName = "evi-children"
    ethernetAutoDiscoveries.EntityData.SegmentPath = "ethernet-auto-discoveries"
    ethernetAutoDiscoveries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetAutoDiscoveries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetAutoDiscoveries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetAutoDiscoveries.EntityData.Children = types.NewOrderedMap()
    ethernetAutoDiscoveries.EntityData.Children.Append("ethernet-auto-discovery", types.YChild{"EthernetAutoDiscovery", nil})
    for i := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
        ethernetAutoDiscoveries.EntityData.Children.Append(types.GetSegmentPath(ethernetAutoDiscoveries.EthernetAutoDiscovery[i]), types.YChild{"EthernetAutoDiscovery", ethernetAutoDiscoveries.EthernetAutoDiscovery[i]})
    }
    ethernetAutoDiscoveries.EntityData.Leafs = types.NewOrderedMap()

    ethernetAutoDiscoveries.EntityData.YListKeys = []string {}

    return &(ethernetAutoDiscoveries.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
// EVPN Ethernet Auto-Discovery Entry
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Local nexthop IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalNextHop interface{}

    // Associated local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Indication of EthernetAutoDiscovery Route is local. The type is bool.
    IsLocalEad interface{}

    // Single-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleActive interface{}

    // Single-flow-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleFlowActive interface{}

    // Number of items in path list buffer. The type is interface{} with range:
    // 0..4294967295.
    NumPaths interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance

    // Ethernet Segment id. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer.
    PathBuffer []*Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
}

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetEntityData() *types.CommonEntityData {
    ethernetAutoDiscovery.EntityData.YFilter = ethernetAutoDiscovery.YFilter
    ethernetAutoDiscovery.EntityData.YangName = "ethernet-auto-discovery"
    ethernetAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    ethernetAutoDiscovery.EntityData.ParentYangName = "ethernet-auto-discoveries"
    ethernetAutoDiscovery.EntityData.SegmentPath = "ethernet-auto-discovery"
    ethernetAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetAutoDiscovery.EntityData.Children = types.NewOrderedMap()
    ethernetAutoDiscovery.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &ethernetAutoDiscovery.EvpnInstance})
    ethernetAutoDiscovery.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
        ethernetAutoDiscovery.EntityData.Children.Append(types.GetSegmentPath(ethernetAutoDiscovery.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", ethernetAutoDiscovery.EthernetSegmentIdentifier[i]})
    }
    ethernetAutoDiscovery.EntityData.Children.Append("path-buffer", types.YChild{"PathBuffer", nil})
    for i := range ethernetAutoDiscovery.PathBuffer {
        ethernetAutoDiscovery.EntityData.Children.Append(types.GetSegmentPath(ethernetAutoDiscovery.PathBuffer[i]), types.YChild{"PathBuffer", ethernetAutoDiscovery.PathBuffer[i]})
    }
    ethernetAutoDiscovery.EntityData.Leafs = types.NewOrderedMap()
    ethernetAutoDiscovery.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", ethernetAutoDiscovery.Evi})
    ethernetAutoDiscovery.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", ethernetAutoDiscovery.Encapsulation})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", ethernetAutoDiscovery.Esi1})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", ethernetAutoDiscovery.Esi2})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", ethernetAutoDiscovery.Esi3})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", ethernetAutoDiscovery.Esi4})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", ethernetAutoDiscovery.Esi5})
    ethernetAutoDiscovery.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", ethernetAutoDiscovery.EthernetTag})
    ethernetAutoDiscovery.EntityData.Leafs.Append("ethernet-tag-xr", types.YLeaf{"EthernetTagXr", ethernetAutoDiscovery.EthernetTagXr})
    ethernetAutoDiscovery.EntityData.Leafs.Append("local-next-hop", types.YLeaf{"LocalNextHop", ethernetAutoDiscovery.LocalNextHop})
    ethernetAutoDiscovery.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", ethernetAutoDiscovery.LocalLabel})
    ethernetAutoDiscovery.EntityData.Leafs.Append("is-local-ead", types.YLeaf{"IsLocalEad", ethernetAutoDiscovery.IsLocalEad})
    ethernetAutoDiscovery.EntityData.Leafs.Append("redundancy-single-active", types.YLeaf{"RedundancySingleActive", ethernetAutoDiscovery.RedundancySingleActive})
    ethernetAutoDiscovery.EntityData.Leafs.Append("redundancy-single-flow-active", types.YLeaf{"RedundancySingleFlowActive", ethernetAutoDiscovery.RedundancySingleFlowActive})
    ethernetAutoDiscovery.EntityData.Leafs.Append("num-paths", types.YLeaf{"NumPaths", ethernetAutoDiscovery.NumPaths})

    ethernetAutoDiscovery.EntityData.YListKeys = []string {}

    return &(ethernetAutoDiscovery.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance
// EVPN Instance summary information
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "ethernet-auto-discovery"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "ethernet-auto-discovery"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
// Path List Buffer
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetEntityData() *types.CommonEntityData {
    pathBuffer.EntityData.YFilter = pathBuffer.YFilter
    pathBuffer.EntityData.YangName = "path-buffer"
    pathBuffer.EntityData.BundleName = "cisco_ios_xr"
    pathBuffer.EntityData.ParentYangName = "ethernet-auto-discovery"
    pathBuffer.EntityData.SegmentPath = "path-buffer"
    pathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathBuffer.EntityData.Children = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", pathBuffer.TunnelEndpointId})
    pathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", pathBuffer.NextHop})
    pathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", pathBuffer.OutputLabel})
    pathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel})

    pathBuffer.EntityData.YListKeys = []string {}

    return &(pathBuffer.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts
// L2VPN EVPN IMCAST table
type Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN IMCAST table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast.
    InclusiveMulticast []*Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
}

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetEntityData() *types.CommonEntityData {
    inclusiveMulticasts.EntityData.YFilter = inclusiveMulticasts.YFilter
    inclusiveMulticasts.EntityData.YangName = "inclusive-multicasts"
    inclusiveMulticasts.EntityData.BundleName = "cisco_ios_xr"
    inclusiveMulticasts.EntityData.ParentYangName = "evi-children"
    inclusiveMulticasts.EntityData.SegmentPath = "inclusive-multicasts"
    inclusiveMulticasts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inclusiveMulticasts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inclusiveMulticasts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inclusiveMulticasts.EntityData.Children = types.NewOrderedMap()
    inclusiveMulticasts.EntityData.Children.Append("inclusive-multicast", types.YChild{"InclusiveMulticast", nil})
    for i := range inclusiveMulticasts.InclusiveMulticast {
        inclusiveMulticasts.EntityData.Children.Append(types.GetSegmentPath(inclusiveMulticasts.InclusiveMulticast[i]), types.YChild{"InclusiveMulticast", inclusiveMulticasts.InclusiveMulticast[i]})
    }
    inclusiveMulticasts.EntityData.Leafs = types.NewOrderedMap()

    inclusiveMulticasts.EntityData.YListKeys = []string {}

    return &(inclusiveMulticasts.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
// L2VPN EVPN IMCAST table
type Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // Originating IP. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIp interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Originating IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIpXr interface{}

    // Tunnel Endpoint ID. The type is interface{} with range: 0..4294967295.
    TunnelEndpointId interface{}

    // IP of nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Local entry. The type is bool.
    IsLocalEntry interface{}

    // Proxy entry. The type is bool.
    IsProxyEntry interface{}

    // SR-TE Policy. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SrtePolicy interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance
}

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetEntityData() *types.CommonEntityData {
    inclusiveMulticast.EntityData.YFilter = inclusiveMulticast.YFilter
    inclusiveMulticast.EntityData.YangName = "inclusive-multicast"
    inclusiveMulticast.EntityData.BundleName = "cisco_ios_xr"
    inclusiveMulticast.EntityData.ParentYangName = "inclusive-multicasts"
    inclusiveMulticast.EntityData.SegmentPath = "inclusive-multicast"
    inclusiveMulticast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inclusiveMulticast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inclusiveMulticast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inclusiveMulticast.EntityData.Children = types.NewOrderedMap()
    inclusiveMulticast.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &inclusiveMulticast.EvpnInstance})
    inclusiveMulticast.EntityData.Leafs = types.NewOrderedMap()
    inclusiveMulticast.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", inclusiveMulticast.Evi})
    inclusiveMulticast.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", inclusiveMulticast.Encapsulation})
    inclusiveMulticast.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", inclusiveMulticast.EthernetTag})
    inclusiveMulticast.EntityData.Leafs.Append("originating-ip", types.YLeaf{"OriginatingIp", inclusiveMulticast.OriginatingIp})
    inclusiveMulticast.EntityData.Leafs.Append("ethernet-tag-xr", types.YLeaf{"EthernetTagXr", inclusiveMulticast.EthernetTagXr})
    inclusiveMulticast.EntityData.Leafs.Append("originating-ip-xr", types.YLeaf{"OriginatingIpXr", inclusiveMulticast.OriginatingIpXr})
    inclusiveMulticast.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", inclusiveMulticast.TunnelEndpointId})
    inclusiveMulticast.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", inclusiveMulticast.NextHop})
    inclusiveMulticast.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", inclusiveMulticast.OutputLabel})
    inclusiveMulticast.EntityData.Leafs.Append("is-local-entry", types.YLeaf{"IsLocalEntry", inclusiveMulticast.IsLocalEntry})
    inclusiveMulticast.EntityData.Leafs.Append("is-proxy-entry", types.YLeaf{"IsProxyEntry", inclusiveMulticast.IsProxyEntry})
    inclusiveMulticast.EntityData.Leafs.Append("srte-policy", types.YLeaf{"SrtePolicy", inclusiveMulticast.SrtePolicy})

    inclusiveMulticast.EntityData.YListKeys = []string {}

    return &(inclusiveMulticast.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance
// EVPN Instance summary information
type Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "inclusive-multicast"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets
// L2VPN EVPN EVI RT Child Table
type Evpn_Active_EviDetail_EviChildren_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI RT Table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget.
    RouteTarget []*Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget
}

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "evi-children"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = types.NewOrderedMap()
    routeTargets.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children.Append(types.GetSegmentPath(routeTargets.RouteTarget[i]), types.YChild{"RouteTarget", routeTargets.RouteTarget[i]})
    }
    routeTargets.EntityData.Leafs = types.NewOrderedMap()

    routeTargets.EntityData.YListKeys = []string {}

    return &(routeTargets.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget
// L2VPN EVPN EVI RT Table
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Role of the route target. The type is BgpRouteTargetRole.
    Role interface{}

    // Format of the route target. The type is BgpRouteTargetFormat.
    Format interface{}

    // Two or Four byte AS Number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // RT AS Index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // RT IP Index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}

    // RT IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // RT Role. The type is L2vpnAdRtRole.
    RouteTargetRole interface{}

    // RT Stitching. The type is bool.
    RouteTargetStitching interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance

    // Route Target.
    RouteTarget Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target"
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &routeTarget.EvpnInstance})
    routeTarget.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", &routeTarget.RouteTarget})
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", routeTarget.Evi})
    routeTarget.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", routeTarget.Encapsulation})
    routeTarget.EntityData.Leafs.Append("role", types.YLeaf{"Role", routeTarget.Role})
    routeTarget.EntityData.Leafs.Append("format", types.YLeaf{"Format", routeTarget.Format})
    routeTarget.EntityData.Leafs.Append("as", types.YLeaf{"As", routeTarget.As})
    routeTarget.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", routeTarget.AsIndex})
    routeTarget.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", routeTarget.AddrIndex})
    routeTarget.EntityData.Leafs.Append("address", types.YLeaf{"Address", routeTarget.Address})
    routeTarget.EntityData.Leafs.Append("route-target-role", types.YLeaf{"RouteTargetRole", routeTarget.RouteTargetRole})
    routeTarget.EntityData.Leafs.Append("route-target-stitching", types.YLeaf{"RouteTargetStitching", routeTarget.RouteTargetStitching})

    routeTarget.EntityData.YListKeys = []string {}

    return &(routeTarget.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance
// EVPN Instance summary information
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "route-target"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
// Route Target
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs

    // v4 addr.
    V4Addr Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr

    // es import.
    EsImport Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-target"
    routeTarget.EntityData.SegmentPath = "route-target"
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &routeTarget.TwoByteAs})
    routeTarget.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &routeTarget.FourByteAs})
    routeTarget.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &routeTarget.V4Addr})
    routeTarget.EntityData.Children.Append("es-import", types.YChild{"EsImport", &routeTarget.EsImport})
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("rt", types.YLeaf{"Rt", routeTarget.Rt})

    routeTarget.EntityData.YListKeys = []string {}

    return &(routeTarget.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "route-target"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs
// four byte as
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "route-target"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr
// v4 addr
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "route-target"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport
// es import
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "route-target"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = types.NewOrderedMap()
    esImport.EntityData.Leafs = types.NewOrderedMap()
    esImport.EntityData.Leafs.Append("high-bytes", types.YLeaf{"HighBytes", esImport.HighBytes})
    esImport.EntityData.Leafs.Append("low-bytes", types.YLeaf{"LowBytes", esImport.LowBytes})

    esImport.EntityData.YListKeys = []string {}

    return &(esImport.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Macs
// L2VPN EVPN EVI MAC table
type Evpn_Active_EviDetail_EviChildren_Macs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN MAC table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Macs_Mac.
    Mac []*Evpn_Active_EviDetail_EviChildren_Macs_Mac
}

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetEntityData() *types.CommonEntityData {
    macs.EntityData.YFilter = macs.YFilter
    macs.EntityData.YangName = "macs"
    macs.EntityData.BundleName = "cisco_ios_xr"
    macs.EntityData.ParentYangName = "evi-children"
    macs.EntityData.SegmentPath = "macs"
    macs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macs.EntityData.Children = types.NewOrderedMap()
    macs.EntityData.Children.Append("mac", types.YChild{"Mac", nil})
    for i := range macs.Mac {
        macs.EntityData.Children.Append(types.GetSegmentPath(macs.Mac[i]), types.YChild{"Mac", macs.Mac[i]})
    }
    macs.EntityData.Leafs = types.NewOrderedMap()

    macs.EntityData.YListKeys = []string {}

    return &(macs.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Macs_Mac
// L2VPN EVPN MAC table
type Evpn_Active_EviDetail_EviChildren_Macs_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // IP Address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddressXr interface{}

    // IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddressXr interface{}

    // Associated local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Number of items in path list buffer. The type is interface{} with range:
    // 0..4294967295.
    NumPaths interface{}

    // Indication of MAC being locally generated. The type is bool.
    IsLocalMac interface{}

    // Proxy entry. The type is bool.
    IsProxyEntry interface{}

    // Indication of MAC being remotely generated. The type is bool.
    IsRemoteMac interface{}

    // SOO nexthop (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SooNexthop interface{}

    // IP nexthop address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpnhAddress interface{}

    // ESI port key. The type is interface{} with range: 0..65535.
    EsiPortKey interface{}

    // Encap type of local MAC. The type is interface{} with range: 0..255.
    LocalEncapType interface{}

    // Encap type of remote MAC. The type is interface{} with range: 0..255.
    RemoteEncapType interface{}

    // Port the MAC was learned on. The type is string.
    LearnedBridgePortName interface{}

    // local seq id. The type is interface{} with range: 0..4294967295.
    LocalSeqId interface{}

    // remote seq id. The type is interface{} with range: 0..4294967295.
    RemoteSeqId interface{}

    // local l3 label. The type is interface{} with range: 0..4294967295.
    LocalL3Label interface{}

    // Router MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    RouterMacAddress interface{}

    // Number of flushes requested . The type is interface{} with range: 0..65535.
    MacFlushRequested interface{}

    // Number of flushes received . The type is interface{} with range: 0..65535.
    MacFlushReceived interface{}

    // MPLS Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Internal Label has resolved per-ES EAD and per-EVI EAD or MAC routes. The
    // type is bool.
    Resolved interface{}

    // Indication if Local MAC is statically configured. The type is bool.
    LocalIsStatic interface{}

    // Indication if Remote MAC is statically configured. The type is bool.
    RemoteIsStatic interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Active_EviDetail_EviChildren_Macs_Mac_EvpnInstance

    // Local Ethernet Segment id. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier.
    LocalEthernetSegmentIdentifier []*Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier

    // Remote Ethernet Segment id. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier.
    RemoteEthernetSegmentIdentifier []*Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer.
    PathBuffer []*Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer
}

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "macs"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &mac.EvpnInstance})
    mac.EntityData.Children.Append("local-ethernet-segment-identifier", types.YChild{"LocalEthernetSegmentIdentifier", nil})
    for i := range mac.LocalEthernetSegmentIdentifier {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.LocalEthernetSegmentIdentifier[i]), types.YChild{"LocalEthernetSegmentIdentifier", mac.LocalEthernetSegmentIdentifier[i]})
    }
    mac.EntityData.Children.Append("remote-ethernet-segment-identifier", types.YChild{"RemoteEthernetSegmentIdentifier", nil})
    for i := range mac.RemoteEthernetSegmentIdentifier {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.RemoteEthernetSegmentIdentifier[i]), types.YChild{"RemoteEthernetSegmentIdentifier", mac.RemoteEthernetSegmentIdentifier[i]})
    }
    mac.EntityData.Children.Append("path-buffer", types.YChild{"PathBuffer", nil})
    for i := range mac.PathBuffer {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.PathBuffer[i]), types.YChild{"PathBuffer", mac.PathBuffer[i]})
    }
    mac.EntityData.Leafs = types.NewOrderedMap()
    mac.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", mac.Evi})
    mac.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", mac.Encapsulation})
    mac.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", mac.EthernetTag})
    mac.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", mac.MacAddress})
    mac.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", mac.IpAddress})
    mac.EntityData.Leafs.Append("ethernet-tag-xr", types.YLeaf{"EthernetTagXr", mac.EthernetTagXr})
    mac.EntityData.Leafs.Append("mac-address-xr", types.YLeaf{"MacAddressXr", mac.MacAddressXr})
    mac.EntityData.Leafs.Append("ip-address-xr", types.YLeaf{"IpAddressXr", mac.IpAddressXr})
    mac.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", mac.LocalLabel})
    mac.EntityData.Leafs.Append("num-paths", types.YLeaf{"NumPaths", mac.NumPaths})
    mac.EntityData.Leafs.Append("is-local-mac", types.YLeaf{"IsLocalMac", mac.IsLocalMac})
    mac.EntityData.Leafs.Append("is-proxy-entry", types.YLeaf{"IsProxyEntry", mac.IsProxyEntry})
    mac.EntityData.Leafs.Append("is-remote-mac", types.YLeaf{"IsRemoteMac", mac.IsRemoteMac})
    mac.EntityData.Leafs.Append("soo-nexthop", types.YLeaf{"SooNexthop", mac.SooNexthop})
    mac.EntityData.Leafs.Append("ipnh-address", types.YLeaf{"IpnhAddress", mac.IpnhAddress})
    mac.EntityData.Leafs.Append("esi-port-key", types.YLeaf{"EsiPortKey", mac.EsiPortKey})
    mac.EntityData.Leafs.Append("local-encap-type", types.YLeaf{"LocalEncapType", mac.LocalEncapType})
    mac.EntityData.Leafs.Append("remote-encap-type", types.YLeaf{"RemoteEncapType", mac.RemoteEncapType})
    mac.EntityData.Leafs.Append("learned-bridge-port-name", types.YLeaf{"LearnedBridgePortName", mac.LearnedBridgePortName})
    mac.EntityData.Leafs.Append("local-seq-id", types.YLeaf{"LocalSeqId", mac.LocalSeqId})
    mac.EntityData.Leafs.Append("remote-seq-id", types.YLeaf{"RemoteSeqId", mac.RemoteSeqId})
    mac.EntityData.Leafs.Append("local-l3-label", types.YLeaf{"LocalL3Label", mac.LocalL3Label})
    mac.EntityData.Leafs.Append("router-mac-address", types.YLeaf{"RouterMacAddress", mac.RouterMacAddress})
    mac.EntityData.Leafs.Append("mac-flush-requested", types.YLeaf{"MacFlushRequested", mac.MacFlushRequested})
    mac.EntityData.Leafs.Append("mac-flush-received", types.YLeaf{"MacFlushReceived", mac.MacFlushReceived})
    mac.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", mac.InternalLabel})
    mac.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", mac.Resolved})
    mac.EntityData.Leafs.Append("local-is-static", types.YLeaf{"LocalIsStatic", mac.LocalIsStatic})
    mac.EntityData.Leafs.Append("remote-is-static", types.YLeaf{"RemoteIsStatic", mac.RemoteIsStatic})

    mac.EntityData.YListKeys = []string {}

    return &(mac.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Macs_Mac_EvpnInstance
// EVPN Instance summary information
type Evpn_Active_EviDetail_EviChildren_Macs_Mac_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Active_EviDetail_EviChildren_Macs_Mac_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "mac"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier
// Local Ethernet Segment id
type Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    localEthernetSegmentIdentifier.EntityData.YFilter = localEthernetSegmentIdentifier.YFilter
    localEthernetSegmentIdentifier.EntityData.YangName = "local-ethernet-segment-identifier"
    localEthernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    localEthernetSegmentIdentifier.EntityData.ParentYangName = "mac"
    localEthernetSegmentIdentifier.EntityData.SegmentPath = "local-ethernet-segment-identifier"
    localEthernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localEthernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localEthernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localEthernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    localEthernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    localEthernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", localEthernetSegmentIdentifier.Entry})

    localEthernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(localEthernetSegmentIdentifier.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier
// Remote Ethernet Segment id
type Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    remoteEthernetSegmentIdentifier.EntityData.YFilter = remoteEthernetSegmentIdentifier.YFilter
    remoteEthernetSegmentIdentifier.EntityData.YangName = "remote-ethernet-segment-identifier"
    remoteEthernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    remoteEthernetSegmentIdentifier.EntityData.ParentYangName = "mac"
    remoteEthernetSegmentIdentifier.EntityData.SegmentPath = "remote-ethernet-segment-identifier"
    remoteEthernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteEthernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteEthernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteEthernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    remoteEthernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    remoteEthernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", remoteEthernetSegmentIdentifier.Entry})

    remoteEthernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(remoteEthernetSegmentIdentifier.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer
// Path List Buffer
type Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetEntityData() *types.CommonEntityData {
    pathBuffer.EntityData.YFilter = pathBuffer.YFilter
    pathBuffer.EntityData.YangName = "path-buffer"
    pathBuffer.EntityData.BundleName = "cisco_ios_xr"
    pathBuffer.EntityData.ParentYangName = "mac"
    pathBuffer.EntityData.SegmentPath = "path-buffer"
    pathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathBuffer.EntityData.Children = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", pathBuffer.TunnelEndpointId})
    pathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", pathBuffer.NextHop})
    pathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", pathBuffer.OutputLabel})
    pathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel})

    pathBuffer.EntityData.YListKeys = []string {}

    return &(pathBuffer.EntityData)
}

// Evpn_Active_Teps
// L2VPN EVPN TEP Table
type Evpn_Active_Teps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN TEP Entry. The type is slice of Evpn_Active_Teps_Tep.
    Tep []*Evpn_Active_Teps_Tep
}

func (teps *Evpn_Active_Teps) GetEntityData() *types.CommonEntityData {
    teps.EntityData.YFilter = teps.YFilter
    teps.EntityData.YangName = "teps"
    teps.EntityData.BundleName = "cisco_ios_xr"
    teps.EntityData.ParentYangName = "active"
    teps.EntityData.SegmentPath = "teps"
    teps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teps.EntityData.Children = types.NewOrderedMap()
    teps.EntityData.Children.Append("tep", types.YChild{"Tep", nil})
    for i := range teps.Tep {
        teps.EntityData.Children.Append(types.GetSegmentPath(teps.Tep[i]), types.YChild{"Tep", teps.Tep[i]})
    }
    teps.EntityData.Leafs = types.NewOrderedMap()

    teps.EntityData.YListKeys = []string {}

    return &(teps.EntityData)
}

// Evpn_Active_Teps_Tep
// L2VPN EVPN TEP Entry
type Evpn_Active_Teps_Tep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TEP id. The type is interface{} with range:
    // 0..4294967295.
    TepId interface{}

    // Tunnel Endpoint id. The type is interface{} with range: 0..4294967295.
    TunnelEndpointId interface{}

    // EVPN Tunnel Endpoint Type. The type is interface{} with range: 0..255.
    Type interface{}

    // in-use counter. The type is interface{} with range: 0..4294967295.
    UseCount interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Table Id in RIB. The type is interface{} with range: 0..4294967295.
    VrfTableId interface{}

    // UDP port. The type is interface{} with range: 0..65535.
    UdpPort interface{}

    // Local TEP Information.
    LocalInfo Evpn_Active_Teps_Tep_LocalInfo

    // Remote TEP Information.
    RemoteInfo Evpn_Active_Teps_Tep_RemoteInfo
}

func (tep *Evpn_Active_Teps_Tep) GetEntityData() *types.CommonEntityData {
    tep.EntityData.YFilter = tep.YFilter
    tep.EntityData.YangName = "tep"
    tep.EntityData.BundleName = "cisco_ios_xr"
    tep.EntityData.ParentYangName = "teps"
    tep.EntityData.SegmentPath = "tep" + types.AddKeyToken(tep.TepId, "tep-id")
    tep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tep.EntityData.Children = types.NewOrderedMap()
    tep.EntityData.Children.Append("local-info", types.YChild{"LocalInfo", &tep.LocalInfo})
    tep.EntityData.Children.Append("remote-info", types.YChild{"RemoteInfo", &tep.RemoteInfo})
    tep.EntityData.Leafs = types.NewOrderedMap()
    tep.EntityData.Leafs.Append("tep-id", types.YLeaf{"TepId", tep.TepId})
    tep.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", tep.TunnelEndpointId})
    tep.EntityData.Leafs.Append("type", types.YLeaf{"Type", tep.Type})
    tep.EntityData.Leafs.Append("use-count", types.YLeaf{"UseCount", tep.UseCount})
    tep.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", tep.VrfName})
    tep.EntityData.Leafs.Append("vrf-table-id", types.YLeaf{"VrfTableId", tep.VrfTableId})
    tep.EntityData.Leafs.Append("udp-port", types.YLeaf{"UdpPort", tep.UdpPort})

    tep.EntityData.YListKeys = []string {"TepId"}

    return &(tep.EntityData)
}

// Evpn_Active_Teps_Tep_LocalInfo
// Local TEP Information
type Evpn_Active_Teps_Tep_LocalInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Tunnel encapsulation. The type is interface{} with range: 0..255.
    Encapsulation interface{}

    // IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}
}

func (localInfo *Evpn_Active_Teps_Tep_LocalInfo) GetEntityData() *types.CommonEntityData {
    localInfo.EntityData.YFilter = localInfo.YFilter
    localInfo.EntityData.YangName = "local-info"
    localInfo.EntityData.BundleName = "cisco_ios_xr"
    localInfo.EntityData.ParentYangName = "tep"
    localInfo.EntityData.SegmentPath = "local-info"
    localInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localInfo.EntityData.Children = types.NewOrderedMap()
    localInfo.EntityData.Leafs = types.NewOrderedMap()
    localInfo.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", localInfo.EthernetVpnId})
    localInfo.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", localInfo.Encapsulation})
    localInfo.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", localInfo.Ip})

    localInfo.EntityData.YListKeys = []string {}

    return &(localInfo.EntityData)
}

// Evpn_Active_Teps_Tep_RemoteInfo
// Remote TEP Information
type Evpn_Active_Teps_Tep_RemoteInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Tunnel encapsulation. The type is interface{} with range: 0..255.
    Encapsulation interface{}

    // IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}
}

func (remoteInfo *Evpn_Active_Teps_Tep_RemoteInfo) GetEntityData() *types.CommonEntityData {
    remoteInfo.EntityData.YFilter = remoteInfo.YFilter
    remoteInfo.EntityData.YangName = "remote-info"
    remoteInfo.EntityData.BundleName = "cisco_ios_xr"
    remoteInfo.EntityData.ParentYangName = "tep"
    remoteInfo.EntityData.SegmentPath = "remote-info"
    remoteInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteInfo.EntityData.Children = types.NewOrderedMap()
    remoteInfo.EntityData.Leafs = types.NewOrderedMap()
    remoteInfo.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", remoteInfo.EthernetVpnId})
    remoteInfo.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", remoteInfo.Encapsulation})
    remoteInfo.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", remoteInfo.Ip})

    remoteInfo.EntityData.YListKeys = []string {}

    return &(remoteInfo.EntityData)
}

// Evpn_Active_InternalLabels
// EVPN Internal Label Table
type Evpn_Active_InternalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN Internal Label. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel.
    InternalLabel []*Evpn_Active_InternalLabels_InternalLabel
}

func (internalLabels *Evpn_Active_InternalLabels) GetEntityData() *types.CommonEntityData {
    internalLabels.EntityData.YFilter = internalLabels.YFilter
    internalLabels.EntityData.YangName = "internal-labels"
    internalLabels.EntityData.BundleName = "cisco_ios_xr"
    internalLabels.EntityData.ParentYangName = "active"
    internalLabels.EntityData.SegmentPath = "internal-labels"
    internalLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalLabels.EntityData.Children = types.NewOrderedMap()
    internalLabels.EntityData.Children.Append("internal-label", types.YChild{"InternalLabel", nil})
    for i := range internalLabels.InternalLabel {
        internalLabels.EntityData.Children.Append(types.GetSegmentPath(internalLabels.InternalLabel[i]), types.YChild{"InternalLabel", internalLabels.InternalLabel[i]})
    }
    internalLabels.EntityData.Leafs = types.NewOrderedMap()

    internalLabels.EntityData.YListKeys = []string {}

    return &(internalLabels.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel
// L2VPN EVPN Internal Label
type Evpn_Active_InternalLabels_InternalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // Ethernet Segment id. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Esi interface{}

    // Label Tag. The type is interface{} with range: 0..4294967295.
    Tag interface{}

    // MPLS Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Number of items in the MAC path buffer. The type is interface{} with range:
    // 0..4294967295.
    MacNumPaths interface{}

    // Number of items in the ead path buffer. The type is interface{} with range:
    // 0..4294967295.
    EadNumPaths interface{}

    // Number of items in the evi path buffer. The type is interface{} with range:
    // 0..4294967295.
    EviNumPaths interface{}

    // Number of items in the sum path buffer. The type is interface{} with range:
    // 0..4294967295.
    SumNumPaths interface{}

    // Number of items in the sum path buffer that are Active Paths. The type is
    // interface{} with range: 0..4294967295.
    SumNumActivePaths interface{}

    // Internal Label has resolved per-ES EAD and per-EVI EAD or MAC routes. The
    // type is bool.
    Resolved interface{}

    // ECMP Disable Per EVI Resolution. The type is bool.
    EcmpDisable interface{}

    // Single-active redundancy configured at remote ES. The type is bool.
    RedundancySingleActive interface{}

    // Single-flow-active redundancy at remote ES (MST-AG). The type is bool.
    RedundancySingleFlowActive interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Active_InternalLabels_InternalLabel_EvpnInstance

    // MAC Path list buffer. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel_MacPathBuffer.
    MacPathBuffer []*Evpn_Active_InternalLabels_InternalLabel_MacPathBuffer

    // EAD/ES Path list buffer. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel_EadPathBuffer.
    EadPathBuffer []*Evpn_Active_InternalLabels_InternalLabel_EadPathBuffer

    // EAD/EVI Path list buffer. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel_EviPathBuffer.
    EviPathBuffer []*Evpn_Active_InternalLabels_InternalLabel_EviPathBuffer

    // Summary Path list buffer. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel_SummaryPathBuffer.
    SummaryPathBuffer []*Evpn_Active_InternalLabels_InternalLabel_SummaryPathBuffer
}

func (internalLabel *Evpn_Active_InternalLabels_InternalLabel) GetEntityData() *types.CommonEntityData {
    internalLabel.EntityData.YFilter = internalLabel.YFilter
    internalLabel.EntityData.YangName = "internal-label"
    internalLabel.EntityData.BundleName = "cisco_ios_xr"
    internalLabel.EntityData.ParentYangName = "internal-labels"
    internalLabel.EntityData.SegmentPath = "internal-label"
    internalLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalLabel.EntityData.Children = types.NewOrderedMap()
    internalLabel.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &internalLabel.EvpnInstance})
    internalLabel.EntityData.Children.Append("mac-path-buffer", types.YChild{"MacPathBuffer", nil})
    for i := range internalLabel.MacPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.MacPathBuffer[i]), types.YChild{"MacPathBuffer", internalLabel.MacPathBuffer[i]})
    }
    internalLabel.EntityData.Children.Append("ead-path-buffer", types.YChild{"EadPathBuffer", nil})
    for i := range internalLabel.EadPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.EadPathBuffer[i]), types.YChild{"EadPathBuffer", internalLabel.EadPathBuffer[i]})
    }
    internalLabel.EntityData.Children.Append("evi-path-buffer", types.YChild{"EviPathBuffer", nil})
    for i := range internalLabel.EviPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.EviPathBuffer[i]), types.YChild{"EviPathBuffer", internalLabel.EviPathBuffer[i]})
    }
    internalLabel.EntityData.Children.Append("summary-path-buffer", types.YChild{"SummaryPathBuffer", nil})
    for i := range internalLabel.SummaryPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.SummaryPathBuffer[i]), types.YChild{"SummaryPathBuffer", internalLabel.SummaryPathBuffer[i]})
    }
    internalLabel.EntityData.Leafs = types.NewOrderedMap()
    internalLabel.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", internalLabel.Evi})
    internalLabel.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", internalLabel.Encapsulation})
    internalLabel.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", internalLabel.Esi1})
    internalLabel.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", internalLabel.Esi2})
    internalLabel.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", internalLabel.Esi3})
    internalLabel.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", internalLabel.Esi4})
    internalLabel.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", internalLabel.Esi5})
    internalLabel.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", internalLabel.EthernetTag})
    internalLabel.EntityData.Leafs.Append("esi", types.YLeaf{"Esi", internalLabel.Esi})
    internalLabel.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", internalLabel.Tag})
    internalLabel.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", internalLabel.InternalLabel})
    internalLabel.EntityData.Leafs.Append("mac-num-paths", types.YLeaf{"MacNumPaths", internalLabel.MacNumPaths})
    internalLabel.EntityData.Leafs.Append("ead-num-paths", types.YLeaf{"EadNumPaths", internalLabel.EadNumPaths})
    internalLabel.EntityData.Leafs.Append("evi-num-paths", types.YLeaf{"EviNumPaths", internalLabel.EviNumPaths})
    internalLabel.EntityData.Leafs.Append("sum-num-paths", types.YLeaf{"SumNumPaths", internalLabel.SumNumPaths})
    internalLabel.EntityData.Leafs.Append("sum-num-active-paths", types.YLeaf{"SumNumActivePaths", internalLabel.SumNumActivePaths})
    internalLabel.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", internalLabel.Resolved})
    internalLabel.EntityData.Leafs.Append("ecmp-disable", types.YLeaf{"EcmpDisable", internalLabel.EcmpDisable})
    internalLabel.EntityData.Leafs.Append("redundancy-single-active", types.YLeaf{"RedundancySingleActive", internalLabel.RedundancySingleActive})
    internalLabel.EntityData.Leafs.Append("redundancy-single-flow-active", types.YLeaf{"RedundancySingleFlowActive", internalLabel.RedundancySingleFlowActive})

    internalLabel.EntityData.YListKeys = []string {}

    return &(internalLabel.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel_EvpnInstance
// EVPN Instance summary information
type Evpn_Active_InternalLabels_InternalLabel_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Active_InternalLabels_InternalLabel_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "internal-label"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel_MacPathBuffer
// MAC Path list buffer
type Evpn_Active_InternalLabels_InternalLabel_MacPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (macPathBuffer *Evpn_Active_InternalLabels_InternalLabel_MacPathBuffer) GetEntityData() *types.CommonEntityData {
    macPathBuffer.EntityData.YFilter = macPathBuffer.YFilter
    macPathBuffer.EntityData.YangName = "mac-path-buffer"
    macPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    macPathBuffer.EntityData.ParentYangName = "internal-label"
    macPathBuffer.EntityData.SegmentPath = "mac-path-buffer"
    macPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macPathBuffer.EntityData.Children = types.NewOrderedMap()
    macPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    macPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", macPathBuffer.TunnelEndpointId})
    macPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", macPathBuffer.NextHop})
    macPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", macPathBuffer.OutputLabel})
    macPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", macPathBuffer.SrteTunnel})

    macPathBuffer.EntityData.YListKeys = []string {}

    return &(macPathBuffer.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel_EadPathBuffer
// EAD/ES Path list buffer
type Evpn_Active_InternalLabels_InternalLabel_EadPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (eadPathBuffer *Evpn_Active_InternalLabels_InternalLabel_EadPathBuffer) GetEntityData() *types.CommonEntityData {
    eadPathBuffer.EntityData.YFilter = eadPathBuffer.YFilter
    eadPathBuffer.EntityData.YangName = "ead-path-buffer"
    eadPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    eadPathBuffer.EntityData.ParentYangName = "internal-label"
    eadPathBuffer.EntityData.SegmentPath = "ead-path-buffer"
    eadPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eadPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eadPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eadPathBuffer.EntityData.Children = types.NewOrderedMap()
    eadPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    eadPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", eadPathBuffer.TunnelEndpointId})
    eadPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", eadPathBuffer.NextHop})
    eadPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", eadPathBuffer.OutputLabel})
    eadPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", eadPathBuffer.SrteTunnel})

    eadPathBuffer.EntityData.YListKeys = []string {}

    return &(eadPathBuffer.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel_EviPathBuffer
// EAD/EVI Path list buffer
type Evpn_Active_InternalLabels_InternalLabel_EviPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (eviPathBuffer *Evpn_Active_InternalLabels_InternalLabel_EviPathBuffer) GetEntityData() *types.CommonEntityData {
    eviPathBuffer.EntityData.YFilter = eviPathBuffer.YFilter
    eviPathBuffer.EntityData.YangName = "evi-path-buffer"
    eviPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    eviPathBuffer.EntityData.ParentYangName = "internal-label"
    eviPathBuffer.EntityData.SegmentPath = "evi-path-buffer"
    eviPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviPathBuffer.EntityData.Children = types.NewOrderedMap()
    eviPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    eviPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", eviPathBuffer.TunnelEndpointId})
    eviPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", eviPathBuffer.NextHop})
    eviPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", eviPathBuffer.OutputLabel})
    eviPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", eviPathBuffer.SrteTunnel})

    eviPathBuffer.EntityData.YListKeys = []string {}

    return &(eviPathBuffer.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel_SummaryPathBuffer
// Summary Path list buffer
type Evpn_Active_InternalLabels_InternalLabel_SummaryPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (summaryPathBuffer *Evpn_Active_InternalLabels_InternalLabel_SummaryPathBuffer) GetEntityData() *types.CommonEntityData {
    summaryPathBuffer.EntityData.YFilter = summaryPathBuffer.YFilter
    summaryPathBuffer.EntityData.YangName = "summary-path-buffer"
    summaryPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    summaryPathBuffer.EntityData.ParentYangName = "internal-label"
    summaryPathBuffer.EntityData.SegmentPath = "summary-path-buffer"
    summaryPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPathBuffer.EntityData.Children = types.NewOrderedMap()
    summaryPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    summaryPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", summaryPathBuffer.TunnelEndpointId})
    summaryPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", summaryPathBuffer.NextHop})
    summaryPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", summaryPathBuffer.OutputLabel})
    summaryPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", summaryPathBuffer.SrteTunnel})

    summaryPathBuffer.EntityData.YListKeys = []string {}

    return &(summaryPathBuffer.EntityData)
}

// Evpn_Active_EthernetSegments
// EVPN Ethernet-Segment Table
type Evpn_Active_EthernetSegments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet-Segment Entry. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment.
    EthernetSegment []*Evpn_Active_EthernetSegments_EthernetSegment
}

func (ethernetSegments *Evpn_Active_EthernetSegments) GetEntityData() *types.CommonEntityData {
    ethernetSegments.EntityData.YFilter = ethernetSegments.YFilter
    ethernetSegments.EntityData.YangName = "ethernet-segments"
    ethernetSegments.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegments.EntityData.ParentYangName = "active"
    ethernetSegments.EntityData.SegmentPath = "ethernet-segments"
    ethernetSegments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegments.EntityData.Children = types.NewOrderedMap()
    ethernetSegments.EntityData.Children.Append("ethernet-segment", types.YChild{"EthernetSegment", nil})
    for i := range ethernetSegments.EthernetSegment {
        ethernetSegments.EntityData.Children.Append(types.GetSegmentPath(ethernetSegments.EthernetSegment[i]), types.YChild{"EthernetSegment", ethernetSegments.EthernetSegment[i]})
    }
    ethernetSegments.EntityData.Leafs = types.NewOrderedMap()

    ethernetSegments.EntityData.YListKeys = []string {}

    return &(ethernetSegments.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment
// EVPN Ethernet-Segment Entry
type Evpn_Active_EthernetSegments_EthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // ESI Type. The type is L2vpnEvpnEsi.
    EsiType interface{}

    // ESI System Identifier. The type is string.
    EsiSystemIdentifier interface{}

    // ESI Port Key. The type is interface{} with range: 0..4294967295.
    EsiPortKey interface{}

    // ESI System Priority. The type is interface{} with range: 0..4294967295.
    EsiSystemPriority interface{}

    // Ethernet Segment Name. The type is string.
    EthernetSegmentName interface{}

    // State of the ethernet segment. The type is interface{} with range:
    // 0..4294967295.
    EthernetSegmentState interface{}

    // Main port ifhandle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    IfHandle interface{}

    // Main port redundancy group role. The type is L2vpnRgRole.
    MainPortRole interface{}

    // Main Port MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MainPortMac interface{}

    // Number of PWs in Up state. The type is interface{} with range:
    // 0..4294967295.
    NumUpPWs interface{}

    // ES-Import Route Target. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    RouteTarget interface{}

    // Origin of operational ES-Import RT. The type is L2vpnEvpnRtOrigin.
    RtOrigin interface{}

    // ES BGP Gates. The type is string.
    EsBgpGates interface{}

    // ES L2FIB Gates. The type is string.
    EsL2fibGates interface{}

    // Configured MAC Flushing mode. The type is L2vpnEvpnMfMode.
    MacFlushingModeConfig interface{}

    // Configured load balancing mode. The type is L2vpnEvpnLbMode.
    LoadBalanceModeConfig interface{}

    // Load balancing mode is default. The type is bool.
    LoadBalanceModeIsDefault interface{}

    // Operational load balancing mode. The type is L2vpnEvpnLbMode.
    LoadBalanceModeOper interface{}

    // Ethernet-Segment forced to single home. The type is bool.
    ForceSingleHome interface{}

    // Operational Source MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacOper interface{}

    // Origin of operational source MAC address. The type is L2vpnEvpnSmacSrc.
    SourceMacOrigin interface{}

    // Configured timer for triggering DF election (seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    PeeringTimer interface{}

    // Milliseconds left on DF election timer. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    PeeringTimerLeft interface{}

    // Configured timer for (STP) recovery (seconds). The type is interface{} with
    // range: 0..4294967295. Units are second.
    RecoveryTimer interface{}

    // Milliseconds left on (STP) recovery timer. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RecoveryTimerLeft interface{}

    // Configured timer for delaying DF election (seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    CarvingTimer interface{}

    // Milliseconds left on carving timer. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    CarvingTimerLeft interface{}

    // Service carving mode. The type is L2vpnEvpnScMode.
    ServiceCarvingMode interface{}

    // Input string of Primary services ESI/I-SIDs. The type is string.
    PrimaryServicesInput interface{}

    // Input string of Secondary services ESI/I-SIDs. The type is string.
    SecondaryServicesInput interface{}

    // Count of Forwarders (AC, AC PW, VFI PW). The type is interface{} with
    // range: 0..4294967295.
    ForwarderPorts interface{}

    // Count of Forwarders with permanent service. The type is interface{} with
    // range: 0..4294967295.
    PermanentForwarderPorts interface{}

    // Count of Forwarders with elected service. The type is interface{} with
    // range: 0..4294967295.
    ElectedForwarderPorts interface{}

    // Count of Forwarders with not elected service. The type is interface{} with
    // range: 0..4294967295.
    NotElectedForwarderPorts interface{}

    // Count of forwarders with missing config detected. The type is interface{}
    // with range: 0..4294967295.
    NotConfigForwarderPorts interface{}

    // MP is protected and not under EVPN control. The type is bool.
    MpProtected interface{}

    // Anycast VTEP mode on NVE main-interface. The type is bool.
    NveAnycastVtep interface{}

    // Ingress-Replication is configured on NVE main-interface. The type is bool.
    NveIngressReplication interface{}

    // Local split horizon group label is valid. The type is bool.
    LocalSplitHorizonGroupLabelValid interface{}

    // Local split horizon group label. The type is interface{} with range:
    // 0..4294967295.
    LocalSplitHorizonGroupLabel interface{}

    // Ethernet Segment id. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier

    // List of Primary services ESI/I-SIDs. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService.
    PrimaryService []*Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService

    // List of Secondary services ESI/I-SIDs. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService.
    SecondaryService []*Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService

    // Elected ISID service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult.
    ServiceCarvingISidelectedResult []*Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult

    // Not elected ISID service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult.
    ServiceCarvingIsidNotElectedResult []*Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult

    // Elected EVI service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult.
    ServiceCarvingEviElectedResult []*Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult

    // Not elected EVI service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult.
    ServiceCarvingEviNotElectedResult []*Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult

    // Elected VNI service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult.
    ServiceCarvingVniElectedResult []*Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult

    // Not elected VNI service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult.
    ServiceCarvingVniNotElectedResult []*Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult

    // List of nexthop IPv6 addresses. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_NextHop.
    NextHop []*Evpn_Active_EthernetSegments_EthernetSegment_NextHop

    // Permanent EVPN VPWS service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult.
    ServiceCarvingVpwsPermanentResult []*Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult

    // Remote split horizon group labels. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel.
    RemoteSplitHorizonGroupLabel []*Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
}

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetEntityData() *types.CommonEntityData {
    ethernetSegment.EntityData.YFilter = ethernetSegment.YFilter
    ethernetSegment.EntityData.YangName = "ethernet-segment"
    ethernetSegment.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegment.EntityData.ParentYangName = "ethernet-segments"
    ethernetSegment.EntityData.SegmentPath = "ethernet-segment"
    ethernetSegment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegment.EntityData.Children = types.NewOrderedMap()
    ethernetSegment.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range ethernetSegment.EthernetSegmentIdentifier {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", ethernetSegment.EthernetSegmentIdentifier[i]})
    }
    ethernetSegment.EntityData.Children.Append("primary-service", types.YChild{"PrimaryService", nil})
    for i := range ethernetSegment.PrimaryService {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.PrimaryService[i]), types.YChild{"PrimaryService", ethernetSegment.PrimaryService[i]})
    }
    ethernetSegment.EntityData.Children.Append("secondary-service", types.YChild{"SecondaryService", nil})
    for i := range ethernetSegment.SecondaryService {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.SecondaryService[i]), types.YChild{"SecondaryService", ethernetSegment.SecondaryService[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-i-sidelected-result", types.YChild{"ServiceCarvingISidelectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingISidelectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingISidelectedResult[i]), types.YChild{"ServiceCarvingISidelectedResult", ethernetSegment.ServiceCarvingISidelectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-isid-not-elected-result", types.YChild{"ServiceCarvingIsidNotElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingIsidNotElectedResult[i]), types.YChild{"ServiceCarvingIsidNotElectedResult", ethernetSegment.ServiceCarvingIsidNotElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-evi-elected-result", types.YChild{"ServiceCarvingEviElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingEviElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingEviElectedResult[i]), types.YChild{"ServiceCarvingEviElectedResult", ethernetSegment.ServiceCarvingEviElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-evi-not-elected-result", types.YChild{"ServiceCarvingEviNotElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingEviNotElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingEviNotElectedResult[i]), types.YChild{"ServiceCarvingEviNotElectedResult", ethernetSegment.ServiceCarvingEviNotElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-vni-elected-result", types.YChild{"ServiceCarvingVniElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingVniElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingVniElectedResult[i]), types.YChild{"ServiceCarvingVniElectedResult", ethernetSegment.ServiceCarvingVniElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-vni-not-elected-result", types.YChild{"ServiceCarvingVniNotElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingVniNotElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingVniNotElectedResult[i]), types.YChild{"ServiceCarvingVniNotElectedResult", ethernetSegment.ServiceCarvingVniNotElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range ethernetSegment.NextHop {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.NextHop[i]), types.YChild{"NextHop", ethernetSegment.NextHop[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-vpws-permanent-result", types.YChild{"ServiceCarvingVpwsPermanentResult", nil})
    for i := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingVpwsPermanentResult[i]), types.YChild{"ServiceCarvingVpwsPermanentResult", ethernetSegment.ServiceCarvingVpwsPermanentResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("remote-split-horizon-group-label", types.YChild{"RemoteSplitHorizonGroupLabel", nil})
    for i := range ethernetSegment.RemoteSplitHorizonGroupLabel {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.RemoteSplitHorizonGroupLabel[i]), types.YChild{"RemoteSplitHorizonGroupLabel", ethernetSegment.RemoteSplitHorizonGroupLabel[i]})
    }
    ethernetSegment.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegment.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ethernetSegment.InterfaceName})
    ethernetSegment.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", ethernetSegment.Esi1})
    ethernetSegment.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", ethernetSegment.Esi2})
    ethernetSegment.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", ethernetSegment.Esi3})
    ethernetSegment.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", ethernetSegment.Esi4})
    ethernetSegment.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", ethernetSegment.Esi5})
    ethernetSegment.EntityData.Leafs.Append("esi-type", types.YLeaf{"EsiType", ethernetSegment.EsiType})
    ethernetSegment.EntityData.Leafs.Append("esi-system-identifier", types.YLeaf{"EsiSystemIdentifier", ethernetSegment.EsiSystemIdentifier})
    ethernetSegment.EntityData.Leafs.Append("esi-port-key", types.YLeaf{"EsiPortKey", ethernetSegment.EsiPortKey})
    ethernetSegment.EntityData.Leafs.Append("esi-system-priority", types.YLeaf{"EsiSystemPriority", ethernetSegment.EsiSystemPriority})
    ethernetSegment.EntityData.Leafs.Append("ethernet-segment-name", types.YLeaf{"EthernetSegmentName", ethernetSegment.EthernetSegmentName})
    ethernetSegment.EntityData.Leafs.Append("ethernet-segment-state", types.YLeaf{"EthernetSegmentState", ethernetSegment.EthernetSegmentState})
    ethernetSegment.EntityData.Leafs.Append("if-handle", types.YLeaf{"IfHandle", ethernetSegment.IfHandle})
    ethernetSegment.EntityData.Leafs.Append("main-port-role", types.YLeaf{"MainPortRole", ethernetSegment.MainPortRole})
    ethernetSegment.EntityData.Leafs.Append("main-port-mac", types.YLeaf{"MainPortMac", ethernetSegment.MainPortMac})
    ethernetSegment.EntityData.Leafs.Append("num-up-p-ws", types.YLeaf{"NumUpPWs", ethernetSegment.NumUpPWs})
    ethernetSegment.EntityData.Leafs.Append("route-target", types.YLeaf{"RouteTarget", ethernetSegment.RouteTarget})
    ethernetSegment.EntityData.Leafs.Append("rt-origin", types.YLeaf{"RtOrigin", ethernetSegment.RtOrigin})
    ethernetSegment.EntityData.Leafs.Append("es-bgp-gates", types.YLeaf{"EsBgpGates", ethernetSegment.EsBgpGates})
    ethernetSegment.EntityData.Leafs.Append("es-l2fib-gates", types.YLeaf{"EsL2fibGates", ethernetSegment.EsL2fibGates})
    ethernetSegment.EntityData.Leafs.Append("mac-flushing-mode-config", types.YLeaf{"MacFlushingModeConfig", ethernetSegment.MacFlushingModeConfig})
    ethernetSegment.EntityData.Leafs.Append("load-balance-mode-config", types.YLeaf{"LoadBalanceModeConfig", ethernetSegment.LoadBalanceModeConfig})
    ethernetSegment.EntityData.Leafs.Append("load-balance-mode-is-default", types.YLeaf{"LoadBalanceModeIsDefault", ethernetSegment.LoadBalanceModeIsDefault})
    ethernetSegment.EntityData.Leafs.Append("load-balance-mode-oper", types.YLeaf{"LoadBalanceModeOper", ethernetSegment.LoadBalanceModeOper})
    ethernetSegment.EntityData.Leafs.Append("force-single-home", types.YLeaf{"ForceSingleHome", ethernetSegment.ForceSingleHome})
    ethernetSegment.EntityData.Leafs.Append("source-mac-oper", types.YLeaf{"SourceMacOper", ethernetSegment.SourceMacOper})
    ethernetSegment.EntityData.Leafs.Append("source-mac-origin", types.YLeaf{"SourceMacOrigin", ethernetSegment.SourceMacOrigin})
    ethernetSegment.EntityData.Leafs.Append("peering-timer", types.YLeaf{"PeeringTimer", ethernetSegment.PeeringTimer})
    ethernetSegment.EntityData.Leafs.Append("peering-timer-left", types.YLeaf{"PeeringTimerLeft", ethernetSegment.PeeringTimerLeft})
    ethernetSegment.EntityData.Leafs.Append("recovery-timer", types.YLeaf{"RecoveryTimer", ethernetSegment.RecoveryTimer})
    ethernetSegment.EntityData.Leafs.Append("recovery-timer-left", types.YLeaf{"RecoveryTimerLeft", ethernetSegment.RecoveryTimerLeft})
    ethernetSegment.EntityData.Leafs.Append("carving-timer", types.YLeaf{"CarvingTimer", ethernetSegment.CarvingTimer})
    ethernetSegment.EntityData.Leafs.Append("carving-timer-left", types.YLeaf{"CarvingTimerLeft", ethernetSegment.CarvingTimerLeft})
    ethernetSegment.EntityData.Leafs.Append("service-carving-mode", types.YLeaf{"ServiceCarvingMode", ethernetSegment.ServiceCarvingMode})
    ethernetSegment.EntityData.Leafs.Append("primary-services-input", types.YLeaf{"PrimaryServicesInput", ethernetSegment.PrimaryServicesInput})
    ethernetSegment.EntityData.Leafs.Append("secondary-services-input", types.YLeaf{"SecondaryServicesInput", ethernetSegment.SecondaryServicesInput})
    ethernetSegment.EntityData.Leafs.Append("forwarder-ports", types.YLeaf{"ForwarderPorts", ethernetSegment.ForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("permanent-forwarder-ports", types.YLeaf{"PermanentForwarderPorts", ethernetSegment.PermanentForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("elected-forwarder-ports", types.YLeaf{"ElectedForwarderPorts", ethernetSegment.ElectedForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("not-elected-forwarder-ports", types.YLeaf{"NotElectedForwarderPorts", ethernetSegment.NotElectedForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("not-config-forwarder-ports", types.YLeaf{"NotConfigForwarderPorts", ethernetSegment.NotConfigForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("mp-protected", types.YLeaf{"MpProtected", ethernetSegment.MpProtected})
    ethernetSegment.EntityData.Leafs.Append("nve-anycast-vtep", types.YLeaf{"NveAnycastVtep", ethernetSegment.NveAnycastVtep})
    ethernetSegment.EntityData.Leafs.Append("nve-ingress-replication", types.YLeaf{"NveIngressReplication", ethernetSegment.NveIngressReplication})
    ethernetSegment.EntityData.Leafs.Append("local-split-horizon-group-label-valid", types.YLeaf{"LocalSplitHorizonGroupLabelValid", ethernetSegment.LocalSplitHorizonGroupLabelValid})
    ethernetSegment.EntityData.Leafs.Append("local-split-horizon-group-label", types.YLeaf{"LocalSplitHorizonGroupLabel", ethernetSegment.LocalSplitHorizonGroupLabel})

    ethernetSegment.EntityData.YListKeys = []string {}

    return &(ethernetSegment.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "ethernet-segment"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService
// List of Primary services ESI/I-SIDs
type Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetEntityData() *types.CommonEntityData {
    primaryService.EntityData.YFilter = primaryService.YFilter
    primaryService.EntityData.YangName = "primary-service"
    primaryService.EntityData.BundleName = "cisco_ios_xr"
    primaryService.EntityData.ParentYangName = "ethernet-segment"
    primaryService.EntityData.SegmentPath = "primary-service"
    primaryService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryService.EntityData.Children = types.NewOrderedMap()
    primaryService.EntityData.Leafs = types.NewOrderedMap()
    primaryService.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", primaryService.Entry})

    primaryService.EntityData.YListKeys = []string {}

    return &(primaryService.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService
// List of Secondary services ESI/I-SIDs
type Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetEntityData() *types.CommonEntityData {
    secondaryService.EntityData.YFilter = secondaryService.YFilter
    secondaryService.EntityData.YangName = "secondary-service"
    secondaryService.EntityData.BundleName = "cisco_ios_xr"
    secondaryService.EntityData.ParentYangName = "ethernet-segment"
    secondaryService.EntityData.SegmentPath = "secondary-service"
    secondaryService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryService.EntityData.Children = types.NewOrderedMap()
    secondaryService.EntityData.Leafs = types.NewOrderedMap()
    secondaryService.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", secondaryService.Entry})

    secondaryService.EntityData.YListKeys = []string {}

    return &(secondaryService.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult
// Elected ISID service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingISidelectedResult.EntityData.YFilter = serviceCarvingISidelectedResult.YFilter
    serviceCarvingISidelectedResult.EntityData.YangName = "service-carving-i-sidelected-result"
    serviceCarvingISidelectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingISidelectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingISidelectedResult.EntityData.SegmentPath = "service-carving-i-sidelected-result"
    serviceCarvingISidelectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingISidelectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingISidelectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingISidelectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingISidelectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingISidelectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingISidelectedResult.Entry})

    serviceCarvingISidelectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingISidelectedResult.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult
// Not elected ISID service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingIsidNotElectedResult.EntityData.YFilter = serviceCarvingIsidNotElectedResult.YFilter
    serviceCarvingIsidNotElectedResult.EntityData.YangName = "service-carving-isid-not-elected-result"
    serviceCarvingIsidNotElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingIsidNotElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingIsidNotElectedResult.EntityData.SegmentPath = "service-carving-isid-not-elected-result"
    serviceCarvingIsidNotElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingIsidNotElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingIsidNotElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingIsidNotElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingIsidNotElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingIsidNotElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingIsidNotElectedResult.Entry})

    serviceCarvingIsidNotElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingIsidNotElectedResult.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult
// Elected EVI service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingEviElectedResult.EntityData.YFilter = serviceCarvingEviElectedResult.YFilter
    serviceCarvingEviElectedResult.EntityData.YangName = "service-carving-evi-elected-result"
    serviceCarvingEviElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingEviElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingEviElectedResult.EntityData.SegmentPath = "service-carving-evi-elected-result"
    serviceCarvingEviElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingEviElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingEviElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingEviElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingEviElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingEviElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingEviElectedResult.Entry})

    serviceCarvingEviElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingEviElectedResult.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult
// Not elected EVI service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingEviNotElectedResult.EntityData.YFilter = serviceCarvingEviNotElectedResult.YFilter
    serviceCarvingEviNotElectedResult.EntityData.YangName = "service-carving-evi-not-elected-result"
    serviceCarvingEviNotElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingEviNotElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingEviNotElectedResult.EntityData.SegmentPath = "service-carving-evi-not-elected-result"
    serviceCarvingEviNotElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingEviNotElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingEviNotElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingEviNotElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingEviNotElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingEviNotElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingEviNotElectedResult.Entry})

    serviceCarvingEviNotElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingEviNotElectedResult.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult
// Elected VNI service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingVniElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingVniElectedResult.EntityData.YFilter = serviceCarvingVniElectedResult.YFilter
    serviceCarvingVniElectedResult.EntityData.YangName = "service-carving-vni-elected-result"
    serviceCarvingVniElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingVniElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingVniElectedResult.EntityData.SegmentPath = "service-carving-vni-elected-result"
    serviceCarvingVniElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingVniElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingVniElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingVniElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingVniElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingVniElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingVniElectedResult.Entry})

    serviceCarvingVniElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingVniElectedResult.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult
// Not elected VNI service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingVniNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingVniNotElectedResult.EntityData.YFilter = serviceCarvingVniNotElectedResult.YFilter
    serviceCarvingVniNotElectedResult.EntityData.YangName = "service-carving-vni-not-elected-result"
    serviceCarvingVniNotElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingVniNotElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingVniNotElectedResult.EntityData.SegmentPath = "service-carving-vni-not-elected-result"
    serviceCarvingVniNotElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingVniNotElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingVniNotElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingVniNotElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingVniNotElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingVniNotElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingVniNotElectedResult.Entry})

    serviceCarvingVniNotElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingVniNotElectedResult.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_NextHop
// List of nexthop IPv6 addresses
type Evpn_Active_EthernetSegments_EthernetSegment_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // DF Dont Premption. The type is bool.
    DfDontPrempt interface{}

    // DF Election Mode Configured. The type is interface{} with range: 0..255.
    DfType interface{}

    // DF Election Preference Set. The type is interface{} with range: 0..65535.
    DfPref interface{}
}

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "ethernet-segment"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", nextHop.NextHop})
    nextHop.EntityData.Leafs.Append("df-dont-prempt", types.YLeaf{"DfDontPrempt", nextHop.DfDontPrempt})
    nextHop.EntityData.Leafs.Append("df-type", types.YLeaf{"DfType", nextHop.DfType})
    nextHop.EntityData.Leafs.Append("df-pref", types.YLeaf{"DfPref", nextHop.DfPref})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult
// Permanent EVPN VPWS service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPN ID. The type is interface{} with range: 0..4294967295.
    VpnId interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}
}

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingVpwsPermanentResult.EntityData.YFilter = serviceCarvingVpwsPermanentResult.YFilter
    serviceCarvingVpwsPermanentResult.EntityData.YangName = "service-carving-vpws-permanent-result"
    serviceCarvingVpwsPermanentResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingVpwsPermanentResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingVpwsPermanentResult.EntityData.SegmentPath = "service-carving-vpws-permanent-result"
    serviceCarvingVpwsPermanentResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingVpwsPermanentResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingVpwsPermanentResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingVpwsPermanentResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingVpwsPermanentResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingVpwsPermanentResult.EntityData.Leafs.Append("vpn-id", types.YLeaf{"VpnId", serviceCarvingVpwsPermanentResult.VpnId})
    serviceCarvingVpwsPermanentResult.EntityData.Leafs.Append("type", types.YLeaf{"Type", serviceCarvingVpwsPermanentResult.Type})
    serviceCarvingVpwsPermanentResult.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", serviceCarvingVpwsPermanentResult.EthernetTag})

    serviceCarvingVpwsPermanentResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingVpwsPermanentResult.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Split horizon label associated with next-hop address. The type is
    // interface{} with range: 0..4294967295.
    Label interface{}
}

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetEntityData() *types.CommonEntityData {
    remoteSplitHorizonGroupLabel.EntityData.YFilter = remoteSplitHorizonGroupLabel.YFilter
    remoteSplitHorizonGroupLabel.EntityData.YangName = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.BundleName = "cisco_ios_xr"
    remoteSplitHorizonGroupLabel.EntityData.ParentYangName = "ethernet-segment"
    remoteSplitHorizonGroupLabel.EntityData.SegmentPath = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteSplitHorizonGroupLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteSplitHorizonGroupLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteSplitHorizonGroupLabel.EntityData.Children = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", remoteSplitHorizonGroupLabel.NextHop})
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("label", types.YLeaf{"Label", remoteSplitHorizonGroupLabel.Label})

    remoteSplitHorizonGroupLabel.EntityData.YListKeys = []string {}

    return &(remoteSplitHorizonGroupLabel.EntityData)
}

// Evpn_Active_AcIds
// EVPN AC ID table
type Evpn_Active_AcIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN AC ID table. The type is slice of Evpn_Active_AcIds_AcId.
    AcId []*Evpn_Active_AcIds_AcId
}

func (acIds *Evpn_Active_AcIds) GetEntityData() *types.CommonEntityData {
    acIds.EntityData.YFilter = acIds.YFilter
    acIds.EntityData.YangName = "ac-ids"
    acIds.EntityData.BundleName = "cisco_ios_xr"
    acIds.EntityData.ParentYangName = "active"
    acIds.EntityData.SegmentPath = "ac-ids"
    acIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acIds.EntityData.Children = types.NewOrderedMap()
    acIds.EntityData.Children.Append("ac-id", types.YChild{"AcId", nil})
    for i := range acIds.AcId {
        acIds.EntityData.Children.Append(types.GetSegmentPath(acIds.AcId[i]), types.YChild{"AcId", acIds.AcId[i]})
    }
    acIds.EntityData.Leafs = types.NewOrderedMap()

    acIds.EntityData.YListKeys = []string {}

    return &(acIds.EntityData)
}

// Evpn_Active_AcIds_AcId
// EVPN AC ID table
type Evpn_Active_AcIds_AcId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // AC ID. The type is interface{} with range: 0..4294967295.
    AcId interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Active_AcIds_AcId_EvpnInstance
}

func (acId *Evpn_Active_AcIds_AcId) GetEntityData() *types.CommonEntityData {
    acId.EntityData.YFilter = acId.YFilter
    acId.EntityData.YangName = "ac-id"
    acId.EntityData.BundleName = "cisco_ios_xr"
    acId.EntityData.ParentYangName = "ac-ids"
    acId.EntityData.SegmentPath = "ac-id"
    acId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acId.EntityData.Children = types.NewOrderedMap()
    acId.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &acId.EvpnInstance})
    acId.EntityData.Leafs = types.NewOrderedMap()
    acId.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", acId.Evi})
    acId.EntityData.Leafs.Append("ac-id", types.YLeaf{"AcId", acId.AcId})
    acId.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", acId.Neighbor})

    acId.EntityData.YListKeys = []string {}

    return &(acId.EntityData)
}

// Evpn_Active_AcIds_AcId_EvpnInstance
// EVPN Instance summary information
type Evpn_Active_AcIds_AcId_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Active_AcIds_AcId_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "ac-id"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Standby
// Standby EVPN operational data
type Evpn_Standby struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Group Table.
    EvpnGroups Evpn_Standby_EvpnGroups

    // EVPN Remote SHG table.
    RemoteShgs Evpn_Standby_RemoteShgs

    // L2VPN EVPN Client.
    Client Evpn_Standby_Client

    // EVPN IGMP table.
    Igmps Evpn_Standby_Igmps

    // L2VPN EVPN EVI Table.
    Evis Evpn_Standby_Evis

    // L2VPN EVPN Summary.
    Summary Evpn_Standby_Summary

    // L2VPN EVI Detail Table.
    EviDetail Evpn_Standby_EviDetail

    // L2VPN EVPN TEP Table.
    Teps Evpn_Standby_Teps

    // EVPN Internal Label Table.
    InternalLabels Evpn_Standby_InternalLabels

    // EVPN Ethernet-Segment Table.
    EthernetSegments Evpn_Standby_EthernetSegments

    // EVPN AC ID table.
    AcIds Evpn_Standby_AcIds
}

func (standby *Evpn_Standby) GetEntityData() *types.CommonEntityData {
    standby.EntityData.YFilter = standby.YFilter
    standby.EntityData.YangName = "standby"
    standby.EntityData.BundleName = "cisco_ios_xr"
    standby.EntityData.ParentYangName = "evpn"
    standby.EntityData.SegmentPath = "standby"
    standby.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standby.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standby.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standby.EntityData.Children = types.NewOrderedMap()
    standby.EntityData.Children.Append("evpn-groups", types.YChild{"EvpnGroups", &standby.EvpnGroups})
    standby.EntityData.Children.Append("remote-shgs", types.YChild{"RemoteShgs", &standby.RemoteShgs})
    standby.EntityData.Children.Append("client", types.YChild{"Client", &standby.Client})
    standby.EntityData.Children.Append("igmps", types.YChild{"Igmps", &standby.Igmps})
    standby.EntityData.Children.Append("evis", types.YChild{"Evis", &standby.Evis})
    standby.EntityData.Children.Append("summary", types.YChild{"Summary", &standby.Summary})
    standby.EntityData.Children.Append("evi-detail", types.YChild{"EviDetail", &standby.EviDetail})
    standby.EntityData.Children.Append("teps", types.YChild{"Teps", &standby.Teps})
    standby.EntityData.Children.Append("internal-labels", types.YChild{"InternalLabels", &standby.InternalLabels})
    standby.EntityData.Children.Append("ethernet-segments", types.YChild{"EthernetSegments", &standby.EthernetSegments})
    standby.EntityData.Children.Append("ac-ids", types.YChild{"AcIds", &standby.AcIds})
    standby.EntityData.Leafs = types.NewOrderedMap()

    standby.EntityData.YListKeys = []string {}

    return &(standby.EntityData)
}

// Evpn_Standby_EvpnGroups
// EVPN Group Table
type Evpn_Standby_EvpnGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Group information. The type is slice of
    // Evpn_Standby_EvpnGroups_EvpnGroup.
    EvpnGroup []*Evpn_Standby_EvpnGroups_EvpnGroup
}

func (evpnGroups *Evpn_Standby_EvpnGroups) GetEntityData() *types.CommonEntityData {
    evpnGroups.EntityData.YFilter = evpnGroups.YFilter
    evpnGroups.EntityData.YangName = "evpn-groups"
    evpnGroups.EntityData.BundleName = "cisco_ios_xr"
    evpnGroups.EntityData.ParentYangName = "standby"
    evpnGroups.EntityData.SegmentPath = "evpn-groups"
    evpnGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroups.EntityData.Children = types.NewOrderedMap()
    evpnGroups.EntityData.Children.Append("evpn-group", types.YChild{"EvpnGroup", nil})
    for i := range evpnGroups.EvpnGroup {
        evpnGroups.EntityData.Children.Append(types.GetSegmentPath(evpnGroups.EvpnGroup[i]), types.YChild{"EvpnGroup", evpnGroups.EvpnGroup[i]})
    }
    evpnGroups.EntityData.Leafs = types.NewOrderedMap()

    evpnGroups.EntityData.YListKeys = []string {}

    return &(evpnGroups.EntityData)
}

// Evpn_Standby_EvpnGroups_EvpnGroup
// EVPN Group information
type Evpn_Standby_EvpnGroups_EvpnGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN group number. The type is interface{} with
    // range: 1..4294967295.
    GroupNumber interface{}

    // EVPN Group ID. The type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // EVPN Group State. The type is EvpnGrp.
    State interface{}

    // EVPN Group Core Interface table. The type is slice of
    // Evpn_Standby_EvpnGroups_EvpnGroup_CoreInterface.
    CoreInterface []*Evpn_Standby_EvpnGroups_EvpnGroup_CoreInterface

    // EVPN Access Core Interface table. The type is slice of
    // Evpn_Standby_EvpnGroups_EvpnGroup_AccessInterface.
    AccessInterface []*Evpn_Standby_EvpnGroups_EvpnGroup_AccessInterface
}

func (evpnGroup *Evpn_Standby_EvpnGroups_EvpnGroup) GetEntityData() *types.CommonEntityData {
    evpnGroup.EntityData.YFilter = evpnGroup.YFilter
    evpnGroup.EntityData.YangName = "evpn-group"
    evpnGroup.EntityData.BundleName = "cisco_ios_xr"
    evpnGroup.EntityData.ParentYangName = "evpn-groups"
    evpnGroup.EntityData.SegmentPath = "evpn-group" + types.AddKeyToken(evpnGroup.GroupNumber, "group-number")
    evpnGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnGroup.EntityData.Children = types.NewOrderedMap()
    evpnGroup.EntityData.Children.Append("core-interface", types.YChild{"CoreInterface", nil})
    for i := range evpnGroup.CoreInterface {
        evpnGroup.EntityData.Children.Append(types.GetSegmentPath(evpnGroup.CoreInterface[i]), types.YChild{"CoreInterface", evpnGroup.CoreInterface[i]})
    }
    evpnGroup.EntityData.Children.Append("access-interface", types.YChild{"AccessInterface", nil})
    for i := range evpnGroup.AccessInterface {
        evpnGroup.EntityData.Children.Append(types.GetSegmentPath(evpnGroup.AccessInterface[i]), types.YChild{"AccessInterface", evpnGroup.AccessInterface[i]})
    }
    evpnGroup.EntityData.Leafs = types.NewOrderedMap()
    evpnGroup.EntityData.Leafs.Append("group-number", types.YLeaf{"GroupNumber", evpnGroup.GroupNumber})
    evpnGroup.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", evpnGroup.GroupId})
    evpnGroup.EntityData.Leafs.Append("state", types.YLeaf{"State", evpnGroup.State})

    evpnGroup.EntityData.YListKeys = []string {"GroupNumber"}

    return &(evpnGroup.EntityData)
}

// Evpn_Standby_EvpnGroups_EvpnGroup_CoreInterface
// EVPN Group Core Interface table
type Evpn_Standby_EvpnGroups_EvpnGroup_CoreInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface State. The type is ImStateEnum.
    State interface{}
}

func (coreInterface *Evpn_Standby_EvpnGroups_EvpnGroup_CoreInterface) GetEntityData() *types.CommonEntityData {
    coreInterface.EntityData.YFilter = coreInterface.YFilter
    coreInterface.EntityData.YangName = "core-interface"
    coreInterface.EntityData.BundleName = "cisco_ios_xr"
    coreInterface.EntityData.ParentYangName = "evpn-group"
    coreInterface.EntityData.SegmentPath = "core-interface"
    coreInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coreInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coreInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coreInterface.EntityData.Children = types.NewOrderedMap()
    coreInterface.EntityData.Leafs = types.NewOrderedMap()
    coreInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", coreInterface.InterfaceName})
    coreInterface.EntityData.Leafs.Append("state", types.YLeaf{"State", coreInterface.State})

    coreInterface.EntityData.YListKeys = []string {}

    return &(coreInterface.EntityData)
}

// Evpn_Standby_EvpnGroups_EvpnGroup_AccessInterface
// EVPN Access Core Interface table
type Evpn_Standby_EvpnGroups_EvpnGroup_AccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface State. The type is ImStateEnum.
    State interface{}
}

func (accessInterface *Evpn_Standby_EvpnGroups_EvpnGroup_AccessInterface) GetEntityData() *types.CommonEntityData {
    accessInterface.EntityData.YFilter = accessInterface.YFilter
    accessInterface.EntityData.YangName = "access-interface"
    accessInterface.EntityData.BundleName = "cisco_ios_xr"
    accessInterface.EntityData.ParentYangName = "evpn-group"
    accessInterface.EntityData.SegmentPath = "access-interface"
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = types.NewOrderedMap()
    accessInterface.EntityData.Leafs = types.NewOrderedMap()
    accessInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterface.InterfaceName})
    accessInterface.EntityData.Leafs.Append("state", types.YLeaf{"State", accessInterface.State})

    accessInterface.EntityData.YListKeys = []string {}

    return &(accessInterface.EntityData)
}

// Evpn_Standby_RemoteShgs
// EVPN Remote SHG table
type Evpn_Standby_RemoteShgs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Remote SHG. The type is slice of Evpn_Standby_RemoteShgs_RemoteShg.
    RemoteShg []*Evpn_Standby_RemoteShgs_RemoteShg
}

func (remoteShgs *Evpn_Standby_RemoteShgs) GetEntityData() *types.CommonEntityData {
    remoteShgs.EntityData.YFilter = remoteShgs.YFilter
    remoteShgs.EntityData.YangName = "remote-shgs"
    remoteShgs.EntityData.BundleName = "cisco_ios_xr"
    remoteShgs.EntityData.ParentYangName = "standby"
    remoteShgs.EntityData.SegmentPath = "remote-shgs"
    remoteShgs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteShgs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteShgs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteShgs.EntityData.Children = types.NewOrderedMap()
    remoteShgs.EntityData.Children.Append("remote-shg", types.YChild{"RemoteShg", nil})
    for i := range remoteShgs.RemoteShg {
        remoteShgs.EntityData.Children.Append(types.GetSegmentPath(remoteShgs.RemoteShg[i]), types.YChild{"RemoteShg", remoteShgs.RemoteShg[i]})
    }
    remoteShgs.EntityData.Leafs = types.NewOrderedMap()

    remoteShgs.EntityData.YListKeys = []string {}

    return &(remoteShgs.EntityData)
}

// Evpn_Standby_RemoteShgs_RemoteShg
// EVPN Remote SHG
type Evpn_Standby_RemoteShgs_RemoteShg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Segment id. The type is slice of
    // Evpn_Standby_RemoteShgs_RemoteShg_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Standby_RemoteShgs_RemoteShg_EthernetSegmentIdentifier

    // Remote split horizon group labels. The type is slice of
    // Evpn_Standby_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel.
    RemoteSplitHorizonGroupLabel []*Evpn_Standby_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel
}

func (remoteShg *Evpn_Standby_RemoteShgs_RemoteShg) GetEntityData() *types.CommonEntityData {
    remoteShg.EntityData.YFilter = remoteShg.YFilter
    remoteShg.EntityData.YangName = "remote-shg"
    remoteShg.EntityData.BundleName = "cisco_ios_xr"
    remoteShg.EntityData.ParentYangName = "remote-shgs"
    remoteShg.EntityData.SegmentPath = "remote-shg"
    remoteShg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteShg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteShg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteShg.EntityData.Children = types.NewOrderedMap()
    remoteShg.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range remoteShg.EthernetSegmentIdentifier {
        remoteShg.EntityData.Children.Append(types.GetSegmentPath(remoteShg.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", remoteShg.EthernetSegmentIdentifier[i]})
    }
    remoteShg.EntityData.Children.Append("remote-split-horizon-group-label", types.YChild{"RemoteSplitHorizonGroupLabel", nil})
    for i := range remoteShg.RemoteSplitHorizonGroupLabel {
        remoteShg.EntityData.Children.Append(types.GetSegmentPath(remoteShg.RemoteSplitHorizonGroupLabel[i]), types.YChild{"RemoteSplitHorizonGroupLabel", remoteShg.RemoteSplitHorizonGroupLabel[i]})
    }
    remoteShg.EntityData.Leafs = types.NewOrderedMap()
    remoteShg.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", remoteShg.Esi1})
    remoteShg.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", remoteShg.Esi2})
    remoteShg.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", remoteShg.Esi3})
    remoteShg.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", remoteShg.Esi4})
    remoteShg.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", remoteShg.Esi5})

    remoteShg.EntityData.YListKeys = []string {}

    return &(remoteShg.EntityData)
}

// Evpn_Standby_RemoteShgs_RemoteShg_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Standby_RemoteShgs_RemoteShg_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Standby_RemoteShgs_RemoteShg_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "remote-shg"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Standby_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Standby_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Split horizon label associated with next-hop address. The type is
    // interface{} with range: 0..4294967295.
    Label interface{}
}

func (remoteSplitHorizonGroupLabel *Evpn_Standby_RemoteShgs_RemoteShg_RemoteSplitHorizonGroupLabel) GetEntityData() *types.CommonEntityData {
    remoteSplitHorizonGroupLabel.EntityData.YFilter = remoteSplitHorizonGroupLabel.YFilter
    remoteSplitHorizonGroupLabel.EntityData.YangName = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.BundleName = "cisco_ios_xr"
    remoteSplitHorizonGroupLabel.EntityData.ParentYangName = "remote-shg"
    remoteSplitHorizonGroupLabel.EntityData.SegmentPath = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteSplitHorizonGroupLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteSplitHorizonGroupLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteSplitHorizonGroupLabel.EntityData.Children = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", remoteSplitHorizonGroupLabel.NextHop})
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("label", types.YLeaf{"Label", remoteSplitHorizonGroupLabel.Label})

    remoteSplitHorizonGroupLabel.EntityData.YListKeys = []string {}

    return &(remoteSplitHorizonGroupLabel.EntityData)
}

// Evpn_Standby_Client
// L2VPN EVPN Client
type Evpn_Standby_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (client *Evpn_Standby_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "standby"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// Evpn_Standby_Igmps
// EVPN IGMP table
type Evpn_Standby_Igmps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Remote. The type is slice of Evpn_Standby_Igmps_Igmp.
    Igmp []*Evpn_Standby_Igmps_Igmp
}

func (igmps *Evpn_Standby_Igmps) GetEntityData() *types.CommonEntityData {
    igmps.EntityData.YFilter = igmps.YFilter
    igmps.EntityData.YangName = "igmps"
    igmps.EntityData.BundleName = "cisco_ios_xr"
    igmps.EntityData.ParentYangName = "standby"
    igmps.EntityData.SegmentPath = "igmps"
    igmps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igmps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igmps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igmps.EntityData.Children = types.NewOrderedMap()
    igmps.EntityData.Children.Append("igmp", types.YChild{"Igmp", nil})
    for i := range igmps.Igmp {
        igmps.EntityData.Children.Append(types.GetSegmentPath(igmps.Igmp[i]), types.YChild{"Igmp", igmps.Igmp[i]})
    }
    igmps.EntityData.Leafs = types.NewOrderedMap()

    igmps.EntityData.YListKeys = []string {}

    return &(igmps.EntityData)
}

// Evpn_Standby_Igmps_Igmp
// EVPN Remote
type Evpn_Standby_Igmps_Igmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Join=0, Leave=1. The type is interface{} with range: 0..4294967295.
    IsLeave interface{}

    // BP xcid. The type is interface{} with range: 0..4294967295.
    Bpxcid interface{}

    // EVI/BD. The type is interface{} with range: 0..4294967295.
    Evibd interface{}

    // Source IP Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SrcIp interface{}

    // Group IP Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GrpIp interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Segment Name. The type is string.
    EthernetSegmentName interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // BD id. The type is interface{} with range: 0..4294967295.
    BdId interface{}

    // Route Type. The type is EvpnIgmpMsg.
    RouteType interface{}

    // Source IP Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddr interface{}

    // Group IP Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GroupAddr interface{}

    // Ethernet Tag id. The type is interface{} with range: 0..4294967295.
    EthernetTagId interface{}

    // IGMP Version. The type is EvpnIgmpVersion.
    IgmpVersion interface{}

    // IGMP Group Type. The type is EvpnIgmpGrp.
    IgmpGroupType interface{}

    // Max Response Time. The type is interface{} with range: 0..255.
    MaXResponseTime interface{}

    // Resolved. The type is bool.
    Resolved interface{}

    // Source Info.
    SourceInfo Evpn_Standby_Igmps_Igmp_SourceInfo

    // Ethernet Segment id. The type is slice of
    // Evpn_Standby_Igmps_Igmp_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Standby_Igmps_Igmp_EthernetSegmentIdentifier

    // List of nexthop IPv6 addresses. The type is slice of
    // Evpn_Standby_Igmps_Igmp_NextHop.
    NextHop []*Evpn_Standby_Igmps_Igmp_NextHop
}

func (igmp *Evpn_Standby_Igmps_Igmp) GetEntityData() *types.CommonEntityData {
    igmp.EntityData.YFilter = igmp.YFilter
    igmp.EntityData.YangName = "igmp"
    igmp.EntityData.BundleName = "cisco_ios_xr"
    igmp.EntityData.ParentYangName = "igmps"
    igmp.EntityData.SegmentPath = "igmp"
    igmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igmp.EntityData.Children = types.NewOrderedMap()
    igmp.EntityData.Children.Append("source-info", types.YChild{"SourceInfo", &igmp.SourceInfo})
    igmp.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range igmp.EthernetSegmentIdentifier {
        igmp.EntityData.Children.Append(types.GetSegmentPath(igmp.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", igmp.EthernetSegmentIdentifier[i]})
    }
    igmp.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range igmp.NextHop {
        igmp.EntityData.Children.Append(types.GetSegmentPath(igmp.NextHop[i]), types.YChild{"NextHop", igmp.NextHop[i]})
    }
    igmp.EntityData.Leafs = types.NewOrderedMap()
    igmp.EntityData.Leafs.Append("is-leave", types.YLeaf{"IsLeave", igmp.IsLeave})
    igmp.EntityData.Leafs.Append("bpxcid", types.YLeaf{"Bpxcid", igmp.Bpxcid})
    igmp.EntityData.Leafs.Append("evibd", types.YLeaf{"Evibd", igmp.Evibd})
    igmp.EntityData.Leafs.Append("src-ip", types.YLeaf{"SrcIp", igmp.SrcIp})
    igmp.EntityData.Leafs.Append("grp-ip", types.YLeaf{"GrpIp", igmp.GrpIp})
    igmp.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", igmp.Esi1})
    igmp.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", igmp.Esi2})
    igmp.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", igmp.Esi3})
    igmp.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", igmp.Esi4})
    igmp.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", igmp.Esi5})
    igmp.EntityData.Leafs.Append("ethernet-segment-name", types.YLeaf{"EthernetSegmentName", igmp.EthernetSegmentName})
    igmp.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", igmp.Evi})
    igmp.EntityData.Leafs.Append("bd-id", types.YLeaf{"BdId", igmp.BdId})
    igmp.EntityData.Leafs.Append("route-type", types.YLeaf{"RouteType", igmp.RouteType})
    igmp.EntityData.Leafs.Append("source-addr", types.YLeaf{"SourceAddr", igmp.SourceAddr})
    igmp.EntityData.Leafs.Append("group-addr", types.YLeaf{"GroupAddr", igmp.GroupAddr})
    igmp.EntityData.Leafs.Append("ethernet-tag-id", types.YLeaf{"EthernetTagId", igmp.EthernetTagId})
    igmp.EntityData.Leafs.Append("igmp-version", types.YLeaf{"IgmpVersion", igmp.IgmpVersion})
    igmp.EntityData.Leafs.Append("igmp-group-type", types.YLeaf{"IgmpGroupType", igmp.IgmpGroupType})
    igmp.EntityData.Leafs.Append("ma-x-response-time", types.YLeaf{"MaXResponseTime", igmp.MaXResponseTime})
    igmp.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", igmp.Resolved})

    igmp.EntityData.YListKeys = []string {}

    return &(igmp.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo
// Source Info
type Evpn_Standby_Igmps_Igmp_SourceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is EvpnIgmpSource.
    Type interface{}

    // remote info. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteInfo interface{}

    // local info.
    LocalInfo Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo
}

func (sourceInfo *Evpn_Standby_Igmps_Igmp_SourceInfo) GetEntityData() *types.CommonEntityData {
    sourceInfo.EntityData.YFilter = sourceInfo.YFilter
    sourceInfo.EntityData.YangName = "source-info"
    sourceInfo.EntityData.BundleName = "cisco_ios_xr"
    sourceInfo.EntityData.ParentYangName = "igmp"
    sourceInfo.EntityData.SegmentPath = "source-info"
    sourceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceInfo.EntityData.Children = types.NewOrderedMap()
    sourceInfo.EntityData.Children.Append("local-info", types.YChild{"LocalInfo", &sourceInfo.LocalInfo})
    sourceInfo.EntityData.Leafs = types.NewOrderedMap()
    sourceInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", sourceInfo.Type})
    sourceInfo.EntityData.Leafs.Append("remote-info", types.YLeaf{"RemoteInfo", sourceInfo.RemoteInfo})

    sourceInfo.EntityData.YListKeys = []string {}

    return &(sourceInfo.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo
// local info
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with length: 0..81.
    Name interface{}

    // Interface MTU. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Payload bytes. The type is interface{} with range: 0..65535. Units are
    // byte.
    PayloadBytes interface{}

    // Interface parameters.
    Parameters Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters
}

func (localInfo *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo) GetEntityData() *types.CommonEntityData {
    localInfo.EntityData.YFilter = localInfo.YFilter
    localInfo.EntityData.YangName = "local-info"
    localInfo.EntityData.BundleName = "cisco_ios_xr"
    localInfo.EntityData.ParentYangName = "source-info"
    localInfo.EntityData.SegmentPath = "local-info"
    localInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localInfo.EntityData.Children = types.NewOrderedMap()
    localInfo.EntityData.Children.Append("parameters", types.YChild{"Parameters", &localInfo.Parameters})
    localInfo.EntityData.Leafs = types.NewOrderedMap()
    localInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", localInfo.Name})
    localInfo.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", localInfo.Mtu})
    localInfo.EntityData.Leafs.Append("payload-bytes", types.YLeaf{"PayloadBytes", localInfo.PayloadBytes})

    localInfo.EntityData.YListKeys = []string {}

    return &(localInfo.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters
// Interface parameters
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is L2vpnInterface.
    Type interface{}

    // Ethernet.
    Ethernet Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet

    // VLAN.
    Vlan Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan

    // TDM.
    Tdm Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm

    // ATM.
    Atm Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm

    // Frame Relay.
    Fr Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr

    // PW Ether.
    PseudowireEther Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther

    // PW IW.
    PseudowireIw Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw
}

func (parameters *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters) GetEntityData() *types.CommonEntityData {
    parameters.EntityData.YFilter = parameters.YFilter
    parameters.EntityData.YangName = "parameters"
    parameters.EntityData.BundleName = "cisco_ios_xr"
    parameters.EntityData.ParentYangName = "local-info"
    parameters.EntityData.SegmentPath = "parameters"
    parameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parameters.EntityData.Children = types.NewOrderedMap()
    parameters.EntityData.Children.Append("ethernet", types.YChild{"Ethernet", &parameters.Ethernet})
    parameters.EntityData.Children.Append("vlan", types.YChild{"Vlan", &parameters.Vlan})
    parameters.EntityData.Children.Append("tdm", types.YChild{"Tdm", &parameters.Tdm})
    parameters.EntityData.Children.Append("atm", types.YChild{"Atm", &parameters.Atm})
    parameters.EntityData.Children.Append("fr", types.YChild{"Fr", &parameters.Fr})
    parameters.EntityData.Children.Append("pseudowire-ether", types.YChild{"PseudowireEther", &parameters.PseudowireEther})
    parameters.EntityData.Children.Append("pseudowire-iw", types.YChild{"PseudowireIw", &parameters.PseudowireIw})
    parameters.EntityData.Leafs = types.NewOrderedMap()
    parameters.EntityData.Leafs.Append("type", types.YLeaf{"Type", parameters.Type})

    parameters.EntityData.YListKeys = []string {}

    return &(parameters.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet
// Ethernet
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XConnect tags. The type is interface{} with range: 0..255.
    XconnectTags interface{}
}

func (ethernet *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "cisco_ios_xr"
    ethernet.EntityData.ParentYangName = "parameters"
    ethernet.EntityData.SegmentPath = "ethernet"
    ethernet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernet.EntityData.Children = types.NewOrderedMap()
    ethernet.EntityData.Leafs = types.NewOrderedMap()
    ethernet.EntityData.Leafs.Append("xconnect-tags", types.YLeaf{"XconnectTags", ethernet.XconnectTags})

    ethernet.EntityData.YListKeys = []string {}

    return &(ethernet.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan
// VLAN
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XConnect tags. The type is interface{} with range: 0..255.
    XconnectTags interface{}

    // VLAN rewrite tag. The type is interface{} with range: 0..65535.
    VlanRewriteTag interface{}

    // Simple EFP. The type is interface{} with range: 0..255.
    SimpleEfp interface{}

    // Encapsulation Type. The type is interface{} with range: 0..255.
    EncapsulationType interface{}

    // Outer Tag. The type is interface{} with range: 0..65535.
    OuterTag interface{}

    // Rewrite Tags. The type is slice of
    // Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag.
    RewriteTag []*Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag

    // vlan range. The type is slice of
    // Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange.
    VlanRange []*Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange
}

func (vlan *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "cisco_ios_xr"
    vlan.EntityData.ParentYangName = "parameters"
    vlan.EntityData.SegmentPath = "vlan"
    vlan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlan.EntityData.Children = types.NewOrderedMap()
    vlan.EntityData.Children.Append("rewrite-tag", types.YChild{"RewriteTag", nil})
    for i := range vlan.RewriteTag {
        vlan.EntityData.Children.Append(types.GetSegmentPath(vlan.RewriteTag[i]), types.YChild{"RewriteTag", vlan.RewriteTag[i]})
    }
    vlan.EntityData.Children.Append("vlan-range", types.YChild{"VlanRange", nil})
    for i := range vlan.VlanRange {
        vlan.EntityData.Children.Append(types.GetSegmentPath(vlan.VlanRange[i]), types.YChild{"VlanRange", vlan.VlanRange[i]})
    }
    vlan.EntityData.Leafs = types.NewOrderedMap()
    vlan.EntityData.Leafs.Append("xconnect-tags", types.YLeaf{"XconnectTags", vlan.XconnectTags})
    vlan.EntityData.Leafs.Append("vlan-rewrite-tag", types.YLeaf{"VlanRewriteTag", vlan.VlanRewriteTag})
    vlan.EntityData.Leafs.Append("simple-efp", types.YLeaf{"SimpleEfp", vlan.SimpleEfp})
    vlan.EntityData.Leafs.Append("encapsulation-type", types.YLeaf{"EncapsulationType", vlan.EncapsulationType})
    vlan.EntityData.Leafs.Append("outer-tag", types.YLeaf{"OuterTag", vlan.OuterTag})

    vlan.EntityData.YListKeys = []string {}

    return &(vlan.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag
// Rewrite Tags
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535.
    Entry interface{}
}

func (rewriteTag *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_RewriteTag) GetEntityData() *types.CommonEntityData {
    rewriteTag.EntityData.YFilter = rewriteTag.YFilter
    rewriteTag.EntityData.YangName = "rewrite-tag"
    rewriteTag.EntityData.BundleName = "cisco_ios_xr"
    rewriteTag.EntityData.ParentYangName = "vlan"
    rewriteTag.EntityData.SegmentPath = "rewrite-tag"
    rewriteTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rewriteTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rewriteTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rewriteTag.EntityData.Children = types.NewOrderedMap()
    rewriteTag.EntityData.Leafs = types.NewOrderedMap()
    rewriteTag.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", rewriteTag.Entry})

    rewriteTag.EntityData.YListKeys = []string {}

    return &(rewriteTag.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange
// vlan range
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lower. The type is interface{} with range: 0..65535.
    Lower interface{}

    // Upper. The type is interface{} with range: 0..65535.
    Upper interface{}
}

func (vlanRange *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Vlan_VlanRange) GetEntityData() *types.CommonEntityData {
    vlanRange.EntityData.YFilter = vlanRange.YFilter
    vlanRange.EntityData.YangName = "vlan-range"
    vlanRange.EntityData.BundleName = "cisco_ios_xr"
    vlanRange.EntityData.ParentYangName = "vlan"
    vlanRange.EntityData.SegmentPath = "vlan-range"
    vlanRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanRange.EntityData.Children = types.NewOrderedMap()
    vlanRange.EntityData.Leafs = types.NewOrderedMap()
    vlanRange.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", vlanRange.Lower})
    vlanRange.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", vlanRange.Upper})

    vlanRange.EntityData.YListKeys = []string {}

    return &(vlanRange.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm
// TDM
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeslots separated by , or - from 1 to 31. : indicates individual timeslot
    // and - represents a range.E.g. 1-3,5 represents timeslots 1, 2, 3, and 5.
    // The type is string.
    TimeslotGroup interface{}

    // Timeslot rate in units of Kbps. The type is interface{} with range: 0..255.
    // Units are kbit/s.
    TimeslotRate interface{}

    // TDM mode. The type is L2vpnTdmMode.
    TdmMode interface{}

    // TDM options.
    TdmOptions Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions
}

func (tdm *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm) GetEntityData() *types.CommonEntityData {
    tdm.EntityData.YFilter = tdm.YFilter
    tdm.EntityData.YangName = "tdm"
    tdm.EntityData.BundleName = "cisco_ios_xr"
    tdm.EntityData.ParentYangName = "parameters"
    tdm.EntityData.SegmentPath = "tdm"
    tdm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tdm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tdm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tdm.EntityData.Children = types.NewOrderedMap()
    tdm.EntityData.Children.Append("tdm-options", types.YChild{"TdmOptions", &tdm.TdmOptions})
    tdm.EntityData.Leafs = types.NewOrderedMap()
    tdm.EntityData.Leafs.Append("timeslot-group", types.YLeaf{"TimeslotGroup", tdm.TimeslotGroup})
    tdm.EntityData.Leafs.Append("timeslot-rate", types.YLeaf{"TimeslotRate", tdm.TimeslotRate})
    tdm.EntityData.Leafs.Append("tdm-mode", types.YLeaf{"TdmMode", tdm.TdmMode})

    tdm.EntityData.YListKeys = []string {}

    return &(tdm.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions
// TDM options
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TDM payload bytes. The type is interface{} with range: 0..65535. Units are
    // byte.
    PayloadBytes interface{}

    // TDM bit rate in units of Kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    BitRate interface{}

    // RTP header. The type is L2vpnTdmRtpOption.
    Rtp interface{}

    // TDM Timestamping mode. The type is L2vpnTimeStampMode.
    TimestampMode interface{}

    // Signalling packets. The type is interface{} with range: 0..255.
    SignallingPackets interface{}

    // CAS. The type is interface{} with range: 0..255.
    Cas interface{}

    // RTP header payload type. The type is interface{} with range: 0..255.
    RtpHeaderPayloadType interface{}

    // Timestamping clock frequency in units of 8Khz. The type is interface{} with
    // range: 0..65535.
    TimestampClockFreq interface{}

    // Synchronization Source identifier. The type is interface{} with range:
    // 0..4294967295.
    Ssrc interface{}
}

func (tdmOptions *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Tdm_TdmOptions) GetEntityData() *types.CommonEntityData {
    tdmOptions.EntityData.YFilter = tdmOptions.YFilter
    tdmOptions.EntityData.YangName = "tdm-options"
    tdmOptions.EntityData.BundleName = "cisco_ios_xr"
    tdmOptions.EntityData.ParentYangName = "tdm"
    tdmOptions.EntityData.SegmentPath = "tdm-options"
    tdmOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tdmOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tdmOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tdmOptions.EntityData.Children = types.NewOrderedMap()
    tdmOptions.EntityData.Leafs = types.NewOrderedMap()
    tdmOptions.EntityData.Leafs.Append("payload-bytes", types.YLeaf{"PayloadBytes", tdmOptions.PayloadBytes})
    tdmOptions.EntityData.Leafs.Append("bit-rate", types.YLeaf{"BitRate", tdmOptions.BitRate})
    tdmOptions.EntityData.Leafs.Append("rtp", types.YLeaf{"Rtp", tdmOptions.Rtp})
    tdmOptions.EntityData.Leafs.Append("timestamp-mode", types.YLeaf{"TimestampMode", tdmOptions.TimestampMode})
    tdmOptions.EntityData.Leafs.Append("signalling-packets", types.YLeaf{"SignallingPackets", tdmOptions.SignallingPackets})
    tdmOptions.EntityData.Leafs.Append("cas", types.YLeaf{"Cas", tdmOptions.Cas})
    tdmOptions.EntityData.Leafs.Append("rtp-header-payload-type", types.YLeaf{"RtpHeaderPayloadType", tdmOptions.RtpHeaderPayloadType})
    tdmOptions.EntityData.Leafs.Append("timestamp-clock-freq", types.YLeaf{"TimestampClockFreq", tdmOptions.TimestampClockFreq})
    tdmOptions.EntityData.Leafs.Append("ssrc", types.YLeaf{"Ssrc", tdmOptions.Ssrc})

    tdmOptions.EntityData.YListKeys = []string {}

    return &(tdmOptions.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm
// ATM
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max number of cells packed. The type is interface{} with range: 0..65535.
    MaximumNumberCellsPacked interface{}

    // Max number of cells unpacked. The type is interface{} with range: 0..65535.
    MaximumNumberCellsUnPacked interface{}

    // ATM mode. The type is L2vpnAtmMode.
    AtmMode interface{}

    // Virtual path identifier. The type is interface{} with range: 0..65535.
    Vpi interface{}

    // Virtual channel identifier. The type is interface{} with range: 0..65535.
    Vci interface{}
}

func (atm *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Atm) GetEntityData() *types.CommonEntityData {
    atm.EntityData.YFilter = atm.YFilter
    atm.EntityData.YangName = "atm"
    atm.EntityData.BundleName = "cisco_ios_xr"
    atm.EntityData.ParentYangName = "parameters"
    atm.EntityData.SegmentPath = "atm"
    atm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    atm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    atm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    atm.EntityData.Children = types.NewOrderedMap()
    atm.EntityData.Leafs = types.NewOrderedMap()
    atm.EntityData.Leafs.Append("maximum-number-cells-packed", types.YLeaf{"MaximumNumberCellsPacked", atm.MaximumNumberCellsPacked})
    atm.EntityData.Leafs.Append("maximum-number-cells-un-packed", types.YLeaf{"MaximumNumberCellsUnPacked", atm.MaximumNumberCellsUnPacked})
    atm.EntityData.Leafs.Append("atm-mode", types.YLeaf{"AtmMode", atm.AtmMode})
    atm.EntityData.Leafs.Append("vpi", types.YLeaf{"Vpi", atm.Vpi})
    atm.EntityData.Leafs.Append("vci", types.YLeaf{"Vci", atm.Vci})

    atm.EntityData.YListKeys = []string {}

    return &(atm.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr
// Frame Relay
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Frame Relay mode. The type is L2vpnFrMode.
    FrMode interface{}

    // Data-link connection identifier. The type is interface{} with range:
    // 0..4294967295.
    Dlci interface{}
}

func (fr *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_Fr) GetEntityData() *types.CommonEntityData {
    fr.EntityData.YFilter = fr.YFilter
    fr.EntityData.YangName = "fr"
    fr.EntityData.BundleName = "cisco_ios_xr"
    fr.EntityData.ParentYangName = "parameters"
    fr.EntityData.SegmentPath = "fr"
    fr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fr.EntityData.Children = types.NewOrderedMap()
    fr.EntityData.Leafs = types.NewOrderedMap()
    fr.EntityData.Leafs.Append("fr-mode", types.YLeaf{"FrMode", fr.FrMode})
    fr.EntityData.Leafs.Append("dlci", types.YLeaf{"Dlci", fr.Dlci})

    fr.EntityData.YListKeys = []string {}

    return &(fr.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther
// PW Ether
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is this Interface list valid. The type is bool.
    IsValid interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Interface list data.
    InterfaceList Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList
}

func (pseudowireEther *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther) GetEntityData() *types.CommonEntityData {
    pseudowireEther.EntityData.YFilter = pseudowireEther.YFilter
    pseudowireEther.EntityData.YangName = "pseudowire-ether"
    pseudowireEther.EntityData.BundleName = "cisco_ios_xr"
    pseudowireEther.EntityData.ParentYangName = "parameters"
    pseudowireEther.EntityData.SegmentPath = "pseudowire-ether"
    pseudowireEther.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireEther.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireEther.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireEther.EntityData.Children = types.NewOrderedMap()
    pseudowireEther.EntityData.Children.Append("interface-list", types.YChild{"InterfaceList", &pseudowireEther.InterfaceList})
    pseudowireEther.EntityData.Leafs = types.NewOrderedMap()
    pseudowireEther.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", pseudowireEther.IsValid})
    pseudowireEther.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", pseudowireEther.InternalLabel})

    pseudowireEther.EntityData.YListKeys = []string {}

    return &(pseudowireEther.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList
// Interface list data
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface-list name. The type is string with length: 0..33.
    InterfaceListName interface{}

    // Interface internal ID. The type is interface{} with range: 0..4294967295.
    InterfaceListId interface{}

    // Interfaces. The type is slice of
    // Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface.
    Interface []*Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface
}

func (interfaceList *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList) GetEntityData() *types.CommonEntityData {
    interfaceList.EntityData.YFilter = interfaceList.YFilter
    interfaceList.EntityData.YangName = "interface-list"
    interfaceList.EntityData.BundleName = "cisco_ios_xr"
    interfaceList.EntityData.ParentYangName = "pseudowire-ether"
    interfaceList.EntityData.SegmentPath = "interface-list"
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = types.NewOrderedMap()
    interfaceList.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceList.Interface {
        interfaceList.EntityData.Children.Append(types.GetSegmentPath(interfaceList.Interface[i]), types.YChild{"Interface", interfaceList.Interface[i]})
    }
    interfaceList.EntityData.Leafs = types.NewOrderedMap()
    interfaceList.EntityData.Leafs.Append("interface-list-name", types.YLeaf{"InterfaceListName", interfaceList.InterfaceListName})
    interfaceList.EntityData.Leafs.Append("interface-list-id", types.YLeaf{"InterfaceListId", interfaceList.InterfaceListId})

    interfaceList.EntityData.YListKeys = []string {}

    return &(interfaceList.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface
// Interfaces
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Replicate status. The type is IflistRepStatus.
    ReplicateStatus interface{}
}

func (self *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireEther_InterfaceList_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interface-list"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("replicate-status", types.YLeaf{"ReplicateStatus", self.ReplicateStatus})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw
// PW IW
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is this Interface list valid. The type is bool.
    IsValid interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Interface list data.
    InterfaceList Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList
}

func (pseudowireIw *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw) GetEntityData() *types.CommonEntityData {
    pseudowireIw.EntityData.YFilter = pseudowireIw.YFilter
    pseudowireIw.EntityData.YangName = "pseudowire-iw"
    pseudowireIw.EntityData.BundleName = "cisco_ios_xr"
    pseudowireIw.EntityData.ParentYangName = "parameters"
    pseudowireIw.EntityData.SegmentPath = "pseudowire-iw"
    pseudowireIw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireIw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireIw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireIw.EntityData.Children = types.NewOrderedMap()
    pseudowireIw.EntityData.Children.Append("interface-list", types.YChild{"InterfaceList", &pseudowireIw.InterfaceList})
    pseudowireIw.EntityData.Leafs = types.NewOrderedMap()
    pseudowireIw.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", pseudowireIw.IsValid})
    pseudowireIw.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", pseudowireIw.InternalLabel})

    pseudowireIw.EntityData.YListKeys = []string {}

    return &(pseudowireIw.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList
// Interface list data
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface-list name. The type is string with length: 0..33.
    InterfaceListName interface{}

    // Interface internal ID. The type is interface{} with range: 0..4294967295.
    InterfaceListId interface{}

    // Interfaces. The type is slice of
    // Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface.
    Interface []*Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface
}

func (interfaceList *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList) GetEntityData() *types.CommonEntityData {
    interfaceList.EntityData.YFilter = interfaceList.YFilter
    interfaceList.EntityData.YangName = "interface-list"
    interfaceList.EntityData.BundleName = "cisco_ios_xr"
    interfaceList.EntityData.ParentYangName = "pseudowire-iw"
    interfaceList.EntityData.SegmentPath = "interface-list"
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = types.NewOrderedMap()
    interfaceList.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceList.Interface {
        interfaceList.EntityData.Children.Append(types.GetSegmentPath(interfaceList.Interface[i]), types.YChild{"Interface", interfaceList.Interface[i]})
    }
    interfaceList.EntityData.Leafs = types.NewOrderedMap()
    interfaceList.EntityData.Leafs.Append("interface-list-name", types.YLeaf{"InterfaceListName", interfaceList.InterfaceListName})
    interfaceList.EntityData.Leafs.Append("interface-list-id", types.YLeaf{"InterfaceListId", interfaceList.InterfaceListId})

    interfaceList.EntityData.YListKeys = []string {}

    return &(interfaceList.EntityData)
}

// Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface
// Interfaces
type Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Replicate status. The type is IflistRepStatus.
    ReplicateStatus interface{}
}

func (self *Evpn_Standby_Igmps_Igmp_SourceInfo_LocalInfo_Parameters_PseudowireIw_InterfaceList_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interface-list"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("replicate-status", types.YLeaf{"ReplicateStatus", self.ReplicateStatus})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Evpn_Standby_Igmps_Igmp_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Standby_Igmps_Igmp_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Standby_Igmps_Igmp_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "igmp"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Standby_Igmps_Igmp_NextHop
// List of nexthop IPv6 addresses
type Evpn_Standby_Igmps_Igmp_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // DF Dont Premption. The type is bool.
    DfDontPrempt interface{}

    // DF Election Mode Configured. The type is interface{} with range: 0..255.
    DfType interface{}

    // DF Election Preference Set. The type is interface{} with range: 0..65535.
    DfPref interface{}
}

func (nextHop *Evpn_Standby_Igmps_Igmp_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "igmp"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", nextHop.NextHop})
    nextHop.EntityData.Leafs.Append("df-dont-prempt", types.YLeaf{"DfDontPrempt", nextHop.DfDontPrempt})
    nextHop.EntityData.Leafs.Append("df-type", types.YLeaf{"DfType", nextHop.DfType})
    nextHop.EntityData.Leafs.Append("df-pref", types.YLeaf{"DfPref", nextHop.DfPref})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// Evpn_Standby_Evis
// L2VPN EVPN EVI Table
type Evpn_Standby_Evis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Entry. The type is slice of Evpn_Standby_Evis_Evi.
    Evi []*Evpn_Standby_Evis_Evi
}

func (evis *Evpn_Standby_Evis) GetEntityData() *types.CommonEntityData {
    evis.EntityData.YFilter = evis.YFilter
    evis.EntityData.YangName = "evis"
    evis.EntityData.BundleName = "cisco_ios_xr"
    evis.EntityData.ParentYangName = "standby"
    evis.EntityData.SegmentPath = "evis"
    evis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evis.EntityData.Children = types.NewOrderedMap()
    evis.EntityData.Children.Append("evi", types.YChild{"Evi", nil})
    for i := range evis.Evi {
        evis.EntityData.Children.Append(types.GetSegmentPath(evis.Evi[i]), types.YChild{"Evi", evis.Evi[i]})
    }
    evis.EntityData.Leafs = types.NewOrderedMap()

    evis.EntityData.YListKeys = []string {}

    return &(evis.EntityData)
}

// Evpn_Standby_Evis_Evi
// L2VPN EVPN EVI Entry
type Evpn_Standby_Evis_Evi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evi *Evpn_Standby_Evis_Evi) GetEntityData() *types.CommonEntityData {
    evi.EntityData.YFilter = evi.YFilter
    evi.EntityData.YangName = "evi"
    evi.EntityData.BundleName = "cisco_ios_xr"
    evi.EntityData.ParentYangName = "evis"
    evi.EntityData.SegmentPath = "evi"
    evi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evi.EntityData.Children = types.NewOrderedMap()
    evi.EntityData.Leafs = types.NewOrderedMap()
    evi.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", evi.Evi})
    evi.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", evi.Encapsulation})
    evi.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evi.EthernetVpnId})
    evi.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evi.EncapsulationXr})
    evi.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evi.BdName})
    evi.EntityData.Leafs.Append("type", types.YLeaf{"Type", evi.Type})

    evi.EntityData.YListKeys = []string {}

    return &(evi.EntityData)
}

// Evpn_Standby_Summary
// L2VPN EVPN Summary
type Evpn_Standby_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Router ID. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // BGP AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // Number of EVI DB Entries. The type is interface{} with range:
    // 0..4294967295.
    EvIs interface{}

    // Number of Tunnel Endpoint DB Entries. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpoints interface{}

    // Number of Local MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalMacRoutes interface{}

    // Number of Local IPv4 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalIpv4MacRoutes interface{}

    // Number of Local IPv6 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalIpv6MacRoutes interface{}

    // Number of ES:Global MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    EsGlobalMacRoutes interface{}

    // Number of Remote MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteMacRoutes interface{}

    // Number of Remote Soo MAC Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteSooMacRoutes interface{}

    // Number of Remote IPv4 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteIpv4MacRoutes interface{}

    // Number of Remote IPv6 MAC-IP Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteIpv6MacRoutes interface{}

    // Number of Local IMCAST Routes. The type is interface{} with range:
    // 0..4294967295.
    LocalImcastRoutes interface{}

    // Number of Remote IMCAST Routes. The type is interface{} with range:
    // 0..4294967295.
    RemoteImcastRoutes interface{}

    // Number of Internal Labels. The type is interface{} with range:
    // 0..4294967295.
    Labels interface{}

    // Number of ES Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    EsEntries interface{}

    // Number of neighbor Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    NeighborEntries interface{}

    // Number of Local EAD Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    LocalEadRoutes interface{}

    // Number of Remote EAD Entries in DB. The type is interface{} with range:
    // 0..4294967295.
    RemoteEadRoutes interface{}

    // Global Source MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    GlobalSourceMac interface{}

    // EVPN ES Peering Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    PeeringTime interface{}

    // EVPN ES Recovery Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    RecoveryTime interface{}

    // EVPN ES Carving Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    CarvingTime interface{}

    // Number of moves within the move interval before locking the MAC. The type
    // is interface{} with range: 0..4294967295.
    MacSecureMoveCount interface{}

    // Interval to watch for subsequent mac moves before locking the MAC. The type
    // is interface{} with range: 0..4294967295.
    MacSecureMoveInterval interface{}

    // Length of time to lock the mac after a MAC security violation. The type is
    // interface{} with range: 0..4294967295.
    MacSecureFreezeTime interface{}

    // Number of times to retry after a MAC un-freezes. The type is interface{}
    // with range: 0..4294967295.
    MacSecureRetryCount interface{}

    // EVPN Node Cost-out. The type is bool.
    CostOut interface{}

    // EVPN Node startup cost-in Time (minutes). The type is interface{} with
    // range: 0..4294967295. Units are minute.
    StartupCostInTime interface{}

    // Send to L2RIB Throttled. The type is bool.
    L2ribThrottle interface{}

    // Logging EVPN Designated Forwarder changes enabled. The type is bool.
    LoggingDfElectionEnabled interface{}
}

func (summary *Evpn_Standby_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "standby"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", summary.RouterId})
    summary.EntityData.Leafs.Append("as", types.YLeaf{"As", summary.As})
    summary.EntityData.Leafs.Append("ev-is", types.YLeaf{"EvIs", summary.EvIs})
    summary.EntityData.Leafs.Append("tunnel-endpoints", types.YLeaf{"TunnelEndpoints", summary.TunnelEndpoints})
    summary.EntityData.Leafs.Append("local-mac-routes", types.YLeaf{"LocalMacRoutes", summary.LocalMacRoutes})
    summary.EntityData.Leafs.Append("local-ipv4-mac-routes", types.YLeaf{"LocalIpv4MacRoutes", summary.LocalIpv4MacRoutes})
    summary.EntityData.Leafs.Append("local-ipv6-mac-routes", types.YLeaf{"LocalIpv6MacRoutes", summary.LocalIpv6MacRoutes})
    summary.EntityData.Leafs.Append("es-global-mac-routes", types.YLeaf{"EsGlobalMacRoutes", summary.EsGlobalMacRoutes})
    summary.EntityData.Leafs.Append("remote-mac-routes", types.YLeaf{"RemoteMacRoutes", summary.RemoteMacRoutes})
    summary.EntityData.Leafs.Append("remote-soo-mac-routes", types.YLeaf{"RemoteSooMacRoutes", summary.RemoteSooMacRoutes})
    summary.EntityData.Leafs.Append("remote-ipv4-mac-routes", types.YLeaf{"RemoteIpv4MacRoutes", summary.RemoteIpv4MacRoutes})
    summary.EntityData.Leafs.Append("remote-ipv6-mac-routes", types.YLeaf{"RemoteIpv6MacRoutes", summary.RemoteIpv6MacRoutes})
    summary.EntityData.Leafs.Append("local-imcast-routes", types.YLeaf{"LocalImcastRoutes", summary.LocalImcastRoutes})
    summary.EntityData.Leafs.Append("remote-imcast-routes", types.YLeaf{"RemoteImcastRoutes", summary.RemoteImcastRoutes})
    summary.EntityData.Leafs.Append("labels", types.YLeaf{"Labels", summary.Labels})
    summary.EntityData.Leafs.Append("es-entries", types.YLeaf{"EsEntries", summary.EsEntries})
    summary.EntityData.Leafs.Append("neighbor-entries", types.YLeaf{"NeighborEntries", summary.NeighborEntries})
    summary.EntityData.Leafs.Append("local-ead-routes", types.YLeaf{"LocalEadRoutes", summary.LocalEadRoutes})
    summary.EntityData.Leafs.Append("remote-ead-routes", types.YLeaf{"RemoteEadRoutes", summary.RemoteEadRoutes})
    summary.EntityData.Leafs.Append("global-source-mac", types.YLeaf{"GlobalSourceMac", summary.GlobalSourceMac})
    summary.EntityData.Leafs.Append("peering-time", types.YLeaf{"PeeringTime", summary.PeeringTime})
    summary.EntityData.Leafs.Append("recovery-time", types.YLeaf{"RecoveryTime", summary.RecoveryTime})
    summary.EntityData.Leafs.Append("carving-time", types.YLeaf{"CarvingTime", summary.CarvingTime})
    summary.EntityData.Leafs.Append("mac-secure-move-count", types.YLeaf{"MacSecureMoveCount", summary.MacSecureMoveCount})
    summary.EntityData.Leafs.Append("mac-secure-move-interval", types.YLeaf{"MacSecureMoveInterval", summary.MacSecureMoveInterval})
    summary.EntityData.Leafs.Append("mac-secure-freeze-time", types.YLeaf{"MacSecureFreezeTime", summary.MacSecureFreezeTime})
    summary.EntityData.Leafs.Append("mac-secure-retry-count", types.YLeaf{"MacSecureRetryCount", summary.MacSecureRetryCount})
    summary.EntityData.Leafs.Append("cost-out", types.YLeaf{"CostOut", summary.CostOut})
    summary.EntityData.Leafs.Append("startup-cost-in-time", types.YLeaf{"StartupCostInTime", summary.StartupCostInTime})
    summary.EntityData.Leafs.Append("l2rib-throttle", types.YLeaf{"L2ribThrottle", summary.L2ribThrottle})
    summary.EntityData.Leafs.Append("logging-df-election-enabled", types.YLeaf{"LoggingDfElectionEnabled", summary.LoggingDfElectionEnabled})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Evpn_Standby_EviDetail
// L2VPN EVI Detail Table
type Evpn_Standby_EviDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info Elements.
    Elements Evpn_Standby_EviDetail_Elements

    // Container for all EVI detail info.
    EviChildren Evpn_Standby_EviDetail_EviChildren
}

func (eviDetail *Evpn_Standby_EviDetail) GetEntityData() *types.CommonEntityData {
    eviDetail.EntityData.YFilter = eviDetail.YFilter
    eviDetail.EntityData.YangName = "evi-detail"
    eviDetail.EntityData.BundleName = "cisco_ios_xr"
    eviDetail.EntityData.ParentYangName = "standby"
    eviDetail.EntityData.SegmentPath = "evi-detail"
    eviDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviDetail.EntityData.Children = types.NewOrderedMap()
    eviDetail.EntityData.Children.Append("elements", types.YChild{"Elements", &eviDetail.Elements})
    eviDetail.EntityData.Children.Append("evi-children", types.YChild{"EviChildren", &eviDetail.EviChildren})
    eviDetail.EntityData.Leafs = types.NewOrderedMap()

    eviDetail.EntityData.YListKeys = []string {}

    return &(eviDetail.EntityData)
}

// Evpn_Standby_EviDetail_Elements
// EVI BGP RT Detail Info Elements
type Evpn_Standby_EviDetail_Elements struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info. The type is slice of
    // Evpn_Standby_EviDetail_Elements_Element.
    Element []*Evpn_Standby_EviDetail_Elements_Element
}

func (elements *Evpn_Standby_EviDetail_Elements) GetEntityData() *types.CommonEntityData {
    elements.EntityData.YFilter = elements.YFilter
    elements.EntityData.YangName = "elements"
    elements.EntityData.BundleName = "cisco_ios_xr"
    elements.EntityData.ParentYangName = "evi-detail"
    elements.EntityData.SegmentPath = "elements"
    elements.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    elements.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    elements.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    elements.EntityData.Children = types.NewOrderedMap()
    elements.EntityData.Children.Append("element", types.YChild{"Element", nil})
    for i := range elements.Element {
        elements.EntityData.Children.Append(types.GetSegmentPath(elements.Element[i]), types.YChild{"Element", elements.Element[i]})
    }
    elements.EntityData.Leafs = types.NewOrderedMap()

    elements.EntityData.YListKeys = []string {}

    return &(elements.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element
// EVI BGP RT Detail Info
type Evpn_Standby_EviDetail_Elements_Element struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // EVPN Instance encapsulation. The type is interface{} with range: 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // EVI description. The type is string.
    Description interface{}

    // Unicast Label. The type is interface{} with range: 0..4294967295.
    UnicastLabel interface{}

    // Multicast Label. The type is interface{} with range: 0..4294967295.
    MulticastLabel interface{}

    // Control-Word Disable. The type is bool.
    CwDisable interface{}

    // Table-policy Name. The type is string.
    TablePolicyName interface{}

    // Forward Class attribute. The type is interface{} with range: 0..255.
    ForwardClass interface{}

    // Is Import RT None set. The type is bool.
    RtImportBlockSet interface{}

    // Is Export RT None set. The type is bool.
    RtExportBlockSet interface{}

    // Advertise MAC-only routes on this EVI. The type is bool.
    AdvertiseMac interface{}

    // Advertise BVI MACs routes on this EVI. The type is bool.
    AdvertiseBviMac interface{}

    // Route Aliasing is disabled. The type is bool.
    AliasingDisabled interface{}

    // Unknown-unicast flooding is disabled. The type is bool.
    UnknownUnicastFloodingDisabled interface{}

    // Route Re-origination is disabled. The type is bool.
    ReoriginateDisabled interface{}

    // EVPN Instance is Regular/Stitching side. The type is bool.
    Stitching interface{}

    // EVI is connected to multicast source. The type is bool.
    MulticastSourceConnected interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Standby_EviDetail_Elements_Element_EvpnInstance

    // Flow Label Information.
    FlowLabel Evpn_Standby_EviDetail_Elements_Element_FlowLabel

    // Automatic Route Distingtuisher.
    RdAuto Evpn_Standby_EviDetail_Elements_Element_RdAuto

    // Configured Route Distinguisher.
    RdConfigured Evpn_Standby_EviDetail_Elements_Element_RdConfigured

    // Automatic Route Target.
    RtAuto Evpn_Standby_EviDetail_Elements_Element_RtAuto
}

func (element *Evpn_Standby_EviDetail_Elements_Element) GetEntityData() *types.CommonEntityData {
    element.EntityData.YFilter = element.YFilter
    element.EntityData.YangName = "element"
    element.EntityData.BundleName = "cisco_ios_xr"
    element.EntityData.ParentYangName = "elements"
    element.EntityData.SegmentPath = "element"
    element.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    element.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    element.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    element.EntityData.Children = types.NewOrderedMap()
    element.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &element.EvpnInstance})
    element.EntityData.Children.Append("flow-label", types.YChild{"FlowLabel", &element.FlowLabel})
    element.EntityData.Children.Append("rd-auto", types.YChild{"RdAuto", &element.RdAuto})
    element.EntityData.Children.Append("rd-configured", types.YChild{"RdConfigured", &element.RdConfigured})
    element.EntityData.Children.Append("rt-auto", types.YChild{"RtAuto", &element.RtAuto})
    element.EntityData.Leafs = types.NewOrderedMap()
    element.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", element.Evi})
    element.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", element.Encapsulation})
    element.EntityData.Leafs.Append("evi-xr", types.YLeaf{"EviXr", element.EviXr})
    element.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", element.EncapsulationXr})
    element.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", element.BdName})
    element.EntityData.Leafs.Append("type", types.YLeaf{"Type", element.Type})
    element.EntityData.Leafs.Append("description", types.YLeaf{"Description", element.Description})
    element.EntityData.Leafs.Append("unicast-label", types.YLeaf{"UnicastLabel", element.UnicastLabel})
    element.EntityData.Leafs.Append("multicast-label", types.YLeaf{"MulticastLabel", element.MulticastLabel})
    element.EntityData.Leafs.Append("cw-disable", types.YLeaf{"CwDisable", element.CwDisable})
    element.EntityData.Leafs.Append("table-policy-name", types.YLeaf{"TablePolicyName", element.TablePolicyName})
    element.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", element.ForwardClass})
    element.EntityData.Leafs.Append("rt-import-block-set", types.YLeaf{"RtImportBlockSet", element.RtImportBlockSet})
    element.EntityData.Leafs.Append("rt-export-block-set", types.YLeaf{"RtExportBlockSet", element.RtExportBlockSet})
    element.EntityData.Leafs.Append("advertise-mac", types.YLeaf{"AdvertiseMac", element.AdvertiseMac})
    element.EntityData.Leafs.Append("advertise-bvi-mac", types.YLeaf{"AdvertiseBviMac", element.AdvertiseBviMac})
    element.EntityData.Leafs.Append("aliasing-disabled", types.YLeaf{"AliasingDisabled", element.AliasingDisabled})
    element.EntityData.Leafs.Append("unknown-unicast-flooding-disabled", types.YLeaf{"UnknownUnicastFloodingDisabled", element.UnknownUnicastFloodingDisabled})
    element.EntityData.Leafs.Append("reoriginate-disabled", types.YLeaf{"ReoriginateDisabled", element.ReoriginateDisabled})
    element.EntityData.Leafs.Append("stitching", types.YLeaf{"Stitching", element.Stitching})
    element.EntityData.Leafs.Append("multicast-source-connected", types.YLeaf{"MulticastSourceConnected", element.MulticastSourceConnected})

    element.EntityData.YListKeys = []string {}

    return &(element.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_EvpnInstance
// EVPN Instance summary information
type Evpn_Standby_EviDetail_Elements_Element_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Standby_EviDetail_Elements_Element_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "element"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_FlowLabel
// Flow Label Information
type Evpn_Standby_EviDetail_Elements_Element_FlowLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static flow label. The type is bool.
    StaticFlowLabel interface{}

    // Globally configured flow label. The type is bool.
    GlobalFlowLabel interface{}
}

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetEntityData() *types.CommonEntityData {
    flowLabel.EntityData.YFilter = flowLabel.YFilter
    flowLabel.EntityData.YangName = "flow-label"
    flowLabel.EntityData.BundleName = "cisco_ios_xr"
    flowLabel.EntityData.ParentYangName = "element"
    flowLabel.EntityData.SegmentPath = "flow-label"
    flowLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowLabel.EntityData.Children = types.NewOrderedMap()
    flowLabel.EntityData.Leafs = types.NewOrderedMap()
    flowLabel.EntityData.Leafs.Append("static-flow-label", types.YLeaf{"StaticFlowLabel", flowLabel.StaticFlowLabel})
    flowLabel.EntityData.Leafs.Append("global-flow-label", types.YLeaf{"GlobalFlowLabel", flowLabel.GlobalFlowLabel})

    flowLabel.EntityData.YListKeys = []string {}

    return &(flowLabel.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdAuto
// Automatic Route Distingtuisher
type Evpn_Standby_EviDetail_Elements_Element_RdAuto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RD. The type is L2vpnAdRd.
    Rd interface{}

    // auto.
    Auto Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto

    // two byte as.
    TwoByteAs Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs

    // v4 addr.
    V4Addr Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr
}

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetEntityData() *types.CommonEntityData {
    rdAuto.EntityData.YFilter = rdAuto.YFilter
    rdAuto.EntityData.YangName = "rd-auto"
    rdAuto.EntityData.BundleName = "cisco_ios_xr"
    rdAuto.EntityData.ParentYangName = "element"
    rdAuto.EntityData.SegmentPath = "rd-auto"
    rdAuto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdAuto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdAuto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdAuto.EntityData.Children = types.NewOrderedMap()
    rdAuto.EntityData.Children.Append("auto", types.YChild{"Auto", &rdAuto.Auto})
    rdAuto.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &rdAuto.TwoByteAs})
    rdAuto.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &rdAuto.FourByteAs})
    rdAuto.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &rdAuto.V4Addr})
    rdAuto.EntityData.Leafs = types.NewOrderedMap()
    rdAuto.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", rdAuto.Rd})

    rdAuto.EntityData.YListKeys = []string {}

    return &(rdAuto.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto
// auto
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetEntityData() *types.CommonEntityData {
    auto.EntityData.YFilter = auto.YFilter
    auto.EntityData.YangName = "auto"
    auto.EntityData.BundleName = "cisco_ios_xr"
    auto.EntityData.ParentYangName = "rd-auto"
    auto.EntityData.SegmentPath = "auto"
    auto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auto.EntityData.Children = types.NewOrderedMap()
    auto.EntityData.Leafs = types.NewOrderedMap()
    auto.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", auto.RouterId})
    auto.EntityData.Leafs.Append("auto-index", types.YLeaf{"AutoIndex", auto.AutoIndex})

    auto.EntityData.YListKeys = []string {}

    return &(auto.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rd-auto"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rd-auto"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rd-auto"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured
// Configured Route Distinguisher
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RD. The type is L2vpnAdRd.
    Rd interface{}

    // auto.
    Auto Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto

    // two byte as.
    TwoByteAs Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs

    // v4 addr.
    V4Addr Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr
}

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetEntityData() *types.CommonEntityData {
    rdConfigured.EntityData.YFilter = rdConfigured.YFilter
    rdConfigured.EntityData.YangName = "rd-configured"
    rdConfigured.EntityData.BundleName = "cisco_ios_xr"
    rdConfigured.EntityData.ParentYangName = "element"
    rdConfigured.EntityData.SegmentPath = "rd-configured"
    rdConfigured.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rdConfigured.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rdConfigured.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rdConfigured.EntityData.Children = types.NewOrderedMap()
    rdConfigured.EntityData.Children.Append("auto", types.YChild{"Auto", &rdConfigured.Auto})
    rdConfigured.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &rdConfigured.TwoByteAs})
    rdConfigured.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &rdConfigured.FourByteAs})
    rdConfigured.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &rdConfigured.V4Addr})
    rdConfigured.EntityData.Leafs = types.NewOrderedMap()
    rdConfigured.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", rdConfigured.Rd})

    rdConfigured.EntityData.YListKeys = []string {}

    return &(rdConfigured.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto
// auto
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetEntityData() *types.CommonEntityData {
    auto.EntityData.YFilter = auto.YFilter
    auto.EntityData.YangName = "auto"
    auto.EntityData.BundleName = "cisco_ios_xr"
    auto.EntityData.ParentYangName = "rd-configured"
    auto.EntityData.SegmentPath = "auto"
    auto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auto.EntityData.Children = types.NewOrderedMap()
    auto.EntityData.Leafs = types.NewOrderedMap()
    auto.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", auto.RouterId})
    auto.EntityData.Leafs.Append("auto-index", types.YLeaf{"AutoIndex", auto.AutoIndex})

    auto.EntityData.YListKeys = []string {}

    return &(auto.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rd-configured"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rd-configured"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rd-configured"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAuto
// Automatic Route Target
type Evpn_Standby_EviDetail_Elements_Element_RtAuto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs

    // v4 addr.
    V4Addr Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr

    // es import.
    EsImport Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport
}

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetEntityData() *types.CommonEntityData {
    rtAuto.EntityData.YFilter = rtAuto.YFilter
    rtAuto.EntityData.YangName = "rt-auto"
    rtAuto.EntityData.BundleName = "cisco_ios_xr"
    rtAuto.EntityData.ParentYangName = "element"
    rtAuto.EntityData.SegmentPath = "rt-auto"
    rtAuto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rtAuto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rtAuto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rtAuto.EntityData.Children = types.NewOrderedMap()
    rtAuto.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &rtAuto.TwoByteAs})
    rtAuto.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &rtAuto.FourByteAs})
    rtAuto.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &rtAuto.V4Addr})
    rtAuto.EntityData.Children.Append("es-import", types.YChild{"EsImport", &rtAuto.EsImport})
    rtAuto.EntityData.Leafs = types.NewOrderedMap()
    rtAuto.EntityData.Leafs.Append("rt", types.YLeaf{"Rt", rtAuto.Rt})

    rtAuto.EntityData.YListKeys = []string {}

    return &(rtAuto.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rt-auto"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rt-auto"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rt-auto"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport
// es import
type Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "rt-auto"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = types.NewOrderedMap()
    esImport.EntityData.Leafs = types.NewOrderedMap()
    esImport.EntityData.Leafs.Append("high-bytes", types.YLeaf{"HighBytes", esImport.HighBytes})
    esImport.EntityData.Leafs.Append("low-bytes", types.YLeaf{"LowBytes", esImport.LowBytes})

    esImport.EntityData.YListKeys = []string {}

    return &(esImport.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren
// Container for all EVI detail info
type Evpn_Standby_EviDetail_EviChildren struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Neighbor table.
    Neighbors Evpn_Standby_EviDetail_EviChildren_Neighbors

    // EVPN Ethernet Auto-Discovery table.
    EthernetAutoDiscoveries Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries

    // L2VPN EVPN IMCAST table.
    InclusiveMulticasts Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts

    // L2VPN EVPN EVI RT Child Table.
    RouteTargets Evpn_Standby_EviDetail_EviChildren_RouteTargets

    // L2VPN EVPN EVI MAC table.
    Macs Evpn_Standby_EviDetail_EviChildren_Macs
}

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetEntityData() *types.CommonEntityData {
    eviChildren.EntityData.YFilter = eviChildren.YFilter
    eviChildren.EntityData.YangName = "evi-children"
    eviChildren.EntityData.BundleName = "cisco_ios_xr"
    eviChildren.EntityData.ParentYangName = "evi-detail"
    eviChildren.EntityData.SegmentPath = "evi-children"
    eviChildren.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviChildren.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviChildren.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviChildren.EntityData.Children = types.NewOrderedMap()
    eviChildren.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &eviChildren.Neighbors})
    eviChildren.EntityData.Children.Append("ethernet-auto-discoveries", types.YChild{"EthernetAutoDiscoveries", &eviChildren.EthernetAutoDiscoveries})
    eviChildren.EntityData.Children.Append("inclusive-multicasts", types.YChild{"InclusiveMulticasts", &eviChildren.InclusiveMulticasts})
    eviChildren.EntityData.Children.Append("route-targets", types.YChild{"RouteTargets", &eviChildren.RouteTargets})
    eviChildren.EntityData.Children.Append("macs", types.YChild{"Macs", &eviChildren.Macs})
    eviChildren.EntityData.Leafs = types.NewOrderedMap()

    eviChildren.EntityData.YListKeys = []string {}

    return &(eviChildren.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Neighbors
// EVPN Neighbor table
type Evpn_Standby_EviDetail_EviChildren_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Neighbor table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor.
    Neighbor []*Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor
}

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "evi-children"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor
// EVPN Neighbor table
type Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Neighbor IP. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborIp interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance
}

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &neighbor.EvpnInstance})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", neighbor.Evi})
    neighbor.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", neighbor.Encapsulation})
    neighbor.EntityData.Leafs.Append("neighbor-ip", types.YLeaf{"NeighborIp", neighbor.NeighborIp})
    neighbor.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", neighbor.Neighbor})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance
// EVPN Instance summary information
type Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "neighbor"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries
// EVPN Ethernet Auto-Discovery table
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet Auto-Discovery Entry. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery.
    EthernetAutoDiscovery []*Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
}

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetEntityData() *types.CommonEntityData {
    ethernetAutoDiscoveries.EntityData.YFilter = ethernetAutoDiscoveries.YFilter
    ethernetAutoDiscoveries.EntityData.YangName = "ethernet-auto-discoveries"
    ethernetAutoDiscoveries.EntityData.BundleName = "cisco_ios_xr"
    ethernetAutoDiscoveries.EntityData.ParentYangName = "evi-children"
    ethernetAutoDiscoveries.EntityData.SegmentPath = "ethernet-auto-discoveries"
    ethernetAutoDiscoveries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetAutoDiscoveries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetAutoDiscoveries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetAutoDiscoveries.EntityData.Children = types.NewOrderedMap()
    ethernetAutoDiscoveries.EntityData.Children.Append("ethernet-auto-discovery", types.YChild{"EthernetAutoDiscovery", nil})
    for i := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
        ethernetAutoDiscoveries.EntityData.Children.Append(types.GetSegmentPath(ethernetAutoDiscoveries.EthernetAutoDiscovery[i]), types.YChild{"EthernetAutoDiscovery", ethernetAutoDiscoveries.EthernetAutoDiscovery[i]})
    }
    ethernetAutoDiscoveries.EntityData.Leafs = types.NewOrderedMap()

    ethernetAutoDiscoveries.EntityData.YListKeys = []string {}

    return &(ethernetAutoDiscoveries.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
// EVPN Ethernet Auto-Discovery Entry
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Local nexthop IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalNextHop interface{}

    // Associated local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Indication of EthernetAutoDiscovery Route is local. The type is bool.
    IsLocalEad interface{}

    // Single-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleActive interface{}

    // Single-flow-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleFlowActive interface{}

    // Number of items in path list buffer. The type is interface{} with range:
    // 0..4294967295.
    NumPaths interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance

    // Ethernet Segment id. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer.
    PathBuffer []*Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
}

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetEntityData() *types.CommonEntityData {
    ethernetAutoDiscovery.EntityData.YFilter = ethernetAutoDiscovery.YFilter
    ethernetAutoDiscovery.EntityData.YangName = "ethernet-auto-discovery"
    ethernetAutoDiscovery.EntityData.BundleName = "cisco_ios_xr"
    ethernetAutoDiscovery.EntityData.ParentYangName = "ethernet-auto-discoveries"
    ethernetAutoDiscovery.EntityData.SegmentPath = "ethernet-auto-discovery"
    ethernetAutoDiscovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetAutoDiscovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetAutoDiscovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetAutoDiscovery.EntityData.Children = types.NewOrderedMap()
    ethernetAutoDiscovery.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &ethernetAutoDiscovery.EvpnInstance})
    ethernetAutoDiscovery.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
        ethernetAutoDiscovery.EntityData.Children.Append(types.GetSegmentPath(ethernetAutoDiscovery.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", ethernetAutoDiscovery.EthernetSegmentIdentifier[i]})
    }
    ethernetAutoDiscovery.EntityData.Children.Append("path-buffer", types.YChild{"PathBuffer", nil})
    for i := range ethernetAutoDiscovery.PathBuffer {
        ethernetAutoDiscovery.EntityData.Children.Append(types.GetSegmentPath(ethernetAutoDiscovery.PathBuffer[i]), types.YChild{"PathBuffer", ethernetAutoDiscovery.PathBuffer[i]})
    }
    ethernetAutoDiscovery.EntityData.Leafs = types.NewOrderedMap()
    ethernetAutoDiscovery.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", ethernetAutoDiscovery.Evi})
    ethernetAutoDiscovery.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", ethernetAutoDiscovery.Encapsulation})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", ethernetAutoDiscovery.Esi1})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", ethernetAutoDiscovery.Esi2})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", ethernetAutoDiscovery.Esi3})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", ethernetAutoDiscovery.Esi4})
    ethernetAutoDiscovery.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", ethernetAutoDiscovery.Esi5})
    ethernetAutoDiscovery.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", ethernetAutoDiscovery.EthernetTag})
    ethernetAutoDiscovery.EntityData.Leafs.Append("ethernet-tag-xr", types.YLeaf{"EthernetTagXr", ethernetAutoDiscovery.EthernetTagXr})
    ethernetAutoDiscovery.EntityData.Leafs.Append("local-next-hop", types.YLeaf{"LocalNextHop", ethernetAutoDiscovery.LocalNextHop})
    ethernetAutoDiscovery.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", ethernetAutoDiscovery.LocalLabel})
    ethernetAutoDiscovery.EntityData.Leafs.Append("is-local-ead", types.YLeaf{"IsLocalEad", ethernetAutoDiscovery.IsLocalEad})
    ethernetAutoDiscovery.EntityData.Leafs.Append("redundancy-single-active", types.YLeaf{"RedundancySingleActive", ethernetAutoDiscovery.RedundancySingleActive})
    ethernetAutoDiscovery.EntityData.Leafs.Append("redundancy-single-flow-active", types.YLeaf{"RedundancySingleFlowActive", ethernetAutoDiscovery.RedundancySingleFlowActive})
    ethernetAutoDiscovery.EntityData.Leafs.Append("num-paths", types.YLeaf{"NumPaths", ethernetAutoDiscovery.NumPaths})

    ethernetAutoDiscovery.EntityData.YListKeys = []string {}

    return &(ethernetAutoDiscovery.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance
// EVPN Instance summary information
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "ethernet-auto-discovery"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "ethernet-auto-discovery"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
// Path List Buffer
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetEntityData() *types.CommonEntityData {
    pathBuffer.EntityData.YFilter = pathBuffer.YFilter
    pathBuffer.EntityData.YangName = "path-buffer"
    pathBuffer.EntityData.BundleName = "cisco_ios_xr"
    pathBuffer.EntityData.ParentYangName = "ethernet-auto-discovery"
    pathBuffer.EntityData.SegmentPath = "path-buffer"
    pathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathBuffer.EntityData.Children = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", pathBuffer.TunnelEndpointId})
    pathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", pathBuffer.NextHop})
    pathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", pathBuffer.OutputLabel})
    pathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel})

    pathBuffer.EntityData.YListKeys = []string {}

    return &(pathBuffer.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts
// L2VPN EVPN IMCAST table
type Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN IMCAST table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast.
    InclusiveMulticast []*Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
}

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetEntityData() *types.CommonEntityData {
    inclusiveMulticasts.EntityData.YFilter = inclusiveMulticasts.YFilter
    inclusiveMulticasts.EntityData.YangName = "inclusive-multicasts"
    inclusiveMulticasts.EntityData.BundleName = "cisco_ios_xr"
    inclusiveMulticasts.EntityData.ParentYangName = "evi-children"
    inclusiveMulticasts.EntityData.SegmentPath = "inclusive-multicasts"
    inclusiveMulticasts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inclusiveMulticasts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inclusiveMulticasts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inclusiveMulticasts.EntityData.Children = types.NewOrderedMap()
    inclusiveMulticasts.EntityData.Children.Append("inclusive-multicast", types.YChild{"InclusiveMulticast", nil})
    for i := range inclusiveMulticasts.InclusiveMulticast {
        inclusiveMulticasts.EntityData.Children.Append(types.GetSegmentPath(inclusiveMulticasts.InclusiveMulticast[i]), types.YChild{"InclusiveMulticast", inclusiveMulticasts.InclusiveMulticast[i]})
    }
    inclusiveMulticasts.EntityData.Leafs = types.NewOrderedMap()

    inclusiveMulticasts.EntityData.YListKeys = []string {}

    return &(inclusiveMulticasts.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
// L2VPN EVPN IMCAST table
type Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // Originating IP. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIp interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Originating IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIpXr interface{}

    // Tunnel Endpoint ID. The type is interface{} with range: 0..4294967295.
    TunnelEndpointId interface{}

    // IP of nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Local entry. The type is bool.
    IsLocalEntry interface{}

    // Proxy entry. The type is bool.
    IsProxyEntry interface{}

    // SR-TE Policy. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SrtePolicy interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance
}

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetEntityData() *types.CommonEntityData {
    inclusiveMulticast.EntityData.YFilter = inclusiveMulticast.YFilter
    inclusiveMulticast.EntityData.YangName = "inclusive-multicast"
    inclusiveMulticast.EntityData.BundleName = "cisco_ios_xr"
    inclusiveMulticast.EntityData.ParentYangName = "inclusive-multicasts"
    inclusiveMulticast.EntityData.SegmentPath = "inclusive-multicast"
    inclusiveMulticast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inclusiveMulticast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inclusiveMulticast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inclusiveMulticast.EntityData.Children = types.NewOrderedMap()
    inclusiveMulticast.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &inclusiveMulticast.EvpnInstance})
    inclusiveMulticast.EntityData.Leafs = types.NewOrderedMap()
    inclusiveMulticast.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", inclusiveMulticast.Evi})
    inclusiveMulticast.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", inclusiveMulticast.Encapsulation})
    inclusiveMulticast.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", inclusiveMulticast.EthernetTag})
    inclusiveMulticast.EntityData.Leafs.Append("originating-ip", types.YLeaf{"OriginatingIp", inclusiveMulticast.OriginatingIp})
    inclusiveMulticast.EntityData.Leafs.Append("ethernet-tag-xr", types.YLeaf{"EthernetTagXr", inclusiveMulticast.EthernetTagXr})
    inclusiveMulticast.EntityData.Leafs.Append("originating-ip-xr", types.YLeaf{"OriginatingIpXr", inclusiveMulticast.OriginatingIpXr})
    inclusiveMulticast.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", inclusiveMulticast.TunnelEndpointId})
    inclusiveMulticast.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", inclusiveMulticast.NextHop})
    inclusiveMulticast.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", inclusiveMulticast.OutputLabel})
    inclusiveMulticast.EntityData.Leafs.Append("is-local-entry", types.YLeaf{"IsLocalEntry", inclusiveMulticast.IsLocalEntry})
    inclusiveMulticast.EntityData.Leafs.Append("is-proxy-entry", types.YLeaf{"IsProxyEntry", inclusiveMulticast.IsProxyEntry})
    inclusiveMulticast.EntityData.Leafs.Append("srte-policy", types.YLeaf{"SrtePolicy", inclusiveMulticast.SrtePolicy})

    inclusiveMulticast.EntityData.YListKeys = []string {}

    return &(inclusiveMulticast.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance
// EVPN Instance summary information
type Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "inclusive-multicast"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets
// L2VPN EVPN EVI RT Child Table
type Evpn_Standby_EviDetail_EviChildren_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI RT Table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget.
    RouteTarget []*Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget
}

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetEntityData() *types.CommonEntityData {
    routeTargets.EntityData.YFilter = routeTargets.YFilter
    routeTargets.EntityData.YangName = "route-targets"
    routeTargets.EntityData.BundleName = "cisco_ios_xr"
    routeTargets.EntityData.ParentYangName = "evi-children"
    routeTargets.EntityData.SegmentPath = "route-targets"
    routeTargets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTargets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTargets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTargets.EntityData.Children = types.NewOrderedMap()
    routeTargets.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", nil})
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children.Append(types.GetSegmentPath(routeTargets.RouteTarget[i]), types.YChild{"RouteTarget", routeTargets.RouteTarget[i]})
    }
    routeTargets.EntityData.Leafs = types.NewOrderedMap()

    routeTargets.EntityData.YListKeys = []string {}

    return &(routeTargets.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget
// L2VPN EVPN EVI RT Table
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Role of the route target. The type is BgpRouteTargetRole.
    Role interface{}

    // Format of the route target. The type is BgpRouteTargetFormat.
    Format interface{}

    // Two or Four byte AS Number. The type is interface{} with range:
    // 1..4294967295.
    As interface{}

    // RT AS Index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // RT IP Index. The type is interface{} with range: 0..65535.
    AddrIndex interface{}

    // RT IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // RT Role. The type is L2vpnAdRtRole.
    RouteTargetRole interface{}

    // RT Stitching. The type is bool.
    RouteTargetStitching interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance

    // Route Target.
    RouteTarget Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-targets"
    routeTarget.EntityData.SegmentPath = "route-target"
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &routeTarget.EvpnInstance})
    routeTarget.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", &routeTarget.RouteTarget})
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", routeTarget.Evi})
    routeTarget.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", routeTarget.Encapsulation})
    routeTarget.EntityData.Leafs.Append("role", types.YLeaf{"Role", routeTarget.Role})
    routeTarget.EntityData.Leafs.Append("format", types.YLeaf{"Format", routeTarget.Format})
    routeTarget.EntityData.Leafs.Append("as", types.YLeaf{"As", routeTarget.As})
    routeTarget.EntityData.Leafs.Append("as-index", types.YLeaf{"AsIndex", routeTarget.AsIndex})
    routeTarget.EntityData.Leafs.Append("addr-index", types.YLeaf{"AddrIndex", routeTarget.AddrIndex})
    routeTarget.EntityData.Leafs.Append("address", types.YLeaf{"Address", routeTarget.Address})
    routeTarget.EntityData.Leafs.Append("route-target-role", types.YLeaf{"RouteTargetRole", routeTarget.RouteTargetRole})
    routeTarget.EntityData.Leafs.Append("route-target-stitching", types.YLeaf{"RouteTargetStitching", routeTarget.RouteTargetStitching})

    routeTarget.EntityData.YListKeys = []string {}

    return &(routeTarget.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance
// EVPN Instance summary information
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "route-target"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
// Route Target
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs

    // v4 addr.
    V4Addr Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr

    // es import.
    EsImport Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "route-target"
    routeTarget.EntityData.SegmentPath = "route-target"
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("two-byte-as", types.YChild{"TwoByteAs", &routeTarget.TwoByteAs})
    routeTarget.EntityData.Children.Append("four-byte-as", types.YChild{"FourByteAs", &routeTarget.FourByteAs})
    routeTarget.EntityData.Children.Append("v4-addr", types.YChild{"V4Addr", &routeTarget.V4Addr})
    routeTarget.EntityData.Children.Append("es-import", types.YChild{"EsImport", &routeTarget.EsImport})
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("rt", types.YLeaf{"Rt", routeTarget.Rt})

    routeTarget.EntityData.YListKeys = []string {}

    return &(routeTarget.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "route-target"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs = types.NewOrderedMap()
    twoByteAs.EntityData.Leafs.Append("two-byte-as", types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs})
    twoByteAs.EntityData.Leafs.Append("four-byte-index", types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex})

    twoByteAs.EntityData.YListKeys = []string {}

    return &(twoByteAs.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "route-target"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs = types.NewOrderedMap()
    fourByteAs.EntityData.Leafs.Append("four-byte-as", types.YLeaf{"FourByteAs", fourByteAs.FourByteAs})
    fourByteAs.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex})

    fourByteAs.EntityData.YListKeys = []string {}

    return &(fourByteAs.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "route-target"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = types.NewOrderedMap()
    v4Addr.EntityData.Leafs = types.NewOrderedMap()
    v4Addr.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address})
    v4Addr.EntityData.Leafs.Append("two-byte-index", types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex})

    v4Addr.EntityData.YListKeys = []string {}

    return &(v4Addr.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport
// es import
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "route-target"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = types.NewOrderedMap()
    esImport.EntityData.Leafs = types.NewOrderedMap()
    esImport.EntityData.Leafs.Append("high-bytes", types.YLeaf{"HighBytes", esImport.HighBytes})
    esImport.EntityData.Leafs.Append("low-bytes", types.YLeaf{"LowBytes", esImport.LowBytes})

    esImport.EntityData.YListKeys = []string {}

    return &(esImport.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Macs
// L2VPN EVPN EVI MAC table
type Evpn_Standby_EviDetail_EviChildren_Macs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN MAC table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Macs_Mac.
    Mac []*Evpn_Standby_EviDetail_EviChildren_Macs_Mac
}

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetEntityData() *types.CommonEntityData {
    macs.EntityData.YFilter = macs.YFilter
    macs.EntityData.YangName = "macs"
    macs.EntityData.BundleName = "cisco_ios_xr"
    macs.EntityData.ParentYangName = "evi-children"
    macs.EntityData.SegmentPath = "macs"
    macs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macs.EntityData.Children = types.NewOrderedMap()
    macs.EntityData.Children.Append("mac", types.YChild{"Mac", nil})
    for i := range macs.Mac {
        macs.EntityData.Children.Append(types.GetSegmentPath(macs.Mac[i]), types.YChild{"Mac", macs.Mac[i]})
    }
    macs.EntityData.Leafs = types.NewOrderedMap()

    macs.EntityData.YListKeys = []string {}

    return &(macs.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac
// L2VPN EVPN MAC table
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // IP Address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddressXr interface{}

    // IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddressXr interface{}

    // Associated local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Number of items in path list buffer. The type is interface{} with range:
    // 0..4294967295.
    NumPaths interface{}

    // Indication of MAC being locally generated. The type is bool.
    IsLocalMac interface{}

    // Proxy entry. The type is bool.
    IsProxyEntry interface{}

    // Indication of MAC being remotely generated. The type is bool.
    IsRemoteMac interface{}

    // SOO nexthop (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SooNexthop interface{}

    // IP nexthop address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpnhAddress interface{}

    // ESI port key. The type is interface{} with range: 0..65535.
    EsiPortKey interface{}

    // Encap type of local MAC. The type is interface{} with range: 0..255.
    LocalEncapType interface{}

    // Encap type of remote MAC. The type is interface{} with range: 0..255.
    RemoteEncapType interface{}

    // Port the MAC was learned on. The type is string.
    LearnedBridgePortName interface{}

    // local seq id. The type is interface{} with range: 0..4294967295.
    LocalSeqId interface{}

    // remote seq id. The type is interface{} with range: 0..4294967295.
    RemoteSeqId interface{}

    // local l3 label. The type is interface{} with range: 0..4294967295.
    LocalL3Label interface{}

    // Router MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    RouterMacAddress interface{}

    // Number of flushes requested . The type is interface{} with range: 0..65535.
    MacFlushRequested interface{}

    // Number of flushes received . The type is interface{} with range: 0..65535.
    MacFlushReceived interface{}

    // MPLS Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Internal Label has resolved per-ES EAD and per-EVI EAD or MAC routes. The
    // type is bool.
    Resolved interface{}

    // Indication if Local MAC is statically configured. The type is bool.
    LocalIsStatic interface{}

    // Indication if Remote MAC is statically configured. The type is bool.
    RemoteIsStatic interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Standby_EviDetail_EviChildren_Macs_Mac_EvpnInstance

    // Local Ethernet Segment id. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier.
    LocalEthernetSegmentIdentifier []*Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier

    // Remote Ethernet Segment id. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier.
    RemoteEthernetSegmentIdentifier []*Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer.
    PathBuffer []*Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer
}

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "macs"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &mac.EvpnInstance})
    mac.EntityData.Children.Append("local-ethernet-segment-identifier", types.YChild{"LocalEthernetSegmentIdentifier", nil})
    for i := range mac.LocalEthernetSegmentIdentifier {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.LocalEthernetSegmentIdentifier[i]), types.YChild{"LocalEthernetSegmentIdentifier", mac.LocalEthernetSegmentIdentifier[i]})
    }
    mac.EntityData.Children.Append("remote-ethernet-segment-identifier", types.YChild{"RemoteEthernetSegmentIdentifier", nil})
    for i := range mac.RemoteEthernetSegmentIdentifier {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.RemoteEthernetSegmentIdentifier[i]), types.YChild{"RemoteEthernetSegmentIdentifier", mac.RemoteEthernetSegmentIdentifier[i]})
    }
    mac.EntityData.Children.Append("path-buffer", types.YChild{"PathBuffer", nil})
    for i := range mac.PathBuffer {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.PathBuffer[i]), types.YChild{"PathBuffer", mac.PathBuffer[i]})
    }
    mac.EntityData.Leafs = types.NewOrderedMap()
    mac.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", mac.Evi})
    mac.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", mac.Encapsulation})
    mac.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", mac.EthernetTag})
    mac.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", mac.MacAddress})
    mac.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", mac.IpAddress})
    mac.EntityData.Leafs.Append("ethernet-tag-xr", types.YLeaf{"EthernetTagXr", mac.EthernetTagXr})
    mac.EntityData.Leafs.Append("mac-address-xr", types.YLeaf{"MacAddressXr", mac.MacAddressXr})
    mac.EntityData.Leafs.Append("ip-address-xr", types.YLeaf{"IpAddressXr", mac.IpAddressXr})
    mac.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", mac.LocalLabel})
    mac.EntityData.Leafs.Append("num-paths", types.YLeaf{"NumPaths", mac.NumPaths})
    mac.EntityData.Leafs.Append("is-local-mac", types.YLeaf{"IsLocalMac", mac.IsLocalMac})
    mac.EntityData.Leafs.Append("is-proxy-entry", types.YLeaf{"IsProxyEntry", mac.IsProxyEntry})
    mac.EntityData.Leafs.Append("is-remote-mac", types.YLeaf{"IsRemoteMac", mac.IsRemoteMac})
    mac.EntityData.Leafs.Append("soo-nexthop", types.YLeaf{"SooNexthop", mac.SooNexthop})
    mac.EntityData.Leafs.Append("ipnh-address", types.YLeaf{"IpnhAddress", mac.IpnhAddress})
    mac.EntityData.Leafs.Append("esi-port-key", types.YLeaf{"EsiPortKey", mac.EsiPortKey})
    mac.EntityData.Leafs.Append("local-encap-type", types.YLeaf{"LocalEncapType", mac.LocalEncapType})
    mac.EntityData.Leafs.Append("remote-encap-type", types.YLeaf{"RemoteEncapType", mac.RemoteEncapType})
    mac.EntityData.Leafs.Append("learned-bridge-port-name", types.YLeaf{"LearnedBridgePortName", mac.LearnedBridgePortName})
    mac.EntityData.Leafs.Append("local-seq-id", types.YLeaf{"LocalSeqId", mac.LocalSeqId})
    mac.EntityData.Leafs.Append("remote-seq-id", types.YLeaf{"RemoteSeqId", mac.RemoteSeqId})
    mac.EntityData.Leafs.Append("local-l3-label", types.YLeaf{"LocalL3Label", mac.LocalL3Label})
    mac.EntityData.Leafs.Append("router-mac-address", types.YLeaf{"RouterMacAddress", mac.RouterMacAddress})
    mac.EntityData.Leafs.Append("mac-flush-requested", types.YLeaf{"MacFlushRequested", mac.MacFlushRequested})
    mac.EntityData.Leafs.Append("mac-flush-received", types.YLeaf{"MacFlushReceived", mac.MacFlushReceived})
    mac.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", mac.InternalLabel})
    mac.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", mac.Resolved})
    mac.EntityData.Leafs.Append("local-is-static", types.YLeaf{"LocalIsStatic", mac.LocalIsStatic})
    mac.EntityData.Leafs.Append("remote-is-static", types.YLeaf{"RemoteIsStatic", mac.RemoteIsStatic})

    mac.EntityData.YListKeys = []string {}

    return &(mac.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac_EvpnInstance
// EVPN Instance summary information
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "mac"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier
// Local Ethernet Segment id
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    localEthernetSegmentIdentifier.EntityData.YFilter = localEthernetSegmentIdentifier.YFilter
    localEthernetSegmentIdentifier.EntityData.YangName = "local-ethernet-segment-identifier"
    localEthernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    localEthernetSegmentIdentifier.EntityData.ParentYangName = "mac"
    localEthernetSegmentIdentifier.EntityData.SegmentPath = "local-ethernet-segment-identifier"
    localEthernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localEthernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localEthernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localEthernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    localEthernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    localEthernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", localEthernetSegmentIdentifier.Entry})

    localEthernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(localEthernetSegmentIdentifier.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier
// Remote Ethernet Segment id
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    remoteEthernetSegmentIdentifier.EntityData.YFilter = remoteEthernetSegmentIdentifier.YFilter
    remoteEthernetSegmentIdentifier.EntityData.YangName = "remote-ethernet-segment-identifier"
    remoteEthernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    remoteEthernetSegmentIdentifier.EntityData.ParentYangName = "mac"
    remoteEthernetSegmentIdentifier.EntityData.SegmentPath = "remote-ethernet-segment-identifier"
    remoteEthernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteEthernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteEthernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteEthernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    remoteEthernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    remoteEthernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", remoteEthernetSegmentIdentifier.Entry})

    remoteEthernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(remoteEthernetSegmentIdentifier.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer
// Path List Buffer
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetEntityData() *types.CommonEntityData {
    pathBuffer.EntityData.YFilter = pathBuffer.YFilter
    pathBuffer.EntityData.YangName = "path-buffer"
    pathBuffer.EntityData.BundleName = "cisco_ios_xr"
    pathBuffer.EntityData.ParentYangName = "mac"
    pathBuffer.EntityData.SegmentPath = "path-buffer"
    pathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathBuffer.EntityData.Children = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs = types.NewOrderedMap()
    pathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", pathBuffer.TunnelEndpointId})
    pathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", pathBuffer.NextHop})
    pathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", pathBuffer.OutputLabel})
    pathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel})

    pathBuffer.EntityData.YListKeys = []string {}

    return &(pathBuffer.EntityData)
}

// Evpn_Standby_Teps
// L2VPN EVPN TEP Table
type Evpn_Standby_Teps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN TEP Entry. The type is slice of Evpn_Standby_Teps_Tep.
    Tep []*Evpn_Standby_Teps_Tep
}

func (teps *Evpn_Standby_Teps) GetEntityData() *types.CommonEntityData {
    teps.EntityData.YFilter = teps.YFilter
    teps.EntityData.YangName = "teps"
    teps.EntityData.BundleName = "cisco_ios_xr"
    teps.EntityData.ParentYangName = "standby"
    teps.EntityData.SegmentPath = "teps"
    teps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teps.EntityData.Children = types.NewOrderedMap()
    teps.EntityData.Children.Append("tep", types.YChild{"Tep", nil})
    for i := range teps.Tep {
        teps.EntityData.Children.Append(types.GetSegmentPath(teps.Tep[i]), types.YChild{"Tep", teps.Tep[i]})
    }
    teps.EntityData.Leafs = types.NewOrderedMap()

    teps.EntityData.YListKeys = []string {}

    return &(teps.EntityData)
}

// Evpn_Standby_Teps_Tep
// L2VPN EVPN TEP Entry
type Evpn_Standby_Teps_Tep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. TEP id. The type is interface{} with range:
    // 0..4294967295.
    TepId interface{}

    // Tunnel Endpoint id. The type is interface{} with range: 0..4294967295.
    TunnelEndpointId interface{}

    // EVPN Tunnel Endpoint Type. The type is interface{} with range: 0..255.
    Type interface{}

    // in-use counter. The type is interface{} with range: 0..4294967295.
    UseCount interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Table Id in RIB. The type is interface{} with range: 0..4294967295.
    VrfTableId interface{}

    // UDP port. The type is interface{} with range: 0..65535.
    UdpPort interface{}

    // Local TEP Information.
    LocalInfo Evpn_Standby_Teps_Tep_LocalInfo

    // Remote TEP Information.
    RemoteInfo Evpn_Standby_Teps_Tep_RemoteInfo
}

func (tep *Evpn_Standby_Teps_Tep) GetEntityData() *types.CommonEntityData {
    tep.EntityData.YFilter = tep.YFilter
    tep.EntityData.YangName = "tep"
    tep.EntityData.BundleName = "cisco_ios_xr"
    tep.EntityData.ParentYangName = "teps"
    tep.EntityData.SegmentPath = "tep" + types.AddKeyToken(tep.TepId, "tep-id")
    tep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tep.EntityData.Children = types.NewOrderedMap()
    tep.EntityData.Children.Append("local-info", types.YChild{"LocalInfo", &tep.LocalInfo})
    tep.EntityData.Children.Append("remote-info", types.YChild{"RemoteInfo", &tep.RemoteInfo})
    tep.EntityData.Leafs = types.NewOrderedMap()
    tep.EntityData.Leafs.Append("tep-id", types.YLeaf{"TepId", tep.TepId})
    tep.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", tep.TunnelEndpointId})
    tep.EntityData.Leafs.Append("type", types.YLeaf{"Type", tep.Type})
    tep.EntityData.Leafs.Append("use-count", types.YLeaf{"UseCount", tep.UseCount})
    tep.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", tep.VrfName})
    tep.EntityData.Leafs.Append("vrf-table-id", types.YLeaf{"VrfTableId", tep.VrfTableId})
    tep.EntityData.Leafs.Append("udp-port", types.YLeaf{"UdpPort", tep.UdpPort})

    tep.EntityData.YListKeys = []string {"TepId"}

    return &(tep.EntityData)
}

// Evpn_Standby_Teps_Tep_LocalInfo
// Local TEP Information
type Evpn_Standby_Teps_Tep_LocalInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Tunnel encapsulation. The type is interface{} with range: 0..255.
    Encapsulation interface{}

    // IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}
}

func (localInfo *Evpn_Standby_Teps_Tep_LocalInfo) GetEntityData() *types.CommonEntityData {
    localInfo.EntityData.YFilter = localInfo.YFilter
    localInfo.EntityData.YangName = "local-info"
    localInfo.EntityData.BundleName = "cisco_ios_xr"
    localInfo.EntityData.ParentYangName = "tep"
    localInfo.EntityData.SegmentPath = "local-info"
    localInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localInfo.EntityData.Children = types.NewOrderedMap()
    localInfo.EntityData.Leafs = types.NewOrderedMap()
    localInfo.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", localInfo.EthernetVpnId})
    localInfo.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", localInfo.Encapsulation})
    localInfo.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", localInfo.Ip})

    localInfo.EntityData.YListKeys = []string {}

    return &(localInfo.EntityData)
}

// Evpn_Standby_Teps_Tep_RemoteInfo
// Remote TEP Information
type Evpn_Standby_Teps_Tep_RemoteInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Tunnel encapsulation. The type is interface{} with range: 0..255.
    Encapsulation interface{}

    // IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}
}

func (remoteInfo *Evpn_Standby_Teps_Tep_RemoteInfo) GetEntityData() *types.CommonEntityData {
    remoteInfo.EntityData.YFilter = remoteInfo.YFilter
    remoteInfo.EntityData.YangName = "remote-info"
    remoteInfo.EntityData.BundleName = "cisco_ios_xr"
    remoteInfo.EntityData.ParentYangName = "tep"
    remoteInfo.EntityData.SegmentPath = "remote-info"
    remoteInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteInfo.EntityData.Children = types.NewOrderedMap()
    remoteInfo.EntityData.Leafs = types.NewOrderedMap()
    remoteInfo.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", remoteInfo.EthernetVpnId})
    remoteInfo.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", remoteInfo.Encapsulation})
    remoteInfo.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", remoteInfo.Ip})

    remoteInfo.EntityData.YListKeys = []string {}

    return &(remoteInfo.EntityData)
}

// Evpn_Standby_InternalLabels
// EVPN Internal Label Table
type Evpn_Standby_InternalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN Internal Label. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel.
    InternalLabel []*Evpn_Standby_InternalLabels_InternalLabel
}

func (internalLabels *Evpn_Standby_InternalLabels) GetEntityData() *types.CommonEntityData {
    internalLabels.EntityData.YFilter = internalLabels.YFilter
    internalLabels.EntityData.YangName = "internal-labels"
    internalLabels.EntityData.BundleName = "cisco_ios_xr"
    internalLabels.EntityData.ParentYangName = "standby"
    internalLabels.EntityData.SegmentPath = "internal-labels"
    internalLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalLabels.EntityData.Children = types.NewOrderedMap()
    internalLabels.EntityData.Children.Append("internal-label", types.YChild{"InternalLabel", nil})
    for i := range internalLabels.InternalLabel {
        internalLabels.EntityData.Children.Append(types.GetSegmentPath(internalLabels.InternalLabel[i]), types.YChild{"InternalLabel", internalLabels.InternalLabel[i]})
    }
    internalLabels.EntityData.Leafs = types.NewOrderedMap()

    internalLabels.EntityData.YListKeys = []string {}

    return &(internalLabels.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel
// L2VPN EVPN Internal Label
type Evpn_Standby_InternalLabels_InternalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // Encap. The type is interface{} with range: 0..4294967295.
    Encapsulation interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}

    // Ethernet Segment id. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Esi interface{}

    // Label Tag. The type is interface{} with range: 0..4294967295.
    Tag interface{}

    // MPLS Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Number of items in the MAC path buffer. The type is interface{} with range:
    // 0..4294967295.
    MacNumPaths interface{}

    // Number of items in the ead path buffer. The type is interface{} with range:
    // 0..4294967295.
    EadNumPaths interface{}

    // Number of items in the evi path buffer. The type is interface{} with range:
    // 0..4294967295.
    EviNumPaths interface{}

    // Number of items in the sum path buffer. The type is interface{} with range:
    // 0..4294967295.
    SumNumPaths interface{}

    // Number of items in the sum path buffer that are Active Paths. The type is
    // interface{} with range: 0..4294967295.
    SumNumActivePaths interface{}

    // Internal Label has resolved per-ES EAD and per-EVI EAD or MAC routes. The
    // type is bool.
    Resolved interface{}

    // ECMP Disable Per EVI Resolution. The type is bool.
    EcmpDisable interface{}

    // Single-active redundancy configured at remote ES. The type is bool.
    RedundancySingleActive interface{}

    // Single-flow-active redundancy at remote ES (MST-AG). The type is bool.
    RedundancySingleFlowActive interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Standby_InternalLabels_InternalLabel_EvpnInstance

    // MAC Path list buffer. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel_MacPathBuffer.
    MacPathBuffer []*Evpn_Standby_InternalLabels_InternalLabel_MacPathBuffer

    // EAD/ES Path list buffer. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel_EadPathBuffer.
    EadPathBuffer []*Evpn_Standby_InternalLabels_InternalLabel_EadPathBuffer

    // EAD/EVI Path list buffer. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel_EviPathBuffer.
    EviPathBuffer []*Evpn_Standby_InternalLabels_InternalLabel_EviPathBuffer

    // Summary Path list buffer. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel_SummaryPathBuffer.
    SummaryPathBuffer []*Evpn_Standby_InternalLabels_InternalLabel_SummaryPathBuffer
}

func (internalLabel *Evpn_Standby_InternalLabels_InternalLabel) GetEntityData() *types.CommonEntityData {
    internalLabel.EntityData.YFilter = internalLabel.YFilter
    internalLabel.EntityData.YangName = "internal-label"
    internalLabel.EntityData.BundleName = "cisco_ios_xr"
    internalLabel.EntityData.ParentYangName = "internal-labels"
    internalLabel.EntityData.SegmentPath = "internal-label"
    internalLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalLabel.EntityData.Children = types.NewOrderedMap()
    internalLabel.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &internalLabel.EvpnInstance})
    internalLabel.EntityData.Children.Append("mac-path-buffer", types.YChild{"MacPathBuffer", nil})
    for i := range internalLabel.MacPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.MacPathBuffer[i]), types.YChild{"MacPathBuffer", internalLabel.MacPathBuffer[i]})
    }
    internalLabel.EntityData.Children.Append("ead-path-buffer", types.YChild{"EadPathBuffer", nil})
    for i := range internalLabel.EadPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.EadPathBuffer[i]), types.YChild{"EadPathBuffer", internalLabel.EadPathBuffer[i]})
    }
    internalLabel.EntityData.Children.Append("evi-path-buffer", types.YChild{"EviPathBuffer", nil})
    for i := range internalLabel.EviPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.EviPathBuffer[i]), types.YChild{"EviPathBuffer", internalLabel.EviPathBuffer[i]})
    }
    internalLabel.EntityData.Children.Append("summary-path-buffer", types.YChild{"SummaryPathBuffer", nil})
    for i := range internalLabel.SummaryPathBuffer {
        internalLabel.EntityData.Children.Append(types.GetSegmentPath(internalLabel.SummaryPathBuffer[i]), types.YChild{"SummaryPathBuffer", internalLabel.SummaryPathBuffer[i]})
    }
    internalLabel.EntityData.Leafs = types.NewOrderedMap()
    internalLabel.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", internalLabel.Evi})
    internalLabel.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", internalLabel.Encapsulation})
    internalLabel.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", internalLabel.Esi1})
    internalLabel.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", internalLabel.Esi2})
    internalLabel.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", internalLabel.Esi3})
    internalLabel.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", internalLabel.Esi4})
    internalLabel.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", internalLabel.Esi5})
    internalLabel.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", internalLabel.EthernetTag})
    internalLabel.EntityData.Leafs.Append("esi", types.YLeaf{"Esi", internalLabel.Esi})
    internalLabel.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", internalLabel.Tag})
    internalLabel.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", internalLabel.InternalLabel})
    internalLabel.EntityData.Leafs.Append("mac-num-paths", types.YLeaf{"MacNumPaths", internalLabel.MacNumPaths})
    internalLabel.EntityData.Leafs.Append("ead-num-paths", types.YLeaf{"EadNumPaths", internalLabel.EadNumPaths})
    internalLabel.EntityData.Leafs.Append("evi-num-paths", types.YLeaf{"EviNumPaths", internalLabel.EviNumPaths})
    internalLabel.EntityData.Leafs.Append("sum-num-paths", types.YLeaf{"SumNumPaths", internalLabel.SumNumPaths})
    internalLabel.EntityData.Leafs.Append("sum-num-active-paths", types.YLeaf{"SumNumActivePaths", internalLabel.SumNumActivePaths})
    internalLabel.EntityData.Leafs.Append("resolved", types.YLeaf{"Resolved", internalLabel.Resolved})
    internalLabel.EntityData.Leafs.Append("ecmp-disable", types.YLeaf{"EcmpDisable", internalLabel.EcmpDisable})
    internalLabel.EntityData.Leafs.Append("redundancy-single-active", types.YLeaf{"RedundancySingleActive", internalLabel.RedundancySingleActive})
    internalLabel.EntityData.Leafs.Append("redundancy-single-flow-active", types.YLeaf{"RedundancySingleFlowActive", internalLabel.RedundancySingleFlowActive})

    internalLabel.EntityData.YListKeys = []string {}

    return &(internalLabel.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel_EvpnInstance
// EVPN Instance summary information
type Evpn_Standby_InternalLabels_InternalLabel_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Standby_InternalLabels_InternalLabel_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "internal-label"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel_MacPathBuffer
// MAC Path list buffer
type Evpn_Standby_InternalLabels_InternalLabel_MacPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (macPathBuffer *Evpn_Standby_InternalLabels_InternalLabel_MacPathBuffer) GetEntityData() *types.CommonEntityData {
    macPathBuffer.EntityData.YFilter = macPathBuffer.YFilter
    macPathBuffer.EntityData.YangName = "mac-path-buffer"
    macPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    macPathBuffer.EntityData.ParentYangName = "internal-label"
    macPathBuffer.EntityData.SegmentPath = "mac-path-buffer"
    macPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macPathBuffer.EntityData.Children = types.NewOrderedMap()
    macPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    macPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", macPathBuffer.TunnelEndpointId})
    macPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", macPathBuffer.NextHop})
    macPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", macPathBuffer.OutputLabel})
    macPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", macPathBuffer.SrteTunnel})

    macPathBuffer.EntityData.YListKeys = []string {}

    return &(macPathBuffer.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel_EadPathBuffer
// EAD/ES Path list buffer
type Evpn_Standby_InternalLabels_InternalLabel_EadPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (eadPathBuffer *Evpn_Standby_InternalLabels_InternalLabel_EadPathBuffer) GetEntityData() *types.CommonEntityData {
    eadPathBuffer.EntityData.YFilter = eadPathBuffer.YFilter
    eadPathBuffer.EntityData.YangName = "ead-path-buffer"
    eadPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    eadPathBuffer.EntityData.ParentYangName = "internal-label"
    eadPathBuffer.EntityData.SegmentPath = "ead-path-buffer"
    eadPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eadPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eadPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eadPathBuffer.EntityData.Children = types.NewOrderedMap()
    eadPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    eadPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", eadPathBuffer.TunnelEndpointId})
    eadPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", eadPathBuffer.NextHop})
    eadPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", eadPathBuffer.OutputLabel})
    eadPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", eadPathBuffer.SrteTunnel})

    eadPathBuffer.EntityData.YListKeys = []string {}

    return &(eadPathBuffer.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel_EviPathBuffer
// EAD/EVI Path list buffer
type Evpn_Standby_InternalLabels_InternalLabel_EviPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (eviPathBuffer *Evpn_Standby_InternalLabels_InternalLabel_EviPathBuffer) GetEntityData() *types.CommonEntityData {
    eviPathBuffer.EntityData.YFilter = eviPathBuffer.YFilter
    eviPathBuffer.EntityData.YangName = "evi-path-buffer"
    eviPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    eviPathBuffer.EntityData.ParentYangName = "internal-label"
    eviPathBuffer.EntityData.SegmentPath = "evi-path-buffer"
    eviPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eviPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eviPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eviPathBuffer.EntityData.Children = types.NewOrderedMap()
    eviPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    eviPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", eviPathBuffer.TunnelEndpointId})
    eviPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", eviPathBuffer.NextHop})
    eviPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", eviPathBuffer.OutputLabel})
    eviPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", eviPathBuffer.SrteTunnel})

    eviPathBuffer.EntityData.YListKeys = []string {}

    return &(eviPathBuffer.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel_SummaryPathBuffer
// Summary Path list buffer
type Evpn_Standby_InternalLabels_InternalLabel_SummaryPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel Endpoint Identifier. The type is interface{} with range:
    // 0..4294967295.
    TunnelEndpointId interface{}

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SrteTunnel interface{}
}

func (summaryPathBuffer *Evpn_Standby_InternalLabels_InternalLabel_SummaryPathBuffer) GetEntityData() *types.CommonEntityData {
    summaryPathBuffer.EntityData.YFilter = summaryPathBuffer.YFilter
    summaryPathBuffer.EntityData.YangName = "summary-path-buffer"
    summaryPathBuffer.EntityData.BundleName = "cisco_ios_xr"
    summaryPathBuffer.EntityData.ParentYangName = "internal-label"
    summaryPathBuffer.EntityData.SegmentPath = "summary-path-buffer"
    summaryPathBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPathBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPathBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPathBuffer.EntityData.Children = types.NewOrderedMap()
    summaryPathBuffer.EntityData.Leafs = types.NewOrderedMap()
    summaryPathBuffer.EntityData.Leafs.Append("tunnel-endpoint-id", types.YLeaf{"TunnelEndpointId", summaryPathBuffer.TunnelEndpointId})
    summaryPathBuffer.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", summaryPathBuffer.NextHop})
    summaryPathBuffer.EntityData.Leafs.Append("output-label", types.YLeaf{"OutputLabel", summaryPathBuffer.OutputLabel})
    summaryPathBuffer.EntityData.Leafs.Append("srte-tunnel", types.YLeaf{"SrteTunnel", summaryPathBuffer.SrteTunnel})

    summaryPathBuffer.EntityData.YListKeys = []string {}

    return &(summaryPathBuffer.EntityData)
}

// Evpn_Standby_EthernetSegments
// EVPN Ethernet-Segment Table
type Evpn_Standby_EthernetSegments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet-Segment Entry. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment.
    EthernetSegment []*Evpn_Standby_EthernetSegments_EthernetSegment
}

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetEntityData() *types.CommonEntityData {
    ethernetSegments.EntityData.YFilter = ethernetSegments.YFilter
    ethernetSegments.EntityData.YangName = "ethernet-segments"
    ethernetSegments.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegments.EntityData.ParentYangName = "standby"
    ethernetSegments.EntityData.SegmentPath = "ethernet-segments"
    ethernetSegments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegments.EntityData.Children = types.NewOrderedMap()
    ethernetSegments.EntityData.Children.Append("ethernet-segment", types.YChild{"EthernetSegment", nil})
    for i := range ethernetSegments.EthernetSegment {
        ethernetSegments.EntityData.Children.Append(types.GetSegmentPath(ethernetSegments.EthernetSegment[i]), types.YChild{"EthernetSegment", ethernetSegments.EthernetSegment[i]})
    }
    ethernetSegments.EntityData.Leafs = types.NewOrderedMap()

    ethernetSegments.EntityData.YListKeys = []string {}

    return &(ethernetSegments.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment
// EVPN Ethernet-Segment Entry
type Evpn_Standby_EthernetSegments_EthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // ES id (part 1/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: [0-9a-fA-F]{1,8}.
    Esi5 interface{}

    // ESI Type. The type is L2vpnEvpnEsi.
    EsiType interface{}

    // ESI System Identifier. The type is string.
    EsiSystemIdentifier interface{}

    // ESI Port Key. The type is interface{} with range: 0..4294967295.
    EsiPortKey interface{}

    // ESI System Priority. The type is interface{} with range: 0..4294967295.
    EsiSystemPriority interface{}

    // Ethernet Segment Name. The type is string.
    EthernetSegmentName interface{}

    // State of the ethernet segment. The type is interface{} with range:
    // 0..4294967295.
    EthernetSegmentState interface{}

    // Main port ifhandle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    IfHandle interface{}

    // Main port redundancy group role. The type is L2vpnRgRole.
    MainPortRole interface{}

    // Main Port MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MainPortMac interface{}

    // Number of PWs in Up state. The type is interface{} with range:
    // 0..4294967295.
    NumUpPWs interface{}

    // ES-Import Route Target. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    RouteTarget interface{}

    // Origin of operational ES-Import RT. The type is L2vpnEvpnRtOrigin.
    RtOrigin interface{}

    // ES BGP Gates. The type is string.
    EsBgpGates interface{}

    // ES L2FIB Gates. The type is string.
    EsL2fibGates interface{}

    // Configured MAC Flushing mode. The type is L2vpnEvpnMfMode.
    MacFlushingModeConfig interface{}

    // Configured load balancing mode. The type is L2vpnEvpnLbMode.
    LoadBalanceModeConfig interface{}

    // Load balancing mode is default. The type is bool.
    LoadBalanceModeIsDefault interface{}

    // Operational load balancing mode. The type is L2vpnEvpnLbMode.
    LoadBalanceModeOper interface{}

    // Ethernet-Segment forced to single home. The type is bool.
    ForceSingleHome interface{}

    // Operational Source MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    SourceMacOper interface{}

    // Origin of operational source MAC address. The type is L2vpnEvpnSmacSrc.
    SourceMacOrigin interface{}

    // Configured timer for triggering DF election (seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    PeeringTimer interface{}

    // Milliseconds left on DF election timer. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    PeeringTimerLeft interface{}

    // Configured timer for (STP) recovery (seconds). The type is interface{} with
    // range: 0..4294967295. Units are second.
    RecoveryTimer interface{}

    // Milliseconds left on (STP) recovery timer. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RecoveryTimerLeft interface{}

    // Configured timer for delaying DF election (seconds). The type is
    // interface{} with range: 0..4294967295. Units are second.
    CarvingTimer interface{}

    // Milliseconds left on carving timer. The type is interface{} with range:
    // 0..4294967295. Units are millisecond.
    CarvingTimerLeft interface{}

    // Service carving mode. The type is L2vpnEvpnScMode.
    ServiceCarvingMode interface{}

    // Input string of Primary services ESI/I-SIDs. The type is string.
    PrimaryServicesInput interface{}

    // Input string of Secondary services ESI/I-SIDs. The type is string.
    SecondaryServicesInput interface{}

    // Count of Forwarders (AC, AC PW, VFI PW). The type is interface{} with
    // range: 0..4294967295.
    ForwarderPorts interface{}

    // Count of Forwarders with permanent service. The type is interface{} with
    // range: 0..4294967295.
    PermanentForwarderPorts interface{}

    // Count of Forwarders with elected service. The type is interface{} with
    // range: 0..4294967295.
    ElectedForwarderPorts interface{}

    // Count of Forwarders with not elected service. The type is interface{} with
    // range: 0..4294967295.
    NotElectedForwarderPorts interface{}

    // Count of forwarders with missing config detected. The type is interface{}
    // with range: 0..4294967295.
    NotConfigForwarderPorts interface{}

    // MP is protected and not under EVPN control. The type is bool.
    MpProtected interface{}

    // Anycast VTEP mode on NVE main-interface. The type is bool.
    NveAnycastVtep interface{}

    // Ingress-Replication is configured on NVE main-interface. The type is bool.
    NveIngressReplication interface{}

    // Local split horizon group label is valid. The type is bool.
    LocalSplitHorizonGroupLabelValid interface{}

    // Local split horizon group label. The type is interface{} with range:
    // 0..4294967295.
    LocalSplitHorizonGroupLabel interface{}

    // Ethernet Segment id. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []*Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier

    // List of Primary services ESI/I-SIDs. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService.
    PrimaryService []*Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService

    // List of Secondary services ESI/I-SIDs. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService.
    SecondaryService []*Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService

    // Elected ISID service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult.
    ServiceCarvingISidelectedResult []*Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult

    // Not elected ISID service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult.
    ServiceCarvingIsidNotElectedResult []*Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult

    // Elected EVI service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult.
    ServiceCarvingEviElectedResult []*Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult

    // Not elected EVI service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult.
    ServiceCarvingEviNotElectedResult []*Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult

    // Elected VNI service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult.
    ServiceCarvingVniElectedResult []*Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult

    // Not elected VNI service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult.
    ServiceCarvingVniNotElectedResult []*Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult

    // List of nexthop IPv6 addresses. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_NextHop.
    NextHop []*Evpn_Standby_EthernetSegments_EthernetSegment_NextHop

    // Permanent EVPN VPWS service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult.
    ServiceCarvingVpwsPermanentResult []*Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult

    // Remote split horizon group labels. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel.
    RemoteSplitHorizonGroupLabel []*Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
}

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetEntityData() *types.CommonEntityData {
    ethernetSegment.EntityData.YFilter = ethernetSegment.YFilter
    ethernetSegment.EntityData.YangName = "ethernet-segment"
    ethernetSegment.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegment.EntityData.ParentYangName = "ethernet-segments"
    ethernetSegment.EntityData.SegmentPath = "ethernet-segment"
    ethernetSegment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegment.EntityData.Children = types.NewOrderedMap()
    ethernetSegment.EntityData.Children.Append("ethernet-segment-identifier", types.YChild{"EthernetSegmentIdentifier", nil})
    for i := range ethernetSegment.EthernetSegmentIdentifier {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.EthernetSegmentIdentifier[i]), types.YChild{"EthernetSegmentIdentifier", ethernetSegment.EthernetSegmentIdentifier[i]})
    }
    ethernetSegment.EntityData.Children.Append("primary-service", types.YChild{"PrimaryService", nil})
    for i := range ethernetSegment.PrimaryService {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.PrimaryService[i]), types.YChild{"PrimaryService", ethernetSegment.PrimaryService[i]})
    }
    ethernetSegment.EntityData.Children.Append("secondary-service", types.YChild{"SecondaryService", nil})
    for i := range ethernetSegment.SecondaryService {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.SecondaryService[i]), types.YChild{"SecondaryService", ethernetSegment.SecondaryService[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-i-sidelected-result", types.YChild{"ServiceCarvingISidelectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingISidelectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingISidelectedResult[i]), types.YChild{"ServiceCarvingISidelectedResult", ethernetSegment.ServiceCarvingISidelectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-isid-not-elected-result", types.YChild{"ServiceCarvingIsidNotElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingIsidNotElectedResult[i]), types.YChild{"ServiceCarvingIsidNotElectedResult", ethernetSegment.ServiceCarvingIsidNotElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-evi-elected-result", types.YChild{"ServiceCarvingEviElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingEviElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingEviElectedResult[i]), types.YChild{"ServiceCarvingEviElectedResult", ethernetSegment.ServiceCarvingEviElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-evi-not-elected-result", types.YChild{"ServiceCarvingEviNotElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingEviNotElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingEviNotElectedResult[i]), types.YChild{"ServiceCarvingEviNotElectedResult", ethernetSegment.ServiceCarvingEviNotElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-vni-elected-result", types.YChild{"ServiceCarvingVniElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingVniElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingVniElectedResult[i]), types.YChild{"ServiceCarvingVniElectedResult", ethernetSegment.ServiceCarvingVniElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-vni-not-elected-result", types.YChild{"ServiceCarvingVniNotElectedResult", nil})
    for i := range ethernetSegment.ServiceCarvingVniNotElectedResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingVniNotElectedResult[i]), types.YChild{"ServiceCarvingVniNotElectedResult", ethernetSegment.ServiceCarvingVniNotElectedResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range ethernetSegment.NextHop {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.NextHop[i]), types.YChild{"NextHop", ethernetSegment.NextHop[i]})
    }
    ethernetSegment.EntityData.Children.Append("service-carving-vpws-permanent-result", types.YChild{"ServiceCarvingVpwsPermanentResult", nil})
    for i := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.ServiceCarvingVpwsPermanentResult[i]), types.YChild{"ServiceCarvingVpwsPermanentResult", ethernetSegment.ServiceCarvingVpwsPermanentResult[i]})
    }
    ethernetSegment.EntityData.Children.Append("remote-split-horizon-group-label", types.YChild{"RemoteSplitHorizonGroupLabel", nil})
    for i := range ethernetSegment.RemoteSplitHorizonGroupLabel {
        ethernetSegment.EntityData.Children.Append(types.GetSegmentPath(ethernetSegment.RemoteSplitHorizonGroupLabel[i]), types.YChild{"RemoteSplitHorizonGroupLabel", ethernetSegment.RemoteSplitHorizonGroupLabel[i]})
    }
    ethernetSegment.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegment.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ethernetSegment.InterfaceName})
    ethernetSegment.EntityData.Leafs.Append("esi1", types.YLeaf{"Esi1", ethernetSegment.Esi1})
    ethernetSegment.EntityData.Leafs.Append("esi2", types.YLeaf{"Esi2", ethernetSegment.Esi2})
    ethernetSegment.EntityData.Leafs.Append("esi3", types.YLeaf{"Esi3", ethernetSegment.Esi3})
    ethernetSegment.EntityData.Leafs.Append("esi4", types.YLeaf{"Esi4", ethernetSegment.Esi4})
    ethernetSegment.EntityData.Leafs.Append("esi5", types.YLeaf{"Esi5", ethernetSegment.Esi5})
    ethernetSegment.EntityData.Leafs.Append("esi-type", types.YLeaf{"EsiType", ethernetSegment.EsiType})
    ethernetSegment.EntityData.Leafs.Append("esi-system-identifier", types.YLeaf{"EsiSystemIdentifier", ethernetSegment.EsiSystemIdentifier})
    ethernetSegment.EntityData.Leafs.Append("esi-port-key", types.YLeaf{"EsiPortKey", ethernetSegment.EsiPortKey})
    ethernetSegment.EntityData.Leafs.Append("esi-system-priority", types.YLeaf{"EsiSystemPriority", ethernetSegment.EsiSystemPriority})
    ethernetSegment.EntityData.Leafs.Append("ethernet-segment-name", types.YLeaf{"EthernetSegmentName", ethernetSegment.EthernetSegmentName})
    ethernetSegment.EntityData.Leafs.Append("ethernet-segment-state", types.YLeaf{"EthernetSegmentState", ethernetSegment.EthernetSegmentState})
    ethernetSegment.EntityData.Leafs.Append("if-handle", types.YLeaf{"IfHandle", ethernetSegment.IfHandle})
    ethernetSegment.EntityData.Leafs.Append("main-port-role", types.YLeaf{"MainPortRole", ethernetSegment.MainPortRole})
    ethernetSegment.EntityData.Leafs.Append("main-port-mac", types.YLeaf{"MainPortMac", ethernetSegment.MainPortMac})
    ethernetSegment.EntityData.Leafs.Append("num-up-p-ws", types.YLeaf{"NumUpPWs", ethernetSegment.NumUpPWs})
    ethernetSegment.EntityData.Leafs.Append("route-target", types.YLeaf{"RouteTarget", ethernetSegment.RouteTarget})
    ethernetSegment.EntityData.Leafs.Append("rt-origin", types.YLeaf{"RtOrigin", ethernetSegment.RtOrigin})
    ethernetSegment.EntityData.Leafs.Append("es-bgp-gates", types.YLeaf{"EsBgpGates", ethernetSegment.EsBgpGates})
    ethernetSegment.EntityData.Leafs.Append("es-l2fib-gates", types.YLeaf{"EsL2fibGates", ethernetSegment.EsL2fibGates})
    ethernetSegment.EntityData.Leafs.Append("mac-flushing-mode-config", types.YLeaf{"MacFlushingModeConfig", ethernetSegment.MacFlushingModeConfig})
    ethernetSegment.EntityData.Leafs.Append("load-balance-mode-config", types.YLeaf{"LoadBalanceModeConfig", ethernetSegment.LoadBalanceModeConfig})
    ethernetSegment.EntityData.Leafs.Append("load-balance-mode-is-default", types.YLeaf{"LoadBalanceModeIsDefault", ethernetSegment.LoadBalanceModeIsDefault})
    ethernetSegment.EntityData.Leafs.Append("load-balance-mode-oper", types.YLeaf{"LoadBalanceModeOper", ethernetSegment.LoadBalanceModeOper})
    ethernetSegment.EntityData.Leafs.Append("force-single-home", types.YLeaf{"ForceSingleHome", ethernetSegment.ForceSingleHome})
    ethernetSegment.EntityData.Leafs.Append("source-mac-oper", types.YLeaf{"SourceMacOper", ethernetSegment.SourceMacOper})
    ethernetSegment.EntityData.Leafs.Append("source-mac-origin", types.YLeaf{"SourceMacOrigin", ethernetSegment.SourceMacOrigin})
    ethernetSegment.EntityData.Leafs.Append("peering-timer", types.YLeaf{"PeeringTimer", ethernetSegment.PeeringTimer})
    ethernetSegment.EntityData.Leafs.Append("peering-timer-left", types.YLeaf{"PeeringTimerLeft", ethernetSegment.PeeringTimerLeft})
    ethernetSegment.EntityData.Leafs.Append("recovery-timer", types.YLeaf{"RecoveryTimer", ethernetSegment.RecoveryTimer})
    ethernetSegment.EntityData.Leafs.Append("recovery-timer-left", types.YLeaf{"RecoveryTimerLeft", ethernetSegment.RecoveryTimerLeft})
    ethernetSegment.EntityData.Leafs.Append("carving-timer", types.YLeaf{"CarvingTimer", ethernetSegment.CarvingTimer})
    ethernetSegment.EntityData.Leafs.Append("carving-timer-left", types.YLeaf{"CarvingTimerLeft", ethernetSegment.CarvingTimerLeft})
    ethernetSegment.EntityData.Leafs.Append("service-carving-mode", types.YLeaf{"ServiceCarvingMode", ethernetSegment.ServiceCarvingMode})
    ethernetSegment.EntityData.Leafs.Append("primary-services-input", types.YLeaf{"PrimaryServicesInput", ethernetSegment.PrimaryServicesInput})
    ethernetSegment.EntityData.Leafs.Append("secondary-services-input", types.YLeaf{"SecondaryServicesInput", ethernetSegment.SecondaryServicesInput})
    ethernetSegment.EntityData.Leafs.Append("forwarder-ports", types.YLeaf{"ForwarderPorts", ethernetSegment.ForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("permanent-forwarder-ports", types.YLeaf{"PermanentForwarderPorts", ethernetSegment.PermanentForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("elected-forwarder-ports", types.YLeaf{"ElectedForwarderPorts", ethernetSegment.ElectedForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("not-elected-forwarder-ports", types.YLeaf{"NotElectedForwarderPorts", ethernetSegment.NotElectedForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("not-config-forwarder-ports", types.YLeaf{"NotConfigForwarderPorts", ethernetSegment.NotConfigForwarderPorts})
    ethernetSegment.EntityData.Leafs.Append("mp-protected", types.YLeaf{"MpProtected", ethernetSegment.MpProtected})
    ethernetSegment.EntityData.Leafs.Append("nve-anycast-vtep", types.YLeaf{"NveAnycastVtep", ethernetSegment.NveAnycastVtep})
    ethernetSegment.EntityData.Leafs.Append("nve-ingress-replication", types.YLeaf{"NveIngressReplication", ethernetSegment.NveIngressReplication})
    ethernetSegment.EntityData.Leafs.Append("local-split-horizon-group-label-valid", types.YLeaf{"LocalSplitHorizonGroupLabelValid", ethernetSegment.LocalSplitHorizonGroupLabelValid})
    ethernetSegment.EntityData.Leafs.Append("local-split-horizon-group-label", types.YLeaf{"LocalSplitHorizonGroupLabel", ethernetSegment.LocalSplitHorizonGroupLabel})

    ethernetSegment.EntityData.YListKeys = []string {}

    return &(ethernetSegment.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetEntityData() *types.CommonEntityData {
    ethernetSegmentIdentifier.EntityData.YFilter = ethernetSegmentIdentifier.YFilter
    ethernetSegmentIdentifier.EntityData.YangName = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.BundleName = "cisco_ios_xr"
    ethernetSegmentIdentifier.EntityData.ParentYangName = "ethernet-segment"
    ethernetSegmentIdentifier.EntityData.SegmentPath = "ethernet-segment-identifier"
    ethernetSegmentIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetSegmentIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetSegmentIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetSegmentIdentifier.EntityData.Children = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs = types.NewOrderedMap()
    ethernetSegmentIdentifier.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry})

    ethernetSegmentIdentifier.EntityData.YListKeys = []string {}

    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService
// List of Primary services ESI/I-SIDs
type Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetEntityData() *types.CommonEntityData {
    primaryService.EntityData.YFilter = primaryService.YFilter
    primaryService.EntityData.YangName = "primary-service"
    primaryService.EntityData.BundleName = "cisco_ios_xr"
    primaryService.EntityData.ParentYangName = "ethernet-segment"
    primaryService.EntityData.SegmentPath = "primary-service"
    primaryService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryService.EntityData.Children = types.NewOrderedMap()
    primaryService.EntityData.Leafs = types.NewOrderedMap()
    primaryService.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", primaryService.Entry})

    primaryService.EntityData.YListKeys = []string {}

    return &(primaryService.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService
// List of Secondary services ESI/I-SIDs
type Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetEntityData() *types.CommonEntityData {
    secondaryService.EntityData.YFilter = secondaryService.YFilter
    secondaryService.EntityData.YangName = "secondary-service"
    secondaryService.EntityData.BundleName = "cisco_ios_xr"
    secondaryService.EntityData.ParentYangName = "ethernet-segment"
    secondaryService.EntityData.SegmentPath = "secondary-service"
    secondaryService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryService.EntityData.Children = types.NewOrderedMap()
    secondaryService.EntityData.Leafs = types.NewOrderedMap()
    secondaryService.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", secondaryService.Entry})

    secondaryService.EntityData.YListKeys = []string {}

    return &(secondaryService.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult
// Elected ISID service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingISidelectedResult.EntityData.YFilter = serviceCarvingISidelectedResult.YFilter
    serviceCarvingISidelectedResult.EntityData.YangName = "service-carving-i-sidelected-result"
    serviceCarvingISidelectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingISidelectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingISidelectedResult.EntityData.SegmentPath = "service-carving-i-sidelected-result"
    serviceCarvingISidelectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingISidelectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingISidelectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingISidelectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingISidelectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingISidelectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingISidelectedResult.Entry})

    serviceCarvingISidelectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingISidelectedResult.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult
// Not elected ISID service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingIsidNotElectedResult.EntityData.YFilter = serviceCarvingIsidNotElectedResult.YFilter
    serviceCarvingIsidNotElectedResult.EntityData.YangName = "service-carving-isid-not-elected-result"
    serviceCarvingIsidNotElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingIsidNotElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingIsidNotElectedResult.EntityData.SegmentPath = "service-carving-isid-not-elected-result"
    serviceCarvingIsidNotElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingIsidNotElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingIsidNotElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingIsidNotElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingIsidNotElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingIsidNotElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingIsidNotElectedResult.Entry})

    serviceCarvingIsidNotElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingIsidNotElectedResult.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult
// Elected EVI service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingEviElectedResult.EntityData.YFilter = serviceCarvingEviElectedResult.YFilter
    serviceCarvingEviElectedResult.EntityData.YangName = "service-carving-evi-elected-result"
    serviceCarvingEviElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingEviElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingEviElectedResult.EntityData.SegmentPath = "service-carving-evi-elected-result"
    serviceCarvingEviElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingEviElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingEviElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingEviElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingEviElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingEviElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingEviElectedResult.Entry})

    serviceCarvingEviElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingEviElectedResult.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult
// Not elected EVI service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingEviNotElectedResult.EntityData.YFilter = serviceCarvingEviNotElectedResult.YFilter
    serviceCarvingEviNotElectedResult.EntityData.YangName = "service-carving-evi-not-elected-result"
    serviceCarvingEviNotElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingEviNotElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingEviNotElectedResult.EntityData.SegmentPath = "service-carving-evi-not-elected-result"
    serviceCarvingEviNotElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingEviNotElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingEviNotElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingEviNotElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingEviNotElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingEviNotElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingEviNotElectedResult.Entry})

    serviceCarvingEviNotElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingEviNotElectedResult.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult
// Elected VNI service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingVniElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingVniElectedResult.EntityData.YFilter = serviceCarvingVniElectedResult.YFilter
    serviceCarvingVniElectedResult.EntityData.YangName = "service-carving-vni-elected-result"
    serviceCarvingVniElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingVniElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingVniElectedResult.EntityData.SegmentPath = "service-carving-vni-elected-result"
    serviceCarvingVniElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingVniElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingVniElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingVniElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingVniElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingVniElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingVniElectedResult.Entry})

    serviceCarvingVniElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingVniElectedResult.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult
// Not elected VNI service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingVniNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVniNotElectedResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingVniNotElectedResult.EntityData.YFilter = serviceCarvingVniNotElectedResult.YFilter
    serviceCarvingVniNotElectedResult.EntityData.YangName = "service-carving-vni-not-elected-result"
    serviceCarvingVniNotElectedResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingVniNotElectedResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingVniNotElectedResult.EntityData.SegmentPath = "service-carving-vni-not-elected-result"
    serviceCarvingVniNotElectedResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingVniNotElectedResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingVniNotElectedResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingVniNotElectedResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingVniNotElectedResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingVniNotElectedResult.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", serviceCarvingVniNotElectedResult.Entry})

    serviceCarvingVniNotElectedResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingVniNotElectedResult.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_NextHop
// List of nexthop IPv6 addresses
type Evpn_Standby_EthernetSegments_EthernetSegment_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // DF Dont Premption. The type is bool.
    DfDontPrempt interface{}

    // DF Election Mode Configured. The type is interface{} with range: 0..255.
    DfType interface{}

    // DF Election Preference Set. The type is interface{} with range: 0..65535.
    DfPref interface{}
}

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "ethernet-segment"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", nextHop.NextHop})
    nextHop.EntityData.Leafs.Append("df-dont-prempt", types.YLeaf{"DfDontPrempt", nextHop.DfDontPrempt})
    nextHop.EntityData.Leafs.Append("df-type", types.YLeaf{"DfType", nextHop.DfType})
    nextHop.EntityData.Leafs.Append("df-pref", types.YLeaf{"DfPref", nextHop.DfPref})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult
// Permanent EVPN VPWS service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPN ID. The type is interface{} with range: 0..4294967295.
    VpnId interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}
}

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetEntityData() *types.CommonEntityData {
    serviceCarvingVpwsPermanentResult.EntityData.YFilter = serviceCarvingVpwsPermanentResult.YFilter
    serviceCarvingVpwsPermanentResult.EntityData.YangName = "service-carving-vpws-permanent-result"
    serviceCarvingVpwsPermanentResult.EntityData.BundleName = "cisco_ios_xr"
    serviceCarvingVpwsPermanentResult.EntityData.ParentYangName = "ethernet-segment"
    serviceCarvingVpwsPermanentResult.EntityData.SegmentPath = "service-carving-vpws-permanent-result"
    serviceCarvingVpwsPermanentResult.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceCarvingVpwsPermanentResult.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceCarvingVpwsPermanentResult.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceCarvingVpwsPermanentResult.EntityData.Children = types.NewOrderedMap()
    serviceCarvingVpwsPermanentResult.EntityData.Leafs = types.NewOrderedMap()
    serviceCarvingVpwsPermanentResult.EntityData.Leafs.Append("vpn-id", types.YLeaf{"VpnId", serviceCarvingVpwsPermanentResult.VpnId})
    serviceCarvingVpwsPermanentResult.EntityData.Leafs.Append("type", types.YLeaf{"Type", serviceCarvingVpwsPermanentResult.Type})
    serviceCarvingVpwsPermanentResult.EntityData.Leafs.Append("ethernet-tag", types.YLeaf{"EthernetTag", serviceCarvingVpwsPermanentResult.EthernetTag})

    serviceCarvingVpwsPermanentResult.EntityData.YListKeys = []string {}

    return &(serviceCarvingVpwsPermanentResult.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Split horizon label associated with next-hop address. The type is
    // interface{} with range: 0..4294967295.
    Label interface{}
}

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetEntityData() *types.CommonEntityData {
    remoteSplitHorizonGroupLabel.EntityData.YFilter = remoteSplitHorizonGroupLabel.YFilter
    remoteSplitHorizonGroupLabel.EntityData.YangName = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.BundleName = "cisco_ios_xr"
    remoteSplitHorizonGroupLabel.EntityData.ParentYangName = "ethernet-segment"
    remoteSplitHorizonGroupLabel.EntityData.SegmentPath = "remote-split-horizon-group-label"
    remoteSplitHorizonGroupLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteSplitHorizonGroupLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteSplitHorizonGroupLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteSplitHorizonGroupLabel.EntityData.Children = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs = types.NewOrderedMap()
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", remoteSplitHorizonGroupLabel.NextHop})
    remoteSplitHorizonGroupLabel.EntityData.Leafs.Append("label", types.YLeaf{"Label", remoteSplitHorizonGroupLabel.Label})

    remoteSplitHorizonGroupLabel.EntityData.YListKeys = []string {}

    return &(remoteSplitHorizonGroupLabel.EntityData)
}

// Evpn_Standby_AcIds
// EVPN AC ID table
type Evpn_Standby_AcIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN AC ID table. The type is slice of Evpn_Standby_AcIds_AcId.
    AcId []*Evpn_Standby_AcIds_AcId
}

func (acIds *Evpn_Standby_AcIds) GetEntityData() *types.CommonEntityData {
    acIds.EntityData.YFilter = acIds.YFilter
    acIds.EntityData.YangName = "ac-ids"
    acIds.EntityData.BundleName = "cisco_ios_xr"
    acIds.EntityData.ParentYangName = "standby"
    acIds.EntityData.SegmentPath = "ac-ids"
    acIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acIds.EntityData.Children = types.NewOrderedMap()
    acIds.EntityData.Children.Append("ac-id", types.YChild{"AcId", nil})
    for i := range acIds.AcId {
        acIds.EntityData.Children.Append(types.GetSegmentPath(acIds.AcId[i]), types.YChild{"AcId", acIds.AcId[i]})
    }
    acIds.EntityData.Leafs = types.NewOrderedMap()

    acIds.EntityData.YListKeys = []string {}

    return &(acIds.EntityData)
}

// Evpn_Standby_AcIds_AcId
// EVPN AC ID table
type Evpn_Standby_AcIds_AcId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: 0..4294967295.
    Evi interface{}

    // AC ID. The type is interface{} with range: 0..4294967295.
    AcId interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}

    // EVPN Instance summary information.
    EvpnInstance Evpn_Standby_AcIds_AcId_EvpnInstance
}

func (acId *Evpn_Standby_AcIds_AcId) GetEntityData() *types.CommonEntityData {
    acId.EntityData.YFilter = acId.YFilter
    acId.EntityData.YangName = "ac-id"
    acId.EntityData.BundleName = "cisco_ios_xr"
    acId.EntityData.ParentYangName = "ac-ids"
    acId.EntityData.SegmentPath = "ac-id"
    acId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acId.EntityData.Children = types.NewOrderedMap()
    acId.EntityData.Children.Append("evpn-instance", types.YChild{"EvpnInstance", &acId.EvpnInstance})
    acId.EntityData.Leafs = types.NewOrderedMap()
    acId.EntityData.Leafs.Append("evi", types.YLeaf{"Evi", acId.Evi})
    acId.EntityData.Leafs.Append("ac-id", types.YLeaf{"AcId", acId.AcId})
    acId.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", acId.Neighbor})

    acId.EntityData.YListKeys = []string {}

    return &(acId.EntityData)
}

// Evpn_Standby_AcIds_AcId_EvpnInstance
// EVPN Instance summary information
type Evpn_Standby_AcIds_AcId_EvpnInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnId interface{}

    // EVPN Instance transport encapsulation. The type is interface{} with range:
    // 0..255.
    EncapsulationXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evpnInstance *Evpn_Standby_AcIds_AcId_EvpnInstance) GetEntityData() *types.CommonEntityData {
    evpnInstance.EntityData.YFilter = evpnInstance.YFilter
    evpnInstance.EntityData.YangName = "evpn-instance"
    evpnInstance.EntityData.BundleName = "cisco_ios_xr"
    evpnInstance.EntityData.ParentYangName = "ac-id"
    evpnInstance.EntityData.SegmentPath = "evpn-instance"
    evpnInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evpnInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evpnInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evpnInstance.EntityData.Children = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs = types.NewOrderedMap()
    evpnInstance.EntityData.Leafs.Append("ethernet-vpn-id", types.YLeaf{"EthernetVpnId", evpnInstance.EthernetVpnId})
    evpnInstance.EntityData.Leafs.Append("encapsulation-xr", types.YLeaf{"EncapsulationXr", evpnInstance.EncapsulationXr})
    evpnInstance.EntityData.Leafs.Append("bd-name", types.YLeaf{"BdName", evpnInstance.BdName})
    evpnInstance.EntityData.Leafs.Append("type", types.YLeaf{"Type", evpnInstance.Type})

    evpnInstance.EntityData.YListKeys = []string {}

    return &(evpnInstance.EntityData)
}

