// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-rip package operational data.
// 
// This module contains definitions
// for the following management objects:
//   rip: RIP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_rip_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_rip_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-rip-oper rip}", reflect.TypeOf(Rip{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-rip-oper:rip", reflect.TypeOf(Rip{}))
}

// RipRouteOrigin represents Rip route origin
type RipRouteOrigin string

const (
    // configured 'network'
    RipRouteOrigin_rip_rt_org_runover RipRouteOrigin = "rip-rt-org-runover"

    // redistributed
    RipRouteOrigin_rip_rt_org_redist RipRouteOrigin = "rip-rt-org-redist"

    // auto summary
    RipRouteOrigin_rip_rt_org_auto_summary RipRouteOrigin = "rip-rt-org-auto-summary"

    // learned via RIP
    RipRouteOrigin_rip_rt_org_rip RipRouteOrigin = "rip-rt-org-rip"

    // interface summary-address
    RipRouteOrigin_rip_rt_org_intsummary RipRouteOrigin = "rip-rt-org-intsummary"

    // route stay in for triggered rip
    RipRouteOrigin_rip_rt_org_unused RipRouteOrigin = "rip-rt-org-unused"
)

// InterfaceState represents Interface state
type InterfaceState string

const (
    // Interface does not exist
    InterfaceState_interface_none InterfaceState = "interface-none"

    // Interface exists but IP is down
    InterfaceState_interface_down InterfaceState = "interface-down"

    // Interface exists and IP is up
    InterfaceState_interface_up InterfaceState = "interface-up"

    // Unknown state
    InterfaceState_interface_unknown InterfaceState = "interface-unknown"
)

// Rip
// RIP operational data
type Rip struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF related operational data.
    Vrfs Rip_Vrfs

    // Protocol operational data.
    Protocol Rip_Protocol

    // RIP operational data for Default VRF.
    DefaultVrf Rip_DefaultVrf
}

func (rip *Rip) GetFilter() yfilter.YFilter { return rip.YFilter }

func (rip *Rip) SetFilter(yf yfilter.YFilter) { rip.YFilter = yf }

func (rip *Rip) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    if yname == "protocol" { return "Protocol" }
    if yname == "default-vrf" { return "DefaultVrf" }
    return ""
}

func (rip *Rip) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-rip-oper:rip"
}

func (rip *Rip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &rip.Vrfs
    }
    if childYangName == "protocol" {
        return &rip.Protocol
    }
    if childYangName == "default-vrf" {
        return &rip.DefaultVrf
    }
    return nil
}

func (rip *Rip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &rip.Vrfs
    children["protocol"] = &rip.Protocol
    children["default-vrf"] = &rip.DefaultVrf
    return children
}

func (rip *Rip) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rip *Rip) GetBundleName() string { return "cisco_ios_xr" }

func (rip *Rip) GetYangName() string { return "rip" }

func (rip *Rip) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rip *Rip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rip *Rip) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rip *Rip) SetParent(parent types.Entity) { rip.parent = parent }

func (rip *Rip) GetParent() types.Entity { return rip.parent }

func (rip *Rip) GetParentYangName() string { return "Cisco-IOS-XR-ip-rip-oper" }

// Rip_Vrfs
// VRF related operational data
type Rip_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for a particular VRF. The type is slice of Rip_Vrfs_Vrf.
    Vrf []Rip_Vrfs_Vrf
}

func (vrfs *Rip_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Rip_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Rip_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Rip_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Rip_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Rip_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Rip_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Rip_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Rip_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Rip_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Rip_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Rip_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Rip_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Rip_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Rip_Vrfs) GetParentYangName() string { return "rip" }

// Rip_Vrfs_Vrf
// Operational data for a particular VRF
type Rip_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // RIP route database.
    Routes Rip_Vrfs_Vrf_Routes

    // RIP global configuration.
    Configuration Rip_Vrfs_Vrf_Configuration

    // RIP statistics information.
    Statistics Rip_Vrfs_Vrf_Statistics

    // RIP interfaces.
    Interfaces Rip_Vrfs_Vrf_Interfaces

    // Global Information .
    Global Rip_Vrfs_Vrf_Global
}

func (vrf *Rip_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Rip_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Rip_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "routes" { return "Routes" }
    if yname == "configuration" { return "Configuration" }
    if yname == "statistics" { return "Statistics" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "global" { return "Global" }
    return ""
}

func (vrf *Rip_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Rip_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &vrf.Routes
    }
    if childYangName == "configuration" {
        return &vrf.Configuration
    }
    if childYangName == "statistics" {
        return &vrf.Statistics
    }
    if childYangName == "interfaces" {
        return &vrf.Interfaces
    }
    if childYangName == "global" {
        return &vrf.Global
    }
    return nil
}

func (vrf *Rip_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &vrf.Routes
    children["configuration"] = &vrf.Configuration
    children["statistics"] = &vrf.Statistics
    children["interfaces"] = &vrf.Interfaces
    children["global"] = &vrf.Global
    return children
}

func (vrf *Rip_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Rip_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Rip_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Rip_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Rip_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Rip_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Rip_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Rip_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Rip_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Rip_Vrfs_Vrf_Routes
// RIP route database
type Rip_Vrfs_Vrf_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A route in the RIP database. The type is slice of
    // Rip_Vrfs_Vrf_Routes_Route.
    Route []Rip_Vrfs_Vrf_Routes_Route
}

func (routes *Rip_Vrfs_Vrf_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *Rip_Vrfs_Vrf_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *Rip_Vrfs_Vrf_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *Rip_Vrfs_Vrf_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *Rip_Vrfs_Vrf_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *Rip_Vrfs_Vrf_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *Rip_Vrfs_Vrf_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *Rip_Vrfs_Vrf_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *Rip_Vrfs_Vrf_Routes) GetYangName() string { return "routes" }

func (routes *Rip_Vrfs_Vrf_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *Rip_Vrfs_Vrf_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *Rip_Vrfs_Vrf_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *Rip_Vrfs_Vrf_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *Rip_Vrfs_Vrf_Routes) GetParent() types.Entity { return routes.parent }

func (routes *Rip_Vrfs_Vrf_Routes) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_Routes_Route
// A route in the RIP database
type Rip_Vrfs_Vrf_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // Destination IP Address for this route. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLengthXr interface{}

    // Route administrative distance. The type is interface{} with range:
    // 0..65535.
    Distance interface{}

    // Hop count for this route. The type is interface{} with range: 0..65535.
    BgpCount interface{}

    // Type of this route. The type is interface{} with range: 0..65535.
    RouteType interface{}

    // Summary route placeholder indicator. The type is bool.
    RouteSummary interface{}

    // Generic route information. The type is interface{} with range: 0..65535.
    RouteTag interface{}

    // RIB supplied version number. The type is interface{} with range: 0..255.
    Version interface{}

    // RIB supplied route attributes. The type is interface{} with range:
    // 0..4294967295.
    Attributes interface{}

    // Active route indicator. The type is bool.
    Active interface{}

    // Where this route was learnt. The type is RipRouteOrigin.
    PathOrigin interface{}

    // Indicates whether route is in holddown. The type is bool.
    HoldDown interface{}

    // The paths for this route. The type is slice of
    // Rip_Vrfs_Vrf_Routes_Route_Paths.
    Paths []Rip_Vrfs_Vrf_Routes_Route_Paths
}

func (route *Rip_Vrfs_Vrf_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *Rip_Vrfs_Vrf_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *Rip_Vrfs_Vrf_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "prefix-length-xr" { return "PrefixLengthXr" }
    if yname == "distance" { return "Distance" }
    if yname == "bgp-count" { return "BgpCount" }
    if yname == "route-type" { return "RouteType" }
    if yname == "route-summary" { return "RouteSummary" }
    if yname == "route-tag" { return "RouteTag" }
    if yname == "version" { return "Version" }
    if yname == "attributes" { return "Attributes" }
    if yname == "active" { return "Active" }
    if yname == "path-origin" { return "PathOrigin" }
    if yname == "hold-down" { return "HoldDown" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (route *Rip_Vrfs_Vrf_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *Rip_Vrfs_Vrf_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "paths" {
        for _, c := range route.Paths {
            if route.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Routes_Route_Paths{}
        route.Paths = append(route.Paths, child)
        return &route.Paths[len(route.Paths)-1]
    }
    return nil
}

func (route *Rip_Vrfs_Vrf_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range route.Paths {
        children[route.Paths[i].GetSegmentPath()] = &route.Paths[i]
    }
    return children
}

func (route *Rip_Vrfs_Vrf_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["prefix-length"] = route.PrefixLength
    leafs["destination-address"] = route.DestinationAddress
    leafs["prefix-length-xr"] = route.PrefixLengthXr
    leafs["distance"] = route.Distance
    leafs["bgp-count"] = route.BgpCount
    leafs["route-type"] = route.RouteType
    leafs["route-summary"] = route.RouteSummary
    leafs["route-tag"] = route.RouteTag
    leafs["version"] = route.Version
    leafs["attributes"] = route.Attributes
    leafs["active"] = route.Active
    leafs["path-origin"] = route.PathOrigin
    leafs["hold-down"] = route.HoldDown
    return leafs
}

func (route *Rip_Vrfs_Vrf_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *Rip_Vrfs_Vrf_Routes_Route) GetYangName() string { return "route" }

func (route *Rip_Vrfs_Vrf_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *Rip_Vrfs_Vrf_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *Rip_Vrfs_Vrf_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *Rip_Vrfs_Vrf_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *Rip_Vrfs_Vrf_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *Rip_Vrfs_Vrf_Routes_Route) GetParentYangName() string { return "routes" }

// Rip_Vrfs_Vrf_Routes_Route_Paths
// The paths for this route
type Rip_Vrfs_Vrf_Routes_Route_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Next hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}

    // Tag. The type is interface{} with range: 0..65535.
    Tag interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Up time. The type is interface{} with range: 0..4294967295.
    Uptime interface{}

    // Permanent indicator. The type is bool.
    IsPermanent interface{}
}

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "metric" { return "Metric" }
    if yname == "tag" { return "Tag" }
    if yname == "interface" { return "Interface" }
    if yname == "uptime" { return "Uptime" }
    if yname == "is-permanent" { return "IsPermanent" }
    return ""
}

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = paths.SourceAddress
    leafs["next-hop-address"] = paths.NextHopAddress
    leafs["metric"] = paths.Metric
    leafs["tag"] = paths.Tag
    leafs["interface"] = paths.Interface
    leafs["uptime"] = paths.Uptime
    leafs["is-permanent"] = paths.IsPermanent
    return leafs
}

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetYangName() string { return "paths" }

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetParent() types.Entity { return paths.parent }

func (paths *Rip_Vrfs_Vrf_Routes_Route_Paths) GetParentYangName() string { return "route" }

// Rip_Vrfs_Vrf_Configuration
// RIP global configuration
type Rip_Vrfs_Vrf_Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF active indicator. The type is bool.
    Active interface{}

    // VRF added to socket indicator. The type is bool.
    VrFisedSocket interface{}

    // Version of RIP configured. The type is interface{} with range:
    // -2147483648..2147483647.
    RipVersion interface{}

    // Default metric for RIP routes. The type is interface{} with range: 0..255.
    DefaultMetric interface{}

    // Maximum number of paths a route can have. The type is interface{} with
    // range: 0..255.
    MaximumPaths interface{}

    // Auto-summarization indicator. The type is bool.
    AutoSummarize interface{}

    // Use broadcast/multicast address indicator. The type is bool.
    MulticastAddress interface{}

    // Flash update threshold. The type is interface{} with range: 0..255.
    FlashThreshold interface{}

    // Length of the input queue. The type is interface{} with range: 0..65535.
    InputQLength interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Incoming packet source validation indicator. The type is bool.
    ValidationIndicator interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // NSF Enable status. The type is bool.
    NsfStatus interface{}

    // NSF life time. The type is interface{} with range: 0..4294967295.
    NsfLifeTime interface{}
}

