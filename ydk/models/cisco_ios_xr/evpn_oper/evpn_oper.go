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

// BgpRouteTarget represents Bgp route target
type BgpRouteTarget string

const (
    // RT is default type
    BgpRouteTarget_no_stitching BgpRouteTarget = "no-stitching"

    // RT is for stitching (Golf-L2)
    BgpRouteTarget_stitching BgpRouteTarget = "stitching"
)

// Evpn
// EVPN Operational Table
type Evpn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of EVPN operational data for a particular node.
    Nodes Evpn_Nodes

    // Active EVPN operational data.
    Active Evpn_Active

    // Standby EVPN operational data.
    Standby Evpn_Standby
}

func (evpn *Evpn) GetFilter() yfilter.YFilter { return evpn.YFilter }

func (evpn *Evpn) SetFilter(yf yfilter.YFilter) { evpn.YFilter = yf }

func (evpn *Evpn) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "active" { return "Active" }
    if yname == "standby" { return "Standby" }
    return ""
}

func (evpn *Evpn) GetSegmentPath() string {
    return "Cisco-IOS-XR-evpn-oper:evpn"
}

func (evpn *Evpn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &evpn.Nodes
    }
    if childYangName == "active" {
        return &evpn.Active
    }
    if childYangName == "standby" {
        return &evpn.Standby
    }
    return nil
}

func (evpn *Evpn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &evpn.Nodes
    children["active"] = &evpn.Active
    children["standby"] = &evpn.Standby
    return children
}

func (evpn *Evpn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (evpn *Evpn) GetBundleName() string { return "cisco_ios_xr" }

func (evpn *Evpn) GetYangName() string { return "evpn" }

func (evpn *Evpn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evpn *Evpn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evpn *Evpn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evpn *Evpn) SetParent(parent types.Entity) { evpn.parent = parent }

func (evpn *Evpn) GetParent() types.Entity { return evpn.parent }

func (evpn *Evpn) GetParentYangName() string { return "Cisco-IOS-XR-evpn-oper" }

// Evpn_Nodes
// Table of EVPN operational data for a particular
// node
type Evpn_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN operational data for a particular node. The type is slice of
    // Evpn_Nodes_Node.
    Node []Evpn_Nodes_Node
}

func (nodes *Evpn_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Evpn_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Evpn_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Evpn_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Evpn_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Evpn_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Evpn_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Evpn_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Evpn_Nodes) GetYangName() string { return "nodes" }

func (nodes *Evpn_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Evpn_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Evpn_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Evpn_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Evpn_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Evpn_Nodes) GetParentYangName() string { return "evpn" }

// Evpn_Nodes_Node
// EVPN operational data for a particular node
type Evpn_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // L2VPN EVPN EVI Table.
    Evis Evpn_Nodes_Node_Evis

    // L2VPN EVPN Summary.
    Summary Evpn_Nodes_Node_Summary

    // L2VPN EVI Detail Table.
    EviDetail Evpn_Nodes_Node_EviDetail

    // EVPN Ethernet-Segment Table.
    EthernetSegments Evpn_Nodes_Node_EthernetSegments

    // EVPN AC ID table.
    AcIds Evpn_Nodes_Node_AcIds
}

func (node *Evpn_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Evpn_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Evpn_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "evis" { return "Evis" }
    if yname == "summary" { return "Summary" }
    if yname == "evi-detail" { return "EviDetail" }
    if yname == "ethernet-segments" { return "EthernetSegments" }
    if yname == "ac-ids" { return "AcIds" }
    return ""
}

func (node *Evpn_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *Evpn_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "evis" {
        return &node.Evis
    }
    if childYangName == "summary" {
        return &node.Summary
    }
    if childYangName == "evi-detail" {
        return &node.EviDetail
    }
    if childYangName == "ethernet-segments" {
        return &node.EthernetSegments
    }
    if childYangName == "ac-ids" {
        return &node.AcIds
    }
    return nil
}

func (node *Evpn_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["evis"] = &node.Evis
    children["summary"] = &node.Summary
    children["evi-detail"] = &node.EviDetail
    children["ethernet-segments"] = &node.EthernetSegments
    children["ac-ids"] = &node.AcIds
    return children
}

func (node *Evpn_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    return leafs
}

func (node *Evpn_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Evpn_Nodes_Node) GetYangName() string { return "node" }

func (node *Evpn_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Evpn_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Evpn_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Evpn_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Evpn_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Evpn_Nodes_Node) GetParentYangName() string { return "nodes" }

// Evpn_Nodes_Node_Evis
// L2VPN EVPN EVI Table
type Evpn_Nodes_Node_Evis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Entry. The type is slice of Evpn_Nodes_Node_Evis_Evi.
    Evi []Evpn_Nodes_Node_Evis_Evi
}

func (evis *Evpn_Nodes_Node_Evis) GetFilter() yfilter.YFilter { return evis.YFilter }

func (evis *Evpn_Nodes_Node_Evis) SetFilter(yf yfilter.YFilter) { evis.YFilter = yf }

func (evis *Evpn_Nodes_Node_Evis) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    return ""
}

func (evis *Evpn_Nodes_Node_Evis) GetSegmentPath() string {
    return "evis"
}

func (evis *Evpn_Nodes_Node_Evis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "evi" {
        for _, c := range evis.Evi {
            if evis.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_Evis_Evi{}
        evis.Evi = append(evis.Evi, child)
        return &evis.Evi[len(evis.Evi)-1]
    }
    return nil
}

func (evis *Evpn_Nodes_Node_Evis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range evis.Evi {
        children[evis.Evi[i].GetSegmentPath()] = &evis.Evi[i]
    }
    return children
}

func (evis *Evpn_Nodes_Node_Evis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (evis *Evpn_Nodes_Node_Evis) GetBundleName() string { return "cisco_ios_xr" }

func (evis *Evpn_Nodes_Node_Evis) GetYangName() string { return "evis" }

func (evis *Evpn_Nodes_Node_Evis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evis *Evpn_Nodes_Node_Evis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evis *Evpn_Nodes_Node_Evis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evis *Evpn_Nodes_Node_Evis) SetParent(parent types.Entity) { evis.parent = parent }

func (evis *Evpn_Nodes_Node_Evis) GetParent() types.Entity { return evis.parent }

func (evis *Evpn_Nodes_Node_Evis) GetParentYangName() string { return "node" }

// Evpn_Nodes_Node_Evis_Evi
// L2VPN EVPN EVI Entry
type Evpn_Nodes_Node_Evis_Evi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN id. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evi *Evpn_Nodes_Node_Evis_Evi) GetFilter() yfilter.YFilter { return evi.YFilter }

func (evi *Evpn_Nodes_Node_Evis_Evi) SetFilter(yf yfilter.YFilter) { evi.YFilter = yf }

func (evi *Evpn_Nodes_Node_Evis_Evi) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "bd-name" { return "BdName" }
    if yname == "type" { return "Type" }
    return ""
}

func (evi *Evpn_Nodes_Node_Evis_Evi) GetSegmentPath() string {
    return "evi" + "[evi='" + fmt.Sprintf("%v", evi.Evi) + "']"
}

func (evi *Evpn_Nodes_Node_Evis_Evi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (evi *Evpn_Nodes_Node_Evis_Evi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (evi *Evpn_Nodes_Node_Evis_Evi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = evi.Evi
    leafs["evi-xr"] = evi.EviXr
    leafs["bd-name"] = evi.BdName
    leafs["type"] = evi.Type
    return leafs
}

func (evi *Evpn_Nodes_Node_Evis_Evi) GetBundleName() string { return "cisco_ios_xr" }

func (evi *Evpn_Nodes_Node_Evis_Evi) GetYangName() string { return "evi" }

func (evi *Evpn_Nodes_Node_Evis_Evi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evi *Evpn_Nodes_Node_Evis_Evi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evi *Evpn_Nodes_Node_Evis_Evi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evi *Evpn_Nodes_Node_Evis_Evi) SetParent(parent types.Entity) { evi.parent = parent }

func (evi *Evpn_Nodes_Node_Evis_Evi) GetParent() types.Entity { return evi.parent }

func (evi *Evpn_Nodes_Node_Evis_Evi) GetParentYangName() string { return "evis" }

// Evpn_Nodes_Node_Summary
// L2VPN EVPN Summary
type Evpn_Nodes_Node_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Router ID. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    GlobalSourceMac interface{}

    // EVPN ES Peering Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    PeeringTime interface{}

    // EVPN ES Recovery Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    RecoveryTime interface{}

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

func (summary *Evpn_Nodes_Node_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Evpn_Nodes_Node_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Evpn_Nodes_Node_Summary) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "as" { return "As" }
    if yname == "ev-is" { return "EvIs" }
    if yname == "local-mac-routes" { return "LocalMacRoutes" }
    if yname == "local-ipv4-mac-routes" { return "LocalIpv4MacRoutes" }
    if yname == "local-ipv6-mac-routes" { return "LocalIpv6MacRoutes" }
    if yname == "es-global-mac-routes" { return "EsGlobalMacRoutes" }
    if yname == "remote-mac-routes" { return "RemoteMacRoutes" }
    if yname == "remote-soo-mac-routes" { return "RemoteSooMacRoutes" }
    if yname == "remote-ipv4-mac-routes" { return "RemoteIpv4MacRoutes" }
    if yname == "remote-ipv6-mac-routes" { return "RemoteIpv6MacRoutes" }
    if yname == "local-imcast-routes" { return "LocalImcastRoutes" }
    if yname == "remote-imcast-routes" { return "RemoteImcastRoutes" }
    if yname == "labels" { return "Labels" }
    if yname == "es-entries" { return "EsEntries" }
    if yname == "neighbor-entries" { return "NeighborEntries" }
    if yname == "local-ead-routes" { return "LocalEadRoutes" }
    if yname == "remote-ead-routes" { return "RemoteEadRoutes" }
    if yname == "global-source-mac" { return "GlobalSourceMac" }
    if yname == "peering-time" { return "PeeringTime" }
    if yname == "recovery-time" { return "RecoveryTime" }
    if yname == "mac-secure-move-count" { return "MacSecureMoveCount" }
    if yname == "mac-secure-move-interval" { return "MacSecureMoveInterval" }
    if yname == "mac-secure-freeze-time" { return "MacSecureFreezeTime" }
    if yname == "mac-secure-retry-count" { return "MacSecureRetryCount" }
    if yname == "cost-out" { return "CostOut" }
    if yname == "startup-cost-in-time" { return "StartupCostInTime" }
    if yname == "l2rib-throttle" { return "L2RibThrottle" }
    if yname == "logging-df-election-enabled" { return "LoggingDfElectionEnabled" }
    return ""
}

func (summary *Evpn_Nodes_Node_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Evpn_Nodes_Node_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Evpn_Nodes_Node_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Evpn_Nodes_Node_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = summary.RouterId
    leafs["as"] = summary.As
    leafs["ev-is"] = summary.EvIs
    leafs["local-mac-routes"] = summary.LocalMacRoutes
    leafs["local-ipv4-mac-routes"] = summary.LocalIpv4MacRoutes
    leafs["local-ipv6-mac-routes"] = summary.LocalIpv6MacRoutes
    leafs["es-global-mac-routes"] = summary.EsGlobalMacRoutes
    leafs["remote-mac-routes"] = summary.RemoteMacRoutes
    leafs["remote-soo-mac-routes"] = summary.RemoteSooMacRoutes
    leafs["remote-ipv4-mac-routes"] = summary.RemoteIpv4MacRoutes
    leafs["remote-ipv6-mac-routes"] = summary.RemoteIpv6MacRoutes
    leafs["local-imcast-routes"] = summary.LocalImcastRoutes
    leafs["remote-imcast-routes"] = summary.RemoteImcastRoutes
    leafs["labels"] = summary.Labels
    leafs["es-entries"] = summary.EsEntries
    leafs["neighbor-entries"] = summary.NeighborEntries
    leafs["local-ead-routes"] = summary.LocalEadRoutes
    leafs["remote-ead-routes"] = summary.RemoteEadRoutes
    leafs["global-source-mac"] = summary.GlobalSourceMac
    leafs["peering-time"] = summary.PeeringTime
    leafs["recovery-time"] = summary.RecoveryTime
    leafs["mac-secure-move-count"] = summary.MacSecureMoveCount
    leafs["mac-secure-move-interval"] = summary.MacSecureMoveInterval
    leafs["mac-secure-freeze-time"] = summary.MacSecureFreezeTime
    leafs["mac-secure-retry-count"] = summary.MacSecureRetryCount
    leafs["cost-out"] = summary.CostOut
    leafs["startup-cost-in-time"] = summary.StartupCostInTime
    leafs["l2rib-throttle"] = summary.L2RibThrottle
    leafs["logging-df-election-enabled"] = summary.LoggingDfElectionEnabled
    return leafs
}

func (summary *Evpn_Nodes_Node_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Evpn_Nodes_Node_Summary) GetYangName() string { return "summary" }

func (summary *Evpn_Nodes_Node_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Evpn_Nodes_Node_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Evpn_Nodes_Node_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Evpn_Nodes_Node_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Evpn_Nodes_Node_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Evpn_Nodes_Node_Summary) GetParentYangName() string { return "node" }

// Evpn_Nodes_Node_EviDetail
// L2VPN EVI Detail Table
type Evpn_Nodes_Node_EviDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info Elements.
    Elements Evpn_Nodes_Node_EviDetail_Elements

    // Container for all EVI detail info.
    EviChildren Evpn_Nodes_Node_EviDetail_EviChildren
}

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetFilter() yfilter.YFilter { return eviDetail.YFilter }

func (eviDetail *Evpn_Nodes_Node_EviDetail) SetFilter(yf yfilter.YFilter) { eviDetail.YFilter = yf }

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetGoName(yname string) string {
    if yname == "elements" { return "Elements" }
    if yname == "evi-children" { return "EviChildren" }
    return ""
}

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetSegmentPath() string {
    return "evi-detail"
}

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "elements" {
        return &eviDetail.Elements
    }
    if childYangName == "evi-children" {
        return &eviDetail.EviChildren
    }
    return nil
}

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["elements"] = &eviDetail.Elements
    children["evi-children"] = &eviDetail.EviChildren
    return children
}

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetBundleName() string { return "cisco_ios_xr" }

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetYangName() string { return "evi-detail" }

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eviDetail *Evpn_Nodes_Node_EviDetail) SetParent(parent types.Entity) { eviDetail.parent = parent }

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetParent() types.Entity { return eviDetail.parent }

func (eviDetail *Evpn_Nodes_Node_EviDetail) GetParentYangName() string { return "node" }

// Evpn_Nodes_Node_EviDetail_Elements
// EVI BGP RT Detail Info Elements
type Evpn_Nodes_Node_EviDetail_Elements struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info. The type is slice of
    // Evpn_Nodes_Node_EviDetail_Elements_Element.
    Element []Evpn_Nodes_Node_EviDetail_Elements_Element
}

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetFilter() yfilter.YFilter { return elements.YFilter }

func (elements *Evpn_Nodes_Node_EviDetail_Elements) SetFilter(yf yfilter.YFilter) { elements.YFilter = yf }

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetGoName(yname string) string {
    if yname == "element" { return "Element" }
    return ""
}

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetSegmentPath() string {
    return "elements"
}

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "element" {
        for _, c := range elements.Element {
            if elements.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_Elements_Element{}
        elements.Element = append(elements.Element, child)
        return &elements.Element[len(elements.Element)-1]
    }
    return nil
}

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range elements.Element {
        children[elements.Element[i].GetSegmentPath()] = &elements.Element[i]
    }
    return children
}

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetBundleName() string { return "cisco_ios_xr" }

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetYangName() string { return "elements" }

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (elements *Evpn_Nodes_Node_EviDetail_Elements) SetParent(parent types.Entity) { elements.parent = parent }

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetParent() types.Entity { return elements.parent }

func (elements *Evpn_Nodes_Node_EviDetail_Elements) GetParentYangName() string { return "evi-detail" }

// Evpn_Nodes_Node_EviDetail_Elements_Element
// EVI BGP RT Detail Info
type Evpn_Nodes_Node_EviDetail_Elements_Element struct {
    parent types.Entity
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
    Type interface{}

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

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetFilter() yfilter.YFilter { return element.YFilter }

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) SetFilter(yf yfilter.YFilter) { element.YFilter = yf }

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "description" { return "Description" }
    if yname == "bd-name" { return "BdName" }
    if yname == "type" { return "Type" }
    if yname == "unicast-label" { return "UnicastLabel" }
    if yname == "multicast-label" { return "MulticastLabel" }
    if yname == "cw-disable" { return "CwDisable" }
    if yname == "table-policy-name" { return "TablePolicyName" }
    if yname == "forward-class" { return "ForwardClass" }
    if yname == "rt-import-block-set" { return "RtImportBlockSet" }
    if yname == "rt-export-block-set" { return "RtExportBlockSet" }
    if yname == "advertise-mac" { return "AdvertiseMac" }
    if yname == "advertise-bvi-mac" { return "AdvertiseBviMac" }
    if yname == "aliasing-disabled" { return "AliasingDisabled" }
    if yname == "unknown-unicast-flooding-disabled" { return "UnknownUnicastFloodingDisabled" }
    if yname == "reoriginate-disabled" { return "ReoriginateDisabled" }
    if yname == "stitching" { return "Stitching" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "flow-label" { return "FlowLabel" }
    if yname == "rd-auto" { return "RdAuto" }
    if yname == "rd-configured" { return "RdConfigured" }
    if yname == "rt-auto" { return "RtAuto" }
    if yname == "rt-auto-stitching" { return "RtAutoStitching" }
    return ""
}

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetSegmentPath() string {
    return "element" + "[evi='" + fmt.Sprintf("%v", element.Evi) + "']"
}

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-label" {
        return &element.FlowLabel
    }
    if childYangName == "rd-auto" {
        return &element.RdAuto
    }
    if childYangName == "rd-configured" {
        return &element.RdConfigured
    }
    if childYangName == "rt-auto" {
        return &element.RtAuto
    }
    if childYangName == "rt-auto-stitching" {
        return &element.RtAutoStitching
    }
    return nil
}

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flow-label"] = &element.FlowLabel
    children["rd-auto"] = &element.RdAuto
    children["rd-configured"] = &element.RdConfigured
    children["rt-auto"] = &element.RtAuto
    children["rt-auto-stitching"] = &element.RtAutoStitching
    return children
}

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = element.Evi
    leafs["evi-xr"] = element.EviXr
    leafs["description"] = element.Description
    leafs["bd-name"] = element.BdName
    leafs["type"] = element.Type
    leafs["unicast-label"] = element.UnicastLabel
    leafs["multicast-label"] = element.MulticastLabel
    leafs["cw-disable"] = element.CwDisable
    leafs["table-policy-name"] = element.TablePolicyName
    leafs["forward-class"] = element.ForwardClass
    leafs["rt-import-block-set"] = element.RtImportBlockSet
    leafs["rt-export-block-set"] = element.RtExportBlockSet
    leafs["advertise-mac"] = element.AdvertiseMac
    leafs["advertise-bvi-mac"] = element.AdvertiseBviMac
    leafs["aliasing-disabled"] = element.AliasingDisabled
    leafs["unknown-unicast-flooding-disabled"] = element.UnknownUnicastFloodingDisabled
    leafs["reoriginate-disabled"] = element.ReoriginateDisabled
    leafs["stitching"] = element.Stitching
    leafs["encapsulation"] = element.Encapsulation
    return leafs
}

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetBundleName() string { return "cisco_ios_xr" }

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetYangName() string { return "element" }

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) SetParent(parent types.Entity) { element.parent = parent }

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetParent() types.Entity { return element.parent }

func (element *Evpn_Nodes_Node_EviDetail_Elements_Element) GetParentYangName() string { return "elements" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel
// Flow Label Information
type Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Static flow label. The type is bool.
    StaticFlowLabel interface{}

    // Globally configured flow label. The type is bool.
    GlobalFlowLabel interface{}
}

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetFilter() yfilter.YFilter { return flowLabel.YFilter }

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) SetFilter(yf yfilter.YFilter) { flowLabel.YFilter = yf }

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetGoName(yname string) string {
    if yname == "static-flow-label" { return "StaticFlowLabel" }
    if yname == "global-flow-label" { return "GlobalFlowLabel" }
    return ""
}

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetSegmentPath() string {
    return "flow-label"
}

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["static-flow-label"] = flowLabel.StaticFlowLabel
    leafs["global-flow-label"] = flowLabel.GlobalFlowLabel
    return leafs
}

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetBundleName() string { return "cisco_ios_xr" }

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetYangName() string { return "flow-label" }

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) SetParent(parent types.Entity) { flowLabel.parent = parent }

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetParent() types.Entity { return flowLabel.parent }

func (flowLabel *Evpn_Nodes_Node_EviDetail_Elements_Element_FlowLabel) GetParentYangName() string { return "element" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto
// Automatic Route Distingtuisher
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto struct {
    parent types.Entity
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

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetFilter() yfilter.YFilter { return rdAuto.YFilter }

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) SetFilter(yf yfilter.YFilter) { rdAuto.YFilter = yf }

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetGoName(yname string) string {
    if yname == "rd" { return "Rd" }
    if yname == "auto" { return "Auto" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    return ""
}

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetSegmentPath() string {
    return "rd-auto"
}

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto" {
        return &rdAuto.Auto
    }
    if childYangName == "two-byte-as" {
        return &rdAuto.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rdAuto.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rdAuto.V4Addr
    }
    return nil
}

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto"] = &rdAuto.Auto
    children["two-byte-as"] = &rdAuto.TwoByteAs
    children["four-byte-as"] = &rdAuto.FourByteAs
    children["v4-addr"] = &rdAuto.V4Addr
    return children
}

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rd"] = rdAuto.Rd
    return leafs
}

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetBundleName() string { return "cisco_ios_xr" }

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetYangName() string { return "rd-auto" }

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) SetParent(parent types.Entity) { rdAuto.parent = parent }

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetParent() types.Entity { return rdAuto.parent }

func (rdAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto) GetParentYangName() string { return "element" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto
// auto
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetFilter() yfilter.YFilter { return auto.YFilter }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) SetFilter(yf yfilter.YFilter) { auto.YFilter = yf }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "auto-index" { return "AutoIndex" }
    return ""
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetSegmentPath() string {
    return "auto"
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = auto.RouterId
    leafs["auto-index"] = auto.AutoIndex
    return leafs
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetBundleName() string { return "cisco_ios_xr" }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetYangName() string { return "auto" }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) SetParent(parent types.Entity) { auto.parent = parent }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetParent() types.Entity { return auto.parent }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_Auto) GetParentYangName() string { return "rd-auto" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetParentYangName() string { return "rd-auto" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_FourByteAs) GetParentYangName() string { return "rd-auto" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdAuto_V4Addr) GetParentYangName() string { return "rd-auto" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured
// Configured Route Distinguisher
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured struct {
    parent types.Entity
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

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetFilter() yfilter.YFilter { return rdConfigured.YFilter }

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) SetFilter(yf yfilter.YFilter) { rdConfigured.YFilter = yf }

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetGoName(yname string) string {
    if yname == "rd" { return "Rd" }
    if yname == "auto" { return "Auto" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    return ""
}

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetSegmentPath() string {
    return "rd-configured"
}

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto" {
        return &rdConfigured.Auto
    }
    if childYangName == "two-byte-as" {
        return &rdConfigured.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rdConfigured.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rdConfigured.V4Addr
    }
    return nil
}

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto"] = &rdConfigured.Auto
    children["two-byte-as"] = &rdConfigured.TwoByteAs
    children["four-byte-as"] = &rdConfigured.FourByteAs
    children["v4-addr"] = &rdConfigured.V4Addr
    return children
}

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rd"] = rdConfigured.Rd
    return leafs
}

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetBundleName() string { return "cisco_ios_xr" }

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetYangName() string { return "rd-configured" }

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) SetParent(parent types.Entity) { rdConfigured.parent = parent }

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetParent() types.Entity { return rdConfigured.parent }

func (rdConfigured *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured) GetParentYangName() string { return "element" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto
// auto
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetFilter() yfilter.YFilter { return auto.YFilter }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) SetFilter(yf yfilter.YFilter) { auto.YFilter = yf }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "auto-index" { return "AutoIndex" }
    return ""
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetSegmentPath() string {
    return "auto"
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = auto.RouterId
    leafs["auto-index"] = auto.AutoIndex
    return leafs
}

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetBundleName() string { return "cisco_ios_xr" }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetYangName() string { return "auto" }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) SetParent(parent types.Entity) { auto.parent = parent }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetParent() types.Entity { return auto.parent }

func (auto *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_Auto) GetParentYangName() string { return "rd-configured" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetParentYangName() string { return "rd-configured" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetParentYangName() string { return "rd-configured" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RdConfigured_V4Addr) GetParentYangName() string { return "rd-configured" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto
// Automatic Route Target
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto struct {
    parent types.Entity
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

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetFilter() yfilter.YFilter { return rtAuto.YFilter }

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) SetFilter(yf yfilter.YFilter) { rtAuto.YFilter = yf }

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetGoName(yname string) string {
    if yname == "rt" { return "Rt" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    if yname == "es-import" { return "EsImport" }
    return ""
}

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetSegmentPath() string {
    return "rt-auto"
}

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "two-byte-as" {
        return &rtAuto.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rtAuto.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rtAuto.V4Addr
    }
    if childYangName == "es-import" {
        return &rtAuto.EsImport
    }
    return nil
}

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["two-byte-as"] = &rtAuto.TwoByteAs
    children["four-byte-as"] = &rtAuto.FourByteAs
    children["v4-addr"] = &rtAuto.V4Addr
    children["es-import"] = &rtAuto.EsImport
    return children
}

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt"] = rtAuto.Rt
    return leafs
}

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetBundleName() string { return "cisco_ios_xr" }

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetYangName() string { return "rt-auto" }

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) SetParent(parent types.Entity) { rtAuto.parent = parent }

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetParent() types.Entity { return rtAuto.parent }

func (rtAuto *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto) GetParentYangName() string { return "element" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetParentYangName() string { return "rt-auto" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_FourByteAs) GetParentYangName() string { return "rt-auto" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_V4Addr) GetParentYangName() string { return "rt-auto" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport
// es import
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetFilter() yfilter.YFilter { return esImport.YFilter }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) SetFilter(yf yfilter.YFilter) { esImport.YFilter = yf }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetGoName(yname string) string {
    if yname == "high-bytes" { return "HighBytes" }
    if yname == "low-bytes" { return "LowBytes" }
    return ""
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetSegmentPath() string {
    return "es-import"
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-bytes"] = esImport.HighBytes
    leafs["low-bytes"] = esImport.LowBytes
    return leafs
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetBundleName() string { return "cisco_ios_xr" }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetYangName() string { return "es-import" }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) SetParent(parent types.Entity) { esImport.parent = parent }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetParent() types.Entity { return esImport.parent }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAuto_EsImport) GetParentYangName() string { return "rt-auto" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching
// Automatic Route Target Stitching
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching struct {
    parent types.Entity
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

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetFilter() yfilter.YFilter { return rtAutoStitching.YFilter }

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) SetFilter(yf yfilter.YFilter) { rtAutoStitching.YFilter = yf }

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetGoName(yname string) string {
    if yname == "rt" { return "Rt" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    if yname == "es-import" { return "EsImport" }
    return ""
}

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetSegmentPath() string {
    return "rt-auto-stitching"
}

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "two-byte-as" {
        return &rtAutoStitching.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rtAutoStitching.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rtAutoStitching.V4Addr
    }
    if childYangName == "es-import" {
        return &rtAutoStitching.EsImport
    }
    return nil
}

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["two-byte-as"] = &rtAutoStitching.TwoByteAs
    children["four-byte-as"] = &rtAutoStitching.FourByteAs
    children["v4-addr"] = &rtAutoStitching.V4Addr
    children["es-import"] = &rtAutoStitching.EsImport
    return children
}

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt"] = rtAutoStitching.Rt
    return leafs
}

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetBundleName() string { return "cisco_ios_xr" }

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetYangName() string { return "rt-auto-stitching" }

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) SetParent(parent types.Entity) { rtAutoStitching.parent = parent }

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetParent() types.Entity { return rtAutoStitching.parent }

func (rtAutoStitching *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching) GetParentYangName() string { return "element" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport
// es import
type Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetFilter() yfilter.YFilter { return esImport.YFilter }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) SetFilter(yf yfilter.YFilter) { esImport.YFilter = yf }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetGoName(yname string) string {
    if yname == "high-bytes" { return "HighBytes" }
    if yname == "low-bytes" { return "LowBytes" }
    return ""
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetSegmentPath() string {
    return "es-import"
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-bytes"] = esImport.HighBytes
    leafs["low-bytes"] = esImport.LowBytes
    return leafs
}

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetBundleName() string { return "cisco_ios_xr" }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetYangName() string { return "es-import" }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) SetParent(parent types.Entity) { esImport.parent = parent }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetParent() types.Entity { return esImport.parent }

func (esImport *Evpn_Nodes_Node_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Nodes_Node_EviDetail_EviChildren
// Container for all EVI detail info
type Evpn_Nodes_Node_EviDetail_EviChildren struct {
    parent types.Entity
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

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetFilter() yfilter.YFilter { return eviChildren.YFilter }

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) SetFilter(yf yfilter.YFilter) { eviChildren.YFilter = yf }

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetGoName(yname string) string {
    if yname == "neighbors" { return "Neighbors" }
    if yname == "ethernet-auto-discoveries" { return "EthernetAutoDiscoveries" }
    if yname == "inclusive-multicasts" { return "InclusiveMulticasts" }
    if yname == "route-targets" { return "RouteTargets" }
    if yname == "macs" { return "Macs" }
    return ""
}

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetSegmentPath() string {
    return "evi-children"
}

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbors" {
        return &eviChildren.Neighbors
    }
    if childYangName == "ethernet-auto-discoveries" {
        return &eviChildren.EthernetAutoDiscoveries
    }
    if childYangName == "inclusive-multicasts" {
        return &eviChildren.InclusiveMulticasts
    }
    if childYangName == "route-targets" {
        return &eviChildren.RouteTargets
    }
    if childYangName == "macs" {
        return &eviChildren.Macs
    }
    return nil
}

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbors"] = &eviChildren.Neighbors
    children["ethernet-auto-discoveries"] = &eviChildren.EthernetAutoDiscoveries
    children["inclusive-multicasts"] = &eviChildren.InclusiveMulticasts
    children["route-targets"] = &eviChildren.RouteTargets
    children["macs"] = &eviChildren.Macs
    return children
}

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetBundleName() string { return "cisco_ios_xr" }

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetYangName() string { return "evi-children" }

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) SetParent(parent types.Entity) { eviChildren.parent = parent }

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetParent() types.Entity { return eviChildren.parent }

