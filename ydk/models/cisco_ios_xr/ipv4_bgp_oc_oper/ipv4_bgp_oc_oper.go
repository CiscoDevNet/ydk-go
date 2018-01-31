// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-bgp-oc package operational data.
// 
// This module contains definitions
// for the following management objects:
//   oc-bgp: OC-BGP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_bgp_oc_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_bgp_oc_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-bgp-oc-oper oc-bgp}", reflect.TypeOf(OcBgp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp", reflect.TypeOf(OcBgp{}))
}

// BgpOcOriginAttr represents Origin Type
type BgpOcOriginAttr string

const (
    // IGP
    BgpOcOriginAttr_igp BgpOcOriginAttr = "igp"

    // EGP
    BgpOcOriginAttr_egp BgpOcOriginAttr = "egp"

    // Incomplete
    BgpOcOriginAttr_incomplete BgpOcOriginAttr = "incomplete"
)

// BgpOcInvalidRouteReason represents Invalid route reason
type BgpOcInvalidRouteReason string

const (
    // Valid route
    BgpOcInvalidRouteReason_valid_route BgpOcInvalidRouteReason = "valid-route"

    // ClusterLoop
    BgpOcInvalidRouteReason_invalid_clsuter_loop BgpOcInvalidRouteReason = "invalid-clsuter-loop"

    // AsPathLoop
    BgpOcInvalidRouteReason_invalid_as_path_loop BgpOcInvalidRouteReason = "invalid-as-path-loop"

    // OriginatorID
    BgpOcInvalidRouteReason_invalid_origin_at_or_id BgpOcInvalidRouteReason = "invalid-origin-at-or-id"

    // ASConfedLoop
    BgpOcInvalidRouteReason_invalid_as_confed_loop BgpOcInvalidRouteReason = "invalid-as-confed-loop"
)

// BgpOcAfi represents BGP Address family
type BgpOcAfi string

const (
    // IPv4 unicast
    BgpOcAfi_ipv4 BgpOcAfi = "ipv4"

    // IPv6 unicast
    BgpOcAfi_ipv6 BgpOcAfi = "ipv6"
)

// OcBgp
// OC-BGP operational data
type OcBgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP-RIB operational data.
    BgpRib OcBgp_BgpRib
}

func (ocBgp *OcBgp) GetFilter() yfilter.YFilter { return ocBgp.YFilter }

func (ocBgp *OcBgp) SetFilter(yf yfilter.YFilter) { ocBgp.YFilter = yf }

func (ocBgp *OcBgp) GetGoName(yname string) string {
    if yname == "bgp-rib" { return "BgpRib" }
    return ""
}

func (ocBgp *OcBgp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp"
}

func (ocBgp *OcBgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-rib" {
        return &ocBgp.BgpRib
    }
    return nil
}

func (ocBgp *OcBgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp-rib"] = &ocBgp.BgpRib
    return children
}

func (ocBgp *OcBgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ocBgp *OcBgp) GetBundleName() string { return "cisco_ios_xr" }

func (ocBgp *OcBgp) GetYangName() string { return "oc-bgp" }

func (ocBgp *OcBgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ocBgp *OcBgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ocBgp *OcBgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ocBgp *OcBgp) SetParent(parent types.Entity) { ocBgp.parent = parent }

func (ocBgp *OcBgp) GetParent() types.Entity { return ocBgp.parent }

func (ocBgp *OcBgp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-bgp-oc-oper" }

// OcBgp_BgpRib
// BGP-RIB operational data
type OcBgp_BgpRib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI-SAFIs information.
    AfiSafiTable OcBgp_BgpRib_AfiSafiTable
}

func (bgpRib *OcBgp_BgpRib) GetFilter() yfilter.YFilter { return bgpRib.YFilter }

func (bgpRib *OcBgp_BgpRib) SetFilter(yf yfilter.YFilter) { bgpRib.YFilter = yf }

func (bgpRib *OcBgp_BgpRib) GetGoName(yname string) string {
    if yname == "afi-safi-table" { return "AfiSafiTable" }
    return ""
}

func (bgpRib *OcBgp_BgpRib) GetSegmentPath() string {
    return "bgp-rib"
}

func (bgpRib *OcBgp_BgpRib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afi-safi-table" {
        return &bgpRib.AfiSafiTable
    }
    return nil
}

func (bgpRib *OcBgp_BgpRib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afi-safi-table"] = &bgpRib.AfiSafiTable
    return children
}

func (bgpRib *OcBgp_BgpRib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpRib *OcBgp_BgpRib) GetBundleName() string { return "cisco_ios_xr" }

