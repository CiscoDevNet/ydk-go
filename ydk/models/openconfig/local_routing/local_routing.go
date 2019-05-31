// This module describes configuration and operational state data
// for routes that are locally generated, i.e., not created by
// dynamic routing protocols.  These include static routes, locally
// created aggregate routes for reducing the number of constituent
// routes that must be advertised, summary routes for IGPs, etc.
// 
// This model expresses locally generated routes as generically as
// possible, avoiding configuration of protocol-specific attributes
// at the time of route creation.  This is primarily to avoid
// assumptions about how underlying router implementations handle
// route attributes in various routing table data structures they
// maintain.  Hence, the definition of locally generated routes
// essentially creates 'bare' routes that do not have any protocol-
// specific attributes.
// 
// When protocol-specific attributes must be attached to a route
// (e.g., communities on a locally defined route meant to be
// advertised via BGP), the attributes should be attached via a
// protocol-specific policy after importing the route into the
// protocol for distribution (again via routing policy).
package local_routing

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package local_routing"))
    ydk.RegisterEntity("{http://openconfig.net/yang/local-routing local-routes}", reflect.TypeOf(LocalRoutes{}))
    ydk.RegisterEntity("openconfig-local-routing:local-routes", reflect.TypeOf(LocalRoutes{}))
}

type LOCALDEFINEDNEXTHOP struct {
}

func (id LOCALDEFINEDNEXTHOP) String() string {
	return "openconfig-local-routing:LOCAL_DEFINED_NEXT_HOP"
}

type DROP struct {
}

func (id DROP) String() string {
	return "openconfig-local-routing:DROP"
}

type LOCALLINK struct {
}

func (id LOCALLINK) String() string {
	return "openconfig-local-routing:LOCAL_LINK"
}

// LocalRoutes
// Top-level container for local routes
type LocalRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for locally defined routes.
    Config LocalRoutes_Config

    // Operational state data for locally defined routes.
    State LocalRoutes_State

    // Enclosing container for the list of static routes.
    StaticRoutes LocalRoutes_StaticRoutes

    // Enclosing container for locally-defined aggregate routes.
    LocalAggregates LocalRoutes_LocalAggregates
}

func (localRoutes *LocalRoutes) GetEntityData() *types.CommonEntityData {
    localRoutes.EntityData.YFilter = localRoutes.YFilter
    localRoutes.EntityData.YangName = "local-routes"
    localRoutes.EntityData.BundleName = "openconfig"
    localRoutes.EntityData.ParentYangName = "openconfig-local-routing"
    localRoutes.EntityData.SegmentPath = "openconfig-local-routing:local-routes"
    localRoutes.EntityData.AbsolutePath = localRoutes.EntityData.SegmentPath
    localRoutes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    localRoutes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    localRoutes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    localRoutes.EntityData.Children = types.NewOrderedMap()
    localRoutes.EntityData.Children.Append("config", types.YChild{"Config", &localRoutes.Config})
    localRoutes.EntityData.Children.Append("state", types.YChild{"State", &localRoutes.State})
    localRoutes.EntityData.Children.Append("static-routes", types.YChild{"StaticRoutes", &localRoutes.StaticRoutes})
    localRoutes.EntityData.Children.Append("local-aggregates", types.YChild{"LocalAggregates", &localRoutes.LocalAggregates})
    localRoutes.EntityData.Leafs = types.NewOrderedMap()

    localRoutes.EntityData.YListKeys = []string {}

    return &(localRoutes.EntityData)
}

// LocalRoutes_Config
// Configuration data for locally defined routes
type LocalRoutes_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *LocalRoutes_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "local-routes"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// LocalRoutes_State
// Operational state data for locally defined routes
type LocalRoutes_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (state *LocalRoutes_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "local-routes"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// LocalRoutes_StaticRoutes
// Enclosing container for the list of static routes
type LocalRoutes_StaticRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of locally configured static routes. The type is slice of
    // LocalRoutes_StaticRoutes_Static.
    Static []*LocalRoutes_StaticRoutes_Static
}

