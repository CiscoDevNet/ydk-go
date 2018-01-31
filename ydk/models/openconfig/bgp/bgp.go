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
    parent types.Entity
    YFilter yfilter.YFilter

    // Global configuration for the BGP router.
    Global Bgp_Global

    // Configuration for BGP neighbors.
    Neighbors Bgp_Neighbors

    // Configuration for BGP peer-groups.
    PeerGroups Bgp_PeerGroups
}

func (bgp *Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Bgp) GetGoName(yname string) string {
    if yname == "global" { return "Global" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "peer-groups" { return "PeerGroups" }
    return ""
}

func (bgp *Bgp) GetSegmentPath() string {
    return "openconfig-bgp:bgp"
}

func (bgp *Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global" {
        return &bgp.Global
    }
    if childYangName == "neighbors" {
        return &bgp.Neighbors
    }
    if childYangName == "peer-groups" {
        return &bgp.PeerGroups
    }
    return nil
}

func (bgp *Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global"] = &bgp.Global
    children["neighbors"] = &bgp.Neighbors
    children["peer-groups"] = &bgp.PeerGroups
    return children
}

func (bgp *Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp *Bgp) GetBundleName() string { return "openconfig" }

func (bgp *Bgp) GetYangName() string { return "bgp" }

func (bgp *Bgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (bgp *Bgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (bgp *Bgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (bgp *Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Bgp) GetParentYangName() string { return "openconfig-bgp" }

// Bgp_Global
// Global configuration for the BGP router
type Bgp_Global struct {
    parent types.Entity
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

func (global *Bgp_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Bgp_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Bgp_Global) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "default-route-distance" { return "DefaultRouteDistance" }
    if yname == "confederation" { return "Confederation" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "use-multiple-paths" { return "UseMultiplePaths" }
    if yname == "route-selection-options" { return "RouteSelectionOptions" }
    if yname == "afi-safis" { return "AfiSafis" }
    if yname == "apply-policy" { return "ApplyPolicy" }
    return ""
}

func (global *Bgp_Global) GetSegmentPath() string {
    return "global"
}

func (global *Bgp_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &global.Config
    }
    if childYangName == "state" {
        return &global.State
    }
    if childYangName == "default-route-distance" {
        return &global.DefaultRouteDistance
    }
    if childYangName == "confederation" {
        return &global.Confederation
    }
    if childYangName == "graceful-restart" {
        return &global.GracefulRestart
    }
    if childYangName == "use-multiple-paths" {
        return &global.UseMultiplePaths
    }
    if childYangName == "route-selection-options" {
        return &global.RouteSelectionOptions
    }
    if childYangName == "afi-safis" {
        return &global.AfiSafis
    }
    if childYangName == "apply-policy" {
        return &global.ApplyPolicy
    }
    return nil
}

func (global *Bgp_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &global.Config
    children["state"] = &global.State
    children["default-route-distance"] = &global.DefaultRouteDistance
    children["confederation"] = &global.Confederation
    children["graceful-restart"] = &global.GracefulRestart
    children["use-multiple-paths"] = &global.UseMultiplePaths
    children["route-selection-options"] = &global.RouteSelectionOptions
    children["afi-safis"] = &global.AfiSafis
    children["apply-policy"] = &global.ApplyPolicy
    return children
}

func (global *Bgp_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *Bgp_Global) GetBundleName() string { return "openconfig" }

func (global *Bgp_Global) GetYangName() string { return "global" }

func (global *Bgp_Global) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (global *Bgp_Global) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (global *Bgp_Global) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (global *Bgp_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Bgp_Global) GetParent() types.Entity { return global.parent }

func (global *Bgp_Global) GetParentYangName() string { return "bgp" }

// Bgp_Global_Config
// Configuration parameters relating to the global BGP router
type Bgp_Global_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local autonomous system number of the router.  Uses the 32-bit as-number
    // type from the model in RFC 6991. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    As interface{}

    // Router id of the router - an unsigned 32-bit integer expressed in dotted
    // quad notation. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    RouterId interface{}
}

func (config *Bgp_Global_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_Config) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "router-id" { return "RouterId" }
    return ""
}

func (config *Bgp_Global_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = config.As
    leafs["router-id"] = config.RouterId
    return leafs
}

func (config *Bgp_Global_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_Config) GetParentYangName() string { return "global" }

// Bgp_Global_State
// State information relating to the global BGP router
type Bgp_Global_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local autonomous system number of the router.  Uses the 32-bit as-number
    // type from the model in RFC 6991. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    As interface{}

    // Router id of the router - an unsigned 32-bit integer expressed in dotted
    // quad notation. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]).
    RouterId interface{}

    // Total number of BGP paths within the context. The type is interface{} with
    // range: 0..4294967295.
    TotalPaths interface{}

    // Total number of BGP prefixes received within the context. The type is
    // interface{} with range: 0..4294967295.
    TotalPrefixes interface{}
}

func (state *Bgp_Global_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_State) GetGoName(yname string) string {
    if yname == "as" { return "As" }
    if yname == "router-id" { return "RouterId" }
    if yname == "total-paths" { return "TotalPaths" }
    if yname == "total-prefixes" { return "TotalPrefixes" }
    return ""
}

func (state *Bgp_Global_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as"] = state.As
    leafs["router-id"] = state.RouterId
    leafs["total-paths"] = state.TotalPaths
    leafs["total-prefixes"] = state.TotalPrefixes
    return leafs
}

func (state *Bgp_Global_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_State) GetYangName() string { return "state" }

func (state *Bgp_Global_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_State) GetParentYangName() string { return "global" }

// Bgp_Global_DefaultRouteDistance
// Administrative distance (or preference) assigned to
// routes received from different sources
// (external, internal, and local).
type Bgp_Global_DefaultRouteDistance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the default route distance.
    Config Bgp_Global_DefaultRouteDistance_Config

    // State information relating to the default route distance.
    State Bgp_Global_DefaultRouteDistance_State
}

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetFilter() yfilter.YFilter { return defaultRouteDistance.YFilter }

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) SetFilter(yf yfilter.YFilter) { defaultRouteDistance.YFilter = yf }

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetSegmentPath() string {
    return "default-route-distance"
}

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &defaultRouteDistance.Config
    }
    if childYangName == "state" {
        return &defaultRouteDistance.State
    }
    return nil
}

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &defaultRouteDistance.Config
    children["state"] = &defaultRouteDistance.State
    return children
}

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetBundleName() string { return "openconfig" }

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetYangName() string { return "default-route-distance" }

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) SetParent(parent types.Entity) { defaultRouteDistance.parent = parent }

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetParent() types.Entity { return defaultRouteDistance.parent }

func (defaultRouteDistance *Bgp_Global_DefaultRouteDistance) GetParentYangName() string { return "global" }

// Bgp_Global_DefaultRouteDistance_Config
// Configuration parameters relating to the default route
// distance
type Bgp_Global_DefaultRouteDistance_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative distance for routes learned from external BGP (eBGP). The
    // type is interface{} with range: 1..255.
    ExternalRouteDistance interface{}

    // Administrative distance for routes learned from internal BGP (iBGP). The
    // type is interface{} with range: 1..255.
    InternalRouteDistance interface{}
}

func (config *Bgp_Global_DefaultRouteDistance_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_DefaultRouteDistance_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_DefaultRouteDistance_Config) GetGoName(yname string) string {
    if yname == "external-route-distance" { return "ExternalRouteDistance" }
    if yname == "internal-route-distance" { return "InternalRouteDistance" }
    return ""
}

func (config *Bgp_Global_DefaultRouteDistance_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_DefaultRouteDistance_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_DefaultRouteDistance_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_DefaultRouteDistance_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["external-route-distance"] = config.ExternalRouteDistance
    leafs["internal-route-distance"] = config.InternalRouteDistance
    return leafs
}

func (config *Bgp_Global_DefaultRouteDistance_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_DefaultRouteDistance_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_DefaultRouteDistance_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_DefaultRouteDistance_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_DefaultRouteDistance_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_DefaultRouteDistance_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_DefaultRouteDistance_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_DefaultRouteDistance_Config) GetParentYangName() string { return "default-route-distance" }

// Bgp_Global_DefaultRouteDistance_State
// State information relating to the default route distance
type Bgp_Global_DefaultRouteDistance_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative distance for routes learned from external BGP (eBGP). The
    // type is interface{} with range: 1..255.
    ExternalRouteDistance interface{}

    // Administrative distance for routes learned from internal BGP (iBGP). The
    // type is interface{} with range: 1..255.
    InternalRouteDistance interface{}
}

func (state *Bgp_Global_DefaultRouteDistance_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_DefaultRouteDistance_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_DefaultRouteDistance_State) GetGoName(yname string) string {
    if yname == "external-route-distance" { return "ExternalRouteDistance" }
    if yname == "internal-route-distance" { return "InternalRouteDistance" }
    return ""
}

func (state *Bgp_Global_DefaultRouteDistance_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_DefaultRouteDistance_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_DefaultRouteDistance_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_DefaultRouteDistance_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["external-route-distance"] = state.ExternalRouteDistance
    leafs["internal-route-distance"] = state.InternalRouteDistance
    return leafs
}

func (state *Bgp_Global_DefaultRouteDistance_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_DefaultRouteDistance_State) GetYangName() string { return "state" }

func (state *Bgp_Global_DefaultRouteDistance_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_DefaultRouteDistance_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_DefaultRouteDistance_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_DefaultRouteDistance_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_DefaultRouteDistance_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_DefaultRouteDistance_State) GetParentYangName() string { return "default-route-distance" }

// Bgp_Global_Confederation
// Parameters indicating whether the local system acts as part
// of a BGP confederation
type Bgp_Global_Confederation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to BGP confederations.
    Config Bgp_Global_Confederation_Config

    // State information relating to the BGP confederations.
    State Bgp_Global_Confederation_State
}

func (confederation *Bgp_Global_Confederation) GetFilter() yfilter.YFilter { return confederation.YFilter }

func (confederation *Bgp_Global_Confederation) SetFilter(yf yfilter.YFilter) { confederation.YFilter = yf }

func (confederation *Bgp_Global_Confederation) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (confederation *Bgp_Global_Confederation) GetSegmentPath() string {
    return "confederation"
}

func (confederation *Bgp_Global_Confederation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &confederation.Config
    }
    if childYangName == "state" {
        return &confederation.State
    }
    return nil
}

func (confederation *Bgp_Global_Confederation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &confederation.Config
    children["state"] = &confederation.State
    return children
}

func (confederation *Bgp_Global_Confederation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (confederation *Bgp_Global_Confederation) GetBundleName() string { return "openconfig" }

func (confederation *Bgp_Global_Confederation) GetYangName() string { return "confederation" }

func (confederation *Bgp_Global_Confederation) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (confederation *Bgp_Global_Confederation) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (confederation *Bgp_Global_Confederation) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (confederation *Bgp_Global_Confederation) SetParent(parent types.Entity) { confederation.parent = parent }

func (confederation *Bgp_Global_Confederation) GetParent() types.Entity { return confederation.parent }

func (confederation *Bgp_Global_Confederation) GetParentYangName() string { return "global" }

// Bgp_Global_Confederation_Config
// Configuration parameters relating to BGP confederations
type Bgp_Global_Confederation_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_Confederation_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_Confederation_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_Confederation_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "identifier" { return "Identifier" }
    if yname == "member-as" { return "MemberAs" }
    return ""
}

func (config *Bgp_Global_Confederation_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_Confederation_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_Confederation_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_Confederation_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["identifier"] = config.Identifier
    leafs["member-as"] = config.MemberAs
    return leafs
}

func (config *Bgp_Global_Confederation_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_Confederation_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_Confederation_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_Confederation_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_Confederation_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_Confederation_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_Confederation_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_Confederation_Config) GetParentYangName() string { return "confederation" }

// Bgp_Global_Confederation_State
// State information relating to the BGP confederations
type Bgp_Global_Confederation_State struct {
    parent types.Entity
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

func (state *Bgp_Global_Confederation_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_Confederation_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_Confederation_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "identifier" { return "Identifier" }
    if yname == "member-as" { return "MemberAs" }
    return ""
}

func (state *Bgp_Global_Confederation_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_Confederation_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_Confederation_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_Confederation_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["identifier"] = state.Identifier
    leafs["member-as"] = state.MemberAs
    return leafs
}

func (state *Bgp_Global_Confederation_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_Confederation_State) GetYangName() string { return "state" }

func (state *Bgp_Global_Confederation_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_Confederation_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_Confederation_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_Confederation_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_Confederation_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_Confederation_State) GetParentYangName() string { return "confederation" }

// Bgp_Global_GracefulRestart
// Parameters relating the graceful restart mechanism for BGP
type Bgp_Global_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to graceful-restart.
    Config Bgp_Global_GracefulRestart_Config

    // State information associated with graceful-restart.
    State Bgp_Global_GracefulRestart_State
}

func (gracefulRestart *Bgp_Global_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *Bgp_Global_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *Bgp_Global_GracefulRestart) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (gracefulRestart *Bgp_Global_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *Bgp_Global_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &gracefulRestart.Config
    }
    if childYangName == "state" {
        return &gracefulRestart.State
    }
    return nil
}

func (gracefulRestart *Bgp_Global_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &gracefulRestart.Config
    children["state"] = &gracefulRestart.State
    return children
}

func (gracefulRestart *Bgp_Global_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gracefulRestart *Bgp_Global_GracefulRestart) GetBundleName() string { return "openconfig" }

func (gracefulRestart *Bgp_Global_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *Bgp_Global_GracefulRestart) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (gracefulRestart *Bgp_Global_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (gracefulRestart *Bgp_Global_GracefulRestart) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (gracefulRestart *Bgp_Global_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *Bgp_Global_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *Bgp_Global_GracefulRestart) GetParentYangName() string { return "global" }

// Bgp_Global_GracefulRestart_Config
// Configuration parameters relating to graceful-restart
type Bgp_Global_GracefulRestart_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_GracefulRestart_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_GracefulRestart_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_GracefulRestart_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "restart-time" { return "RestartTime" }
    if yname == "stale-routes-time" { return "StaleRoutesTime" }
    if yname == "helper-only" { return "HelperOnly" }
    return ""
}

func (config *Bgp_Global_GracefulRestart_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_GracefulRestart_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_GracefulRestart_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_GracefulRestart_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["restart-time"] = config.RestartTime
    leafs["stale-routes-time"] = config.StaleRoutesTime
    leafs["helper-only"] = config.HelperOnly
    return leafs
}

func (config *Bgp_Global_GracefulRestart_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_GracefulRestart_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_GracefulRestart_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_GracefulRestart_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_GracefulRestart_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_GracefulRestart_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_GracefulRestart_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_GracefulRestart_Config) GetParentYangName() string { return "graceful-restart" }

// Bgp_Global_GracefulRestart_State
// State information associated with graceful-restart
type Bgp_Global_GracefulRestart_State struct {
    parent types.Entity
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

func (state *Bgp_Global_GracefulRestart_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_GracefulRestart_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_GracefulRestart_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "restart-time" { return "RestartTime" }
    if yname == "stale-routes-time" { return "StaleRoutesTime" }
    if yname == "helper-only" { return "HelperOnly" }
    return ""
}

func (state *Bgp_Global_GracefulRestart_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_GracefulRestart_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_GracefulRestart_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_GracefulRestart_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["restart-time"] = state.RestartTime
    leafs["stale-routes-time"] = state.StaleRoutesTime
    leafs["helper-only"] = state.HelperOnly
    return leafs
}

func (state *Bgp_Global_GracefulRestart_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_GracefulRestart_State) GetYangName() string { return "state" }

func (state *Bgp_Global_GracefulRestart_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_GracefulRestart_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_GracefulRestart_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_GracefulRestart_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_GracefulRestart_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_GracefulRestart_State) GetParentYangName() string { return "graceful-restart" }

// Bgp_Global_UseMultiplePaths
// Parameters related to the use of multiple paths for the
// same NLRI
type Bgp_Global_UseMultiplePaths struct {
    parent types.Entity
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

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetFilter() yfilter.YFilter { return useMultiplePaths.YFilter }

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) SetFilter(yf yfilter.YFilter) { useMultiplePaths.YFilter = yf }

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "ebgp" { return "Ebgp" }
    if yname == "ibgp" { return "Ibgp" }
    return ""
}

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetSegmentPath() string {
    return "use-multiple-paths"
}

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &useMultiplePaths.Config
    }
    if childYangName == "state" {
        return &useMultiplePaths.State
    }
    if childYangName == "ebgp" {
        return &useMultiplePaths.Ebgp
    }
    if childYangName == "ibgp" {
        return &useMultiplePaths.Ibgp
    }
    return nil
}

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &useMultiplePaths.Config
    children["state"] = &useMultiplePaths.State
    children["ebgp"] = &useMultiplePaths.Ebgp
    children["ibgp"] = &useMultiplePaths.Ibgp
    return children
}

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetBundleName() string { return "openconfig" }

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetYangName() string { return "use-multiple-paths" }

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) SetParent(parent types.Entity) { useMultiplePaths.parent = parent }

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetParent() types.Entity { return useMultiplePaths.parent }

func (useMultiplePaths *Bgp_Global_UseMultiplePaths) GetParentYangName() string { return "global" }

// Bgp_Global_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_Global_UseMultiplePaths_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Global_UseMultiplePaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_UseMultiplePaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_UseMultiplePaths_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_Global_UseMultiplePaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_UseMultiplePaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_UseMultiplePaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_UseMultiplePaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_Global_UseMultiplePaths_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_UseMultiplePaths_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_UseMultiplePaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_UseMultiplePaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_UseMultiplePaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_UseMultiplePaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_UseMultiplePaths_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_UseMultiplePaths_Config) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Global_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_Global_UseMultiplePaths_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Global_UseMultiplePaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_UseMultiplePaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_UseMultiplePaths_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Bgp_Global_UseMultiplePaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_UseMultiplePaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_UseMultiplePaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_UseMultiplePaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Bgp_Global_UseMultiplePaths_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_UseMultiplePaths_State) GetYangName() string { return "state" }

func (state *Bgp_Global_UseMultiplePaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_UseMultiplePaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_UseMultiplePaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_UseMultiplePaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_UseMultiplePaths_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_UseMultiplePaths_State) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Global_UseMultiplePaths_Ebgp
// Multipath parameters for eBGP
type Bgp_Global_UseMultiplePaths_Ebgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_Global_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_Global_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetFilter() yfilter.YFilter { return ebgp.YFilter }

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) SetFilter(yf yfilter.YFilter) { ebgp.YFilter = yf }

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetSegmentPath() string {
    return "ebgp"
}

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ebgp.Config
    }
    if childYangName == "state" {
        return &ebgp.State
    }
    return nil
}

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ebgp.Config
    children["state"] = &ebgp.State
    return children
}

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetBundleName() string { return "openconfig" }

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetYangName() string { return "ebgp" }

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) SetParent(parent types.Entity) { ebgp.parent = parent }

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetParent() types.Entity { return ebgp.parent }

func (ebgp *Bgp_Global_UseMultiplePaths_Ebgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Global_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_Global_UseMultiplePaths_Ebgp_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = config.AllowMultipleAs
    leafs["maximum-paths"] = config.MaximumPaths
    return leafs
}

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_UseMultiplePaths_Ebgp_Config) GetParentYangName() string { return "ebgp" }

// Bgp_Global_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_Global_UseMultiplePaths_Ebgp_State struct {
    parent types.Entity
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

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = state.AllowMultipleAs
    leafs["maximum-paths"] = state.MaximumPaths
    return leafs
}

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetYangName() string { return "state" }

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_UseMultiplePaths_Ebgp_State) GetParentYangName() string { return "ebgp" }

// Bgp_Global_UseMultiplePaths_Ibgp
// Multipath parameters for iBGP
type Bgp_Global_UseMultiplePaths_Ibgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to iBGP multipath.
    Config Bgp_Global_UseMultiplePaths_Ibgp_Config

    // State information relating to iBGP multipath.
    State Bgp_Global_UseMultiplePaths_Ibgp_State
}

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetFilter() yfilter.YFilter { return ibgp.YFilter }

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) SetFilter(yf yfilter.YFilter) { ibgp.YFilter = yf }

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetSegmentPath() string {
    return "ibgp"
}

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ibgp.Config
    }
    if childYangName == "state" {
        return &ibgp.State
    }
    return nil
}

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ibgp.Config
    children["state"] = &ibgp.State
    return children
}

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetBundleName() string { return "openconfig" }

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetYangName() string { return "ibgp" }

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) SetParent(parent types.Entity) { ibgp.parent = parent }

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetParent() types.Entity { return ibgp.parent }

func (ibgp *Bgp_Global_UseMultiplePaths_Ibgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Global_UseMultiplePaths_Ibgp_Config
// Configuration parameters relating to iBGP multipath
type Bgp_Global_UseMultiplePaths_Ibgp_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetGoName(yname string) string {
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-paths"] = config.MaximumPaths
    return leafs
}

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_UseMultiplePaths_Ibgp_Config) GetParentYangName() string { return "ibgp" }

// Bgp_Global_UseMultiplePaths_Ibgp_State
// State information relating to iBGP multipath
type Bgp_Global_UseMultiplePaths_Ibgp_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetGoName(yname string) string {
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-paths"] = state.MaximumPaths
    return leafs
}

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetYangName() string { return "state" }

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_UseMultiplePaths_Ibgp_State) GetParentYangName() string { return "ibgp" }

// Bgp_Global_RouteSelectionOptions
// Parameters relating to options for route selection
type Bgp_Global_RouteSelectionOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to route selection options.
    Config Bgp_Global_RouteSelectionOptions_Config

    // State information for the route selection options.
    State Bgp_Global_RouteSelectionOptions_State
}

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetFilter() yfilter.YFilter { return routeSelectionOptions.YFilter }

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) SetFilter(yf yfilter.YFilter) { routeSelectionOptions.YFilter = yf }

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetSegmentPath() string {
    return "route-selection-options"
}

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &routeSelectionOptions.Config
    }
    if childYangName == "state" {
        return &routeSelectionOptions.State
    }
    return nil
}

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &routeSelectionOptions.Config
    children["state"] = &routeSelectionOptions.State
    return children
}

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetBundleName() string { return "openconfig" }

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetYangName() string { return "route-selection-options" }

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) SetParent(parent types.Entity) { routeSelectionOptions.parent = parent }

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetParent() types.Entity { return routeSelectionOptions.parent }

