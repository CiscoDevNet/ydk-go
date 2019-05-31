// This module describes a YANG model for BGP protocol
// configuration.It is a limited subset of all of the configuration
// parameters available in the variety of vendor implementations,
// hence it is expected that it would be augmented with vendor-
// specific configuration data as needed. Additional modules or
// submodules to handle other aspects of BGP configuration,
// including policy, VRFs, VPNs, and additional address families
// are also expected.
// 
// This model supports the following BGP configuration level
// hierarchy:
// 
//  BGP
//    |
//    +-> [ global BGP configuration ]
//      +-> AFI / SAFI global
//    +-> peer group
//      +-> [ peer group config ]
//      +-> AFI / SAFI [ per-AFI overrides ]
//    +-> neighbor
//      +-> [ neighbor config ]
//      +-> [ optional pointer to peer-group ]
//      +-> AFI / SAFI [ per-AFI overrides ]
package bgp

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bgp"))
    ydk.RegisterEntity("{http://openconfig.net/yang/bgp bgp}", reflect.TypeOf(Bgp{}))
    ydk.RegisterEntity("openconfig-bgp:bgp", reflect.TypeOf(Bgp{}))
}

// Bgp
// Top-level configuration and state for the BGP router
type Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global configuration for the BGP router.
    Global Bgp_Global

    // Configuration for BGP neighbors.
    Neighbors Bgp_Neighbors

    // Configuration for BGP peer-groups.
    PeerGroups Bgp_PeerGroups
}

func (bgp *Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "openconfig"
    bgp.EntityData.ParentYangName = "openconfig-bgp"
    bgp.EntityData.SegmentPath = "openconfig-bgp:bgp"
    bgp.EntityData.AbsolutePath = bgp.EntityData.SegmentPath
    bgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("global", types.YChild{"Global", &bgp.Global})
    bgp.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &bgp.Neighbors})
    bgp.EntityData.Children.Append("peer-groups", types.YChild{"PeerGroups", &bgp.PeerGroups})
    bgp.EntityData.Leafs = types.NewOrderedMap()

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Bgp_Global
// Global configuration for the BGP router
type Bgp_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the global BGP router.
    Config Bgp_Global_Config

    // State information relating to the global BGP router.
    State Bgp_Global_State

    // Administrative distance (or preference) assigned to routes received from
    // different sources (external, internal, and local).
    DefaultRouteDistance Bgp_Global_DefaultRouteDistance

    // Parameters indicating whether the local system acts as part of a BGP
    // confederation.
    Confederation Bgp_Global_Confederation

    // Parameters relating the graceful restart mechanism for BGP.
    GracefulRestart Bgp_Global_GracefulRestart

    // Parameters related to the use of multiple paths for the same NLRI.
    UseMultiplePaths Bgp_Global_UseMultiplePaths

    // Parameters relating to options for route selection.
    RouteSelectionOptions Bgp_Global_RouteSelectionOptions

    // Address family specific configuration.
    AfiSafis Bgp_Global_AfiSafis

    // A list of IP prefixes from which the system should:  - Accept connections
    // to the BGP daemon  - Dynamically configure a BGP neighbor corresponding to
    // the    source address of the remote system, using the parameters    of the
    // specified peer-group. For such neighbors, an entry within the neighbor list
    // should be created, indicating that the peer was dynamically configured, and
    // referencing the peer-group from which the configuration was derived.
    DynamicNeighborPrefixes Bgp_Global_DynamicNeighborPrefixes
}

func (global *Bgp_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "openconfig"
    global.EntityData.ParentYangName = "bgp"
    global.EntityData.SegmentPath = "global"
    global.EntityData.AbsolutePath = "openconfig-bgp:bgp/" + global.EntityData.SegmentPath
    global.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    global.EntityData.NamespaceTable = openconfig.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("config", types.YChild{"Config", &global.Config})
    global.EntityData.Children.Append("state", types.YChild{"State", &global.State})
    global.EntityData.Children.Append("default-route-distance", types.YChild{"DefaultRouteDistance", &global.DefaultRouteDistance})
    global.EntityData.Children.Append("confederation", types.YChild{"Confederation", &global.Confederation})
    global.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &global.GracefulRestart})
    global.EntityData.Children.Append("use-multiple-paths", types.YChild{"UseMultiplePaths", &global.UseMultiplePaths})
    global.EntityData.Children.Append("route-selection-options", types.YChild{"RouteSelectionOptions", &global.RouteSelectionOptions})
    global.EntityData.Children.Append("afi-safis", types.YChild{"AfiSafis", &global.AfiSafis})
    global.EntityData.Children.Append("dynamic-neighbor-prefixes", types.YChild{"DynamicNeighborPrefixes", &global.DynamicNeighborPrefixes})
    global.EntityData.Leafs = types.NewOrderedMap()

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// Bgp_Global_Config
// Configuration parameters relating to the global BGP router
type Bgp_Global_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local autonomous system number of the router.  Uses the 32-bit as-number
    // type from the model in RFC 6991. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    As interface{}

    // Router id of the router - an unsigned 32-bit integer expressed in dotted
    // quad notation. The type is string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$'.
    RouterId interface{}
}

func (config *Bgp_Global_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "global"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("as", types.YLeaf{"As", config.As})
    config.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", config.RouterId})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_State
// State information relating to the global BGP router
type Bgp_Global_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local autonomous system number of the router.  Uses the 32-bit as-number
    // type from the model in RFC 6991. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    As interface{}

    // Router id of the router - an unsigned 32-bit integer expressed in dotted
    // quad notation. The type is string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$'.
    RouterId interface{}

    // Total number of BGP paths within the context. The type is interface{} with
    // range: 0..4294967295.
    TotalPaths interface{}

    // Total number of BGP prefixes received within the context. The type is
    // interface{} with range: 0..4294967295.
    TotalPrefixes interface{}
}

func (state *Bgp_Global_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "global"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("as", types.YLeaf{"As", state.As})
    state.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", state.RouterId})
    state.EntityData.Leafs.Append("total-paths", types.YLeaf{"TotalPaths", state.TotalPaths})
    state.EntityData.Leafs.Append("total-prefixes", types.YLeaf{"TotalPrefixes", state.TotalPrefixes})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_DefaultRouteDistance
// Administrative distance (or preference) assigned to
// routes received from different sources
// (external, internal, and local).
type Bgp_Global_DefaultRouteDistance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the default route distance.
    Config Bgp_Global_DefaultRouteDistance_Config

    // State information relating to the default route distance.
    State Bgp_Global_DefaultRouteDistance_State
}

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetEntityData() *types.CommonEntityData {
    defaultRouteDistance.EntityData.YFilter = defaultRouteDistance.YFilter
    defaultRouteDistance.EntityData.YangName = "default-route-distance"
    defaultRouteDistance.EntityData.BundleName = "openconfig"
    defaultRouteDistance.EntityData.ParentYangName = "global"
    defaultRouteDistance.EntityData.SegmentPath = "default-route-distance"
    defaultRouteDistance.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/" + defaultRouteDistance.EntityData.SegmentPath
    defaultRouteDistance.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    defaultRouteDistance.EntityData.NamespaceTable = openconfig.GetNamespaces()
    defaultRouteDistance.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    defaultRouteDistance.EntityData.Children = types.NewOrderedMap()
    defaultRouteDistance.EntityData.Children.Append("config", types.YChild{"Config", &defaultRouteDistance.Config})
    defaultRouteDistance.EntityData.Children.Append("state", types.YChild{"State", &defaultRouteDistance.State})
    defaultRouteDistance.EntityData.Leafs = types.NewOrderedMap()

    defaultRouteDistance.EntityData.YListKeys = []string {}

    return &(defaultRouteDistance.EntityData)
}

// Bgp_Global_DefaultRouteDistance_Config
// Configuration parameters relating to the default route
// distance
type Bgp_Global_DefaultRouteDistance_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative distance for routes learned from external BGP (eBGP). The
    // type is interface{} with range: 1..255.
    ExternalRouteDistance interface{}

    // Administrative distance for routes learned from internal BGP (iBGP). The
    // type is interface{} with range: 1..255.
    InternalRouteDistance interface{}
}

func (config *Bgp_Global_DefaultRouteDistance_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "default-route-distance"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/default-route-distance/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("external-route-distance", types.YLeaf{"ExternalRouteDistance", config.ExternalRouteDistance})
    config.EntityData.Leafs.Append("internal-route-distance", types.YLeaf{"InternalRouteDistance", config.InternalRouteDistance})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_DefaultRouteDistance_State
// State information relating to the default route distance
type Bgp_Global_DefaultRouteDistance_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative distance for routes learned from external BGP (eBGP). The
    // type is interface{} with range: 1..255.
    ExternalRouteDistance interface{}

    // Administrative distance for routes learned from internal BGP (iBGP). The
    // type is interface{} with range: 1..255.
    InternalRouteDistance interface{}
}

func (state *Bgp_Global_DefaultRouteDistance_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "default-route-distance"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/default-route-distance/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("external-route-distance", types.YLeaf{"ExternalRouteDistance", state.ExternalRouteDistance})
    state.EntityData.Leafs.Append("internal-route-distance", types.YLeaf{"InternalRouteDistance", state.InternalRouteDistance})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_Confederation
// Parameters indicating whether the local system acts as part
// of a BGP confederation
type Bgp_Global_Confederation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to BGP confederations.
    Config Bgp_Global_Confederation_Config

    // State information relating to the BGP confederations.
    State Bgp_Global_Confederation_State
}

func (confederation *Bgp_Global_Confederation) GetEntityData() *types.CommonEntityData {
    confederation.EntityData.YFilter = confederation.YFilter
    confederation.EntityData.YangName = "confederation"
    confederation.EntityData.BundleName = "openconfig"
    confederation.EntityData.ParentYangName = "global"
    confederation.EntityData.SegmentPath = "confederation"
    confederation.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/" + confederation.EntityData.SegmentPath
    confederation.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    confederation.EntityData.NamespaceTable = openconfig.GetNamespaces()
    confederation.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    confederation.EntityData.Children = types.NewOrderedMap()
    confederation.EntityData.Children.Append("config", types.YChild{"Config", &confederation.Config})
    confederation.EntityData.Children.Append("state", types.YChild{"State", &confederation.State})
    confederation.EntityData.Leafs = types.NewOrderedMap()

    confederation.EntityData.YListKeys = []string {}

    return &(confederation.EntityData)
}

// Bgp_Global_Confederation_Config
// Configuration parameters relating to BGP confederations
type Bgp_Global_Confederation_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When this leaf is set to true it indicates that the local-AS is part of a
    // BGP confederation. The type is bool.
    Enabled interface{}

    // Confederation identifier for the autonomous system. The type is interface{}
    // with range: 0..4294967295.
    Identifier interface{}

    // Remote autonomous systems that are to be treated as part of the local
    // confederation. The type is slice of interface{} with range: 0..4294967295.
    MemberAs []interface{}
}

func (config *Bgp_Global_Confederation_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "confederation"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/confederation/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", config.Identifier})
    config.EntityData.Leafs.Append("member-as", types.YLeaf{"MemberAs", config.MemberAs})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_Confederation_State
// State information relating to the BGP confederations
type Bgp_Global_Confederation_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When this leaf is set to true it indicates that the local-AS is part of a
    // BGP confederation. The type is bool.
    Enabled interface{}

    // Confederation identifier for the autonomous system. The type is interface{}
    // with range: 0..4294967295.
    Identifier interface{}

    // Remote autonomous systems that are to be treated as part of the local
    // confederation. The type is slice of interface{} with range: 0..4294967295.
    MemberAs []interface{}
}

func (state *Bgp_Global_Confederation_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "confederation"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/confederation/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", state.Identifier})
    state.EntityData.Leafs.Append("member-as", types.YLeaf{"MemberAs", state.MemberAs})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_GracefulRestart
// Parameters relating the graceful restart mechanism for BGP
type Bgp_Global_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to graceful-restart.
    Config Bgp_Global_GracefulRestart_Config

    // State information associated with graceful-restart.
    State Bgp_Global_GracefulRestart_State
}

func (gracefulRestart *Bgp_Global_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "openconfig"
    gracefulRestart.EntityData.ParentYangName = "global"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("config", types.YChild{"Config", &gracefulRestart.Config})
    gracefulRestart.EntityData.Children.Append("state", types.YChild{"State", &gracefulRestart.State})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// Bgp_Global_GracefulRestart_Config
// Configuration parameters relating to graceful-restart
type Bgp_Global_GracefulRestart_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable the graceful-restart capability. The type is bool.
    Enabled interface{}

    // Estimated time (in seconds) for the local BGP speaker to restart a session.
    // This value is advertise in the graceful restart BGP capability.  This is a
    // 12-bit value, referred to as Restart Time in RFC4724.  Per RFC4724, the
    // suggested default value is <= the hold-time value. The type is interface{}
    // with range: 0..4096.
    RestartTime interface{}

    // An upper-bound on the time thate stale routes will be retained by a router
    // after a session is restarted. If an End-of-RIB (EOR) marker is received
    // prior to this timer expiring stale-routes will be flushed upon its receipt
    // - if no EOR is received, then when this timer expires stale paths will be
    // purged. This timer is referred to as the Selection_Deferral_Timer in
    // RFC4724. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    StaleRoutesTime interface{}

    // Enable graceful-restart in helper mode only. When this leaf is set, the
    // local system does not retain forwarding its own state during a restart, but
    // supports procedures for the receiving speaker, as defined in RFC4724. The
    // type is bool.
    HelperOnly interface{}
}

func (config *Bgp_Global_GracefulRestart_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "graceful-restart"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/graceful-restart/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", config.RestartTime})
    config.EntityData.Leafs.Append("stale-routes-time", types.YLeaf{"StaleRoutesTime", config.StaleRoutesTime})
    config.EntityData.Leafs.Append("helper-only", types.YLeaf{"HelperOnly", config.HelperOnly})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_GracefulRestart_State
// State information associated with graceful-restart
type Bgp_Global_GracefulRestart_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable the graceful-restart capability. The type is bool.
    Enabled interface{}

    // Estimated time (in seconds) for the local BGP speaker to restart a session.
    // This value is advertise in the graceful restart BGP capability.  This is a
    // 12-bit value, referred to as Restart Time in RFC4724.  Per RFC4724, the
    // suggested default value is <= the hold-time value. The type is interface{}
    // with range: 0..4096.
    RestartTime interface{}

    // An upper-bound on the time thate stale routes will be retained by a router
    // after a session is restarted. If an End-of-RIB (EOR) marker is received
    // prior to this timer expiring stale-routes will be flushed upon its receipt
    // - if no EOR is received, then when this timer expires stale paths will be
    // purged. This timer is referred to as the Selection_Deferral_Timer in
    // RFC4724. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    StaleRoutesTime interface{}

    // Enable graceful-restart in helper mode only. When this leaf is set, the
    // local system does not retain forwarding its own state during a restart, but
    // supports procedures for the receiving speaker, as defined in RFC4724. The
    // type is bool.
    HelperOnly interface{}
}

func (state *Bgp_Global_GracefulRestart_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "graceful-restart"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/graceful-restart/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", state.RestartTime})
    state.EntityData.Leafs.Append("stale-routes-time", types.YLeaf{"StaleRoutesTime", state.StaleRoutesTime})
    state.EntityData.Leafs.Append("helper-only", types.YLeaf{"HelperOnly", state.HelperOnly})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_UseMultiplePaths
// Parameters related to the use of multiple paths for the
// same NLRI
type Bgp_Global_UseMultiplePaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to multipath.
    Config Bgp_Global_UseMultiplePaths_Config

    // State parameters relating to multipath.
    State Bgp_Global_UseMultiplePaths_State

    // Multipath parameters for eBGP.
    Ebgp Bgp_Global_UseMultiplePaths_Ebgp

    // Multipath parameters for iBGP.
    Ibgp Bgp_Global_UseMultiplePaths_Ibgp
}

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetEntityData() *types.CommonEntityData {
    useMultiplePaths.EntityData.YFilter = useMultiplePaths.YFilter
    useMultiplePaths.EntityData.YangName = "use-multiple-paths"
    useMultiplePaths.EntityData.BundleName = "openconfig"
    useMultiplePaths.EntityData.ParentYangName = "global"
    useMultiplePaths.EntityData.SegmentPath = "use-multiple-paths"
    useMultiplePaths.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/" + useMultiplePaths.EntityData.SegmentPath
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = types.NewOrderedMap()
    useMultiplePaths.EntityData.Children.Append("config", types.YChild{"Config", &useMultiplePaths.Config})
    useMultiplePaths.EntityData.Children.Append("state", types.YChild{"State", &useMultiplePaths.State})
    useMultiplePaths.EntityData.Children.Append("ebgp", types.YChild{"Ebgp", &useMultiplePaths.Ebgp})
    useMultiplePaths.EntityData.Children.Append("ibgp", types.YChild{"Ibgp", &useMultiplePaths.Ibgp})
    useMultiplePaths.EntityData.Leafs = types.NewOrderedMap()

    useMultiplePaths.EntityData.YListKeys = []string {}

    return &(useMultiplePaths.EntityData)
}

// Bgp_Global_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_Global_UseMultiplePaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Global_UseMultiplePaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "use-multiple-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/use-multiple-paths/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_Global_UseMultiplePaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Global_UseMultiplePaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "use-multiple-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/use-multiple-paths/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_UseMultiplePaths_Ebgp
// Multipath parameters for eBGP
type Bgp_Global_UseMultiplePaths_Ebgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_Global_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_Global_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetEntityData() *types.CommonEntityData {
    ebgp.EntityData.YFilter = ebgp.YFilter
    ebgp.EntityData.YangName = "ebgp"
    ebgp.EntityData.BundleName = "openconfig"
    ebgp.EntityData.ParentYangName = "use-multiple-paths"
    ebgp.EntityData.SegmentPath = "ebgp"
    ebgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/use-multiple-paths/" + ebgp.EntityData.SegmentPath
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = types.NewOrderedMap()
    ebgp.EntityData.Children.Append("config", types.YChild{"Config", &ebgp.Config})
    ebgp.EntityData.Children.Append("state", types.YChild{"State", &ebgp.State})
    ebgp.EntityData.Leafs = types.NewOrderedMap()

    ebgp.EntityData.YListKeys = []string {}

    return &(ebgp.EntityData)
}

// Bgp_Global_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_Global_UseMultiplePaths_Ebgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}

    // Maximum number of parallel paths to consider when using BGP multipath. The
    // default is use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ebgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/use-multiple-paths/ebgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs})
    config.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", config.MaximumPaths})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_Global_UseMultiplePaths_Ebgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}

    // Maximum number of parallel paths to consider when using BGP multipath. The
    // default is use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ebgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/use-multiple-paths/ebgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs})
    state.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", state.MaximumPaths})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_UseMultiplePaths_Ibgp
// Multipath parameters for iBGP
type Bgp_Global_UseMultiplePaths_Ibgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to iBGP multipath.
    Config Bgp_Global_UseMultiplePaths_Ibgp_Config

    // State information relating to iBGP multipath.
    State Bgp_Global_UseMultiplePaths_Ibgp_State
}

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetEntityData() *types.CommonEntityData {
    ibgp.EntityData.YFilter = ibgp.YFilter
    ibgp.EntityData.YangName = "ibgp"
    ibgp.EntityData.BundleName = "openconfig"
    ibgp.EntityData.ParentYangName = "use-multiple-paths"
    ibgp.EntityData.SegmentPath = "ibgp"
    ibgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/use-multiple-paths/" + ibgp.EntityData.SegmentPath
    ibgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ibgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ibgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ibgp.EntityData.Children = types.NewOrderedMap()
    ibgp.EntityData.Children.Append("config", types.YChild{"Config", &ibgp.Config})
    ibgp.EntityData.Children.Append("state", types.YChild{"State", &ibgp.State})
    ibgp.EntityData.Leafs = types.NewOrderedMap()

    ibgp.EntityData.YListKeys = []string {}

    return &(ibgp.EntityData)
}

// Bgp_Global_UseMultiplePaths_Ibgp_Config
// Configuration parameters relating to iBGP multipath
type Bgp_Global_UseMultiplePaths_Ibgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ibgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/use-multiple-paths/ibgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", config.MaximumPaths})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_UseMultiplePaths_Ibgp_State
// State information relating to iBGP multipath
type Bgp_Global_UseMultiplePaths_Ibgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ibgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/use-multiple-paths/ibgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", state.MaximumPaths})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_RouteSelectionOptions
// Parameters relating to options for route selection
type Bgp_Global_RouteSelectionOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to route selection options.
    Config Bgp_Global_RouteSelectionOptions_Config

    // State information for the route selection options.
    State Bgp_Global_RouteSelectionOptions_State
}

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetEntityData() *types.CommonEntityData {
    routeSelectionOptions.EntityData.YFilter = routeSelectionOptions.YFilter
    routeSelectionOptions.EntityData.YangName = "route-selection-options"
    routeSelectionOptions.EntityData.BundleName = "openconfig"
    routeSelectionOptions.EntityData.ParentYangName = "global"
    routeSelectionOptions.EntityData.SegmentPath = "route-selection-options"
    routeSelectionOptions.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/" + routeSelectionOptions.EntityData.SegmentPath
    routeSelectionOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeSelectionOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeSelectionOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeSelectionOptions.EntityData.Children = types.NewOrderedMap()
    routeSelectionOptions.EntityData.Children.Append("config", types.YChild{"Config", &routeSelectionOptions.Config})
    routeSelectionOptions.EntityData.Children.Append("state", types.YChild{"State", &routeSelectionOptions.State})
    routeSelectionOptions.EntityData.Leafs = types.NewOrderedMap()

    routeSelectionOptions.EntityData.YListKeys = []string {}

    return &(routeSelectionOptions.EntityData)
}

// Bgp_Global_RouteSelectionOptions_Config
// Configuration parameters relating to route selection
// options
type Bgp_Global_RouteSelectionOptions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Compare multi-exit discriminator (MED) value from different ASes when
    // selecting the best route.  The default behavior is to only compare MEDs for
    // paths received from the same AS. The type is bool. The default value is
    // false.
    AlwaysCompareMed interface{}

    // Ignore the AS path length when selecting the best path. The default is to
    // use the AS path length and prefer paths with shorter length. The type is
    // bool. The default value is false.
    IgnoreAsPathLength interface{}

    // When comparing similar routes received from external BGP peers, use the
    // router-id as a criterion to select the active path. The type is bool. The
    // default value is true.
    ExternalCompareRouterId interface{}

    // Advertise inactive routes to external peers.  The default is to only
    // advertise active routes. The type is bool. The default value is false.
    AdvertiseInactiveRoutes interface{}

    // Flag to enable sending / receiving accumulated IGP attribute in routing
    // updates. The type is bool. The default value is false.
    EnableAigp interface{}

    // Ignore the IGP metric to the next-hop when calculating BGP best-path. The
    // default is to select the route for which the metric to the next-hop is
    // lowest. The type is bool. The default value is false.
    IgnoreNextHopIgpMetric interface{}
}

func (config *Bgp_Global_RouteSelectionOptions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "route-selection-options"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/route-selection-options/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("always-compare-med", types.YLeaf{"AlwaysCompareMed", config.AlwaysCompareMed})
    config.EntityData.Leafs.Append("ignore-as-path-length", types.YLeaf{"IgnoreAsPathLength", config.IgnoreAsPathLength})
    config.EntityData.Leafs.Append("external-compare-router-id", types.YLeaf{"ExternalCompareRouterId", config.ExternalCompareRouterId})
    config.EntityData.Leafs.Append("advertise-inactive-routes", types.YLeaf{"AdvertiseInactiveRoutes", config.AdvertiseInactiveRoutes})
    config.EntityData.Leafs.Append("enable-aigp", types.YLeaf{"EnableAigp", config.EnableAigp})
    config.EntityData.Leafs.Append("ignore-next-hop-igp-metric", types.YLeaf{"IgnoreNextHopIgpMetric", config.IgnoreNextHopIgpMetric})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_RouteSelectionOptions_State
// State information for the route selection options
type Bgp_Global_RouteSelectionOptions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Compare multi-exit discriminator (MED) value from different ASes when
    // selecting the best route.  The default behavior is to only compare MEDs for
    // paths received from the same AS. The type is bool. The default value is
    // false.
    AlwaysCompareMed interface{}

    // Ignore the AS path length when selecting the best path. The default is to
    // use the AS path length and prefer paths with shorter length. The type is
    // bool. The default value is false.
    IgnoreAsPathLength interface{}

    // When comparing similar routes received from external BGP peers, use the
    // router-id as a criterion to select the active path. The type is bool. The
    // default value is true.
    ExternalCompareRouterId interface{}

    // Advertise inactive routes to external peers.  The default is to only
    // advertise active routes. The type is bool. The default value is false.
    AdvertiseInactiveRoutes interface{}

    // Flag to enable sending / receiving accumulated IGP attribute in routing
    // updates. The type is bool. The default value is false.
    EnableAigp interface{}

    // Ignore the IGP metric to the next-hop when calculating BGP best-path. The
    // default is to select the route for which the metric to the next-hop is
    // lowest. The type is bool. The default value is false.
    IgnoreNextHopIgpMetric interface{}
}

func (state *Bgp_Global_RouteSelectionOptions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "route-selection-options"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/route-selection-options/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("always-compare-med", types.YLeaf{"AlwaysCompareMed", state.AlwaysCompareMed})
    state.EntityData.Leafs.Append("ignore-as-path-length", types.YLeaf{"IgnoreAsPathLength", state.IgnoreAsPathLength})
    state.EntityData.Leafs.Append("external-compare-router-id", types.YLeaf{"ExternalCompareRouterId", state.ExternalCompareRouterId})
    state.EntityData.Leafs.Append("advertise-inactive-routes", types.YLeaf{"AdvertiseInactiveRoutes", state.AdvertiseInactiveRoutes})
    state.EntityData.Leafs.Append("enable-aigp", types.YLeaf{"EnableAigp", state.EnableAigp})
    state.EntityData.Leafs.Append("ignore-next-hop-igp-metric", types.YLeaf{"IgnoreNextHopIgpMetric", state.IgnoreNextHopIgpMetric})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis
// Address family specific configuration
type Bgp_Global_AfiSafis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI configuration available for the neighbour or group. The type is
    // slice of Bgp_Global_AfiSafis_AfiSafi.
    AfiSafi []*Bgp_Global_AfiSafis_AfiSafi
}

