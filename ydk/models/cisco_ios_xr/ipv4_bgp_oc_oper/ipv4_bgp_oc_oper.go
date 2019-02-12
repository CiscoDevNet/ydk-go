// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-bgp-oc package operational data.
// 
// This module contains definitions
// for the following management objects:
//   oc-bgp: OC-BGP operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP-RIB operational data.
    BgpRib OcBgp_BgpRib
}

func (ocBgp *OcBgp) GetEntityData() *types.CommonEntityData {
    ocBgp.EntityData.YFilter = ocBgp.YFilter
    ocBgp.EntityData.YangName = "oc-bgp"
    ocBgp.EntityData.BundleName = "cisco_ios_xr"
    ocBgp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-bgp-oc-oper"
    ocBgp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp"
    ocBgp.EntityData.AbsolutePath = ocBgp.EntityData.SegmentPath
    ocBgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ocBgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ocBgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ocBgp.EntityData.Children = types.NewOrderedMap()
    ocBgp.EntityData.Children.Append("bgp-rib", types.YChild{"BgpRib", &ocBgp.BgpRib})
    ocBgp.EntityData.Leafs = types.NewOrderedMap()

    ocBgp.EntityData.YListKeys = []string {}

    return &(ocBgp.EntityData)
}

// OcBgp_BgpRib
// BGP-RIB operational data
type OcBgp_BgpRib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI-SAFIs information.
    AfiSafiTable OcBgp_BgpRib_AfiSafiTable
}