func (routeSelectionOptions *Bgp_Global_RouteSelectionOptions) GetParentYangName() string { return "global" }

// Bgp_Global_RouteSelectionOptions_Config
// Configuration parameters relating to route selection
// options
type Bgp_Global_RouteSelectionOptions_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_RouteSelectionOptions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_RouteSelectionOptions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_RouteSelectionOptions_Config) GetGoName(yname string) string {
    if yname == "always-compare-med" { return "AlwaysCompareMed" }
    if yname == "ignore-as-path-length" { return "IgnoreAsPathLength" }
    if yname == "external-compare-router-id" { return "ExternalCompareRouterId" }
    if yname == "advertise-inactive-routes" { return "AdvertiseInactiveRoutes" }
    if yname == "enable-aigp" { return "EnableAigp" }
    if yname == "ignore-next-hop-igp-metric" { return "IgnoreNextHopIgpMetric" }
    return ""
}

func (config *Bgp_Global_RouteSelectionOptions_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_RouteSelectionOptions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_RouteSelectionOptions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_RouteSelectionOptions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["always-compare-med"] = config.AlwaysCompareMed
    leafs["ignore-as-path-length"] = config.IgnoreAsPathLength
    leafs["external-compare-router-id"] = config.ExternalCompareRouterId
    leafs["advertise-inactive-routes"] = config.AdvertiseInactiveRoutes
    leafs["enable-aigp"] = config.EnableAigp
    leafs["ignore-next-hop-igp-metric"] = config.IgnoreNextHopIgpMetric
    return leafs
}

func (config *Bgp_Global_RouteSelectionOptions_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_RouteSelectionOptions_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_RouteSelectionOptions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_RouteSelectionOptions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_RouteSelectionOptions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_RouteSelectionOptions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_RouteSelectionOptions_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_RouteSelectionOptions_Config) GetParentYangName() string { return "route-selection-options" }

// Bgp_Global_RouteSelectionOptions_State
// State information for the route selection options
type Bgp_Global_RouteSelectionOptions_State struct {
    parent types.Entity
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

func (state *Bgp_Global_RouteSelectionOptions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_RouteSelectionOptions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_RouteSelectionOptions_State) GetGoName(yname string) string {
    if yname == "always-compare-med" { return "AlwaysCompareMed" }
    if yname == "ignore-as-path-length" { return "IgnoreAsPathLength" }
    if yname == "external-compare-router-id" { return "ExternalCompareRouterId" }
    if yname == "advertise-inactive-routes" { return "AdvertiseInactiveRoutes" }
    if yname == "enable-aigp" { return "EnableAigp" }
    if yname == "ignore-next-hop-igp-metric" { return "IgnoreNextHopIgpMetric" }
    return ""
}

func (state *Bgp_Global_RouteSelectionOptions_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_RouteSelectionOptions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_RouteSelectionOptions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_RouteSelectionOptions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["always-compare-med"] = state.AlwaysCompareMed
    leafs["ignore-as-path-length"] = state.IgnoreAsPathLength
    leafs["external-compare-router-id"] = state.ExternalCompareRouterId
    leafs["advertise-inactive-routes"] = state.AdvertiseInactiveRoutes
    leafs["enable-aigp"] = state.EnableAigp
    leafs["ignore-next-hop-igp-metric"] = state.IgnoreNextHopIgpMetric
    return leafs
}

func (state *Bgp_Global_RouteSelectionOptions_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_RouteSelectionOptions_State) GetYangName() string { return "state" }

func (state *Bgp_Global_RouteSelectionOptions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_RouteSelectionOptions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_RouteSelectionOptions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_RouteSelectionOptions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_RouteSelectionOptions_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_RouteSelectionOptions_State) GetParentYangName() string { return "route-selection-options" }

// Bgp_Global_AfiSafis
// Address family specific configuration
type Bgp_Global_AfiSafis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI,SAFI configuration available for the neighbour or group. The type is
    // slice of Bgp_Global_AfiSafis_AfiSafi.
    AfiSafi []Bgp_Global_AfiSafis_AfiSafi
}

func (afiSafis *Bgp_Global_AfiSafis) GetFilter() yfilter.YFilter { return afiSafis.YFilter }

func (afiSafis *Bgp_Global_AfiSafis) SetFilter(yf yfilter.YFilter) { afiSafis.YFilter = yf }

func (afiSafis *Bgp_Global_AfiSafis) GetGoName(yname string) string {
    if yname == "afi-safi" { return "AfiSafi" }
    return ""
}

func (afiSafis *Bgp_Global_AfiSafis) GetSegmentPath() string {
    return "afi-safis"
}

func (afiSafis *Bgp_Global_AfiSafis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afi-safi" {
        for _, c := range afiSafis.AfiSafi {
            if afiSafis.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Bgp_Global_AfiSafis_AfiSafi{}
        afiSafis.AfiSafi = append(afiSafis.AfiSafi, child)
        return &afiSafis.AfiSafi[len(afiSafis.AfiSafi)-1]
    }
    return nil
}

func (afiSafis *Bgp_Global_AfiSafis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afiSafis.AfiSafi {
        children[afiSafis.AfiSafi[i].GetSegmentPath()] = &afiSafis.AfiSafi[i]
    }
    return children
}

func (afiSafis *Bgp_Global_AfiSafis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afiSafis *Bgp_Global_AfiSafis) GetBundleName() string { return "openconfig" }

func (afiSafis *Bgp_Global_AfiSafis) GetYangName() string { return "afi-safis" }

func (afiSafis *Bgp_Global_AfiSafis) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (afiSafis *Bgp_Global_AfiSafis) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (afiSafis *Bgp_Global_AfiSafis) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (afiSafis *Bgp_Global_AfiSafis) SetParent(parent types.Entity) { afiSafis.parent = parent }

func (afiSafis *Bgp_Global_AfiSafis) GetParent() types.Entity { return afiSafis.parent }

func (afiSafis *Bgp_Global_AfiSafis) GetParentYangName() string { return "global" }

// Bgp_Global_AfiSafis_AfiSafi
// AFI,SAFI configuration available for the
// neighbour or group
type Bgp_Global_AfiSafis_AfiSafi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the AFI-SAFI name used as a key for
    // the AFI-SAFI list. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
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

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetFilter() yfilter.YFilter { return afiSafi.YFilter }

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) SetFilter(yf yfilter.YFilter) { afiSafi.YFilter = yf }

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "route-selection-options" { return "RouteSelectionOptions" }
    if yname == "use-multiple-paths" { return "UseMultiplePaths" }
    if yname == "apply-policy" { return "ApplyPolicy" }
    if yname == "ipv4-unicast" { return "Ipv4Unicast" }
    if yname == "ipv6-unicast" { return "Ipv6Unicast" }
    if yname == "ipv4-labeled-unicast" { return "Ipv4LabeledUnicast" }
    if yname == "ipv6-labeled-unicast" { return "Ipv6LabeledUnicast" }
    if yname == "l3vpn-ipv4-unicast" { return "L3VpnIpv4Unicast" }
    if yname == "l3vpn-ipv6-unicast" { return "L3VpnIpv6Unicast" }
    if yname == "l3vpn-ipv4-multicast" { return "L3VpnIpv4Multicast" }
    if yname == "l3vpn-ipv6-multicast" { return "L3VpnIpv6Multicast" }
    if yname == "l2vpn-vpls" { return "L2VpnVpls" }
    if yname == "l2vpn-evpn" { return "L2VpnEvpn" }
    return ""
}

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetSegmentPath() string {
    return "afi-safi" + "[afi-safi-name='" + fmt.Sprintf("%v", afiSafi.AfiSafiName) + "']"
}

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &afiSafi.Config
    }
    if childYangName == "state" {
        return &afiSafi.State
    }
    if childYangName == "graceful-restart" {
        return &afiSafi.GracefulRestart
    }
    if childYangName == "route-selection-options" {
        return &afiSafi.RouteSelectionOptions
    }
    if childYangName == "use-multiple-paths" {
        return &afiSafi.UseMultiplePaths
    }
    if childYangName == "apply-policy" {
        return &afiSafi.ApplyPolicy
    }
    if childYangName == "ipv4-unicast" {
        return &afiSafi.Ipv4Unicast
    }
    if childYangName == "ipv6-unicast" {
        return &afiSafi.Ipv6Unicast
    }
    if childYangName == "ipv4-labeled-unicast" {
        return &afiSafi.Ipv4LabeledUnicast
    }
    if childYangName == "ipv6-labeled-unicast" {
        return &afiSafi.Ipv6LabeledUnicast
    }
    if childYangName == "l3vpn-ipv4-unicast" {
        return &afiSafi.L3VpnIpv4Unicast
    }
    if childYangName == "l3vpn-ipv6-unicast" {
        return &afiSafi.L3VpnIpv6Unicast
    }
    if childYangName == "l3vpn-ipv4-multicast" {
        return &afiSafi.L3VpnIpv4Multicast
    }
    if childYangName == "l3vpn-ipv6-multicast" {
        return &afiSafi.L3VpnIpv6Multicast
    }
    if childYangName == "l2vpn-vpls" {
        return &afiSafi.L2VpnVpls
    }
    if childYangName == "l2vpn-evpn" {
        return &afiSafi.L2VpnEvpn
    }
    return nil
}

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &afiSafi.Config
    children["state"] = &afiSafi.State
    children["graceful-restart"] = &afiSafi.GracefulRestart
    children["route-selection-options"] = &afiSafi.RouteSelectionOptions
    children["use-multiple-paths"] = &afiSafi.UseMultiplePaths
    children["apply-policy"] = &afiSafi.ApplyPolicy
    children["ipv4-unicast"] = &afiSafi.Ipv4Unicast
    children["ipv6-unicast"] = &afiSafi.Ipv6Unicast
    children["ipv4-labeled-unicast"] = &afiSafi.Ipv4LabeledUnicast
    children["ipv6-labeled-unicast"] = &afiSafi.Ipv6LabeledUnicast
    children["l3vpn-ipv4-unicast"] = &afiSafi.L3VpnIpv4Unicast
    children["l3vpn-ipv6-unicast"] = &afiSafi.L3VpnIpv6Unicast
    children["l3vpn-ipv4-multicast"] = &afiSafi.L3VpnIpv4Multicast
    children["l3vpn-ipv6-multicast"] = &afiSafi.L3VpnIpv6Multicast
    children["l2vpn-vpls"] = &afiSafi.L2VpnVpls
    children["l2vpn-evpn"] = &afiSafi.L2VpnEvpn
    return children
}

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = afiSafi.AfiSafiName
    return leafs
}

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetBundleName() string { return "openconfig" }

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetYangName() string { return "afi-safi" }

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) SetParent(parent types.Entity) { afiSafi.parent = parent }

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetParent() types.Entity { return afiSafi.parent }

func (afiSafi *Bgp_Global_AfiSafis_AfiSafi) GetParentYangName() string { return "afi-safis" }

// Bgp_Global_AfiSafis_AfiSafi_Config
// Configuration parameters for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = config.AfiSafiName
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Config) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_State
// State information relating to the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
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

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "enabled" { return "Enabled" }
    if yname == "total-paths" { return "TotalPaths" }
    if yname == "total-prefixes" { return "TotalPrefixes" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = state.AfiSafiName
    leafs["enabled"] = state.Enabled
    leafs["total-paths"] = state.TotalPaths
    leafs["total-prefixes"] = state.TotalPrefixes
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_State) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_GracefulRestart
// Parameters relating to BGP graceful-restart
type Bgp_Global_AfiSafis_AfiSafi_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration options for BGP graceful-restart.
    Config Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config

    // State information for BGP graceful-restart.
    State Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State
}

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &gracefulRestart.Config
    }
    if childYangName == "state" {
        return &gracefulRestart.State
    }
    return nil
}

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &gracefulRestart.Config
    children["state"] = &gracefulRestart.State
    return children
}

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetBundleName() string { return "openconfig" }

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config
// Configuration options for BGP graceful-restart
type Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_Config) GetParentYangName() string { return "graceful-restart" }

// Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State
// State information for BGP graceful-restart
type Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_GracefulRestart_State) GetParentYangName() string { return "graceful-restart" }

// Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions
// Parameters relating to options for route selection
type Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to route selection options.
    Config Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config

    // State information for the route selection options.
    State Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State
}

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetFilter() yfilter.YFilter { return routeSelectionOptions.YFilter }

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) SetFilter(yf yfilter.YFilter) { routeSelectionOptions.YFilter = yf }

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetSegmentPath() string {
    return "route-selection-options"
}

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &routeSelectionOptions.Config
    }
    if childYangName == "state" {
        return &routeSelectionOptions.State
    }
    return nil
}

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &routeSelectionOptions.Config
    children["state"] = &routeSelectionOptions.State
    return children
}

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetBundleName() string { return "openconfig" }

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetYangName() string { return "route-selection-options" }

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) SetParent(parent types.Entity) { routeSelectionOptions.parent = parent }

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetParent() types.Entity { return routeSelectionOptions.parent }

func (routeSelectionOptions *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config
// Configuration parameters relating to route selection
// options
type Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetGoName(yname string) string {
    if yname == "always-compare-med" { return "AlwaysCompareMed" }
    if yname == "ignore-as-path-length" { return "IgnoreAsPathLength" }
    if yname == "external-compare-router-id" { return "ExternalCompareRouterId" }
    if yname == "advertise-inactive-routes" { return "AdvertiseInactiveRoutes" }
    if yname == "enable-aigp" { return "EnableAigp" }
    if yname == "ignore-next-hop-igp-metric" { return "IgnoreNextHopIgpMetric" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["always-compare-med"] = config.AlwaysCompareMed
    leafs["ignore-as-path-length"] = config.IgnoreAsPathLength
    leafs["external-compare-router-id"] = config.ExternalCompareRouterId
    leafs["advertise-inactive-routes"] = config.AdvertiseInactiveRoutes
    leafs["enable-aigp"] = config.EnableAigp
    leafs["ignore-next-hop-igp-metric"] = config.IgnoreNextHopIgpMetric
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetParentYangName() string { return "route-selection-options" }

// Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State
// State information for the route selection options
type Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetGoName(yname string) string {
    if yname == "always-compare-med" { return "AlwaysCompareMed" }
    if yname == "ignore-as-path-length" { return "IgnoreAsPathLength" }
    if yname == "external-compare-router-id" { return "ExternalCompareRouterId" }
    if yname == "advertise-inactive-routes" { return "AdvertiseInactiveRoutes" }
    if yname == "enable-aigp" { return "EnableAigp" }
    if yname == "ignore-next-hop-igp-metric" { return "IgnoreNextHopIgpMetric" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["always-compare-med"] = state.AlwaysCompareMed
    leafs["ignore-as-path-length"] = state.IgnoreAsPathLength
    leafs["external-compare-router-id"] = state.ExternalCompareRouterId
    leafs["advertise-inactive-routes"] = state.AdvertiseInactiveRoutes
    leafs["enable-aigp"] = state.EnableAigp
    leafs["ignore-next-hop-igp-metric"] = state.IgnoreNextHopIgpMetric
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetParentYangName() string { return "route-selection-options" }

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths
// Parameters related to the use of multiple paths for the
// same NLRI
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths struct {
    parent types.Entity
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

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetFilter() yfilter.YFilter { return useMultiplePaths.YFilter }

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) SetFilter(yf yfilter.YFilter) { useMultiplePaths.YFilter = yf }

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "ebgp" { return "Ebgp" }
    if yname == "ibgp" { return "Ibgp" }
    return ""
}

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetSegmentPath() string {
    return "use-multiple-paths"
}

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &useMultiplePaths.Config
    }
    if childYangName == "state" {
        return &useMultiplePaths.State
    }
    if childYangName == "ebgp" {
        return &useMultiplePaths.Ebgp
    }
    if childYangName == "ibgp" {
        return &useMultiplePaths.Ibgp
    }
    return nil
}

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &useMultiplePaths.Config
    children["state"] = &useMultiplePaths.State
    children["ebgp"] = &useMultiplePaths.Ebgp
    children["ibgp"] = &useMultiplePaths.Ibgp
    return children
}

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetBundleName() string { return "openconfig" }

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetYangName() string { return "use-multiple-paths" }

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) SetParent(parent types.Entity) { useMultiplePaths.parent = parent }

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetParent() types.Entity { return useMultiplePaths.parent }

func (useMultiplePaths *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_State) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp
// Multipath parameters for eBGP
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetFilter() yfilter.YFilter { return ebgp.YFilter }

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) SetFilter(yf yfilter.YFilter) { ebgp.YFilter = yf }

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetSegmentPath() string {
    return "ebgp"
}

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ebgp.Config
    }
    if childYangName == "state" {
        return &ebgp.State
    }
    return nil
}

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ebgp.Config
    children["state"] = &ebgp.State
    return children
}

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetBundleName() string { return "openconfig" }

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetYangName() string { return "ebgp" }

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) SetParent(parent types.Entity) { ebgp.parent = parent }

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetParent() types.Entity { return ebgp.parent }

func (ebgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = config.AllowMultipleAs
    leafs["maximum-paths"] = config.MaximumPaths
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetParentYangName() string { return "ebgp" }

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = state.AllowMultipleAs
    leafs["maximum-paths"] = state.MaximumPaths
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetParentYangName() string { return "ebgp" }

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp
// Multipath parameters for iBGP
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to iBGP multipath.
    Config Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config

    // State information relating to iBGP multipath.
    State Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State
}

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetFilter() yfilter.YFilter { return ibgp.YFilter }

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) SetFilter(yf yfilter.YFilter) { ibgp.YFilter = yf }

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetSegmentPath() string {
    return "ibgp"
}

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ibgp.Config
    }
    if childYangName == "state" {
        return &ibgp.State
    }
    return nil
}

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ibgp.Config
    children["state"] = &ibgp.State
    return children
}

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetBundleName() string { return "openconfig" }

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetYangName() string { return "ibgp" }

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) SetParent(parent types.Entity) { ibgp.parent = parent }

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetParent() types.Entity { return ibgp.parent }

func (ibgp *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config
// Configuration parameters relating to iBGP multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetGoName(yname string) string {
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-paths"] = config.MaximumPaths
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetParentYangName() string { return "ibgp" }

// Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State
// State information relating to iBGP multipath
type Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetGoName(yname string) string {
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-paths"] = state.MaximumPaths
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetParentYangName() string { return "ibgp" }

// Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State
}

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetFilter() yfilter.YFilter { return applyPolicy.YFilter }

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) SetFilter(yf yfilter.YFilter) { applyPolicy.YFilter = yf }

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetSegmentPath() string {
    return "apply-policy"
}

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &applyPolicy.Config
    }
    if childYangName == "state" {
        return &applyPolicy.State
    }
    return nil
}

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &applyPolicy.Config
    children["state"] = &applyPolicy.State
    return children
}

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetBundleName() string { return "openconfig" }

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetYangName() string { return "apply-policy" }

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) SetParent(parent types.Entity) { applyPolicy.parent = parent }

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetParent() types.Entity { return applyPolicy.parent }

func (applyPolicy *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config
// Policy configuration data.
type Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = config.ImportPolicy
    leafs["default-import-policy"] = config.DefaultImportPolicy
    leafs["export-policy"] = config.ExportPolicy
    leafs["default-export-policy"] = config.DefaultExportPolicy
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_Config) GetParentYangName() string { return "apply-policy" }

// Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State
// Operational state for routing policy
type Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = state.ImportPolicy
    leafs["default-import-policy"] = state.DefaultImportPolicy
    leafs["export-policy"] = state.ExportPolicy
    leafs["default-export-policy"] = state.DefaultExportPolicy
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_ApplyPolicy_State) GetParentYangName() string { return "apply-policy" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast
// IPv4 unicast configuration options
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State
}

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetFilter() yfilter.YFilter { return ipv4Unicast.YFilter }

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) SetFilter(yf yfilter.YFilter) { ipv4Unicast.YFilter = yf }

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetSegmentPath() string {
    return "ipv4-unicast"
}

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv4Unicast.PrefixLimit
    }
    if childYangName == "config" {
        return &ipv4Unicast.Config
    }
    if childYangName == "state" {
        return &ipv4Unicast.State
    }
    return nil
}

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv4Unicast.PrefixLimit
    children["config"] = &ipv4Unicast.Config
    children["state"] = &ipv4Unicast.State
    return children
}

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetBundleName() string { return "openconfig" }

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetYangName() string { return "ipv4-unicast" }

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) SetParent(parent types.Entity) { ipv4Unicast.parent = parent }

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetParent() types.Entity { return ipv4Unicast.parent }

func (ipv4Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetParentYangName() string { return "ipv4-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = config.SendDefaultRoute
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetParentYangName() string { return "ipv4-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = state.SendDefaultRoute
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4Unicast_State) GetParentYangName() string { return "ipv4-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast
// IPv6 unicast configuration options
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State
}

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetFilter() yfilter.YFilter { return ipv6Unicast.YFilter }

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) SetFilter(yf yfilter.YFilter) { ipv6Unicast.YFilter = yf }

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetSegmentPath() string {
    return "ipv6-unicast"
}

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv6Unicast.PrefixLimit
    }
    if childYangName == "config" {
        return &ipv6Unicast.Config
    }
    if childYangName == "state" {
        return &ipv6Unicast.State
    }
    return nil
}

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv6Unicast.PrefixLimit
    children["config"] = &ipv6Unicast.Config
    children["state"] = &ipv6Unicast.State
    return children
}

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetBundleName() string { return "openconfig" }

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetYangName() string { return "ipv6-unicast" }

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) SetParent(parent types.Entity) { ipv6Unicast.parent = parent }

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetParent() types.Entity { return ipv6Unicast.parent }