func (staticRoutes *LocalRoutes_StaticRoutes) GetEntityData() *types.CommonEntityData {
    staticRoutes.EntityData.YFilter = staticRoutes.YFilter
    staticRoutes.EntityData.YangName = "static-routes"
    staticRoutes.EntityData.BundleName = "openconfig"
    staticRoutes.EntityData.ParentYangName = "local-routes"
    staticRoutes.EntityData.SegmentPath = "static-routes"
    staticRoutes.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/" + staticRoutes.EntityData.SegmentPath
    staticRoutes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    staticRoutes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    staticRoutes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    staticRoutes.EntityData.Children = types.NewOrderedMap()
    staticRoutes.EntityData.Children.Append("static", types.YChild{"Static", nil})
    for i := range staticRoutes.Static {
        staticRoutes.EntityData.Children.Append(types.GetSegmentPath(staticRoutes.Static[i]), types.YChild{"Static", staticRoutes.Static[i]})
    }
    staticRoutes.EntityData.Leafs = types.NewOrderedMap()

    staticRoutes.EntityData.YListKeys = []string {}

    return &(staticRoutes.EntityData)
}

// LocalRoutes_StaticRoutes_Static
// List of locally configured static routes
type LocalRoutes_StaticRoutes_Static struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the destination prefix list key. The
    // type is one of the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$'.
    Prefix interface{}

    // Configuration data for static routes.
    Config LocalRoutes_StaticRoutes_Static_Config

    // Operational state data for static routes.
    State LocalRoutes_StaticRoutes_Static_State

    // Configuration and state parameters relating to the next-hops that are to be
    // utilised for the static route being specified.
    NextHops LocalRoutes_StaticRoutes_Static_NextHops
}

func (static *LocalRoutes_StaticRoutes_Static) GetEntityData() *types.CommonEntityData {
    static.EntityData.YFilter = static.YFilter
    static.EntityData.YangName = "static"
    static.EntityData.BundleName = "openconfig"
    static.EntityData.ParentYangName = "static-routes"
    static.EntityData.SegmentPath = "static" + types.AddKeyToken(static.Prefix, "prefix")
    static.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/" + static.EntityData.SegmentPath
    static.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    static.EntityData.NamespaceTable = openconfig.GetNamespaces()
    static.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    static.EntityData.Children = types.NewOrderedMap()
    static.EntityData.Children.Append("config", types.YChild{"Config", &static.Config})
    static.EntityData.Children.Append("state", types.YChild{"State", &static.State})
    static.EntityData.Children.Append("next-hops", types.YChild{"NextHops", &static.NextHops})
    static.EntityData.Leafs = types.NewOrderedMap()
    static.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", static.Prefix})

    static.EntityData.YListKeys = []string {"Prefix"}

    return &(static.EntityData)
}

// LocalRoutes_StaticRoutes_Static_Config
// Configuration data for static routes
type LocalRoutes_StaticRoutes_Static_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination prefix for the static route, either IPv4 or IPv6. The type is
    // one of the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$'.
    Prefix interface{}

    // Set a generic tag value on the route. This tag can be used for filtering
    // routes that are distributed to other routing protocols. The type is one of
    // the following types: int with range: 0..4294967295, or string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    SetTag interface{}
}

func (config *LocalRoutes_StaticRoutes_Static_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "static"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/static/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", config.Prefix})
    config.EntityData.Leafs.Append("set-tag", types.YLeaf{"SetTag", config.SetTag})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// LocalRoutes_StaticRoutes_Static_State
// Operational state data for static routes
type LocalRoutes_StaticRoutes_Static_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination prefix for the static route, either IPv4 or IPv6. The type is
    // one of the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$'.
    Prefix interface{}

    // Set a generic tag value on the route. This tag can be used for filtering
    // routes that are distributed to other routing protocols. The type is one of
    // the following types: int with range: 0..4294967295, or string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    SetTag interface{}
}