func (bgpRib *OcBgp_BgpRib) GetEntityData() *types.CommonEntityData {
    bgpRib.EntityData.YFilter = bgpRib.YFilter
    bgpRib.EntityData.YangName = "bgp-rib"
    bgpRib.EntityData.BundleName = "cisco_ios_xr"
    bgpRib.EntityData.ParentYangName = "oc-bgp"
    bgpRib.EntityData.SegmentPath = "bgp-rib"
    bgpRib.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/" + bgpRib.EntityData.SegmentPath
    bgpRib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpRib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpRib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpRib.EntityData.Children = types.NewOrderedMap()
    bgpRib.EntityData.Children.Append("afi-safi-table", types.YChild{"AfiSafiTable", &bgpRib.AfiSafiTable})
    bgpRib.EntityData.Leafs = types.NewOrderedMap()

    bgpRib.EntityData.YListKeys = []string {}

    return &(bgpRib.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable
// AFI-SAFIs information
type OcBgp_BgpRib_AfiSafiTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Unicast.
    Ipv4Unicast OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast

    // IPv6 Unicast.
    Ipv6Unicast OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast
}

func (afiSafiTable *OcBgp_BgpRib_AfiSafiTable) GetEntityData() *types.CommonEntityData {
    afiSafiTable.EntityData.YFilter = afiSafiTable.YFilter
    afiSafiTable.EntityData.YangName = "afi-safi-table"
    afiSafiTable.EntityData.BundleName = "cisco_ios_xr"
    afiSafiTable.EntityData.ParentYangName = "bgp-rib"
    afiSafiTable.EntityData.SegmentPath = "afi-safi-table"
    afiSafiTable.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/" + afiSafiTable.EntityData.SegmentPath
    afiSafiTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afiSafiTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afiSafiTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afiSafiTable.EntityData.Children = types.NewOrderedMap()
    afiSafiTable.EntityData.Children.Append("ipv4-unicast", types.YChild{"Ipv4Unicast", &afiSafiTable.Ipv4Unicast})
    afiSafiTable.EntityData.Children.Append("ipv6-unicast", types.YChild{"Ipv6Unicast", &afiSafiTable.Ipv6Unicast})
    afiSafiTable.EntityData.Leafs = types.NewOrderedMap()

    afiSafiTable.EntityData.YListKeys = []string {}

    return &(afiSafiTable.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast
// IPv4 Unicast
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local rib route table.
    LocRib OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib

    // Neighbor list.
    OpenConfigNeighbors OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors
}

func (ipv4Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast) GetEntityData() *types.CommonEntityData {
    ipv4Unicast.EntityData.YFilter = ipv4Unicast.YFilter
    ipv4Unicast.EntityData.YangName = "ipv4-unicast"
    ipv4Unicast.EntityData.BundleName = "cisco_ios_xr"
    ipv4Unicast.EntityData.ParentYangName = "afi-safi-table"
    ipv4Unicast.EntityData.SegmentPath = "ipv4-unicast"
    ipv4Unicast.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/" + ipv4Unicast.EntityData.SegmentPath
    ipv4Unicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Unicast.EntityData.Children = types.NewOrderedMap()
    ipv4Unicast.EntityData.Children.Append("loc-rib", types.YChild{"LocRib", &ipv4Unicast.LocRib})
    ipv4Unicast.EntityData.Children.Append("open-config-neighbors", types.YChild{"OpenConfigNeighbors", &ipv4Unicast.OpenConfigNeighbors})
    ipv4Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv4Unicast.EntityData.YListKeys = []string {}

    return &(ipv4Unicast.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib
// Local rib route table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib) GetEntityData() *types.CommonEntityData {
    locRib.EntityData.YFilter = locRib.YFilter
    locRib.EntityData.YangName = "loc-rib"
    locRib.EntityData.BundleName = "cisco_ios_xr"
    locRib.EntityData.ParentYangName = "ipv4-unicast"
    locRib.EntityData.SegmentPath = "loc-rib"
    locRib.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/" + locRib.EntityData.SegmentPath
    locRib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locRib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locRib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locRib.EntityData.Children = types.NewOrderedMap()
    locRib.EntityData.Children.Append("routes", types.YChild{"Routes", &locRib.Routes})
    locRib.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &locRib.NumRoutes})
    locRib.EntityData.Leafs = types.NewOrderedMap()

    locRib.EntityData.YListKeys = []string {}

    return &(locRib.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "loc-rib"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_LocRib_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "loc-rib"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/loc-rib/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors
// Neighbor list
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor name. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor.
    OpenConfigNeighbor []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors) GetEntityData() *types.CommonEntityData {
    openConfigNeighbors.EntityData.YFilter = openConfigNeighbors.YFilter
    openConfigNeighbors.EntityData.YangName = "open-config-neighbors"
    openConfigNeighbors.EntityData.BundleName = "cisco_ios_xr"
    openConfigNeighbors.EntityData.ParentYangName = "ipv4-unicast"
    openConfigNeighbors.EntityData.SegmentPath = "open-config-neighbors"
    openConfigNeighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/" + openConfigNeighbors.EntityData.SegmentPath
    openConfigNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    openConfigNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    openConfigNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    openConfigNeighbors.EntityData.Children = types.NewOrderedMap()
    openConfigNeighbors.EntityData.Children.Append("open-config-neighbor", types.YChild{"OpenConfigNeighbor", nil})
    for i := range openConfigNeighbors.OpenConfigNeighbor {
        openConfigNeighbors.EntityData.Children.Append(types.GetSegmentPath(openConfigNeighbors.OpenConfigNeighbor[i]), types.YChild{"OpenConfigNeighbor", openConfigNeighbors.OpenConfigNeighbor[i]})
    }
    openConfigNeighbors.EntityData.Leafs = types.NewOrderedMap()

    openConfigNeighbors.EntityData.YListKeys = []string {}

    return &(openConfigNeighbors.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor
// Neighbor name
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetEntityData() *types.CommonEntityData {
    openConfigNeighbor.EntityData.YFilter = openConfigNeighbor.YFilter
    openConfigNeighbor.EntityData.YangName = "open-config-neighbor"
    openConfigNeighbor.EntityData.BundleName = "cisco_ios_xr"
    openConfigNeighbor.EntityData.ParentYangName = "open-config-neighbors"
    openConfigNeighbor.EntityData.SegmentPath = "open-config-neighbor" + types.AddKeyToken(openConfigNeighbor.NeighborAddress, "neighbor-address")
    openConfigNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/" + openConfigNeighbor.EntityData.SegmentPath
    openConfigNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    openConfigNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    openConfigNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    openConfigNeighbor.EntityData.Children = types.NewOrderedMap()
    openConfigNeighbor.EntityData.Children.Append("adj-rib-in-post", types.YChild{"AdjRibInPost", &openConfigNeighbor.AdjRibInPost})
    openConfigNeighbor.EntityData.Children.Append("adj-rib-out-post", types.YChild{"AdjRibOutPost", &openConfigNeighbor.AdjRibOutPost})
    openConfigNeighbor.EntityData.Children.Append("adj-rib-out-pre", types.YChild{"AdjRibOutPre", &openConfigNeighbor.AdjRibOutPre})
    openConfigNeighbor.EntityData.Children.Append("adj-rib-in-pre", types.YChild{"AdjRibInPre", &openConfigNeighbor.AdjRibInPre})
    openConfigNeighbor.EntityData.Leafs = types.NewOrderedMap()
    openConfigNeighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", openConfigNeighbor.NeighborAddress})

    openConfigNeighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(openConfigNeighbor.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost
// Adjacency rib in-bound post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetEntityData() *types.CommonEntityData {
    adjRibInPost.EntityData.YFilter = adjRibInPost.YFilter
    adjRibInPost.EntityData.YangName = "adj-rib-in-post"
    adjRibInPost.EntityData.BundleName = "cisco_ios_xr"
    adjRibInPost.EntityData.ParentYangName = "open-config-neighbor"
    adjRibInPost.EntityData.SegmentPath = "adj-rib-in-post"
    adjRibInPost.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/" + adjRibInPost.EntityData.SegmentPath
    adjRibInPost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjRibInPost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjRibInPost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjRibInPost.EntityData.Children = types.NewOrderedMap()
    adjRibInPost.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibInPost.Routes})
    adjRibInPost.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &adjRibInPost.NumRoutes})
    adjRibInPost.EntityData.Leafs = types.NewOrderedMap()

    adjRibInPost.EntityData.YListKeys = []string {}

    return &(adjRibInPost.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "adj-rib-in-post"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "adj-rib-in-post"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost
// Adjacency rib out-bound post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetEntityData() *types.CommonEntityData {
    adjRibOutPost.EntityData.YFilter = adjRibOutPost.YFilter
    adjRibOutPost.EntityData.YangName = "adj-rib-out-post"
    adjRibOutPost.EntityData.BundleName = "cisco_ios_xr"
    adjRibOutPost.EntityData.ParentYangName = "open-config-neighbor"
    adjRibOutPost.EntityData.SegmentPath = "adj-rib-out-post"
    adjRibOutPost.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/" + adjRibOutPost.EntityData.SegmentPath
    adjRibOutPost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjRibOutPost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjRibOutPost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjRibOutPost.EntityData.Children = types.NewOrderedMap()
    adjRibOutPost.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibOutPost.Routes})
    adjRibOutPost.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &adjRibOutPost.NumRoutes})
    adjRibOutPost.EntityData.Leafs = types.NewOrderedMap()

    adjRibOutPost.EntityData.YListKeys = []string {}

    return &(adjRibOutPost.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "adj-rib-out-post"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "adj-rib-out-post"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre
// Adjacency rib out-bound pre-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetEntityData() *types.CommonEntityData {
    adjRibOutPre.EntityData.YFilter = adjRibOutPre.YFilter
    adjRibOutPre.EntityData.YangName = "adj-rib-out-pre"
    adjRibOutPre.EntityData.BundleName = "cisco_ios_xr"
    adjRibOutPre.EntityData.ParentYangName = "open-config-neighbor"
    adjRibOutPre.EntityData.SegmentPath = "adj-rib-out-pre"
    adjRibOutPre.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/" + adjRibOutPre.EntityData.SegmentPath
    adjRibOutPre.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjRibOutPre.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjRibOutPre.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjRibOutPre.EntityData.Children = types.NewOrderedMap()
    adjRibOutPre.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibOutPre.Routes})
    adjRibOutPre.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &adjRibOutPre.NumRoutes})
    adjRibOutPre.EntityData.Leafs = types.NewOrderedMap()

    adjRibOutPre.EntityData.YListKeys = []string {}

    return &(adjRibOutPre.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "adj-rib-out-pre"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "adj-rib-out-pre"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre
// Adjacency rib in-bound pre-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetEntityData() *types.CommonEntityData {
    adjRibInPre.EntityData.YFilter = adjRibInPre.YFilter
    adjRibInPre.EntityData.YangName = "adj-rib-in-pre"
    adjRibInPre.EntityData.BundleName = "cisco_ios_xr"
    adjRibInPre.EntityData.ParentYangName = "open-config-neighbor"
    adjRibInPre.EntityData.SegmentPath = "adj-rib-in-pre"
    adjRibInPre.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/" + adjRibInPre.EntityData.SegmentPath
    adjRibInPre.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjRibInPre.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjRibInPre.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjRibInPre.EntityData.Children = types.NewOrderedMap()
    adjRibInPre.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibInPre.Routes})
    adjRibInPre.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &adjRibInPre.NumRoutes})
    adjRibInPre.EntityData.Leafs = types.NewOrderedMap()

    adjRibInPre.EntityData.YListKeys = []string {}

    return &(adjRibInPre.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "adj-rib-in-pre"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv4Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "adj-rib-in-pre"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv4-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast
// IPv6 Unicast
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local rib route table.
    LocRib OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib

    // Neighbor list.
    OpenConfigNeighbors OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors
}