func (ipv6Unicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetParentYangName() string { return "ipv6-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = config.SendDefaultRoute
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetParentYangName() string { return "ipv6-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = state.SendDefaultRoute
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6Unicast_State) GetParentYangName() string { return "ipv6-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast
// IPv4 Labeled Unicast configuration options
type Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetFilter() yfilter.YFilter { return ipv4LabeledUnicast.YFilter }

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) SetFilter(yf yfilter.YFilter) { ipv4LabeledUnicast.YFilter = yf }

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetSegmentPath() string {
    return "ipv4-labeled-unicast"
}

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv4LabeledUnicast.PrefixLimit
    }
    return nil
}

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv4LabeledUnicast.PrefixLimit
    return children
}

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetBundleName() string { return "openconfig" }

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetYangName() string { return "ipv4-labeled-unicast" }

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) SetParent(parent types.Entity) { ipv4LabeledUnicast.parent = parent }

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetParent() types.Entity { return ipv4LabeledUnicast.parent }

func (ipv4LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetParentYangName() string { return "ipv4-labeled-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast
// IPv6 Labeled Unicast configuration options
type Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetFilter() yfilter.YFilter { return ipv6LabeledUnicast.YFilter }

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) SetFilter(yf yfilter.YFilter) { ipv6LabeledUnicast.YFilter = yf }

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetSegmentPath() string {
    return "ipv6-labeled-unicast"
}

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv6LabeledUnicast.PrefixLimit
    }
    return nil
}

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv6LabeledUnicast.PrefixLimit
    return children
}

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetBundleName() string { return "openconfig" }

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetYangName() string { return "ipv6-labeled-unicast" }

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) SetParent(parent types.Entity) { ipv6LabeledUnicast.parent = parent }

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetParent() types.Entity { return ipv6LabeledUnicast.parent }

func (ipv6LabeledUnicast *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetParentYangName() string { return "ipv6-labeled-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast
// Unicast IPv4 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetFilter() yfilter.YFilter { return l3VpnIpv4Unicast.YFilter }

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv4Unicast.YFilter = yf }

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetSegmentPath() string {
    return "l3vpn-ipv4-unicast"
}

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv4Unicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv4Unicast.PrefixLimit
    return children
}

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetYangName() string { return "l3vpn-ipv4-unicast" }

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) SetParent(parent types.Entity) { l3VpnIpv4Unicast.parent = parent }

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetParent() types.Entity { return l3VpnIpv4Unicast.parent }

func (l3VpnIpv4Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv4-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast
// Unicast IPv6 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetFilter() yfilter.YFilter { return l3VpnIpv6Unicast.YFilter }

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv6Unicast.YFilter = yf }

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetSegmentPath() string {
    return "l3vpn-ipv6-unicast"
}

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv6Unicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv6Unicast.PrefixLimit
    return children
}

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetYangName() string { return "l3vpn-ipv6-unicast" }

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) SetParent(parent types.Entity) { l3VpnIpv6Unicast.parent = parent }

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetParent() types.Entity { return l3VpnIpv6Unicast.parent }

func (l3VpnIpv6Unicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv6-unicast" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast
// Multicast IPv4 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetFilter() yfilter.YFilter { return l3VpnIpv4Multicast.YFilter }

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv4Multicast.YFilter = yf }

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetSegmentPath() string {
    return "l3vpn-ipv4-multicast"
}

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv4Multicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv4Multicast.PrefixLimit
    return children
}

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetYangName() string { return "l3vpn-ipv4-multicast" }

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) SetParent(parent types.Entity) { l3VpnIpv4Multicast.parent = parent }

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetParent() types.Entity { return l3VpnIpv4Multicast.parent }

func (l3VpnIpv4Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv4-multicast" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast
// Multicast IPv6 L3VPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetFilter() yfilter.YFilter { return l3VpnIpv6Multicast.YFilter }

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv6Multicast.YFilter = yf }

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetSegmentPath() string {
    return "l3vpn-ipv6-multicast"
}

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv6Multicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv6Multicast.PrefixLimit
    return children
}

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetYangName() string { return "l3vpn-ipv6-multicast" }

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) SetParent(parent types.Entity) { l3VpnIpv6Multicast.parent = parent }

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetParent() types.Entity { return l3VpnIpv6Multicast.parent }

func (l3VpnIpv6Multicast *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv6-multicast" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls
// BGP-signalled VPLS configuration options
type Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
}

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetFilter() yfilter.YFilter { return l2VpnVpls.YFilter }

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) SetFilter(yf yfilter.YFilter) { l2VpnVpls.YFilter = yf }

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetSegmentPath() string {
    return "l2vpn-vpls"
}

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l2VpnVpls.PrefixLimit
    }
    return nil
}

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l2VpnVpls.PrefixLimit
    return children
}

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetBundleName() string { return "openconfig" }

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetYangName() string { return "l2vpn-vpls" }

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) SetParent(parent types.Entity) { l2VpnVpls.parent = parent }

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetParent() types.Entity { return l2VpnVpls.parent }

func (l2VpnVpls *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetParentYangName() string { return "l2vpn-vpls" }

// Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn
// BGP EVPN configuration options
type Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
}

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetFilter() yfilter.YFilter { return l2VpnEvpn.YFilter }

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) SetFilter(yf yfilter.YFilter) { l2VpnEvpn.YFilter = yf }

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetSegmentPath() string {
    return "l2vpn-evpn"
}

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l2VpnEvpn.PrefixLimit
    }
    return nil
}

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l2VpnEvpn.PrefixLimit
    return children
}

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetBundleName() string { return "openconfig" }

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetYangName() string { return "l2vpn-evpn" }

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) SetParent(parent types.Entity) { l2VpnEvpn.parent = parent }

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetParent() types.Entity { return l2VpnEvpn.parent }

func (l2VpnEvpn *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn) GetParentYangName() string { return "afi-safi" }

// Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetParentYangName() string { return "l2vpn-evpn" }

// Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Global_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_Global_ApplyPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_Global_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_Global_ApplyPolicy_State
}

func (applyPolicy *Bgp_Global_ApplyPolicy) GetFilter() yfilter.YFilter { return applyPolicy.YFilter }

func (applyPolicy *Bgp_Global_ApplyPolicy) SetFilter(yf yfilter.YFilter) { applyPolicy.YFilter = yf }

func (applyPolicy *Bgp_Global_ApplyPolicy) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (applyPolicy *Bgp_Global_ApplyPolicy) GetSegmentPath() string {
    return "apply-policy"
}

func (applyPolicy *Bgp_Global_ApplyPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &applyPolicy.Config
    }
    if childYangName == "state" {
        return &applyPolicy.State
    }
    return nil
}

func (applyPolicy *Bgp_Global_ApplyPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &applyPolicy.Config
    children["state"] = &applyPolicy.State
    return children
}

func (applyPolicy *Bgp_Global_ApplyPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (applyPolicy *Bgp_Global_ApplyPolicy) GetBundleName() string { return "openconfig" }

func (applyPolicy *Bgp_Global_ApplyPolicy) GetYangName() string { return "apply-policy" }

func (applyPolicy *Bgp_Global_ApplyPolicy) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (applyPolicy *Bgp_Global_ApplyPolicy) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (applyPolicy *Bgp_Global_ApplyPolicy) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (applyPolicy *Bgp_Global_ApplyPolicy) SetParent(parent types.Entity) { applyPolicy.parent = parent }

func (applyPolicy *Bgp_Global_ApplyPolicy) GetParent() types.Entity { return applyPolicy.parent }

func (applyPolicy *Bgp_Global_ApplyPolicy) GetParentYangName() string { return "global" }

// Bgp_Global_ApplyPolicy_Config
// Policy configuration data.
type Bgp_Global_ApplyPolicy_Config struct {
    parent types.Entity
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

func (config *Bgp_Global_ApplyPolicy_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Global_ApplyPolicy_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Global_ApplyPolicy_Config) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (config *Bgp_Global_ApplyPolicy_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Global_ApplyPolicy_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Global_ApplyPolicy_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Global_ApplyPolicy_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = config.ImportPolicy
    leafs["default-import-policy"] = config.DefaultImportPolicy
    leafs["export-policy"] = config.ExportPolicy
    leafs["default-export-policy"] = config.DefaultExportPolicy
    return leafs
}

func (config *Bgp_Global_ApplyPolicy_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Global_ApplyPolicy_Config) GetYangName() string { return "config" }

func (config *Bgp_Global_ApplyPolicy_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Global_ApplyPolicy_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Global_ApplyPolicy_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Global_ApplyPolicy_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Global_ApplyPolicy_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Global_ApplyPolicy_Config) GetParentYangName() string { return "apply-policy" }

// Bgp_Global_ApplyPolicy_State
// Operational state for routing policy
type Bgp_Global_ApplyPolicy_State struct {
    parent types.Entity
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

func (state *Bgp_Global_ApplyPolicy_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Global_ApplyPolicy_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Global_ApplyPolicy_State) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (state *Bgp_Global_ApplyPolicy_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Global_ApplyPolicy_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Global_ApplyPolicy_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Global_ApplyPolicy_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = state.ImportPolicy
    leafs["default-import-policy"] = state.DefaultImportPolicy
    leafs["export-policy"] = state.ExportPolicy
    leafs["default-export-policy"] = state.DefaultExportPolicy
    return leafs
}

func (state *Bgp_Global_ApplyPolicy_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Global_ApplyPolicy_State) GetYangName() string { return "state" }

func (state *Bgp_Global_ApplyPolicy_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Global_ApplyPolicy_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Global_ApplyPolicy_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Global_ApplyPolicy_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Global_ApplyPolicy_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Global_ApplyPolicy_State) GetParentYangName() string { return "apply-policy" }

// Bgp_Neighbors
// Configuration for BGP neighbors
type Bgp_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of BGP neighbors configured on the local system, uniquely identified
    // by peer IPv[46] address. The type is slice of Bgp_Neighbors_Neighbor.
    Neighbor []Bgp_Neighbors_Neighbor
}

func (neighbors *Bgp_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Bgp_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Bgp_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Bgp_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Bgp_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Bgp_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Bgp_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Bgp_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Bgp_Neighbors) GetBundleName() string { return "openconfig" }

func (neighbors *Bgp_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Bgp_Neighbors) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbors *Bgp_Neighbors) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbors *Bgp_Neighbors) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbors *Bgp_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Bgp_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Bgp_Neighbors) GetParentYangName() string { return "bgp" }

// Bgp_Neighbors_Neighbor
// List of BGP neighbors configured on the local system,
// uniquely identified by peer IPv[46] address
type Bgp_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the address of the BGP neighbor used
    // as a key in the neighbor list. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (neighbor *Bgp_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Bgp_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Bgp_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "timers" { return "Timers" }
    if yname == "transport" { return "Transport" }
    if yname == "error-handling" { return "ErrorHandling" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "logging-options" { return "LoggingOptions" }
    if yname == "ebgp-multihop" { return "EbgpMultihop" }
    if yname == "route-reflector" { return "RouteReflector" }
    if yname == "as-path-options" { return "AsPathOptions" }
    if yname == "add-paths" { return "AddPaths" }
    if yname == "use-multiple-paths" { return "UseMultiplePaths" }
    if yname == "apply-policy" { return "ApplyPolicy" }
    if yname == "afi-safis" { return "AfiSafis" }
    return ""
}

func (neighbor *Bgp_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", neighbor.NeighborAddress) + "']"
}

func (neighbor *Bgp_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &neighbor.Config
    }
    if childYangName == "state" {
        return &neighbor.State
    }
    if childYangName == "timers" {
        return &neighbor.Timers
    }
    if childYangName == "transport" {
        return &neighbor.Transport
    }
    if childYangName == "error-handling" {
        return &neighbor.ErrorHandling
    }
    if childYangName == "graceful-restart" {
        return &neighbor.GracefulRestart
    }
    if childYangName == "logging-options" {
        return &neighbor.LoggingOptions
    }
    if childYangName == "ebgp-multihop" {
        return &neighbor.EbgpMultihop
    }
    if childYangName == "route-reflector" {
        return &neighbor.RouteReflector
    }
    if childYangName == "as-path-options" {
        return &neighbor.AsPathOptions
    }
    if childYangName == "add-paths" {
        return &neighbor.AddPaths
    }
    if childYangName == "use-multiple-paths" {
        return &neighbor.UseMultiplePaths
    }
    if childYangName == "apply-policy" {
        return &neighbor.ApplyPolicy
    }
    if childYangName == "afi-safis" {
        return &neighbor.AfiSafis
    }
    return nil
}

func (neighbor *Bgp_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &neighbor.Config
    children["state"] = &neighbor.State
    children["timers"] = &neighbor.Timers
    children["transport"] = &neighbor.Transport
    children["error-handling"] = &neighbor.ErrorHandling
    children["graceful-restart"] = &neighbor.GracefulRestart
    children["logging-options"] = &neighbor.LoggingOptions
    children["ebgp-multihop"] = &neighbor.EbgpMultihop
    children["route-reflector"] = &neighbor.RouteReflector
    children["as-path-options"] = &neighbor.AsPathOptions
    children["add-paths"] = &neighbor.AddPaths
    children["use-multiple-paths"] = &neighbor.UseMultiplePaths
    children["apply-policy"] = &neighbor.ApplyPolicy
    children["afi-safis"] = &neighbor.AfiSafis
    return children
}

func (neighbor *Bgp_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-address"] = neighbor.NeighborAddress
    return leafs
}

func (neighbor *Bgp_Neighbors_Neighbor) GetBundleName() string { return "openconfig" }

func (neighbor *Bgp_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Bgp_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbor *Bgp_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbor *Bgp_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbor *Bgp_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Bgp_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Bgp_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Bgp_Neighbors_Neighbor_Config
// Configuration parameters relating to the BGP neighbor or
// group
type Bgp_Neighbors_Neighbor_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The peer-group with which this neighbor is associated. The type is string.
    // Refers to bgp.Bgp_PeerGroups_PeerGroup_PeerGroupName
    PeerGroup interface{}

    // Address of the BGP peer, either in IPv4 or IPv6. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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
    // The type is one of the following: PRIVATEASREPLACEALLPRIVATEASREMOVEALL.
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

func (config *Bgp_Neighbors_Neighbor_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_Config) GetGoName(yname string) string {
    if yname == "peer-group" { return "PeerGroup" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "enabled" { return "Enabled" }
    if yname == "peer-as" { return "PeerAs" }
    if yname == "local-as" { return "LocalAs" }
    if yname == "peer-type" { return "PeerType" }
    if yname == "auth-password" { return "AuthPassword" }
    if yname == "remove-private-as" { return "RemovePrivateAs" }
    if yname == "route-flap-damping" { return "RouteFlapDamping" }
    if yname == "send-community" { return "SendCommunity" }
    if yname == "description" { return "Description" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-group"] = config.PeerGroup
    leafs["neighbor-address"] = config.NeighborAddress
    leafs["enabled"] = config.Enabled
    leafs["peer-as"] = config.PeerAs
    leafs["local-as"] = config.LocalAs
    leafs["peer-type"] = config.PeerType
    leafs["auth-password"] = config.AuthPassword
    leafs["remove-private-as"] = config.RemovePrivateAs
    leafs["route-flap-damping"] = config.RouteFlapDamping
    leafs["send-community"] = config.SendCommunity
    leafs["description"] = config.Description
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_Config) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_State
// State information relating to the BGP neighbor
type Bgp_Neighbors_Neighbor_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The peer-group with which this neighbor is associated. The type is string.
    // Refers to bgp.Bgp_PeerGroups_PeerGroup_PeerGroupName
    PeerGroup interface{}

    // Address of the BGP peer, either in IPv4 or IPv6. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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
    // The type is one of the following: PRIVATEASREPLACEALLPRIVATEASREMOVEALL.
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
    // of [u'GRACEFULRESTART', u'ROUTEREFRESH', u'MPBGP', u'ASN32', u'ADDPATHS'].
    SupportedCapabilities []interface{}

    // Counters for BGP messages sent and received from the neighbor.
    Messages Bgp_Neighbors_Neighbor_State_Messages

    // Counters related to queued messages associated with the BGP neighbor.
    Queues Bgp_Neighbors_Neighbor_State_Queues
}

func (state *Bgp_Neighbors_Neighbor_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_State) GetGoName(yname string) string {
    if yname == "peer-group" { return "PeerGroup" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "enabled" { return "Enabled" }
    if yname == "peer-as" { return "PeerAs" }
    if yname == "local-as" { return "LocalAs" }
    if yname == "peer-type" { return "PeerType" }
    if yname == "auth-password" { return "AuthPassword" }
    if yname == "remove-private-as" { return "RemovePrivateAs" }
    if yname == "route-flap-damping" { return "RouteFlapDamping" }
    if yname == "send-community" { return "SendCommunity" }
    if yname == "description" { return "Description" }
    if yname == "session-state" { return "SessionState" }
    if yname == "last-established" { return "LastEstablished" }
    if yname == "established-transitions" { return "EstablishedTransitions" }
    if yname == "supported-capabilities" { return "SupportedCapabilities" }
    if yname == "messages" { return "Messages" }
    if yname == "queues" { return "Queues" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "messages" {
        return &state.Messages
    }
    if childYangName == "queues" {
        return &state.Queues
    }
    return nil
}

func (state *Bgp_Neighbors_Neighbor_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["messages"] = &state.Messages
    children["queues"] = &state.Queues
    return children
}

func (state *Bgp_Neighbors_Neighbor_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-group"] = state.PeerGroup
    leafs["neighbor-address"] = state.NeighborAddress
    leafs["enabled"] = state.Enabled
    leafs["peer-as"] = state.PeerAs
    leafs["local-as"] = state.LocalAs
    leafs["peer-type"] = state.PeerType
    leafs["auth-password"] = state.AuthPassword
    leafs["remove-private-as"] = state.RemovePrivateAs
    leafs["route-flap-damping"] = state.RouteFlapDamping
    leafs["send-community"] = state.SendCommunity
    leafs["description"] = state.Description
    leafs["session-state"] = state.SessionState
    leafs["last-established"] = state.LastEstablished
    leafs["established-transitions"] = state.EstablishedTransitions
    leafs["supported-capabilities"] = state.SupportedCapabilities
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_State) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_State_Messages
// Counters for BGP messages sent and received from the
// neighbor
type Bgp_Neighbors_Neighbor_State_Messages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Counters relating to BGP messages sent to the neighbor.
    Sent Bgp_Neighbors_Neighbor_State_Messages_Sent

    // Counters for BGP messages received from the neighbor.
    Received Bgp_Neighbors_Neighbor_State_Messages_Received
}

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetFilter() yfilter.YFilter { return messages.YFilter }

func (messages *Bgp_Neighbors_Neighbor_State_Messages) SetFilter(yf yfilter.YFilter) { messages.YFilter = yf }

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetGoName(yname string) string {
    if yname == "sent" { return "Sent" }
    if yname == "received" { return "Received" }
    return ""
}

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetSegmentPath() string {
    return "messages"
}

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sent" {
        return &messages.Sent
    }
    if childYangName == "received" {
        return &messages.Received
    }
    return nil
}

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sent"] = &messages.Sent
    children["received"] = &messages.Received
    return children
}

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetBundleName() string { return "openconfig" }

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetYangName() string { return "messages" }

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (messages *Bgp_Neighbors_Neighbor_State_Messages) SetParent(parent types.Entity) { messages.parent = parent }

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetParent() types.Entity { return messages.parent }

func (messages *Bgp_Neighbors_Neighbor_State_Messages) GetParentYangName() string { return "state" }

// Bgp_Neighbors_Neighbor_State_Messages_Sent
// Counters relating to BGP messages sent to the neighbor
type Bgp_Neighbors_Neighbor_State_Messages_Sent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of BGP UPDATE messages announcing, withdrawing or modifying paths
    // exchanged. The type is interface{} with range: 0..18446744073709551615.
    Update interface{}

    // Number of BGP NOTIFICATION messages indicating an error condition has
    // occurred exchanged. The type is interface{} with range:
    // 0..18446744073709551615.
    Notification interface{}
}

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetFilter() yfilter.YFilter { return sent.YFilter }

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) SetFilter(yf yfilter.YFilter) { sent.YFilter = yf }

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetGoName(yname string) string {
    if yname == "UPDATE" { return "Update" }
    if yname == "NOTIFICATION" { return "Notification" }
    return ""
}

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetSegmentPath() string {
    return "sent"
}

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["UPDATE"] = sent.Update
    leafs["NOTIFICATION"] = sent.Notification
    return leafs
}

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetBundleName() string { return "openconfig" }

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetYangName() string { return "sent" }

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) SetParent(parent types.Entity) { sent.parent = parent }

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetParent() types.Entity { return sent.parent }

func (sent *Bgp_Neighbors_Neighbor_State_Messages_Sent) GetParentYangName() string { return "messages" }

// Bgp_Neighbors_Neighbor_State_Messages_Received
// Counters for BGP messages received from the neighbor
type Bgp_Neighbors_Neighbor_State_Messages_Received struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of BGP UPDATE messages announcing, withdrawing or modifying paths
    // exchanged. The type is interface{} with range: 0..18446744073709551615.
    Update interface{}

    // Number of BGP NOTIFICATION messages indicating an error condition has
    // occurred exchanged. The type is interface{} with range:
    // 0..18446744073709551615.
    Notification interface{}
}

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetFilter() yfilter.YFilter { return received.YFilter }

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) SetFilter(yf yfilter.YFilter) { received.YFilter = yf }

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetGoName(yname string) string {
    if yname == "UPDATE" { return "Update" }
    if yname == "NOTIFICATION" { return "Notification" }
    return ""
}

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetSegmentPath() string {
    return "received"
}

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["UPDATE"] = received.Update
    leafs["NOTIFICATION"] = received.Notification
    return leafs
}

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetBundleName() string { return "openconfig" }

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetYangName() string { return "received" }

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) SetParent(parent types.Entity) { received.parent = parent }

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetParent() types.Entity { return received.parent }

func (received *Bgp_Neighbors_Neighbor_State_Messages_Received) GetParentYangName() string { return "messages" }

// Bgp_Neighbors_Neighbor_State_Queues
// Counters related to queued messages associated with the
// BGP neighbor
type Bgp_Neighbors_Neighbor_State_Queues struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of messages received from the peer currently queued. The type is
    // interface{} with range: 0..4294967295.
    Input interface{}

    // The number of messages queued to be sent to the peer. The type is
    // interface{} with range: 0..4294967295.
    Output interface{}
}

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetFilter() yfilter.YFilter { return queues.YFilter }

