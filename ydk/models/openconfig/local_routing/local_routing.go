// This module describes configuration and operational state data
// for routes that are locally generated, i.e., not created by
// dynamic routing protocols.  These include static routes, locally
// created aggregate routes for reducing the number of constituent
// routes that must be advertised, summary routes for IGPs, etc.
// This model expresses locally generated routes as generically as
// possible, avoiding configuration of protocol-specific attributes
// at the time of route creation.  This is primarily to avoid
// assumptions about how underlying router implementations handle
// route attributes in various routing table data structures they
// maintain.  Hence, the definition of locally generated routes
// essentially creates 'bare' routes that do not have any protocol-
// specific attributes.
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

type DROP struct {
}

func (id DROP) String() string {
	return "openconfig-local-routing:DROP"
}

type LOCALDEFINEDNEXTHOP struct {
}

func (id LOCALDEFINEDNEXTHOP) String() string {
	return "openconfig-local-routing:LOCAL_DEFINED_NEXT_HOP"
}

type LOCALLINK struct {
}

func (id LOCALLINK) String() string {
	return "openconfig-local-routing:LOCAL_LINK"
}

// LocalRoutes
// Top-level container for local routes
type LocalRoutes struct {
    parent types.Entity
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

func (localRoutes *LocalRoutes) GetFilter() yfilter.YFilter { return localRoutes.YFilter }

func (localRoutes *LocalRoutes) SetFilter(yf yfilter.YFilter) { localRoutes.YFilter = yf }

func (localRoutes *LocalRoutes) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "static-routes" { return "StaticRoutes" }
    if yname == "local-aggregates" { return "LocalAggregates" }
    return ""
}

func (localRoutes *LocalRoutes) GetSegmentPath() string {
    return "openconfig-local-routing:local-routes"
}

func (localRoutes *LocalRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &localRoutes.Config
    }
    if childYangName == "state" {
        return &localRoutes.State
    }
    if childYangName == "static-routes" {
        return &localRoutes.StaticRoutes
    }
    if childYangName == "local-aggregates" {
        return &localRoutes.LocalAggregates
    }
    return nil
}

func (localRoutes *LocalRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &localRoutes.Config
    children["state"] = &localRoutes.State
    children["static-routes"] = &localRoutes.StaticRoutes
    children["local-aggregates"] = &localRoutes.LocalAggregates
    return children
}

func (localRoutes *LocalRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localRoutes *LocalRoutes) GetBundleName() string { return "openconfig" }

func (localRoutes *LocalRoutes) GetYangName() string { return "local-routes" }

