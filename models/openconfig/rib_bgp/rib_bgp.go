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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for address family list.
    AfiSafis BgpRib_AfiSafis
}

func (bgpRib *BgpRib) GetEntityData() *types.CommonEntityData {
    bgpRib.EntityData.YFilter = bgpRib.YFilter
    bgpRib.EntityData.YangName = "bgp-rib"
    bgpRib.EntityData.BundleName = "openconfig"
    bgpRib.EntityData.ParentYangName = "openconfig-rib-bgp"
    bgpRib.EntityData.SegmentPath = "openconfig-rib-bgp:bgp-rib"
    bgpRib.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bgpRib.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bgpRib.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bgpRib.EntityData.Children = types.NewOrderedMap()
    bgpRib.EntityData.Children.Append("afi-safis", types.YChild{"AfiSafis", &bgpRib.AfiSafis})
    bgpRib.EntityData.Leafs = types.NewOrderedMap()

    bgpRib.EntityData.YListKeys = []string {}

    return &(bgpRib.EntityData)
}

// BgpRib_AfiSafis
// Enclosing container for address family list
type BgpRib_AfiSafis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of afi-safi types. The type is slice of BgpRib_AfiSafis_AfiSafi.
    AfiSafi []*BgpRib_AfiSafis_AfiSafi
}

func (afiSafis *BgpRib_AfiSafis) GetEntityData() *types.CommonEntityData {
    afiSafis.EntityData.YFilter = afiSafis.YFilter
    afiSafis.EntityData.YangName = "afi-safis"
    afiSafis.EntityData.BundleName = "openconfig"
    afiSafis.EntityData.ParentYangName = "bgp-rib"
    afiSafis.EntityData.SegmentPath = "afi-safis"
    afiSafis.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafis.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafis.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafis.EntityData.Children = types.NewOrderedMap()
    afiSafis.EntityData.Children.Append("afi-safi", types.YChild{"AfiSafi", nil})
    for i := range afiSafis.AfiSafi {
        afiSafis.EntityData.Children.Append(types.GetSegmentPath(afiSafis.AfiSafi[i]), types.YChild{"AfiSafi", afiSafis.AfiSafi[i]})
    }
    afiSafis.EntityData.Leafs = types.NewOrderedMap()

    afiSafis.EntityData.YListKeys = []string {}

    return &(afiSafis.EntityData)
}