func (queues *Bgp_Neighbors_Neighbor_State_Queues) SetFilter(yf yfilter.YFilter) { queues.YFilter = yf }

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetSegmentPath() string {
    return "queues"
}

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input"] = queues.Input
    leafs["output"] = queues.Output
    return leafs
}

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetBundleName() string { return "openconfig" }

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetYangName() string { return "queues" }

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (queues *Bgp_Neighbors_Neighbor_State_Queues) SetParent(parent types.Entity) { queues.parent = parent }

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetParent() types.Entity { return queues.parent }

func (queues *Bgp_Neighbors_Neighbor_State_Queues) GetParentYangName() string { return "state" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to timers used for the BGP neighbor.
    Config Bgp_Neighbors_Neighbor_Timers_Config

    // State information relating to the timers used for the BGP neighbor.
    State Bgp_Neighbors_Neighbor_Timers_State
}

func (timers *Bgp_Neighbors_Neighbor_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Bgp_Neighbors_Neighbor_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Bgp_Neighbors_Neighbor_Timers) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (timers *Bgp_Neighbors_Neighbor_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Bgp_Neighbors_Neighbor_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &timers.Config
    }
    if childYangName == "state" {
        return &timers.State
    }
    return nil
}

func (timers *Bgp_Neighbors_Neighbor_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &timers.Config
    children["state"] = &timers.State
    return children
}

func (timers *Bgp_Neighbors_Neighbor_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (timers *Bgp_Neighbors_Neighbor_Timers) GetBundleName() string { return "openconfig" }

func (timers *Bgp_Neighbors_Neighbor_Timers) GetYangName() string { return "timers" }

func (timers *Bgp_Neighbors_Neighbor_Timers) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (timers *Bgp_Neighbors_Neighbor_Timers) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (timers *Bgp_Neighbors_Neighbor_Timers) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (timers *Bgp_Neighbors_Neighbor_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Bgp_Neighbors_Neighbor_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Bgp_Neighbors_Neighbor_Timers) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_Timers_Config
// Configuration parameters relating to timers used for the
// BGP neighbor
type Bgp_Neighbors_Neighbor_Timers_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_Timers_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetGoName(yname string) string {
    if yname == "connect-retry" { return "ConnectRetry" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "keepalive-interval" { return "KeepaliveInterval" }
    if yname == "minimum-advertisement-interval" { return "MinimumAdvertisementInterval" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["connect-retry"] = config.ConnectRetry
    leafs["hold-time"] = config.HoldTime
    leafs["keepalive-interval"] = config.KeepaliveInterval
    leafs["minimum-advertisement-interval"] = config.MinimumAdvertisementInterval
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_Timers_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_Timers_Config) GetParentYangName() string { return "timers" }

// Bgp_Neighbors_Neighbor_Timers_State
// State information relating to the timers used for the BGP
// neighbor
type Bgp_Neighbors_Neighbor_Timers_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_Timers_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetGoName(yname string) string {
    if yname == "connect-retry" { return "ConnectRetry" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "keepalive-interval" { return "KeepaliveInterval" }
    if yname == "minimum-advertisement-interval" { return "MinimumAdvertisementInterval" }
    if yname == "negotiated-hold-time" { return "NegotiatedHoldTime" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["connect-retry"] = state.ConnectRetry
    leafs["hold-time"] = state.HoldTime
    leafs["keepalive-interval"] = state.KeepaliveInterval
    leafs["minimum-advertisement-interval"] = state.MinimumAdvertisementInterval
    leafs["negotiated-hold-time"] = state.NegotiatedHoldTime
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_Timers_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_Timers_State) GetParentYangName() string { return "timers" }

// Bgp_Neighbors_Neighbor_Transport
// Transport session parameters for the BGP neighbor
type Bgp_Neighbors_Neighbor_Transport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the transport session(s) used for the
    // BGP neighbor.
    Config Bgp_Neighbors_Neighbor_Transport_Config

    // State information relating to the transport session(s) used for the BGP
    // neighbor.
    State Bgp_Neighbors_Neighbor_Transport_State
}

func (transport *Bgp_Neighbors_Neighbor_Transport) GetFilter() yfilter.YFilter { return transport.YFilter }

func (transport *Bgp_Neighbors_Neighbor_Transport) SetFilter(yf yfilter.YFilter) { transport.YFilter = yf }

func (transport *Bgp_Neighbors_Neighbor_Transport) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (transport *Bgp_Neighbors_Neighbor_Transport) GetSegmentPath() string {
    return "transport"
}

func (transport *Bgp_Neighbors_Neighbor_Transport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &transport.Config
    }
    if childYangName == "state" {
        return &transport.State
    }
    return nil
}

func (transport *Bgp_Neighbors_Neighbor_Transport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &transport.Config
    children["state"] = &transport.State
    return children
}

func (transport *Bgp_Neighbors_Neighbor_Transport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (transport *Bgp_Neighbors_Neighbor_Transport) GetBundleName() string { return "openconfig" }

func (transport *Bgp_Neighbors_Neighbor_Transport) GetYangName() string { return "transport" }

func (transport *Bgp_Neighbors_Neighbor_Transport) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (transport *Bgp_Neighbors_Neighbor_Transport) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (transport *Bgp_Neighbors_Neighbor_Transport) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (transport *Bgp_Neighbors_Neighbor_Transport) SetParent(parent types.Entity) { transport.parent = parent }

func (transport *Bgp_Neighbors_Neighbor_Transport) GetParent() types.Entity { return transport.parent }

func (transport *Bgp_Neighbors_Neighbor_Transport) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_Transport_Config
// Configuration parameters relating to the transport
// session(s) used for the BGP neighbor
type Bgp_Neighbors_Neighbor_Transport_Config struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
    // or string.
    LocalAddress interface{}
}

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_Transport_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetGoName(yname string) string {
    if yname == "tcp-mss" { return "TcpMss" }
    if yname == "mtu-discovery" { return "MtuDiscovery" }
    if yname == "passive-mode" { return "PassiveMode" }
    if yname == "local-address" { return "LocalAddress" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-mss"] = config.TcpMss
    leafs["mtu-discovery"] = config.MtuDiscovery
    leafs["passive-mode"] = config.PassiveMode
    leafs["local-address"] = config.LocalAddress
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_Transport_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_Transport_Config) GetParentYangName() string { return "transport" }

// Bgp_Neighbors_Neighbor_Transport_State
// State information relating to the transport session(s)
// used for the BGP neighbor
type Bgp_Neighbors_Neighbor_Transport_State struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
    // or string.
    LocalAddress interface{}

    // Local TCP port being used for the TCP session supporting the BGP session.
    // The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Remote address to which the BGP session has been established. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteAddress interface{}

    // Remote port being used by the peer for the TCP session supporting the BGP
    // session. The type is interface{} with range: 0..65535.
    RemotePort interface{}
}

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_Transport_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetGoName(yname string) string {
    if yname == "tcp-mss" { return "TcpMss" }
    if yname == "mtu-discovery" { return "MtuDiscovery" }
    if yname == "passive-mode" { return "PassiveMode" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "remote-address" { return "RemoteAddress" }
    if yname == "remote-port" { return "RemotePort" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-mss"] = state.TcpMss
    leafs["mtu-discovery"] = state.MtuDiscovery
    leafs["passive-mode"] = state.PassiveMode
    leafs["local-address"] = state.LocalAddress
    leafs["local-port"] = state.LocalPort
    leafs["remote-address"] = state.RemoteAddress
    leafs["remote-port"] = state.RemotePort
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_Transport_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_Transport_State) GetParentYangName() string { return "transport" }

// Bgp_Neighbors_Neighbor_ErrorHandling
// Error handling parameters used for the BGP neighbor or
// group
type Bgp_Neighbors_Neighbor_ErrorHandling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters enabling or modifying the behavior or enhanced
    // error handling mechanisms for the BGP neighbor.
    Config Bgp_Neighbors_Neighbor_ErrorHandling_Config

    // State information relating to enhanced error handling mechanisms for the
    // BGP neighbor.
    State Bgp_Neighbors_Neighbor_ErrorHandling_State
}

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetFilter() yfilter.YFilter { return errorHandling.YFilter }

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) SetFilter(yf yfilter.YFilter) { errorHandling.YFilter = yf }

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetSegmentPath() string {
    return "error-handling"
}

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &errorHandling.Config
    }
    if childYangName == "state" {
        return &errorHandling.State
    }
    return nil
}

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &errorHandling.Config
    children["state"] = &errorHandling.State
    return children
}

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetBundleName() string { return "openconfig" }

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetYangName() string { return "error-handling" }

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) SetParent(parent types.Entity) { errorHandling.parent = parent }

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetParent() types.Entity { return errorHandling.parent }

func (errorHandling *Bgp_Neighbors_Neighbor_ErrorHandling) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_ErrorHandling_Config
// Configuration parameters enabling or modifying the
// behavior or enhanced error handling mechanisms for the BGP
// neighbor
type Bgp_Neighbors_Neighbor_ErrorHandling_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify whether erroneous UPDATE messages for which the NLRI can be
    // extracted are reated as though the NLRI is withdrawn - avoiding session
    // reset. The type is bool. The default value is false.
    TreatAsWithdraw interface{}
}

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetGoName(yname string) string {
    if yname == "treat-as-withdraw" { return "TreatAsWithdraw" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["treat-as-withdraw"] = config.TreatAsWithdraw
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_ErrorHandling_Config) GetParentYangName() string { return "error-handling" }

// Bgp_Neighbors_Neighbor_ErrorHandling_State
// State information relating to enhanced error handling
// mechanisms for the BGP neighbor
type Bgp_Neighbors_Neighbor_ErrorHandling_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetGoName(yname string) string {
    if yname == "treat-as-withdraw" { return "TreatAsWithdraw" }
    if yname == "erroneous-update-messages" { return "ErroneousUpdateMessages" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["treat-as-withdraw"] = state.TreatAsWithdraw
    leafs["erroneous-update-messages"] = state.ErroneousUpdateMessages
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_ErrorHandling_State) GetParentYangName() string { return "error-handling" }

// Bgp_Neighbors_Neighbor_GracefulRestart
// Parameters relating the graceful restart mechanism for BGP
type Bgp_Neighbors_Neighbor_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to graceful-restart.
    Config Bgp_Neighbors_Neighbor_GracefulRestart_Config

    // State information associated with graceful-restart.
    State Bgp_Neighbors_Neighbor_GracefulRestart_State
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &gracefulRestart.Config
    }
    if childYangName == "state" {
        return &gracefulRestart.State
    }
    return nil
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &gracefulRestart.Config
    children["state"] = &gracefulRestart.State
    return children
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetBundleName() string { return "openconfig" }

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *Bgp_Neighbors_Neighbor_GracefulRestart) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_GracefulRestart_Config
// Configuration parameters relating to graceful-restart
type Bgp_Neighbors_Neighbor_GracefulRestart_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "restart-time" { return "RestartTime" }
    if yname == "stale-routes-time" { return "StaleRoutesTime" }
    if yname == "helper-only" { return "HelperOnly" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["restart-time"] = config.RestartTime
    leafs["stale-routes-time"] = config.StaleRoutesTime
    leafs["helper-only"] = config.HelperOnly
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_GracefulRestart_Config) GetParentYangName() string { return "graceful-restart" }

// Bgp_Neighbors_Neighbor_GracefulRestart_State
// State information associated with graceful-restart
type Bgp_Neighbors_Neighbor_GracefulRestart_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "restart-time" { return "RestartTime" }
    if yname == "stale-routes-time" { return "StaleRoutesTime" }
    if yname == "helper-only" { return "HelperOnly" }
    if yname == "peer-restart-time" { return "PeerRestartTime" }
    if yname == "peer-restarting" { return "PeerRestarting" }
    if yname == "local-restarting" { return "LocalRestarting" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["restart-time"] = state.RestartTime
    leafs["stale-routes-time"] = state.StaleRoutesTime
    leafs["helper-only"] = state.HelperOnly
    leafs["peer-restart-time"] = state.PeerRestartTime
    leafs["peer-restarting"] = state.PeerRestarting
    leafs["local-restarting"] = state.LocalRestarting
    leafs["mode"] = state.Mode
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_GracefulRestart_State) GetParentYangName() string { return "graceful-restart" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters enabling or modifying logging for events relating
    // to the BGPgroup.
    Config Bgp_Neighbors_Neighbor_LoggingOptions_Config

    // State information relating to logging for the BGP neighbor or group.
    State Bgp_Neighbors_Neighbor_LoggingOptions_State
}

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetFilter() yfilter.YFilter { return loggingOptions.YFilter }

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) SetFilter(yf yfilter.YFilter) { loggingOptions.YFilter = yf }

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetSegmentPath() string {
    return "logging-options"
}

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &loggingOptions.Config
    }
    if childYangName == "state" {
        return &loggingOptions.State
    }
    return nil
}

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &loggingOptions.Config
    children["state"] = &loggingOptions.State
    return children
}

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetBundleName() string { return "openconfig" }

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetYangName() string { return "logging-options" }

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) SetParent(parent types.Entity) { loggingOptions.parent = parent }

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetParent() types.Entity { return loggingOptions.parent }

func (loggingOptions *Bgp_Neighbors_Neighbor_LoggingOptions) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_LoggingOptions_Config
// Configuration parameters enabling or modifying logging
// for events relating to the BGPgroup
type Bgp_Neighbors_Neighbor_LoggingOptions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure logging of peer state changes.  Default is to enable logging of
    // peer state changes. The type is bool. The default value is true.
    LogNeighborStateChanges interface{}
}

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetGoName(yname string) string {
    if yname == "log-neighbor-state-changes" { return "LogNeighborStateChanges" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-neighbor-state-changes"] = config.LogNeighborStateChanges
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_LoggingOptions_Config) GetParentYangName() string { return "logging-options" }

// Bgp_Neighbors_Neighbor_LoggingOptions_State
// State information relating to logging for the BGP neighbor
// or group
type Bgp_Neighbors_Neighbor_LoggingOptions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure logging of peer state changes.  Default is to enable logging of
    // peer state changes. The type is bool. The default value is true.
    LogNeighborStateChanges interface{}
}

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetGoName(yname string) string {
    if yname == "log-neighbor-state-changes" { return "LogNeighborStateChanges" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-neighbor-state-changes"] = state.LogNeighborStateChanges
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_LoggingOptions_State) GetParentYangName() string { return "logging-options" }

// Bgp_Neighbors_Neighbor_EbgpMultihop
// eBGP multi-hop parameters for the BGPgroup
type Bgp_Neighbors_Neighbor_EbgpMultihop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multihop for the BGP group.
    Config Bgp_Neighbors_Neighbor_EbgpMultihop_Config

    // State information for eBGP multihop, for the BGP neighbor or group.
    State Bgp_Neighbors_Neighbor_EbgpMultihop_State
}

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetFilter() yfilter.YFilter { return ebgpMultihop.YFilter }

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) SetFilter(yf yfilter.YFilter) { ebgpMultihop.YFilter = yf }

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetSegmentPath() string {
    return "ebgp-multihop"
}

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ebgpMultihop.Config
    }
    if childYangName == "state" {
        return &ebgpMultihop.State
    }
    return nil
}

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ebgpMultihop.Config
    children["state"] = &ebgpMultihop.State
    return children
}

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetBundleName() string { return "openconfig" }

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetYangName() string { return "ebgp-multihop" }

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) SetParent(parent types.Entity) { ebgpMultihop.parent = parent }

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetParent() types.Entity { return ebgpMultihop.parent }

func (ebgpMultihop *Bgp_Neighbors_Neighbor_EbgpMultihop) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_EbgpMultihop_Config
// Configuration parameters relating to eBGP multihop for the
// BGP group
type Bgp_Neighbors_Neighbor_EbgpMultihop_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "multihop-ttl" { return "MultihopTtl" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["multihop-ttl"] = config.MultihopTtl
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_EbgpMultihop_Config) GetParentYangName() string { return "ebgp-multihop" }

// Bgp_Neighbors_Neighbor_EbgpMultihop_State
// State information for eBGP multihop, for the BGP neighbor
// or group
type Bgp_Neighbors_Neighbor_EbgpMultihop_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "multihop-ttl" { return "MultihopTtl" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["multihop-ttl"] = state.MultihopTtl
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_EbgpMultihop_State) GetParentYangName() string { return "ebgp-multihop" }

// Bgp_Neighbors_Neighbor_RouteReflector
// Route reflector parameters for the BGPgroup
type Bgp_Neighbors_Neighbor_RouteReflector struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuraton parameters relating to route reflection for the BGPgroup.
    Config Bgp_Neighbors_Neighbor_RouteReflector_Config

    // State information relating to route reflection for the BGPgroup.
    State Bgp_Neighbors_Neighbor_RouteReflector_State
}

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetFilter() yfilter.YFilter { return routeReflector.YFilter }

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) SetFilter(yf yfilter.YFilter) { routeReflector.YFilter = yf }

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetSegmentPath() string {
    return "route-reflector"
}

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &routeReflector.Config
    }
    if childYangName == "state" {
        return &routeReflector.State
    }
    return nil
}

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &routeReflector.Config
    children["state"] = &routeReflector.State
    return children
}

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetBundleName() string { return "openconfig" }

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetYangName() string { return "route-reflector" }

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) SetParent(parent types.Entity) { routeReflector.parent = parent }

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetParent() types.Entity { return routeReflector.parent }

func (routeReflector *Bgp_Neighbors_Neighbor_RouteReflector) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_RouteReflector_Config
// Configuraton parameters relating to route reflection
// for the BGPgroup
type Bgp_Neighbors_Neighbor_RouteReflector_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route-reflector cluster id to use when local router is configured as a
    // route reflector.  Commonly set at the group level, but allows a different
    // cluster id to be set for each neighbor. The type is one of the following
    // types: int with range: 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouteReflectorClusterId interface{}

    // Configure the neighbor as a route reflector client. The type is bool. The
    // default value is false.
    RouteReflectorClient interface{}
}

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetGoName(yname string) string {
    if yname == "route-reflector-cluster-id" { return "RouteReflectorClusterId" }
    if yname == "route-reflector-client" { return "RouteReflectorClient" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-reflector-cluster-id"] = config.RouteReflectorClusterId
    leafs["route-reflector-client"] = config.RouteReflectorClient
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_RouteReflector_Config) GetParentYangName() string { return "route-reflector" }

// Bgp_Neighbors_Neighbor_RouteReflector_State
// State information relating to route reflection for the
// BGPgroup
type Bgp_Neighbors_Neighbor_RouteReflector_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route-reflector cluster id to use when local router is configured as a
    // route reflector.  Commonly set at the group level, but allows a different
    // cluster id to be set for each neighbor. The type is one of the following
    // types: int with range: 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouteReflectorClusterId interface{}

    // Configure the neighbor as a route reflector client. The type is bool. The
    // default value is false.
    RouteReflectorClient interface{}
}

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetGoName(yname string) string {
    if yname == "route-reflector-cluster-id" { return "RouteReflectorClusterId" }
    if yname == "route-reflector-client" { return "RouteReflectorClient" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-reflector-cluster-id"] = state.RouteReflectorClusterId
    leafs["route-reflector-client"] = state.RouteReflectorClient
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_RouteReflector_State) GetParentYangName() string { return "route-reflector" }

// Bgp_Neighbors_Neighbor_AsPathOptions
// AS_PATH manipulation parameters for the BGP neighbor or
// group
type Bgp_Neighbors_Neighbor_AsPathOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to AS_PATH manipulation for the BGP peer
    // or group.
    Config Bgp_Neighbors_Neighbor_AsPathOptions_Config

    // State information relating to the AS_PATH manipulation mechanisms for the
    // BGP peer or group.
    State Bgp_Neighbors_Neighbor_AsPathOptions_State
}

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetFilter() yfilter.YFilter { return asPathOptions.YFilter }

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) SetFilter(yf yfilter.YFilter) { asPathOptions.YFilter = yf }

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetSegmentPath() string {
    return "as-path-options"
}

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &asPathOptions.Config
    }
    if childYangName == "state" {
        return &asPathOptions.State
    }
    return nil
}

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &asPathOptions.Config
    children["state"] = &asPathOptions.State
    return children
}

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetBundleName() string { return "openconfig" }

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetYangName() string { return "as-path-options" }

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) SetParent(parent types.Entity) { asPathOptions.parent = parent }

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetParent() types.Entity { return asPathOptions.parent }

func (asPathOptions *Bgp_Neighbors_Neighbor_AsPathOptions) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_AsPathOptions_Config
// Configuration parameters relating to AS_PATH manipulation
// for the BGP peer or group
type Bgp_Neighbors_Neighbor_AsPathOptions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the number of occurrences of the local BGP speaker's AS that can
    // occur within the AS_PATH before it is rejected. The type is interface{}
    // with range: 0..255. The default value is 0.
    AllowOwnAs interface{}

    // Replace occurrences of the peer's AS in the AS_PATH with the local
    // autonomous system number. The type is bool. The default value is false.
    ReplacePeerAs interface{}
}

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetGoName(yname string) string {
    if yname == "allow-own-as" { return "AllowOwnAs" }
    if yname == "replace-peer-as" { return "ReplacePeerAs" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-own-as"] = config.AllowOwnAs
    leafs["replace-peer-as"] = config.ReplacePeerAs
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AsPathOptions_Config) GetParentYangName() string { return "as-path-options" }

// Bgp_Neighbors_Neighbor_AsPathOptions_State
// State information relating to the AS_PATH manipulation
// mechanisms for the BGP peer or group
type Bgp_Neighbors_Neighbor_AsPathOptions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the number of occurrences of the local BGP speaker's AS that can
    // occur within the AS_PATH before it is rejected. The type is interface{}
    // with range: 0..255. The default value is 0.
    AllowOwnAs interface{}

    // Replace occurrences of the peer's AS in the AS_PATH with the local
    // autonomous system number. The type is bool. The default value is false.
    ReplacePeerAs interface{}
}

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetGoName(yname string) string {
    if yname == "allow-own-as" { return "AllowOwnAs" }
    if yname == "replace-peer-as" { return "ReplacePeerAs" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-own-as"] = state.AllowOwnAs
    leafs["replace-peer-as"] = state.ReplacePeerAs
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AsPathOptions_State) GetParentYangName() string { return "as-path-options" }

// Bgp_Neighbors_Neighbor_AddPaths
// Parameters relating to the advertisement and receipt of
// multiple paths for a single NLRI (add-paths)
type Bgp_Neighbors_Neighbor_AddPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to ADD_PATHS.
    Config Bgp_Neighbors_Neighbor_AddPaths_Config

    // State information associated with ADD_PATHS.
    State Bgp_Neighbors_Neighbor_AddPaths_State
}

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetFilter() yfilter.YFilter { return addPaths.YFilter }

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) SetFilter(yf yfilter.YFilter) { addPaths.YFilter = yf }

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetSegmentPath() string {
    return "add-paths"
}

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &addPaths.Config
    }
    if childYangName == "state" {
        return &addPaths.State
    }
    return nil
}

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &addPaths.Config
    children["state"] = &addPaths.State
    return children
}

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetBundleName() string { return "openconfig" }

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetYangName() string { return "add-paths" }

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) SetParent(parent types.Entity) { addPaths.parent = parent }

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetParent() types.Entity { return addPaths.parent }