func (configuration *Rip_Vrfs_Vrf_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *Rip_Vrfs_Vrf_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *Rip_Vrfs_Vrf_Configuration) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    if yname == "vr-fised-socket" { return "VrFisedSocket" }
    if yname == "rip-version" { return "RipVersion" }
    if yname == "default-metric" { return "DefaultMetric" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    if yname == "auto-summarize" { return "AutoSummarize" }
    if yname == "multicast-address" { return "MulticastAddress" }
    if yname == "flash-threshold" { return "FlashThreshold" }
    if yname == "input-q-length" { return "InputQLength" }
    if yname == "triggered-rip" { return "TriggeredRip" }
    if yname == "validation-indicator" { return "ValidationIndicator" }
    if yname == "update-timer" { return "UpdateTimer" }
    if yname == "next-update-time" { return "NextUpdateTime" }
    if yname == "invalid-timer" { return "InvalidTimer" }
    if yname == "hold-down-timer" { return "HoldDownTimer" }
    if yname == "flush-timer" { return "FlushTimer" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "nsf-status" { return "NsfStatus" }
    if yname == "nsf-life-time" { return "NsfLifeTime" }
    return ""
}

func (configuration *Rip_Vrfs_Vrf_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *Rip_Vrfs_Vrf_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuration *Rip_Vrfs_Vrf_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuration *Rip_Vrfs_Vrf_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active"] = configuration.Active
    leafs["vr-fised-socket"] = configuration.VrFisedSocket
    leafs["rip-version"] = configuration.RipVersion
    leafs["default-metric"] = configuration.DefaultMetric
    leafs["maximum-paths"] = configuration.MaximumPaths
    leafs["auto-summarize"] = configuration.AutoSummarize
    leafs["multicast-address"] = configuration.MulticastAddress
    leafs["flash-threshold"] = configuration.FlashThreshold
    leafs["input-q-length"] = configuration.InputQLength
    leafs["triggered-rip"] = configuration.TriggeredRip
    leafs["validation-indicator"] = configuration.ValidationIndicator
    leafs["update-timer"] = configuration.UpdateTimer
    leafs["next-update-time"] = configuration.NextUpdateTime
    leafs["invalid-timer"] = configuration.InvalidTimer
    leafs["hold-down-timer"] = configuration.HoldDownTimer
    leafs["flush-timer"] = configuration.FlushTimer
    leafs["oom-flags"] = configuration.OomFlags
    leafs["nsf-status"] = configuration.NsfStatus
    leafs["nsf-life-time"] = configuration.NsfLifeTime
    return leafs
}

func (configuration *Rip_Vrfs_Vrf_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *Rip_Vrfs_Vrf_Configuration) GetYangName() string { return "configuration" }

func (configuration *Rip_Vrfs_Vrf_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *Rip_Vrfs_Vrf_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *Rip_Vrfs_Vrf_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *Rip_Vrfs_Vrf_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *Rip_Vrfs_Vrf_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *Rip_Vrfs_Vrf_Configuration) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_Statistics
// RIP statistics information
type Rip_Vrfs_Vrf_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total packets received. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Total discarded packets. The type is interface{} with range: 0..4294967295.
    DiscardedPackets interface{}

    // Total discarded routes. The type is interface{} with range: 0..4294967295.
    DiscardedRoutes interface{}

    // Packets rx in SRP. The type is interface{} with range: 0..4294967295.
    StandbyPacketsReceived interface{}

    // Number of messages sent. The type is interface{} with range: 0..4294967295.
    SentMessages interface{}

    // Number of message send failures. The type is interface{} with range:
    // 0..4294967295.
    SentMessageFailures interface{}

    // Number of RIP queries responded to. The type is interface{} with range:
    // 0..4294967295.
    QueryResponses interface{}

    // Number of periodic RIP updates. The type is interface{} with range:
    // 0..4294967295.
    PeriodicUpdates interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Number of failures to allocate memory for a route. The type is interface{}
    // with range: 0..4294967295.
    RouteMallocFailures interface{}

    // Number of failures to allocate memory for a path. The type is interface{}
    // with range: 0..4294967295.
    PathMallocFailures interface{}

    // Number of route updates to RIB made by RIP. The type is interface{} with
    // range: 0..4294967295.
    RibUpdates interface{}
}

func (statistics *Rip_Vrfs_Vrf_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Rip_Vrfs_Vrf_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Rip_Vrfs_Vrf_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "discarded-packets" { return "DiscardedPackets" }
    if yname == "discarded-routes" { return "DiscardedRoutes" }
    if yname == "standby-packets-received" { return "StandbyPacketsReceived" }
    if yname == "sent-messages" { return "SentMessages" }
    if yname == "sent-message-failures" { return "SentMessageFailures" }
    if yname == "query-responses" { return "QueryResponses" }
    if yname == "periodic-updates" { return "PeriodicUpdates" }
    if yname == "route-count" { return "RouteCount" }
    if yname == "path-count" { return "PathCount" }
    if yname == "route-malloc-failures" { return "RouteMallocFailures" }
    if yname == "path-malloc-failures" { return "PathMallocFailures" }
    if yname == "rib-updates" { return "RibUpdates" }
    return ""
}

func (statistics *Rip_Vrfs_Vrf_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Rip_Vrfs_Vrf_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Rip_Vrfs_Vrf_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Rip_Vrfs_Vrf_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["discarded-packets"] = statistics.DiscardedPackets
    leafs["discarded-routes"] = statistics.DiscardedRoutes
    leafs["standby-packets-received"] = statistics.StandbyPacketsReceived
    leafs["sent-messages"] = statistics.SentMessages
    leafs["sent-message-failures"] = statistics.SentMessageFailures
    leafs["query-responses"] = statistics.QueryResponses
    leafs["periodic-updates"] = statistics.PeriodicUpdates
    leafs["route-count"] = statistics.RouteCount
    leafs["path-count"] = statistics.PathCount
    leafs["route-malloc-failures"] = statistics.RouteMallocFailures
    leafs["path-malloc-failures"] = statistics.PathMallocFailures
    leafs["rib-updates"] = statistics.RibUpdates
    return leafs
}

func (statistics *Rip_Vrfs_Vrf_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Rip_Vrfs_Vrf_Statistics) GetYangName() string { return "statistics" }

func (statistics *Rip_Vrfs_Vrf_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Rip_Vrfs_Vrf_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Rip_Vrfs_Vrf_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Rip_Vrfs_Vrf_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Rip_Vrfs_Vrf_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Rip_Vrfs_Vrf_Statistics) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_Interfaces
// RIP interfaces
type Rip_Vrfs_Vrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular RIP interface. The type is slice of
    // Rip_Vrfs_Vrf_Interfaces_Interface.
    Interface []Rip_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_Interfaces_Interface
// Information about a particular RIP interface
type Rip_Vrfs_Vrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    Interface interface{}

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    IfHandle interface{}

    // Whether RIP is enabled on this interface. The type is bool.
    RipEnabled interface{}

    // Passive interface indicator. The type is bool.
    IsPassiveInterface interface{}

    // Use broadcast address for v2 packets. The type is bool.
    MulticastAddress interface{}

    // Accept routes of metric 0 indicator. The type is bool.
    AcceptMetric interface{}

    // Versions that the interface is sending. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // Versions that the interface will recieve. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Current state of the interface. The type is InterfaceState.
    State interface{}

    // IP Address of this interface. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Prefix length of the IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Cost added to routes through this interface. The type is interface{} with
    // range: 0..4294967295.
    MetricCost interface{}

    // Split horizon enabled indicator. The type is bool.
    SplitHorizon interface{}

    // Poisoned reverse enabled indicator. The type is bool.
    PoisonHorizon interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Interface's triggered RIP neighbor. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // Multicast group join status. The type is bool.
    JoinStatus interface{}

    // LPTSState. The type is bool.
    LptsState interface{}

    // Authentication Mode. The type is interface{} with range: 0..4294967295.
    AuthMode interface{}

    // Authentication Keychain Name. The type is string.
    AuthKeychain interface{}

    // Authentication send key exists. The type is bool.
    SendAuthKeyExists interface{}

    // Authentication key programmed with MD5 algorithm. The type is bool.
    AuthKeyMd5 interface{}

    // Current active Send Authentication Key Id. The type is interface{} with
    // range: 0..18446744073709551615.
    AuthKeySendId interface{}

    // Total packets received. The type is interface{} with range: 0..4294967295.
    TotalPktRecvd interface{}

    // Packets dropped due to wrong keychain configured. The type is interface{}
    // with range: 0..4294967295.
    PktDropWrongKc interface{}

    // Packets dropped due to missing authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropNoAuth interface{}

    // Packets dropped due to invalid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropInvalidAuth interface{}

    // Packets accepted with valid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktAcceptedValidAuth interface{}

    // User defined summary addresses. The type is slice of
    // Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary.
    RipSummary []Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary

    // Neighbors on this interface. The type is slice of
    // Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer.
    RipPeer []Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "rip-enabled" { return "RipEnabled" }
    if yname == "is-passive-interface" { return "IsPassiveInterface" }
    if yname == "multicast-address" { return "MulticastAddress" }
    if yname == "accept-metric" { return "AcceptMetric" }
    if yname == "send-version" { return "SendVersion" }
    if yname == "receive-version" { return "ReceiveVersion" }
    if yname == "state" { return "State" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "metric-cost" { return "MetricCost" }
    if yname == "split-horizon" { return "SplitHorizon" }
    if yname == "poison-horizon" { return "PoisonHorizon" }
    if yname == "triggered-rip" { return "TriggeredRip" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "join-status" { return "JoinStatus" }
    if yname == "lpts-state" { return "LptsState" }
    if yname == "auth-mode" { return "AuthMode" }
    if yname == "auth-keychain" { return "AuthKeychain" }
    if yname == "send-auth-key-exists" { return "SendAuthKeyExists" }
    if yname == "auth-key-md5" { return "AuthKeyMd5" }
    if yname == "auth-key-send-id" { return "AuthKeySendId" }
    if yname == "total-pkt-recvd" { return "TotalPktRecvd" }
    if yname == "pkt-drop-wrong-kc" { return "PktDropWrongKc" }
    if yname == "pkt-drop-no-auth" { return "PktDropNoAuth" }
    if yname == "pkt-drop-invalid-auth" { return "PktDropInvalidAuth" }
    if yname == "pkt-accepted-valid-auth" { return "PktAcceptedValidAuth" }
    if yname == "rip-summary" { return "RipSummary" }
    if yname == "rip-peer" { return "RipPeer" }
    return ""
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rip-summary" {
        for _, c := range self.RipSummary {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary{}
        self.RipSummary = append(self.RipSummary, child)
        return &self.RipSummary[len(self.RipSummary)-1]
    }
    if childYangName == "rip-peer" {
        for _, c := range self.RipPeer {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer{}
        self.RipPeer = append(self.RipPeer, child)
        return &self.RipPeer[len(self.RipPeer)-1]
    }
    return nil
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.RipSummary {
        children[self.RipSummary[i].GetSegmentPath()] = &self.RipSummary[i]
    }
    for i := range self.RipPeer {
        children[self.RipPeer[i].GetSegmentPath()] = &self.RipPeer[i]
    }
    return children
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["if-handle"] = self.IfHandle
    leafs["rip-enabled"] = self.RipEnabled
    leafs["is-passive-interface"] = self.IsPassiveInterface
    leafs["multicast-address"] = self.MulticastAddress
    leafs["accept-metric"] = self.AcceptMetric
    leafs["send-version"] = self.SendVersion
    leafs["receive-version"] = self.ReceiveVersion
    leafs["state"] = self.State
    leafs["destination-address"] = self.DestinationAddress
    leafs["prefix-length"] = self.PrefixLength
    leafs["metric-cost"] = self.MetricCost
    leafs["split-horizon"] = self.SplitHorizon
    leafs["poison-horizon"] = self.PoisonHorizon
    leafs["triggered-rip"] = self.TriggeredRip
    leafs["neighbor-address"] = self.NeighborAddress
    leafs["oom-flags"] = self.OomFlags
    leafs["join-status"] = self.JoinStatus
    leafs["lpts-state"] = self.LptsState
    leafs["auth-mode"] = self.AuthMode
    leafs["auth-keychain"] = self.AuthKeychain
    leafs["send-auth-key-exists"] = self.SendAuthKeyExists
    leafs["auth-key-md5"] = self.AuthKeyMd5
    leafs["auth-key-send-id"] = self.AuthKeySendId
    leafs["total-pkt-recvd"] = self.TotalPktRecvd
    leafs["pkt-drop-wrong-kc"] = self.PktDropWrongKc
    leafs["pkt-drop-no-auth"] = self.PktDropNoAuth
    leafs["pkt-drop-invalid-auth"] = self.PktDropInvalidAuth
    leafs["pkt-accepted-valid-auth"] = self.PktAcceptedValidAuth
    return leafs
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary
// User defined summary addresses
type Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary address prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Summary address prefix length. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Summary address next hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Summary metric. The type is interface{} with range:
    // -2147483648..2147483647.
    Metric interface{}
}

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetFilter() yfilter.YFilter { return ripSummary.YFilter }

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) SetFilter(yf yfilter.YFilter) { ripSummary.YFilter = yf }

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetSegmentPath() string {
    return "rip-summary"
}

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = ripSummary.Prefix
    leafs["prefix-length"] = ripSummary.PrefixLength
    leafs["next-hop-address"] = ripSummary.NextHopAddress
    leafs["metric"] = ripSummary.Metric
    return leafs
}

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetYangName() string { return "rip-summary" }

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) SetParent(parent types.Entity) { ripSummary.parent = parent }

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetParent() types.Entity { return ripSummary.parent }

func (ripSummary *Rip_Vrfs_Vrf_Interfaces_Interface_RipSummary) GetParentYangName() string { return "interface" }

// Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer
// Neighbors on this interface
type Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Uptime of this peer. The type is interface{} with range: 0..4294967295.
    PeerUptime interface{}

    // IP Address of this peer. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // RIP version for this peer. The type is interface{} with range: 0..255.
    PeerVersion interface{}

    // Discarded packets from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerPackets interface{}

    // Discarded routes from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerRoutes interface{}
}

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetFilter() yfilter.YFilter { return ripPeer.YFilter }

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) SetFilter(yf yfilter.YFilter) { ripPeer.YFilter = yf }

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetGoName(yname string) string {
    if yname == "peer-uptime" { return "PeerUptime" }
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "peer-version" { return "PeerVersion" }
    if yname == "discarded-peer-packets" { return "DiscardedPeerPackets" }
    if yname == "discarded-peer-routes" { return "DiscardedPeerRoutes" }
    return ""
}

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetSegmentPath() string {
    return "rip-peer"
}

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-uptime"] = ripPeer.PeerUptime
    leafs["peer-address"] = ripPeer.PeerAddress
    leafs["peer-version"] = ripPeer.PeerVersion
    leafs["discarded-peer-packets"] = ripPeer.DiscardedPeerPackets
    leafs["discarded-peer-routes"] = ripPeer.DiscardedPeerRoutes
    return leafs
}

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetBundleName() string { return "cisco_ios_xr" }

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetYangName() string { return "rip-peer" }

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) SetParent(parent types.Entity) { ripPeer.parent = parent }

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetParent() types.Entity { return ripPeer.parent }

func (ripPeer *Rip_Vrfs_Vrf_Interfaces_Interface_RipPeer) GetParentYangName() string { return "interface" }

// Rip_Vrfs_Vrf_Global
// Global Information 
type Rip_Vrfs_Vrf_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF summary data.
    VrfSummary Rip_Vrfs_Vrf_Global_VrfSummary

    // List of Interfaces configured. The type is slice of
    // Rip_Vrfs_Vrf_Global_InterfaceSummary.
    InterfaceSummary []Rip_Vrfs_Vrf_Global_InterfaceSummary
}

func (global *Rip_Vrfs_Vrf_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Rip_Vrfs_Vrf_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Rip_Vrfs_Vrf_Global) GetGoName(yname string) string {
    if yname == "vrf-summary" { return "VrfSummary" }
    if yname == "interface-summary" { return "InterfaceSummary" }
    return ""
}

func (global *Rip_Vrfs_Vrf_Global) GetSegmentPath() string {
    return "global"
}

func (global *Rip_Vrfs_Vrf_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-summary" {
        return &global.VrfSummary
    }
    if childYangName == "interface-summary" {
        for _, c := range global.InterfaceSummary {
            if global.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Vrfs_Vrf_Global_InterfaceSummary{}
        global.InterfaceSummary = append(global.InterfaceSummary, child)
        return &global.InterfaceSummary[len(global.InterfaceSummary)-1]
    }
    return nil
}

func (global *Rip_Vrfs_Vrf_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf-summary"] = &global.VrfSummary
    for i := range global.InterfaceSummary {
        children[global.InterfaceSummary[i].GetSegmentPath()] = &global.InterfaceSummary[i]
    }
    return children
}

func (global *Rip_Vrfs_Vrf_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *Rip_Vrfs_Vrf_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *Rip_Vrfs_Vrf_Global) GetYangName() string { return "global" }

func (global *Rip_Vrfs_Vrf_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *Rip_Vrfs_Vrf_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *Rip_Vrfs_Vrf_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *Rip_Vrfs_Vrf_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Rip_Vrfs_Vrf_Global) GetParent() types.Entity { return global.parent }

func (global *Rip_Vrfs_Vrf_Global) GetParentYangName() string { return "vrf" }

// Rip_Vrfs_Vrf_Global_VrfSummary
// VRF summary data
type Rip_Vrfs_Vrf_Global_VrfSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Active indicator. The type is bool.
    Active interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Number of interfaces configured. The type is interface{} with range:
    // 0..4294967295.
    InterfaceConfiguredCount interface{}

    // Number of active interfaces. The type is interface{} with range:
    // 0..4294967295.
    ActiveInterfaceCount interface{}
}

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetFilter() yfilter.YFilter { return vrfSummary.YFilter }

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) SetFilter(yf yfilter.YFilter) { vrfSummary.YFilter = yf }

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "active" { return "Active" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "route-count" { return "RouteCount" }
    if yname == "path-count" { return "PathCount" }
    if yname == "update-timer" { return "UpdateTimer" }
    if yname == "next-update-time" { return "NextUpdateTime" }
    if yname == "invalid-timer" { return "InvalidTimer" }
    if yname == "hold-down-timer" { return "HoldDownTimer" }
    if yname == "flush-timer" { return "FlushTimer" }
    if yname == "interface-configured-count" { return "InterfaceConfiguredCount" }
    if yname == "active-interface-count" { return "ActiveInterfaceCount" }
    return ""
}

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetSegmentPath() string {
    return "vrf-summary"
}

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrfSummary.VrfName
    leafs["active"] = vrfSummary.Active
    leafs["oom-flags"] = vrfSummary.OomFlags
    leafs["route-count"] = vrfSummary.RouteCount
    leafs["path-count"] = vrfSummary.PathCount
    leafs["update-timer"] = vrfSummary.UpdateTimer
    leafs["next-update-time"] = vrfSummary.NextUpdateTime
    leafs["invalid-timer"] = vrfSummary.InvalidTimer
    leafs["hold-down-timer"] = vrfSummary.HoldDownTimer
    leafs["flush-timer"] = vrfSummary.FlushTimer
    leafs["interface-configured-count"] = vrfSummary.InterfaceConfiguredCount
    leafs["active-interface-count"] = vrfSummary.ActiveInterfaceCount
    return leafs
}

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetBundleName() string { return "cisco_ios_xr" }

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetYangName() string { return "vrf-summary" }

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) SetParent(parent types.Entity) { vrfSummary.parent = parent }

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetParent() types.Entity { return vrfSummary.parent }

func (vrfSummary *Rip_Vrfs_Vrf_Global_VrfSummary) GetParentYangName() string { return "global" }

// Rip_Vrfs_Vrf_Global_InterfaceSummary
// List of Interfaces configured
type Rip_Vrfs_Vrf_Global_InterfaceSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // RIP enabled indicator. The type is bool.
    Enabled interface{}

    // Interface state. The type is InterfaceState.
    State interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // RIP versions this interface sends out. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // RIP versions this interface will receive. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Number of neighbors seen. The type is interface{} with range:
    // 0..4294967295.
    NeighborCount interface{}
}

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetFilter() yfilter.YFilter { return interfaceSummary.YFilter }

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) SetFilter(yf yfilter.YFilter) { interfaceSummary.YFilter = yf }

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enabled" { return "Enabled" }
    if yname == "state" { return "State" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "send-version" { return "SendVersion" }
    if yname == "receive-version" { return "ReceiveVersion" }
    if yname == "neighbor-count" { return "NeighborCount" }
    return ""
}

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetSegmentPath() string {
    return "interface-summary"
}

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceSummary.InterfaceName
    leafs["enabled"] = interfaceSummary.Enabled
    leafs["state"] = interfaceSummary.State
    leafs["destination-address"] = interfaceSummary.DestinationAddress
    leafs["prefix-length"] = interfaceSummary.PrefixLength
    leafs["oom-flags"] = interfaceSummary.OomFlags
    leafs["send-version"] = interfaceSummary.SendVersion
    leafs["receive-version"] = interfaceSummary.ReceiveVersion
    leafs["neighbor-count"] = interfaceSummary.NeighborCount
    return leafs
}

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetYangName() string { return "interface-summary" }

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) SetParent(parent types.Entity) { interfaceSummary.parent = parent }

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetParent() types.Entity { return interfaceSummary.parent }

func (interfaceSummary *Rip_Vrfs_Vrf_Global_InterfaceSummary) GetParentYangName() string { return "global" }

// Rip_Protocol
// Protocol operational data
type Rip_Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RIP global process .
    Process Rip_Protocol_Process

    // RIP operational data for Default VRF.
    DefaultVrf Rip_Protocol_DefaultVrf
}

func (protocol *Rip_Protocol) GetFilter() yfilter.YFilter { return protocol.YFilter }

func (protocol *Rip_Protocol) SetFilter(yf yfilter.YFilter) { protocol.YFilter = yf }

func (protocol *Rip_Protocol) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    if yname == "default-vrf" { return "DefaultVrf" }
    return ""
}

func (protocol *Rip_Protocol) GetSegmentPath() string {
    return "protocol"
}

func (protocol *Rip_Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process" {
        return &protocol.Process
    }
    if childYangName == "default-vrf" {
        return &protocol.DefaultVrf
    }
    return nil
}

func (protocol *Rip_Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["process"] = &protocol.Process
    children["default-vrf"] = &protocol.DefaultVrf
    return children
}

func (protocol *Rip_Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocol *Rip_Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (protocol *Rip_Protocol) GetYangName() string { return "protocol" }

func (protocol *Rip_Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocol *Rip_Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocol *Rip_Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocol *Rip_Protocol) SetParent(parent types.Entity) { protocol.parent = parent }

func (protocol *Rip_Protocol) GetParent() types.Entity { return protocol.parent }

func (protocol *Rip_Protocol) GetParentYangName() string { return "rip" }

// Rip_Protocol_Process
// RIP global process 
type Rip_Protocol_Process struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of VRFs configured. The type is interface{} with range:
    // 0..4294967295.
    VrfConfigCount interface{}

    // Number of active VRFs. The type is interface{} with range: 0..4294967295.
    VrfActiveCount interface{}

    // Socket descriptior. The type is interface{} with range:
    // -2147483648..2147483647.
    SocketDescriptor interface{}

    // Current OOM state. The type is interface{} with range:
    // -2147483648..2147483647.
    CurrentOomState interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // List of VRFs configured. The type is slice of
    // Rip_Protocol_Process_VrfSummary.
    VrfSummary []Rip_Protocol_Process_VrfSummary
}

func (process *Rip_Protocol_Process) GetFilter() yfilter.YFilter { return process.YFilter }

func (process *Rip_Protocol_Process) SetFilter(yf yfilter.YFilter) { process.YFilter = yf }

func (process *Rip_Protocol_Process) GetGoName(yname string) string {
    if yname == "vrf-config-count" { return "VrfConfigCount" }
    if yname == "vrf-active-count" { return "VrfActiveCount" }
    if yname == "socket-descriptor" { return "SocketDescriptor" }
    if yname == "current-oom-state" { return "CurrentOomState" }
    if yname == "route-count" { return "RouteCount" }
    if yname == "path-count" { return "PathCount" }
    if yname == "vrf-summary" { return "VrfSummary" }
    return ""
}

func (process *Rip_Protocol_Process) GetSegmentPath() string {
    return "process"
}

func (process *Rip_Protocol_Process) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-summary" {
        for _, c := range process.VrfSummary {
            if process.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Protocol_Process_VrfSummary{}
        process.VrfSummary = append(process.VrfSummary, child)
        return &process.VrfSummary[len(process.VrfSummary)-1]
    }
    return nil
}

func (process *Rip_Protocol_Process) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range process.VrfSummary {
        children[process.VrfSummary[i].GetSegmentPath()] = &process.VrfSummary[i]
    }
    return children
}

func (process *Rip_Protocol_Process) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-config-count"] = process.VrfConfigCount
    leafs["vrf-active-count"] = process.VrfActiveCount
    leafs["socket-descriptor"] = process.SocketDescriptor
    leafs["current-oom-state"] = process.CurrentOomState
    leafs["route-count"] = process.RouteCount
    leafs["path-count"] = process.PathCount
    return leafs
}

func (process *Rip_Protocol_Process) GetBundleName() string { return "cisco_ios_xr" }

func (process *Rip_Protocol_Process) GetYangName() string { return "process" }

func (process *Rip_Protocol_Process) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (process *Rip_Protocol_Process) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (process *Rip_Protocol_Process) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (process *Rip_Protocol_Process) SetParent(parent types.Entity) { process.parent = parent }

func (process *Rip_Protocol_Process) GetParent() types.Entity { return process.parent }

func (process *Rip_Protocol_Process) GetParentYangName() string { return "protocol" }

// Rip_Protocol_Process_VrfSummary
// List of VRFs configured
type Rip_Protocol_Process_VrfSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Active indicator. The type is bool.
    Active interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Number of interfaces configured. The type is interface{} with range:
    // 0..4294967295.
    InterfaceConfiguredCount interface{}

    // Number of active interfaces. The type is interface{} with range:
    // 0..4294967295.
    ActiveInterfaceCount interface{}
}

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetFilter() yfilter.YFilter { return vrfSummary.YFilter }

func (vrfSummary *Rip_Protocol_Process_VrfSummary) SetFilter(yf yfilter.YFilter) { vrfSummary.YFilter = yf }

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "active" { return "Active" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "route-count" { return "RouteCount" }
    if yname == "path-count" { return "PathCount" }
    if yname == "update-timer" { return "UpdateTimer" }
    if yname == "next-update-time" { return "NextUpdateTime" }
    if yname == "invalid-timer" { return "InvalidTimer" }
    if yname == "hold-down-timer" { return "HoldDownTimer" }
    if yname == "flush-timer" { return "FlushTimer" }
    if yname == "interface-configured-count" { return "InterfaceConfiguredCount" }
    if yname == "active-interface-count" { return "ActiveInterfaceCount" }
    return ""
}

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetSegmentPath() string {
    return "vrf-summary"
}

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrfSummary.VrfName
    leafs["active"] = vrfSummary.Active
    leafs["oom-flags"] = vrfSummary.OomFlags
    leafs["route-count"] = vrfSummary.RouteCount
    leafs["path-count"] = vrfSummary.PathCount
    leafs["update-timer"] = vrfSummary.UpdateTimer
    leafs["next-update-time"] = vrfSummary.NextUpdateTime
    leafs["invalid-timer"] = vrfSummary.InvalidTimer
    leafs["hold-down-timer"] = vrfSummary.HoldDownTimer
    leafs["flush-timer"] = vrfSummary.FlushTimer
    leafs["interface-configured-count"] = vrfSummary.InterfaceConfiguredCount
    leafs["active-interface-count"] = vrfSummary.ActiveInterfaceCount
    return leafs
}

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetBundleName() string { return "cisco_ios_xr" }

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetYangName() string { return "vrf-summary" }

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfSummary *Rip_Protocol_Process_VrfSummary) SetParent(parent types.Entity) { vrfSummary.parent = parent }

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetParent() types.Entity { return vrfSummary.parent }