func (ipv6Unicast *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast) GetEntityData() *types.CommonEntityData {
    ipv6Unicast.EntityData.YFilter = ipv6Unicast.YFilter
    ipv6Unicast.EntityData.YangName = "ipv6-unicast"
    ipv6Unicast.EntityData.BundleName = "cisco_ios_xr"
    ipv6Unicast.EntityData.ParentYangName = "afi-safi-table"
    ipv6Unicast.EntityData.SegmentPath = "ipv6-unicast"
    ipv6Unicast.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/" + ipv6Unicast.EntityData.SegmentPath
    ipv6Unicast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Unicast.EntityData.Children = types.NewOrderedMap()
    ipv6Unicast.EntityData.Children.Append("loc-rib", types.YChild{"LocRib", &ipv6Unicast.LocRib})
    ipv6Unicast.EntityData.Children.Append("open-config-neighbors", types.YChild{"OpenConfigNeighbors", &ipv6Unicast.OpenConfigNeighbors})
    ipv6Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv6Unicast.EntityData.YListKeys = []string {}

    return &(ipv6Unicast.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib
// Local rib route table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes
}

func (locRib *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib) GetEntityData() *types.CommonEntityData {
    locRib.EntityData.YFilter = locRib.YFilter
    locRib.EntityData.YangName = "loc-rib"
    locRib.EntityData.BundleName = "cisco_ios_xr"
    locRib.EntityData.ParentYangName = "ipv6-unicast"
    locRib.EntityData.SegmentPath = "loc-rib"
    locRib.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/" + locRib.EntityData.SegmentPath
    locRib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locRib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locRib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locRib.EntityData.Children = types.NewOrderedMap()
    locRib.EntityData.Children.Append("routes", types.YChild{"Routes", &locRib.Routes})
    locRib.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &locRib.NumRoutes})
    locRib.EntityData.Leafs = types.NewOrderedMap()

    locRib.EntityData.YListKeys = []string {}

    return &(locRib.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "loc-rib"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_LocRib_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "loc-rib"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/loc-rib/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors
// Neighbor list
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor name. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor.
    OpenConfigNeighbor []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor
}