func (afiSafis *Bgp_Global_AfiSafis) GetEntityData() *types.CommonEntityData {
    afiSafis.EntityData.YFilter = afiSafis.YFilter
    afiSafis.EntityData.YangName = "afi-safis"
    afiSafis.EntityData.BundleName = "openconfig"
    afiSafis.EntityData.ParentYangName = "global"
    afiSafis.EntityData.SegmentPath = "afi-safis"
    afiSafis.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/" + afiSafis.EntityData.SegmentPath
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

// Bgp_Global_AfiSafis_AfiSafi
// AFI,SAFI configuration available for the
// neighbour or group
type Bgp_Global_AfiSafis_AfiSafi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the AFI-SAFI name used as a key for
    // the AFI-SAFI list. The type is one of the following:
    // IPV4UNICASTIPV6UNICASTIPV4LABELEDUNICASTIPV6LABELEDUNICASTL3VPNIPV4UNICASTL3VPNIPV6UNICASTL3VPNIPV4MULTICASTL3VPNIPV6MULTICASTL2VPNVPLSL2VPNEVPN.
    AfiSafiName interface{}

    // Configuration parameters for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_Config

    // State information relating to the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_State

    // Parameters relating to BGP graceful-restart.
    GracefulRestart Bgp_Global_AfiSafis_AfiSafi_GracefulRestart

    // Parameters relating to options for route selection.
    RouteSelectionOptions Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions

    // Parameters related to the use of multiple paths for the same NLRI.
    UseMultiplePaths Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths

    // IPv4 unicast configuration options.
    Ipv4Unicast Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast

    // IPv6 unicast configuration options.
    Ipv6Unicast Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast

    // IPv4 Labeled Unicast configuration options.
    Ipv4LabeledUnicast Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast

    // IPv6 Labeled Unicast configuration options.
    Ipv6LabeledUnicast Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast

    // Unicast IPv4 L3VPN configuration options.
    L3vpnIpv4Unicast Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast

    // Unicast IPv6 L3VPN configuration options.
    L3vpnIpv6Unicast Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast

    // Multicast IPv4 L3VPN configuration options.
    L3vpnIpv4Multicast Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast

    // Multicast IPv6 L3VPN configuration options.
    L3vpnIpv6Multicast Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast

    // BGP-signalled VPLS configuration options.
    L2vpnVpls Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls

    // BGP EVPN configuration options.
    L2vpnEvpn Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn
}

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetEntityData() *types.CommonEntityData {
    afiSafi.EntityData.YFilter = afiSafi.YFilter
    afiSafi.EntityData.YangName = "afi-safi"
    afiSafi.EntityData.BundleName = "openconfig"
    afiSafi.EntityData.ParentYangName = "afi-safis"
    afiSafi.EntityData.SegmentPath = "afi-safi" + types.AddKeyToken(afiSafi.AfiSafiName, "afi-safi-name")
    afiSafi.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/" + afiSafi.EntityData.SegmentPath
    afiSafi.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafi.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafi.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafi.EntityData.Children = types.NewOrderedMap()
    afiSafi.EntityData.Children.Append("config", types.YChild{"Config", &afiSafi.Config})
    afiSafi.EntityData.Children.Append("state", types.YChild{"State", &afiSafi.State})
    afiSafi.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &afiSafi.GracefulRestart})
    afiSafi.EntityData.Children.Append("route-selection-options", types.YChild{"RouteSelectionOptions", &afiSafi.RouteSelectionOptions})
    afiSafi.EntityData.Children.Append("use-multiple-paths", types.YChild{"UseMultiplePaths", &afiSafi.UseMultiplePaths})
    afiSafi.EntityData.Children.Append("ipv4-unicast", types.YChild{"Ipv4Unicast", &afiSafi.Ipv4Unicast})
    afiSafi.EntityData.Children.Append("ipv6-unicast", types.YChild{"Ipv6Unicast", &afiSafi.Ipv6Unicast})
    afiSafi.EntityData.Children.Append("ipv4-labeled-unicast", types.YChild{"Ipv4LabeledUnicast", &afiSafi.Ipv4LabeledUnicast})
    afiSafi.EntityData.Children.Append("ipv6-labeled-unicast", types.YChild{"Ipv6LabeledUnicast", &afiSafi.Ipv6LabeledUnicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv4-unicast", types.YChild{"L3vpnIpv4Unicast", &afiSafi.L3vpnIpv4Unicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv6-unicast", types.YChild{"L3vpnIpv6Unicast", &afiSafi.L3vpnIpv6Unicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv4-multicast", types.YChild{"L3vpnIpv4Multicast", &afiSafi.L3vpnIpv4Multicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv6-multicast", types.YChild{"L3vpnIpv6Multicast", &afiSafi.L3vpnIpv6Multicast})
    afiSafi.EntityData.Children.Append("l2vpn-vpls", types.YChild{"L2vpnVpls", &afiSafi.L2vpnVpls})
    afiSafi.EntityData.Children.Append("l2vpn-evpn", types.YChild{"L2vpnEvpn", &afiSafi.L2vpnEvpn})
    afiSafi.EntityData.Leafs = types.NewOrderedMap()
    afiSafi.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", afiSafi.AfiSafiName})

    afiSafi.EntityData.YListKeys = []string {"AfiSafiName"}

    return &(afiSafi.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Config
// Configuration parameters for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // IPV4UNICASTIPV6UNICASTIPV4LABELEDUNICASTIPV6LABELEDUNICASTL3VPNIPV4UNICASTL3VPNIPV6UNICASTL3VPNIPV4MULTICASTL3VPNIPV6MULTICASTL2VPNVPLSL2VPNEVPN.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "afi-safi"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", config.AfiSafiName})
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_State
// State information relating to the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // IPV4UNICASTIPV6UNICASTIPV4LABELEDUNICASTIPV6LABELEDUNICASTL3VPNIPV4UNICASTL3VPNIPV6UNICASTL3VPNIPV4MULTICASTL3VPNIPV6MULTICASTL2VPNVPLSL2VPNEVPN.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}

    // Total number of BGP paths within the context. The type is interface{} with
    // range: 0..4294967295.
    TotalPaths interface{}

    // Total number of BGP prefixes received within the context. The type is
    // interface{} with range: 0..4294967295.
    TotalPrefixes interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "afi-safi"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", state.AfiSafiName})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("total-paths", types.YLeaf{"TotalPaths", state.TotalPaths})
    state.EntityData.Leafs.Append("total-prefixes", types.YLeaf{"TotalPrefixes", state.TotalPrefixes})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_GracefulRestart
// Parameters relating to BGP graceful-restart
type Bgp_Global_AfiSafis_AfiSafi_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration options for BGP graceful-restart.
    Config Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config

    // State information for BGP graceful-restart.
    State Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State
}

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "openconfig"
    gracefulRestart.EntityData.ParentYangName = "afi-safi"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("config", types.YChild{"Config", &gracefulRestart.Config})
    gracefulRestart.EntityData.Children.Append("state", types.YChild{"State", &gracefulRestart.State})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config
// Configuration options for BGP graceful-restart
type Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "graceful-restart"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/graceful-restart/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State
// State information for BGP graceful-restart
type Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "graceful-restart"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/graceful-restart/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions
// Parameters relating to options for route selection
type Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to route selection options.
    Config Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config

    // State information for the route selection options.
    State Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State
}

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetEntityData() *types.CommonEntityData {
    routeSelectionOptions.EntityData.YFilter = routeSelectionOptions.YFilter
    routeSelectionOptions.EntityData.YangName = "route-selection-options"
    routeSelectionOptions.EntityData.BundleName = "openconfig"
    routeSelectionOptions.EntityData.ParentYangName = "afi-safi"
    routeSelectionOptions.EntityData.SegmentPath = "route-selection-options"
    routeSelectionOptions.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + routeSelectionOptions.EntityData.SegmentPath
    routeSelectionOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeSelectionOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeSelectionOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeSelectionOptions.EntityData.Children = types.NewOrderedMap()
    routeSelectionOptions.EntityData.Children.Append("config", types.YChild{"Config", &routeSelectionOptions.Config})
    routeSelectionOptions.EntityData.Children.Append("state", types.YChild{"State", &routeSelectionOptions.State})
    routeSelectionOptions.EntityData.Leafs = types.NewOrderedMap()

    routeSelectionOptions.EntityData.YListKeys = []string {}

    return &(routeSelectionOptions.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config
// Configuration parameters relating to route selection
// options
type Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Compare multi-exit discriminator (MED) value from different ASes when
    // selecting the best route.  The default behavior is to only compare MEDs for
    // paths received from the same AS. The type is bool. The default value is
    // false.
    AlwaysCompareMed interface{}

    // Ignore the AS path length when selecting the best path. The default is to
    // use the AS path length and prefer paths with shorter length. The type is
    // bool. The default value is false.
    IgnoreAsPathLength interface{}

    // When comparing similar routes received from external BGP peers, use the
    // router-id as a criterion to select the active path. The type is bool. The
    // default value is true.
    ExternalCompareRouterId interface{}

    // Advertise inactive routes to external peers.  The default is to only
    // advertise active routes. The type is bool. The default value is false.
    AdvertiseInactiveRoutes interface{}

    // Flag to enable sending / receiving accumulated IGP attribute in routing
    // updates. The type is bool. The default value is false.
    EnableAigp interface{}

    // Ignore the IGP metric to the next-hop when calculating BGP best-path. The
    // default is to select the route for which the metric to the next-hop is
    // lowest. The type is bool. The default value is false.
    IgnoreNextHopIgpMetric interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "route-selection-options"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/route-selection-options/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("always-compare-med", types.YLeaf{"AlwaysCompareMed", config.AlwaysCompareMed})
    config.EntityData.Leafs.Append("ignore-as-path-length", types.YLeaf{"IgnoreAsPathLength", config.IgnoreAsPathLength})
    config.EntityData.Leafs.Append("external-compare-router-id", types.YLeaf{"ExternalCompareRouterId", config.ExternalCompareRouterId})
    config.EntityData.Leafs.Append("advertise-inactive-routes", types.YLeaf{"AdvertiseInactiveRoutes", config.AdvertiseInactiveRoutes})
    config.EntityData.Leafs.Append("enable-aigp", types.YLeaf{"EnableAigp", config.EnableAigp})
    config.EntityData.Leafs.Append("ignore-next-hop-igp-metric", types.YLeaf{"IgnoreNextHopIgpMetric", config.IgnoreNextHopIgpMetric})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State
// State information for the route selection options
type Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Compare multi-exit discriminator (MED) value from different ASes when
    // selecting the best route.  The default behavior is to only compare MEDs for
    // paths received from the same AS. The type is bool. The default value is
    // false.
    AlwaysCompareMed interface{}

    // Ignore the AS path length when selecting the best path. The default is to
    // use the AS path length and prefer paths with shorter length. The type is
    // bool. The default value is false.
    IgnoreAsPathLength interface{}

    // When comparing similar routes received from external BGP peers, use the
    // router-id as a criterion to select the active path. The type is bool. The
    // default value is true.
    ExternalCompareRouterId interface{}

    // Advertise inactive routes to external peers.  The default is to only
    // advertise active routes. The type is bool. The default value is false.
    AdvertiseInactiveRoutes interface{}

    // Flag to enable sending / receiving accumulated IGP attribute in routing
    // updates. The type is bool. The default value is false.
    EnableAigp interface{}

    // Ignore the IGP metric to the next-hop when calculating BGP best-path. The
    // default is to select the route for which the metric to the next-hop is
    // lowest. The type is bool. The default value is false.
    IgnoreNextHopIgpMetric interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "route-selection-options"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/route-selection-options/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("always-compare-med", types.YLeaf{"AlwaysCompareMed", state.AlwaysCompareMed})
    state.EntityData.Leafs.Append("ignore-as-path-length", types.YLeaf{"IgnoreAsPathLength", state.IgnoreAsPathLength})
    state.EntityData.Leafs.Append("external-compare-router-id", types.YLeaf{"ExternalCompareRouterId", state.ExternalCompareRouterId})
    state.EntityData.Leafs.Append("advertise-inactive-routes", types.YLeaf{"AdvertiseInactiveRoutes", state.AdvertiseInactiveRoutes})
    state.EntityData.Leafs.Append("enable-aigp", types.YLeaf{"EnableAigp", state.EnableAigp})
    state.EntityData.Leafs.Append("ignore-next-hop-igp-metric", types.YLeaf{"IgnoreNextHopIgpMetric", state.IgnoreNextHopIgpMetric})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths
// Parameters related to the use of multiple paths for the
// same NLRI
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to multipath.
    Config Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config

    // State parameters relating to multipath.
    State Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State

    // Multipath parameters for eBGP.
    Ebgp Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp

    // Multipath parameters for iBGP.
    Ibgp Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp
}

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetEntityData() *types.CommonEntityData {
    useMultiplePaths.EntityData.YFilter = useMultiplePaths.YFilter
    useMultiplePaths.EntityData.YangName = "use-multiple-paths"
    useMultiplePaths.EntityData.BundleName = "openconfig"
    useMultiplePaths.EntityData.ParentYangName = "afi-safi"
    useMultiplePaths.EntityData.SegmentPath = "use-multiple-paths"
    useMultiplePaths.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + useMultiplePaths.EntityData.SegmentPath
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = types.NewOrderedMap()
    useMultiplePaths.EntityData.Children.Append("config", types.YChild{"Config", &useMultiplePaths.Config})
    useMultiplePaths.EntityData.Children.Append("state", types.YChild{"State", &useMultiplePaths.State})
    useMultiplePaths.EntityData.Children.Append("ebgp", types.YChild{"Ebgp", &useMultiplePaths.Ebgp})
    useMultiplePaths.EntityData.Children.Append("ibgp", types.YChild{"Ibgp", &useMultiplePaths.Ibgp})
    useMultiplePaths.EntityData.Leafs = types.NewOrderedMap()

    useMultiplePaths.EntityData.YListKeys = []string {}

    return &(useMultiplePaths.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "use-multiple-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/use-multiple-paths/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "use-multiple-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/use-multiple-paths/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp
// Multipath parameters for eBGP
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetEntityData() *types.CommonEntityData {
    ebgp.EntityData.YFilter = ebgp.YFilter
    ebgp.EntityData.YangName = "ebgp"
    ebgp.EntityData.BundleName = "openconfig"
    ebgp.EntityData.ParentYangName = "use-multiple-paths"
    ebgp.EntityData.SegmentPath = "ebgp"
    ebgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/use-multiple-paths/" + ebgp.EntityData.SegmentPath
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = types.NewOrderedMap()
    ebgp.EntityData.Children.Append("config", types.YChild{"Config", &ebgp.Config})
    ebgp.EntityData.Children.Append("state", types.YChild{"State", &ebgp.State})
    ebgp.EntityData.Leafs = types.NewOrderedMap()

    ebgp.EntityData.YListKeys = []string {}

    return &(ebgp.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}

    // Maximum number of parallel paths to consider when using BGP multipath. The
    // default is use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ebgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/use-multiple-paths/ebgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs})
    config.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", config.MaximumPaths})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}

    // Maximum number of parallel paths to consider when using BGP multipath. The
    // default is use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ebgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/use-multiple-paths/ebgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs})
    state.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", state.MaximumPaths})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp
// Multipath parameters for iBGP
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to iBGP multipath.
    Config Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config

    // State information relating to iBGP multipath.
    State Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State
}

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetEntityData() *types.CommonEntityData {
    ibgp.EntityData.YFilter = ibgp.YFilter
    ibgp.EntityData.YangName = "ibgp"
    ibgp.EntityData.BundleName = "openconfig"
    ibgp.EntityData.ParentYangName = "use-multiple-paths"
    ibgp.EntityData.SegmentPath = "ibgp"
    ibgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/use-multiple-paths/" + ibgp.EntityData.SegmentPath
    ibgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ibgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ibgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ibgp.EntityData.Children = types.NewOrderedMap()
    ibgp.EntityData.Children.Append("config", types.YChild{"Config", &ibgp.Config})
    ibgp.EntityData.Children.Append("state", types.YChild{"State", &ibgp.State})
    ibgp.EntityData.Leafs = types.NewOrderedMap()

    ibgp.EntityData.YListKeys = []string {}

    return &(ibgp.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config
// Configuration parameters relating to iBGP multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ibgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/use-multiple-paths/ibgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", config.MaximumPaths})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State
// State information relating to iBGP multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ibgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/use-multiple-paths/ibgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", state.MaximumPaths})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast
// IPv4 unicast configuration options
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State
}

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetEntityData() *types.CommonEntityData {
    ipv4Unicast.EntityData.YFilter = ipv4Unicast.YFilter
    ipv4Unicast.EntityData.YangName = "ipv4-unicast"
    ipv4Unicast.EntityData.BundleName = "openconfig"
    ipv4Unicast.EntityData.ParentYangName = "afi-safi"
    ipv4Unicast.EntityData.SegmentPath = "ipv4-unicast"
    ipv4Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + ipv4Unicast.EntityData.SegmentPath
    ipv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4Unicast.EntityData.Children = types.NewOrderedMap()
    ipv4Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv4Unicast.PrefixLimit})
    ipv4Unicast.EntityData.Children.Append("config", types.YChild{"Config", &ipv4Unicast.Config})
    ipv4Unicast.EntityData.Children.Append("state", types.YChild{"State", &ipv4Unicast.State})
    ipv4Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv4Unicast.EntityData.YListKeys = []string {}

    return &(ipv4Unicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv4-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv4-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv4-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv4-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv4-unicast"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv4-unicast/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv4-unicast"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv4-unicast/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast
// IPv6 unicast configuration options
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State
}

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetEntityData() *types.CommonEntityData {
    ipv6Unicast.EntityData.YFilter = ipv6Unicast.YFilter
    ipv6Unicast.EntityData.YangName = "ipv6-unicast"
    ipv6Unicast.EntityData.BundleName = "openconfig"
    ipv6Unicast.EntityData.ParentYangName = "afi-safi"
    ipv6Unicast.EntityData.SegmentPath = "ipv6-unicast"
    ipv6Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + ipv6Unicast.EntityData.SegmentPath
    ipv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6Unicast.EntityData.Children = types.NewOrderedMap()
    ipv6Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv6Unicast.PrefixLimit})
    ipv6Unicast.EntityData.Children.Append("config", types.YChild{"Config", &ipv6Unicast.Config})
    ipv6Unicast.EntityData.Children.Append("state", types.YChild{"State", &ipv6Unicast.State})
    ipv6Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv6Unicast.EntityData.YListKeys = []string {}

    return &(ipv6Unicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv6-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv6-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv6-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv6-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv6-unicast"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv6-unicast/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv6-unicast"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv6-unicast/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast
// IPv4 Labeled Unicast configuration options
type Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetEntityData() *types.CommonEntityData {
    ipv4LabeledUnicast.EntityData.YFilter = ipv4LabeledUnicast.YFilter
    ipv4LabeledUnicast.EntityData.YangName = "ipv4-labeled-unicast"
    ipv4LabeledUnicast.EntityData.BundleName = "openconfig"
    ipv4LabeledUnicast.EntityData.ParentYangName = "afi-safi"
    ipv4LabeledUnicast.EntityData.SegmentPath = "ipv4-labeled-unicast"
    ipv4LabeledUnicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + ipv4LabeledUnicast.EntityData.SegmentPath
    ipv4LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4LabeledUnicast.EntityData.Children = types.NewOrderedMap()
    ipv4LabeledUnicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv4LabeledUnicast.PrefixLimit})
    ipv4LabeledUnicast.EntityData.Leafs = types.NewOrderedMap()

    ipv4LabeledUnicast.EntityData.YListKeys = []string {}

    return &(ipv4LabeledUnicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv4-labeled-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv4-labeled-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv4-labeled-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv4-labeled-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast
// IPv6 Labeled Unicast configuration options
type Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetEntityData() *types.CommonEntityData {
    ipv6LabeledUnicast.EntityData.YFilter = ipv6LabeledUnicast.YFilter
    ipv6LabeledUnicast.EntityData.YangName = "ipv6-labeled-unicast"
    ipv6LabeledUnicast.EntityData.BundleName = "openconfig"
    ipv6LabeledUnicast.EntityData.ParentYangName = "afi-safi"
    ipv6LabeledUnicast.EntityData.SegmentPath = "ipv6-labeled-unicast"
    ipv6LabeledUnicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + ipv6LabeledUnicast.EntityData.SegmentPath
    ipv6LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6LabeledUnicast.EntityData.Children = types.NewOrderedMap()
    ipv6LabeledUnicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv6LabeledUnicast.PrefixLimit})
    ipv6LabeledUnicast.EntityData.Leafs = types.NewOrderedMap()

    ipv6LabeledUnicast.EntityData.YListKeys = []string {}

    return &(ipv6LabeledUnicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv6-labeled-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv6-labeled-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv6-labeled-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/ipv6-labeled-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast
// Unicast IPv4 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit
}

func (l3vpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv4Unicast.EntityData.YFilter = l3vpnIpv4Unicast.YFilter
    l3vpnIpv4Unicast.EntityData.YangName = "l3vpn-ipv4-unicast"
    l3vpnIpv4Unicast.EntityData.BundleName = "openconfig"
    l3vpnIpv4Unicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv4Unicast.EntityData.SegmentPath = "l3vpn-ipv4-unicast"
    l3vpnIpv4Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + l3vpnIpv4Unicast.EntityData.SegmentPath
    l3vpnIpv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv4Unicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv4Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv4Unicast.PrefixLimit})
    l3vpnIpv4Unicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv4Unicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv4Unicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv4-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv4-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv4-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast
// Unicast IPv6 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit
}

func (l3vpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv6Unicast.EntityData.YFilter = l3vpnIpv6Unicast.YFilter
    l3vpnIpv6Unicast.EntityData.YangName = "l3vpn-ipv6-unicast"
    l3vpnIpv6Unicast.EntityData.BundleName = "openconfig"
    l3vpnIpv6Unicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv6Unicast.EntityData.SegmentPath = "l3vpn-ipv6-unicast"
    l3vpnIpv6Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + l3vpnIpv6Unicast.EntityData.SegmentPath
    l3vpnIpv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv6Unicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv6Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv6Unicast.PrefixLimit})
    l3vpnIpv6Unicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv6Unicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv6Unicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv6-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv6-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv6-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast
// Multicast IPv4 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit
}

func (l3vpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv4Multicast.EntityData.YFilter = l3vpnIpv4Multicast.YFilter
    l3vpnIpv4Multicast.EntityData.YangName = "l3vpn-ipv4-multicast"
    l3vpnIpv4Multicast.EntityData.BundleName = "openconfig"
    l3vpnIpv4Multicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv4Multicast.EntityData.SegmentPath = "l3vpn-ipv4-multicast"
    l3vpnIpv4Multicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + l3vpnIpv4Multicast.EntityData.SegmentPath
    l3vpnIpv4Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv4Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv4Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv4Multicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv4Multicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv4Multicast.PrefixLimit})
    l3vpnIpv4Multicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv4Multicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv4Multicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv4-multicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv4-multicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv4-multicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast
// Multicast IPv6 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit
}

func (l3vpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv6Multicast.EntityData.YFilter = l3vpnIpv6Multicast.YFilter
    l3vpnIpv6Multicast.EntityData.YangName = "l3vpn-ipv6-multicast"
    l3vpnIpv6Multicast.EntityData.BundleName = "openconfig"
    l3vpnIpv6Multicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv6Multicast.EntityData.SegmentPath = "l3vpn-ipv6-multicast"
    l3vpnIpv6Multicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + l3vpnIpv6Multicast.EntityData.SegmentPath
    l3vpnIpv6Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv6Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv6Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv6Multicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv6Multicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv6Multicast.PrefixLimit})
    l3vpnIpv6Multicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv6Multicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv6Multicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv6-multicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv6-multicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l3vpn-ipv6-multicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls
// BGP-signalled VPLS configuration options
type Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit
}

func (l2vpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls) GetEntityData() *types.CommonEntityData {
    l2vpnVpls.EntityData.YFilter = l2vpnVpls.YFilter
    l2vpnVpls.EntityData.YangName = "l2vpn-vpls"
    l2vpnVpls.EntityData.BundleName = "openconfig"
    l2vpnVpls.EntityData.ParentYangName = "afi-safi"
    l2vpnVpls.EntityData.SegmentPath = "l2vpn-vpls"
    l2vpnVpls.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + l2vpnVpls.EntityData.SegmentPath
    l2vpnVpls.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2vpnVpls.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2vpnVpls.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2vpnVpls.EntityData.Children = types.NewOrderedMap()
    l2vpnVpls.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l2vpnVpls.PrefixLimit})
    l2vpnVpls.EntityData.Leafs = types.NewOrderedMap()

    l2vpnVpls.EntityData.YListKeys = []string {}

    return &(l2vpnVpls.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-vpls"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l2vpn-vpls/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l2vpn-vpls/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l2vpn-vpls/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn
// BGP EVPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit
}

func (l2vpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn) GetEntityData() *types.CommonEntityData {
    l2vpnEvpn.EntityData.YFilter = l2vpnEvpn.YFilter
    l2vpnEvpn.EntityData.YangName = "l2vpn-evpn"
    l2vpnEvpn.EntityData.BundleName = "openconfig"
    l2vpnEvpn.EntityData.ParentYangName = "afi-safi"
    l2vpnEvpn.EntityData.SegmentPath = "l2vpn-evpn"
    l2vpnEvpn.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/" + l2vpnEvpn.EntityData.SegmentPath
    l2vpnEvpn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2vpnEvpn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2vpnEvpn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2vpnEvpn.EntityData.Children = types.NewOrderedMap()
    l2vpnEvpn.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l2vpnEvpn.PrefixLimit})
    l2vpnEvpn.EntityData.Leafs = types.NewOrderedMap()

    l2vpnEvpn.EntityData.YListKeys = []string {}

    return &(l2vpnEvpn.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-evpn"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l2vpn-evpn/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l2vpn-evpn/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/afi-safis/afi-safi/l2vpn-evpn/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Global_DynamicNeighborPrefixes
// A list of IP prefixes from which the system should:
//  - Accept connections to the BGP daemon
//  - Dynamically configure a BGP neighbor corresponding to the
//    source address of the remote system, using the parameters
//    of the specified peer-group.
// For such neighbors, an entry within the neighbor list should
// be created, indicating that the peer was dynamically
// configured, and referencing the peer-group from which the
// configuration was derived.
type Bgp_Global_DynamicNeighborPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An individual prefix from which dynamic neighbor connections are allowed.
    // The type is slice of
    // Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix.
    DynamicNeighborPrefix []*Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix
}

