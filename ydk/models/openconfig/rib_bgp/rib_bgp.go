// Defines a data model for representing BGP routing table (RIB)
// contents.  The model supports 5 logical RIBs per address family:
// 
// loc-rib: This is the main BGP routing table for the local routing
// instance, containing best-path selections for each prefix. The
// loc-rib table may contain multiple routes for a given prefix,
// with an attribute to indicate which was selected as the best
// path.
// 
// adj-rib-in-pre: This is a per-neighbor table containing the NLRI
// updates received from the neighbor before any local input policy
// rules or filters have been applied.  This can be considered the
// 'raw' updates from a given neighbor.
// 
// adj-rib-in-post: This is a per-neighbor table containing the
// routes received from the neighbor that are eligible for
// best-path selection after local input policy rules have been
// applied.
// 
// adj-rib-out-pre: This is a per-neighbor table containing routes
// eligible for sending (advertising) to the neighbor before output
// policy rules have been applied.
// 
// adj-rib-out-post: This is a per-neighbor table containing routes
// eligible for sending (advertising) to the neighbor after output
// policy rules have been applied.
package rib_bgp

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rib_bgp"))
    ydk.RegisterEntity("{http://openconfig.net/yang/rib/bgp bgp-rib}", reflect.TypeOf(BgpRib{}))
    ydk.RegisterEntity("openconfig-rib-bgp:bgp-rib", reflect.TypeOf(BgpRib{}))
}

// BgpRib
// Top level container for BGP RIBs
type BgpRib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enclosing container for address family list.
    AfiSafis BgpRib_AfiSafis
}

func (bgpRib *BgpRib) GetFilter() yfilter.YFilter { return bgpRib.YFilter }

func (bgpRib *BgpRib) SetFilter(yf yfilter.YFilter) { bgpRib.YFilter = yf }

func (bgpRib *BgpRib) GetGoName(yname string) string {
    if yname == "afi-safis" { return "AfiSafis" }
    return ""
}

func (bgpRib *BgpRib) GetSegmentPath() string {
    return "openconfig-rib-bgp:bgp-rib"
}

func (bgpRib *BgpRib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afi-safis" {
        return &bgpRib.AfiSafis
    }
    return nil
}

func (bgpRib *BgpRib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afi-safis"] = &bgpRib.AfiSafis
    return children
}

func (bgpRib *BgpRib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpRib *BgpRib) GetBundleName() string { return "openconfig" }

func (bgpRib *BgpRib) GetYangName() string { return "bgp-rib" }

func (bgpRib *BgpRib) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (bgpRib *BgpRib) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (bgpRib *BgpRib) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (bgpRib *BgpRib) SetParent(parent types.Entity) { bgpRib.parent = parent }

func (bgpRib *BgpRib) GetParent() types.Entity { return bgpRib.parent }

func (bgpRib *BgpRib) GetParentYangName() string { return "openconfig-rib-bgp" }

// BgpRib_AfiSafis
// Enclosing container for address family list
type BgpRib_AfiSafis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // list of afi-safi types. The type is slice of BgpRib_AfiSafis_AfiSafi.
    AfiSafi []BgpRib_AfiSafis_AfiSafi
}

func (afiSafis *BgpRib_AfiSafis) GetFilter() yfilter.YFilter { return afiSafis.YFilter }

func (afiSafis *BgpRib_AfiSafis) SetFilter(yf yfilter.YFilter) { afiSafis.YFilter = yf }

func (afiSafis *BgpRib_AfiSafis) GetGoName(yname string) string {
    if yname == "afi-safi" { return "AfiSafi" }
    return ""
}

func (afiSafis *BgpRib_AfiSafis) GetSegmentPath() string {
    return "afi-safis"
}

func (afiSafis *BgpRib_AfiSafis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afi-safi" {
        for _, c := range afiSafis.AfiSafi {
            if afiSafis.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi{}
        afiSafis.AfiSafi = append(afiSafis.AfiSafi, child)
        return &afiSafis.AfiSafi[len(afiSafis.AfiSafi)-1]
    }
    return nil
}

func (afiSafis *BgpRib_AfiSafis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afiSafis.AfiSafi {
        children[afiSafis.AfiSafi[i].GetSegmentPath()] = &afiSafis.AfiSafi[i]
    }
    return children
}

func (afiSafis *BgpRib_AfiSafis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afiSafis *BgpRib_AfiSafis) GetBundleName() string { return "openconfig" }