func (openConfigNeighbors *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors) GetEntityData() *types.CommonEntityData {
    openConfigNeighbors.EntityData.YFilter = openConfigNeighbors.YFilter
    openConfigNeighbors.EntityData.YangName = "open-config-neighbors"
    openConfigNeighbors.EntityData.BundleName = "cisco_ios_xr"
    openConfigNeighbors.EntityData.ParentYangName = "ipv6-unicast"
    openConfigNeighbors.EntityData.SegmentPath = "open-config-neighbors"
    openConfigNeighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/" + openConfigNeighbors.EntityData.SegmentPath
    openConfigNeighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    openConfigNeighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    openConfigNeighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    openConfigNeighbors.EntityData.Children = types.NewOrderedMap()
    openConfigNeighbors.EntityData.Children.Append("open-config-neighbor", types.YChild{"OpenConfigNeighbor", nil})
    for i := range openConfigNeighbors.OpenConfigNeighbor {
        openConfigNeighbors.EntityData.Children.Append(types.GetSegmentPath(openConfigNeighbors.OpenConfigNeighbor[i]), types.YChild{"OpenConfigNeighbor", openConfigNeighbors.OpenConfigNeighbor[i]})
    }
    openConfigNeighbors.EntityData.Leafs = types.NewOrderedMap()

    openConfigNeighbors.EntityData.YListKeys = []string {}

    return &(openConfigNeighbors.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor
// Neighbor name
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (openConfigNeighbor *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor) GetEntityData() *types.CommonEntityData {
    openConfigNeighbor.EntityData.YFilter = openConfigNeighbor.YFilter
    openConfigNeighbor.EntityData.YangName = "open-config-neighbor"
    openConfigNeighbor.EntityData.BundleName = "cisco_ios_xr"
    openConfigNeighbor.EntityData.ParentYangName = "open-config-neighbors"
    openConfigNeighbor.EntityData.SegmentPath = "open-config-neighbor" + types.AddKeyToken(openConfigNeighbor.NeighborAddress, "neighbor-address")
    openConfigNeighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/" + openConfigNeighbor.EntityData.SegmentPath
    openConfigNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    openConfigNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    openConfigNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    openConfigNeighbor.EntityData.Children = types.NewOrderedMap()
    openConfigNeighbor.EntityData.Children.Append("adj-rib-in-post", types.YChild{"AdjRibInPost", &openConfigNeighbor.AdjRibInPost})
    openConfigNeighbor.EntityData.Children.Append("adj-rib-out-post", types.YChild{"AdjRibOutPost", &openConfigNeighbor.AdjRibOutPost})
    openConfigNeighbor.EntityData.Children.Append("adj-rib-out-pre", types.YChild{"AdjRibOutPre", &openConfigNeighbor.AdjRibOutPre})
    openConfigNeighbor.EntityData.Children.Append("adj-rib-in-pre", types.YChild{"AdjRibInPre", &openConfigNeighbor.AdjRibInPre})
    openConfigNeighbor.EntityData.Leafs = types.NewOrderedMap()
    openConfigNeighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", openConfigNeighbor.NeighborAddress})

    openConfigNeighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(openConfigNeighbor.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost
// Adjacency rib in-bound post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes
}

func (adjRibInPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost) GetEntityData() *types.CommonEntityData {
    adjRibInPost.EntityData.YFilter = adjRibInPost.YFilter
    adjRibInPost.EntityData.YangName = "adj-rib-in-post"
    adjRibInPost.EntityData.BundleName = "cisco_ios_xr"
    adjRibInPost.EntityData.ParentYangName = "open-config-neighbor"
    adjRibInPost.EntityData.SegmentPath = "adj-rib-in-post"
    adjRibInPost.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/" + adjRibInPost.EntityData.SegmentPath
    adjRibInPost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjRibInPost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjRibInPost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjRibInPost.EntityData.Children = types.NewOrderedMap()
    adjRibInPost.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibInPost.Routes})
    adjRibInPost.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &adjRibInPost.NumRoutes})
    adjRibInPost.EntityData.Leafs = types.NewOrderedMap()

    adjRibInPost.EntityData.YListKeys = []string {}

    return &(adjRibInPost.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "adj-rib-in-post"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPost_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "adj-rib-in-post"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-post/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost
// Adjacency rib out-bound post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes
}