func (localRoutes *LocalRoutes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (localRoutes *LocalRoutes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (localRoutes *LocalRoutes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (localRoutes *LocalRoutes) SetParent(parent types.Entity) { localRoutes.parent = parent }

func (localRoutes *LocalRoutes) GetParent() types.Entity { return localRoutes.parent }

func (localRoutes *LocalRoutes) GetParentYangName() string { return "openconfig-local-routing" }

// LocalRoutes_Config
// Configuration data for locally defined routes
type LocalRoutes_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (config *LocalRoutes_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *LocalRoutes_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *LocalRoutes_Config) GetGoName(yname string) string {
    return ""
}

func (config *LocalRoutes_Config) GetSegmentPath() string {
    return "config"
}

func (config *LocalRoutes_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *LocalRoutes_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *LocalRoutes_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (config *LocalRoutes_Config) GetBundleName() string { return "openconfig" }

func (config *LocalRoutes_Config) GetYangName() string { return "config" }

func (config *LocalRoutes_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *LocalRoutes_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *LocalRoutes_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *LocalRoutes_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *LocalRoutes_Config) GetParent() types.Entity { return config.parent }

func (config *LocalRoutes_Config) GetParentYangName() string { return "local-routes" }

// LocalRoutes_State
// Operational state data for locally defined routes
type LocalRoutes_State struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (state *LocalRoutes_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *LocalRoutes_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *LocalRoutes_State) GetGoName(yname string) string {
    return ""
}

func (state *LocalRoutes_State) GetSegmentPath() string {
    return "state"
}

func (state *LocalRoutes_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *LocalRoutes_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *LocalRoutes_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (state *LocalRoutes_State) GetBundleName() string { return "openconfig" }

func (state *LocalRoutes_State) GetYangName() string { return "state" }

func (state *LocalRoutes_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *LocalRoutes_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *LocalRoutes_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *LocalRoutes_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *LocalRoutes_State) GetParent() types.Entity { return state.parent }

func (state *LocalRoutes_State) GetParentYangName() string { return "local-routes" }

// LocalRoutes_StaticRoutes
// Enclosing container for the list of static routes
type LocalRoutes_StaticRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of locally configured static routes. The type is slice of
    // LocalRoutes_StaticRoutes_Static.
    Static []LocalRoutes_StaticRoutes_Static
}

func (staticRoutes *LocalRoutes_StaticRoutes) GetFilter() yfilter.YFilter { return staticRoutes.YFilter }

func (staticRoutes *LocalRoutes_StaticRoutes) SetFilter(yf yfilter.YFilter) { staticRoutes.YFilter = yf }

func (staticRoutes *LocalRoutes_StaticRoutes) GetGoName(yname string) string {
    if yname == "static" { return "Static" }
    return ""
}

func (staticRoutes *LocalRoutes_StaticRoutes) GetSegmentPath() string {
    return "static-routes"
}

func (staticRoutes *LocalRoutes_StaticRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static" {
        for _, c := range staticRoutes.Static {
            if staticRoutes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LocalRoutes_StaticRoutes_Static{}
        staticRoutes.Static = append(staticRoutes.Static, child)
        return &staticRoutes.Static[len(staticRoutes.Static)-1]
    }
    return nil
}

func (staticRoutes *LocalRoutes_StaticRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticRoutes.Static {
        children[staticRoutes.Static[i].GetSegmentPath()] = &staticRoutes.Static[i]
    }
    return children
}

func (staticRoutes *LocalRoutes_StaticRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticRoutes *LocalRoutes_StaticRoutes) GetBundleName() string { return "openconfig" }

func (staticRoutes *LocalRoutes_StaticRoutes) GetYangName() string { return "static-routes" }

func (staticRoutes *LocalRoutes_StaticRoutes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (staticRoutes *LocalRoutes_StaticRoutes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (staticRoutes *LocalRoutes_StaticRoutes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (staticRoutes *LocalRoutes_StaticRoutes) SetParent(parent types.Entity) { staticRoutes.parent = parent }

func (staticRoutes *LocalRoutes_StaticRoutes) GetParent() types.Entity { return staticRoutes.parent }

func (staticRoutes *LocalRoutes_StaticRoutes) GetParentYangName() string { return "local-routes" }

// LocalRoutes_StaticRoutes_Static
// List of locally configured static routes
type LocalRoutes_StaticRoutes_Static struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the destination prefix list key. The
    // type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // Configuration data for static routes.
    Config LocalRoutes_StaticRoutes_Static_Config

    // Operational state data for static routes.
    State LocalRoutes_StaticRoutes_Static_State

    // Configuration and state parameters relating to the next-hops that are to be
    // utilised for the static route being specified.
    NextHops LocalRoutes_StaticRoutes_Static_NextHops
}

func (static *LocalRoutes_StaticRoutes_Static) GetFilter() yfilter.YFilter { return static.YFilter }

func (static *LocalRoutes_StaticRoutes_Static) SetFilter(yf yfilter.YFilter) { static.YFilter = yf }

func (static *LocalRoutes_StaticRoutes_Static) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "next-hops" { return "NextHops" }
    return ""
}

func (static *LocalRoutes_StaticRoutes_Static) GetSegmentPath() string {
    return "static" + "[prefix='" + fmt.Sprintf("%v", static.Prefix) + "']"
}

func (static *LocalRoutes_StaticRoutes_Static) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &static.Config
    }
    if childYangName == "state" {
        return &static.State
    }
    if childYangName == "next-hops" {
        return &static.NextHops
    }
    return nil
}

func (static *LocalRoutes_StaticRoutes_Static) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &static.Config
    children["state"] = &static.State
    children["next-hops"] = &static.NextHops
    return children
}

func (static *LocalRoutes_StaticRoutes_Static) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = static.Prefix
    return leafs
}

func (static *LocalRoutes_StaticRoutes_Static) GetBundleName() string { return "openconfig" }

func (static *LocalRoutes_StaticRoutes_Static) GetYangName() string { return "static" }

func (static *LocalRoutes_StaticRoutes_Static) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (static *LocalRoutes_StaticRoutes_Static) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (static *LocalRoutes_StaticRoutes_Static) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (static *LocalRoutes_StaticRoutes_Static) SetParent(parent types.Entity) { static.parent = parent }

func (static *LocalRoutes_StaticRoutes_Static) GetParent() types.Entity { return static.parent }

func (static *LocalRoutes_StaticRoutes_Static) GetParentYangName() string { return "static-routes" }

// LocalRoutes_StaticRoutes_Static_Config
// Configuration data for static routes
type LocalRoutes_StaticRoutes_Static_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination prefix for the static route, either IPv4 or IPv6. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // Set a generic tag value on the route. This tag can be used for filtering
    // routes that are distributed to other routing protocols. The type is one of
    // the following types: int with range: 0..4294967295, or string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    SetTag interface{}
}