// BgpRib_AfiSafis_AfiSafi
// list of afi-safi types
type BgpRib_AfiSafis_AfiSafi struct {
    EntityData types.CommonEntityData
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

func (afiSafi *BgpRib_AfiSafis_AfiSafi) GetEntityData() *types.CommonEntityData {
    afiSafi.EntityData.YFilter = afiSafi.YFilter
    afiSafi.EntityData.YangName = "afi-safi"
    afiSafi.EntityData.BundleName = "openconfig"
    afiSafi.EntityData.ParentYangName = "afi-safis"
    afiSafi.EntityData.SegmentPath = "afi-safi" + types.AddKeyToken(afiSafi.AfiSafiName, "afi-safi-name")
    afiSafi.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafi.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafi.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafi.EntityData.Children = types.NewOrderedMap()
    afiSafi.EntityData.Children.Append("ipv4-unicast", types.YChild{"Ipv4Unicast", &afiSafi.Ipv4Unicast})
    afiSafi.EntityData.Children.Append("ipv6-unicast", types.YChild{"Ipv6Unicast", &afiSafi.Ipv6Unicast})
    afiSafi.EntityData.Leafs = types.NewOrderedMap()
    afiSafi.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", afiSafi.AfiSafiName})

    afiSafi.EntityData.YListKeys = []string {"AfiSafiName"}

    return &(afiSafi.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast
// Routing tables for IPv4 unicast -- active when the
// afi-safi name is ipv4-unicast
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast struct {
    EntityData types.CommonEntityData
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

func (ipv4Unicast *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast) GetEntityData() *types.CommonEntityData {
    ipv4Unicast.EntityData.YFilter = ipv4Unicast.YFilter
    ipv4Unicast.EntityData.YangName = "ipv4-unicast"
    ipv4Unicast.EntityData.BundleName = "openconfig"
    ipv4Unicast.EntityData.ParentYangName = "afi-safi"
    ipv4Unicast.EntityData.SegmentPath = "ipv4-unicast"
    ipv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4Unicast.EntityData.Children = types.NewOrderedMap()
    ipv4Unicast.EntityData.Children.Append("loc-rib", types.YChild{"LocRib", &ipv4Unicast.LocRib})
    ipv4Unicast.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &ipv4Unicast.Neighbors})
    ipv4Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv4Unicast.EntityData.YListKeys = []string {}

    return &(ipv4Unicast.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib) GetEntityData() *types.CommonEntityData {
    locRib.EntityData.YFilter = locRib.YFilter
    locRib.EntityData.YangName = "loc-rib"
    locRib.EntityData.BundleName = "openconfig"
    locRib.EntityData.ParentYangName = "ipv4-unicast"
    locRib.EntityData.SegmentPath = "loc-rib"
    locRib.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    locRib.EntityData.NamespaceTable = openconfig.GetNamespaces()
    locRib.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    locRib.EntityData.Children = types.NewOrderedMap()
    locRib.EntityData.Children.Append("routes", types.YChild{"Routes", &locRib.Routes})
    locRib.EntityData.Leafs = types.NewOrderedMap()
    locRib.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", locRib.NumRoutes})

    locRib.EntityData.YListKeys = []string {}

    return &(locRib.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "loc-rib"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors
// Enclosing container for neighbor list
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of neighbors (peers) of the local BGP speaker. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor.
    Neighbor []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "ipv4-unicast"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbors.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor
// List of neighbors (peers) of the local BGP speaker
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
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

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborAddress, "neighbor-address")
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("adj-rib-in-pre", types.YChild{"AdjRibInPre", &neighbor.AdjRibInPre})
    neighbor.EntityData.Children.Append("adj-rib-in-post", types.YChild{"AdjRibInPost", &neighbor.AdjRibInPost})
    neighbor.EntityData.Children.Append("adj-rib-out-pre", types.YChild{"AdjRibOutPre", &neighbor.AdjRibOutPre})
    neighbor.EntityData.Children.Append("adj-rib-out-post", types.YChild{"AdjRibOutPost", &neighbor.AdjRibOutPost})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighbor.NeighborAddress})

    neighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(neighbor.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre
// Per-neighbor table containing the NLRI updates
// received from the neighbor before any local input
// policy rules or filters have been applied.  This can
// be considered the 'raw' updates from the neighbor.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre) GetEntityData() *types.CommonEntityData {
    adjRibInPre.EntityData.YFilter = adjRibInPre.YFilter
    adjRibInPre.EntityData.YangName = "adj-rib-in-pre"
    adjRibInPre.EntityData.BundleName = "openconfig"
    adjRibInPre.EntityData.ParentYangName = "neighbor"
    adjRibInPre.EntityData.SegmentPath = "adj-rib-in-pre"
    adjRibInPre.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adjRibInPre.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adjRibInPre.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adjRibInPre.EntityData.Children = types.NewOrderedMap()
    adjRibInPre.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibInPre.Routes})
    adjRibInPre.EntityData.Leafs = types.NewOrderedMap()
    adjRibInPre.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", adjRibInPre.NumRoutes})

    adjRibInPre.EntityData.YListKeys = []string {}

    return &(adjRibInPre.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "adj-rib-in-pre"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost
// Per-neighbor table containing the paths received from
// the neighbor that are eligible for best-path selection
// after local input policy rules have been applied.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost) GetEntityData() *types.CommonEntityData {
    adjRibInPost.EntityData.YFilter = adjRibInPost.YFilter
    adjRibInPost.EntityData.YangName = "adj-rib-in-post"
    adjRibInPost.EntityData.BundleName = "openconfig"
    adjRibInPost.EntityData.ParentYangName = "neighbor"
    adjRibInPost.EntityData.SegmentPath = "adj-rib-in-post"
    adjRibInPost.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adjRibInPost.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adjRibInPost.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adjRibInPost.EntityData.Children = types.NewOrderedMap()
    adjRibInPost.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibInPost.Routes})
    adjRibInPost.EntityData.Leafs = types.NewOrderedMap()
    adjRibInPost.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", adjRibInPost.NumRoutes})

    adjRibInPost.EntityData.YListKeys = []string {}

    return &(adjRibInPost.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "adj-rib-in-post"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre
// Per-neighbor table containing paths eligble for
// sending (advertising) to the neighbor before output
// policy rules have been applied
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre) GetEntityData() *types.CommonEntityData {
    adjRibOutPre.EntityData.YFilter = adjRibOutPre.YFilter
    adjRibOutPre.EntityData.YangName = "adj-rib-out-pre"
    adjRibOutPre.EntityData.BundleName = "openconfig"
    adjRibOutPre.EntityData.ParentYangName = "neighbor"
    adjRibOutPre.EntityData.SegmentPath = "adj-rib-out-pre"
    adjRibOutPre.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adjRibOutPre.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adjRibOutPre.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adjRibOutPre.EntityData.Children = types.NewOrderedMap()
    adjRibOutPre.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibOutPre.Routes})
    adjRibOutPre.EntityData.Leafs = types.NewOrderedMap()
    adjRibOutPre.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", adjRibOutPre.NumRoutes})

    adjRibOutPre.EntityData.YListKeys = []string {}

    return &(adjRibOutPre.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "adj-rib-out-pre"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost
// Per-neighbor table containing paths eligble for
// sending (advertising) to the neighbor after output
// policy rules have been applied
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost) GetEntityData() *types.CommonEntityData {
    adjRibOutPost.EntityData.YFilter = adjRibOutPost.YFilter
    adjRibOutPost.EntityData.YangName = "adj-rib-out-post"
    adjRibOutPost.EntityData.BundleName = "openconfig"
    adjRibOutPost.EntityData.ParentYangName = "neighbor"
    adjRibOutPost.EntityData.SegmentPath = "adj-rib-out-post"
    adjRibOutPost.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adjRibOutPost.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adjRibOutPost.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adjRibOutPost.EntityData.Children = types.NewOrderedMap()
    adjRibOutPost.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibOutPost.Routes})
    adjRibOutPost.EntityData.Leafs = types.NewOrderedMap()
    adjRibOutPost.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", adjRibOutPost.NumRoutes})

    adjRibOutPost.EntityData.YListKeys = []string {}

    return &(adjRibOutPost.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "adj-rib-out-post"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv4Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast
// Routing tables for IPv6 unicast -- active when the
// afi-safi name is ipv6-unicast
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast struct {
    EntityData types.CommonEntityData
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

func (ipv6Unicast *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast) GetEntityData() *types.CommonEntityData {
    ipv6Unicast.EntityData.YFilter = ipv6Unicast.YFilter
    ipv6Unicast.EntityData.YangName = "ipv6-unicast"
    ipv6Unicast.EntityData.BundleName = "openconfig"
    ipv6Unicast.EntityData.ParentYangName = "afi-safi"
    ipv6Unicast.EntityData.SegmentPath = "ipv6-unicast"
    ipv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6Unicast.EntityData.Children = types.NewOrderedMap()
    ipv6Unicast.EntityData.Children.Append("loc-rib", types.YChild{"LocRib", &ipv6Unicast.LocRib})
    ipv6Unicast.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &ipv6Unicast.Neighbors})
    ipv6Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv6Unicast.EntityData.YListKeys = []string {}

    return &(ipv6Unicast.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes
}

func (locRib *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib) GetEntityData() *types.CommonEntityData {
    locRib.EntityData.YFilter = locRib.YFilter
    locRib.EntityData.YangName = "loc-rib"
    locRib.EntityData.BundleName = "openconfig"
    locRib.EntityData.ParentYangName = "ipv6-unicast"
    locRib.EntityData.SegmentPath = "loc-rib"
    locRib.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    locRib.EntityData.NamespaceTable = openconfig.GetNamespaces()
    locRib.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    locRib.EntityData.Children = types.NewOrderedMap()
    locRib.EntityData.Children.Append("routes", types.YChild{"Routes", &locRib.Routes})
    locRib.EntityData.Leafs = types.NewOrderedMap()
    locRib.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", locRib.NumRoutes})

    locRib.EntityData.YListKeys = []string {}

    return &(locRib.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "loc-rib"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_LocRib_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors
// Enclosing container for neighbor list
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of neighbors (peers) of the local BGP speaker. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor.
    Neighbor []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor
}

func (neighbors *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "ipv6-unicast"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbors.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor
// List of neighbors (peers) of the local BGP speaker
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
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

func (neighbor *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborAddress, "neighbor-address")
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("adj-rib-in-pre", types.YChild{"AdjRibInPre", &neighbor.AdjRibInPre})
    neighbor.EntityData.Children.Append("adj-rib-in-post", types.YChild{"AdjRibInPost", &neighbor.AdjRibInPost})
    neighbor.EntityData.Children.Append("adj-rib-out-pre", types.YChild{"AdjRibOutPre", &neighbor.AdjRibOutPre})
    neighbor.EntityData.Children.Append("adj-rib-out-post", types.YChild{"AdjRibOutPost", &neighbor.AdjRibOutPost})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighbor.NeighborAddress})

    neighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(neighbor.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre
// Per-neighbor table containing the NLRI updates
// received from the neighbor before any local input
// policy rules or filters have been applied.  This can
// be considered the 'raw' updates from the neighbor.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes
}

func (adjRibInPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre) GetEntityData() *types.CommonEntityData {
    adjRibInPre.EntityData.YFilter = adjRibInPre.YFilter
    adjRibInPre.EntityData.YangName = "adj-rib-in-pre"
    adjRibInPre.EntityData.BundleName = "openconfig"
    adjRibInPre.EntityData.ParentYangName = "neighbor"
    adjRibInPre.EntityData.SegmentPath = "adj-rib-in-pre"
    adjRibInPre.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adjRibInPre.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adjRibInPre.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adjRibInPre.EntityData.Children = types.NewOrderedMap()
    adjRibInPre.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibInPre.Routes})
    adjRibInPre.EntityData.Leafs = types.NewOrderedMap()
    adjRibInPre.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", adjRibInPre.NumRoutes})

    adjRibInPre.EntityData.YListKeys = []string {}

    return &(adjRibInPre.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "adj-rib-in-pre"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPre_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost
// Per-neighbor table containing the paths received from
// the neighbor that are eligible for best-path selection
// after local input policy rules have been applied.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes
}