func (dynamicNeighborPrefixes *Bgp_Global_DynamicNeighborPrefixes) GetEntityData() *types.CommonEntityData {
    dynamicNeighborPrefixes.EntityData.YFilter = dynamicNeighborPrefixes.YFilter
    dynamicNeighborPrefixes.EntityData.YangName = "dynamic-neighbor-prefixes"
    dynamicNeighborPrefixes.EntityData.BundleName = "openconfig"
    dynamicNeighborPrefixes.EntityData.ParentYangName = "global"
    dynamicNeighborPrefixes.EntityData.SegmentPath = "dynamic-neighbor-prefixes"
    dynamicNeighborPrefixes.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/" + dynamicNeighborPrefixes.EntityData.SegmentPath
    dynamicNeighborPrefixes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    dynamicNeighborPrefixes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    dynamicNeighborPrefixes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    dynamicNeighborPrefixes.EntityData.Children = types.NewOrderedMap()
    dynamicNeighborPrefixes.EntityData.Children.Append("dynamic-neighbor-prefix", types.YChild{"DynamicNeighborPrefix", nil})
    for i := range dynamicNeighborPrefixes.DynamicNeighborPrefix {
        dynamicNeighborPrefixes.EntityData.Children.Append(types.GetSegmentPath(dynamicNeighborPrefixes.DynamicNeighborPrefix[i]), types.YChild{"DynamicNeighborPrefix", dynamicNeighborPrefixes.DynamicNeighborPrefix[i]})
    }
    dynamicNeighborPrefixes.EntityData.Leafs = types.NewOrderedMap()

    dynamicNeighborPrefixes.EntityData.YListKeys = []string {}

    return &(dynamicNeighborPrefixes.EntityData)
}

// Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix
// An individual prefix from which dynamic neighbor
// connections are allowed.
type Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the IP prefix from which source
    // connections are allowed for the dynamic neighbor group. The type is one of
    // the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$'.
    Prefix interface{}

    // Configuration parameters relating to the source prefix for the dynamic BGP
    // neighbor connections.
    Config Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix_Config

    // Operational state parameters relating to the source prefix for the dynamic
    // BGP neighbor connections.
    State Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix_State
}

func (dynamicNeighborPrefix *Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix) GetEntityData() *types.CommonEntityData {
    dynamicNeighborPrefix.EntityData.YFilter = dynamicNeighborPrefix.YFilter
    dynamicNeighborPrefix.EntityData.YangName = "dynamic-neighbor-prefix"
    dynamicNeighborPrefix.EntityData.BundleName = "openconfig"
    dynamicNeighborPrefix.EntityData.ParentYangName = "dynamic-neighbor-prefixes"
    dynamicNeighborPrefix.EntityData.SegmentPath = "dynamic-neighbor-prefix" + types.AddKeyToken(dynamicNeighborPrefix.Prefix, "prefix")
    dynamicNeighborPrefix.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/dynamic-neighbor-prefixes/" + dynamicNeighborPrefix.EntityData.SegmentPath
    dynamicNeighborPrefix.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    dynamicNeighborPrefix.EntityData.NamespaceTable = openconfig.GetNamespaces()
    dynamicNeighborPrefix.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    dynamicNeighborPrefix.EntityData.Children = types.NewOrderedMap()
    dynamicNeighborPrefix.EntityData.Children.Append("config", types.YChild{"Config", &dynamicNeighborPrefix.Config})
    dynamicNeighborPrefix.EntityData.Children.Append("state", types.YChild{"State", &dynamicNeighborPrefix.State})
    dynamicNeighborPrefix.EntityData.Leafs = types.NewOrderedMap()
    dynamicNeighborPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", dynamicNeighborPrefix.Prefix})

    dynamicNeighborPrefix.EntityData.YListKeys = []string {"Prefix"}

    return &(dynamicNeighborPrefix.EntityData)
}

// Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix_Config
// Configuration parameters relating to the source prefix
// for the dynamic BGP neighbor connections.
type Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP prefix within which the source address of the remote BGP speaker
    // must fall to be considered eligible to the dynamically configured. The type
    // is one of the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$'.
    Prefix interface{}

    // The peer-group within which the dynamic neighbor will be configured.  The
    // configuration parameters used for the dynamic neighbor are those specified
    // within the referenced peer group. The type is string. Refers to
    // bgp.Bgp_PeerGroups_PeerGroup_Config_PeerGroupName
    PeerGroup interface{}
}

func (config *Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "dynamic-neighbor-prefix"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/dynamic-neighbor-prefixes/dynamic-neighbor-prefix/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", config.Prefix})
    config.EntityData.Leafs.Append("peer-group", types.YLeaf{"PeerGroup", config.PeerGroup})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix_State
// Operational state parameters relating to the source
// prefix for the dynamic BGP neighbor connections.
type Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IP prefix within which the source address of the remote BGP speaker
    // must fall to be considered eligible to the dynamically configured. The type
    // is one of the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/(12[0-8]|1[0-1][0-9]|[1-9][0-9]|[0-9])$'.
    Prefix interface{}

    // The peer-group within which the dynamic neighbor will be configured.  The
    // configuration parameters used for the dynamic neighbor are those specified
    // within the referenced peer group. The type is string. Refers to
    // bgp.Bgp_PeerGroups_PeerGroup_Config_PeerGroupName
    PeerGroup interface{}
}

func (state *Bgp_Global_DynamicNeighborPrefixes_DynamicNeighborPrefix_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "dynamic-neighbor-prefix"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/global/dynamic-neighbor-prefixes/dynamic-neighbor-prefix/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", state.Prefix})
    state.EntityData.Leafs.Append("peer-group", types.YLeaf{"PeerGroup", state.PeerGroup})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors
// Configuration for BGP neighbors
type Bgp_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP neighbors configured on the local system, uniquely identified
    // by peer IPv[46] address. The type is slice of Bgp_Neighbors_Neighbor.
    Neighbor []*Bgp_Neighbors_Neighbor
}

func (neighbors *Bgp_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "bgp"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "openconfig-bgp:bgp/" + neighbors.EntityData.SegmentPath
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

// Bgp_Neighbors_Neighbor
// List of BGP neighbors configured on the local system,
// uniquely identified by peer IPv[46] address
type Bgp_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the address of the BGP neighbor used
    // as a key in the neighbor list. The type is one of the following types:
    // string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.
    NeighborAddress interface{}

    // Configuration parameters relating to the BGP neighbor or group.
    Config Bgp_Neighbors_Neighbor_Config

    // State information relating to the BGP neighbor.
    State Bgp_Neighbors_Neighbor_State

    // Timers related to a BGP neighbor.
    Timers Bgp_Neighbors_Neighbor_Timers

    // Transport session parameters for the BGP neighbor.
    Transport Bgp_Neighbors_Neighbor_Transport

    // Error handling parameters used for the BGP neighbor or group.
    ErrorHandling Bgp_Neighbors_Neighbor_ErrorHandling

    // Parameters relating the graceful restart mechanism for BGP.
    GracefulRestart Bgp_Neighbors_Neighbor_GracefulRestart

    // Logging options for events related to the BGP neighbor or group.
    LoggingOptions Bgp_Neighbors_Neighbor_LoggingOptions

    // eBGP multi-hop parameters for the BGPgroup.
    EbgpMultihop Bgp_Neighbors_Neighbor_EbgpMultihop

    // Route reflector parameters for the BGPgroup.
    RouteReflector Bgp_Neighbors_Neighbor_RouteReflector

    // AS_PATH manipulation parameters for the BGP neighbor or group.
    AsPathOptions Bgp_Neighbors_Neighbor_AsPathOptions

    // Parameters relating to the advertisement and receipt of multiple paths for
    // a single NLRI (add-paths).
    AddPaths Bgp_Neighbors_Neighbor_AddPaths

    // Parameters related to the use of multiple-paths for the same NLRI when they
    // are received only from this neighbor.
    UseMultiplePaths Bgp_Neighbors_Neighbor_UseMultiplePaths

    // Anchor point for routing policies in the model. Import and export policies
    // are with respect to the local routing table, i.e., export (send) and import
    // (receive), depending on the context.
    ApplyPolicy Bgp_Neighbors_Neighbor_ApplyPolicy

    // Per-address-family configuration parameters associated with the neighbor.
    AfiSafis Bgp_Neighbors_Neighbor_AfiSafis
}

func (neighbor *Bgp_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborAddress, "neighbor-address")
    neighbor.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("config", types.YChild{"Config", &neighbor.Config})
    neighbor.EntityData.Children.Append("state", types.YChild{"State", &neighbor.State})
    neighbor.EntityData.Children.Append("timers", types.YChild{"Timers", &neighbor.Timers})
    neighbor.EntityData.Children.Append("transport", types.YChild{"Transport", &neighbor.Transport})
    neighbor.EntityData.Children.Append("error-handling", types.YChild{"ErrorHandling", &neighbor.ErrorHandling})
    neighbor.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &neighbor.GracefulRestart})
    neighbor.EntityData.Children.Append("logging-options", types.YChild{"LoggingOptions", &neighbor.LoggingOptions})
    neighbor.EntityData.Children.Append("ebgp-multihop", types.YChild{"EbgpMultihop", &neighbor.EbgpMultihop})
    neighbor.EntityData.Children.Append("route-reflector", types.YChild{"RouteReflector", &neighbor.RouteReflector})
    neighbor.EntityData.Children.Append("as-path-options", types.YChild{"AsPathOptions", &neighbor.AsPathOptions})
    neighbor.EntityData.Children.Append("add-paths", types.YChild{"AddPaths", &neighbor.AddPaths})
    neighbor.EntityData.Children.Append("use-multiple-paths", types.YChild{"UseMultiplePaths", &neighbor.UseMultiplePaths})
    neighbor.EntityData.Children.Append("apply-policy", types.YChild{"ApplyPolicy", &neighbor.ApplyPolicy})
    neighbor.EntityData.Children.Append("afi-safis", types.YChild{"AfiSafis", &neighbor.AfiSafis})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighbor.NeighborAddress})

    neighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(neighbor.EntityData)
}

// Bgp_Neighbors_Neighbor_Config
// Configuration parameters relating to the BGP neighbor or
// group
type Bgp_Neighbors_Neighbor_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The peer-group with which this neighbor is associated. The type is string.
    // Refers to bgp.Bgp_PeerGroups_PeerGroup_PeerGroupName
    PeerGroup interface{}

    // Address of the BGP peer, either in IPv4 or IPv6. The type is one of the
    // following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.
    NeighborAddress interface{}

    // Whether the BGP peer is enabled. In cases where the enabled leaf is set to
    // false, the local system should not initiate connections to the neighbor,
    // and should not respond to TCP connections attempts from the neighbor. If
    // the state of the BGP session is ESTABLISHED at the time that this leaf is
    // set to false, the BGP session should be ceased. The type is bool. The
    // default value is true.
    Enabled interface{}

    // AS number of the peer. The type is interface{} with range: 0..4294967295.
    PeerAs interface{}

    // The local autonomous system number that is to be used when establishing
    // sessions with the remote peer or peer group, if this differs from the
    // global BGP router autonomous system number. The type is interface{} with
    // range: 0..4294967295.
    LocalAs interface{}

    // Explicitly designate the peer or peer group as internal (iBGP) or external
    // (eBGP). The type is PeerType.
    PeerType interface{}

    // Configures an MD5 authentication password for use with neighboring devices.
    // The type is string.
    AuthPassword interface{}

    // Remove private AS numbers from updates sent to peers - when this leaf is
    // not specified, the AS_PATH attribute should be sent to the peer unchanged.
    // The type is one of the following: PRIVATEASREMOVEALLPRIVATEASREPLACEALL.
    RemovePrivateAs interface{}

    // Enable route flap damping. The type is bool. The default value is false.
    RouteFlapDamping interface{}

    // Specify which types of community should be sent to the neighbor or group.
    // The default is to not send the community attribute. The type is
    // CommunityType. The default value is NONE.
    SendCommunity interface{}

    // An optional textual description (intended primarily for use with a peer or
    // group. The type is string.
    Description interface{}
}

func (config *Bgp_Neighbors_Neighbor_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "neighbor"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("peer-group", types.YLeaf{"PeerGroup", config.PeerGroup})
    config.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", config.NeighborAddress})
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("peer-as", types.YLeaf{"PeerAs", config.PeerAs})
    config.EntityData.Leafs.Append("local-as", types.YLeaf{"LocalAs", config.LocalAs})
    config.EntityData.Leafs.Append("peer-type", types.YLeaf{"PeerType", config.PeerType})
    config.EntityData.Leafs.Append("auth-password", types.YLeaf{"AuthPassword", config.AuthPassword})
    config.EntityData.Leafs.Append("remove-private-as", types.YLeaf{"RemovePrivateAs", config.RemovePrivateAs})
    config.EntityData.Leafs.Append("route-flap-damping", types.YLeaf{"RouteFlapDamping", config.RouteFlapDamping})
    config.EntityData.Leafs.Append("send-community", types.YLeaf{"SendCommunity", config.SendCommunity})
    config.EntityData.Leafs.Append("description", types.YLeaf{"Description", config.Description})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_State
// State information relating to the BGP neighbor
type Bgp_Neighbors_Neighbor_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The peer-group with which this neighbor is associated. The type is string.
    // Refers to bgp.Bgp_PeerGroups_PeerGroup_PeerGroupName
    PeerGroup interface{}

    // Address of the BGP peer, either in IPv4 or IPv6. The type is one of the
    // following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.
    NeighborAddress interface{}

    // Whether the BGP peer is enabled. In cases where the enabled leaf is set to
    // false, the local system should not initiate connections to the neighbor,
    // and should not respond to TCP connections attempts from the neighbor. If
    // the state of the BGP session is ESTABLISHED at the time that this leaf is
    // set to false, the BGP session should be ceased. The type is bool. The
    // default value is true.
    Enabled interface{}

    // AS number of the peer. The type is interface{} with range: 0..4294967295.
    PeerAs interface{}

    // The local autonomous system number that is to be used when establishing
    // sessions with the remote peer or peer group, if this differs from the
    // global BGP router autonomous system number. The type is interface{} with
    // range: 0..4294967295.
    LocalAs interface{}

    // Explicitly designate the peer or peer group as internal (iBGP) or external
    // (eBGP). The type is PeerType.
    PeerType interface{}

    // Configures an MD5 authentication password for use with neighboring devices.
    // The type is string.
    AuthPassword interface{}

    // Remove private AS numbers from updates sent to peers - when this leaf is
    // not specified, the AS_PATH attribute should be sent to the peer unchanged.
    // The type is one of the following: PRIVATEASREMOVEALLPRIVATEASREPLACEALL.
    RemovePrivateAs interface{}

    // Enable route flap damping. The type is bool. The default value is false.
    RouteFlapDamping interface{}

    // Specify which types of community should be sent to the neighbor or group.
    // The default is to not send the community attribute. The type is
    // CommunityType. The default value is NONE.
    SendCommunity interface{}

    // An optional textual description (intended primarily for use with a peer or
    // group. The type is string.
    Description interface{}

    // Operational state of the BGP peer. The type is SessionState.
    SessionState interface{}

    // This timestamp indicates the time that the BGP session last transitioned in
    // or out of the Established state.  The value is the timestamp in seconds
    // relative to the Unix Epoch (Jan 1, 1970 00:00:00 UTC).  The BGP session
    // uptime can be computed by clients as the difference between this value and
    // the current time in UTC (assuming the session is in the ESTABLISHED state,
    // per the session-state leaf). The type is interface{} with range:
    // 0..18446744073709551615.
    LastEstablished interface{}

    // Number of transitions to the Established state for the neighbor session. 
    // This value is analogous to the bgpPeerFsmEstablishedTransitions object from
    // the standard BGP-4 MIB. The type is interface{} with range:
    // 0..18446744073709551615.
    EstablishedTransitions interface{}

    // BGP capabilities negotiated as supported with the peer. The type is slice
    // of ['MPBGP', 'ROUTEREFRESH', 'ASN32', 'GRACEFULRESTART', 'ADDPATHS'].
    SupportedCapabilities []interface{}

    // When this leaf is set to true, the peer was configured dynamically due to
    // an inbound connection request from a specified source prefix within a
    // dynamic-neighbor-prefix. The type is bool. The default value is false.
    DynamicallyConfigured interface{}

    // Counters for BGP messages sent and received from the neighbor.
    Messages Bgp_Neighbors_Neighbor_State_Messages

    // Counters related to queued messages associated with the BGP neighbor.
    Queues Bgp_Neighbors_Neighbor_State_Queues
}

func (state *Bgp_Neighbors_Neighbor_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "neighbor"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("messages", types.YChild{"Messages", &state.Messages})
    state.EntityData.Children.Append("queues", types.YChild{"Queues", &state.Queues})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("peer-group", types.YLeaf{"PeerGroup", state.PeerGroup})
    state.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", state.NeighborAddress})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("peer-as", types.YLeaf{"PeerAs", state.PeerAs})
    state.EntityData.Leafs.Append("local-as", types.YLeaf{"LocalAs", state.LocalAs})
    state.EntityData.Leafs.Append("peer-type", types.YLeaf{"PeerType", state.PeerType})
    state.EntityData.Leafs.Append("auth-password", types.YLeaf{"AuthPassword", state.AuthPassword})
    state.EntityData.Leafs.Append("remove-private-as", types.YLeaf{"RemovePrivateAs", state.RemovePrivateAs})
    state.EntityData.Leafs.Append("route-flap-damping", types.YLeaf{"RouteFlapDamping", state.RouteFlapDamping})
    state.EntityData.Leafs.Append("send-community", types.YLeaf{"SendCommunity", state.SendCommunity})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("session-state", types.YLeaf{"SessionState", state.SessionState})
    state.EntityData.Leafs.Append("last-established", types.YLeaf{"LastEstablished", state.LastEstablished})
    state.EntityData.Leafs.Append("established-transitions", types.YLeaf{"EstablishedTransitions", state.EstablishedTransitions})
    state.EntityData.Leafs.Append("supported-capabilities", types.YLeaf{"SupportedCapabilities", state.SupportedCapabilities})
    state.EntityData.Leafs.Append("dynamically-configured", types.YLeaf{"DynamicallyConfigured", state.DynamicallyConfigured})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_State_Messages
// Counters for BGP messages sent and received from the
// neighbor
type Bgp_Neighbors_Neighbor_State_Messages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Counters relating to BGP messages sent to the neighbor.
    Sent Bgp_Neighbors_Neighbor_State_Messages_Sent

    // Counters for BGP messages received from the neighbor.
    Received Bgp_Neighbors_Neighbor_State_Messages_Received
}

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetEntityData() *types.CommonEntityData {
    messages.EntityData.YFilter = messages.YFilter
    messages.EntityData.YangName = "messages"
    messages.EntityData.BundleName = "openconfig"
    messages.EntityData.ParentYangName = "state"
    messages.EntityData.SegmentPath = "messages"
    messages.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/state/" + messages.EntityData.SegmentPath
    messages.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    messages.EntityData.NamespaceTable = openconfig.GetNamespaces()
    messages.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    messages.EntityData.Children = types.NewOrderedMap()
    messages.EntityData.Children.Append("sent", types.YChild{"Sent", &messages.Sent})
    messages.EntityData.Children.Append("received", types.YChild{"Received", &messages.Received})
    messages.EntityData.Leafs = types.NewOrderedMap()

    messages.EntityData.YListKeys = []string {}

    return &(messages.EntityData)
}

// Bgp_Neighbors_Neighbor_State_Messages_Sent
// Counters relating to BGP messages sent to the neighbor
type Bgp_Neighbors_Neighbor_State_Messages_Sent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of BGP UPDATE messages announcing, withdrawing or modifying paths
    // exchanged. The type is interface{} with range: 0..18446744073709551615.
    UPDATE interface{}

    // Number of BGP NOTIFICATION messages indicating an error condition has
    // occurred exchanged. The type is interface{} with range:
    // 0..18446744073709551615.
    NOTIFICATION interface{}
}

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetEntityData() *types.CommonEntityData {
    sent.EntityData.YFilter = sent.YFilter
    sent.EntityData.YangName = "sent"
    sent.EntityData.BundleName = "openconfig"
    sent.EntityData.ParentYangName = "messages"
    sent.EntityData.SegmentPath = "sent"
    sent.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/state/messages/" + sent.EntityData.SegmentPath
    sent.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sent.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sent.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sent.EntityData.Children = types.NewOrderedMap()
    sent.EntityData.Leafs = types.NewOrderedMap()
    sent.EntityData.Leafs.Append("UPDATE", types.YLeaf{"UPDATE", sent.UPDATE})
    sent.EntityData.Leafs.Append("NOTIFICATION", types.YLeaf{"NOTIFICATION", sent.NOTIFICATION})

    sent.EntityData.YListKeys = []string {}

    return &(sent.EntityData)
}

// Bgp_Neighbors_Neighbor_State_Messages_Received
// Counters for BGP messages received from the neighbor
type Bgp_Neighbors_Neighbor_State_Messages_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of BGP UPDATE messages announcing, withdrawing or modifying paths
    // exchanged. The type is interface{} with range: 0..18446744073709551615.
    UPDATE interface{}

    // Number of BGP NOTIFICATION messages indicating an error condition has
    // occurred exchanged. The type is interface{} with range:
    // 0..18446744073709551615.
    NOTIFICATION interface{}
}

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "openconfig"
    received.EntityData.ParentYangName = "messages"
    received.EntityData.SegmentPath = "received"
    received.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/state/messages/" + received.EntityData.SegmentPath
    received.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    received.EntityData.NamespaceTable = openconfig.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    received.EntityData.Children = types.NewOrderedMap()
    received.EntityData.Leafs = types.NewOrderedMap()
    received.EntityData.Leafs.Append("UPDATE", types.YLeaf{"UPDATE", received.UPDATE})
    received.EntityData.Leafs.Append("NOTIFICATION", types.YLeaf{"NOTIFICATION", received.NOTIFICATION})

    received.EntityData.YListKeys = []string {}

    return &(received.EntityData)
}

// Bgp_Neighbors_Neighbor_State_Queues
// Counters related to queued messages associated with the
// BGP neighbor
type Bgp_Neighbors_Neighbor_State_Queues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of messages received from the peer currently queued. The type is
    // interface{} with range: 0..4294967295.
    Input interface{}

    // The number of messages queued to be sent to the peer. The type is
    // interface{} with range: 0..4294967295.
    Output interface{}
}

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetEntityData() *types.CommonEntityData {
    queues.EntityData.YFilter = queues.YFilter
    queues.EntityData.YangName = "queues"
    queues.EntityData.BundleName = "openconfig"
    queues.EntityData.ParentYangName = "state"
    queues.EntityData.SegmentPath = "queues"
    queues.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/state/" + queues.EntityData.SegmentPath
    queues.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    queues.EntityData.NamespaceTable = openconfig.GetNamespaces()
    queues.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    queues.EntityData.Children = types.NewOrderedMap()
    queues.EntityData.Leafs = types.NewOrderedMap()
    queues.EntityData.Leafs.Append("input", types.YLeaf{"Input", queues.Input})
    queues.EntityData.Leafs.Append("output", types.YLeaf{"Output", queues.Output})

    queues.EntityData.YListKeys = []string {}

    return &(queues.EntityData)
}

// Bgp_Neighbors_Neighbor_State_SessionState represents Operational state of the BGP peer
type Bgp_Neighbors_Neighbor_State_SessionState string

const (
    // neighbor is down, and in the Idle state of the
    // FSM
    Bgp_Neighbors_Neighbor_State_SessionState_IDLE Bgp_Neighbors_Neighbor_State_SessionState = "IDLE"

    // neighbor is down, and the session is waiting for
    // the underlying transport session to be established
    Bgp_Neighbors_Neighbor_State_SessionState_CONNECT Bgp_Neighbors_Neighbor_State_SessionState = "CONNECT"

    // neighbor is down, and the local system is awaiting
    // a conncetion from the remote peer
    Bgp_Neighbors_Neighbor_State_SessionState_ACTIVE Bgp_Neighbors_Neighbor_State_SessionState = "ACTIVE"

    // neighbor is in the process of being established.
    // The local system has sent an OPEN message
    Bgp_Neighbors_Neighbor_State_SessionState_OPENSENT Bgp_Neighbors_Neighbor_State_SessionState = "OPENSENT"

    // neighbor is in the process of being established.
    // The local system is awaiting a NOTIFICATION or
    // KEEPALIVE message
    Bgp_Neighbors_Neighbor_State_SessionState_OPENCONFIRM Bgp_Neighbors_Neighbor_State_SessionState = "OPENCONFIRM"

    // neighbor is up - the BGP session with the peer is
    // established
    Bgp_Neighbors_Neighbor_State_SessionState_ESTABLISHED Bgp_Neighbors_Neighbor_State_SessionState = "ESTABLISHED"
)

// Bgp_Neighbors_Neighbor_Timers
// Timers related to a BGP neighbor
type Bgp_Neighbors_Neighbor_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to timers used for the BGP neighbor.
    Config Bgp_Neighbors_Neighbor_Timers_Config

    // State information relating to the timers used for the BGP neighbor.
    State Bgp_Neighbors_Neighbor_Timers_State
}

func (timers *Bgp_Neighbors_Neighbor_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "openconfig"
    timers.EntityData.ParentYangName = "neighbor"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + timers.EntityData.SegmentPath
    timers.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    timers.EntityData.NamespaceTable = openconfig.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("config", types.YChild{"Config", &timers.Config})
    timers.EntityData.Children.Append("state", types.YChild{"State", &timers.State})
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// Bgp_Neighbors_Neighbor_Timers_Config
// Configuration parameters relating to timers used for the
// BGP neighbor
type Bgp_Neighbors_Neighbor_Timers_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time interval in seconds between attempts to establish a session with the
    // peer. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 30.
    ConnectRetry interface{}

    // Time interval in seconds that a BGP session will be considered active in
    // the absence of keepalive or other messages from the peer.  The hold-time is
    // typically set to 3x the keepalive-interval. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 90.
    HoldTime interface{}

    // Time interval in seconds between transmission of keepalive messages to the
    // neighbor.  Typically set to 1/3 the hold-time. The type is string with
    // range: -92233720368547758.08..92233720368547758.07. The default value is
    // 30.
    KeepaliveInterval interface{}

    // Minimum time which must elapse between subsequent UPDATE messages relating
    // to a common set of NLRI being transmitted to a peer. This timer is referred
    // to as MinRouteAdvertisementIntervalTimer by RFC 4721 and serves to reduce
    // the number of UPDATE messages transmitted when a particular set of NLRI
    // exhibit instability. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 30.
    MinimumAdvertisementInterval interface{}
}

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "timers"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/timers/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("connect-retry", types.YLeaf{"ConnectRetry", config.ConnectRetry})
    config.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", config.HoldTime})
    config.EntityData.Leafs.Append("keepalive-interval", types.YLeaf{"KeepaliveInterval", config.KeepaliveInterval})
    config.EntityData.Leafs.Append("minimum-advertisement-interval", types.YLeaf{"MinimumAdvertisementInterval", config.MinimumAdvertisementInterval})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_Timers_State