func (eviChildren *Evpn_Nodes_Node_EviDetail_EviChildren) GetParentYangName() string { return "evi-detail" }

// Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors
// EVPN Neighbor table
type Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Neighbor table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor.
    Neighbor []Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor
}

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors) GetParentYangName() string { return "evi-children" }

// Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor
// EVPN Neighbor table
type Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Neighbor IP. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}
}

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "neighbor-ip" { return "NeighborIp" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = neighbor.Evi
    leafs["neighbor-ip"] = neighbor.NeighborIp
    leafs["evi-xr"] = neighbor.EviXr
    leafs["neighbor"] = neighbor.Neighbor
    return leafs
}

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Evpn_Nodes_Node_EviDetail_EviChildren_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries
// EVPN Ethernet Auto-Discovery table
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Ethernet Auto-Discovery Entry. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery.
    EthernetAutoDiscovery []Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
}

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetFilter() yfilter.YFilter { return ethernetAutoDiscoveries.YFilter }

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) SetFilter(yf yfilter.YFilter) { ethernetAutoDiscoveries.YFilter = yf }

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetGoName(yname string) string {
    if yname == "ethernet-auto-discovery" { return "EthernetAutoDiscovery" }
    return ""
}

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetSegmentPath() string {
    return "ethernet-auto-discoveries"
}

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-auto-discovery" {
        for _, c := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
            if ethernetAutoDiscoveries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery{}
        ethernetAutoDiscoveries.EthernetAutoDiscovery = append(ethernetAutoDiscoveries.EthernetAutoDiscovery, child)
        return &ethernetAutoDiscoveries.EthernetAutoDiscovery[len(ethernetAutoDiscoveries.EthernetAutoDiscovery)-1]
    }
    return nil
}

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
        children[ethernetAutoDiscoveries.EthernetAutoDiscovery[i].GetSegmentPath()] = &ethernetAutoDiscoveries.EthernetAutoDiscovery[i]
    }
    return children
}

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetYangName() string { return "ethernet-auto-discoveries" }

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) SetParent(parent types.Entity) { ethernetAutoDiscoveries.parent = parent }

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetParent() types.Entity { return ethernetAutoDiscoveries.parent }

func (ethernetAutoDiscoveries *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries) GetParentYangName() string { return "evi-children" }

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
// EVPN Ethernet Auto-Discovery Entry
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

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

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnid interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Local nexthop IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetFilter() yfilter.YFilter { return ethernetAutoDiscovery.YFilter }

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) SetFilter(yf yfilter.YFilter) { ethernetAutoDiscovery.YFilter = yf }

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "esi1" { return "Esi1" }
    if yname == "esi2" { return "Esi2" }
    if yname == "esi3" { return "Esi3" }
    if yname == "esi4" { return "Esi4" }
    if yname == "esi5" { return "Esi5" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    if yname == "ethernet-vpnid" { return "EthernetVpnid" }
    if yname == "type" { return "Type" }
    if yname == "ethernet-tag-xr" { return "EthernetTagXr" }
    if yname == "local-next-hop" { return "LocalNextHop" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "is-local-ead" { return "IsLocalEad" }
    if yname == "encap" { return "Encap" }
    if yname == "redundancy-single-active" { return "RedundancySingleActive" }
    if yname == "num-paths" { return "NumPaths" }
    if yname == "ethernet-segment-identifier" { return "EthernetSegmentIdentifier" }
    if yname == "path-buffer" { return "PathBuffer" }
    return ""
}

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetSegmentPath() string {
    return "ethernet-auto-discovery"
}

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-identifier" {
        for _, c := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
            if ethernetAutoDiscovery.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier{}
        ethernetAutoDiscovery.EthernetSegmentIdentifier = append(ethernetAutoDiscovery.EthernetSegmentIdentifier, child)
        return &ethernetAutoDiscovery.EthernetSegmentIdentifier[len(ethernetAutoDiscovery.EthernetSegmentIdentifier)-1]
    }
    if childYangName == "path-buffer" {
        for _, c := range ethernetAutoDiscovery.PathBuffer {
            if ethernetAutoDiscovery.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer{}
        ethernetAutoDiscovery.PathBuffer = append(ethernetAutoDiscovery.PathBuffer, child)
        return &ethernetAutoDiscovery.PathBuffer[len(ethernetAutoDiscovery.PathBuffer)-1]
    }
    return nil
}

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
        children[ethernetAutoDiscovery.EthernetSegmentIdentifier[i].GetSegmentPath()] = &ethernetAutoDiscovery.EthernetSegmentIdentifier[i]
    }
    for i := range ethernetAutoDiscovery.PathBuffer {
        children[ethernetAutoDiscovery.PathBuffer[i].GetSegmentPath()] = &ethernetAutoDiscovery.PathBuffer[i]
    }
    return children
}

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = ethernetAutoDiscovery.Evi
    leafs["esi1"] = ethernetAutoDiscovery.Esi1
    leafs["esi2"] = ethernetAutoDiscovery.Esi2
    leafs["esi3"] = ethernetAutoDiscovery.Esi3
    leafs["esi4"] = ethernetAutoDiscovery.Esi4
    leafs["esi5"] = ethernetAutoDiscovery.Esi5
    leafs["ethernet-tag"] = ethernetAutoDiscovery.EthernetTag
    leafs["ethernet-vpnid"] = ethernetAutoDiscovery.EthernetVpnid
    leafs["type"] = ethernetAutoDiscovery.Type
    leafs["ethernet-tag-xr"] = ethernetAutoDiscovery.EthernetTagXr
    leafs["local-next-hop"] = ethernetAutoDiscovery.LocalNextHop
    leafs["local-label"] = ethernetAutoDiscovery.LocalLabel
    leafs["is-local-ead"] = ethernetAutoDiscovery.IsLocalEad
    leafs["encap"] = ethernetAutoDiscovery.Encap
    leafs["redundancy-single-active"] = ethernetAutoDiscovery.RedundancySingleActive
    leafs["num-paths"] = ethernetAutoDiscovery.NumPaths
    return leafs
}

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetYangName() string { return "ethernet-auto-discovery" }

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) SetParent(parent types.Entity) { ethernetAutoDiscovery.parent = parent }

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetParent() types.Entity { return ethernetAutoDiscovery.parent }

func (ethernetAutoDiscovery *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetParentYangName() string { return "ethernet-auto-discoveries" }

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return ethernetSegmentIdentifier.YFilter }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { ethernetSegmentIdentifier.YFilter = yf }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetSegmentPath() string {
    return "ethernet-segment-identifier"
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = ethernetSegmentIdentifier.Entry
    return leafs
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetYangName() string { return "ethernet-segment-identifier" }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) SetParent(parent types.Entity) { ethernetSegmentIdentifier.parent = parent }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetParent() types.Entity { return ethernetSegmentIdentifier.parent }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetParentYangName() string { return "ethernet-auto-discovery" }

// Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
// Path List Buffer
type Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetFilter() yfilter.YFilter { return pathBuffer.YFilter }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) SetFilter(yf yfilter.YFilter) { pathBuffer.YFilter = yf }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "output-label" { return "OutputLabel" }
    if yname == "srte-tunnel" { return "SrteTunnel" }
    return ""
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetSegmentPath() string {
    return "path-buffer"
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = pathBuffer.NextHop
    leafs["output-label"] = pathBuffer.OutputLabel
    leafs["srte-tunnel"] = pathBuffer.SrteTunnel
    return leafs
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetYangName() string { return "path-buffer" }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) SetParent(parent types.Entity) { pathBuffer.parent = parent }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetParent() types.Entity { return pathBuffer.parent }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetParentYangName() string { return "ethernet-auto-discovery" }

// Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts
// L2VPN EVPN IMCAST table
type Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN IMCAST table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast.
    InclusiveMulticast []Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
}

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetFilter() yfilter.YFilter { return inclusiveMulticasts.YFilter }

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) SetFilter(yf yfilter.YFilter) { inclusiveMulticasts.YFilter = yf }

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetGoName(yname string) string {
    if yname == "inclusive-multicast" { return "InclusiveMulticast" }
    return ""
}

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetSegmentPath() string {
    return "inclusive-multicasts"
}

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inclusive-multicast" {
        for _, c := range inclusiveMulticasts.InclusiveMulticast {
            if inclusiveMulticasts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast{}
        inclusiveMulticasts.InclusiveMulticast = append(inclusiveMulticasts.InclusiveMulticast, child)
        return &inclusiveMulticasts.InclusiveMulticast[len(inclusiveMulticasts.InclusiveMulticast)-1]
    }
    return nil
}

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inclusiveMulticasts.InclusiveMulticast {
        children[inclusiveMulticasts.InclusiveMulticast[i].GetSegmentPath()] = &inclusiveMulticasts.InclusiveMulticast[i]
    }
    return children
}

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetBundleName() string { return "cisco_ios_xr" }

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetYangName() string { return "inclusive-multicasts" }

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) SetParent(parent types.Entity) { inclusiveMulticasts.parent = parent }

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetParent() types.Entity { return inclusiveMulticasts.parent }

func (inclusiveMulticasts *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts) GetParentYangName() string { return "evi-children" }

// Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
// L2VPN EVPN IMCAST table
type Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag. The type is interface{} with range: -2147483648..2147483647.
    EthernetTag interface{}

    // Originating IP. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Originating IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIpXr interface{}

    // IP of nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetFilter() yfilter.YFilter { return inclusiveMulticast.YFilter }

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) SetFilter(yf yfilter.YFilter) { inclusiveMulticast.YFilter = yf }

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    if yname == "originating-ip" { return "OriginatingIp" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "ethernet-tag-xr" { return "EthernetTagXr" }
    if yname == "originating-ip-xr" { return "OriginatingIpXr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "output-label" { return "OutputLabel" }
    if yname == "is-local-entry" { return "IsLocalEntry" }
    if yname == "is-proxy-entry" { return "IsProxyEntry" }
    if yname == "encap-type" { return "EncapType" }
    return ""
}

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetSegmentPath() string {
    return "inclusive-multicast"
}

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = inclusiveMulticast.Evi
    leafs["ethernet-tag"] = inclusiveMulticast.EthernetTag
    leafs["originating-ip"] = inclusiveMulticast.OriginatingIp
    leafs["evi-xr"] = inclusiveMulticast.EviXr
    leafs["ethernet-tag-xr"] = inclusiveMulticast.EthernetTagXr
    leafs["originating-ip-xr"] = inclusiveMulticast.OriginatingIpXr
    leafs["next-hop"] = inclusiveMulticast.NextHop
    leafs["output-label"] = inclusiveMulticast.OutputLabel
    leafs["is-local-entry"] = inclusiveMulticast.IsLocalEntry
    leafs["is-proxy-entry"] = inclusiveMulticast.IsProxyEntry
    leafs["encap-type"] = inclusiveMulticast.EncapType
    return leafs
}

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetBundleName() string { return "cisco_ios_xr" }

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetYangName() string { return "inclusive-multicast" }

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) SetParent(parent types.Entity) { inclusiveMulticast.parent = parent }

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetParent() types.Entity { return inclusiveMulticast.parent }

func (inclusiveMulticast *Evpn_Nodes_Node_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetParentYangName() string { return "inclusive-multicasts" }

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets
// L2VPN EVPN EVI RT Child Table
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI RT Table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget.
    RouteTarget []Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget
}

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetFilter() yfilter.YFilter { return routeTargets.YFilter }

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) SetFilter(yf yfilter.YFilter) { routeTargets.YFilter = yf }

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetGoName(yname string) string {
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetSegmentPath() string {
    return "route-targets"
}

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        for _, c := range routeTargets.RouteTarget {
            if routeTargets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget{}
        routeTargets.RouteTarget = append(routeTargets.RouteTarget, child)
        return &routeTargets.RouteTarget[len(routeTargets.RouteTarget)-1]
    }
    return nil
}

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routeTargets.RouteTarget {
        children[routeTargets.RouteTarget[i].GetSegmentPath()] = &routeTargets.RouteTarget[i]
    }
    return children
}

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetBundleName() string { return "cisco_ios_xr" }

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetYangName() string { return "route-targets" }

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) SetParent(parent types.Entity) { routeTargets.parent = parent }

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetParent() types.Entity { return routeTargets.parent }

func (routeTargets *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets) GetParentYangName() string { return "evi-children" }

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget
// L2VPN EVPN EVI RT Table
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Role of the route target. The type is BgpRouteTargetRole.
    Role interface{}

    // Type of the route target. The type is BgpRouteTarget.
    Type interface{}

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

    // Bridge Domain Name. The type is string.
    BdName interface{}

    // VPN ID. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // RT Role. The type is L2vpnAdRtRole.
    RouteTargetRole interface{}

    // RT Stitching. The type is bool.
    RouteTargetStitching interface{}

    // Route Target.
    RouteTarget Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "role" { return "Role" }
    if yname == "type" { return "Type" }
    if yname == "format" { return "Format" }
    if yname == "as" { return "As" }
    if yname == "as-index" { return "AsIndex" }
    if yname == "addr-index" { return "AddrIndex" }
    if yname == "address" { return "Address" }
    if yname == "bd-name" { return "BdName" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "route-target-role" { return "RouteTargetRole" }
    if yname == "route-target-stitching" { return "RouteTargetStitching" }
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetSegmentPath() string {
    return "route-target"
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        return &routeTarget.RouteTarget
    }
    return nil
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route-target"] = &routeTarget.RouteTarget
    return children
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = routeTarget.Evi
    leafs["role"] = routeTarget.Role
    leafs["type"] = routeTarget.Type
    leafs["format"] = routeTarget.Format
    leafs["as"] = routeTarget.As
    leafs["as-index"] = routeTarget.AsIndex
    leafs["addr-index"] = routeTarget.AddrIndex
    leafs["address"] = routeTarget.Address
    leafs["bd-name"] = routeTarget.BdName
    leafs["evi-xr"] = routeTarget.EviXr
    leafs["route-target-role"] = routeTarget.RouteTargetRole
    leafs["route-target-stitching"] = routeTarget.RouteTargetStitching
    return leafs
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget) GetParentYangName() string { return "route-targets" }

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
// Route Target
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget struct {
    parent types.Entity
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

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetGoName(yname string) string {
    if yname == "rt" { return "Rt" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    if yname == "es-import" { return "EsImport" }
    return ""
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetSegmentPath() string {
    return "route-target"
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "two-byte-as" {
        return &routeTarget.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &routeTarget.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &routeTarget.V4Addr
    }
    if childYangName == "es-import" {
        return &routeTarget.EsImport
    }
    return nil
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["two-byte-as"] = &routeTarget.TwoByteAs
    children["four-byte-as"] = &routeTarget.FourByteAs
    children["v4-addr"] = &routeTarget.V4Addr
    children["es-import"] = &routeTarget.EsImport
    return children
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt"] = routeTarget.Rt
    return leafs
}

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetParentYangName() string { return "route-target" }

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs
// two byte as
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetParentYangName() string { return "route-target" }

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs
// four byte as
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetParentYangName() string { return "route-target" }

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr
// v4 addr
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetParentYangName() string { return "route-target" }

// Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport
// es import
type Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetFilter() yfilter.YFilter { return esImport.YFilter }

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) SetFilter(yf yfilter.YFilter) { esImport.YFilter = yf }

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetGoName(yname string) string {
    if yname == "high-bytes" { return "HighBytes" }
    if yname == "low-bytes" { return "LowBytes" }
    return ""
}

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetSegmentPath() string {
    return "es-import"
}

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-bytes"] = esImport.HighBytes
    leafs["low-bytes"] = esImport.LowBytes
    return leafs
}

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetBundleName() string { return "cisco_ios_xr" }

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetYangName() string { return "es-import" }

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) SetParent(parent types.Entity) { esImport.parent = parent }

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetParent() types.Entity { return esImport.parent }

func (esImport *Evpn_Nodes_Node_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetParentYangName() string { return "route-target" }

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs
// L2VPN EVPN EVI MAC table
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN MAC table. The type is slice of
    // Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac.
    Mac []Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac
}

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetFilter() yfilter.YFilter { return macs.YFilter }

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) SetFilter(yf yfilter.YFilter) { macs.YFilter = yf }

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetGoName(yname string) string {
    if yname == "mac" { return "Mac" }
    return ""
}

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetSegmentPath() string {
    return "macs"
}

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac" {
        for _, c := range macs.Mac {
            if macs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac{}
        macs.Mac = append(macs.Mac, child)
        return &macs.Mac[len(macs.Mac)-1]
    }
    return nil
}

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macs.Mac {
        children[macs.Mac[i].GetSegmentPath()] = &macs.Mac[i]
    }
    return children
}

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetBundleName() string { return "cisco_ios_xr" }

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetYangName() string { return "macs" }

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) SetParent(parent types.Entity) { macs.parent = parent }

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetParent() types.Entity { return macs.parent }

func (macs *Evpn_Nodes_Node_EviDetail_EviChildren_Macs) GetParentYangName() string { return "evi-children" }

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac
// L2VPN EVPN MAC table
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
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

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "ethernet-tag-xr" { return "EthernetTagXr" }
    if yname == "mac-address-xr" { return "MacAddressXr" }
    if yname == "ip-address-xr" { return "IpAddressXr" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "num-paths" { return "NumPaths" }
    if yname == "is-local-mac" { return "IsLocalMac" }
    if yname == "is-proxy-entry" { return "IsProxyEntry" }
    if yname == "is-remote-mac" { return "IsRemoteMac" }
    if yname == "soo-nexthop" { return "SooNexthop" }
    if yname == "ipnh-address" { return "IpnhAddress" }
    if yname == "esi-port-key" { return "EsiPortKey" }
    if yname == "local-encap-type" { return "LocalEncapType" }
    if yname == "remote-encap-type" { return "RemoteEncapType" }
    if yname == "learned-bridge-port-name" { return "LearnedBridgePortName" }
    if yname == "local-seq-id" { return "LocalSeqId" }
    if yname == "remote-seq-id" { return "RemoteSeqId" }
    if yname == "local-l3-label" { return "LocalL3Label" }
    if yname == "router-mac-address" { return "RouterMacAddress" }
    if yname == "mac-flush-requested" { return "MacFlushRequested" }
    if yname == "mac-flush-received" { return "MacFlushReceived" }
    if yname == "internal-label" { return "InternalLabel" }
    if yname == "resolved" { return "Resolved" }
    if yname == "local-is-static" { return "LocalIsStatic" }
    if yname == "remote-is-static" { return "RemoteIsStatic" }
    if yname == "local-ethernet-segment-identifier" { return "LocalEthernetSegmentIdentifier" }
    if yname == "remote-ethernet-segment-identifier" { return "RemoteEthernetSegmentIdentifier" }
    if yname == "path-buffer" { return "PathBuffer" }
    return ""
}

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-ethernet-segment-identifier" {
        for _, c := range mac.LocalEthernetSegmentIdentifier {
            if mac.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier{}
        mac.LocalEthernetSegmentIdentifier = append(mac.LocalEthernetSegmentIdentifier, child)
        return &mac.LocalEthernetSegmentIdentifier[len(mac.LocalEthernetSegmentIdentifier)-1]
    }
    if childYangName == "remote-ethernet-segment-identifier" {
        for _, c := range mac.RemoteEthernetSegmentIdentifier {
            if mac.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier{}
        mac.RemoteEthernetSegmentIdentifier = append(mac.RemoteEthernetSegmentIdentifier, child)
        return &mac.RemoteEthernetSegmentIdentifier[len(mac.RemoteEthernetSegmentIdentifier)-1]
    }
    if childYangName == "path-buffer" {
        for _, c := range mac.PathBuffer {
            if mac.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer{}
        mac.PathBuffer = append(mac.PathBuffer, child)
        return &mac.PathBuffer[len(mac.PathBuffer)-1]
    }
    return nil
}

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mac.LocalEthernetSegmentIdentifier {
        children[mac.LocalEthernetSegmentIdentifier[i].GetSegmentPath()] = &mac.LocalEthernetSegmentIdentifier[i]
    }
    for i := range mac.RemoteEthernetSegmentIdentifier {
        children[mac.RemoteEthernetSegmentIdentifier[i].GetSegmentPath()] = &mac.RemoteEthernetSegmentIdentifier[i]
    }
    for i := range mac.PathBuffer {
        children[mac.PathBuffer[i].GetSegmentPath()] = &mac.PathBuffer[i]
    }
    return children
}

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = mac.Evi
    leafs["ethernet-tag"] = mac.EthernetTag
    leafs["mac-address"] = mac.MacAddress
    leafs["ip-address"] = mac.IpAddress
    leafs["ethernet-tag-xr"] = mac.EthernetTagXr
    leafs["mac-address-xr"] = mac.MacAddressXr
    leafs["ip-address-xr"] = mac.IpAddressXr
    leafs["local-label"] = mac.LocalLabel
    leafs["num-paths"] = mac.NumPaths
    leafs["is-local-mac"] = mac.IsLocalMac
    leafs["is-proxy-entry"] = mac.IsProxyEntry
    leafs["is-remote-mac"] = mac.IsRemoteMac
    leafs["soo-nexthop"] = mac.SooNexthop
    leafs["ipnh-address"] = mac.IpnhAddress
    leafs["esi-port-key"] = mac.EsiPortKey
    leafs["local-encap-type"] = mac.LocalEncapType
    leafs["remote-encap-type"] = mac.RemoteEncapType
    leafs["learned-bridge-port-name"] = mac.LearnedBridgePortName
    leafs["local-seq-id"] = mac.LocalSeqId
    leafs["remote-seq-id"] = mac.RemoteSeqId
    leafs["local-l3-label"] = mac.LocalL3Label
    leafs["router-mac-address"] = mac.RouterMacAddress
    leafs["mac-flush-requested"] = mac.MacFlushRequested
    leafs["mac-flush-received"] = mac.MacFlushReceived
    leafs["internal-label"] = mac.InternalLabel
    leafs["resolved"] = mac.Resolved
    leafs["local-is-static"] = mac.LocalIsStatic
    leafs["remote-is-static"] = mac.RemoteIsStatic
    return leafs
}

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetYangName() string { return "mac" }

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetParent() types.Entity { return mac.parent }

func (mac *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac) GetParentYangName() string { return "macs" }

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier
// Local Ethernet Segment id
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return localEthernetSegmentIdentifier.YFilter }

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { localEthernetSegmentIdentifier.YFilter = yf }

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetSegmentPath() string {
    return "local-ethernet-segment-identifier"
}

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = localEthernetSegmentIdentifier.Entry
    return leafs
}

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetYangName() string { return "local-ethernet-segment-identifier" }

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) SetParent(parent types.Entity) { localEthernetSegmentIdentifier.parent = parent }

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetParent() types.Entity { return localEthernetSegmentIdentifier.parent }

func (localEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetParentYangName() string { return "mac" }

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier
// Remote Ethernet Segment id
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return remoteEthernetSegmentIdentifier.YFilter }

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { remoteEthernetSegmentIdentifier.YFilter = yf }

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetSegmentPath() string {
    return "remote-ethernet-segment-identifier"
}

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = remoteEthernetSegmentIdentifier.Entry
    return leafs
}

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetYangName() string { return "remote-ethernet-segment-identifier" }

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) SetParent(parent types.Entity) { remoteEthernetSegmentIdentifier.parent = parent }

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetParent() types.Entity { return remoteEthernetSegmentIdentifier.parent }

func (remoteEthernetSegmentIdentifier *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetParentYangName() string { return "mac" }

// Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer
// Path List Buffer
type Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetFilter() yfilter.YFilter { return pathBuffer.YFilter }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) SetFilter(yf yfilter.YFilter) { pathBuffer.YFilter = yf }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "output-label" { return "OutputLabel" }
    if yname == "srte-tunnel" { return "SrteTunnel" }
    return ""
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetSegmentPath() string {
    return "path-buffer"
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = pathBuffer.NextHop
    leafs["output-label"] = pathBuffer.OutputLabel
    leafs["srte-tunnel"] = pathBuffer.SrteTunnel
    return leafs
}

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetYangName() string { return "path-buffer" }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) SetParent(parent types.Entity) { pathBuffer.parent = parent }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetParent() types.Entity { return pathBuffer.parent }

func (pathBuffer *Evpn_Nodes_Node_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetParentYangName() string { return "mac" }

// Evpn_Nodes_Node_EthernetSegments
// EVPN Ethernet-Segment Table
type Evpn_Nodes_Node_EthernetSegments struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Ethernet-Segment Entry. The type is slice of
    // Evpn_Nodes_Node_EthernetSegments_EthernetSegment.
    EthernetSegment []Evpn_Nodes_Node_EthernetSegments_EthernetSegment
}

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetFilter() yfilter.YFilter { return ethernetSegments.YFilter }

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) SetFilter(yf yfilter.YFilter) { ethernetSegments.YFilter = yf }

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetGoName(yname string) string {
    if yname == "ethernet-segment" { return "EthernetSegment" }
    return ""
}

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetSegmentPath() string {
    return "ethernet-segments"
}

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment" {
        for _, c := range ethernetSegments.EthernetSegment {
            if ethernetSegments.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment{}
        ethernetSegments.EthernetSegment = append(ethernetSegments.EthernetSegment, child)
        return &ethernetSegments.EthernetSegment[len(ethernetSegments.EthernetSegment)-1]
    }
    return nil
}

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetSegments.EthernetSegment {
        children[ethernetSegments.EthernetSegment[i].GetSegmentPath()] = &ethernetSegments.EthernetSegment[i]
    }
    return children
}

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetYangName() string { return "ethernet-segments" }

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) SetParent(parent types.Entity) { ethernetSegments.parent = parent }

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetParent() types.Entity { return ethernetSegments.parent }