func (addPaths *Bgp_Neighbors_Neighbor_AddPaths) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_AddPaths_Config
// Configuration parameters relating to ADD_PATHS
type Bgp_Neighbors_Neighbor_AddPaths_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetGoName(yname string) string {
    if yname == "receive" { return "Receive" }
    if yname == "send-max" { return "SendMax" }
    if yname == "eligible-prefix-policy" { return "EligiblePrefixPolicy" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receive"] = config.Receive
    leafs["send-max"] = config.SendMax
    leafs["eligible-prefix-policy"] = config.EligiblePrefixPolicy
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AddPaths_Config) GetParentYangName() string { return "add-paths" }

// Bgp_Neighbors_Neighbor_AddPaths_State
// State information associated with ADD_PATHS
type Bgp_Neighbors_Neighbor_AddPaths_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetGoName(yname string) string {
    if yname == "receive" { return "Receive" }
    if yname == "send-max" { return "SendMax" }
    if yname == "eligible-prefix-policy" { return "EligiblePrefixPolicy" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receive"] = state.Receive
    leafs["send-max"] = state.SendMax
    leafs["eligible-prefix-policy"] = state.EligiblePrefixPolicy
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AddPaths_State) GetParentYangName() string { return "add-paths" }

// Bgp_Neighbors_Neighbor_UseMultiplePaths
// Parameters related to the use of multiple-paths for the same
// NLRI when they are received only from this neighbor
type Bgp_Neighbors_Neighbor_UseMultiplePaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to multipath.
    Config Bgp_Neighbors_Neighbor_UseMultiplePaths_Config

    // State parameters relating to multipath.
    State Bgp_Neighbors_Neighbor_UseMultiplePaths_State

    // Multipath configuration for eBGP.
    Ebgp Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetFilter() yfilter.YFilter { return useMultiplePaths.YFilter }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) SetFilter(yf yfilter.YFilter) { useMultiplePaths.YFilter = yf }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "ebgp" { return "Ebgp" }
    return ""
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetSegmentPath() string {
    return "use-multiple-paths"
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &useMultiplePaths.Config
    }
    if childYangName == "state" {
        return &useMultiplePaths.State
    }
    if childYangName == "ebgp" {
        return &useMultiplePaths.Ebgp
    }
    return nil
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &useMultiplePaths.Config
    children["state"] = &useMultiplePaths.State
    children["ebgp"] = &useMultiplePaths.Ebgp
    return children
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetBundleName() string { return "openconfig" }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetYangName() string { return "use-multiple-paths" }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) SetParent(parent types.Entity) { useMultiplePaths.parent = parent }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetParent() types.Entity { return useMultiplePaths.parent }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_UseMultiplePaths) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_Neighbors_Neighbor_UseMultiplePaths_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Config) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Neighbors_Neighbor_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_Neighbors_Neighbor_UseMultiplePaths_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_State) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp
// Multipath configuration for eBGP
type Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetFilter() yfilter.YFilter { return ebgp.YFilter }

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) SetFilter(yf yfilter.YFilter) { ebgp.YFilter = yf }

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetSegmentPath() string {
    return "ebgp"
}

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ebgp.Config
    }
    if childYangName == "state" {
        return &ebgp.State
    }
    return nil
}

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ebgp.Config
    children["state"] = &ebgp.State
    return children
}

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetBundleName() string { return "openconfig" }

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetYangName() string { return "ebgp" }

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) SetParent(parent types.Entity) { ebgp.parent = parent }

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetParent() types.Entity { return ebgp.parent }

func (ebgp *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = config.AllowMultipleAs
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_Config) GetParentYangName() string { return "ebgp" }

// Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = state.AllowMultipleAs
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_UseMultiplePaths_Ebgp_State) GetParentYangName() string { return "ebgp" }

// Bgp_Neighbors_Neighbor_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_Neighbors_Neighbor_ApplyPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_Neighbors_Neighbor_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_Neighbors_Neighbor_ApplyPolicy_State
}

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetFilter() yfilter.YFilter { return applyPolicy.YFilter }

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) SetFilter(yf yfilter.YFilter) { applyPolicy.YFilter = yf }

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetSegmentPath() string {
    return "apply-policy"
}

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &applyPolicy.Config
    }
    if childYangName == "state" {
        return &applyPolicy.State
    }
    return nil
}

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &applyPolicy.Config
    children["state"] = &applyPolicy.State
    return children
}

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetBundleName() string { return "openconfig" }

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetYangName() string { return "apply-policy" }

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) SetParent(parent types.Entity) { applyPolicy.parent = parent }

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetParent() types.Entity { return applyPolicy.parent }

func (applyPolicy *Bgp_Neighbors_Neighbor_ApplyPolicy) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_ApplyPolicy_Config
// Policy configuration data.
type Bgp_Neighbors_Neighbor_ApplyPolicy_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = config.ImportPolicy
    leafs["default-import-policy"] = config.DefaultImportPolicy
    leafs["export-policy"] = config.ExportPolicy
    leafs["default-export-policy"] = config.DefaultExportPolicy
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_ApplyPolicy_Config) GetParentYangName() string { return "apply-policy" }

// Bgp_Neighbors_Neighbor_ApplyPolicy_State
// Operational state for routing policy
type Bgp_Neighbors_Neighbor_ApplyPolicy_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = state.ImportPolicy
    leafs["default-import-policy"] = state.DefaultImportPolicy
    leafs["export-policy"] = state.ExportPolicy
    leafs["default-export-policy"] = state.DefaultExportPolicy
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_ApplyPolicy_State) GetParentYangName() string { return "apply-policy" }

// Bgp_Neighbors_Neighbor_AfiSafis
// Per-address-family configuration parameters associated with
// the neighbor
type Bgp_Neighbors_Neighbor_AfiSafis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI,SAFI configuration available for the neighbour or group. The type is
    // slice of Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi.
    AfiSafi []Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi
}

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetFilter() yfilter.YFilter { return afiSafis.YFilter }

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) SetFilter(yf yfilter.YFilter) { afiSafis.YFilter = yf }

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetGoName(yname string) string {
    if yname == "afi-safi" { return "AfiSafi" }
    return ""
}

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetSegmentPath() string {
    return "afi-safis"
}

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afi-safi" {
        for _, c := range afiSafis.AfiSafi {
            if afiSafis.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi{}
        afiSafis.AfiSafi = append(afiSafis.AfiSafi, child)
        return &afiSafis.AfiSafi[len(afiSafis.AfiSafi)-1]
    }
    return nil
}

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afiSafis.AfiSafi {
        children[afiSafis.AfiSafi[i].GetSegmentPath()] = &afiSafis.AfiSafi[i]
    }
    return children
}

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetBundleName() string { return "openconfig" }

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetYangName() string { return "afi-safis" }

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) SetParent(parent types.Entity) { afiSafis.parent = parent }

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetParent() types.Entity { return afiSafis.parent }

func (afiSafis *Bgp_Neighbors_Neighbor_AfiSafis) GetParentYangName() string { return "neighbor" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi
// AFI,SAFI configuration available for the
// neighbour or group
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the AFI-SAFI name used as a key for
    // the AFI-SAFI list. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
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

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetFilter() yfilter.YFilter { return afiSafi.YFilter }

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) SetFilter(yf yfilter.YFilter) { afiSafi.YFilter = yf }

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "apply-policy" { return "ApplyPolicy" }
    if yname == "ipv4-unicast" { return "Ipv4Unicast" }
    if yname == "ipv6-unicast" { return "Ipv6Unicast" }
    if yname == "ipv4-labeled-unicast" { return "Ipv4LabeledUnicast" }
    if yname == "ipv6-labeled-unicast" { return "Ipv6LabeledUnicast" }
    if yname == "l3vpn-ipv4-unicast" { return "L3VpnIpv4Unicast" }
    if yname == "l3vpn-ipv6-unicast" { return "L3VpnIpv6Unicast" }
    if yname == "l3vpn-ipv4-multicast" { return "L3VpnIpv4Multicast" }
    if yname == "l3vpn-ipv6-multicast" { return "L3VpnIpv6Multicast" }
    if yname == "l2vpn-vpls" { return "L2VpnVpls" }
    if yname == "l2vpn-evpn" { return "L2VpnEvpn" }
    if yname == "use-multiple-paths" { return "UseMultiplePaths" }
    return ""
}

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetSegmentPath() string {
    return "afi-safi" + "[afi-safi-name='" + fmt.Sprintf("%v", afiSafi.AfiSafiName) + "']"
}

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &afiSafi.Config
    }
    if childYangName == "state" {
        return &afiSafi.State
    }
    if childYangName == "graceful-restart" {
        return &afiSafi.GracefulRestart
    }
    if childYangName == "apply-policy" {
        return &afiSafi.ApplyPolicy
    }
    if childYangName == "ipv4-unicast" {
        return &afiSafi.Ipv4Unicast
    }
    if childYangName == "ipv6-unicast" {
        return &afiSafi.Ipv6Unicast
    }
    if childYangName == "ipv4-labeled-unicast" {
        return &afiSafi.Ipv4LabeledUnicast
    }
    if childYangName == "ipv6-labeled-unicast" {
        return &afiSafi.Ipv6LabeledUnicast
    }
    if childYangName == "l3vpn-ipv4-unicast" {
        return &afiSafi.L3VpnIpv4Unicast
    }
    if childYangName == "l3vpn-ipv6-unicast" {
        return &afiSafi.L3VpnIpv6Unicast
    }
    if childYangName == "l3vpn-ipv4-multicast" {
        return &afiSafi.L3VpnIpv4Multicast
    }
    if childYangName == "l3vpn-ipv6-multicast" {
        return &afiSafi.L3VpnIpv6Multicast
    }
    if childYangName == "l2vpn-vpls" {
        return &afiSafi.L2VpnVpls
    }
    if childYangName == "l2vpn-evpn" {
        return &afiSafi.L2VpnEvpn
    }
    if childYangName == "use-multiple-paths" {
        return &afiSafi.UseMultiplePaths
    }
    return nil
}

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &afiSafi.Config
    children["state"] = &afiSafi.State
    children["graceful-restart"] = &afiSafi.GracefulRestart
    children["apply-policy"] = &afiSafi.ApplyPolicy
    children["ipv4-unicast"] = &afiSafi.Ipv4Unicast
    children["ipv6-unicast"] = &afiSafi.Ipv6Unicast
    children["ipv4-labeled-unicast"] = &afiSafi.Ipv4LabeledUnicast
    children["ipv6-labeled-unicast"] = &afiSafi.Ipv6LabeledUnicast
    children["l3vpn-ipv4-unicast"] = &afiSafi.L3VpnIpv4Unicast
    children["l3vpn-ipv6-unicast"] = &afiSafi.L3VpnIpv6Unicast
    children["l3vpn-ipv4-multicast"] = &afiSafi.L3VpnIpv4Multicast
    children["l3vpn-ipv6-multicast"] = &afiSafi.L3VpnIpv6Multicast
    children["l2vpn-vpls"] = &afiSafi.L2VpnVpls
    children["l2vpn-evpn"] = &afiSafi.L2VpnEvpn
    children["use-multiple-paths"] = &afiSafi.UseMultiplePaths
    return children
}

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = afiSafi.AfiSafiName
    return leafs
}

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetBundleName() string { return "openconfig" }

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetYangName() string { return "afi-safi" }

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) SetParent(parent types.Entity) { afiSafi.parent = parent }

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetParent() types.Entity { return afiSafi.parent }

func (afiSafi *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi) GetParentYangName() string { return "afi-safis" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config
// Configuration parameters for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = config.AfiSafiName
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Config) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State
// State information relating to the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "enabled" { return "Enabled" }
    if yname == "active" { return "Active" }
    if yname == "prefixes" { return "Prefixes" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefixes" {
        return &state.Prefixes
    }
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefixes"] = &state.Prefixes
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = state.AfiSafiName
    leafs["enabled"] = state.Enabled
    leafs["active"] = state.Active
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes
// Prefix counters for the BGP session
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes struct {
    parent types.Entity
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

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetFilter() yfilter.YFilter { return prefixes.YFilter }

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) SetFilter(yf yfilter.YFilter) { prefixes.YFilter = yf }

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetGoName(yname string) string {
    if yname == "received" { return "Received" }
    if yname == "sent" { return "Sent" }
    if yname == "installed" { return "Installed" }
    return ""
}

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetSegmentPath() string {
    return "prefixes"
}

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received"] = prefixes.Received
    leafs["sent"] = prefixes.Sent
    leafs["installed"] = prefixes.Installed
    return leafs
}

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetBundleName() string { return "openconfig" }

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetYangName() string { return "prefixes" }

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) SetParent(parent types.Entity) { prefixes.parent = parent }

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetParent() types.Entity { return prefixes.parent }

func (prefixes *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_State_Prefixes) GetParentYangName() string { return "state" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart
// Parameters relating to BGP graceful-restart
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration options for BGP graceful-restart.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config

    // State information for BGP graceful-restart.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &gracefulRestart.Config
    }
    if childYangName == "state" {
        return &gracefulRestart.State
    }
    return nil
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &gracefulRestart.Config
    children["state"] = &gracefulRestart.State
    return children
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetBundleName() string { return "openconfig" }

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config
// Configuration options for BGP graceful-restart
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_Config) GetParentYangName() string { return "graceful-restart" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State
// State information for BGP graceful-restart
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "received" { return "Received" }
    if yname == "advertised" { return "Advertised" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["received"] = state.Received
    leafs["advertised"] = state.Advertised
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_GracefulRestart_State) GetParentYangName() string { return "graceful-restart" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State
}

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetFilter() yfilter.YFilter { return applyPolicy.YFilter }

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) SetFilter(yf yfilter.YFilter) { applyPolicy.YFilter = yf }

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetSegmentPath() string {
    return "apply-policy"
}

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &applyPolicy.Config
    }
    if childYangName == "state" {
        return &applyPolicy.State
    }
    return nil
}

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &applyPolicy.Config
    children["state"] = &applyPolicy.State
    return children
}

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetBundleName() string { return "openconfig" }

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetYangName() string { return "apply-policy" }

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) SetParent(parent types.Entity) { applyPolicy.parent = parent }

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetParent() types.Entity { return applyPolicy.parent }

func (applyPolicy *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config
// Policy configuration data.
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = config.ImportPolicy
    leafs["default-import-policy"] = config.DefaultImportPolicy
    leafs["export-policy"] = config.ExportPolicy
    leafs["default-export-policy"] = config.DefaultExportPolicy
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_Config) GetParentYangName() string { return "apply-policy" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State
// Operational state for routing policy
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = state.ImportPolicy
    leafs["default-import-policy"] = state.DefaultImportPolicy
    leafs["export-policy"] = state.ExportPolicy
    leafs["default-export-policy"] = state.DefaultExportPolicy
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_ApplyPolicy_State) GetParentYangName() string { return "apply-policy" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast
// IPv4 unicast configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State
}

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetFilter() yfilter.YFilter { return ipv4Unicast.YFilter }

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) SetFilter(yf yfilter.YFilter) { ipv4Unicast.YFilter = yf }

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetSegmentPath() string {
    return "ipv4-unicast"
}

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv4Unicast.PrefixLimit
    }
    if childYangName == "config" {
        return &ipv4Unicast.Config
    }
    if childYangName == "state" {
        return &ipv4Unicast.State
    }
    return nil
}

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv4Unicast.PrefixLimit
    children["config"] = &ipv4Unicast.Config
    children["state"] = &ipv4Unicast.State
    return children
}

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetBundleName() string { return "openconfig" }

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetYangName() string { return "ipv4-unicast" }

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) SetParent(parent types.Entity) { ipv4Unicast.parent = parent }

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetParent() types.Entity { return ipv4Unicast.parent }

func (ipv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetParentYangName() string { return "ipv4-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = config.SendDefaultRoute
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetParentYangName() string { return "ipv4-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = state.SendDefaultRoute
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4Unicast_State) GetParentYangName() string { return "ipv4-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast
// IPv6 unicast configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State
}

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetFilter() yfilter.YFilter { return ipv6Unicast.YFilter }

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) SetFilter(yf yfilter.YFilter) { ipv6Unicast.YFilter = yf }

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetSegmentPath() string {
    return "ipv6-unicast"
}

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv6Unicast.PrefixLimit
    }
    if childYangName == "config" {
        return &ipv6Unicast.Config
    }
    if childYangName == "state" {
        return &ipv6Unicast.State
    }
    return nil
}

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv6Unicast.PrefixLimit
    children["config"] = &ipv6Unicast.Config
    children["state"] = &ipv6Unicast.State
    return children
}

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetBundleName() string { return "openconfig" }

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetYangName() string { return "ipv6-unicast" }

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) SetParent(parent types.Entity) { ipv6Unicast.parent = parent }

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetParent() types.Entity { return ipv6Unicast.parent }

func (ipv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetParentYangName() string { return "ipv6-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = config.SendDefaultRoute
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetParentYangName() string { return "ipv6-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = state.SendDefaultRoute
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6Unicast_State) GetParentYangName() string { return "ipv6-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast
// IPv4 Labeled Unicast configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetFilter() yfilter.YFilter { return ipv4LabeledUnicast.YFilter }

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) SetFilter(yf yfilter.YFilter) { ipv4LabeledUnicast.YFilter = yf }

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetSegmentPath() string {
    return "ipv4-labeled-unicast"
}

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv4LabeledUnicast.PrefixLimit
    }
    return nil
}

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv4LabeledUnicast.PrefixLimit
    return children
}

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetBundleName() string { return "openconfig" }

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetYangName() string { return "ipv4-labeled-unicast" }

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) SetParent(parent types.Entity) { ipv4LabeledUnicast.parent = parent }

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetParent() types.Entity { return ipv4LabeledUnicast.parent }

func (ipv4LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetParentYangName() string { return "ipv4-labeled-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast
// IPv6 Labeled Unicast configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetFilter() yfilter.YFilter { return ipv6LabeledUnicast.YFilter }

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) SetFilter(yf yfilter.YFilter) { ipv6LabeledUnicast.YFilter = yf }

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetSegmentPath() string {
    return "ipv6-labeled-unicast"
}

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv6LabeledUnicast.PrefixLimit
    }
    return nil
}

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv6LabeledUnicast.PrefixLimit
    return children
}

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetBundleName() string { return "openconfig" }

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetYangName() string { return "ipv6-labeled-unicast" }

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) SetParent(parent types.Entity) { ipv6LabeledUnicast.parent = parent }

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetParent() types.Entity { return ipv6LabeledUnicast.parent }

func (ipv6LabeledUnicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetParentYangName() string { return "ipv6-labeled-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast
// Unicast IPv4 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetFilter() yfilter.YFilter { return l3VpnIpv4Unicast.YFilter }

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv4Unicast.YFilter = yf }

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetSegmentPath() string {
    return "l3vpn-ipv4-unicast"
}

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv4Unicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv4Unicast.PrefixLimit
    return children
}

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetYangName() string { return "l3vpn-ipv4-unicast" }

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) SetParent(parent types.Entity) { l3VpnIpv4Unicast.parent = parent }

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetParent() types.Entity { return l3VpnIpv4Unicast.parent }

func (l3VpnIpv4Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv4-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast
// Unicast IPv6 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetFilter() yfilter.YFilter { return l3VpnIpv6Unicast.YFilter }

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv6Unicast.YFilter = yf }

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetSegmentPath() string {
    return "l3vpn-ipv6-unicast"
}

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv6Unicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv6Unicast.PrefixLimit
    return children
}

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetYangName() string { return "l3vpn-ipv6-unicast" }

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) SetParent(parent types.Entity) { l3VpnIpv6Unicast.parent = parent }

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetParent() types.Entity { return l3VpnIpv6Unicast.parent }

func (l3VpnIpv6Unicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv6-unicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast
// Multicast IPv4 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetFilter() yfilter.YFilter { return l3VpnIpv4Multicast.YFilter }

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv4Multicast.YFilter = yf }

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetSegmentPath() string {
    return "l3vpn-ipv4-multicast"
}

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv4Multicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv4Multicast.PrefixLimit
    return children
}

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetYangName() string { return "l3vpn-ipv4-multicast" }

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) SetParent(parent types.Entity) { l3VpnIpv4Multicast.parent = parent }

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetParent() types.Entity { return l3VpnIpv4Multicast.parent }

func (l3VpnIpv4Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv4-multicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast
// Multicast IPv6 L3VPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetFilter() yfilter.YFilter { return l3VpnIpv6Multicast.YFilter }

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv6Multicast.YFilter = yf }

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetSegmentPath() string {
    return "l3vpn-ipv6-multicast"
}

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv6Multicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv6Multicast.PrefixLimit
    return children
}

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetYangName() string { return "l3vpn-ipv6-multicast" }

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) SetParent(parent types.Entity) { l3VpnIpv6Multicast.parent = parent }

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetParent() types.Entity { return l3VpnIpv6Multicast.parent }