// State information relating to the timers used for the BGP
// neighbor
type Bgp_Neighbors_Neighbor_Timers_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time interval in seconds between attempts to establish a session with the
    // peer. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 30.
    ConnectRetry interface{}

    // Time interval in seconds that a BGP session will be considered active in
    // the absence of keepalive or other messages from the peer.  The hold-time is
    // typically set to 3x the keepalive-interval. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 90.
    HoldTime interface{}

    // Time interval in seconds between transmission of keepalive messages to the
    // neighbor.  Typically set to 1/3 the hold-time. The type is string with
    // range: -92233720368547758.08..92233720368547758.07. The default value is
    // 30.
    KeepaliveInterval interface{}

    // Minimum time which must elapse between subsequent UPDATE messages relating
    // to a common set of NLRI being transmitted to a peer. This timer is referred
    // to as MinRouteAdvertisementIntervalTimer by RFC 4721 and serves to reduce
    // the number of UPDATE messages transmitted when a particular set of NLRI
    // exhibit instability. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 30.
    MinimumAdvertisementInterval interface{}

    // The negotiated hold-time for the BGP session. The type is string with
    // range: -92233720368547758.08..92233720368547758.07.
    NegotiatedHoldTime interface{}
}

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "timers"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/timers/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("connect-retry", types.YLeaf{"ConnectRetry", state.ConnectRetry})
    state.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", state.HoldTime})
    state.EntityData.Leafs.Append("keepalive-interval", types.YLeaf{"KeepaliveInterval", state.KeepaliveInterval})
    state.EntityData.Leafs.Append("minimum-advertisement-interval", types.YLeaf{"MinimumAdvertisementInterval", state.MinimumAdvertisementInterval})
    state.EntityData.Leafs.Append("negotiated-hold-time", types.YLeaf{"NegotiatedHoldTime", state.NegotiatedHoldTime})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_Transport
// Transport session parameters for the BGP neighbor
type Bgp_Neighbors_Neighbor_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the transport session(s) used for the
    // BGP neighbor.
    Config Bgp_Neighbors_Neighbor_Transport_Config

    // State information relating to the transport session(s) used for the BGP
    // neighbor.
    State Bgp_Neighbors_Neighbor_Transport_State
}

func (transport *Bgp_Neighbors_Neighbor_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "openconfig"
    transport.EntityData.ParentYangName = "neighbor"
    transport.EntityData.SegmentPath = "transport"
    transport.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + transport.EntityData.SegmentPath
    transport.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    transport.EntityData.NamespaceTable = openconfig.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    transport.EntityData.Children = types.NewOrderedMap()
    transport.EntityData.Children.Append("config", types.YChild{"Config", &transport.Config})
    transport.EntityData.Children.Append("state", types.YChild{"State", &transport.State})
    transport.EntityData.Leafs = types.NewOrderedMap()

    transport.EntityData.YListKeys = []string {}

    return &(transport.EntityData)
}

// Bgp_Neighbors_Neighbor_Transport_Config
// Configuration parameters relating to the transport
// session(s) used for the BGP neighbor
type Bgp_Neighbors_Neighbor_Transport_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets the max segment size for BGP TCP sessions. The type is interface{}
    // with range: 0..65535.
    TcpMss interface{}

    // Turns path mtu discovery for BGP TCP sessions on (true) or off (false). The
    // type is bool. The default value is false.
    MtuDiscovery interface{}

    // Wait for peers to issue requests to open a BGP session, rather than
    // initiating sessions from the local router. The type is bool. The default
    // value is false.
    PassiveMode interface{}

    // Set the local IP (either IPv4 or IPv6) address to use for the session when
    // sending BGP update messages.  This may be expressed as either an IP address
    // or reference to the name of an interface. The type is one of the following
    // types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.,
    // or string.
    LocalAddress interface{}
}

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "transport"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/transport/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("tcp-mss", types.YLeaf{"TcpMss", config.TcpMss})
    config.EntityData.Leafs.Append("mtu-discovery", types.YLeaf{"MtuDiscovery", config.MtuDiscovery})
    config.EntityData.Leafs.Append("passive-mode", types.YLeaf{"PassiveMode", config.PassiveMode})
    config.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", config.LocalAddress})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_Transport_State
// State information relating to the transport session(s)
// used for the BGP neighbor
type Bgp_Neighbors_Neighbor_Transport_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets the max segment size for BGP TCP sessions. The type is interface{}
    // with range: 0..65535.
    TcpMss interface{}

    // Turns path mtu discovery for BGP TCP sessions on (true) or off (false). The
    // type is bool. The default value is false.
    MtuDiscovery interface{}

    // Wait for peers to issue requests to open a BGP session, rather than
    // initiating sessions from the local router. The type is bool. The default
    // value is false.
    PassiveMode interface{}

    // Set the local IP (either IPv4 or IPv6) address to use for the session when
    // sending BGP update messages.  This may be expressed as either an IP address
    // or reference to the name of an interface. The type is one of the following
    // types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.,
    // or string.
    LocalAddress interface{}

    // Local TCP port being used for the TCP session supporting the BGP session.
    // The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Remote address to which the BGP session has been established. The type is
    // one of the following types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.
    RemoteAddress interface{}

    // Remote port being used by the peer for the TCP session supporting the BGP
    // session. The type is interface{} with range: 0..65535.
    RemotePort interface{}
}

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "transport"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/transport/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("tcp-mss", types.YLeaf{"TcpMss", state.TcpMss})
    state.EntityData.Leafs.Append("mtu-discovery", types.YLeaf{"MtuDiscovery", state.MtuDiscovery})
    state.EntityData.Leafs.Append("passive-mode", types.YLeaf{"PassiveMode", state.PassiveMode})
    state.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", state.LocalAddress})
    state.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", state.LocalPort})
    state.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", state.RemoteAddress})
    state.EntityData.Leafs.Append("remote-port", types.YLeaf{"RemotePort", state.RemotePort})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_ErrorHandling
// Error handling parameters used for the BGP neighbor or
// group
type Bgp_Neighbors_Neighbor_ErrorHandling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters enabling or modifying the behavior or enhanced
    // error handling mechanisms for the BGP neighbor.
    Config Bgp_Neighbors_Neighbor_ErrorHandling_Config

    // State information relating to enhanced error handling mechanisms for the
    // BGP neighbor.
    State Bgp_Neighbors_Neighbor_ErrorHandling_State
}

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetEntityData() *types.CommonEntityData {
    errorHandling.EntityData.YFilter = errorHandling.YFilter
    errorHandling.EntityData.YangName = "error-handling"
    errorHandling.EntityData.BundleName = "openconfig"
    errorHandling.EntityData.ParentYangName = "neighbor"
    errorHandling.EntityData.SegmentPath = "error-handling"
    errorHandling.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + errorHandling.EntityData.SegmentPath
    errorHandling.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    errorHandling.EntityData.NamespaceTable = openconfig.GetNamespaces()
    errorHandling.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    errorHandling.EntityData.Children = types.NewOrderedMap()
    errorHandling.EntityData.Children.Append("config", types.YChild{"Config", &errorHandling.Config})
    errorHandling.EntityData.Children.Append("state", types.YChild{"State", &errorHandling.State})
    errorHandling.EntityData.Leafs = types.NewOrderedMap()

    errorHandling.EntityData.YListKeys = []string {}

    return &(errorHandling.EntityData)
}

// Bgp_Neighbors_Neighbor_ErrorHandling_Config
// Configuration parameters enabling or modifying the
// behavior or enhanced error handling mechanisms for the BGP
// neighbor
type Bgp_Neighbors_Neighbor_ErrorHandling_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether erroneous UPDATE messages for which the NLRI can be
    // extracted are reated as though the NLRI is withdrawn - avoiding session
    // reset. The type is bool. The default value is false.
    TreatAsWithdraw interface{}
}

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "error-handling"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/error-handling/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("treat-as-withdraw", types.YLeaf{"TreatAsWithdraw", config.TreatAsWithdraw})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_ErrorHandling_State
// State information relating to enhanced error handling
// mechanisms for the BGP neighbor
type Bgp_Neighbors_Neighbor_ErrorHandling_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether erroneous UPDATE messages for which the NLRI can be
    // extracted are reated as though the NLRI is withdrawn - avoiding session
    // reset. The type is bool. The default value is false.
    TreatAsWithdraw interface{}

    // The number of BGP UPDATE messages for which the treat-as-withdraw mechanism
    // has been applied based on erroneous message contents. The type is
    // interface{} with range: 0..4294967295.
    ErroneousUpdateMessages interface{}
}

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "error-handling"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/error-handling/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("treat-as-withdraw", types.YLeaf{"TreatAsWithdraw", state.TreatAsWithdraw})
    state.EntityData.Leafs.Append("erroneous-update-messages", types.YLeaf{"ErroneousUpdateMessages", state.ErroneousUpdateMessages})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_GracefulRestart
// Parameters relating the graceful restart mechanism for BGP
type Bgp_Neighbors_Neighbor_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to graceful-restart.
    Config Bgp_Neighbors_Neighbor_GracefulRestart_Config

    // State information associated with graceful-restart.
    State Bgp_Neighbors_Neighbor_GracefulRestart_State
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "openconfig"
    gracefulRestart.EntityData.ParentYangName = "neighbor"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("config", types.YChild{"Config", &gracefulRestart.Config})
    gracefulRestart.EntityData.Children.Append("state", types.YChild{"State", &gracefulRestart.State})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// Bgp_Neighbors_Neighbor_GracefulRestart_Config
// Configuration parameters relating to graceful-restart
type Bgp_Neighbors_Neighbor_GracefulRestart_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable the graceful-restart capability. The type is bool.
    Enabled interface{}

    // Estimated time (in seconds) for the local BGP speaker to restart a session.
    // This value is advertise in the graceful restart BGP capability.  This is a
    // 12-bit value, referred to as Restart Time in RFC4724.  Per RFC4724, the
    // suggested default value is <= the hold-time value. The type is interface{}
    // with range: 0..4096.
    RestartTime interface{}

    // An upper-bound on the time thate stale routes will be retained by a router
    // after a session is restarted. If an End-of-RIB (EOR) marker is received
    // prior to this timer expiring stale-routes will be flushed upon its receipt
    // - if no EOR is received, then when this timer expires stale paths will be
    // purged. This timer is referred to as the Selection_Deferral_Timer in
    // RFC4724. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    StaleRoutesTime interface{}

    // Enable graceful-restart in helper mode only. When this leaf is set, the
    // local system does not retain forwarding its own state during a restart, but
    // supports procedures for the receiving speaker, as defined in RFC4724. The
    // type is bool.
    HelperOnly interface{}
}

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "graceful-restart"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/graceful-restart/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", config.RestartTime})
    config.EntityData.Leafs.Append("stale-routes-time", types.YLeaf{"StaleRoutesTime", config.StaleRoutesTime})
    config.EntityData.Leafs.Append("helper-only", types.YLeaf{"HelperOnly", config.HelperOnly})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_GracefulRestart_State
// State information associated with graceful-restart
type Bgp_Neighbors_Neighbor_GracefulRestart_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable the graceful-restart capability. The type is bool.
    Enabled interface{}

    // Estimated time (in seconds) for the local BGP speaker to restart a session.
    // This value is advertise in the graceful restart BGP capability.  This is a
    // 12-bit value, referred to as Restart Time in RFC4724.  Per RFC4724, the
    // suggested default value is <= the hold-time value. The type is interface{}
    // with range: 0..4096.
    RestartTime interface{}

    // An upper-bound on the time thate stale routes will be retained by a router
    // after a session is restarted. If an End-of-RIB (EOR) marker is received
    // prior to this timer expiring stale-routes will be flushed upon its receipt
    // - if no EOR is received, then when this timer expires stale paths will be
    // purged. This timer is referred to as the Selection_Deferral_Timer in
    // RFC4724. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    StaleRoutesTime interface{}

    // Enable graceful-restart in helper mode only. When this leaf is set, the
    // local system does not retain forwarding its own state during a restart, but
    // supports procedures for the receiving speaker, as defined in RFC4724. The
    // type is bool.
    HelperOnly interface{}

    // The period of time (advertised by the peer) that the peer expects a restart
    // of a BGP session to take. The type is interface{} with range: 0..4096.
    PeerRestartTime interface{}

    // This flag indicates whether the remote neighbor is currently in the process
    // of restarting, and hence received routes are currently stale. The type is
    // bool.
    PeerRestarting interface{}

    // This flag indicates whether the local neighbor is currently restarting. The
    // flag is unset after all NLRI have been advertised to the peer, and the
    // End-of-RIB (EOR) marker has been unset. The type is bool.
    LocalRestarting interface{}

    // Ths leaf indicates the mode of operation of BGP graceful restart with the
    // peer. The type is Mode.
    Mode interface{}
}

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "graceful-restart"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/graceful-restart/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", state.RestartTime})
    state.EntityData.Leafs.Append("stale-routes-time", types.YLeaf{"StaleRoutesTime", state.StaleRoutesTime})
    state.EntityData.Leafs.Append("helper-only", types.YLeaf{"HelperOnly", state.HelperOnly})
    state.EntityData.Leafs.Append("peer-restart-time", types.YLeaf{"PeerRestartTime", state.PeerRestartTime})
    state.EntityData.Leafs.Append("peer-restarting", types.YLeaf{"PeerRestarting", state.PeerRestarting})
    state.EntityData.Leafs.Append("local-restarting", types.YLeaf{"LocalRestarting", state.LocalRestarting})
    state.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", state.Mode})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_GracefulRestart_State_Mode represents restart with the peer
type Bgp_Neighbors_Neighbor_GracefulRestart_State_Mode string

const (
    // The local router is operating in helper-only mode, and
    // hence will not retain forwarding state during a local
    // session restart, but will do so during a restart of the
    // remote peer
    Bgp_Neighbors_Neighbor_GracefulRestart_State_Mode_HELPER_ONLY Bgp_Neighbors_Neighbor_GracefulRestart_State_Mode = "HELPER_ONLY"

    // The local router is operating in both helper mode, and
    // hence retains forwarding state during a remote restart,
    // and also maintains forwarding state during local session
    // restart
    Bgp_Neighbors_Neighbor_GracefulRestart_State_Mode_BILATERAL Bgp_Neighbors_Neighbor_GracefulRestart_State_Mode = "BILATERAL"

    // The local system is able to retain routes during restart
    // but the remote system is only able to act as a helper
    Bgp_Neighbors_Neighbor_GracefulRestart_State_Mode_REMOTE_HELPER Bgp_Neighbors_Neighbor_GracefulRestart_State_Mode = "REMOTE_HELPER"
)

// Bgp_Neighbors_Neighbor_LoggingOptions
// Logging options for events related to the BGP neighbor or
// group
type Bgp_Neighbors_Neighbor_LoggingOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters enabling or modifying logging for events relating
    // to the BGPgroup.
    Config Bgp_Neighbors_Neighbor_LoggingOptions_Config

    // State information relating to logging for the BGP neighbor or group.
    State Bgp_Neighbors_Neighbor_LoggingOptions_State
}

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetEntityData() *types.CommonEntityData {
    loggingOptions.EntityData.YFilter = loggingOptions.YFilter
    loggingOptions.EntityData.YangName = "logging-options"
    loggingOptions.EntityData.BundleName = "openconfig"
    loggingOptions.EntityData.ParentYangName = "neighbor"
    loggingOptions.EntityData.SegmentPath = "logging-options"
    loggingOptions.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + loggingOptions.EntityData.SegmentPath
    loggingOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    loggingOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    loggingOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    loggingOptions.EntityData.Children = types.NewOrderedMap()
    loggingOptions.EntityData.Children.Append("config", types.YChild{"Config", &loggingOptions.Config})
    loggingOptions.EntityData.Children.Append("state", types.YChild{"State", &loggingOptions.State})
    loggingOptions.EntityData.Leafs = types.NewOrderedMap()

    loggingOptions.EntityData.YListKeys = []string {}

    return &(loggingOptions.EntityData)
}

// Bgp_Neighbors_Neighbor_LoggingOptions_Config
// Configuration parameters enabling or modifying logging
// for events relating to the BGPgroup
type Bgp_Neighbors_Neighbor_LoggingOptions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure logging of peer state changes.  Default is to enable logging of
    // peer state changes. The type is bool. The default value is true.
    LogNeighborStateChanges interface{}
}

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "logging-options"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/logging-options/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("log-neighbor-state-changes", types.YLeaf{"LogNeighborStateChanges", config.LogNeighborStateChanges})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_LoggingOptions_State
// State information relating to logging for the BGP neighbor
// or group
type Bgp_Neighbors_Neighbor_LoggingOptions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure logging of peer state changes.  Default is to enable logging of
    // peer state changes. The type is bool. The default value is true.
    LogNeighborStateChanges interface{}
}

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "logging-options"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/logging-options/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("log-neighbor-state-changes", types.YLeaf{"LogNeighborStateChanges", state.LogNeighborStateChanges})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_EbgpMultihop
// eBGP multi-hop parameters for the BGPgroup
type Bgp_Neighbors_Neighbor_EbgpMultihop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multihop for the BGP group.
    Config Bgp_Neighbors_Neighbor_EbgpMultihop_Config

    // State information for eBGP multihop, for the BGP neighbor or group.
    State Bgp_Neighbors_Neighbor_EbgpMultihop_State
}

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetEntityData() *types.CommonEntityData {
    ebgpMultihop.EntityData.YFilter = ebgpMultihop.YFilter
    ebgpMultihop.EntityData.YangName = "ebgp-multihop"
    ebgpMultihop.EntityData.BundleName = "openconfig"
    ebgpMultihop.EntityData.ParentYangName = "neighbor"
    ebgpMultihop.EntityData.SegmentPath = "ebgp-multihop"
    ebgpMultihop.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + ebgpMultihop.EntityData.SegmentPath
    ebgpMultihop.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgpMultihop.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgpMultihop.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgpMultihop.EntityData.Children = types.NewOrderedMap()
    ebgpMultihop.EntityData.Children.Append("config", types.YChild{"Config", &ebgpMultihop.Config})
    ebgpMultihop.EntityData.Children.Append("state", types.YChild{"State", &ebgpMultihop.State})
    ebgpMultihop.EntityData.Leafs = types.NewOrderedMap()

    ebgpMultihop.EntityData.YListKeys = []string {}

    return &(ebgpMultihop.EntityData)
}

// Bgp_Neighbors_Neighbor_EbgpMultihop_Config
// Configuration parameters relating to eBGP multihop for the
// BGP group
type Bgp_Neighbors_Neighbor_EbgpMultihop_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When enabled the referenced group or neighbors are permitted to be
    // indirectly connected - including cases where the TTL can be decremented
    // between the BGP peers. The type is bool. The default value is false.
    Enabled interface{}

    // Time-to-live value to use when packets are sent to the referenced group or
    // neighbors and ebgp-multihop is enabled. The type is interface{} with range:
    // 0..255.
    MultihopTtl interface{}
}

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ebgp-multihop"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/ebgp-multihop/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("multihop-ttl", types.YLeaf{"MultihopTtl", config.MultihopTtl})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_EbgpMultihop_State
// State information for eBGP multihop, for the BGP neighbor
// or group
type Bgp_Neighbors_Neighbor_EbgpMultihop_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When enabled the referenced group or neighbors are permitted to be
    // indirectly connected - including cases where the TTL can be decremented
    // between the BGP peers. The type is bool. The default value is false.
    Enabled interface{}

    // Time-to-live value to use when packets are sent to the referenced group or
    // neighbors and ebgp-multihop is enabled. The type is interface{} with range:
    // 0..255.
    MultihopTtl interface{}
}

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ebgp-multihop"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/ebgp-multihop/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("multihop-ttl", types.YLeaf{"MultihopTtl", state.MultihopTtl})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_RouteReflector
// Route reflector parameters for the BGPgroup
type Bgp_Neighbors_Neighbor_RouteReflector struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuraton parameters relating to route reflection for the BGPgroup.
    Config Bgp_Neighbors_Neighbor_RouteReflector_Config

    // State information relating to route reflection for the BGPgroup.
    State Bgp_Neighbors_Neighbor_RouteReflector_State
}

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetEntityData() *types.CommonEntityData {
    routeReflector.EntityData.YFilter = routeReflector.YFilter
    routeReflector.EntityData.YangName = "route-reflector"
    routeReflector.EntityData.BundleName = "openconfig"
    routeReflector.EntityData.ParentYangName = "neighbor"
    routeReflector.EntityData.SegmentPath = "route-reflector"
    routeReflector.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + routeReflector.EntityData.SegmentPath
    routeReflector.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeReflector.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeReflector.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeReflector.EntityData.Children = types.NewOrderedMap()
    routeReflector.EntityData.Children.Append("config", types.YChild{"Config", &routeReflector.Config})
    routeReflector.EntityData.Children.Append("state", types.YChild{"State", &routeReflector.State})
    routeReflector.EntityData.Leafs = types.NewOrderedMap()

    routeReflector.EntityData.YListKeys = []string {}

    return &(routeReflector.EntityData)
}

// Bgp_Neighbors_Neighbor_RouteReflector_Config
// Configuraton parameters relating to route reflection
// for the BGPgroup
type Bgp_Neighbors_Neighbor_RouteReflector_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route-reflector cluster id to use when local router is configured as a
    // route reflector.  Commonly set at the group level, but allows a different
    // cluster id to be set for each neighbor. The type is one of the following
    // types: int with range: 0..4294967295, or string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$'.
    RouteReflectorClusterId interface{}

    // Configure the neighbor as a route reflector client. The type is bool. The
    // default value is false.
    RouteReflectorClient interface{}
}

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "route-reflector"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/route-reflector/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("route-reflector-cluster-id", types.YLeaf{"RouteReflectorClusterId", config.RouteReflectorClusterId})
    config.EntityData.Leafs.Append("route-reflector-client", types.YLeaf{"RouteReflectorClient", config.RouteReflectorClient})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_RouteReflector_State
// State information relating to route reflection for the
// BGPgroup
type Bgp_Neighbors_Neighbor_RouteReflector_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route-reflector cluster id to use when local router is configured as a
    // route reflector.  Commonly set at the group level, but allows a different
    // cluster id to be set for each neighbor. The type is one of the following
    // types: int with range: 0..4294967295, or string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$'.
    RouteReflectorClusterId interface{}

    // Configure the neighbor as a route reflector client. The type is bool. The
    // default value is false.
    RouteReflectorClient interface{}
}

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "route-reflector"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/route-reflector/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("route-reflector-cluster-id", types.YLeaf{"RouteReflectorClusterId", state.RouteReflectorClusterId})
    state.EntityData.Leafs.Append("route-reflector-client", types.YLeaf{"RouteReflectorClient", state.RouteReflectorClient})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AsPathOptions
// AS_PATH manipulation parameters for the BGP neighbor or
// group
type Bgp_Neighbors_Neighbor_AsPathOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to AS_PATH manipulation for the BGP peer
    // or group.
    Config Bgp_Neighbors_Neighbor_AsPathOptions_Config

    // State information relating to the AS_PATH manipulation mechanisms for the
    // BGP peer or group.
    State Bgp_Neighbors_Neighbor_AsPathOptions_State
}

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetEntityData() *types.CommonEntityData {
    asPathOptions.EntityData.YFilter = asPathOptions.YFilter
    asPathOptions.EntityData.YangName = "as-path-options"
    asPathOptions.EntityData.BundleName = "openconfig"
    asPathOptions.EntityData.ParentYangName = "neighbor"
    asPathOptions.EntityData.SegmentPath = "as-path-options"
    asPathOptions.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + asPathOptions.EntityData.SegmentPath
    asPathOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    asPathOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    asPathOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    asPathOptions.EntityData.Children = types.NewOrderedMap()
    asPathOptions.EntityData.Children.Append("config", types.YChild{"Config", &asPathOptions.Config})
    asPathOptions.EntityData.Children.Append("state", types.YChild{"State", &asPathOptions.State})
    asPathOptions.EntityData.Leafs = types.NewOrderedMap()

    asPathOptions.EntityData.YListKeys = []string {}

    return &(asPathOptions.EntityData)
}

// Bgp_Neighbors_Neighbor_AsPathOptions_Config
// Configuration parameters relating to AS_PATH manipulation
// for the BGP peer or group
type Bgp_Neighbors_Neighbor_AsPathOptions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the number of occurrences of the local BGP speaker's AS that can
    // occur within the AS_PATH before it is rejected. The type is interface{}
    // with range: 0..255. The default value is 0.
    AllowOwnAs interface{}

    // Replace occurrences of the peer's AS in the AS_PATH with the local
    // autonomous system number. The type is bool. The default value is false.
    ReplacePeerAs interface{}
}

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "as-path-options"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/as-path-options/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("allow-own-as", types.YLeaf{"AllowOwnAs", config.AllowOwnAs})
    config.EntityData.Leafs.Append("replace-peer-as", types.YLeaf{"ReplacePeerAs", config.ReplacePeerAs})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AsPathOptions_State
// State information relating to the AS_PATH manipulation
// mechanisms for the BGP peer or group
type Bgp_Neighbors_Neighbor_AsPathOptions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the number of occurrences of the local BGP speaker's AS that can
    // occur within the AS_PATH before it is rejected. The type is interface{}
    // with range: 0..255. The default value is 0.
    AllowOwnAs interface{}

    // Replace occurrences of the peer's AS in the AS_PATH with the local
    // autonomous system number. The type is bool. The default value is false.
    ReplacePeerAs interface{}
}

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "as-path-options"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/as-path-options/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("allow-own-as", types.YLeaf{"AllowOwnAs", state.AllowOwnAs})
    state.EntityData.Leafs.Append("replace-peer-as", types.YLeaf{"ReplacePeerAs", state.ReplacePeerAs})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AddPaths
// Parameters relating to the advertisement and receipt of
// multiple paths for a single NLRI (add-paths)
type Bgp_Neighbors_Neighbor_AddPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to ADD_PATHS.
    Config Bgp_Neighbors_Neighbor_AddPaths_Config

    // State information associated with ADD_PATHS.
    State Bgp_Neighbors_Neighbor_AddPaths_State
}

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetEntityData() *types.CommonEntityData {
    addPaths.EntityData.YFilter = addPaths.YFilter
    addPaths.EntityData.YangName = "add-paths"
    addPaths.EntityData.BundleName = "openconfig"
    addPaths.EntityData.ParentYangName = "neighbor"
    addPaths.EntityData.SegmentPath = "add-paths"
    addPaths.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + addPaths.EntityData.SegmentPath
    addPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    addPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    addPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    addPaths.EntityData.Children = types.NewOrderedMap()
    addPaths.EntityData.Children.Append("config", types.YChild{"Config", &addPaths.Config})
    addPaths.EntityData.Children.Append("state", types.YChild{"State", &addPaths.State})
    addPaths.EntityData.Leafs = types.NewOrderedMap()

    addPaths.EntityData.YListKeys = []string {}

    return &(addPaths.EntityData)
}