func (ethernetSegments *Evpn_Nodes_Node_EthernetSegments) GetParentYangName() string { return "node" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment
// EVPN Ethernet-Segment Entry
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
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

    // Ethernet Segment Name. The type is string.
    EthernetSegmentName interface{}

    // State of the ethernet segment. The type is interface{} with range:
    // 0..4294967295.
    EthernetSegmentState interface{}

    // Main port ifhandle. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetFilter() yfilter.YFilter { return ethernetSegment.YFilter }

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) SetFilter(yf yfilter.YFilter) { ethernetSegment.YFilter = yf }

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "esi1" { return "Esi1" }
    if yname == "esi2" { return "Esi2" }
    if yname == "esi3" { return "Esi3" }
    if yname == "esi4" { return "Esi4" }
    if yname == "esi5" { return "Esi5" }
    if yname == "esi-type" { return "EsiType" }
    if yname == "ethernet-segment-name" { return "EthernetSegmentName" }
    if yname == "ethernet-segment-state" { return "EthernetSegmentState" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "main-port-role" { return "MainPortRole" }
    if yname == "main-port-mac" { return "MainPortMac" }
    if yname == "num-up-p-ws" { return "NumUpPWs" }
    if yname == "route-target" { return "RouteTarget" }
    if yname == "rt-origin" { return "RtOrigin" }
    if yname == "es-bgp-gates" { return "EsBgpGates" }
    if yname == "es-l2fib-gates" { return "EsL2FibGates" }
    if yname == "mac-flushing-mode-config" { return "MacFlushingModeConfig" }
    if yname == "load-balance-mode-config" { return "LoadBalanceModeConfig" }
    if yname == "load-balance-mode-is-default" { return "LoadBalanceModeIsDefault" }
    if yname == "load-balance-mode-oper" { return "LoadBalanceModeOper" }
    if yname == "force-single-home" { return "ForceSingleHome" }
    if yname == "source-mac-oper" { return "SourceMacOper" }
    if yname == "source-mac-origin" { return "SourceMacOrigin" }
    if yname == "peering-timer" { return "PeeringTimer" }
    if yname == "peering-timer-left" { return "PeeringTimerLeft" }
    if yname == "recovery-timer" { return "RecoveryTimer" }
    if yname == "recovery-timer-left" { return "RecoveryTimerLeft" }
    if yname == "service-carving-mode" { return "ServiceCarvingMode" }
    if yname == "primary-services-input" { return "PrimaryServicesInput" }
    if yname == "secondary-services-input" { return "SecondaryServicesInput" }
    if yname == "forwarder-ports" { return "ForwarderPorts" }
    if yname == "permanent-forwarder-ports" { return "PermanentForwarderPorts" }
    if yname == "elected-forwarder-ports" { return "ElectedForwarderPorts" }
    if yname == "not-elected-forwarder-ports" { return "NotElectedForwarderPorts" }
    if yname == "not-config-forwarder-ports" { return "NotConfigForwarderPorts" }
    if yname == "mp-protected" { return "MpProtected" }
    if yname == "nve-anycast-vtep" { return "NveAnycastVtep" }
    if yname == "nve-ingress-replication" { return "NveIngressReplication" }
    if yname == "local-split-horizon-group-label" { return "LocalSplitHorizonGroupLabel" }
    if yname == "ethernet-segment-identifier" { return "EthernetSegmentIdentifier" }
    if yname == "primary-service" { return "PrimaryService" }
    if yname == "secondary-service" { return "SecondaryService" }
    if yname == "service-carving-i-sidelected-result" { return "ServiceCarvingISidelectedResult" }
    if yname == "service-carving-isid-not-elected-result" { return "ServiceCarvingIsidNotElectedResult" }
    if yname == "service-carving-evi-elected-result" { return "ServiceCarvingEviElectedResult" }
    if yname == "service-carving-evi-not-elected-result" { return "ServiceCarvingEviNotElectedResult" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "service-carving-vpws-permanent-result" { return "ServiceCarvingVpwsPermanentResult" }
    if yname == "remote-split-horizon-group-label" { return "RemoteSplitHorizonGroupLabel" }
    return ""
}

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetSegmentPath() string {
    return "ethernet-segment"
}

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-identifier" {
        for _, c := range ethernetSegment.EthernetSegmentIdentifier {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier{}
        ethernetSegment.EthernetSegmentIdentifier = append(ethernetSegment.EthernetSegmentIdentifier, child)
        return &ethernetSegment.EthernetSegmentIdentifier[len(ethernetSegment.EthernetSegmentIdentifier)-1]
    }
    if childYangName == "primary-service" {
        for _, c := range ethernetSegment.PrimaryService {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService{}
        ethernetSegment.PrimaryService = append(ethernetSegment.PrimaryService, child)
        return &ethernetSegment.PrimaryService[len(ethernetSegment.PrimaryService)-1]
    }
    if childYangName == "secondary-service" {
        for _, c := range ethernetSegment.SecondaryService {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService{}
        ethernetSegment.SecondaryService = append(ethernetSegment.SecondaryService, child)
        return &ethernetSegment.SecondaryService[len(ethernetSegment.SecondaryService)-1]
    }
    if childYangName == "service-carving-i-sidelected-result" {
        for _, c := range ethernetSegment.ServiceCarvingISidelectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult{}
        ethernetSegment.ServiceCarvingISidelectedResult = append(ethernetSegment.ServiceCarvingISidelectedResult, child)
        return &ethernetSegment.ServiceCarvingISidelectedResult[len(ethernetSegment.ServiceCarvingISidelectedResult)-1]
    }
    if childYangName == "service-carving-isid-not-elected-result" {
        for _, c := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult{}
        ethernetSegment.ServiceCarvingIsidNotElectedResult = append(ethernetSegment.ServiceCarvingIsidNotElectedResult, child)
        return &ethernetSegment.ServiceCarvingIsidNotElectedResult[len(ethernetSegment.ServiceCarvingIsidNotElectedResult)-1]
    }
    if childYangName == "service-carving-evi-elected-result" {
        for _, c := range ethernetSegment.ServiceCarvingEviElectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult{}
        ethernetSegment.ServiceCarvingEviElectedResult = append(ethernetSegment.ServiceCarvingEviElectedResult, child)
        return &ethernetSegment.ServiceCarvingEviElectedResult[len(ethernetSegment.ServiceCarvingEviElectedResult)-1]
    }
    if childYangName == "service-carving-evi-not-elected-result" {
        for _, c := range ethernetSegment.ServiceCarvingEviNotElectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult{}
        ethernetSegment.ServiceCarvingEviNotElectedResult = append(ethernetSegment.ServiceCarvingEviNotElectedResult, child)
        return &ethernetSegment.ServiceCarvingEviNotElectedResult[len(ethernetSegment.ServiceCarvingEviNotElectedResult)-1]
    }
    if childYangName == "next-hop" {
        for _, c := range ethernetSegment.NextHop {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop{}
        ethernetSegment.NextHop = append(ethernetSegment.NextHop, child)
        return &ethernetSegment.NextHop[len(ethernetSegment.NextHop)-1]
    }
    if childYangName == "service-carving-vpws-permanent-result" {
        for _, c := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult{}
        ethernetSegment.ServiceCarvingVpwsPermanentResult = append(ethernetSegment.ServiceCarvingVpwsPermanentResult, child)
        return &ethernetSegment.ServiceCarvingVpwsPermanentResult[len(ethernetSegment.ServiceCarvingVpwsPermanentResult)-1]
    }
    if childYangName == "remote-split-horizon-group-label" {
        for _, c := range ethernetSegment.RemoteSplitHorizonGroupLabel {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel{}
        ethernetSegment.RemoteSplitHorizonGroupLabel = append(ethernetSegment.RemoteSplitHorizonGroupLabel, child)
        return &ethernetSegment.RemoteSplitHorizonGroupLabel[len(ethernetSegment.RemoteSplitHorizonGroupLabel)-1]
    }
    return nil
}

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetSegment.EthernetSegmentIdentifier {
        children[ethernetSegment.EthernetSegmentIdentifier[i].GetSegmentPath()] = &ethernetSegment.EthernetSegmentIdentifier[i]
    }
    for i := range ethernetSegment.PrimaryService {
        children[ethernetSegment.PrimaryService[i].GetSegmentPath()] = &ethernetSegment.PrimaryService[i]
    }
    for i := range ethernetSegment.SecondaryService {
        children[ethernetSegment.SecondaryService[i].GetSegmentPath()] = &ethernetSegment.SecondaryService[i]
    }
    for i := range ethernetSegment.ServiceCarvingISidelectedResult {
        children[ethernetSegment.ServiceCarvingISidelectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingISidelectedResult[i]
    }
    for i := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
        children[ethernetSegment.ServiceCarvingIsidNotElectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingIsidNotElectedResult[i]
    }
    for i := range ethernetSegment.ServiceCarvingEviElectedResult {
        children[ethernetSegment.ServiceCarvingEviElectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingEviElectedResult[i]
    }
    for i := range ethernetSegment.ServiceCarvingEviNotElectedResult {
        children[ethernetSegment.ServiceCarvingEviNotElectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingEviNotElectedResult[i]
    }
    for i := range ethernetSegment.NextHop {
        children[ethernetSegment.NextHop[i].GetSegmentPath()] = &ethernetSegment.NextHop[i]
    }
    for i := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
        children[ethernetSegment.ServiceCarvingVpwsPermanentResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingVpwsPermanentResult[i]
    }
    for i := range ethernetSegment.RemoteSplitHorizonGroupLabel {
        children[ethernetSegment.RemoteSplitHorizonGroupLabel[i].GetSegmentPath()] = &ethernetSegment.RemoteSplitHorizonGroupLabel[i]
    }
    return children
}

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = ethernetSegment.InterfaceName
    leafs["esi1"] = ethernetSegment.Esi1
    leafs["esi2"] = ethernetSegment.Esi2
    leafs["esi3"] = ethernetSegment.Esi3
    leafs["esi4"] = ethernetSegment.Esi4
    leafs["esi5"] = ethernetSegment.Esi5
    leafs["esi-type"] = ethernetSegment.EsiType
    leafs["ethernet-segment-name"] = ethernetSegment.EthernetSegmentName
    leafs["ethernet-segment-state"] = ethernetSegment.EthernetSegmentState
    leafs["if-handle"] = ethernetSegment.IfHandle
    leafs["main-port-role"] = ethernetSegment.MainPortRole
    leafs["main-port-mac"] = ethernetSegment.MainPortMac
    leafs["num-up-p-ws"] = ethernetSegment.NumUpPWs
    leafs["route-target"] = ethernetSegment.RouteTarget
    leafs["rt-origin"] = ethernetSegment.RtOrigin
    leafs["es-bgp-gates"] = ethernetSegment.EsBgpGates
    leafs["es-l2fib-gates"] = ethernetSegment.EsL2FibGates
    leafs["mac-flushing-mode-config"] = ethernetSegment.MacFlushingModeConfig
    leafs["load-balance-mode-config"] = ethernetSegment.LoadBalanceModeConfig
    leafs["load-balance-mode-is-default"] = ethernetSegment.LoadBalanceModeIsDefault
    leafs["load-balance-mode-oper"] = ethernetSegment.LoadBalanceModeOper
    leafs["force-single-home"] = ethernetSegment.ForceSingleHome
    leafs["source-mac-oper"] = ethernetSegment.SourceMacOper
    leafs["source-mac-origin"] = ethernetSegment.SourceMacOrigin
    leafs["peering-timer"] = ethernetSegment.PeeringTimer
    leafs["peering-timer-left"] = ethernetSegment.PeeringTimerLeft
    leafs["recovery-timer"] = ethernetSegment.RecoveryTimer
    leafs["recovery-timer-left"] = ethernetSegment.RecoveryTimerLeft
    leafs["service-carving-mode"] = ethernetSegment.ServiceCarvingMode
    leafs["primary-services-input"] = ethernetSegment.PrimaryServicesInput
    leafs["secondary-services-input"] = ethernetSegment.SecondaryServicesInput
    leafs["forwarder-ports"] = ethernetSegment.ForwarderPorts
    leafs["permanent-forwarder-ports"] = ethernetSegment.PermanentForwarderPorts
    leafs["elected-forwarder-ports"] = ethernetSegment.ElectedForwarderPorts
    leafs["not-elected-forwarder-ports"] = ethernetSegment.NotElectedForwarderPorts
    leafs["not-config-forwarder-ports"] = ethernetSegment.NotConfigForwarderPorts
    leafs["mp-protected"] = ethernetSegment.MpProtected
    leafs["nve-anycast-vtep"] = ethernetSegment.NveAnycastVtep
    leafs["nve-ingress-replication"] = ethernetSegment.NveIngressReplication
    leafs["local-split-horizon-group-label"] = ethernetSegment.LocalSplitHorizonGroupLabel
    return leafs
}

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetYangName() string { return "ethernet-segment" }

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) SetParent(parent types.Entity) { ethernetSegment.parent = parent }

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetParent() types.Entity { return ethernetSegment.parent }

func (ethernetSegment *Evpn_Nodes_Node_EthernetSegments_EthernetSegment) GetParentYangName() string { return "ethernet-segments" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return ethernetSegmentIdentifier.YFilter }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { ethernetSegmentIdentifier.YFilter = yf }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetSegmentPath() string {
    return "ethernet-segment-identifier"
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = ethernetSegmentIdentifier.Entry
    return leafs
}

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetYangName() string { return "ethernet-segment-identifier" }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) SetParent(parent types.Entity) { ethernetSegmentIdentifier.parent = parent }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetParent() types.Entity { return ethernetSegmentIdentifier.parent }

func (ethernetSegmentIdentifier *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService
// List of Primary services ESI/I-SIDs
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetFilter() yfilter.YFilter { return primaryService.YFilter }

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) SetFilter(yf yfilter.YFilter) { primaryService.YFilter = yf }

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetSegmentPath() string {
    return "primary-service"
}

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = primaryService.Entry
    return leafs
}

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetBundleName() string { return "cisco_ios_xr" }

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetYangName() string { return "primary-service" }

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) SetParent(parent types.Entity) { primaryService.parent = parent }

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetParent() types.Entity { return primaryService.parent }

func (primaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_PrimaryService) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService
// List of Secondary services ESI/I-SIDs
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetFilter() yfilter.YFilter { return secondaryService.YFilter }

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) SetFilter(yf yfilter.YFilter) { secondaryService.YFilter = yf }

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetSegmentPath() string {
    return "secondary-service"
}

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = secondaryService.Entry
    return leafs
}

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetYangName() string { return "secondary-service" }

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) SetParent(parent types.Entity) { secondaryService.parent = parent }

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetParent() types.Entity { return secondaryService.parent }

func (secondaryService *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_SecondaryService) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult
// Elected ISID service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetFilter() yfilter.YFilter { return serviceCarvingISidelectedResult.YFilter }

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingISidelectedResult.YFilter = yf }

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetSegmentPath() string {
    return "service-carving-i-sidelected-result"
}

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingISidelectedResult.Entry
    return leafs
}

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetYangName() string { return "service-carving-i-sidelected-result" }

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) SetParent(parent types.Entity) { serviceCarvingISidelectedResult.parent = parent }

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetParent() types.Entity { return serviceCarvingISidelectedResult.parent }

func (serviceCarvingISidelectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult
// Not elected ISID service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetFilter() yfilter.YFilter { return serviceCarvingIsidNotElectedResult.YFilter }

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingIsidNotElectedResult.YFilter = yf }

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetSegmentPath() string {
    return "service-carving-isid-not-elected-result"
}

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingIsidNotElectedResult.Entry
    return leafs
}

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetYangName() string { return "service-carving-isid-not-elected-result" }

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) SetParent(parent types.Entity) { serviceCarvingIsidNotElectedResult.parent = parent }

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetParent() types.Entity { return serviceCarvingIsidNotElectedResult.parent }

func (serviceCarvingIsidNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult
// Elected EVI service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetFilter() yfilter.YFilter { return serviceCarvingEviElectedResult.YFilter }

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingEviElectedResult.YFilter = yf }

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetSegmentPath() string {
    return "service-carving-evi-elected-result"
}

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingEviElectedResult.Entry
    return leafs
}

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetYangName() string { return "service-carving-evi-elected-result" }

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) SetParent(parent types.Entity) { serviceCarvingEviElectedResult.parent = parent }

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetParent() types.Entity { return serviceCarvingEviElectedResult.parent }

func (serviceCarvingEviElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult
// Not elected EVI service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetFilter() yfilter.YFilter { return serviceCarvingEviNotElectedResult.YFilter }

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingEviNotElectedResult.YFilter = yf }

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetSegmentPath() string {
    return "service-carving-evi-not-elected-result"
}

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingEviNotElectedResult.Entry
    return leafs
}

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetYangName() string { return "service-carving-evi-not-elected-result" }

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) SetParent(parent types.Entity) { serviceCarvingEviNotElectedResult.parent = parent }

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetParent() types.Entity { return serviceCarvingEviNotElectedResult.parent }

func (serviceCarvingEviNotElectedResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop
// List of nexthop IPv6 addresses
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}
}

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHop.NextHop
    return leafs
}

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_NextHop) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult
// Permanent EVPN VPWS service carving results
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPN ID. The type is interface{} with range: 0..4294967295.
    VpnId interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}
}

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetFilter() yfilter.YFilter { return serviceCarvingVpwsPermanentResult.YFilter }

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) SetFilter(yf yfilter.YFilter) { serviceCarvingVpwsPermanentResult.YFilter = yf }

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetGoName(yname string) string {
    if yname == "vpn-id" { return "VpnId" }
    if yname == "type" { return "Type" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    return ""
}

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetSegmentPath() string {
    return "service-carving-vpws-permanent-result"
}

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vpn-id"] = serviceCarvingVpwsPermanentResult.VpnId
    leafs["type"] = serviceCarvingVpwsPermanentResult.Type
    leafs["ethernet-tag"] = serviceCarvingVpwsPermanentResult.EthernetTag
    return leafs
}

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetYangName() string { return "service-carving-vpws-permanent-result" }

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) SetParent(parent types.Entity) { serviceCarvingVpwsPermanentResult.parent = parent }

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetParent() types.Entity { return serviceCarvingVpwsPermanentResult.parent }

func (serviceCarvingVpwsPermanentResult *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Split horizon label associated with next-hop address. The type is
    // interface{} with range: 0..4294967295.
    Label interface{}
}

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetFilter() yfilter.YFilter { return remoteSplitHorizonGroupLabel.YFilter }

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) SetFilter(yf yfilter.YFilter) { remoteSplitHorizonGroupLabel.YFilter = yf }

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "label" { return "Label" }
    return ""
}

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetSegmentPath() string {
    return "remote-split-horizon-group-label"
}

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = remoteSplitHorizonGroupLabel.NextHop
    leafs["label"] = remoteSplitHorizonGroupLabel.Label
    return leafs
}

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetBundleName() string { return "cisco_ios_xr" }

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetYangName() string { return "remote-split-horizon-group-label" }

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) SetParent(parent types.Entity) { remoteSplitHorizonGroupLabel.parent = parent }

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetParent() types.Entity { return remoteSplitHorizonGroupLabel.parent }

func (remoteSplitHorizonGroupLabel *Evpn_Nodes_Node_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Nodes_Node_AcIds
// EVPN AC ID table
type Evpn_Nodes_Node_AcIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN AC ID table. The type is slice of Evpn_Nodes_Node_AcIds_AcId.
    AcId []Evpn_Nodes_Node_AcIds_AcId
}

func (acIds *Evpn_Nodes_Node_AcIds) GetFilter() yfilter.YFilter { return acIds.YFilter }

func (acIds *Evpn_Nodes_Node_AcIds) SetFilter(yf yfilter.YFilter) { acIds.YFilter = yf }

func (acIds *Evpn_Nodes_Node_AcIds) GetGoName(yname string) string {
    if yname == "ac-id" { return "AcId" }
    return ""
}

func (acIds *Evpn_Nodes_Node_AcIds) GetSegmentPath() string {
    return "ac-ids"
}

func (acIds *Evpn_Nodes_Node_AcIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ac-id" {
        for _, c := range acIds.AcId {
            if acIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Nodes_Node_AcIds_AcId{}
        acIds.AcId = append(acIds.AcId, child)
        return &acIds.AcId[len(acIds.AcId)-1]
    }
    return nil
}

func (acIds *Evpn_Nodes_Node_AcIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range acIds.AcId {
        children[acIds.AcId[i].GetSegmentPath()] = &acIds.AcId[i]
    }
    return children
}

func (acIds *Evpn_Nodes_Node_AcIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (acIds *Evpn_Nodes_Node_AcIds) GetBundleName() string { return "cisco_ios_xr" }

func (acIds *Evpn_Nodes_Node_AcIds) GetYangName() string { return "ac-ids" }

func (acIds *Evpn_Nodes_Node_AcIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acIds *Evpn_Nodes_Node_AcIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acIds *Evpn_Nodes_Node_AcIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acIds *Evpn_Nodes_Node_AcIds) SetParent(parent types.Entity) { acIds.parent = parent }

func (acIds *Evpn_Nodes_Node_AcIds) GetParent() types.Entity { return acIds.parent }

func (acIds *Evpn_Nodes_Node_AcIds) GetParentYangName() string { return "node" }

// Evpn_Nodes_Node_AcIds_AcId
// EVPN AC ID table
type Evpn_Nodes_Node_AcIds_AcId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // AC ID. The type is interface{} with range: -2147483648..2147483647.
    AcId interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}
}

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetFilter() yfilter.YFilter { return acId.YFilter }

func (acId *Evpn_Nodes_Node_AcIds_AcId) SetFilter(yf yfilter.YFilter) { acId.YFilter = yf }

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "ac-id" { return "AcId" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetSegmentPath() string {
    return "ac-id"
}

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = acId.Evi
    leafs["ac-id"] = acId.AcId
    leafs["evi-xr"] = acId.EviXr
    leafs["neighbor"] = acId.Neighbor
    return leafs
}

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetBundleName() string { return "cisco_ios_xr" }

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetYangName() string { return "ac-id" }

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acId *Evpn_Nodes_Node_AcIds_AcId) SetParent(parent types.Entity) { acId.parent = parent }

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetParent() types.Entity { return acId.parent }

func (acId *Evpn_Nodes_Node_AcIds_AcId) GetParentYangName() string { return "ac-ids" }

// Evpn_Active
// Active EVPN operational data
type Evpn_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Table.
    Evis Evpn_Active_Evis

    // L2VPN EVPN Summary.
    Summary Evpn_Active_Summary

    // L2VPN EVI Detail Table.
    EviDetail Evpn_Active_EviDetail

    // EVPN Ethernet-Segment Table.
    EthernetSegments Evpn_Active_EthernetSegments

    // EVPN AC ID table.
    AcIds Evpn_Active_AcIds
}

func (active *Evpn_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *Evpn_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *Evpn_Active) GetGoName(yname string) string {
    if yname == "evis" { return "Evis" }
    if yname == "summary" { return "Summary" }
    if yname == "evi-detail" { return "EviDetail" }
    if yname == "ethernet-segments" { return "EthernetSegments" }
    if yname == "ac-ids" { return "AcIds" }
    return ""
}

func (active *Evpn_Active) GetSegmentPath() string {
    return "active"
}

func (active *Evpn_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "evis" {
        return &active.Evis
    }
    if childYangName == "summary" {
        return &active.Summary
    }
    if childYangName == "evi-detail" {
        return &active.EviDetail
    }
    if childYangName == "ethernet-segments" {
        return &active.EthernetSegments
    }
    if childYangName == "ac-ids" {
        return &active.AcIds
    }
    return nil
}

func (active *Evpn_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["evis"] = &active.Evis
    children["summary"] = &active.Summary
    children["evi-detail"] = &active.EviDetail
    children["ethernet-segments"] = &active.EthernetSegments
    children["ac-ids"] = &active.AcIds
    return children
}

func (active *Evpn_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (active *Evpn_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *Evpn_Active) GetYangName() string { return "active" }

func (active *Evpn_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *Evpn_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *Evpn_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *Evpn_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *Evpn_Active) GetParent() types.Entity { return active.parent }

func (active *Evpn_Active) GetParentYangName() string { return "evpn" }

// Evpn_Active_Evis
// L2VPN EVPN EVI Table
type Evpn_Active_Evis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Entry. The type is slice of Evpn_Active_Evis_Evi.
    Evi []Evpn_Active_Evis_Evi
}

func (evis *Evpn_Active_Evis) GetFilter() yfilter.YFilter { return evis.YFilter }

func (evis *Evpn_Active_Evis) SetFilter(yf yfilter.YFilter) { evis.YFilter = yf }

func (evis *Evpn_Active_Evis) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    return ""
}

func (evis *Evpn_Active_Evis) GetSegmentPath() string {
    return "evis"
}

func (evis *Evpn_Active_Evis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "evi" {
        for _, c := range evis.Evi {
            if evis.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_Evis_Evi{}
        evis.Evi = append(evis.Evi, child)
        return &evis.Evi[len(evis.Evi)-1]
    }
    return nil
}

func (evis *Evpn_Active_Evis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range evis.Evi {
        children[evis.Evi[i].GetSegmentPath()] = &evis.Evi[i]
    }
    return children
}

func (evis *Evpn_Active_Evis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (evis *Evpn_Active_Evis) GetBundleName() string { return "cisco_ios_xr" }

func (evis *Evpn_Active_Evis) GetYangName() string { return "evis" }

func (evis *Evpn_Active_Evis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evis *Evpn_Active_Evis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evis *Evpn_Active_Evis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evis *Evpn_Active_Evis) SetParent(parent types.Entity) { evis.parent = parent }

func (evis *Evpn_Active_Evis) GetParent() types.Entity { return evis.parent }

func (evis *Evpn_Active_Evis) GetParentYangName() string { return "active" }

// Evpn_Active_Evis_Evi
// L2VPN EVPN EVI Entry
type Evpn_Active_Evis_Evi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN id. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evi *Evpn_Active_Evis_Evi) GetFilter() yfilter.YFilter { return evi.YFilter }

func (evi *Evpn_Active_Evis_Evi) SetFilter(yf yfilter.YFilter) { evi.YFilter = yf }

func (evi *Evpn_Active_Evis_Evi) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "bd-name" { return "BdName" }
    if yname == "type" { return "Type" }
    return ""
}

func (evi *Evpn_Active_Evis_Evi) GetSegmentPath() string {
    return "evi" + "[evi='" + fmt.Sprintf("%v", evi.Evi) + "']"
}

func (evi *Evpn_Active_Evis_Evi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (evi *Evpn_Active_Evis_Evi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (evi *Evpn_Active_Evis_Evi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = evi.Evi
    leafs["evi-xr"] = evi.EviXr
    leafs["bd-name"] = evi.BdName
    leafs["type"] = evi.Type
    return leafs
}

func (evi *Evpn_Active_Evis_Evi) GetBundleName() string { return "cisco_ios_xr" }

func (evi *Evpn_Active_Evis_Evi) GetYangName() string { return "evi" }

func (evi *Evpn_Active_Evis_Evi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evi *Evpn_Active_Evis_Evi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evi *Evpn_Active_Evis_Evi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evi *Evpn_Active_Evis_Evi) SetParent(parent types.Entity) { evi.parent = parent }

func (evi *Evpn_Active_Evis_Evi) GetParent() types.Entity { return evi.parent }

func (evi *Evpn_Active_Evis_Evi) GetParentYangName() string { return "evis" }

// Evpn_Active_Summary
// L2VPN EVPN Summary
type Evpn_Active_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Router ID. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    GlobalSourceMac interface{}

    // EVPN ES Peering Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    PeeringTime interface{}

    // EVPN ES Recovery Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    RecoveryTime interface{}

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

func (summary *Evpn_Active_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Evpn_Active_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Evpn_Active_Summary) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "as" { return "As" }
    if yname == "ev-is" { return "EvIs" }
    if yname == "local-mac-routes" { return "LocalMacRoutes" }
    if yname == "local-ipv4-mac-routes" { return "LocalIpv4MacRoutes" }
    if yname == "local-ipv6-mac-routes" { return "LocalIpv6MacRoutes" }
    if yname == "es-global-mac-routes" { return "EsGlobalMacRoutes" }
    if yname == "remote-mac-routes" { return "RemoteMacRoutes" }
    if yname == "remote-soo-mac-routes" { return "RemoteSooMacRoutes" }
    if yname == "remote-ipv4-mac-routes" { return "RemoteIpv4MacRoutes" }
    if yname == "remote-ipv6-mac-routes" { return "RemoteIpv6MacRoutes" }
    if yname == "local-imcast-routes" { return "LocalImcastRoutes" }
    if yname == "remote-imcast-routes" { return "RemoteImcastRoutes" }
    if yname == "labels" { return "Labels" }
    if yname == "es-entries" { return "EsEntries" }
    if yname == "neighbor-entries" { return "NeighborEntries" }
    if yname == "local-ead-routes" { return "LocalEadRoutes" }
    if yname == "remote-ead-routes" { return "RemoteEadRoutes" }
    if yname == "global-source-mac" { return "GlobalSourceMac" }
    if yname == "peering-time" { return "PeeringTime" }
    if yname == "recovery-time" { return "RecoveryTime" }
    if yname == "mac-secure-move-count" { return "MacSecureMoveCount" }
    if yname == "mac-secure-move-interval" { return "MacSecureMoveInterval" }
    if yname == "mac-secure-freeze-time" { return "MacSecureFreezeTime" }
    if yname == "mac-secure-retry-count" { return "MacSecureRetryCount" }
    if yname == "cost-out" { return "CostOut" }
    if yname == "startup-cost-in-time" { return "StartupCostInTime" }
    if yname == "l2rib-throttle" { return "L2RibThrottle" }
    if yname == "logging-df-election-enabled" { return "LoggingDfElectionEnabled" }
    return ""
}

func (summary *Evpn_Active_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Evpn_Active_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Evpn_Active_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Evpn_Active_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = summary.RouterId
    leafs["as"] = summary.As
    leafs["ev-is"] = summary.EvIs
    leafs["local-mac-routes"] = summary.LocalMacRoutes
    leafs["local-ipv4-mac-routes"] = summary.LocalIpv4MacRoutes
    leafs["local-ipv6-mac-routes"] = summary.LocalIpv6MacRoutes
    leafs["es-global-mac-routes"] = summary.EsGlobalMacRoutes
    leafs["remote-mac-routes"] = summary.RemoteMacRoutes
    leafs["remote-soo-mac-routes"] = summary.RemoteSooMacRoutes
    leafs["remote-ipv4-mac-routes"] = summary.RemoteIpv4MacRoutes
    leafs["remote-ipv6-mac-routes"] = summary.RemoteIpv6MacRoutes
    leafs["local-imcast-routes"] = summary.LocalImcastRoutes
    leafs["remote-imcast-routes"] = summary.RemoteImcastRoutes
    leafs["labels"] = summary.Labels
    leafs["es-entries"] = summary.EsEntries
    leafs["neighbor-entries"] = summary.NeighborEntries
    leafs["local-ead-routes"] = summary.LocalEadRoutes
    leafs["remote-ead-routes"] = summary.RemoteEadRoutes
    leafs["global-source-mac"] = summary.GlobalSourceMac
    leafs["peering-time"] = summary.PeeringTime
    leafs["recovery-time"] = summary.RecoveryTime
    leafs["mac-secure-move-count"] = summary.MacSecureMoveCount
    leafs["mac-secure-move-interval"] = summary.MacSecureMoveInterval
    leafs["mac-secure-freeze-time"] = summary.MacSecureFreezeTime
    leafs["mac-secure-retry-count"] = summary.MacSecureRetryCount
    leafs["cost-out"] = summary.CostOut
    leafs["startup-cost-in-time"] = summary.StartupCostInTime
    leafs["l2rib-throttle"] = summary.L2RibThrottle
    leafs["logging-df-election-enabled"] = summary.LoggingDfElectionEnabled
    return leafs
}

func (summary *Evpn_Active_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Evpn_Active_Summary) GetYangName() string { return "summary" }

func (summary *Evpn_Active_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Evpn_Active_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Evpn_Active_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Evpn_Active_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Evpn_Active_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Evpn_Active_Summary) GetParentYangName() string { return "active" }

// Evpn_Active_EviDetail
// L2VPN EVI Detail Table
type Evpn_Active_EviDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info Elements.
    Elements Evpn_Active_EviDetail_Elements

    // Container for all EVI detail info.
    EviChildren Evpn_Active_EviDetail_EviChildren
}

func (eviDetail *Evpn_Active_EviDetail) GetFilter() yfilter.YFilter { return eviDetail.YFilter }

func (eviDetail *Evpn_Active_EviDetail) SetFilter(yf yfilter.YFilter) { eviDetail.YFilter = yf }

func (eviDetail *Evpn_Active_EviDetail) GetGoName(yname string) string {
    if yname == "elements" { return "Elements" }
    if yname == "evi-children" { return "EviChildren" }
    return ""
}

func (eviDetail *Evpn_Active_EviDetail) GetSegmentPath() string {
    return "evi-detail"
}

func (eviDetail *Evpn_Active_EviDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "elements" {
        return &eviDetail.Elements
    }
    if childYangName == "evi-children" {
        return &eviDetail.EviChildren
    }
    return nil
}

func (eviDetail *Evpn_Active_EviDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["elements"] = &eviDetail.Elements
    children["evi-children"] = &eviDetail.EviChildren
    return children
}

func (eviDetail *Evpn_Active_EviDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eviDetail *Evpn_Active_EviDetail) GetBundleName() string { return "cisco_ios_xr" }

func (eviDetail *Evpn_Active_EviDetail) GetYangName() string { return "evi-detail" }

func (eviDetail *Evpn_Active_EviDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eviDetail *Evpn_Active_EviDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eviDetail *Evpn_Active_EviDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eviDetail *Evpn_Active_EviDetail) SetParent(parent types.Entity) { eviDetail.parent = parent }

func (eviDetail *Evpn_Active_EviDetail) GetParent() types.Entity { return eviDetail.parent }

func (eviDetail *Evpn_Active_EviDetail) GetParentYangName() string { return "active" }

// Evpn_Active_EviDetail_Elements
// EVI BGP RT Detail Info Elements
type Evpn_Active_EviDetail_Elements struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info. The type is slice of
    // Evpn_Active_EviDetail_Elements_Element.
    Element []Evpn_Active_EviDetail_Elements_Element
}

func (elements *Evpn_Active_EviDetail_Elements) GetFilter() yfilter.YFilter { return elements.YFilter }

func (elements *Evpn_Active_EviDetail_Elements) SetFilter(yf yfilter.YFilter) { elements.YFilter = yf }

func (elements *Evpn_Active_EviDetail_Elements) GetGoName(yname string) string {
    if yname == "element" { return "Element" }
    return ""
}

func (elements *Evpn_Active_EviDetail_Elements) GetSegmentPath() string {
    return "elements"
}

func (elements *Evpn_Active_EviDetail_Elements) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "element" {
        for _, c := range elements.Element {
            if elements.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_Elements_Element{}
        elements.Element = append(elements.Element, child)
        return &elements.Element[len(elements.Element)-1]
    }
    return nil
}

func (elements *Evpn_Active_EviDetail_Elements) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range elements.Element {
        children[elements.Element[i].GetSegmentPath()] = &elements.Element[i]
    }
    return children
}

func (elements *Evpn_Active_EviDetail_Elements) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (elements *Evpn_Active_EviDetail_Elements) GetBundleName() string { return "cisco_ios_xr" }

func (elements *Evpn_Active_EviDetail_Elements) GetYangName() string { return "elements" }

func (elements *Evpn_Active_EviDetail_Elements) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (elements *Evpn_Active_EviDetail_Elements) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (elements *Evpn_Active_EviDetail_Elements) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (elements *Evpn_Active_EviDetail_Elements) SetParent(parent types.Entity) { elements.parent = parent }

func (elements *Evpn_Active_EviDetail_Elements) GetParent() types.Entity { return elements.parent }