func (state *LocalRoutes_StaticRoutes_Static_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "static"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/static/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", state.Prefix})
    state.EntityData.Leafs.Append("set-tag", types.YLeaf{"SetTag", state.SetTag})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// LocalRoutes_StaticRoutes_Static_NextHops
// Configuration and state parameters relating to the
// next-hops that are to be utilised for the static
// route being specified
type LocalRoutes_StaticRoutes_Static_NextHops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of next-hops to be utilised for the static route being specified.
    // The type is slice of LocalRoutes_StaticRoutes_Static_NextHops_NextHop.
    NextHop []*LocalRoutes_StaticRoutes_Static_NextHops_NextHop
}

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetEntityData() *types.CommonEntityData {
    nextHops.EntityData.YFilter = nextHops.YFilter
    nextHops.EntityData.YangName = "next-hops"
    nextHops.EntityData.BundleName = "openconfig"
    nextHops.EntityData.ParentYangName = "static"
    nextHops.EntityData.SegmentPath = "next-hops"
    nextHops.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/static/" + nextHops.EntityData.SegmentPath
    nextHops.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    nextHops.EntityData.NamespaceTable = openconfig.GetNamespaces()
    nextHops.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    nextHops.EntityData.Children = types.NewOrderedMap()
    nextHops.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range nextHops.NextHop {
        nextHops.EntityData.Children.Append(types.GetSegmentPath(nextHops.NextHop[i]), types.YChild{"NextHop", nextHops.NextHop[i]})
    }
    nextHops.EntityData.Leafs = types.NewOrderedMap()

    nextHops.EntityData.YListKeys = []string {}

    return &(nextHops.EntityData)
}

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop
// A list of next-hops to be utilised for the static
// route being specified.
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A reference to the index of the current next-hop.
    // The index is intended to be a user-specified value which can be used to
    // reference the next-hop in question, without any other semantics being
    // assigned to it. The type is string. Refers to
    // local_routing.LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config_Index
    Index interface{}

    // Configuration parameters relating to the next-hop entry.
    Config LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config

    // Operational state parameters relating to the next-hop entry.
    State LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State

    // Reference to an interface or subinterface.
    InterfaceRef LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef
}

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "openconfig"
    nextHop.EntityData.ParentYangName = "next-hops"
    nextHop.EntityData.SegmentPath = "next-hop" + types.AddKeyToken(nextHop.Index, "index")
    nextHop.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/static/next-hops/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    nextHop.EntityData.NamespaceTable = openconfig.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("config", types.YChild{"Config", &nextHop.Config})
    nextHop.EntityData.Children.Append("state", types.YChild{"State", &nextHop.State})
    nextHop.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &nextHop.InterfaceRef})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("index", types.YLeaf{"Index", nextHop.Index})

    nextHop.EntityData.YListKeys = []string {"Index"}

    return &(nextHop.EntityData)
}

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config
// Configuration parameters relating to the next-hop
// entry
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An user-specified identifier utilised to uniquely reference the next-hop
    // entry in the next-hop list. The value of this index has no semantic meaning
    // other than for referencing the entry. The type is string.
    Index interface{}

    // The next-hop that is to be used for the static route - this may be
    // specified as an IP address, an interface or a pre-defined next-hop type -
    // for instance, DROP or LOCAL_LINK. When this leaf is not set, and the
    // interface-ref value is specified for the next-hop, then the system should
    // treat the prefix as though it is directly connected to the interface. The
    // type is one of the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.,
    // or :go:struct:`LOCALDEFINEDNEXTHOP
    // <ydk/models/local_routing/LOCALDEFINEDNEXTHOP>`.
    NextHop interface{}

    // A metric which is utilised to specify the preference of the next-hop entry
    // when it is injected into the RIB. The lower the metric, the more preferable
    // the prefix is. When this value is not specified the metric is inherited
    // from the default metric utilised for static routes within the network
    // instance that the static routes are being instantiated. When multiple
    // next-hops are specified for a static route, the metric is utilised to
    // determine which of the next-hops is to be installed in the RIB. When
    // multiple next-hops have the same metric (be it specified, or simply the
    // default) then these next-hops should all be installed in the RIB. The type
    // is interface{} with range: 0..4294967295.
    Metric interface{}

    // Determines whether the next-hop should be allowed to be looked up
    // recursively - i.e., via a RIB entry which has been installed by a routing
    // protocol, or another static route - rather than needing to be connected
    // directly to an interface of the local system within the current network
    // instance. When the interface reference specified within the next-hop entry
    // is set (i.e., is not null) then forwarding is restricted to being via the
    // interface specified - and recursion is hence disabled. The type is bool.
    // The default value is false.
    Recurse interface{}
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "next-hop"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/static/next-hops/next-hop/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("index", types.YLeaf{"Index", config.Index})
    config.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", config.NextHop})
    config.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", config.Metric})
    config.EntityData.Leafs.Append("recurse", types.YLeaf{"Recurse", config.Recurse})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State