// Bgp_Neighbors_Neighbor_AddPaths_Config
// Configuration parameters relating to ADD_PATHS
type Bgp_Neighbors_Neighbor_AddPaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ability to receive multiple path advertisements for an NLRI from the
    // neighbor or group. The type is bool. The default value is false.
    Receive interface{}

    // The maximum number of paths to advertise to neighbors for a single NLRI.
    // The type is interface{} with range: 0..255.
    SendMax interface{}

    // A reference to a routing policy which can be used to restrict the prefixes
    // for which add-paths is enabled. The type is string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    EligiblePrefixPolicy interface{}
}

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "add-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/add-paths/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("receive", types.YLeaf{"Receive", config.Receive})
    config.EntityData.Leafs.Append("send-max", types.YLeaf{"SendMax", config.SendMax})
    config.EntityData.Leafs.Append("eligible-prefix-policy", types.YLeaf{"EligiblePrefixPolicy", config.EligiblePrefixPolicy})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AddPaths_State
// State information associated with ADD_PATHS
type Bgp_Neighbors_Neighbor_AddPaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ability to receive multiple path advertisements for an NLRI from the
    // neighbor or group. The type is bool. The default value is false.
    Receive interface{}

    // The maximum number of paths to advertise to neighbors for a single NLRI.
    // The type is interface{} with range: 0..255.
    SendMax interface{}

    // A reference to a routing policy which can be used to restrict the prefixes
    // for which add-paths is enabled. The type is string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    EligiblePrefixPolicy interface{}
}

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "add-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/add-paths/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("receive", types.YLeaf{"Receive", state.Receive})
    state.EntityData.Leafs.Append("send-max", types.YLeaf{"SendMax", state.SendMax})
    state.EntityData.Leafs.Append("eligible-prefix-policy", types.YLeaf{"EligiblePrefixPolicy", state.EligiblePrefixPolicy})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_UseMultiplePaths
// Parameters related to the use of multiple-paths for the same
// NLRI when they are received only from this neighbor
type Bgp_Neighbors_Neighbor_UseMultiplePaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to multipath.
    Config Bgp_Neighbors_Neighbor_UseMultiplePaths_Config

    // State parameters relating to multipath.
    State Bgp_Neighbors_Neighbor_UseMultiplePaths_State

    // Multipath configuration for eBGP.
    Ebgp Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetEntityData() *types.CommonEntityData {
    useMultiplePaths.EntityData.YFilter = useMultiplePaths.YFilter
    useMultiplePaths.EntityData.YangName = "use-multiple-paths"
    useMultiplePaths.EntityData.BundleName = "openconfig"
    useMultiplePaths.EntityData.ParentYangName = "neighbor"
    useMultiplePaths.EntityData.SegmentPath = "use-multiple-paths"
    useMultiplePaths.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + useMultiplePaths.EntityData.SegmentPath
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = types.NewOrderedMap()
    useMultiplePaths.EntityData.Children.Append("config", types.YChild{"Config", &useMultiplePaths.Config})
    useMultiplePaths.EntityData.Children.Append("state", types.YChild{"State", &useMultiplePaths.State})
    useMultiplePaths.EntityData.Children.Append("ebgp", types.YChild{"Ebgp", &useMultiplePaths.Ebgp})
    useMultiplePaths.EntityData.Leafs = types.NewOrderedMap()

    useMultiplePaths.EntityData.YListKeys = []string {}

    return &(useMultiplePaths.EntityData)
}

// Bgp_Neighbors_Neighbor_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_Neighbors_Neighbor_UseMultiplePaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "use-multiple-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/use-multiple-paths/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_Neighbors_Neighbor_UseMultiplePaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "use-multiple-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/use-multiple-paths/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp
// Multipath configuration for eBGP
type Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetEntityData() *types.CommonEntityData {
    ebgp.EntityData.YFilter = ebgp.YFilter
    ebgp.EntityData.YangName = "ebgp"
    ebgp.EntityData.BundleName = "openconfig"
    ebgp.EntityData.ParentYangName = "use-multiple-paths"
    ebgp.EntityData.SegmentPath = "ebgp"
    ebgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/use-multiple-paths/" + ebgp.EntityData.SegmentPath
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = types.NewOrderedMap()
    ebgp.EntityData.Children.Append("config", types.YChild{"Config", &ebgp.Config})
    ebgp.EntityData.Children.Append("state", types.YChild{"State", &ebgp.State})
    ebgp.EntityData.Leafs = types.NewOrderedMap()

    ebgp.EntityData.YListKeys = []string {}

    return &(ebgp.EntityData)
}

// Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ebgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/use-multiple-paths/ebgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ebgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/use-multiple-paths/ebgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_Neighbors_Neighbor_ApplyPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_Neighbors_Neighbor_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_Neighbors_Neighbor_ApplyPolicy_State
}

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetEntityData() *types.CommonEntityData {
    applyPolicy.EntityData.YFilter = applyPolicy.YFilter
    applyPolicy.EntityData.YangName = "apply-policy"
    applyPolicy.EntityData.BundleName = "openconfig"
    applyPolicy.EntityData.ParentYangName = "neighbor"
    applyPolicy.EntityData.SegmentPath = "apply-policy"
    applyPolicy.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + applyPolicy.EntityData.SegmentPath
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = types.NewOrderedMap()
    applyPolicy.EntityData.Children.Append("config", types.YChild{"Config", &applyPolicy.Config})
    applyPolicy.EntityData.Children.Append("state", types.YChild{"State", &applyPolicy.State})
    applyPolicy.EntityData.Leafs = types.NewOrderedMap()

    applyPolicy.EntityData.YListKeys = []string {}

    return &(applyPolicy.EntityData)
}

// Bgp_Neighbors_Neighbor_ApplyPolicy_Config
// Policy configuration data.
type Bgp_Neighbors_Neighbor_ApplyPolicy_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of policy names in sequence to be applied on receiving a routing
    // update in the current context, e.g., for the current peer group, neighbor,
    // address family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ImportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the import
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultImportPolicy interface{}

    // list of policy names in sequence to be applied on sending a routing update
    // in the current context, e.g., for the current peer group, neighbor, address
    // family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ExportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the export
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultExportPolicy interface{}
}

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "apply-policy"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/apply-policy/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("import-policy", types.YLeaf{"ImportPolicy", config.ImportPolicy})
    config.EntityData.Leafs.Append("default-import-policy", types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy})
    config.EntityData.Leafs.Append("export-policy", types.YLeaf{"ExportPolicy", config.ExportPolicy})
    config.EntityData.Leafs.Append("default-export-policy", types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_ApplyPolicy_State
// Operational state for routing policy
type Bgp_Neighbors_Neighbor_ApplyPolicy_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of policy names in sequence to be applied on receiving a routing
    // update in the current context, e.g., for the current peer group, neighbor,
    // address family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ImportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the import
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultImportPolicy interface{}

    // list of policy names in sequence to be applied on sending a routing update
    // in the current context, e.g., for the current peer group, neighbor, address
    // family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ExportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the export
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultExportPolicy interface{}
}

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "apply-policy"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/apply-policy/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("import-policy", types.YLeaf{"ImportPolicy", state.ImportPolicy})
    state.EntityData.Leafs.Append("default-import-policy", types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy})
    state.EntityData.Leafs.Append("export-policy", types.YLeaf{"ExportPolicy", state.ExportPolicy})
    state.EntityData.Leafs.Append("default-export-policy", types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis
// Per-address-family configuration parameters associated with
// the neighbor
type Bgp_Neighbors_Neighbor_AfiSafis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI configuration available for the neighbour or group. The type is
    // slice of Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi.
    AfiSafi []*Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi
}

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetEntityData() *types.CommonEntityData {
    afiSafis.EntityData.YFilter = afiSafis.YFilter
    afiSafis.EntityData.YangName = "afi-safis"
    afiSafis.EntityData.BundleName = "openconfig"
    afiSafis.EntityData.ParentYangName = "neighbor"
    afiSafis.EntityData.SegmentPath = "afi-safis"
    afiSafis.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/" + afiSafis.EntityData.SegmentPath
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

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi
// AFI,SAFI configuration available for the
// neighbour or group
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the AFI-SAFI name used as a key for
    // the AFI-SAFI list. The type is one of the following:
    // IPV4UNICASTIPV6UNICASTIPV4LABELEDUNICASTIPV6LABELEDUNICASTL3VPNIPV4UNICASTL3VPNIPV6UNICASTL3VPNIPV4MULTICASTL3VPNIPV6MULTICASTL2VPNVPLSL2VPNEVPN.
    AfiSafiName interface{}

    // Configuration parameters for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config

    // State information relating to the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State

    // Parameters relating to BGP graceful-restart.
    GracefulRestart Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart

    // Anchor point for routing policies in the model. Import and export policies
    // are with respect to the local routing table, i.e., export (send) and import
    // (receive), depending on the context.
    ApplyPolicy Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy

    // IPv4 unicast configuration options.
    Ipv4Unicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast

    // IPv6 unicast configuration options.
    Ipv6Unicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast

    // IPv4 Labeled Unicast configuration options.
    Ipv4LabeledUnicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast

    // IPv6 Labeled Unicast configuration options.
    Ipv6LabeledUnicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast

    // Unicast IPv4 L3VPN configuration options.
    L3vpnIpv4Unicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast

    // Unicast IPv6 L3VPN configuration options.
    L3vpnIpv6Unicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast

    // Multicast IPv4 L3VPN configuration options.
    L3vpnIpv4Multicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast

    // Multicast IPv6 L3VPN configuration options.
    L3vpnIpv6Multicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast

    // BGP-signalled VPLS configuration options.
    L2vpnVpls Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls

    // BGP EVPN configuration options.
    L2vpnEvpn Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn

    // Parameters related to the use of multiple-paths for the same NLRI when they
    // are received only from this neighbor.
    UseMultiplePaths Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths
}

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetEntityData() *types.CommonEntityData {
    afiSafi.EntityData.YFilter = afiSafi.YFilter
    afiSafi.EntityData.YangName = "afi-safi"
    afiSafi.EntityData.BundleName = "openconfig"
    afiSafi.EntityData.ParentYangName = "afi-safis"
    afiSafi.EntityData.SegmentPath = "afi-safi" + types.AddKeyToken(afiSafi.AfiSafiName, "afi-safi-name")
    afiSafi.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/" + afiSafi.EntityData.SegmentPath
    afiSafi.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafi.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafi.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafi.EntityData.Children = types.NewOrderedMap()
    afiSafi.EntityData.Children.Append("config", types.YChild{"Config", &afiSafi.Config})
    afiSafi.EntityData.Children.Append("state", types.YChild{"State", &afiSafi.State})
    afiSafi.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &afiSafi.GracefulRestart})
    afiSafi.EntityData.Children.Append("apply-policy", types.YChild{"ApplyPolicy", &afiSafi.ApplyPolicy})
    afiSafi.EntityData.Children.Append("ipv4-unicast", types.YChild{"Ipv4Unicast", &afiSafi.Ipv4Unicast})
    afiSafi.EntityData.Children.Append("ipv6-unicast", types.YChild{"Ipv6Unicast", &afiSafi.Ipv6Unicast})
    afiSafi.EntityData.Children.Append("ipv4-labeled-unicast", types.YChild{"Ipv4LabeledUnicast", &afiSafi.Ipv4LabeledUnicast})
    afiSafi.EntityData.Children.Append("ipv6-labeled-unicast", types.YChild{"Ipv6LabeledUnicast", &afiSafi.Ipv6LabeledUnicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv4-unicast", types.YChild{"L3vpnIpv4Unicast", &afiSafi.L3vpnIpv4Unicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv6-unicast", types.YChild{"L3vpnIpv6Unicast", &afiSafi.L3vpnIpv6Unicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv4-multicast", types.YChild{"L3vpnIpv4Multicast", &afiSafi.L3vpnIpv4Multicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv6-multicast", types.YChild{"L3vpnIpv6Multicast", &afiSafi.L3vpnIpv6Multicast})
    afiSafi.EntityData.Children.Append("l2vpn-vpls", types.YChild{"L2vpnVpls", &afiSafi.L2vpnVpls})
    afiSafi.EntityData.Children.Append("l2vpn-evpn", types.YChild{"L2vpnEvpn", &afiSafi.L2vpnEvpn})
    afiSafi.EntityData.Children.Append("use-multiple-paths", types.YChild{"UseMultiplePaths", &afiSafi.UseMultiplePaths})
    afiSafi.EntityData.Leafs = types.NewOrderedMap()
    afiSafi.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", afiSafi.AfiSafiName})

    afiSafi.EntityData.YListKeys = []string {"AfiSafiName"}

    return &(afiSafi.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config
// Configuration parameters for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // IPV4UNICASTIPV6UNICASTIPV4LABELEDUNICASTIPV6LABELEDUNICASTL3VPNIPV4UNICASTL3VPNIPV6UNICASTL3VPNIPV4MULTICASTL3VPNIPV6MULTICASTL2VPNVPLSL2VPNEVPN.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "afi-safi"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", config.AfiSafiName})
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State
// State information relating to the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // IPV4UNICASTIPV6UNICASTIPV4LABELEDUNICASTIPV6LABELEDUNICASTL3VPNIPV4UNICASTL3VPNIPV6UNICASTL3VPNIPV4MULTICASTL3VPNIPV6MULTICASTL2VPNVPLSL2VPNEVPN.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}

    // This value indicates whether a particular AFI-SAFI has been succesfully
    // negotiated with the peer. An AFI-SAFI may be enabled in the current running
    // configuration, but a session restart may be required in order to negotiate
    // the new capability. The type is bool.
    Active interface{}

    // Prefix counters for the BGP session.
    Prefixes Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "afi-safi"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("prefixes", types.YChild{"Prefixes", &state.Prefixes})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", state.AfiSafiName})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("active", types.YLeaf{"Active", state.Active})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes
// Prefix counters for the BGP session
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of prefixes received from the neighbor. The type is interface{}
    // with range: 0..4294967295.
    Received interface{}

    // The number of prefixes advertised to the neighbor. The type is interface{}
    // with range: 0..4294967295.
    Sent interface{}

    // The number of advertised prefixes installed in the Loc-RIB. The type is
    // interface{} with range: 0..4294967295.
    Installed interface{}
}

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "openconfig"
    prefixes.EntityData.ParentYangName = "state"
    prefixes.EntityData.SegmentPath = "prefixes"
    prefixes.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/state/" + prefixes.EntityData.SegmentPath
    prefixes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixes.EntityData.Children = types.NewOrderedMap()
    prefixes.EntityData.Leafs = types.NewOrderedMap()
    prefixes.EntityData.Leafs.Append("received", types.YLeaf{"Received", prefixes.Received})
    prefixes.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", prefixes.Sent})
    prefixes.EntityData.Leafs.Append("installed", types.YLeaf{"Installed", prefixes.Installed})

    prefixes.EntityData.YListKeys = []string {}

    return &(prefixes.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart
// Parameters relating to BGP graceful-restart
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration options for BGP graceful-restart.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config

    // State information for BGP graceful-restart.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "openconfig"
    gracefulRestart.EntityData.ParentYangName = "afi-safi"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("config", types.YChild{"Config", &gracefulRestart.Config})
    gracefulRestart.EntityData.Children.Append("state", types.YChild{"State", &gracefulRestart.State})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config
// Configuration options for BGP graceful-restart
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "graceful-restart"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/graceful-restart/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State
// State information for BGP graceful-restart
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}

    // This leaf indicates whether the neighbor advertised the ability to support
    // graceful-restart for this AFI-SAFI. The type is bool.
    Received interface{}

    // This leaf indicates whether the ability to support graceful-restart has
    // been advertised to the peer. The type is bool.
    Advertised interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "graceful-restart"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/graceful-restart/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("received", types.YLeaf{"Received", state.Received})
    state.EntityData.Leafs.Append("advertised", types.YLeaf{"Advertised", state.Advertised})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State
}

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetEntityData() *types.CommonEntityData {
    applyPolicy.EntityData.YFilter = applyPolicy.YFilter
    applyPolicy.EntityData.YangName = "apply-policy"
    applyPolicy.EntityData.BundleName = "openconfig"
    applyPolicy.EntityData.ParentYangName = "afi-safi"
    applyPolicy.EntityData.SegmentPath = "apply-policy"
    applyPolicy.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + applyPolicy.EntityData.SegmentPath
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = types.NewOrderedMap()
    applyPolicy.EntityData.Children.Append("config", types.YChild{"Config", &applyPolicy.Config})
    applyPolicy.EntityData.Children.Append("state", types.YChild{"State", &applyPolicy.State})
    applyPolicy.EntityData.Leafs = types.NewOrderedMap()

    applyPolicy.EntityData.YListKeys = []string {}

    return &(applyPolicy.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config
// Policy configuration data.
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of policy names in sequence to be applied on receiving a routing
    // update in the current context, e.g., for the current peer group, neighbor,
    // address family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ImportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the import
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultImportPolicy interface{}

    // list of policy names in sequence to be applied on sending a routing update
    // in the current context, e.g., for the current peer group, neighbor, address
    // family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ExportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the export
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultExportPolicy interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "apply-policy"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/apply-policy/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("import-policy", types.YLeaf{"ImportPolicy", config.ImportPolicy})
    config.EntityData.Leafs.Append("default-import-policy", types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy})
    config.EntityData.Leafs.Append("export-policy", types.YLeaf{"ExportPolicy", config.ExportPolicy})
    config.EntityData.Leafs.Append("default-export-policy", types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State
// Operational state for routing policy
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of policy names in sequence to be applied on receiving a routing
    // update in the current context, e.g., for the current peer group, neighbor,
    // address family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ImportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the import
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultImportPolicy interface{}

    // list of policy names in sequence to be applied on sending a routing update
    // in the current context, e.g., for the current peer group, neighbor, address
    // family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ExportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the export
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultExportPolicy interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "apply-policy"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/apply-policy/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("import-policy", types.YLeaf{"ImportPolicy", state.ImportPolicy})
    state.EntityData.Leafs.Append("default-import-policy", types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy})
    state.EntityData.Leafs.Append("export-policy", types.YLeaf{"ExportPolicy", state.ExportPolicy})
    state.EntityData.Leafs.Append("default-export-policy", types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast
// IPv4 unicast configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State
}

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetEntityData() *types.CommonEntityData {
    ipv4Unicast.EntityData.YFilter = ipv4Unicast.YFilter
    ipv4Unicast.EntityData.YangName = "ipv4-unicast"
    ipv4Unicast.EntityData.BundleName = "openconfig"
    ipv4Unicast.EntityData.ParentYangName = "afi-safi"
    ipv4Unicast.EntityData.SegmentPath = "ipv4-unicast"
    ipv4Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + ipv4Unicast.EntityData.SegmentPath
    ipv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4Unicast.EntityData.Children = types.NewOrderedMap()
    ipv4Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv4Unicast.PrefixLimit})
    ipv4Unicast.EntityData.Children.Append("config", types.YChild{"Config", &ipv4Unicast.Config})
    ipv4Unicast.EntityData.Children.Append("state", types.YChild{"State", &ipv4Unicast.State})
    ipv4Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv4Unicast.EntityData.YListKeys = []string {}

    return &(ipv4Unicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv4-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv4-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv4-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv4-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv4-unicast"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv4-unicast/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv4-unicast"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv4-unicast/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast
// IPv6 unicast configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State
}

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetEntityData() *types.CommonEntityData {
    ipv6Unicast.EntityData.YFilter = ipv6Unicast.YFilter
    ipv6Unicast.EntityData.YangName = "ipv6-unicast"
    ipv6Unicast.EntityData.BundleName = "openconfig"
    ipv6Unicast.EntityData.ParentYangName = "afi-safi"
    ipv6Unicast.EntityData.SegmentPath = "ipv6-unicast"
    ipv6Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + ipv6Unicast.EntityData.SegmentPath
    ipv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6Unicast.EntityData.Children = types.NewOrderedMap()
    ipv6Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv6Unicast.PrefixLimit})
    ipv6Unicast.EntityData.Children.Append("config", types.YChild{"Config", &ipv6Unicast.Config})
    ipv6Unicast.EntityData.Children.Append("state", types.YChild{"State", &ipv6Unicast.State})
    ipv6Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv6Unicast.EntityData.YListKeys = []string {}

    return &(ipv6Unicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv6-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv6-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv6-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv6-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv6-unicast"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv6-unicast/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv6-unicast"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv6-unicast/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast
// IPv4 Labeled Unicast configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetEntityData() *types.CommonEntityData {
    ipv4LabeledUnicast.EntityData.YFilter = ipv4LabeledUnicast.YFilter
    ipv4LabeledUnicast.EntityData.YangName = "ipv4-labeled-unicast"
    ipv4LabeledUnicast.EntityData.BundleName = "openconfig"
    ipv4LabeledUnicast.EntityData.ParentYangName = "afi-safi"
    ipv4LabeledUnicast.EntityData.SegmentPath = "ipv4-labeled-unicast"
    ipv4LabeledUnicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + ipv4LabeledUnicast.EntityData.SegmentPath
    ipv4LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4LabeledUnicast.EntityData.Children = types.NewOrderedMap()
    ipv4LabeledUnicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv4LabeledUnicast.PrefixLimit})
    ipv4LabeledUnicast.EntityData.Leafs = types.NewOrderedMap()

    ipv4LabeledUnicast.EntityData.YListKeys = []string {}

    return &(ipv4LabeledUnicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv4-labeled-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv4-labeled-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv4-labeled-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv4-labeled-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast
// IPv6 Labeled Unicast configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetEntityData() *types.CommonEntityData {
    ipv6LabeledUnicast.EntityData.YFilter = ipv6LabeledUnicast.YFilter
    ipv6LabeledUnicast.EntityData.YangName = "ipv6-labeled-unicast"
    ipv6LabeledUnicast.EntityData.BundleName = "openconfig"
    ipv6LabeledUnicast.EntityData.ParentYangName = "afi-safi"
    ipv6LabeledUnicast.EntityData.SegmentPath = "ipv6-labeled-unicast"
    ipv6LabeledUnicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + ipv6LabeledUnicast.EntityData.SegmentPath
    ipv6LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6LabeledUnicast.EntityData.Children = types.NewOrderedMap()
    ipv6LabeledUnicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv6LabeledUnicast.PrefixLimit})
    ipv6LabeledUnicast.EntityData.Leafs = types.NewOrderedMap()

    ipv6LabeledUnicast.EntityData.YListKeys = []string {}

    return &(ipv6LabeledUnicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv6-labeled-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv6-labeled-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv6-labeled-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/ipv6-labeled-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast
// Unicast IPv4 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit
}

func (l3vpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv4Unicast.EntityData.YFilter = l3vpnIpv4Unicast.YFilter
    l3vpnIpv4Unicast.EntityData.YangName = "l3vpn-ipv4-unicast"
    l3vpnIpv4Unicast.EntityData.BundleName = "openconfig"
    l3vpnIpv4Unicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv4Unicast.EntityData.SegmentPath = "l3vpn-ipv4-unicast"
    l3vpnIpv4Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + l3vpnIpv4Unicast.EntityData.SegmentPath
    l3vpnIpv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv4Unicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv4Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv4Unicast.PrefixLimit})
    l3vpnIpv4Unicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv4Unicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv4Unicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv4-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv4-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv4-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast
// Unicast IPv6 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit
}

func (l3vpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv6Unicast.EntityData.YFilter = l3vpnIpv6Unicast.YFilter
    l3vpnIpv6Unicast.EntityData.YangName = "l3vpn-ipv6-unicast"
    l3vpnIpv6Unicast.EntityData.BundleName = "openconfig"
    l3vpnIpv6Unicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv6Unicast.EntityData.SegmentPath = "l3vpn-ipv6-unicast"
    l3vpnIpv6Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + l3vpnIpv6Unicast.EntityData.SegmentPath
    l3vpnIpv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv6Unicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv6Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv6Unicast.PrefixLimit})
    l3vpnIpv6Unicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv6Unicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv6Unicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv6-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv6-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv6-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast
// Multicast IPv4 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit
}

func (l3vpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv4Multicast.EntityData.YFilter = l3vpnIpv4Multicast.YFilter
    l3vpnIpv4Multicast.EntityData.YangName = "l3vpn-ipv4-multicast"
    l3vpnIpv4Multicast.EntityData.BundleName = "openconfig"
    l3vpnIpv4Multicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv4Multicast.EntityData.SegmentPath = "l3vpn-ipv4-multicast"
    l3vpnIpv4Multicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + l3vpnIpv4Multicast.EntityData.SegmentPath
    l3vpnIpv4Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv4Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv4Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv4Multicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv4Multicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv4Multicast.PrefixLimit})
    l3vpnIpv4Multicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv4Multicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv4Multicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv4-multicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv4-multicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv4-multicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast
// Multicast IPv6 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit
}

func (l3vpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv6Multicast.EntityData.YFilter = l3vpnIpv6Multicast.YFilter
    l3vpnIpv6Multicast.EntityData.YangName = "l3vpn-ipv6-multicast"
    l3vpnIpv6Multicast.EntityData.BundleName = "openconfig"
    l3vpnIpv6Multicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv6Multicast.EntityData.SegmentPath = "l3vpn-ipv6-multicast"
    l3vpnIpv6Multicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + l3vpnIpv6Multicast.EntityData.SegmentPath
    l3vpnIpv6Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv6Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv6Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv6Multicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv6Multicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv6Multicast.PrefixLimit})
    l3vpnIpv6Multicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv6Multicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv6Multicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv6-multicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv6-multicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l3vpn-ipv6-multicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls
// BGP-signalled VPLS configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit
}

func (l2vpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls) GetEntityData() *types.CommonEntityData {
    l2vpnVpls.EntityData.YFilter = l2vpnVpls.YFilter
    l2vpnVpls.EntityData.YangName = "l2vpn-vpls"
    l2vpnVpls.EntityData.BundleName = "openconfig"
    l2vpnVpls.EntityData.ParentYangName = "afi-safi"
    l2vpnVpls.EntityData.SegmentPath = "l2vpn-vpls"
    l2vpnVpls.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + l2vpnVpls.EntityData.SegmentPath
    l2vpnVpls.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2vpnVpls.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2vpnVpls.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2vpnVpls.EntityData.Children = types.NewOrderedMap()
    l2vpnVpls.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l2vpnVpls.PrefixLimit})
    l2vpnVpls.EntityData.Leafs = types.NewOrderedMap()

    l2vpnVpls.EntityData.YListKeys = []string {}

    return &(l2vpnVpls.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-vpls"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l2vpn-vpls/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l2vpn-vpls/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l2vpn-vpls/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn
// BGP EVPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit
}