func (l3VpnIpv6Multicast *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv6-multicast" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls
// BGP-signalled VPLS configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
}

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetFilter() yfilter.YFilter { return l2VpnVpls.YFilter }

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) SetFilter(yf yfilter.YFilter) { l2VpnVpls.YFilter = yf }

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetSegmentPath() string {
    return "l2vpn-vpls"
}

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l2VpnVpls.PrefixLimit
    }
    return nil
}

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l2VpnVpls.PrefixLimit
    return children
}

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetBundleName() string { return "openconfig" }

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetYangName() string { return "l2vpn-vpls" }

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) SetParent(parent types.Entity) { l2VpnVpls.parent = parent }

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetParent() types.Entity { return l2VpnVpls.parent }

func (l2VpnVpls *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetParentYangName() string { return "l2vpn-vpls" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn
// BGP EVPN configuration options
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
}

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetFilter() yfilter.YFilter { return l2VpnEvpn.YFilter }

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) SetFilter(yf yfilter.YFilter) { l2VpnEvpn.YFilter = yf }

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetSegmentPath() string {
    return "l2vpn-evpn"
}

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l2VpnEvpn.PrefixLimit
    }
    return nil
}

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l2VpnEvpn.PrefixLimit
    return children
}

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetBundleName() string { return "openconfig" }

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetYangName() string { return "l2vpn-evpn" }

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) SetParent(parent types.Entity) { l2VpnEvpn.parent = parent }

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetParent() types.Entity { return l2VpnEvpn.parent }

func (l2VpnEvpn *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetParentYangName() string { return "l2vpn-evpn" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths
// Parameters related to the use of multiple-paths for the same
// NLRI when they are received only from this neighbor
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to multipath.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config

    // State parameters relating to multipath.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State

    // Multipath configuration for eBGP.
    Ebgp Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetFilter() yfilter.YFilter { return useMultiplePaths.YFilter }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) SetFilter(yf yfilter.YFilter) { useMultiplePaths.YFilter = yf }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "ebgp" { return "Ebgp" }
    return ""
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetSegmentPath() string {
    return "use-multiple-paths"
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &useMultiplePaths.Config
    }
    if childYangName == "state" {
        return &useMultiplePaths.State
    }
    if childYangName == "ebgp" {
        return &useMultiplePaths.Ebgp
    }
    return nil
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &useMultiplePaths.Config
    children["state"] = &useMultiplePaths.State
    children["ebgp"] = &useMultiplePaths.Ebgp
    return children
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetBundleName() string { return "openconfig" }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetYangName() string { return "use-multiple-paths" }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) SetParent(parent types.Entity) { useMultiplePaths.parent = parent }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetParent() types.Entity { return useMultiplePaths.parent }

func (useMultiplePaths *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths) GetParentYangName() string { return "afi-safi" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_State) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp
// Multipath configuration for eBGP
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetFilter() yfilter.YFilter { return ebgp.YFilter }

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) SetFilter(yf yfilter.YFilter) { ebgp.YFilter = yf }

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetSegmentPath() string {
    return "ebgp"
}

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ebgp.Config
    }
    if childYangName == "state" {
        return &ebgp.State
    }
    return nil
}

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ebgp.Config
    children["state"] = &ebgp.State
    return children
}

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetBundleName() string { return "openconfig" }

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetYangName() string { return "ebgp" }

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) SetParent(parent types.Entity) { ebgp.parent = parent }

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetParent() types.Entity { return ebgp.parent }

func (ebgp *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    return ""
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = config.AllowMultipleAs
    return leafs
}

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetYangName() string { return "config" }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetParentYangName() string { return "ebgp" }

// Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Allow multipath to use paths from different neighbouring ASes.  The default
    // is to only consider multiple paths from the same neighbouring AS. The type
    // is bool. The default value is false.
    AllowMultipleAs interface{}
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    return ""
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = state.AllowMultipleAs
    return leafs
}

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetYangName() string { return "state" }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_Neighbors_Neighbor_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetParentYangName() string { return "ebgp" }

// Bgp_PeerGroups
// Configuration for BGP peer-groups
type Bgp_PeerGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of BGP peer-groups configured on the local system - uniquely
    // identified by peer-group name. The type is slice of
    // Bgp_PeerGroups_PeerGroup.
    PeerGroup []Bgp_PeerGroups_PeerGroup
}

func (peerGroups *Bgp_PeerGroups) GetFilter() yfilter.YFilter { return peerGroups.YFilter }

func (peerGroups *Bgp_PeerGroups) SetFilter(yf yfilter.YFilter) { peerGroups.YFilter = yf }

func (peerGroups *Bgp_PeerGroups) GetGoName(yname string) string {
    if yname == "peer-group" { return "PeerGroup" }
    return ""
}

func (peerGroups *Bgp_PeerGroups) GetSegmentPath() string {
    return "peer-groups"
}

func (peerGroups *Bgp_PeerGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-group" {
        for _, c := range peerGroups.PeerGroup {
            if peerGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Bgp_PeerGroups_PeerGroup{}
        peerGroups.PeerGroup = append(peerGroups.PeerGroup, child)
        return &peerGroups.PeerGroup[len(peerGroups.PeerGroup)-1]
    }
    return nil
}

func (peerGroups *Bgp_PeerGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerGroups.PeerGroup {
        children[peerGroups.PeerGroup[i].GetSegmentPath()] = &peerGroups.PeerGroup[i]
    }
    return children
}

func (peerGroups *Bgp_PeerGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerGroups *Bgp_PeerGroups) GetBundleName() string { return "openconfig" }

func (peerGroups *Bgp_PeerGroups) GetYangName() string { return "peer-groups" }

func (peerGroups *Bgp_PeerGroups) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (peerGroups *Bgp_PeerGroups) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (peerGroups *Bgp_PeerGroups) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (peerGroups *Bgp_PeerGroups) SetParent(parent types.Entity) { peerGroups.parent = parent }

func (peerGroups *Bgp_PeerGroups) GetParent() types.Entity { return peerGroups.parent }

func (peerGroups *Bgp_PeerGroups) GetParentYangName() string { return "bgp" }

// Bgp_PeerGroups_PeerGroup
// List of BGP peer-groups configured on the local system -
// uniquely identified by peer-group name
type Bgp_PeerGroups_PeerGroup struct {
    parent types.Entity
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

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetFilter() yfilter.YFilter { return peerGroup.YFilter }

func (peerGroup *Bgp_PeerGroups_PeerGroup) SetFilter(yf yfilter.YFilter) { peerGroup.YFilter = yf }

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetGoName(yname string) string {
    if yname == "peer-group-name" { return "PeerGroupName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "timers" { return "Timers" }
    if yname == "transport" { return "Transport" }
    if yname == "error-handling" { return "ErrorHandling" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "logging-options" { return "LoggingOptions" }
    if yname == "ebgp-multihop" { return "EbgpMultihop" }
    if yname == "route-reflector" { return "RouteReflector" }
    if yname == "as-path-options" { return "AsPathOptions" }
    if yname == "add-paths" { return "AddPaths" }
    if yname == "use-multiple-paths" { return "UseMultiplePaths" }
    if yname == "apply-policy" { return "ApplyPolicy" }
    if yname == "afi-safis" { return "AfiSafis" }
    return ""
}

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetSegmentPath() string {
    return "peer-group" + "[peer-group-name='" + fmt.Sprintf("%v", peerGroup.PeerGroupName) + "']"
}

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &peerGroup.Config
    }
    if childYangName == "state" {
        return &peerGroup.State
    }
    if childYangName == "timers" {
        return &peerGroup.Timers
    }
    if childYangName == "transport" {
        return &peerGroup.Transport
    }
    if childYangName == "error-handling" {
        return &peerGroup.ErrorHandling
    }
    if childYangName == "graceful-restart" {
        return &peerGroup.GracefulRestart
    }
    if childYangName == "logging-options" {
        return &peerGroup.LoggingOptions
    }
    if childYangName == "ebgp-multihop" {
        return &peerGroup.EbgpMultihop
    }
    if childYangName == "route-reflector" {
        return &peerGroup.RouteReflector
    }
    if childYangName == "as-path-options" {
        return &peerGroup.AsPathOptions
    }
    if childYangName == "add-paths" {
        return &peerGroup.AddPaths
    }
    if childYangName == "use-multiple-paths" {
        return &peerGroup.UseMultiplePaths
    }
    if childYangName == "apply-policy" {
        return &peerGroup.ApplyPolicy
    }
    if childYangName == "afi-safis" {
        return &peerGroup.AfiSafis
    }
    return nil
}

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &peerGroup.Config
    children["state"] = &peerGroup.State
    children["timers"] = &peerGroup.Timers
    children["transport"] = &peerGroup.Transport
    children["error-handling"] = &peerGroup.ErrorHandling
    children["graceful-restart"] = &peerGroup.GracefulRestart
    children["logging-options"] = &peerGroup.LoggingOptions
    children["ebgp-multihop"] = &peerGroup.EbgpMultihop
    children["route-reflector"] = &peerGroup.RouteReflector
    children["as-path-options"] = &peerGroup.AsPathOptions
    children["add-paths"] = &peerGroup.AddPaths
    children["use-multiple-paths"] = &peerGroup.UseMultiplePaths
    children["apply-policy"] = &peerGroup.ApplyPolicy
    children["afi-safis"] = &peerGroup.AfiSafis
    return children
}

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-group-name"] = peerGroup.PeerGroupName
    return leafs
}

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetBundleName() string { return "openconfig" }

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetYangName() string { return "peer-group" }

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (peerGroup *Bgp_PeerGroups_PeerGroup) SetParent(parent types.Entity) { peerGroup.parent = parent }

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetParent() types.Entity { return peerGroup.parent }

func (peerGroup *Bgp_PeerGroups_PeerGroup) GetParentYangName() string { return "peer-groups" }

// Bgp_PeerGroups_PeerGroup_Config
// Configuration parameters relating to the BGP neighbor or
// group
type Bgp_PeerGroups_PeerGroup_Config struct {
    parent types.Entity
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
    // The type is one of the following: PRIVATEASREPLACEALLPRIVATEASREMOVEALL.
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

func (config *Bgp_PeerGroups_PeerGroup_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_Config) GetGoName(yname string) string {
    if yname == "peer-group-name" { return "PeerGroupName" }
    if yname == "peer-as" { return "PeerAs" }
    if yname == "local-as" { return "LocalAs" }
    if yname == "peer-type" { return "PeerType" }
    if yname == "auth-password" { return "AuthPassword" }
    if yname == "remove-private-as" { return "RemovePrivateAs" }
    if yname == "route-flap-damping" { return "RouteFlapDamping" }
    if yname == "send-community" { return "SendCommunity" }
    if yname == "description" { return "Description" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-group-name"] = config.PeerGroupName
    leafs["peer-as"] = config.PeerAs
    leafs["local-as"] = config.LocalAs
    leafs["peer-type"] = config.PeerType
    leafs["auth-password"] = config.AuthPassword
    leafs["remove-private-as"] = config.RemovePrivateAs
    leafs["route-flap-damping"] = config.RouteFlapDamping
    leafs["send-community"] = config.SendCommunity
    leafs["description"] = config.Description
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_Config) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_State
// State information relating to the BGP peer-group
type Bgp_PeerGroups_PeerGroup_State struct {
    parent types.Entity
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
    // The type is one of the following: PRIVATEASREPLACEALLPRIVATEASREMOVEALL.
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

func (state *Bgp_PeerGroups_PeerGroup_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_State) GetGoName(yname string) string {
    if yname == "peer-group-name" { return "PeerGroupName" }
    if yname == "peer-as" { return "PeerAs" }
    if yname == "local-as" { return "LocalAs" }
    if yname == "peer-type" { return "PeerType" }
    if yname == "auth-password" { return "AuthPassword" }
    if yname == "remove-private-as" { return "RemovePrivateAs" }
    if yname == "route-flap-damping" { return "RouteFlapDamping" }
    if yname == "send-community" { return "SendCommunity" }
    if yname == "description" { return "Description" }
    if yname == "total-paths" { return "TotalPaths" }
    if yname == "total-prefixes" { return "TotalPrefixes" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-group-name"] = state.PeerGroupName
    leafs["peer-as"] = state.PeerAs
    leafs["local-as"] = state.LocalAs
    leafs["peer-type"] = state.PeerType
    leafs["auth-password"] = state.AuthPassword
    leafs["remove-private-as"] = state.RemovePrivateAs
    leafs["route-flap-damping"] = state.RouteFlapDamping
    leafs["send-community"] = state.SendCommunity
    leafs["description"] = state.Description
    leafs["total-paths"] = state.TotalPaths
    leafs["total-prefixes"] = state.TotalPrefixes
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_State) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_Timers
// Timers related to a BGP peer-group
type Bgp_PeerGroups_PeerGroup_Timers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to timers used for the BGP neighbor or
    // peer group.
    Config Bgp_PeerGroups_PeerGroup_Timers_Config

    // State information relating to the timers used for the BGP group.
    State Bgp_PeerGroups_PeerGroup_Timers_State
}

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Bgp_PeerGroups_PeerGroup_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &timers.Config
    }
    if childYangName == "state" {
        return &timers.State
    }
    return nil
}

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &timers.Config
    children["state"] = &timers.State
    return children
}

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetBundleName() string { return "openconfig" }

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetYangName() string { return "timers" }

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (timers *Bgp_PeerGroups_PeerGroup_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Bgp_PeerGroups_PeerGroup_Timers) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_Timers_Config
// Configuration parameters relating to timers used for the
// BGP neighbor or peer group
type Bgp_PeerGroups_PeerGroup_Timers_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetGoName(yname string) string {
    if yname == "connect-retry" { return "ConnectRetry" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "keepalive-interval" { return "KeepaliveInterval" }
    if yname == "minimum-advertisement-interval" { return "MinimumAdvertisementInterval" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["connect-retry"] = config.ConnectRetry
    leafs["hold-time"] = config.HoldTime
    leafs["keepalive-interval"] = config.KeepaliveInterval
    leafs["minimum-advertisement-interval"] = config.MinimumAdvertisementInterval
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_Timers_Config) GetParentYangName() string { return "timers" }

// Bgp_PeerGroups_PeerGroup_Timers_State
// State information relating to the timers used for the BGP
// group
type Bgp_PeerGroups_PeerGroup_Timers_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetGoName(yname string) string {
    if yname == "connect-retry" { return "ConnectRetry" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "keepalive-interval" { return "KeepaliveInterval" }
    if yname == "minimum-advertisement-interval" { return "MinimumAdvertisementInterval" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["connect-retry"] = state.ConnectRetry
    leafs["hold-time"] = state.HoldTime
    leafs["keepalive-interval"] = state.KeepaliveInterval
    leafs["minimum-advertisement-interval"] = state.MinimumAdvertisementInterval
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_Timers_State) GetParentYangName() string { return "timers" }

// Bgp_PeerGroups_PeerGroup_Transport
// Transport session parameters for the BGP peer-group
type Bgp_PeerGroups_PeerGroup_Transport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the transport session(s) used for the
    // BGP neighbor or group.
    Config Bgp_PeerGroups_PeerGroup_Transport_Config

    // State information relating to the transport session(s) used for the BGP
    // neighbor or group.
    State Bgp_PeerGroups_PeerGroup_Transport_State
}

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetFilter() yfilter.YFilter { return transport.YFilter }

func (transport *Bgp_PeerGroups_PeerGroup_Transport) SetFilter(yf yfilter.YFilter) { transport.YFilter = yf }

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetSegmentPath() string {
    return "transport"
}

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &transport.Config
    }
    if childYangName == "state" {
        return &transport.State
    }
    return nil
}

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &transport.Config
    children["state"] = &transport.State
    return children
}

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetBundleName() string { return "openconfig" }

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetYangName() string { return "transport" }

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (transport *Bgp_PeerGroups_PeerGroup_Transport) SetParent(parent types.Entity) { transport.parent = parent }

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetParent() types.Entity { return transport.parent }

func (transport *Bgp_PeerGroups_PeerGroup_Transport) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_Transport_Config
// Configuration parameters relating to the transport
// session(s) used for the BGP neighbor or group
type Bgp_PeerGroups_PeerGroup_Transport_Config struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
    // or string.
    LocalAddress interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetGoName(yname string) string {
    if yname == "tcp-mss" { return "TcpMss" }
    if yname == "mtu-discovery" { return "MtuDiscovery" }
    if yname == "passive-mode" { return "PassiveMode" }
    if yname == "local-address" { return "LocalAddress" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-mss"] = config.TcpMss
    leafs["mtu-discovery"] = config.MtuDiscovery
    leafs["passive-mode"] = config.PassiveMode
    leafs["local-address"] = config.LocalAddress
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_Transport_Config) GetParentYangName() string { return "transport" }

// Bgp_PeerGroups_PeerGroup_Transport_State
// State information relating to the transport session(s)
// used for the BGP neighbor or group
type Bgp_PeerGroups_PeerGroup_Transport_State struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
    // or string.
    LocalAddress interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetGoName(yname string) string {
    if yname == "tcp-mss" { return "TcpMss" }
    if yname == "mtu-discovery" { return "MtuDiscovery" }
    if yname == "passive-mode" { return "PassiveMode" }
    if yname == "local-address" { return "LocalAddress" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-mss"] = state.TcpMss
    leafs["mtu-discovery"] = state.MtuDiscovery
    leafs["passive-mode"] = state.PassiveMode
    leafs["local-address"] = state.LocalAddress
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_Transport_State) GetParentYangName() string { return "transport" }

// Bgp_PeerGroups_PeerGroup_ErrorHandling
// Error handling parameters used for the BGP peer-group
type Bgp_PeerGroups_PeerGroup_ErrorHandling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters enabling or modifying the behavior or enhanced
    // error handling mechanisms for the BGP group.
    Config Bgp_PeerGroups_PeerGroup_ErrorHandling_Config

    // State information relating to enhanced error handling mechanisms for the
    // BGP group.
    State Bgp_PeerGroups_PeerGroup_ErrorHandling_State
}

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetFilter() yfilter.YFilter { return errorHandling.YFilter }

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) SetFilter(yf yfilter.YFilter) { errorHandling.YFilter = yf }

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetSegmentPath() string {
    return "error-handling"
}

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &errorHandling.Config
    }
    if childYangName == "state" {
        return &errorHandling.State
    }
    return nil
}

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &errorHandling.Config
    children["state"] = &errorHandling.State
    return children
}

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetBundleName() string { return "openconfig" }

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetYangName() string { return "error-handling" }

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) SetParent(parent types.Entity) { errorHandling.parent = parent }

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetParent() types.Entity { return errorHandling.parent }

func (errorHandling *Bgp_PeerGroups_PeerGroup_ErrorHandling) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_ErrorHandling_Config
// Configuration parameters enabling or modifying the
// behavior or enhanced error handling mechanisms for the BGP
// group
type Bgp_PeerGroups_PeerGroup_ErrorHandling_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify whether erroneous UPDATE messages for which the NLRI can be
    // extracted are reated as though the NLRI is withdrawn - avoiding session
    // reset. The type is bool. The default value is false.
    TreatAsWithdraw interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetGoName(yname string) string {
    if yname == "treat-as-withdraw" { return "TreatAsWithdraw" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["treat-as-withdraw"] = config.TreatAsWithdraw
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_ErrorHandling_Config) GetParentYangName() string { return "error-handling" }

// Bgp_PeerGroups_PeerGroup_ErrorHandling_State
// State information relating to enhanced error handling
// mechanisms for the BGP group
type Bgp_PeerGroups_PeerGroup_ErrorHandling_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify whether erroneous UPDATE messages for which the NLRI can be
    // extracted are reated as though the NLRI is withdrawn - avoiding session
    // reset. The type is bool. The default value is false.
    TreatAsWithdraw interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetGoName(yname string) string {
    if yname == "treat-as-withdraw" { return "TreatAsWithdraw" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["treat-as-withdraw"] = state.TreatAsWithdraw
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_ErrorHandling_State) GetParentYangName() string { return "error-handling" }

// Bgp_PeerGroups_PeerGroup_GracefulRestart
// Parameters relating the graceful restart mechanism for BGP
type Bgp_PeerGroups_PeerGroup_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to graceful-restart.
    Config Bgp_PeerGroups_PeerGroup_GracefulRestart_Config

    // State information associated with graceful-restart.
    State Bgp_PeerGroups_PeerGroup_GracefulRestart_State
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &gracefulRestart.Config
    }
    if childYangName == "state" {
        return &gracefulRestart.State
    }
    return nil
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &gracefulRestart.Config
    children["state"] = &gracefulRestart.State
    return children
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetBundleName() string { return "openconfig" }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_GracefulRestart) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_GracefulRestart_Config
// Configuration parameters relating to graceful-restart
type Bgp_PeerGroups_PeerGroup_GracefulRestart_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "restart-time" { return "RestartTime" }
    if yname == "stale-routes-time" { return "StaleRoutesTime" }
    if yname == "helper-only" { return "HelperOnly" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["restart-time"] = config.RestartTime
    leafs["stale-routes-time"] = config.StaleRoutesTime
    leafs["helper-only"] = config.HelperOnly
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_GracefulRestart_Config) GetParentYangName() string { return "graceful-restart" }

// Bgp_PeerGroups_PeerGroup_GracefulRestart_State
// State information associated with graceful-restart
type Bgp_PeerGroups_PeerGroup_GracefulRestart_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "restart-time" { return "RestartTime" }
    if yname == "stale-routes-time" { return "StaleRoutesTime" }
    if yname == "helper-only" { return "HelperOnly" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["restart-time"] = state.RestartTime
    leafs["stale-routes-time"] = state.StaleRoutesTime
    leafs["helper-only"] = state.HelperOnly
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_GracefulRestart_State) GetParentYangName() string { return "graceful-restart" }

// Bgp_PeerGroups_PeerGroup_LoggingOptions
// Logging options for events related to the BGP neighbor or
// group
type Bgp_PeerGroups_PeerGroup_LoggingOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters enabling or modifying logging for events relating
    // to the BGPgroup.
    Config Bgp_PeerGroups_PeerGroup_LoggingOptions_Config

    // State information relating to logging for the BGP neighbor or group.
    State Bgp_PeerGroups_PeerGroup_LoggingOptions_State
}

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetFilter() yfilter.YFilter { return loggingOptions.YFilter }

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) SetFilter(yf yfilter.YFilter) { loggingOptions.YFilter = yf }

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetSegmentPath() string {
    return "logging-options"
}

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &loggingOptions.Config
    }
    if childYangName == "state" {
        return &loggingOptions.State
    }
    return nil
}

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &loggingOptions.Config
    children["state"] = &loggingOptions.State
    return children
}

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetBundleName() string { return "openconfig" }

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetYangName() string { return "logging-options" }

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) SetParent(parent types.Entity) { loggingOptions.parent = parent }

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetParent() types.Entity { return loggingOptions.parent }