func (vrfSummary *Rip_Protocol_Process_VrfSummary) GetParentYangName() string { return "process" }

// Rip_Protocol_DefaultVrf
// RIP operational data for Default VRF
type Rip_Protocol_DefaultVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RIP route database.
    Routes Rip_Protocol_DefaultVrf_Routes

    // RIP global configuration.
    Configuration Rip_Protocol_DefaultVrf_Configuration

    // RIP statistics information.
    Statistics Rip_Protocol_DefaultVrf_Statistics

    // RIP interfaces.
    Interfaces Rip_Protocol_DefaultVrf_Interfaces

    // Global Information .
    Global Rip_Protocol_DefaultVrf_Global
}

func (defaultVrf *Rip_Protocol_DefaultVrf) GetFilter() yfilter.YFilter { return defaultVrf.YFilter }

func (defaultVrf *Rip_Protocol_DefaultVrf) SetFilter(yf yfilter.YFilter) { defaultVrf.YFilter = yf }

func (defaultVrf *Rip_Protocol_DefaultVrf) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "configuration" { return "Configuration" }
    if yname == "statistics" { return "Statistics" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "global" { return "Global" }
    return ""
}

func (defaultVrf *Rip_Protocol_DefaultVrf) GetSegmentPath() string {
    return "default-vrf"
}

func (defaultVrf *Rip_Protocol_DefaultVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &defaultVrf.Routes
    }
    if childYangName == "configuration" {
        return &defaultVrf.Configuration
    }
    if childYangName == "statistics" {
        return &defaultVrf.Statistics
    }
    if childYangName == "interfaces" {
        return &defaultVrf.Interfaces
    }
    if childYangName == "global" {
        return &defaultVrf.Global
    }
    return nil
}

func (defaultVrf *Rip_Protocol_DefaultVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &defaultVrf.Routes
    children["configuration"] = &defaultVrf.Configuration
    children["statistics"] = &defaultVrf.Statistics
    children["interfaces"] = &defaultVrf.Interfaces
    children["global"] = &defaultVrf.Global
    return children
}

func (defaultVrf *Rip_Protocol_DefaultVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultVrf *Rip_Protocol_DefaultVrf) GetBundleName() string { return "cisco_ios_xr" }

func (defaultVrf *Rip_Protocol_DefaultVrf) GetYangName() string { return "default-vrf" }

func (defaultVrf *Rip_Protocol_DefaultVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultVrf *Rip_Protocol_DefaultVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultVrf *Rip_Protocol_DefaultVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultVrf *Rip_Protocol_DefaultVrf) SetParent(parent types.Entity) { defaultVrf.parent = parent }

func (defaultVrf *Rip_Protocol_DefaultVrf) GetParent() types.Entity { return defaultVrf.parent }

func (defaultVrf *Rip_Protocol_DefaultVrf) GetParentYangName() string { return "protocol" }

// Rip_Protocol_DefaultVrf_Routes
// RIP route database
type Rip_Protocol_DefaultVrf_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A route in the RIP database. The type is slice of
    // Rip_Protocol_DefaultVrf_Routes_Route.
    Route []Rip_Protocol_DefaultVrf_Routes_Route
}

func (routes *Rip_Protocol_DefaultVrf_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *Rip_Protocol_DefaultVrf_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *Rip_Protocol_DefaultVrf_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *Rip_Protocol_DefaultVrf_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *Rip_Protocol_DefaultVrf_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Protocol_DefaultVrf_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *Rip_Protocol_DefaultVrf_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *Rip_Protocol_DefaultVrf_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *Rip_Protocol_DefaultVrf_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *Rip_Protocol_DefaultVrf_Routes) GetYangName() string { return "routes" }

func (routes *Rip_Protocol_DefaultVrf_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *Rip_Protocol_DefaultVrf_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *Rip_Protocol_DefaultVrf_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *Rip_Protocol_DefaultVrf_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *Rip_Protocol_DefaultVrf_Routes) GetParent() types.Entity { return routes.parent }

func (routes *Rip_Protocol_DefaultVrf_Routes) GetParentYangName() string { return "default-vrf" }

// Rip_Protocol_DefaultVrf_Routes_Route
// A route in the RIP database
type Rip_Protocol_DefaultVrf_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // Destination IP Address for this route. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLengthXr interface{}

    // Route administrative distance. The type is interface{} with range:
    // 0..65535.
    Distance interface{}

    // Hop count for this route. The type is interface{} with range: 0..65535.
    BgpCount interface{}

    // Type of this route. The type is interface{} with range: 0..65535.
    RouteType interface{}

    // Summary route placeholder indicator. The type is bool.
    RouteSummary interface{}

    // Generic route information. The type is interface{} with range: 0..65535.
    RouteTag interface{}

    // RIB supplied version number. The type is interface{} with range: 0..255.
    Version interface{}

    // RIB supplied route attributes. The type is interface{} with range:
    // 0..4294967295.
    Attributes interface{}

    // Active route indicator. The type is bool.
    Active interface{}

    // Where this route was learnt. The type is RipRouteOrigin.
    PathOrigin interface{}

    // Indicates whether route is in holddown. The type is bool.
    HoldDown interface{}

    // The paths for this route. The type is slice of
    // Rip_Protocol_DefaultVrf_Routes_Route_Paths.
    Paths []Rip_Protocol_DefaultVrf_Routes_Route_Paths
}

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *Rip_Protocol_DefaultVrf_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "prefix-length-xr" { return "PrefixLengthXr" }
    if yname == "distance" { return "Distance" }
    if yname == "bgp-count" { return "BgpCount" }
    if yname == "route-type" { return "RouteType" }
    if yname == "route-summary" { return "RouteSummary" }
    if yname == "route-tag" { return "RouteTag" }
    if yname == "version" { return "Version" }
    if yname == "attributes" { return "Attributes" }
    if yname == "active" { return "Active" }
    if yname == "path-origin" { return "PathOrigin" }
    if yname == "hold-down" { return "HoldDown" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "paths" {
        for _, c := range route.Paths {
            if route.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Protocol_DefaultVrf_Routes_Route_Paths{}
        route.Paths = append(route.Paths, child)
        return &route.Paths[len(route.Paths)-1]
    }
    return nil
}

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range route.Paths {
        children[route.Paths[i].GetSegmentPath()] = &route.Paths[i]
    }
    return children
}

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["prefix-length"] = route.PrefixLength
    leafs["destination-address"] = route.DestinationAddress
    leafs["prefix-length-xr"] = route.PrefixLengthXr
    leafs["distance"] = route.Distance
    leafs["bgp-count"] = route.BgpCount
    leafs["route-type"] = route.RouteType
    leafs["route-summary"] = route.RouteSummary
    leafs["route-tag"] = route.RouteTag
    leafs["version"] = route.Version
    leafs["attributes"] = route.Attributes
    leafs["active"] = route.Active
    leafs["path-origin"] = route.PathOrigin
    leafs["hold-down"] = route.HoldDown
    return leafs
}

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetYangName() string { return "route" }

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *Rip_Protocol_DefaultVrf_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *Rip_Protocol_DefaultVrf_Routes_Route) GetParentYangName() string { return "routes" }

// Rip_Protocol_DefaultVrf_Routes_Route_Paths
// The paths for this route
type Rip_Protocol_DefaultVrf_Routes_Route_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Next hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}

    // Tag. The type is interface{} with range: 0..65535.
    Tag interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Up time. The type is interface{} with range: 0..4294967295.
    Uptime interface{}

    // Permanent indicator. The type is bool.
    IsPermanent interface{}
}

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "metric" { return "Metric" }
    if yname == "tag" { return "Tag" }
    if yname == "interface" { return "Interface" }
    if yname == "uptime" { return "Uptime" }
    if yname == "is-permanent" { return "IsPermanent" }
    return ""
}

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = paths.SourceAddress
    leafs["next-hop-address"] = paths.NextHopAddress
    leafs["metric"] = paths.Metric
    leafs["tag"] = paths.Tag
    leafs["interface"] = paths.Interface
    leafs["uptime"] = paths.Uptime
    leafs["is-permanent"] = paths.IsPermanent
    return leafs
}

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetYangName() string { return "paths" }

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetParent() types.Entity { return paths.parent }

func (paths *Rip_Protocol_DefaultVrf_Routes_Route_Paths) GetParentYangName() string { return "route" }

// Rip_Protocol_DefaultVrf_Configuration
// RIP global configuration
type Rip_Protocol_DefaultVrf_Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF active indicator. The type is bool.
    Active interface{}

    // VRF added to socket indicator. The type is bool.
    VrFisedSocket interface{}

    // Version of RIP configured. The type is interface{} with range:
    // -2147483648..2147483647.
    RipVersion interface{}

    // Default metric for RIP routes. The type is interface{} with range: 0..255.
    DefaultMetric interface{}

    // Maximum number of paths a route can have. The type is interface{} with
    // range: 0..255.
    MaximumPaths interface{}

    // Auto-summarization indicator. The type is bool.
    AutoSummarize interface{}

    // Use broadcast/multicast address indicator. The type is bool.
    MulticastAddress interface{}

    // Flash update threshold. The type is interface{} with range: 0..255.
    FlashThreshold interface{}

    // Length of the input queue. The type is interface{} with range: 0..65535.
    InputQLength interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Incoming packet source validation indicator. The type is bool.
    ValidationIndicator interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // NSF Enable status. The type is bool.
    NsfStatus interface{}

    // NSF life time. The type is interface{} with range: 0..4294967295.
    NsfLifeTime interface{}
}

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *Rip_Protocol_DefaultVrf_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    if yname == "vr-fised-socket" { return "VrFisedSocket" }
    if yname == "rip-version" { return "RipVersion" }
    if yname == "default-metric" { return "DefaultMetric" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    if yname == "auto-summarize" { return "AutoSummarize" }
    if yname == "multicast-address" { return "MulticastAddress" }
    if yname == "flash-threshold" { return "FlashThreshold" }
    if yname == "input-q-length" { return "InputQLength" }
    if yname == "triggered-rip" { return "TriggeredRip" }
    if yname == "validation-indicator" { return "ValidationIndicator" }
    if yname == "update-timer" { return "UpdateTimer" }
    if yname == "next-update-time" { return "NextUpdateTime" }
    if yname == "invalid-timer" { return "InvalidTimer" }
    if yname == "hold-down-timer" { return "HoldDownTimer" }
    if yname == "flush-timer" { return "FlushTimer" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "nsf-status" { return "NsfStatus" }
    if yname == "nsf-life-time" { return "NsfLifeTime" }
    return ""
}

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active"] = configuration.Active
    leafs["vr-fised-socket"] = configuration.VrFisedSocket
    leafs["rip-version"] = configuration.RipVersion
    leafs["default-metric"] = configuration.DefaultMetric
    leafs["maximum-paths"] = configuration.MaximumPaths
    leafs["auto-summarize"] = configuration.AutoSummarize
    leafs["multicast-address"] = configuration.MulticastAddress
    leafs["flash-threshold"] = configuration.FlashThreshold
    leafs["input-q-length"] = configuration.InputQLength
    leafs["triggered-rip"] = configuration.TriggeredRip
    leafs["validation-indicator"] = configuration.ValidationIndicator
    leafs["update-timer"] = configuration.UpdateTimer
    leafs["next-update-time"] = configuration.NextUpdateTime
    leafs["invalid-timer"] = configuration.InvalidTimer
    leafs["hold-down-timer"] = configuration.HoldDownTimer
    leafs["flush-timer"] = configuration.FlushTimer
    leafs["oom-flags"] = configuration.OomFlags
    leafs["nsf-status"] = configuration.NsfStatus
    leafs["nsf-life-time"] = configuration.NsfLifeTime
    return leafs
}

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetYangName() string { return "configuration" }

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *Rip_Protocol_DefaultVrf_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *Rip_Protocol_DefaultVrf_Configuration) GetParentYangName() string { return "default-vrf" }