func (l2vpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn) GetEntityData() *types.CommonEntityData {
    l2vpnEvpn.EntityData.YFilter = l2vpnEvpn.YFilter
    l2vpnEvpn.EntityData.YangName = "l2vpn-evpn"
    l2vpnEvpn.EntityData.BundleName = "openconfig"
    l2vpnEvpn.EntityData.ParentYangName = "afi-safi"
    l2vpnEvpn.EntityData.SegmentPath = "l2vpn-evpn"
    l2vpnEvpn.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + l2vpnEvpn.EntityData.SegmentPath
    l2vpnEvpn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2vpnEvpn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2vpnEvpn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2vpnEvpn.EntityData.Children = types.NewOrderedMap()
    l2vpnEvpn.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l2vpnEvpn.PrefixLimit})
    l2vpnEvpn.EntityData.Leafs = types.NewOrderedMap()

    l2vpnEvpn.EntityData.YListKeys = []string {}

    return &(l2vpnEvpn.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-evpn"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l2vpn-evpn/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l2vpn-evpn/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/l2vpn-evpn/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths
// Parameters related to the use of multiple-paths for the same
// NLRI when they are received only from this neighbor
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to multipath.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config

    // State parameters relating to multipath.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State

    // Multipath configuration for eBGP.
    Ebgp Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetEntityData() *types.CommonEntityData {
    useMultiplePaths.EntityData.YFilter = useMultiplePaths.YFilter
    useMultiplePaths.EntityData.YangName = "use-multiple-paths"
    useMultiplePaths.EntityData.BundleName = "openconfig"
    useMultiplePaths.EntityData.ParentYangName = "afi-safi"
    useMultiplePaths.EntityData.SegmentPath = "use-multiple-paths"
    useMultiplePaths.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/" + useMultiplePaths.EntityData.SegmentPath
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = types.NewOrderedMap()
    useMultiplePaths.EntityData.Children.Append("config", types.YChild{"Config", &useMultiplePaths.Config})
    useMultiplePaths.EntityData.Children.Append("state", types.YChild{"State", &useMultiplePaths.State})
    useMultiplePaths.EntityData.Children.Append("ebgp", types.YChild{"Ebgp", &useMultiplePaths.Ebgp})
    useMultiplePaths.EntityData.Leafs = types.NewOrderedMap()

    useMultiplePaths.EntityData.YListKeys = []string {}

    return &(useMultiplePaths.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "use-multiple-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/use-multiple-paths/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "use-multiple-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/use-multiple-paths/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp
// Multipath configuration for eBGP
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetEntityData() *types.CommonEntityData {
    ebgp.EntityData.YFilter = ebgp.YFilter
    ebgp.EntityData.YangName = "ebgp"
    ebgp.EntityData.BundleName = "openconfig"
    ebgp.EntityData.ParentYangName = "use-multiple-paths"
    ebgp.EntityData.SegmentPath = "ebgp"
    ebgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/use-multiple-paths/" + ebgp.EntityData.SegmentPath
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = types.NewOrderedMap()
    ebgp.EntityData.Children.Append("config", types.YChild{"Config", &ebgp.Config})
    ebgp.EntityData.Children.Append("state", types.YChild{"State", &ebgp.State})
    ebgp.EntityData.Leafs = types.NewOrderedMap()

    ebgp.EntityData.YListKeys = []string {}

    return &(ebgp.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ebgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/use-multiple-paths/ebgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ebgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/neighbors/neighbor/afi-safis/afi-safi/use-multiple-paths/ebgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups
// Configuration for BGP peer-groups
type Bgp_PeerGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP peer-groups configured on the local system - uniquely
    // identified by peer-group name. The type is slice of
    // Bgp_PeerGroups_PeerGroup.
    PeerGroup []*Bgp_PeerGroups_PeerGroup
}

func (peerGroups *Bgp_PeerGroups) GetEntityData() *types.CommonEntityData {
    peerGroups.EntityData.YFilter = peerGroups.YFilter
    peerGroups.EntityData.YangName = "peer-groups"
    peerGroups.EntityData.BundleName = "openconfig"
    peerGroups.EntityData.ParentYangName = "bgp"
    peerGroups.EntityData.SegmentPath = "peer-groups"
    peerGroups.EntityData.AbsolutePath = "openconfig-bgp:bgp/" + peerGroups.EntityData.SegmentPath
    peerGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    peerGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    peerGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    peerGroups.EntityData.Children = types.NewOrderedMap()
    peerGroups.EntityData.Children.Append("peer-group", types.YChild{"PeerGroup", nil})
    for i := range peerGroups.PeerGroup {
        peerGroups.EntityData.Children.Append(types.GetSegmentPath(peerGroups.PeerGroup[i]), types.YChild{"PeerGroup", peerGroups.PeerGroup[i]})
    }
    peerGroups.EntityData.Leafs = types.NewOrderedMap()

    peerGroups.EntityData.YListKeys = []string {}

    return &(peerGroups.EntityData)
}

// Bgp_PeerGroups_PeerGroup
// List of BGP peer-groups configured on the local system -
// uniquely identified by peer-group name
type Bgp_PeerGroups_PeerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the name of the BGP peer-group used
    // as a key in the peer-group list. The type is string. Refers to
    // bgp.Bgp_PeerGroups_PeerGroup_Config_PeerGroupName
    PeerGroupName interface{}

    // Configuration parameters relating to the BGP neighbor or group.
    Config Bgp_PeerGroups_PeerGroup_Config

    // State information relating to the BGP peer-group.
    State Bgp_PeerGroups_PeerGroup_State

    // Timers related to a BGP peer-group.
    Timers Bgp_PeerGroups_PeerGroup_Timers

    // Transport session parameters for the BGP peer-group.
    Transport Bgp_PeerGroups_PeerGroup_Transport

    // Error handling parameters used for the BGP peer-group.
    ErrorHandling Bgp_PeerGroups_PeerGroup_ErrorHandling

    // Parameters relating the graceful restart mechanism for BGP.
    GracefulRestart Bgp_PeerGroups_PeerGroup_GracefulRestart

    // Logging options for events related to the BGP neighbor or group.
    LoggingOptions Bgp_PeerGroups_PeerGroup_LoggingOptions

    // eBGP multi-hop parameters for the BGPgroup.
    EbgpMultihop Bgp_PeerGroups_PeerGroup_EbgpMultihop

    // Route reflector parameters for the BGPgroup.
    RouteReflector Bgp_PeerGroups_PeerGroup_RouteReflector

    // AS_PATH manipulation parameters for the BGP neighbor or group.
    AsPathOptions Bgp_PeerGroups_PeerGroup_AsPathOptions

    // Parameters relating to the advertisement and receipt of multiple paths for
    // a single NLRI (add-paths).
    AddPaths Bgp_PeerGroups_PeerGroup_AddPaths

    // Parameters related to the use of multiple paths for the same NLRI.
    UseMultiplePaths Bgp_PeerGroups_PeerGroup_UseMultiplePaths

    // Anchor point for routing policies in the model. Import and export policies
    // are with respect to the local routing table, i.e., export (send) and import
    // (receive), depending on the context.
    ApplyPolicy Bgp_PeerGroups_PeerGroup_ApplyPolicy

    // Per-address-family configuration parameters associated with thegroup.
    AfiSafis Bgp_PeerGroups_PeerGroup_AfiSafis
}

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetEntityData() *types.CommonEntityData {
    peerGroup.EntityData.YFilter = peerGroup.YFilter
    peerGroup.EntityData.YangName = "peer-group"
    peerGroup.EntityData.BundleName = "openconfig"
    peerGroup.EntityData.ParentYangName = "peer-groups"
    peerGroup.EntityData.SegmentPath = "peer-group" + types.AddKeyToken(peerGroup.PeerGroupName, "peer-group-name")
    peerGroup.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/" + peerGroup.EntityData.SegmentPath
    peerGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    peerGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    peerGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    peerGroup.EntityData.Children = types.NewOrderedMap()
    peerGroup.EntityData.Children.Append("config", types.YChild{"Config", &peerGroup.Config})
    peerGroup.EntityData.Children.Append("state", types.YChild{"State", &peerGroup.State})
    peerGroup.EntityData.Children.Append("timers", types.YChild{"Timers", &peerGroup.Timers})
    peerGroup.EntityData.Children.Append("transport", types.YChild{"Transport", &peerGroup.Transport})
    peerGroup.EntityData.Children.Append("error-handling", types.YChild{"ErrorHandling", &peerGroup.ErrorHandling})
    peerGroup.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &peerGroup.GracefulRestart})
    peerGroup.EntityData.Children.Append("logging-options", types.YChild{"LoggingOptions", &peerGroup.LoggingOptions})
    peerGroup.EntityData.Children.Append("ebgp-multihop", types.YChild{"EbgpMultihop", &peerGroup.EbgpMultihop})
    peerGroup.EntityData.Children.Append("route-reflector", types.YChild{"RouteReflector", &peerGroup.RouteReflector})
    peerGroup.EntityData.Children.Append("as-path-options", types.YChild{"AsPathOptions", &peerGroup.AsPathOptions})
    peerGroup.EntityData.Children.Append("add-paths", types.YChild{"AddPaths", &peerGroup.AddPaths})
    peerGroup.EntityData.Children.Append("use-multiple-paths", types.YChild{"UseMultiplePaths", &peerGroup.UseMultiplePaths})
    peerGroup.EntityData.Children.Append("apply-policy", types.YChild{"ApplyPolicy", &peerGroup.ApplyPolicy})
    peerGroup.EntityData.Children.Append("afi-safis", types.YChild{"AfiSafis", &peerGroup.AfiSafis})
    peerGroup.EntityData.Leafs = types.NewOrderedMap()
    peerGroup.EntityData.Leafs.Append("peer-group-name", types.YLeaf{"PeerGroupName", peerGroup.PeerGroupName})

    peerGroup.EntityData.YListKeys = []string {"PeerGroupName"}

    return &(peerGroup.EntityData)
}

// Bgp_PeerGroups_PeerGroup_Config
// Configuration parameters relating to the BGP neighbor or
// group
type Bgp_PeerGroups_PeerGroup_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the BGP peer-group. The type is string.
    PeerGroupName interface{}

    // AS number of the peer. The type is interface{} with range: 0..4294967295.
    PeerAs interface{}

    // The local autonomous system number that is to be used when establishing
    // sessions with the remote peer or peer group, if this differs from the
    // global BGP router autonomous system number. The type is interface{} with
    // range: 0..4294967295.
    LocalAs interface{}

    // Explicitly designate the peer or peer group as internal (iBGP) or external
    // (eBGP). The type is PeerType.
    PeerType interface{}

    // Configures an MD5 authentication password for use with neighboring devices.
    // The type is string.
    AuthPassword interface{}

    // Remove private AS numbers from updates sent to peers - when this leaf is
    // not specified, the AS_PATH attribute should be sent to the peer unchanged.
    // The type is one of the following: PRIVATEASREMOVEALLPRIVATEASREPLACEALL.
    RemovePrivateAs interface{}

    // Enable route flap damping. The type is bool. The default value is false.
    RouteFlapDamping interface{}

    // Specify which types of community should be sent to the neighbor or group.
    // The default is to not send the community attribute. The type is
    // CommunityType. The default value is NONE.
    SendCommunity interface{}

    // An optional textual description (intended primarily for use with a peer or
    // group. The type is string.
    Description interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "peer-group"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("peer-group-name", types.YLeaf{"PeerGroupName", config.PeerGroupName})
    config.EntityData.Leafs.Append("peer-as", types.YLeaf{"PeerAs", config.PeerAs})
    config.EntityData.Leafs.Append("local-as", types.YLeaf{"LocalAs", config.LocalAs})
    config.EntityData.Leafs.Append("peer-type", types.YLeaf{"PeerType", config.PeerType})
    config.EntityData.Leafs.Append("auth-password", types.YLeaf{"AuthPassword", config.AuthPassword})
    config.EntityData.Leafs.Append("remove-private-as", types.YLeaf{"RemovePrivateAs", config.RemovePrivateAs})
    config.EntityData.Leafs.Append("route-flap-damping", types.YLeaf{"RouteFlapDamping", config.RouteFlapDamping})
    config.EntityData.Leafs.Append("send-community", types.YLeaf{"SendCommunity", config.SendCommunity})
    config.EntityData.Leafs.Append("description", types.YLeaf{"Description", config.Description})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_State
// State information relating to the BGP peer-group
type Bgp_PeerGroups_PeerGroup_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the BGP peer-group. The type is string.
    PeerGroupName interface{}

    // AS number of the peer. The type is interface{} with range: 0..4294967295.
    PeerAs interface{}

    // The local autonomous system number that is to be used when establishing
    // sessions with the remote peer or peer group, if this differs from the
    // global BGP router autonomous system number. The type is interface{} with
    // range: 0..4294967295.
    LocalAs interface{}

    // Explicitly designate the peer or peer group as internal (iBGP) or external
    // (eBGP). The type is PeerType.
    PeerType interface{}

    // Configures an MD5 authentication password for use with neighboring devices.
    // The type is string.
    AuthPassword interface{}

    // Remove private AS numbers from updates sent to peers - when this leaf is
    // not specified, the AS_PATH attribute should be sent to the peer unchanged.
    // The type is one of the following: PRIVATEASREMOVEALLPRIVATEASREPLACEALL.
    RemovePrivateAs interface{}

    // Enable route flap damping. The type is bool. The default value is false.
    RouteFlapDamping interface{}

    // Specify which types of community should be sent to the neighbor or group.
    // The default is to not send the community attribute. The type is
    // CommunityType. The default value is NONE.
    SendCommunity interface{}

    // An optional textual description (intended primarily for use with a peer or
    // group. The type is string.
    Description interface{}

    // Total number of BGP paths within the context. The type is interface{} with
    // range: 0..4294967295.
    TotalPaths interface{}

    // Total number of BGP prefixes received within the context. The type is
    // interface{} with range: 0..4294967295.
    TotalPrefixes interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "peer-group"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("peer-group-name", types.YLeaf{"PeerGroupName", state.PeerGroupName})
    state.EntityData.Leafs.Append("peer-as", types.YLeaf{"PeerAs", state.PeerAs})
    state.EntityData.Leafs.Append("local-as", types.YLeaf{"LocalAs", state.LocalAs})
    state.EntityData.Leafs.Append("peer-type", types.YLeaf{"PeerType", state.PeerType})
    state.EntityData.Leafs.Append("auth-password", types.YLeaf{"AuthPassword", state.AuthPassword})
    state.EntityData.Leafs.Append("remove-private-as", types.YLeaf{"RemovePrivateAs", state.RemovePrivateAs})
    state.EntityData.Leafs.Append("route-flap-damping", types.YLeaf{"RouteFlapDamping", state.RouteFlapDamping})
    state.EntityData.Leafs.Append("send-community", types.YLeaf{"SendCommunity", state.SendCommunity})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("total-paths", types.YLeaf{"TotalPaths", state.TotalPaths})
    state.EntityData.Leafs.Append("total-prefixes", types.YLeaf{"TotalPrefixes", state.TotalPrefixes})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_Timers
// Timers related to a BGP peer-group
type Bgp_PeerGroups_PeerGroup_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to timers used for the BGP neighbor or
    // peer group.
    Config Bgp_PeerGroups_PeerGroup_Timers_Config

    // State information relating to the timers used for the BGP group.
    State Bgp_PeerGroups_PeerGroup_Timers_State
}

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "openconfig"
    timers.EntityData.ParentYangName = "peer-group"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + timers.EntityData.SegmentPath
    timers.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    timers.EntityData.NamespaceTable = openconfig.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("config", types.YChild{"Config", &timers.Config})
    timers.EntityData.Children.Append("state", types.YChild{"State", &timers.State})
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// Bgp_PeerGroups_PeerGroup_Timers_Config
// Configuration parameters relating to timers used for the
// BGP neighbor or peer group
type Bgp_PeerGroups_PeerGroup_Timers_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time interval in seconds between attempts to establish a session with the
    // peer. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 30.
    ConnectRetry interface{}

    // Time interval in seconds that a BGP session will be considered active in
    // the absence of keepalive or other messages from the peer.  The hold-time is
    // typically set to 3x the keepalive-interval. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 90.
    HoldTime interface{}

    // Time interval in seconds between transmission of keepalive messages to the
    // neighbor.  Typically set to 1/3 the hold-time. The type is string with
    // range: -92233720368547758.08..92233720368547758.07. The default value is
    // 30.
    KeepaliveInterval interface{}

    // Minimum time which must elapse between subsequent UPDATE messages relating
    // to a common set of NLRI being transmitted to a peer. This timer is referred
    // to as MinRouteAdvertisementIntervalTimer by RFC 4721 and serves to reduce
    // the number of UPDATE messages transmitted when a particular set of NLRI
    // exhibit instability. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 30.
    MinimumAdvertisementInterval interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "timers"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/timers/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("connect-retry", types.YLeaf{"ConnectRetry", config.ConnectRetry})
    config.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", config.HoldTime})
    config.EntityData.Leafs.Append("keepalive-interval", types.YLeaf{"KeepaliveInterval", config.KeepaliveInterval})
    config.EntityData.Leafs.Append("minimum-advertisement-interval", types.YLeaf{"MinimumAdvertisementInterval", config.MinimumAdvertisementInterval})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_Timers_State
// State information relating to the timers used for the BGP
// group
type Bgp_PeerGroups_PeerGroup_Timers_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time interval in seconds between attempts to establish a session with the
    // peer. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 30.
    ConnectRetry interface{}

    // Time interval in seconds that a BGP session will be considered active in
    // the absence of keepalive or other messages from the peer.  The hold-time is
    // typically set to 3x the keepalive-interval. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 90.
    HoldTime interface{}

    // Time interval in seconds between transmission of keepalive messages to the
    // neighbor.  Typically set to 1/3 the hold-time. The type is string with
    // range: -92233720368547758.08..92233720368547758.07. The default value is
    // 30.
    KeepaliveInterval interface{}

    // Minimum time which must elapse between subsequent UPDATE messages relating
    // to a common set of NLRI being transmitted to a peer. This timer is referred
    // to as MinRouteAdvertisementIntervalTimer by RFC 4721 and serves to reduce
    // the number of UPDATE messages transmitted when a particular set of NLRI
    // exhibit instability. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. The default value is 30.
    MinimumAdvertisementInterval interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "timers"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/timers/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("connect-retry", types.YLeaf{"ConnectRetry", state.ConnectRetry})
    state.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", state.HoldTime})
    state.EntityData.Leafs.Append("keepalive-interval", types.YLeaf{"KeepaliveInterval", state.KeepaliveInterval})
    state.EntityData.Leafs.Append("minimum-advertisement-interval", types.YLeaf{"MinimumAdvertisementInterval", state.MinimumAdvertisementInterval})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_Transport
// Transport session parameters for the BGP peer-group
type Bgp_PeerGroups_PeerGroup_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the transport session(s) used for the
    // BGP neighbor or group.
    Config Bgp_PeerGroups_PeerGroup_Transport_Config

    // State information relating to the transport session(s) used for the BGP
    // neighbor or group.
    State Bgp_PeerGroups_PeerGroup_Transport_State
}

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "openconfig"
    transport.EntityData.ParentYangName = "peer-group"
    transport.EntityData.SegmentPath = "transport"
    transport.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + transport.EntityData.SegmentPath
    transport.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    transport.EntityData.NamespaceTable = openconfig.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    transport.EntityData.Children = types.NewOrderedMap()
    transport.EntityData.Children.Append("config", types.YChild{"Config", &transport.Config})
    transport.EntityData.Children.Append("state", types.YChild{"State", &transport.State})
    transport.EntityData.Leafs = types.NewOrderedMap()

    transport.EntityData.YListKeys = []string {}

    return &(transport.EntityData)
}

// Bgp_PeerGroups_PeerGroup_Transport_Config
// Configuration parameters relating to the transport
// session(s) used for the BGP neighbor or group
type Bgp_PeerGroups_PeerGroup_Transport_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets the max segment size for BGP TCP sessions. The type is interface{}
    // with range: 0..65535.
    TcpMss interface{}

    // Turns path mtu discovery for BGP TCP sessions on (true) or off (false). The
    // type is bool. The default value is false.
    MtuDiscovery interface{}

    // Wait for peers to issue requests to open a BGP session, rather than
    // initiating sessions from the local router. The type is bool. The default
    // value is false.
    PassiveMode interface{}

    // Set the local IP (either IPv4 or IPv6) address to use for the session when
    // sending BGP update messages.  This may be expressed as either an IP address
    // or reference to the name of an interface. The type is one of the following
    // types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.,
    // or string.
    LocalAddress interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "transport"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/transport/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("tcp-mss", types.YLeaf{"TcpMss", config.TcpMss})
    config.EntityData.Leafs.Append("mtu-discovery", types.YLeaf{"MtuDiscovery", config.MtuDiscovery})
    config.EntityData.Leafs.Append("passive-mode", types.YLeaf{"PassiveMode", config.PassiveMode})
    config.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", config.LocalAddress})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_Transport_State
// State information relating to the transport session(s)
// used for the BGP neighbor or group
type Bgp_PeerGroups_PeerGroup_Transport_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets the max segment size for BGP TCP sessions. The type is interface{}
    // with range: 0..65535.
    TcpMss interface{}

    // Turns path mtu discovery for BGP TCP sessions on (true) or off (false). The
    // type is bool. The default value is false.
    MtuDiscovery interface{}

    // Wait for peers to issue requests to open a BGP session, rather than
    // initiating sessions from the local router. The type is bool. The default
    // value is false.
    PassiveMode interface{}

    // Set the local IP (either IPv4 or IPv6) address to use for the session when
    // sending BGP update messages.  This may be expressed as either an IP address
    // or reference to the name of an interface. The type is one of the following
    // types: string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$',
    // or string with pattern:
    // b'^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$'.,
    // or string.
    LocalAddress interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "transport"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/transport/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("tcp-mss", types.YLeaf{"TcpMss", state.TcpMss})
    state.EntityData.Leafs.Append("mtu-discovery", types.YLeaf{"MtuDiscovery", state.MtuDiscovery})
    state.EntityData.Leafs.Append("passive-mode", types.YLeaf{"PassiveMode", state.PassiveMode})
    state.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", state.LocalAddress})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_ErrorHandling
// Error handling parameters used for the BGP peer-group
type Bgp_PeerGroups_PeerGroup_ErrorHandling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters enabling or modifying the behavior or enhanced
    // error handling mechanisms for the BGP group.
    Config Bgp_PeerGroups_PeerGroup_ErrorHandling_Config

    // State information relating to enhanced error handling mechanisms for the
    // BGP group.
    State Bgp_PeerGroups_PeerGroup_ErrorHandling_State
}

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetEntityData() *types.CommonEntityData {
    errorHandling.EntityData.YFilter = errorHandling.YFilter
    errorHandling.EntityData.YangName = "error-handling"
    errorHandling.EntityData.BundleName = "openconfig"
    errorHandling.EntityData.ParentYangName = "peer-group"
    errorHandling.EntityData.SegmentPath = "error-handling"
    errorHandling.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + errorHandling.EntityData.SegmentPath
    errorHandling.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    errorHandling.EntityData.NamespaceTable = openconfig.GetNamespaces()
    errorHandling.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    errorHandling.EntityData.Children = types.NewOrderedMap()
    errorHandling.EntityData.Children.Append("config", types.YChild{"Config", &errorHandling.Config})
    errorHandling.EntityData.Children.Append("state", types.YChild{"State", &errorHandling.State})
    errorHandling.EntityData.Leafs = types.NewOrderedMap()

    errorHandling.EntityData.YListKeys = []string {}

    return &(errorHandling.EntityData)
}

// Bgp_PeerGroups_PeerGroup_ErrorHandling_Config
// Configuration parameters enabling or modifying the
// behavior or enhanced error handling mechanisms for the BGP
// group
type Bgp_PeerGroups_PeerGroup_ErrorHandling_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether erroneous UPDATE messages for which the NLRI can be
    // extracted are reated as though the NLRI is withdrawn - avoiding session
    // reset. The type is bool. The default value is false.
    TreatAsWithdraw interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "error-handling"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/error-handling/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("treat-as-withdraw", types.YLeaf{"TreatAsWithdraw", config.TreatAsWithdraw})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_ErrorHandling_State
// State information relating to enhanced error handling
// mechanisms for the BGP group
type Bgp_PeerGroups_PeerGroup_ErrorHandling_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether erroneous UPDATE messages for which the NLRI can be
    // extracted are reated as though the NLRI is withdrawn - avoiding session
    // reset. The type is bool. The default value is false.
    TreatAsWithdraw interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "error-handling"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/error-handling/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("treat-as-withdraw", types.YLeaf{"TreatAsWithdraw", state.TreatAsWithdraw})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_GracefulRestart
// Parameters relating the graceful restart mechanism for BGP
type Bgp_PeerGroups_PeerGroup_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to graceful-restart.
    Config Bgp_PeerGroups_PeerGroup_GracefulRestart_Config

    // State information associated with graceful-restart.
    State Bgp_PeerGroups_PeerGroup_GracefulRestart_State
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "openconfig"
    gracefulRestart.EntityData.ParentYangName = "peer-group"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("config", types.YChild{"Config", &gracefulRestart.Config})
    gracefulRestart.EntityData.Children.Append("state", types.YChild{"State", &gracefulRestart.State})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// Bgp_PeerGroups_PeerGroup_GracefulRestart_Config
// Configuration parameters relating to graceful-restart
type Bgp_PeerGroups_PeerGroup_GracefulRestart_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable the graceful-restart capability. The type is bool.
    Enabled interface{}

    // Estimated time (in seconds) for the local BGP speaker to restart a session.
    // This value is advertise in the graceful restart BGP capability.  This is a
    // 12-bit value, referred to as Restart Time in RFC4724.  Per RFC4724, the
    // suggested default value is <= the hold-time value. The type is interface{}
    // with range: 0..4096.
    RestartTime interface{}

    // An upper-bound on the time thate stale routes will be retained by a router
    // after a session is restarted. If an End-of-RIB (EOR) marker is received
    // prior to this timer expiring stale-routes will be flushed upon its receipt
    // - if no EOR is received, then when this timer expires stale paths will be
    // purged. This timer is referred to as the Selection_Deferral_Timer in
    // RFC4724. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    StaleRoutesTime interface{}

    // Enable graceful-restart in helper mode only. When this leaf is set, the
    // local system does not retain forwarding its own state during a restart, but
    // supports procedures for the receiving speaker, as defined in RFC4724. The
    // type is bool.
    HelperOnly interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "graceful-restart"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/graceful-restart/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", config.RestartTime})
    config.EntityData.Leafs.Append("stale-routes-time", types.YLeaf{"StaleRoutesTime", config.StaleRoutesTime})
    config.EntityData.Leafs.Append("helper-only", types.YLeaf{"HelperOnly", config.HelperOnly})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_GracefulRestart_State