func (elements *Evpn_Active_EviDetail_Elements) GetParentYangName() string { return "evi-detail" }

// Evpn_Active_EviDetail_Elements_Element
// EVI BGP RT Detail Info
type Evpn_Active_EviDetail_Elements_Element struct {
    parent types.Entity
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
    Type interface{}

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

func (element *Evpn_Active_EviDetail_Elements_Element) GetFilter() yfilter.YFilter { return element.YFilter }

func (element *Evpn_Active_EviDetail_Elements_Element) SetFilter(yf yfilter.YFilter) { element.YFilter = yf }

func (element *Evpn_Active_EviDetail_Elements_Element) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "description" { return "Description" }
    if yname == "bd-name" { return "BdName" }
    if yname == "type" { return "Type" }
    if yname == "unicast-label" { return "UnicastLabel" }
    if yname == "multicast-label" { return "MulticastLabel" }
    if yname == "cw-disable" { return "CwDisable" }
    if yname == "table-policy-name" { return "TablePolicyName" }
    if yname == "forward-class" { return "ForwardClass" }
    if yname == "rt-import-block-set" { return "RtImportBlockSet" }
    if yname == "rt-export-block-set" { return "RtExportBlockSet" }
    if yname == "advertise-mac" { return "AdvertiseMac" }
    if yname == "advertise-bvi-mac" { return "AdvertiseBviMac" }
    if yname == "aliasing-disabled" { return "AliasingDisabled" }
    if yname == "unknown-unicast-flooding-disabled" { return "UnknownUnicastFloodingDisabled" }
    if yname == "reoriginate-disabled" { return "ReoriginateDisabled" }
    if yname == "stitching" { return "Stitching" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "flow-label" { return "FlowLabel" }
    if yname == "rd-auto" { return "RdAuto" }
    if yname == "rd-configured" { return "RdConfigured" }
    if yname == "rt-auto" { return "RtAuto" }
    if yname == "rt-auto-stitching" { return "RtAutoStitching" }
    return ""
}

func (element *Evpn_Active_EviDetail_Elements_Element) GetSegmentPath() string {
    return "element" + "[evi='" + fmt.Sprintf("%v", element.Evi) + "']"
}

func (element *Evpn_Active_EviDetail_Elements_Element) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-label" {
        return &element.FlowLabel
    }
    if childYangName == "rd-auto" {
        return &element.RdAuto
    }
    if childYangName == "rd-configured" {
        return &element.RdConfigured
    }
    if childYangName == "rt-auto" {
        return &element.RtAuto
    }
    if childYangName == "rt-auto-stitching" {
        return &element.RtAutoStitching
    }
    return nil
}

func (element *Evpn_Active_EviDetail_Elements_Element) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flow-label"] = &element.FlowLabel
    children["rd-auto"] = &element.RdAuto
    children["rd-configured"] = &element.RdConfigured
    children["rt-auto"] = &element.RtAuto
    children["rt-auto-stitching"] = &element.RtAutoStitching
    return children
}

func (element *Evpn_Active_EviDetail_Elements_Element) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = element.Evi
    leafs["evi-xr"] = element.EviXr
    leafs["description"] = element.Description
    leafs["bd-name"] = element.BdName
    leafs["type"] = element.Type
    leafs["unicast-label"] = element.UnicastLabel
    leafs["multicast-label"] = element.MulticastLabel
    leafs["cw-disable"] = element.CwDisable
    leafs["table-policy-name"] = element.TablePolicyName
    leafs["forward-class"] = element.ForwardClass
    leafs["rt-import-block-set"] = element.RtImportBlockSet
    leafs["rt-export-block-set"] = element.RtExportBlockSet
    leafs["advertise-mac"] = element.AdvertiseMac
    leafs["advertise-bvi-mac"] = element.AdvertiseBviMac
    leafs["aliasing-disabled"] = element.AliasingDisabled
    leafs["unknown-unicast-flooding-disabled"] = element.UnknownUnicastFloodingDisabled
    leafs["reoriginate-disabled"] = element.ReoriginateDisabled
    leafs["stitching"] = element.Stitching
    leafs["encapsulation"] = element.Encapsulation
    return leafs
}

func (element *Evpn_Active_EviDetail_Elements_Element) GetBundleName() string { return "cisco_ios_xr" }

func (element *Evpn_Active_EviDetail_Elements_Element) GetYangName() string { return "element" }

func (element *Evpn_Active_EviDetail_Elements_Element) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (element *Evpn_Active_EviDetail_Elements_Element) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (element *Evpn_Active_EviDetail_Elements_Element) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (element *Evpn_Active_EviDetail_Elements_Element) SetParent(parent types.Entity) { element.parent = parent }

func (element *Evpn_Active_EviDetail_Elements_Element) GetParent() types.Entity { return element.parent }

func (element *Evpn_Active_EviDetail_Elements_Element) GetParentYangName() string { return "elements" }

// Evpn_Active_EviDetail_Elements_Element_FlowLabel
// Flow Label Information
type Evpn_Active_EviDetail_Elements_Element_FlowLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Static flow label. The type is bool.
    StaticFlowLabel interface{}

    // Globally configured flow label. The type is bool.
    GlobalFlowLabel interface{}
}

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetFilter() yfilter.YFilter { return flowLabel.YFilter }

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) SetFilter(yf yfilter.YFilter) { flowLabel.YFilter = yf }

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetGoName(yname string) string {
    if yname == "static-flow-label" { return "StaticFlowLabel" }
    if yname == "global-flow-label" { return "GlobalFlowLabel" }
    return ""
}

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetSegmentPath() string {
    return "flow-label"
}

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["static-flow-label"] = flowLabel.StaticFlowLabel
    leafs["global-flow-label"] = flowLabel.GlobalFlowLabel
    return leafs
}

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetBundleName() string { return "cisco_ios_xr" }

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetYangName() string { return "flow-label" }

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) SetParent(parent types.Entity) { flowLabel.parent = parent }

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetParent() types.Entity { return flowLabel.parent }

func (flowLabel *Evpn_Active_EviDetail_Elements_Element_FlowLabel) GetParentYangName() string { return "element" }

// Evpn_Active_EviDetail_Elements_Element_RdAuto
// Automatic Route Distingtuisher
type Evpn_Active_EviDetail_Elements_Element_RdAuto struct {
    parent types.Entity
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

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetFilter() yfilter.YFilter { return rdAuto.YFilter }

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) SetFilter(yf yfilter.YFilter) { rdAuto.YFilter = yf }

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetGoName(yname string) string {
    if yname == "rd" { return "Rd" }
    if yname == "auto" { return "Auto" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    return ""
}

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetSegmentPath() string {
    return "rd-auto"
}

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto" {
        return &rdAuto.Auto
    }
    if childYangName == "two-byte-as" {
        return &rdAuto.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rdAuto.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rdAuto.V4Addr
    }
    return nil
}

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto"] = &rdAuto.Auto
    children["two-byte-as"] = &rdAuto.TwoByteAs
    children["four-byte-as"] = &rdAuto.FourByteAs
    children["v4-addr"] = &rdAuto.V4Addr
    return children
}

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rd"] = rdAuto.Rd
    return leafs
}

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetBundleName() string { return "cisco_ios_xr" }

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetYangName() string { return "rd-auto" }

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) SetParent(parent types.Entity) { rdAuto.parent = parent }

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetParent() types.Entity { return rdAuto.parent }

func (rdAuto *Evpn_Active_EviDetail_Elements_Element_RdAuto) GetParentYangName() string { return "element" }

// Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto
// auto
type Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetFilter() yfilter.YFilter { return auto.YFilter }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) SetFilter(yf yfilter.YFilter) { auto.YFilter = yf }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "auto-index" { return "AutoIndex" }
    return ""
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetSegmentPath() string {
    return "auto"
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = auto.RouterId
    leafs["auto-index"] = auto.AutoIndex
    return leafs
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetBundleName() string { return "cisco_ios_xr" }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetYangName() string { return "auto" }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) SetParent(parent types.Entity) { auto.parent = parent }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetParent() types.Entity { return auto.parent }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdAuto_Auto) GetParentYangName() string { return "rd-auto" }

// Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetParentYangName() string { return "rd-auto" }

// Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs
// four byte as
type Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdAuto_FourByteAs) GetParentYangName() string { return "rd-auto" }

// Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdAuto_V4Addr) GetParentYangName() string { return "rd-auto" }

// Evpn_Active_EviDetail_Elements_Element_RdConfigured
// Configured Route Distinguisher
type Evpn_Active_EviDetail_Elements_Element_RdConfigured struct {
    parent types.Entity
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

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetFilter() yfilter.YFilter { return rdConfigured.YFilter }

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) SetFilter(yf yfilter.YFilter) { rdConfigured.YFilter = yf }

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetGoName(yname string) string {
    if yname == "rd" { return "Rd" }
    if yname == "auto" { return "Auto" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    return ""
}

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetSegmentPath() string {
    return "rd-configured"
}

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto" {
        return &rdConfigured.Auto
    }
    if childYangName == "two-byte-as" {
        return &rdConfigured.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rdConfigured.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rdConfigured.V4Addr
    }
    return nil
}

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto"] = &rdConfigured.Auto
    children["two-byte-as"] = &rdConfigured.TwoByteAs
    children["four-byte-as"] = &rdConfigured.FourByteAs
    children["v4-addr"] = &rdConfigured.V4Addr
    return children
}

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rd"] = rdConfigured.Rd
    return leafs
}

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetBundleName() string { return "cisco_ios_xr" }

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetYangName() string { return "rd-configured" }

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) SetParent(parent types.Entity) { rdConfigured.parent = parent }

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetParent() types.Entity { return rdConfigured.parent }

func (rdConfigured *Evpn_Active_EviDetail_Elements_Element_RdConfigured) GetParentYangName() string { return "element" }

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto
// auto
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetFilter() yfilter.YFilter { return auto.YFilter }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) SetFilter(yf yfilter.YFilter) { auto.YFilter = yf }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "auto-index" { return "AutoIndex" }
    return ""
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetSegmentPath() string {
    return "auto"
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = auto.RouterId
    leafs["auto-index"] = auto.AutoIndex
    return leafs
}

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetBundleName() string { return "cisco_ios_xr" }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetYangName() string { return "auto" }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) SetParent(parent types.Entity) { auto.parent = parent }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetParent() types.Entity { return auto.parent }

func (auto *Evpn_Active_EviDetail_Elements_Element_RdConfigured_Auto) GetParentYangName() string { return "rd-configured" }

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetParentYangName() string { return "rd-configured" }

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs
// four byte as
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetParentYangName() string { return "rd-configured" }

// Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RdConfigured_V4Addr) GetParentYangName() string { return "rd-configured" }

// Evpn_Active_EviDetail_Elements_Element_RtAuto
// Automatic Route Target
type Evpn_Active_EviDetail_Elements_Element_RtAuto struct {
    parent types.Entity
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

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetFilter() yfilter.YFilter { return rtAuto.YFilter }

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) SetFilter(yf yfilter.YFilter) { rtAuto.YFilter = yf }

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetGoName(yname string) string {
    if yname == "rt" { return "Rt" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    if yname == "es-import" { return "EsImport" }
    return ""
}

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetSegmentPath() string {
    return "rt-auto"
}

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "two-byte-as" {
        return &rtAuto.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rtAuto.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rtAuto.V4Addr
    }
    if childYangName == "es-import" {
        return &rtAuto.EsImport
    }
    return nil
}

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["two-byte-as"] = &rtAuto.TwoByteAs
    children["four-byte-as"] = &rtAuto.FourByteAs
    children["v4-addr"] = &rtAuto.V4Addr
    children["es-import"] = &rtAuto.EsImport
    return children
}

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt"] = rtAuto.Rt
    return leafs
}

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetBundleName() string { return "cisco_ios_xr" }

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetYangName() string { return "rt-auto" }

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) SetParent(parent types.Entity) { rtAuto.parent = parent }

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetParent() types.Entity { return rtAuto.parent }

func (rtAuto *Evpn_Active_EviDetail_Elements_Element_RtAuto) GetParentYangName() string { return "element" }

// Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetParentYangName() string { return "rt-auto" }

// Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs
// four byte as
type Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAuto_FourByteAs) GetParentYangName() string { return "rt-auto" }

// Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAuto_V4Addr) GetParentYangName() string { return "rt-auto" }

// Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport
// es import
type Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetFilter() yfilter.YFilter { return esImport.YFilter }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) SetFilter(yf yfilter.YFilter) { esImport.YFilter = yf }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetGoName(yname string) string {
    if yname == "high-bytes" { return "HighBytes" }
    if yname == "low-bytes" { return "LowBytes" }
    return ""
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetSegmentPath() string {
    return "es-import"
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-bytes"] = esImport.HighBytes
    leafs["low-bytes"] = esImport.LowBytes
    return leafs
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetBundleName() string { return "cisco_ios_xr" }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetYangName() string { return "es-import" }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) SetParent(parent types.Entity) { esImport.parent = parent }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetParent() types.Entity { return esImport.parent }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAuto_EsImport) GetParentYangName() string { return "rt-auto" }

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching
// Automatic Route Target Stitching
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching struct {
    parent types.Entity
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

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetFilter() yfilter.YFilter { return rtAutoStitching.YFilter }

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) SetFilter(yf yfilter.YFilter) { rtAutoStitching.YFilter = yf }

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetGoName(yname string) string {
    if yname == "rt" { return "Rt" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    if yname == "es-import" { return "EsImport" }
    return ""
}

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetSegmentPath() string {
    return "rt-auto-stitching"
}

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "two-byte-as" {
        return &rtAutoStitching.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rtAutoStitching.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rtAutoStitching.V4Addr
    }
    if childYangName == "es-import" {
        return &rtAutoStitching.EsImport
    }
    return nil
}

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["two-byte-as"] = &rtAutoStitching.TwoByteAs
    children["four-byte-as"] = &rtAutoStitching.FourByteAs
    children["v4-addr"] = &rtAutoStitching.V4Addr
    children["es-import"] = &rtAutoStitching.EsImport
    return children
}

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt"] = rtAutoStitching.Rt
    return leafs
}

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetBundleName() string { return "cisco_ios_xr" }

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetYangName() string { return "rt-auto-stitching" }

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) SetParent(parent types.Entity) { rtAutoStitching.parent = parent }

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetParent() types.Entity { return rtAutoStitching.parent }

func (rtAutoStitching *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching) GetParentYangName() string { return "element" }

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs
// four byte as
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr
// v4 addr
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport
// es import
type Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetFilter() yfilter.YFilter { return esImport.YFilter }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) SetFilter(yf yfilter.YFilter) { esImport.YFilter = yf }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetGoName(yname string) string {
    if yname == "high-bytes" { return "HighBytes" }
    if yname == "low-bytes" { return "LowBytes" }
    return ""
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetSegmentPath() string {
    return "es-import"
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-bytes"] = esImport.HighBytes
    leafs["low-bytes"] = esImport.LowBytes
    return leafs
}

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetBundleName() string { return "cisco_ios_xr" }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetYangName() string { return "es-import" }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) SetParent(parent types.Entity) { esImport.parent = parent }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetParent() types.Entity { return esImport.parent }

func (esImport *Evpn_Active_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Active_EviDetail_EviChildren
// Container for all EVI detail info
type Evpn_Active_EviDetail_EviChildren struct {
    parent types.Entity
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

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetFilter() yfilter.YFilter { return eviChildren.YFilter }

func (eviChildren *Evpn_Active_EviDetail_EviChildren) SetFilter(yf yfilter.YFilter) { eviChildren.YFilter = yf }

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetGoName(yname string) string {
    if yname == "neighbors" { return "Neighbors" }
    if yname == "ethernet-auto-discoveries" { return "EthernetAutoDiscoveries" }
    if yname == "inclusive-multicasts" { return "InclusiveMulticasts" }
    if yname == "route-targets" { return "RouteTargets" }
    if yname == "macs" { return "Macs" }
    return ""
}

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetSegmentPath() string {
    return "evi-children"
}

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbors" {
        return &eviChildren.Neighbors
    }
    if childYangName == "ethernet-auto-discoveries" {
        return &eviChildren.EthernetAutoDiscoveries
    }
    if childYangName == "inclusive-multicasts" {
        return &eviChildren.InclusiveMulticasts
    }
    if childYangName == "route-targets" {
        return &eviChildren.RouteTargets
    }
    if childYangName == "macs" {
        return &eviChildren.Macs
    }
    return nil
}

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbors"] = &eviChildren.Neighbors
    children["ethernet-auto-discoveries"] = &eviChildren.EthernetAutoDiscoveries
    children["inclusive-multicasts"] = &eviChildren.InclusiveMulticasts
    children["route-targets"] = &eviChildren.RouteTargets
    children["macs"] = &eviChildren.Macs
    return children
}

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetBundleName() string { return "cisco_ios_xr" }

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetYangName() string { return "evi-children" }

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eviChildren *Evpn_Active_EviDetail_EviChildren) SetParent(parent types.Entity) { eviChildren.parent = parent }

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetParent() types.Entity { return eviChildren.parent }

func (eviChildren *Evpn_Active_EviDetail_EviChildren) GetParentYangName() string { return "evi-detail" }

// Evpn_Active_EviDetail_EviChildren_Neighbors
// EVPN Neighbor table
type Evpn_Active_EviDetail_EviChildren_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Neighbor table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor.
    Neighbor []Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor
}

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Evpn_Active_EviDetail_EviChildren_Neighbors) GetParentYangName() string { return "evi-children" }

// Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor
// EVPN Neighbor table
type Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Neighbor IP. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}
}

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "neighbor-ip" { return "NeighborIp" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = neighbor.Evi
    leafs["neighbor-ip"] = neighbor.NeighborIp
    leafs["evi-xr"] = neighbor.EviXr
    leafs["neighbor"] = neighbor.Neighbor
    return leafs
}

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Evpn_Active_EviDetail_EviChildren_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries
// EVPN Ethernet Auto-Discovery table
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Ethernet Auto-Discovery Entry. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery.
    EthernetAutoDiscovery []Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
}

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetFilter() yfilter.YFilter { return ethernetAutoDiscoveries.YFilter }

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) SetFilter(yf yfilter.YFilter) { ethernetAutoDiscoveries.YFilter = yf }

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetGoName(yname string) string {
    if yname == "ethernet-auto-discovery" { return "EthernetAutoDiscovery" }
    return ""
}

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetSegmentPath() string {
    return "ethernet-auto-discoveries"
}

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-auto-discovery" {
        for _, c := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
            if ethernetAutoDiscoveries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery{}
        ethernetAutoDiscoveries.EthernetAutoDiscovery = append(ethernetAutoDiscoveries.EthernetAutoDiscovery, child)
        return &ethernetAutoDiscoveries.EthernetAutoDiscovery[len(ethernetAutoDiscoveries.EthernetAutoDiscovery)-1]
    }
    return nil
}

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
        children[ethernetAutoDiscoveries.EthernetAutoDiscovery[i].GetSegmentPath()] = &ethernetAutoDiscoveries.EthernetAutoDiscovery[i]
    }
    return children
}

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetYangName() string { return "ethernet-auto-discoveries" }

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) SetParent(parent types.Entity) { ethernetAutoDiscoveries.parent = parent }

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetParent() types.Entity { return ethernetAutoDiscoveries.parent }

func (ethernetAutoDiscoveries *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries) GetParentYangName() string { return "evi-children" }

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
// EVPN Ethernet Auto-Discovery Entry
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

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

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnid interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Local nexthop IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetFilter() yfilter.YFilter { return ethernetAutoDiscovery.YFilter }

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) SetFilter(yf yfilter.YFilter) { ethernetAutoDiscovery.YFilter = yf }

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "esi1" { return "Esi1" }
    if yname == "esi2" { return "Esi2" }
    if yname == "esi3" { return "Esi3" }
    if yname == "esi4" { return "Esi4" }
    if yname == "esi5" { return "Esi5" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    if yname == "ethernet-vpnid" { return "EthernetVpnid" }
    if yname == "type" { return "Type" }
    if yname == "ethernet-tag-xr" { return "EthernetTagXr" }
    if yname == "local-next-hop" { return "LocalNextHop" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "is-local-ead" { return "IsLocalEad" }
    if yname == "encap" { return "Encap" }
    if yname == "redundancy-single-active" { return "RedundancySingleActive" }
    if yname == "num-paths" { return "NumPaths" }
    if yname == "ethernet-segment-identifier" { return "EthernetSegmentIdentifier" }
    if yname == "path-buffer" { return "PathBuffer" }
    return ""
}

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetSegmentPath() string {
    return "ethernet-auto-discovery"
}

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-identifier" {
        for _, c := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
            if ethernetAutoDiscovery.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier{}
        ethernetAutoDiscovery.EthernetSegmentIdentifier = append(ethernetAutoDiscovery.EthernetSegmentIdentifier, child)
        return &ethernetAutoDiscovery.EthernetSegmentIdentifier[len(ethernetAutoDiscovery.EthernetSegmentIdentifier)-1]
    }
    if childYangName == "path-buffer" {
        for _, c := range ethernetAutoDiscovery.PathBuffer {
            if ethernetAutoDiscovery.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer{}
        ethernetAutoDiscovery.PathBuffer = append(ethernetAutoDiscovery.PathBuffer, child)
        return &ethernetAutoDiscovery.PathBuffer[len(ethernetAutoDiscovery.PathBuffer)-1]
    }
    return nil
}

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
        children[ethernetAutoDiscovery.EthernetSegmentIdentifier[i].GetSegmentPath()] = &ethernetAutoDiscovery.EthernetSegmentIdentifier[i]
    }
    for i := range ethernetAutoDiscovery.PathBuffer {
        children[ethernetAutoDiscovery.PathBuffer[i].GetSegmentPath()] = &ethernetAutoDiscovery.PathBuffer[i]
    }
    return children
}

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = ethernetAutoDiscovery.Evi
    leafs["esi1"] = ethernetAutoDiscovery.Esi1
    leafs["esi2"] = ethernetAutoDiscovery.Esi2
    leafs["esi3"] = ethernetAutoDiscovery.Esi3
    leafs["esi4"] = ethernetAutoDiscovery.Esi4
    leafs["esi5"] = ethernetAutoDiscovery.Esi5
    leafs["ethernet-tag"] = ethernetAutoDiscovery.EthernetTag
    leafs["ethernet-vpnid"] = ethernetAutoDiscovery.EthernetVpnid
    leafs["type"] = ethernetAutoDiscovery.Type
    leafs["ethernet-tag-xr"] = ethernetAutoDiscovery.EthernetTagXr
    leafs["local-next-hop"] = ethernetAutoDiscovery.LocalNextHop
    leafs["local-label"] = ethernetAutoDiscovery.LocalLabel
    leafs["is-local-ead"] = ethernetAutoDiscovery.IsLocalEad
    leafs["encap"] = ethernetAutoDiscovery.Encap
    leafs["redundancy-single-active"] = ethernetAutoDiscovery.RedundancySingleActive
    leafs["num-paths"] = ethernetAutoDiscovery.NumPaths
    return leafs
}

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetYangName() string { return "ethernet-auto-discovery" }

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) SetParent(parent types.Entity) { ethernetAutoDiscovery.parent = parent }

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetParent() types.Entity { return ethernetAutoDiscovery.parent }

func (ethernetAutoDiscovery *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetParentYangName() string { return "ethernet-auto-discoveries" }

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return ethernetSegmentIdentifier.YFilter }

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { ethernetSegmentIdentifier.YFilter = yf }

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetSegmentPath() string {
    return "ethernet-segment-identifier"
}

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = ethernetSegmentIdentifier.Entry
    return leafs
}

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetYangName() string { return "ethernet-segment-identifier" }

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) SetParent(parent types.Entity) { ethernetSegmentIdentifier.parent = parent }

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetParent() types.Entity { return ethernetSegmentIdentifier.parent }

func (ethernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetParentYangName() string { return "ethernet-auto-discovery" }

// Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
// Path List Buffer
type Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetFilter() yfilter.YFilter { return pathBuffer.YFilter }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) SetFilter(yf yfilter.YFilter) { pathBuffer.YFilter = yf }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "output-label" { return "OutputLabel" }
    if yname == "srte-tunnel" { return "SrteTunnel" }
    return ""
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetSegmentPath() string {
    return "path-buffer"
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = pathBuffer.NextHop
    leafs["output-label"] = pathBuffer.OutputLabel
    leafs["srte-tunnel"] = pathBuffer.SrteTunnel
    return leafs
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetYangName() string { return "path-buffer" }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) SetParent(parent types.Entity) { pathBuffer.parent = parent }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetParent() types.Entity { return pathBuffer.parent }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetParentYangName() string { return "ethernet-auto-discovery" }

// Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts
// L2VPN EVPN IMCAST table
type Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN IMCAST table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast.
    InclusiveMulticast []Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
}

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetFilter() yfilter.YFilter { return inclusiveMulticasts.YFilter }

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) SetFilter(yf yfilter.YFilter) { inclusiveMulticasts.YFilter = yf }

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetGoName(yname string) string {
    if yname == "inclusive-multicast" { return "InclusiveMulticast" }
    return ""
}

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetSegmentPath() string {
    return "inclusive-multicasts"
}

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inclusive-multicast" {
        for _, c := range inclusiveMulticasts.InclusiveMulticast {
            if inclusiveMulticasts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast{}
        inclusiveMulticasts.InclusiveMulticast = append(inclusiveMulticasts.InclusiveMulticast, child)
        return &inclusiveMulticasts.InclusiveMulticast[len(inclusiveMulticasts.InclusiveMulticast)-1]
    }
    return nil
}

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inclusiveMulticasts.InclusiveMulticast {
        children[inclusiveMulticasts.InclusiveMulticast[i].GetSegmentPath()] = &inclusiveMulticasts.InclusiveMulticast[i]
    }
    return children
}

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetBundleName() string { return "cisco_ios_xr" }

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetYangName() string { return "inclusive-multicasts" }

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) SetParent(parent types.Entity) { inclusiveMulticasts.parent = parent }

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetParent() types.Entity { return inclusiveMulticasts.parent }

func (inclusiveMulticasts *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts) GetParentYangName() string { return "evi-children" }

// Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
// L2VPN EVPN IMCAST table
type Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag. The type is interface{} with range: -2147483648..2147483647.
    EthernetTag interface{}

    // Originating IP. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Originating IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIpXr interface{}

    // IP of nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetFilter() yfilter.YFilter { return inclusiveMulticast.YFilter }

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) SetFilter(yf yfilter.YFilter) { inclusiveMulticast.YFilter = yf }

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    if yname == "originating-ip" { return "OriginatingIp" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "ethernet-tag-xr" { return "EthernetTagXr" }
    if yname == "originating-ip-xr" { return "OriginatingIpXr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "output-label" { return "OutputLabel" }
    if yname == "is-local-entry" { return "IsLocalEntry" }
    if yname == "is-proxy-entry" { return "IsProxyEntry" }
    if yname == "encap-type" { return "EncapType" }
    return ""
}

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetSegmentPath() string {
    return "inclusive-multicast"
}

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = inclusiveMulticast.Evi
    leafs["ethernet-tag"] = inclusiveMulticast.EthernetTag
    leafs["originating-ip"] = inclusiveMulticast.OriginatingIp
    leafs["evi-xr"] = inclusiveMulticast.EviXr
    leafs["ethernet-tag-xr"] = inclusiveMulticast.EthernetTagXr
    leafs["originating-ip-xr"] = inclusiveMulticast.OriginatingIpXr
    leafs["next-hop"] = inclusiveMulticast.NextHop
    leafs["output-label"] = inclusiveMulticast.OutputLabel
    leafs["is-local-entry"] = inclusiveMulticast.IsLocalEntry
    leafs["is-proxy-entry"] = inclusiveMulticast.IsProxyEntry
    leafs["encap-type"] = inclusiveMulticast.EncapType
    return leafs
}

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetBundleName() string { return "cisco_ios_xr" }

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetYangName() string { return "inclusive-multicast" }

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) SetParent(parent types.Entity) { inclusiveMulticast.parent = parent }

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetParent() types.Entity { return inclusiveMulticast.parent }

func (inclusiveMulticast *Evpn_Active_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetParentYangName() string { return "inclusive-multicasts" }

// Evpn_Active_EviDetail_EviChildren_RouteTargets
// L2VPN EVPN EVI RT Child Table
type Evpn_Active_EviDetail_EviChildren_RouteTargets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI RT Table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget.
    RouteTarget []Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget
}

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetFilter() yfilter.YFilter { return routeTargets.YFilter }

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) SetFilter(yf yfilter.YFilter) { routeTargets.YFilter = yf }

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetGoName(yname string) string {
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetSegmentPath() string {
    return "route-targets"
}

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        for _, c := range routeTargets.RouteTarget {
            if routeTargets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget{}
        routeTargets.RouteTarget = append(routeTargets.RouteTarget, child)
        return &routeTargets.RouteTarget[len(routeTargets.RouteTarget)-1]
    }
    return nil
}

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routeTargets.RouteTarget {
        children[routeTargets.RouteTarget[i].GetSegmentPath()] = &routeTargets.RouteTarget[i]
    }
    return children
}

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetBundleName() string { return "cisco_ios_xr" }

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetYangName() string { return "route-targets" }

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) SetParent(parent types.Entity) { routeTargets.parent = parent }

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetParent() types.Entity { return routeTargets.parent }