func (adjRibOutPost *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost) GetEntityData() *types.CommonEntityData {
    adjRibOutPost.EntityData.YFilter = adjRibOutPost.YFilter
    adjRibOutPost.EntityData.YangName = "adj-rib-out-post"
    adjRibOutPost.EntityData.BundleName = "cisco_ios_xr"
    adjRibOutPost.EntityData.ParentYangName = "open-config-neighbor"
    adjRibOutPost.EntityData.SegmentPath = "adj-rib-out-post"
    adjRibOutPost.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/" + adjRibOutPost.EntityData.SegmentPath
    adjRibOutPost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjRibOutPost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjRibOutPost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjRibOutPost.EntityData.Children = types.NewOrderedMap()
    adjRibOutPost.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibOutPost.Routes})
    adjRibOutPost.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &adjRibOutPost.NumRoutes})
    adjRibOutPost.EntityData.Leafs = types.NewOrderedMap()

    adjRibOutPost.EntityData.YListKeys = []string {}

    return &(adjRibOutPost.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "adj-rib-out-post"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPost_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "adj-rib-out-post"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-post/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre
// Adjacency rib out-bound pre-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes
}

func (adjRibOutPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre) GetEntityData() *types.CommonEntityData {
    adjRibOutPre.EntityData.YFilter = adjRibOutPre.YFilter
    adjRibOutPre.EntityData.YangName = "adj-rib-out-pre"
    adjRibOutPre.EntityData.BundleName = "cisco_ios_xr"
    adjRibOutPre.EntityData.ParentYangName = "open-config-neighbor"
    adjRibOutPre.EntityData.SegmentPath = "adj-rib-out-pre"
    adjRibOutPre.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/" + adjRibOutPre.EntityData.SegmentPath
    adjRibOutPre.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjRibOutPre.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjRibOutPre.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjRibOutPre.EntityData.Children = types.NewOrderedMap()
    adjRibOutPre.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibOutPre.Routes})
    adjRibOutPre.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &adjRibOutPre.NumRoutes})
    adjRibOutPre.EntityData.Leafs = types.NewOrderedMap()

    adjRibOutPre.EntityData.YListKeys = []string {}

    return &(adjRibOutPre.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "adj-rib-out-pre"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibOutPre_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "adj-rib-out-pre"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-out-pre/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre
// Adjacency rib in-bound pre-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // routes table.
    Routes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes

    // Number of routes in adjacency rib out-bound post-policy table.
    NumRoutes OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes
}

func (adjRibInPre *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre) GetEntityData() *types.CommonEntityData {
    adjRibInPre.EntityData.YFilter = adjRibInPre.YFilter
    adjRibInPre.EntityData.YangName = "adj-rib-in-pre"
    adjRibInPre.EntityData.BundleName = "cisco_ios_xr"
    adjRibInPre.EntityData.ParentYangName = "open-config-neighbor"
    adjRibInPre.EntityData.SegmentPath = "adj-rib-in-pre"
    adjRibInPre.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/" + adjRibInPre.EntityData.SegmentPath
    adjRibInPre.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjRibInPre.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjRibInPre.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjRibInPre.EntityData.Children = types.NewOrderedMap()
    adjRibInPre.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibInPre.Routes})
    adjRibInPre.EntityData.Children.Append("num-routes", types.YChild{"NumRoutes", &adjRibInPre.NumRoutes})
    adjRibInPre.EntityData.Leafs = types.NewOrderedMap()

    adjRibInPre.EntityData.YListKeys = []string {}

    return &(adjRibInPre.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes
// routes table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route entry. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route.
    Route []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route
}