// Rip_Protocol_DefaultVrf_Statistics
// RIP statistics information
type Rip_Protocol_DefaultVrf_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total packets received. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Total discarded packets. The type is interface{} with range: 0..4294967295.
    DiscardedPackets interface{}

    // Total discarded routes. The type is interface{} with range: 0..4294967295.
    DiscardedRoutes interface{}

    // Packets rx in SRP. The type is interface{} with range: 0..4294967295.
    StandbyPacketsReceived interface{}

    // Number of messages sent. The type is interface{} with range: 0..4294967295.
    SentMessages interface{}

    // Number of message send failures. The type is interface{} with range:
    // 0..4294967295.
    SentMessageFailures interface{}

    // Number of RIP queries responded to. The type is interface{} with range:
    // 0..4294967295.
    QueryResponses interface{}

    // Number of periodic RIP updates. The type is interface{} with range:
    // 0..4294967295.
    PeriodicUpdates interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Number of failures to allocate memory for a route. The type is interface{}
    // with range: 0..4294967295.
    RouteMallocFailures interface{}

    // Number of failures to allocate memory for a path. The type is interface{}
    // with range: 0..4294967295.
    PathMallocFailures interface{}

    // Number of route updates to RIB made by RIP. The type is interface{} with
    // range: 0..4294967295.
    RibUpdates interface{}
}

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Rip_Protocol_DefaultVrf_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "discarded-packets" { return "DiscardedPackets" }
    if yname == "discarded-routes" { return "DiscardedRoutes" }
    if yname == "standby-packets-received" { return "StandbyPacketsReceived" }
    if yname == "sent-messages" { return "SentMessages" }
    if yname == "sent-message-failures" { return "SentMessageFailures" }
    if yname == "query-responses" { return "QueryResponses" }
    if yname == "periodic-updates" { return "PeriodicUpdates" }
    if yname == "route-count" { return "RouteCount" }
    if yname == "path-count" { return "PathCount" }
    if yname == "route-malloc-failures" { return "RouteMallocFailures" }
    if yname == "path-malloc-failures" { return "PathMallocFailures" }
    if yname == "rib-updates" { return "RibUpdates" }
    return ""
}

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["discarded-packets"] = statistics.DiscardedPackets
    leafs["discarded-routes"] = statistics.DiscardedRoutes
    leafs["standby-packets-received"] = statistics.StandbyPacketsReceived
    leafs["sent-messages"] = statistics.SentMessages
    leafs["sent-message-failures"] = statistics.SentMessageFailures
    leafs["query-responses"] = statistics.QueryResponses
    leafs["periodic-updates"] = statistics.PeriodicUpdates
    leafs["route-count"] = statistics.RouteCount
    leafs["path-count"] = statistics.PathCount
    leafs["route-malloc-failures"] = statistics.RouteMallocFailures
    leafs["path-malloc-failures"] = statistics.PathMallocFailures
    leafs["rib-updates"] = statistics.RibUpdates
    return leafs
}

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetYangName() string { return "statistics" }

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Rip_Protocol_DefaultVrf_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Rip_Protocol_DefaultVrf_Statistics) GetParentYangName() string { return "default-vrf" }

// Rip_Protocol_DefaultVrf_Interfaces
// RIP interfaces
type Rip_Protocol_DefaultVrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular RIP interface. The type is slice of
    // Rip_Protocol_DefaultVrf_Interfaces_Interface.
    Interface []Rip_Protocol_DefaultVrf_Interfaces_Interface
}

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Protocol_DefaultVrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Rip_Protocol_DefaultVrf_Interfaces) GetParentYangName() string { return "default-vrf" }

// Rip_Protocol_DefaultVrf_Interfaces_Interface
// Information about a particular RIP interface
type Rip_Protocol_DefaultVrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    Interface interface{}

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    IfHandle interface{}

    // Whether RIP is enabled on this interface. The type is bool.
    RipEnabled interface{}

    // Passive interface indicator. The type is bool.
    IsPassiveInterface interface{}

    // Use broadcast address for v2 packets. The type is bool.
    MulticastAddress interface{}

    // Accept routes of metric 0 indicator. The type is bool.
    AcceptMetric interface{}

    // Versions that the interface is sending. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // Versions that the interface will recieve. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Current state of the interface. The type is InterfaceState.
    State interface{}

    // IP Address of this interface. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Prefix length of the IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Cost added to routes through this interface. The type is interface{} with
    // range: 0..4294967295.
    MetricCost interface{}

    // Split horizon enabled indicator. The type is bool.
    SplitHorizon interface{}

    // Poisoned reverse enabled indicator. The type is bool.
    PoisonHorizon interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Interface's triggered RIP neighbor. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // Multicast group join status. The type is bool.
    JoinStatus interface{}

    // LPTSState. The type is bool.
    LptsState interface{}

    // Authentication Mode. The type is interface{} with range: 0..4294967295.
    AuthMode interface{}

    // Authentication Keychain Name. The type is string.
    AuthKeychain interface{}

    // Authentication send key exists. The type is bool.
    SendAuthKeyExists interface{}

    // Authentication key programmed with MD5 algorithm. The type is bool.
    AuthKeyMd5 interface{}

    // Current active Send Authentication Key Id. The type is interface{} with
    // range: 0..18446744073709551615.
    AuthKeySendId interface{}

    // Total packets received. The type is interface{} with range: 0..4294967295.
    TotalPktRecvd interface{}

    // Packets dropped due to wrong keychain configured. The type is interface{}
    // with range: 0..4294967295.
    PktDropWrongKc interface{}

    // Packets dropped due to missing authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropNoAuth interface{}

    // Packets dropped due to invalid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropInvalidAuth interface{}

    // Packets accepted with valid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktAcceptedValidAuth interface{}

    // User defined summary addresses. The type is slice of
    // Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary.
    RipSummary []Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary

    // Neighbors on this interface. The type is slice of
    // Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer.
    RipPeer []Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer
}

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "rip-enabled" { return "RipEnabled" }
    if yname == "is-passive-interface" { return "IsPassiveInterface" }
    if yname == "multicast-address" { return "MulticastAddress" }
    if yname == "accept-metric" { return "AcceptMetric" }
    if yname == "send-version" { return "SendVersion" }
    if yname == "receive-version" { return "ReceiveVersion" }
    if yname == "state" { return "State" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "metric-cost" { return "MetricCost" }
    if yname == "split-horizon" { return "SplitHorizon" }
    if yname == "poison-horizon" { return "PoisonHorizon" }
    if yname == "triggered-rip" { return "TriggeredRip" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "join-status" { return "JoinStatus" }
    if yname == "lpts-state" { return "LptsState" }
    if yname == "auth-mode" { return "AuthMode" }
    if yname == "auth-keychain" { return "AuthKeychain" }
    if yname == "send-auth-key-exists" { return "SendAuthKeyExists" }
    if yname == "auth-key-md5" { return "AuthKeyMd5" }
    if yname == "auth-key-send-id" { return "AuthKeySendId" }
    if yname == "total-pkt-recvd" { return "TotalPktRecvd" }
    if yname == "pkt-drop-wrong-kc" { return "PktDropWrongKc" }
    if yname == "pkt-drop-no-auth" { return "PktDropNoAuth" }
    if yname == "pkt-drop-invalid-auth" { return "PktDropInvalidAuth" }
    if yname == "pkt-accepted-valid-auth" { return "PktAcceptedValidAuth" }
    if yname == "rip-summary" { return "RipSummary" }
    if yname == "rip-peer" { return "RipPeer" }
    return ""
}

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rip-summary" {
        for _, c := range self.RipSummary {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary{}
        self.RipSummary = append(self.RipSummary, child)
        return &self.RipSummary[len(self.RipSummary)-1]
    }
    if childYangName == "rip-peer" {
        for _, c := range self.RipPeer {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer{}
        self.RipPeer = append(self.RipPeer, child)
        return &self.RipPeer[len(self.RipPeer)-1]
    }
    return nil
}

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.RipSummary {
        children[self.RipSummary[i].GetSegmentPath()] = &self.RipSummary[i]
    }
    for i := range self.RipPeer {
        children[self.RipPeer[i].GetSegmentPath()] = &self.RipPeer[i]
    }
    return children
}

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["if-handle"] = self.IfHandle
    leafs["rip-enabled"] = self.RipEnabled
    leafs["is-passive-interface"] = self.IsPassiveInterface
    leafs["multicast-address"] = self.MulticastAddress
    leafs["accept-metric"] = self.AcceptMetric
    leafs["send-version"] = self.SendVersion
    leafs["receive-version"] = self.ReceiveVersion
    leafs["state"] = self.State
    leafs["destination-address"] = self.DestinationAddress
    leafs["prefix-length"] = self.PrefixLength
    leafs["metric-cost"] = self.MetricCost
    leafs["split-horizon"] = self.SplitHorizon
    leafs["poison-horizon"] = self.PoisonHorizon
    leafs["triggered-rip"] = self.TriggeredRip
    leafs["neighbor-address"] = self.NeighborAddress
    leafs["oom-flags"] = self.OomFlags
    leafs["join-status"] = self.JoinStatus
    leafs["lpts-state"] = self.LptsState
    leafs["auth-mode"] = self.AuthMode
    leafs["auth-keychain"] = self.AuthKeychain
    leafs["send-auth-key-exists"] = self.SendAuthKeyExists
    leafs["auth-key-md5"] = self.AuthKeyMd5
    leafs["auth-key-send-id"] = self.AuthKeySendId
    leafs["total-pkt-recvd"] = self.TotalPktRecvd
    leafs["pkt-drop-wrong-kc"] = self.PktDropWrongKc
    leafs["pkt-drop-no-auth"] = self.PktDropNoAuth
    leafs["pkt-drop-invalid-auth"] = self.PktDropInvalidAuth
    leafs["pkt-accepted-valid-auth"] = self.PktAcceptedValidAuth
    return leafs
}

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Rip_Protocol_DefaultVrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary
// User defined summary addresses
type Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary address prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Summary address prefix length. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Summary address next hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Summary metric. The type is interface{} with range:
    // -2147483648..2147483647.
    Metric interface{}
}

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetFilter() yfilter.YFilter { return ripSummary.YFilter }

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) SetFilter(yf yfilter.YFilter) { ripSummary.YFilter = yf }

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetSegmentPath() string {
    return "rip-summary"
}

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = ripSummary.Prefix
    leafs["prefix-length"] = ripSummary.PrefixLength
    leafs["next-hop-address"] = ripSummary.NextHopAddress
    leafs["metric"] = ripSummary.Metric
    return leafs
}

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetYangName() string { return "rip-summary" }

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) SetParent(parent types.Entity) { ripSummary.parent = parent }

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetParent() types.Entity { return ripSummary.parent }

func (ripSummary *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipSummary) GetParentYangName() string { return "interface" }

// Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer
// Neighbors on this interface
type Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Uptime of this peer. The type is interface{} with range: 0..4294967295.
    PeerUptime interface{}

    // IP Address of this peer. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // RIP version for this peer. The type is interface{} with range: 0..255.
    PeerVersion interface{}

    // Discarded packets from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerPackets interface{}

    // Discarded routes from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerRoutes interface{}
}

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetFilter() yfilter.YFilter { return ripPeer.YFilter }

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) SetFilter(yf yfilter.YFilter) { ripPeer.YFilter = yf }

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetGoName(yname string) string {
    if yname == "peer-uptime" { return "PeerUptime" }
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "peer-version" { return "PeerVersion" }
    if yname == "discarded-peer-packets" { return "DiscardedPeerPackets" }
    if yname == "discarded-peer-routes" { return "DiscardedPeerRoutes" }
    return ""
}

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetSegmentPath() string {
    return "rip-peer"
}

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-uptime"] = ripPeer.PeerUptime
    leafs["peer-address"] = ripPeer.PeerAddress
    leafs["peer-version"] = ripPeer.PeerVersion
    leafs["discarded-peer-packets"] = ripPeer.DiscardedPeerPackets
    leafs["discarded-peer-routes"] = ripPeer.DiscardedPeerRoutes
    return leafs
}

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetBundleName() string { return "cisco_ios_xr" }

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetYangName() string { return "rip-peer" }

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) SetParent(parent types.Entity) { ripPeer.parent = parent }

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetParent() types.Entity { return ripPeer.parent }

func (ripPeer *Rip_Protocol_DefaultVrf_Interfaces_Interface_RipPeer) GetParentYangName() string { return "interface" }

// Rip_Protocol_DefaultVrf_Global
// Global Information 
type Rip_Protocol_DefaultVrf_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF summary data.
    VrfSummary Rip_Protocol_DefaultVrf_Global_VrfSummary

    // List of Interfaces configured. The type is slice of
    // Rip_Protocol_DefaultVrf_Global_InterfaceSummary.
    InterfaceSummary []Rip_Protocol_DefaultVrf_Global_InterfaceSummary
}

func (global *Rip_Protocol_DefaultVrf_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Rip_Protocol_DefaultVrf_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Rip_Protocol_DefaultVrf_Global) GetGoName(yname string) string {
    if yname == "vrf-summary" { return "VrfSummary" }
    if yname == "interface-summary" { return "InterfaceSummary" }
    return ""
}

func (global *Rip_Protocol_DefaultVrf_Global) GetSegmentPath() string {
    return "global"
}

func (global *Rip_Protocol_DefaultVrf_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-summary" {
        return &global.VrfSummary
    }
    if childYangName == "interface-summary" {
        for _, c := range global.InterfaceSummary {
            if global.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_Protocol_DefaultVrf_Global_InterfaceSummary{}
        global.InterfaceSummary = append(global.InterfaceSummary, child)
        return &global.InterfaceSummary[len(global.InterfaceSummary)-1]
    }
    return nil
}

