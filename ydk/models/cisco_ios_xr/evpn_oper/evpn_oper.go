// This module contains a collection of YANG definitions
// for Cisco IOS-XR evpn package operational data.
// 
// This module contains definitions
// for the following management objects:
//   evpn: EVPN Operational Table
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

// BgpRouteTarget represents Bgp route target
type BgpRouteTarget string

const (
    // RT is default type
    BgpRouteTarget_no_stitching BgpRouteTarget = "no-stitching"

    // RT is for stitching (Golf-L2)
    BgpRouteTarget_stitching BgpRouteTarget = "stitching"
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

// L2vpnEvpnScMode represents EVPN Ethernet-Segment service carving mode
type L2vpnEvpnScMode string

const (
    // Invalid service carving mode
    L2vpnEvpnScMode_invalid L2vpnEvpnScMode = "invalid"

    // Auto service carving mode
    L2vpnEvpnScMode_auto L2vpnEvpnScMode = "auto"

    // Manual service carving
    L2vpnEvpnScMode_manual L2vpnEvpnScMode = "manual"
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

    evpn.EntityData.Children = make(map[string]types.YChild)
    evpn.EntityData.Children["nodes"] = types.YChild{"Nodes", &evpn.Nodes}
    evpn.EntityData.Children["active"] = types.YChild{"Active", &evpn.Active}
    evpn.EntityData.Children["standby"] = types.YChild{"Standby", &evpn.Standby}
    evpn.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Node []Evpn_Nodes_Node
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

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Evpn_Nodes_Node
// EVPN operational data for a particular node
type Evpn_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // L2VPN EVPN EVI Table.
    Evis Evpn_Nodes_Node_Evis

    // L2VPN EVPN Summary.
    Summary Evpn_Nodes_Node_Summary

    // L2VPN EVI Detail Table.
    EviDetail Evpn_Nodes_Node_EviDetail

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
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["evis"] = types.YChild{"Evis", &node.Evis}
    node.EntityData.Children["summary"] = types.YChild{"Summary", &node.Summary}
    node.EntityData.Children["evi-detail"] = types.YChild{"EviDetail", &node.EviDetail}
    node.EntityData.Children["internal-labels"] = types.YChild{"InternalLabels", &node.InternalLabels}
    node.EntityData.Children["ethernet-segments"] = types.YChild{"EthernetSegments", &node.EthernetSegments}
    node.EntityData.Children["ac-ids"] = types.YChild{"AcIds", &node.AcIds}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    return &(node.EntityData)
}

// Evpn_Nodes_Node_Evis
// L2VPN EVPN EVI Table
type Evpn_Nodes_Node_Evis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Entry. The type is slice of Evpn_Nodes_Node_Evis_Evi.
    Evi []Evpn_Nodes_Node_Evis_Evi
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

    evis.EntityData.Children = make(map[string]types.YChild)
    evis.EntityData.Children["evi"] = types.YChild{"Evi", nil}
    for i := range evis.Evi {
        evis.EntityData.Children[types.GetSegmentPath(&evis.Evi[i])] = types.YChild{"Evi", &evis.Evi[i]}
    }
    evis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(evis.EntityData)
}

// Evpn_Nodes_Node_Evis_Evi
// L2VPN EVPN EVI Entry
type Evpn_Nodes_Node_Evis_Evi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN id. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type_ interface{}
}

func (evi *Evpn_Nodes_Node_Evis_Evi) GetEntityData() *types.CommonEntityData {
    evi.EntityData.YFilter = evi.YFilter
    evi.EntityData.YangName = "evi"
    evi.EntityData.BundleName = "cisco_ios_xr"
    evi.EntityData.ParentYangName = "evis"
    evi.EntityData.SegmentPath = "evi" + "[evi='" + fmt.Sprintf("%v", evi.Evi) + "']"
    evi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evi.EntityData.Children = make(map[string]types.YChild)
    evi.EntityData.Leafs = make(map[string]types.YLeaf)
    evi.EntityData.Leafs["evi"] = types.YLeaf{"Evi", evi.Evi}
    evi.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", evi.EviXr}
    evi.EntityData.Leafs["bd-name"] = types.YLeaf{"BdName", evi.BdName}
    evi.EntityData.Leafs["type"] = types.YLeaf{"Type_", evi.Type_}
    return &(evi.EntityData)
}

// Evpn_Nodes_Node_Summary
// L2VPN EVPN Summary
type Evpn_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Router ID. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // BGP AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // Number of EVI DB Entries. The type is interface{} with range:
    // 0..4294967295.
    EvIs interface{}

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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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
    L2RibThrottle interface{}

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

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", summary.RouterId}
    summary.EntityData.Leafs["as"] = types.YLeaf{"As", summary.As}
    summary.EntityData.Leafs["ev-is"] = types.YLeaf{"EvIs", summary.EvIs}
    summary.EntityData.Leafs["local-mac-routes"] = types.YLeaf{"LocalMacRoutes", summary.LocalMacRoutes}
    summary.EntityData.Leafs["local-ipv4-mac-routes"] = types.YLeaf{"LocalIpv4MacRoutes", summary.LocalIpv4MacRoutes}
    summary.EntityData.Leafs["local-ipv6-mac-routes"] = types.YLeaf{"LocalIpv6MacRoutes", summary.LocalIpv6MacRoutes}
    summary.EntityData.Leafs["es-global-mac-routes"] = types.YLeaf{"EsGlobalMacRoutes", summary.EsGlobalMacRoutes}
    summary.EntityData.Leafs["remote-mac-routes"] = types.YLeaf{"RemoteMacRoutes", summary.RemoteMacRoutes}
    summary.EntityData.Leafs["remote-soo-mac-routes"] = types.YLeaf{"RemoteSooMacRoutes", summary.RemoteSooMacRoutes}
    summary.EntityData.Leafs["remote-ipv4-mac-routes"] = types.YLeaf{"RemoteIpv4MacRoutes", summary.RemoteIpv4MacRoutes}
    summary.EntityData.Leafs["remote-ipv6-mac-routes"] = types.YLeaf{"RemoteIpv6MacRoutes", summary.RemoteIpv6MacRoutes}
    summary.EntityData.Leafs["local-imcast-routes"] = types.YLeaf{"LocalImcastRoutes", summary.LocalImcastRoutes}
    summary.EntityData.Leafs["remote-imcast-routes"] = types.YLeaf{"RemoteImcastRoutes", summary.RemoteImcastRoutes}
    summary.EntityData.Leafs["labels"] = types.YLeaf{"Labels", summary.Labels}
    summary.EntityData.Leafs["es-entries"] = types.YLeaf{"EsEntries", summary.EsEntries}
    summary.EntityData.Leafs["neighbor-entries"] = types.YLeaf{"NeighborEntries", summary.NeighborEntries}
    summary.EntityData.Leafs["local-ead-routes"] = types.YLeaf{"LocalEadRoutes", summary.LocalEadRoutes}
    summary.EntityData.Leafs["remote-ead-routes"] = types.YLeaf{"RemoteEadRoutes", summary.RemoteEadRoutes}
    summary.EntityData.Leafs["global-source-mac"] = types.YLeaf{"GlobalSourceMac", summary.GlobalSourceMac}
    summary.EntityData.Leafs["peering-time"] = types.YLeaf{"PeeringTime", summary.PeeringTime}
    summary.EntityData.Leafs["recovery-time"] = types.YLeaf{"RecoveryTime", summary.RecoveryTime}
    summary.EntityData.Leafs["carving-time"] = types.YLeaf{"CarvingTime", summary.CarvingTime}
    summary.EntityData.Leafs["mac-secure-move-count"] = types.YLeaf{"MacSecureMoveCount", summary.MacSecureMoveCount}
    summary.EntityData.Leafs["mac-secure-move-interval"] = types.YLeaf{"MacSecureMoveInterval", summary.MacSecureMoveInterval}
    summary.EntityData.Leafs["mac-secure-freeze-time"] = types.YLeaf{"MacSecureFreezeTime", summary.MacSecureFreezeTime}
    summary.EntityData.Leafs["mac-secure-retry-count"] = types.YLeaf{"MacSecureRetryCount", summary.MacSecureRetryCount}
    summary.EntityData.Leafs["cost-out"] = types.YLeaf{"CostOut", summary.CostOut}
    summary.EntityData.Leafs["startup-cost-in-time"] = types.YLeaf{"StartupCostInTime", summary.StartupCostInTime}
    summary.EntityData.Leafs["l2rib-throttle"] = types.YLeaf{"L2RibThrottle", summary.L2RibThrottle}
    summary.EntityData.Leafs["logging-df-election-enabled"] = types.YLeaf{"LoggingDfElectionEnabled", summary.LoggingDfElectionEnabled}
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

    eviDetail.EntityData.Children = make(map[string]types.YChild)
    eviDetail.EntityData.Children["elements"] = types.YChild{"Elements", &eviDetail.Elements}
    eviDetail.EntityData.Children["evi-children"] = types.YChild{"EviChildren", &eviDetail.EviChildren}
    eviDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eviDetail.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements
// EVI BGP RT Detail Info Elements
type Evpn_Nodes_Node_EviDetail_Elements struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info. The type is slice of
    // Evpn_Nodes_Node_EviDetail_Elements_Element.
    Element []Evpn_Nodes_Node_EviDetail_Elements_Element
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

    elements.EntityData.Children = make(map[string]types.YChild)
    elements.EntityData.Children["element"] = types.YChild{"Element", nil}
    for i := range elements.Element {
        elements.EntityData.Children[types.GetSegmentPath(&elements.Element[i])] = types.YChild{"Element", &elements.Element[i]}
    }
    elements.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(elements.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element
// EVI BGP RT Detail Info
type Evpn_Nodes_Node_EviDetail_Elements_Element struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN id. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // EVI description. The type is string.
    Description interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type_ interface{}

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

    // EVPN Instance is Regular/Stitching side. The type is interface{} with
    // range: 0..255.
    Stitching interface{}

    // EVPN Instance encapsulation. The type is interface{} with range: 0..255.
    Encapsulation interface{}

    // Flow Label Information.
    FlowLabel Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel

    // Automatic Route Distingtuisher.
    RdAuto Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto

    // Configured Route Distinguisher.
    RdConfigured Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured

    // Automatic Route Target.
    RtAuto Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto

    // Automatic Route Target Stitching.
    RtAutoStitching Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching
}

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetEntityData() *types.CommonEntityData {
    element.EntityData.YFilter = element.YFilter
    element.EntityData.YangName = "element"
    element.EntityData.BundleName = "cisco_ios_xr"
    element.EntityData.ParentYangName = "elements"
    element.EntityData.SegmentPath = "element" + "[evi='" + fmt.Sprintf("%v", element.Evi) + "']"
    element.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    element.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    element.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    element.EntityData.Children = make(map[string]types.YChild)
    element.EntityData.Children["flow-label"] = types.YChild{"FlowLabel", &element.FlowLabel}
    element.EntityData.Children["rd-auto"] = types.YChild{"RdAuto", &element.RdAuto}
    element.EntityData.Children["rd-configured"] = types.YChild{"RdConfigured", &element.RdConfigured}
    element.EntityData.Children["rt-auto"] = types.YChild{"RtAuto", &element.RtAuto}
    element.EntityData.Children["rt-auto-stitching"] = types.YChild{"RtAutoStitching", &element.RtAutoStitching}
    element.EntityData.Leafs = make(map[string]types.YLeaf)
    element.EntityData.Leafs["evi"] = types.YLeaf{"Evi", element.Evi}
    element.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", element.EviXr}
    element.EntityData.Leafs["description"] = types.YLeaf{"Description", element.Description}
    element.EntityData.Leafs["bd-name"] = types.YLeaf{"BdName", element.BdName}
    element.EntityData.Leafs["type"] = types.YLeaf{"Type_", element.Type_}
    element.EntityData.Leafs["unicast-label"] = types.YLeaf{"UnicastLabel", element.UnicastLabel}
    element.EntityData.Leafs["multicast-label"] = types.YLeaf{"MulticastLabel", element.MulticastLabel}
    element.EntityData.Leafs["cw-disable"] = types.YLeaf{"CwDisable", element.CwDisable}
    element.EntityData.Leafs["table-policy-name"] = types.YLeaf{"TablePolicyName", element.TablePolicyName}
    element.EntityData.Leafs["forward-class"] = types.YLeaf{"ForwardClass", element.ForwardClass}
    element.EntityData.Leafs["rt-import-block-set"] = types.YLeaf{"RtImportBlockSet", element.RtImportBlockSet}
    element.EntityData.Leafs["rt-export-block-set"] = types.YLeaf{"RtExportBlockSet", element.RtExportBlockSet}
    element.EntityData.Leafs["advertise-mac"] = types.YLeaf{"AdvertiseMac", element.AdvertiseMac}
    element.EntityData.Leafs["advertise-bvi-mac"] = types.YLeaf{"AdvertiseBviMac", element.AdvertiseBviMac}
    element.EntityData.Leafs["aliasing-disabled"] = types.YLeaf{"AliasingDisabled", element.AliasingDisabled}
    element.EntityData.Leafs["unknown-unicast-flooding-disabled"] = types.YLeaf{"UnknownUnicastFloodingDisabled", element.UnknownUnicastFloodingDisabled}
    element.EntityData.Leafs["reoriginate-disabled"] = types.YLeaf{"ReoriginateDisabled", element.ReoriginateDisabled}
    element.EntityData.Leafs["stitching"] = types.YLeaf{"Stitching", element.Stitching}
    element.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", element.Encapsulation}
    return &(element.EntityData)
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

    flowLabel.EntityData.Children = make(map[string]types.YChild)
    flowLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    flowLabel.EntityData.Leafs["static-flow-label"] = types.YLeaf{"StaticFlowLabel", flowLabel.StaticFlowLabel}
    flowLabel.EntityData.Leafs["global-flow-label"] = types.YLeaf{"GlobalFlowLabel", flowLabel.GlobalFlowLabel}
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

    rdAuto.EntityData.Children = make(map[string]types.YChild)
    rdAuto.EntityData.Children["auto"] = types.YChild{"Auto", &rdAuto.Auto}
    rdAuto.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rdAuto.TwoByteAs}
    rdAuto.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rdAuto.FourByteAs}
    rdAuto.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rdAuto.V4Addr}
    rdAuto.EntityData.Leafs = make(map[string]types.YLeaf)
    rdAuto.EntityData.Leafs["rd"] = types.YLeaf{"Rd", rdAuto.Rd}
    return &(rdAuto.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto
// auto
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    auto.EntityData.Children = make(map[string]types.YChild)
    auto.EntityData.Leafs = make(map[string]types.YLeaf)
    auto.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", auto.RouterId}
    auto.EntityData.Leafs["auto-index"] = types.YLeaf{"AutoIndex", auto.AutoIndex}
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

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
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

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
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

    rdConfigured.EntityData.Children = make(map[string]types.YChild)
    rdConfigured.EntityData.Children["auto"] = types.YChild{"Auto", &rdConfigured.Auto}
    rdConfigured.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rdConfigured.TwoByteAs}
    rdConfigured.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rdConfigured.FourByteAs}
    rdConfigured.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rdConfigured.V4Addr}
    rdConfigured.EntityData.Leafs = make(map[string]types.YLeaf)
    rdConfigured.EntityData.Leafs["rd"] = types.YLeaf{"Rd", rdConfigured.Rd}
    return &(rdConfigured.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto
// auto
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    auto.EntityData.Children = make(map[string]types.YChild)
    auto.EntityData.Leafs = make(map[string]types.YLeaf)
    auto.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", auto.RouterId}
    auto.EntityData.Leafs["auto-index"] = types.YLeaf{"AutoIndex", auto.AutoIndex}
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

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
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

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
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

    rtAuto.EntityData.Children = make(map[string]types.YChild)
    rtAuto.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rtAuto.TwoByteAs}
    rtAuto.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rtAuto.FourByteAs}
    rtAuto.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rtAuto.V4Addr}
    rtAuto.EntityData.Children["es-import"] = types.YChild{"EsImport", &rtAuto.EsImport}
    rtAuto.EntityData.Leafs = make(map[string]types.YLeaf)
    rtAuto.EntityData.Leafs["rt"] = types.YLeaf{"Rt", rtAuto.Rt}
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

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
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

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
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

    esImport.EntityData.Children = make(map[string]types.YChild)
    esImport.EntityData.Leafs = make(map[string]types.YLeaf)
    esImport.EntityData.Leafs["high-bytes"] = types.YLeaf{"HighBytes", esImport.HighBytes}
    esImport.EntityData.Leafs["low-bytes"] = types.YLeaf{"LowBytes", esImport.LowBytes}
    return &(esImport.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching
// Automatic Route Target Stitching
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs

    // v4 addr.
    V4Addr Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr

    // es import.
    EsImport Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport
}

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetEntityData() *types.CommonEntityData {
    rtAutoStitching.EntityData.YFilter = rtAutoStitching.YFilter
    rtAutoStitching.EntityData.YangName = "rt-auto-stitching"
    rtAutoStitching.EntityData.BundleName = "cisco_ios_xr"
    rtAutoStitching.EntityData.ParentYangName = "element"
    rtAutoStitching.EntityData.SegmentPath = "rt-auto-stitching"
    rtAutoStitching.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rtAutoStitching.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rtAutoStitching.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rtAutoStitching.EntityData.Children = make(map[string]types.YChild)
    rtAutoStitching.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rtAutoStitching.TwoByteAs}
    rtAutoStitching.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rtAutoStitching.FourByteAs}
    rtAutoStitching.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rtAutoStitching.V4Addr}
    rtAutoStitching.EntityData.Children["es-import"] = types.YChild{"EsImport", &rtAutoStitching.EsImport}
    rtAutoStitching.EntityData.Leafs = make(map[string]types.YLeaf)
    rtAutoStitching.EntityData.Leafs["rt"] = types.YLeaf{"Rt", rtAutoStitching.Rt}
    return &(rtAutoStitching.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rt-auto-stitching"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
    return &(twoByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rt-auto-stitching"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rt-auto-stitching"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
    return &(v4Addr.EntityData)
}

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport
// es import
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "rt-auto-stitching"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = make(map[string]types.YChild)
    esImport.EntityData.Leafs = make(map[string]types.YLeaf)
    esImport.EntityData.Leafs["high-bytes"] = types.YLeaf{"HighBytes", esImport.HighBytes}
    esImport.EntityData.Leafs["low-bytes"] = types.YLeaf{"LowBytes", esImport.LowBytes}
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

    eviChildren.EntityData.Children = make(map[string]types.YChild)
    eviChildren.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &eviChildren.Neighbors}
    eviChildren.EntityData.Children["ethernet-auto-discoveries"] = types.YChild{"EthernetAutoDiscoveries", &eviChildren.EthernetAutoDiscoveries}
    eviChildren.EntityData.Children["inclusive-multicasts"] = types.YChild{"InclusiveMulticasts", &eviChildren.InclusiveMulticasts}
    eviChildren.EntityData.Children["route-targets"] = types.YChild{"RouteTargets", &eviChildren.RouteTargets}
    eviChildren.EntityData.Children["macs"] = types.YChild{"Macs", &eviChildren.Macs}
    eviChildren.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eviChildren.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors
// EVPN Neighbor table
type Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Neighbor table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor.
    Neighbor []Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor
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

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children[types.GetSegmentPath(&neighbors.Neighbor[i])] = types.YChild{"Neighbor", &neighbors.Neighbor[i]}
    }
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor
// EVPN Neighbor table
type Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Neighbor IP. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}
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

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["evi"] = types.YLeaf{"Evi", neighbor.Evi}
    neighbor.EntityData.Leafs["neighbor-ip"] = types.YLeaf{"NeighborIp", neighbor.NeighborIp}
    neighbor.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", neighbor.EviXr}
    neighbor.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", neighbor.Neighbor}
    return &(neighbor.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries
// EVPN Ethernet Auto-Discovery table
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet Auto-Discovery Entry. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery.
    EthernetAutoDiscovery []Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
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

    ethernetAutoDiscoveries.EntityData.Children = make(map[string]types.YChild)
    ethernetAutoDiscoveries.EntityData.Children["ethernet-auto-discovery"] = types.YChild{"EthernetAutoDiscovery", nil}
    for i := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
        ethernetAutoDiscoveries.EntityData.Children[types.GetSegmentPath(&ethernetAutoDiscoveries.EthernetAutoDiscovery[i])] = types.YChild{"EthernetAutoDiscovery", &ethernetAutoDiscoveries.EthernetAutoDiscovery[i]}
    }
    ethernetAutoDiscoveries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetAutoDiscoveries.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
// EVPN Ethernet Auto-Discovery Entry
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // ES id (part 1/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnid interface{}

    // Service Type. The type is L2vpnEvpn.
    Type_ interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Local nexthop IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalNextHop interface{}

    // Associated local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Indication of EthernetAutoDiscovery Route is local. The type is bool.
    IsLocalEad interface{}

    // Encap type of local or remote EAD. The type is interface{} with range:
    // 0..255.
    Encap interface{}

    // Single-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleActive interface{}

    // Single-flow-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleFlowActive interface{}

    // Number of items in path list buffer. The type is interface{} with range:
    // 0..4294967295.
    NumPaths interface{}

    // Ethernet Segment id. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer.
    PathBuffer []Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
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

    ethernetAutoDiscovery.EntityData.Children = make(map[string]types.YChild)
    ethernetAutoDiscovery.EntityData.Children["ethernet-segment-identifier"] = types.YChild{"EthernetSegmentIdentifier", nil}
    for i := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
        ethernetAutoDiscovery.EntityData.Children[types.GetSegmentPath(&ethernetAutoDiscovery.EthernetSegmentIdentifier[i])] = types.YChild{"EthernetSegmentIdentifier", &ethernetAutoDiscovery.EthernetSegmentIdentifier[i]}
    }
    ethernetAutoDiscovery.EntityData.Children["path-buffer"] = types.YChild{"PathBuffer", nil}
    for i := range ethernetAutoDiscovery.PathBuffer {
        ethernetAutoDiscovery.EntityData.Children[types.GetSegmentPath(&ethernetAutoDiscovery.PathBuffer[i])] = types.YChild{"PathBuffer", &ethernetAutoDiscovery.PathBuffer[i]}
    }
    ethernetAutoDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetAutoDiscovery.EntityData.Leafs["evi"] = types.YLeaf{"Evi", ethernetAutoDiscovery.Evi}
    ethernetAutoDiscovery.EntityData.Leafs["esi1"] = types.YLeaf{"Esi1", ethernetAutoDiscovery.Esi1}
    ethernetAutoDiscovery.EntityData.Leafs["esi2"] = types.YLeaf{"Esi2", ethernetAutoDiscovery.Esi2}
    ethernetAutoDiscovery.EntityData.Leafs["esi3"] = types.YLeaf{"Esi3", ethernetAutoDiscovery.Esi3}
    ethernetAutoDiscovery.EntityData.Leafs["esi4"] = types.YLeaf{"Esi4", ethernetAutoDiscovery.Esi4}
    ethernetAutoDiscovery.EntityData.Leafs["esi5"] = types.YLeaf{"Esi5", ethernetAutoDiscovery.Esi5}
    ethernetAutoDiscovery.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", ethernetAutoDiscovery.EthernetTag}
    ethernetAutoDiscovery.EntityData.Leafs["ethernet-vpnid"] = types.YLeaf{"EthernetVpnid", ethernetAutoDiscovery.EthernetVpnid}
    ethernetAutoDiscovery.EntityData.Leafs["type"] = types.YLeaf{"Type_", ethernetAutoDiscovery.Type_}
    ethernetAutoDiscovery.EntityData.Leafs["ethernet-tag-xr"] = types.YLeaf{"EthernetTagXr", ethernetAutoDiscovery.EthernetTagXr}
    ethernetAutoDiscovery.EntityData.Leafs["local-next-hop"] = types.YLeaf{"LocalNextHop", ethernetAutoDiscovery.LocalNextHop}
    ethernetAutoDiscovery.EntityData.Leafs["local-label"] = types.YLeaf{"LocalLabel", ethernetAutoDiscovery.LocalLabel}
    ethernetAutoDiscovery.EntityData.Leafs["is-local-ead"] = types.YLeaf{"IsLocalEad", ethernetAutoDiscovery.IsLocalEad}
    ethernetAutoDiscovery.EntityData.Leafs["encap"] = types.YLeaf{"Encap", ethernetAutoDiscovery.Encap}
    ethernetAutoDiscovery.EntityData.Leafs["redundancy-single-active"] = types.YLeaf{"RedundancySingleActive", ethernetAutoDiscovery.RedundancySingleActive}
    ethernetAutoDiscovery.EntityData.Leafs["redundancy-single-flow-active"] = types.YLeaf{"RedundancySingleFlowActive", ethernetAutoDiscovery.RedundancySingleFlowActive}
    ethernetAutoDiscovery.EntityData.Leafs["num-paths"] = types.YLeaf{"NumPaths", ethernetAutoDiscovery.NumPaths}
    return &(ethernetAutoDiscovery.EntityData)
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

    ethernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    ethernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry}
    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