func (routes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "cisco_ios_xr"
    routes.EntityData.ParentYangName = "adj-rib-in-pre"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/" + routes.EntityData.SegmentPath
    routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        types.SetYListKey(routes.Route[i], i)
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route
// route entry
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Path ID. The type is interface{} with range: 0..4294967295.
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

func (route *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xr"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route" + types.AddNoKeyToken(route)
    route.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/" + route.EntityData.SegmentPath
    route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("prefix-name", types.YChild{"PrefixName", &route.PrefixName})
    route.EntityData.Children.Append("route-attr-list", types.YChild{"RouteAttrList", &route.RouteAttrList})
    route.EntityData.Children.Append("ext-attributes-list", types.YChild{"ExtAttributesList", &route.ExtAttributesList})
    route.EntityData.Children.Append("last-modified-date", types.YChild{"LastModifiedDate", &route.LastModifiedDate})
    route.EntityData.Children.Append("last-update-recieved", types.YChild{"LastUpdateRecieved", &route.LastUpdateRecieved})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("route", types.YLeaf{"Route", route.Route})
    route.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", route.NeighborAddress})
    route.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", route.PathId})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix
}

func (prefixName *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName) GetEntityData() *types.CommonEntityData {
    prefixName.EntityData.YFilter = prefixName.YFilter
    prefixName.EntityData.YangName = "prefix-name"
    prefixName.EntityData.BundleName = "cisco_ios_xr"
    prefixName.EntityData.ParentYangName = "route"
    prefixName.EntityData.SegmentPath = "prefix-name"
    prefixName.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + prefixName.EntityData.SegmentPath
    prefixName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixName.EntityData.Children = types.NewOrderedMap()
    prefixName.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefixName.Prefix})
    prefixName.EntityData.Leafs = types.NewOrderedMap()
    prefixName.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixName.PrefixLength})

    prefixName.EntityData.YListKeys = []string {}

    return &(prefixName.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix
// Prefix
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_PrefixName_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix-name"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/prefix-name/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", prefix.Afi})
    prefix.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", prefix.Ipv4Address})
    prefix.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", prefix.Ipv6Address})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList
// RouteAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList struct {
    EntityData types.CommonEntityData
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
    Community []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community
}