func (config *LocalRoutes_StaticRoutes_Static_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *LocalRoutes_StaticRoutes_Static_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *LocalRoutes_StaticRoutes_Static_Config) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "set-tag" { return "SetTag" }
    return ""
}

func (config *LocalRoutes_StaticRoutes_Static_Config) GetSegmentPath() string {
    return "config"
}

func (config *LocalRoutes_StaticRoutes_Static_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *LocalRoutes_StaticRoutes_Static_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *LocalRoutes_StaticRoutes_Static_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = config.Prefix
    leafs["set-tag"] = config.SetTag
    return leafs
}

func (config *LocalRoutes_StaticRoutes_Static_Config) GetBundleName() string { return "openconfig" }

func (config *LocalRoutes_StaticRoutes_Static_Config) GetYangName() string { return "config" }

func (config *LocalRoutes_StaticRoutes_Static_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *LocalRoutes_StaticRoutes_Static_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *LocalRoutes_StaticRoutes_Static_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *LocalRoutes_StaticRoutes_Static_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *LocalRoutes_StaticRoutes_Static_Config) GetParent() types.Entity { return config.parent }

func (config *LocalRoutes_StaticRoutes_Static_Config) GetParentYangName() string { return "static" }

// LocalRoutes_StaticRoutes_Static_State
// Operational state data for static routes
type LocalRoutes_StaticRoutes_Static_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination prefix for the static route, either IPv4 or IPv6. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // Set a generic tag value on the route. This tag can be used for filtering
    // routes that are distributed to other routing protocols. The type is one of
    // the following types: int with range: 0..4294967295, or string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    SetTag interface{}
}

func (state *LocalRoutes_StaticRoutes_Static_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *LocalRoutes_StaticRoutes_Static_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *LocalRoutes_StaticRoutes_Static_State) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "set-tag" { return "SetTag" }
    return ""
}

func (state *LocalRoutes_StaticRoutes_Static_State) GetSegmentPath() string {
    return "state"
}

func (state *LocalRoutes_StaticRoutes_Static_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *LocalRoutes_StaticRoutes_Static_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *LocalRoutes_StaticRoutes_Static_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = state.Prefix
    leafs["set-tag"] = state.SetTag
    return leafs
}

func (state *LocalRoutes_StaticRoutes_Static_State) GetBundleName() string { return "openconfig" }

func (state *LocalRoutes_StaticRoutes_Static_State) GetYangName() string { return "state" }

func (state *LocalRoutes_StaticRoutes_Static_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *LocalRoutes_StaticRoutes_Static_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *LocalRoutes_StaticRoutes_Static_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *LocalRoutes_StaticRoutes_Static_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *LocalRoutes_StaticRoutes_Static_State) GetParent() types.Entity { return state.parent }

func (state *LocalRoutes_StaticRoutes_Static_State) GetParentYangName() string { return "static" }

// LocalRoutes_StaticRoutes_Static_NextHops
// Configuration and state parameters relating to the
// next-hops that are to be utilised for the static
// route being specified
type LocalRoutes_StaticRoutes_Static_NextHops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of next-hops to be utilised for the static route being specified.
    // The type is slice of LocalRoutes_StaticRoutes_Static_NextHops_NextHop.
    NextHop []LocalRoutes_StaticRoutes_Static_NextHops_NextHop
}

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetFilter() yfilter.YFilter { return nextHops.YFilter }

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) SetFilter(yf yfilter.YFilter) { nextHops.YFilter = yf }

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetSegmentPath() string {
    return "next-hops"
}

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        for _, c := range nextHops.NextHop {
            if nextHops.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LocalRoutes_StaticRoutes_Static_NextHops_NextHop{}
        nextHops.NextHop = append(nextHops.NextHop, child)
        return &nextHops.NextHop[len(nextHops.NextHop)-1]
    }
    return nil
}

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nextHops.NextHop {
        children[nextHops.NextHop[i].GetSegmentPath()] = &nextHops.NextHop[i]
    }
    return children
}

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetBundleName() string { return "openconfig" }

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetYangName() string { return "next-hops" }

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) SetParent(parent types.Entity) { nextHops.parent = parent }

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetParent() types.Entity { return nextHops.parent }