func (routeTargets *Evpn_Active_EviDetail_EviChildren_RouteTargets) GetParentYangName() string { return "evi-children" }

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget
// L2VPN EVPN EVI RT Table
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Role of the route target. The type is BgpRouteTargetRole.
    Role interface{}

    // Type of the route target. The type is BgpRouteTarget.
    Type interface{}

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

    // Bridge Domain Name. The type is string.
    BdName interface{}

    // VPN ID. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // RT Role. The type is L2vpnAdRtRole.
    RouteTargetRole interface{}

    // RT Stitching. The type is bool.
    RouteTargetStitching interface{}

    // Route Target.
    RouteTarget Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "role" { return "Role" }
    if yname == "type" { return "Type" }
    if yname == "format" { return "Format" }
    if yname == "as" { return "As" }
    if yname == "as-index" { return "AsIndex" }
    if yname == "addr-index" { return "AddrIndex" }
    if yname == "address" { return "Address" }
    if yname == "bd-name" { return "BdName" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "route-target-role" { return "RouteTargetRole" }
    if yname == "route-target-stitching" { return "RouteTargetStitching" }
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetSegmentPath() string {
    return "route-target"
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        return &routeTarget.RouteTarget
    }
    return nil
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route-target"] = &routeTarget.RouteTarget
    return children
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = routeTarget.Evi
    leafs["role"] = routeTarget.Role
    leafs["type"] = routeTarget.Type
    leafs["format"] = routeTarget.Format
    leafs["as"] = routeTarget.As
    leafs["as-index"] = routeTarget.AsIndex
    leafs["addr-index"] = routeTarget.AddrIndex
    leafs["address"] = routeTarget.Address
    leafs["bd-name"] = routeTarget.BdName
    leafs["evi-xr"] = routeTarget.EviXr
    leafs["route-target-role"] = routeTarget.RouteTargetRole
    leafs["route-target-stitching"] = routeTarget.RouteTargetStitching
    return leafs
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget) GetParentYangName() string { return "route-targets" }

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
// Route Target
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget struct {
    parent types.Entity
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

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetGoName(yname string) string {
    if yname == "rt" { return "Rt" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    if yname == "es-import" { return "EsImport" }
    return ""
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetSegmentPath() string {
    return "route-target"
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "two-byte-as" {
        return &routeTarget.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &routeTarget.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &routeTarget.V4Addr
    }
    if childYangName == "es-import" {
        return &routeTarget.EsImport
    }
    return nil
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["two-byte-as"] = &routeTarget.TwoByteAs
    children["four-byte-as"] = &routeTarget.FourByteAs
    children["v4-addr"] = &routeTarget.V4Addr
    children["es-import"] = &routeTarget.EsImport
    return children
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt"] = routeTarget.Rt
    return leafs
}

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetParentYangName() string { return "route-target" }

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs
// two byte as
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetParentYangName() string { return "route-target" }

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs
// four byte as
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetParentYangName() string { return "route-target" }

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr
// v4 addr
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetParentYangName() string { return "route-target" }

// Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport
// es import
type Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetFilter() yfilter.YFilter { return esImport.YFilter }

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) SetFilter(yf yfilter.YFilter) { esImport.YFilter = yf }

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetGoName(yname string) string {
    if yname == "high-bytes" { return "HighBytes" }
    if yname == "low-bytes" { return "LowBytes" }
    return ""
}

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetSegmentPath() string {
    return "es-import"
}

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-bytes"] = esImport.HighBytes
    leafs["low-bytes"] = esImport.LowBytes
    return leafs
}

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetBundleName() string { return "cisco_ios_xr" }

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetYangName() string { return "es-import" }

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) SetParent(parent types.Entity) { esImport.parent = parent }

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetParent() types.Entity { return esImport.parent }

func (esImport *Evpn_Active_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetParentYangName() string { return "route-target" }

// Evpn_Active_EviDetail_EviChildren_Macs
// L2VPN EVPN EVI MAC table
type Evpn_Active_EviDetail_EviChildren_Macs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN MAC table. The type is slice of
    // Evpn_Active_EviDetail_EviChildren_Macs_Mac.
    Mac []Evpn_Active_EviDetail_EviChildren_Macs_Mac
}

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetFilter() yfilter.YFilter { return macs.YFilter }

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) SetFilter(yf yfilter.YFilter) { macs.YFilter = yf }

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetGoName(yname string) string {
    if yname == "mac" { return "Mac" }
    return ""
}

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetSegmentPath() string {
    return "macs"
}

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac" {
        for _, c := range macs.Mac {
            if macs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_Macs_Mac{}
        macs.Mac = append(macs.Mac, child)
        return &macs.Mac[len(macs.Mac)-1]
    }
    return nil
}

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macs.Mac {
        children[macs.Mac[i].GetSegmentPath()] = &macs.Mac[i]
    }
    return children
}

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetBundleName() string { return "cisco_ios_xr" }

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetYangName() string { return "macs" }

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) SetParent(parent types.Entity) { macs.parent = parent }

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetParent() types.Entity { return macs.parent }

func (macs *Evpn_Active_EviDetail_EviChildren_Macs) GetParentYangName() string { return "evi-children" }

// Evpn_Active_EviDetail_EviChildren_Macs_Mac
// L2VPN EVPN MAC table
type Evpn_Active_EviDetail_EviChildren_Macs_Mac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
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

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "ethernet-tag-xr" { return "EthernetTagXr" }
    if yname == "mac-address-xr" { return "MacAddressXr" }
    if yname == "ip-address-xr" { return "IpAddressXr" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "num-paths" { return "NumPaths" }
    if yname == "is-local-mac" { return "IsLocalMac" }
    if yname == "is-proxy-entry" { return "IsProxyEntry" }
    if yname == "is-remote-mac" { return "IsRemoteMac" }
    if yname == "soo-nexthop" { return "SooNexthop" }
    if yname == "ipnh-address" { return "IpnhAddress" }
    if yname == "esi-port-key" { return "EsiPortKey" }
    if yname == "local-encap-type" { return "LocalEncapType" }
    if yname == "remote-encap-type" { return "RemoteEncapType" }
    if yname == "learned-bridge-port-name" { return "LearnedBridgePortName" }
    if yname == "local-seq-id" { return "LocalSeqId" }
    if yname == "remote-seq-id" { return "RemoteSeqId" }
    if yname == "local-l3-label" { return "LocalL3Label" }
    if yname == "router-mac-address" { return "RouterMacAddress" }
    if yname == "mac-flush-requested" { return "MacFlushRequested" }
    if yname == "mac-flush-received" { return "MacFlushReceived" }
    if yname == "internal-label" { return "InternalLabel" }
    if yname == "resolved" { return "Resolved" }
    if yname == "local-is-static" { return "LocalIsStatic" }
    if yname == "remote-is-static" { return "RemoteIsStatic" }
    if yname == "local-ethernet-segment-identifier" { return "LocalEthernetSegmentIdentifier" }
    if yname == "remote-ethernet-segment-identifier" { return "RemoteEthernetSegmentIdentifier" }
    if yname == "path-buffer" { return "PathBuffer" }
    return ""
}

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-ethernet-segment-identifier" {
        for _, c := range mac.LocalEthernetSegmentIdentifier {
            if mac.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier{}
        mac.LocalEthernetSegmentIdentifier = append(mac.LocalEthernetSegmentIdentifier, child)
        return &mac.LocalEthernetSegmentIdentifier[len(mac.LocalEthernetSegmentIdentifier)-1]
    }
    if childYangName == "remote-ethernet-segment-identifier" {
        for _, c := range mac.RemoteEthernetSegmentIdentifier {
            if mac.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier{}
        mac.RemoteEthernetSegmentIdentifier = append(mac.RemoteEthernetSegmentIdentifier, child)
        return &mac.RemoteEthernetSegmentIdentifier[len(mac.RemoteEthernetSegmentIdentifier)-1]
    }
    if childYangName == "path-buffer" {
        for _, c := range mac.PathBuffer {
            if mac.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer{}
        mac.PathBuffer = append(mac.PathBuffer, child)
        return &mac.PathBuffer[len(mac.PathBuffer)-1]
    }
    return nil
}

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mac.LocalEthernetSegmentIdentifier {
        children[mac.LocalEthernetSegmentIdentifier[i].GetSegmentPath()] = &mac.LocalEthernetSegmentIdentifier[i]
    }
    for i := range mac.RemoteEthernetSegmentIdentifier {
        children[mac.RemoteEthernetSegmentIdentifier[i].GetSegmentPath()] = &mac.RemoteEthernetSegmentIdentifier[i]
    }
    for i := range mac.PathBuffer {
        children[mac.PathBuffer[i].GetSegmentPath()] = &mac.PathBuffer[i]
    }
    return children
}

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = mac.Evi
    leafs["ethernet-tag"] = mac.EthernetTag
    leafs["mac-address"] = mac.MacAddress
    leafs["ip-address"] = mac.IpAddress
    leafs["ethernet-tag-xr"] = mac.EthernetTagXr
    leafs["mac-address-xr"] = mac.MacAddressXr
    leafs["ip-address-xr"] = mac.IpAddressXr
    leafs["local-label"] = mac.LocalLabel
    leafs["num-paths"] = mac.NumPaths
    leafs["is-local-mac"] = mac.IsLocalMac
    leafs["is-proxy-entry"] = mac.IsProxyEntry
    leafs["is-remote-mac"] = mac.IsRemoteMac
    leafs["soo-nexthop"] = mac.SooNexthop
    leafs["ipnh-address"] = mac.IpnhAddress
    leafs["esi-port-key"] = mac.EsiPortKey
    leafs["local-encap-type"] = mac.LocalEncapType
    leafs["remote-encap-type"] = mac.RemoteEncapType
    leafs["learned-bridge-port-name"] = mac.LearnedBridgePortName
    leafs["local-seq-id"] = mac.LocalSeqId
    leafs["remote-seq-id"] = mac.RemoteSeqId
    leafs["local-l3-label"] = mac.LocalL3Label
    leafs["router-mac-address"] = mac.RouterMacAddress
    leafs["mac-flush-requested"] = mac.MacFlushRequested
    leafs["mac-flush-received"] = mac.MacFlushReceived
    leafs["internal-label"] = mac.InternalLabel
    leafs["resolved"] = mac.Resolved
    leafs["local-is-static"] = mac.LocalIsStatic
    leafs["remote-is-static"] = mac.RemoteIsStatic
    return leafs
}

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetYangName() string { return "mac" }

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetParent() types.Entity { return mac.parent }

func (mac *Evpn_Active_EviDetail_EviChildren_Macs_Mac) GetParentYangName() string { return "macs" }

// Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier
// Local Ethernet Segment id
type Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return localEthernetSegmentIdentifier.YFilter }

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { localEthernetSegmentIdentifier.YFilter = yf }

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetSegmentPath() string {
    return "local-ethernet-segment-identifier"
}

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = localEthernetSegmentIdentifier.Entry
    return leafs
}

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetYangName() string { return "local-ethernet-segment-identifier" }

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) SetParent(parent types.Entity) { localEthernetSegmentIdentifier.parent = parent }

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetParent() types.Entity { return localEthernetSegmentIdentifier.parent }

func (localEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetParentYangName() string { return "mac" }

// Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier
// Remote Ethernet Segment id
type Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return remoteEthernetSegmentIdentifier.YFilter }

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { remoteEthernetSegmentIdentifier.YFilter = yf }

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetSegmentPath() string {
    return "remote-ethernet-segment-identifier"
}

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = remoteEthernetSegmentIdentifier.Entry
    return leafs
}

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetYangName() string { return "remote-ethernet-segment-identifier" }

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) SetParent(parent types.Entity) { remoteEthernetSegmentIdentifier.parent = parent }

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetParent() types.Entity { return remoteEthernetSegmentIdentifier.parent }

func (remoteEthernetSegmentIdentifier *Evpn_Active_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetParentYangName() string { return "mac" }

// Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer
// Path List Buffer
type Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetFilter() yfilter.YFilter { return pathBuffer.YFilter }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) SetFilter(yf yfilter.YFilter) { pathBuffer.YFilter = yf }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "output-label" { return "OutputLabel" }
    if yname == "srte-tunnel" { return "SrteTunnel" }
    return ""
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetSegmentPath() string {
    return "path-buffer"
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = pathBuffer.NextHop
    leafs["output-label"] = pathBuffer.OutputLabel
    leafs["srte-tunnel"] = pathBuffer.SrteTunnel
    return leafs
}

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetYangName() string { return "path-buffer" }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) SetParent(parent types.Entity) { pathBuffer.parent = parent }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetParent() types.Entity { return pathBuffer.parent }

func (pathBuffer *Evpn_Active_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetParentYangName() string { return "mac" }

// Evpn_Active_EthernetSegments
// EVPN Ethernet-Segment Table
type Evpn_Active_EthernetSegments struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Ethernet-Segment Entry. The type is slice of
    // Evpn_Active_EthernetSegments_EthernetSegment.
    EthernetSegment []Evpn_Active_EthernetSegments_EthernetSegment
}

func (ethernetSegments *Evpn_Active_EthernetSegments) GetFilter() yfilter.YFilter { return ethernetSegments.YFilter }

func (ethernetSegments *Evpn_Active_EthernetSegments) SetFilter(yf yfilter.YFilter) { ethernetSegments.YFilter = yf }

func (ethernetSegments *Evpn_Active_EthernetSegments) GetGoName(yname string) string {
    if yname == "ethernet-segment" { return "EthernetSegment" }
    return ""
}

func (ethernetSegments *Evpn_Active_EthernetSegments) GetSegmentPath() string {
    return "ethernet-segments"
}

func (ethernetSegments *Evpn_Active_EthernetSegments) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment" {
        for _, c := range ethernetSegments.EthernetSegment {
            if ethernetSegments.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment{}
        ethernetSegments.EthernetSegment = append(ethernetSegments.EthernetSegment, child)
        return &ethernetSegments.EthernetSegment[len(ethernetSegments.EthernetSegment)-1]
    }
    return nil
}

func (ethernetSegments *Evpn_Active_EthernetSegments) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetSegments.EthernetSegment {
        children[ethernetSegments.EthernetSegment[i].GetSegmentPath()] = &ethernetSegments.EthernetSegment[i]
    }
    return children
}

func (ethernetSegments *Evpn_Active_EthernetSegments) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetSegments *Evpn_Active_EthernetSegments) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegments *Evpn_Active_EthernetSegments) GetYangName() string { return "ethernet-segments" }

func (ethernetSegments *Evpn_Active_EthernetSegments) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegments *Evpn_Active_EthernetSegments) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegments *Evpn_Active_EthernetSegments) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegments *Evpn_Active_EthernetSegments) SetParent(parent types.Entity) { ethernetSegments.parent = parent }

func (ethernetSegments *Evpn_Active_EthernetSegments) GetParent() types.Entity { return ethernetSegments.parent }

func (ethernetSegments *Evpn_Active_EthernetSegments) GetParentYangName() string { return "active" }

// Evpn_Active_EthernetSegments_EthernetSegment
// EVPN Ethernet-Segment Entry
type Evpn_Active_EthernetSegments_EthernetSegment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
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

    // Ethernet Segment Name. The type is string.
    EthernetSegmentName interface{}

    // State of the ethernet segment. The type is interface{} with range:
    // 0..4294967295.
    EthernetSegmentState interface{}

    // Main port ifhandle. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetFilter() yfilter.YFilter { return ethernetSegment.YFilter }

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) SetFilter(yf yfilter.YFilter) { ethernetSegment.YFilter = yf }

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "esi1" { return "Esi1" }
    if yname == "esi2" { return "Esi2" }
    if yname == "esi3" { return "Esi3" }
    if yname == "esi4" { return "Esi4" }
    if yname == "esi5" { return "Esi5" }
    if yname == "esi-type" { return "EsiType" }
    if yname == "ethernet-segment-name" { return "EthernetSegmentName" }
    if yname == "ethernet-segment-state" { return "EthernetSegmentState" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "main-port-role" { return "MainPortRole" }
    if yname == "main-port-mac" { return "MainPortMac" }
    if yname == "num-up-p-ws" { return "NumUpPWs" }
    if yname == "route-target" { return "RouteTarget" }
    if yname == "rt-origin" { return "RtOrigin" }
    if yname == "es-bgp-gates" { return "EsBgpGates" }
    if yname == "es-l2fib-gates" { return "EsL2FibGates" }
    if yname == "mac-flushing-mode-config" { return "MacFlushingModeConfig" }
    if yname == "load-balance-mode-config" { return "LoadBalanceModeConfig" }
    if yname == "load-balance-mode-is-default" { return "LoadBalanceModeIsDefault" }
    if yname == "load-balance-mode-oper" { return "LoadBalanceModeOper" }
    if yname == "force-single-home" { return "ForceSingleHome" }
    if yname == "source-mac-oper" { return "SourceMacOper" }
    if yname == "source-mac-origin" { return "SourceMacOrigin" }
    if yname == "peering-timer" { return "PeeringTimer" }
    if yname == "peering-timer-left" { return "PeeringTimerLeft" }
    if yname == "recovery-timer" { return "RecoveryTimer" }
    if yname == "recovery-timer-left" { return "RecoveryTimerLeft" }
    if yname == "service-carving-mode" { return "ServiceCarvingMode" }
    if yname == "primary-services-input" { return "PrimaryServicesInput" }
    if yname == "secondary-services-input" { return "SecondaryServicesInput" }
    if yname == "forwarder-ports" { return "ForwarderPorts" }
    if yname == "permanent-forwarder-ports" { return "PermanentForwarderPorts" }
    if yname == "elected-forwarder-ports" { return "ElectedForwarderPorts" }
    if yname == "not-elected-forwarder-ports" { return "NotElectedForwarderPorts" }
    if yname == "not-config-forwarder-ports" { return "NotConfigForwarderPorts" }
    if yname == "mp-protected" { return "MpProtected" }
    if yname == "nve-anycast-vtep" { return "NveAnycastVtep" }
    if yname == "nve-ingress-replication" { return "NveIngressReplication" }
    if yname == "local-split-horizon-group-label" { return "LocalSplitHorizonGroupLabel" }
    if yname == "ethernet-segment-identifier" { return "EthernetSegmentIdentifier" }
    if yname == "primary-service" { return "PrimaryService" }
    if yname == "secondary-service" { return "SecondaryService" }
    if yname == "service-carving-i-sidelected-result" { return "ServiceCarvingISidelectedResult" }
    if yname == "service-carving-isid-not-elected-result" { return "ServiceCarvingIsidNotElectedResult" }
    if yname == "service-carving-evi-elected-result" { return "ServiceCarvingEviElectedResult" }
    if yname == "service-carving-evi-not-elected-result" { return "ServiceCarvingEviNotElectedResult" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "service-carving-vpws-permanent-result" { return "ServiceCarvingVpwsPermanentResult" }
    if yname == "remote-split-horizon-group-label" { return "RemoteSplitHorizonGroupLabel" }
    return ""
}

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetSegmentPath() string {
    return "ethernet-segment"
}

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-identifier" {
        for _, c := range ethernetSegment.EthernetSegmentIdentifier {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier{}
        ethernetSegment.EthernetSegmentIdentifier = append(ethernetSegment.EthernetSegmentIdentifier, child)
        return &ethernetSegment.EthernetSegmentIdentifier[len(ethernetSegment.EthernetSegmentIdentifier)-1]
    }
    if childYangName == "primary-service" {
        for _, c := range ethernetSegment.PrimaryService {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService{}
        ethernetSegment.PrimaryService = append(ethernetSegment.PrimaryService, child)
        return &ethernetSegment.PrimaryService[len(ethernetSegment.PrimaryService)-1]
    }
    if childYangName == "secondary-service" {
        for _, c := range ethernetSegment.SecondaryService {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService{}
        ethernetSegment.SecondaryService = append(ethernetSegment.SecondaryService, child)
        return &ethernetSegment.SecondaryService[len(ethernetSegment.SecondaryService)-1]
    }
    if childYangName == "service-carving-i-sidelected-result" {
        for _, c := range ethernetSegment.ServiceCarvingISidelectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult{}
        ethernetSegment.ServiceCarvingISidelectedResult = append(ethernetSegment.ServiceCarvingISidelectedResult, child)
        return &ethernetSegment.ServiceCarvingISidelectedResult[len(ethernetSegment.ServiceCarvingISidelectedResult)-1]
    }
    if childYangName == "service-carving-isid-not-elected-result" {
        for _, c := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult{}
        ethernetSegment.ServiceCarvingIsidNotElectedResult = append(ethernetSegment.ServiceCarvingIsidNotElectedResult, child)
        return &ethernetSegment.ServiceCarvingIsidNotElectedResult[len(ethernetSegment.ServiceCarvingIsidNotElectedResult)-1]
    }
    if childYangName == "service-carving-evi-elected-result" {
        for _, c := range ethernetSegment.ServiceCarvingEviElectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult{}
        ethernetSegment.ServiceCarvingEviElectedResult = append(ethernetSegment.ServiceCarvingEviElectedResult, child)
        return &ethernetSegment.ServiceCarvingEviElectedResult[len(ethernetSegment.ServiceCarvingEviElectedResult)-1]
    }
    if childYangName == "service-carving-evi-not-elected-result" {
        for _, c := range ethernetSegment.ServiceCarvingEviNotElectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult{}
        ethernetSegment.ServiceCarvingEviNotElectedResult = append(ethernetSegment.ServiceCarvingEviNotElectedResult, child)
        return &ethernetSegment.ServiceCarvingEviNotElectedResult[len(ethernetSegment.ServiceCarvingEviNotElectedResult)-1]
    }
    if childYangName == "next-hop" {
        for _, c := range ethernetSegment.NextHop {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_NextHop{}
        ethernetSegment.NextHop = append(ethernetSegment.NextHop, child)
        return &ethernetSegment.NextHop[len(ethernetSegment.NextHop)-1]
    }
    if childYangName == "service-carving-vpws-permanent-result" {
        for _, c := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult{}
        ethernetSegment.ServiceCarvingVpwsPermanentResult = append(ethernetSegment.ServiceCarvingVpwsPermanentResult, child)
        return &ethernetSegment.ServiceCarvingVpwsPermanentResult[len(ethernetSegment.ServiceCarvingVpwsPermanentResult)-1]
    }
    if childYangName == "remote-split-horizon-group-label" {
        for _, c := range ethernetSegment.RemoteSplitHorizonGroupLabel {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel{}
        ethernetSegment.RemoteSplitHorizonGroupLabel = append(ethernetSegment.RemoteSplitHorizonGroupLabel, child)
        return &ethernetSegment.RemoteSplitHorizonGroupLabel[len(ethernetSegment.RemoteSplitHorizonGroupLabel)-1]
    }
    return nil
}

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetSegment.EthernetSegmentIdentifier {
        children[ethernetSegment.EthernetSegmentIdentifier[i].GetSegmentPath()] = &ethernetSegment.EthernetSegmentIdentifier[i]
    }
    for i := range ethernetSegment.PrimaryService {
        children[ethernetSegment.PrimaryService[i].GetSegmentPath()] = &ethernetSegment.PrimaryService[i]
    }
    for i := range ethernetSegment.SecondaryService {
        children[ethernetSegment.SecondaryService[i].GetSegmentPath()] = &ethernetSegment.SecondaryService[i]
    }
    for i := range ethernetSegment.ServiceCarvingISidelectedResult {
        children[ethernetSegment.ServiceCarvingISidelectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingISidelectedResult[i]
    }
    for i := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
        children[ethernetSegment.ServiceCarvingIsidNotElectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingIsidNotElectedResult[i]
    }
    for i := range ethernetSegment.ServiceCarvingEviElectedResult {
        children[ethernetSegment.ServiceCarvingEviElectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingEviElectedResult[i]
    }
    for i := range ethernetSegment.ServiceCarvingEviNotElectedResult {
        children[ethernetSegment.ServiceCarvingEviNotElectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingEviNotElectedResult[i]
    }
    for i := range ethernetSegment.NextHop {
        children[ethernetSegment.NextHop[i].GetSegmentPath()] = &ethernetSegment.NextHop[i]
    }
    for i := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
        children[ethernetSegment.ServiceCarvingVpwsPermanentResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingVpwsPermanentResult[i]
    }
    for i := range ethernetSegment.RemoteSplitHorizonGroupLabel {
        children[ethernetSegment.RemoteSplitHorizonGroupLabel[i].GetSegmentPath()] = &ethernetSegment.RemoteSplitHorizonGroupLabel[i]
    }
    return children
}

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = ethernetSegment.InterfaceName
    leafs["esi1"] = ethernetSegment.Esi1
    leafs["esi2"] = ethernetSegment.Esi2
    leafs["esi3"] = ethernetSegment.Esi3
    leafs["esi4"] = ethernetSegment.Esi4
    leafs["esi5"] = ethernetSegment.Esi5
    leafs["esi-type"] = ethernetSegment.EsiType
    leafs["ethernet-segment-name"] = ethernetSegment.EthernetSegmentName
    leafs["ethernet-segment-state"] = ethernetSegment.EthernetSegmentState
    leafs["if-handle"] = ethernetSegment.IfHandle
    leafs["main-port-role"] = ethernetSegment.MainPortRole
    leafs["main-port-mac"] = ethernetSegment.MainPortMac
    leafs["num-up-p-ws"] = ethernetSegment.NumUpPWs
    leafs["route-target"] = ethernetSegment.RouteTarget
    leafs["rt-origin"] = ethernetSegment.RtOrigin
    leafs["es-bgp-gates"] = ethernetSegment.EsBgpGates
    leafs["es-l2fib-gates"] = ethernetSegment.EsL2FibGates
    leafs["mac-flushing-mode-config"] = ethernetSegment.MacFlushingModeConfig
    leafs["load-balance-mode-config"] = ethernetSegment.LoadBalanceModeConfig
    leafs["load-balance-mode-is-default"] = ethernetSegment.LoadBalanceModeIsDefault
    leafs["load-balance-mode-oper"] = ethernetSegment.LoadBalanceModeOper
    leafs["force-single-home"] = ethernetSegment.ForceSingleHome
    leafs["source-mac-oper"] = ethernetSegment.SourceMacOper
    leafs["source-mac-origin"] = ethernetSegment.SourceMacOrigin
    leafs["peering-timer"] = ethernetSegment.PeeringTimer
    leafs["peering-timer-left"] = ethernetSegment.PeeringTimerLeft
    leafs["recovery-timer"] = ethernetSegment.RecoveryTimer
    leafs["recovery-timer-left"] = ethernetSegment.RecoveryTimerLeft
    leafs["service-carving-mode"] = ethernetSegment.ServiceCarvingMode
    leafs["primary-services-input"] = ethernetSegment.PrimaryServicesInput
    leafs["secondary-services-input"] = ethernetSegment.SecondaryServicesInput
    leafs["forwarder-ports"] = ethernetSegment.ForwarderPorts
    leafs["permanent-forwarder-ports"] = ethernetSegment.PermanentForwarderPorts
    leafs["elected-forwarder-ports"] = ethernetSegment.ElectedForwarderPorts
    leafs["not-elected-forwarder-ports"] = ethernetSegment.NotElectedForwarderPorts
    leafs["not-config-forwarder-ports"] = ethernetSegment.NotConfigForwarderPorts
    leafs["mp-protected"] = ethernetSegment.MpProtected
    leafs["nve-anycast-vtep"] = ethernetSegment.NveAnycastVtep
    leafs["nve-ingress-replication"] = ethernetSegment.NveIngressReplication
    leafs["local-split-horizon-group-label"] = ethernetSegment.LocalSplitHorizonGroupLabel
    return leafs
}

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetYangName() string { return "ethernet-segment" }

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) SetParent(parent types.Entity) { ethernetSegment.parent = parent }

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetParent() types.Entity { return ethernetSegment.parent }

func (ethernetSegment *Evpn_Active_EthernetSegments_EthernetSegment) GetParentYangName() string { return "ethernet-segments" }

// Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return ethernetSegmentIdentifier.YFilter }

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { ethernetSegmentIdentifier.YFilter = yf }

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetSegmentPath() string {
    return "ethernet-segment-identifier"
}

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = ethernetSegmentIdentifier.Entry
    return leafs
}

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetYangName() string { return "ethernet-segment-identifier" }

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) SetParent(parent types.Entity) { ethernetSegmentIdentifier.parent = parent }

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetParent() types.Entity { return ethernetSegmentIdentifier.parent }

func (ethernetSegmentIdentifier *Evpn_Active_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService
// List of Primary services ESI/I-SIDs
type Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetFilter() yfilter.YFilter { return primaryService.YFilter }

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) SetFilter(yf yfilter.YFilter) { primaryService.YFilter = yf }

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetSegmentPath() string {
    return "primary-service"
}

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = primaryService.Entry
    return leafs
}

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetBundleName() string { return "cisco_ios_xr" }

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetYangName() string { return "primary-service" }

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) SetParent(parent types.Entity) { primaryService.parent = parent }

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetParent() types.Entity { return primaryService.parent }

func (primaryService *Evpn_Active_EthernetSegments_EthernetSegment_PrimaryService) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService
// List of Secondary services ESI/I-SIDs
type Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetFilter() yfilter.YFilter { return secondaryService.YFilter }

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) SetFilter(yf yfilter.YFilter) { secondaryService.YFilter = yf }

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetSegmentPath() string {
    return "secondary-service"
}

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = secondaryService.Entry
    return leafs
}

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetYangName() string { return "secondary-service" }

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) SetParent(parent types.Entity) { secondaryService.parent = parent }

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetParent() types.Entity { return secondaryService.parent }

func (secondaryService *Evpn_Active_EthernetSegments_EthernetSegment_SecondaryService) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult
// Elected ISID service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetFilter() yfilter.YFilter { return serviceCarvingISidelectedResult.YFilter }

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingISidelectedResult.YFilter = yf }

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetSegmentPath() string {
    return "service-carving-i-sidelected-result"
}

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingISidelectedResult.Entry
    return leafs
}

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetYangName() string { return "service-carving-i-sidelected-result" }

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) SetParent(parent types.Entity) { serviceCarvingISidelectedResult.parent = parent }

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetParent() types.Entity { return serviceCarvingISidelectedResult.parent }

func (serviceCarvingISidelectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult
// Not elected ISID service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetFilter() yfilter.YFilter { return serviceCarvingIsidNotElectedResult.YFilter }

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingIsidNotElectedResult.YFilter = yf }

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetSegmentPath() string {
    return "service-carving-isid-not-elected-result"
}

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingIsidNotElectedResult.Entry
    return leafs
}

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetYangName() string { return "service-carving-isid-not-elected-result" }

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) SetParent(parent types.Entity) { serviceCarvingIsidNotElectedResult.parent = parent }

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetParent() types.Entity { return serviceCarvingIsidNotElectedResult.parent }

func (serviceCarvingIsidNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult
// Elected EVI service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetFilter() yfilter.YFilter { return serviceCarvingEviElectedResult.YFilter }

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingEviElectedResult.YFilter = yf }

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetSegmentPath() string {
    return "service-carving-evi-elected-result"
}

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingEviElectedResult.Entry
    return leafs
}

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetYangName() string { return "service-carving-evi-elected-result" }

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) SetParent(parent types.Entity) { serviceCarvingEviElectedResult.parent = parent }

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetParent() types.Entity { return serviceCarvingEviElectedResult.parent }

func (serviceCarvingEviElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult
// Not elected EVI service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetFilter() yfilter.YFilter { return serviceCarvingEviNotElectedResult.YFilter }

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingEviNotElectedResult.YFilter = yf }

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetSegmentPath() string {
    return "service-carving-evi-not-elected-result"
}

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingEviNotElectedResult.Entry
    return leafs
}

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetYangName() string { return "service-carving-evi-not-elected-result" }

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) SetParent(parent types.Entity) { serviceCarvingEviNotElectedResult.parent = parent }

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetParent() types.Entity { return serviceCarvingEviNotElectedResult.parent }

func (serviceCarvingEviNotElectedResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_EthernetSegments_EthernetSegment_NextHop
// List of nexthop IPv6 addresses
type Evpn_Active_EthernetSegments_EthernetSegment_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}
}

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHop.NextHop
    return leafs
}

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *Evpn_Active_EthernetSegments_EthernetSegment_NextHop) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult
// Permanent EVPN VPWS service carving results
type Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPN ID. The type is interface{} with range: 0..4294967295.
    VpnId interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}
}

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetFilter() yfilter.YFilter { return serviceCarvingVpwsPermanentResult.YFilter }

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) SetFilter(yf yfilter.YFilter) { serviceCarvingVpwsPermanentResult.YFilter = yf }

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetGoName(yname string) string {
    if yname == "vpn-id" { return "VpnId" }
    if yname == "type" { return "Type" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    return ""
}

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetSegmentPath() string {
    return "service-carving-vpws-permanent-result"
}

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vpn-id"] = serviceCarvingVpwsPermanentResult.VpnId
    leafs["type"] = serviceCarvingVpwsPermanentResult.Type
    leafs["ethernet-tag"] = serviceCarvingVpwsPermanentResult.EthernetTag
    return leafs
}

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetYangName() string { return "service-carving-vpws-permanent-result" }

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) SetParent(parent types.Entity) { serviceCarvingVpwsPermanentResult.parent = parent }

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetParent() types.Entity { return serviceCarvingVpwsPermanentResult.parent }

func (serviceCarvingVpwsPermanentResult *Evpn_Active_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Split horizon label associated with next-hop address. The type is
    // interface{} with range: 0..4294967295.
    Label interface{}
}

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetFilter() yfilter.YFilter { return remoteSplitHorizonGroupLabel.YFilter }

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) SetFilter(yf yfilter.YFilter) { remoteSplitHorizonGroupLabel.YFilter = yf }

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "label" { return "Label" }
    return ""
}

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetSegmentPath() string {
    return "remote-split-horizon-group-label"
}

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = remoteSplitHorizonGroupLabel.NextHop
    leafs["label"] = remoteSplitHorizonGroupLabel.Label
    return leafs
}

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetBundleName() string { return "cisco_ios_xr" }

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetYangName() string { return "remote-split-horizon-group-label" }

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) SetParent(parent types.Entity) { remoteSplitHorizonGroupLabel.parent = parent }

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetParent() types.Entity { return remoteSplitHorizonGroupLabel.parent }

func (remoteSplitHorizonGroupLabel *Evpn_Active_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Active_AcIds
// EVPN AC ID table
type Evpn_Active_AcIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN AC ID table. The type is slice of Evpn_Active_AcIds_AcId.
    AcId []Evpn_Active_AcIds_AcId
}

func (acIds *Evpn_Active_AcIds) GetFilter() yfilter.YFilter { return acIds.YFilter }

func (acIds *Evpn_Active_AcIds) SetFilter(yf yfilter.YFilter) { acIds.YFilter = yf }

func (acIds *Evpn_Active_AcIds) GetGoName(yname string) string {
    if yname == "ac-id" { return "AcId" }
    return ""
}

func (acIds *Evpn_Active_AcIds) GetSegmentPath() string {
    return "ac-ids"
}

func (acIds *Evpn_Active_AcIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ac-id" {
        for _, c := range acIds.AcId {
            if acIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Active_AcIds_AcId{}
        acIds.AcId = append(acIds.AcId, child)
        return &acIds.AcId[len(acIds.AcId)-1]
    }
    return nil
}

func (acIds *Evpn_Active_AcIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range acIds.AcId {
        children[acIds.AcId[i].GetSegmentPath()] = &acIds.AcId[i]
    }
    return children
}

func (acIds *Evpn_Active_AcIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (acIds *Evpn_Active_AcIds) GetBundleName() string { return "cisco_ios_xr" }

func (acIds *Evpn_Active_AcIds) GetYangName() string { return "ac-ids" }

func (acIds *Evpn_Active_AcIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acIds *Evpn_Active_AcIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acIds *Evpn_Active_AcIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acIds *Evpn_Active_AcIds) SetParent(parent types.Entity) { acIds.parent = parent }

func (acIds *Evpn_Active_AcIds) GetParent() types.Entity { return acIds.parent }

func (acIds *Evpn_Active_AcIds) GetParentYangName() string { return "active" }

// Evpn_Active_AcIds_AcId
// EVPN AC ID table
type Evpn_Active_AcIds_AcId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // AC ID. The type is interface{} with range: -2147483648..2147483647.
    AcId interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}
}

func (acId *Evpn_Active_AcIds_AcId) GetFilter() yfilter.YFilter { return acId.YFilter }

func (acId *Evpn_Active_AcIds_AcId) SetFilter(yf yfilter.YFilter) { acId.YFilter = yf }

func (acId *Evpn_Active_AcIds_AcId) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "ac-id" { return "AcId" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (acId *Evpn_Active_AcIds_AcId) GetSegmentPath() string {
    return "ac-id"
}

func (acId *Evpn_Active_AcIds_AcId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acId *Evpn_Active_AcIds_AcId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acId *Evpn_Active_AcIds_AcId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = acId.Evi
    leafs["ac-id"] = acId.AcId
    leafs["evi-xr"] = acId.EviXr
    leafs["neighbor"] = acId.Neighbor
    return leafs
}

func (acId *Evpn_Active_AcIds_AcId) GetBundleName() string { return "cisco_ios_xr" }

func (acId *Evpn_Active_AcIds_AcId) GetYangName() string { return "ac-id" }

func (acId *Evpn_Active_AcIds_AcId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acId *Evpn_Active_AcIds_AcId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acId *Evpn_Active_AcIds_AcId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acId *Evpn_Active_AcIds_AcId) SetParent(parent types.Entity) { acId.parent = parent }

func (acId *Evpn_Active_AcIds_AcId) GetParent() types.Entity { return acId.parent }

func (acId *Evpn_Active_AcIds_AcId) GetParentYangName() string { return "ac-ids" }

// Evpn_Standby
// Standby EVPN operational data
type Evpn_Standby struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Table.
    Evis Evpn_Standby_Evis

    // L2VPN EVPN Summary.
    Summary Evpn_Standby_Summary

    // L2VPN EVI Detail Table.
    EviDetail Evpn_Standby_EviDetail

    // EVPN Ethernet-Segment Table.
    EthernetSegments Evpn_Standby_EthernetSegments

    // EVPN AC ID table.
    AcIds Evpn_Standby_AcIds
}

func (standby *Evpn_Standby) GetFilter() yfilter.YFilter { return standby.YFilter }

func (standby *Evpn_Standby) SetFilter(yf yfilter.YFilter) { standby.YFilter = yf }

func (standby *Evpn_Standby) GetGoName(yname string) string {
    if yname == "evis" { return "Evis" }
    if yname == "summary" { return "Summary" }
    if yname == "evi-detail" { return "EviDetail" }
    if yname == "ethernet-segments" { return "EthernetSegments" }
    if yname == "ac-ids" { return "AcIds" }
    return ""
}

func (standby *Evpn_Standby) GetSegmentPath() string {
    return "standby"
}

func (standby *Evpn_Standby) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "evis" {
        return &standby.Evis
    }
    if childYangName == "summary" {
        return &standby.Summary
    }
    if childYangName == "evi-detail" {
        return &standby.EviDetail
    }
    if childYangName == "ethernet-segments" {
        return &standby.EthernetSegments
    }
    if childYangName == "ac-ids" {
        return &standby.AcIds
    }
    return nil
}

func (standby *Evpn_Standby) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["evis"] = &standby.Evis
    children["summary"] = &standby.Summary
    children["evi-detail"] = &standby.EviDetail
    children["ethernet-segments"] = &standby.EthernetSegments
    children["ac-ids"] = &standby.AcIds
    return children
}

func (standby *Evpn_Standby) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (standby *Evpn_Standby) GetBundleName() string { return "cisco_ios_xr" }

func (standby *Evpn_Standby) GetYangName() string { return "standby" }

func (standby *Evpn_Standby) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standby *Evpn_Standby) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standby *Evpn_Standby) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standby *Evpn_Standby) SetParent(parent types.Entity) { standby.parent = parent }

func (standby *Evpn_Standby) GetParent() types.Entity { return standby.parent }

func (standby *Evpn_Standby) GetParentYangName() string { return "evpn" }

// Evpn_Standby_Evis
// L2VPN EVPN EVI Table
type Evpn_Standby_Evis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI Entry. The type is slice of Evpn_Standby_Evis_Evi.
    Evi []Evpn_Standby_Evis_Evi
}

func (evis *Evpn_Standby_Evis) GetFilter() yfilter.YFilter { return evis.YFilter }

func (evis *Evpn_Standby_Evis) SetFilter(yf yfilter.YFilter) { evis.YFilter = yf }

func (evis *Evpn_Standby_Evis) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    return ""
}

func (evis *Evpn_Standby_Evis) GetSegmentPath() string {
    return "evis"
}

func (evis *Evpn_Standby_Evis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "evi" {
        for _, c := range evis.Evi {
            if evis.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_Evis_Evi{}
        evis.Evi = append(evis.Evi, child)
        return &evis.Evi[len(evis.Evi)-1]
    }
    return nil
}

func (evis *Evpn_Standby_Evis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range evis.Evi {
        children[evis.Evi[i].GetSegmentPath()] = &evis.Evi[i]
    }
    return children
}

func (evis *Evpn_Standby_Evis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (evis *Evpn_Standby_Evis) GetBundleName() string { return "cisco_ios_xr" }

func (evis *Evpn_Standby_Evis) GetYangName() string { return "evis" }

func (evis *Evpn_Standby_Evis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evis *Evpn_Standby_Evis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evis *Evpn_Standby_Evis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evis *Evpn_Standby_Evis) SetParent(parent types.Entity) { evis.parent = parent }

func (evis *Evpn_Standby_Evis) GetParent() types.Entity { return evis.parent }

func (evis *Evpn_Standby_Evis) GetParentYangName() string { return "standby" }

// Evpn_Standby_Evis_Evi
// L2VPN EVPN EVI Entry
type Evpn_Standby_Evis_Evi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. EVPN id. The type is interface{} with range:
    // -2147483648..2147483647.
    Evi interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Bridge domain name. The type is string.
    BdName interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}
}

func (evi *Evpn_Standby_Evis_Evi) GetFilter() yfilter.YFilter { return evi.YFilter }

func (evi *Evpn_Standby_Evis_Evi) SetFilter(yf yfilter.YFilter) { evi.YFilter = yf }

func (evi *Evpn_Standby_Evis_Evi) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "bd-name" { return "BdName" }
    if yname == "type" { return "Type" }
    return ""
}

func (evi *Evpn_Standby_Evis_Evi) GetSegmentPath() string {
    return "evi" + "[evi='" + fmt.Sprintf("%v", evi.Evi) + "']"
}

func (evi *Evpn_Standby_Evis_Evi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (evi *Evpn_Standby_Evis_Evi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (evi *Evpn_Standby_Evis_Evi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = evi.Evi
    leafs["evi-xr"] = evi.EviXr
    leafs["bd-name"] = evi.BdName
    leafs["type"] = evi.Type
    return leafs
}

func (evi *Evpn_Standby_Evis_Evi) GetBundleName() string { return "cisco_ios_xr" }

func (evi *Evpn_Standby_Evis_Evi) GetYangName() string { return "evi" }

func (evi *Evpn_Standby_Evis_Evi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evi *Evpn_Standby_Evis_Evi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evi *Evpn_Standby_Evis_Evi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evi *Evpn_Standby_Evis_Evi) SetParent(parent types.Entity) { evi.parent = parent }

func (evi *Evpn_Standby_Evis_Evi) GetParent() types.Entity { return evi.parent }

func (evi *Evpn_Standby_Evis_Evi) GetParentYangName() string { return "evis" }

// Evpn_Standby_Summary
// L2VPN EVPN Summary
type Evpn_Standby_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Router ID. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    GlobalSourceMac interface{}

    // EVPN ES Peering Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    PeeringTime interface{}

    // EVPN ES Recovery Time (seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    RecoveryTime interface{}

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

func (summary *Evpn_Standby_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Evpn_Standby_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Evpn_Standby_Summary) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "as" { return "As" }
    if yname == "ev-is" { return "EvIs" }
    if yname == "local-mac-routes" { return "LocalMacRoutes" }
    if yname == "local-ipv4-mac-routes" { return "LocalIpv4MacRoutes" }
    if yname == "local-ipv6-mac-routes" { return "LocalIpv6MacRoutes" }
    if yname == "es-global-mac-routes" { return "EsGlobalMacRoutes" }
    if yname == "remote-mac-routes" { return "RemoteMacRoutes" }
    if yname == "remote-soo-mac-routes" { return "RemoteSooMacRoutes" }
    if yname == "remote-ipv4-mac-routes" { return "RemoteIpv4MacRoutes" }
    if yname == "remote-ipv6-mac-routes" { return "RemoteIpv6MacRoutes" }
    if yname == "local-imcast-routes" { return "LocalImcastRoutes" }
    if yname == "remote-imcast-routes" { return "RemoteImcastRoutes" }
    if yname == "labels" { return "Labels" }
    if yname == "es-entries" { return "EsEntries" }
    if yname == "neighbor-entries" { return "NeighborEntries" }
    if yname == "local-ead-routes" { return "LocalEadRoutes" }
    if yname == "remote-ead-routes" { return "RemoteEadRoutes" }
    if yname == "global-source-mac" { return "GlobalSourceMac" }
    if yname == "peering-time" { return "PeeringTime" }
    if yname == "recovery-time" { return "RecoveryTime" }
    if yname == "mac-secure-move-count" { return "MacSecureMoveCount" }
    if yname == "mac-secure-move-interval" { return "MacSecureMoveInterval" }
    if yname == "mac-secure-freeze-time" { return "MacSecureFreezeTime" }
    if yname == "mac-secure-retry-count" { return "MacSecureRetryCount" }
    if yname == "cost-out" { return "CostOut" }
    if yname == "startup-cost-in-time" { return "StartupCostInTime" }
    if yname == "l2rib-throttle" { return "L2RibThrottle" }
    if yname == "logging-df-election-enabled" { return "LoggingDfElectionEnabled" }
    return ""
}

func (summary *Evpn_Standby_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Evpn_Standby_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *Evpn_Standby_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *Evpn_Standby_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = summary.RouterId
    leafs["as"] = summary.As
    leafs["ev-is"] = summary.EvIs
    leafs["local-mac-routes"] = summary.LocalMacRoutes
    leafs["local-ipv4-mac-routes"] = summary.LocalIpv4MacRoutes
    leafs["local-ipv6-mac-routes"] = summary.LocalIpv6MacRoutes
    leafs["es-global-mac-routes"] = summary.EsGlobalMacRoutes
    leafs["remote-mac-routes"] = summary.RemoteMacRoutes
    leafs["remote-soo-mac-routes"] = summary.RemoteSooMacRoutes
    leafs["remote-ipv4-mac-routes"] = summary.RemoteIpv4MacRoutes
    leafs["remote-ipv6-mac-routes"] = summary.RemoteIpv6MacRoutes
    leafs["local-imcast-routes"] = summary.LocalImcastRoutes
    leafs["remote-imcast-routes"] = summary.RemoteImcastRoutes
    leafs["labels"] = summary.Labels
    leafs["es-entries"] = summary.EsEntries
    leafs["neighbor-entries"] = summary.NeighborEntries
    leafs["local-ead-routes"] = summary.LocalEadRoutes
    leafs["remote-ead-routes"] = summary.RemoteEadRoutes
    leafs["global-source-mac"] = summary.GlobalSourceMac
    leafs["peering-time"] = summary.PeeringTime
    leafs["recovery-time"] = summary.RecoveryTime
    leafs["mac-secure-move-count"] = summary.MacSecureMoveCount
    leafs["mac-secure-move-interval"] = summary.MacSecureMoveInterval
    leafs["mac-secure-freeze-time"] = summary.MacSecureFreezeTime
    leafs["mac-secure-retry-count"] = summary.MacSecureRetryCount
    leafs["cost-out"] = summary.CostOut
    leafs["startup-cost-in-time"] = summary.StartupCostInTime
    leafs["l2rib-throttle"] = summary.L2RibThrottle
    leafs["logging-df-election-enabled"] = summary.LoggingDfElectionEnabled
    return leafs
}

func (summary *Evpn_Standby_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Evpn_Standby_Summary) GetYangName() string { return "summary" }

func (summary *Evpn_Standby_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Evpn_Standby_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Evpn_Standby_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Evpn_Standby_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Evpn_Standby_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Evpn_Standby_Summary) GetParentYangName() string { return "standby" }

// Evpn_Standby_EviDetail
// L2VPN EVI Detail Table
type Evpn_Standby_EviDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info Elements.
    Elements Evpn_Standby_EviDetail_Elements

    // Container for all EVI detail info.
    EviChildren Evpn_Standby_EviDetail_EviChildren
}

func (eviDetail *Evpn_Standby_EviDetail) GetFilter() yfilter.YFilter { return eviDetail.YFilter }

func (eviDetail *Evpn_Standby_EviDetail) SetFilter(yf yfilter.YFilter) { eviDetail.YFilter = yf }

func (eviDetail *Evpn_Standby_EviDetail) GetGoName(yname string) string {
    if yname == "elements" { return "Elements" }
    if yname == "evi-children" { return "EviChildren" }
    return ""
}

func (eviDetail *Evpn_Standby_EviDetail) GetSegmentPath() string {
    return "evi-detail"
}

func (eviDetail *Evpn_Standby_EviDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "elements" {
        return &eviDetail.Elements
    }
    if childYangName == "evi-children" {
        return &eviDetail.EviChildren
    }
    return nil
}

func (eviDetail *Evpn_Standby_EviDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["elements"] = &eviDetail.Elements
    children["evi-children"] = &eviDetail.EviChildren
    return children
}

func (eviDetail *Evpn_Standby_EviDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eviDetail *Evpn_Standby_EviDetail) GetBundleName() string { return "cisco_ios_xr" }

func (eviDetail *Evpn_Standby_EviDetail) GetYangName() string { return "evi-detail" }

func (eviDetail *Evpn_Standby_EviDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eviDetail *Evpn_Standby_EviDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eviDetail *Evpn_Standby_EviDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eviDetail *Evpn_Standby_EviDetail) SetParent(parent types.Entity) { eviDetail.parent = parent }

func (eviDetail *Evpn_Standby_EviDetail) GetParent() types.Entity { return eviDetail.parent }

func (eviDetail *Evpn_Standby_EviDetail) GetParentYangName() string { return "standby" }

// Evpn_Standby_EviDetail_Elements
// EVI BGP RT Detail Info Elements
type Evpn_Standby_EviDetail_Elements struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVI BGP RT Detail Info. The type is slice of
    // Evpn_Standby_EviDetail_Elements_Element.
    Element []Evpn_Standby_EviDetail_Elements_Element
}

func (elements *Evpn_Standby_EviDetail_Elements) GetFilter() yfilter.YFilter { return elements.YFilter }

func (elements *Evpn_Standby_EviDetail_Elements) SetFilter(yf yfilter.YFilter) { elements.YFilter = yf }

func (elements *Evpn_Standby_EviDetail_Elements) GetGoName(yname string) string {
    if yname == "element" { return "Element" }
    return ""
}

func (elements *Evpn_Standby_EviDetail_Elements) GetSegmentPath() string {
    return "elements"
}

func (elements *Evpn_Standby_EviDetail_Elements) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "element" {
        for _, c := range elements.Element {
            if elements.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_Elements_Element{}
        elements.Element = append(elements.Element, child)
        return &elements.Element[len(elements.Element)-1]
    }
    return nil
}

func (elements *Evpn_Standby_EviDetail_Elements) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range elements.Element {
        children[elements.Element[i].GetSegmentPath()] = &elements.Element[i]
    }
    return children
}

func (elements *Evpn_Standby_EviDetail_Elements) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (elements *Evpn_Standby_EviDetail_Elements) GetBundleName() string { return "cisco_ios_xr" }

func (elements *Evpn_Standby_EviDetail_Elements) GetYangName() string { return "elements" }

func (elements *Evpn_Standby_EviDetail_Elements) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (elements *Evpn_Standby_EviDetail_Elements) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (elements *Evpn_Standby_EviDetail_Elements) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (elements *Evpn_Standby_EviDetail_Elements) SetParent(parent types.Entity) { elements.parent = parent }

func (elements *Evpn_Standby_EviDetail_Elements) GetParent() types.Entity { return elements.parent }

func (elements *Evpn_Standby_EviDetail_Elements) GetParentYangName() string { return "evi-detail" }

// Evpn_Standby_EviDetail_Elements_Element
// EVI BGP RT Detail Info
type Evpn_Standby_EviDetail_Elements_Element struct {
    parent types.Entity
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
    Type interface{}

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

func (element *Evpn_Standby_EviDetail_Elements_Element) GetFilter() yfilter.YFilter { return element.YFilter }

func (element *Evpn_Standby_EviDetail_Elements_Element) SetFilter(yf yfilter.YFilter) { element.YFilter = yf }

func (element *Evpn_Standby_EviDetail_Elements_Element) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "description" { return "Description" }
    if yname == "bd-name" { return "BdName" }
    if yname == "type" { return "Type" }
    if yname == "unicast-label" { return "UnicastLabel" }
    if yname == "multicast-label" { return "MulticastLabel" }
    if yname == "cw-disable" { return "CwDisable" }
    if yname == "table-policy-name" { return "TablePolicyName" }
    if yname == "forward-class" { return "ForwardClass" }
    if yname == "rt-import-block-set" { return "RtImportBlockSet" }
    if yname == "rt-export-block-set" { return "RtExportBlockSet" }
    if yname == "advertise-mac" { return "AdvertiseMac" }
    if yname == "advertise-bvi-mac" { return "AdvertiseBviMac" }
    if yname == "aliasing-disabled" { return "AliasingDisabled" }
    if yname == "unknown-unicast-flooding-disabled" { return "UnknownUnicastFloodingDisabled" }
    if yname == "reoriginate-disabled" { return "ReoriginateDisabled" }
    if yname == "stitching" { return "Stitching" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "flow-label" { return "FlowLabel" }
    if yname == "rd-auto" { return "RdAuto" }
    if yname == "rd-configured" { return "RdConfigured" }
    if yname == "rt-auto" { return "RtAuto" }
    if yname == "rt-auto-stitching" { return "RtAutoStitching" }
    return ""
}

func (element *Evpn_Standby_EviDetail_Elements_Element) GetSegmentPath() string {
    return "element" + "[evi='" + fmt.Sprintf("%v", element.Evi) + "']"
}

func (element *Evpn_Standby_EviDetail_Elements_Element) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-label" {
        return &element.FlowLabel
    }
    if childYangName == "rd-auto" {
        return &element.RdAuto
    }
    if childYangName == "rd-configured" {
        return &element.RdConfigured
    }
    if childYangName == "rt-auto" {
        return &element.RtAuto
    }
    if childYangName == "rt-auto-stitching" {
        return &element.RtAutoStitching
    }
    return nil
}

func (element *Evpn_Standby_EviDetail_Elements_Element) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flow-label"] = &element.FlowLabel
    children["rd-auto"] = &element.RdAuto
    children["rd-configured"] = &element.RdConfigured
    children["rt-auto"] = &element.RtAuto
    children["rt-auto-stitching"] = &element.RtAutoStitching
    return children
}

func (element *Evpn_Standby_EviDetail_Elements_Element) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = element.Evi
    leafs["evi-xr"] = element.EviXr
    leafs["description"] = element.Description
    leafs["bd-name"] = element.BdName
    leafs["type"] = element.Type
    leafs["unicast-label"] = element.UnicastLabel
    leafs["multicast-label"] = element.MulticastLabel
    leafs["cw-disable"] = element.CwDisable
    leafs["table-policy-name"] = element.TablePolicyName
    leafs["forward-class"] = element.ForwardClass
    leafs["rt-import-block-set"] = element.RtImportBlockSet
    leafs["rt-export-block-set"] = element.RtExportBlockSet
    leafs["advertise-mac"] = element.AdvertiseMac
    leafs["advertise-bvi-mac"] = element.AdvertiseBviMac
    leafs["aliasing-disabled"] = element.AliasingDisabled
    leafs["unknown-unicast-flooding-disabled"] = element.UnknownUnicastFloodingDisabled
    leafs["reoriginate-disabled"] = element.ReoriginateDisabled
    leafs["stitching"] = element.Stitching
    leafs["encapsulation"] = element.Encapsulation
    return leafs
}

func (element *Evpn_Standby_EviDetail_Elements_Element) GetBundleName() string { return "cisco_ios_xr" }

func (element *Evpn_Standby_EviDetail_Elements_Element) GetYangName() string { return "element" }

func (element *Evpn_Standby_EviDetail_Elements_Element) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (element *Evpn_Standby_EviDetail_Elements_Element) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (element *Evpn_Standby_EviDetail_Elements_Element) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (element *Evpn_Standby_EviDetail_Elements_Element) SetParent(parent types.Entity) { element.parent = parent }

func (element *Evpn_Standby_EviDetail_Elements_Element) GetParent() types.Entity { return element.parent }

func (element *Evpn_Standby_EviDetail_Elements_Element) GetParentYangName() string { return "elements" }

// Evpn_Standby_EviDetail_Elements_Element_FlowLabel
// Flow Label Information
type Evpn_Standby_EviDetail_Elements_Element_FlowLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Static flow label. The type is bool.
    StaticFlowLabel interface{}

    // Globally configured flow label. The type is bool.
    GlobalFlowLabel interface{}
}

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetFilter() yfilter.YFilter { return flowLabel.YFilter }

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) SetFilter(yf yfilter.YFilter) { flowLabel.YFilter = yf }

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetGoName(yname string) string {
    if yname == "static-flow-label" { return "StaticFlowLabel" }
    if yname == "global-flow-label" { return "GlobalFlowLabel" }
    return ""
}

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetSegmentPath() string {
    return "flow-label"
}

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["static-flow-label"] = flowLabel.StaticFlowLabel
    leafs["global-flow-label"] = flowLabel.GlobalFlowLabel
    return leafs
}

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetBundleName() string { return "cisco_ios_xr" }

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetYangName() string { return "flow-label" }

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) SetParent(parent types.Entity) { flowLabel.parent = parent }

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetParent() types.Entity { return flowLabel.parent }

func (flowLabel *Evpn_Standby_EviDetail_Elements_Element_FlowLabel) GetParentYangName() string { return "element" }

// Evpn_Standby_EviDetail_Elements_Element_RdAuto
// Automatic Route Distingtuisher
type Evpn_Standby_EviDetail_Elements_Element_RdAuto struct {
    parent types.Entity
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

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetFilter() yfilter.YFilter { return rdAuto.YFilter }

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) SetFilter(yf yfilter.YFilter) { rdAuto.YFilter = yf }

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetGoName(yname string) string {
    if yname == "rd" { return "Rd" }
    if yname == "auto" { return "Auto" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    return ""
}

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetSegmentPath() string {
    return "rd-auto"
}

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto" {
        return &rdAuto.Auto
    }
    if childYangName == "two-byte-as" {
        return &rdAuto.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rdAuto.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rdAuto.V4Addr
    }
    return nil
}

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto"] = &rdAuto.Auto
    children["two-byte-as"] = &rdAuto.TwoByteAs
    children["four-byte-as"] = &rdAuto.FourByteAs
    children["v4-addr"] = &rdAuto.V4Addr
    return children
}

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rd"] = rdAuto.Rd
    return leafs
}

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetBundleName() string { return "cisco_ios_xr" }

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetYangName() string { return "rd-auto" }

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) SetParent(parent types.Entity) { rdAuto.parent = parent }

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetParent() types.Entity { return rdAuto.parent }

func (rdAuto *Evpn_Standby_EviDetail_Elements_Element_RdAuto) GetParentYangName() string { return "element" }

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto
// auto
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetFilter() yfilter.YFilter { return auto.YFilter }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) SetFilter(yf yfilter.YFilter) { auto.YFilter = yf }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "auto-index" { return "AutoIndex" }
    return ""
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetSegmentPath() string {
    return "auto"
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = auto.RouterId
    leafs["auto-index"] = auto.AutoIndex
    return leafs
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetBundleName() string { return "cisco_ios_xr" }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetYangName() string { return "auto" }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) SetParent(parent types.Entity) { auto.parent = parent }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetParent() types.Entity { return auto.parent }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdAuto_Auto) GetParentYangName() string { return "rd-auto" }

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_TwoByteAs) GetParentYangName() string { return "rd-auto" }

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdAuto_FourByteAs) GetParentYangName() string { return "rd-auto" }

// Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdAuto_V4Addr) GetParentYangName() string { return "rd-auto" }

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured
// Configured Route Distinguisher
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured struct {
    parent types.Entity
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

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetFilter() yfilter.YFilter { return rdConfigured.YFilter }

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) SetFilter(yf yfilter.YFilter) { rdConfigured.YFilter = yf }

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetGoName(yname string) string {
    if yname == "rd" { return "Rd" }
    if yname == "auto" { return "Auto" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    return ""
}

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetSegmentPath() string {
    return "rd-configured"
}

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto" {
        return &rdConfigured.Auto
    }
    if childYangName == "two-byte-as" {
        return &rdConfigured.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rdConfigured.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rdConfigured.V4Addr
    }
    return nil
}

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto"] = &rdConfigured.Auto
    children["two-byte-as"] = &rdConfigured.TwoByteAs
    children["four-byte-as"] = &rdConfigured.FourByteAs
    children["v4-addr"] = &rdConfigured.V4Addr
    return children
}

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rd"] = rdConfigured.Rd
    return leafs
}

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetBundleName() string { return "cisco_ios_xr" }

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetYangName() string { return "rd-configured" }

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) SetParent(parent types.Entity) { rdConfigured.parent = parent }

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetParent() types.Entity { return rdConfigured.parent }

func (rdConfigured *Evpn_Standby_EviDetail_Elements_Element_RdConfigured) GetParentYangName() string { return "element" }

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto
// auto
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP Router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Auto-generated Index. The type is interface{} with range: 0..65535.
    AutoIndex interface{}
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetFilter() yfilter.YFilter { return auto.YFilter }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) SetFilter(yf yfilter.YFilter) { auto.YFilter = yf }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "auto-index" { return "AutoIndex" }
    return ""
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetSegmentPath() string {
    return "auto"
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = auto.RouterId
    leafs["auto-index"] = auto.AutoIndex
    return leafs
}

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetBundleName() string { return "cisco_ios_xr" }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetYangName() string { return "auto" }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) SetParent(parent types.Entity) { auto.parent = parent }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetParent() types.Entity { return auto.parent }

func (auto *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_Auto) GetParentYangName() string { return "rd-configured" }

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_TwoByteAs) GetParentYangName() string { return "rd-configured" }

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_FourByteAs) GetParentYangName() string { return "rd-configured" }

// Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RdConfigured_V4Addr) GetParentYangName() string { return "rd-configured" }