// Operational state parameters relating to the
// next-hop entry
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An user-specified identifier utilised to uniquely reference the next-hop
    // entry in the next-hop list. The value of this index has no semantic meaning
    // other than for referencing the entry. The type is string.
    Index interface{}

    // The next-hop that is to be used for the static route - this may be
    // specified as an IP address, an interface or a pre-defined next-hop type -
    // for instance, DROP or LOCAL_LINK. When this leaf is not set, and the
    // interface-ref value is specified for the next-hop, then the system should
    // treat the prefix as though it is directly connected to the interface. The
    // type is one of the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.,
    // or :go:struct:`LOCALDEFINEDNEXTHOP
    // <ydk/models/local_routing/LOCALDEFINEDNEXTHOP>`.
    NextHop interface{}

    // A metric which is utilised to specify the preference of the next-hop entry
    // when it is injected into the RIB. The lower the metric, the more preferable
    // the prefix is. When this value is not specified the metric is inherited
    // from the default metric utilised for static routes within the network
    // instance that the static routes are being instantiated. When multiple
    // next-hops are specified for a static route, the metric is utilised to
    // determine which of the next-hops is to be installed in the RIB. When
    // multiple next-hops have the same metric (be it specified, or simply the
    // default) then these next-hops should all be installed in the RIB. The type
    // is interface{} with range: 0..4294967295.
    Metric interface{}

    // Determines whether the next-hop should be allowed to be looked up
    // recursively - i.e., via a RIB entry which has been installed by a routing
    // protocol, or another static route - rather than needing to be connected
    // directly to an interface of the local system within the current network
    // instance. When the interface reference specified within the next-hop entry
    // is set (i.e., is not null) then forwarding is restricted to being via the
    // interface specified - and recursion is hence disabled. The type is bool.
    // The default value is false.
    Recurse interface{}
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "next-hop"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/static/next-hops/next-hop/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("index", types.YLeaf{"Index", state.Index})
    state.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", state.NextHop})
    state.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", state.Metric})
    state.EntityData.Leafs.Append("recurse", types.YLeaf{"Recurse", state.Recurse})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef
// Reference to an interface or subinterface
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config

    // Operational state for interface-ref.
    State LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State
}

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "next-hop"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/static/next-hops/next-hop/" + interfaceRef.EntityData.SegmentPath
    interfaceRef.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("config", types.YChild{"Config", &interfaceRef.Config})
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config
// Configured reference to interface / subinterface
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-ref"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/static/next-hops/next-hop/interface-ref/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", config.Interface})
    config.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", config.Subinterface})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State
// Operational state for interface-ref
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/static-routes/static/next-hops/next-hop/interface-ref/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// LocalRoutes_LocalAggregates
// Enclosing container for locally-defined aggregate
// routes
type LocalRoutes_LocalAggregates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of aggregates. The type is slice of
    // LocalRoutes_LocalAggregates_Aggregate.
    Aggregate []*LocalRoutes_LocalAggregates_Aggregate
}

func (localAggregates *LocalRoutes_LocalAggregates) GetEntityData() *types.CommonEntityData {
    localAggregates.EntityData.YFilter = localAggregates.YFilter
    localAggregates.EntityData.YangName = "local-aggregates"
    localAggregates.EntityData.BundleName = "openconfig"
    localAggregates.EntityData.ParentYangName = "local-routes"
    localAggregates.EntityData.SegmentPath = "local-aggregates"
    localAggregates.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/" + localAggregates.EntityData.SegmentPath
    localAggregates.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    localAggregates.EntityData.NamespaceTable = openconfig.GetNamespaces()
    localAggregates.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    localAggregates.EntityData.Children = types.NewOrderedMap()
    localAggregates.EntityData.Children.Append("aggregate", types.YChild{"Aggregate", nil})
    for i := range localAggregates.Aggregate {
        localAggregates.EntityData.Children.Append(types.GetSegmentPath(localAggregates.Aggregate[i]), types.YChild{"Aggregate", localAggregates.Aggregate[i]})
    }
    localAggregates.EntityData.Leafs = types.NewOrderedMap()

    localAggregates.EntityData.YListKeys = []string {}

    return &(localAggregates.EntityData)
}