func (nextHops *LocalRoutes_StaticRoutes_Static_NextHops) GetParentYangName() string { return "static" }

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop
// A list of next-hops to be utilised for the static
// route being specified.
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interface-ref" { return "InterfaceRef" }
    return ""
}

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetSegmentPath() string {
    return "next-hop" + "[index='" + fmt.Sprintf("%v", nextHop.Index) + "']"
}

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &nextHop.Config
    }
    if childYangName == "state" {
        return &nextHop.State
    }
    if childYangName == "interface-ref" {
        return &nextHop.InterfaceRef
    }
    return nil
}

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &nextHop.Config
    children["state"] = &nextHop.State
    children["interface-ref"] = &nextHop.InterfaceRef
    return children
}

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = nextHop.Index
    return leafs
}

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetBundleName() string { return "openconfig" }

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *LocalRoutes_StaticRoutes_Static_NextHops_NextHop) GetParentYangName() string { return "next-hops" }

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config
// Configuration parameters relating to the next-hop
// entry
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
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

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "metric" { return "Metric" }
    if yname == "recurse" { return "Recurse" }
    return ""
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetSegmentPath() string {
    return "config"
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = config.Index
    leafs["next-hop"] = config.NextHop
    leafs["metric"] = config.Metric
    leafs["recurse"] = config.Recurse
    return leafs
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetBundleName() string { return "openconfig" }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetYangName() string { return "config" }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetParent() types.Entity { return config.parent }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_Config) GetParentYangName() string { return "next-hop" }

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State
// Operational state parameters relating to the
// next-hop entry
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
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

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "metric" { return "Metric" }
    if yname == "recurse" { return "Recurse" }
    return ""
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetSegmentPath() string {
    return "state"
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = state.Index
    leafs["next-hop"] = state.NextHop
    leafs["metric"] = state.Metric
    leafs["recurse"] = state.Recurse
    return leafs
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetBundleName() string { return "openconfig" }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetYangName() string { return "state" }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetParent() types.Entity { return state.parent }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_State) GetParentYangName() string { return "next-hop" }

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef
// Reference to an interface or subinterface
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config

    // Operational state for interface-ref.
    State LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State
}

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetFilter() yfilter.YFilter { return interfaceRef.YFilter }

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) SetFilter(yf yfilter.YFilter) { interfaceRef.YFilter = yf }

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetSegmentPath() string {
    return "interface-ref"
}

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaceRef.Config
    }
    if childYangName == "state" {
        return &interfaceRef.State
    }
    return nil
}

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaceRef.Config
    children["state"] = &interfaceRef.State
    return children
}

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetBundleName() string { return "openconfig" }

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetYangName() string { return "interface-ref" }

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) SetParent(parent types.Entity) { interfaceRef.parent = parent }

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetParent() types.Entity { return interfaceRef.parent }