func (global *Rip_Protocol_DefaultVrf_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf-summary"] = &global.VrfSummary
    for i := range global.InterfaceSummary {
        children[global.InterfaceSummary[i].GetSegmentPath()] = &global.InterfaceSummary[i]
    }
    return children
}

func (global *Rip_Protocol_DefaultVrf_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *Rip_Protocol_DefaultVrf_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *Rip_Protocol_DefaultVrf_Global) GetYangName() string { return "global" }

func (global *Rip_Protocol_DefaultVrf_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *Rip_Protocol_DefaultVrf_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *Rip_Protocol_DefaultVrf_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *Rip_Protocol_DefaultVrf_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Rip_Protocol_DefaultVrf_Global) GetParent() types.Entity { return global.parent }

func (global *Rip_Protocol_DefaultVrf_Global) GetParentYangName() string { return "default-vrf" }

// Rip_Protocol_DefaultVrf_Global_VrfSummary
// VRF summary data
type Rip_Protocol_DefaultVrf_Global_VrfSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Active indicator. The type is bool.
    Active interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Number of interfaces configured. The type is interface{} with range:
    // 0..4294967295.
    InterfaceConfiguredCount interface{}

    // Number of active interfaces. The type is interface{} with range:
    // 0..4294967295.
    ActiveInterfaceCount interface{}
}

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetFilter() yfilter.YFilter { return vrfSummary.YFilter }

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) SetFilter(yf yfilter.YFilter) { vrfSummary.YFilter = yf }

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "active" { return "Active" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "route-count" { return "RouteCount" }
    if yname == "path-count" { return "PathCount" }
    if yname == "update-timer" { return "UpdateTimer" }
    if yname == "next-update-time" { return "NextUpdateTime" }
    if yname == "invalid-timer" { return "InvalidTimer" }
    if yname == "hold-down-timer" { return "HoldDownTimer" }
    if yname == "flush-timer" { return "FlushTimer" }
    if yname == "interface-configured-count" { return "InterfaceConfiguredCount" }
    if yname == "active-interface-count" { return "ActiveInterfaceCount" }
    return ""
}

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetSegmentPath() string {
    return "vrf-summary"
}

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrfSummary.VrfName
    leafs["active"] = vrfSummary.Active
    leafs["oom-flags"] = vrfSummary.OomFlags
    leafs["route-count"] = vrfSummary.RouteCount
    leafs["path-count"] = vrfSummary.PathCount
    leafs["update-timer"] = vrfSummary.UpdateTimer
    leafs["next-update-time"] = vrfSummary.NextUpdateTime
    leafs["invalid-timer"] = vrfSummary.InvalidTimer
    leafs["hold-down-timer"] = vrfSummary.HoldDownTimer
    leafs["flush-timer"] = vrfSummary.FlushTimer
    leafs["interface-configured-count"] = vrfSummary.InterfaceConfiguredCount
    leafs["active-interface-count"] = vrfSummary.ActiveInterfaceCount
    return leafs
}

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetBundleName() string { return "cisco_ios_xr" }

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetYangName() string { return "vrf-summary" }

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) SetParent(parent types.Entity) { vrfSummary.parent = parent }

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetParent() types.Entity { return vrfSummary.parent }

func (vrfSummary *Rip_Protocol_DefaultVrf_Global_VrfSummary) GetParentYangName() string { return "global" }

// Rip_Protocol_DefaultVrf_Global_InterfaceSummary
// List of Interfaces configured
type Rip_Protocol_DefaultVrf_Global_InterfaceSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // RIP enabled indicator. The type is bool.
    Enabled interface{}

    // Interface state. The type is InterfaceState.
    State interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // RIP versions this interface sends out. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // RIP versions this interface will receive. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Number of neighbors seen. The type is interface{} with range:
    // 0..4294967295.
    NeighborCount interface{}
}

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetFilter() yfilter.YFilter { return interfaceSummary.YFilter }

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) SetFilter(yf yfilter.YFilter) { interfaceSummary.YFilter = yf }

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enabled" { return "Enabled" }
    if yname == "state" { return "State" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "send-version" { return "SendVersion" }
    if yname == "receive-version" { return "ReceiveVersion" }
    if yname == "neighbor-count" { return "NeighborCount" }
    return ""
}

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetSegmentPath() string {
    return "interface-summary"
}

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceSummary.InterfaceName
    leafs["enabled"] = interfaceSummary.Enabled
    leafs["state"] = interfaceSummary.State
    leafs["destination-address"] = interfaceSummary.DestinationAddress
    leafs["prefix-length"] = interfaceSummary.PrefixLength
    leafs["oom-flags"] = interfaceSummary.OomFlags
    leafs["send-version"] = interfaceSummary.SendVersion
    leafs["receive-version"] = interfaceSummary.ReceiveVersion
    leafs["neighbor-count"] = interfaceSummary.NeighborCount
    return leafs
}

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetYangName() string { return "interface-summary" }

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) SetParent(parent types.Entity) { interfaceSummary.parent = parent }

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetParent() types.Entity { return interfaceSummary.parent }

func (interfaceSummary *Rip_Protocol_DefaultVrf_Global_InterfaceSummary) GetParentYangName() string { return "global" }

// Rip_DefaultVrf
// RIP operational data for Default VRF
type Rip_DefaultVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RIP route database.
    Routes Rip_DefaultVrf_Routes

    // RIP global configuration.
    Configuration Rip_DefaultVrf_Configuration

    // RIP statistics information.
    Statistics Rip_DefaultVrf_Statistics

    // RIP interfaces.
    Interfaces Rip_DefaultVrf_Interfaces

    // Global Information .
    Global Rip_DefaultVrf_Global
}

func (defaultVrf *Rip_DefaultVrf) GetFilter() yfilter.YFilter { return defaultVrf.YFilter }

func (defaultVrf *Rip_DefaultVrf) SetFilter(yf yfilter.YFilter) { defaultVrf.YFilter = yf }

func (defaultVrf *Rip_DefaultVrf) GetGoName(yname string) string {
    if yname == "routes" { return "Routes" }
    if yname == "configuration" { return "Configuration" }
    if yname == "statistics" { return "Statistics" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "global" { return "Global" }
    return ""
}

func (defaultVrf *Rip_DefaultVrf) GetSegmentPath() string {
    return "default-vrf"
}

func (defaultVrf *Rip_DefaultVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routes" {
        return &defaultVrf.Routes
    }
    if childYangName == "configuration" {
        return &defaultVrf.Configuration
    }
    if childYangName == "statistics" {
        return &defaultVrf.Statistics
    }
    if childYangName == "interfaces" {
        return &defaultVrf.Interfaces
    }
    if childYangName == "global" {
        return &defaultVrf.Global
    }
    return nil
}

func (defaultVrf *Rip_DefaultVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routes"] = &defaultVrf.Routes
    children["configuration"] = &defaultVrf.Configuration
    children["statistics"] = &defaultVrf.Statistics
    children["interfaces"] = &defaultVrf.Interfaces
    children["global"] = &defaultVrf.Global
    return children
}

func (defaultVrf *Rip_DefaultVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultVrf *Rip_DefaultVrf) GetBundleName() string { return "cisco_ios_xr" }

func (defaultVrf *Rip_DefaultVrf) GetYangName() string { return "default-vrf" }

func (defaultVrf *Rip_DefaultVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultVrf *Rip_DefaultVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultVrf *Rip_DefaultVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultVrf *Rip_DefaultVrf) SetParent(parent types.Entity) { defaultVrf.parent = parent }

func (defaultVrf *Rip_DefaultVrf) GetParent() types.Entity { return defaultVrf.parent }

func (defaultVrf *Rip_DefaultVrf) GetParentYangName() string { return "rip" }

// Rip_DefaultVrf_Routes
// RIP route database
type Rip_DefaultVrf_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A route in the RIP database. The type is slice of
    // Rip_DefaultVrf_Routes_Route.
    Route []Rip_DefaultVrf_Routes_Route
}

func (routes *Rip_DefaultVrf_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *Rip_DefaultVrf_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *Rip_DefaultVrf_Routes) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (routes *Rip_DefaultVrf_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *Rip_DefaultVrf_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        for _, c := range routes.Route {
            if routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Routes_Route{}
        routes.Route = append(routes.Route, child)
        return &routes.Route[len(routes.Route)-1]
    }
    return nil
}

func (routes *Rip_DefaultVrf_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routes.Route {
        children[routes.Route[i].GetSegmentPath()] = &routes.Route[i]
    }
    return children
}

func (routes *Rip_DefaultVrf_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routes *Rip_DefaultVrf_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *Rip_DefaultVrf_Routes) GetYangName() string { return "routes" }

func (routes *Rip_DefaultVrf_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *Rip_DefaultVrf_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *Rip_DefaultVrf_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *Rip_DefaultVrf_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *Rip_DefaultVrf_Routes) GetParent() types.Entity { return routes.parent }

func (routes *Rip_DefaultVrf_Routes) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_Routes_Route
// A route in the RIP database
type Rip_DefaultVrf_Routes_Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // Destination IP Address for this route. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLengthXr interface{}

    // Route administrative distance. The type is interface{} with range:
    // 0..65535.
    Distance interface{}

    // Hop count for this route. The type is interface{} with range: 0..65535.
    BgpCount interface{}

    // Type of this route. The type is interface{} with range: 0..65535.
    RouteType interface{}

    // Summary route placeholder indicator. The type is bool.
    RouteSummary interface{}

    // Generic route information. The type is interface{} with range: 0..65535.
    RouteTag interface{}

    // RIB supplied version number. The type is interface{} with range: 0..255.
    Version interface{}

    // RIB supplied route attributes. The type is interface{} with range:
    // 0..4294967295.
    Attributes interface{}

    // Active route indicator. The type is bool.
    Active interface{}

    // Where this route was learnt. The type is RipRouteOrigin.
    PathOrigin interface{}

    // Indicates whether route is in holddown. The type is bool.
    HoldDown interface{}

    // The paths for this route. The type is slice of
    // Rip_DefaultVrf_Routes_Route_Paths.
    Paths []Rip_DefaultVrf_Routes_Route_Paths
}