// Evpn_Standby_EviDetail_Elements_Element_RtAuto
// Automatic Route Target
type Evpn_Standby_EviDetail_Elements_Element_RtAuto struct {
    parent types.Entity
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

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetFilter() yfilter.YFilter { return rtAuto.YFilter }

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) SetFilter(yf yfilter.YFilter) { rtAuto.YFilter = yf }

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetGoName(yname string) string {
    if yname == "rt" { return "Rt" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    if yname == "es-import" { return "EsImport" }
    return ""
}

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetSegmentPath() string {
    return "rt-auto"
}

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "two-byte-as" {
        return &rtAuto.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rtAuto.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rtAuto.V4Addr
    }
    if childYangName == "es-import" {
        return &rtAuto.EsImport
    }
    return nil
}

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["two-byte-as"] = &rtAuto.TwoByteAs
    children["four-byte-as"] = &rtAuto.FourByteAs
    children["v4-addr"] = &rtAuto.V4Addr
    children["es-import"] = &rtAuto.EsImport
    return children
}

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt"] = rtAuto.Rt
    return leafs
}

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetBundleName() string { return "cisco_ios_xr" }

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetYangName() string { return "rt-auto" }

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) SetParent(parent types.Entity) { rtAuto.parent = parent }

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetParent() types.Entity { return rtAuto.parent }

func (rtAuto *Evpn_Standby_EviDetail_Elements_Element_RtAuto) GetParentYangName() string { return "element" }

// Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_TwoByteAs) GetParentYangName() string { return "rt-auto" }

// Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAuto_FourByteAs) GetParentYangName() string { return "rt-auto" }

// Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAuto_V4Addr) GetParentYangName() string { return "rt-auto" }

// Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport
// es import
type Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetFilter() yfilter.YFilter { return esImport.YFilter }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) SetFilter(yf yfilter.YFilter) { esImport.YFilter = yf }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetGoName(yname string) string {
    if yname == "high-bytes" { return "HighBytes" }
    if yname == "low-bytes" { return "LowBytes" }
    return ""
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetSegmentPath() string {
    return "es-import"
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-bytes"] = esImport.HighBytes
    leafs["low-bytes"] = esImport.LowBytes
    return leafs
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetBundleName() string { return "cisco_ios_xr" }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetYangName() string { return "es-import" }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) SetParent(parent types.Entity) { esImport.parent = parent }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetParent() types.Entity { return esImport.parent }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAuto_EsImport) GetParentYangName() string { return "rt-auto" }

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching
// Automatic Route Target Stitching
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching struct {
    parent types.Entity
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

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetFilter() yfilter.YFilter { return rtAutoStitching.YFilter }

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) SetFilter(yf yfilter.YFilter) { rtAutoStitching.YFilter = yf }

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetGoName(yname string) string {
    if yname == "rt" { return "Rt" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    if yname == "es-import" { return "EsImport" }
    return ""
}

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetSegmentPath() string {
    return "rt-auto-stitching"
}

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "two-byte-as" {
        return &rtAutoStitching.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &rtAutoStitching.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &rtAutoStitching.V4Addr
    }
    if childYangName == "es-import" {
        return &rtAutoStitching.EsImport
    }
    return nil
}

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["two-byte-as"] = &rtAutoStitching.TwoByteAs
    children["four-byte-as"] = &rtAutoStitching.FourByteAs
    children["v4-addr"] = &rtAutoStitching.V4Addr
    children["es-import"] = &rtAutoStitching.EsImport
    return children
}

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt"] = rtAutoStitching.Rt
    return leafs
}

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetBundleName() string { return "cisco_ios_xr" }

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetYangName() string { return "rt-auto-stitching" }

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) SetParent(parent types.Entity) { rtAutoStitching.parent = parent }

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetParent() types.Entity { return rtAutoStitching.parent }

func (rtAutoStitching *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching) GetParentYangName() string { return "element" }

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_TwoByteAs) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_FourByteAs) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_V4Addr) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport
// es import
type Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetFilter() yfilter.YFilter { return esImport.YFilter }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) SetFilter(yf yfilter.YFilter) { esImport.YFilter = yf }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetGoName(yname string) string {
    if yname == "high-bytes" { return "HighBytes" }
    if yname == "low-bytes" { return "LowBytes" }
    return ""
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetSegmentPath() string {
    return "es-import"
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-bytes"] = esImport.HighBytes
    leafs["low-bytes"] = esImport.LowBytes
    return leafs
}

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetBundleName() string { return "cisco_ios_xr" }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetYangName() string { return "es-import" }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) SetParent(parent types.Entity) { esImport.parent = parent }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetParent() types.Entity { return esImport.parent }

func (esImport *Evpn_Standby_EviDetail_Elements_Element_RtAutoStitching_EsImport) GetParentYangName() string { return "rt-auto-stitching" }

// Evpn_Standby_EviDetail_EviChildren
// Container for all EVI detail info
type Evpn_Standby_EviDetail_EviChildren struct {
    parent types.Entity
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

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetFilter() yfilter.YFilter { return eviChildren.YFilter }

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) SetFilter(yf yfilter.YFilter) { eviChildren.YFilter = yf }

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetGoName(yname string) string {
    if yname == "neighbors" { return "Neighbors" }
    if yname == "ethernet-auto-discoveries" { return "EthernetAutoDiscoveries" }
    if yname == "inclusive-multicasts" { return "InclusiveMulticasts" }
    if yname == "route-targets" { return "RouteTargets" }
    if yname == "macs" { return "Macs" }
    return ""
}

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetSegmentPath() string {
    return "evi-children"
}

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbors" {
        return &eviChildren.Neighbors
    }
    if childYangName == "ethernet-auto-discoveries" {
        return &eviChildren.EthernetAutoDiscoveries
    }
    if childYangName == "inclusive-multicasts" {
        return &eviChildren.InclusiveMulticasts
    }
    if childYangName == "route-targets" {
        return &eviChildren.RouteTargets
    }
    if childYangName == "macs" {
        return &eviChildren.Macs
    }
    return nil
}

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbors"] = &eviChildren.Neighbors
    children["ethernet-auto-discoveries"] = &eviChildren.EthernetAutoDiscoveries
    children["inclusive-multicasts"] = &eviChildren.InclusiveMulticasts
    children["route-targets"] = &eviChildren.RouteTargets
    children["macs"] = &eviChildren.Macs
    return children
}

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetBundleName() string { return "cisco_ios_xr" }

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetYangName() string { return "evi-children" }

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) SetParent(parent types.Entity) { eviChildren.parent = parent }

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetParent() types.Entity { return eviChildren.parent }

func (eviChildren *Evpn_Standby_EviDetail_EviChildren) GetParentYangName() string { return "evi-detail" }

// Evpn_Standby_EviDetail_EviChildren_Neighbors
// EVPN Neighbor table
type Evpn_Standby_EviDetail_EviChildren_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Neighbor table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor.
    Neighbor []Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor
}

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Evpn_Standby_EviDetail_EviChildren_Neighbors) GetParentYangName() string { return "evi-children" }

// Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor
// EVPN Neighbor table
type Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Neighbor IP. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}
}

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "neighbor-ip" { return "NeighborIp" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = neighbor.Evi
    leafs["neighbor-ip"] = neighbor.NeighborIp
    leafs["evi-xr"] = neighbor.EviXr
    leafs["neighbor"] = neighbor.Neighbor
    return leafs
}

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Evpn_Standby_EviDetail_EviChildren_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries
// EVPN Ethernet Auto-Discovery table
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Ethernet Auto-Discovery Entry. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery.
    EthernetAutoDiscovery []Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
}

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetFilter() yfilter.YFilter { return ethernetAutoDiscoveries.YFilter }

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) SetFilter(yf yfilter.YFilter) { ethernetAutoDiscoveries.YFilter = yf }

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetGoName(yname string) string {
    if yname == "ethernet-auto-discovery" { return "EthernetAutoDiscovery" }
    return ""
}

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetSegmentPath() string {
    return "ethernet-auto-discoveries"
}

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-auto-discovery" {
        for _, c := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
            if ethernetAutoDiscoveries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery{}
        ethernetAutoDiscoveries.EthernetAutoDiscovery = append(ethernetAutoDiscoveries.EthernetAutoDiscovery, child)
        return &ethernetAutoDiscoveries.EthernetAutoDiscovery[len(ethernetAutoDiscoveries.EthernetAutoDiscovery)-1]
    }
    return nil
}

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetAutoDiscoveries.EthernetAutoDiscovery {
        children[ethernetAutoDiscoveries.EthernetAutoDiscovery[i].GetSegmentPath()] = &ethernetAutoDiscoveries.EthernetAutoDiscovery[i]
    }
    return children
}

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetYangName() string { return "ethernet-auto-discoveries" }

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) SetParent(parent types.Entity) { ethernetAutoDiscoveries.parent = parent }

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetParent() types.Entity { return ethernetAutoDiscoveries.parent }

func (ethernetAutoDiscoveries *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries) GetParentYangName() string { return "evi-children" }

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery
// EVPN Ethernet Auto-Discovery Entry
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

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

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EthernetTag interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EthernetVpnid interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Local nexthop IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetFilter() yfilter.YFilter { return ethernetAutoDiscovery.YFilter }

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) SetFilter(yf yfilter.YFilter) { ethernetAutoDiscovery.YFilter = yf }

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "esi1" { return "Esi1" }
    if yname == "esi2" { return "Esi2" }
    if yname == "esi3" { return "Esi3" }
    if yname == "esi4" { return "Esi4" }
    if yname == "esi5" { return "Esi5" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    if yname == "ethernet-vpnid" { return "EthernetVpnid" }
    if yname == "type" { return "Type" }
    if yname == "ethernet-tag-xr" { return "EthernetTagXr" }
    if yname == "local-next-hop" { return "LocalNextHop" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "is-local-ead" { return "IsLocalEad" }
    if yname == "encap" { return "Encap" }
    if yname == "redundancy-single-active" { return "RedundancySingleActive" }
    if yname == "num-paths" { return "NumPaths" }
    if yname == "ethernet-segment-identifier" { return "EthernetSegmentIdentifier" }
    if yname == "path-buffer" { return "PathBuffer" }
    return ""
}

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetSegmentPath() string {
    return "ethernet-auto-discovery"
}

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-identifier" {
        for _, c := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
            if ethernetAutoDiscovery.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier{}
        ethernetAutoDiscovery.EthernetSegmentIdentifier = append(ethernetAutoDiscovery.EthernetSegmentIdentifier, child)
        return &ethernetAutoDiscovery.EthernetSegmentIdentifier[len(ethernetAutoDiscovery.EthernetSegmentIdentifier)-1]
    }
    if childYangName == "path-buffer" {
        for _, c := range ethernetAutoDiscovery.PathBuffer {
            if ethernetAutoDiscovery.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer{}
        ethernetAutoDiscovery.PathBuffer = append(ethernetAutoDiscovery.PathBuffer, child)
        return &ethernetAutoDiscovery.PathBuffer[len(ethernetAutoDiscovery.PathBuffer)-1]
    }
    return nil
}

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetAutoDiscovery.EthernetSegmentIdentifier {
        children[ethernetAutoDiscovery.EthernetSegmentIdentifier[i].GetSegmentPath()] = &ethernetAutoDiscovery.EthernetSegmentIdentifier[i]
    }
    for i := range ethernetAutoDiscovery.PathBuffer {
        children[ethernetAutoDiscovery.PathBuffer[i].GetSegmentPath()] = &ethernetAutoDiscovery.PathBuffer[i]
    }
    return children
}

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = ethernetAutoDiscovery.Evi
    leafs["esi1"] = ethernetAutoDiscovery.Esi1
    leafs["esi2"] = ethernetAutoDiscovery.Esi2
    leafs["esi3"] = ethernetAutoDiscovery.Esi3
    leafs["esi4"] = ethernetAutoDiscovery.Esi4
    leafs["esi5"] = ethernetAutoDiscovery.Esi5
    leafs["ethernet-tag"] = ethernetAutoDiscovery.EthernetTag
    leafs["ethernet-vpnid"] = ethernetAutoDiscovery.EthernetVpnid
    leafs["type"] = ethernetAutoDiscovery.Type
    leafs["ethernet-tag-xr"] = ethernetAutoDiscovery.EthernetTagXr
    leafs["local-next-hop"] = ethernetAutoDiscovery.LocalNextHop
    leafs["local-label"] = ethernetAutoDiscovery.LocalLabel
    leafs["is-local-ead"] = ethernetAutoDiscovery.IsLocalEad
    leafs["encap"] = ethernetAutoDiscovery.Encap
    leafs["redundancy-single-active"] = ethernetAutoDiscovery.RedundancySingleActive
    leafs["num-paths"] = ethernetAutoDiscovery.NumPaths
    return leafs
}

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetYangName() string { return "ethernet-auto-discovery" }

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) SetParent(parent types.Entity) { ethernetAutoDiscovery.parent = parent }

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetParent() types.Entity { return ethernetAutoDiscovery.parent }

func (ethernetAutoDiscovery *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery) GetParentYangName() string { return "ethernet-auto-discoveries" }

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return ethernetSegmentIdentifier.YFilter }

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { ethernetSegmentIdentifier.YFilter = yf }

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetSegmentPath() string {
    return "ethernet-segment-identifier"
}

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = ethernetSegmentIdentifier.Entry
    return leafs
}

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetYangName() string { return "ethernet-segment-identifier" }

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) SetParent(parent types.Entity) { ethernetSegmentIdentifier.parent = parent }

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetParent() types.Entity { return ethernetSegmentIdentifier.parent }

func (ethernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_EthernetSegmentIdentifier) GetParentYangName() string { return "ethernet-auto-discovery" }

// Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer
// Path List Buffer
type Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetFilter() yfilter.YFilter { return pathBuffer.YFilter }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) SetFilter(yf yfilter.YFilter) { pathBuffer.YFilter = yf }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "output-label" { return "OutputLabel" }
    if yname == "srte-tunnel" { return "SrteTunnel" }
    return ""
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetSegmentPath() string {
    return "path-buffer"
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = pathBuffer.NextHop
    leafs["output-label"] = pathBuffer.OutputLabel
    leafs["srte-tunnel"] = pathBuffer.SrteTunnel
    return leafs
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetYangName() string { return "path-buffer" }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) SetParent(parent types.Entity) { pathBuffer.parent = parent }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetParent() types.Entity { return pathBuffer.parent }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_EthernetAutoDiscoveries_EthernetAutoDiscovery_PathBuffer) GetParentYangName() string { return "ethernet-auto-discovery" }

// Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts
// L2VPN EVPN IMCAST table
type Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN IMCAST table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast.
    InclusiveMulticast []Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
}

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetFilter() yfilter.YFilter { return inclusiveMulticasts.YFilter }

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) SetFilter(yf yfilter.YFilter) { inclusiveMulticasts.YFilter = yf }

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetGoName(yname string) string {
    if yname == "inclusive-multicast" { return "InclusiveMulticast" }
    return ""
}

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetSegmentPath() string {
    return "inclusive-multicasts"
}

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inclusive-multicast" {
        for _, c := range inclusiveMulticasts.InclusiveMulticast {
            if inclusiveMulticasts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast{}
        inclusiveMulticasts.InclusiveMulticast = append(inclusiveMulticasts.InclusiveMulticast, child)
        return &inclusiveMulticasts.InclusiveMulticast[len(inclusiveMulticasts.InclusiveMulticast)-1]
    }
    return nil
}

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range inclusiveMulticasts.InclusiveMulticast {
        children[inclusiveMulticasts.InclusiveMulticast[i].GetSegmentPath()] = &inclusiveMulticasts.InclusiveMulticast[i]
    }
    return children
}

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetBundleName() string { return "cisco_ios_xr" }

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetYangName() string { return "inclusive-multicasts" }

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) SetParent(parent types.Entity) { inclusiveMulticasts.parent = parent }

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetParent() types.Entity { return inclusiveMulticasts.parent }

func (inclusiveMulticasts *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts) GetParentYangName() string { return "evi-children" }

// Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast
// L2VPN EVPN IMCAST table
type Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag. The type is interface{} with range: -2147483648..2147483647.
    EthernetTag interface{}

    // Originating IP. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIp interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTagXr interface{}

    // Originating IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    OriginatingIpXr interface{}

    // IP of nexthop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetFilter() yfilter.YFilter { return inclusiveMulticast.YFilter }

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) SetFilter(yf yfilter.YFilter) { inclusiveMulticast.YFilter = yf }

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    if yname == "originating-ip" { return "OriginatingIp" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "ethernet-tag-xr" { return "EthernetTagXr" }
    if yname == "originating-ip-xr" { return "OriginatingIpXr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "output-label" { return "OutputLabel" }
    if yname == "is-local-entry" { return "IsLocalEntry" }
    if yname == "is-proxy-entry" { return "IsProxyEntry" }
    if yname == "encap-type" { return "EncapType" }
    return ""
}

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetSegmentPath() string {
    return "inclusive-multicast"
}

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = inclusiveMulticast.Evi
    leafs["ethernet-tag"] = inclusiveMulticast.EthernetTag
    leafs["originating-ip"] = inclusiveMulticast.OriginatingIp
    leafs["evi-xr"] = inclusiveMulticast.EviXr
    leafs["ethernet-tag-xr"] = inclusiveMulticast.EthernetTagXr
    leafs["originating-ip-xr"] = inclusiveMulticast.OriginatingIpXr
    leafs["next-hop"] = inclusiveMulticast.NextHop
    leafs["output-label"] = inclusiveMulticast.OutputLabel
    leafs["is-local-entry"] = inclusiveMulticast.IsLocalEntry
    leafs["is-proxy-entry"] = inclusiveMulticast.IsProxyEntry
    leafs["encap-type"] = inclusiveMulticast.EncapType
    return leafs
}

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetBundleName() string { return "cisco_ios_xr" }

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetYangName() string { return "inclusive-multicast" }

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) SetParent(parent types.Entity) { inclusiveMulticast.parent = parent }

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetParent() types.Entity { return inclusiveMulticast.parent }

func (inclusiveMulticast *Evpn_Standby_EviDetail_EviChildren_InclusiveMulticasts_InclusiveMulticast) GetParentYangName() string { return "inclusive-multicasts" }

// Evpn_Standby_EviDetail_EviChildren_RouteTargets
// L2VPN EVPN EVI RT Child Table
type Evpn_Standby_EviDetail_EviChildren_RouteTargets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN EVI RT Table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget.
    RouteTarget []Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget
}

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetFilter() yfilter.YFilter { return routeTargets.YFilter }

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) SetFilter(yf yfilter.YFilter) { routeTargets.YFilter = yf }

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetGoName(yname string) string {
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetSegmentPath() string {
    return "route-targets"
}

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        for _, c := range routeTargets.RouteTarget {
            if routeTargets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget{}
        routeTargets.RouteTarget = append(routeTargets.RouteTarget, child)
        return &routeTargets.RouteTarget[len(routeTargets.RouteTarget)-1]
    }
    return nil
}

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routeTargets.RouteTarget {
        children[routeTargets.RouteTarget[i].GetSegmentPath()] = &routeTargets.RouteTarget[i]
    }
    return children
}

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetBundleName() string { return "cisco_ios_xr" }

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetYangName() string { return "route-targets" }

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) SetParent(parent types.Entity) { routeTargets.parent = parent }

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetParent() types.Entity { return routeTargets.parent }

func (routeTargets *Evpn_Standby_EviDetail_EviChildren_RouteTargets) GetParentYangName() string { return "evi-children" }

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget
// L2VPN EVPN EVI RT Table
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Role of the route target. The type is BgpRouteTargetRole.
    Role interface{}

    // Type of the route target. The type is BgpRouteTarget.
    Type interface{}

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

    // Bridge Domain Name. The type is string.
    BdName interface{}

    // VPN ID. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // RT Role. The type is L2vpnAdRtRole.
    RouteTargetRole interface{}

    // RT Stitching. The type is bool.
    RouteTargetStitching interface{}

    // Route Target.
    RouteTarget Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "role" { return "Role" }
    if yname == "type" { return "Type" }
    if yname == "format" { return "Format" }
    if yname == "as" { return "As" }
    if yname == "as-index" { return "AsIndex" }
    if yname == "addr-index" { return "AddrIndex" }
    if yname == "address" { return "Address" }
    if yname == "bd-name" { return "BdName" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "route-target-role" { return "RouteTargetRole" }
    if yname == "route-target-stitching" { return "RouteTargetStitching" }
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetSegmentPath() string {
    return "route-target"
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        return &routeTarget.RouteTarget
    }
    return nil
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route-target"] = &routeTarget.RouteTarget
    return children
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = routeTarget.Evi
    leafs["role"] = routeTarget.Role
    leafs["type"] = routeTarget.Type
    leafs["format"] = routeTarget.Format
    leafs["as"] = routeTarget.As
    leafs["as-index"] = routeTarget.AsIndex
    leafs["addr-index"] = routeTarget.AddrIndex
    leafs["address"] = routeTarget.Address
    leafs["bd-name"] = routeTarget.BdName
    leafs["evi-xr"] = routeTarget.EviXr
    leafs["route-target-role"] = routeTarget.RouteTargetRole
    leafs["route-target-stitching"] = routeTarget.RouteTargetStitching
    return leafs
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget) GetParentYangName() string { return "route-targets" }

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget
// Route Target
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget struct {
    parent types.Entity
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

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetGoName(yname string) string {
    if yname == "rt" { return "Rt" }
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "v4-addr" { return "V4Addr" }
    if yname == "es-import" { return "EsImport" }
    return ""
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetSegmentPath() string {
    return "route-target"
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "two-byte-as" {
        return &routeTarget.TwoByteAs
    }
    if childYangName == "four-byte-as" {
        return &routeTarget.FourByteAs
    }
    if childYangName == "v4-addr" {
        return &routeTarget.V4Addr
    }
    if childYangName == "es-import" {
        return &routeTarget.EsImport
    }
    return nil
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["two-byte-as"] = &routeTarget.TwoByteAs
    children["four-byte-as"] = &routeTarget.FourByteAs
    children["v4-addr"] = &routeTarget.V4Addr
    children["es-import"] = &routeTarget.EsImport
    return children
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt"] = routeTarget.Rt
    return leafs
}

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget) GetParentYangName() string { return "route-target" }

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs
// two byte as
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2 Byte AS Number. The type is interface{} with range: 0..65535.
    TwoByteAs interface{}

    // 4 Byte Index. The type is interface{} with range: 0..4294967295.
    FourByteIndex interface{}
}

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetFilter() yfilter.YFilter { return twoByteAs.YFilter }

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) SetFilter(yf yfilter.YFilter) { twoByteAs.YFilter = yf }

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetGoName(yname string) string {
    if yname == "two-byte-as" { return "TwoByteAs" }
    if yname == "four-byte-index" { return "FourByteIndex" }
    return ""
}

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetSegmentPath() string {
    return "two-byte-as"
}

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["two-byte-as"] = twoByteAs.TwoByteAs
    leafs["four-byte-index"] = twoByteAs.FourByteIndex
    return leafs
}

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetYangName() string { return "two-byte-as" }

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) SetParent(parent types.Entity) { twoByteAs.parent = parent }

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetParent() types.Entity { return twoByteAs.parent }

func (twoByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_TwoByteAs) GetParentYangName() string { return "route-target" }

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs
// four byte as
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 4 Byte AS Number. The type is interface{} with range: 0..4294967295.
    FourByteAs interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetFilter() yfilter.YFilter { return fourByteAs.YFilter }

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) SetFilter(yf yfilter.YFilter) { fourByteAs.YFilter = yf }

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetGoName(yname string) string {
    if yname == "four-byte-as" { return "FourByteAs" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetSegmentPath() string {
    return "four-byte-as"
}

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["four-byte-as"] = fourByteAs.FourByteAs
    leafs["two-byte-index"] = fourByteAs.TwoByteIndex
    return leafs
}

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetBundleName() string { return "cisco_ios_xr" }

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetYangName() string { return "four-byte-as" }

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) SetParent(parent types.Entity) { fourByteAs.parent = parent }

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetParent() types.Entity { return fourByteAs.parent }

func (fourByteAs *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_FourByteAs) GetParentYangName() string { return "route-target" }

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr
// v4 addr
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // 2 Byte Index. The type is interface{} with range: 0..65535.
    TwoByteIndex interface{}
}

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetFilter() yfilter.YFilter { return v4Addr.YFilter }

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) SetFilter(yf yfilter.YFilter) { v4Addr.YFilter = yf }

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "two-byte-index" { return "TwoByteIndex" }
    return ""
}

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetSegmentPath() string {
    return "v4-addr"
}

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = v4Addr.Ipv4Address
    leafs["two-byte-index"] = v4Addr.TwoByteIndex
    return leafs
}

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetBundleName() string { return "cisco_ios_xr" }

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetYangName() string { return "v4-addr" }

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) SetParent(parent types.Entity) { v4Addr.parent = parent }

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetParent() types.Entity { return v4Addr.parent }

func (v4Addr *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_V4Addr) GetParentYangName() string { return "route-target" }

// Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport
// es import
type Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top 4 bytes of ES Import. The type is interface{} with range:
    // 0..4294967295.
    HighBytes interface{}

    // Low 2 bytes of ES Import. The type is interface{} with range: 0..65535.
    LowBytes interface{}
}

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetFilter() yfilter.YFilter { return esImport.YFilter }

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) SetFilter(yf yfilter.YFilter) { esImport.YFilter = yf }

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetGoName(yname string) string {
    if yname == "high-bytes" { return "HighBytes" }
    if yname == "low-bytes" { return "LowBytes" }
    return ""
}

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetSegmentPath() string {
    return "es-import"
}

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high-bytes"] = esImport.HighBytes
    leafs["low-bytes"] = esImport.LowBytes
    return leafs
}

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetBundleName() string { return "cisco_ios_xr" }

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetYangName() string { return "es-import" }

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) SetParent(parent types.Entity) { esImport.parent = parent }

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetParent() types.Entity { return esImport.parent }

func (esImport *Evpn_Standby_EviDetail_EviChildren_RouteTargets_RouteTarget_RouteTarget_EsImport) GetParentYangName() string { return "route-target" }

// Evpn_Standby_EviDetail_EviChildren_Macs
// L2VPN EVPN EVI MAC table
type Evpn_Standby_EviDetail_EviChildren_Macs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // L2VPN EVPN MAC table. The type is slice of
    // Evpn_Standby_EviDetail_EviChildren_Macs_Mac.
    Mac []Evpn_Standby_EviDetail_EviChildren_Macs_Mac
}

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetFilter() yfilter.YFilter { return macs.YFilter }

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) SetFilter(yf yfilter.YFilter) { macs.YFilter = yf }

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetGoName(yname string) string {
    if yname == "mac" { return "Mac" }
    return ""
}

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetSegmentPath() string {
    return "macs"
}

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac" {
        for _, c := range macs.Mac {
            if macs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_Macs_Mac{}
        macs.Mac = append(macs.Mac, child)
        return &macs.Mac[len(macs.Mac)-1]
    }
    return nil
}

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macs.Mac {
        children[macs.Mac[i].GetSegmentPath()] = &macs.Mac[i]
    }
    return children
}

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetBundleName() string { return "cisco_ios_xr" }

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetYangName() string { return "macs" }

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) SetParent(parent types.Entity) { macs.parent = parent }

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetParent() types.Entity { return macs.parent }

func (macs *Evpn_Standby_EviDetail_EviChildren_Macs) GetParentYangName() string { return "evi-children" }

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac
// L2VPN EVPN MAC table
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // Ethernet Tag ID. The type is interface{} with range:
    // -2147483648..2147483647.
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

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "ethernet-tag-xr" { return "EthernetTagXr" }
    if yname == "mac-address-xr" { return "MacAddressXr" }
    if yname == "ip-address-xr" { return "IpAddressXr" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "num-paths" { return "NumPaths" }
    if yname == "is-local-mac" { return "IsLocalMac" }
    if yname == "is-proxy-entry" { return "IsProxyEntry" }
    if yname == "is-remote-mac" { return "IsRemoteMac" }
    if yname == "soo-nexthop" { return "SooNexthop" }
    if yname == "ipnh-address" { return "IpnhAddress" }
    if yname == "esi-port-key" { return "EsiPortKey" }
    if yname == "local-encap-type" { return "LocalEncapType" }
    if yname == "remote-encap-type" { return "RemoteEncapType" }
    if yname == "learned-bridge-port-name" { return "LearnedBridgePortName" }
    if yname == "local-seq-id" { return "LocalSeqId" }
    if yname == "remote-seq-id" { return "RemoteSeqId" }
    if yname == "local-l3-label" { return "LocalL3Label" }
    if yname == "router-mac-address" { return "RouterMacAddress" }
    if yname == "mac-flush-requested" { return "MacFlushRequested" }
    if yname == "mac-flush-received" { return "MacFlushReceived" }
    if yname == "internal-label" { return "InternalLabel" }
    if yname == "resolved" { return "Resolved" }
    if yname == "local-is-static" { return "LocalIsStatic" }
    if yname == "remote-is-static" { return "RemoteIsStatic" }
    if yname == "local-ethernet-segment-identifier" { return "LocalEthernetSegmentIdentifier" }
    if yname == "remote-ethernet-segment-identifier" { return "RemoteEthernetSegmentIdentifier" }
    if yname == "path-buffer" { return "PathBuffer" }
    return ""
}

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-ethernet-segment-identifier" {
        for _, c := range mac.LocalEthernetSegmentIdentifier {
            if mac.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier{}
        mac.LocalEthernetSegmentIdentifier = append(mac.LocalEthernetSegmentIdentifier, child)
        return &mac.LocalEthernetSegmentIdentifier[len(mac.LocalEthernetSegmentIdentifier)-1]
    }
    if childYangName == "remote-ethernet-segment-identifier" {
        for _, c := range mac.RemoteEthernetSegmentIdentifier {
            if mac.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier{}
        mac.RemoteEthernetSegmentIdentifier = append(mac.RemoteEthernetSegmentIdentifier, child)
        return &mac.RemoteEthernetSegmentIdentifier[len(mac.RemoteEthernetSegmentIdentifier)-1]
    }
    if childYangName == "path-buffer" {
        for _, c := range mac.PathBuffer {
            if mac.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer{}
        mac.PathBuffer = append(mac.PathBuffer, child)
        return &mac.PathBuffer[len(mac.PathBuffer)-1]
    }
    return nil
}

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mac.LocalEthernetSegmentIdentifier {
        children[mac.LocalEthernetSegmentIdentifier[i].GetSegmentPath()] = &mac.LocalEthernetSegmentIdentifier[i]
    }
    for i := range mac.RemoteEthernetSegmentIdentifier {
        children[mac.RemoteEthernetSegmentIdentifier[i].GetSegmentPath()] = &mac.RemoteEthernetSegmentIdentifier[i]
    }
    for i := range mac.PathBuffer {
        children[mac.PathBuffer[i].GetSegmentPath()] = &mac.PathBuffer[i]
    }
    return children
}

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = mac.Evi
    leafs["ethernet-tag"] = mac.EthernetTag
    leafs["mac-address"] = mac.MacAddress
    leafs["ip-address"] = mac.IpAddress
    leafs["ethernet-tag-xr"] = mac.EthernetTagXr
    leafs["mac-address-xr"] = mac.MacAddressXr
    leafs["ip-address-xr"] = mac.IpAddressXr
    leafs["local-label"] = mac.LocalLabel
    leafs["num-paths"] = mac.NumPaths
    leafs["is-local-mac"] = mac.IsLocalMac
    leafs["is-proxy-entry"] = mac.IsProxyEntry
    leafs["is-remote-mac"] = mac.IsRemoteMac
    leafs["soo-nexthop"] = mac.SooNexthop
    leafs["ipnh-address"] = mac.IpnhAddress
    leafs["esi-port-key"] = mac.EsiPortKey
    leafs["local-encap-type"] = mac.LocalEncapType
    leafs["remote-encap-type"] = mac.RemoteEncapType
    leafs["learned-bridge-port-name"] = mac.LearnedBridgePortName
    leafs["local-seq-id"] = mac.LocalSeqId
    leafs["remote-seq-id"] = mac.RemoteSeqId
    leafs["local-l3-label"] = mac.LocalL3Label
    leafs["router-mac-address"] = mac.RouterMacAddress
    leafs["mac-flush-requested"] = mac.MacFlushRequested
    leafs["mac-flush-received"] = mac.MacFlushReceived
    leafs["internal-label"] = mac.InternalLabel
    leafs["resolved"] = mac.Resolved
    leafs["local-is-static"] = mac.LocalIsStatic
    leafs["remote-is-static"] = mac.RemoteIsStatic
    return leafs
}

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetYangName() string { return "mac" }

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetParent() types.Entity { return mac.parent }