// Path List Buffer
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    pathBuffer.EntityData.Children = make(map[string]types.YChild)
    pathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    pathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", pathBuffer.NextHop}
    pathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", pathBuffer.OutputLabel}
    pathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel}
    return &(pathBuffer.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts
// L2VPN EVPN IMCAST table
type Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN IMCAST table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast.
    InclusiveMulticast []Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
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

    inclusiveMulticasts.EntityData.Children = make(map[string]types.YChild)
    inclusiveMulticasts.EntityData.Children["inclusive-multicast"] = types.YChild{"InclusiveMulticast", nil}
    for i := range inclusiveMulticasts.InclusiveMulticast {
        inclusiveMulticasts.EntityData.Children[types.GetSegmentPath(&inclusiveMulticasts.InclusiveMulticast[i])] = types.YChild{"InclusiveMulticast", &inclusiveMulticasts.InclusiveMulticast[i]}
    }
    inclusiveMulticasts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inclusiveMulticasts.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
// L2VPN EVPN IMCAST table
type Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag. The type is interface{} with range: -2147483648..2147483647.
    EthernetTag interface{}

    // Originating IP. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    OriginatingIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Originating IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    OriginatingIpXr interface{}

    // IP of nexthop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Local entry. The type is bool.
    IsLocalEntry interface{}

    // Proxy entry. The type is bool.
    IsProxyEntry interface{}

    // Encap type of local or remote IMCAST route. The type is interface{} with
    // range: 0..255.
    EncapType interface{}
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

    inclusiveMulticast.EntityData.Children = make(map[string]types.YChild)
    inclusiveMulticast.EntityData.Leafs = make(map[string]types.YLeaf)
    inclusiveMulticast.EntityData.Leafs["evi"] = types.YLeaf{"Evi", inclusiveMulticast.Evi}
    inclusiveMulticast.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", inclusiveMulticast.EthernetTag}
    inclusiveMulticast.EntityData.Leafs["originating-ip"] = types.YLeaf{"OriginatingIp", inclusiveMulticast.OriginatingIp}
    inclusiveMulticast.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", inclusiveMulticast.EviXr}
    inclusiveMulticast.EntityData.Leafs["ethernet-tag-xr"] = types.YLeaf{"EthernetTagXr", inclusiveMulticast.EthernetTagXr}
    inclusiveMulticast.EntityData.Leafs["originating-ip-xr"] = types.YLeaf{"OriginatingIpXr", inclusiveMulticast.OriginatingIpXr}
    inclusiveMulticast.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", inclusiveMulticast.NextHop}
    inclusiveMulticast.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", inclusiveMulticast.OutputLabel}
    inclusiveMulticast.EntityData.Leafs["is-local-entry"] = types.YLeaf{"IsLocalEntry", inclusiveMulticast.IsLocalEntry}
    inclusiveMulticast.EntityData.Leafs["is-proxy-entry"] = types.YLeaf{"IsProxyEntry", inclusiveMulticast.IsProxyEntry}
    inclusiveMulticast.EntityData.Leafs["encap-type"] = types.YLeaf{"EncapType", inclusiveMulticast.EncapType}
    return &(inclusiveMulticast.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets
// L2VPN EVPN EVI RT Child Table
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI RT Table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget.
    RouteTarget []Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget
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

    routeTargets.EntityData.Children = make(map[string]types.YChild)
    routeTargets.EntityData.Children["route-target"] = types.YChild{"RouteTarget", nil}
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children[types.GetSegmentPath(&routeTargets.RouteTarget[i])] = types.YChild{"RouteTarget", &routeTargets.RouteTarget[i]}
    }
    routeTargets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routeTargets.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget
// L2VPN EVPN EVI RT Table
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Role of the route target. The type is BgpRouteTargetRole.
    Role interface{}

    // Type of the route target. The type is BgpRouteTarget.
    Type_ interface{}

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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Bridge Domain Name. The type is string.
    BdName interface{}

    // VPN ID. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // RT Role. The type is L2vpnAdRtRole.
    RouteTargetRole interface{}

    // RT Stitching. The type is bool.
    RouteTargetStitching interface{}

    // Route Target.
    RouteTarget Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_
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

    routeTarget.EntityData.Children = make(map[string]types.YChild)
    routeTarget.EntityData.Children["route-target"] = types.YChild{"RouteTarget", &routeTarget.RouteTarget}
    routeTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTarget.EntityData.Leafs["evi"] = types.YLeaf{"Evi", routeTarget.Evi}
    routeTarget.EntityData.Leafs["role"] = types.YLeaf{"Role", routeTarget.Role}
    routeTarget.EntityData.Leafs["type"] = types.YLeaf{"Type_", routeTarget.Type_}
    routeTarget.EntityData.Leafs["format"] = types.YLeaf{"Format", routeTarget.Format}
    routeTarget.EntityData.Leafs["as"] = types.YLeaf{"As", routeTarget.As}
    routeTarget.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", routeTarget.AsIndex}
    routeTarget.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", routeTarget.AddrIndex}
    routeTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", routeTarget.Address}
    routeTarget.EntityData.Leafs["bd-name"] = types.YLeaf{"BdName", routeTarget.BdName}
    routeTarget.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", routeTarget.EviXr}
    routeTarget.EntityData.Leafs["route-target-role"] = types.YLeaf{"RouteTargetRole", routeTarget.RouteTargetRole}
    routeTarget.EntityData.Leafs["route-target-stitching"] = types.YLeaf{"RouteTargetStitching", routeTarget.RouteTargetStitching}
    return &(routeTarget.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_
// Route Target
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs

    // four byte as.
    FourByteAs Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs

    // v4 addr.
    V4Addr Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr

    // es import.
    EsImport Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport
}

func (routeTarget_ *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_) GetEntityData() *types.CommonEntityData {
    routeTarget_.EntityData.YFilter = routeTarget_.YFilter
    routeTarget_.EntityData.YangName = "route-target"
    routeTarget_.EntityData.BundleName = "cisco_ios_xr"
    routeTarget_.EntityData.ParentYangName = "route-target"
    routeTarget_.EntityData.SegmentPath = "route-target"
    routeTarget_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget_.EntityData.Children = make(map[string]types.YChild)
    routeTarget_.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &routeTarget_.TwoByteAs}
    routeTarget_.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &routeTarget_.FourByteAs}
    routeTarget_.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &routeTarget_.V4Addr}
    routeTarget_.EntityData.Children["es-import"] = types.YChild{"EsImport", &routeTarget_.EsImport}
    routeTarget_.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTarget_.EntityData.Leafs["rt"] = types.YLeaf{"Rt", routeTarget_.Rt}
    return &(routeTarget_.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "route-target"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
    return &(twoByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "route-target"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "route-target"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
    return &(v4Addr.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport
// es import
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "route-target"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = make(map[string]types.YChild)
    esImport.EntityData.Leafs = make(map[string]types.YLeaf)
    esImport.EntityData.Leafs["high-bytes"] = types.YLeaf{"HighBytes", esImport.HighBytes}
    esImport.EntityData.Leafs["low-bytes"] = types.YLeaf{"LowBytes", esImport.LowBytes}
    return &(esImport.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs
// L2VPN EVPN EVI MAC table
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN MAC table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac.
    Mac []Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac
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

    macs.EntityData.Children = make(map[string]types.YChild)
    macs.EntityData.Children["mac"] = types.YChild{"Mac", nil}
    for i := range macs.Mac {
        macs.EntityData.Children[types.GetSegmentPath(&macs.Mac[i])] = types.YChild{"Mac", &macs.Mac[i]}
    }
    macs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macs.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac
// L2VPN EVPN MAC table
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // IP Address. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddressXr interface{}

    // IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SooNexthop interface{}

    // IP nexthop address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

    // Local Ethernet Segment id. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier.
    LocalEthernetSegmentIdentifier []Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier

    // Remote Ethernet Segment id. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier.
    RemoteEthernetSegmentIdentifier []Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer.
    PathBuffer []Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer
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

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["local-ethernet-segment-identifier"] = types.YChild{"LocalEthernetSegmentIdentifier", nil}
    for i := range mac.LocalEthernetSegmentIdentifier {
        mac.EntityData.Children[types.GetSegmentPath(&mac.LocalEthernetSegmentIdentifier[i])] = types.YChild{"LocalEthernetSegmentIdentifier", &mac.LocalEthernetSegmentIdentifier[i]}
    }
    mac.EntityData.Children["remote-ethernet-segment-identifier"] = types.YChild{"RemoteEthernetSegmentIdentifier", nil}
    for i := range mac.RemoteEthernetSegmentIdentifier {
        mac.EntityData.Children[types.GetSegmentPath(&mac.RemoteEthernetSegmentIdentifier[i])] = types.YChild{"RemoteEthernetSegmentIdentifier", &mac.RemoteEthernetSegmentIdentifier[i]}
    }
    mac.EntityData.Children["path-buffer"] = types.YChild{"PathBuffer", nil}
    for i := range mac.PathBuffer {
        mac.EntityData.Children[types.GetSegmentPath(&mac.PathBuffer[i])] = types.YChild{"PathBuffer", &mac.PathBuffer[i]}
    }
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    mac.EntityData.Leafs["evi"] = types.YLeaf{"Evi", mac.Evi}
    mac.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", mac.EthernetTag}
    mac.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", mac.MacAddress}
    mac.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", mac.IpAddress}
    mac.EntityData.Leafs["ethernet-tag-xr"] = types.YLeaf{"EthernetTagXr", mac.EthernetTagXr}
    mac.EntityData.Leafs["mac-address-xr"] = types.YLeaf{"MacAddressXr", mac.MacAddressXr}
    mac.EntityData.Leafs["ip-address-xr"] = types.YLeaf{"IpAddressXr", mac.IpAddressXr}
    mac.EntityData.Leafs["local-label"] = types.YLeaf{"LocalLabel", mac.LocalLabel}
    mac.EntityData.Leafs["num-paths"] = types.YLeaf{"NumPaths", mac.NumPaths}
    mac.EntityData.Leafs["is-local-mac"] = types.YLeaf{"IsLocalMac", mac.IsLocalMac}
    mac.EntityData.Leafs["is-proxy-entry"] = types.YLeaf{"IsProxyEntry", mac.IsProxyEntry}
    mac.EntityData.Leafs["is-remote-mac"] = types.YLeaf{"IsRemoteMac", mac.IsRemoteMac}
    mac.EntityData.Leafs["soo-nexthop"] = types.YLeaf{"SooNexthop", mac.SooNexthop}
    mac.EntityData.Leafs["ipnh-address"] = types.YLeaf{"IpnhAddress", mac.IpnhAddress}
    mac.EntityData.Leafs["esi-port-key"] = types.YLeaf{"EsiPortKey", mac.EsiPortKey}
    mac.EntityData.Leafs["local-encap-type"] = types.YLeaf{"LocalEncapType", mac.LocalEncapType}
    mac.EntityData.Leafs["remote-encap-type"] = types.YLeaf{"RemoteEncapType", mac.RemoteEncapType}
    mac.EntityData.Leafs["learned-bridge-port-name"] = types.YLeaf{"LearnedBridgePortName", mac.LearnedBridgePortName}
    mac.EntityData.Leafs["local-seq-id"] = types.YLeaf{"LocalSeqId", mac.LocalSeqId}
    mac.EntityData.Leafs["remote-seq-id"] = types.YLeaf{"RemoteSeqId", mac.RemoteSeqId}
    mac.EntityData.Leafs["local-l3-label"] = types.YLeaf{"LocalL3Label", mac.LocalL3Label}
    mac.EntityData.Leafs["router-mac-address"] = types.YLeaf{"RouterMacAddress", mac.RouterMacAddress}
    mac.EntityData.Leafs["mac-flush-requested"] = types.YLeaf{"MacFlushRequested", mac.MacFlushRequested}
    mac.EntityData.Leafs["mac-flush-received"] = types.YLeaf{"MacFlushReceived", mac.MacFlushReceived}
    mac.EntityData.Leafs["internal-label"] = types.YLeaf{"InternalLabel", mac.InternalLabel}
    mac.EntityData.Leafs["resolved"] = types.YLeaf{"Resolved", mac.Resolved}
    mac.EntityData.Leafs["local-is-static"] = types.YLeaf{"LocalIsStatic", mac.LocalIsStatic}
    mac.EntityData.Leafs["remote-is-static"] = types.YLeaf{"RemoteIsStatic", mac.RemoteIsStatic}
    return &(mac.EntityData)
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

    localEthernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    localEthernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    localEthernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", localEthernetSegmentIdentifier.Entry}
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

    remoteEthernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    remoteEthernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteEthernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", remoteEthernetSegmentIdentifier.Entry}
    return &(remoteEthernetSegmentIdentifier.EntityData)
}

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer
// Path List Buffer
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    pathBuffer.EntityData.Children = make(map[string]types.YChild)
    pathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    pathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", pathBuffer.NextHop}
    pathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", pathBuffer.OutputLabel}
    pathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel}
    return &(pathBuffer.EntityData)
}

// Evpn_Nodes_Node_InternalLabels
// EVPN Internal Label Table
type Evpn_Nodes_Node_InternalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN Internal Label. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel.
    InternalLabel []Evpn_Nodes_Node_InternalLabels_InternalLabel
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

    internalLabels.EntityData.Children = make(map[string]types.YChild)
    internalLabels.EntityData.Children["internal-label"] = types.YChild{"InternalLabel", nil}
    for i := range internalLabels.InternalLabel {
        internalLabels.EntityData.Children[types.GetSegmentPath(&internalLabels.InternalLabel[i])] = types.YChild{"InternalLabel", &internalLabels.InternalLabel[i]}
    }
    internalLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(internalLabels.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel
// L2VPN EVPN Internal Label
type Evpn_Nodes_Node_InternalLabels_InternalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // ES id (part 1/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // EVPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Ethernet Segment id. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Esi interface{}

    // Label Tag. The type is interface{} with range: 0..4294967295.
    Tag interface{}

    // MPLS Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Encap type of remote EAD/ES, EAD/EVI and MAC routes. The type is
    // interface{} with range: 0..255.
    Encap interface{}

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

    // MAC Path list buffer. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel_MacPathBuffer.
    MacPathBuffer []Evpn_Nodes_Node_InternalLabels_InternalLabel_MacPathBuffer

    // EAD/ES Path list buffer. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel_EadPathBuffer.
    EadPathBuffer []Evpn_Nodes_Node_InternalLabels_InternalLabel_EadPathBuffer

    // EAD/EVI Path list buffer. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel_EviPathBuffer.
    EviPathBuffer []Evpn_Nodes_Node_InternalLabels_InternalLabel_EviPathBuffer

    // Summary Path list buffer. The type is slice of
    // Evpn_Nodes_Node_InternalLabels_InternalLabel_SummaryPathBuffer.
    SummaryPathBuffer []Evpn_Nodes_Node_InternalLabels_InternalLabel_SummaryPathBuffer
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

    internalLabel.EntityData.Children = make(map[string]types.YChild)
    internalLabel.EntityData.Children["mac-path-buffer"] = types.YChild{"MacPathBuffer", nil}
    for i := range internalLabel.MacPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.MacPathBuffer[i])] = types.YChild{"MacPathBuffer", &internalLabel.MacPathBuffer[i]}
    }
    internalLabel.EntityData.Children["ead-path-buffer"] = types.YChild{"EadPathBuffer", nil}
    for i := range internalLabel.EadPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.EadPathBuffer[i])] = types.YChild{"EadPathBuffer", &internalLabel.EadPathBuffer[i]}
    }
    internalLabel.EntityData.Children["evi-path-buffer"] = types.YChild{"EviPathBuffer", nil}
    for i := range internalLabel.EviPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.EviPathBuffer[i])] = types.YChild{"EviPathBuffer", &internalLabel.EviPathBuffer[i]}
    }
    internalLabel.EntityData.Children["summary-path-buffer"] = types.YChild{"SummaryPathBuffer", nil}
    for i := range internalLabel.SummaryPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.SummaryPathBuffer[i])] = types.YChild{"SummaryPathBuffer", &internalLabel.SummaryPathBuffer[i]}
    }
    internalLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    internalLabel.EntityData.Leafs["evi"] = types.YLeaf{"Evi", internalLabel.Evi}
    internalLabel.EntityData.Leafs["esi1"] = types.YLeaf{"Esi1", internalLabel.Esi1}
    internalLabel.EntityData.Leafs["esi2"] = types.YLeaf{"Esi2", internalLabel.Esi2}
    internalLabel.EntityData.Leafs["esi3"] = types.YLeaf{"Esi3", internalLabel.Esi3}
    internalLabel.EntityData.Leafs["esi4"] = types.YLeaf{"Esi4", internalLabel.Esi4}
    internalLabel.EntityData.Leafs["esi5"] = types.YLeaf{"Esi5", internalLabel.Esi5}
    internalLabel.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", internalLabel.EthernetTag}
    internalLabel.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", internalLabel.EviXr}
    internalLabel.EntityData.Leafs["esi"] = types.YLeaf{"Esi", internalLabel.Esi}
    internalLabel.EntityData.Leafs["tag"] = types.YLeaf{"Tag", internalLabel.Tag}
    internalLabel.EntityData.Leafs["internal-label"] = types.YLeaf{"InternalLabel", internalLabel.InternalLabel}
    internalLabel.EntityData.Leafs["encap"] = types.YLeaf{"Encap", internalLabel.Encap}
    internalLabel.EntityData.Leafs["mac-num-paths"] = types.YLeaf{"MacNumPaths", internalLabel.MacNumPaths}
    internalLabel.EntityData.Leafs["ead-num-paths"] = types.YLeaf{"EadNumPaths", internalLabel.EadNumPaths}
    internalLabel.EntityData.Leafs["evi-num-paths"] = types.YLeaf{"EviNumPaths", internalLabel.EviNumPaths}
    internalLabel.EntityData.Leafs["sum-num-paths"] = types.YLeaf{"SumNumPaths", internalLabel.SumNumPaths}
    internalLabel.EntityData.Leafs["sum-num-active-paths"] = types.YLeaf{"SumNumActivePaths", internalLabel.SumNumActivePaths}
    internalLabel.EntityData.Leafs["resolved"] = types.YLeaf{"Resolved", internalLabel.Resolved}
    internalLabel.EntityData.Leafs["ecmp-disable"] = types.YLeaf{"EcmpDisable", internalLabel.EcmpDisable}
    internalLabel.EntityData.Leafs["redundancy-single-active"] = types.YLeaf{"RedundancySingleActive", internalLabel.RedundancySingleActive}
    internalLabel.EntityData.Leafs["redundancy-single-flow-active"] = types.YLeaf{"RedundancySingleFlowActive", internalLabel.RedundancySingleFlowActive}
    return &(internalLabel.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel_MacPathBuffer
// MAC Path list buffer
type Evpn_Nodes_Node_InternalLabels_InternalLabel_MacPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    macPathBuffer.EntityData.Children = make(map[string]types.YChild)
    macPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    macPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", macPathBuffer.NextHop}
    macPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", macPathBuffer.OutputLabel}
    macPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", macPathBuffer.SrteTunnel}
    return &(macPathBuffer.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel_EadPathBuffer
// EAD/ES Path list buffer
type Evpn_Nodes_Node_InternalLabels_InternalLabel_EadPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    eadPathBuffer.EntityData.Children = make(map[string]types.YChild)
    eadPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    eadPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", eadPathBuffer.NextHop}
    eadPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", eadPathBuffer.OutputLabel}
    eadPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", eadPathBuffer.SrteTunnel}
    return &(eadPathBuffer.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel_EviPathBuffer
// EAD/EVI Path list buffer
type Evpn_Nodes_Node_InternalLabels_InternalLabel_EviPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    eviPathBuffer.EntityData.Children = make(map[string]types.YChild)
    eviPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    eviPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", eviPathBuffer.NextHop}
    eviPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", eviPathBuffer.OutputLabel}
    eviPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", eviPathBuffer.SrteTunnel}
    return &(eviPathBuffer.EntityData)
}

// Evpn_Nodes_Node_InternalLabels_InternalLabel_SummaryPathBuffer
// Summary Path list buffer
type Evpn_Nodes_Node_InternalLabels_InternalLabel_SummaryPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    summaryPathBuffer.EntityData.Children = make(map[string]types.YChild)
    summaryPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    summaryPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", summaryPathBuffer.NextHop}
    summaryPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", summaryPathBuffer.OutputLabel}
    summaryPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", summaryPathBuffer.SrteTunnel}
    return &(summaryPathBuffer.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments
// EVPN Ethernet-Segment Table
type Evpn_Nodes_Node_EthernetSegments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet-Segment Entry. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment.
    EthernetSegment []Evpn_Nodes_Node_EthernetSegments_EthernetSegment
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

    ethernetSegments.EntityData.Children = make(map[string]types.YChild)
    ethernetSegments.EntityData.Children["ethernet-segment"] = types.YChild{"EthernetSegment", nil}
    for i := range ethernetSegments.EthernetSegment {
        ethernetSegments.EntityData.Children[types.GetSegmentPath(&ethernetSegments.EthernetSegment[i])] = types.YChild{"EthernetSegment", &ethernetSegments.EthernetSegment[i]}
    }
    ethernetSegments.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetSegments.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment
// EVPN Ethernet-Segment Entry
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // ES id (part 1/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
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

    // Main port ifhandle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    IfHandle interface{}

    // Main port redundancy group role. The type is L2vpnRgRole.
    MainPortRole interface{}

    // Main Port MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MainPortMac interface{}

    // Number of PWs in Up state. The type is interface{} with range:
    // 0..4294967295.
    NumUpPWs interface{}

    // ES-Import Route Target. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    RouteTarget interface{}

    // Origin of operational ES-Import RT. The type is L2vpnEvpnRtOrigin.
    RtOrigin interface{}

    // ES BGP Gates. The type is string.
    EsBgpGates interface{}

    // ES L2FIB Gates. The type is string.
    EsL2FibGates interface{}

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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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
    EthernetSegmentIdentifier []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier

    // List of Primary services ESI/I-SIDs. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService.
    PrimaryService []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService

    // List of Secondary services ESI/I-SIDs. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService.
    SecondaryService []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService

    // Elected ISID service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult.
    ServiceCarvingISidelectedResult []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult

    // Not elected ISID service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult.
    ServiceCarvingIsidNotElectedResult []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult

    // Elected EVI service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult.
    ServiceCarvingEviElectedResult []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult

    // Not elected EVI service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult.
    ServiceCarvingEviNotElectedResult []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult

    // List of nexthop IPv6 addresses. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop.
    NextHop []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop

    // Permanent EVPN VPWS service carving results. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult.
    ServiceCarvingVpwsPermanentResult []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult

    // Remote split horizon group labels. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel.
    RemoteSplitHorizonGroupLabel []Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
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

    ethernetSegment.EntityData.Children = make(map[string]types.YChild)
    ethernetSegment.EntityData.Children["ethernet-segment-identifier"] = types.YChild{"EthernetSegmentIdentifier", nil}
    for i := range ethernetSegment.EthernetSegmentIdentifier {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.EthernetSegmentIdentifier[i])] = types.YChild{"EthernetSegmentIdentifier", &ethernetSegment.EthernetSegmentIdentifier[i]}
    }
    ethernetSegment.EntityData.Children["primary-service"] = types.YChild{"PrimaryService", nil}
    for i := range ethernetSegment.PrimaryService {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.PrimaryService[i])] = types.YChild{"PrimaryService", &ethernetSegment.PrimaryService[i]}
    }
    ethernetSegment.EntityData.Children["secondary-service"] = types.YChild{"SecondaryService", nil}
    for i := range ethernetSegment.SecondaryService {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.SecondaryService[i])] = types.YChild{"SecondaryService", &ethernetSegment.SecondaryService[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-i-sidelected-result"] = types.YChild{"ServiceCarvingISidelectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingISidelectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingISidelectedResult[i])] = types.YChild{"ServiceCarvingISidelectedResult", &ethernetSegment.ServiceCarvingISidelectedResult[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-isid-not-elected-result"] = types.YChild{"ServiceCarvingIsidNotElectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingIsidNotElectedResult[i])] = types.YChild{"ServiceCarvingIsidNotElectedResult", &ethernetSegment.ServiceCarvingIsidNotElectedResult[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-evi-elected-result"] = types.YChild{"ServiceCarvingEviElectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingEviElectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingEviElectedResult[i])] = types.YChild{"ServiceCarvingEviElectedResult", &ethernetSegment.ServiceCarvingEviElectedResult[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-evi-not-elected-result"] = types.YChild{"ServiceCarvingEviNotElectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingEviNotElectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingEviNotElectedResult[i])] = types.YChild{"ServiceCarvingEviNotElectedResult", &ethernetSegment.ServiceCarvingEviNotElectedResult[i]}
    }
    ethernetSegment.EntityData.Children["next-hop"] = types.YChild{"NextHop", nil}
    for i := range ethernetSegment.NextHop {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.NextHop[i])] = types.YChild{"NextHop", &ethernetSegment.NextHop[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-vpws-permanent-result"] = types.YChild{"ServiceCarvingVpwsPermanentResult", nil}
    for i := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingVpwsPermanentResult[i])] = types.YChild{"ServiceCarvingVpwsPermanentResult", &ethernetSegment.ServiceCarvingVpwsPermanentResult[i]}
    }
    ethernetSegment.EntityData.Children["remote-split-horizon-group-label"] = types.YChild{"RemoteSplitHorizonGroupLabel", nil}
    for i := range ethernetSegment.RemoteSplitHorizonGroupLabel {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.RemoteSplitHorizonGroupLabel[i])] = types.YChild{"RemoteSplitHorizonGroupLabel", &ethernetSegment.RemoteSplitHorizonGroupLabel[i]}
    }
    ethernetSegment.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegment.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", ethernetSegment.InterfaceName}
    ethernetSegment.EntityData.Leafs["esi1"] = types.YLeaf{"Esi1", ethernetSegment.Esi1}
    ethernetSegment.EntityData.Leafs["esi2"] = types.YLeaf{"Esi2", ethernetSegment.Esi2}
    ethernetSegment.EntityData.Leafs["esi3"] = types.YLeaf{"Esi3", ethernetSegment.Esi3}
    ethernetSegment.EntityData.Leafs["esi4"] = types.YLeaf{"Esi4", ethernetSegment.Esi4}
    ethernetSegment.EntityData.Leafs["esi5"] = types.YLeaf{"Esi5", ethernetSegment.Esi5}
    ethernetSegment.EntityData.Leafs["esi-type"] = types.YLeaf{"EsiType", ethernetSegment.EsiType}
    ethernetSegment.EntityData.Leafs["esi-system-identifier"] = types.YLeaf{"EsiSystemIdentifier", ethernetSegment.EsiSystemIdentifier}
    ethernetSegment.EntityData.Leafs["esi-port-key"] = types.YLeaf{"EsiPortKey", ethernetSegment.EsiPortKey}
    ethernetSegment.EntityData.Leafs["esi-system-priority"] = types.YLeaf{"EsiSystemPriority", ethernetSegment.EsiSystemPriority}
    ethernetSegment.EntityData.Leafs["ethernet-segment-name"] = types.YLeaf{"EthernetSegmentName", ethernetSegment.EthernetSegmentName}
    ethernetSegment.EntityData.Leafs["ethernet-segment-state"] = types.YLeaf{"EthernetSegmentState", ethernetSegment.EthernetSegmentState}
    ethernetSegment.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", ethernetSegment.IfHandle}
    ethernetSegment.EntityData.Leafs["main-port-role"] = types.YLeaf{"MainPortRole", ethernetSegment.MainPortRole}
    ethernetSegment.EntityData.Leafs["main-port-mac"] = types.YLeaf{"MainPortMac", ethernetSegment.MainPortMac}
    ethernetSegment.EntityData.Leafs["num-up-p-ws"] = types.YLeaf{"NumUpPWs", ethernetSegment.NumUpPWs}
    ethernetSegment.EntityData.Leafs["route-target"] = types.YLeaf{"RouteTarget", ethernetSegment.RouteTarget}
    ethernetSegment.EntityData.Leafs["rt-origin"] = types.YLeaf{"RtOrigin", ethernetSegment.RtOrigin}
    ethernetSegment.EntityData.Leafs["es-bgp-gates"] = types.YLeaf{"EsBgpGates", ethernetSegment.EsBgpGates}
    ethernetSegment.EntityData.Leafs["es-l2fib-gates"] = types.YLeaf{"EsL2FibGates", ethernetSegment.EsL2FibGates}
    ethernetSegment.EntityData.Leafs["mac-flushing-mode-config"] = types.YLeaf{"MacFlushingModeConfig", ethernetSegment.MacFlushingModeConfig}
    ethernetSegment.EntityData.Leafs["load-balance-mode-config"] = types.YLeaf{"LoadBalanceModeConfig", ethernetSegment.LoadBalanceModeConfig}
    ethernetSegment.EntityData.Leafs["load-balance-mode-is-default"] = types.YLeaf{"LoadBalanceModeIsDefault", ethernetSegment.LoadBalanceModeIsDefault}
    ethernetSegment.EntityData.Leafs["load-balance-mode-oper"] = types.YLeaf{"LoadBalanceModeOper", ethernetSegment.LoadBalanceModeOper}
    ethernetSegment.EntityData.Leafs["force-single-home"] = types.YLeaf{"ForceSingleHome", ethernetSegment.ForceSingleHome}
    ethernetSegment.EntityData.Leafs["source-mac-oper"] = types.YLeaf{"SourceMacOper", ethernetSegment.SourceMacOper}
    ethernetSegment.EntityData.Leafs["source-mac-origin"] = types.YLeaf{"SourceMacOrigin", ethernetSegment.SourceMacOrigin}
    ethernetSegment.EntityData.Leafs["peering-timer"] = types.YLeaf{"PeeringTimer", ethernetSegment.PeeringTimer}
    ethernetSegment.EntityData.Leafs["peering-timer-left"] = types.YLeaf{"PeeringTimerLeft", ethernetSegment.PeeringTimerLeft}
    ethernetSegment.EntityData.Leafs["recovery-timer"] = types.YLeaf{"RecoveryTimer", ethernetSegment.RecoveryTimer}
    ethernetSegment.EntityData.Leafs["recovery-timer-left"] = types.YLeaf{"RecoveryTimerLeft", ethernetSegment.RecoveryTimerLeft}
    ethernetSegment.EntityData.Leafs["carving-timer"] = types.YLeaf{"CarvingTimer", ethernetSegment.CarvingTimer}
    ethernetSegment.EntityData.Leafs["carving-timer-left"] = types.YLeaf{"CarvingTimerLeft", ethernetSegment.CarvingTimerLeft}
    ethernetSegment.EntityData.Leafs["service-carving-mode"] = types.YLeaf{"ServiceCarvingMode", ethernetSegment.ServiceCarvingMode}
    ethernetSegment.EntityData.Leafs["primary-services-input"] = types.YLeaf{"PrimaryServicesInput", ethernetSegment.PrimaryServicesInput}
    ethernetSegment.EntityData.Leafs["secondary-services-input"] = types.YLeaf{"SecondaryServicesInput", ethernetSegment.SecondaryServicesInput}
    ethernetSegment.EntityData.Leafs["forwarder-ports"] = types.YLeaf{"ForwarderPorts", ethernetSegment.ForwarderPorts}
    ethernetSegment.EntityData.Leafs["permanent-forwarder-ports"] = types.YLeaf{"PermanentForwarderPorts", ethernetSegment.PermanentForwarderPorts}
    ethernetSegment.EntityData.Leafs["elected-forwarder-ports"] = types.YLeaf{"ElectedForwarderPorts", ethernetSegment.ElectedForwarderPorts}
    ethernetSegment.EntityData.Leafs["not-elected-forwarder-ports"] = types.YLeaf{"NotElectedForwarderPorts", ethernetSegment.NotElectedForwarderPorts}
    ethernetSegment.EntityData.Leafs["not-config-forwarder-ports"] = types.YLeaf{"NotConfigForwarderPorts", ethernetSegment.NotConfigForwarderPorts}
    ethernetSegment.EntityData.Leafs["mp-protected"] = types.YLeaf{"MpProtected", ethernetSegment.MpProtected}
    ethernetSegment.EntityData.Leafs["nve-anycast-vtep"] = types.YLeaf{"NveAnycastVtep", ethernetSegment.NveAnycastVtep}
    ethernetSegment.EntityData.Leafs["nve-ingress-replication"] = types.YLeaf{"NveIngressReplication", ethernetSegment.NveIngressReplication}
    ethernetSegment.EntityData.Leafs["local-split-horizon-group-label-valid"] = types.YLeaf{"LocalSplitHorizonGroupLabelValid", ethernetSegment.LocalSplitHorizonGroupLabelValid}
    ethernetSegment.EntityData.Leafs["local-split-horizon-group-label"] = types.YLeaf{"LocalSplitHorizonGroupLabel", ethernetSegment.LocalSplitHorizonGroupLabel}
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

    ethernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    ethernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry}
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

    primaryService.EntityData.Children = make(map[string]types.YChild)
    primaryService.EntityData.Leafs = make(map[string]types.YLeaf)
    primaryService.EntityData.Leafs["entry"] = types.YLeaf{"Entry", primaryService.Entry}
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

    secondaryService.EntityData.Children = make(map[string]types.YChild)
    secondaryService.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryService.EntityData.Leafs["entry"] = types.YLeaf{"Entry", secondaryService.Entry}
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

    serviceCarvingISidelectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingISidelectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingISidelectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingISidelectedResult.Entry}
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

    serviceCarvingIsidNotElectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingIsidNotElectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingIsidNotElectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingIsidNotElectedResult.Entry}
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

    serviceCarvingEviElectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingEviElectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingEviElectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingEviElectedResult.Entry}
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

    serviceCarvingEviNotElectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingEviNotElectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingEviNotElectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingEviNotElectedResult.Entry}
    return &(serviceCarvingEviNotElectedResult.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop
// List of nexthop IPv6 addresses
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}
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

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", nextHop.NextHop}
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
    Type_ interface{}

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

    serviceCarvingVpwsPermanentResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingVpwsPermanentResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingVpwsPermanentResult.EntityData.Leafs["vpn-id"] = types.YLeaf{"VpnId", serviceCarvingVpwsPermanentResult.VpnId}
    serviceCarvingVpwsPermanentResult.EntityData.Leafs["type"] = types.YLeaf{"Type_", serviceCarvingVpwsPermanentResult.Type_}
    serviceCarvingVpwsPermanentResult.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", serviceCarvingVpwsPermanentResult.EthernetTag}
    return &(serviceCarvingVpwsPermanentResult.EntityData)
}

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

    remoteSplitHorizonGroupLabel.EntityData.Children = make(map[string]types.YChild)
    remoteSplitHorizonGroupLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteSplitHorizonGroupLabel.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", remoteSplitHorizonGroupLabel.NextHop}
    remoteSplitHorizonGroupLabel.EntityData.Leafs["label"] = types.YLeaf{"Label", remoteSplitHorizonGroupLabel.Label}
    return &(remoteSplitHorizonGroupLabel.EntityData)
}