// LocalRoutes_LocalAggregates_Aggregate
// List of aggregates
type LocalRoutes_LocalAggregates_Aggregate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the configured prefix for this
    // aggregate. The type is one of the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$'.
    Prefix interface{}

    // Configuration data for aggregate advertisements.
    Config LocalRoutes_LocalAggregates_Aggregate_Config

    // Operational state data for aggregate advertisements.
    State LocalRoutes_LocalAggregates_Aggregate_State
}

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetEntityData() *types.CommonEntityData {
    aggregate.EntityData.YFilter = aggregate.YFilter
    aggregate.EntityData.YangName = "aggregate"
    aggregate.EntityData.BundleName = "openconfig"
    aggregate.EntityData.ParentYangName = "local-aggregates"
    aggregate.EntityData.SegmentPath = "aggregate" + types.AddKeyToken(aggregate.Prefix, "prefix")
    aggregate.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/local-aggregates/" + aggregate.EntityData.SegmentPath
    aggregate.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregate.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregate.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregate.EntityData.Children = types.NewOrderedMap()
    aggregate.EntityData.Children.Append("config", types.YChild{"Config", &aggregate.Config})
    aggregate.EntityData.Children.Append("state", types.YChild{"State", &aggregate.State})
    aggregate.EntityData.Leafs = types.NewOrderedMap()
    aggregate.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", aggregate.Prefix})

    aggregate.EntityData.YListKeys = []string {"Prefix"}

    return &(aggregate.EntityData)
}

// LocalRoutes_LocalAggregates_Aggregate_Config
// Configuration data for aggregate advertisements
type LocalRoutes_LocalAggregates_Aggregate_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Aggregate prefix to be advertised. The type is one of the following types:
    // string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$'.
    Prefix interface{}

    // When true, install the aggregate route with a discard next-hop -- traffic
    // destined to the aggregate will be discarded with no ICMP message generated.
    // When false, traffic destined to an aggregate address when no constituent
    // routes are present will generate an ICMP unreachable message. The type is
    // bool. The default value is false.
    Discard interface{}

    // Set a generic tag value on the route. This tag can be used for filtering
    // routes that are distributed to other routing protocols. The type is one of
    // the following types: int with range: 0..4294967295, or string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    SetTag interface{}
}

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "aggregate"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/local-aggregates/aggregate/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", config.Prefix})
    config.EntityData.Leafs.Append("discard", types.YLeaf{"Discard", config.Discard})
    config.EntityData.Leafs.Append("set-tag", types.YLeaf{"SetTag", config.SetTag})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// LocalRoutes_LocalAggregates_Aggregate_State
// Operational state data for aggregate
// advertisements
type LocalRoutes_LocalAggregates_Aggregate_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Aggregate prefix to be advertised. The type is one of the following types:
    // string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$'.
    Prefix interface{}

    // When true, install the aggregate route with a discard next-hop -- traffic
    // destined to the aggregate will be discarded with no ICMP message generated.
    // When false, traffic destined to an aggregate address when no constituent
    // routes are present will generate an ICMP unreachable message. The type is
    // bool. The default value is false.
    Discard interface{}

    // Set a generic tag value on the route. This tag can be used for filtering
    // routes that are distributed to other routing protocols. The type is one of
    // the following types: int with range: 0..4294967295, or string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    SetTag interface{}
}

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "aggregate"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-local-routing:local-routes/local-aggregates/aggregate/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", state.Prefix})
    state.EntityData.Leafs.Append("discard", types.YLeaf{"Discard", state.Discard})
    state.EntityData.Leafs.Append("set-tag", types.YLeaf{"SetTag", state.SetTag})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