func (interfaceRef *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef) GetParentYangName() string { return "next-hop" }

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config
// Configured reference to interface / subinterface
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config struct {
    parent types.Entity
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

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetSegmentPath() string {
    return "config"
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = config.Interface
    leafs["subinterface"] = config.Subinterface
    return leafs
}

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetBundleName() string { return "openconfig" }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetYangName() string { return "config" }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetParent() types.Entity { return config.parent }

func (config *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_Config) GetParentYangName() string { return "interface-ref" }

// LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State
// Operational state for interface-ref
type LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State struct {
    parent types.Entity
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

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetSegmentPath() string {
    return "state"
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    leafs["subinterface"] = state.Subinterface
    return leafs
}

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetBundleName() string { return "openconfig" }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetYangName() string { return "state" }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetParent() types.Entity { return state.parent }

func (state *LocalRoutes_StaticRoutes_Static_NextHops_NextHop_InterfaceRef_State) GetParentYangName() string { return "interface-ref" }

// LocalRoutes_LocalAggregates
// Enclosing container for locally-defined aggregate
// routes
type LocalRoutes_LocalAggregates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of aggregates. The type is slice of
    // LocalRoutes_LocalAggregates_Aggregate.
    Aggregate []LocalRoutes_LocalAggregates_Aggregate
}

func (localAggregates *LocalRoutes_LocalAggregates) GetFilter() yfilter.YFilter { return localAggregates.YFilter }

func (localAggregates *LocalRoutes_LocalAggregates) SetFilter(yf yfilter.YFilter) { localAggregates.YFilter = yf }

func (localAggregates *LocalRoutes_LocalAggregates) GetGoName(yname string) string {
    if yname == "aggregate" { return "Aggregate" }
    return ""
}

func (localAggregates *LocalRoutes_LocalAggregates) GetSegmentPath() string {
    return "local-aggregates"
}

func (localAggregates *LocalRoutes_LocalAggregates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregate" {
        for _, c := range localAggregates.Aggregate {
            if localAggregates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := LocalRoutes_LocalAggregates_Aggregate{}
        localAggregates.Aggregate = append(localAggregates.Aggregate, child)
        return &localAggregates.Aggregate[len(localAggregates.Aggregate)-1]
    }
    return nil
}

func (localAggregates *LocalRoutes_LocalAggregates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localAggregates.Aggregate {
        children[localAggregates.Aggregate[i].GetSegmentPath()] = &localAggregates.Aggregate[i]
    }
    return children
}

func (localAggregates *LocalRoutes_LocalAggregates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localAggregates *LocalRoutes_LocalAggregates) GetBundleName() string { return "openconfig" }

func (localAggregates *LocalRoutes_LocalAggregates) GetYangName() string { return "local-aggregates" }

func (localAggregates *LocalRoutes_LocalAggregates) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (localAggregates *LocalRoutes_LocalAggregates) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (localAggregates *LocalRoutes_LocalAggregates) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (localAggregates *LocalRoutes_LocalAggregates) SetParent(parent types.Entity) { localAggregates.parent = parent }

func (localAggregates *LocalRoutes_LocalAggregates) GetParent() types.Entity { return localAggregates.parent }

func (localAggregates *LocalRoutes_LocalAggregates) GetParentYangName() string { return "local-routes" }

// LocalRoutes_LocalAggregates_Aggregate
// List of aggregates
type LocalRoutes_LocalAggregates_Aggregate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the configured prefix for this
    // aggregate. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // Configuration data for aggregate advertisements.
    Config LocalRoutes_LocalAggregates_Aggregate_Config

    // Operational state data for aggregate advertisements.
    State LocalRoutes_LocalAggregates_Aggregate_State
}

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetFilter() yfilter.YFilter { return aggregate.YFilter }

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) SetFilter(yf yfilter.YFilter) { aggregate.YFilter = yf }

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetSegmentPath() string {
    return "aggregate" + "[prefix='" + fmt.Sprintf("%v", aggregate.Prefix) + "']"
}

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &aggregate.Config
    }
    if childYangName == "state" {
        return &aggregate.State
    }
    return nil
}

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &aggregate.Config
    children["state"] = &aggregate.State
    return children
}

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = aggregate.Prefix
    return leafs
}

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetBundleName() string { return "openconfig" }

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetYangName() string { return "aggregate" }

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) SetParent(parent types.Entity) { aggregate.parent = parent }

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetParent() types.Entity { return aggregate.parent }

func (aggregate *LocalRoutes_LocalAggregates_Aggregate) GetParentYangName() string { return "local-aggregates" }

// LocalRoutes_LocalAggregates_Aggregate_Config
// Configuration data for aggregate advertisements
type LocalRoutes_LocalAggregates_Aggregate_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Aggregate prefix to be advertised. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
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
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    SetTag interface{}
}

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "discard" { return "Discard" }
    if yname == "set-tag" { return "SetTag" }
    return ""
}

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetSegmentPath() string {
    return "config"
}

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = config.Prefix
    leafs["discard"] = config.Discard
    leafs["set-tag"] = config.SetTag
    return leafs
}

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetBundleName() string { return "openconfig" }

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetYangName() string { return "config" }

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetParent() types.Entity { return config.parent }

func (config *LocalRoutes_LocalAggregates_Aggregate_Config) GetParentYangName() string { return "aggregate" }

// LocalRoutes_LocalAggregates_Aggregate_State
// Operational state data for aggregate
// advertisements
type LocalRoutes_LocalAggregates_Aggregate_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Aggregate prefix to be advertised. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
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
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    SetTag interface{}
}

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *LocalRoutes_LocalAggregates_Aggregate_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "discard" { return "Discard" }
    if yname == "set-tag" { return "SetTag" }
    return ""
}

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetSegmentPath() string {
    return "state"
}

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = state.Prefix
    leafs["discard"] = state.Discard
    leafs["set-tag"] = state.SetTag
    return leafs
}

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetBundleName() string { return "openconfig" }

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetYangName() string { return "state" }

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *LocalRoutes_LocalAggregates_Aggregate_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetParent() types.Entity { return state.parent }

func (state *LocalRoutes_LocalAggregates_Aggregate_State) GetParentYangName() string { return "aggregate" }