// Evpn_Nodes_Node_AcIds
// EVPN AC ID table
type Evpn_Nodes_Node_AcIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN AC ID table. The type is slice of Evpn_Nodes_Node_AcIds_AcId.
    AcId []Evpn_Nodes_Node_AcIds_AcId
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

    acIds.EntityData.Children = make(map[string]types.YChild)
    acIds.EntityData.Children["ac-id"] = types.YChild{"AcId", nil}
    for i := range acIds.AcId {
        acIds.EntityData.Children[types.GetSegmentPath(&acIds.AcId[i])] = types.YChild{"AcId", &acIds.AcId[i]}
    }
    acIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(acIds.EntityData)
}

// Evpn_Nodes_Node_AcIds_AcId
// EVPN AC ID table
type Evpn_Nodes_Node_AcIds_AcId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // AC ID. The type is interface{} with range: -2147483648..2147483647.
    AcId interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}
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

    acId.EntityData.Children = make(map[string]types.YChild)
    acId.EntityData.Leafs = make(map[string]types.YLeaf)
    acId.EntityData.Leafs["evi"] = types.YLeaf{"Evi", acId.Evi}
    acId.EntityData.Leafs["ac-id"] = types.YLeaf{"AcId", acId.AcId}
    acId.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", acId.EviXr}
    acId.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", acId.Neighbor}
    return &(acId.EntityData)
}

// Evpn_Active
// Active EVPN operational data
type Evpn_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Table.
    Evis Evpn_Active_Evis

    // L2VPN EVPN Summary.
    Summary Evpn_Active_Summary

    // L2VPN EVI Detail Table.
    EviDetail Evpn_Active_EviDetail

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

    active.EntityData.Children = make(map[string]types.YChild)
    active.EntityData.Children["evis"] = types.YChild{"Evis", &active.Evis}
    active.EntityData.Children["summary"] = types.YChild{"Summary", &active.Summary}
    active.EntityData.Children["evi-detail"] = types.YChild{"EviDetail", &active.EviDetail}
    active.EntityData.Children["internal-labels"] = types.YChild{"InternalLabels", &active.InternalLabels}
    active.EntityData.Children["ethernet-segments"] = types.YChild{"EthernetSegments", &active.EthernetSegments}
    active.EntityData.Children["ac-ids"] = types.YChild{"AcIds", &active.AcIds}
    active.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(active.EntityData)
}

// Evpn_Active_Evis
// L2VPN EVPN EVI Table
type Evpn_Active_Evis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Entry. The type is slice of Evpn_Active_Evis_Evi.
    Evi []Evpn_Active_Evis_Evi
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

    evis.EntityData.Children = make(map[string]types.YChild)
    evis.EntityData.Children["evi"] = types.YChild{"Evi", nil}
    for i := range evis.Evi {
        evis.EntityData.Children[types.GetSegmentPath(&evis.Evi[i])] = types.YChild{"Evi", &evis.Evi[i]}
    }
    evis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(evis.EntityData)
}

// Evpn_Active_Evis_Evi
// L2VPN EVPN EVI Entry
type Evpn_Active_Evis_Evi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN id. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type_ interface{}
}

func (evi *Evpn_Active_Evis_Evi) GetEntityData() *types.CommonEntityData {
    evi.EntityData.YFilter = evi.YFilter
    evi.EntityData.YangName = "evi"
    evi.EntityData.BundleName = "cisco_ios_xr"
    evi.EntityData.ParentYangName = "evis"
    evi.EntityData.SegmentPath = "evi" + "[evi='" + fmt.Sprintf("%v", evi.Evi) + "']"
    evi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evi.EntityData.Children = make(map[string]types.YChild)
    evi.EntityData.Leafs = make(map[string]types.YLeaf)
    evi.EntityData.Leafs["evi"] = types.YLeaf{"Evi", evi.Evi}
    evi.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", evi.EviXr}
    evi.EntityData.Leafs["bd-name"] = types.YLeaf{"BdName", evi.BdName}
    evi.EntityData.Leafs["type"] = types.YLeaf{"Type_", evi.Type_}
    return &(evi.EntityData)
}

// Evpn_Active_Summary
// L2VPN EVPN Summary
type Evpn_Active_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Router ID. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // BGP AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // Number of EVI DB Entries. The type is interface{} with range:
    // 0..4294967295.
    EvIs interface{}

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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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
    L2RibThrottle interface{}

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

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", summary.RouterId}
    summary.EntityData.Leafs["as"] = types.YLeaf{"As", summary.As}
    summary.EntityData.Leafs["ev-is"] = types.YLeaf{"EvIs", summary.EvIs}
    summary.EntityData.Leafs["local-mac-routes"] = types.YLeaf{"LocalMacRoutes", summary.LocalMacRoutes}
    summary.EntityData.Leafs["local-ipv4-mac-routes"] = types.YLeaf{"LocalIpv4MacRoutes", summary.LocalIpv4MacRoutes}
    summary.EntityData.Leafs["local-ipv6-mac-routes"] = types.YLeaf{"LocalIpv6MacRoutes", summary.LocalIpv6MacRoutes}
    summary.EntityData.Leafs["es-global-mac-routes"] = types.YLeaf{"EsGlobalMacRoutes", summary.EsGlobalMacRoutes}
    summary.EntityData.Leafs["remote-mac-routes"] = types.YLeaf{"RemoteMacRoutes", summary.RemoteMacRoutes}
    summary.EntityData.Leafs["remote-soo-mac-routes"] = types.YLeaf{"RemoteSooMacRoutes", summary.RemoteSooMacRoutes}
    summary.EntityData.Leafs["remote-ipv4-mac-routes"] = types.YLeaf{"RemoteIpv4MacRoutes", summary.RemoteIpv4MacRoutes}
    summary.EntityData.Leafs["remote-ipv6-mac-routes"] = types.YLeaf{"RemoteIpv6MacRoutes", summary.RemoteIpv6MacRoutes}
    summary.EntityData.Leafs["local-imcast-routes"] = types.YLeaf{"LocalImcastRoutes", summary.LocalImcastRoutes}
    summary.EntityData.Leafs["remote-imcast-routes"] = types.YLeaf{"RemoteImcastRoutes", summary.RemoteImcastRoutes}
    summary.EntityData.Leafs["labels"] = types.YLeaf{"Labels", summary.Labels}
    summary.EntityData.Leafs["es-entries"] = types.YLeaf{"EsEntries", summary.EsEntries}
    summary.EntityData.Leafs["neighbor-entries"] = types.YLeaf{"NeighborEntries", summary.NeighborEntries}
    summary.EntityData.Leafs["local-ead-routes"] = types.YLeaf{"LocalEadRoutes", summary.LocalEadRoutes}
    summary.EntityData.Leafs["remote-ead-routes"] = types.YLeaf{"RemoteEadRoutes", summary.RemoteEadRoutes}
    summary.EntityData.Leafs["global-source-mac"] = types.YLeaf{"GlobalSourceMac", summary.GlobalSourceMac}
    summary.EntityData.Leafs["peering-time"] = types.YLeaf{"PeeringTime", summary.PeeringTime}
    summary.EntityData.Leafs["recovery-time"] = types.YLeaf{"RecoveryTime", summary.RecoveryTime}
    summary.EntityData.Leafs["carving-time"] = types.YLeaf{"CarvingTime", summary.CarvingTime}
    summary.EntityData.Leafs["mac-secure-move-count"] = types.YLeaf{"MacSecureMoveCount", summary.MacSecureMoveCount}
    summary.EntityData.Leafs["mac-secure-move-interval"] = types.YLeaf{"MacSecureMoveInterval", summary.MacSecureMoveInterval}
    summary.EntityData.Leafs["mac-secure-freeze-time"] = types.YLeaf{"MacSecureFreezeTime", summary.MacSecureFreezeTime}
    summary.EntityData.Leafs["mac-secure-retry-count"] = types.YLeaf{"MacSecureRetryCount", summary.MacSecureRetryCount}
    summary.EntityData.Leafs["cost-out"] = types.YLeaf{"CostOut", summary.CostOut}
    summary.EntityData.Leafs["startup-cost-in-time"] = types.YLeaf{"StartupCostInTime", summary.StartupCostInTime}
    summary.EntityData.Leafs["l2rib-throttle"] = types.YLeaf{"L2RibThrottle", summary.L2RibThrottle}
    summary.EntityData.Leafs["logging-df-election-enabled"] = types.YLeaf{"LoggingDfElectionEnabled", summary.LoggingDfElectionEnabled}
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

    eviDetail.EntityData.Children = make(map[string]types.YChild)
    eviDetail.EntityData.Children["elements"] = types.YChild{"Elements", &eviDetail.Elements}
    eviDetail.EntityData.Children["evi-children"] = types.YChild{"EviChildren", &eviDetail.EviChildren}
    eviDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eviDetail.EntityData)
}

// Evpn_Active_EviDetail_Elements
// EVI BGP RT Detail Info Elements
type Evpn_Active_EviDetail_Elements struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info. The type is slice of
    // Evpn_Active_EviDetail_Elements_Element.
    Element []Evpn_Active_EviDetail_Elements_Element
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

    elements.EntityData.Children = make(map[string]types.YChild)
    elements.EntityData.Children["element"] = types.YChild{"Element", nil}
    for i := range elements.Element {
        elements.EntityData.Children[types.GetSegmentPath(&elements.Element[i])] = types.YChild{"Element", &elements.Element[i]}
    }
    elements.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(elements.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element
// EVI BGP RT Detail Info
type Evpn_Active_EviDetail_Elements_Element struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN id. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // EVI description. The type is string.
    Description interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type_ interface{}

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

    // EVPN Instance is Regular/Stitching side. The type is interface{} with
    // range: 0..255.
    Stitching interface{}

    // EVPN Instance encapsulation. The type is interface{} with range: 0..255.
    Encapsulation interface{}

    // Flow Label Information.
    FlowLabel Evpn_Active_EviDetail_Elements_Element_FlowLabel

    // Automatic Route Distingtuisher.
    RdAuto Evpn_Active_EviDetail_Elements_Element_RdAuto

    // Configured Route Distinguisher.
    RdConfigured Evpn_Active_EviDetail_Elements_Element_RdConfigured

    // Automatic Route Target.
    RtAuto Evpn_Active_EviDetail_Elements_Element_RtAuto

    // Automatic Route Target Stitching.
    RtAutoStitching Evpn_Active_EviDetail_Elements_Element_RtAutoStitching
}