func (loggingOptions *Bgp_PeerGroups_PeerGroup_LoggingOptions) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_LoggingOptions_Config
// Configuration parameters enabling or modifying logging
// for events relating to the BGPgroup
type Bgp_PeerGroups_PeerGroup_LoggingOptions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure logging of peer state changes.  Default is to enable logging of
    // peer state changes. The type is bool. The default value is true.
    LogNeighborStateChanges interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetGoName(yname string) string {
    if yname == "log-neighbor-state-changes" { return "LogNeighborStateChanges" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-neighbor-state-changes"] = config.LogNeighborStateChanges
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_LoggingOptions_Config) GetParentYangName() string { return "logging-options" }

// Bgp_PeerGroups_PeerGroup_LoggingOptions_State
// State information relating to logging for the BGP neighbor
// or group
type Bgp_PeerGroups_PeerGroup_LoggingOptions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure logging of peer state changes.  Default is to enable logging of
    // peer state changes. The type is bool. The default value is true.
    LogNeighborStateChanges interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetGoName(yname string) string {
    if yname == "log-neighbor-state-changes" { return "LogNeighborStateChanges" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["log-neighbor-state-changes"] = state.LogNeighborStateChanges
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_LoggingOptions_State) GetParentYangName() string { return "logging-options" }

// Bgp_PeerGroups_PeerGroup_EbgpMultihop
// eBGP multi-hop parameters for the BGPgroup
type Bgp_PeerGroups_PeerGroup_EbgpMultihop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multihop for the BGP group.
    Config Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config

    // State information for eBGP multihop, for the BGP neighbor or group.
    State Bgp_PeerGroups_PeerGroup_EbgpMultihop_State
}

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetFilter() yfilter.YFilter { return ebgpMultihop.YFilter }

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) SetFilter(yf yfilter.YFilter) { ebgpMultihop.YFilter = yf }

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetSegmentPath() string {
    return "ebgp-multihop"
}

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ebgpMultihop.Config
    }
    if childYangName == "state" {
        return &ebgpMultihop.State
    }
    return nil
}

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ebgpMultihop.Config
    children["state"] = &ebgpMultihop.State
    return children
}

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetBundleName() string { return "openconfig" }

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetYangName() string { return "ebgp-multihop" }

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) SetParent(parent types.Entity) { ebgpMultihop.parent = parent }

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetParent() types.Entity { return ebgpMultihop.parent }

func (ebgpMultihop *Bgp_PeerGroups_PeerGroup_EbgpMultihop) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config
// Configuration parameters relating to eBGP multihop for the
// BGP group
type Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "multihop-ttl" { return "MultihopTtl" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["multihop-ttl"] = config.MultihopTtl
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_EbgpMultihop_Config) GetParentYangName() string { return "ebgp-multihop" }

// Bgp_PeerGroups_PeerGroup_EbgpMultihop_State
// State information for eBGP multihop, for the BGP neighbor
// or group
type Bgp_PeerGroups_PeerGroup_EbgpMultihop_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "multihop-ttl" { return "MultihopTtl" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["multihop-ttl"] = state.MultihopTtl
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_EbgpMultihop_State) GetParentYangName() string { return "ebgp-multihop" }

// Bgp_PeerGroups_PeerGroup_RouteReflector
// Route reflector parameters for the BGPgroup
type Bgp_PeerGroups_PeerGroup_RouteReflector struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuraton parameters relating to route reflection for the BGPgroup.
    Config Bgp_PeerGroups_PeerGroup_RouteReflector_Config

    // State information relating to route reflection for the BGPgroup.
    State Bgp_PeerGroups_PeerGroup_RouteReflector_State
}

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetFilter() yfilter.YFilter { return routeReflector.YFilter }

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) SetFilter(yf yfilter.YFilter) { routeReflector.YFilter = yf }

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetSegmentPath() string {
    return "route-reflector"
}

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &routeReflector.Config
    }
    if childYangName == "state" {
        return &routeReflector.State
    }
    return nil
}

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &routeReflector.Config
    children["state"] = &routeReflector.State
    return children
}

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetBundleName() string { return "openconfig" }

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetYangName() string { return "route-reflector" }

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) SetParent(parent types.Entity) { routeReflector.parent = parent }

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetParent() types.Entity { return routeReflector.parent }

func (routeReflector *Bgp_PeerGroups_PeerGroup_RouteReflector) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_RouteReflector_Config
// Configuraton parameters relating to route reflection
// for the BGPgroup
type Bgp_PeerGroups_PeerGroup_RouteReflector_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route-reflector cluster id to use when local router is configured as a
    // route reflector.  Commonly set at the group level, but allows a different
    // cluster id to be set for each neighbor. The type is one of the following
    // types: int with range: 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouteReflectorClusterId interface{}

    // Configure the neighbor as a route reflector client. The type is bool. The
    // default value is false.
    RouteReflectorClient interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetGoName(yname string) string {
    if yname == "route-reflector-cluster-id" { return "RouteReflectorClusterId" }
    if yname == "route-reflector-client" { return "RouteReflectorClient" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-reflector-cluster-id"] = config.RouteReflectorClusterId
    leafs["route-reflector-client"] = config.RouteReflectorClient
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_RouteReflector_Config) GetParentYangName() string { return "route-reflector" }

// Bgp_PeerGroups_PeerGroup_RouteReflector_State
// State information relating to route reflection for the
// BGPgroup
type Bgp_PeerGroups_PeerGroup_RouteReflector_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // route-reflector cluster id to use when local router is configured as a
    // route reflector.  Commonly set at the group level, but allows a different
    // cluster id to be set for each neighbor. The type is one of the following
    // types: int with range: 0..4294967295, or string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouteReflectorClusterId interface{}

    // Configure the neighbor as a route reflector client. The type is bool. The
    // default value is false.
    RouteReflectorClient interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetGoName(yname string) string {
    if yname == "route-reflector-cluster-id" { return "RouteReflectorClusterId" }
    if yname == "route-reflector-client" { return "RouteReflectorClient" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-reflector-cluster-id"] = state.RouteReflectorClusterId
    leafs["route-reflector-client"] = state.RouteReflectorClient
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_RouteReflector_State) GetParentYangName() string { return "route-reflector" }

// Bgp_PeerGroups_PeerGroup_AsPathOptions
// AS_PATH manipulation parameters for the BGP neighbor or
// group
type Bgp_PeerGroups_PeerGroup_AsPathOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to AS_PATH manipulation for the BGP peer
    // or group.
    Config Bgp_PeerGroups_PeerGroup_AsPathOptions_Config

    // State information relating to the AS_PATH manipulation mechanisms for the
    // BGP peer or group.
    State Bgp_PeerGroups_PeerGroup_AsPathOptions_State
}

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetFilter() yfilter.YFilter { return asPathOptions.YFilter }

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) SetFilter(yf yfilter.YFilter) { asPathOptions.YFilter = yf }

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetSegmentPath() string {
    return "as-path-options"
}

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &asPathOptions.Config
    }
    if childYangName == "state" {
        return &asPathOptions.State
    }
    return nil
}

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &asPathOptions.Config
    children["state"] = &asPathOptions.State
    return children
}

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetBundleName() string { return "openconfig" }

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetYangName() string { return "as-path-options" }

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) SetParent(parent types.Entity) { asPathOptions.parent = parent }

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetParent() types.Entity { return asPathOptions.parent }

func (asPathOptions *Bgp_PeerGroups_PeerGroup_AsPathOptions) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_AsPathOptions_Config
// Configuration parameters relating to AS_PATH manipulation
// for the BGP peer or group
type Bgp_PeerGroups_PeerGroup_AsPathOptions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the number of occurrences of the local BGP speaker's AS that can
    // occur within the AS_PATH before it is rejected. The type is interface{}
    // with range: 0..255. The default value is 0.
    AllowOwnAs interface{}

    // Replace occurrences of the peer's AS in the AS_PATH with the local
    // autonomous system number. The type is bool. The default value is false.
    ReplacePeerAs interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetGoName(yname string) string {
    if yname == "allow-own-as" { return "AllowOwnAs" }
    if yname == "replace-peer-as" { return "ReplacePeerAs" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-own-as"] = config.AllowOwnAs
    leafs["replace-peer-as"] = config.ReplacePeerAs
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AsPathOptions_Config) GetParentYangName() string { return "as-path-options" }

// Bgp_PeerGroups_PeerGroup_AsPathOptions_State
// State information relating to the AS_PATH manipulation
// mechanisms for the BGP peer or group
type Bgp_PeerGroups_PeerGroup_AsPathOptions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the number of occurrences of the local BGP speaker's AS that can
    // occur within the AS_PATH before it is rejected. The type is interface{}
    // with range: 0..255. The default value is 0.
    AllowOwnAs interface{}

    // Replace occurrences of the peer's AS in the AS_PATH with the local
    // autonomous system number. The type is bool. The default value is false.
    ReplacePeerAs interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetGoName(yname string) string {
    if yname == "allow-own-as" { return "AllowOwnAs" }
    if yname == "replace-peer-as" { return "ReplacePeerAs" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-own-as"] = state.AllowOwnAs
    leafs["replace-peer-as"] = state.ReplacePeerAs
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AsPathOptions_State) GetParentYangName() string { return "as-path-options" }

// Bgp_PeerGroups_PeerGroup_AddPaths
// Parameters relating to the advertisement and receipt of
// multiple paths for a single NLRI (add-paths)
type Bgp_PeerGroups_PeerGroup_AddPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to ADD_PATHS.
    Config Bgp_PeerGroups_PeerGroup_AddPaths_Config

    // State information associated with ADD_PATHS.
    State Bgp_PeerGroups_PeerGroup_AddPaths_State
}

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetFilter() yfilter.YFilter { return addPaths.YFilter }

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) SetFilter(yf yfilter.YFilter) { addPaths.YFilter = yf }

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetSegmentPath() string {
    return "add-paths"
}

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &addPaths.Config
    }
    if childYangName == "state" {
        return &addPaths.State
    }
    return nil
}

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &addPaths.Config
    children["state"] = &addPaths.State
    return children
}

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetBundleName() string { return "openconfig" }

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetYangName() string { return "add-paths" }

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) SetParent(parent types.Entity) { addPaths.parent = parent }

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetParent() types.Entity { return addPaths.parent }

func (addPaths *Bgp_PeerGroups_PeerGroup_AddPaths) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_AddPaths_Config
// Configuration parameters relating to ADD_PATHS
type Bgp_PeerGroups_PeerGroup_AddPaths_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetGoName(yname string) string {
    if yname == "receive" { return "Receive" }
    if yname == "send-max" { return "SendMax" }
    if yname == "eligible-prefix-policy" { return "EligiblePrefixPolicy" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receive"] = config.Receive
    leafs["send-max"] = config.SendMax
    leafs["eligible-prefix-policy"] = config.EligiblePrefixPolicy
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AddPaths_Config) GetParentYangName() string { return "add-paths" }

// Bgp_PeerGroups_PeerGroup_AddPaths_State
// State information associated with ADD_PATHS
type Bgp_PeerGroups_PeerGroup_AddPaths_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetGoName(yname string) string {
    if yname == "receive" { return "Receive" }
    if yname == "send-max" { return "SendMax" }
    if yname == "eligible-prefix-policy" { return "EligiblePrefixPolicy" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receive"] = state.Receive
    leafs["send-max"] = state.SendMax
    leafs["eligible-prefix-policy"] = state.EligiblePrefixPolicy
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AddPaths_State) GetParentYangName() string { return "add-paths" }

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths
// Parameters related to the use of multiple paths for the
// same NLRI
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths struct {
    parent types.Entity
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

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetFilter() yfilter.YFilter { return useMultiplePaths.YFilter }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) SetFilter(yf yfilter.YFilter) { useMultiplePaths.YFilter = yf }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "ebgp" { return "Ebgp" }
    if yname == "ibgp" { return "Ibgp" }
    return ""
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetSegmentPath() string {
    return "use-multiple-paths"
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &useMultiplePaths.Config
    }
    if childYangName == "state" {
        return &useMultiplePaths.State
    }
    if childYangName == "ebgp" {
        return &useMultiplePaths.Ebgp
    }
    if childYangName == "ibgp" {
        return &useMultiplePaths.Ibgp
    }
    return nil
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &useMultiplePaths.Config
    children["state"] = &useMultiplePaths.State
    children["ebgp"] = &useMultiplePaths.Ebgp
    children["ibgp"] = &useMultiplePaths.Ibgp
    return children
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetBundleName() string { return "openconfig" }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetYangName() string { return "use-multiple-paths" }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) SetParent(parent types.Entity) { useMultiplePaths.parent = parent }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetParent() types.Entity { return useMultiplePaths.parent }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_UseMultiplePaths) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Config) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_State) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp
// Multipath parameters for eBGP
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetFilter() yfilter.YFilter { return ebgp.YFilter }

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) SetFilter(yf yfilter.YFilter) { ebgp.YFilter = yf }

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetSegmentPath() string {
    return "ebgp"
}

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ebgp.Config
    }
    if childYangName == "state" {
        return &ebgp.State
    }
    return nil
}

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ebgp.Config
    children["state"] = &ebgp.State
    return children
}

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetBundleName() string { return "openconfig" }

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetYangName() string { return "ebgp" }

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) SetParent(parent types.Entity) { ebgp.parent = parent }

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetParent() types.Entity { return ebgp.parent }

func (ebgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = config.AllowMultipleAs
    leafs["maximum-paths"] = config.MaximumPaths
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_Config) GetParentYangName() string { return "ebgp" }

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = state.AllowMultipleAs
    leafs["maximum-paths"] = state.MaximumPaths
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ebgp_State) GetParentYangName() string { return "ebgp" }

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp
// Multipath parameters for iBGP
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to iBGP multipath.
    Config Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config

    // State information relating to iBGP multipath.
    State Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State
}

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetFilter() yfilter.YFilter { return ibgp.YFilter }

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) SetFilter(yf yfilter.YFilter) { ibgp.YFilter = yf }

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetSegmentPath() string {
    return "ibgp"
}

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ibgp.Config
    }
    if childYangName == "state" {
        return &ibgp.State
    }
    return nil
}

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ibgp.Config
    children["state"] = &ibgp.State
    return children
}

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetBundleName() string { return "openconfig" }

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetYangName() string { return "ibgp" }

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) SetParent(parent types.Entity) { ibgp.parent = parent }

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetParent() types.Entity { return ibgp.parent }

func (ibgp *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config
// Configuration parameters relating to iBGP multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetGoName(yname string) string {
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-paths"] = config.MaximumPaths
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_Config) GetParentYangName() string { return "ibgp" }

// Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State
// State information relating to iBGP multipath
type Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetGoName(yname string) string {
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-paths"] = state.MaximumPaths
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_UseMultiplePaths_Ibgp_State) GetParentYangName() string { return "ibgp" }

// Bgp_PeerGroups_PeerGroup_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_PeerGroups_PeerGroup_ApplyPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_PeerGroups_PeerGroup_ApplyPolicy_State
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetFilter() yfilter.YFilter { return applyPolicy.YFilter }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) SetFilter(yf yfilter.YFilter) { applyPolicy.YFilter = yf }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetSegmentPath() string {
    return "apply-policy"
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &applyPolicy.Config
    }
    if childYangName == "state" {
        return &applyPolicy.State
    }
    return nil
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &applyPolicy.Config
    children["state"] = &applyPolicy.State
    return children
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetBundleName() string { return "openconfig" }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetYangName() string { return "apply-policy" }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) SetParent(parent types.Entity) { applyPolicy.parent = parent }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetParent() types.Entity { return applyPolicy.parent }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_ApplyPolicy) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config
// Policy configuration data.
type Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = config.ImportPolicy
    leafs["default-import-policy"] = config.DefaultImportPolicy
    leafs["export-policy"] = config.ExportPolicy
    leafs["default-export-policy"] = config.DefaultExportPolicy
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_ApplyPolicy_Config) GetParentYangName() string { return "apply-policy" }

// Bgp_PeerGroups_PeerGroup_ApplyPolicy_State
// Operational state for routing policy
type Bgp_PeerGroups_PeerGroup_ApplyPolicy_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = state.ImportPolicy
    leafs["default-import-policy"] = state.DefaultImportPolicy
    leafs["export-policy"] = state.ExportPolicy
    leafs["default-export-policy"] = state.DefaultExportPolicy
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_ApplyPolicy_State) GetParentYangName() string { return "apply-policy" }

// Bgp_PeerGroups_PeerGroup_AfiSafis
// Per-address-family configuration parameters associated with
// thegroup
type Bgp_PeerGroups_PeerGroup_AfiSafis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI,SAFI configuration available for the neighbour or group. The type is
    // slice of Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi.
    AfiSafi []Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi
}

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetFilter() yfilter.YFilter { return afiSafis.YFilter }

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) SetFilter(yf yfilter.YFilter) { afiSafis.YFilter = yf }

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetGoName(yname string) string {
    if yname == "afi-safi" { return "AfiSafi" }
    return ""
}

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetSegmentPath() string {
    return "afi-safis"
}

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afi-safi" {
        for _, c := range afiSafis.AfiSafi {
            if afiSafis.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi{}
        afiSafis.AfiSafi = append(afiSafis.AfiSafi, child)
        return &afiSafis.AfiSafi[len(afiSafis.AfiSafi)-1]
    }
    return nil
}

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afiSafis.AfiSafi {
        children[afiSafis.AfiSafi[i].GetSegmentPath()] = &afiSafis.AfiSafi[i]
    }
    return children
}

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetBundleName() string { return "openconfig" }

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetYangName() string { return "afi-safis" }

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) SetParent(parent types.Entity) { afiSafis.parent = parent }

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetParent() types.Entity { return afiSafis.parent }

func (afiSafis *Bgp_PeerGroups_PeerGroup_AfiSafis) GetParentYangName() string { return "peer-group" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi
// AFI,SAFI configuration available for the
// neighbour or group
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the AFI-SAFI name used as a key for
    // the AFI-SAFI list. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
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

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetFilter() yfilter.YFilter { return afiSafi.YFilter }

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) SetFilter(yf yfilter.YFilter) { afiSafi.YFilter = yf }

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "route-selection-options" { return "RouteSelectionOptions" }
    if yname == "use-multiple-paths" { return "UseMultiplePaths" }
    if yname == "apply-policy" { return "ApplyPolicy" }
    if yname == "ipv4-unicast" { return "Ipv4Unicast" }
    if yname == "ipv6-unicast" { return "Ipv6Unicast" }
    if yname == "ipv4-labeled-unicast" { return "Ipv4LabeledUnicast" }
    if yname == "ipv6-labeled-unicast" { return "Ipv6LabeledUnicast" }
    if yname == "l3vpn-ipv4-unicast" { return "L3VpnIpv4Unicast" }
    if yname == "l3vpn-ipv6-unicast" { return "L3VpnIpv6Unicast" }
    if yname == "l3vpn-ipv4-multicast" { return "L3VpnIpv4Multicast" }
    if yname == "l3vpn-ipv6-multicast" { return "L3VpnIpv6Multicast" }
    if yname == "l2vpn-vpls" { return "L2VpnVpls" }
    if yname == "l2vpn-evpn" { return "L2VpnEvpn" }
    return ""
}

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetSegmentPath() string {
    return "afi-safi" + "[afi-safi-name='" + fmt.Sprintf("%v", afiSafi.AfiSafiName) + "']"
}

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &afiSafi.Config
    }
    if childYangName == "state" {
        return &afiSafi.State
    }
    if childYangName == "graceful-restart" {
        return &afiSafi.GracefulRestart
    }
    if childYangName == "route-selection-options" {
        return &afiSafi.RouteSelectionOptions
    }
    if childYangName == "use-multiple-paths" {
        return &afiSafi.UseMultiplePaths
    }
    if childYangName == "apply-policy" {
        return &afiSafi.ApplyPolicy
    }
    if childYangName == "ipv4-unicast" {
        return &afiSafi.Ipv4Unicast
    }
    if childYangName == "ipv6-unicast" {
        return &afiSafi.Ipv6Unicast
    }
    if childYangName == "ipv4-labeled-unicast" {
        return &afiSafi.Ipv4LabeledUnicast
    }
    if childYangName == "ipv6-labeled-unicast" {
        return &afiSafi.Ipv6LabeledUnicast
    }
    if childYangName == "l3vpn-ipv4-unicast" {
        return &afiSafi.L3VpnIpv4Unicast
    }
    if childYangName == "l3vpn-ipv6-unicast" {
        return &afiSafi.L3VpnIpv6Unicast
    }
    if childYangName == "l3vpn-ipv4-multicast" {
        return &afiSafi.L3VpnIpv4Multicast
    }
    if childYangName == "l3vpn-ipv6-multicast" {
        return &afiSafi.L3VpnIpv6Multicast
    }
    if childYangName == "l2vpn-vpls" {
        return &afiSafi.L2VpnVpls
    }
    if childYangName == "l2vpn-evpn" {
        return &afiSafi.L2VpnEvpn
    }
    return nil
}

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &afiSafi.Config
    children["state"] = &afiSafi.State
    children["graceful-restart"] = &afiSafi.GracefulRestart
    children["route-selection-options"] = &afiSafi.RouteSelectionOptions
    children["use-multiple-paths"] = &afiSafi.UseMultiplePaths
    children["apply-policy"] = &afiSafi.ApplyPolicy
    children["ipv4-unicast"] = &afiSafi.Ipv4Unicast
    children["ipv6-unicast"] = &afiSafi.Ipv6Unicast
    children["ipv4-labeled-unicast"] = &afiSafi.Ipv4LabeledUnicast
    children["ipv6-labeled-unicast"] = &afiSafi.Ipv6LabeledUnicast
    children["l3vpn-ipv4-unicast"] = &afiSafi.L3VpnIpv4Unicast
    children["l3vpn-ipv6-unicast"] = &afiSafi.L3VpnIpv6Unicast
    children["l3vpn-ipv4-multicast"] = &afiSafi.L3VpnIpv4Multicast
    children["l3vpn-ipv6-multicast"] = &afiSafi.L3VpnIpv6Multicast
    children["l2vpn-vpls"] = &afiSafi.L2VpnVpls
    children["l2vpn-evpn"] = &afiSafi.L2VpnEvpn
    return children
}

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = afiSafi.AfiSafiName
    return leafs
}

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetBundleName() string { return "openconfig" }

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetYangName() string { return "afi-safi" }

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) SetParent(parent types.Entity) { afiSafi.parent = parent }

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetParent() types.Entity { return afiSafi.parent }