func (mac *Evpn_Standby_EviDetail_EviChildren_Macs_Mac) GetParentYangName() string { return "macs" }

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier
// Local Ethernet Segment id
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return localEthernetSegmentIdentifier.YFilter }

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { localEthernetSegmentIdentifier.YFilter = yf }

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetSegmentPath() string {
    return "local-ethernet-segment-identifier"
}

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = localEthernetSegmentIdentifier.Entry
    return leafs
}

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetYangName() string { return "local-ethernet-segment-identifier" }

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) SetParent(parent types.Entity) { localEthernetSegmentIdentifier.parent = parent }

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetParent() types.Entity { return localEthernetSegmentIdentifier.parent }

func (localEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_LocalEthernetSegmentIdentifier) GetParentYangName() string { return "mac" }

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier
// Remote Ethernet Segment id
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return remoteEthernetSegmentIdentifier.YFilter }

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { remoteEthernetSegmentIdentifier.YFilter = yf }

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetSegmentPath() string {
    return "remote-ethernet-segment-identifier"
}

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = remoteEthernetSegmentIdentifier.Entry
    return leafs
}

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetYangName() string { return "remote-ethernet-segment-identifier" }

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) SetParent(parent types.Entity) { remoteEthernetSegmentIdentifier.parent = parent }

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetParent() types.Entity { return remoteEthernetSegmentIdentifier.parent }

func (remoteEthernetSegmentIdentifier *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_RemoteEthernetSegmentIdentifier) GetParentYangName() string { return "mac" }

// Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer
// Path List Buffer
type Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Output Label. The type is interface{} with range: 0..4294967295.
    OutputLabel interface{}

    // Segment-Routing Traffic Engineering Tunnel Interface Handle. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    SrteTunnel interface{}
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetFilter() yfilter.YFilter { return pathBuffer.YFilter }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) SetFilter(yf yfilter.YFilter) { pathBuffer.YFilter = yf }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "output-label" { return "OutputLabel" }
    if yname == "srte-tunnel" { return "SrteTunnel" }
    return ""
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetSegmentPath() string {
    return "path-buffer"
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = pathBuffer.NextHop
    leafs["output-label"] = pathBuffer.OutputLabel
    leafs["srte-tunnel"] = pathBuffer.SrteTunnel
    return leafs
}

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetYangName() string { return "path-buffer" }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) SetParent(parent types.Entity) { pathBuffer.parent = parent }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetParent() types.Entity { return pathBuffer.parent }

func (pathBuffer *Evpn_Standby_EviDetail_EviChildren_Macs_Mac_PathBuffer) GetParentYangName() string { return "mac" }

// Evpn_Standby_EthernetSegments
// EVPN Ethernet-Segment Table
type Evpn_Standby_EthernetSegments struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN Ethernet-Segment Entry. The type is slice of
    // Evpn_Standby_EthernetSegments_EthernetSegment.
    EthernetSegment []Evpn_Standby_EthernetSegments_EthernetSegment
}

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetFilter() yfilter.YFilter { return ethernetSegments.YFilter }

func (ethernetSegments *Evpn_Standby_EthernetSegments) SetFilter(yf yfilter.YFilter) { ethernetSegments.YFilter = yf }

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetGoName(yname string) string {
    if yname == "ethernet-segment" { return "EthernetSegment" }
    return ""
}

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetSegmentPath() string {
    return "ethernet-segments"
}

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment" {
        for _, c := range ethernetSegments.EthernetSegment {
            if ethernetSegments.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment{}
        ethernetSegments.EthernetSegment = append(ethernetSegments.EthernetSegment, child)
        return &ethernetSegments.EthernetSegment[len(ethernetSegments.EthernetSegment)-1]
    }
    return nil
}

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetSegments.EthernetSegment {
        children[ethernetSegments.EthernetSegment[i].GetSegmentPath()] = &ethernetSegments.EthernetSegment[i]
    }
    return children
}

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetYangName() string { return "ethernet-segments" }

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegments *Evpn_Standby_EthernetSegments) SetParent(parent types.Entity) { ethernetSegments.parent = parent }

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetParent() types.Entity { return ethernetSegments.parent }

func (ethernetSegments *Evpn_Standby_EthernetSegments) GetParentYangName() string { return "standby" }

// Evpn_Standby_EthernetSegments_EthernetSegment
// EVPN Ethernet-Segment Entry
type Evpn_Standby_EthernetSegments_EthernetSegment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
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

    // Ethernet Segment Name. The type is string.
    EthernetSegmentName interface{}

    // State of the ethernet segment. The type is interface{} with range:
    // 0..4294967295.
    EthernetSegmentState interface{}

    // Main port ifhandle. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetFilter() yfilter.YFilter { return ethernetSegment.YFilter }

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) SetFilter(yf yfilter.YFilter) { ethernetSegment.YFilter = yf }

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "esi1" { return "Esi1" }
    if yname == "esi2" { return "Esi2" }
    if yname == "esi3" { return "Esi3" }
    if yname == "esi4" { return "Esi4" }
    if yname == "esi5" { return "Esi5" }
    if yname == "esi-type" { return "EsiType" }
    if yname == "ethernet-segment-name" { return "EthernetSegmentName" }
    if yname == "ethernet-segment-state" { return "EthernetSegmentState" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "main-port-role" { return "MainPortRole" }
    if yname == "main-port-mac" { return "MainPortMac" }
    if yname == "num-up-p-ws" { return "NumUpPWs" }
    if yname == "route-target" { return "RouteTarget" }
    if yname == "rt-origin" { return "RtOrigin" }
    if yname == "es-bgp-gates" { return "EsBgpGates" }
    if yname == "es-l2fib-gates" { return "EsL2FibGates" }
    if yname == "mac-flushing-mode-config" { return "MacFlushingModeConfig" }
    if yname == "load-balance-mode-config" { return "LoadBalanceModeConfig" }
    if yname == "load-balance-mode-is-default" { return "LoadBalanceModeIsDefault" }
    if yname == "load-balance-mode-oper" { return "LoadBalanceModeOper" }
    if yname == "force-single-home" { return "ForceSingleHome" }
    if yname == "source-mac-oper" { return "SourceMacOper" }
    if yname == "source-mac-origin" { return "SourceMacOrigin" }
    if yname == "peering-timer" { return "PeeringTimer" }
    if yname == "peering-timer-left" { return "PeeringTimerLeft" }
    if yname == "recovery-timer" { return "RecoveryTimer" }
    if yname == "recovery-timer-left" { return "RecoveryTimerLeft" }
    if yname == "service-carving-mode" { return "ServiceCarvingMode" }
    if yname == "primary-services-input" { return "PrimaryServicesInput" }
    if yname == "secondary-services-input" { return "SecondaryServicesInput" }
    if yname == "forwarder-ports" { return "ForwarderPorts" }
    if yname == "permanent-forwarder-ports" { return "PermanentForwarderPorts" }
    if yname == "elected-forwarder-ports" { return "ElectedForwarderPorts" }
    if yname == "not-elected-forwarder-ports" { return "NotElectedForwarderPorts" }
    if yname == "not-config-forwarder-ports" { return "NotConfigForwarderPorts" }
    if yname == "mp-protected" { return "MpProtected" }
    if yname == "nve-anycast-vtep" { return "NveAnycastVtep" }
    if yname == "nve-ingress-replication" { return "NveIngressReplication" }
    if yname == "local-split-horizon-group-label" { return "LocalSplitHorizonGroupLabel" }
    if yname == "ethernet-segment-identifier" { return "EthernetSegmentIdentifier" }
    if yname == "primary-service" { return "PrimaryService" }
    if yname == "secondary-service" { return "SecondaryService" }
    if yname == "service-carving-i-sidelected-result" { return "ServiceCarvingISidelectedResult" }
    if yname == "service-carving-isid-not-elected-result" { return "ServiceCarvingIsidNotElectedResult" }
    if yname == "service-carving-evi-elected-result" { return "ServiceCarvingEviElectedResult" }
    if yname == "service-carving-evi-not-elected-result" { return "ServiceCarvingEviNotElectedResult" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "service-carving-vpws-permanent-result" { return "ServiceCarvingVpwsPermanentResult" }
    if yname == "remote-split-horizon-group-label" { return "RemoteSplitHorizonGroupLabel" }
    return ""
}

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetSegmentPath() string {
    return "ethernet-segment"
}

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-segment-identifier" {
        for _, c := range ethernetSegment.EthernetSegmentIdentifier {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier{}
        ethernetSegment.EthernetSegmentIdentifier = append(ethernetSegment.EthernetSegmentIdentifier, child)
        return &ethernetSegment.EthernetSegmentIdentifier[len(ethernetSegment.EthernetSegmentIdentifier)-1]
    }
    if childYangName == "primary-service" {
        for _, c := range ethernetSegment.PrimaryService {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService{}
        ethernetSegment.PrimaryService = append(ethernetSegment.PrimaryService, child)
        return &ethernetSegment.PrimaryService[len(ethernetSegment.PrimaryService)-1]
    }
    if childYangName == "secondary-service" {
        for _, c := range ethernetSegment.SecondaryService {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService{}
        ethernetSegment.SecondaryService = append(ethernetSegment.SecondaryService, child)
        return &ethernetSegment.SecondaryService[len(ethernetSegment.SecondaryService)-1]
    }
    if childYangName == "service-carving-i-sidelected-result" {
        for _, c := range ethernetSegment.ServiceCarvingISidelectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult{}
        ethernetSegment.ServiceCarvingISidelectedResult = append(ethernetSegment.ServiceCarvingISidelectedResult, child)
        return &ethernetSegment.ServiceCarvingISidelectedResult[len(ethernetSegment.ServiceCarvingISidelectedResult)-1]
    }
    if childYangName == "service-carving-isid-not-elected-result" {
        for _, c := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult{}
        ethernetSegment.ServiceCarvingIsidNotElectedResult = append(ethernetSegment.ServiceCarvingIsidNotElectedResult, child)
        return &ethernetSegment.ServiceCarvingIsidNotElectedResult[len(ethernetSegment.ServiceCarvingIsidNotElectedResult)-1]
    }
    if childYangName == "service-carving-evi-elected-result" {
        for _, c := range ethernetSegment.ServiceCarvingEviElectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult{}
        ethernetSegment.ServiceCarvingEviElectedResult = append(ethernetSegment.ServiceCarvingEviElectedResult, child)
        return &ethernetSegment.ServiceCarvingEviElectedResult[len(ethernetSegment.ServiceCarvingEviElectedResult)-1]
    }
    if childYangName == "service-carving-evi-not-elected-result" {
        for _, c := range ethernetSegment.ServiceCarvingEviNotElectedResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult{}
        ethernetSegment.ServiceCarvingEviNotElectedResult = append(ethernetSegment.ServiceCarvingEviNotElectedResult, child)
        return &ethernetSegment.ServiceCarvingEviNotElectedResult[len(ethernetSegment.ServiceCarvingEviNotElectedResult)-1]
    }
    if childYangName == "next-hop" {
        for _, c := range ethernetSegment.NextHop {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_NextHop{}
        ethernetSegment.NextHop = append(ethernetSegment.NextHop, child)
        return &ethernetSegment.NextHop[len(ethernetSegment.NextHop)-1]
    }
    if childYangName == "service-carving-vpws-permanent-result" {
        for _, c := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult{}
        ethernetSegment.ServiceCarvingVpwsPermanentResult = append(ethernetSegment.ServiceCarvingVpwsPermanentResult, child)
        return &ethernetSegment.ServiceCarvingVpwsPermanentResult[len(ethernetSegment.ServiceCarvingVpwsPermanentResult)-1]
    }
    if childYangName == "remote-split-horizon-group-label" {
        for _, c := range ethernetSegment.RemoteSplitHorizonGroupLabel {
            if ethernetSegment.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel{}
        ethernetSegment.RemoteSplitHorizonGroupLabel = append(ethernetSegment.RemoteSplitHorizonGroupLabel, child)
        return &ethernetSegment.RemoteSplitHorizonGroupLabel[len(ethernetSegment.RemoteSplitHorizonGroupLabel)-1]
    }
    return nil
}

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetSegment.EthernetSegmentIdentifier {
        children[ethernetSegment.EthernetSegmentIdentifier[i].GetSegmentPath()] = &ethernetSegment.EthernetSegmentIdentifier[i]
    }
    for i := range ethernetSegment.PrimaryService {
        children[ethernetSegment.PrimaryService[i].GetSegmentPath()] = &ethernetSegment.PrimaryService[i]
    }
    for i := range ethernetSegment.SecondaryService {
        children[ethernetSegment.SecondaryService[i].GetSegmentPath()] = &ethernetSegment.SecondaryService[i]
    }
    for i := range ethernetSegment.ServiceCarvingISidelectedResult {
        children[ethernetSegment.ServiceCarvingISidelectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingISidelectedResult[i]
    }
    for i := range ethernetSegment.ServiceCarvingIsidNotElectedResult {
        children[ethernetSegment.ServiceCarvingIsidNotElectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingIsidNotElectedResult[i]
    }
    for i := range ethernetSegment.ServiceCarvingEviElectedResult {
        children[ethernetSegment.ServiceCarvingEviElectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingEviElectedResult[i]
    }
    for i := range ethernetSegment.ServiceCarvingEviNotElectedResult {
        children[ethernetSegment.ServiceCarvingEviNotElectedResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingEviNotElectedResult[i]
    }
    for i := range ethernetSegment.NextHop {
        children[ethernetSegment.NextHop[i].GetSegmentPath()] = &ethernetSegment.NextHop[i]
    }
    for i := range ethernetSegment.ServiceCarvingVpwsPermanentResult {
        children[ethernetSegment.ServiceCarvingVpwsPermanentResult[i].GetSegmentPath()] = &ethernetSegment.ServiceCarvingVpwsPermanentResult[i]
    }
    for i := range ethernetSegment.RemoteSplitHorizonGroupLabel {
        children[ethernetSegment.RemoteSplitHorizonGroupLabel[i].GetSegmentPath()] = &ethernetSegment.RemoteSplitHorizonGroupLabel[i]
    }
    return children
}

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = ethernetSegment.InterfaceName
    leafs["esi1"] = ethernetSegment.Esi1
    leafs["esi2"] = ethernetSegment.Esi2
    leafs["esi3"] = ethernetSegment.Esi3
    leafs["esi4"] = ethernetSegment.Esi4
    leafs["esi5"] = ethernetSegment.Esi5
    leafs["esi-type"] = ethernetSegment.EsiType
    leafs["ethernet-segment-name"] = ethernetSegment.EthernetSegmentName
    leafs["ethernet-segment-state"] = ethernetSegment.EthernetSegmentState
    leafs["if-handle"] = ethernetSegment.IfHandle
    leafs["main-port-role"] = ethernetSegment.MainPortRole
    leafs["main-port-mac"] = ethernetSegment.MainPortMac
    leafs["num-up-p-ws"] = ethernetSegment.NumUpPWs
    leafs["route-target"] = ethernetSegment.RouteTarget
    leafs["rt-origin"] = ethernetSegment.RtOrigin
    leafs["es-bgp-gates"] = ethernetSegment.EsBgpGates
    leafs["es-l2fib-gates"] = ethernetSegment.EsL2FibGates
    leafs["mac-flushing-mode-config"] = ethernetSegment.MacFlushingModeConfig
    leafs["load-balance-mode-config"] = ethernetSegment.LoadBalanceModeConfig
    leafs["load-balance-mode-is-default"] = ethernetSegment.LoadBalanceModeIsDefault
    leafs["load-balance-mode-oper"] = ethernetSegment.LoadBalanceModeOper
    leafs["force-single-home"] = ethernetSegment.ForceSingleHome
    leafs["source-mac-oper"] = ethernetSegment.SourceMacOper
    leafs["source-mac-origin"] = ethernetSegment.SourceMacOrigin
    leafs["peering-timer"] = ethernetSegment.PeeringTimer
    leafs["peering-timer-left"] = ethernetSegment.PeeringTimerLeft
    leafs["recovery-timer"] = ethernetSegment.RecoveryTimer
    leafs["recovery-timer-left"] = ethernetSegment.RecoveryTimerLeft
    leafs["service-carving-mode"] = ethernetSegment.ServiceCarvingMode
    leafs["primary-services-input"] = ethernetSegment.PrimaryServicesInput
    leafs["secondary-services-input"] = ethernetSegment.SecondaryServicesInput
    leafs["forwarder-ports"] = ethernetSegment.ForwarderPorts
    leafs["permanent-forwarder-ports"] = ethernetSegment.PermanentForwarderPorts
    leafs["elected-forwarder-ports"] = ethernetSegment.ElectedForwarderPorts
    leafs["not-elected-forwarder-ports"] = ethernetSegment.NotElectedForwarderPorts
    leafs["not-config-forwarder-ports"] = ethernetSegment.NotConfigForwarderPorts
    leafs["mp-protected"] = ethernetSegment.MpProtected
    leafs["nve-anycast-vtep"] = ethernetSegment.NveAnycastVtep
    leafs["nve-ingress-replication"] = ethernetSegment.NveIngressReplication
    leafs["local-split-horizon-group-label"] = ethernetSegment.LocalSplitHorizonGroupLabel
    return leafs
}

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetYangName() string { return "ethernet-segment" }

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) SetParent(parent types.Entity) { ethernetSegment.parent = parent }

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetParent() types.Entity { return ethernetSegment.parent }

func (ethernetSegment *Evpn_Standby_EthernetSegments_EthernetSegment) GetParentYangName() string { return "ethernet-segments" }

// Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier
// Ethernet Segment id
type Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..255.
    Entry interface{}
}

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetFilter() yfilter.YFilter { return ethernetSegmentIdentifier.YFilter }

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) SetFilter(yf yfilter.YFilter) { ethernetSegmentIdentifier.YFilter = yf }

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetSegmentPath() string {
    return "ethernet-segment-identifier"
}

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = ethernetSegmentIdentifier.Entry
    return leafs
}

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetYangName() string { return "ethernet-segment-identifier" }

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) SetParent(parent types.Entity) { ethernetSegmentIdentifier.parent = parent }

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetParent() types.Entity { return ethernetSegmentIdentifier.parent }

func (ethernetSegmentIdentifier *Evpn_Standby_EthernetSegments_EthernetSegment_EthernetSegmentIdentifier) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService
// List of Primary services ESI/I-SIDs
type Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetFilter() yfilter.YFilter { return primaryService.YFilter }

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) SetFilter(yf yfilter.YFilter) { primaryService.YFilter = yf }

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetSegmentPath() string {
    return "primary-service"
}

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = primaryService.Entry
    return leafs
}

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetBundleName() string { return "cisco_ios_xr" }

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetYangName() string { return "primary-service" }

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) SetParent(parent types.Entity) { primaryService.parent = parent }

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetParent() types.Entity { return primaryService.parent }

func (primaryService *Evpn_Standby_EthernetSegments_EthernetSegment_PrimaryService) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService
// List of Secondary services ESI/I-SIDs
type Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetFilter() yfilter.YFilter { return secondaryService.YFilter }

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) SetFilter(yf yfilter.YFilter) { secondaryService.YFilter = yf }

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetSegmentPath() string {
    return "secondary-service"
}

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = secondaryService.Entry
    return leafs
}

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetYangName() string { return "secondary-service" }

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) SetParent(parent types.Entity) { secondaryService.parent = parent }

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetParent() types.Entity { return secondaryService.parent }

func (secondaryService *Evpn_Standby_EthernetSegments_EthernetSegment_SecondaryService) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult
// Elected ISID service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetFilter() yfilter.YFilter { return serviceCarvingISidelectedResult.YFilter }

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingISidelectedResult.YFilter = yf }

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetSegmentPath() string {
    return "service-carving-i-sidelected-result"
}

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingISidelectedResult.Entry
    return leafs
}

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetYangName() string { return "service-carving-i-sidelected-result" }

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) SetParent(parent types.Entity) { serviceCarvingISidelectedResult.parent = parent }

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetParent() types.Entity { return serviceCarvingISidelectedResult.parent }

func (serviceCarvingISidelectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingISidelectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult
// Not elected ISID service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetFilter() yfilter.YFilter { return serviceCarvingIsidNotElectedResult.YFilter }

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingIsidNotElectedResult.YFilter = yf }

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetSegmentPath() string {
    return "service-carving-isid-not-elected-result"
}

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingIsidNotElectedResult.Entry
    return leafs
}

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetYangName() string { return "service-carving-isid-not-elected-result" }

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) SetParent(parent types.Entity) { serviceCarvingIsidNotElectedResult.parent = parent }

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetParent() types.Entity { return serviceCarvingIsidNotElectedResult.parent }

func (serviceCarvingIsidNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingIsidNotElectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult
// Elected EVI service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetFilter() yfilter.YFilter { return serviceCarvingEviElectedResult.YFilter }

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingEviElectedResult.YFilter = yf }

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetSegmentPath() string {
    return "service-carving-evi-elected-result"
}

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingEviElectedResult.Entry
    return leafs
}

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetYangName() string { return "service-carving-evi-elected-result" }

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) SetParent(parent types.Entity) { serviceCarvingEviElectedResult.parent = parent }

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetParent() types.Entity { return serviceCarvingEviElectedResult.parent }

func (serviceCarvingEviElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviElectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult
// Not elected EVI service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetFilter() yfilter.YFilter { return serviceCarvingEviNotElectedResult.YFilter }

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) SetFilter(yf yfilter.YFilter) { serviceCarvingEviNotElectedResult.YFilter = yf }

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    return ""
}

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetSegmentPath() string {
    return "service-carving-evi-not-elected-result"
}

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = serviceCarvingEviNotElectedResult.Entry
    return leafs
}

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetYangName() string { return "service-carving-evi-not-elected-result" }

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) SetParent(parent types.Entity) { serviceCarvingEviNotElectedResult.parent = parent }

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetParent() types.Entity { return serviceCarvingEviNotElectedResult.parent }

func (serviceCarvingEviNotElectedResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingEviNotElectedResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_EthernetSegments_EthernetSegment_NextHop
// List of nexthop IPv6 addresses
type Evpn_Standby_EthernetSegments_EthernetSegment_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}
}

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHop.NextHop
    return leafs
}

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *Evpn_Standby_EthernetSegments_EthernetSegment_NextHop) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult
// Permanent EVPN VPWS service carving results
type Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VPN ID. The type is interface{} with range: 0..4294967295.
    VpnId interface{}

    // Service Type. The type is L2vpnEvpn.
    Type interface{}

    // Ethernet Tag. The type is interface{} with range: 0..4294967295.
    EthernetTag interface{}
}

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetFilter() yfilter.YFilter { return serviceCarvingVpwsPermanentResult.YFilter }

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) SetFilter(yf yfilter.YFilter) { serviceCarvingVpwsPermanentResult.YFilter = yf }

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetGoName(yname string) string {
    if yname == "vpn-id" { return "VpnId" }
    if yname == "type" { return "Type" }
    if yname == "ethernet-tag" { return "EthernetTag" }
    return ""
}

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetSegmentPath() string {
    return "service-carving-vpws-permanent-result"
}

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vpn-id"] = serviceCarvingVpwsPermanentResult.VpnId
    leafs["type"] = serviceCarvingVpwsPermanentResult.Type
    leafs["ethernet-tag"] = serviceCarvingVpwsPermanentResult.EthernetTag
    return leafs
}

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetBundleName() string { return "cisco_ios_xr" }

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetYangName() string { return "service-carving-vpws-permanent-result" }

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) SetParent(parent types.Entity) { serviceCarvingVpwsPermanentResult.parent = parent }

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetParent() types.Entity { return serviceCarvingVpwsPermanentResult.parent }

func (serviceCarvingVpwsPermanentResult *Evpn_Standby_EthernetSegments_EthernetSegment_ServiceCarvingVpwsPermanentResult) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel
// Remote split horizon group labels
type Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-hop IP address (v6 format). The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Split horizon label associated with next-hop address. The type is
    // interface{} with range: 0..4294967295.
    Label interface{}
}

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetFilter() yfilter.YFilter { return remoteSplitHorizonGroupLabel.YFilter }

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) SetFilter(yf yfilter.YFilter) { remoteSplitHorizonGroupLabel.YFilter = yf }

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "label" { return "Label" }
    return ""
}

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetSegmentPath() string {
    return "remote-split-horizon-group-label"
}

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = remoteSplitHorizonGroupLabel.NextHop
    leafs["label"] = remoteSplitHorizonGroupLabel.Label
    return leafs
}

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetBundleName() string { return "cisco_ios_xr" }

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetYangName() string { return "remote-split-horizon-group-label" }

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) SetParent(parent types.Entity) { remoteSplitHorizonGroupLabel.parent = parent }

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetParent() types.Entity { return remoteSplitHorizonGroupLabel.parent }

func (remoteSplitHorizonGroupLabel *Evpn_Standby_EthernetSegments_EthernetSegment_RemoteSplitHorizonGroupLabel) GetParentYangName() string { return "ethernet-segment" }

// Evpn_Standby_AcIds
// EVPN AC ID table
type Evpn_Standby_AcIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN AC ID table. The type is slice of Evpn_Standby_AcIds_AcId.
    AcId []Evpn_Standby_AcIds_AcId
}

func (acIds *Evpn_Standby_AcIds) GetFilter() yfilter.YFilter { return acIds.YFilter }

func (acIds *Evpn_Standby_AcIds) SetFilter(yf yfilter.YFilter) { acIds.YFilter = yf }

func (acIds *Evpn_Standby_AcIds) GetGoName(yname string) string {
    if yname == "ac-id" { return "AcId" }
    return ""
}

func (acIds *Evpn_Standby_AcIds) GetSegmentPath() string {
    return "ac-ids"
}

func (acIds *Evpn_Standby_AcIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ac-id" {
        for _, c := range acIds.AcId {
            if acIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Evpn_Standby_AcIds_AcId{}
        acIds.AcId = append(acIds.AcId, child)
        return &acIds.AcId[len(acIds.AcId)-1]
    }
    return nil
}

func (acIds *Evpn_Standby_AcIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range acIds.AcId {
        children[acIds.AcId[i].GetSegmentPath()] = &acIds.AcId[i]
    }
    return children
}

func (acIds *Evpn_Standby_AcIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (acIds *Evpn_Standby_AcIds) GetBundleName() string { return "cisco_ios_xr" }

func (acIds *Evpn_Standby_AcIds) GetYangName() string { return "ac-ids" }

func (acIds *Evpn_Standby_AcIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acIds *Evpn_Standby_AcIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acIds *Evpn_Standby_AcIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acIds *Evpn_Standby_AcIds) SetParent(parent types.Entity) { acIds.parent = parent }

func (acIds *Evpn_Standby_AcIds) GetParent() types.Entity { return acIds.parent }

func (acIds *Evpn_Standby_AcIds) GetParentYangName() string { return "standby" }

// Evpn_Standby_AcIds_AcId
// EVPN AC ID table
type Evpn_Standby_AcIds_AcId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EVPN id. The type is interface{} with range: -2147483648..2147483647.
    Evi interface{}

    // AC ID. The type is interface{} with range: -2147483648..2147483647.
    AcId interface{}

    // E-VPN id. The type is interface{} with range: 0..4294967295.
    EviXr interface{}

    // Neighbor IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Neighbor interface{}
}

func (acId *Evpn_Standby_AcIds_AcId) GetFilter() yfilter.YFilter { return acId.YFilter }

func (acId *Evpn_Standby_AcIds_AcId) SetFilter(yf yfilter.YFilter) { acId.YFilter = yf }

func (acId *Evpn_Standby_AcIds_AcId) GetGoName(yname string) string {
    if yname == "evi" { return "Evi" }
    if yname == "ac-id" { return "AcId" }
    if yname == "evi-xr" { return "EviXr" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (acId *Evpn_Standby_AcIds_AcId) GetSegmentPath() string {
    return "ac-id"
}

func (acId *Evpn_Standby_AcIds_AcId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acId *Evpn_Standby_AcIds_AcId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acId *Evpn_Standby_AcIds_AcId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["evi"] = acId.Evi
    leafs["ac-id"] = acId.AcId
    leafs["evi-xr"] = acId.EviXr
    leafs["neighbor"] = acId.Neighbor
    return leafs
}

func (acId *Evpn_Standby_AcIds_AcId) GetBundleName() string { return "cisco_ios_xr" }

func (acId *Evpn_Standby_AcIds_AcId) GetYangName() string { return "ac-id" }

func (acId *Evpn_Standby_AcIds_AcId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acId *Evpn_Standby_AcIds_AcId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acId *Evpn_Standby_AcIds_AcId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acId *Evpn_Standby_AcIds_AcId) SetParent(parent types.Entity) { acId.parent = parent }

func (acId *Evpn_Standby_AcIds_AcId) GetParent() types.Entity { return acId.parent }

func (acId *Evpn_Standby_AcIds_AcId) GetParentYangName() string { return "ac-ids" }