func (element *Evpn_Active_EviDetail_Elements_Element) GetEntityData() *types.CommonEntityData {
    element.EntityData.YFilter = element.YFilter
    element.EntityData.YangName = "element"
    element.EntityData.BundleName = "cisco_ios_xr"
    element.EntityData.ParentYangName = "elements"
    element.EntityData.SegmentPath = "element" + "[evi='" + fmt.Sprintf("%v", element.Evi) + "']"
    element.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    element.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    element.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    element.EntityData.Children = make(map[string]types.YChild)
    element.EntityData.Children["flow-label"] = types.YChild{"FlowLabel", &element.FlowLabel}
    element.EntityData.Children["rd-auto"] = types.YChild{"RdAuto", &element.RdAuto}
    element.EntityData.Children["rd-configured"] = types.YChild{"RdConfigured", &element.RdConfigured}
    element.EntityData.Children["rt-auto"] = types.YChild{"RtAuto", &element.RtAuto}
    element.EntityData.Children["rt-auto-stitching"] = types.YChild{"RtAutoStitching", &element.RtAutoStitching}
    element.EntityData.Leafs = make(map[string]types.YLeaf)
    element.EntityData.Leafs["evi"] = types.YLeaf{"Evi", element.Evi}
    element.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", element.EviXr}
    element.EntityData.Leafs["description"] = types.YLeaf{"Description", element.Description}
    element.EntityData.Leafs["bd-name"] = types.YLeaf{"BdName", element.BdName}
    element.EntityData.Leafs["type"] = types.YLeaf{"Type_", element.Type_}
    element.EntityData.Leafs["unicast-label"] = types.YLeaf{"UnicastLabel", element.UnicastLabel}
    element.EntityData.Leafs["multicast-label"] = types.YLeaf{"MulticastLabel", element.MulticastLabel}
    element.EntityData.Leafs["cw-disable"] = types.YLeaf{"CwDisable", element.CwDisable}
    element.EntityData.Leafs["table-policy-name"] = types.YLeaf{"TablePolicyName", element.TablePolicyName}
    element.EntityData.Leafs["forward-class"] = types.YLeaf{"ForwardClass", element.ForwardClass}
    element.EntityData.Leafs["rt-import-block-set"] = types.YLeaf{"RtImportBlockSet", element.RtImportBlockSet}
    element.EntityData.Leafs["rt-export-block-set"] = types.YLeaf{"RtExportBlockSet", element.RtExportBlockSet}
    element.EntityData.Leafs["advertise-mac"] = types.YLeaf{"AdvertiseMac", element.AdvertiseMac}
    element.EntityData.Leafs["advertise-bvi-mac"] = types.YLeaf{"AdvertiseBviMac", element.AdvertiseBviMac}
    element.EntityData.Leafs["aliasing-disabled"] = types.YLeaf{"AliasingDisabled", element.AliasingDisabled}
    element.EntityData.Leafs["unknown-unicast-flooding-disabled"] = types.YLeaf{"UnknownUnicastFloodingDisabled", element.UnknownUnicastFloodingDisabled}
    element.EntityData.Leafs["reoriginate-disabled"] = types.YLeaf{"ReoriginateDisabled", element.ReoriginateDisabled}
    element.EntityData.Leafs["stitching"] = types.YLeaf{"Stitching", element.Stitching}
    element.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", element.Encapsulation}
    return &(element.EntityData)
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

    flowLabel.EntityData.Children = make(map[string]types.YChild)
    flowLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    flowLabel.EntityData.Leafs["static-flow-label"] = types.YLeaf{"StaticFlowLabel", flowLabel.StaticFlowLabel}
    flowLabel.EntityData.Leafs["global-flow-label"] = types.YLeaf{"GlobalFlowLabel", flowLabel.GlobalFlowLabel}
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

    rdAuto.EntityData.Children = make(map[string]types.YChild)
    rdAuto.EntityData.Children["auto"] = types.YChild{"Auto", &rdAuto.Auto}
    rdAuto.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rdAuto.TwoByteAs}
    rdAuto.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rdAuto.FourByteAs}
    rdAuto.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rdAuto.V4Addr}
    rdAuto.EntityData.Leafs = make(map[string]types.YLeaf)
    rdAuto.EntityData.Leafs["rd"] = types.YLeaf{"Rd", rdAuto.Rd}
    return &(rdAuto.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto
// auto
type Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    auto.EntityData.Children = make(map[string]types.YChild)
    auto.EntityData.Leafs = make(map[string]types.YLeaf)
    auto.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", auto.RouterId}
    auto.EntityData.Leafs["auto-index"] = types.YLeaf{"AutoIndex", auto.AutoIndex}
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

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
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

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
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

    rdConfigured.EntityData.Children = make(map[string]types.YChild)
    rdConfigured.EntityData.Children["auto"] = types.YChild{"Auto", &rdConfigured.Auto}
    rdConfigured.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rdConfigured.TwoByteAs}
    rdConfigured.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rdConfigured.FourByteAs}
    rdConfigured.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rdConfigured.V4Addr}
    rdConfigured.EntityData.Leafs = make(map[string]types.YLeaf)
    rdConfigured.EntityData.Leafs["rd"] = types.YLeaf{"Rd", rdConfigured.Rd}
    return &(rdConfigured.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto
// auto
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    auto.EntityData.Children = make(map[string]types.YChild)
    auto.EntityData.Leafs = make(map[string]types.YLeaf)
    auto.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", auto.RouterId}
    auto.EntityData.Leafs["auto-index"] = types.YLeaf{"AutoIndex", auto.AutoIndex}
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

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
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

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
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

    rtAuto.EntityData.Children = make(map[string]types.YChild)
    rtAuto.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rtAuto.TwoByteAs}
    rtAuto.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rtAuto.FourByteAs}
    rtAuto.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rtAuto.V4Addr}
    rtAuto.EntityData.Children["es-import"] = types.YChild{"EsImport", &rtAuto.EsImport}
    rtAuto.EntityData.Leafs = make(map[string]types.YLeaf)
    rtAuto.EntityData.Leafs["rt"] = types.YLeaf{"Rt", rtAuto.Rt}
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

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
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

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
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

    esImport.EntityData.Children = make(map[string]types.YChild)
    esImport.EntityData.Leafs = make(map[string]types.YLeaf)
    esImport.EntityData.Leafs["high-bytes"] = types.YLeaf{"HighBytes", esImport.HighBytes}
    esImport.EntityData.Leafs["low-bytes"] = types.YLeaf{"LowBytes", esImport.LowBytes}
    return &(esImport.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching
// Automatic Route Target Stitching
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs

    // v4 addr.
    V4Addr Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr

    // es import.
    EsImport Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport
}

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetEntityData() *types.CommonEntityData {
    rtAutoStitching.EntityData.YFilter = rtAutoStitching.YFilter
    rtAutoStitching.EntityData.YangName = "rt-auto-stitching"
    rtAutoStitching.EntityData.BundleName = "cisco_ios_xr"
    rtAutoStitching.EntityData.ParentYangName = "element"
    rtAutoStitching.EntityData.SegmentPath = "rt-auto-stitching"
    rtAutoStitching.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rtAutoStitching.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rtAutoStitching.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rtAutoStitching.EntityData.Children = make(map[string]types.YChild)
    rtAutoStitching.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rtAutoStitching.TwoByteAs}
    rtAutoStitching.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rtAutoStitching.FourByteAs}
    rtAutoStitching.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rtAutoStitching.V4Addr}
    rtAutoStitching.EntityData.Children["es-import"] = types.YChild{"EsImport", &rtAutoStitching.EsImport}
    rtAutoStitching.EntityData.Leafs = make(map[string]types.YLeaf)
    rtAutoStitching.EntityData.Leafs["rt"] = types.YLeaf{"Rt", rtAutoStitching.Rt}
    return &(rtAutoStitching.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rt-auto-stitching"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
    return &(twoByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs
// four byte as
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rt-auto-stitching"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rt-auto-stitching"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
    return &(v4Addr.EntityData)
}

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport
// es import
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "rt-auto-stitching"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = make(map[string]types.YChild)
    esImport.EntityData.Leafs = make(map[string]types.YLeaf)
    esImport.EntityData.Leafs["high-bytes"] = types.YLeaf{"HighBytes", esImport.HighBytes}
    esImport.EntityData.Leafs["low-bytes"] = types.YLeaf{"LowBytes", esImport.LowBytes}
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

    eviChildren.EntityData.Children = make(map[string]types.YChild)
    eviChildren.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &eviChildren.Neighbors}
    eviChildren.EntityData.Children["ethernet-auto-discoveries"] = types.YChild{"EthernetAutoDiscoveries", &eviChildren.EthernetAutoDiscoveries}
    eviChildren.EntityData.Children["inclusive-multicasts"] = types.YChild{"InclusiveMulticasts", &eviChildren.InclusiveMulticasts}
    eviChildren.EntityData.Children["route-targets"] = types.YChild{"RouteTargets", &eviChildren.RouteTargets}
    eviChildren.EntityData.Children["macs"] = types.YChild{"Macs", &eviChildren.Macs}
    eviChildren.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eviChildren.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Neighbors
// EVPN Neighbor table
type Evpn_Active_EviDetail_EviChildren_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Neighbor table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor.
    Neighbor []Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor
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

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children[types.GetSegmentPath(&neighbors.Neighbor[i])] = types.YChild{"Neighbor", &neighbors.Neighbor[i]}
    }
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor
// EVPN Neighbor table
type Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Neighbor IP. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}
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

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["evi"] = types.YLeaf{"Evi", neighbor.Evi}
    neighbor.EntityData.Leafs["neighbor-ip"] = types.YLeaf{"NeighborIp", neighbor.NeighborIp}
    neighbor.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", neighbor.EviXr}
    neighbor.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", neighbor.Neighbor}
    return &(neighbor.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries
// EVPN Ethernet Auto-Discovery table
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet Auto-Discovery Entry. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery.
    EthernetAutoDiscovery []Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
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

    ethernetAutoDiscoveries.EntityData.Children = make(map[string]types.YChild)
    ethernetAutoDiscoveries.EntityData.Children["ethernet-auto-discovery"] = types.YChild{"EthernetAutoDiscovery", nil}
    for i := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
        ethernetAutoDiscoveries.EntityData.Children[types.GetSegmentPath(&ethernetAutoDiscoveries.EthernetAutoDiscovery[i])] = types.YChild{"EthernetAutoDiscovery", &ethernetAutoDiscoveries.EthernetAutoDiscovery[i]}
    }
    ethernetAutoDiscoveries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetAutoDiscoveries.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
// EVPN Ethernet Auto-Discovery Entry
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // ES id (part 1/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnid interface{}

    // Service Type. The type is L2vpnEvpn.
    Type_ interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Local nexthop IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalNextHop interface{}

    // Associated local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Indication of EthernetAutoDiscovery Route is local. The type is bool.
    IsLocalEad interface{}

    // Encap type of local or remote EAD. The type is interface{} with range:
    // 0..255.
    Encap interface{}

    // Single-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleActive interface{}

    // Single-flow-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleFlowActive interface{}

    // Number of items in path list buffer. The type is interface{} with range:
    // 0..4294967295.
    NumPaths interface{}

    // Ethernet Segment id. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer.
    PathBuffer []Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
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

    ethernetAutoDiscovery.EntityData.Children = make(map[string]types.YChild)
    ethernetAutoDiscovery.EntityData.Children["ethernet-segment-identifier"] = types.YChild{"EthernetSegmentIdentifier", nil}
    for i := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
        ethernetAutoDiscovery.EntityData.Children[types.GetSegmentPath(&ethernetAutoDiscovery.EthernetSegmentIdentifier[i])] = types.YChild{"EthernetSegmentIdentifier", &ethernetAutoDiscovery.EthernetSegmentIdentifier[i]}
    }
    ethernetAutoDiscovery.EntityData.Children["path-buffer"] = types.YChild{"PathBuffer", nil}
    for i := range ethernetAutoDiscovery.PathBuffer {
        ethernetAutoDiscovery.EntityData.Children[types.GetSegmentPath(&ethernetAutoDiscovery.PathBuffer[i])] = types.YChild{"PathBuffer", &ethernetAutoDiscovery.PathBuffer[i]}
    }
    ethernetAutoDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetAutoDiscovery.EntityData.Leafs["evi"] = types.YLeaf{"Evi", ethernetAutoDiscovery.Evi}
    ethernetAutoDiscovery.EntityData.Leafs["esi1"] = types.YLeaf{"Esi1", ethernetAutoDiscovery.Esi1}
    ethernetAutoDiscovery.EntityData.Leafs["esi2"] = types.YLeaf{"Esi2", ethernetAutoDiscovery.Esi2}
    ethernetAutoDiscovery.EntityData.Leafs["esi3"] = types.YLeaf{"Esi3", ethernetAutoDiscovery.Esi3}
    ethernetAutoDiscovery.EntityData.Leafs["esi4"] = types.YLeaf{"Esi4", ethernetAutoDiscovery.Esi4}
    ethernetAutoDiscovery.EntityData.Leafs["esi5"] = types.YLeaf{"Esi5", ethernetAutoDiscovery.Esi5}
    ethernetAutoDiscovery.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", ethernetAutoDiscovery.EthernetTag}
    ethernetAutoDiscovery.EntityData.Leafs["ethernet-vpnid"] = types.YLeaf{"EthernetVpnid", ethernetAutoDiscovery.EthernetVpnid}
    ethernetAutoDiscovery.EntityData.Leafs["type"] = types.YLeaf{"Type_", ethernetAutoDiscovery.Type_}
    ethernetAutoDiscovery.EntityData.Leafs["ethernet-tag-xr"] = types.YLeaf{"EthernetTagXr", ethernetAutoDiscovery.EthernetTagXr}
    ethernetAutoDiscovery.EntityData.Leafs["local-next-hop"] = types.YLeaf{"LocalNextHop", ethernetAutoDiscovery.LocalNextHop}
    ethernetAutoDiscovery.EntityData.Leafs["local-label"] = types.YLeaf{"LocalLabel", ethernetAutoDiscovery.LocalLabel}
    ethernetAutoDiscovery.EntityData.Leafs["is-local-ead"] = types.YLeaf{"IsLocalEad", ethernetAutoDiscovery.IsLocalEad}
    ethernetAutoDiscovery.EntityData.Leafs["encap"] = types.YLeaf{"Encap", ethernetAutoDiscovery.Encap}
    ethernetAutoDiscovery.EntityData.Leafs["redundancy-single-active"] = types.YLeaf{"RedundancySingleActive", ethernetAutoDiscovery.RedundancySingleActive}
    ethernetAutoDiscovery.EntityData.Leafs["redundancy-single-flow-active"] = types.YLeaf{"RedundancySingleFlowActive", ethernetAutoDiscovery.RedundancySingleFlowActive}
    ethernetAutoDiscovery.EntityData.Leafs["num-paths"] = types.YLeaf{"NumPaths", ethernetAutoDiscovery.NumPaths}
    return &(ethernetAutoDiscovery.EntityData)
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

    ethernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    ethernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry}
    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