func (adjRibInPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost) GetEntityData() *types.CommonEntityData {
    adjRibInPost.EntityData.YFilter = adjRibInPost.YFilter
    adjRibInPost.EntityData.YangName = "adj-rib-in-post"
    adjRibInPost.EntityData.BundleName = "openconfig"
    adjRibInPost.EntityData.ParentYangName = "neighbor"
    adjRibInPost.EntityData.SegmentPath = "adj-rib-in-post"
    adjRibInPost.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adjRibInPost.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adjRibInPost.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adjRibInPost.EntityData.Children = types.NewOrderedMap()
    adjRibInPost.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibInPost.Routes})
    adjRibInPost.EntityData.Leafs = types.NewOrderedMap()
    adjRibInPost.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", adjRibInPost.NumRoutes})

    adjRibInPost.EntityData.YListKeys = []string {}

    return &(adjRibInPost.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "adj-rib-in-post"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibInPost_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre
// Per-neighbor table containing paths eligble for
// sending (advertising) to the neighbor before output
// policy rules have been applied
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes
}

func (adjRibOutPre *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre) GetEntityData() *types.CommonEntityData {
    adjRibOutPre.EntityData.YFilter = adjRibOutPre.YFilter
    adjRibOutPre.EntityData.YangName = "adj-rib-out-pre"
    adjRibOutPre.EntityData.BundleName = "openconfig"
    adjRibOutPre.EntityData.ParentYangName = "neighbor"
    adjRibOutPre.EntityData.SegmentPath = "adj-rib-out-pre"
    adjRibOutPre.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adjRibOutPre.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adjRibOutPre.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adjRibOutPre.EntityData.Children = types.NewOrderedMap()
    adjRibOutPre.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibOutPre.Routes})
    adjRibOutPre.EntityData.Leafs = types.NewOrderedMap()
    adjRibOutPre.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", adjRibOutPre.NumRoutes})

    adjRibOutPre.EntityData.YListKeys = []string {}

    return &(adjRibOutPre.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "adj-rib-out-pre"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPre_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost
// Per-neighbor table containing paths eligble for
// sending (advertising) to the neighbor after output
// policy rules have been applied
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of route entries in the table. The type is interface{} with range:
    // 0..18446744073709551615.
    NumRoutes interface{}

    // Enclosing container for list of routes in the routing table.
    Routes BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes
}

