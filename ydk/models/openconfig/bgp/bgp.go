// This module describes a YANG model for BGP protocol
// configuration.It is a limited subset of all of the configuration
// parameters available in the variety of vendor implementations,
// hence it is expected that it would be augmented with vendor-
// specific configuration data as needed. Additional modules or
// submodules to handle other aspects of BGP configuration,
// including policy, VRFs, VPNs, and additional address families
// are also expected.
// This model supports the following BGP configuration level
// hierarchy:
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
    bgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Children["global"] = types.YChild{"Global", &bgp.Global}
    bgp.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &bgp.Neighbors}
    bgp.EntityData.Children["peer-groups"] = types.YChild{"PeerGroups", &bgp.PeerGroups}
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
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

    // Anchor point for routing policies in the model. Import and export policies
    // are with respect to the local routing table, i.e., export (send) and import
    // (receive), depending on the context.
    ApplyPolicy Bgp_Global_ApplyPolicy
}

func (global *Bgp_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "openconfig"
    global.EntityData.ParentYangName = "bgp"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    global.EntityData.NamespaceTable = openconfig.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["config"] = types.YChild{"Config", &global.Config}
    global.EntityData.Children["state"] = types.YChild{"State", &global.State}
    global.EntityData.Children["default-route-distance"] = types.YChild{"DefaultRouteDistance", &global.DefaultRouteDistance}
    global.EntityData.Children["confederation"] = types.YChild{"Confederation", &global.Confederation}
    global.EntityData.Children["graceful-restart"] = types.YChild{"GracefulRestart", &global.GracefulRestart}
    global.EntityData.Children["use-multiple-paths"] = types.YChild{"UseMultiplePaths", &global.UseMultiplePaths}
    global.EntityData.Children["route-selection-options"] = types.YChild{"RouteSelectionOptions", &global.RouteSelectionOptions}
    global.EntityData.Children["afi-safis"] = types.YChild{"AfiSafis", &global.AfiSafis}
    global.EntityData.Children["apply-policy"] = types.YChild{"ApplyPolicy", &global.ApplyPolicy}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
    RouterId interface{}
}

func (config *Bgp_Global_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "global"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["as"] = types.YLeaf{"As", config.As}
    config.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", config.RouterId}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])'.
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["as"] = types.YLeaf{"As", state.As}
    state.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", state.RouterId}
    state.EntityData.Leafs["total-paths"] = types.YLeaf{"TotalPaths", state.TotalPaths}
    state.EntityData.Leafs["total-prefixes"] = types.YLeaf{"TotalPrefixes", state.TotalPrefixes}
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
    defaultRouteDistance.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    defaultRouteDistance.EntityData.NamespaceTable = openconfig.GetNamespaces()
    defaultRouteDistance.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    defaultRouteDistance.EntityData.Children = make(map[string]types.YChild)
    defaultRouteDistance.EntityData.Children["config"] = types.YChild{"Config", &defaultRouteDistance.Config}
    defaultRouteDistance.EntityData.Children["state"] = types.YChild{"State", &defaultRouteDistance.State}
    defaultRouteDistance.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["external-route-distance"] = types.YLeaf{"ExternalRouteDistance", config.ExternalRouteDistance}
    config.EntityData.Leafs["internal-route-distance"] = types.YLeaf{"InternalRouteDistance", config.InternalRouteDistance}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["external-route-distance"] = types.YLeaf{"ExternalRouteDistance", state.ExternalRouteDistance}
    state.EntityData.Leafs["internal-route-distance"] = types.YLeaf{"InternalRouteDistance", state.InternalRouteDistance}
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
    confederation.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    confederation.EntityData.NamespaceTable = openconfig.GetNamespaces()
    confederation.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    confederation.EntityData.Children = make(map[string]types.YChild)
    confederation.EntityData.Children["config"] = types.YChild{"Config", &confederation.Config}
    confederation.EntityData.Children["state"] = types.YChild{"State", &confederation.State}
    confederation.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["identifier"] = types.YLeaf{"Identifier", config.Identifier}
    config.EntityData.Leafs["member-as"] = types.YLeaf{"MemberAs", config.MemberAs}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["identifier"] = types.YLeaf{"Identifier", state.Identifier}
    state.EntityData.Leafs["member-as"] = types.YLeaf{"MemberAs", state.MemberAs}
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
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = make(map[string]types.YChild)
    gracefulRestart.EntityData.Children["config"] = types.YChild{"Config", &gracefulRestart.Config}
    gracefulRestart.EntityData.Children["state"] = types.YChild{"State", &gracefulRestart.State}
    gracefulRestart.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["restart-time"] = types.YLeaf{"RestartTime", config.RestartTime}
    config.EntityData.Leafs["stale-routes-time"] = types.YLeaf{"StaleRoutesTime", config.StaleRoutesTime}
    config.EntityData.Leafs["helper-only"] = types.YLeaf{"HelperOnly", config.HelperOnly}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["restart-time"] = types.YLeaf{"RestartTime", state.RestartTime}
    state.EntityData.Leafs["stale-routes-time"] = types.YLeaf{"StaleRoutesTime", state.StaleRoutesTime}
    state.EntityData.Leafs["helper-only"] = types.YLeaf{"HelperOnly", state.HelperOnly}
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
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = make(map[string]types.YChild)
    useMultiplePaths.EntityData.Children["config"] = types.YChild{"Config", &useMultiplePaths.Config}
    useMultiplePaths.EntityData.Children["state"] = types.YChild{"State", &useMultiplePaths.State}
    useMultiplePaths.EntityData.Children["ebgp"] = types.YChild{"Ebgp", &useMultiplePaths.Ebgp}
    useMultiplePaths.EntityData.Children["ibgp"] = types.YChild{"Ibgp", &useMultiplePaths.Ibgp}
    useMultiplePaths.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
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
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = make(map[string]types.YChild)
    ebgp.EntityData.Children["config"] = types.YChild{"Config", &ebgp.Config}
    ebgp.EntityData.Children["state"] = types.YChild{"State", &ebgp.State}
    ebgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs}
    config.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", config.MaximumPaths}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs}
    state.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", state.MaximumPaths}
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
    ibgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ibgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ibgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ibgp.EntityData.Children = make(map[string]types.YChild)
    ibgp.EntityData.Children["config"] = types.YChild{"Config", &ibgp.Config}
    ibgp.EntityData.Children["state"] = types.YChild{"State", &ibgp.State}
    ibgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", config.MaximumPaths}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", state.MaximumPaths}
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
    routeSelectionOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeSelectionOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeSelectionOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeSelectionOptions.EntityData.Children = make(map[string]types.YChild)
    routeSelectionOptions.EntityData.Children["config"] = types.YChild{"Config", &routeSelectionOptions.Config}
    routeSelectionOptions.EntityData.Children["state"] = types.YChild{"State", &routeSelectionOptions.State}
    routeSelectionOptions.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["always-compare-med"] = types.YLeaf{"AlwaysCompareMed", config.AlwaysCompareMed}
    config.EntityData.Leafs["ignore-as-path-length"] = types.YLeaf{"IgnoreAsPathLength", config.IgnoreAsPathLength}
    config.EntityData.Leafs["external-compare-router-id"] = types.YLeaf{"ExternalCompareRouterId", config.ExternalCompareRouterId}
    config.EntityData.Leafs["advertise-inactive-routes"] = types.YLeaf{"AdvertiseInactiveRoutes", config.AdvertiseInactiveRoutes}
    config.EntityData.Leafs["enable-aigp"] = types.YLeaf{"EnableAigp", config.EnableAigp}
    config.EntityData.Leafs["ignore-next-hop-igp-metric"] = types.YLeaf{"IgnoreNextHopIgpMetric", config.IgnoreNextHopIgpMetric}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["always-compare-med"] = types.YLeaf{"AlwaysCompareMed", state.AlwaysCompareMed}
    state.EntityData.Leafs["ignore-as-path-length"] = types.YLeaf{"IgnoreAsPathLength", state.IgnoreAsPathLength}
    state.EntityData.Leafs["external-compare-router-id"] = types.YLeaf{"ExternalCompareRouterId", state.ExternalCompareRouterId}
    state.EntityData.Leafs["advertise-inactive-routes"] = types.YLeaf{"AdvertiseInactiveRoutes", state.AdvertiseInactiveRoutes}
    state.EntityData.Leafs["enable-aigp"] = types.YLeaf{"EnableAigp", state.EnableAigp}
    state.EntityData.Leafs["ignore-next-hop-igp-metric"] = types.YLeaf{"IgnoreNextHopIgpMetric", state.IgnoreNextHopIgpMetric}
    return &(state.EntityData)
}

// Bgp_Global_AfiSafis
// Address family specific configuration
type Bgp_Global_AfiSafis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI,SAFI configuration available for the neighbour or group. The type is
    // slice of Bgp_Global_AfiSafis_AfiSafi.
    AfiSafi []Bgp_Global_AfiSafis_AfiSafi
}