// Path List Buffer
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    pathBuffer.EntityData.Children = make(map[string]types.YChild)
    pathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    pathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", pathBuffer.NextHop}
    pathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", pathBuffer.OutputLabel}
    pathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel}
    return &(pathBuffer.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts
// L2VPN EVPN IMCAST table
type Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN IMCAST table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast.
    InclusiveMulticast []Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
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

    inclusiveMulticasts.EntityData.Children = make(map[string]types.YChild)
    inclusiveMulticasts.EntityData.Children["inclusive-multicast"] = types.YChild{"InclusiveMulticast", nil}
    for i := range inclusiveMulticasts.InclusiveMulticast {
        inclusiveMulticasts.EntityData.Children[types.GetSegmentPath(&inclusiveMulticasts.InclusiveMulticast[i])] = types.YChild{"InclusiveMulticast", &inclusiveMulticasts.InclusiveMulticast[i]}
    }
    inclusiveMulticasts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inclusiveMulticasts.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
// L2VPN EVPN IMCAST table
type Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag. The type is interface{} with range: -2147483648..2147483647.
    EthernetTag interface{}

    // Originating IP. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    OriginatingIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Originating IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    OriginatingIpXr interface{}

    // IP of nexthop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Local entry. The type is bool.
    IsLocalEntry interface{}

    // Proxy entry. The type is bool.
    IsProxyEntry interface{}

    // Encap type of local or remote IMCAST route. The type is interface{} with
    // range: 0..255.
    EncapType interface{}
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

    inclusiveMulticast.EntityData.Children = make(map[string]types.YChild)
    inclusiveMulticast.EntityData.Leafs = make(map[string]types.YLeaf)
    inclusiveMulticast.EntityData.Leafs["evi"] = types.YLeaf{"Evi", inclusiveMulticast.Evi}
    inclusiveMulticast.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", inclusiveMulticast.EthernetTag}
    inclusiveMulticast.EntityData.Leafs["originating-ip"] = types.YLeaf{"OriginatingIp", inclusiveMulticast.OriginatingIp}
    inclusiveMulticast.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", inclusiveMulticast.EviXr}
    inclusiveMulticast.EntityData.Leafs["ethernet-tag-xr"] = types.YLeaf{"EthernetTagXr", inclusiveMulticast.EthernetTagXr}
    inclusiveMulticast.EntityData.Leafs["originating-ip-xr"] = types.YLeaf{"OriginatingIpXr", inclusiveMulticast.OriginatingIpXr}
    inclusiveMulticast.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", inclusiveMulticast.NextHop}
    inclusiveMulticast.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", inclusiveMulticast.OutputLabel}
    inclusiveMulticast.EntityData.Leafs["is-local-entry"] = types.YLeaf{"IsLocalEntry", inclusiveMulticast.IsLocalEntry}
    inclusiveMulticast.EntityData.Leafs["is-proxy-entry"] = types.YLeaf{"IsProxyEntry", inclusiveMulticast.IsProxyEntry}
    inclusiveMulticast.EntityData.Leafs["encap-type"] = types.YLeaf{"EncapType", inclusiveMulticast.EncapType}
    return &(inclusiveMulticast.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets
// L2VPN EVPN EVI RT Child Table
type Evpn_Active_EviDetail_EviChildren_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI RT Table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget.
    RouteTarget []Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget
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

    routeTargets.EntityData.Children = make(map[string]types.YChild)
    routeTargets.EntityData.Children["route-target"] = types.YChild{"RouteTarget", nil}
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children[types.GetSegmentPath(&routeTargets.RouteTarget[i])] = types.YChild{"RouteTarget", &routeTargets.RouteTarget[i]}
    }
    routeTargets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routeTargets.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget
// L2VPN EVPN EVI RT Table
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Role of the route target. The type is BgpRouteTargetRole.
    Role interface{}

    // Type of the route target. The type is BgpRouteTarget.
    Type_ interface{}

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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Bridge Domain Name. The type is string.
    BdName interface{}

    // VPN ID. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // RT Role. The type is L2vpnAdRtRole.
    RouteTargetRole interface{}

    // RT Stitching. The type is bool.
    RouteTargetStitching interface{}

    // Route Target.
    RouteTarget Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_
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

    routeTarget.EntityData.Children = make(map[string]types.YChild)
    routeTarget.EntityData.Children["route-target"] = types.YChild{"RouteTarget", &routeTarget.RouteTarget}
    routeTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTarget.EntityData.Leafs["evi"] = types.YLeaf{"Evi", routeTarget.Evi}
    routeTarget.EntityData.Leafs["role"] = types.YLeaf{"Role", routeTarget.Role}
    routeTarget.EntityData.Leafs["type"] = types.YLeaf{"Type_", routeTarget.Type_}
    routeTarget.EntityData.Leafs["format"] = types.YLeaf{"Format", routeTarget.Format}
    routeTarget.EntityData.Leafs["as"] = types.YLeaf{"As", routeTarget.As}
    routeTarget.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", routeTarget.AsIndex}
    routeTarget.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", routeTarget.AddrIndex}
    routeTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", routeTarget.Address}
    routeTarget.EntityData.Leafs["bd-name"] = types.YLeaf{"BdName", routeTarget.BdName}
    routeTarget.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", routeTarget.EviXr}
    routeTarget.EntityData.Leafs["route-target-role"] = types.YLeaf{"RouteTargetRole", routeTarget.RouteTargetRole}
    routeTarget.EntityData.Leafs["route-target-stitching"] = types.YLeaf{"RouteTargetStitching", routeTarget.RouteTargetStitching}
    return &(routeTarget.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_
// Route Target
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs

    // four byte as.
    FourByteAs Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs

    // v4 addr.
    V4Addr Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr

    // es import.
    EsImport Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport
}

func (routeTarget_ *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_) GetEntityData() *types.CommonEntityData {
    routeTarget_.EntityData.YFilter = routeTarget_.YFilter
    routeTarget_.EntityData.YangName = "route-target"
    routeTarget_.EntityData.BundleName = "cisco_ios_xr"
    routeTarget_.EntityData.ParentYangName = "route-target"
    routeTarget_.EntityData.SegmentPath = "route-target"
    routeTarget_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget_.EntityData.Children = make(map[string]types.YChild)
    routeTarget_.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &routeTarget_.TwoByteAs}
    routeTarget_.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &routeTarget_.FourByteAs}
    routeTarget_.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &routeTarget_.V4Addr}
    routeTarget_.EntityData.Children["es-import"] = types.YChild{"EsImport", &routeTarget_.EsImport}
    routeTarget_.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTarget_.EntityData.Leafs["rt"] = types.YLeaf{"Rt", routeTarget_.Rt}
    return &(routeTarget_.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs
// two byte as
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "route-target"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
    return &(twoByteAs.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs
// four byte as
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "route-target"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr
// v4 addr
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "route-target"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
    return &(v4Addr.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport
// es import
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "route-target"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = make(map[string]types.YChild)
    esImport.EntityData.Leafs = make(map[string]types.YLeaf)
    esImport.EntityData.Leafs["high-bytes"] = types.YLeaf{"HighBytes", esImport.HighBytes}
    esImport.EntityData.Leafs["low-bytes"] = types.YLeaf{"LowBytes", esImport.LowBytes}
    return &(esImport.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Macs
// L2VPN EVPN EVI MAC table
type Evpn_Active_EviDetail_EviChildren_Macs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN MAC table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Macs_Mac.
    Mac []Evpn_Active_EviDetail_EviChildren_Macs_Mac
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

    macs.EntityData.Children = make(map[string]types.YChild)
    macs.EntityData.Children["mac"] = types.YChild{"Mac", nil}
    for i := range macs.Mac {
        macs.EntityData.Children[types.GetSegmentPath(&macs.Mac[i])] = types.YChild{"Mac", &macs.Mac[i]}
    }
    macs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macs.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Macs_Mac
// L2VPN EVPN MAC table
type Evpn_Active_EviDetail_EviChildren_Macs_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // IP Address. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddressXr interface{}

    // IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SooNexthop interface{}

    // IP nexthop address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

    // Local Ethernet Segment id. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier.
    LocalEthernetSegmentIdentifier []Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier

    // Remote Ethernet Segment id. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier.
    RemoteEthernetSegmentIdentifier []Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer.
    PathBuffer []Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer
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

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["local-ethernet-segment-identifier"] = types.YChild{"LocalEthernetSegmentIdentifier", nil}
    for i := range mac.LocalEthernetSegmentIdentifier {
        mac.EntityData.Children[types.GetSegmentPath(&mac.LocalEthernetSegmentIdentifier[i])] = types.YChild{"LocalEthernetSegmentIdentifier", &mac.LocalEthernetSegmentIdentifier[i]}
    }
    mac.EntityData.Children["remote-ethernet-segment-identifier"] = types.YChild{"RemoteEthernetSegmentIdentifier", nil}
    for i := range mac.RemoteEthernetSegmentIdentifier {
        mac.EntityData.Children[types.GetSegmentPath(&mac.RemoteEthernetSegmentIdentifier[i])] = types.YChild{"RemoteEthernetSegmentIdentifier", &mac.RemoteEthernetSegmentIdentifier[i]}
    }
    mac.EntityData.Children["path-buffer"] = types.YChild{"PathBuffer", nil}
    for i := range mac.PathBuffer {
        mac.EntityData.Children[types.GetSegmentPath(&mac.PathBuffer[i])] = types.YChild{"PathBuffer", &mac.PathBuffer[i]}
    }
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    mac.EntityData.Leafs["evi"] = types.YLeaf{"Evi", mac.Evi}
    mac.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", mac.EthernetTag}
    mac.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", mac.MacAddress}
    mac.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", mac.IpAddress}
    mac.EntityData.Leafs["ethernet-tag-xr"] = types.YLeaf{"EthernetTagXr", mac.EthernetTagXr}
    mac.EntityData.Leafs["mac-address-xr"] = types.YLeaf{"MacAddressXr", mac.MacAddressXr}
    mac.EntityData.Leafs["ip-address-xr"] = types.YLeaf{"IpAddressXr", mac.IpAddressXr}
    mac.EntityData.Leafs["local-label"] = types.YLeaf{"LocalLabel", mac.LocalLabel}
    mac.EntityData.Leafs["num-paths"] = types.YLeaf{"NumPaths", mac.NumPaths}
    mac.EntityData.Leafs["is-local-mac"] = types.YLeaf{"IsLocalMac", mac.IsLocalMac}
    mac.EntityData.Leafs["is-proxy-entry"] = types.YLeaf{"IsProxyEntry", mac.IsProxyEntry}
    mac.EntityData.Leafs["is-remote-mac"] = types.YLeaf{"IsRemoteMac", mac.IsRemoteMac}
    mac.EntityData.Leafs["soo-nexthop"] = types.YLeaf{"SooNexthop", mac.SooNexthop}
    mac.EntityData.Leafs["ipnh-address"] = types.YLeaf{"IpnhAddress", mac.IpnhAddress}
    mac.EntityData.Leafs["esi-port-key"] = types.YLeaf{"EsiPortKey", mac.EsiPortKey}
    mac.EntityData.Leafs["local-encap-type"] = types.YLeaf{"LocalEncapType", mac.LocalEncapType}
    mac.EntityData.Leafs["remote-encap-type"] = types.YLeaf{"RemoteEncapType", mac.RemoteEncapType}
    mac.EntityData.Leafs["learned-bridge-port-name"] = types.YLeaf{"LearnedBridgePortName", mac.LearnedBridgePortName}
    mac.EntityData.Leafs["local-seq-id"] = types.YLeaf{"LocalSeqId", mac.LocalSeqId}
    mac.EntityData.Leafs["remote-seq-id"] = types.YLeaf{"RemoteSeqId", mac.RemoteSeqId}
    mac.EntityData.Leafs["local-l3-label"] = types.YLeaf{"LocalL3Label", mac.LocalL3Label}
    mac.EntityData.Leafs["router-mac-address"] = types.YLeaf{"RouterMacAddress", mac.RouterMacAddress}
    mac.EntityData.Leafs["mac-flush-requested"] = types.YLeaf{"MacFlushRequested", mac.MacFlushRequested}
    mac.EntityData.Leafs["mac-flush-received"] = types.YLeaf{"MacFlushReceived", mac.MacFlushReceived}
    mac.EntityData.Leafs["internal-label"] = types.YLeaf{"InternalLabel", mac.InternalLabel}
    mac.EntityData.Leafs["resolved"] = types.YLeaf{"Resolved", mac.Resolved}
    mac.EntityData.Leafs["local-is-static"] = types.YLeaf{"LocalIsStatic", mac.LocalIsStatic}
    mac.EntityData.Leafs["remote-is-static"] = types.YLeaf{"RemoteIsStatic", mac.RemoteIsStatic}
    return &(mac.EntityData)
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

    localEthernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    localEthernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    localEthernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", localEthernetSegmentIdentifier.Entry}
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

    remoteEthernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    remoteEthernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteEthernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", remoteEthernetSegmentIdentifier.Entry}
    return &(remoteEthernetSegmentIdentifier.EntityData)
}

// Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer
// Path List Buffer
type Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    pathBuffer.EntityData.Children = make(map[string]types.YChild)
    pathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    pathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", pathBuffer.NextHop}
    pathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", pathBuffer.OutputLabel}
    pathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel}
    return &(pathBuffer.EntityData)
}

// Evpn_Active_InternalLabels
// EVPN Internal Label Table
type Evpn_Active_InternalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN Internal Label. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel.
    InternalLabel []Evpn_Active_InternalLabels_InternalLabel
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

    internalLabels.EntityData.Children = make(map[string]types.YChild)
    internalLabels.EntityData.Children["internal-label"] = types.YChild{"InternalLabel", nil}
    for i := range internalLabels.InternalLabel {
        internalLabels.EntityData.Children[types.GetSegmentPath(&internalLabels.InternalLabel[i])] = types.YChild{"InternalLabel", &internalLabels.InternalLabel[i]}
    }
    internalLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(internalLabels.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel
// L2VPN EVPN Internal Label
type Evpn_Active_InternalLabels_InternalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // ES id (part 1/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // EVPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Ethernet Segment id. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Esi interface{}

    // Label Tag. The type is interface{} with range: 0..4294967295.
    Tag interface{}

    // MPLS Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Encap type of remote EAD/ES, EAD/EVI and MAC routes. The type is
    // interface{} with range: 0..255.
    Encap interface{}

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

    // MAC Path list buffer. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel_MacPathBuffer.
    MacPathBuffer []Evpn_Active_InternalLabels_InternalLabel_MacPathBuffer

    // EAD/ES Path list buffer. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel_EadPathBuffer.
    EadPathBuffer []Evpn_Active_InternalLabels_InternalLabel_EadPathBuffer

    // EAD/EVI Path list buffer. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel_EviPathBuffer.
    EviPathBuffer []Evpn_Active_InternalLabels_InternalLabel_EviPathBuffer

    // Summary Path list buffer. The type is slice of
    // Evpn_Active_InternalLabels_InternalLabel_SummaryPathBuffer.
    SummaryPathBuffer []Evpn_Active_InternalLabels_InternalLabel_SummaryPathBuffer
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

    internalLabel.EntityData.Children = make(map[string]types.YChild)
    internalLabel.EntityData.Children["mac-path-buffer"] = types.YChild{"MacPathBuffer", nil}
    for i := range internalLabel.MacPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.MacPathBuffer[i])] = types.YChild{"MacPathBuffer", &internalLabel.MacPathBuffer[i]}
    }
    internalLabel.EntityData.Children["ead-path-buffer"] = types.YChild{"EadPathBuffer", nil}
    for i := range internalLabel.EadPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.EadPathBuffer[i])] = types.YChild{"EadPathBuffer", &internalLabel.EadPathBuffer[i]}
    }
    internalLabel.EntityData.Children["evi-path-buffer"] = types.YChild{"EviPathBuffer", nil}
    for i := range internalLabel.EviPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.EviPathBuffer[i])] = types.YChild{"EviPathBuffer", &internalLabel.EviPathBuffer[i]}
    }
    internalLabel.EntityData.Children["summary-path-buffer"] = types.YChild{"SummaryPathBuffer", nil}
    for i := range internalLabel.SummaryPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.SummaryPathBuffer[i])] = types.YChild{"SummaryPathBuffer", &internalLabel.SummaryPathBuffer[i]}
    }
    internalLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    internalLabel.EntityData.Leafs["evi"] = types.YLeaf{"Evi", internalLabel.Evi}
    internalLabel.EntityData.Leafs["esi1"] = types.YLeaf{"Esi1", internalLabel.Esi1}
    internalLabel.EntityData.Leafs["esi2"] = types.YLeaf{"Esi2", internalLabel.Esi2}
    internalLabel.EntityData.Leafs["esi3"] = types.YLeaf{"Esi3", internalLabel.Esi3}
    internalLabel.EntityData.Leafs["esi4"] = types.YLeaf{"Esi4", internalLabel.Esi4}
    internalLabel.EntityData.Leafs["esi5"] = types.YLeaf{"Esi5", internalLabel.Esi5}
    internalLabel.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", internalLabel.EthernetTag}
    internalLabel.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", internalLabel.EviXr}
    internalLabel.EntityData.Leafs["esi"] = types.YLeaf{"Esi", internalLabel.Esi}
    internalLabel.EntityData.Leafs["tag"] = types.YLeaf{"Tag", internalLabel.Tag}
    internalLabel.EntityData.Leafs["internal-label"] = types.YLeaf{"InternalLabel", internalLabel.InternalLabel}
    internalLabel.EntityData.Leafs["encap"] = types.YLeaf{"Encap", internalLabel.Encap}
    internalLabel.EntityData.Leafs["mac-num-paths"] = types.YLeaf{"MacNumPaths", internalLabel.MacNumPaths}
    internalLabel.EntityData.Leafs["ead-num-paths"] = types.YLeaf{"EadNumPaths", internalLabel.EadNumPaths}
    internalLabel.EntityData.Leafs["evi-num-paths"] = types.YLeaf{"EviNumPaths", internalLabel.EviNumPaths}
    internalLabel.EntityData.Leafs["sum-num-paths"] = types.YLeaf{"SumNumPaths", internalLabel.SumNumPaths}
    internalLabel.EntityData.Leafs["sum-num-active-paths"] = types.YLeaf{"SumNumActivePaths", internalLabel.SumNumActivePaths}
    internalLabel.EntityData.Leafs["resolved"] = types.YLeaf{"Resolved", internalLabel.Resolved}
    internalLabel.EntityData.Leafs["ecmp-disable"] = types.YLeaf{"EcmpDisable", internalLabel.EcmpDisable}
    internalLabel.EntityData.Leafs["redundancy-single-active"] = types.YLeaf{"RedundancySingleActive", internalLabel.RedundancySingleActive}
    internalLabel.EntityData.Leafs["redundancy-single-flow-active"] = types.YLeaf{"RedundancySingleFlowActive", internalLabel.RedundancySingleFlowActive}
    return &(internalLabel.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel_MacPathBuffer
// MAC Path list buffer
type Evpn_Active_InternalLabels_InternalLabel_MacPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    macPathBuffer.EntityData.Children = make(map[string]types.YChild)
    macPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    macPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", macPathBuffer.NextHop}
    macPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", macPathBuffer.OutputLabel}
    macPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", macPathBuffer.SrteTunnel}
    return &(macPathBuffer.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel_EadPathBuffer
// EAD/ES Path list buffer
type Evpn_Active_InternalLabels_InternalLabel_EadPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    eadPathBuffer.EntityData.Children = make(map[string]types.YChild)
    eadPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    eadPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", eadPathBuffer.NextHop}
    eadPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", eadPathBuffer.OutputLabel}
    eadPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", eadPathBuffer.SrteTunnel}
    return &(eadPathBuffer.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel_EviPathBuffer
// EAD/EVI Path list buffer
type Evpn_Active_InternalLabels_InternalLabel_EviPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    eviPathBuffer.EntityData.Children = make(map[string]types.YChild)
    eviPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    eviPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", eviPathBuffer.NextHop}
    eviPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", eviPathBuffer.OutputLabel}
    eviPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", eviPathBuffer.SrteTunnel}
    return &(eviPathBuffer.EntityData)
}

// Evpn_Active_InternalLabels_InternalLabel_SummaryPathBuffer
// Summary Path list buffer
type Evpn_Active_InternalLabels_InternalLabel_SummaryPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    summaryPathBuffer.EntityData.Children = make(map[string]types.YChild)
    summaryPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    summaryPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", summaryPathBuffer.NextHop}
    summaryPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", summaryPathBuffer.OutputLabel}
    summaryPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", summaryPathBuffer.SrteTunnel}
    return &(summaryPathBuffer.EntityData)
}

// Evpn_Active_EthernetSegments
// EVPN Ethernet-Segment Table
type Evpn_Active_EthernetSegments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet-Segment Entry. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment.
    EthernetSegment []Evpn_Active_EthernetSegments_EthernetSegment
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

    ethernetSegments.EntityData.Children = make(map[string]types.YChild)
    ethernetSegments.EntityData.Children["ethernet-segment"] = types.YChild{"EthernetSegment", nil}
    for i := range ethernetSegments.EthernetSegment {
        ethernetSegments.EntityData.Children[types.GetSegmentPath(&ethernetSegments.EthernetSegment[i])] = types.YChild{"EthernetSegment", &ethernetSegments.EthernetSegment[i]}
    }
    ethernetSegments.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetSegments.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment
// EVPN Ethernet-Segment Entry
type Evpn_Active_EthernetSegments_EthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // ES id (part 1/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
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

    // Main port ifhandle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    IfHandle interface{}

    // Main port redundancy group role. The type is L2vpnRgRole.
    MainPortRole interface{}

    // Main Port MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MainPortMac interface{}

    // Number of PWs in Up state. The type is interface{} with range:
    // 0..4294967295.
    NumUpPWs interface{}

    // ES-Import Route Target. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    RouteTarget interface{}

    // Origin of operational ES-Import RT. The type is L2vpnEvpnRtOrigin.
    RtOrigin interface{}

    // ES BGP Gates. The type is string.
    EsBgpGates interface{}

    // ES L2FIB Gates. The type is string.
    EsL2FibGates interface{}

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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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
    EthernetSegmentIdentifier []Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier

    // List of Primary services ESI/I-SIDs. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService.
    PrimaryService []Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService

    // List of Secondary services ESI/I-SIDs. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService.
    SecondaryService []Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService

    // Elected ISID service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult.
    ServiceCarvingISidelectedResult []Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult

    // Not elected ISID service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult.
    ServiceCarvingIsidNotElectedResult []Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult

    // Elected EVI service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult.
    ServiceCarvingEviElectedResult []Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult

    // Not elected EVI service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult.
    ServiceCarvingEviNotElectedResult []Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult

    // List of nexthop IPv6 addresses. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_NextHop.
    NextHop []Evpn_Active_EthernetSegments_EthernetSegment_NextHop

    // Permanent EVPN VPWS service carving results. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult.
    ServiceCarvingVpwsPermanentResult []Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult

    // Remote split horizon group labels. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel.
    RemoteSplitHorizonGroupLabel []Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
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

    ethernetSegment.EntityData.Children = make(map[string]types.YChild)
    ethernetSegment.EntityData.Children["ethernet-segment-identifier"] = types.YChild{"EthernetSegmentIdentifier", nil}
    for i := range ethernetSegment.EthernetSegmentIdentifier {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.EthernetSegmentIdentifier[i])] = types.YChild{"EthernetSegmentIdentifier", &ethernetSegment.EthernetSegmentIdentifier[i]}
    }
    ethernetSegment.EntityData.Children["primary-service"] = types.YChild{"PrimaryService", nil}
    for i := range ethernetSegment.PrimaryService {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.PrimaryService[i])] = types.YChild{"PrimaryService", &ethernetSegment.PrimaryService[i]}
    }
    ethernetSegment.EntityData.Children["secondary-service"] = types.YChild{"SecondaryService", nil}
    for i := range ethernetSegment.SecondaryService {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.SecondaryService[i])] = types.YChild{"SecondaryService", &ethernetSegment.SecondaryService[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-i-sidelected-result"] = types.YChild{"ServiceCarvingISidelectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingISidelectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingISidelectedResult[i])] = types.YChild{"ServiceCarvingISidelectedResult", &ethernetSegment.ServiceCarvingISidelectedResult[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-isid-not-elected-result"] = types.YChild{"ServiceCarvingIsidNotElectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingIsidNotElectedResult[i])] = types.YChild{"ServiceCarvingIsidNotElectedResult", &ethernetSegment.ServiceCarvingIsidNotElectedResult[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-evi-elected-result"] = types.YChild{"ServiceCarvingEviElectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingEviElectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingEviElectedResult[i])] = types.YChild{"ServiceCarvingEviElectedResult", &ethernetSegment.ServiceCarvingEviElectedResult[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-evi-not-elected-result"] = types.YChild{"ServiceCarvingEviNotElectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingEviNotElectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingEviNotElectedResult[i])] = types.YChild{"ServiceCarvingEviNotElectedResult", &ethernetSegment.ServiceCarvingEviNotElectedResult[i]}
    }
    ethernetSegment.EntityData.Children["next-hop"] = types.YChild{"NextHop", nil}
    for i := range ethernetSegment.NextHop {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.NextHop[i])] = types.YChild{"NextHop", &ethernetSegment.NextHop[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-vpws-permanent-result"] = types.YChild{"ServiceCarvingVpwsPermanentResult", nil}
    for i := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingVpwsPermanentResult[i])] = types.YChild{"ServiceCarvingVpwsPermanentResult", &ethernetSegment.ServiceCarvingVpwsPermanentResult[i]}
    }
    ethernetSegment.EntityData.Children["remote-split-horizon-group-label"] = types.YChild{"RemoteSplitHorizonGroupLabel", nil}
    for i := range ethernetSegment.RemoteSplitHorizonGroupLabel {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.RemoteSplitHorizonGroupLabel[i])] = types.YChild{"RemoteSplitHorizonGroupLabel", &ethernetSegment.RemoteSplitHorizonGroupLabel[i]}
    }
    ethernetSegment.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegment.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", ethernetSegment.InterfaceName}
    ethernetSegment.EntityData.Leafs["esi1"] = types.YLeaf{"Esi1", ethernetSegment.Esi1}
    ethernetSegment.EntityData.Leafs["esi2"] = types.YLeaf{"Esi2", ethernetSegment.Esi2}
    ethernetSegment.EntityData.Leafs["esi3"] = types.YLeaf{"Esi3", ethernetSegment.Esi3}
    ethernetSegment.EntityData.Leafs["esi4"] = types.YLeaf{"Esi4", ethernetSegment.Esi4}
    ethernetSegment.EntityData.Leafs["esi5"] = types.YLeaf{"Esi5", ethernetSegment.Esi5}
    ethernetSegment.EntityData.Leafs["esi-type"] = types.YLeaf{"EsiType", ethernetSegment.EsiType}
    ethernetSegment.EntityData.Leafs["esi-system-identifier"] = types.YLeaf{"EsiSystemIdentifier", ethernetSegment.EsiSystemIdentifier}
    ethernetSegment.EntityData.Leafs["esi-port-key"] = types.YLeaf{"EsiPortKey", ethernetSegment.EsiPortKey}
    ethernetSegment.EntityData.Leafs["esi-system-priority"] = types.YLeaf{"EsiSystemPriority", ethernetSegment.EsiSystemPriority}
    ethernetSegment.EntityData.Leafs["ethernet-segment-name"] = types.YLeaf{"EthernetSegmentName", ethernetSegment.EthernetSegmentName}
    ethernetSegment.EntityData.Leafs["ethernet-segment-state"] = types.YLeaf{"EthernetSegmentState", ethernetSegment.EthernetSegmentState}
    ethernetSegment.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", ethernetSegment.IfHandle}
    ethernetSegment.EntityData.Leafs["main-port-role"] = types.YLeaf{"MainPortRole", ethernetSegment.MainPortRole}
    ethernetSegment.EntityData.Leafs["main-port-mac"] = types.YLeaf{"MainPortMac", ethernetSegment.MainPortMac}
    ethernetSegment.EntityData.Leafs["num-up-p-ws"] = types.YLeaf{"NumUpPWs", ethernetSegment.NumUpPWs}
    ethernetSegment.EntityData.Leafs["route-target"] = types.YLeaf{"RouteTarget", ethernetSegment.RouteTarget}
    ethernetSegment.EntityData.Leafs["rt-origin"] = types.YLeaf{"RtOrigin", ethernetSegment.RtOrigin}
    ethernetSegment.EntityData.Leafs["es-bgp-gates"] = types.YLeaf{"EsBgpGates", ethernetSegment.EsBgpGates}
    ethernetSegment.EntityData.Leafs["es-l2fib-gates"] = types.YLeaf{"EsL2FibGates", ethernetSegment.EsL2FibGates}
    ethernetSegment.EntityData.Leafs["mac-flushing-mode-config"] = types.YLeaf{"MacFlushingModeConfig", ethernetSegment.MacFlushingModeConfig}
    ethernetSegment.EntityData.Leafs["load-balance-mode-config"] = types.YLeaf{"LoadBalanceModeConfig", ethernetSegment.LoadBalanceModeConfig}
    ethernetSegment.EntityData.Leafs["load-balance-mode-is-default"] = types.YLeaf{"LoadBalanceModeIsDefault", ethernetSegment.LoadBalanceModeIsDefault}
    ethernetSegment.EntityData.Leafs["load-balance-mode-oper"] = types.YLeaf{"LoadBalanceModeOper", ethernetSegment.LoadBalanceModeOper}
    ethernetSegment.EntityData.Leafs["force-single-home"] = types.YLeaf{"ForceSingleHome", ethernetSegment.ForceSingleHome}
    ethernetSegment.EntityData.Leafs["source-mac-oper"] = types.YLeaf{"SourceMacOper", ethernetSegment.SourceMacOper}
    ethernetSegment.EntityData.Leafs["source-mac-origin"] = types.YLeaf{"SourceMacOrigin", ethernetSegment.SourceMacOrigin}
    ethernetSegment.EntityData.Leafs["peering-timer"] = types.YLeaf{"PeeringTimer", ethernetSegment.PeeringTimer}
    ethernetSegment.EntityData.Leafs["peering-timer-left"] = types.YLeaf{"PeeringTimerLeft", ethernetSegment.PeeringTimerLeft}
    ethernetSegment.EntityData.Leafs["recovery-timer"] = types.YLeaf{"RecoveryTimer", ethernetSegment.RecoveryTimer}
    ethernetSegment.EntityData.Leafs["recovery-timer-left"] = types.YLeaf{"RecoveryTimerLeft", ethernetSegment.RecoveryTimerLeft}
    ethernetSegment.EntityData.Leafs["carving-timer"] = types.YLeaf{"CarvingTimer", ethernetSegment.CarvingTimer}
    ethernetSegment.EntityData.Leafs["carving-timer-left"] = types.YLeaf{"CarvingTimerLeft", ethernetSegment.CarvingTimerLeft}
    ethernetSegment.EntityData.Leafs["service-carving-mode"] = types.YLeaf{"ServiceCarvingMode", ethernetSegment.ServiceCarvingMode}
    ethernetSegment.EntityData.Leafs["primary-services-input"] = types.YLeaf{"PrimaryServicesInput", ethernetSegment.PrimaryServicesInput}
    ethernetSegment.EntityData.Leafs["secondary-services-input"] = types.YLeaf{"SecondaryServicesInput", ethernetSegment.SecondaryServicesInput}
    ethernetSegment.EntityData.Leafs["forwarder-ports"] = types.YLeaf{"ForwarderPorts", ethernetSegment.ForwarderPorts}
    ethernetSegment.EntityData.Leafs["permanent-forwarder-ports"] = types.YLeaf{"PermanentForwarderPorts", ethernetSegment.PermanentForwarderPorts}
    ethernetSegment.EntityData.Leafs["elected-forwarder-ports"] = types.YLeaf{"ElectedForwarderPorts", ethernetSegment.ElectedForwarderPorts}
    ethernetSegment.EntityData.Leafs["not-elected-forwarder-ports"] = types.YLeaf{"NotElectedForwarderPorts", ethernetSegment.NotElectedForwarderPorts}
    ethernetSegment.EntityData.Leafs["not-config-forwarder-ports"] = types.YLeaf{"NotConfigForwarderPorts", ethernetSegment.NotConfigForwarderPorts}
    ethernetSegment.EntityData.Leafs["mp-protected"] = types.YLeaf{"MpProtected", ethernetSegment.MpProtected}
    ethernetSegment.EntityData.Leafs["nve-anycast-vtep"] = types.YLeaf{"NveAnycastVtep", ethernetSegment.NveAnycastVtep}
    ethernetSegment.EntityData.Leafs["nve-ingress-replication"] = types.YLeaf{"NveIngressReplication", ethernetSegment.NveIngressReplication}
    ethernetSegment.EntityData.Leafs["local-split-horizon-group-label-valid"] = types.YLeaf{"LocalSplitHorizonGroupLabelValid", ethernetSegment.LocalSplitHorizonGroupLabelValid}
    ethernetSegment.EntityData.Leafs["local-split-horizon-group-label"] = types.YLeaf{"LocalSplitHorizonGroupLabel", ethernetSegment.LocalSplitHorizonGroupLabel}
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

    ethernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    ethernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry}
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

    primaryService.EntityData.Children = make(map[string]types.YChild)
    primaryService.EntityData.Leafs = make(map[string]types.YLeaf)
    primaryService.EntityData.Leafs["entry"] = types.YLeaf{"Entry", primaryService.Entry}
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

    secondaryService.EntityData.Children = make(map[string]types.YChild)
    secondaryService.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryService.EntityData.Leafs["entry"] = types.YLeaf{"Entry", secondaryService.Entry}
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

    serviceCarvingISidelectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingISidelectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingISidelectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingISidelectedResult.Entry}
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

    serviceCarvingIsidNotElectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingIsidNotElectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingIsidNotElectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingIsidNotElectedResult.Entry}
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

    serviceCarvingEviElectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingEviElectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingEviElectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingEviElectedResult.Entry}
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

    serviceCarvingEviNotElectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingEviNotElectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingEviNotElectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingEviNotElectedResult.Entry}
    return &(serviceCarvingEviNotElectedResult.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_NextHop
// List of nexthop IPv6 addresses
type Evpn_Active_EthernetSegments_EthernetSegment_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}
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

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", nextHop.NextHop}
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
    Type_ interface{}

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

    serviceCarvingVpwsPermanentResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingVpwsPermanentResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingVpwsPermanentResult.EntityData.Leafs["vpn-id"] = types.YLeaf{"VpnId", serviceCarvingVpwsPermanentResult.VpnId}
    serviceCarvingVpwsPermanentResult.EntityData.Leafs["type"] = types.YLeaf{"Type_", serviceCarvingVpwsPermanentResult.Type_}
    serviceCarvingVpwsPermanentResult.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", serviceCarvingVpwsPermanentResult.EthernetTag}
    return &(serviceCarvingVpwsPermanentResult.EntityData)
}

// Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

    remoteSplitHorizonGroupLabel.EntityData.Children = make(map[string]types.YChild)
    remoteSplitHorizonGroupLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteSplitHorizonGroupLabel.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", remoteSplitHorizonGroupLabel.NextHop}
    remoteSplitHorizonGroupLabel.EntityData.Leafs["label"] = types.YLeaf{"Label", remoteSplitHorizonGroupLabel.Label}
    return &(remoteSplitHorizonGroupLabel.EntityData)
}

// Evpn_Active_AcIds
// EVPN AC ID table
type Evpn_Active_AcIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN AC ID table. The type is slice of Evpn_Active_AcIds_AcId.
    AcId []Evpn_Active_AcIds_AcId
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

    acIds.EntityData.Children = make(map[string]types.YChild)
    acIds.EntityData.Children["ac-id"] = types.YChild{"AcId", nil}
    for i := range acIds.AcId {
        acIds.EntityData.Children[types.GetSegmentPath(&acIds.AcId[i])] = types.YChild{"AcId", &acIds.AcId[i]}
    }
    acIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(acIds.EntityData)
}

// Evpn_Active_AcIds_AcId
// EVPN AC ID table
type Evpn_Active_AcIds_AcId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // AC ID. The type is interface{} with range: -2147483648..2147483647.
    AcId interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}
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

    acId.EntityData.Children = make(map[string]types.YChild)
    acId.EntityData.Leafs = make(map[string]types.YLeaf)
    acId.EntityData.Leafs["evi"] = types.YLeaf{"Evi", acId.Evi}
    acId.EntityData.Leafs["ac-id"] = types.YLeaf{"AcId", acId.AcId}
    acId.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", acId.EviXr}
    acId.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", acId.Neighbor}
    return &(acId.EntityData)
}

// Evpn_Standby
// Standby EVPN operational data
type Evpn_Standby struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Table.
    Evis Evpn_Standby_Evis

    // L2VPN EVPN Summary.
    Summary Evpn_Standby_Summary

    // L2VPN EVI Detail Table.
    EviDetail Evpn_Standby_EviDetail

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

    standby.EntityData.Children = make(map[string]types.YChild)
    standby.EntityData.Children["evis"] = types.YChild{"Evis", &standby.Evis}
    standby.EntityData.Children["summary"] = types.YChild{"Summary", &standby.Summary}
    standby.EntityData.Children["evi-detail"] = types.YChild{"EviDetail", &standby.EviDetail}
    standby.EntityData.Children["internal-labels"] = types.YChild{"InternalLabels", &standby.InternalLabels}
    standby.EntityData.Children["ethernet-segments"] = types.YChild{"EthernetSegments", &standby.EthernetSegments}
    standby.EntityData.Children["ac-ids"] = types.YChild{"AcIds", &standby.AcIds}
    standby.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(standby.EntityData)
}

// Evpn_Standby_Evis
// L2VPN EVPN EVI Table
type Evpn_Standby_Evis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Entry. The type is slice of Evpn_Standby_Evis_Evi.
    Evi []Evpn_Standby_Evis_Evi
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

    evis.EntityData.Children = make(map[string]types.YChild)
    evis.EntityData.Children["evi"] = types.YChild{"Evi", nil}
    for i := range evis.Evi {
        evis.EntityData.Children[types.GetSegmentPath(&evis.Evi[i])] = types.YChild{"Evi", &evis.Evi[i]}
    }
    evis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(evis.EntityData)
}

// Evpn_Standby_Evis_Evi
// L2VPN EVPN EVI Entry
type Evpn_Standby_Evis_Evi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN id. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type_ interface{}
}

func (evi *Evpn_Standby_Evis_Evi) GetEntityData() *types.CommonEntityData {
    evi.EntityData.YFilter = evi.YFilter
    evi.EntityData.YangName = "evi"
    evi.EntityData.BundleName = "cisco_ios_xr"
    evi.EntityData.ParentYangName = "evis"
    evi.EntityData.SegmentPath = "evi" + "[evi='" + fmt.Sprintf("%v", evi.Evi) + "']"
    evi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evi.EntityData.Children = make(map[string]types.YChild)
    evi.EntityData.Leafs = make(map[string]types.YLeaf)
    evi.EntityData.Leafs["evi"] = types.YLeaf{"Evi", evi.Evi}
    evi.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", evi.EviXr}
    evi.EntityData.Leafs["bd-name"] = types.YLeaf{"BdName", evi.BdName}
    evi.EntityData.Leafs["type"] = types.YLeaf{"Type_", evi.Type_}
    return &(evi.EntityData)
}

// Evpn_Standby_Summary
// L2VPN EVPN Summary
type Evpn_Standby_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Router ID. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // BGP AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // Number of EVI DB Entries. The type is interface{} with range:
    // 0..4294967295.
    EvIs interface{}

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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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
    L2RibThrottle interface{}

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

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", summary.RouterId}
    summary.EntityData.Leafs["as"] = types.YLeaf{"As", summary.As}
    summary.EntityData.Leafs["ev-is"] = types.YLeaf{"EvIs", summary.EvIs}
    summary.EntityData.Leafs["local-mac-routes"] = types.YLeaf{"LocalMacRoutes", summary.LocalMacRoutes}
    summary.EntityData.Leafs["local-ipv4-mac-routes"] = types.YLeaf{"LocalIpv4MacRoutes", summary.LocalIpv4MacRoutes}
    summary.EntityData.Leafs["local-ipv6-mac-routes"] = types.YLeaf{"LocalIpv6MacRoutes", summary.LocalIpv6MacRoutes}
    summary.EntityData.Leafs["es-global-mac-routes"] = types.YLeaf{"EsGlobalMacRoutes", summary.EsGlobalMacRoutes}
    summary.EntityData.Leafs["remote-mac-routes"] = types.YLeaf{"RemoteMacRoutes", summary.RemoteMacRoutes}
    summary.EntityData.Leafs["remote-soo-mac-routes"] = types.YLeaf{"RemoteSooMacRoutes", summary.RemoteSooMacRoutes}
    summary.EntityData.Leafs["remote-ipv4-mac-routes"] = types.YLeaf{"RemoteIpv4MacRoutes", summary.RemoteIpv4MacRoutes}
    summary.EntityData.Leafs["remote-ipv6-mac-routes"] = types.YLeaf{"RemoteIpv6MacRoutes", summary.RemoteIpv6MacRoutes}
    summary.EntityData.Leafs["local-imcast-routes"] = types.YLeaf{"LocalImcastRoutes", summary.LocalImcastRoutes}
    summary.EntityData.Leafs["remote-imcast-routes"] = types.YLeaf{"RemoteImcastRoutes", summary.RemoteImcastRoutes}
    summary.EntityData.Leafs["labels"] = types.YLeaf{"Labels", summary.Labels}
    summary.EntityData.Leafs["es-entries"] = types.YLeaf{"EsEntries", summary.EsEntries}
    summary.EntityData.Leafs["neighbor-entries"] = types.YLeaf{"NeighborEntries", summary.NeighborEntries}
    summary.EntityData.Leafs["local-ead-routes"] = types.YLeaf{"LocalEadRoutes", summary.LocalEadRoutes}
    summary.EntityData.Leafs["remote-ead-routes"] = types.YLeaf{"RemoteEadRoutes", summary.RemoteEadRoutes}
    summary.EntityData.Leafs["global-source-mac"] = types.YLeaf{"GlobalSourceMac", summary.GlobalSourceMac}
    summary.EntityData.Leafs["peering-time"] = types.YLeaf{"PeeringTime", summary.PeeringTime}
    summary.EntityData.Leafs["recovery-time"] = types.YLeaf{"RecoveryTime", summary.RecoveryTime}
    summary.EntityData.Leafs["carving-time"] = types.YLeaf{"CarvingTime", summary.CarvingTime}
    summary.EntityData.Leafs["mac-secure-move-count"] = types.YLeaf{"MacSecureMoveCount", summary.MacSecureMoveCount}
    summary.EntityData.Leafs["mac-secure-move-interval"] = types.YLeaf{"MacSecureMoveInterval", summary.MacSecureMoveInterval}
    summary.EntityData.Leafs["mac-secure-freeze-time"] = types.YLeaf{"MacSecureFreezeTime", summary.MacSecureFreezeTime}
    summary.EntityData.Leafs["mac-secure-retry-count"] = types.YLeaf{"MacSecureRetryCount", summary.MacSecureRetryCount}
    summary.EntityData.Leafs["cost-out"] = types.YLeaf{"CostOut", summary.CostOut}
    summary.EntityData.Leafs["startup-cost-in-time"] = types.YLeaf{"StartupCostInTime", summary.StartupCostInTime}
    summary.EntityData.Leafs["l2rib-throttle"] = types.YLeaf{"L2RibThrottle", summary.L2RibThrottle}
    summary.EntityData.Leafs["logging-df-election-enabled"] = types.YLeaf{"LoggingDfElectionEnabled", summary.LoggingDfElectionEnabled}
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

    eviDetail.EntityData.Children = make(map[string]types.YChild)
    eviDetail.EntityData.Children["elements"] = types.YChild{"Elements", &eviDetail.Elements}
    eviDetail.EntityData.Children["evi-children"] = types.YChild{"EviChildren", &eviDetail.EviChildren}
    eviDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eviDetail.EntityData)
}

// Evpn_Standby_EviDetail_Elements
// EVI BGP RT Detail Info Elements
type Evpn_Standby_EviDetail_Elements struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info. The type is slice of
    // Evpn_Standby_EviDetail_Elements_Element.
    Element []Evpn_Standby_EviDetail_Elements_Element
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

    elements.EntityData.Children = make(map[string]types.YChild)
    elements.EntityData.Children["element"] = types.YChild{"Element", nil}
    for i := range elements.Element {
        elements.EntityData.Children[types.GetSegmentPath(&elements.Element[i])] = types.YChild{"Element", &elements.Element[i]}
    }
    elements.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(elements.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element
// EVI BGP RT Detail Info
type Evpn_Standby_EviDetail_Elements_Element struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN id. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // EVI description. The type is string.
    Description interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type_ interface{}

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

    // EVPN Instance is Regular/Stitching side. The type is interface{} with
    // range: 0..255.
    Stitching interface{}

    // EVPN Instance encapsulation. The type is interface{} with range: 0..255.
    Encapsulation interface{}

    // Flow Label Information.
    FlowLabel Evpn_Standby_EviDetail_Elements_Element_FlowLabel

    // Automatic Route Distingtuisher.
    RdAuto Evpn_Standby_EviDetail_Elements_Element_RdAuto

    // Configured Route Distinguisher.
    RdConfigured Evpn_Standby_EviDetail_Elements_Element_RdConfigured

    // Automatic Route Target.
    RtAuto Evpn_Standby_EviDetail_Elements_Element_RtAuto

    // Automatic Route Target Stitching.
    RtAutoStitching Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching
}

func (element *Evpn_Standby_EviDetail_Elements_Element) GetEntityData() *types.CommonEntityData {
    element.EntityData.YFilter = element.YFilter
    element.EntityData.YangName = "element"
    element.EntityData.BundleName = "cisco_ios_xr"
    element.EntityData.ParentYangName = "elements"
    element.EntityData.SegmentPath = "element" + "[evi='" + fmt.Sprintf("%v", element.Evi) + "']"
    element.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    element.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    element.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    element.EntityData.Children = make(map[string]types.YChild)
    element.EntityData.Children["flow-label"] = types.YChild{"FlowLabel", &element.FlowLabel}
    element.EntityData.Children["rd-auto"] = types.YChild{"RdAuto", &element.RdAuto}
    element.EntityData.Children["rd-configured"] = types.YChild{"RdConfigured", &element.RdConfigured}
    element.EntityData.Children["rt-auto"] = types.YChild{"RtAuto", &element.RtAuto}
    element.EntityData.Children["rt-auto-stitching"] = types.YChild{"RtAutoStitching", &element.RtAutoStitching}
    element.EntityData.Leafs = make(map[string]types.YLeaf)
    element.EntityData.Leafs["evi"] = types.YLeaf{"Evi", element.Evi}
    element.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", element.EviXr}
    element.EntityData.Leafs["description"] = types.YLeaf{"Description", element.Description}
    element.EntityData.Leafs["bd-name"] = types.YLeaf{"BdName", element.BdName}
    element.EntityData.Leafs["type"] = types.YLeaf{"Type_", element.Type_}
    element.EntityData.Leafs["unicast-label"] = types.YLeaf{"UnicastLabel", element.UnicastLabel}
    element.EntityData.Leafs["multicast-label"] = types.YLeaf{"MulticastLabel", element.MulticastLabel}
    element.EntityData.Leafs["cw-disable"] = types.YLeaf{"CwDisable", element.CwDisable}
    element.EntityData.Leafs["table-policy-name"] = types.YLeaf{"TablePolicyName", element.TablePolicyName}
    element.EntityData.Leafs["forward-class"] = types.YLeaf{"ForwardClass", element.ForwardClass}
    element.EntityData.Leafs["rt-import-block-set"] = types.YLeaf{"RtImportBlockSet", element.RtImportBlockSet}
    element.EntityData.Leafs["rt-export-block-set"] = types.YLeaf{"RtExportBlockSet", element.RtExportBlockSet}
    element.EntityData.Leafs["advertise-mac"] = types.YLeaf{"AdvertiseMac", element.AdvertiseMac}
    element.EntityData.Leafs["advertise-bvi-mac"] = types.YLeaf{"AdvertiseBviMac", element.AdvertiseBviMac}
    element.EntityData.Leafs["aliasing-disabled"] = types.YLeaf{"AliasingDisabled", element.AliasingDisabled}
    element.EntityData.Leafs["unknown-unicast-flooding-disabled"] = types.YLeaf{"UnknownUnicastFloodingDisabled", element.UnknownUnicastFloodingDisabled}
    element.EntityData.Leafs["reoriginate-disabled"] = types.YLeaf{"ReoriginateDisabled", element.ReoriginateDisabled}
    element.EntityData.Leafs["stitching"] = types.YLeaf{"Stitching", element.Stitching}
    element.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", element.Encapsulation}
    return &(element.EntityData)
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

    flowLabel.EntityData.Children = make(map[string]types.YChild)
    flowLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    flowLabel.EntityData.Leafs["static-flow-label"] = types.YLeaf{"StaticFlowLabel", flowLabel.StaticFlowLabel}
    flowLabel.EntityData.Leafs["global-flow-label"] = types.YLeaf{"GlobalFlowLabel", flowLabel.GlobalFlowLabel}
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

    rdAuto.EntityData.Children = make(map[string]types.YChild)
    rdAuto.EntityData.Children["auto"] = types.YChild{"Auto", &rdAuto.Auto}
    rdAuto.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rdAuto.TwoByteAs}
    rdAuto.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rdAuto.FourByteAs}
    rdAuto.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rdAuto.V4Addr}
    rdAuto.EntityData.Leafs = make(map[string]types.YLeaf)
    rdAuto.EntityData.Leafs["rd"] = types.YLeaf{"Rd", rdAuto.Rd}
    return &(rdAuto.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto
// auto
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    auto.EntityData.Children = make(map[string]types.YChild)
    auto.EntityData.Leafs = make(map[string]types.YLeaf)
    auto.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", auto.RouterId}
    auto.EntityData.Leafs["auto-index"] = types.YLeaf{"AutoIndex", auto.AutoIndex}
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

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
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

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
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

    rdConfigured.EntityData.Children = make(map[string]types.YChild)
    rdConfigured.EntityData.Children["auto"] = types.YChild{"Auto", &rdConfigured.Auto}
    rdConfigured.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rdConfigured.TwoByteAs}
    rdConfigured.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rdConfigured.FourByteAs}
    rdConfigured.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rdConfigured.V4Addr}
    rdConfigured.EntityData.Leafs = make(map[string]types.YLeaf)
    rdConfigured.EntityData.Leafs["rd"] = types.YLeaf{"Rd", rdConfigured.Rd}
    return &(rdConfigured.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto
// auto
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    auto.EntityData.Children = make(map[string]types.YChild)
    auto.EntityData.Leafs = make(map[string]types.YLeaf)
    auto.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", auto.RouterId}
    auto.EntityData.Leafs["auto-index"] = types.YLeaf{"AutoIndex", auto.AutoIndex}
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

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
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

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
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

    rtAuto.EntityData.Children = make(map[string]types.YChild)
    rtAuto.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rtAuto.TwoByteAs}
    rtAuto.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rtAuto.FourByteAs}
    rtAuto.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rtAuto.V4Addr}
    rtAuto.EntityData.Children["es-import"] = types.YChild{"EsImport", &rtAuto.EsImport}
    rtAuto.EntityData.Leafs = make(map[string]types.YLeaf)
    rtAuto.EntityData.Leafs["rt"] = types.YLeaf{"Rt", rtAuto.Rt}
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

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
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

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
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

    esImport.EntityData.Children = make(map[string]types.YChild)
    esImport.EntityData.Leafs = make(map[string]types.YLeaf)
    esImport.EntityData.Leafs["high-bytes"] = types.YLeaf{"HighBytes", esImport.HighBytes}
    esImport.EntityData.Leafs["low-bytes"] = types.YLeaf{"LowBytes", esImport.LowBytes}
    return &(esImport.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching
// Automatic Route Target Stitching
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs

    // four byte as.
    FourByteAs Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs

    // v4 addr.
    V4Addr Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr

    // es import.
    EsImport Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport
}

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetEntityData() *types.CommonEntityData {
    rtAutoStitching.EntityData.YFilter = rtAutoStitching.YFilter
    rtAutoStitching.EntityData.YangName = "rt-auto-stitching"
    rtAutoStitching.EntityData.BundleName = "cisco_ios_xr"
    rtAutoStitching.EntityData.ParentYangName = "element"
    rtAutoStitching.EntityData.SegmentPath = "rt-auto-stitching"
    rtAutoStitching.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rtAutoStitching.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rtAutoStitching.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rtAutoStitching.EntityData.Children = make(map[string]types.YChild)
    rtAutoStitching.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &rtAutoStitching.TwoByteAs}
    rtAutoStitching.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &rtAutoStitching.FourByteAs}
    rtAutoStitching.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &rtAutoStitching.V4Addr}
    rtAutoStitching.EntityData.Children["es-import"] = types.YChild{"EsImport", &rtAutoStitching.EsImport}
    rtAutoStitching.EntityData.Leafs = make(map[string]types.YLeaf)
    rtAutoStitching.EntityData.Leafs["rt"] = types.YLeaf{"Rt", rtAutoStitching.Rt}
    return &(rtAutoStitching.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "rt-auto-stitching"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
    return &(twoByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "rt-auto-stitching"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "rt-auto-stitching"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
    return &(v4Addr.EntityData)
}

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport
// es import
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "rt-auto-stitching"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = make(map[string]types.YChild)
    esImport.EntityData.Leafs = make(map[string]types.YLeaf)
    esImport.EntityData.Leafs["high-bytes"] = types.YLeaf{"HighBytes", esImport.HighBytes}
    esImport.EntityData.Leafs["low-bytes"] = types.YLeaf{"LowBytes", esImport.LowBytes}
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

    eviChildren.EntityData.Children = make(map[string]types.YChild)
    eviChildren.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &eviChildren.Neighbors}
    eviChildren.EntityData.Children["ethernet-auto-discoveries"] = types.YChild{"EthernetAutoDiscoveries", &eviChildren.EthernetAutoDiscoveries}
    eviChildren.EntityData.Children["inclusive-multicasts"] = types.YChild{"InclusiveMulticasts", &eviChildren.InclusiveMulticasts}
    eviChildren.EntityData.Children["route-targets"] = types.YChild{"RouteTargets", &eviChildren.RouteTargets}
    eviChildren.EntityData.Children["macs"] = types.YChild{"Macs", &eviChildren.Macs}
    eviChildren.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eviChildren.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Neighbors
// EVPN Neighbor table
type Evpn_Standby_EviDetail_EviChildren_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Neighbor table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor.
    Neighbor []Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor
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

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children[types.GetSegmentPath(&neighbors.Neighbor[i])] = types.YChild{"Neighbor", &neighbors.Neighbor[i]}
    }
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor
// EVPN Neighbor table
type Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Neighbor IP. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}
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

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["evi"] = types.YLeaf{"Evi", neighbor.Evi}
    neighbor.EntityData.Leafs["neighbor-ip"] = types.YLeaf{"NeighborIp", neighbor.NeighborIp}
    neighbor.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", neighbor.EviXr}
    neighbor.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", neighbor.Neighbor}
    return &(neighbor.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries
// EVPN Ethernet Auto-Discovery table
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet Auto-Discovery Entry. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery.
    EthernetAutoDiscovery []Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
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

    ethernetAutoDiscoveries.EntityData.Children = make(map[string]types.YChild)
    ethernetAutoDiscoveries.EntityData.Children["ethernet-auto-discovery"] = types.YChild{"EthernetAutoDiscovery", nil}
    for i := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
        ethernetAutoDiscoveries.EntityData.Children[types.GetSegmentPath(&ethernetAutoDiscoveries.EthernetAutoDiscovery[i])] = types.YChild{"EthernetAutoDiscovery", &ethernetAutoDiscoveries.EthernetAutoDiscovery[i]}
    }
    ethernetAutoDiscoveries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetAutoDiscoveries.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