// State information associated with graceful-restart
type Bgp_PeerGroups_PeerGroup_GracefulRestart_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable the graceful-restart capability. The type is bool.
    Enabled interface{}

    // Estimated time (in seconds) for the local BGP speaker to restart a session.
    // This value is advertise in the graceful restart BGP capability.  This is a
    // 12-bit value, referred to as Restart Time in RFC4724.  Per RFC4724, the
    // suggested default value is <= the hold-time value. The type is interface{}
    // with range: 0..4096.
    RestartTime interface{}

    // An upper-bound on the time thate stale routes will be retained by a router
    // after a session is restarted. If an End-of-RIB (EOR) marker is received
    // prior to this timer expiring stale-routes will be flushed upon its receipt
    // - if no EOR is received, then when this timer expires stale paths will be
    // purged. This timer is referred to as the Selection_Deferral_Timer in
    // RFC4724. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    StaleRoutesTime interface{}

    // Enable graceful-restart in helper mode only. When this leaf is set, the
    // local system does not retain forwarding its own state during a restart, but
    // supports procedures for the receiving speaker, as defined in RFC4724. The
    // type is bool.
    HelperOnly interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "graceful-restart"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/graceful-restart/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", state.RestartTime})
    state.EntityData.Leafs.Append("stale-routes-time", types.YLeaf{"StaleRoutesTime", state.StaleRoutesTime})
    state.EntityData.Leafs.Append("helper-only", types.YLeaf{"HelperOnly", state.HelperOnly})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_LoggingOptions
// Logging options for events related to the BGP neighbor or
// group
type Bgp_PeerGroups_PeerGroup_LoggingOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters enabling or modifying logging for events relating
    // to the BGPgroup.
    Config Bgp_PeerGroups_PeerGroup_LoggingOptions_Config

    // State information relating to logging for the BGP neighbor or group.
    State Bgp_PeerGroups_PeerGroup_LoggingOptions_State
}

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetEntityData() *types.CommonEntityData {
    loggingOptions.EntityData.YFilter = loggingOptions.YFilter
    loggingOptions.EntityData.YangName = "logging-options"
    loggingOptions.EntityData.BundleName = "openconfig"
    loggingOptions.EntityData.ParentYangName = "peer-group"
    loggingOptions.EntityData.SegmentPath = "logging-options"
    loggingOptions.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + loggingOptions.EntityData.SegmentPath
    loggingOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    loggingOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    loggingOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    loggingOptions.EntityData.Children = types.NewOrderedMap()
    loggingOptions.EntityData.Children.Append("config", types.YChild{"Config", &loggingOptions.Config})
    loggingOptions.EntityData.Children.Append("state", types.YChild{"State", &loggingOptions.State})
    loggingOptions.EntityData.Leafs = types.NewOrderedMap()

    loggingOptions.EntityData.YListKeys = []string {}

    return &(loggingOptions.EntityData)
}

// Bgp_PeerGroups_PeerGroup_LoggingOptions_Config
// Configuration parameters enabling or modifying logging
// for events relating to the BGPgroup
type Bgp_PeerGroups_PeerGroup_LoggingOptions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure logging of peer state changes.  Default is to enable logging of
    // peer state changes. The type is bool. The default value is true.
    LogNeighborStateChanges interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "logging-options"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/logging-options/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("log-neighbor-state-changes", types.YLeaf{"LogNeighborStateChanges", config.LogNeighborStateChanges})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_LoggingOptions_State
// State information relating to logging for the BGP neighbor
// or group
type Bgp_PeerGroups_PeerGroup_LoggingOptions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure logging of peer state changes.  Default is to enable logging of
    // peer state changes. The type is bool. The default value is true.
    LogNeighborStateChanges interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "logging-options"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/logging-options/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("log-neighbor-state-changes", types.YLeaf{"LogNeighborStateChanges", state.LogNeighborStateChanges})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_EbgpMultihop
// eBGP multi-hop parameters for the BGPgroup
type Bgp_PeerGroups_PeerGroup_EbgpMultihop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multihop for the BGP group.
    Config Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config

    // State information for eBGP multihop, for the BGP neighbor or group.
    State Bgp_PeerGroups_PeerGroup_EbgpMultihop_State
}

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetEntityData() *types.CommonEntityData {
    ebgpMultihop.EntityData.YFilter = ebgpMultihop.YFilter
    ebgpMultihop.EntityData.YangName = "ebgp-multihop"
    ebgpMultihop.EntityData.BundleName = "openconfig"
    ebgpMultihop.EntityData.ParentYangName = "peer-group"
    ebgpMultihop.EntityData.SegmentPath = "ebgp-multihop"
    ebgpMultihop.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + ebgpMultihop.EntityData.SegmentPath
    ebgpMultihop.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgpMultihop.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgpMultihop.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgpMultihop.EntityData.Children = types.NewOrderedMap()
    ebgpMultihop.EntityData.Children.Append("config", types.YChild{"Config", &ebgpMultihop.Config})
    ebgpMultihop.EntityData.Children.Append("state", types.YChild{"State", &ebgpMultihop.State})
    ebgpMultihop.EntityData.Leafs = types.NewOrderedMap()

    ebgpMultihop.EntityData.YListKeys = []string {}

    return &(ebgpMultihop.EntityData)
}

// Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config
// Configuration parameters relating to eBGP multihop for the
// BGP group
type Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When enabled the referenced group or neighbors are permitted to be
    // indirectly connected - including cases where the TTL can be decremented
    // between the BGP peers. The type is bool. The default value is false.
    Enabled interface{}

    // Time-to-live value to use when packets are sent to the referenced group or
    // neighbors and ebgp-multihop is enabled. The type is interface{} with range:
    // 0..255.
    MultihopTtl interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ebgp-multihop"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/ebgp-multihop/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("multihop-ttl", types.YLeaf{"MultihopTtl", config.MultihopTtl})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_EbgpMultihop_State
// State information for eBGP multihop, for the BGP neighbor
// or group
type Bgp_PeerGroups_PeerGroup_EbgpMultihop_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When enabled the referenced group or neighbors are permitted to be
    // indirectly connected - including cases where the TTL can be decremented
    // between the BGP peers. The type is bool. The default value is false.
    Enabled interface{}

    // Time-to-live value to use when packets are sent to the referenced group or
    // neighbors and ebgp-multihop is enabled. The type is interface{} with range:
    // 0..255.
    MultihopTtl interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ebgp-multihop"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/ebgp-multihop/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("multihop-ttl", types.YLeaf{"MultihopTtl", state.MultihopTtl})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_RouteReflector
// Route reflector parameters for the BGPgroup
type Bgp_PeerGroups_PeerGroup_RouteReflector struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuraton parameters relating to route reflection for the BGPgroup.
    Config Bgp_PeerGroups_PeerGroup_RouteReflector_Config

    // State information relating to route reflection for the BGPgroup.
    State Bgp_PeerGroups_PeerGroup_RouteReflector_State
}

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetEntityData() *types.CommonEntityData {
    routeReflector.EntityData.YFilter = routeReflector.YFilter
    routeReflector.EntityData.YangName = "route-reflector"
    routeReflector.EntityData.BundleName = "openconfig"
    routeReflector.EntityData.ParentYangName = "peer-group"
    routeReflector.EntityData.SegmentPath = "route-reflector"
    routeReflector.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + routeReflector.EntityData.SegmentPath
    routeReflector.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeReflector.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeReflector.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeReflector.EntityData.Children = types.NewOrderedMap()
    routeReflector.EntityData.Children.Append("config", types.YChild{"Config", &routeReflector.Config})
    routeReflector.EntityData.Children.Append("state", types.YChild{"State", &routeReflector.State})
    routeReflector.EntityData.Leafs = types.NewOrderedMap()

    routeReflector.EntityData.YListKeys = []string {}

    return &(routeReflector.EntityData)
}

// Bgp_PeerGroups_PeerGroup_RouteReflector_Config
// Configuraton parameters relating to route reflection
// for the BGPgroup
type Bgp_PeerGroups_PeerGroup_RouteReflector_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route-reflector cluster id to use when local router is configured as a
    // route reflector.  Commonly set at the group level, but allows a different
    // cluster id to be set for each neighbor. The type is one of the following
    // types: int with range: 0..4294967295, or string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$'.
    RouteReflectorClusterId interface{}

    // Configure the neighbor as a route reflector client. The type is bool. The
    // default value is false.
    RouteReflectorClient interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "route-reflector"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/route-reflector/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("route-reflector-cluster-id", types.YLeaf{"RouteReflectorClusterId", config.RouteReflectorClusterId})
    config.EntityData.Leafs.Append("route-reflector-client", types.YLeaf{"RouteReflectorClient", config.RouteReflectorClient})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_RouteReflector_State
// State information relating to route reflection for the
// BGPgroup
type Bgp_PeerGroups_PeerGroup_RouteReflector_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // route-reflector cluster id to use when local router is configured as a
    // route reflector.  Commonly set at the group level, but allows a different
    // cluster id to be set for each neighbor. The type is one of the following
    // types: int with range: 0..4294967295, or string with pattern:
    // b'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$'.
    RouteReflectorClusterId interface{}

    // Configure the neighbor as a route reflector client. The type is bool. The
    // default value is false.
    RouteReflectorClient interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "route-reflector"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/route-reflector/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("route-reflector-cluster-id", types.YLeaf{"RouteReflectorClusterId", state.RouteReflectorClusterId})
    state.EntityData.Leafs.Append("route-reflector-client", types.YLeaf{"RouteReflectorClient", state.RouteReflectorClient})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AsPathOptions