func (afiSafis *Bgp_Global_AfiSafis) GetEntityData() *types.CommonEntityData {
    afiSafis.EntityData.YFilter = afiSafis.YFilter
    afiSafis.EntityData.YangName = "afi-safis"
    afiSafis.EntityData.BundleName = "openconfig"
    afiSafis.EntityData.ParentYangName = "global"
    afiSafis.EntityData.SegmentPath = "afi-safis"
    afiSafis.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafis.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafis.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafis.EntityData.Children = make(map[string]types.YChild)
    afiSafis.EntityData.Children["afi-safi"] = types.YChild{"AfiSafi", nil}
    for i := range afiSafis.AfiSafi {
        afiSafis.EntityData.Children[types.GetSegmentPath(&afiSafis.AfiSafi[i])] = types.YChild{"AfiSafi", &afiSafis.AfiSafi[i]}
    }
    afiSafis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(afiSafis.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi
// AFI,SAFI configuration available for the
// neighbour or group
type Bgp_Global_AfiSafis_AfiSafi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    // Anchor point for routing policies in the model. Import and export policies
    // are with respect to the local routing table, i.e., export (send) and import
    // (receive), depending on the context.
    ApplyPolicy Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy

    // IPv4 unicast configuration options.
    Ipv4Unicast Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast

    // IPv6 unicast configuration options.
    Ipv6Unicast Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast

    // IPv4 Labeled Unicast configuration options.
    Ipv4LabeledUnicast Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast

    // IPv6 Labeled Unicast configuration options.
    Ipv6LabeledUnicast Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast

    // Unicast IPv4 L3VPN configuration options.
    L3VpnIpv4Unicast Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast

    // Unicast IPv6 L3VPN configuration options.
    L3VpnIpv6Unicast Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast

    // Multicast IPv4 L3VPN configuration options.
    L3VpnIpv4Multicast Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast

    // Multicast IPv6 L3VPN configuration options.
    L3VpnIpv6Multicast Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast

    // BGP-signalled VPLS configuration options.
    L2VpnVpls Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls

    // BGP EVPN configuration options.
    L2VpnEvpn Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn
}

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetEntityData() *types.CommonEntityData {
    afiSafi.EntityData.YFilter = afiSafi.YFilter
    afiSafi.EntityData.YangName = "afi-safi"
    afiSafi.EntityData.BundleName = "openconfig"
    afiSafi.EntityData.ParentYangName = "afi-safis"
    afiSafi.EntityData.SegmentPath = "afi-safi" + "[afi-safi-name='" + fmt.Sprintf("%v", afiSafi.AfiSafiName) + "']"
    afiSafi.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafi.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafi.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafi.EntityData.Children = make(map[string]types.YChild)
    afiSafi.EntityData.Children["config"] = types.YChild{"Config", &afiSafi.Config}
    afiSafi.EntityData.Children["state"] = types.YChild{"State", &afiSafi.State}
    afiSafi.EntityData.Children["graceful-restart"] = types.YChild{"GracefulRestart", &afiSafi.GracefulRestart}
    afiSafi.EntityData.Children["route-selection-options"] = types.YChild{"RouteSelectionOptions", &afiSafi.RouteSelectionOptions}
    afiSafi.EntityData.Children["use-multiple-paths"] = types.YChild{"UseMultiplePaths", &afiSafi.UseMultiplePaths}
    afiSafi.EntityData.Children["apply-policy"] = types.YChild{"ApplyPolicy", &afiSafi.ApplyPolicy}
    afiSafi.EntityData.Children["ipv4-unicast"] = types.YChild{"Ipv4Unicast", &afiSafi.Ipv4Unicast}
    afiSafi.EntityData.Children["ipv6-unicast"] = types.YChild{"Ipv6Unicast", &afiSafi.Ipv6Unicast}
    afiSafi.EntityData.Children["ipv4-labeled-unicast"] = types.YChild{"Ipv4LabeledUnicast", &afiSafi.Ipv4LabeledUnicast}
    afiSafi.EntityData.Children["ipv6-labeled-unicast"] = types.YChild{"Ipv6LabeledUnicast", &afiSafi.Ipv6LabeledUnicast}
    afiSafi.EntityData.Children["l3vpn-ipv4-unicast"] = types.YChild{"L3VpnIpv4Unicast", &afiSafi.L3VpnIpv4Unicast}
    afiSafi.EntityData.Children["l3vpn-ipv6-unicast"] = types.YChild{"L3VpnIpv6Unicast", &afiSafi.L3VpnIpv6Unicast}
    afiSafi.EntityData.Children["l3vpn-ipv4-multicast"] = types.YChild{"L3VpnIpv4Multicast", &afiSafi.L3VpnIpv4Multicast}
    afiSafi.EntityData.Children["l3vpn-ipv6-multicast"] = types.YChild{"L3VpnIpv6Multicast", &afiSafi.L3VpnIpv6Multicast}
    afiSafi.EntityData.Children["l2vpn-vpls"] = types.YChild{"L2VpnVpls", &afiSafi.L2VpnVpls}
    afiSafi.EntityData.Children["l2vpn-evpn"] = types.YChild{"L2VpnEvpn", &afiSafi.L2VpnEvpn}
    afiSafi.EntityData.Leafs = make(map[string]types.YLeaf)
    afiSafi.EntityData.Leafs["afi-safi-name"] = types.YLeaf{"AfiSafiName", afiSafi.AfiSafiName}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["afi-safi-name"] = types.YLeaf{"AfiSafiName", config.AfiSafiName}
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["afi-safi-name"] = types.YLeaf{"AfiSafiName", state.AfiSafiName}
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["total-paths"] = types.YLeaf{"TotalPaths", state.TotalPaths}
    state.EntityData.Leafs["total-prefixes"] = types.YLeaf{"TotalPrefixes", state.TotalPrefixes}
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
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = make(map[string]types.YChild)
    gracefulRestart.EntityData.Children["config"] = types.YChild{"Config", &gracefulRestart.Config}
    gracefulRestart.EntityData.Children["state"] = types.YChild{"State", &gracefulRestart.State}
    gracefulRestart.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
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
    routeSelectionOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeSelectionOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeSelectionOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeSelectionOptions.EntityData.Children = make(map[string]types.YChild)
    routeSelectionOptions.EntityData.Children["config"] = types.YChild{"Config", &routeSelectionOptions.Config}
    routeSelectionOptions.EntityData.Children["state"] = types.YChild{"State", &routeSelectionOptions.State}
    routeSelectionOptions.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["always-compare-med"] = types.YLeaf{"AlwaysCompareMed", config.AlwaysCompareMed}
    config.EntityData.Leafs["ignore-as-path-length"] = types.YLeaf{"IgnoreAsPathLength", config.IgnoreAsPathLength}
    config.EntityData.Leafs["external-compare-router-id"] = types.YLeaf{"ExternalCompareRouterId", config.ExternalCompareRouterId}
    config.EntityData.Leafs["advertise-inactive-routes"] = types.YLeaf{"AdvertiseInactiveRoutes", config.AdvertiseInactiveRoutes}
    config.EntityData.Leafs["enable-aigp"] = types.YLeaf{"EnableAigp", config.EnableAigp}
    config.EntityData.Leafs["ignore-next-hop-igp-metric"] = types.YLeaf{"IgnoreNextHopIgpMetric", config.IgnoreNextHopIgpMetric}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["always-compare-med"] = types.YLeaf{"AlwaysCompareMed", state.AlwaysCompareMed}
    state.EntityData.Leafs["ignore-as-path-length"] = types.YLeaf{"IgnoreAsPathLength", state.IgnoreAsPathLength}
    state.EntityData.Leafs["external-compare-router-id"] = types.YLeaf{"ExternalCompareRouterId", state.ExternalCompareRouterId}
    state.EntityData.Leafs["advertise-inactive-routes"] = types.YLeaf{"AdvertiseInactiveRoutes", state.AdvertiseInactiveRoutes}
    state.EntityData.Leafs["enable-aigp"] = types.YLeaf{"EnableAigp", state.EnableAigp}
    state.EntityData.Leafs["ignore-next-hop-igp-metric"] = types.YLeaf{"IgnoreNextHopIgpMetric", state.IgnoreNextHopIgpMetric}
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
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = make(map[string]types.YChild)
    useMultiplePaths.EntityData.Children["config"] = types.YChild{"Config", &useMultiplePaths.Config}
    useMultiplePaths.EntityData.Children["state"] = types.YChild{"State", &useMultiplePaths.State}
    useMultiplePaths.EntityData.Children["ebgp"] = types.YChild{"Ebgp", &useMultiplePaths.Ebgp}
    useMultiplePaths.EntityData.Children["ibgp"] = types.YChild{"Ibgp", &useMultiplePaths.Ibgp}
    useMultiplePaths.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
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
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = make(map[string]types.YChild)
    ebgp.EntityData.Children["config"] = types.YChild{"Config", &ebgp.Config}
    ebgp.EntityData.Children["state"] = types.YChild{"State", &ebgp.State}
    ebgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs}
    config.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", config.MaximumPaths}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs}
    state.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", state.MaximumPaths}
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
    ibgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ibgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ibgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ibgp.EntityData.Children = make(map[string]types.YChild)
    ibgp.EntityData.Children["config"] = types.YChild{"Config", &ibgp.Config}
    ibgp.EntityData.Children["state"] = types.YChild{"State", &ibgp.State}
    ibgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", config.MaximumPaths}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", state.MaximumPaths}
    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State
}

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetEntityData() *types.CommonEntityData {
    applyPolicy.EntityData.YFilter = applyPolicy.YFilter
    applyPolicy.EntityData.YangName = "apply-policy"
    applyPolicy.EntityData.BundleName = "openconfig"
    applyPolicy.EntityData.ParentYangName = "afi-safi"
    applyPolicy.EntityData.SegmentPath = "apply-policy"
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = make(map[string]types.YChild)
    applyPolicy.EntityData.Children["config"] = types.YChild{"Config", &applyPolicy.Config}
    applyPolicy.EntityData.Children["state"] = types.YChild{"State", &applyPolicy.State}
    applyPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(applyPolicy.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config
// Policy configuration data.
type Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config struct {
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

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "apply-policy"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", config.ImportPolicy}
    config.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy}
    config.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", config.ExportPolicy}
    config.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy}
    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State
// Operational state for routing policy
type Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State struct {
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

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "apply-policy"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", state.ImportPolicy}
    state.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy}
    state.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", state.ExportPolicy}
    state.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy}
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
    ipv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4Unicast.EntityData.Children = make(map[string]types.YChild)
    ipv4Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv4Unicast.PrefixLimit}
    ipv4Unicast.EntityData.Children["config"] = types.YChild{"Config", &ipv4Unicast.Config}
    ipv4Unicast.EntityData.Children["state"] = types.YChild{"State", &ipv4Unicast.State}
    ipv4Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute}
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
    ipv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6Unicast.EntityData.Children = make(map[string]types.YChild)
    ipv6Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv6Unicast.PrefixLimit}
    ipv6Unicast.EntityData.Children["config"] = types.YChild{"Config", &ipv6Unicast.Config}
    ipv6Unicast.EntityData.Children["state"] = types.YChild{"State", &ipv6Unicast.State}
    ipv6Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute}
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
    ipv4LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4LabeledUnicast.EntityData.Children = make(map[string]types.YChild)
    ipv4LabeledUnicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv4LabeledUnicast.PrefixLimit}
    ipv4LabeledUnicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    ipv6LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6LabeledUnicast.EntityData.Children = make(map[string]types.YChild)
    ipv6LabeledUnicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv6LabeledUnicast.PrefixLimit}
    ipv6LabeledUnicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast
// Unicast IPv4 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv4Unicast.EntityData.YFilter = l3VpnIpv4Unicast.YFilter
    l3VpnIpv4Unicast.EntityData.YangName = "l3vpn-ipv4-unicast"
    l3VpnIpv4Unicast.EntityData.BundleName = "openconfig"
    l3VpnIpv4Unicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv4Unicast.EntityData.SegmentPath = "l3vpn-ipv4-unicast"
    l3VpnIpv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv4Unicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv4Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv4Unicast.PrefixLimit}
    l3VpnIpv4Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv4Unicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast
// Unicast IPv6 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv6Unicast.EntityData.YFilter = l3VpnIpv6Unicast.YFilter
    l3VpnIpv6Unicast.EntityData.YangName = "l3vpn-ipv6-unicast"
    l3VpnIpv6Unicast.EntityData.BundleName = "openconfig"
    l3VpnIpv6Unicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv6Unicast.EntityData.SegmentPath = "l3vpn-ipv6-unicast"
    l3VpnIpv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv6Unicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv6Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv6Unicast.PrefixLimit}
    l3VpnIpv6Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv6Unicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast
// Multicast IPv4 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv4Multicast.EntityData.YFilter = l3VpnIpv4Multicast.YFilter
    l3VpnIpv4Multicast.EntityData.YangName = "l3vpn-ipv4-multicast"
    l3VpnIpv4Multicast.EntityData.BundleName = "openconfig"
    l3VpnIpv4Multicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv4Multicast.EntityData.SegmentPath = "l3vpn-ipv4-multicast"
    l3VpnIpv4Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv4Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv4Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv4Multicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv4Multicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv4Multicast.PrefixLimit}
    l3VpnIpv4Multicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv4Multicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast
// Multicast IPv6 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv6Multicast.EntityData.YFilter = l3VpnIpv6Multicast.YFilter
    l3VpnIpv6Multicast.EntityData.YangName = "l3vpn-ipv6-multicast"
    l3VpnIpv6Multicast.EntityData.BundleName = "openconfig"
    l3VpnIpv6Multicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv6Multicast.EntityData.SegmentPath = "l3vpn-ipv6-multicast"
    l3VpnIpv6Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv6Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv6Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv6Multicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv6Multicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv6Multicast.PrefixLimit}
    l3VpnIpv6Multicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv6Multicast.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls
// BGP-signalled VPLS configuration options
type Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
}

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetEntityData() *types.CommonEntityData {
    l2VpnVpls.EntityData.YFilter = l2VpnVpls.YFilter
    l2VpnVpls.EntityData.YangName = "l2vpn-vpls"
    l2VpnVpls.EntityData.BundleName = "openconfig"
    l2VpnVpls.EntityData.ParentYangName = "afi-safi"
    l2VpnVpls.EntityData.SegmentPath = "l2vpn-vpls"
    l2VpnVpls.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2VpnVpls.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2VpnVpls.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2VpnVpls.EntityData.Children = make(map[string]types.YChild)
    l2VpnVpls.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l2VpnVpls.PrefixLimit}
    l2VpnVpls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l2VpnVpls.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-vpls"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn
// BGP EVPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
}

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetEntityData() *types.CommonEntityData {
    l2VpnEvpn.EntityData.YFilter = l2VpnEvpn.YFilter
    l2VpnEvpn.EntityData.YangName = "l2vpn-evpn"
    l2VpnEvpn.EntityData.BundleName = "openconfig"
    l2VpnEvpn.EntityData.ParentYangName = "afi-safi"
    l2VpnEvpn.EntityData.SegmentPath = "l2vpn-evpn"
    l2VpnEvpn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2VpnEvpn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2VpnEvpn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2VpnEvpn.EntityData.Children = make(map[string]types.YChild)
    l2VpnEvpn.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l2VpnEvpn.PrefixLimit}
    l2VpnEvpn.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l2VpnEvpn.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-evpn"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Global_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_Global_ApplyPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_Global_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_Global_ApplyPolicy_State
}

func (applyPolicy *Bgp_Global_ApplyPolicy) GetEntityData() *types.CommonEntityData {
    applyPolicy.EntityData.YFilter = applyPolicy.YFilter
    applyPolicy.EntityData.YangName = "apply-policy"
    applyPolicy.EntityData.BundleName = "openconfig"
    applyPolicy.EntityData.ParentYangName = "global"
    applyPolicy.EntityData.SegmentPath = "apply-policy"
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = make(map[string]types.YChild)
    applyPolicy.EntityData.Children["config"] = types.YChild{"Config", &applyPolicy.Config}
    applyPolicy.EntityData.Children["state"] = types.YChild{"State", &applyPolicy.State}
    applyPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(applyPolicy.EntityData)
}

// Bgp_Global_ApplyPolicy_Config
// Policy configuration data.
type Bgp_Global_ApplyPolicy_Config struct {
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

func (config *Bgp_Global_ApplyPolicy_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "apply-policy"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", config.ImportPolicy}
    config.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy}
    config.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", config.ExportPolicy}
    config.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy}
    return &(config.EntityData)
}

// Bgp_Global_ApplyPolicy_State
// Operational state for routing policy
type Bgp_Global_ApplyPolicy_State struct {
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

func (state *Bgp_Global_ApplyPolicy_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "apply-policy"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", state.ImportPolicy}
    state.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy}
    state.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", state.ExportPolicy}
    state.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy}
    return &(state.EntityData)
}

// Bgp_Neighbors
// Configuration for BGP neighbors
type Bgp_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BGP neighbors configured on the local system, uniquely identified
    // by peer IPv[46] address. The type is slice of Bgp_Neighbors_Neighbor.
    Neighbor []Bgp_Neighbors_Neighbor
}

func (neighbors *Bgp_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "bgp"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbors.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children[types.GetSegmentPath(&neighbors.Neighbor[i])] = types.YChild{"Neighbor", &neighbors.Neighbor[i]}
    }
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// Bgp_Neighbors_Neighbor
// List of BGP neighbors configured on the local system,
// uniquely identified by peer IPv[46] address
type Bgp_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the address of the BGP neighbor used
    // as a key in the neighbor list. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    neighbor.EntityData.SegmentPath = "neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", neighbor.NeighborAddress) + "']"
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Children["config"] = types.YChild{"Config", &neighbor.Config}
    neighbor.EntityData.Children["state"] = types.YChild{"State", &neighbor.State}
    neighbor.EntityData.Children["timers"] = types.YChild{"Timers", &neighbor.Timers}
    neighbor.EntityData.Children["transport"] = types.YChild{"Transport", &neighbor.Transport}
    neighbor.EntityData.Children["error-handling"] = types.YChild{"ErrorHandling", &neighbor.ErrorHandling}
    neighbor.EntityData.Children["graceful-restart"] = types.YChild{"GracefulRestart", &neighbor.GracefulRestart}
    neighbor.EntityData.Children["logging-options"] = types.YChild{"LoggingOptions", &neighbor.LoggingOptions}
    neighbor.EntityData.Children["ebgp-multihop"] = types.YChild{"EbgpMultihop", &neighbor.EbgpMultihop}
    neighbor.EntityData.Children["route-reflector"] = types.YChild{"RouteReflector", &neighbor.RouteReflector}
    neighbor.EntityData.Children["as-path-options"] = types.YChild{"AsPathOptions", &neighbor.AsPathOptions}
    neighbor.EntityData.Children["add-paths"] = types.YChild{"AddPaths", &neighbor.AddPaths}
    neighbor.EntityData.Children["use-multiple-paths"] = types.YChild{"UseMultiplePaths", &neighbor.UseMultiplePaths}
    neighbor.EntityData.Children["apply-policy"] = types.YChild{"ApplyPolicy", &neighbor.ApplyPolicy}
    neighbor.EntityData.Children["afi-safis"] = types.YChild{"AfiSafis", &neighbor.AfiSafis}
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", neighbor.NeighborAddress}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["peer-group"] = types.YLeaf{"PeerGroup", config.PeerGroup}
    config.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", config.NeighborAddress}
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["peer-as"] = types.YLeaf{"PeerAs", config.PeerAs}
    config.EntityData.Leafs["local-as"] = types.YLeaf{"LocalAs", config.LocalAs}
    config.EntityData.Leafs["peer-type"] = types.YLeaf{"PeerType", config.PeerType}
    config.EntityData.Leafs["auth-password"] = types.YLeaf{"AuthPassword", config.AuthPassword}
    config.EntityData.Leafs["remove-private-as"] = types.YLeaf{"RemovePrivateAs", config.RemovePrivateAs}
    config.EntityData.Leafs["route-flap-damping"] = types.YLeaf{"RouteFlapDamping", config.RouteFlapDamping}
    config.EntityData.Leafs["send-community"] = types.YLeaf{"SendCommunity", config.SendCommunity}
    config.EntityData.Leafs["description"] = types.YLeaf{"Description", config.Description}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // relative to the Unix Epoch (Jan 1, 1970 00:00:00 UTC). The BGP session
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["messages"] = types.YChild{"Messages", &state.Messages}
    state.EntityData.Children["queues"] = types.YChild{"Queues", &state.Queues}
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["peer-group"] = types.YLeaf{"PeerGroup", state.PeerGroup}
    state.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", state.NeighborAddress}
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["peer-as"] = types.YLeaf{"PeerAs", state.PeerAs}
    state.EntityData.Leafs["local-as"] = types.YLeaf{"LocalAs", state.LocalAs}
    state.EntityData.Leafs["peer-type"] = types.YLeaf{"PeerType", state.PeerType}
    state.EntityData.Leafs["auth-password"] = types.YLeaf{"AuthPassword", state.AuthPassword}
    state.EntityData.Leafs["remove-private-as"] = types.YLeaf{"RemovePrivateAs", state.RemovePrivateAs}
    state.EntityData.Leafs["route-flap-damping"] = types.YLeaf{"RouteFlapDamping", state.RouteFlapDamping}
    state.EntityData.Leafs["send-community"] = types.YLeaf{"SendCommunity", state.SendCommunity}
    state.EntityData.Leafs["description"] = types.YLeaf{"Description", state.Description}
    state.EntityData.Leafs["session-state"] = types.YLeaf{"SessionState", state.SessionState}
    state.EntityData.Leafs["last-established"] = types.YLeaf{"LastEstablished", state.LastEstablished}
    state.EntityData.Leafs["established-transitions"] = types.YLeaf{"EstablishedTransitions", state.EstablishedTransitions}
    state.EntityData.Leafs["supported-capabilities"] = types.YLeaf{"SupportedCapabilities", state.SupportedCapabilities}
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
    messages.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    messages.EntityData.NamespaceTable = openconfig.GetNamespaces()
    messages.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    messages.EntityData.Children = make(map[string]types.YChild)
    messages.EntityData.Children["sent"] = types.YChild{"Sent", &messages.Sent}
    messages.EntityData.Children["received"] = types.YChild{"Received", &messages.Received}
    messages.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(messages.EntityData)
}

// Bgp_Neighbors_Neighbor_State_Messages_Sent
// Counters relating to BGP messages sent to the neighbor
type Bgp_Neighbors_Neighbor_State_Messages_Sent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of BGP UPDATE messages announcing, withdrawing or modifying paths
    // exchanged. The type is interface{} with range: 0..18446744073709551615.
    Update interface{}

    // Number of BGP NOTIFICATION messages indicating an error condition has
    // occurred exchanged. The type is interface{} with range:
    // 0..18446744073709551615.
    Notification interface{}
}

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetEntityData() *types.CommonEntityData {
    sent.EntityData.YFilter = sent.YFilter
    sent.EntityData.YangName = "sent"
    sent.EntityData.BundleName = "openconfig"
    sent.EntityData.ParentYangName = "messages"
    sent.EntityData.SegmentPath = "sent"
    sent.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sent.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sent.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sent.EntityData.Children = make(map[string]types.YChild)
    sent.EntityData.Leafs = make(map[string]types.YLeaf)
    sent.EntityData.Leafs["UPDATE"] = types.YLeaf{"Update", sent.Update}
    sent.EntityData.Leafs["NOTIFICATION"] = types.YLeaf{"Notification", sent.Notification}
    return &(sent.EntityData)
}