func (routeAttrList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList) GetEntityData() *types.CommonEntityData {
    routeAttrList.EntityData.YFilter = routeAttrList.YFilter
    routeAttrList.EntityData.YangName = "route-attr-list"
    routeAttrList.EntityData.BundleName = "cisco_ios_xr"
    routeAttrList.EntityData.ParentYangName = "route"
    routeAttrList.EntityData.SegmentPath = "route-attr-list"
    routeAttrList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + routeAttrList.EntityData.SegmentPath
    routeAttrList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeAttrList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeAttrList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeAttrList.EntityData.Children = types.NewOrderedMap()
    routeAttrList.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &routeAttrList.NextHop})
    routeAttrList.EntityData.Children.Append("aggregrator-attributes", types.YChild{"AggregratorAttributes", &routeAttrList.AggregratorAttributes})
    routeAttrList.EntityData.Children.Append("community", types.YChild{"Community", nil})
    for i := range routeAttrList.Community {
        types.SetYListKey(routeAttrList.Community[i], i)
        routeAttrList.EntityData.Children.Append(types.GetSegmentPath(routeAttrList.Community[i]), types.YChild{"Community", routeAttrList.Community[i]})
    }
    routeAttrList.EntityData.Leafs = types.NewOrderedMap()
    routeAttrList.EntityData.Leafs.Append("origin-type", types.YLeaf{"OriginType", routeAttrList.OriginType})
    routeAttrList.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", routeAttrList.AsPath})
    routeAttrList.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", routeAttrList.As4Path})
    routeAttrList.EntityData.Leafs.Append("med", types.YLeaf{"Med", routeAttrList.Med})
    routeAttrList.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", routeAttrList.LocalPref})
    routeAttrList.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", routeAttrList.AtomicAggr})

    routeAttrList.EntityData.YListKeys = []string {}

    return &(routeAttrList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop
// NextHopAddress
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "route-attr-list"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/route-attr-list/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nextHop.Afi})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes
// AggregatorList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AS number. The type is interface{} with range: 0..4294967295.
    As interface{}

    // AS4 number. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregratorAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_AggregratorAttributes) GetEntityData() *types.CommonEntityData {
    aggregratorAttributes.EntityData.YFilter = aggregratorAttributes.YFilter
    aggregratorAttributes.EntityData.YangName = "aggregrator-attributes"
    aggregratorAttributes.EntityData.BundleName = "cisco_ios_xr"
    aggregratorAttributes.EntityData.ParentYangName = "route-attr-list"
    aggregratorAttributes.EntityData.SegmentPath = "aggregrator-attributes"
    aggregratorAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/route-attr-list/" + aggregratorAttributes.EntityData.SegmentPath
    aggregratorAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregratorAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregratorAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregratorAttributes.EntityData.Children = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs = types.NewOrderedMap()
    aggregratorAttributes.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregratorAttributes.As})
    aggregratorAttributes.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregratorAttributes.As4})
    aggregratorAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregratorAttributes.Address})

    aggregratorAttributes.EntityData.YListKeys = []string {}

    return &(aggregratorAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community
// CommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (community *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_RouteAttrList_Community) GetEntityData() *types.CommonEntityData {
    community.EntityData.YFilter = community.YFilter
    community.EntityData.YangName = "community"
    community.EntityData.BundleName = "cisco_ios_xr"
    community.EntityData.ParentYangName = "route-attr-list"
    community.EntityData.SegmentPath = "community" + types.AddNoKeyToken(community)
    community.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/route-attr-list/" + community.EntityData.SegmentPath
    community.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    community.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    community.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    community.EntityData.Children = types.NewOrderedMap()
    community.EntityData.Leafs = types.NewOrderedMap()
    community.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", community.Objects})

    community.EntityData.YListKeys = []string {}

    return &(community.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList
// ExtAttributesList
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList struct {
    EntityData types.CommonEntityData
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
    ExtCommunity []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity

    // UnknownAttributes. The type is slice of
    // OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes.
    UnknownAttributes []*OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes
}

func (extAttributesList *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList) GetEntityData() *types.CommonEntityData {
    extAttributesList.EntityData.YFilter = extAttributesList.YFilter
    extAttributesList.EntityData.YangName = "ext-attributes-list"
    extAttributesList.EntityData.BundleName = "cisco_ios_xr"
    extAttributesList.EntityData.ParentYangName = "route"
    extAttributesList.EntityData.SegmentPath = "ext-attributes-list"
    extAttributesList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + extAttributesList.EntityData.SegmentPath
    extAttributesList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extAttributesList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extAttributesList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extAttributesList.EntityData.Children = types.NewOrderedMap()
    extAttributesList.EntityData.Children.Append("ext-community", types.YChild{"ExtCommunity", nil})
    for i := range extAttributesList.ExtCommunity {
        types.SetYListKey(extAttributesList.ExtCommunity[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.ExtCommunity[i]), types.YChild{"ExtCommunity", extAttributesList.ExtCommunity[i]})
    }
    extAttributesList.EntityData.Children.Append("unknown-attributes", types.YChild{"UnknownAttributes", nil})
    for i := range extAttributesList.UnknownAttributes {
        types.SetYListKey(extAttributesList.UnknownAttributes[i], i)
        extAttributesList.EntityData.Children.Append(types.GetSegmentPath(extAttributesList.UnknownAttributes[i]), types.YChild{"UnknownAttributes", extAttributesList.UnknownAttributes[i]})
    }
    extAttributesList.EntityData.Leafs = types.NewOrderedMap()
    extAttributesList.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributesList.OriginatorId})
    extAttributesList.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributesList.Aigp})
    extAttributesList.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributesList.PathId})
    extAttributesList.EntityData.Leafs.Append("cluster", types.YLeaf{"Cluster", extAttributesList.Cluster})

    extAttributesList.EntityData.YListKeys = []string {}

    return &(extAttributesList.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity
// ExtendedCommunityArray
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // BGP OC objects. The type is string.
    Objects interface{}
}