// EVPN Ethernet Auto-Discovery Entry
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // ES id (part 1/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnid interface{}

    // Service Type. The type is L2vpnEvpn.
    Type_ interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Local nexthop IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalNextHop interface{}

    // Associated local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Indication of EthernetAutoDiscovery Route is local. The type is bool.
    IsLocalEad interface{}

    // Encap type of local or remote EAD. The type is interface{} with range:
    // 0..255.
    Encap interface{}

    // Single-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleActive interface{}

    // Single-flow-active redundancy configured at remote EAD. The type is bool.
    RedundancySingleFlowActive interface{}

    // Number of items in path list buffer. The type is interface{} with range:
    // 0..4294967295.
    NumPaths interface{}

    // Ethernet Segment id. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier.
    EthernetSegmentIdentifier []Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer.
    PathBuffer []Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
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

    ethernetAutoDiscovery.EntityData.Children = make(map[string]types.YChild)
    ethernetAutoDiscovery.EntityData.Children["ethernet-segment-identifier"] = types.YChild{"EthernetSegmentIdentifier", nil}
    for i := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
        ethernetAutoDiscovery.EntityData.Children[types.GetSegmentPath(&ethernetAutoDiscovery.EthernetSegmentIdentifier[i])] = types.YChild{"EthernetSegmentIdentifier", &ethernetAutoDiscovery.EthernetSegmentIdentifier[i]}
    }
    ethernetAutoDiscovery.EntityData.Children["path-buffer"] = types.YChild{"PathBuffer", nil}
    for i := range ethernetAutoDiscovery.PathBuffer {
        ethernetAutoDiscovery.EntityData.Children[types.GetSegmentPath(&ethernetAutoDiscovery.PathBuffer[i])] = types.YChild{"PathBuffer", &ethernetAutoDiscovery.PathBuffer[i]}
    }
    ethernetAutoDiscovery.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetAutoDiscovery.EntityData.Leafs["evi"] = types.YLeaf{"Evi", ethernetAutoDiscovery.Evi}
    ethernetAutoDiscovery.EntityData.Leafs["esi1"] = types.YLeaf{"Esi1", ethernetAutoDiscovery.Esi1}
    ethernetAutoDiscovery.EntityData.Leafs["esi2"] = types.YLeaf{"Esi2", ethernetAutoDiscovery.Esi2}
    ethernetAutoDiscovery.EntityData.Leafs["esi3"] = types.YLeaf{"Esi3", ethernetAutoDiscovery.Esi3}
    ethernetAutoDiscovery.EntityData.Leafs["esi4"] = types.YLeaf{"Esi4", ethernetAutoDiscovery.Esi4}
    ethernetAutoDiscovery.EntityData.Leafs["esi5"] = types.YLeaf{"Esi5", ethernetAutoDiscovery.Esi5}
    ethernetAutoDiscovery.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", ethernetAutoDiscovery.EthernetTag}
    ethernetAutoDiscovery.EntityData.Leafs["ethernet-vpnid"] = types.YLeaf{"EthernetVpnid", ethernetAutoDiscovery.EthernetVpnid}
    ethernetAutoDiscovery.EntityData.Leafs["type"] = types.YLeaf{"Type_", ethernetAutoDiscovery.Type_}
    ethernetAutoDiscovery.EntityData.Leafs["ethernet-tag-xr"] = types.YLeaf{"EthernetTagXr", ethernetAutoDiscovery.EthernetTagXr}
    ethernetAutoDiscovery.EntityData.Leafs["local-next-hop"] = types.YLeaf{"LocalNextHop", ethernetAutoDiscovery.LocalNextHop}
    ethernetAutoDiscovery.EntityData.Leafs["local-label"] = types.YLeaf{"LocalLabel", ethernetAutoDiscovery.LocalLabel}
    ethernetAutoDiscovery.EntityData.Leafs["is-local-ead"] = types.YLeaf{"IsLocalEad", ethernetAutoDiscovery.IsLocalEad}
    ethernetAutoDiscovery.EntityData.Leafs["encap"] = types.YLeaf{"Encap", ethernetAutoDiscovery.Encap}
    ethernetAutoDiscovery.EntityData.Leafs["redundancy-single-active"] = types.YLeaf{"RedundancySingleActive", ethernetAutoDiscovery.RedundancySingleActive}
    ethernetAutoDiscovery.EntityData.Leafs["redundancy-single-flow-active"] = types.YLeaf{"RedundancySingleFlowActive", ethernetAutoDiscovery.RedundancySingleFlowActive}
    ethernetAutoDiscovery.EntityData.Leafs["num-paths"] = types.YLeaf{"NumPaths", ethernetAutoDiscovery.NumPaths}
    return &(ethernetAutoDiscovery.EntityData)
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

    ethernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    ethernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry}
    return &(ethernetSegmentIdentifier.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
// Path List Buffer
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    pathBuffer.EntityData.Children = make(map[string]types.YChild)
    pathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    pathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", pathBuffer.NextHop}
    pathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", pathBuffer.OutputLabel}
    pathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel}
    return &(pathBuffer.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts
// L2VPN EVPN IMCAST table
type Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN IMCAST table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast.
    InclusiveMulticast []Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
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

    inclusiveMulticasts.EntityData.Children = make(map[string]types.YChild)
    inclusiveMulticasts.EntityData.Children["inclusive-multicast"] = types.YChild{"InclusiveMulticast", nil}
    for i := range inclusiveMulticasts.InclusiveMulticast {
        inclusiveMulticasts.EntityData.Children[types.GetSegmentPath(&inclusiveMulticasts.InclusiveMulticast[i])] = types.YChild{"InclusiveMulticast", &inclusiveMulticasts.InclusiveMulticast[i]}
    }
    inclusiveMulticasts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inclusiveMulticasts.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
// L2VPN EVPN IMCAST table
type Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag. The type is interface{} with range: -2147483648..2147483647.
    EthernetTag interface{}

    // Originating IP. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    OriginatingIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Originating IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    OriginatingIpXr interface{}

    // IP of nexthop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Local entry. The type is bool.
    IsLocalEntry interface{}

    // Proxy entry. The type is bool.
    IsProxyEntry interface{}

    // Encap type of local or remote IMCAST route. The type is interface{} with
    // range: 0..255.
    EncapType interface{}
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

    inclusiveMulticast.EntityData.Children = make(map[string]types.YChild)
    inclusiveMulticast.EntityData.Leafs = make(map[string]types.YLeaf)
    inclusiveMulticast.EntityData.Leafs["evi"] = types.YLeaf{"Evi", inclusiveMulticast.Evi}
    inclusiveMulticast.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", inclusiveMulticast.EthernetTag}
    inclusiveMulticast.EntityData.Leafs["originating-ip"] = types.YLeaf{"OriginatingIp", inclusiveMulticast.OriginatingIp}
    inclusiveMulticast.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", inclusiveMulticast.EviXr}
    inclusiveMulticast.EntityData.Leafs["ethernet-tag-xr"] = types.YLeaf{"EthernetTagXr", inclusiveMulticast.EthernetTagXr}
    inclusiveMulticast.EntityData.Leafs["originating-ip-xr"] = types.YLeaf{"OriginatingIpXr", inclusiveMulticast.OriginatingIpXr}
    inclusiveMulticast.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", inclusiveMulticast.NextHop}
    inclusiveMulticast.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", inclusiveMulticast.OutputLabel}
    inclusiveMulticast.EntityData.Leafs["is-local-entry"] = types.YLeaf{"IsLocalEntry", inclusiveMulticast.IsLocalEntry}
    inclusiveMulticast.EntityData.Leafs["is-proxy-entry"] = types.YLeaf{"IsProxyEntry", inclusiveMulticast.IsProxyEntry}
    inclusiveMulticast.EntityData.Leafs["encap-type"] = types.YLeaf{"EncapType", inclusiveMulticast.EncapType}
    return &(inclusiveMulticast.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets
// L2VPN EVPN EVI RT Child Table
type Evpn_Standby_EviDetail_EviChildren_RouteTargets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI RT Table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget.
    RouteTarget []Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget
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

    routeTargets.EntityData.Children = make(map[string]types.YChild)
    routeTargets.EntityData.Children["route-target"] = types.YChild{"RouteTarget", nil}
    for i := range routeTargets.RouteTarget {
        routeTargets.EntityData.Children[types.GetSegmentPath(&routeTargets.RouteTarget[i])] = types.YChild{"RouteTarget", &routeTargets.RouteTarget[i]}
    }
    routeTargets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(routeTargets.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget
// L2VPN EVPN EVI RT Table
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Role of the route target. The type is BgpRouteTargetRole.
    Role interface{}

    // Type of the route target. The type is BgpRouteTarget.
    Type_ interface{}

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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Bridge Domain Name. The type is string.
    BdName interface{}

    // VPN ID. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // RT Role. The type is L2vpnAdRtRole.
    RouteTargetRole interface{}

    // RT Stitching. The type is bool.
    RouteTargetStitching interface{}

    // Route Target.
    RouteTarget Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_
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

    routeTarget.EntityData.Children = make(map[string]types.YChild)
    routeTarget.EntityData.Children["route-target"] = types.YChild{"RouteTarget", &routeTarget.RouteTarget}
    routeTarget.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTarget.EntityData.Leafs["evi"] = types.YLeaf{"Evi", routeTarget.Evi}
    routeTarget.EntityData.Leafs["role"] = types.YLeaf{"Role", routeTarget.Role}
    routeTarget.EntityData.Leafs["type"] = types.YLeaf{"Type_", routeTarget.Type_}
    routeTarget.EntityData.Leafs["format"] = types.YLeaf{"Format", routeTarget.Format}
    routeTarget.EntityData.Leafs["as"] = types.YLeaf{"As", routeTarget.As}
    routeTarget.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", routeTarget.AsIndex}
    routeTarget.EntityData.Leafs["addr-index"] = types.YLeaf{"AddrIndex", routeTarget.AddrIndex}
    routeTarget.EntityData.Leafs["address"] = types.YLeaf{"Address", routeTarget.Address}
    routeTarget.EntityData.Leafs["bd-name"] = types.YLeaf{"BdName", routeTarget.BdName}
    routeTarget.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", routeTarget.EviXr}
    routeTarget.EntityData.Leafs["route-target-role"] = types.YLeaf{"RouteTargetRole", routeTarget.RouteTargetRole}
    routeTarget.EntityData.Leafs["route-target-stitching"] = types.YLeaf{"RouteTargetStitching", routeTarget.RouteTargetStitching}
    return &(routeTarget.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_
// Route Target
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RT. The type is L2vpnAdRt.
    Rt interface{}

    // two byte as.
    TwoByteAs Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs

    // four byte as.
    FourByteAs Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs

    // v4 addr.
    V4Addr Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr

    // es import.
    EsImport Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport
}

func (routeTarget_ *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_) GetEntityData() *types.CommonEntityData {
    routeTarget_.EntityData.YFilter = routeTarget_.YFilter
    routeTarget_.EntityData.YangName = "route-target"
    routeTarget_.EntityData.BundleName = "cisco_ios_xr"
    routeTarget_.EntityData.ParentYangName = "route-target"
    routeTarget_.EntityData.SegmentPath = "route-target"
    routeTarget_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget_.EntityData.Children = make(map[string]types.YChild)
    routeTarget_.EntityData.Children["two-byte-as"] = types.YChild{"TwoByteAs", &routeTarget_.TwoByteAs}
    routeTarget_.EntityData.Children["four-byte-as"] = types.YChild{"FourByteAs", &routeTarget_.FourByteAs}
    routeTarget_.EntityData.Children["v4-addr"] = types.YChild{"V4Addr", &routeTarget_.V4Addr}
    routeTarget_.EntityData.Children["es-import"] = types.YChild{"EsImport", &routeTarget_.EsImport}
    routeTarget_.EntityData.Leafs = make(map[string]types.YLeaf)
    routeTarget_.EntityData.Leafs["rt"] = types.YLeaf{"Rt", routeTarget_.Rt}
    return &(routeTarget_.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__TwoByteAs) GetEntityData() *types.CommonEntityData {
    twoByteAs.EntityData.YFilter = twoByteAs.YFilter
    twoByteAs.EntityData.YangName = "two-byte-as"
    twoByteAs.EntityData.BundleName = "cisco_ios_xr"
    twoByteAs.EntityData.ParentYangName = "route-target"
    twoByteAs.EntityData.SegmentPath = "two-byte-as"
    twoByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    twoByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    twoByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    twoByteAs.EntityData.Children = make(map[string]types.YChild)
    twoByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    twoByteAs.EntityData.Leafs["two-byte-as"] = types.YLeaf{"TwoByteAs", twoByteAs.TwoByteAs}
    twoByteAs.EntityData.Leafs["four-byte-index"] = types.YLeaf{"FourByteIndex", twoByteAs.FourByteIndex}
    return &(twoByteAs.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs
// four byte as
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__FourByteAs) GetEntityData() *types.CommonEntityData {
    fourByteAs.EntityData.YFilter = fourByteAs.YFilter
    fourByteAs.EntityData.YangName = "four-byte-as"
    fourByteAs.EntityData.BundleName = "cisco_ios_xr"
    fourByteAs.EntityData.ParentYangName = "route-target"
    fourByteAs.EntityData.SegmentPath = "four-byte-as"
    fourByteAs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fourByteAs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fourByteAs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fourByteAs.EntityData.Children = make(map[string]types.YChild)
    fourByteAs.EntityData.Leafs = make(map[string]types.YLeaf)
    fourByteAs.EntityData.Leafs["four-byte-as"] = types.YLeaf{"FourByteAs", fourByteAs.FourByteAs}
    fourByteAs.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", fourByteAs.TwoByteIndex}
    return &(fourByteAs.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr
// v4 addr
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__V4Addr) GetEntityData() *types.CommonEntityData {
    v4Addr.EntityData.YFilter = v4Addr.YFilter
    v4Addr.EntityData.YangName = "v4-addr"
    v4Addr.EntityData.BundleName = "cisco_ios_xr"
    v4Addr.EntityData.ParentYangName = "route-target"
    v4Addr.EntityData.SegmentPath = "v4-addr"
    v4Addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v4Addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v4Addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v4Addr.EntityData.Children = make(map[string]types.YChild)
    v4Addr.EntityData.Leafs = make(map[string]types.YLeaf)
    v4Addr.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", v4Addr.Ipv4Address}
    v4Addr.EntityData.Leafs["two-byte-index"] = types.YLeaf{"TwoByteIndex", v4Addr.TwoByteIndex}
    return &(v4Addr.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport
// es import
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget__EsImport) GetEntityData() *types.CommonEntityData {
    esImport.EntityData.YFilter = esImport.YFilter
    esImport.EntityData.YangName = "es-import"
    esImport.EntityData.BundleName = "cisco_ios_xr"
    esImport.EntityData.ParentYangName = "route-target"
    esImport.EntityData.SegmentPath = "es-import"
    esImport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esImport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esImport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esImport.EntityData.Children = make(map[string]types.YChild)
    esImport.EntityData.Leafs = make(map[string]types.YLeaf)
    esImport.EntityData.Leafs["high-bytes"] = types.YLeaf{"HighBytes", esImport.HighBytes}
    esImport.EntityData.Leafs["low-bytes"] = types.YLeaf{"LowBytes", esImport.LowBytes}
    return &(esImport.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Macs
// L2VPN EVPN EVI MAC table
type Evpn_Standby_EviDetail_EviChildren_Macs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN MAC table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Macs_Mac.
    Mac []Evpn_Standby_EviDetail_EviChildren_Macs_Mac
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

    macs.EntityData.Children = make(map[string]types.YChild)
    macs.EntityData.Children["mac"] = types.YChild{"Mac", nil}
    for i := range macs.Mac {
        macs.EntityData.Children[types.GetSegmentPath(&macs.Mac[i])] = types.YChild{"Mac", &macs.Mac[i]}
    }
    macs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macs.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac
// L2VPN EVPN MAC table
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // IP Address. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddressXr interface{}

    // IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SooNexthop interface{}

    // IP nexthop address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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

    // Local Ethernet Segment id. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier.
    LocalEthernetSegmentIdentifier []Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier

    // Remote Ethernet Segment id. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier.
    RemoteEthernetSegmentIdentifier []Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier

    // Path List Buffer. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer.
    PathBuffer []Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer
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

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["local-ethernet-segment-identifier"] = types.YChild{"LocalEthernetSegmentIdentifier", nil}
    for i := range mac.LocalEthernetSegmentIdentifier {
        mac.EntityData.Children[types.GetSegmentPath(&mac.LocalEthernetSegmentIdentifier[i])] = types.YChild{"LocalEthernetSegmentIdentifier", &mac.LocalEthernetSegmentIdentifier[i]}
    }
    mac.EntityData.Children["remote-ethernet-segment-identifier"] = types.YChild{"RemoteEthernetSegmentIdentifier", nil}
    for i := range mac.RemoteEthernetSegmentIdentifier {
        mac.EntityData.Children[types.GetSegmentPath(&mac.RemoteEthernetSegmentIdentifier[i])] = types.YChild{"RemoteEthernetSegmentIdentifier", &mac.RemoteEthernetSegmentIdentifier[i]}
    }
    mac.EntityData.Children["path-buffer"] = types.YChild{"PathBuffer", nil}
    for i := range mac.PathBuffer {
        mac.EntityData.Children[types.GetSegmentPath(&mac.PathBuffer[i])] = types.YChild{"PathBuffer", &mac.PathBuffer[i]}
    }
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    mac.EntityData.Leafs["evi"] = types.YLeaf{"Evi", mac.Evi}
    mac.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", mac.EthernetTag}
    mac.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", mac.MacAddress}
    mac.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", mac.IpAddress}
    mac.EntityData.Leafs["ethernet-tag-xr"] = types.YLeaf{"EthernetTagXr", mac.EthernetTagXr}
    mac.EntityData.Leafs["mac-address-xr"] = types.YLeaf{"MacAddressXr", mac.MacAddressXr}
    mac.EntityData.Leafs["ip-address-xr"] = types.YLeaf{"IpAddressXr", mac.IpAddressXr}
    mac.EntityData.Leafs["local-label"] = types.YLeaf{"LocalLabel", mac.LocalLabel}
    mac.EntityData.Leafs["num-paths"] = types.YLeaf{"NumPaths", mac.NumPaths}
    mac.EntityData.Leafs["is-local-mac"] = types.YLeaf{"IsLocalMac", mac.IsLocalMac}
    mac.EntityData.Leafs["is-proxy-entry"] = types.YLeaf{"IsProxyEntry", mac.IsProxyEntry}
    mac.EntityData.Leafs["is-remote-mac"] = types.YLeaf{"IsRemoteMac", mac.IsRemoteMac}
    mac.EntityData.Leafs["soo-nexthop"] = types.YLeaf{"SooNexthop", mac.SooNexthop}
    mac.EntityData.Leafs["ipnh-address"] = types.YLeaf{"IpnhAddress", mac.IpnhAddress}
    mac.EntityData.Leafs["esi-port-key"] = types.YLeaf{"EsiPortKey", mac.EsiPortKey}
    mac.EntityData.Leafs["local-encap-type"] = types.YLeaf{"LocalEncapType", mac.LocalEncapType}
    mac.EntityData.Leafs["remote-encap-type"] = types.YLeaf{"RemoteEncapType", mac.RemoteEncapType}
    mac.EntityData.Leafs["learned-bridge-port-name"] = types.YLeaf{"LearnedBridgePortName", mac.LearnedBridgePortName}
    mac.EntityData.Leafs["local-seq-id"] = types.YLeaf{"LocalSeqId", mac.LocalSeqId}
    mac.EntityData.Leafs["remote-seq-id"] = types.YLeaf{"RemoteSeqId", mac.RemoteSeqId}
    mac.EntityData.Leafs["local-l3-label"] = types.YLeaf{"LocalL3Label", mac.LocalL3Label}
    mac.EntityData.Leafs["router-mac-address"] = types.YLeaf{"RouterMacAddress", mac.RouterMacAddress}
    mac.EntityData.Leafs["mac-flush-requested"] = types.YLeaf{"MacFlushRequested", mac.MacFlushRequested}
    mac.EntityData.Leafs["mac-flush-received"] = types.YLeaf{"MacFlushReceived", mac.MacFlushReceived}
    mac.EntityData.Leafs["internal-label"] = types.YLeaf{"InternalLabel", mac.InternalLabel}
    mac.EntityData.Leafs["resolved"] = types.YLeaf{"Resolved", mac.Resolved}
    mac.EntityData.Leafs["local-is-static"] = types.YLeaf{"LocalIsStatic", mac.LocalIsStatic}
    mac.EntityData.Leafs["remote-is-static"] = types.YLeaf{"RemoteIsStatic", mac.RemoteIsStatic}
    return &(mac.EntityData)
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

    localEthernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    localEthernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    localEthernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", localEthernetSegmentIdentifier.Entry}
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

    remoteEthernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    remoteEthernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteEthernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", remoteEthernetSegmentIdentifier.Entry}
    return &(remoteEthernetSegmentIdentifier.EntityData)
}

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer
// Path List Buffer
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    pathBuffer.EntityData.Children = make(map[string]types.YChild)
    pathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    pathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", pathBuffer.NextHop}
    pathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", pathBuffer.OutputLabel}
    pathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", pathBuffer.SrteTunnel}
    return &(pathBuffer.EntityData)
}

// Evpn_Standby_InternalLabels
// EVPN Internal Label Table
type Evpn_Standby_InternalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2VPN EVPN Internal Label. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel.
    InternalLabel []Evpn_Standby_InternalLabels_InternalLabel
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

    internalLabels.EntityData.Children = make(map[string]types.YChild)
    internalLabels.EntityData.Children["internal-label"] = types.YChild{"InternalLabel", nil}
    for i := range internalLabels.InternalLabel {
        internalLabels.EntityData.Children[types.GetSegmentPath(&internalLabels.InternalLabel[i])] = types.YChild{"InternalLabel", &internalLabels.InternalLabel[i]}
    }
    internalLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(internalLabels.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel
// L2VPN EVPN Internal Label
type Evpn_Standby_InternalLabels_InternalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // ES id (part 1/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi5 interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // EVPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Ethernet Segment id. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Esi interface{}

    // Label Tag. The type is interface{} with range: 0..4294967295.
    Tag interface{}

    // MPLS Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Encap type of remote EAD/ES, EAD/EVI and MAC routes. The type is
    // interface{} with range: 0..255.
    Encap interface{}

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

    // MAC Path list buffer. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel_MacPathBuffer.
    MacPathBuffer []Evpn_Standby_InternalLabels_InternalLabel_MacPathBuffer

    // EAD/ES Path list buffer. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel_EadPathBuffer.
    EadPathBuffer []Evpn_Standby_InternalLabels_InternalLabel_EadPathBuffer

    // EAD/EVI Path list buffer. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel_EviPathBuffer.
    EviPathBuffer []Evpn_Standby_InternalLabels_InternalLabel_EviPathBuffer

    // Summary Path list buffer. The type is slice of
    // Evpn_Standby_InternalLabels_InternalLabel_SummaryPathBuffer.
    SummaryPathBuffer []Evpn_Standby_InternalLabels_InternalLabel_SummaryPathBuffer
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

    internalLabel.EntityData.Children = make(map[string]types.YChild)
    internalLabel.EntityData.Children["mac-path-buffer"] = types.YChild{"MacPathBuffer", nil}
    for i := range internalLabel.MacPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.MacPathBuffer[i])] = types.YChild{"MacPathBuffer", &internalLabel.MacPathBuffer[i]}
    }
    internalLabel.EntityData.Children["ead-path-buffer"] = types.YChild{"EadPathBuffer", nil}
    for i := range internalLabel.EadPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.EadPathBuffer[i])] = types.YChild{"EadPathBuffer", &internalLabel.EadPathBuffer[i]}
    }
    internalLabel.EntityData.Children["evi-path-buffer"] = types.YChild{"EviPathBuffer", nil}
    for i := range internalLabel.EviPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.EviPathBuffer[i])] = types.YChild{"EviPathBuffer", &internalLabel.EviPathBuffer[i]}
    }
    internalLabel.EntityData.Children["summary-path-buffer"] = types.YChild{"SummaryPathBuffer", nil}
    for i := range internalLabel.SummaryPathBuffer {
        internalLabel.EntityData.Children[types.GetSegmentPath(&internalLabel.SummaryPathBuffer[i])] = types.YChild{"SummaryPathBuffer", &internalLabel.SummaryPathBuffer[i]}
    }
    internalLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    internalLabel.EntityData.Leafs["evi"] = types.YLeaf{"Evi", internalLabel.Evi}
    internalLabel.EntityData.Leafs["esi1"] = types.YLeaf{"Esi1", internalLabel.Esi1}
    internalLabel.EntityData.Leafs["esi2"] = types.YLeaf{"Esi2", internalLabel.Esi2}
    internalLabel.EntityData.Leafs["esi3"] = types.YLeaf{"Esi3", internalLabel.Esi3}
    internalLabel.EntityData.Leafs["esi4"] = types.YLeaf{"Esi4", internalLabel.Esi4}
    internalLabel.EntityData.Leafs["esi5"] = types.YLeaf{"Esi5", internalLabel.Esi5}
    internalLabel.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", internalLabel.EthernetTag}
    internalLabel.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", internalLabel.EviXr}
    internalLabel.EntityData.Leafs["esi"] = types.YLeaf{"Esi", internalLabel.Esi}
    internalLabel.EntityData.Leafs["tag"] = types.YLeaf{"Tag", internalLabel.Tag}
    internalLabel.EntityData.Leafs["internal-label"] = types.YLeaf{"InternalLabel", internalLabel.InternalLabel}
    internalLabel.EntityData.Leafs["encap"] = types.YLeaf{"Encap", internalLabel.Encap}
    internalLabel.EntityData.Leafs["mac-num-paths"] = types.YLeaf{"MacNumPaths", internalLabel.MacNumPaths}
    internalLabel.EntityData.Leafs["ead-num-paths"] = types.YLeaf{"EadNumPaths", internalLabel.EadNumPaths}
    internalLabel.EntityData.Leafs["evi-num-paths"] = types.YLeaf{"EviNumPaths", internalLabel.EviNumPaths}
    internalLabel.EntityData.Leafs["sum-num-paths"] = types.YLeaf{"SumNumPaths", internalLabel.SumNumPaths}
    internalLabel.EntityData.Leafs["sum-num-active-paths"] = types.YLeaf{"SumNumActivePaths", internalLabel.SumNumActivePaths}
    internalLabel.EntityData.Leafs["resolved"] = types.YLeaf{"Resolved", internalLabel.Resolved}
    internalLabel.EntityData.Leafs["ecmp-disable"] = types.YLeaf{"EcmpDisable", internalLabel.EcmpDisable}
    internalLabel.EntityData.Leafs["redundancy-single-active"] = types.YLeaf{"RedundancySingleActive", internalLabel.RedundancySingleActive}
    internalLabel.EntityData.Leafs["redundancy-single-flow-active"] = types.YLeaf{"RedundancySingleFlowActive", internalLabel.RedundancySingleFlowActive}
    return &(internalLabel.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel_MacPathBuffer
// MAC Path list buffer
type Evpn_Standby_InternalLabels_InternalLabel_MacPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    macPathBuffer.EntityData.Children = make(map[string]types.YChild)
    macPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    macPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", macPathBuffer.NextHop}
    macPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", macPathBuffer.OutputLabel}
    macPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", macPathBuffer.SrteTunnel}
    return &(macPathBuffer.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel_EadPathBuffer
// EAD/ES Path list buffer
type Evpn_Standby_InternalLabels_InternalLabel_EadPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    eadPathBuffer.EntityData.Children = make(map[string]types.YChild)
    eadPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    eadPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", eadPathBuffer.NextHop}
    eadPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", eadPathBuffer.OutputLabel}
    eadPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", eadPathBuffer.SrteTunnel}
    return &(eadPathBuffer.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel_EviPathBuffer
// EAD/EVI Path list buffer
type Evpn_Standby_InternalLabels_InternalLabel_EviPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    eviPathBuffer.EntityData.Children = make(map[string]types.YChild)
    eviPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    eviPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", eviPathBuffer.NextHop}
    eviPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", eviPathBuffer.OutputLabel}
    eviPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", eviPathBuffer.SrteTunnel}
    return &(eviPathBuffer.EntityData)
}

// Evpn_Standby_InternalLabels_InternalLabel_SummaryPathBuffer
// Summary Path list buffer
type Evpn_Standby_InternalLabels_InternalLabel_SummaryPathBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

    summaryPathBuffer.EntityData.Children = make(map[string]types.YChild)
    summaryPathBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    summaryPathBuffer.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", summaryPathBuffer.NextHop}
    summaryPathBuffer.EntityData.Leafs["output-label"] = types.YLeaf{"OutputLabel", summaryPathBuffer.OutputLabel}
    summaryPathBuffer.EntityData.Leafs["srte-tunnel"] = types.YLeaf{"SrteTunnel", summaryPathBuffer.SrteTunnel}
    return &(summaryPathBuffer.EntityData)
}

// Evpn_Standby_EthernetSegments
// EVPN Ethernet-Segment Table
type Evpn_Standby_EthernetSegments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN Ethernet-Segment Entry. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment.
    EthernetSegment []Evpn_Standby_EthernetSegments_EthernetSegment
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

    ethernetSegments.EntityData.Children = make(map[string]types.YChild)
    ethernetSegments.EntityData.Children["ethernet-segment"] = types.YChild{"EthernetSegment", nil}
    for i := range ethernetSegments.EthernetSegment {
        ethernetSegments.EntityData.Children[types.GetSegmentPath(&ethernetSegments.EthernetSegment[i])] = types.YChild{"EthernetSegment", &ethernetSegments.EthernetSegment[i]}
    }
    ethernetSegments.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetSegments.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment
// EVPN Ethernet-Segment Entry
type Evpn_Standby_EthernetSegments_EthernetSegment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // ES id (part 1/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi1 interface{}

    // ES id (part 2/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi2 interface{}

    // ES id (part 3/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi3 interface{}

    // ES id (part 4/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    Esi4 interface{}

    // ES id (part 5/5). The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
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

    // Main port ifhandle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    IfHandle interface{}

    // Main port redundancy group role. The type is L2vpnRgRole.
    MainPortRole interface{}

    // Main Port MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MainPortMac interface{}

    // Number of PWs in Up state. The type is interface{} with range:
    // 0..4294967295.
    NumUpPWs interface{}

    // ES-Import Route Target. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    RouteTarget interface{}

    // Origin of operational ES-Import RT. The type is L2vpnEvpnRtOrigin.
    RtOrigin interface{}

    // ES BGP Gates. The type is string.
    EsBgpGates interface{}

    // ES L2FIB Gates. The type is string.
    EsL2FibGates interface{}

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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
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
    EthernetSegmentIdentifier []Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier

    // List of Primary services ESI/I-SIDs. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService.
    PrimaryService []Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService

    // List of Secondary services ESI/I-SIDs. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService.
    SecondaryService []Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService

    // Elected ISID service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult.
    ServiceCarvingISidelectedResult []Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult

    // Not elected ISID service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult.
    ServiceCarvingIsidNotElectedResult []Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult

    // Elected EVI service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult.
    ServiceCarvingEviElectedResult []Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult

    // Not elected EVI service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult.
    ServiceCarvingEviNotElectedResult []Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult

    // List of nexthop IPv6 addresses. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_NextHop.
    NextHop []Evpn_Standby_EthernetSegments_EthernetSegment_NextHop

    // Permanent EVPN VPWS service carving results. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult.
    ServiceCarvingVpwsPermanentResult []Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult

    // Remote split horizon group labels. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel.
    RemoteSplitHorizonGroupLabel []Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
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

    ethernetSegment.EntityData.Children = make(map[string]types.YChild)
    ethernetSegment.EntityData.Children["ethernet-segment-identifier"] = types.YChild{"EthernetSegmentIdentifier", nil}
    for i := range ethernetSegment.EthernetSegmentIdentifier {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.EthernetSegmentIdentifier[i])] = types.YChild{"EthernetSegmentIdentifier", &ethernetSegment.EthernetSegmentIdentifier[i]}
    }
    ethernetSegment.EntityData.Children["primary-service"] = types.YChild{"PrimaryService", nil}
    for i := range ethernetSegment.PrimaryService {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.PrimaryService[i])] = types.YChild{"PrimaryService", &ethernetSegment.PrimaryService[i]}
    }
    ethernetSegment.EntityData.Children["secondary-service"] = types.YChild{"SecondaryService", nil}
    for i := range ethernetSegment.SecondaryService {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.SecondaryService[i])] = types.YChild{"SecondaryService", &ethernetSegment.SecondaryService[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-i-sidelected-result"] = types.YChild{"ServiceCarvingISidelectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingISidelectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingISidelectedResult[i])] = types.YChild{"ServiceCarvingISidelectedResult", &ethernetSegment.ServiceCarvingISidelectedResult[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-isid-not-elected-result"] = types.YChild{"ServiceCarvingIsidNotElectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingIsidNotElectedResult[i])] = types.YChild{"ServiceCarvingIsidNotElectedResult", &ethernetSegment.ServiceCarvingIsidNotElectedResult[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-evi-elected-result"] = types.YChild{"ServiceCarvingEviElectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingEviElectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingEviElectedResult[i])] = types.YChild{"ServiceCarvingEviElectedResult", &ethernetSegment.ServiceCarvingEviElectedResult[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-evi-not-elected-result"] = types.YChild{"ServiceCarvingEviNotElectedResult", nil}
    for i := range ethernetSegment.ServiceCarvingEviNotElectedResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingEviNotElectedResult[i])] = types.YChild{"ServiceCarvingEviNotElectedResult", &ethernetSegment.ServiceCarvingEviNotElectedResult[i]}
    }
    ethernetSegment.EntityData.Children["next-hop"] = types.YChild{"NextHop", nil}
    for i := range ethernetSegment.NextHop {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.NextHop[i])] = types.YChild{"NextHop", &ethernetSegment.NextHop[i]}
    }
    ethernetSegment.EntityData.Children["service-carving-vpws-permanent-result"] = types.YChild{"ServiceCarvingVpwsPermanentResult", nil}
    for i := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.ServiceCarvingVpwsPermanentResult[i])] = types.YChild{"ServiceCarvingVpwsPermanentResult", &ethernetSegment.ServiceCarvingVpwsPermanentResult[i]}
    }
    ethernetSegment.EntityData.Children["remote-split-horizon-group-label"] = types.YChild{"RemoteSplitHorizonGroupLabel", nil}
    for i := range ethernetSegment.RemoteSplitHorizonGroupLabel {
        ethernetSegment.EntityData.Children[types.GetSegmentPath(&ethernetSegment.RemoteSplitHorizonGroupLabel[i])] = types.YChild{"RemoteSplitHorizonGroupLabel", &ethernetSegment.RemoteSplitHorizonGroupLabel[i]}
    }
    ethernetSegment.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegment.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", ethernetSegment.InterfaceName}
    ethernetSegment.EntityData.Leafs["esi1"] = types.YLeaf{"Esi1", ethernetSegment.Esi1}
    ethernetSegment.EntityData.Leafs["esi2"] = types.YLeaf{"Esi2", ethernetSegment.Esi2}
    ethernetSegment.EntityData.Leafs["esi3"] = types.YLeaf{"Esi3", ethernetSegment.Esi3}
    ethernetSegment.EntityData.Leafs["esi4"] = types.YLeaf{"Esi4", ethernetSegment.Esi4}
    ethernetSegment.EntityData.Leafs["esi5"] = types.YLeaf{"Esi5", ethernetSegment.Esi5}
    ethernetSegment.EntityData.Leafs["esi-type"] = types.YLeaf{"EsiType", ethernetSegment.EsiType}
    ethernetSegment.EntityData.Leafs["esi-system-identifier"] = types.YLeaf{"EsiSystemIdentifier", ethernetSegment.EsiSystemIdentifier}
    ethernetSegment.EntityData.Leafs["esi-port-key"] = types.YLeaf{"EsiPortKey", ethernetSegment.EsiPortKey}
    ethernetSegment.EntityData.Leafs["esi-system-priority"] = types.YLeaf{"EsiSystemPriority", ethernetSegment.EsiSystemPriority}
    ethernetSegment.EntityData.Leafs["ethernet-segment-name"] = types.YLeaf{"EthernetSegmentName", ethernetSegment.EthernetSegmentName}
    ethernetSegment.EntityData.Leafs["ethernet-segment-state"] = types.YLeaf{"EthernetSegmentState", ethernetSegment.EthernetSegmentState}
    ethernetSegment.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", ethernetSegment.IfHandle}
    ethernetSegment.EntityData.Leafs["main-port-role"] = types.YLeaf{"MainPortRole", ethernetSegment.MainPortRole}
    ethernetSegment.EntityData.Leafs["main-port-mac"] = types.YLeaf{"MainPortMac", ethernetSegment.MainPortMac}
    ethernetSegment.EntityData.Leafs["num-up-p-ws"] = types.YLeaf{"NumUpPWs", ethernetSegment.NumUpPWs}
    ethernetSegment.EntityData.Leafs["route-target"] = types.YLeaf{"RouteTarget", ethernetSegment.RouteTarget}
    ethernetSegment.EntityData.Leafs["rt-origin"] = types.YLeaf{"RtOrigin", ethernetSegment.RtOrigin}
    ethernetSegment.EntityData.Leafs["es-bgp-gates"] = types.YLeaf{"EsBgpGates", ethernetSegment.EsBgpGates}
    ethernetSegment.EntityData.Leafs["es-l2fib-gates"] = types.YLeaf{"EsL2FibGates", ethernetSegment.EsL2FibGates}
    ethernetSegment.EntityData.Leafs["mac-flushing-mode-config"] = types.YLeaf{"MacFlushingModeConfig", ethernetSegment.MacFlushingModeConfig}
    ethernetSegment.EntityData.Leafs["load-balance-mode-config"] = types.YLeaf{"LoadBalanceModeConfig", ethernetSegment.LoadBalanceModeConfig}
    ethernetSegment.EntityData.Leafs["load-balance-mode-is-default"] = types.YLeaf{"LoadBalanceModeIsDefault", ethernetSegment.LoadBalanceModeIsDefault}
    ethernetSegment.EntityData.Leafs["load-balance-mode-oper"] = types.YLeaf{"LoadBalanceModeOper", ethernetSegment.LoadBalanceModeOper}
    ethernetSegment.EntityData.Leafs["force-single-home"] = types.YLeaf{"ForceSingleHome", ethernetSegment.ForceSingleHome}
    ethernetSegment.EntityData.Leafs["source-mac-oper"] = types.YLeaf{"SourceMacOper", ethernetSegment.SourceMacOper}
    ethernetSegment.EntityData.Leafs["source-mac-origin"] = types.YLeaf{"SourceMacOrigin", ethernetSegment.SourceMacOrigin}
    ethernetSegment.EntityData.Leafs["peering-timer"] = types.YLeaf{"PeeringTimer", ethernetSegment.PeeringTimer}
    ethernetSegment.EntityData.Leafs["peering-timer-left"] = types.YLeaf{"PeeringTimerLeft", ethernetSegment.PeeringTimerLeft}
    ethernetSegment.EntityData.Leafs["recovery-timer"] = types.YLeaf{"RecoveryTimer", ethernetSegment.RecoveryTimer}
    ethernetSegment.EntityData.Leafs["recovery-timer-left"] = types.YLeaf{"RecoveryTimerLeft", ethernetSegment.RecoveryTimerLeft}
    ethernetSegment.EntityData.Leafs["carving-timer"] = types.YLeaf{"CarvingTimer", ethernetSegment.CarvingTimer}
    ethernetSegment.EntityData.Leafs["carving-timer-left"] = types.YLeaf{"CarvingTimerLeft", ethernetSegment.CarvingTimerLeft}
    ethernetSegment.EntityData.Leafs["service-carving-mode"] = types.YLeaf{"ServiceCarvingMode", ethernetSegment.ServiceCarvingMode}
    ethernetSegment.EntityData.Leafs["primary-services-input"] = types.YLeaf{"PrimaryServicesInput", ethernetSegment.PrimaryServicesInput}
    ethernetSegment.EntityData.Leafs["secondary-services-input"] = types.YLeaf{"SecondaryServicesInput", ethernetSegment.SecondaryServicesInput}
    ethernetSegment.EntityData.Leafs["forwarder-ports"] = types.YLeaf{"ForwarderPorts", ethernetSegment.ForwarderPorts}
    ethernetSegment.EntityData.Leafs["permanent-forwarder-ports"] = types.YLeaf{"PermanentForwarderPorts", ethernetSegment.PermanentForwarderPorts}
    ethernetSegment.EntityData.Leafs["elected-forwarder-ports"] = types.YLeaf{"ElectedForwarderPorts", ethernetSegment.ElectedForwarderPorts}
    ethernetSegment.EntityData.Leafs["not-elected-forwarder-ports"] = types.YLeaf{"NotElectedForwarderPorts", ethernetSegment.NotElectedForwarderPorts}
    ethernetSegment.EntityData.Leafs["not-config-forwarder-ports"] = types.YLeaf{"NotConfigForwarderPorts", ethernetSegment.NotConfigForwarderPorts}
    ethernetSegment.EntityData.Leafs["mp-protected"] = types.YLeaf{"MpProtected", ethernetSegment.MpProtected}
    ethernetSegment.EntityData.Leafs["nve-anycast-vtep"] = types.YLeaf{"NveAnycastVtep", ethernetSegment.NveAnycastVtep}
    ethernetSegment.EntityData.Leafs["nve-ingress-replication"] = types.YLeaf{"NveIngressReplication", ethernetSegment.NveIngressReplication}
    ethernetSegment.EntityData.Leafs["local-split-horizon-group-label-valid"] = types.YLeaf{"LocalSplitHorizonGroupLabelValid", ethernetSegment.LocalSplitHorizonGroupLabelValid}
    ethernetSegment.EntityData.Leafs["local-split-horizon-group-label"] = types.YLeaf{"LocalSplitHorizonGroupLabel", ethernetSegment.LocalSplitHorizonGroupLabel}
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

    ethernetSegmentIdentifier.EntityData.Children = make(map[string]types.YChild)
    ethernetSegmentIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetSegmentIdentifier.EntityData.Leafs["entry"] = types.YLeaf{"Entry", ethernetSegmentIdentifier.Entry}
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

    primaryService.EntityData.Children = make(map[string]types.YChild)
    primaryService.EntityData.Leafs = make(map[string]types.YLeaf)
    primaryService.EntityData.Leafs["entry"] = types.YLeaf{"Entry", primaryService.Entry}
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

    secondaryService.EntityData.Children = make(map[string]types.YChild)
    secondaryService.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryService.EntityData.Leafs["entry"] = types.YLeaf{"Entry", secondaryService.Entry}
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

    serviceCarvingISidelectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingISidelectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingISidelectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingISidelectedResult.Entry}
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

    serviceCarvingIsidNotElectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingIsidNotElectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingIsidNotElectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingIsidNotElectedResult.Entry}
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

    serviceCarvingEviElectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingEviElectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingEviElectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingEviElectedResult.Entry}
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

    serviceCarvingEviNotElectedResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingEviNotElectedResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingEviNotElectedResult.EntityData.Leafs["entry"] = types.YLeaf{"Entry", serviceCarvingEviNotElectedResult.Entry}
    return &(serviceCarvingEviNotElectedResult.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_NextHop
// List of nexthop IPv6 addresses
type Evpn_Standby_EthernetSegments_EthernetSegment_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}
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

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", nextHop.NextHop}
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
    Type_ interface{}

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

    serviceCarvingVpwsPermanentResult.EntityData.Children = make(map[string]types.YChild)
    serviceCarvingVpwsPermanentResult.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceCarvingVpwsPermanentResult.EntityData.Leafs["vpn-id"] = types.YLeaf{"VpnId", serviceCarvingVpwsPermanentResult.VpnId}
    serviceCarvingVpwsPermanentResult.EntityData.Leafs["type"] = types.YLeaf{"Type_", serviceCarvingVpwsPermanentResult.Type_}
    serviceCarvingVpwsPermanentResult.EntityData.Leafs["ethernet-tag"] = types.YLeaf{"EthernetTag", serviceCarvingVpwsPermanentResult.EthernetTag}
    return &(serviceCarvingVpwsPermanentResult.EntityData)
}

// Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

    remoteSplitHorizonGroupLabel.EntityData.Children = make(map[string]types.YChild)
    remoteSplitHorizonGroupLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteSplitHorizonGroupLabel.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", remoteSplitHorizonGroupLabel.NextHop}
    remoteSplitHorizonGroupLabel.EntityData.Leafs["label"] = types.YLeaf{"Label", remoteSplitHorizonGroupLabel.Label}
    return &(remoteSplitHorizonGroupLabel.EntityData)
}

// Evpn_Standby_AcIds
// EVPN AC ID table
type Evpn_Standby_AcIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN AC ID table. The type is slice of Evpn_Standby_AcIds_AcId.
    AcId []Evpn_Standby_AcIds_AcId
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

    acIds.EntityData.Children = make(map[string]types.YChild)
    acIds.EntityData.Children["ac-id"] = types.YChild{"AcId", nil}
    for i := range acIds.AcId {
        acIds.EntityData.Children[types.GetSegmentPath(&acIds.AcId[i])] = types.YChild{"AcId", &acIds.AcId[i]}
    }
    acIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(acIds.EntityData)
}

// Evpn_Standby_AcIds_AcId
// EVPN AC ID table
type Evpn_Standby_AcIds_AcId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // AC ID. The type is interface{} with range: -2147483648..2147483647.
    AcId interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Neighbor interface{}
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

    acId.EntityData.Children = make(map[string]types.YChild)
    acId.EntityData.Leafs = make(map[string]types.YLeaf)
    acId.EntityData.Leafs["evi"] = types.YLeaf{"Evi", acId.Evi}
    acId.EntityData.Leafs["ac-id"] = types.YLeaf{"AcId", acId.AcId}
    acId.EntityData.Leafs["evi-xr"] = types.YLeaf{"EviXr", acId.EviXr}
    acId.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", acId.Neighbor}
    return &(acId.EntityData)
}