func (afiSafi *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi) GetParentYangName() string { return "afi-safis" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config
// Configuration parameters for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = config.AfiSafiName
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Config) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State
// State information relating to the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI,SAFI. The type is one of the following:
    // L2VPNEVPNL2VPNVPLSIPV4UNICASTL3VPNIPV6MULTICASTL3VPNIPV6UNICASTL3VPNIPV4UNICASTL3VPNIPV4MULTICASTIPV4LABELEDUNICASTIPV6UNICASTIPV6LABELEDUNICAST.
    AfiSafiName interface{}

    // This leaf indicates whether the IPv4 Unicast AFI,SAFI is enabled for the
    // neighbour or group. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetGoName(yname string) string {
    if yname == "afi-safi-name" { return "AfiSafiName" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi-safi-name"] = state.AfiSafiName
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_State) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart
// Parameters relating to BGP graceful-restart
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration options for BGP graceful-restart.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config

    // State information for BGP graceful-restart.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &gracefulRestart.Config
    }
    if childYangName == "state" {
        return &gracefulRestart.State
    }
    return nil
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &gracefulRestart.Config
    children["state"] = &gracefulRestart.State
    return children
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetBundleName() string { return "openconfig" }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config
// Configuration options for BGP graceful-restart
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_Config) GetParentYangName() string { return "graceful-restart" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State
// State information for BGP graceful-restart
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This leaf indicates whether graceful-restart is enabled for this AFI-SAFI.
    // The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_GracefulRestart_State) GetParentYangName() string { return "graceful-restart" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions
// Parameters relating to options for route selection
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to route selection options.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config

    // State information for the route selection options.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State
}

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetFilter() yfilter.YFilter { return routeSelectionOptions.YFilter }

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) SetFilter(yf yfilter.YFilter) { routeSelectionOptions.YFilter = yf }

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetSegmentPath() string {
    return "route-selection-options"
}

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &routeSelectionOptions.Config
    }
    if childYangName == "state" {
        return &routeSelectionOptions.State
    }
    return nil
}

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &routeSelectionOptions.Config
    children["state"] = &routeSelectionOptions.State
    return children
}

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetBundleName() string { return "openconfig" }

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetYangName() string { return "route-selection-options" }

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) SetParent(parent types.Entity) { routeSelectionOptions.parent = parent }

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetParent() types.Entity { return routeSelectionOptions.parent }

func (routeSelectionOptions *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config
// Configuration parameters relating to route selection
// options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetGoName(yname string) string {
    if yname == "always-compare-med" { return "AlwaysCompareMed" }
    if yname == "ignore-as-path-length" { return "IgnoreAsPathLength" }
    if yname == "external-compare-router-id" { return "ExternalCompareRouterId" }
    if yname == "advertise-inactive-routes" { return "AdvertiseInactiveRoutes" }
    if yname == "enable-aigp" { return "EnableAigp" }
    if yname == "ignore-next-hop-igp-metric" { return "IgnoreNextHopIgpMetric" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["always-compare-med"] = config.AlwaysCompareMed
    leafs["ignore-as-path-length"] = config.IgnoreAsPathLength
    leafs["external-compare-router-id"] = config.ExternalCompareRouterId
    leafs["advertise-inactive-routes"] = config.AdvertiseInactiveRoutes
    leafs["enable-aigp"] = config.EnableAigp
    leafs["ignore-next-hop-igp-metric"] = config.IgnoreNextHopIgpMetric
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_Config) GetParentYangName() string { return "route-selection-options" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State
// State information for the route selection options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetGoName(yname string) string {
    if yname == "always-compare-med" { return "AlwaysCompareMed" }
    if yname == "ignore-as-path-length" { return "IgnoreAsPathLength" }
    if yname == "external-compare-router-id" { return "ExternalCompareRouterId" }
    if yname == "advertise-inactive-routes" { return "AdvertiseInactiveRoutes" }
    if yname == "enable-aigp" { return "EnableAigp" }
    if yname == "ignore-next-hop-igp-metric" { return "IgnoreNextHopIgpMetric" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["always-compare-med"] = state.AlwaysCompareMed
    leafs["ignore-as-path-length"] = state.IgnoreAsPathLength
    leafs["external-compare-router-id"] = state.ExternalCompareRouterId
    leafs["advertise-inactive-routes"] = state.AdvertiseInactiveRoutes
    leafs["enable-aigp"] = state.EnableAigp
    leafs["ignore-next-hop-igp-metric"] = state.IgnoreNextHopIgpMetric
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_RouteSelectionOptions_State) GetParentYangName() string { return "route-selection-options" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths
// Parameters related to the use of multiple paths for the
// same NLRI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths struct {
    parent types.Entity
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

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetFilter() yfilter.YFilter { return useMultiplePaths.YFilter }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) SetFilter(yf yfilter.YFilter) { useMultiplePaths.YFilter = yf }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "ebgp" { return "Ebgp" }
    if yname == "ibgp" { return "Ibgp" }
    return ""
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetSegmentPath() string {
    return "use-multiple-paths"
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &useMultiplePaths.Config
    }
    if childYangName == "state" {
        return &useMultiplePaths.State
    }
    if childYangName == "ebgp" {
        return &useMultiplePaths.Ebgp
    }
    if childYangName == "ibgp" {
        return &useMultiplePaths.Ibgp
    }
    return nil
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &useMultiplePaths.Config
    children["state"] = &useMultiplePaths.State
    children["ebgp"] = &useMultiplePaths.Ebgp
    children["ibgp"] = &useMultiplePaths.Ibgp
    return children
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetBundleName() string { return "openconfig" }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetYangName() string { return "use-multiple-paths" }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) SetParent(parent types.Entity) { useMultiplePaths.parent = parent }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetParent() types.Entity { return useMultiplePaths.parent }

func (useMultiplePaths *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config
// Configuration parameters relating to multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Config) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State
// State parameters relating to multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether the use of multiple paths for the same NLRI is enabled for the
    // neighbor. This value is overridden by any more specific configuration
    // value. The type is bool. The default value is false.
    Enabled interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_State) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp
// Multipath parameters for eBGP
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to eBGP multipath.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config

    // State information relating to eBGP multipath.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
}

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetFilter() yfilter.YFilter { return ebgp.YFilter }

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) SetFilter(yf yfilter.YFilter) { ebgp.YFilter = yf }

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetSegmentPath() string {
    return "ebgp"
}

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ebgp.Config
    }
    if childYangName == "state" {
        return &ebgp.State
    }
    return nil
}

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ebgp.Config
    children["state"] = &ebgp.State
    return children
}

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetBundleName() string { return "openconfig" }

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetYangName() string { return "ebgp" }

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) SetParent(parent types.Entity) { ebgp.parent = parent }

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetParent() types.Entity { return ebgp.parent }

func (ebgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config
// Configuration parameters relating to eBGP multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = config.AllowMultipleAs
    leafs["maximum-paths"] = config.MaximumPaths
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_Config) GetParentYangName() string { return "ebgp" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State
// State information relating to eBGP multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetGoName(yname string) string {
    if yname == "allow-multiple-as" { return "AllowMultipleAs" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-multiple-as"] = state.AllowMultipleAs
    leafs["maximum-paths"] = state.MaximumPaths
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ebgp_State) GetParentYangName() string { return "ebgp" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp
// Multipath parameters for iBGP
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to iBGP multipath.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config

    // State information relating to iBGP multipath.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State
}

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetFilter() yfilter.YFilter { return ibgp.YFilter }

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) SetFilter(yf yfilter.YFilter) { ibgp.YFilter = yf }

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetSegmentPath() string {
    return "ibgp"
}

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ibgp.Config
    }
    if childYangName == "state" {
        return &ibgp.State
    }
    return nil
}

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ibgp.Config
    children["state"] = &ibgp.State
    return children
}

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetBundleName() string { return "openconfig" }

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetYangName() string { return "ibgp" }

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) SetParent(parent types.Entity) { ibgp.parent = parent }

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetParent() types.Entity { return ibgp.parent }

func (ibgp *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp) GetParentYangName() string { return "use-multiple-paths" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config
// Configuration parameters relating to iBGP multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetGoName(yname string) string {
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-paths"] = config.MaximumPaths
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_Config) GetParentYangName() string { return "ibgp" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State
// State information relating to iBGP multipath
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of parallel paths to consider when using iBGP multipath. The
    // default is to use a single path. The type is interface{} with range:
    // 0..4294967295. The default value is 1.
    MaximumPaths interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetGoName(yname string) string {
    if yname == "maximum-paths" { return "MaximumPaths" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-paths"] = state.MaximumPaths
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_UseMultiplePaths_Ibgp_State) GetParentYangName() string { return "ibgp" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy
// Anchor point for routing policies in the model.
// Import and export policies are with respect to the local
// routing table, i.e., export (send) and import (receive),
// depending on the context.
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy configuration data.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config

    // Operational state for routing policy.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetFilter() yfilter.YFilter { return applyPolicy.YFilter }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) SetFilter(yf yfilter.YFilter) { applyPolicy.YFilter = yf }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetSegmentPath() string {
    return "apply-policy"
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &applyPolicy.Config
    }
    if childYangName == "state" {
        return &applyPolicy.State
    }
    return nil
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &applyPolicy.Config
    children["state"] = &applyPolicy.State
    return children
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetBundleName() string { return "openconfig" }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetYangName() string { return "apply-policy" }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) SetParent(parent types.Entity) { applyPolicy.parent = parent }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetParent() types.Entity { return applyPolicy.parent }

func (applyPolicy *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config
// Policy configuration data.
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = config.ImportPolicy
    leafs["default-import-policy"] = config.DefaultImportPolicy
    leafs["export-policy"] = config.ExportPolicy
    leafs["default-export-policy"] = config.DefaultExportPolicy
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_Config) GetParentYangName() string { return "apply-policy" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State
// Operational state for routing policy
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetGoName(yname string) string {
    if yname == "import-policy" { return "ImportPolicy" }
    if yname == "default-import-policy" { return "DefaultImportPolicy" }
    if yname == "export-policy" { return "ExportPolicy" }
    if yname == "default-export-policy" { return "DefaultExportPolicy" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["import-policy"] = state.ImportPolicy
    leafs["default-import-policy"] = state.DefaultImportPolicy
    leafs["export-policy"] = state.ExportPolicy
    leafs["default-export-policy"] = state.DefaultExportPolicy
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_ApplyPolicy_State) GetParentYangName() string { return "apply-policy" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast
// IPv4 unicast configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State
}

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetFilter() yfilter.YFilter { return ipv4Unicast.YFilter }

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) SetFilter(yf yfilter.YFilter) { ipv4Unicast.YFilter = yf }

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetSegmentPath() string {
    return "ipv4-unicast"
}

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv4Unicast.PrefixLimit
    }
    if childYangName == "config" {
        return &ipv4Unicast.Config
    }
    if childYangName == "state" {
        return &ipv4Unicast.State
    }
    return nil
}

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv4Unicast.PrefixLimit
    children["config"] = &ipv4Unicast.Config
    children["state"] = &ipv4Unicast.State
    return children
}

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetBundleName() string { return "openconfig" }

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetYangName() string { return "ipv4-unicast" }

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) SetParent(parent types.Entity) { ipv4Unicast.parent = parent }

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetParent() types.Entity { return ipv4Unicast.parent }

func (ipv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit) GetParentYangName() string { return "ipv4-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = config.SendDefaultRoute
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_Config) GetParentYangName() string { return "ipv4-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = state.SendDefaultRoute
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4Unicast_State) GetParentYangName() string { return "ipv4-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast
// IPv6 unicast configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit

    // Configuration parameters for common IPv4 and IPv6 unicast AFI-SAFI options.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config

    // State information for common IPv4 and IPv6 unicast parameters.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State
}

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetFilter() yfilter.YFilter { return ipv6Unicast.YFilter }

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) SetFilter(yf yfilter.YFilter) { ipv6Unicast.YFilter = yf }

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetSegmentPath() string {
    return "ipv6-unicast"
}

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv6Unicast.PrefixLimit
    }
    if childYangName == "config" {
        return &ipv6Unicast.Config
    }
    if childYangName == "state" {
        return &ipv6Unicast.State
    }
    return nil
}

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv6Unicast.PrefixLimit
    children["config"] = &ipv6Unicast.Config
    children["state"] = &ipv6Unicast.State
    return children
}

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetBundleName() string { return "openconfig" }

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetYangName() string { return "ipv6-unicast" }

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) SetParent(parent types.Entity) { ipv6Unicast.parent = parent }

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetParent() types.Entity { return ipv6Unicast.parent }

func (ipv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit) GetParentYangName() string { return "ipv6-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config
// Configuration parameters for common IPv4 and IPv6 unicast
// AFI-SAFI options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = config.SendDefaultRoute
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_Config) GetParentYangName() string { return "ipv6-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State
// State information for common IPv4 and IPv6 unicast
// parameters
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set to true, send the default-route to the neighbour(s). The type is
    // bool. The default value is false.
    SendDefaultRoute interface{}
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetGoName(yname string) string {
    if yname == "send-default-route" { return "SendDefaultRoute" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["send-default-route"] = state.SendDefaultRoute
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6Unicast_State) GetParentYangName() string { return "ipv6-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast
// IPv4 Labeled Unicast configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetFilter() yfilter.YFilter { return ipv4LabeledUnicast.YFilter }

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) SetFilter(yf yfilter.YFilter) { ipv4LabeledUnicast.YFilter = yf }

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetSegmentPath() string {
    return "ipv4-labeled-unicast"
}

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv4LabeledUnicast.PrefixLimit
    }
    return nil
}

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv4LabeledUnicast.PrefixLimit
    return children
}

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetBundleName() string { return "openconfig" }

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetYangName() string { return "ipv4-labeled-unicast" }

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) SetParent(parent types.Entity) { ipv4LabeledUnicast.parent = parent }

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetParent() types.Entity { return ipv4LabeledUnicast.parent }

func (ipv4LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) GetParentYangName() string { return "ipv4-labeled-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv4LabeledUnicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast
// IPv6 Labeled Unicast configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetFilter() yfilter.YFilter { return ipv6LabeledUnicast.YFilter }

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) SetFilter(yf yfilter.YFilter) { ipv6LabeledUnicast.YFilter = yf }

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetSegmentPath() string {
    return "ipv6-labeled-unicast"
}

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &ipv6LabeledUnicast.PrefixLimit
    }
    return nil
}

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &ipv6LabeledUnicast.PrefixLimit
    return children
}

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetBundleName() string { return "openconfig" }

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetYangName() string { return "ipv6-labeled-unicast" }

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) SetParent(parent types.Entity) { ipv6LabeledUnicast.parent = parent }

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetParent() types.Entity { return ipv6LabeledUnicast.parent }

func (ipv6LabeledUnicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) GetParentYangName() string { return "ipv6-labeled-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_Ipv6LabeledUnicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast
// Unicast IPv4 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetFilter() yfilter.YFilter { return l3VpnIpv4Unicast.YFilter }

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv4Unicast.YFilter = yf }

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetSegmentPath() string {
    return "l3vpn-ipv4-unicast"
}

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv4Unicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv4Unicast.PrefixLimit
    return children
}

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetYangName() string { return "l3vpn-ipv4-unicast" }

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) SetParent(parent types.Entity) { l3VpnIpv4Unicast.parent = parent }

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetParent() types.Entity { return l3VpnIpv4Unicast.parent }

func (l3VpnIpv4Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv4-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast
// Unicast IPv6 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetFilter() yfilter.YFilter { return l3VpnIpv6Unicast.YFilter }

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv6Unicast.YFilter = yf }

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetSegmentPath() string {
    return "l3vpn-ipv6-unicast"
}

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv6Unicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv6Unicast.PrefixLimit
    return children
}

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetYangName() string { return "l3vpn-ipv6-unicast" }

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) SetParent(parent types.Entity) { l3VpnIpv6Unicast.parent = parent }

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetParent() types.Entity { return l3VpnIpv6Unicast.parent }

func (l3VpnIpv6Unicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv6-unicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Unicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast
// Multicast IPv4 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetFilter() yfilter.YFilter { return l3VpnIpv4Multicast.YFilter }

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv4Multicast.YFilter = yf }

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetSegmentPath() string {
    return "l3vpn-ipv4-multicast"
}

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv4Multicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv4Multicast.PrefixLimit
    return children
}

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetYangName() string { return "l3vpn-ipv4-multicast" }

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) SetParent(parent types.Entity) { l3VpnIpv4Multicast.parent = parent }

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetParent() types.Entity { return l3VpnIpv4Multicast.parent }

func (l3VpnIpv4Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv4-multicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv4Multicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast
// Multicast IPv6 L3VPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetFilter() yfilter.YFilter { return l3VpnIpv6Multicast.YFilter }

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) SetFilter(yf yfilter.YFilter) { l3VpnIpv6Multicast.YFilter = yf }

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetSegmentPath() string {
    return "l3vpn-ipv6-multicast"
}

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l3VpnIpv6Multicast.PrefixLimit
    }
    return nil
}

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l3VpnIpv6Multicast.PrefixLimit
    return children
}

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetBundleName() string { return "openconfig" }

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetYangName() string { return "l3vpn-ipv6-multicast" }

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) SetParent(parent types.Entity) { l3VpnIpv6Multicast.parent = parent }

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetParent() types.Entity { return l3VpnIpv6Multicast.parent }

func (l3VpnIpv6Multicast *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) GetParentYangName() string { return "l3vpn-ipv6-multicast" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L3VpnIpv6Multicast_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls
// BGP-signalled VPLS configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
}

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetFilter() yfilter.YFilter { return l2VpnVpls.YFilter }

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) SetFilter(yf yfilter.YFilter) { l2VpnVpls.YFilter = yf }

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetSegmentPath() string {
    return "l2vpn-vpls"
}

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l2VpnVpls.PrefixLimit
    }
    return nil
}

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l2VpnVpls.PrefixLimit
    return children
}

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetBundleName() string { return "openconfig" }

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetYangName() string { return "l2vpn-vpls" }

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) SetParent(parent types.Entity) { l2VpnVpls.parent = parent }

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetParent() types.Entity { return l2VpnVpls.parent }

func (l2VpnVpls *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit) GetParentYangName() string { return "l2vpn-vpls" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnVpls_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn
// BGP EVPN configuration options
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum number of prefixes that will be accepted from a peer.
    PrefixLimit Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
}

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetFilter() yfilter.YFilter { return l2VpnEvpn.YFilter }

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) SetFilter(yf yfilter.YFilter) { l2VpnEvpn.YFilter = yf }

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetGoName(yname string) string {
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetSegmentPath() string {
    return "l2vpn-evpn"
}

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-limit" {
        return &l2VpnEvpn.PrefixLimit
    }
    return nil
}

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-limit"] = &l2VpnEvpn.PrefixLimit
    return children
}

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetBundleName() string { return "openconfig" }

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetYangName() string { return "l2vpn-evpn" }

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) SetParent(parent types.Entity) { l2VpnEvpn.parent = parent }

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetParent() types.Entity { return l2VpnEvpn.parent }

func (l2VpnEvpn *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn) GetParentYangName() string { return "afi-safi" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit
// Configure the maximum number of prefixes that will be
// accepted from a peer
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the prefix limit for the AFI-SAFI.
    Config Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config

    // State information relating to the prefix-limit for the AFI-SAFI.
    State Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetFilter() yfilter.YFilter { return prefixLimit.YFilter }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) SetFilter(yf yfilter.YFilter) { prefixLimit.YFilter = yf }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetSegmentPath() string {
    return "prefix-limit"
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixLimit.Config
    }
    if childYangName == "state" {
        return &prefixLimit.State
    }
    return nil
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixLimit.Config
    children["state"] = &prefixLimit.State
    return children
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetBundleName() string { return "openconfig" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetYangName() string { return "prefix-limit" }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) SetParent(parent types.Entity) { prefixLimit.parent = parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetParent() types.Entity { return prefixLimit.parent }

func (prefixLimit *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit) GetParentYangName() string { return "l2vpn-evpn" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config
// Configuration parameters relating to the prefix
// limit for the AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config struct {
    parent types.Entity
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

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetSegmentPath() string {
    return "config"
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = config.MaxPrefixes
    leafs["shutdown-threshold-pct"] = config.ShutdownThresholdPct
    leafs["restart-timer"] = config.RestartTimer
    return leafs
}

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetBundleName() string { return "openconfig" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetYangName() string { return "config" }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetParent() types.Entity { return config.parent }

func (config *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_Config) GetParentYangName() string { return "prefix-limit" }

// Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State
// State information relating to the prefix-limit for the
// AFI-SAFI
type Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State struct {
    parent types.Entity
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

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetGoName(yname string) string {
    if yname == "max-prefixes" { return "MaxPrefixes" }
    if yname == "shutdown-threshold-pct" { return "ShutdownThresholdPct" }
    if yname == "restart-timer" { return "RestartTimer" }
    return ""
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetSegmentPath() string {
    return "state"
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-prefixes"] = state.MaxPrefixes
    leafs["shutdown-threshold-pct"] = state.ShutdownThresholdPct
    leafs["restart-timer"] = state.RestartTimer
    return leafs
}

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetBundleName() string { return "openconfig" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetYangName() string { return "state" }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetParent() types.Entity { return state.parent }

func (state *Bgp_PeerGroups_PeerGroup_AfiSafis_AfiSafi_L2VpnEvpn_PrefixLimit_State) GetParentYangName() string { return "prefix-limit" }