func (bgpRib *OcBgp_BgpRib) GetYangName() string { return "bgp-rib" }

func (bgpRib *OcBgp_BgpRib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpRib *OcBgp_BgpRib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpRib *OcBgp_BgpRib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpRib *OcBgp_BgpRib) SetParent(parent types.Entity) { bgpRib.parent = parent }

func (bgpRib *OcBgp_BgpRib) GetParent() types.Entity { return bgpRib.parent }

func (bgpRib *OcBgp_BgpRib) GetParentYangName() string { return "oc-bgp" }

// OcBgp_BgpRib_AfiSafiTable
// AFI-SAFIs information
type OcBgp_BgpRib_AfiSafiTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Unicast.
    Ipv4Unicast OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast

    // IPv6 Unicast.
    Ipv6Unicast OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast
}

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetFilter() yfilter.YFilter { return afiSafiTable.YFilter }

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) SetFilter(yf yfilter.YFilter) { afiSafiTable.YFilter = yf }

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetGoName(yname string) string {
    if yname == "ipv4-unicast" { return "Ipv4Unicast" }
    if yname == "ipv6-unicast" { return "Ipv6Unicast" }
    return ""
}

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetSegmentPath() string {
    return "afi-safi-table"
}

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-unicast" {
        return &afiSafiTable.Ipv4Unicast
    }
    if childYangName == "ipv6-unicast" {
        return &afiSafiTable.Ipv6Unicast
    }
    return nil
}

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-unicast"] = &afiSafiTable.Ipv4Unicast
    children["ipv6-unicast"] = &afiSafiTable.Ipv6Unicast
    return children
}

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetBundleName() string { return "cisco_ios_xr" }

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetYangName() string { return "afi-safi-table" }

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) SetParent(parent types.Entity) { afiSafiTable.parent = parent }

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetParent() types.Entity { return afiSafiTable.parent }

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetParentYangName() string { return "bgp-rib" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast
// IPv4 Unicast
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local rib route table.
    LocRib OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib

    // Neighbor list.
    OpenConfigNeighbors OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors
}

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetFilter() yfilter.YFilter { return ipv4Unicast.YFilter }

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) SetFilter(yf yfilter.YFilter) { ipv4Unicast.YFilter = yf }

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetGoName(yname string) string {
    if yname == "loc-rib" { return "LocRib" }
    if yname == "open-config-neighbors" { return "OpenConfigNeighbors" }
    return ""
}

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetSegmentPath() string {
    return "ipv4-unicast"
}

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "loc-rib" {
        return &ipv4Unicast.LocRib
    }
    if childYangName == "open-config-neighbors" {
        return &ipv4Unicast.OpenConfigNeighbors
    }
    return nil
}

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["loc-rib"] = &ipv4Unicast.LocRib
    children["open-config-neighbors"] = &ipv4Unicast.OpenConfigNeighbors
    return children
}

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetYangName() string { return "ipv4-unicast" }

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) SetParent(parent types.Entity) { ipv4Unicast.parent = parent }

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetParent() types.Entity { return ipv4Unicast.parent }

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetParentYangName() string { return "afi-safi-table" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib
// Local rib route table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetFilter() yfilter.YFilter { return locRib.YFilter }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) SetFilter(yf yfilter.YFilter) { locRib.YFilter = yf }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetSegmentPath() string {
    return "loc-rib"
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &locRib.Routes
    }
    if childYangName == "num-routes" {
        return &locRib.NumRoutes
    }
    return nil
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &locRib.Routes
    children["num-routes"] = &locRib.NumRoutes
    return children
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetBundleName() string { return "cisco_ios_xr" }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetYangName() string { return "loc-rib" }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) SetParent(parent types.Entity) { locRib.parent = parent }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetParent() types.Entity { return locRib.parent }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetParentYangName() string { return "ipv4-unicast" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetParentYangName() string { return "loc-rib" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetParentYangName() string { return "loc-rib" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors
// Neighbor list
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor name. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor.
    OpenConfigNeighbor []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetFilter() yfilter.YFilter { return openConfigNeighbors.YFilter }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) SetFilter(yf yfilter.YFilter) { openConfigNeighbors.YFilter = yf }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetGoName(yname string) string {
    if yname == "open-config-neighbor" { return "OpenConfigNeighbor" }
    return ""
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetSegmentPath() string {
    return "open-config-neighbors"
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "open-config-neighbor" {
        for _, c := range openConfigNeighbors.OpenConfigNeighbor {
            if openConfigNeighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor{}
        openConfigNeighbors.OpenConfigNeighbor = append(openConfigNeighbors.OpenConfigNeighbor, child)
        return &openConfigNeighbors.OpenConfigNeighbor[len(openConfigNeighbors.OpenConfigNeighbor)-1]
    }
    return nil
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range openConfigNeighbors.OpenConfigNeighbor {
        children[openConfigNeighbors.OpenConfigNeighbor[i].GetSegmentPath()] = &openConfigNeighbors.OpenConfigNeighbor[i]
    }
    return children
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetBundleName() string { return "cisco_ios_xr" }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetYangName() string { return "open-config-neighbors" }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) SetParent(parent types.Entity) { openConfigNeighbors.parent = parent }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetParent() types.Entity { return openConfigNeighbors.parent }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetParentYangName() string { return "ipv4-unicast" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor
// Neighbor name
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Adjacency rib in-bound post-policy table.
    AdjRibInPost OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost

    // Adjacency rib out-bound post-policy table.
    AdjRibOutPost OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost

    // Adjacency rib out-bound pre-policy table.
    AdjRibOutPre OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre

    // Adjacency rib in-bound pre-policy table.
    AdjRibInPre OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetFilter() yfilter.YFilter { return openConfigNeighbor.YFilter }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) SetFilter(yf yfilter.YFilter) { openConfigNeighbor.YFilter = yf }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetGoName(yname string) string {
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "adj-rib-in-post" { return "AdjRibInPost" }
    if yname == "adj-rib-out-post" { return "AdjRibOutPost" }
    if yname == "adj-rib-out-pre" { return "AdjRibOutPre" }
    if yname == "adj-rib-in-pre" { return "AdjRibInPre" }
    return ""
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetSegmentPath() string {
    return "open-config-neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", openConfigNeighbor.NeighborAddress) + "']"
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "adj-rib-in-post" {
        return &openConfigNeighbor.AdjRibInPost
    }
    if childYangName == "adj-rib-out-post" {
        return &openConfigNeighbor.AdjRibOutPost
    }
    if childYangName == "adj-rib-out-pre" {
        return &openConfigNeighbor.AdjRibOutPre
    }
    if childYangName == "adj-rib-in-pre" {
        return &openConfigNeighbor.AdjRibInPre
    }
    return nil
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["adj-rib-in-post"] = &openConfigNeighbor.AdjRibInPost
    children["adj-rib-out-post"] = &openConfigNeighbor.AdjRibOutPost
    children["adj-rib-out-pre"] = &openConfigNeighbor.AdjRibOutPre
    children["adj-rib-in-pre"] = &openConfigNeighbor.AdjRibInPre
    return children
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-address"] = openConfigNeighbor.NeighborAddress
    return leafs
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetYangName() string { return "open-config-neighbor" }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) SetParent(parent types.Entity) { openConfigNeighbor.parent = parent }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetParent() types.Entity { return openConfigNeighbor.parent }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetParentYangName() string { return "open-config-neighbors" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost
// Adjacency rib in-bound post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetFilter() yfilter.YFilter { return adjRibInPost.YFilter }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) SetFilter(yf yfilter.YFilter) { adjRibInPost.YFilter = yf }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetSegmentPath() string {
    return "adj-rib-in-post"
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibInPost.Routes
    }
    if childYangName == "num-routes" {
        return &adjRibInPost.NumRoutes
    }
    return nil
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibInPost.Routes
    children["num-routes"] = &adjRibInPost.NumRoutes
    return children
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetBundleName() string { return "cisco_ios_xr" }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetYangName() string { return "adj-rib-in-post" }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) SetParent(parent types.Entity) { adjRibInPost.parent = parent }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetParent() types.Entity { return adjRibInPost.parent }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetParentYangName() string { return "open-config-neighbor" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetParentYangName() string { return "adj-rib-in-post" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetParentYangName() string { return "adj-rib-in-post" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost
// Adjacency rib out-bound post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetFilter() yfilter.YFilter { return adjRibOutPost.YFilter }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) SetFilter(yf yfilter.YFilter) { adjRibOutPost.YFilter = yf }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetSegmentPath() string {
    return "adj-rib-out-post"
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibOutPost.Routes
    }
    if childYangName == "num-routes" {
        return &adjRibOutPost.NumRoutes
    }
    return nil
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibOutPost.Routes
    children["num-routes"] = &adjRibOutPost.NumRoutes
    return children
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetBundleName() string { return "cisco_ios_xr" }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetYangName() string { return "adj-rib-out-post" }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) SetParent(parent types.Entity) { adjRibOutPost.parent = parent }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetParent() types.Entity { return adjRibOutPost.parent }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetParentYangName() string { return "open-config-neighbor" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetParentYangName() string { return "adj-rib-out-post" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetParentYangName() string { return "adj-rib-out-post" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre
// Adjacency rib out-bound pre-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetFilter() yfilter.YFilter { return adjRibOutPre.YFilter }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) SetFilter(yf yfilter.YFilter) { adjRibOutPre.YFilter = yf }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetSegmentPath() string {
    return "adj-rib-out-pre"
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibOutPre.Routes
    }
    if childYangName == "num-routes" {
        return &adjRibOutPre.NumRoutes
    }
    return nil
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibOutPre.Routes
    children["num-routes"] = &adjRibOutPre.NumRoutes
    return children
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetBundleName() string { return "cisco_ios_xr" }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetYangName() string { return "adj-rib-out-pre" }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) SetParent(parent types.Entity) { adjRibOutPre.parent = parent }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetParent() types.Entity { return adjRibOutPre.parent }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetParentYangName() string { return "open-config-neighbor" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetParentYangName() string { return "adj-rib-out-pre" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetParentYangName() string { return "adj-rib-out-pre" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre
// Adjacency rib in-bound pre-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetFilter() yfilter.YFilter { return adjRibInPre.YFilter }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) SetFilter(yf yfilter.YFilter) { adjRibInPre.YFilter = yf }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetSegmentPath() string {
    return "adj-rib-in-pre"
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibInPre.Routes
    }
    if childYangName == "num-routes" {
        return &adjRibInPre.NumRoutes
    }
    return nil
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibInPre.Routes
    children["num-routes"] = &adjRibInPre.NumRoutes
    return children
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetBundleName() string { return "cisco_ios_xr" }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetYangName() string { return "adj-rib-in-pre" }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) SetParent(parent types.Entity) { adjRibInPre.parent = parent }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetParent() types.Entity { return adjRibInPre.parent }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetParentYangName() string { return "open-config-neighbor" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetParentYangName() string { return "adj-rib-in-pre" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetParentYangName() string { return "adj-rib-in-pre" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast
// IPv6 Unicast
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local rib route table.
    LocRib OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib

    // Neighbor list.
    OpenConfigNeighbors OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors
}

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetFilter() yfilter.YFilter { return ipv6Unicast.YFilter }

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) SetFilter(yf yfilter.YFilter) { ipv6Unicast.YFilter = yf }

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetGoName(yname string) string {
    if yname == "loc-rib" { return "LocRib" }
    if yname == "open-config-neighbors" { return "OpenConfigNeighbors" }
    return ""
}

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetSegmentPath() string {
    return "ipv6-unicast"
}

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "loc-rib" {
        return &ipv6Unicast.LocRib
    }
    if childYangName == "open-config-neighbors" {
        return &ipv6Unicast.OpenConfigNeighbors
    }
    return nil
}

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["loc-rib"] = &ipv6Unicast.LocRib
    children["open-config-neighbors"] = &ipv6Unicast.OpenConfigNeighbors
    return children
}

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetYangName() string { return "ipv6-unicast" }

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) SetParent(parent types.Entity) { ipv6Unicast.parent = parent }

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetParent() types.Entity { return ipv6Unicast.parent }

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetParentYangName() string { return "afi-safi-table" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib
// Local rib route table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetFilter() yfilter.YFilter { return locRib.YFilter }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) SetFilter(yf yfilter.YFilter) { locRib.YFilter = yf }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetSegmentPath() string {
    return "loc-rib"
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &locRib.Routes
    }
    if childYangName == "num-routes" {
        return &locRib.NumRoutes
    }
    return nil
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &locRib.Routes
    children["num-routes"] = &locRib.NumRoutes
    return children
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetBundleName() string { return "cisco_ios_xr" }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetYangName() string { return "loc-rib" }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) SetParent(parent types.Entity) { locRib.parent = parent }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetParent() types.Entity { return locRib.parent }

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetParentYangName() string { return "ipv6-unicast" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetParentYangName() string { return "loc-rib" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetParentYangName() string { return "loc-rib" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors
// Neighbor list
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor name. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor.
    OpenConfigNeighbor []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetFilter() yfilter.YFilter { return openConfigNeighbors.YFilter }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) SetFilter(yf yfilter.YFilter) { openConfigNeighbors.YFilter = yf }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetGoName(yname string) string {
    if yname == "open-config-neighbor" { return "OpenConfigNeighbor" }
    return ""
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetSegmentPath() string {
    return "open-config-neighbors"
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "open-config-neighbor" {
        for _, c := range openConfigNeighbors.OpenConfigNeighbor {
            if openConfigNeighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor{}
        openConfigNeighbors.OpenConfigNeighbor = append(openConfigNeighbors.OpenConfigNeighbor, child)
        return &openConfigNeighbors.OpenConfigNeighbor[len(openConfigNeighbors.OpenConfigNeighbor)-1]
    }
    return nil
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range openConfigNeighbors.OpenConfigNeighbor {
        children[openConfigNeighbors.OpenConfigNeighbor[i].GetSegmentPath()] = &openConfigNeighbors.OpenConfigNeighbor[i]
    }
    return children
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetBundleName() string { return "cisco_ios_xr" }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetYangName() string { return "open-config-neighbors" }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) SetParent(parent types.Entity) { openConfigNeighbors.parent = parent }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetParent() types.Entity { return openConfigNeighbors.parent }

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetParentYangName() string { return "ipv6-unicast" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor
// Neighbor name
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Adjacency rib in-bound post-policy table.
    AdjRibInPost OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost

    // Adjacency rib out-bound post-policy table.
    AdjRibOutPost OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost

    // Adjacency rib out-bound pre-policy table.
    AdjRibOutPre OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre

    // Adjacency rib in-bound pre-policy table.
    AdjRibInPre OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetFilter() yfilter.YFilter { return openConfigNeighbor.YFilter }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) SetFilter(yf yfilter.YFilter) { openConfigNeighbor.YFilter = yf }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetGoName(yname string) string {
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "adj-rib-in-post" { return "AdjRibInPost" }
    if yname == "adj-rib-out-post" { return "AdjRibOutPost" }
    if yname == "adj-rib-out-pre" { return "AdjRibOutPre" }
    if yname == "adj-rib-in-pre" { return "AdjRibInPre" }
    return ""
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetSegmentPath() string {
    return "open-config-neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", openConfigNeighbor.NeighborAddress) + "']"
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "adj-rib-in-post" {
        return &openConfigNeighbor.AdjRibInPost
    }
    if childYangName == "adj-rib-out-post" {
        return &openConfigNeighbor.AdjRibOutPost
    }
    if childYangName == "adj-rib-out-pre" {
        return &openConfigNeighbor.AdjRibOutPre
    }
    if childYangName == "adj-rib-in-pre" {
        return &openConfigNeighbor.AdjRibInPre
    }
    return nil
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["adj-rib-in-post"] = &openConfigNeighbor.AdjRibInPost
    children["adj-rib-out-post"] = &openConfigNeighbor.AdjRibOutPost
    children["adj-rib-out-pre"] = &openConfigNeighbor.AdjRibOutPre
    children["adj-rib-in-pre"] = &openConfigNeighbor.AdjRibInPre
    return children
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-address"] = openConfigNeighbor.NeighborAddress
    return leafs
}

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetYangName() string { return "open-config-neighbor" }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) SetParent(parent types.Entity) { openConfigNeighbor.parent = parent }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetParent() types.Entity { return openConfigNeighbor.parent }

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetParentYangName() string { return "open-config-neighbors" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost
// Adjacency rib in-bound post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetFilter() yfilter.YFilter { return adjRibInPost.YFilter }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) SetFilter(yf yfilter.YFilter) { adjRibInPost.YFilter = yf }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetSegmentPath() string {
    return "adj-rib-in-post"
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibInPost.Routes
    }
    if childYangName == "num-routes" {
        return &adjRibInPost.NumRoutes
    }
    return nil
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibInPost.Routes
    children["num-routes"] = &adjRibInPost.NumRoutes
    return children
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetBundleName() string { return "cisco_ios_xr" }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetYangName() string { return "adj-rib-in-post" }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) SetParent(parent types.Entity) { adjRibInPost.parent = parent }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetParent() types.Entity { return adjRibInPost.parent }

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetParentYangName() string { return "open-config-neighbor" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetParentYangName() string { return "adj-rib-in-post" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetParentYangName() string { return "adj-rib-in-post" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost
// Adjacency rib out-bound post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetFilter() yfilter.YFilter { return adjRibOutPost.YFilter }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) SetFilter(yf yfilter.YFilter) { adjRibOutPost.YFilter = yf }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetSegmentPath() string {
    return "adj-rib-out-post"
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibOutPost.Routes
    }
    if childYangName == "num-routes" {
        return &adjRibOutPost.NumRoutes
    }
    return nil
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibOutPost.Routes
    children["num-routes"] = &adjRibOutPost.NumRoutes
    return children
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetBundleName() string { return "cisco_ios_xr" }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetYangName() string { return "adj-rib-out-post" }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) SetParent(parent types.Entity) { adjRibOutPost.parent = parent }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetParent() types.Entity { return adjRibOutPost.parent }

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetParentYangName() string { return "open-config-neighbor" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetParentYangName() string { return "adj-rib-out-post" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetParentYangName() string { return "adj-rib-out-post" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre
// Adjacency rib out-bound pre-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetFilter() yfilter.YFilter { return adjRibOutPre.YFilter }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) SetFilter(yf yfilter.YFilter) { adjRibOutPre.YFilter = yf }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetSegmentPath() string {
    return "adj-rib-out-pre"
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibOutPre.Routes
    }
    if childYangName == "num-routes" {
        return &adjRibOutPre.NumRoutes
    }
    return nil
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibOutPre.Routes
    children["num-routes"] = &adjRibOutPre.NumRoutes
    return children
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetBundleName() string { return "cisco_ios_xr" }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetYangName() string { return "adj-rib-out-pre" }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) SetParent(parent types.Entity) { adjRibOutPre.parent = parent }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetParent() types.Entity { return adjRibOutPre.parent }

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetParentYangName() string { return "open-config-neighbor" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetParentYangName() string { return "adj-rib-out-pre" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetParentYangName() string { return "adj-rib-out-pre" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre
// Adjacency rib in-bound pre-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetFilter() yfilter.YFilter { return adjRibInPre.YFilter }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) SetFilter(yf yfilter.YFilter) { adjRibInPre.YFilter = yf }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetSegmentPath() string {
    return "adj-rib-in-pre"
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibInPre.Routes
    }
    if childYangName == "num-routes" {
        return &adjRibInPre.NumRoutes
    }
    return nil
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibInPre.Routes
    children["num-routes"] = &adjRibInPre.NumRoutes
    return children
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetBundleName() string { return "cisco_ios_xr" }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetYangName() string { return "adj-rib-in-pre" }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) SetParent(parent types.Entity) { adjRibInPre.parent = parent }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetParent() types.Entity { return adjRibInPre.parent }

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetParentYangName() string { return "open-config-neighbor" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route.
    Route []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetYangName() string { return "routes" }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetParent() types.Entity { return routes.parent }

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetParentYangName() string { return "adj-rib-in-pre" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network in prefix/length format. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Route interface{}

    // Neighbor address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Path ID. The type is interface{} with range: -2147483648..2147483647.
    PathId interface{}

    // ValidRoute. The type is bool.
    ValidRoute interface{}

    // IndentityRef. The type is BgpOcInvalidRouteReason.
    InvalidReason interface{}

    // BestPath. The type is bool.
    BestPath interface{}

    // Prefix.
    PrefixName OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName

    // RouteAttributesList.
    RouteAttrList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList

    // ExtAttributesList.
    ExtAttributesList OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList

    // LastModifiedDate.
    LastModifiedDate OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate

    // LastUpdateRecieved.
    LastUpdateRecieved OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "path-id" { return "PathId" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "prefix-name" { return "PrefixName" }
    if yname == "route-attr-list" { return "RouteAttrList" }
    if yname == "ext-attributes-list" { return "ExtAttributesList" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-recieved" { return "LastUpdateRecieved" }
    return ""
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-name" {
        return &route.PrefixName
    }
    if childYangName == "route-attr-list" {
        return &route.RouteAttrList
    }
    if childYangName == "ext-attributes-list" {
        return &route.ExtAttributesList
    }
    if childYangName == "last-modified-date" {
        return &route.LastModifiedDate
    }
    if childYangName == "last-update-recieved" {
        return &route.LastUpdateRecieved
    }
    return nil
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-name"] = &route.PrefixName
    children["route-attr-list"] = &route.RouteAttrList
    children["ext-attributes-list"] = &route.ExtAttributesList
    children["last-modified-date"] = &route.LastModifiedDate
    children["last-update-recieved"] = &route.LastUpdateRecieved
    return children
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = route.Route
    leafs["neighbor-address"] = route.NeighborAddress
    leafs["path-id"] = route.PathId
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetYangName() string { return "route" }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetParentYangName() string { return "routes" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetFilter() yfilter.YFilter { return prefixName.YFilter }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) SetFilter(yf yfilter.YFilter) { prefixName.YFilter = yf }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetSegmentPath() string {
    return "prefix-name"
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefixName.Prefix
    }
    return nil
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefixName.Prefix
    return children
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefixName.PrefixLength
    return leafs
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetBundleName() string { return "cisco_ios_xr" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetYangName() string { return "prefix-name" }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) SetParent(parent types.Entity) { prefixName.parent = parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetParent() types.Entity { return prefixName.parent }

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = prefix.Afi
    leafs["ipv4-address"] = prefix.Ipv4Address
    leafs["ipv6-address"] = prefix.Ipv6Address
    return leafs
}

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetYangName() string { return "prefix" }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetParentYangName() string { return "prefix-name" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Origin Attribute Type. The type is BgpOcOriginAttr.
    OriginType interface{}

    // AS Path. The type is string.
    AsPath interface{}

    // AS4 Path. The type is string.
    As4Path interface{}

    // Med. The type is interface{} with range: 0..4294967295.
    Med interface{}

    // LocalPref. The type is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // AtomicAggr. The type is bool.
    AtomicAggr interface{}

    // NextHopAddress.
    NextHop OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop

    // AggregatorList.
    AggregratorAttributes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes

    // CommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community.
    Community []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetFilter() yfilter.YFilter { return routeAttrList.YFilter }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) SetFilter(yf yfilter.YFilter) { routeAttrList.YFilter = yf }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetGoName(yname string) string {
    if yname == "origin-type" { return "OriginType" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "aggregrator-attributes" { return "AggregratorAttributes" }
    if yname == "community" { return "Community" }
    return ""
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetSegmentPath() string {
    return "route-attr-list"
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &routeAttrList.NextHop
    }
    if childYangName == "aggregrator-attributes" {
        return &routeAttrList.AggregratorAttributes
    }
    if childYangName == "community" {
        for _, c := range routeAttrList.Community {
            if routeAttrList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community{}
        routeAttrList.Community = append(routeAttrList.Community, child)
        return &routeAttrList.Community[len(routeAttrList.Community)-1]
    }
    return nil
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &routeAttrList.NextHop
    children["aggregrator-attributes"] = &routeAttrList.AggregratorAttributes
    for i := range routeAttrList.Community {
        children[routeAttrList.Community[i].GetSegmentPath()] = &routeAttrList.Community[i]
    }
    return children
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-type"] = routeAttrList.OriginType
    leafs["as-path"] = routeAttrList.AsPath
    leafs["as4-path"] = routeAttrList.As4Path
    leafs["med"] = routeAttrList.Med
    leafs["local-pref"] = routeAttrList.LocalPref
    leafs["atomic-aggr"] = routeAttrList.AtomicAggr
    return leafs
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetBundleName() string { return "cisco_ios_xr" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetYangName() string { return "route-attr-list" }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) SetParent(parent types.Entity) { routeAttrList.parent = parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetParent() types.Entity { return routeAttrList.parent }

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BgpOcAfi.
    Afi interface{}

    // IPv4 Addr. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Addr. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = nextHop.Afi
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetFilter() yfilter.YFilter { return aggregratorAttributes.YFilter }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) SetFilter(yf yfilter.YFilter) { aggregratorAttributes.YFilter = yf }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetSegmentPath() string {
    return "aggregrator-attributes"
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregratorAttributes.As
    leafs["as4"] = aggregratorAttributes.As4
    leafs["address"] = aggregratorAttributes.Address
    return leafs
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetYangName() string { return "aggregrator-attributes" }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) SetParent(parent types.Entity) { aggregratorAttributes.parent = parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetParent() types.Entity { return aggregratorAttributes.parent }

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetFilter() yfilter.YFilter { return community.YFilter }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) SetFilter(yf yfilter.YFilter) { community.YFilter = yf }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetSegmentPath() string {
    return "community"
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = community.Objects
    return leafs
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetBundleName() string { return "cisco_ios_xr" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetYangName() string { return "community" }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) SetParent(parent types.Entity) { community.parent = parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetParent() types.Entity { return community.parent }

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetParentYangName() string { return "route-attr-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OriginatorID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // AIGP. The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // PathId. The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // ClusterList. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cluster []interface{}

    // ExtendedCommunityArray. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity.
    ExtCommunity []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetFilter() yfilter.YFilter { return extAttributesList.YFilter }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) SetFilter(yf yfilter.YFilter) { extAttributesList.YFilter = yf }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "cluster" { return "Cluster" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "unknown-attributes" { return "UnknownAttributes" }
    return ""
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetSegmentPath() string {
    return "ext-attributes-list"
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ext-community" {
        for _, c := range extAttributesList.ExtCommunity {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity{}
        extAttributesList.ExtCommunity = append(extAttributesList.ExtCommunity, child)
        return &extAttributesList.ExtCommunity[len(extAttributesList.ExtCommunity)-1]
    }
    if childYangName == "unknown-attributes" {
        for _, c := range extAttributesList.UnknownAttributes {
            if extAttributesList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes{}
        extAttributesList.UnknownAttributes = append(extAttributesList.UnknownAttributes, child)
        return &extAttributesList.UnknownAttributes[len(extAttributesList.UnknownAttributes)-1]
    }
    return nil
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributesList.ExtCommunity {
        children[extAttributesList.ExtCommunity[i].GetSegmentPath()] = &extAttributesList.ExtCommunity[i]
    }
    for i := range extAttributesList.UnknownAttributes {
        children[extAttributesList.UnknownAttributes[i].GetSegmentPath()] = &extAttributesList.UnknownAttributes[i]
    }
    return children
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributesList.OriginatorId
    leafs["aigp"] = extAttributesList.Aigp
    leafs["path-id"] = extAttributesList.PathId
    leafs["cluster"] = extAttributesList.Cluster
    return leafs
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetBundleName() string { return "cisco_ios_xr" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetYangName() string { return "ext-attributes-list" }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) SetParent(parent types.Entity) { extAttributesList.parent = parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetParent() types.Entity { return extAttributesList.parent }

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetFilter() yfilter.YFilter { return extCommunity.YFilter }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) SetFilter(yf yfilter.YFilter) { extCommunity.YFilter = yf }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetGoName(yname string) string {
    if yname == "objects" { return "Objects" }
    return ""
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetSegmentPath() string {
    return "ext-community"
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["objects"] = extCommunity.Objects
    return leafs
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetYangName() string { return "ext-community" }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) SetParent(parent types.Entity) { extCommunity.parent = parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetParent() types.Entity { return extCommunity.parent }

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetFilter() yfilter.YFilter { return unknownAttributes.YFilter }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) SetFilter(yf yfilter.YFilter) { unknownAttributes.YFilter = yf }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetGoName(yname string) string {
    if yname == "attribute-type" { return "AttributeType" }
    if yname == "attribute-length" { return "AttributeLength" }
    if yname == "attribute-value" { return "AttributeValue" }
    return ""
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetSegmentPath() string {
    return "unknown-attributes"
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attribute-type"] = unknownAttributes.AttributeType
    leafs["attribute-length"] = unknownAttributes.AttributeLength
    leafs["attribute-value"] = unknownAttributes.AttributeValue
    return leafs
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetYangName() string { return "unknown-attributes" }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) SetParent(parent types.Entity) { unknownAttributes.parent = parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetParent() types.Entity { return unknownAttributes.parent }

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetParentYangName() string { return "ext-attributes-list" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetFilter() yfilter.YFilter { return lastModifiedDate.YFilter }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) SetFilter(yf yfilter.YFilter) { lastModifiedDate.YFilter = yf }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetSegmentPath() string {
    return "last-modified-date"
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastModifiedDate.TimeValue
    return leafs
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetBundleName() string { return "cisco_ios_xr" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetYangName() string { return "last-modified-date" }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) SetParent(parent types.Entity) { lastModifiedDate.parent = parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetParent() types.Entity { return lastModifiedDate.parent }

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetFilter() yfilter.YFilter { return lastUpdateRecieved.YFilter }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) SetFilter(yf yfilter.YFilter) { lastUpdateRecieved.YFilter = yf }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetGoName(yname string) string {
    if yname == "time-value" { return "TimeValue" }
    return ""
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetSegmentPath() string {
    return "last-update-recieved"
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-value"] = lastUpdateRecieved.TimeValue
    return leafs
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetBundleName() string { return "cisco_ios_xr" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetYangName() string { return "last-update-recieved" }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) SetParent(parent types.Entity) { lastUpdateRecieved.parent = parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetParent() types.Entity { return lastUpdateRecieved.parent }

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetParentYangName() string { return "route" }

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetFilter() yfilter.YFilter { return numRoutes.YFilter }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) SetFilter(yf yfilter.YFilter) { numRoutes.YFilter = yf }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    return ""
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetSegmentPath() string {
    return "num-routes"
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = numRoutes.NumRoutes
    return leafs
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetYangName() string { return "num-routes" }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) SetParent(parent types.Entity) { numRoutes.parent = parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetParent() types.Entity { return numRoutes.parent }

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetParentYangName() string { return "adj-rib-in-pre" }