// AS_PATH manipulation parameters for the BGP neighbor or
// group
type Bgp_PeerGroups_PeerGroup_AsPathOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to AS_PATH manipulation for the BGP peer
    // or group.
    Config Bgp_PeerGroups_PeerGroup_AsPathOptions_Config

    // State information relating to the AS_PATH manipulation mechanisms for the
    // BGP peer or group.
    State Bgp_PeerGroups_PeerGroup_AsPathOptions_State
}

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetEntityData() *types.CommonEntityData {
    asPathOptions.EntityData.YFilter = asPathOptions.YFilter
    asPathOptions.EntityData.YangName = "as-path-options"
    asPathOptions.EntityData.BundleName = "openconfig"
    asPathOptions.EntityData.ParentYangName = "peer-group"
    asPathOptions.EntityData.SegmentPath = "as-path-options"
    asPathOptions.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + asPathOptions.EntityData.SegmentPath
    asPathOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    asPathOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    asPathOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    asPathOptions.EntityData.Children = types.NewOrderedMap()
    asPathOptions.EntityData.Children.Append("config", types.YChild{"Config", &asPathOptions.Config})
    asPathOptions.EntityData.Children.Append("state", types.YChild{"State", &asPathOptions.State})
    asPathOptions.EntityData.Leafs = types.NewOrderedMap()

    asPathOptions.EntityData.YListKeys = []string {}

    return &(asPathOptions.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AsPathOptions_Config
// Configuration parameters relating to AS_PATH manipulation
// for the BGP peer or group
type Bgp_PeerGroups_PeerGroup_AsPathOptions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the number of occurrences of the local BGP speaker's AS that can
    // occur within the AS_PATH before it is rejected. The type is interface{}
    // with range: 0..255. The default value is 0.
    AllowOwnAs interface{}

    // Replace occurrences of the peer's AS in the AS_PATH with the local
    // autonomous system number. The type is bool. The default value is false.
    ReplacePeerAs interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "as-path-options"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/as-path-options/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("allow-own-as", types.YLeaf{"AllowOwnAs", config.AllowOwnAs})
    config.EntityData.Leafs.Append("replace-peer-as", types.YLeaf{"ReplacePeerAs", config.ReplacePeerAs})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AsPathOptions_State
// State information relating to the AS_PATH manipulation
// mechanisms for the BGP peer or group
type Bgp_PeerGroups_PeerGroup_AsPathOptions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the number of occurrences of the local BGP speaker's AS that can
    // occur within the AS_PATH before it is rejected. The type is interface{}
    // with range: 0..255. The default value is 0.
    AllowOwnAs interface{}

    // Replace occurrences of the peer's AS in the AS_PATH with the local
    // autonomous system number. The type is bool. The default value is false.
    ReplacePeerAs interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "as-path-options"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/as-path-options/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("allow-own-as", types.YLeaf{"AllowOwnAs", state.AllowOwnAs})
    state.EntityData.Leafs.Append("replace-peer-as", types.YLeaf{"ReplacePeerAs", state.ReplacePeerAs})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AddPaths
// Parameters relating to the advertisement and receipt of
// multiple paths for a single NLRI (add-paths)
type Bgp_PeerGroups_PeerGroup_AddPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to ADD_PATHS.
    Config Bgp_PeerGroups_PeerGroup_AddPaths_Config

    // State information associated with ADD_PATHS.
    State Bgp_PeerGroups_PeerGroup_AddPaths_State
}

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetEntityData() *types.CommonEntityData {
    addPaths.EntityData.YFilter = addPaths.YFilter
    addPaths.EntityData.YangName = "add-paths"
    addPaths.EntityData.BundleName = "openconfig"
    addPaths.EntityData.ParentYangName = "peer-group"
    addPaths.EntityData.SegmentPath = "add-paths"
    addPaths.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + addPaths.EntityData.SegmentPath
    addPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    addPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    addPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    addPaths.EntityData.Children = types.NewOrderedMap()
    addPaths.EntityData.Children.Append("config", types.YChild{"Config", &addPaths.Config})
    addPaths.EntityData.Children.Append("state", types.YChild{"State", &addPaths.State})
    addPaths.EntityData.Leafs = types.NewOrderedMap()

    addPaths.EntityData.YListKeys = []string {}

    return &(addPaths.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AddPaths_Config
// Configuration parameters relating to ADD_PATHS
type Bgp_PeerGroups_PeerGroup_AddPaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ability to receive multiple path advertisements for an NLRI from the
    // neighbor or group. The type is bool. The default value is false.
    Receive interface{}

    // The maximum number of paths to advertise to neighbors for a single NLRI.
    // The type is interface{} with range: 0..255.
    SendMax interface{}

    // A reference to a routing policy which can be used to restrict the prefixes
    // for which add-paths is enabled. The type is string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    EligiblePrefixPolicy interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "add-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/add-paths/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("receive", types.YLeaf{"Receive", config.Receive})
    config.EntityData.Leafs.Append("send-max", types.YLeaf{"SendMax", config.SendMax})
    config.EntityData.Leafs.Append("eligible-prefix-policy", types.YLeaf{"EligiblePrefixPolicy", config.EligiblePrefixPolicy})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AddPaths_State
// State information associated with ADD_PATHS
type Bgp_PeerGroups_PeerGroup_AddPaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ability to receive multiple path advertisements for an NLRI from the
    // neighbor or group. The type is bool. The default value is false.
    Receive interface{}

    // The maximum number of paths to advertise to neighbors for a single NLRI.
    // The type is interface{} with range: 0..255.
    SendMax interface{}

    // A reference to a routing policy which can be used to restrict the prefixes
    // for which add-paths is enabled. The type is string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    EligiblePrefixPolicy interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "add-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/add-paths/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("receive", types.YLeaf{"Receive", state.Receive})
    state.EntityData.Leafs.Append("send-max", types.YLeaf{"SendMax", state.SendMax})
    state.EntityData.Leafs.Append("eligible-prefix-policy", types.YLeaf{"EligiblePrefixPolicy", state.EligiblePrefixPolicy})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths
// Parameters related to the use of multiple paths for the
// same NLRI
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to multipath.
    Config Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config

    // State parameters relating to multipath.
    State Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State

    // Multipath parameters for eBGP.
    Ebgp Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp

    // Multipath parameters for iBGP.
    Ibgp Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetEntityData() *types.CommonEntityData {
    useMultiplePaths.EntityData.YFilter = useMultiplePaths.YFilter
    useMultiplePaths.EntityData.YangName = "use-multiple-paths"
    useMultiplePaths.EntityData.BundleName = "openconfig"
    useMultiplePaths.EntityData.ParentYangName = "peer-group"
    useMultiplePaths.EntityData.SegmentPath = "use-multiple-paths"
    useMultiplePaths.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + useMultiplePaths.EntityData.SegmentPath
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = types.NewOrderedMap()
    useMultiplePaths.EntityData.Children.Append("config", types.YChild{"Config", &useMultiplePaths.Config})
    useMultiplePaths.EntityData.Children.Append("state", types.YChild{"State", &useMultiplePaths.State})
    useMultiplePaths.EntityData.Children.Append("ebgp", types.YChild{"Ebgp", &useMultiplePaths.Ebgp})
    useMultiplePaths.EntityData.Children.Append("ibgp", types.YChild{"Ibgp", &useMultiplePaths.Ibgp})
    useMultiplePaths.EntityData.Leafs = types.NewOrderedMap()

    useMultiplePaths.EntityData.YListKeys = []string {}

    return &(useMultiplePaths.EntityData)
}

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "use-multiple-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/use-multiple-paths/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "use-multiple-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/use-multiple-paths/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp
// Multipath parameters for eBGP
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetEntityData() *types.CommonEntityData {
    ebgp.EntityData.YFilter = ebgp.YFilter
    ebgp.EntityData.YangName = "ebgp"
    ebgp.EntityData.BundleName = "openconfig"
    ebgp.EntityData.ParentYangName = "use-multiple-paths"
    ebgp.EntityData.SegmentPath = "ebgp"
    ebgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/use-multiple-paths/" + ebgp.EntityData.SegmentPath
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = types.NewOrderedMap()
    ebgp.EntityData.Children.Append("config", types.YChild{"Config", &ebgp.Config})
    ebgp.EntityData.Children.Append("state", types.YChild{"State", &ebgp.State})
    ebgp.EntityData.Leafs = types.NewOrderedMap()

    ebgp.EntityData.YListKeys = []string {}

    return &(ebgp.EntityData)
}

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}

    // Maximum number of parallel paths to consider when using BGP multipath. The
    // default is use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ebgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/use-multiple-paths/ebgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs})
    config.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", config.MaximumPaths})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}

    // Maximum number of parallel paths to consider when using BGP multipath. The
    // default is use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ebgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/use-multiple-paths/ebgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs})
    state.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", state.MaximumPaths})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp
// Multipath parameters for iBGP
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to iBGP multipath.
    Config Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config

    // State information relating to iBGP multipath.
    State Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State
}

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetEntityData() *types.CommonEntityData {
    ibgp.EntityData.YFilter = ibgp.YFilter
    ibgp.EntityData.YangName = "ibgp"
    ibgp.EntityData.BundleName = "openconfig"
    ibgp.EntityData.ParentYangName = "use-multiple-paths"
    ibgp.EntityData.SegmentPath = "ibgp"
    ibgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/use-multiple-paths/" + ibgp.EntityData.SegmentPath
    ibgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ibgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ibgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ibgp.EntityData.Children = types.NewOrderedMap()
    ibgp.EntityData.Children.Append("config", types.YChild{"Config", &ibgp.Config})
    ibgp.EntityData.Children.Append("state", types.YChild{"State", &ibgp.State})
    ibgp.EntityData.Leafs = types.NewOrderedMap()

    ibgp.EntityData.YListKeys = []string {}

    return &(ibgp.EntityData)
}

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config
// Configuration parameters relating to iBGP multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ibgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/use-multiple-paths/ibgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", config.MaximumPaths})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State
// State information relating to iBGP multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ibgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/use-multiple-paths/ibgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", state.MaximumPaths})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_PeerGroups_PeerGroup_ApplyPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_PeerGroups_PeerGroup_ApplyPolicy_State
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetEntityData() *types.CommonEntityData {
    applyPolicy.EntityData.YFilter = applyPolicy.YFilter
    applyPolicy.EntityData.YangName = "apply-policy"
    applyPolicy.EntityData.BundleName = "openconfig"
    applyPolicy.EntityData.ParentYangName = "peer-group"
    applyPolicy.EntityData.SegmentPath = "apply-policy"
    applyPolicy.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + applyPolicy.EntityData.SegmentPath
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = types.NewOrderedMap()
    applyPolicy.EntityData.Children.Append("config", types.YChild{"Config", &applyPolicy.Config})
    applyPolicy.EntityData.Children.Append("state", types.YChild{"State", &applyPolicy.State})
    applyPolicy.EntityData.Leafs = types.NewOrderedMap()

    applyPolicy.EntityData.YListKeys = []string {}

    return &(applyPolicy.EntityData)
}

// Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config
// Policy configuration data.
type Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of policy names in sequence to be applied on receiving a routing
    // update in the current context, e.g., for the current peer group, neighbor,
    // address family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ImportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the import
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultImportPolicy interface{}

    // list of policy names in sequence to be applied on sending a routing update
    // in the current context, e.g., for the current peer group, neighbor, address
    // family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ExportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the export
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultExportPolicy interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "apply-policy"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/apply-policy/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("import-policy", types.YLeaf{"ImportPolicy", config.ImportPolicy})
    config.EntityData.Leafs.Append("default-import-policy", types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy})
    config.EntityData.Leafs.Append("export-policy", types.YLeaf{"ExportPolicy", config.ExportPolicy})
    config.EntityData.Leafs.Append("default-export-policy", types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_ApplyPolicy_State
// Operational state for routing policy
type Bgp_PeerGroups_PeerGroup_ApplyPolicy_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of policy names in sequence to be applied on receiving a routing
    // update in the current context, e.g., for the current peer group, neighbor,
    // address family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ImportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the import
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultImportPolicy interface{}

    // list of policy names in sequence to be applied on sending a routing update
    // in the current context, e.g., for the current peer group, neighbor, address
    // family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ExportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the export
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultExportPolicy interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "apply-policy"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/apply-policy/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("import-policy", types.YLeaf{"ImportPolicy", state.ImportPolicy})
    state.EntityData.Leafs.Append("default-import-policy", types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy})
    state.EntityData.Leafs.Append("export-policy", types.YLeaf{"ExportPolicy", state.ExportPolicy})
    state.EntityData.Leafs.Append("default-export-policy", types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis
// Per-address-family configuration parameters associated with
// thegroup
type Bgp_PeerGroups_PeerGroup_AfiSafis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI configuration available for the neighbour or group. The type is
    // slice of Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi.
    AfiSafi []*Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi
}

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetEntityData() *types.CommonEntityData {
    afiSafis.EntityData.YFilter = afiSafis.YFilter
    afiSafis.EntityData.YangName = "afi-safis"
    afiSafis.EntityData.BundleName = "openconfig"
    afiSafis.EntityData.ParentYangName = "peer-group"
    afiSafis.EntityData.SegmentPath = "afi-safis"
    afiSafis.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/" + afiSafis.EntityData.SegmentPath
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

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi
// AFI,SAFI configuration available for the
// neighbour or group
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to the AFI-SAFI name used as a key for
    // the AFI-SAFI list. The type is one of the following:
    // IPV4UNICASTIPV6UNICASTIPV4LABELEDUNICASTIPV6LABELEDUNICASTL3VPNIPV4UNICASTL3VPNIPV6UNICASTL3VPNIPV4MULTICASTL3VPNIPV6MULTICASTL2VPNVPLSL2VPNEVPN.
    AfiSafiName interface{}

    // Configuration parameters for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config

    // State information relating to the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State

    // Parameters relating to BGP graceful-restart.
    GracefulRestart Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart

    // Parameters relating to options for route selection.
    RouteSelectionOptions Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions

    // Parameters related to the use of multiple paths for the same NLRI.
    UseMultiplePaths Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths

    // Anchor point for routing policies in the model. Import and export policies
    // are with respect to the local routing table, i.e., export (send) and import
    // (receive), depending on the context.
    ApplyPolicy Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy

    // IPv4 unicast configuration options.
    Ipv4Unicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast

    // IPv6 unicast configuration options.
    Ipv6Unicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast

    // IPv4 Labeled Unicast configuration options.
    Ipv4LabeledUnicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast

    // IPv6 Labeled Unicast configuration options.
    Ipv6LabeledUnicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast

    // Unicast IPv4 L3VPN configuration options.
    L3vpnIpv4Unicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast

    // Unicast IPv6 L3VPN configuration options.
    L3vpnIpv6Unicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast

    // Multicast IPv4 L3VPN configuration options.
    L3vpnIpv4Multicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast

    // Multicast IPv6 L3VPN configuration options.
    L3vpnIpv6Multicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast

    // BGP-signalled VPLS configuration options.
    L2vpnVpls Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls

    // BGP EVPN configuration options.
    L2vpnEvpn Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn
}

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetEntityData() *types.CommonEntityData {
    afiSafi.EntityData.YFilter = afiSafi.YFilter
    afiSafi.EntityData.YangName = "afi-safi"
    afiSafi.EntityData.BundleName = "openconfig"
    afiSafi.EntityData.ParentYangName = "afi-safis"
    afiSafi.EntityData.SegmentPath = "afi-safi" + types.AddKeyToken(afiSafi.AfiSafiName, "afi-safi-name")
    afiSafi.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/" + afiSafi.EntityData.SegmentPath
    afiSafi.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafi.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafi.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafi.EntityData.Children = types.NewOrderedMap()
    afiSafi.EntityData.Children.Append("config", types.YChild{"Config", &afiSafi.Config})
    afiSafi.EntityData.Children.Append("state", types.YChild{"State", &afiSafi.State})
    afiSafi.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &afiSafi.GracefulRestart})
    afiSafi.EntityData.Children.Append("route-selection-options", types.YChild{"RouteSelectionOptions", &afiSafi.RouteSelectionOptions})
    afiSafi.EntityData.Children.Append("use-multiple-paths", types.YChild{"UseMultiplePaths", &afiSafi.UseMultiplePaths})
    afiSafi.EntityData.Children.Append("apply-policy", types.YChild{"ApplyPolicy", &afiSafi.ApplyPolicy})
    afiSafi.EntityData.Children.Append("ipv4-unicast", types.YChild{"Ipv4Unicast", &afiSafi.Ipv4Unicast})
    afiSafi.EntityData.Children.Append("ipv6-unicast", types.YChild{"Ipv6Unicast", &afiSafi.Ipv6Unicast})
    afiSafi.EntityData.Children.Append("ipv4-labeled-unicast", types.YChild{"Ipv4LabeledUnicast", &afiSafi.Ipv4LabeledUnicast})
    afiSafi.EntityData.Children.Append("ipv6-labeled-unicast", types.YChild{"Ipv6LabeledUnicast", &afiSafi.Ipv6LabeledUnicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv4-unicast", types.YChild{"L3vpnIpv4Unicast", &afiSafi.L3vpnIpv4Unicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv6-unicast", types.YChild{"L3vpnIpv6Unicast", &afiSafi.L3vpnIpv6Unicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv4-multicast", types.YChild{"L3vpnIpv4Multicast", &afiSafi.L3vpnIpv4Multicast})
    afiSafi.EntityData.Children.Append("l3vpn-ipv6-multicast", types.YChild{"L3vpnIpv6Multicast", &afiSafi.L3vpnIpv6Multicast})
    afiSafi.EntityData.Children.Append("l2vpn-vpls", types.YChild{"L2vpnVpls", &afiSafi.L2vpnVpls})
    afiSafi.EntityData.Children.Append("l2vpn-evpn", types.YChild{"L2vpnEvpn", &afiSafi.L2vpnEvpn})
    afiSafi.EntityData.Leafs = types.NewOrderedMap()
    afiSafi.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", afiSafi.AfiSafiName})

    afiSafi.EntityData.YListKeys = []string {"AfiSafiName"}

    return &(afiSafi.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config
// Configuration parameters for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // IPV4UNICASTIPV6UNICASTIPV4LABELEDUNICASTIPV6LABELEDUNICASTL3VPNIPV4UNICASTL3VPNIPV6UNICASTL3VPNIPV4MULTICASTL3VPNIPV6MULTICASTL2VPNVPLSL2VPNEVPN.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "afi-safi"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", config.AfiSafiName})
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State
// State information relating to the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // IPV4UNICASTIPV6UNICASTIPV4LABELEDUNICASTIPV6LABELEDUNICASTL3VPNIPV4UNICASTL3VPNIPV6UNICASTL3VPNIPV4MULTICASTL3VPNIPV6MULTICASTL2VPNVPLSL2VPNEVPN.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "afi-safi"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("afi-safi-name", types.YLeaf{"AfiSafiName", state.AfiSafiName})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart
// Parameters relating to BGP graceful-restart
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration options for BGP graceful-restart.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config

    // State information for BGP graceful-restart.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "openconfig"
    gracefulRestart.EntityData.ParentYangName = "afi-safi"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("config", types.YChild{"Config", &gracefulRestart.Config})
    gracefulRestart.EntityData.Children.Append("state", types.YChild{"State", &gracefulRestart.State})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config
// Configuration options for BGP graceful-restart
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "graceful-restart"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/graceful-restart/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State
// State information for BGP graceful-restart
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "graceful-restart"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/graceful-restart/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions
// Parameters relating to options for route selection
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to route selection options.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config

    // State information for the route selection options.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State
}

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetEntityData() *types.CommonEntityData {
    routeSelectionOptions.EntityData.YFilter = routeSelectionOptions.YFilter
    routeSelectionOptions.EntityData.YangName = "route-selection-options"
    routeSelectionOptions.EntityData.BundleName = "openconfig"
    routeSelectionOptions.EntityData.ParentYangName = "afi-safi"
    routeSelectionOptions.EntityData.SegmentPath = "route-selection-options"
    routeSelectionOptions.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + routeSelectionOptions.EntityData.SegmentPath
    routeSelectionOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeSelectionOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeSelectionOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeSelectionOptions.EntityData.Children = types.NewOrderedMap()
    routeSelectionOptions.EntityData.Children.Append("config", types.YChild{"Config", &routeSelectionOptions.Config})
    routeSelectionOptions.EntityData.Children.Append("state", types.YChild{"State", &routeSelectionOptions.State})
    routeSelectionOptions.EntityData.Leafs = types.NewOrderedMap()

    routeSelectionOptions.EntityData.YListKeys = []string {}

    return &(routeSelectionOptions.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config
// Configuration parameters relating to route selection
// options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Compare multi-exit discriminator (MED) value from different ASes when
    // selecting the best route.  The default behavior is to only compare MEDs for
    // paths received from the same AS. The type is bool. The default value is
    // false.
    AlwaysCompareMed interface{}

    // Ignore the AS path length when selecting the best path. The default is to
    // use the AS path length and prefer paths with shorter length. The type is
    // bool. The default value is false.
    IgnoreAsPathLength interface{}

    // When comparing similar routes received from external BGP peers, use the
    // router-id as a criterion to select the active path. The type is bool. The
    // default value is true.
    ExternalCompareRouterId interface{}

    // Advertise inactive routes to external peers.  The default is to only
    // advertise active routes. The type is bool. The default value is false.
    AdvertiseInactiveRoutes interface{}

    // Flag to enable sending / receiving accumulated IGP attribute in routing
    // updates. The type is bool. The default value is false.
    EnableAigp interface{}

    // Ignore the IGP metric to the next-hop when calculating BGP best-path. The
    // default is to select the route for which the metric to the next-hop is
    // lowest. The type is bool. The default value is false.
    IgnoreNextHopIgpMetric interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "route-selection-options"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/route-selection-options/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("always-compare-med", types.YLeaf{"AlwaysCompareMed", config.AlwaysCompareMed})
    config.EntityData.Leafs.Append("ignore-as-path-length", types.YLeaf{"IgnoreAsPathLength", config.IgnoreAsPathLength})
    config.EntityData.Leafs.Append("external-compare-router-id", types.YLeaf{"ExternalCompareRouterId", config.ExternalCompareRouterId})
    config.EntityData.Leafs.Append("advertise-inactive-routes", types.YLeaf{"AdvertiseInactiveRoutes", config.AdvertiseInactiveRoutes})
    config.EntityData.Leafs.Append("enable-aigp", types.YLeaf{"EnableAigp", config.EnableAigp})
    config.EntityData.Leafs.Append("ignore-next-hop-igp-metric", types.YLeaf{"IgnoreNextHopIgpMetric", config.IgnoreNextHopIgpMetric})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State
// State information for the route selection options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Compare multi-exit discriminator (MED) value from different ASes when
    // selecting the best route.  The default behavior is to only compare MEDs for
    // paths received from the same AS. The type is bool. The default value is
    // false.
    AlwaysCompareMed interface{}

    // Ignore the AS path length when selecting the best path. The default is to
    // use the AS path length and prefer paths with shorter length. The type is
    // bool. The default value is false.
    IgnoreAsPathLength interface{}

    // When comparing similar routes received from external BGP peers, use the
    // router-id as a criterion to select the active path. The type is bool. The
    // default value is true.
    ExternalCompareRouterId interface{}

    // Advertise inactive routes to external peers.  The default is to only
    // advertise active routes. The type is bool. The default value is false.
    AdvertiseInactiveRoutes interface{}

    // Flag to enable sending / receiving accumulated IGP attribute in routing
    // updates. The type is bool. The default value is false.
    EnableAigp interface{}

    // Ignore the IGP metric to the next-hop when calculating BGP best-path. The
    // default is to select the route for which the metric to the next-hop is
    // lowest. The type is bool. The default value is false.
    IgnoreNextHopIgpMetric interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "route-selection-options"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/route-selection-options/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("always-compare-med", types.YLeaf{"AlwaysCompareMed", state.AlwaysCompareMed})
    state.EntityData.Leafs.Append("ignore-as-path-length", types.YLeaf{"IgnoreAsPathLength", state.IgnoreAsPathLength})
    state.EntityData.Leafs.Append("external-compare-router-id", types.YLeaf{"ExternalCompareRouterId", state.ExternalCompareRouterId})
    state.EntityData.Leafs.Append("advertise-inactive-routes", types.YLeaf{"AdvertiseInactiveRoutes", state.AdvertiseInactiveRoutes})
    state.EntityData.Leafs.Append("enable-aigp", types.YLeaf{"EnableAigp", state.EnableAigp})
    state.EntityData.Leafs.Append("ignore-next-hop-igp-metric", types.YLeaf{"IgnoreNextHopIgpMetric", state.IgnoreNextHopIgpMetric})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths
// Parameters related to the use of multiple paths for the
// same NLRI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to multipath.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config

    // State parameters relating to multipath.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State

    // Multipath parameters for eBGP.
    Ebgp Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp

    // Multipath parameters for iBGP.
    Ibgp Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetEntityData() *types.CommonEntityData {
    useMultiplePaths.EntityData.YFilter = useMultiplePaths.YFilter
    useMultiplePaths.EntityData.YangName = "use-multiple-paths"
    useMultiplePaths.EntityData.BundleName = "openconfig"
    useMultiplePaths.EntityData.ParentYangName = "afi-safi"
    useMultiplePaths.EntityData.SegmentPath = "use-multiple-paths"
    useMultiplePaths.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + useMultiplePaths.EntityData.SegmentPath
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = types.NewOrderedMap()
    useMultiplePaths.EntityData.Children.Append("config", types.YChild{"Config", &useMultiplePaths.Config})
    useMultiplePaths.EntityData.Children.Append("state", types.YChild{"State", &useMultiplePaths.State})
    useMultiplePaths.EntityData.Children.Append("ebgp", types.YChild{"Ebgp", &useMultiplePaths.Ebgp})
    useMultiplePaths.EntityData.Children.Append("ibgp", types.YChild{"Ibgp", &useMultiplePaths.Ibgp})
    useMultiplePaths.EntityData.Leafs = types.NewOrderedMap()

    useMultiplePaths.EntityData.YListKeys = []string {}

    return &(useMultiplePaths.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "use-multiple-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/use-multiple-paths/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "use-multiple-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/use-multiple-paths/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp
// Multipath parameters for eBGP
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetEntityData() *types.CommonEntityData {
    ebgp.EntityData.YFilter = ebgp.YFilter
    ebgp.EntityData.YangName = "ebgp"
    ebgp.EntityData.BundleName = "openconfig"
    ebgp.EntityData.ParentYangName = "use-multiple-paths"
    ebgp.EntityData.SegmentPath = "ebgp"
    ebgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/use-multiple-paths/" + ebgp.EntityData.SegmentPath
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = types.NewOrderedMap()
    ebgp.EntityData.Children.Append("config", types.YChild{"Config", &ebgp.Config})
    ebgp.EntityData.Children.Append("state", types.YChild{"State", &ebgp.State})
    ebgp.EntityData.Leafs = types.NewOrderedMap()

    ebgp.EntityData.YListKeys = []string {}

    return &(ebgp.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}

    // Maximum number of parallel paths to consider when using BGP multipath. The
    // default is use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ebgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/use-multiple-paths/ebgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs})
    config.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", config.MaximumPaths})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}

    // Maximum number of parallel paths to consider when using BGP multipath. The
    // default is use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ebgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/use-multiple-paths/ebgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("allow-multiple-as", types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs})
    state.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", state.MaximumPaths})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp
// Multipath parameters for iBGP
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to iBGP multipath.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config

    // State information relating to iBGP multipath.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State
}

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetEntityData() *types.CommonEntityData {
    ibgp.EntityData.YFilter = ibgp.YFilter
    ibgp.EntityData.YangName = "ibgp"
    ibgp.EntityData.BundleName = "openconfig"
    ibgp.EntityData.ParentYangName = "use-multiple-paths"
    ibgp.EntityData.SegmentPath = "ibgp"
    ibgp.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/use-multiple-paths/" + ibgp.EntityData.SegmentPath
    ibgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ibgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ibgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ibgp.EntityData.Children = types.NewOrderedMap()
    ibgp.EntityData.Children.Append("config", types.YChild{"Config", &ibgp.Config})
    ibgp.EntityData.Children.Append("state", types.YChild{"State", &ibgp.State})
    ibgp.EntityData.Leafs = types.NewOrderedMap()

    ibgp.EntityData.YListKeys = []string {}

    return &(ibgp.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config
// Configuration parameters relating to iBGP multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ibgp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/use-multiple-paths/ibgp/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", config.MaximumPaths})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State
// State information relating to iBGP multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ibgp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/use-multiple-paths/ibgp/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", state.MaximumPaths})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetEntityData() *types.CommonEntityData {
    applyPolicy.EntityData.YFilter = applyPolicy.YFilter
    applyPolicy.EntityData.YangName = "apply-policy"
    applyPolicy.EntityData.BundleName = "openconfig"
    applyPolicy.EntityData.ParentYangName = "afi-safi"
    applyPolicy.EntityData.SegmentPath = "apply-policy"
    applyPolicy.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + applyPolicy.EntityData.SegmentPath
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = types.NewOrderedMap()
    applyPolicy.EntityData.Children.Append("config", types.YChild{"Config", &applyPolicy.Config})
    applyPolicy.EntityData.Children.Append("state", types.YChild{"State", &applyPolicy.State})
    applyPolicy.EntityData.Leafs = types.NewOrderedMap()

    applyPolicy.EntityData.YListKeys = []string {}

    return &(applyPolicy.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config
// Policy configuration data.
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of policy names in sequence to be applied on receiving a routing
    // update in the current context, e.g., for the current peer group, neighbor,
    // address family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ImportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the import
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultImportPolicy interface{}

    // list of policy names in sequence to be applied on sending a routing update
    // in the current context, e.g., for the current peer group, neighbor, address
    // family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ExportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the export
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultExportPolicy interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "apply-policy"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/apply-policy/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("import-policy", types.YLeaf{"ImportPolicy", config.ImportPolicy})
    config.EntityData.Leafs.Append("default-import-policy", types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy})
    config.EntityData.Leafs.Append("export-policy", types.YLeaf{"ExportPolicy", config.ExportPolicy})
    config.EntityData.Leafs.Append("default-export-policy", types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State
// Operational state for routing policy
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of policy names in sequence to be applied on receiving a routing
    // update in the current context, e.g., for the current peer group, neighbor,
    // address family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ImportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the import
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultImportPolicy interface{}

    // list of policy names in sequence to be applied on sending a routing update
    // in the current context, e.g., for the current peer group, neighbor, address
    // family, etc. The type is slice of string. Refers to
    // routing_policy.RoutingPolicy_PolicyDefinitions_PolicyDefinition_Name
    ExportPolicy []interface{}

    // explicitly set a default policy if no policy definition in the export
    // policy chain is satisfied. The type is DefaultPolicyType. The default value
    // is REJECT_ROUTE.
    DefaultExportPolicy interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "apply-policy"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/apply-policy/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("import-policy", types.YLeaf{"ImportPolicy", state.ImportPolicy})
    state.EntityData.Leafs.Append("default-import-policy", types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy})
    state.EntityData.Leafs.Append("export-policy", types.YLeaf{"ExportPolicy", state.ExportPolicy})
    state.EntityData.Leafs.Append("default-export-policy", types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast
// IPv4 unicast configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State
}

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetEntityData() *types.CommonEntityData {
    ipv4Unicast.EntityData.YFilter = ipv4Unicast.YFilter
    ipv4Unicast.EntityData.YangName = "ipv4-unicast"
    ipv4Unicast.EntityData.BundleName = "openconfig"
    ipv4Unicast.EntityData.ParentYangName = "afi-safi"
    ipv4Unicast.EntityData.SegmentPath = "ipv4-unicast"
    ipv4Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + ipv4Unicast.EntityData.SegmentPath
    ipv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4Unicast.EntityData.Children = types.NewOrderedMap()
    ipv4Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv4Unicast.PrefixLimit})
    ipv4Unicast.EntityData.Children.Append("config", types.YChild{"Config", &ipv4Unicast.Config})
    ipv4Unicast.EntityData.Children.Append("state", types.YChild{"State", &ipv4Unicast.State})
    ipv4Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv4Unicast.EntityData.YListKeys = []string {}

    return &(ipv4Unicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv4-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv4-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv4-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv4-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv4-unicast"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv4-unicast/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv4-unicast"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv4-unicast/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast
// IPv6 unicast configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State
}

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetEntityData() *types.CommonEntityData {
    ipv6Unicast.EntityData.YFilter = ipv6Unicast.YFilter
    ipv6Unicast.EntityData.YangName = "ipv6-unicast"
    ipv6Unicast.EntityData.BundleName = "openconfig"
    ipv6Unicast.EntityData.ParentYangName = "afi-safi"
    ipv6Unicast.EntityData.SegmentPath = "ipv6-unicast"
    ipv6Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + ipv6Unicast.EntityData.SegmentPath
    ipv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6Unicast.EntityData.Children = types.NewOrderedMap()
    ipv6Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv6Unicast.PrefixLimit})
    ipv6Unicast.EntityData.Children.Append("config", types.YChild{"Config", &ipv6Unicast.Config})
    ipv6Unicast.EntityData.Children.Append("state", types.YChild{"State", &ipv6Unicast.State})
    ipv6Unicast.EntityData.Leafs = types.NewOrderedMap()

    ipv6Unicast.EntityData.YListKeys = []string {}

    return &(ipv6Unicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv6-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv6-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv6-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv6-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv6-unicast"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv6-unicast/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv6-unicast"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv6-unicast/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("send-default-route", types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast
// IPv4 Labeled Unicast configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetEntityData() *types.CommonEntityData {
    ipv4LabeledUnicast.EntityData.YFilter = ipv4LabeledUnicast.YFilter
    ipv4LabeledUnicast.EntityData.YangName = "ipv4-labeled-unicast"
    ipv4LabeledUnicast.EntityData.BundleName = "openconfig"
    ipv4LabeledUnicast.EntityData.ParentYangName = "afi-safi"
    ipv4LabeledUnicast.EntityData.SegmentPath = "ipv4-labeled-unicast"
    ipv4LabeledUnicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + ipv4LabeledUnicast.EntityData.SegmentPath
    ipv4LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4LabeledUnicast.EntityData.Children = types.NewOrderedMap()
    ipv4LabeledUnicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv4LabeledUnicast.PrefixLimit})
    ipv4LabeledUnicast.EntityData.Leafs = types.NewOrderedMap()

    ipv4LabeledUnicast.EntityData.YListKeys = []string {}

    return &(ipv4LabeledUnicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv4-labeled-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv4-labeled-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv4-labeled-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv4-labeled-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast
// IPv6 Labeled Unicast configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetEntityData() *types.CommonEntityData {
    ipv6LabeledUnicast.EntityData.YFilter = ipv6LabeledUnicast.YFilter
    ipv6LabeledUnicast.EntityData.YangName = "ipv6-labeled-unicast"
    ipv6LabeledUnicast.EntityData.BundleName = "openconfig"
    ipv6LabeledUnicast.EntityData.ParentYangName = "afi-safi"
    ipv6LabeledUnicast.EntityData.SegmentPath = "ipv6-labeled-unicast"
    ipv6LabeledUnicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + ipv6LabeledUnicast.EntityData.SegmentPath
    ipv6LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6LabeledUnicast.EntityData.Children = types.NewOrderedMap()
    ipv6LabeledUnicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &ipv6LabeledUnicast.PrefixLimit})
    ipv6LabeledUnicast.EntityData.Leafs = types.NewOrderedMap()

    ipv6LabeledUnicast.EntityData.YListKeys = []string {}

    return &(ipv6LabeledUnicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "ipv6-labeled-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv6-labeled-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv6-labeled-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/ipv6-labeled-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast
// Unicast IPv4 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit
}

func (l3vpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv4Unicast.EntityData.YFilter = l3vpnIpv4Unicast.YFilter
    l3vpnIpv4Unicast.EntityData.YangName = "l3vpn-ipv4-unicast"
    l3vpnIpv4Unicast.EntityData.BundleName = "openconfig"
    l3vpnIpv4Unicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv4Unicast.EntityData.SegmentPath = "l3vpn-ipv4-unicast"
    l3vpnIpv4Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + l3vpnIpv4Unicast.EntityData.SegmentPath
    l3vpnIpv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv4Unicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv4Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv4Unicast.PrefixLimit})
    l3vpnIpv4Unicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv4Unicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv4Unicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv4-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv4-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv4-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast
// Unicast IPv6 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit
}

func (l3vpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv6Unicast.EntityData.YFilter = l3vpnIpv6Unicast.YFilter
    l3vpnIpv6Unicast.EntityData.YangName = "l3vpn-ipv6-unicast"
    l3vpnIpv6Unicast.EntityData.BundleName = "openconfig"
    l3vpnIpv6Unicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv6Unicast.EntityData.SegmentPath = "l3vpn-ipv6-unicast"
    l3vpnIpv6Unicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + l3vpnIpv6Unicast.EntityData.SegmentPath
    l3vpnIpv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv6Unicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv6Unicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv6Unicast.PrefixLimit})
    l3vpnIpv6Unicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv6Unicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv6Unicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv6-unicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv6-unicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv6-unicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast
// Multicast IPv4 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit
}

func (l3vpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv4Multicast.EntityData.YFilter = l3vpnIpv4Multicast.YFilter
    l3vpnIpv4Multicast.EntityData.YangName = "l3vpn-ipv4-multicast"
    l3vpnIpv4Multicast.EntityData.BundleName = "openconfig"
    l3vpnIpv4Multicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv4Multicast.EntityData.SegmentPath = "l3vpn-ipv4-multicast"
    l3vpnIpv4Multicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + l3vpnIpv4Multicast.EntityData.SegmentPath
    l3vpnIpv4Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv4Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv4Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv4Multicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv4Multicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv4Multicast.PrefixLimit})
    l3vpnIpv4Multicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv4Multicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv4Multicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv4-multicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv4-multicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv4Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv4-multicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast
// Multicast IPv6 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit
}

func (l3vpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast) GetEntityData() *types.CommonEntityData {
    l3vpnIpv6Multicast.EntityData.YFilter = l3vpnIpv6Multicast.YFilter
    l3vpnIpv6Multicast.EntityData.YangName = "l3vpn-ipv6-multicast"
    l3vpnIpv6Multicast.EntityData.BundleName = "openconfig"
    l3vpnIpv6Multicast.EntityData.ParentYangName = "afi-safi"
    l3vpnIpv6Multicast.EntityData.SegmentPath = "l3vpn-ipv6-multicast"
    l3vpnIpv6Multicast.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + l3vpnIpv6Multicast.EntityData.SegmentPath
    l3vpnIpv6Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3vpnIpv6Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3vpnIpv6Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3vpnIpv6Multicast.EntityData.Children = types.NewOrderedMap()
    l3vpnIpv6Multicast.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l3vpnIpv6Multicast.PrefixLimit})
    l3vpnIpv6Multicast.EntityData.Leafs = types.NewOrderedMap()

    l3vpnIpv6Multicast.EntityData.YListKeys = []string {}

    return &(l3vpnIpv6Multicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv6-multicast/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv6-multicast/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3vpnIpv6Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l3vpn-ipv6-multicast/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls
// BGP-signalled VPLS configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit
}

func (l2vpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls) GetEntityData() *types.CommonEntityData {
    l2vpnVpls.EntityData.YFilter = l2vpnVpls.YFilter
    l2vpnVpls.EntityData.YangName = "l2vpn-vpls"
    l2vpnVpls.EntityData.BundleName = "openconfig"
    l2vpnVpls.EntityData.ParentYangName = "afi-safi"
    l2vpnVpls.EntityData.SegmentPath = "l2vpn-vpls"
    l2vpnVpls.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + l2vpnVpls.EntityData.SegmentPath
    l2vpnVpls.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2vpnVpls.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2vpnVpls.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2vpnVpls.EntityData.Children = types.NewOrderedMap()
    l2vpnVpls.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l2vpnVpls.PrefixLimit})
    l2vpnVpls.EntityData.Leafs = types.NewOrderedMap()

    l2vpnVpls.EntityData.YListKeys = []string {}

    return &(l2vpnVpls.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-vpls"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l2vpn-vpls/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l2vpn-vpls/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnVpls_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l2vpn-vpls/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn
// BGP EVPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit
}

func (l2vpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn) GetEntityData() *types.CommonEntityData {
    l2vpnEvpn.EntityData.YFilter = l2vpnEvpn.YFilter
    l2vpnEvpn.EntityData.YangName = "l2vpn-evpn"
    l2vpnEvpn.EntityData.BundleName = "openconfig"
    l2vpnEvpn.EntityData.ParentYangName = "afi-safi"
    l2vpnEvpn.EntityData.SegmentPath = "l2vpn-evpn"
    l2vpnEvpn.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/" + l2vpnEvpn.EntityData.SegmentPath
    l2vpnEvpn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2vpnEvpn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2vpnEvpn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2vpnEvpn.EntityData.Children = types.NewOrderedMap()
    l2vpnEvpn.EntityData.Children.Append("prefix-limit", types.YChild{"PrefixLimit", &l2vpnEvpn.PrefixLimit})
    l2vpnEvpn.EntityData.Leafs = types.NewOrderedMap()

    l2vpnEvpn.EntityData.YListKeys = []string {}

    return &(l2vpnEvpn.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-evpn"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l2vpn-evpn/" + prefixLimit.EntityData.SegmentPath
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = types.NewOrderedMap()
    prefixLimit.EntityData.Children.Append("config", types.YChild{"Config", &prefixLimit.Config})
    prefixLimit.EntityData.Children.Append("state", types.YChild{"State", &prefixLimit.State})
    prefixLimit.EntityData.Leafs = types.NewOrderedMap()

    prefixLimit.EntityData.YListKeys = []string {}

    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l2vpn-evpn/prefix-limit/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", config.MaxPrefixes})
    config.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", config.PreventTeardown})
    config.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct})
    config.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", config.RestartTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

    // Do not tear down the BGP session when the maximum prefix limit is exceeded,
    // but rather only log a warning. The default of this leaf is false, such that
    // when it is not specified, the session is torn down. The type is bool. The
    // default value is false.
    PreventTeardown interface{}

    // Threshold on number of prefixes that can be received from a neighbour
    // before generation of warning messages or log entries. Expressed as a
    // percentage of max-prefixes. The type is interface{} with range: 0..100.
    ShutdownThresholdPct interface{}

    // Time interval in seconds after which the BGP session is re-established
    // after being torn down due to exceeding the max-prefix limit. The type is
    // string with range: -92233720368547758.08..92233720368547758.07. Units are
    // seconds.
    RestartTimer interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2vpnEvpn_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "openconfig-bgp:bgp/peer-groups/peer-group/afi-safis/afi-safi/l2vpn-evpn/prefix-limit/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("max-prefixes", types.YLeaf{"MaxPrefixes", state.MaxPrefixes})
    state.EntityData.Leafs.Append("prevent-teardown", types.YLeaf{"PreventTeardown", state.PreventTeardown})
    state.EntityData.Leafs.Append("shutdown-threshold-pct", types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct})
    state.EntityData.Leafs.Append("restart-timer", types.YLeaf{"RestartTimer", state.RestartTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