func (afiSafis *BgpRib_AfiSafis) GetYangName() string { return "afi-safis" }

func (afiSafis *BgpRib_AfiSafis) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (afiSafis *BgpRib_AfiSafis) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (afiSafis *BgpRib_AfiSafis) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (afiSafis *BgpRib_AfiSafis) SetParent(parent types.Entity) { afiSafis.parent = parent }

func (afiSafis *BgpRib_AfiSafis) GetParent() types.Entity { return afiSafis.parent }

func (afiSafis *BgpRib_AfiSafis) GetParentYangName() string { return "bgp-rib" }

// BgpRib_AfiSafis_AfiSafi
// list of afi-safi types
type BgpRib_AfiSafis_AfiSafi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. AFI,SAFI. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
    AfiSafiName interface{}

    // Routing tables for IPv4 unicast -- active when the afi-safi name is
    // ipv4-unicast.
    Ipv4Unicast BgpRib_AfiSafis_AfiSafi_Ipv4Unicast

    // Routing tables for IPv6 unicast -- active when the afi-safi name is
    // ipv6-unicast.
    Ipv6Unicast BgpRib_AfiSafis_AfiSafi_Ipv6Unicast
}

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetFilter() yfilter.YFilter { return afiSafi.YFilter }

func (afiSafi *BgpRib_AfiSafis_AfiSafi) SetFilter(yf yfilter.YFilter) { afiSafi.YFilter = yf }

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "ipv4-unicast" { return "Ipv4Unicast" }
    if yname == "ipv6-unicast" { return "Ipv6Unicast" }
    return ""
}

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetSegmentPath() string {
    return "afi-safi" + "[afi-safi-name='" + fmt.Sprintf("%v", afiSafi.AfiSafiName) + "']"
}

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-unicast" {
        return &afiSafi.Ipv4Unicast
    }
    if childYangName == "ipv6-unicast" {
        return &afiSafi.Ipv6Unicast
    }
    return nil
}

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-unicast"] = &afiSafi.Ipv4Unicast
    children["ipv6-unicast"] = &afiSafi.Ipv6Unicast
    return children
}

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = afiSafi.AfiSafiName
    return leafs
}

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetBundleName() string { return "openconfig" }

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetYangName() string { return "afi-safi" }

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (afiSafi *BgpRib_AfiSafis_AfiSafi) SetParent(parent types.Entity) { afiSafi.parent = parent }

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetParent() types.Entity { return afiSafi.parent }

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetParentYangName() string { return "afi-safis" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast
// Routing tables for IPv4 unicast -- active when the
// afi-safi name is ipv4-unicast
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Main routing table on the router, containing best-path selections for each
    // prefix.  The loc-rib may contain multiple routes for the same prefix (it is
    // a read-only, unkeyed list).  The best-path leaf should be set to true for
    // the route selected by the best-path selection process. Note that multiple
    // paths may be used or advertised even if only one path is marked as best,
    // e.g., when using BGP add-paths.  An implementation may choose to mark
    // multiple paths in the RIB as best path by setting the flag to true for
    // multiple entries.
    LocRib BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib

    // Enclosing container for neighbor list.
    Neighbors BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors
}

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetFilter() yfilter.YFilter { return ipv4Unicast.YFilter }

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) SetFilter(yf yfilter.YFilter) { ipv4Unicast.YFilter = yf }

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetGoName(yname string) string {
    if yname == "loc-rib" { return "LocRib" }
    if yname == "neighbors" { return "Neighbors" }
    return ""
}

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetSegmentPath() string {
    return "ipv4-unicast"
}

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "loc-rib" {
        return &ipv4Unicast.LocRib
    }
    if childYangName == "neighbors" {
        return &ipv4Unicast.Neighbors
    }
    return nil
}

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["loc-rib"] = &ipv4Unicast.LocRib
    children["neighbors"] = &ipv4Unicast.Neighbors
    return children
}

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetBundleName() string { return "openconfig" }

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetYangName() string { return "ipv4-unicast" }

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) SetParent(parent types.Entity) { ipv4Unicast.parent = parent }

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetParent() types.Entity { return ipv4Unicast.parent }

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetParentYangName() string { return "afi-safi" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib
// Main routing table on the router, containing best-path
// selections for each prefix.  The loc-rib may contain
// multiple routes for the same prefix (it is a read-only,
// unkeyed list).  The best-path leaf should be set to true
// for the route selected by the best-path selection process.
// Note that multiple paths may be used or advertised even if
// only one path is marked as best, e.g., when using BGP
// add-paths.  An implementation may choose to mark multiple
// paths in the RIB as best path by setting the flag to true for
// multiple entries.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetFilter() yfilter.YFilter { return locRib.YFilter }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) SetFilter(yf yfilter.YFilter) { locRib.YFilter = yf }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetSegmentPath() string {
    return "loc-rib"
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &locRib.Routes
    }
    return nil
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &locRib.Routes
    return children
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = locRib.NumRoutes
    return leafs
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetBundleName() string { return "openconfig" }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetYangName() string { return "loc-rib" }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) SetParent(parent types.Entity) { locRib.parent = parent }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetParent() types.Entity { return locRib.parent }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetParentYangName() string { return "ipv4-unicast" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetParentYangName() string { return "loc-rib" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors
// Enclosing container for neighbor list
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of neighbors (peers) of the local BGP speaker. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor.
    Neighbor []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetBundleName() string { return "openconfig" }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetParentYangName() string { return "ipv4-unicast" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor
// List of neighbors (peers) of the local BGP speaker
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the BGP neighbor or peer. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Per-neighbor table containing the NLRI updates received from the neighbor
    // before any local input policy rules or filters have been applied.  This can
    // be considered the 'raw' updates from the neighbor.
    AdjRibInPre BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre

    // Per-neighbor table containing the paths received from the neighbor that are
    // eligible for best-path selection after local input policy rules have been
    // applied.
    AdjRibInPost BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost

    // Per-neighbor table containing paths eligble for sending (advertising) to
    // the neighbor before output policy rules have been applied.
    AdjRibOutPre BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre

    // Per-neighbor table containing paths eligble for sending (advertising) to
    // the neighbor after output policy rules have been applied.
    AdjRibOutPost BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "adj-rib-in-pre" { return "AdjRibInPre" }
    if yname == "adj-rib-in-post" { return "AdjRibInPost" }
    if yname == "adj-rib-out-pre" { return "AdjRibOutPre" }
    if yname == "adj-rib-out-post" { return "AdjRibOutPost" }
    return ""
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", neighbor.NeighborAddress) + "']"
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "adj-rib-in-pre" {
        return &neighbor.AdjRibInPre
    }
    if childYangName == "adj-rib-in-post" {
        return &neighbor.AdjRibInPost
    }
    if childYangName == "adj-rib-out-pre" {
        return &neighbor.AdjRibOutPre
    }
    if childYangName == "adj-rib-out-post" {
        return &neighbor.AdjRibOutPost
    }
    return nil
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["adj-rib-in-pre"] = &neighbor.AdjRibInPre
    children["adj-rib-in-post"] = &neighbor.AdjRibInPost
    children["adj-rib-out-pre"] = &neighbor.AdjRibOutPre
    children["adj-rib-out-post"] = &neighbor.AdjRibOutPost
    return children
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-address"] = neighbor.NeighborAddress
    return leafs
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetBundleName() string { return "openconfig" }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre
// Per-neighbor table containing the NLRI updates
// received from the neighbor before any local input
// policy rules or filters have been applied.  This can
// be considered the 'raw' updates from the neighbor.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetFilter() yfilter.YFilter { return adjRibInPre.YFilter }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) SetFilter(yf yfilter.YFilter) { adjRibInPre.YFilter = yf }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetSegmentPath() string {
    return "adj-rib-in-pre"
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibInPre.Routes
    }
    return nil
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibInPre.Routes
    return children
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = adjRibInPre.NumRoutes
    return leafs
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetBundleName() string { return "openconfig" }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetYangName() string { return "adj-rib-in-pre" }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) SetParent(parent types.Entity) { adjRibInPre.parent = parent }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetParent() types.Entity { return adjRibInPre.parent }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetParentYangName() string { return "neighbor" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetParentYangName() string { return "adj-rib-in-pre" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost
// Per-neighbor table containing the paths received from
// the neighbor that are eligible for best-path selection
// after local input policy rules have been applied.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetFilter() yfilter.YFilter { return adjRibInPost.YFilter }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) SetFilter(yf yfilter.YFilter) { adjRibInPost.YFilter = yf }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetSegmentPath() string {
    return "adj-rib-in-post"
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibInPost.Routes
    }
    return nil
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibInPost.Routes
    return children
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = adjRibInPost.NumRoutes
    return leafs
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetBundleName() string { return "openconfig" }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetYangName() string { return "adj-rib-in-post" }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) SetParent(parent types.Entity) { adjRibInPost.parent = parent }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetParent() types.Entity { return adjRibInPost.parent }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetParentYangName() string { return "neighbor" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetParentYangName() string { return "adj-rib-in-post" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre
// Per-neighbor table containing paths eligble for
// sending (advertising) to the neighbor before output
// policy rules have been applied
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetFilter() yfilter.YFilter { return adjRibOutPre.YFilter }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) SetFilter(yf yfilter.YFilter) { adjRibOutPre.YFilter = yf }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetSegmentPath() string {
    return "adj-rib-out-pre"
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibOutPre.Routes
    }
    return nil
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibOutPre.Routes
    return children
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = adjRibOutPre.NumRoutes
    return leafs
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetBundleName() string { return "openconfig" }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetYangName() string { return "adj-rib-out-pre" }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) SetParent(parent types.Entity) { adjRibOutPre.parent = parent }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetParent() types.Entity { return adjRibOutPre.parent }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetParentYangName() string { return "neighbor" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetParentYangName() string { return "adj-rib-out-pre" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost
// Per-neighbor table containing paths eligble for
// sending (advertising) to the neighbor after output
// policy rules have been applied
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetFilter() yfilter.YFilter { return adjRibOutPost.YFilter }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) SetFilter(yf yfilter.YFilter) { adjRibOutPost.YFilter = yf }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetSegmentPath() string {
    return "adj-rib-out-post"
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibOutPost.Routes
    }
    return nil
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibOutPost.Routes
    return children
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = adjRibOutPost.NumRoutes
    return leafs
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetBundleName() string { return "openconfig" }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetYangName() string { return "adj-rib-out-post" }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) SetParent(parent types.Entity) { adjRibOutPost.parent = parent }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetParent() types.Entity { return adjRibOutPost.parent }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetParentYangName() string { return "neighbor" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetParentYangName() string { return "adj-rib-out-post" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast
// Routing tables for IPv6 unicast -- active when the
// afi-safi name is ipv6-unicast
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Main routing table on the router, containing best-path selections for each
    // prefix.  The loc-rib may contain multiple routes for the same prefix (it is
    // a read-only, unkeyed list).  The best-path leaf should be set to true for
    // the route selected by the best-path selection process. Note that multiple
    // paths may be used or advertised even if only one path is marked as best,
    // e.g., when using BGP add-paths.  An implementation may choose to mark
    // multiple paths in the RIB as best path by setting the flag to true for
    // multiple entries.
    LocRib BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib

    // Enclosing container for neighbor list.
    Neighbors BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors
}

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetFilter() yfilter.YFilter { return ipv6Unicast.YFilter }

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) SetFilter(yf yfilter.YFilter) { ipv6Unicast.YFilter = yf }

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetGoName(yname string) string {
    if yname == "loc-rib" { return "LocRib" }
    if yname == "neighbors" { return "Neighbors" }
    return ""
}

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetSegmentPath() string {
    return "ipv6-unicast"
}

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "loc-rib" {
        return &ipv6Unicast.LocRib
    }
    if childYangName == "neighbors" {
        return &ipv6Unicast.Neighbors
    }
    return nil
}

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["loc-rib"] = &ipv6Unicast.LocRib
    children["neighbors"] = &ipv6Unicast.Neighbors
    return children
}

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetBundleName() string { return "openconfig" }

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetYangName() string { return "ipv6-unicast" }

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) SetParent(parent types.Entity) { ipv6Unicast.parent = parent }

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetParent() types.Entity { return ipv6Unicast.parent }

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetParentYangName() string { return "afi-safi" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib
// Main routing table on the router, containing best-path
// selections for each prefix.  The loc-rib may contain
// multiple routes for the same prefix (it is a read-only,
// unkeyed list).  The best-path leaf should be set to true
// for the route selected by the best-path selection process.
// Note that multiple paths may be used or advertised even if
// only one path is marked as best, e.g., when using BGP
// add-paths.  An implementation may choose to mark multiple
// paths in the RIB as best path by setting the flag to true for
// multiple entries.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetFilter() yfilter.YFilter { return locRib.YFilter }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) SetFilter(yf yfilter.YFilter) { locRib.YFilter = yf }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetSegmentPath() string {
    return "loc-rib"
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &locRib.Routes
    }
    return nil
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &locRib.Routes
    return children
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = locRib.NumRoutes
    return leafs
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetBundleName() string { return "openconfig" }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetYangName() string { return "loc-rib" }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) SetParent(parent types.Entity) { locRib.parent = parent }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetParent() types.Entity { return locRib.parent }

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetParentYangName() string { return "ipv6-unicast" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetParentYangName() string { return "loc-rib" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors
// Enclosing container for neighbor list
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of neighbors (peers) of the local BGP speaker. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor.
    Neighbor []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetBundleName() string { return "openconfig" }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetParentYangName() string { return "ipv6-unicast" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor
// List of neighbors (peers) of the local BGP speaker
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the BGP neighbor or peer. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Per-neighbor table containing the NLRI updates received from the neighbor
    // before any local input policy rules or filters have been applied.  This can
    // be considered the 'raw' updates from the neighbor.
    AdjRibInPre BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre

    // Per-neighbor table containing the paths received from the neighbor that are
    // eligible for best-path selection after local input policy rules have been
    // applied.
    AdjRibInPost BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost

    // Per-neighbor table containing paths eligble for sending (advertising) to
    // the neighbor before output policy rules have been applied.
    AdjRibOutPre BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre

    // Per-neighbor table containing paths eligble for sending (advertising) to
    // the neighbor after output policy rules have been applied.
    AdjRibOutPost BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "adj-rib-in-pre" { return "AdjRibInPre" }
    if yname == "adj-rib-in-post" { return "AdjRibInPost" }
    if yname == "adj-rib-out-pre" { return "AdjRibOutPre" }
    if yname == "adj-rib-out-post" { return "AdjRibOutPost" }
    return ""
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", neighbor.NeighborAddress) + "']"
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "adj-rib-in-pre" {
        return &neighbor.AdjRibInPre
    }
    if childYangName == "adj-rib-in-post" {
        return &neighbor.AdjRibInPost
    }
    if childYangName == "adj-rib-out-pre" {
        return &neighbor.AdjRibOutPre
    }
    if childYangName == "adj-rib-out-post" {
        return &neighbor.AdjRibOutPost
    }
    return nil
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["adj-rib-in-pre"] = &neighbor.AdjRibInPre
    children["adj-rib-in-post"] = &neighbor.AdjRibInPost
    children["adj-rib-out-pre"] = &neighbor.AdjRibOutPre
    children["adj-rib-out-post"] = &neighbor.AdjRibOutPost
    return children
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-address"] = neighbor.NeighborAddress
    return leafs
}

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetBundleName() string { return "openconfig" }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre
// Per-neighbor table containing the NLRI updates
// received from the neighbor before any local input
// policy rules or filters have been applied.  This can
// be considered the 'raw' updates from the neighbor.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetFilter() yfilter.YFilter { return adjRibInPre.YFilter }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) SetFilter(yf yfilter.YFilter) { adjRibInPre.YFilter = yf }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetSegmentPath() string {
    return "adj-rib-in-pre"
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibInPre.Routes
    }
    return nil
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibInPre.Routes
    return children
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = adjRibInPre.NumRoutes
    return leafs
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetBundleName() string { return "openconfig" }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetYangName() string { return "adj-rib-in-pre" }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) SetParent(parent types.Entity) { adjRibInPre.parent = parent }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetParent() types.Entity { return adjRibInPre.parent }

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetParentYangName() string { return "neighbor" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetParentYangName() string { return "adj-rib-in-pre" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost
// Per-neighbor table containing the paths received from
// the neighbor that are eligible for best-path selection
// after local input policy rules have been applied.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetFilter() yfilter.YFilter { return adjRibInPost.YFilter }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) SetFilter(yf yfilter.YFilter) { adjRibInPost.YFilter = yf }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetSegmentPath() string {
    return "adj-rib-in-post"
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibInPost.Routes
    }
    return nil
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibInPost.Routes
    return children
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = adjRibInPost.NumRoutes
    return leafs
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetBundleName() string { return "openconfig" }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetYangName() string { return "adj-rib-in-post" }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) SetParent(parent types.Entity) { adjRibInPost.parent = parent }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetParent() types.Entity { return adjRibInPost.parent }

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetParentYangName() string { return "neighbor" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetParentYangName() string { return "adj-rib-in-post" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre
// Per-neighbor table containing paths eligble for
// sending (advertising) to the neighbor before output
// policy rules have been applied
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetFilter() yfilter.YFilter { return adjRibOutPre.YFilter }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) SetFilter(yf yfilter.YFilter) { adjRibOutPre.YFilter = yf }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetSegmentPath() string {
    return "adj-rib-out-pre"
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibOutPre.Routes
    }
    return nil
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibOutPre.Routes
    return children
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = adjRibOutPre.NumRoutes
    return leafs
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetBundleName() string { return "openconfig" }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetYangName() string { return "adj-rib-out-pre" }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) SetParent(parent types.Entity) { adjRibOutPre.parent = parent }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetParent() types.Entity { return adjRibOutPre.parent }

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetParentYangName() string { return "neighbor" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetParentYangName() string { return "adj-rib-out-pre" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost
// Per-neighbor table containing paths eligble for
// sending (advertising) to the neighbor after output
// policy rules have been applied
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetFilter() yfilter.YFilter { return adjRibOutPost.YFilter }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) SetFilter(yf yfilter.YFilter) { adjRibOutPost.YFilter = yf }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetGoName(yname string) string {
    if yname == "num-routes" { return "NumRoutes" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetSegmentPath() string {
    return "adj-rib-out-post"
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &adjRibOutPost.Routes
    }
    return nil
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &adjRibOutPost.Routes
    return children
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-routes"] = adjRibOutPost.NumRoutes
    return leafs
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetBundleName() string { return "openconfig" }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetYangName() string { return "adj-rib-out-post" }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) SetParent(parent types.Entity) { adjRibOutPost.parent = parent }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetParent() types.Entity { return adjRibOutPost.parent }

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetParentYangName() string { return "neighbor" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route.
    Route []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetBundleName() string { return "openconfig" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetYangName() string { return "routes" }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetParent() types.Entity { return routes.parent }

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetParentYangName() string { return "adj-rib-out-post" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix for the route. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // Timestamp of when this path was last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastModifiedDate interface{}

    // Timestamp of when the last BGP update message was received for this path /
    // prefix. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastUpdateReceived interface{}

    // Indicates that the route is considered valid by the local router. The type
    // is bool.
    ValidRoute interface{}

    // If the route is rejected as invalid, this indicates the reason. The type is
    // one of the following:
    // INVALIDORIGINATORINVALIDCLUSTERLOOPINVALIDASLOOPINVALIDCONFED.
    InvalidReason interface{}

    // Current path was selected as the best path. The type is bool.
    BestPath interface{}

    // Base BGP route attributes associated with this route.
    Attributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes

    // Extended BGP route attributes associated with this route.
    ExtAttributes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "last-modified-date" { return "LastModifiedDate" }
    if yname == "last-update-received" { return "LastUpdateReceived" }
    if yname == "valid-route" { return "ValidRoute" }
    if yname == "invalid-reason" { return "InvalidReason" }
    if yname == "best-path" { return "BestPath" }
    if yname == "attributes" { return "Attributes" }
    if yname == "ext-attributes" { return "ExtAttributes" }
    return ""
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attributes" {
        return &route.Attributes
    }
    if childYangName == "ext-attributes" {
        return &route.ExtAttributes
    }
    return nil
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attributes"] = &route.Attributes
    children["ext-attributes"] = &route.ExtAttributes
    return children
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["last-modified-date"] = route.LastModifiedDate
    leafs["last-update-received"] = route.LastUpdateReceived
    leafs["valid-route"] = route.ValidRoute
    leafs["invalid-reason"] = route.InvalidReason
    leafs["best-path"] = route.BestPath
    return leafs
}

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetBundleName() string { return "openconfig" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetYangName() string { return "route" }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetParentYangName() string { return "routes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute defining the origin of the path information. The type is
    // BgpOriginAttrType.
    Origin interface{}

    // String representation of the BGP AS path attribute as concatenated AS path
    // segments.  Each segment of the AS_PATH should be formatted as follows based
    // on the segment type (#### denotes a single AS number):   AS_SEQ: #### ####
    // #####  AS_SET: { #### #### }  AS_CONFED_SEQUENCE: ( #### #### ) 
    // AS_CONFED_SET: [ #### #### ]  AS_PATH segment types are described in RFC
    // 5065.  In the Adj-RIB-In or Adj-RIB-Out, this leaf should show the received
    // or sent AS_PATH value, respectively.  For example, if the local router is
    // not 4-byte capable, this value should consist of 2-octet ASNs or the
    // AS_TRANS (AS 23456) values received or sent in route updates.  In the
    // Loc-RIB, this leaf should reflect the effective AS path for the route,
    // e.g., a 4-octet value if the local router is 4-octet capable. The type is
    // string.
    AsPath interface{}

    // This string represents the AS path encoded with 4-octet AS numbers in the
    // optional transitive AS4_PATH attribute. This value is populated with the
    // received or sent attribute in Adj-RIB-In or Adj-RIB-Out, respectively.  It
    // should not be populated in Loc-RIB since the Loc-RIB is expected to store
    // the effective AS-Path in the as-path leaf regardless of being 4-octet or
    // 2-octet. The type is string.
    As4Path interface{}

    // BGP next hop attribute defining the IP address of the router that should be
    // used as the next hop to the destination. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // BGP multi-exit discriminator attribute used in BGP route selection process.
    // The type is interface{} with range: 0..4294967295.
    Med interface{}

    // BGP local preference attribute sent to internal peers to indicate. The type
    // is interface{} with range: 0..4294967295.
    LocalPref interface{}

    // BGP attribute indicating that the prefix is an atomic aggregate, i.e., the
    // peer selected a less specific route without selecting a more specific route
    // that is included in it. The type is bool.
    AtomicAggr interface{}

    // List of standard BGP community attributes. The type is one of the following
    // types: slice of int with range: 65536..4294901759, or slice of string with
    // pattern: ([0-9]+:[0-9]+).
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetFilter() yfilter.YFilter { return attributes.YFilter }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) SetFilter(yf yfilter.YFilter) { attributes.YFilter = yf }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetGoName(yname string) string {
    if yname == "origin" { return "Origin" }
    if yname == "as-path" { return "AsPath" }
    if yname == "as4-path" { return "As4Path" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "med" { return "Med" }
    if yname == "local-pref" { return "LocalPref" }
    if yname == "atomic-aggr" { return "AtomicAggr" }
    if yname == "community" { return "Community" }
    if yname == "aggregator" { return "Aggregator" }
    return ""
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetSegmentPath() string {
    return "attributes"
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregator" {
        return &attributes.Aggregator
    }
    return nil
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregator"] = &attributes.Aggregator
    return children
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin"] = attributes.Origin
    leafs["as-path"] = attributes.AsPath
    leafs["as4-path"] = attributes.As4Path
    leafs["next-hop"] = attributes.NextHop
    leafs["med"] = attributes.Med
    leafs["local-pref"] = attributes.LocalPref
    leafs["atomic-aggr"] = attributes.AtomicAggr
    leafs["community"] = attributes.Community
    return leafs
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetBundleName() string { return "openconfig" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetYangName() string { return "attributes" }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) SetParent(parent types.Entity) { attributes.parent = parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetParent() types.Entity { return attributes.parent }

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AS number of the autnonomous system that performed the aggregation. The
    // type is interface{} with range: 0..4294967295.
    As interface{}

    // AS number of the autnonomous system that performed the aggregation (4-octet
    // representation).  This value is populated if an upstream router is not
    // 4-octet capable. Its semantics are similar to the AS4_PATH optional
    // transitive attribute. The type is interface{} with range: 0..4294967295.
    As4 interface{}

    // IP address of the router that performed the aggregation. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetFilter() yfilter.YFilter { return aggregator.YFilter }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) SetFilter(yf yfilter.YFilter) { aggregator.YFilter = yf }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "as4" { return "As4" }
    if yname == "address" { return "Address" }
    return ""
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetSegmentPath() string {
    return "aggregator"
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = aggregator.As
    leafs["as4"] = aggregator.As4
    leafs["address"] = aggregator.Address
    return leafs
}

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetBundleName() string { return "openconfig" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetYangName() string { return "aggregator" }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) SetParent(parent types.Entity) { aggregator.parent = parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetParent() types.Entity { return aggregator.parent }

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetParentYangName() string { return "attributes" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP attribute that provides the id as an IPv4 address of the route
    // reflector that created the announcement. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OriginatorId interface{}

    // Represents the reflection path that the route has passed. The type is slice
    // of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ClusterList []interface{}

    // List of BGP extended community attributes. The type is one of the following
    // types: slice of string with pattern:
    // (6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-target:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]),
    // or slice of string with pattern:
    // route\-origin:(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]):(4[0-2][0-9][0-4][0-9][0-6][0-7][0-2][0-9][0-6]|[1-3][0-9]{9}|[1-9]([0-9]{1,7})?[0-9]|[1-9]),
    // or slice of string with pattern:
    // route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6[0-5][0-5][0-3][0-5]|[1-5][0-9]{4}|[1-9][0-9]{1,4}|[0-9]).
    ExtCommunity []interface{}

    // BGP path attribute representing the accumulated IGP metric for the path.
    // The type is interface{} with range: 0..18446744073709551615.
    Aigp interface{}

    // When the BGP speaker supports advertisement of multiple paths for a prefix,
    // the path identifier is used to uniquely identify a route based on the
    // combination of the prefix and path id.  In the Adj-RIB-In, the path-id
    // value is the value received in the update message.   In the Loc-RIB, if
    // used, it should represent a locally generated path-id value for the
    // corresponding route.  In Adj-RIB-Out, it should be the value sent to a
    // neighbor when add-paths is used, i.e., the capability has been negotiated.
    // The type is interface{} with range: 0..4294967295.
    PathId interface{}

    // This list contains received attributes that are unrecognized or unsupported
    // by the local router.  The list may be empty. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute.
    UnknownAttribute []BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetFilter() yfilter.YFilter { return extAttributes.YFilter }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) SetFilter(yf yfilter.YFilter) { extAttributes.YFilter = yf }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetGoName(yname string) string {
    if yname == "originator-id" { return "OriginatorId" }
    if yname == "cluster-list" { return "ClusterList" }
    if yname == "ext-community" { return "ExtCommunity" }
    if yname == "aigp" { return "Aigp" }
    if yname == "path-id" { return "PathId" }
    if yname == "unknown-attribute" { return "UnknownAttribute" }
    return ""
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetSegmentPath() string {
    return "ext-attributes"
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-attribute" {
        for _, c := range extAttributes.UnknownAttribute {
            if extAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute{}
        extAttributes.UnknownAttribute = append(extAttributes.UnknownAttribute, child)
        return &extAttributes.UnknownAttribute[len(extAttributes.UnknownAttribute)-1]
    }
    return nil
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range extAttributes.UnknownAttribute {
        children[extAttributes.UnknownAttribute[i].GetSegmentPath()] = &extAttributes.UnknownAttribute[i]
    }
    return children
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["originator-id"] = extAttributes.OriginatorId
    leafs["cluster-list"] = extAttributes.ClusterList
    leafs["ext-community"] = extAttributes.ExtCommunity
    leafs["aigp"] = extAttributes.Aigp
    leafs["path-id"] = extAttributes.PathId
    return leafs
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetBundleName() string { return "openconfig" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetYangName() string { return "ext-attributes" }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) SetParent(parent types.Entity) { extAttributes.parent = parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetParent() types.Entity { return extAttributes.parent }

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetParentYangName() string { return "route" }

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. 2-octet value encoding the attribute flags and the
    // attribute type code. The type is interface{} with range: 0..65535.
    AttrType interface{}

    // One or two octet attribute length field indicating the length of the
    // attribute data in octets.  If the Extended Length attribute flag in the
    // attribute type field is set, the length field is 2 octets, otherwise it is
    // 1 octet. The type is interface{} with range: 0..65535.
    AttrLen interface{}

    // Raw attribute value data, not to exceed the length indicated in the
    // attr-len field.  The maximum length of the attribute data is 2^16-1 per the
    // max value of the attr-len field (2 octets). The type is string with length:
    // 1..65535.
    AttrValue interface{}
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetFilter() yfilter.YFilter { return unknownAttribute.YFilter }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) SetFilter(yf yfilter.YFilter) { unknownAttribute.YFilter = yf }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetGoName(yname string) string {
    if yname == "attr-type" { return "AttrType" }
    if yname == "attr-len" { return "AttrLen" }
    if yname == "attr-value" { return "AttrValue" }
    return ""
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetSegmentPath() string {
    return "unknown-attribute" + "[attr-type='" + fmt.Sprintf("%v", unknownAttribute.AttrType) + "']"
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attr-type"] = unknownAttribute.AttrType
    leafs["attr-len"] = unknownAttribute.AttrLen
    leafs["attr-value"] = unknownAttribute.AttrValue
    return leafs
}

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleName() string { return "openconfig" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetYangName() string { return "unknown-attribute" }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) SetParent(parent types.Entity) { unknownAttribute.parent = parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetParent() types.Entity { return unknownAttribute.parent }

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetParentYangName() string { return "ext-attributes" }