func (route *Rip_DefaultVrf_Routes_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *Rip_DefaultVrf_Routes_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *Rip_DefaultVrf_Routes_Route) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "prefix-length-xr" { return "PrefixLengthXr" }
    if yname == "distance" { return "Distance" }
    if yname == "bgp-count" { return "BgpCount" }
    if yname == "route-type" { return "RouteType" }
    if yname == "route-summary" { return "RouteSummary" }
    if yname == "route-tag" { return "RouteTag" }
    if yname == "version" { return "Version" }
    if yname == "attributes" { return "Attributes" }
    if yname == "active" { return "Active" }
    if yname == "path-origin" { return "PathOrigin" }
    if yname == "hold-down" { return "HoldDown" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (route *Rip_DefaultVrf_Routes_Route) GetSegmentPath() string {
    return "route"
}

func (route *Rip_DefaultVrf_Routes_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "paths" {
        for _, c := range route.Paths {
            if route.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Routes_Route_Paths{}
        route.Paths = append(route.Paths, child)
        return &route.Paths[len(route.Paths)-1]
    }
    return nil
}

func (route *Rip_DefaultVrf_Routes_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range route.Paths {
        children[route.Paths[i].GetSegmentPath()] = &route.Paths[i]
    }
    return children
}

func (route *Rip_DefaultVrf_Routes_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = route.Prefix
    leafs["prefix-length"] = route.PrefixLength
    leafs["destination-address"] = route.DestinationAddress
    leafs["prefix-length-xr"] = route.PrefixLengthXr
    leafs["distance"] = route.Distance
    leafs["bgp-count"] = route.BgpCount
    leafs["route-type"] = route.RouteType
    leafs["route-summary"] = route.RouteSummary
    leafs["route-tag"] = route.RouteTag
    leafs["version"] = route.Version
    leafs["attributes"] = route.Attributes
    leafs["active"] = route.Active
    leafs["path-origin"] = route.PathOrigin
    leafs["hold-down"] = route.HoldDown
    return leafs
}

func (route *Rip_DefaultVrf_Routes_Route) GetBundleName() string { return "cisco_ios_xr" }

func (route *Rip_DefaultVrf_Routes_Route) GetYangName() string { return "route" }

func (route *Rip_DefaultVrf_Routes_Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (route *Rip_DefaultVrf_Routes_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (route *Rip_DefaultVrf_Routes_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (route *Rip_DefaultVrf_Routes_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *Rip_DefaultVrf_Routes_Route) GetParent() types.Entity { return route.parent }

func (route *Rip_DefaultVrf_Routes_Route) GetParentYangName() string { return "routes" }

// Rip_DefaultVrf_Routes_Route_Paths
// The paths for this route
type Rip_DefaultVrf_Routes_Route_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Next hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Metric. The type is interface{} with range: 0..65535.
    Metric interface{}

    // Tag. The type is interface{} with range: 0..65535.
    Tag interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Up time. The type is interface{} with range: 0..4294967295.
    Uptime interface{}

    // Permanent indicator. The type is bool.
    IsPermanent interface{}
}

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *Rip_DefaultVrf_Routes_Route_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "metric" { return "Metric" }
    if yname == "tag" { return "Tag" }
    if yname == "interface" { return "Interface" }
    if yname == "uptime" { return "Uptime" }
    if yname == "is-permanent" { return "IsPermanent" }
    return ""
}

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = paths.SourceAddress
    leafs["next-hop-address"] = paths.NextHopAddress
    leafs["metric"] = paths.Metric
    leafs["tag"] = paths.Tag
    leafs["interface"] = paths.Interface
    leafs["uptime"] = paths.Uptime
    leafs["is-permanent"] = paths.IsPermanent
    return leafs
}

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetYangName() string { return "paths" }

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *Rip_DefaultVrf_Routes_Route_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetParent() types.Entity { return paths.parent }

func (paths *Rip_DefaultVrf_Routes_Route_Paths) GetParentYangName() string { return "route" }

// Rip_DefaultVrf_Configuration
// RIP global configuration
type Rip_DefaultVrf_Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF active indicator. The type is bool.
    Active interface{}

    // VRF added to socket indicator. The type is bool.
    VrFisedSocket interface{}

    // Version of RIP configured. The type is interface{} with range:
    // -2147483648..2147483647.
    RipVersion interface{}

    // Default metric for RIP routes. The type is interface{} with range: 0..255.
    DefaultMetric interface{}

    // Maximum number of paths a route can have. The type is interface{} with
    // range: 0..255.
    MaximumPaths interface{}

    // Auto-summarization indicator. The type is bool.
    AutoSummarize interface{}

    // Use broadcast/multicast address indicator. The type is bool.
    MulticastAddress interface{}

    // Flash update threshold. The type is interface{} with range: 0..255.
    FlashThreshold interface{}

    // Length of the input queue. The type is interface{} with range: 0..65535.
    InputQLength interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Incoming packet source validation indicator. The type is bool.
    ValidationIndicator interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // NSF Enable status. The type is bool.
    NsfStatus interface{}

    // NSF life time. The type is interface{} with range: 0..4294967295.
    NsfLifeTime interface{}
}

func (configuration *Rip_DefaultVrf_Configuration) GetFilter() yfilter.YFilter { return configuration.YFilter }

func (configuration *Rip_DefaultVrf_Configuration) SetFilter(yf yfilter.YFilter) { configuration.YFilter = yf }

func (configuration *Rip_DefaultVrf_Configuration) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    if yname == "vr-fised-socket" { return "VrFisedSocket" }
    if yname == "rip-version" { return "RipVersion" }
    if yname == "default-metric" { return "DefaultMetric" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    if yname == "auto-summarize" { return "AutoSummarize" }
    if yname == "multicast-address" { return "MulticastAddress" }
    if yname == "flash-threshold" { return "FlashThreshold" }
    if yname == "input-q-length" { return "InputQLength" }
    if yname == "triggered-rip" { return "TriggeredRip" }
    if yname == "validation-indicator" { return "ValidationIndicator" }
    if yname == "update-timer" { return "UpdateTimer" }
    if yname == "next-update-time" { return "NextUpdateTime" }
    if yname == "invalid-timer" { return "InvalidTimer" }
    if yname == "hold-down-timer" { return "HoldDownTimer" }
    if yname == "flush-timer" { return "FlushTimer" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "nsf-status" { return "NsfStatus" }
    if yname == "nsf-life-time" { return "NsfLifeTime" }
    return ""
}

func (configuration *Rip_DefaultVrf_Configuration) GetSegmentPath() string {
    return "configuration"
}

func (configuration *Rip_DefaultVrf_Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuration *Rip_DefaultVrf_Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuration *Rip_DefaultVrf_Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active"] = configuration.Active
    leafs["vr-fised-socket"] = configuration.VrFisedSocket
    leafs["rip-version"] = configuration.RipVersion
    leafs["default-metric"] = configuration.DefaultMetric
    leafs["maximum-paths"] = configuration.MaximumPaths
    leafs["auto-summarize"] = configuration.AutoSummarize
    leafs["multicast-address"] = configuration.MulticastAddress
    leafs["flash-threshold"] = configuration.FlashThreshold
    leafs["input-q-length"] = configuration.InputQLength
    leafs["triggered-rip"] = configuration.TriggeredRip
    leafs["validation-indicator"] = configuration.ValidationIndicator
    leafs["update-timer"] = configuration.UpdateTimer
    leafs["next-update-time"] = configuration.NextUpdateTime
    leafs["invalid-timer"] = configuration.InvalidTimer
    leafs["hold-down-timer"] = configuration.HoldDownTimer
    leafs["flush-timer"] = configuration.FlushTimer
    leafs["oom-flags"] = configuration.OomFlags
    leafs["nsf-status"] = configuration.NsfStatus
    leafs["nsf-life-time"] = configuration.NsfLifeTime
    return leafs
}

func (configuration *Rip_DefaultVrf_Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (configuration *Rip_DefaultVrf_Configuration) GetYangName() string { return "configuration" }

func (configuration *Rip_DefaultVrf_Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuration *Rip_DefaultVrf_Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuration *Rip_DefaultVrf_Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuration *Rip_DefaultVrf_Configuration) SetParent(parent types.Entity) { configuration.parent = parent }

func (configuration *Rip_DefaultVrf_Configuration) GetParent() types.Entity { return configuration.parent }

func (configuration *Rip_DefaultVrf_Configuration) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_Statistics
// RIP statistics information
type Rip_DefaultVrf_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total packets received. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Total discarded packets. The type is interface{} with range: 0..4294967295.
    DiscardedPackets interface{}

    // Total discarded routes. The type is interface{} with range: 0..4294967295.
    DiscardedRoutes interface{}

    // Packets rx in SRP. The type is interface{} with range: 0..4294967295.
    StandbyPacketsReceived interface{}

    // Number of messages sent. The type is interface{} with range: 0..4294967295.
    SentMessages interface{}

    // Number of message send failures. The type is interface{} with range:
    // 0..4294967295.
    SentMessageFailures interface{}

    // Number of RIP queries responded to. The type is interface{} with range:
    // 0..4294967295.
    QueryResponses interface{}

    // Number of periodic RIP updates. The type is interface{} with range:
    // 0..4294967295.
    PeriodicUpdates interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Number of failures to allocate memory for a route. The type is interface{}
    // with range: 0..4294967295.
    RouteMallocFailures interface{}

    // Number of failures to allocate memory for a path. The type is interface{}
    // with range: 0..4294967295.
    PathMallocFailures interface{}

    // Number of route updates to RIB made by RIP. The type is interface{} with
    // range: 0..4294967295.
    RibUpdates interface{}
}

func (statistics *Rip_DefaultVrf_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Rip_DefaultVrf_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Rip_DefaultVrf_Statistics) GetGoName(yname string) string {
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "discarded-packets" { return "DiscardedPackets" }
    if yname == "discarded-routes" { return "DiscardedRoutes" }
    if yname == "standby-packets-received" { return "StandbyPacketsReceived" }
    if yname == "sent-messages" { return "SentMessages" }
    if yname == "sent-message-failures" { return "SentMessageFailures" }
    if yname == "query-responses" { return "QueryResponses" }
    if yname == "periodic-updates" { return "PeriodicUpdates" }
    if yname == "route-count" { return "RouteCount" }
    if yname == "path-count" { return "PathCount" }
    if yname == "route-malloc-failures" { return "RouteMallocFailures" }
    if yname == "path-malloc-failures" { return "PathMallocFailures" }
    if yname == "rib-updates" { return "RibUpdates" }
    return ""
}

func (statistics *Rip_DefaultVrf_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Rip_DefaultVrf_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Rip_DefaultVrf_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Rip_DefaultVrf_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["discarded-packets"] = statistics.DiscardedPackets
    leafs["discarded-routes"] = statistics.DiscardedRoutes
    leafs["standby-packets-received"] = statistics.StandbyPacketsReceived
    leafs["sent-messages"] = statistics.SentMessages
    leafs["sent-message-failures"] = statistics.SentMessageFailures
    leafs["query-responses"] = statistics.QueryResponses
    leafs["periodic-updates"] = statistics.PeriodicUpdates
    leafs["route-count"] = statistics.RouteCount
    leafs["path-count"] = statistics.PathCount
    leafs["route-malloc-failures"] = statistics.RouteMallocFailures
    leafs["path-malloc-failures"] = statistics.PathMallocFailures
    leafs["rib-updates"] = statistics.RibUpdates
    return leafs
}

func (statistics *Rip_DefaultVrf_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Rip_DefaultVrf_Statistics) GetYangName() string { return "statistics" }

func (statistics *Rip_DefaultVrf_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Rip_DefaultVrf_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Rip_DefaultVrf_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Rip_DefaultVrf_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Rip_DefaultVrf_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Rip_DefaultVrf_Statistics) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_Interfaces
// RIP interfaces
type Rip_DefaultVrf_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular RIP interface. The type is slice of
    // Rip_DefaultVrf_Interfaces_Interface.
    Interface []Rip_DefaultVrf_Interfaces_Interface
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Rip_DefaultVrf_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Rip_DefaultVrf_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Rip_DefaultVrf_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Rip_DefaultVrf_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Rip_DefaultVrf_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Rip_DefaultVrf_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Rip_DefaultVrf_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Rip_DefaultVrf_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Rip_DefaultVrf_Interfaces) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_Interfaces_Interface
// Information about a particular RIP interface
type Rip_DefaultVrf_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    Interface interface{}

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    IfHandle interface{}

    // Whether RIP is enabled on this interface. The type is bool.
    RipEnabled interface{}

    // Passive interface indicator. The type is bool.
    IsPassiveInterface interface{}

    // Use broadcast address for v2 packets. The type is bool.
    MulticastAddress interface{}

    // Accept routes of metric 0 indicator. The type is bool.
    AcceptMetric interface{}

    // Versions that the interface is sending. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // Versions that the interface will recieve. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Current state of the interface. The type is InterfaceState.
    State interface{}

    // IP Address of this interface. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Prefix length of the IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Cost added to routes through this interface. The type is interface{} with
    // range: 0..4294967295.
    MetricCost interface{}

    // Split horizon enabled indicator. The type is bool.
    SplitHorizon interface{}

    // Poisoned reverse enabled indicator. The type is bool.
    PoisonHorizon interface{}

    // Triggered RIP enabled indicator. The type is bool.
    TriggeredRip interface{}

    // Interface's triggered RIP neighbor. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborAddress interface{}

    // Out-of-memory status flags. The type is interface{} with range:
    // 0..4294967295.
    OomFlags interface{}

    // Multicast group join status. The type is bool.
    JoinStatus interface{}

    // LPTSState. The type is bool.
    LptsState interface{}

    // Authentication Mode. The type is interface{} with range: 0..4294967295.
    AuthMode interface{}

    // Authentication Keychain Name. The type is string.
    AuthKeychain interface{}

    // Authentication send key exists. The type is bool.
    SendAuthKeyExists interface{}

    // Authentication key programmed with MD5 algorithm. The type is bool.
    AuthKeyMd5 interface{}

    // Current active Send Authentication Key Id. The type is interface{} with
    // range: 0..18446744073709551615.
    AuthKeySendId interface{}

    // Total packets received. The type is interface{} with range: 0..4294967295.
    TotalPktRecvd interface{}

    // Packets dropped due to wrong keychain configured. The type is interface{}
    // with range: 0..4294967295.
    PktDropWrongKc interface{}

    // Packets dropped due to missing authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropNoAuth interface{}

    // Packets dropped due to invalid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktDropInvalidAuth interface{}

    // Packets accepted with valid authentication data. The type is interface{}
    // with range: 0..4294967295.
    PktAcceptedValidAuth interface{}

    // User defined summary addresses. The type is slice of
    // Rip_DefaultVrf_Interfaces_Interface_RipSummary.
    RipSummary []Rip_DefaultVrf_Interfaces_Interface_RipSummary

    // Neighbors on this interface. The type is slice of
    // Rip_DefaultVrf_Interfaces_Interface_RipPeer.
    RipPeer []Rip_DefaultVrf_Interfaces_Interface_RipPeer
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Rip_DefaultVrf_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "rip-enabled" { return "RipEnabled" }
    if yname == "is-passive-interface" { return "IsPassiveInterface" }
    if yname == "multicast-address" { return "MulticastAddress" }
    if yname == "accept-metric" { return "AcceptMetric" }
    if yname == "send-version" { return "SendVersion" }
    if yname == "receive-version" { return "ReceiveVersion" }
    if yname == "state" { return "State" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "metric-cost" { return "MetricCost" }
    if yname == "split-horizon" { return "SplitHorizon" }
    if yname == "poison-horizon" { return "PoisonHorizon" }
    if yname == "triggered-rip" { return "TriggeredRip" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "join-status" { return "JoinStatus" }
    if yname == "lpts-state" { return "LptsState" }
    if yname == "auth-mode" { return "AuthMode" }
    if yname == "auth-keychain" { return "AuthKeychain" }
    if yname == "send-auth-key-exists" { return "SendAuthKeyExists" }
    if yname == "auth-key-md5" { return "AuthKeyMd5" }
    if yname == "auth-key-send-id" { return "AuthKeySendId" }
    if yname == "total-pkt-recvd" { return "TotalPktRecvd" }
    if yname == "pkt-drop-wrong-kc" { return "PktDropWrongKc" }
    if yname == "pkt-drop-no-auth" { return "PktDropNoAuth" }
    if yname == "pkt-drop-invalid-auth" { return "PktDropInvalidAuth" }
    if yname == "pkt-accepted-valid-auth" { return "PktAcceptedValidAuth" }
    if yname == "rip-summary" { return "RipSummary" }
    if yname == "rip-peer" { return "RipPeer" }
    return ""
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rip-summary" {
        for _, c := range self.RipSummary {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Interfaces_Interface_RipSummary{}
        self.RipSummary = append(self.RipSummary, child)
        return &self.RipSummary[len(self.RipSummary)-1]
    }
    if childYangName == "rip-peer" {
        for _, c := range self.RipPeer {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Interfaces_Interface_RipPeer{}
        self.RipPeer = append(self.RipPeer, child)
        return &self.RipPeer[len(self.RipPeer)-1]
    }
    return nil
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.RipSummary {
        children[self.RipSummary[i].GetSegmentPath()] = &self.RipSummary[i]
    }
    for i := range self.RipPeer {
        children[self.RipPeer[i].GetSegmentPath()] = &self.RipPeer[i]
    }
    return children
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["if-handle"] = self.IfHandle
    leafs["rip-enabled"] = self.RipEnabled
    leafs["is-passive-interface"] = self.IsPassiveInterface
    leafs["multicast-address"] = self.MulticastAddress
    leafs["accept-metric"] = self.AcceptMetric
    leafs["send-version"] = self.SendVersion
    leafs["receive-version"] = self.ReceiveVersion
    leafs["state"] = self.State
    leafs["destination-address"] = self.DestinationAddress
    leafs["prefix-length"] = self.PrefixLength
    leafs["metric-cost"] = self.MetricCost
    leafs["split-horizon"] = self.SplitHorizon
    leafs["poison-horizon"] = self.PoisonHorizon
    leafs["triggered-rip"] = self.TriggeredRip
    leafs["neighbor-address"] = self.NeighborAddress
    leafs["oom-flags"] = self.OomFlags
    leafs["join-status"] = self.JoinStatus
    leafs["lpts-state"] = self.LptsState
    leafs["auth-mode"] = self.AuthMode
    leafs["auth-keychain"] = self.AuthKeychain
    leafs["send-auth-key-exists"] = self.SendAuthKeyExists
    leafs["auth-key-md5"] = self.AuthKeyMd5
    leafs["auth-key-send-id"] = self.AuthKeySendId
    leafs["total-pkt-recvd"] = self.TotalPktRecvd
    leafs["pkt-drop-wrong-kc"] = self.PktDropWrongKc
    leafs["pkt-drop-no-auth"] = self.PktDropNoAuth
    leafs["pkt-drop-invalid-auth"] = self.PktDropInvalidAuth
    leafs["pkt-accepted-valid-auth"] = self.PktAcceptedValidAuth
    return leafs
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Rip_DefaultVrf_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Rip_DefaultVrf_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Rip_DefaultVrf_Interfaces_Interface_RipSummary
// User defined summary addresses
type Rip_DefaultVrf_Interfaces_Interface_RipSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary address prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Summary address prefix length. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Summary address next hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}

    // Summary metric. The type is interface{} with range:
    // -2147483648..2147483647.
    Metric interface{}
}

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetFilter() yfilter.YFilter { return ripSummary.YFilter }

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) SetFilter(yf yfilter.YFilter) { ripSummary.YFilter = yf }

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetSegmentPath() string {
    return "rip-summary"
}

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = ripSummary.Prefix
    leafs["prefix-length"] = ripSummary.PrefixLength
    leafs["next-hop-address"] = ripSummary.NextHopAddress
    leafs["metric"] = ripSummary.Metric
    return leafs
}

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetYangName() string { return "rip-summary" }

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) SetParent(parent types.Entity) { ripSummary.parent = parent }

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetParent() types.Entity { return ripSummary.parent }