func (extCommunity *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_ExtCommunity) GetEntityData() *types.CommonEntityData {
    extCommunity.EntityData.YFilter = extCommunity.YFilter
    extCommunity.EntityData.YangName = "ext-community"
    extCommunity.EntityData.BundleName = "cisco_ios_xr"
    extCommunity.EntityData.ParentYangName = "ext-attributes-list"
    extCommunity.EntityData.SegmentPath = "ext-community" + types.AddNoKeyToken(extCommunity)
    extCommunity.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/ext-attributes-list/" + extCommunity.EntityData.SegmentPath
    extCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extCommunity.EntityData.Children = types.NewOrderedMap()
    extCommunity.EntityData.Leafs = types.NewOrderedMap()
    extCommunity.EntityData.Leafs.Append("objects", types.YLeaf{"Objects", extCommunity.Objects})

    extCommunity.EntityData.YListKeys = []string {}

    return &(extCommunity.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes
// UnknownAttributes
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // AttributeType. The type is interface{} with range: 0..65535.
    AttributeType interface{}

    // AttributeLength. The type is interface{} with range: 0..65535.
    AttributeLength interface{}

    // Atributevalue. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AttributeValue interface{}
}

func (unknownAttributes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_ExtAttributesList_UnknownAttributes) GetEntityData() *types.CommonEntityData {
    unknownAttributes.EntityData.YFilter = unknownAttributes.YFilter
    unknownAttributes.EntityData.YangName = "unknown-attributes"
    unknownAttributes.EntityData.BundleName = "cisco_ios_xr"
    unknownAttributes.EntityData.ParentYangName = "ext-attributes-list"
    unknownAttributes.EntityData.SegmentPath = "unknown-attributes" + types.AddNoKeyToken(unknownAttributes)
    unknownAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/ext-attributes-list/" + unknownAttributes.EntityData.SegmentPath
    unknownAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownAttributes.EntityData.Children = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs = types.NewOrderedMap()
    unknownAttributes.EntityData.Leafs.Append("attribute-type", types.YLeaf{"AttributeType", unknownAttributes.AttributeType})
    unknownAttributes.EntityData.Leafs.Append("attribute-length", types.YLeaf{"AttributeLength", unknownAttributes.AttributeLength})
    unknownAttributes.EntityData.Leafs.Append("attribute-value", types.YLeaf{"AttributeValue", unknownAttributes.AttributeValue})

    unknownAttributes.EntityData.YListKeys = []string {}

    return &(unknownAttributes.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate
// LastModifiedDate
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastModifiedDate *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastModifiedDate) GetEntityData() *types.CommonEntityData {
    lastModifiedDate.EntityData.YFilter = lastModifiedDate.YFilter
    lastModifiedDate.EntityData.YangName = "last-modified-date"
    lastModifiedDate.EntityData.BundleName = "cisco_ios_xr"
    lastModifiedDate.EntityData.ParentYangName = "route"
    lastModifiedDate.EntityData.SegmentPath = "last-modified-date"
    lastModifiedDate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + lastModifiedDate.EntityData.SegmentPath
    lastModifiedDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastModifiedDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastModifiedDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastModifiedDate.EntityData.Children = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs = types.NewOrderedMap()
    lastModifiedDate.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastModifiedDate.TimeValue})

    lastModifiedDate.EntityData.YListKeys = []string {}

    return &(lastModifiedDate.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved
// LastUpdateRecieved
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeValue. The type is string.
    TimeValue interface{}
}

func (lastUpdateRecieved *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_Routes_Route_LastUpdateRecieved) GetEntityData() *types.CommonEntityData {
    lastUpdateRecieved.EntityData.YFilter = lastUpdateRecieved.YFilter
    lastUpdateRecieved.EntityData.YangName = "last-update-recieved"
    lastUpdateRecieved.EntityData.BundleName = "cisco_ios_xr"
    lastUpdateRecieved.EntityData.ParentYangName = "route"
    lastUpdateRecieved.EntityData.SegmentPath = "last-update-recieved"
    lastUpdateRecieved.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/routes/route/" + lastUpdateRecieved.EntityData.SegmentPath
    lastUpdateRecieved.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastUpdateRecieved.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastUpdateRecieved.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastUpdateRecieved.EntityData.Children = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs = types.NewOrderedMap()
    lastUpdateRecieved.EntityData.Leafs.Append("time-value", types.YLeaf{"TimeValue", lastUpdateRecieved.TimeValue})

    lastUpdateRecieved.EntityData.YListKeys = []string {}

    return &(lastUpdateRecieved.EntityData)
}

// OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes
// Number of routes in adjacency rib out-bound
// post-policy table
type OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NumRoutes. The type is interface{} with range: 0..18446744073709551615.
    NumRoutes interface{}
}

func (numRoutes *OcBgp_BgpRib_AfiSafiTable_Ipv6Unicast_OpenConfigNeighbors_OpenConfigNeighbor_AdjRibInPre_NumRoutes) GetEntityData() *types.CommonEntityData {
    numRoutes.EntityData.YFilter = numRoutes.YFilter
    numRoutes.EntityData.YangName = "num-routes"
    numRoutes.EntityData.BundleName = "cisco_ios_xr"
    numRoutes.EntityData.ParentYangName = "adj-rib-in-pre"
    numRoutes.EntityData.SegmentPath = "num-routes"
    numRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-bgp-oc-oper:oc-bgp/bgp-rib/afi-safi-table/ipv6-unicast/open-config-neighbors/open-config-neighbor/adj-rib-in-pre/" + numRoutes.EntityData.SegmentPath
    numRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numRoutes.EntityData.Children = types.NewOrderedMap()
    numRoutes.EntityData.Leafs = types.NewOrderedMap()
    numRoutes.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", numRoutes.NumRoutes})

    numRoutes.EntityData.YListKeys = []string {}

    return &(numRoutes.EntityData)
}