// Bgp_Neighbors_Neighbor_State_Messages_Received
// Counters for BGP messages received from the neighbor
type Bgp_Neighbors_Neighbor_State_Messages_Received struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of BGP UPDATE messages announcing, withdrawing or modifying paths
    // exchanged. The type is interface{} with range: 0..18446744073709551615.
    Update interface{}

    // Number of BGP NOTIFICATION messages indicating an error condition has
    // occurred exchanged. The type is interface{} with range:
    // 0..18446744073709551615.
    Notification interface{}
}

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetEntityData() *types.CommonEntityData {
    received.EntityData.YFilter = received.YFilter
    received.EntityData.YangName = "received"
    received.EntityData.BundleName = "openconfig"
    received.EntityData.ParentYangName = "messages"
    received.EntityData.SegmentPath = "received"
    received.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    received.EntityData.NamespaceTable = openconfig.GetNamespaces()
    received.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    received.EntityData.Children = make(map[string]types.YChild)
    received.EntityData.Leafs = make(map[string]types.YLeaf)
    received.EntityData.Leafs["UPDATE"] = types.YLeaf{"Update", received.Update}
    received.EntityData.Leafs["NOTIFICATION"] = types.YLeaf{"Notification", received.Notification}
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
    queues.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    queues.EntityData.NamespaceTable = openconfig.GetNamespaces()
    queues.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    queues.EntityData.Children = make(map[string]types.YChild)
    queues.EntityData.Leafs = make(map[string]types.YLeaf)
    queues.EntityData.Leafs["input"] = types.YLeaf{"Input", queues.Input}
    queues.EntityData.Leafs["output"] = types.YLeaf{"Output", queues.Output}
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
    timers.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    timers.EntityData.NamespaceTable = openconfig.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Children["config"] = types.YChild{"Config", &timers.Config}
    timers.EntityData.Children["state"] = types.YChild{"State", &timers.State}
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["connect-retry"] = types.YLeaf{"ConnectRetry", config.ConnectRetry}
    config.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", config.HoldTime}
    config.EntityData.Leafs["keepalive-interval"] = types.YLeaf{"KeepaliveInterval", config.KeepaliveInterval}
    config.EntityData.Leafs["minimum-advertisement-interval"] = types.YLeaf{"MinimumAdvertisementInterval", config.MinimumAdvertisementInterval}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["connect-retry"] = types.YLeaf{"ConnectRetry", state.ConnectRetry}
    state.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", state.HoldTime}
    state.EntityData.Leafs["keepalive-interval"] = types.YLeaf{"KeepaliveInterval", state.KeepaliveInterval}
    state.EntityData.Leafs["minimum-advertisement-interval"] = types.YLeaf{"MinimumAdvertisementInterval", state.MinimumAdvertisementInterval}
    state.EntityData.Leafs["negotiated-hold-time"] = types.YLeaf{"NegotiatedHoldTime", state.NegotiatedHoldTime}
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
    transport.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    transport.EntityData.NamespaceTable = openconfig.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    transport.EntityData.Children = make(map[string]types.YChild)
    transport.EntityData.Children["config"] = types.YChild{"Config", &transport.Config}
    transport.EntityData.Children["state"] = types.YChild{"State", &transport.State}
    transport.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.,
    // or string.
    LocalAddress interface{}
}

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "transport"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["tcp-mss"] = types.YLeaf{"TcpMss", config.TcpMss}
    config.EntityData.Leafs["mtu-discovery"] = types.YLeaf{"MtuDiscovery", config.MtuDiscovery}
    config.EntityData.Leafs["passive-mode"] = types.YLeaf{"PassiveMode", config.PassiveMode}
    config.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", config.LocalAddress}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.,
    // or string.
    LocalAddress interface{}

    // Local TCP port being used for the TCP session supporting the BGP session.
    // The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Remote address to which the BGP session has been established. The type is
    // one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["tcp-mss"] = types.YLeaf{"TcpMss", state.TcpMss}
    state.EntityData.Leafs["mtu-discovery"] = types.YLeaf{"MtuDiscovery", state.MtuDiscovery}
    state.EntityData.Leafs["passive-mode"] = types.YLeaf{"PassiveMode", state.PassiveMode}
    state.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", state.LocalAddress}
    state.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", state.LocalPort}
    state.EntityData.Leafs["remote-address"] = types.YLeaf{"RemoteAddress", state.RemoteAddress}
    state.EntityData.Leafs["remote-port"] = types.YLeaf{"RemotePort", state.RemotePort}
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
    errorHandling.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    errorHandling.EntityData.NamespaceTable = openconfig.GetNamespaces()
    errorHandling.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    errorHandling.EntityData.Children = make(map[string]types.YChild)
    errorHandling.EntityData.Children["config"] = types.YChild{"Config", &errorHandling.Config}
    errorHandling.EntityData.Children["state"] = types.YChild{"State", &errorHandling.State}
    errorHandling.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["treat-as-withdraw"] = types.YLeaf{"TreatAsWithdraw", config.TreatAsWithdraw}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["treat-as-withdraw"] = types.YLeaf{"TreatAsWithdraw", state.TreatAsWithdraw}
    state.EntityData.Leafs["erroneous-update-messages"] = types.YLeaf{"ErroneousUpdateMessages", state.ErroneousUpdateMessages}
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
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = make(map[string]types.YChild)
    gracefulRestart.EntityData.Children["config"] = types.YChild{"Config", &gracefulRestart.Config}
    gracefulRestart.EntityData.Children["state"] = types.YChild{"State", &gracefulRestart.State}
    gracefulRestart.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["restart-time"] = types.YLeaf{"RestartTime", config.RestartTime}
    config.EntityData.Leafs["stale-routes-time"] = types.YLeaf{"StaleRoutesTime", config.StaleRoutesTime}
    config.EntityData.Leafs["helper-only"] = types.YLeaf{"HelperOnly", config.HelperOnly}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["restart-time"] = types.YLeaf{"RestartTime", state.RestartTime}
    state.EntityData.Leafs["stale-routes-time"] = types.YLeaf{"StaleRoutesTime", state.StaleRoutesTime}
    state.EntityData.Leafs["helper-only"] = types.YLeaf{"HelperOnly", state.HelperOnly}
    state.EntityData.Leafs["peer-restart-time"] = types.YLeaf{"PeerRestartTime", state.PeerRestartTime}
    state.EntityData.Leafs["peer-restarting"] = types.YLeaf{"PeerRestarting", state.PeerRestarting}
    state.EntityData.Leafs["local-restarting"] = types.YLeaf{"LocalRestarting", state.LocalRestarting}
    state.EntityData.Leafs["mode"] = types.YLeaf{"Mode", state.Mode}
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
    loggingOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    loggingOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    loggingOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    loggingOptions.EntityData.Children = make(map[string]types.YChild)
    loggingOptions.EntityData.Children["config"] = types.YChild{"Config", &loggingOptions.Config}
    loggingOptions.EntityData.Children["state"] = types.YChild{"State", &loggingOptions.State}
    loggingOptions.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["log-neighbor-state-changes"] = types.YLeaf{"LogNeighborStateChanges", config.LogNeighborStateChanges}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["log-neighbor-state-changes"] = types.YLeaf{"LogNeighborStateChanges", state.LogNeighborStateChanges}
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
    ebgpMultihop.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgpMultihop.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgpMultihop.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgpMultihop.EntityData.Children = make(map[string]types.YChild)
    ebgpMultihop.EntityData.Children["config"] = types.YChild{"Config", &ebgpMultihop.Config}
    ebgpMultihop.EntityData.Children["state"] = types.YChild{"State", &ebgpMultihop.State}
    ebgpMultihop.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["multihop-ttl"] = types.YLeaf{"MultihopTtl", config.MultihopTtl}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["multihop-ttl"] = types.YLeaf{"MultihopTtl", state.MultihopTtl}
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
    routeReflector.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeReflector.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeReflector.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeReflector.EntityData.Children = make(map[string]types.YChild)
    routeReflector.EntityData.Children["config"] = types.YChild{"Config", &routeReflector.Config}
    routeReflector.EntityData.Children["state"] = types.YChild{"State", &routeReflector.State}
    routeReflector.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["route-reflector-cluster-id"] = types.YLeaf{"RouteReflectorClusterId", config.RouteReflectorClusterId}
    config.EntityData.Leafs["route-reflector-client"] = types.YLeaf{"RouteReflectorClient", config.RouteReflectorClient}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["route-reflector-cluster-id"] = types.YLeaf{"RouteReflectorClusterId", state.RouteReflectorClusterId}
    state.EntityData.Leafs["route-reflector-client"] = types.YLeaf{"RouteReflectorClient", state.RouteReflectorClient}
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
    asPathOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    asPathOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    asPathOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    asPathOptions.EntityData.Children = make(map[string]types.YChild)
    asPathOptions.EntityData.Children["config"] = types.YChild{"Config", &asPathOptions.Config}
    asPathOptions.EntityData.Children["state"] = types.YChild{"State", &asPathOptions.State}
    asPathOptions.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["allow-own-as"] = types.YLeaf{"AllowOwnAs", config.AllowOwnAs}
    config.EntityData.Leafs["replace-peer-as"] = types.YLeaf{"ReplacePeerAs", config.ReplacePeerAs}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["allow-own-as"] = types.YLeaf{"AllowOwnAs", state.AllowOwnAs}
    state.EntityData.Leafs["replace-peer-as"] = types.YLeaf{"ReplacePeerAs", state.ReplacePeerAs}
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
    addPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    addPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    addPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    addPaths.EntityData.Children = make(map[string]types.YChild)
    addPaths.EntityData.Children["config"] = types.YChild{"Config", &addPaths.Config}
    addPaths.EntityData.Children["state"] = types.YChild{"State", &addPaths.State}
    addPaths.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["receive"] = types.YLeaf{"Receive", config.Receive}
    config.EntityData.Leafs["send-max"] = types.YLeaf{"SendMax", config.SendMax}
    config.EntityData.Leafs["eligible-prefix-policy"] = types.YLeaf{"EligiblePrefixPolicy", config.EligiblePrefixPolicy}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["receive"] = types.YLeaf{"Receive", state.Receive}
    state.EntityData.Leafs["send-max"] = types.YLeaf{"SendMax", state.SendMax}
    state.EntityData.Leafs["eligible-prefix-policy"] = types.YLeaf{"EligiblePrefixPolicy", state.EligiblePrefixPolicy}
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
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = make(map[string]types.YChild)
    useMultiplePaths.EntityData.Children["config"] = types.YChild{"Config", &useMultiplePaths.Config}
    useMultiplePaths.EntityData.Children["state"] = types.YChild{"State", &useMultiplePaths.State}
    useMultiplePaths.EntityData.Children["ebgp"] = types.YChild{"Ebgp", &useMultiplePaths.Ebgp}
    useMultiplePaths.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
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
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = make(map[string]types.YChild)
    ebgp.EntityData.Children["config"] = types.YChild{"Config", &ebgp.Config}
    ebgp.EntityData.Children["state"] = types.YChild{"State", &ebgp.State}
    ebgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs}
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
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = make(map[string]types.YChild)
    applyPolicy.EntityData.Children["config"] = types.YChild{"Config", &applyPolicy.Config}
    applyPolicy.EntityData.Children["state"] = types.YChild{"State", &applyPolicy.State}
    applyPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", config.ImportPolicy}
    config.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy}
    config.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", config.ExportPolicy}
    config.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", state.ImportPolicy}
    state.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy}
    state.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", state.ExportPolicy}
    state.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy}
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
    AfiSafi []Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi
}

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetEntityData() *types.CommonEntityData {
    afiSafis.EntityData.YFilter = afiSafis.YFilter
    afiSafis.EntityData.YangName = "afi-safis"
    afiSafis.EntityData.BundleName = "openconfig"
    afiSafis.EntityData.ParentYangName = "neighbor"
    afiSafis.EntityData.SegmentPath = "afi-safis"
    afiSafis.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafis.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafis.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafis.EntityData.Children = make(map[string]types.YChild)
    afiSafis.EntityData.Children["afi-safi"] = types.YChild{"AfiSafi", nil}
    for i := range afiSafis.AfiSafi {
        afiSafis.EntityData.Children[types.GetSegmentPath(&afiSafis.AfiSafi[i])] = types.YChild{"AfiSafi", &afiSafis.AfiSafi[i]}
    }
    afiSafis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(afiSafis.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi
// AFI,SAFI configuration available for the
// neighbour or group
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    L3VpnIpv4Unicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast

    // Unicast IPv6 L3VPN configuration options.
    L3VpnIpv6Unicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast

    // Multicast IPv4 L3VPN configuration options.
    L3VpnIpv4Multicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast

    // Multicast IPv6 L3VPN configuration options.
    L3VpnIpv6Multicast Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast

    // BGP-signalled VPLS configuration options.
    L2VpnVpls Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls

    // BGP EVPN configuration options.
    L2VpnEvpn Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn

    // Parameters related to the use of multiple-paths for the same NLRI when they
    // are received only from this neighbor.
    UseMultiplePaths Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths
}

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetEntityData() *types.CommonEntityData {
    afiSafi.EntityData.YFilter = afiSafi.YFilter
    afiSafi.EntityData.YangName = "afi-safi"
    afiSafi.EntityData.BundleName = "openconfig"
    afiSafi.EntityData.ParentYangName = "afi-safis"
    afiSafi.EntityData.SegmentPath = "afi-safi" + "[afi-safi-name='" + fmt.Sprintf("%v", afiSafi.AfiSafiName) + "']"
    afiSafi.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafi.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafi.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafi.EntityData.Children = make(map[string]types.YChild)
    afiSafi.EntityData.Children["config"] = types.YChild{"Config", &afiSafi.Config}
    afiSafi.EntityData.Children["state"] = types.YChild{"State", &afiSafi.State}
    afiSafi.EntityData.Children["graceful-restart"] = types.YChild{"GracefulRestart", &afiSafi.GracefulRestart}
    afiSafi.EntityData.Children["apply-policy"] = types.YChild{"ApplyPolicy", &afiSafi.ApplyPolicy}
    afiSafi.EntityData.Children["ipv4-unicast"] = types.YChild{"Ipv4Unicast", &afiSafi.Ipv4Unicast}
    afiSafi.EntityData.Children["ipv6-unicast"] = types.YChild{"Ipv6Unicast", &afiSafi.Ipv6Unicast}
    afiSafi.EntityData.Children["ipv4-labeled-unicast"] = types.YChild{"Ipv4LabeledUnicast", &afiSafi.Ipv4LabeledUnicast}
    afiSafi.EntityData.Children["ipv6-labeled-unicast"] = types.YChild{"Ipv6LabeledUnicast", &afiSafi.Ipv6LabeledUnicast}
    afiSafi.EntityData.Children["l3vpn-ipv4-unicast"] = types.YChild{"L3VpnIpv4Unicast", &afiSafi.L3VpnIpv4Unicast}
    afiSafi.EntityData.Children["l3vpn-ipv6-unicast"] = types.YChild{"L3VpnIpv6Unicast", &afiSafi.L3VpnIpv6Unicast}
    afiSafi.EntityData.Children["l3vpn-ipv4-multicast"] = types.YChild{"L3VpnIpv4Multicast", &afiSafi.L3VpnIpv4Multicast}
    afiSafi.EntityData.Children["l3vpn-ipv6-multicast"] = types.YChild{"L3VpnIpv6Multicast", &afiSafi.L3VpnIpv6Multicast}
    afiSafi.EntityData.Children["l2vpn-vpls"] = types.YChild{"L2VpnVpls", &afiSafi.L2VpnVpls}
    afiSafi.EntityData.Children["l2vpn-evpn"] = types.YChild{"L2VpnEvpn", &afiSafi.L2VpnEvpn}
    afiSafi.EntityData.Children["use-multiple-paths"] = types.YChild{"UseMultiplePaths", &afiSafi.UseMultiplePaths}
    afiSafi.EntityData.Leafs = make(map[string]types.YLeaf)
    afiSafi.EntityData.Leafs["afi-safi-name"] = types.YLeaf{"AfiSafiName", afiSafi.AfiSafiName}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["afi-safi-name"] = types.YLeaf{"AfiSafiName", config.AfiSafiName}
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["prefixes"] = types.YChild{"Prefixes", &state.Prefixes}
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["afi-safi-name"] = types.YLeaf{"AfiSafiName", state.AfiSafiName}
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["active"] = types.YLeaf{"Active", state.Active}
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
    prefixes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixes.EntityData.Children = make(map[string]types.YChild)
    prefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixes.EntityData.Leafs["received"] = types.YLeaf{"Received", prefixes.Received}
    prefixes.EntityData.Leafs["sent"] = types.YLeaf{"Sent", prefixes.Sent}
    prefixes.EntityData.Leafs["installed"] = types.YLeaf{"Installed", prefixes.Installed}
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
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = make(map[string]types.YChild)
    gracefulRestart.EntityData.Children["config"] = types.YChild{"Config", &gracefulRestart.Config}
    gracefulRestart.EntityData.Children["state"] = types.YChild{"State", &gracefulRestart.State}
    gracefulRestart.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["received"] = types.YLeaf{"Received", state.Received}
    state.EntityData.Leafs["advertised"] = types.YLeaf{"Advertised", state.Advertised}
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
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = make(map[string]types.YChild)
    applyPolicy.EntityData.Children["config"] = types.YChild{"Config", &applyPolicy.Config}
    applyPolicy.EntityData.Children["state"] = types.YChild{"State", &applyPolicy.State}
    applyPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", config.ImportPolicy}
    config.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy}
    config.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", config.ExportPolicy}
    config.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", state.ImportPolicy}
    state.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy}
    state.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", state.ExportPolicy}
    state.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy}
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
    ipv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4Unicast.EntityData.Children = make(map[string]types.YChild)
    ipv4Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv4Unicast.PrefixLimit}
    ipv4Unicast.EntityData.Children["config"] = types.YChild{"Config", &ipv4Unicast.Config}
    ipv4Unicast.EntityData.Children["state"] = types.YChild{"State", &ipv4Unicast.State}
    ipv4Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute}
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
    ipv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6Unicast.EntityData.Children = make(map[string]types.YChild)
    ipv6Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv6Unicast.PrefixLimit}
    ipv6Unicast.EntityData.Children["config"] = types.YChild{"Config", &ipv6Unicast.Config}
    ipv6Unicast.EntityData.Children["state"] = types.YChild{"State", &ipv6Unicast.State}
    ipv6Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute}
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
    ipv4LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4LabeledUnicast.EntityData.Children = make(map[string]types.YChild)
    ipv4LabeledUnicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv4LabeledUnicast.PrefixLimit}
    ipv4LabeledUnicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    ipv6LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6LabeledUnicast.EntityData.Children = make(map[string]types.YChild)
    ipv6LabeledUnicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv6LabeledUnicast.PrefixLimit}
    ipv6LabeledUnicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast
// Unicast IPv4 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv4Unicast.EntityData.YFilter = l3VpnIpv4Unicast.YFilter
    l3VpnIpv4Unicast.EntityData.YangName = "l3vpn-ipv4-unicast"
    l3VpnIpv4Unicast.EntityData.BundleName = "openconfig"
    l3VpnIpv4Unicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv4Unicast.EntityData.SegmentPath = "l3vpn-ipv4-unicast"
    l3VpnIpv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv4Unicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv4Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv4Unicast.PrefixLimit}
    l3VpnIpv4Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv4Unicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast
// Unicast IPv6 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv6Unicast.EntityData.YFilter = l3VpnIpv6Unicast.YFilter
    l3VpnIpv6Unicast.EntityData.YangName = "l3vpn-ipv6-unicast"
    l3VpnIpv6Unicast.EntityData.BundleName = "openconfig"
    l3VpnIpv6Unicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv6Unicast.EntityData.SegmentPath = "l3vpn-ipv6-unicast"
    l3VpnIpv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv6Unicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv6Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv6Unicast.PrefixLimit}
    l3VpnIpv6Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv6Unicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast
// Multicast IPv4 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv4Multicast.EntityData.YFilter = l3VpnIpv4Multicast.YFilter
    l3VpnIpv4Multicast.EntityData.YangName = "l3vpn-ipv4-multicast"
    l3VpnIpv4Multicast.EntityData.BundleName = "openconfig"
    l3VpnIpv4Multicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv4Multicast.EntityData.SegmentPath = "l3vpn-ipv4-multicast"
    l3VpnIpv4Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv4Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv4Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv4Multicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv4Multicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv4Multicast.PrefixLimit}
    l3VpnIpv4Multicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv4Multicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast
// Multicast IPv6 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv6Multicast.EntityData.YFilter = l3VpnIpv6Multicast.YFilter
    l3VpnIpv6Multicast.EntityData.YangName = "l3vpn-ipv6-multicast"
    l3VpnIpv6Multicast.EntityData.BundleName = "openconfig"
    l3VpnIpv6Multicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv6Multicast.EntityData.SegmentPath = "l3vpn-ipv6-multicast"
    l3VpnIpv6Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv6Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv6Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv6Multicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv6Multicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv6Multicast.PrefixLimit}
    l3VpnIpv6Multicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv6Multicast.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls
// BGP-signalled VPLS configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
}

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetEntityData() *types.CommonEntityData {
    l2VpnVpls.EntityData.YFilter = l2VpnVpls.YFilter
    l2VpnVpls.EntityData.YangName = "l2vpn-vpls"
    l2VpnVpls.EntityData.BundleName = "openconfig"
    l2VpnVpls.EntityData.ParentYangName = "afi-safi"
    l2VpnVpls.EntityData.SegmentPath = "l2vpn-vpls"
    l2VpnVpls.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2VpnVpls.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2VpnVpls.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2VpnVpls.EntityData.Children = make(map[string]types.YChild)
    l2VpnVpls.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l2VpnVpls.PrefixLimit}
    l2VpnVpls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l2VpnVpls.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-vpls"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn
// BGP EVPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
}

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetEntityData() *types.CommonEntityData {
    l2VpnEvpn.EntityData.YFilter = l2VpnEvpn.YFilter
    l2VpnEvpn.EntityData.YangName = "l2vpn-evpn"
    l2VpnEvpn.EntityData.BundleName = "openconfig"
    l2VpnEvpn.EntityData.ParentYangName = "afi-safi"
    l2VpnEvpn.EntityData.SegmentPath = "l2vpn-evpn"
    l2VpnEvpn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2VpnEvpn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2VpnEvpn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2VpnEvpn.EntityData.Children = make(map[string]types.YChild)
    l2VpnEvpn.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l2VpnEvpn.PrefixLimit}
    l2VpnEvpn.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l2VpnEvpn.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-evpn"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = make(map[string]types.YChild)
    useMultiplePaths.EntityData.Children["config"] = types.YChild{"Config", &useMultiplePaths.Config}
    useMultiplePaths.EntityData.Children["state"] = types.YChild{"State", &useMultiplePaths.State}
    useMultiplePaths.EntityData.Children["ebgp"] = types.YChild{"Ebgp", &useMultiplePaths.Ebgp}
    useMultiplePaths.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
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
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = make(map[string]types.YChild)
    ebgp.EntityData.Children["config"] = types.YChild{"Config", &ebgp.Config}
    ebgp.EntityData.Children["state"] = types.YChild{"State", &ebgp.State}
    ebgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs}
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
    PeerGroup []Bgp_PeerGroups_PeerGroup
}

func (peerGroups *Bgp_PeerGroups) GetEntityData() *types.CommonEntityData {
    peerGroups.EntityData.YFilter = peerGroups.YFilter
    peerGroups.EntityData.YangName = "peer-groups"
    peerGroups.EntityData.BundleName = "openconfig"
    peerGroups.EntityData.ParentYangName = "bgp"
    peerGroups.EntityData.SegmentPath = "peer-groups"
    peerGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    peerGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    peerGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    peerGroups.EntityData.Children = make(map[string]types.YChild)
    peerGroups.EntityData.Children["peer-group"] = types.YChild{"PeerGroup", nil}
    for i := range peerGroups.PeerGroup {
        peerGroups.EntityData.Children[types.GetSegmentPath(&peerGroups.PeerGroup[i])] = types.YChild{"PeerGroup", &peerGroups.PeerGroup[i]}
    }
    peerGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerGroups.EntityData)
}