func (ripSummary *Rip_DefaultVrf_Interfaces_Interface_RipSummary) GetParentYangName() string { return "interface" }

// Rip_DefaultVrf_Interfaces_Interface_RipPeer
// Neighbors on this interface
type Rip_DefaultVrf_Interfaces_Interface_RipPeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Uptime of this peer. The type is interface{} with range: 0..4294967295.
    PeerUptime interface{}

    // IP Address of this peer. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // RIP version for this peer. The type is interface{} with range: 0..255.
    PeerVersion interface{}

    // Discarded packets from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerPackets interface{}

    // Discarded routes from this peer. The type is interface{} with range:
    // 0..4294967295.
    DiscardedPeerRoutes interface{}
}

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetFilter() yfilter.YFilter { return ripPeer.YFilter }

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) SetFilter(yf yfilter.YFilter) { ripPeer.YFilter = yf }

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetGoName(yname string) string {
    if yname == "peer-uptime" { return "PeerUptime" }
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "peer-version" { return "PeerVersion" }
    if yname == "discarded-peer-packets" { return "DiscardedPeerPackets" }
    if yname == "discarded-peer-routes" { return "DiscardedPeerRoutes" }
    return ""
}

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetSegmentPath() string {
    return "rip-peer"
}

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-uptime"] = ripPeer.PeerUptime
    leafs["peer-address"] = ripPeer.PeerAddress
    leafs["peer-version"] = ripPeer.PeerVersion
    leafs["discarded-peer-packets"] = ripPeer.DiscardedPeerPackets
    leafs["discarded-peer-routes"] = ripPeer.DiscardedPeerRoutes
    return leafs
}

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetBundleName() string { return "cisco_ios_xr" }

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetYangName() string { return "rip-peer" }

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) SetParent(parent types.Entity) { ripPeer.parent = parent }

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetParent() types.Entity { return ripPeer.parent }

func (ripPeer *Rip_DefaultVrf_Interfaces_Interface_RipPeer) GetParentYangName() string { return "interface" }

// Rip_DefaultVrf_Global
// Global Information 
type Rip_DefaultVrf_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF summary data.
    VrfSummary Rip_DefaultVrf_Global_VrfSummary

    // List of Interfaces configured. The type is slice of
    // Rip_DefaultVrf_Global_InterfaceSummary.
    InterfaceSummary []Rip_DefaultVrf_Global_InterfaceSummary
}

func (global *Rip_DefaultVrf_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Rip_DefaultVrf_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Rip_DefaultVrf_Global) GetGoName(yname string) string {
    if yname == "vrf-summary" { return "VrfSummary" }
    if yname == "interface-summary" { return "InterfaceSummary" }
    return ""
}

func (global *Rip_DefaultVrf_Global) GetSegmentPath() string {
    return "global"
}

func (global *Rip_DefaultVrf_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-summary" {
        return &global.VrfSummary
    }
    if childYangName == "interface-summary" {
        for _, c := range global.InterfaceSummary {
            if global.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Rip_DefaultVrf_Global_InterfaceSummary{}
        global.InterfaceSummary = append(global.InterfaceSummary, child)
        return &global.InterfaceSummary[len(global.InterfaceSummary)-1]
    }
    return nil
}

func (global *Rip_DefaultVrf_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf-summary"] = &global.VrfSummary
    for i := range global.InterfaceSummary {
        children[global.InterfaceSummary[i].GetSegmentPath()] = &global.InterfaceSummary[i]
    }
    return children
}

func (global *Rip_DefaultVrf_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *Rip_DefaultVrf_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *Rip_DefaultVrf_Global) GetYangName() string { return "global" }

func (global *Rip_DefaultVrf_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *Rip_DefaultVrf_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *Rip_DefaultVrf_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *Rip_DefaultVrf_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Rip_DefaultVrf_Global) GetParent() types.Entity { return global.parent }

func (global *Rip_DefaultVrf_Global) GetParentYangName() string { return "default-vrf" }

// Rip_DefaultVrf_Global_VrfSummary
// VRF summary data
type Rip_DefaultVrf_Global_VrfSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF Active indicator. The type is bool.
    Active interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // Number of routes allocated. The type is interface{} with range:
    // 0..4294967295.
    RouteCount interface{}

    // Number of paths allocated. The type is interface{} with range:
    // 0..4294967295.
    PathCount interface{}

    // Update timer. The type is interface{} with range: 0..4294967295.
    UpdateTimer interface{}

    // Time left for next update. The type is interface{} with range:
    // 0..4294967295.
    NextUpdateTime interface{}

    // Invalid timer. The type is interface{} with range: 0..4294967295.
    InvalidTimer interface{}

    // Holddown timer. The type is interface{} with range: 0..4294967295.
    HoldDownTimer interface{}

    // Flush timer. The type is interface{} with range: 0..4294967295.
    FlushTimer interface{}

    // Number of interfaces configured. The type is interface{} with range:
    // 0..4294967295.
    InterfaceConfiguredCount interface{}

    // Number of active interfaces. The type is interface{} with range:
    // 0..4294967295.
    ActiveInterfaceCount interface{}
}

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetFilter() yfilter.YFilter { return vrfSummary.YFilter }

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) SetFilter(yf yfilter.YFilter) { vrfSummary.YFilter = yf }

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "active" { return "Active" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "route-count" { return "RouteCount" }
    if yname == "path-count" { return "PathCount" }
    if yname == "update-timer" { return "UpdateTimer" }
    if yname == "next-update-time" { return "NextUpdateTime" }
    if yname == "invalid-timer" { return "InvalidTimer" }
    if yname == "hold-down-timer" { return "HoldDownTimer" }
    if yname == "flush-timer" { return "FlushTimer" }
    if yname == "interface-configured-count" { return "InterfaceConfiguredCount" }
    if yname == "active-interface-count" { return "ActiveInterfaceCount" }
    return ""
}

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetSegmentPath() string {
    return "vrf-summary"
}

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrfSummary.VrfName
    leafs["active"] = vrfSummary.Active
    leafs["oom-flags"] = vrfSummary.OomFlags
    leafs["route-count"] = vrfSummary.RouteCount
    leafs["path-count"] = vrfSummary.PathCount
    leafs["update-timer"] = vrfSummary.UpdateTimer
    leafs["next-update-time"] = vrfSummary.NextUpdateTime
    leafs["invalid-timer"] = vrfSummary.InvalidTimer
    leafs["hold-down-timer"] = vrfSummary.HoldDownTimer
    leafs["flush-timer"] = vrfSummary.FlushTimer
    leafs["interface-configured-count"] = vrfSummary.InterfaceConfiguredCount
    leafs["active-interface-count"] = vrfSummary.ActiveInterfaceCount
    return leafs
}

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetBundleName() string { return "cisco_ios_xr" }

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetYangName() string { return "vrf-summary" }

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) SetParent(parent types.Entity) { vrfSummary.parent = parent }

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetParent() types.Entity { return vrfSummary.parent }

func (vrfSummary *Rip_DefaultVrf_Global_VrfSummary) GetParentYangName() string { return "global" }

// Rip_DefaultVrf_Global_InterfaceSummary
// List of Interfaces configured
type Rip_DefaultVrf_Global_InterfaceSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // RIP enabled indicator. The type is bool.
    Enabled interface{}

    // Interface state. The type is InterfaceState.
    State interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Prefix length of IP address. The type is interface{} with range:
    // 0..4294967295.
    PrefixLength interface{}

    // Current OOM flags. The type is interface{} with range: 0..4294967295.
    OomFlags interface{}

    // RIP versions this interface sends out. The type is interface{} with range:
    // 0..4294967295.
    SendVersion interface{}

    // RIP versions this interface will receive. The type is interface{} with
    // range: 0..4294967295.
    ReceiveVersion interface{}

    // Number of neighbors seen. The type is interface{} with range:
    // 0..4294967295.
    NeighborCount interface{}
}

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetFilter() yfilter.YFilter { return interfaceSummary.YFilter }

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) SetFilter(yf yfilter.YFilter) { interfaceSummary.YFilter = yf }

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enabled" { return "Enabled" }
    if yname == "state" { return "State" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "oom-flags" { return "OomFlags" }
    if yname == "send-version" { return "SendVersion" }
    if yname == "receive-version" { return "ReceiveVersion" }
    if yname == "neighbor-count" { return "NeighborCount" }
    return ""
}

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetSegmentPath() string {
    return "interface-summary"
}

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceSummary.InterfaceName
    leafs["enabled"] = interfaceSummary.Enabled
    leafs["state"] = interfaceSummary.State
    leafs["destination-address"] = interfaceSummary.DestinationAddress
    leafs["prefix-length"] = interfaceSummary.PrefixLength
    leafs["oom-flags"] = interfaceSummary.OomFlags
    leafs["send-version"] = interfaceSummary.SendVersion
    leafs["receive-version"] = interfaceSummary.ReceiveVersion
    leafs["neighbor-count"] = interfaceSummary.NeighborCount
    return leafs
}

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetYangName() string { return "interface-summary" }

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) SetParent(parent types.Entity) { interfaceSummary.parent = parent }

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetParent() types.Entity { return interfaceSummary.parent }

func (interfaceSummary *Rip_DefaultVrf_Global_InterfaceSummary) GetParentYangName() string { return "global" }