func (adjRibOutPost *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost) GetEntityData() *types.CommonEntityData {
    adjRibOutPost.EntityData.YFilter = adjRibOutPost.YFilter
    adjRibOutPost.EntityData.YangName = "adj-rib-out-post"
    adjRibOutPost.EntityData.BundleName = "openconfig"
    adjRibOutPost.EntityData.ParentYangName = "neighbor"
    adjRibOutPost.EntityData.SegmentPath = "adj-rib-out-post"
    adjRibOutPost.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adjRibOutPost.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adjRibOutPost.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adjRibOutPost.EntityData.Children = types.NewOrderedMap()
    adjRibOutPost.EntityData.Children.Append("routes", types.YChild{"Routes", &adjRibOutPost.Routes})
    adjRibOutPost.EntityData.Leafs = types.NewOrderedMap()
    adjRibOutPost.EntityData.Leafs.Append("num-routes", types.YLeaf{"NumRoutes", adjRibOutPost.NumRoutes})

    adjRibOutPost.EntityData.YListKeys = []string {}

    return &(adjRibOutPost.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes
// Enclosing container for list of routes in the routing
// table.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of routes in the table. The type is slice of
    // BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route.
    Route []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route
}

func (routes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes) GetEntityData() *types.CommonEntityData {
    routes.EntityData.YFilter = routes.YFilter
    routes.EntityData.YangName = "routes"
    routes.EntityData.BundleName = "openconfig"
    routes.EntityData.ParentYangName = "adj-rib-out-post"
    routes.EntityData.SegmentPath = "routes"
    routes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routes.EntityData.Children = types.NewOrderedMap()
    routes.EntityData.Children.Append("route", types.YChild{"Route", nil})
    for i := range routes.Route {
        routes.EntityData.Children.Append(types.GetSegmentPath(routes.Route[i]), types.YChild{"Route", routes.Route[i]})
    }
    routes.EntityData.Leafs = types.NewOrderedMap()

    routes.EntityData.YListKeys = []string {}

    return &(routes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route
// List of routes in the table
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route struct {
    EntityData types.CommonEntityData
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

func (route *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "openconfig"
    route.EntityData.ParentYangName = "routes"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    route.EntityData.NamespaceTable = openconfig.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Children.Append("attributes", types.YChild{"Attributes", &route.Attributes})
    route.EntityData.Children.Append("ext-attributes", types.YChild{"ExtAttributes", &route.ExtAttributes})
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", route.Prefix})
    route.EntityData.Leafs.Append("last-modified-date", types.YLeaf{"LastModifiedDate", route.LastModifiedDate})
    route.EntityData.Leafs.Append("last-update-received", types.YLeaf{"LastUpdateReceived", route.LastUpdateReceived})
    route.EntityData.Leafs.Append("valid-route", types.YLeaf{"ValidRoute", route.ValidRoute})
    route.EntityData.Leafs.Append("invalid-reason", types.YLeaf{"InvalidReason", route.InvalidReason})
    route.EntityData.Leafs.Append("best-path", types.YLeaf{"BestPath", route.BestPath})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes
// Base BGP route attributes associated with this route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes struct {
    EntityData types.CommonEntityData
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
    // pattern:
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
    Community []interface{}

    // BGP attribute indicating the prefix has been aggregated by the specified AS
    // and router.
    Aggregator BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator
}

func (attributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "openconfig"
    attributes.EntityData.ParentYangName = "route"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    attributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("aggregator", types.YChild{"Aggregator", &attributes.Aggregator})
    attributes.EntityData.Leafs = types.NewOrderedMap()
    attributes.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", attributes.Origin})
    attributes.EntityData.Leafs.Append("as-path", types.YLeaf{"AsPath", attributes.AsPath})
    attributes.EntityData.Leafs.Append("as4-path", types.YLeaf{"As4Path", attributes.As4Path})
    attributes.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", attributes.NextHop})
    attributes.EntityData.Leafs.Append("med", types.YLeaf{"Med", attributes.Med})
    attributes.EntityData.Leafs.Append("local-pref", types.YLeaf{"LocalPref", attributes.LocalPref})
    attributes.EntityData.Leafs.Append("atomic-aggr", types.YLeaf{"AtomicAggr", attributes.AtomicAggr})
    attributes.EntityData.Leafs.Append("community", types.YLeaf{"Community", attributes.Community})

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator
// BGP attribute indicating the prefix has been aggregated by
// the specified AS and router.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator struct {
    EntityData types.CommonEntityData
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

func (aggregator *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_Attributes_Aggregator) GetEntityData() *types.CommonEntityData {
    aggregator.EntityData.YFilter = aggregator.YFilter
    aggregator.EntityData.YangName = "aggregator"
    aggregator.EntityData.BundleName = "openconfig"
    aggregator.EntityData.ParentYangName = "attributes"
    aggregator.EntityData.SegmentPath = "aggregator"
    aggregator.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregator.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregator.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregator.EntityData.Children = types.NewOrderedMap()
    aggregator.EntityData.Leafs = types.NewOrderedMap()
    aggregator.EntityData.Leafs.Append("as", types.YLeaf{"As", aggregator.As})
    aggregator.EntityData.Leafs.Append("as4", types.YLeaf{"As4", aggregator.As4})
    aggregator.EntityData.Leafs.Append("address", types.YLeaf{"Address", aggregator.Address})

    aggregator.EntityData.YListKeys = []string {}

    return &(aggregator.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes
// Extended BGP route attributes associated with this
// route
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes struct {
    EntityData types.CommonEntityData
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
    // ^(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-target:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9]):(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$,
    // or slice of string with pattern:
    // ^route\-origin:(429496729[0-5]|42949672[0-8][0-9]|4294967[0-1][0-9]{2}|429496[0-6][0-9]{3}|42949[0-5][0-9]{4}|4294[0-8][0-9]{5}|429[0-3][0-9]{6}|4[0-1][0-9]{7}|[1-3][0-9]{9}|[1-9][0-9]{1,8}|[0-9]):(6553[0-5]|655[0-2][0-9]|654[0-9]{2}|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[1-9][0-9]{1,3}|[0-9])$.
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
    UnknownAttribute []*BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute
}

func (extAttributes *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes) GetEntityData() *types.CommonEntityData {
    extAttributes.EntityData.YFilter = extAttributes.YFilter
    extAttributes.EntityData.YangName = "ext-attributes"
    extAttributes.EntityData.BundleName = "openconfig"
    extAttributes.EntityData.ParentYangName = "route"
    extAttributes.EntityData.SegmentPath = "ext-attributes"
    extAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    extAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    extAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    extAttributes.EntityData.Children = types.NewOrderedMap()
    extAttributes.EntityData.Children.Append("unknown-attribute", types.YChild{"UnknownAttribute", nil})
    for i := range extAttributes.UnknownAttribute {
        extAttributes.EntityData.Children.Append(types.GetSegmentPath(extAttributes.UnknownAttribute[i]), types.YChild{"UnknownAttribute", extAttributes.UnknownAttribute[i]})
    }
    extAttributes.EntityData.Leafs = types.NewOrderedMap()
    extAttributes.EntityData.Leafs.Append("originator-id", types.YLeaf{"OriginatorId", extAttributes.OriginatorId})
    extAttributes.EntityData.Leafs.Append("cluster-list", types.YLeaf{"ClusterList", extAttributes.ClusterList})
    extAttributes.EntityData.Leafs.Append("ext-community", types.YLeaf{"ExtCommunity", extAttributes.ExtCommunity})
    extAttributes.EntityData.Leafs.Append("aigp", types.YLeaf{"Aigp", extAttributes.Aigp})
    extAttributes.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", extAttributes.PathId})

    extAttributes.EntityData.YListKeys = []string {}

    return &(extAttributes.EntityData)
}

// BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute
// This list contains received attributes that are unrecognized
// or unsupported by the local router.  The list may be empty.
type BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute struct {
    EntityData types.CommonEntityData
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

func (unknownAttribute *BgpRib_AfiSafis_AfiSafi_Ipv6Unicast_Neighbors_Neighbor_AdjRibOutPost_Routes_Route_ExtAttributes_UnknownAttribute) GetEntityData() *types.CommonEntityData {
    unknownAttribute.EntityData.YFilter = unknownAttribute.YFilter
    unknownAttribute.EntityData.YangName = "unknown-attribute"
    unknownAttribute.EntityData.BundleName = "openconfig"
    unknownAttribute.EntityData.ParentYangName = "ext-attributes"
    unknownAttribute.EntityData.SegmentPath = "unknown-attribute" + types.AddKeyToken(unknownAttribute.AttrType, "attr-type")
    unknownAttribute.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unknownAttribute.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unknownAttribute.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unknownAttribute.EntityData.Children = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs = types.NewOrderedMap()
    unknownAttribute.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", unknownAttribute.AttrType})
    unknownAttribute.EntityData.Leafs.Append("attr-len", types.YLeaf{"AttrLen", unknownAttribute.AttrLen})
    unknownAttribute.EntityData.Leafs.Append("attr-value", types.YLeaf{"AttrValue", unknownAttribute.AttrValue})

    unknownAttribute.EntityData.YListKeys = []string {"AttrType"}

    return &(unknownAttribute.EntityData)
}