// Bgp_PeerGroups_PeerGroup
// List of BGP peer-groups configured on the local system -
// uniquely identified by peer-group name
type Bgp_PeerGroups_PeerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    peerGroup.EntityData.SegmentPath = "peer-group" + "[peer-group-name='" + fmt.Sprintf("%v", peerGroup.PeerGroupName) + "']"
    peerGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    peerGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    peerGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    peerGroup.EntityData.Children = make(map[string]types.YChild)
    peerGroup.EntityData.Children["config"] = types.YChild{"Config", &peerGroup.Config}
    peerGroup.EntityData.Children["state"] = types.YChild{"State", &peerGroup.State}
    peerGroup.EntityData.Children["timers"] = types.YChild{"Timers", &peerGroup.Timers}
    peerGroup.EntityData.Children["transport"] = types.YChild{"Transport", &peerGroup.Transport}
    peerGroup.EntityData.Children["error-handling"] = types.YChild{"ErrorHandling", &peerGroup.ErrorHandling}
    peerGroup.EntityData.Children["graceful-restart"] = types.YChild{"GracefulRestart", &peerGroup.GracefulRestart}
    peerGroup.EntityData.Children["logging-options"] = types.YChild{"LoggingOptions", &peerGroup.LoggingOptions}
    peerGroup.EntityData.Children["ebgp-multihop"] = types.YChild{"EbgpMultihop", &peerGroup.EbgpMultihop}
    peerGroup.EntityData.Children["route-reflector"] = types.YChild{"RouteReflector", &peerGroup.RouteReflector}
    peerGroup.EntityData.Children["as-path-options"] = types.YChild{"AsPathOptions", &peerGroup.AsPathOptions}
    peerGroup.EntityData.Children["add-paths"] = types.YChild{"AddPaths", &peerGroup.AddPaths}
    peerGroup.EntityData.Children["use-multiple-paths"] = types.YChild{"UseMultiplePaths", &peerGroup.UseMultiplePaths}
    peerGroup.EntityData.Children["apply-policy"] = types.YChild{"ApplyPolicy", &peerGroup.ApplyPolicy}
    peerGroup.EntityData.Children["afi-safis"] = types.YChild{"AfiSafis", &peerGroup.AfiSafis}
    peerGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    peerGroup.EntityData.Leafs["peer-group-name"] = types.YLeaf{"PeerGroupName", peerGroup.PeerGroupName}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["peer-group-name"] = types.YLeaf{"PeerGroupName", config.PeerGroupName}
    config.EntityData.Leafs["peer-as"] = types.YLeaf{"PeerAs", config.PeerAs}
    config.EntityData.Leafs["local-as"] = types.YLeaf{"LocalAs", config.LocalAs}
    config.EntityData.Leafs["peer-type"] = types.YLeaf{"PeerType", config.PeerType}
    config.EntityData.Leafs["auth-password"] = types.YLeaf{"AuthPassword", config.AuthPassword}
    config.EntityData.Leafs["remove-private-as"] = types.YLeaf{"RemovePrivateAs", config.RemovePrivateAs}
    config.EntityData.Leafs["route-flap-damping"] = types.YLeaf{"RouteFlapDamping", config.RouteFlapDamping}
    config.EntityData.Leafs["send-community"] = types.YLeaf{"SendCommunity", config.SendCommunity}
    config.EntityData.Leafs["description"] = types.YLeaf{"Description", config.Description}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["peer-group-name"] = types.YLeaf{"PeerGroupName", state.PeerGroupName}
    state.EntityData.Leafs["peer-as"] = types.YLeaf{"PeerAs", state.PeerAs}
    state.EntityData.Leafs["local-as"] = types.YLeaf{"LocalAs", state.LocalAs}
    state.EntityData.Leafs["peer-type"] = types.YLeaf{"PeerType", state.PeerType}
    state.EntityData.Leafs["auth-password"] = types.YLeaf{"AuthPassword", state.AuthPassword}
    state.EntityData.Leafs["remove-private-as"] = types.YLeaf{"RemovePrivateAs", state.RemovePrivateAs}
    state.EntityData.Leafs["route-flap-damping"] = types.YLeaf{"RouteFlapDamping", state.RouteFlapDamping}
    state.EntityData.Leafs["send-community"] = types.YLeaf{"SendCommunity", state.SendCommunity}
    state.EntityData.Leafs["description"] = types.YLeaf{"Description", state.Description}
    state.EntityData.Leafs["total-paths"] = types.YLeaf{"TotalPaths", state.TotalPaths}
    state.EntityData.Leafs["total-prefixes"] = types.YLeaf{"TotalPrefixes", state.TotalPrefixes}
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
    timers.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    timers.EntityData.NamespaceTable = openconfig.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Children["config"] = types.YChild{"Config", &timers.Config}
    timers.EntityData.Children["state"] = types.YChild{"State", &timers.State}
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["connect-retry"] = types.YLeaf{"ConnectRetry", config.ConnectRetry}
    config.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", config.HoldTime}
    config.EntityData.Leafs["keepalive-interval"] = types.YLeaf{"KeepaliveInterval", config.KeepaliveInterval}
    config.EntityData.Leafs["minimum-advertisement-interval"] = types.YLeaf{"MinimumAdvertisementInterval", config.MinimumAdvertisementInterval}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["connect-retry"] = types.YLeaf{"ConnectRetry", state.ConnectRetry}
    state.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", state.HoldTime}
    state.EntityData.Leafs["keepalive-interval"] = types.YLeaf{"KeepaliveInterval", state.KeepaliveInterval}
    state.EntityData.Leafs["minimum-advertisement-interval"] = types.YLeaf{"MinimumAdvertisementInterval", state.MinimumAdvertisementInterval}
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
    transport.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    transport.EntityData.NamespaceTable = openconfig.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    transport.EntityData.Children = make(map[string]types.YChild)
    transport.EntityData.Children["config"] = types.YChild{"Config", &transport.Config}
    transport.EntityData.Children["state"] = types.YChild{"State", &transport.State}
    transport.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.,
    // or string.
    LocalAddress interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "transport"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["tcp-mss"] = types.YLeaf{"TcpMss", config.TcpMss}
    config.EntityData.Leafs["mtu-discovery"] = types.YLeaf{"MtuDiscovery", config.MtuDiscovery}
    config.EntityData.Leafs["passive-mode"] = types.YLeaf{"PassiveMode", config.PassiveMode}
    config.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", config.LocalAddress}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.,
    // or string.
    LocalAddress interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "transport"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["tcp-mss"] = types.YLeaf{"TcpMss", state.TcpMss}
    state.EntityData.Leafs["mtu-discovery"] = types.YLeaf{"MtuDiscovery", state.MtuDiscovery}
    state.EntityData.Leafs["passive-mode"] = types.YLeaf{"PassiveMode", state.PassiveMode}
    state.EntityData.Leafs["local-address"] = types.YLeaf{"LocalAddress", state.LocalAddress}
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
    errorHandling.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    errorHandling.EntityData.NamespaceTable = openconfig.GetNamespaces()
    errorHandling.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    errorHandling.EntityData.Children = make(map[string]types.YChild)
    errorHandling.EntityData.Children["config"] = types.YChild{"Config", &errorHandling.Config}
    errorHandling.EntityData.Children["state"] = types.YChild{"State", &errorHandling.State}
    errorHandling.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["treat-as-withdraw"] = types.YLeaf{"TreatAsWithdraw", config.TreatAsWithdraw}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["treat-as-withdraw"] = types.YLeaf{"TreatAsWithdraw", state.TreatAsWithdraw}
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
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = make(map[string]types.YChild)
    gracefulRestart.EntityData.Children["config"] = types.YChild{"Config", &gracefulRestart.Config}
    gracefulRestart.EntityData.Children["state"] = types.YChild{"State", &gracefulRestart.State}
    gracefulRestart.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["restart-time"] = types.YLeaf{"RestartTime", config.RestartTime}
    config.EntityData.Leafs["stale-routes-time"] = types.YLeaf{"StaleRoutesTime", config.StaleRoutesTime}
    config.EntityData.Leafs["helper-only"] = types.YLeaf{"HelperOnly", config.HelperOnly}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["restart-time"] = types.YLeaf{"RestartTime", state.RestartTime}
    state.EntityData.Leafs["stale-routes-time"] = types.YLeaf{"StaleRoutesTime", state.StaleRoutesTime}
    state.EntityData.Leafs["helper-only"] = types.YLeaf{"HelperOnly", state.HelperOnly}
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
    loggingOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    loggingOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    loggingOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    loggingOptions.EntityData.Children = make(map[string]types.YChild)
    loggingOptions.EntityData.Children["config"] = types.YChild{"Config", &loggingOptions.Config}
    loggingOptions.EntityData.Children["state"] = types.YChild{"State", &loggingOptions.State}
    loggingOptions.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["log-neighbor-state-changes"] = types.YLeaf{"LogNeighborStateChanges", config.LogNeighborStateChanges}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["log-neighbor-state-changes"] = types.YLeaf{"LogNeighborStateChanges", state.LogNeighborStateChanges}
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
    ebgpMultihop.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgpMultihop.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgpMultihop.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgpMultihop.EntityData.Children = make(map[string]types.YChild)
    ebgpMultihop.EntityData.Children["config"] = types.YChild{"Config", &ebgpMultihop.Config}
    ebgpMultihop.EntityData.Children["state"] = types.YChild{"State", &ebgpMultihop.State}
    ebgpMultihop.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["multihop-ttl"] = types.YLeaf{"MultihopTtl", config.MultihopTtl}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["multihop-ttl"] = types.YLeaf{"MultihopTtl", state.MultihopTtl}
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
    routeReflector.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeReflector.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeReflector.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeReflector.EntityData.Children = make(map[string]types.YChild)
    routeReflector.EntityData.Children["config"] = types.YChild{"Config", &routeReflector.Config}
    routeReflector.EntityData.Children["state"] = types.YChild{"State", &routeReflector.State}
    routeReflector.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["route-reflector-cluster-id"] = types.YLeaf{"RouteReflectorClusterId", config.RouteReflectorClusterId}
    config.EntityData.Leafs["route-reflector-client"] = types.YLeaf{"RouteReflectorClient", config.RouteReflectorClient}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["route-reflector-cluster-id"] = types.YLeaf{"RouteReflectorClusterId", state.RouteReflectorClusterId}
    state.EntityData.Leafs["route-reflector-client"] = types.YLeaf{"RouteReflectorClient", state.RouteReflectorClient}
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
    asPathOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    asPathOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    asPathOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    asPathOptions.EntityData.Children = make(map[string]types.YChild)
    asPathOptions.EntityData.Children["config"] = types.YChild{"Config", &asPathOptions.Config}
    asPathOptions.EntityData.Children["state"] = types.YChild{"State", &asPathOptions.State}
    asPathOptions.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["allow-own-as"] = types.YLeaf{"AllowOwnAs", config.AllowOwnAs}
    config.EntityData.Leafs["replace-peer-as"] = types.YLeaf{"ReplacePeerAs", config.ReplacePeerAs}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["allow-own-as"] = types.YLeaf{"AllowOwnAs", state.AllowOwnAs}
    state.EntityData.Leafs["replace-peer-as"] = types.YLeaf{"ReplacePeerAs", state.ReplacePeerAs}
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
    addPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    addPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    addPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    addPaths.EntityData.Children = make(map[string]types.YChild)
    addPaths.EntityData.Children["config"] = types.YChild{"Config", &addPaths.Config}
    addPaths.EntityData.Children["state"] = types.YChild{"State", &addPaths.State}
    addPaths.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["receive"] = types.YLeaf{"Receive", config.Receive}
    config.EntityData.Leafs["send-max"] = types.YLeaf{"SendMax", config.SendMax}
    config.EntityData.Leafs["eligible-prefix-policy"] = types.YLeaf{"EligiblePrefixPolicy", config.EligiblePrefixPolicy}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["receive"] = types.YLeaf{"Receive", state.Receive}
    state.EntityData.Leafs["send-max"] = types.YLeaf{"SendMax", state.SendMax}
    state.EntityData.Leafs["eligible-prefix-policy"] = types.YLeaf{"EligiblePrefixPolicy", state.EligiblePrefixPolicy}
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
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = make(map[string]types.YChild)
    useMultiplePaths.EntityData.Children["config"] = types.YChild{"Config", &useMultiplePaths.Config}
    useMultiplePaths.EntityData.Children["state"] = types.YChild{"State", &useMultiplePaths.State}
    useMultiplePaths.EntityData.Children["ebgp"] = types.YChild{"Ebgp", &useMultiplePaths.Ebgp}
    useMultiplePaths.EntityData.Children["ibgp"] = types.YChild{"Ibgp", &useMultiplePaths.Ibgp}
    useMultiplePaths.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
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
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = make(map[string]types.YChild)
    ebgp.EntityData.Children["config"] = types.YChild{"Config", &ebgp.Config}
    ebgp.EntityData.Children["state"] = types.YChild{"State", &ebgp.State}
    ebgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs}
    config.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", config.MaximumPaths}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs}
    state.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", state.MaximumPaths}
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
    ibgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ibgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ibgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ibgp.EntityData.Children = make(map[string]types.YChild)
    ibgp.EntityData.Children["config"] = types.YChild{"Config", &ibgp.Config}
    ibgp.EntityData.Children["state"] = types.YChild{"State", &ibgp.State}
    ibgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", config.MaximumPaths}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", state.MaximumPaths}
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
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = make(map[string]types.YChild)
    applyPolicy.EntityData.Children["config"] = types.YChild{"Config", &applyPolicy.Config}
    applyPolicy.EntityData.Children["state"] = types.YChild{"State", &applyPolicy.State}
    applyPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", config.ImportPolicy}
    config.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy}
    config.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", config.ExportPolicy}
    config.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", state.ImportPolicy}
    state.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy}
    state.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", state.ExportPolicy}
    state.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy}
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
    AfiSafi []Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi
}

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetEntityData() *types.CommonEntityData {
    afiSafis.EntityData.YFilter = afiSafis.YFilter
    afiSafis.EntityData.YangName = "afi-safis"
    afiSafis.EntityData.BundleName = "openconfig"
    afiSafis.EntityData.ParentYangName = "peer-group"
    afiSafis.EntityData.SegmentPath = "afi-safis"
    afiSafis.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafis.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafis.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafis.EntityData.Children = make(map[string]types.YChild)
    afiSafis.EntityData.Children["afi-safi"] = types.YChild{"AfiSafi", nil}
    for i := range afiSafis.AfiSafi {
        afiSafis.EntityData.Children[types.GetSegmentPath(&afiSafis.AfiSafi[i])] = types.YChild{"AfiSafi", &afiSafis.AfiSafi[i]}
    }
    afiSafis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(afiSafis.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi
// AFI,SAFI configuration available for the
// neighbour or group
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    L3VpnIpv4Unicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast

    // Unicast IPv6 L3VPN configuration options.
    L3VpnIpv6Unicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast

    // Multicast IPv4 L3VPN configuration options.
    L3VpnIpv4Multicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast

    // Multicast IPv6 L3VPN configuration options.
    L3VpnIpv6Multicast Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast

    // BGP-signalled VPLS configuration options.
    L2VpnVpls Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls

    // BGP EVPN configuration options.
    L2VpnEvpn Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn
}

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetEntityData() *types.CommonEntityData {
    afiSafi.EntityData.YFilter = afiSafi.YFilter
    afiSafi.EntityData.YangName = "afi-safi"
    afiSafi.EntityData.BundleName = "openconfig"
    afiSafi.EntityData.ParentYangName = "afi-safis"
    afiSafi.EntityData.SegmentPath = "afi-safi" + "[afi-safi-name='" + fmt.Sprintf("%v", afiSafi.AfiSafiName) + "']"
    afiSafi.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    afiSafi.EntityData.NamespaceTable = openconfig.GetNamespaces()
    afiSafi.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    afiSafi.EntityData.Children = make(map[string]types.YChild)
    afiSafi.EntityData.Children["config"] = types.YChild{"Config", &afiSafi.Config}
    afiSafi.EntityData.Children["state"] = types.YChild{"State", &afiSafi.State}
    afiSafi.EntityData.Children["graceful-restart"] = types.YChild{"GracefulRestart", &afiSafi.GracefulRestart}
    afiSafi.EntityData.Children["route-selection-options"] = types.YChild{"RouteSelectionOptions", &afiSafi.RouteSelectionOptions}
    afiSafi.EntityData.Children["use-multiple-paths"] = types.YChild{"UseMultiplePaths", &afiSafi.UseMultiplePaths}
    afiSafi.EntityData.Children["apply-policy"] = types.YChild{"ApplyPolicy", &afiSafi.ApplyPolicy}
    afiSafi.EntityData.Children["ipv4-unicast"] = types.YChild{"Ipv4Unicast", &afiSafi.Ipv4Unicast}
    afiSafi.EntityData.Children["ipv6-unicast"] = types.YChild{"Ipv6Unicast", &afiSafi.Ipv6Unicast}
    afiSafi.EntityData.Children["ipv4-labeled-unicast"] = types.YChild{"Ipv4LabeledUnicast", &afiSafi.Ipv4LabeledUnicast}
    afiSafi.EntityData.Children["ipv6-labeled-unicast"] = types.YChild{"Ipv6LabeledUnicast", &afiSafi.Ipv6LabeledUnicast}
    afiSafi.EntityData.Children["l3vpn-ipv4-unicast"] = types.YChild{"L3VpnIpv4Unicast", &afiSafi.L3VpnIpv4Unicast}
    afiSafi.EntityData.Children["l3vpn-ipv6-unicast"] = types.YChild{"L3VpnIpv6Unicast", &afiSafi.L3VpnIpv6Unicast}
    afiSafi.EntityData.Children["l3vpn-ipv4-multicast"] = types.YChild{"L3VpnIpv4Multicast", &afiSafi.L3VpnIpv4Multicast}
    afiSafi.EntityData.Children["l3vpn-ipv6-multicast"] = types.YChild{"L3VpnIpv6Multicast", &afiSafi.L3VpnIpv6Multicast}
    afiSafi.EntityData.Children["l2vpn-vpls"] = types.YChild{"L2VpnVpls", &afiSafi.L2VpnVpls}
    afiSafi.EntityData.Children["l2vpn-evpn"] = types.YChild{"L2VpnEvpn", &afiSafi.L2VpnEvpn}
    afiSafi.EntityData.Leafs = make(map[string]types.YLeaf)
    afiSafi.EntityData.Leafs["afi-safi-name"] = types.YLeaf{"AfiSafiName", afiSafi.AfiSafiName}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["afi-safi-name"] = types.YLeaf{"AfiSafiName", config.AfiSafiName}
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["afi-safi-name"] = types.YLeaf{"AfiSafiName", state.AfiSafiName}
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
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
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = make(map[string]types.YChild)
    gracefulRestart.EntityData.Children["config"] = types.YChild{"Config", &gracefulRestart.Config}
    gracefulRestart.EntityData.Children["state"] = types.YChild{"State", &gracefulRestart.State}
    gracefulRestart.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
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
    routeSelectionOptions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routeSelectionOptions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routeSelectionOptions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routeSelectionOptions.EntityData.Children = make(map[string]types.YChild)
    routeSelectionOptions.EntityData.Children["config"] = types.YChild{"Config", &routeSelectionOptions.Config}
    routeSelectionOptions.EntityData.Children["state"] = types.YChild{"State", &routeSelectionOptions.State}
    routeSelectionOptions.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["always-compare-med"] = types.YLeaf{"AlwaysCompareMed", config.AlwaysCompareMed}
    config.EntityData.Leafs["ignore-as-path-length"] = types.YLeaf{"IgnoreAsPathLength", config.IgnoreAsPathLength}
    config.EntityData.Leafs["external-compare-router-id"] = types.YLeaf{"ExternalCompareRouterId", config.ExternalCompareRouterId}
    config.EntityData.Leafs["advertise-inactive-routes"] = types.YLeaf{"AdvertiseInactiveRoutes", config.AdvertiseInactiveRoutes}
    config.EntityData.Leafs["enable-aigp"] = types.YLeaf{"EnableAigp", config.EnableAigp}
    config.EntityData.Leafs["ignore-next-hop-igp-metric"] = types.YLeaf{"IgnoreNextHopIgpMetric", config.IgnoreNextHopIgpMetric}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["always-compare-med"] = types.YLeaf{"AlwaysCompareMed", state.AlwaysCompareMed}
    state.EntityData.Leafs["ignore-as-path-length"] = types.YLeaf{"IgnoreAsPathLength", state.IgnoreAsPathLength}
    state.EntityData.Leafs["external-compare-router-id"] = types.YLeaf{"ExternalCompareRouterId", state.ExternalCompareRouterId}
    state.EntityData.Leafs["advertise-inactive-routes"] = types.YLeaf{"AdvertiseInactiveRoutes", state.AdvertiseInactiveRoutes}
    state.EntityData.Leafs["enable-aigp"] = types.YLeaf{"EnableAigp", state.EnableAigp}
    state.EntityData.Leafs["ignore-next-hop-igp-metric"] = types.YLeaf{"IgnoreNextHopIgpMetric", state.IgnoreNextHopIgpMetric}
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
    useMultiplePaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    useMultiplePaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    useMultiplePaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    useMultiplePaths.EntityData.Children = make(map[string]types.YChild)
    useMultiplePaths.EntityData.Children["config"] = types.YChild{"Config", &useMultiplePaths.Config}
    useMultiplePaths.EntityData.Children["state"] = types.YChild{"State", &useMultiplePaths.State}
    useMultiplePaths.EntityData.Children["ebgp"] = types.YChild{"Ebgp", &useMultiplePaths.Ebgp}
    useMultiplePaths.EntityData.Children["ibgp"] = types.YChild{"Ibgp", &useMultiplePaths.Ibgp}
    useMultiplePaths.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
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
    ebgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ebgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ebgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ebgp.EntityData.Children = make(map[string]types.YChild)
    ebgp.EntityData.Children["config"] = types.YChild{"Config", &ebgp.Config}
    ebgp.EntityData.Children["state"] = types.YChild{"State", &ebgp.State}
    ebgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", config.AllowMultipleAs}
    config.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", config.MaximumPaths}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["allow-multiple-as"] = types.YLeaf{"AllowMultipleAs", state.AllowMultipleAs}
    state.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", state.MaximumPaths}
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
    ibgp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ibgp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ibgp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ibgp.EntityData.Children = make(map[string]types.YChild)
    ibgp.EntityData.Children["config"] = types.YChild{"Config", &ibgp.Config}
    ibgp.EntityData.Children["state"] = types.YChild{"State", &ibgp.State}
    ibgp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", config.MaximumPaths}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", state.MaximumPaths}
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
    applyPolicy.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    applyPolicy.EntityData.NamespaceTable = openconfig.GetNamespaces()
    applyPolicy.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    applyPolicy.EntityData.Children = make(map[string]types.YChild)
    applyPolicy.EntityData.Children["config"] = types.YChild{"Config", &applyPolicy.Config}
    applyPolicy.EntityData.Children["state"] = types.YChild{"State", &applyPolicy.State}
    applyPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", config.ImportPolicy}
    config.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", config.DefaultImportPolicy}
    config.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", config.ExportPolicy}
    config.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", config.DefaultExportPolicy}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["import-policy"] = types.YLeaf{"ImportPolicy", state.ImportPolicy}
    state.EntityData.Leafs["default-import-policy"] = types.YLeaf{"DefaultImportPolicy", state.DefaultImportPolicy}
    state.EntityData.Leafs["export-policy"] = types.YLeaf{"ExportPolicy", state.ExportPolicy}
    state.EntityData.Leafs["default-export-policy"] = types.YLeaf{"DefaultExportPolicy", state.DefaultExportPolicy}
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
    ipv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4Unicast.EntityData.Children = make(map[string]types.YChild)
    ipv4Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv4Unicast.PrefixLimit}
    ipv4Unicast.EntityData.Children["config"] = types.YChild{"Config", &ipv4Unicast.Config}
    ipv4Unicast.EntityData.Children["state"] = types.YChild{"State", &ipv4Unicast.State}
    ipv4Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute}
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
    ipv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6Unicast.EntityData.Children = make(map[string]types.YChild)
    ipv6Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv6Unicast.PrefixLimit}
    ipv6Unicast.EntityData.Children["config"] = types.YChild{"Config", &ipv6Unicast.Config}
    ipv6Unicast.EntityData.Children["state"] = types.YChild{"State", &ipv6Unicast.State}
    ipv6Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", config.SendDefaultRoute}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["send-default-route"] = types.YLeaf{"SendDefaultRoute", state.SendDefaultRoute}
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
    ipv4LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4LabeledUnicast.EntityData.Children = make(map[string]types.YChild)
    ipv4LabeledUnicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv4LabeledUnicast.PrefixLimit}
    ipv4LabeledUnicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
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
    ipv6LabeledUnicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6LabeledUnicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6LabeledUnicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6LabeledUnicast.EntityData.Children = make(map[string]types.YChild)
    ipv6LabeledUnicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &ipv6LabeledUnicast.PrefixLimit}
    ipv6LabeledUnicast.EntityData.Leafs = make(map[string]types.YLeaf)
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
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
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
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
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
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast
// Unicast IPv4 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv4Unicast.EntityData.YFilter = l3VpnIpv4Unicast.YFilter
    l3VpnIpv4Unicast.EntityData.YangName = "l3vpn-ipv4-unicast"
    l3VpnIpv4Unicast.EntityData.BundleName = "openconfig"
    l3VpnIpv4Unicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv4Unicast.EntityData.SegmentPath = "l3vpn-ipv4-unicast"
    l3VpnIpv4Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv4Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv4Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv4Unicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv4Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv4Unicast.PrefixLimit}
    l3VpnIpv4Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv4Unicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast
// Unicast IPv6 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv6Unicast.EntityData.YFilter = l3VpnIpv6Unicast.YFilter
    l3VpnIpv6Unicast.EntityData.YangName = "l3vpn-ipv6-unicast"
    l3VpnIpv6Unicast.EntityData.BundleName = "openconfig"
    l3VpnIpv6Unicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv6Unicast.EntityData.SegmentPath = "l3vpn-ipv6-unicast"
    l3VpnIpv6Unicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv6Unicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv6Unicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv6Unicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv6Unicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv6Unicast.PrefixLimit}
    l3VpnIpv6Unicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv6Unicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-unicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast
// Multicast IPv4 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv4Multicast.EntityData.YFilter = l3VpnIpv4Multicast.YFilter
    l3VpnIpv4Multicast.EntityData.YangName = "l3vpn-ipv4-multicast"
    l3VpnIpv4Multicast.EntityData.BundleName = "openconfig"
    l3VpnIpv4Multicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv4Multicast.EntityData.SegmentPath = "l3vpn-ipv4-multicast"
    l3VpnIpv4Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv4Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv4Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv4Multicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv4Multicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv4Multicast.PrefixLimit}
    l3VpnIpv4Multicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv4Multicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv4-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast
// Multicast IPv6 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetEntityData() *types.CommonEntityData {
    l3VpnIpv6Multicast.EntityData.YFilter = l3VpnIpv6Multicast.YFilter
    l3VpnIpv6Multicast.EntityData.YangName = "l3vpn-ipv6-multicast"
    l3VpnIpv6Multicast.EntityData.BundleName = "openconfig"
    l3VpnIpv6Multicast.EntityData.ParentYangName = "afi-safi"
    l3VpnIpv6Multicast.EntityData.SegmentPath = "l3vpn-ipv6-multicast"
    l3VpnIpv6Multicast.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l3VpnIpv6Multicast.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l3VpnIpv6Multicast.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l3VpnIpv6Multicast.EntityData.Children = make(map[string]types.YChild)
    l3VpnIpv6Multicast.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l3VpnIpv6Multicast.PrefixLimit}
    l3VpnIpv6Multicast.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l3VpnIpv6Multicast.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l3vpn-ipv6-multicast"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls
// BGP-signalled VPLS configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
}

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetEntityData() *types.CommonEntityData {
    l2VpnVpls.EntityData.YFilter = l2VpnVpls.YFilter
    l2VpnVpls.EntityData.YangName = "l2vpn-vpls"
    l2VpnVpls.EntityData.BundleName = "openconfig"
    l2VpnVpls.EntityData.ParentYangName = "afi-safi"
    l2VpnVpls.EntityData.SegmentPath = "l2vpn-vpls"
    l2VpnVpls.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2VpnVpls.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2VpnVpls.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2VpnVpls.EntityData.Children = make(map[string]types.YChild)
    l2VpnVpls.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l2VpnVpls.PrefixLimit}
    l2VpnVpls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l2VpnVpls.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-vpls"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn
// BGP EVPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
}

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetEntityData() *types.CommonEntityData {
    l2VpnEvpn.EntityData.YFilter = l2VpnEvpn.YFilter
    l2VpnEvpn.EntityData.YangName = "l2vpn-evpn"
    l2VpnEvpn.EntityData.BundleName = "openconfig"
    l2VpnEvpn.EntityData.ParentYangName = "afi-safi"
    l2VpnEvpn.EntityData.SegmentPath = "l2vpn-evpn"
    l2VpnEvpn.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    l2VpnEvpn.EntityData.NamespaceTable = openconfig.GetNamespaces()
    l2VpnEvpn.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    l2VpnEvpn.EntityData.Children = make(map[string]types.YChild)
    l2VpnEvpn.EntityData.Children["prefix-limit"] = types.YChild{"PrefixLimit", &l2VpnEvpn.PrefixLimit}
    l2VpnEvpn.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l2VpnEvpn.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetEntityData() *types.CommonEntityData {
    prefixLimit.EntityData.YFilter = prefixLimit.YFilter
    prefixLimit.EntityData.YangName = "prefix-limit"
    prefixLimit.EntityData.BundleName = "openconfig"
    prefixLimit.EntityData.ParentYangName = "l2vpn-evpn"
    prefixLimit.EntityData.SegmentPath = "prefix-limit"
    prefixLimit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixLimit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixLimit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixLimit.EntityData.Children = make(map[string]types.YChild)
    prefixLimit.EntityData.Children["config"] = types.YChild{"Config", &prefixLimit.Config}
    prefixLimit.EntityData.Children["state"] = types.YChild{"State", &prefixLimit.State}
    prefixLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixLimit.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-limit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", config.MaxPrefixes}
    config.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", config.ShutdownThresholdPct}
    config.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", config.RestartTimer}
    return &(config.EntityData)
}

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes that will be accepted from the neighbour. The
    // type is interface{} with range: 0..4294967295.
    MaxPrefixes interface{}

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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-limit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["max-prefixes"] = types.YLeaf{"MaxPrefixes", state.MaxPrefixes}
    state.EntityData.Leafs["shutdown-threshold-pct"] = types.YLeaf{"ShutdownThresholdPct", state.ShutdownThresholdPct}
    state.EntityData.Leafs["restart-timer"] = types.YLeaf{"RestartTimer", state.RestartTimer}
    return &(state.EntityData)
}

