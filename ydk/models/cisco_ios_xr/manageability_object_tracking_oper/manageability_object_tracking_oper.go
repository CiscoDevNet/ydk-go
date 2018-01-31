// This module contains a collection of YANG definitions
// for Cisco IOS-XR manageability-object-tracking package operational data.
// 
// This module contains definitions
// for the following management objects:
//   object-tracking: Object Tracking operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package manageability_object_tracking_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package manageability_object_tracking_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-manageability-object-tracking-oper object-tracking}", reflect.TypeOf(ObjectTracking{}))
    ydk.RegisterEntity("Cisco-IOS-XR-manageability-object-tracking-oper:object-tracking", reflect.TypeOf(ObjectTracking{}))
}

// Track represents Track
type Track string

const (
    // Line protocol type
    Track_interface_type Track = "interface-type"

    // Route type
    Track_route_type Track = "route-type"

    // Boolean and type
    Track_bool_and_type Track = "bool-and-type"

    // Boolean or type
    Track_bool_or_type Track = "bool-or-type"

    // Ipsla track type
    Track_ipsla_type Track = "ipsla-type"

    // type undefined
    Track_undefined_type Track = "undefined-type"

    // type threshold weight
    Track_threshold_weight Track = "threshold-weight"

    // type threshold percentage
    Track_threshold_percentage Track = "threshold-percentage"

    // type bfd rtr
    Track_bfd_type Track = "bfd-type"
)

// ObjectTracking
// Object Tracking operational data
type ObjectTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object Tracking Type interface info.
    TrackTypeInterface ObjectTracking_TrackTypeInterface

    // Object Tracking Track table brief.
    TrackBriefs ObjectTracking_TrackBriefs

    // Object Tracking Type RTR Reachability info.
    TrackTypeRtrReachability ObjectTracking_TrackTypeRtrReachability

    // Object Tracking Type RTR Reachability brief info.
    TrackTypeRtrReachabilityBrief ObjectTracking_TrackTypeRtrReachabilityBrief

    // Object Tracking Track table.
    Tracks ObjectTracking_Tracks

    // Object Tracking Type Ipv4 Route brief info.
    TrackTypeIpv4RouteBrief ObjectTracking_TrackTypeIpv4RouteBrief

    // Object Tracking Type IPV4 route info.
    TrackTypeIpv4Route ObjectTracking_TrackTypeIpv4Route

    // Object Tracking Type Interface brief info.
    TrackTypeInterfaceBrief ObjectTracking_TrackTypeInterfaceBrief
}

func (objectTracking *ObjectTracking) GetFilter() yfilter.YFilter { return objectTracking.YFilter }

func (objectTracking *ObjectTracking) SetFilter(yf yfilter.YFilter) { objectTracking.YFilter = yf }

func (objectTracking *ObjectTracking) GetGoName(yname string) string {
    if yname == "track-type-interface" { return "TrackTypeInterface" }
    if yname == "track-briefs" { return "TrackBriefs" }
    if yname == "track-type-rtr-reachability" { return "TrackTypeRtrReachability" }
    if yname == "track-type-rtr-reachability-brief" { return "TrackTypeRtrReachabilityBrief" }
    if yname == "tracks" { return "Tracks" }
    if yname == "track-type-ipv4-route-brief" { return "TrackTypeIpv4RouteBrief" }
    if yname == "track-type-ipv4-route" { return "TrackTypeIpv4Route" }
    if yname == "track-type-interface-brief" { return "TrackTypeInterfaceBrief" }
    return ""
}

func (objectTracking *ObjectTracking) GetSegmentPath() string {
    return "Cisco-IOS-XR-manageability-object-tracking-oper:object-tracking"
}

func (objectTracking *ObjectTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-type-interface" {
        return &objectTracking.TrackTypeInterface
    }
    if childYangName == "track-briefs" {
        return &objectTracking.TrackBriefs
    }
    if childYangName == "track-type-rtr-reachability" {
        return &objectTracking.TrackTypeRtrReachability
    }
    if childYangName == "track-type-rtr-reachability-brief" {
        return &objectTracking.TrackTypeRtrReachabilityBrief
    }
    if childYangName == "tracks" {
        return &objectTracking.Tracks
    }
    if childYangName == "track-type-ipv4-route-brief" {
        return &objectTracking.TrackTypeIpv4RouteBrief
    }
    if childYangName == "track-type-ipv4-route" {
        return &objectTracking.TrackTypeIpv4Route
    }
    if childYangName == "track-type-interface-brief" {
        return &objectTracking.TrackTypeInterfaceBrief
    }
    return nil
}

func (objectTracking *ObjectTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-type-interface"] = &objectTracking.TrackTypeInterface
    children["track-briefs"] = &objectTracking.TrackBriefs
    children["track-type-rtr-reachability"] = &objectTracking.TrackTypeRtrReachability
    children["track-type-rtr-reachability-brief"] = &objectTracking.TrackTypeRtrReachabilityBrief
    children["tracks"] = &objectTracking.Tracks
    children["track-type-ipv4-route-brief"] = &objectTracking.TrackTypeIpv4RouteBrief
    children["track-type-ipv4-route"] = &objectTracking.TrackTypeIpv4Route
    children["track-type-interface-brief"] = &objectTracking.TrackTypeInterfaceBrief
    return children
}

func (objectTracking *ObjectTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (objectTracking *ObjectTracking) GetBundleName() string { return "cisco_ios_xr" }

func (objectTracking *ObjectTracking) GetYangName() string { return "object-tracking" }

func (objectTracking *ObjectTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (objectTracking *ObjectTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (objectTracking *ObjectTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (objectTracking *ObjectTracking) SetParent(parent types.Entity) { objectTracking.parent = parent }

func (objectTracking *ObjectTracking) GetParent() types.Entity { return objectTracking.parent }

func (objectTracking *ObjectTracking) GetParentYangName() string { return "Cisco-IOS-XR-manageability-object-tracking-oper" }

// ObjectTracking_TrackTypeInterface
// Object Tracking Type interface info
type ObjectTracking_TrackTypeInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // track info. The type is slice of
    // ObjectTracking_TrackTypeInterface_TrackInfo.
    TrackInfo []ObjectTracking_TrackTypeInterface_TrackInfo
}

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetFilter() yfilter.YFilter { return trackTypeInterface.YFilter }

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) SetFilter(yf yfilter.YFilter) { trackTypeInterface.YFilter = yf }

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetGoName(yname string) string {
    if yname == "track-info" { return "TrackInfo" }
    return ""
}

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetSegmentPath() string {
    return "track-type-interface"
}

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-info" {
        for _, c := range trackTypeInterface.TrackInfo {
            if trackTypeInterface.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeInterface_TrackInfo{}
        trackTypeInterface.TrackInfo = append(trackTypeInterface.TrackInfo, child)
        return &trackTypeInterface.TrackInfo[len(trackTypeInterface.TrackInfo)-1]
    }
    return nil
}

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackTypeInterface.TrackInfo {
        children[trackTypeInterface.TrackInfo[i].GetSegmentPath()] = &trackTypeInterface.TrackInfo[i]
    }
    return children
}

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetYangName() string { return "track-type-interface" }

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) SetParent(parent types.Entity) { trackTypeInterface.parent = parent }

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetParent() types.Entity { return trackTypeInterface.parent }

func (trackTypeInterface *ObjectTracking_TrackTypeInterface) GetParentYangName() string { return "object-tracking" }

// ObjectTracking_TrackTypeInterface_TrackInfo
// track info
type ObjectTracking_TrackTypeInterface_TrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // State Change Counter. The type is interface{} with range: 0..4294967295.
    StateChangeCounter interface{}

    // Seconds Last Change. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecondsLastChange interface{}

    // User specified threshold upper limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdUp interface{}

    // User specified threshold lower limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdDown interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo

    // boolean objects.
    BoolTracks ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks

    // Threshold objects.
    ThresholdTracks ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks

    // Tracking Interfaces.
    TrackingInteraces ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces

    // Is the state change delay counter in progress.
    Delayed ObjectTracking_TrackTypeInterface_TrackInfo_Delayed
}

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetFilter() yfilter.YFilter { return trackInfo.YFilter }

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) SetFilter(yf yfilter.YFilter) { trackInfo.YFilter = yf }

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetGoName(yname string) string {
    if yname == "tracke-name" { return "TrackeName" }
    if yname == "type" { return "Type" }
    if yname == "track-state" { return "TrackState" }
    if yname == "state-change-counter" { return "StateChangeCounter" }
    if yname == "seconds-last-change" { return "SecondsLastChange" }
    if yname == "threshold-up" { return "ThresholdUp" }
    if yname == "threshold-down" { return "ThresholdDown" }
    if yname == "track-type-info" { return "TrackTypeInfo" }
    if yname == "bool-tracks" { return "BoolTracks" }
    if yname == "threshold-tracks" { return "ThresholdTracks" }
    if yname == "tracking-interaces" { return "TrackingInteraces" }
    if yname == "delayed" { return "Delayed" }
    return ""
}

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetSegmentPath() string {
    return "track-info"
}

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-type-info" {
        return &trackInfo.TrackTypeInfo
    }
    if childYangName == "bool-tracks" {
        return &trackInfo.BoolTracks
    }
    if childYangName == "threshold-tracks" {
        return &trackInfo.ThresholdTracks
    }
    if childYangName == "tracking-interaces" {
        return &trackInfo.TrackingInteraces
    }
    if childYangName == "delayed" {
        return &trackInfo.Delayed
    }
    return nil
}

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-type-info"] = &trackInfo.TrackTypeInfo
    children["bool-tracks"] = &trackInfo.BoolTracks
    children["threshold-tracks"] = &trackInfo.ThresholdTracks
    children["tracking-interaces"] = &trackInfo.TrackingInteraces
    children["delayed"] = &trackInfo.Delayed
    return children
}

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tracke-name"] = trackInfo.TrackeName
    leafs["type"] = trackInfo.Type
    leafs["track-state"] = trackInfo.TrackState
    leafs["state-change-counter"] = trackInfo.StateChangeCounter
    leafs["seconds-last-change"] = trackInfo.SecondsLastChange
    leafs["threshold-up"] = trackInfo.ThresholdUp
    leafs["threshold-down"] = trackInfo.ThresholdDown
    return leafs
}

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetYangName() string { return "track-info" }

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) SetParent(parent types.Entity) { trackInfo.parent = parent }

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetParent() types.Entity { return trackInfo.parent }

func (trackInfo *ObjectTracking_TrackTypeInterface_TrackInfo) GetParentYangName() string { return "track-type-interface" }

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetFilter() yfilter.YFilter { return trackTypeInfo.YFilter }

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) SetFilter(yf yfilter.YFilter) { trackTypeInfo.YFilter = yf }

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetGoName(yname string) string {
    if yname == "discriminant" { return "Discriminant" }
    if yname == "interface-tracks" { return "InterfaceTracks" }
    if yname == "route-tracks" { return "RouteTracks" }
    if yname == "ipsla-tracks" { return "IpslaTracks" }
    if yname == "bfd-tracks" { return "BfdTracks" }
    return ""
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetSegmentPath() string {
    return "track-type-info"
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracks" {
        return &trackTypeInfo.InterfaceTracks
    }
    if childYangName == "route-tracks" {
        return &trackTypeInfo.RouteTracks
    }
    if childYangName == "ipsla-tracks" {
        return &trackTypeInfo.IpslaTracks
    }
    if childYangName == "bfd-tracks" {
        return &trackTypeInfo.BfdTracks
    }
    return nil
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-tracks"] = &trackTypeInfo.InterfaceTracks
    children["route-tracks"] = &trackTypeInfo.RouteTracks
    children["ipsla-tracks"] = &trackTypeInfo.IpslaTracks
    children["bfd-tracks"] = &trackTypeInfo.BfdTracks
    return children
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminant"] = trackTypeInfo.Discriminant
    return leafs
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetYangName() string { return "track-type-info" }

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) SetParent(parent types.Entity) { trackTypeInfo.parent = parent }

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetParent() types.Entity { return trackTypeInfo.parent }

func (trackTypeInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetFilter() yfilter.YFilter { return interfaceTracks.YFilter }

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) SetFilter(yf yfilter.YFilter) { interfaceTracks.YFilter = yf }

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetSegmentPath() string {
    return "interface-tracks"
}

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTracks.InterfaceName
    return leafs
}

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetYangName() string { return "interface-tracks" }

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) SetParent(parent types.Entity) { interfaceTracks.parent = parent }

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetParent() types.Entity { return interfaceTracks.parent }

func (interfaceTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_InterfaceTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetFilter() yfilter.YFilter { return routeTracks.YFilter }

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) SetFilter(yf yfilter.YFilter) { routeTracks.YFilter = yf }

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "vrf" { return "Vrf" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetSegmentPath() string {
    return "route-tracks"
}

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = routeTracks.Prefix
    leafs["prefix-length"] = routeTracks.PrefixLength
    leafs["vrf"] = routeTracks.Vrf
    leafs["next-hop"] = routeTracks.NextHop
    return leafs
}

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetBundleName() string { return "cisco_ios_xr" }

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetYangName() string { return "route-tracks" }

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) SetParent(parent types.Entity) { routeTracks.parent = parent }

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetParent() types.Entity { return routeTracks.parent }

func (routeTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_RouteTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetFilter() yfilter.YFilter { return ipslaTracks.YFilter }

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) SetFilter(yf yfilter.YFilter) { ipslaTracks.YFilter = yf }

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetGoName(yname string) string {
    if yname == "ipsla-op-id" { return "IpslaOpId" }
    if yname == "rtt" { return "Rtt" }
    if yname == "return-code" { return "ReturnCode" }
    return ""
}

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetSegmentPath() string {
    return "ipsla-tracks"
}

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipsla-op-id"] = ipslaTracks.IpslaOpId
    leafs["rtt"] = ipslaTracks.Rtt
    leafs["return-code"] = ipslaTracks.ReturnCode
    return leafs
}

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetBundleName() string { return "cisco_ios_xr" }

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetYangName() string { return "ipsla-tracks" }

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) SetParent(parent types.Entity) { ipslaTracks.parent = parent }

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetParent() types.Entity { return ipslaTracks.parent }

func (ipslaTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_IpslaTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetFilter() yfilter.YFilter { return bfdTracks.YFilter }

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) SetFilter(yf yfilter.YFilter) { bfdTracks.YFilter = yf }

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "rate" { return "Rate" }
    if yname == "debounce-count" { return "DebounceCount" }
    return ""
}

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetSegmentPath() string {
    return "bfd-tracks"
}

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bfdTracks.InterfaceName
    leafs["destination-address"] = bfdTracks.DestinationAddress
    leafs["rate"] = bfdTracks.Rate
    leafs["debounce-count"] = bfdTracks.DebounceCount
    return leafs
}

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetYangName() string { return "bfd-tracks" }

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) SetParent(parent types.Entity) { bfdTracks.parent = parent }

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetParent() types.Entity { return bfdTracks.parent }

func (bfdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_TrackTypeInfo_BfdTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks
// boolean objects
type ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bool track info. The type is slice of
    // ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo.
    BoolTrackInfo []ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo
}

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetFilter() yfilter.YFilter { return boolTracks.YFilter }

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) SetFilter(yf yfilter.YFilter) { boolTracks.YFilter = yf }

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetGoName(yname string) string {
    if yname == "bool-track-info" { return "BoolTrackInfo" }
    return ""
}

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetSegmentPath() string {
    return "bool-tracks"
}

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bool-track-info" {
        for _, c := range boolTracks.BoolTrackInfo {
            if boolTracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo{}
        boolTracks.BoolTrackInfo = append(boolTracks.BoolTrackInfo, child)
        return &boolTracks.BoolTrackInfo[len(boolTracks.BoolTrackInfo)-1]
    }
    return nil
}

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range boolTracks.BoolTrackInfo {
        children[boolTracks.BoolTrackInfo[i].GetSegmentPath()] = &boolTracks.BoolTrackInfo[i]
    }
    return children
}

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetBundleName() string { return "cisco_ios_xr" }

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetYangName() string { return "bool-tracks" }

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) SetParent(parent types.Entity) { boolTracks.parent = parent }

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetParent() types.Entity { return boolTracks.parent }

func (boolTracks *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo
// bool track info
type ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object Name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track object with Not. The type is bool.
    WithNot interface{}
}

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetFilter() yfilter.YFilter { return boolTrackInfo.YFilter }

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) SetFilter(yf yfilter.YFilter) { boolTrackInfo.YFilter = yf }

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "track-state" { return "TrackState" }
    if yname == "with-not" { return "WithNot" }
    return ""
}

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetSegmentPath() string {
    return "bool-track-info"
}

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = boolTrackInfo.ObjectName
    leafs["track-state"] = boolTrackInfo.TrackState
    leafs["with-not"] = boolTrackInfo.WithNot
    return leafs
}

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetYangName() string { return "bool-track-info" }

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) SetParent(parent types.Entity) { boolTrackInfo.parent = parent }

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetParent() types.Entity { return boolTrackInfo.parent }

func (boolTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_BoolTracks_BoolTrackInfo) GetParentYangName() string { return "bool-tracks" }

// ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks
// Threshold objects
type ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // threshold track info. The type is slice of
    // ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo.
    ThresholdTrackInfo []ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo
}

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetFilter() yfilter.YFilter { return thresholdTracks.YFilter }

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) SetFilter(yf yfilter.YFilter) { thresholdTracks.YFilter = yf }

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetGoName(yname string) string {
    if yname == "threshold-track-info" { return "ThresholdTrackInfo" }
    return ""
}

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetSegmentPath() string {
    return "threshold-tracks"
}

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-track-info" {
        for _, c := range thresholdTracks.ThresholdTrackInfo {
            if thresholdTracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo{}
        thresholdTracks.ThresholdTrackInfo = append(thresholdTracks.ThresholdTrackInfo, child)
        return &thresholdTracks.ThresholdTrackInfo[len(thresholdTracks.ThresholdTrackInfo)-1]
    }
    return nil
}

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholdTracks.ThresholdTrackInfo {
        children[thresholdTracks.ThresholdTrackInfo[i].GetSegmentPath()] = &thresholdTracks.ThresholdTrackInfo[i]
    }
    return children
}

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetYangName() string { return "threshold-tracks" }

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) SetParent(parent types.Entity) { thresholdTracks.parent = parent }

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetParent() types.Entity { return thresholdTracks.parent }

func (thresholdTracks *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo
// threshold track info
type ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. True means track is up; False means track is down. The type is
    // bool.
    TrackState interface{}

    // Weight is the number assigned to a track object . In case of a type
    // threshold weight( i.e. weighted sum list), weight is asigned by User at the
    // time of configuration. In case of a type threshold percentage (i.e.
    // percentage based list), weight is internally computed by (1/N)x100, where N
    // is the number of objects in the list. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Weight interface{}
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetFilter() yfilter.YFilter { return thresholdTrackInfo.YFilter }

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) SetFilter(yf yfilter.YFilter) { thresholdTrackInfo.YFilter = yf }

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "track-state" { return "TrackState" }
    if yname == "weight" { return "Weight" }
    return ""
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetSegmentPath() string {
    return "threshold-track-info"
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = thresholdTrackInfo.ObjectName
    leafs["track-state"] = thresholdTrackInfo.TrackState
    leafs["weight"] = thresholdTrackInfo.Weight
    return leafs
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetYangName() string { return "threshold-track-info" }

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) SetParent(parent types.Entity) { thresholdTrackInfo.parent = parent }

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetParent() types.Entity { return thresholdTrackInfo.parent }

func (thresholdTrackInfo *ObjectTracking_TrackTypeInterface_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetParentYangName() string { return "threshold-tracks" }

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces
// Tracking Interfaces
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface tracking info. The type is slice of
    // ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo.
    InterfaceTrackingInfo []ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
}

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetFilter() yfilter.YFilter { return trackingInteraces.YFilter }

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) SetFilter(yf yfilter.YFilter) { trackingInteraces.YFilter = yf }

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetGoName(yname string) string {
    if yname == "interface-tracking-info" { return "InterfaceTrackingInfo" }
    return ""
}

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetSegmentPath() string {
    return "tracking-interaces"
}

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracking-info" {
        for _, c := range trackingInteraces.InterfaceTrackingInfo {
            if trackingInteraces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo{}
        trackingInteraces.InterfaceTrackingInfo = append(trackingInteraces.InterfaceTrackingInfo, child)
        return &trackingInteraces.InterfaceTrackingInfo[len(trackingInteraces.InterfaceTrackingInfo)-1]
    }
    return nil
}

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackingInteraces.InterfaceTrackingInfo {
        children[trackingInteraces.InterfaceTrackingInfo[i].GetSegmentPath()] = &trackingInteraces.InterfaceTrackingInfo[i]
    }
    return children
}

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetBundleName() string { return "cisco_ios_xr" }

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetYangName() string { return "tracking-interaces" }

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) SetParent(parent types.Entity) { trackingInteraces.parent = parent }

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetParent() types.Entity { return trackingInteraces.parent }

func (trackingInteraces *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
// interface tracking info
type ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetFilter() yfilter.YFilter { return interfaceTrackingInfo.YFilter }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) SetFilter(yf yfilter.YFilter) { interfaceTrackingInfo.YFilter = yf }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetSegmentPath() string {
    return "interface-tracking-info"
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTrackingInfo.InterfaceName
    return leafs
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetYangName() string { return "interface-tracking-info" }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) SetParent(parent types.Entity) { interfaceTrackingInfo.parent = parent }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetParent() types.Entity { return interfaceTrackingInfo.parent }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeInterface_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetParentYangName() string { return "tracking-interaces" }

// ObjectTracking_TrackTypeInterface_TrackInfo_Delayed
// Is the state change delay counter in progress
type ObjectTracking_TrackTypeInterface_TrackInfo_Delayed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The time remaining in seconds for the counter to trigger state change. The
    // type is interface{} with range: 0..4294967295. Units are second.
    TimeRemaining interface{}

    // State the track will transition to. Track state. True means track is up;
    // False means track is down. The type is bool.
    TrackState interface{}
}

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetFilter() yfilter.YFilter { return delayed.YFilter }

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) SetFilter(yf yfilter.YFilter) { delayed.YFilter = yf }

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetGoName(yname string) string {
    if yname == "time-remaining" { return "TimeRemaining" }
    if yname == "track-state" { return "TrackState" }
    return ""
}

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetSegmentPath() string {
    return "delayed"
}

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-remaining"] = delayed.TimeRemaining
    leafs["track-state"] = delayed.TrackState
    return leafs
}

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetBundleName() string { return "cisco_ios_xr" }

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetYangName() string { return "delayed" }

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) SetParent(parent types.Entity) { delayed.parent = parent }

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetParent() types.Entity { return delayed.parent }

func (delayed *ObjectTracking_TrackTypeInterface_TrackInfo_Delayed) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackBriefs
// Object Tracking Track table brief
type ObjectTracking_TrackBriefs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTracking_TrackBriefs_TrackBrief.
    TrackBrief []ObjectTracking_TrackBriefs_TrackBrief
}

func (trackBriefs *ObjectTracking_TrackBriefs) GetFilter() yfilter.YFilter { return trackBriefs.YFilter }

func (trackBriefs *ObjectTracking_TrackBriefs) SetFilter(yf yfilter.YFilter) { trackBriefs.YFilter = yf }

func (trackBriefs *ObjectTracking_TrackBriefs) GetGoName(yname string) string {
    if yname == "track-brief" { return "TrackBrief" }
    return ""
}

func (trackBriefs *ObjectTracking_TrackBriefs) GetSegmentPath() string {
    return "track-briefs"
}

func (trackBriefs *ObjectTracking_TrackBriefs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-brief" {
        for _, c := range trackBriefs.TrackBrief {
            if trackBriefs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackBriefs_TrackBrief{}
        trackBriefs.TrackBrief = append(trackBriefs.TrackBrief, child)
        return &trackBriefs.TrackBrief[len(trackBriefs.TrackBrief)-1]
    }
    return nil
}

func (trackBriefs *ObjectTracking_TrackBriefs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackBriefs.TrackBrief {
        children[trackBriefs.TrackBrief[i].GetSegmentPath()] = &trackBriefs.TrackBrief[i]
    }
    return children
}

func (trackBriefs *ObjectTracking_TrackBriefs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackBriefs *ObjectTracking_TrackBriefs) GetBundleName() string { return "cisco_ios_xr" }

func (trackBriefs *ObjectTracking_TrackBriefs) GetYangName() string { return "track-briefs" }

func (trackBriefs *ObjectTracking_TrackBriefs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackBriefs *ObjectTracking_TrackBriefs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackBriefs *ObjectTracking_TrackBriefs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackBriefs *ObjectTracking_TrackBriefs) SetParent(parent types.Entity) { trackBriefs.parent = parent }

func (trackBriefs *ObjectTracking_TrackBriefs) GetParent() types.Entity { return trackBriefs.parent }

func (trackBriefs *ObjectTracking_TrackBriefs) GetParentYangName() string { return "object-tracking" }

// ObjectTracking_TrackBriefs_TrackBrief
// Track name - maximum 32 characters
type ObjectTracking_TrackBriefs_TrackBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Track name. The type is string with length: 1..32.
    TrackName interface{}

    // track info brief. The type is slice of
    // ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief.
    TrackInfoBrief []ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief
}

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetFilter() yfilter.YFilter { return trackBrief.YFilter }

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) SetFilter(yf yfilter.YFilter) { trackBrief.YFilter = yf }

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetGoName(yname string) string {
    if yname == "track-name" { return "TrackName" }
    if yname == "track-info-brief" { return "TrackInfoBrief" }
    return ""
}

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetSegmentPath() string {
    return "track-brief" + "[track-name='" + fmt.Sprintf("%v", trackBrief.TrackName) + "']"
}

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-info-brief" {
        for _, c := range trackBrief.TrackInfoBrief {
            if trackBrief.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief{}
        trackBrief.TrackInfoBrief = append(trackBrief.TrackInfoBrief, child)
        return &trackBrief.TrackInfoBrief[len(trackBrief.TrackInfoBrief)-1]
    }
    return nil
}

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackBrief.TrackInfoBrief {
        children[trackBrief.TrackInfoBrief[i].GetSegmentPath()] = &trackBrief.TrackInfoBrief[i]
    }
    return children
}

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-name"] = trackBrief.TrackName
    return leafs
}

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetBundleName() string { return "cisco_ios_xr" }

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetYangName() string { return "track-brief" }

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) SetParent(parent types.Entity) { trackBrief.parent = parent }

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetParent() types.Entity { return trackBrief.parent }

func (trackBrief *ObjectTracking_TrackBriefs_TrackBrief) GetParentYangName() string { return "track-briefs" }

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief
// track info brief
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo
}

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetFilter() yfilter.YFilter { return trackInfoBrief.YFilter }

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) SetFilter(yf yfilter.YFilter) { trackInfoBrief.YFilter = yf }

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetGoName(yname string) string {
    if yname == "tracke-name" { return "TrackeName" }
    if yname == "type" { return "Type" }
    if yname == "track-state" { return "TrackState" }
    if yname == "track-type-info" { return "TrackTypeInfo" }
    return ""
}

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetSegmentPath() string {
    return "track-info-brief"
}

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-type-info" {
        return &trackInfoBrief.TrackTypeInfo
    }
    return nil
}

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-type-info"] = &trackInfoBrief.TrackTypeInfo
    return children
}

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tracke-name"] = trackInfoBrief.TrackeName
    leafs["type"] = trackInfoBrief.Type
    leafs["track-state"] = trackInfoBrief.TrackState
    return leafs
}

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetBundleName() string { return "cisco_ios_xr" }

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetYangName() string { return "track-info-brief" }

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) SetParent(parent types.Entity) { trackInfoBrief.parent = parent }

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetParent() types.Entity { return trackInfoBrief.parent }

func (trackInfoBrief *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief) GetParentYangName() string { return "track-brief" }

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo
// Track type information
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetFilter() yfilter.YFilter { return trackTypeInfo.YFilter }

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) SetFilter(yf yfilter.YFilter) { trackTypeInfo.YFilter = yf }

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetGoName(yname string) string {
    if yname == "discriminant" { return "Discriminant" }
    if yname == "interface-tracks" { return "InterfaceTracks" }
    if yname == "route-tracks" { return "RouteTracks" }
    if yname == "ipsla-tracks" { return "IpslaTracks" }
    if yname == "bfd-tracks" { return "BfdTracks" }
    return ""
}

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetSegmentPath() string {
    return "track-type-info"
}

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracks" {
        return &trackTypeInfo.InterfaceTracks
    }
    if childYangName == "route-tracks" {
        return &trackTypeInfo.RouteTracks
    }
    if childYangName == "ipsla-tracks" {
        return &trackTypeInfo.IpslaTracks
    }
    if childYangName == "bfd-tracks" {
        return &trackTypeInfo.BfdTracks
    }
    return nil
}

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-tracks"] = &trackTypeInfo.InterfaceTracks
    children["route-tracks"] = &trackTypeInfo.RouteTracks
    children["ipsla-tracks"] = &trackTypeInfo.IpslaTracks
    children["bfd-tracks"] = &trackTypeInfo.BfdTracks
    return children
}

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminant"] = trackTypeInfo.Discriminant
    return leafs
}

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetYangName() string { return "track-type-info" }

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) SetParent(parent types.Entity) { trackTypeInfo.parent = parent }

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetParent() types.Entity { return trackTypeInfo.parent }

func (trackTypeInfo *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo) GetParentYangName() string { return "track-info-brief" }

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetFilter() yfilter.YFilter { return interfaceTracks.YFilter }

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) SetFilter(yf yfilter.YFilter) { interfaceTracks.YFilter = yf }

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetSegmentPath() string {
    return "interface-tracks"
}

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTracks.InterfaceName
    return leafs
}

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetYangName() string { return "interface-tracks" }

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) SetParent(parent types.Entity) { interfaceTracks.parent = parent }

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetParent() types.Entity { return interfaceTracks.parent }

func (interfaceTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetFilter() yfilter.YFilter { return routeTracks.YFilter }

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) SetFilter(yf yfilter.YFilter) { routeTracks.YFilter = yf }

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "vrf" { return "Vrf" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetSegmentPath() string {
    return "route-tracks"
}

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = routeTracks.Prefix
    leafs["prefix-length"] = routeTracks.PrefixLength
    leafs["vrf"] = routeTracks.Vrf
    leafs["next-hop"] = routeTracks.NextHop
    return leafs
}

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetBundleName() string { return "cisco_ios_xr" }

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetYangName() string { return "route-tracks" }

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) SetParent(parent types.Entity) { routeTracks.parent = parent }

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetParent() types.Entity { return routeTracks.parent }

func (routeTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetFilter() yfilter.YFilter { return ipslaTracks.YFilter }

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) SetFilter(yf yfilter.YFilter) { ipslaTracks.YFilter = yf }

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetGoName(yname string) string {
    if yname == "ipsla-op-id" { return "IpslaOpId" }
    if yname == "rtt" { return "Rtt" }
    if yname == "return-code" { return "ReturnCode" }
    return ""
}

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetSegmentPath() string {
    return "ipsla-tracks"
}

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipsla-op-id"] = ipslaTracks.IpslaOpId
    leafs["rtt"] = ipslaTracks.Rtt
    leafs["return-code"] = ipslaTracks.ReturnCode
    return leafs
}

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetBundleName() string { return "cisco_ios_xr" }

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetYangName() string { return "ipsla-tracks" }

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) SetParent(parent types.Entity) { ipslaTracks.parent = parent }

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetParent() types.Entity { return ipslaTracks.parent }

func (ipslaTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetFilter() yfilter.YFilter { return bfdTracks.YFilter }

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) SetFilter(yf yfilter.YFilter) { bfdTracks.YFilter = yf }

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "rate" { return "Rate" }
    if yname == "debounce-count" { return "DebounceCount" }
    return ""
}

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetSegmentPath() string {
    return "bfd-tracks"
}

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bfdTracks.InterfaceName
    leafs["destination-address"] = bfdTracks.DestinationAddress
    leafs["rate"] = bfdTracks.Rate
    leafs["debounce-count"] = bfdTracks.DebounceCount
    return leafs
}

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetYangName() string { return "bfd-tracks" }

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) SetParent(parent types.Entity) { bfdTracks.parent = parent }

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetParent() types.Entity { return bfdTracks.parent }

func (bfdTracks *ObjectTracking_TrackBriefs_TrackBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeRtrReachability
// Object Tracking Type RTR Reachability info
type ObjectTracking_TrackTypeRtrReachability struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // track info. The type is slice of
    // ObjectTracking_TrackTypeRtrReachability_TrackInfo.
    TrackInfo []ObjectTracking_TrackTypeRtrReachability_TrackInfo
}

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetFilter() yfilter.YFilter { return trackTypeRtrReachability.YFilter }

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) SetFilter(yf yfilter.YFilter) { trackTypeRtrReachability.YFilter = yf }

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetGoName(yname string) string {
    if yname == "track-info" { return "TrackInfo" }
    return ""
}

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetSegmentPath() string {
    return "track-type-rtr-reachability"
}

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-info" {
        for _, c := range trackTypeRtrReachability.TrackInfo {
            if trackTypeRtrReachability.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeRtrReachability_TrackInfo{}
        trackTypeRtrReachability.TrackInfo = append(trackTypeRtrReachability.TrackInfo, child)
        return &trackTypeRtrReachability.TrackInfo[len(trackTypeRtrReachability.TrackInfo)-1]
    }
    return nil
}

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackTypeRtrReachability.TrackInfo {
        children[trackTypeRtrReachability.TrackInfo[i].GetSegmentPath()] = &trackTypeRtrReachability.TrackInfo[i]
    }
    return children
}

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetYangName() string { return "track-type-rtr-reachability" }

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) SetParent(parent types.Entity) { trackTypeRtrReachability.parent = parent }

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetParent() types.Entity { return trackTypeRtrReachability.parent }

func (trackTypeRtrReachability *ObjectTracking_TrackTypeRtrReachability) GetParentYangName() string { return "object-tracking" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo
// track info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // State Change Counter. The type is interface{} with range: 0..4294967295.
    StateChangeCounter interface{}

    // Seconds Last Change. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecondsLastChange interface{}

    // User specified threshold upper limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdUp interface{}

    // User specified threshold lower limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdDown interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo

    // boolean objects.
    BoolTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks

    // Threshold objects.
    ThresholdTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks

    // Tracking Interfaces.
    TrackingInteraces ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces

    // Is the state change delay counter in progress.
    Delayed ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed
}

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetFilter() yfilter.YFilter { return trackInfo.YFilter }

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) SetFilter(yf yfilter.YFilter) { trackInfo.YFilter = yf }

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetGoName(yname string) string {
    if yname == "tracke-name" { return "TrackeName" }
    if yname == "type" { return "Type" }
    if yname == "track-state" { return "TrackState" }
    if yname == "state-change-counter" { return "StateChangeCounter" }
    if yname == "seconds-last-change" { return "SecondsLastChange" }
    if yname == "threshold-up" { return "ThresholdUp" }
    if yname == "threshold-down" { return "ThresholdDown" }
    if yname == "track-type-info" { return "TrackTypeInfo" }
    if yname == "bool-tracks" { return "BoolTracks" }
    if yname == "threshold-tracks" { return "ThresholdTracks" }
    if yname == "tracking-interaces" { return "TrackingInteraces" }
    if yname == "delayed" { return "Delayed" }
    return ""
}

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetSegmentPath() string {
    return "track-info"
}

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-type-info" {
        return &trackInfo.TrackTypeInfo
    }
    if childYangName == "bool-tracks" {
        return &trackInfo.BoolTracks
    }
    if childYangName == "threshold-tracks" {
        return &trackInfo.ThresholdTracks
    }
    if childYangName == "tracking-interaces" {
        return &trackInfo.TrackingInteraces
    }
    if childYangName == "delayed" {
        return &trackInfo.Delayed
    }
    return nil
}

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-type-info"] = &trackInfo.TrackTypeInfo
    children["bool-tracks"] = &trackInfo.BoolTracks
    children["threshold-tracks"] = &trackInfo.ThresholdTracks
    children["tracking-interaces"] = &trackInfo.TrackingInteraces
    children["delayed"] = &trackInfo.Delayed
    return children
}

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tracke-name"] = trackInfo.TrackeName
    leafs["type"] = trackInfo.Type
    leafs["track-state"] = trackInfo.TrackState
    leafs["state-change-counter"] = trackInfo.StateChangeCounter
    leafs["seconds-last-change"] = trackInfo.SecondsLastChange
    leafs["threshold-up"] = trackInfo.ThresholdUp
    leafs["threshold-down"] = trackInfo.ThresholdDown
    return leafs
}

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetYangName() string { return "track-info" }

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) SetParent(parent types.Entity) { trackInfo.parent = parent }

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetParent() types.Entity { return trackInfo.parent }

func (trackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo) GetParentYangName() string { return "track-type-rtr-reachability" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetFilter() yfilter.YFilter { return trackTypeInfo.YFilter }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) SetFilter(yf yfilter.YFilter) { trackTypeInfo.YFilter = yf }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetGoName(yname string) string {
    if yname == "discriminant" { return "Discriminant" }
    if yname == "interface-tracks" { return "InterfaceTracks" }
    if yname == "route-tracks" { return "RouteTracks" }
    if yname == "ipsla-tracks" { return "IpslaTracks" }
    if yname == "bfd-tracks" { return "BfdTracks" }
    return ""
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetSegmentPath() string {
    return "track-type-info"
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracks" {
        return &trackTypeInfo.InterfaceTracks
    }
    if childYangName == "route-tracks" {
        return &trackTypeInfo.RouteTracks
    }
    if childYangName == "ipsla-tracks" {
        return &trackTypeInfo.IpslaTracks
    }
    if childYangName == "bfd-tracks" {
        return &trackTypeInfo.BfdTracks
    }
    return nil
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-tracks"] = &trackTypeInfo.InterfaceTracks
    children["route-tracks"] = &trackTypeInfo.RouteTracks
    children["ipsla-tracks"] = &trackTypeInfo.IpslaTracks
    children["bfd-tracks"] = &trackTypeInfo.BfdTracks
    return children
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminant"] = trackTypeInfo.Discriminant
    return leafs
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetYangName() string { return "track-type-info" }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) SetParent(parent types.Entity) { trackTypeInfo.parent = parent }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetParent() types.Entity { return trackTypeInfo.parent }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetFilter() yfilter.YFilter { return interfaceTracks.YFilter }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) SetFilter(yf yfilter.YFilter) { interfaceTracks.YFilter = yf }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetSegmentPath() string {
    return "interface-tracks"
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTracks.InterfaceName
    return leafs
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetYangName() string { return "interface-tracks" }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) SetParent(parent types.Entity) { interfaceTracks.parent = parent }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetParent() types.Entity { return interfaceTracks.parent }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_InterfaceTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetFilter() yfilter.YFilter { return routeTracks.YFilter }

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) SetFilter(yf yfilter.YFilter) { routeTracks.YFilter = yf }

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "vrf" { return "Vrf" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetSegmentPath() string {
    return "route-tracks"
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = routeTracks.Prefix
    leafs["prefix-length"] = routeTracks.PrefixLength
    leafs["vrf"] = routeTracks.Vrf
    leafs["next-hop"] = routeTracks.NextHop
    return leafs
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetBundleName() string { return "cisco_ios_xr" }

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetYangName() string { return "route-tracks" }

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) SetParent(parent types.Entity) { routeTracks.parent = parent }

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetParent() types.Entity { return routeTracks.parent }

func (routeTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_RouteTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetFilter() yfilter.YFilter { return ipslaTracks.YFilter }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) SetFilter(yf yfilter.YFilter) { ipslaTracks.YFilter = yf }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetGoName(yname string) string {
    if yname == "ipsla-op-id" { return "IpslaOpId" }
    if yname == "rtt" { return "Rtt" }
    if yname == "return-code" { return "ReturnCode" }
    return ""
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetSegmentPath() string {
    return "ipsla-tracks"
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipsla-op-id"] = ipslaTracks.IpslaOpId
    leafs["rtt"] = ipslaTracks.Rtt
    leafs["return-code"] = ipslaTracks.ReturnCode
    return leafs
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetBundleName() string { return "cisco_ios_xr" }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetYangName() string { return "ipsla-tracks" }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) SetParent(parent types.Entity) { ipslaTracks.parent = parent }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetParent() types.Entity { return ipslaTracks.parent }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_IpslaTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetFilter() yfilter.YFilter { return bfdTracks.YFilter }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) SetFilter(yf yfilter.YFilter) { bfdTracks.YFilter = yf }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "rate" { return "Rate" }
    if yname == "debounce-count" { return "DebounceCount" }
    return ""
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetSegmentPath() string {
    return "bfd-tracks"
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bfdTracks.InterfaceName
    leafs["destination-address"] = bfdTracks.DestinationAddress
    leafs["rate"] = bfdTracks.Rate
    leafs["debounce-count"] = bfdTracks.DebounceCount
    return leafs
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetYangName() string { return "bfd-tracks" }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) SetParent(parent types.Entity) { bfdTracks.parent = parent }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetParent() types.Entity { return bfdTracks.parent }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackTypeInfo_BfdTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks
// boolean objects
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bool track info. The type is slice of
    // ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo.
    BoolTrackInfo []ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo
}

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetFilter() yfilter.YFilter { return boolTracks.YFilter }

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) SetFilter(yf yfilter.YFilter) { boolTracks.YFilter = yf }

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetGoName(yname string) string {
    if yname == "bool-track-info" { return "BoolTrackInfo" }
    return ""
}

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetSegmentPath() string {
    return "bool-tracks"
}

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bool-track-info" {
        for _, c := range boolTracks.BoolTrackInfo {
            if boolTracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo{}
        boolTracks.BoolTrackInfo = append(boolTracks.BoolTrackInfo, child)
        return &boolTracks.BoolTrackInfo[len(boolTracks.BoolTrackInfo)-1]
    }
    return nil
}

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range boolTracks.BoolTrackInfo {
        children[boolTracks.BoolTrackInfo[i].GetSegmentPath()] = &boolTracks.BoolTrackInfo[i]
    }
    return children
}

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetBundleName() string { return "cisco_ios_xr" }

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetYangName() string { return "bool-tracks" }

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) SetParent(parent types.Entity) { boolTracks.parent = parent }

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetParent() types.Entity { return boolTracks.parent }

func (boolTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo
// bool track info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object Name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track object with Not. The type is bool.
    WithNot interface{}
}

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetFilter() yfilter.YFilter { return boolTrackInfo.YFilter }

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) SetFilter(yf yfilter.YFilter) { boolTrackInfo.YFilter = yf }

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "track-state" { return "TrackState" }
    if yname == "with-not" { return "WithNot" }
    return ""
}

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetSegmentPath() string {
    return "bool-track-info"
}

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = boolTrackInfo.ObjectName
    leafs["track-state"] = boolTrackInfo.TrackState
    leafs["with-not"] = boolTrackInfo.WithNot
    return leafs
}

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetYangName() string { return "bool-track-info" }

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) SetParent(parent types.Entity) { boolTrackInfo.parent = parent }

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetParent() types.Entity { return boolTrackInfo.parent }

func (boolTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_BoolTracks_BoolTrackInfo) GetParentYangName() string { return "bool-tracks" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks
// Threshold objects
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // threshold track info. The type is slice of
    // ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo.
    ThresholdTrackInfo []ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo
}

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetFilter() yfilter.YFilter { return thresholdTracks.YFilter }

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) SetFilter(yf yfilter.YFilter) { thresholdTracks.YFilter = yf }

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetGoName(yname string) string {
    if yname == "threshold-track-info" { return "ThresholdTrackInfo" }
    return ""
}

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetSegmentPath() string {
    return "threshold-tracks"
}

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-track-info" {
        for _, c := range thresholdTracks.ThresholdTrackInfo {
            if thresholdTracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo{}
        thresholdTracks.ThresholdTrackInfo = append(thresholdTracks.ThresholdTrackInfo, child)
        return &thresholdTracks.ThresholdTrackInfo[len(thresholdTracks.ThresholdTrackInfo)-1]
    }
    return nil
}

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholdTracks.ThresholdTrackInfo {
        children[thresholdTracks.ThresholdTrackInfo[i].GetSegmentPath()] = &thresholdTracks.ThresholdTrackInfo[i]
    }
    return children
}

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetYangName() string { return "threshold-tracks" }

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) SetParent(parent types.Entity) { thresholdTracks.parent = parent }

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetParent() types.Entity { return thresholdTracks.parent }

func (thresholdTracks *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo
// threshold track info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. True means track is up; False means track is down. The type is
    // bool.
    TrackState interface{}

    // Weight is the number assigned to a track object . In case of a type
    // threshold weight( i.e. weighted sum list), weight is asigned by User at the
    // time of configuration. In case of a type threshold percentage (i.e.
    // percentage based list), weight is internally computed by (1/N)x100, where N
    // is the number of objects in the list. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Weight interface{}
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetFilter() yfilter.YFilter { return thresholdTrackInfo.YFilter }

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) SetFilter(yf yfilter.YFilter) { thresholdTrackInfo.YFilter = yf }

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "track-state" { return "TrackState" }
    if yname == "weight" { return "Weight" }
    return ""
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetSegmentPath() string {
    return "threshold-track-info"
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = thresholdTrackInfo.ObjectName
    leafs["track-state"] = thresholdTrackInfo.TrackState
    leafs["weight"] = thresholdTrackInfo.Weight
    return leafs
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetYangName() string { return "threshold-track-info" }

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) SetParent(parent types.Entity) { thresholdTrackInfo.parent = parent }

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetParent() types.Entity { return thresholdTrackInfo.parent }

func (thresholdTrackInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetParentYangName() string { return "threshold-tracks" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces
// Tracking Interfaces
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface tracking info. The type is slice of
    // ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo.
    InterfaceTrackingInfo []ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
}

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetFilter() yfilter.YFilter { return trackingInteraces.YFilter }

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) SetFilter(yf yfilter.YFilter) { trackingInteraces.YFilter = yf }

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetGoName(yname string) string {
    if yname == "interface-tracking-info" { return "InterfaceTrackingInfo" }
    return ""
}

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetSegmentPath() string {
    return "tracking-interaces"
}

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracking-info" {
        for _, c := range trackingInteraces.InterfaceTrackingInfo {
            if trackingInteraces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo{}
        trackingInteraces.InterfaceTrackingInfo = append(trackingInteraces.InterfaceTrackingInfo, child)
        return &trackingInteraces.InterfaceTrackingInfo[len(trackingInteraces.InterfaceTrackingInfo)-1]
    }
    return nil
}

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackingInteraces.InterfaceTrackingInfo {
        children[trackingInteraces.InterfaceTrackingInfo[i].GetSegmentPath()] = &trackingInteraces.InterfaceTrackingInfo[i]
    }
    return children
}

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetBundleName() string { return "cisco_ios_xr" }

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetYangName() string { return "tracking-interaces" }

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) SetParent(parent types.Entity) { trackingInteraces.parent = parent }

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetParent() types.Entity { return trackingInteraces.parent }

func (trackingInteraces *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
// interface tracking info
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetFilter() yfilter.YFilter { return interfaceTrackingInfo.YFilter }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) SetFilter(yf yfilter.YFilter) { interfaceTrackingInfo.YFilter = yf }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetSegmentPath() string {
    return "interface-tracking-info"
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTrackingInfo.InterfaceName
    return leafs
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetYangName() string { return "interface-tracking-info" }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) SetParent(parent types.Entity) { interfaceTrackingInfo.parent = parent }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetParent() types.Entity { return interfaceTrackingInfo.parent }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeRtrReachability_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetParentYangName() string { return "tracking-interaces" }

// ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed
// Is the state change delay counter in progress
type ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The time remaining in seconds for the counter to trigger state change. The
    // type is interface{} with range: 0..4294967295. Units are second.
    TimeRemaining interface{}

    // State the track will transition to. Track state. True means track is up;
    // False means track is down. The type is bool.
    TrackState interface{}
}

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetFilter() yfilter.YFilter { return delayed.YFilter }

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) SetFilter(yf yfilter.YFilter) { delayed.YFilter = yf }

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetGoName(yname string) string {
    if yname == "time-remaining" { return "TimeRemaining" }
    if yname == "track-state" { return "TrackState" }
    return ""
}

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetSegmentPath() string {
    return "delayed"
}

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-remaining"] = delayed.TimeRemaining
    leafs["track-state"] = delayed.TrackState
    return leafs
}

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetBundleName() string { return "cisco_ios_xr" }

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetYangName() string { return "delayed" }

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) SetParent(parent types.Entity) { delayed.parent = parent }

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetParent() types.Entity { return delayed.parent }

func (delayed *ObjectTracking_TrackTypeRtrReachability_TrackInfo_Delayed) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeRtrReachabilityBrief
// Object Tracking Type RTR Reachability brief info
type ObjectTracking_TrackTypeRtrReachabilityBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // track info brief. The type is slice of
    // ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief.
    TrackInfoBrief []ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief
}

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetFilter() yfilter.YFilter { return trackTypeRtrReachabilityBrief.YFilter }

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) SetFilter(yf yfilter.YFilter) { trackTypeRtrReachabilityBrief.YFilter = yf }

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetGoName(yname string) string {
    if yname == "track-info-brief" { return "TrackInfoBrief" }
    return ""
}

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetSegmentPath() string {
    return "track-type-rtr-reachability-brief"
}

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-info-brief" {
        for _, c := range trackTypeRtrReachabilityBrief.TrackInfoBrief {
            if trackTypeRtrReachabilityBrief.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief{}
        trackTypeRtrReachabilityBrief.TrackInfoBrief = append(trackTypeRtrReachabilityBrief.TrackInfoBrief, child)
        return &trackTypeRtrReachabilityBrief.TrackInfoBrief[len(trackTypeRtrReachabilityBrief.TrackInfoBrief)-1]
    }
    return nil
}

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackTypeRtrReachabilityBrief.TrackInfoBrief {
        children[trackTypeRtrReachabilityBrief.TrackInfoBrief[i].GetSegmentPath()] = &trackTypeRtrReachabilityBrief.TrackInfoBrief[i]
    }
    return children
}

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetYangName() string { return "track-type-rtr-reachability-brief" }

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) SetParent(parent types.Entity) { trackTypeRtrReachabilityBrief.parent = parent }

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetParent() types.Entity { return trackTypeRtrReachabilityBrief.parent }

func (trackTypeRtrReachabilityBrief *ObjectTracking_TrackTypeRtrReachabilityBrief) GetParentYangName() string { return "object-tracking" }

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief
// track info brief
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo
}

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetFilter() yfilter.YFilter { return trackInfoBrief.YFilter }

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) SetFilter(yf yfilter.YFilter) { trackInfoBrief.YFilter = yf }

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetGoName(yname string) string {
    if yname == "tracke-name" { return "TrackeName" }
    if yname == "type" { return "Type" }
    if yname == "track-state" { return "TrackState" }
    if yname == "track-type-info" { return "TrackTypeInfo" }
    return ""
}

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetSegmentPath() string {
    return "track-info-brief"
}

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-type-info" {
        return &trackInfoBrief.TrackTypeInfo
    }
    return nil
}

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-type-info"] = &trackInfoBrief.TrackTypeInfo
    return children
}

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tracke-name"] = trackInfoBrief.TrackeName
    leafs["type"] = trackInfoBrief.Type
    leafs["track-state"] = trackInfoBrief.TrackState
    return leafs
}

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetBundleName() string { return "cisco_ios_xr" }

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetYangName() string { return "track-info-brief" }

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) SetParent(parent types.Entity) { trackInfoBrief.parent = parent }

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetParent() types.Entity { return trackInfoBrief.parent }

func (trackInfoBrief *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief) GetParentYangName() string { return "track-type-rtr-reachability-brief" }

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetFilter() yfilter.YFilter { return trackTypeInfo.YFilter }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) SetFilter(yf yfilter.YFilter) { trackTypeInfo.YFilter = yf }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetGoName(yname string) string {
    if yname == "discriminant" { return "Discriminant" }
    if yname == "interface-tracks" { return "InterfaceTracks" }
    if yname == "route-tracks" { return "RouteTracks" }
    if yname == "ipsla-tracks" { return "IpslaTracks" }
    if yname == "bfd-tracks" { return "BfdTracks" }
    return ""
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetSegmentPath() string {
    return "track-type-info"
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracks" {
        return &trackTypeInfo.InterfaceTracks
    }
    if childYangName == "route-tracks" {
        return &trackTypeInfo.RouteTracks
    }
    if childYangName == "ipsla-tracks" {
        return &trackTypeInfo.IpslaTracks
    }
    if childYangName == "bfd-tracks" {
        return &trackTypeInfo.BfdTracks
    }
    return nil
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-tracks"] = &trackTypeInfo.InterfaceTracks
    children["route-tracks"] = &trackTypeInfo.RouteTracks
    children["ipsla-tracks"] = &trackTypeInfo.IpslaTracks
    children["bfd-tracks"] = &trackTypeInfo.BfdTracks
    return children
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminant"] = trackTypeInfo.Discriminant
    return leafs
}

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetYangName() string { return "track-type-info" }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) SetParent(parent types.Entity) { trackTypeInfo.parent = parent }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetParent() types.Entity { return trackTypeInfo.parent }

func (trackTypeInfo *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo) GetParentYangName() string { return "track-info-brief" }

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetFilter() yfilter.YFilter { return interfaceTracks.YFilter }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) SetFilter(yf yfilter.YFilter) { interfaceTracks.YFilter = yf }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetSegmentPath() string {
    return "interface-tracks"
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTracks.InterfaceName
    return leafs
}

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetYangName() string { return "interface-tracks" }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) SetParent(parent types.Entity) { interfaceTracks.parent = parent }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetParent() types.Entity { return interfaceTracks.parent }

func (interfaceTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetFilter() yfilter.YFilter { return routeTracks.YFilter }

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) SetFilter(yf yfilter.YFilter) { routeTracks.YFilter = yf }

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "vrf" { return "Vrf" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetSegmentPath() string {
    return "route-tracks"
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = routeTracks.Prefix
    leafs["prefix-length"] = routeTracks.PrefixLength
    leafs["vrf"] = routeTracks.Vrf
    leafs["next-hop"] = routeTracks.NextHop
    return leafs
}

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetBundleName() string { return "cisco_ios_xr" }

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetYangName() string { return "route-tracks" }

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) SetParent(parent types.Entity) { routeTracks.parent = parent }

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetParent() types.Entity { return routeTracks.parent }

func (routeTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetFilter() yfilter.YFilter { return ipslaTracks.YFilter }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) SetFilter(yf yfilter.YFilter) { ipslaTracks.YFilter = yf }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetGoName(yname string) string {
    if yname == "ipsla-op-id" { return "IpslaOpId" }
    if yname == "rtt" { return "Rtt" }
    if yname == "return-code" { return "ReturnCode" }
    return ""
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetSegmentPath() string {
    return "ipsla-tracks"
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipsla-op-id"] = ipslaTracks.IpslaOpId
    leafs["rtt"] = ipslaTracks.Rtt
    leafs["return-code"] = ipslaTracks.ReturnCode
    return leafs
}

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetBundleName() string { return "cisco_ios_xr" }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetYangName() string { return "ipsla-tracks" }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) SetParent(parent types.Entity) { ipslaTracks.parent = parent }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetParent() types.Entity { return ipslaTracks.parent }

func (ipslaTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetFilter() yfilter.YFilter { return bfdTracks.YFilter }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) SetFilter(yf yfilter.YFilter) { bfdTracks.YFilter = yf }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "rate" { return "Rate" }
    if yname == "debounce-count" { return "DebounceCount" }
    return ""
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetSegmentPath() string {
    return "bfd-tracks"
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bfdTracks.InterfaceName
    leafs["destination-address"] = bfdTracks.DestinationAddress
    leafs["rate"] = bfdTracks.Rate
    leafs["debounce-count"] = bfdTracks.DebounceCount
    return leafs
}

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetYangName() string { return "bfd-tracks" }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) SetParent(parent types.Entity) { bfdTracks.parent = parent }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetParent() types.Entity { return bfdTracks.parent }

func (bfdTracks *ObjectTracking_TrackTypeRtrReachabilityBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_Tracks
// Object Tracking Track table
type ObjectTracking_Tracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTracking_Tracks_Track.
    Track []ObjectTracking_Tracks_Track
}

func (tracks *ObjectTracking_Tracks) GetFilter() yfilter.YFilter { return tracks.YFilter }

func (tracks *ObjectTracking_Tracks) SetFilter(yf yfilter.YFilter) { tracks.YFilter = yf }

func (tracks *ObjectTracking_Tracks) GetGoName(yname string) string {
    if yname == "track" { return "Track" }
    return ""
}

func (tracks *ObjectTracking_Tracks) GetSegmentPath() string {
    return "tracks"
}

func (tracks *ObjectTracking_Tracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track" {
        for _, c := range tracks.Track {
            if tracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_Tracks_Track{}
        tracks.Track = append(tracks.Track, child)
        return &tracks.Track[len(tracks.Track)-1]
    }
    return nil
}

func (tracks *ObjectTracking_Tracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tracks.Track {
        children[tracks.Track[i].GetSegmentPath()] = &tracks.Track[i]
    }
    return children
}

func (tracks *ObjectTracking_Tracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tracks *ObjectTracking_Tracks) GetBundleName() string { return "cisco_ios_xr" }

func (tracks *ObjectTracking_Tracks) GetYangName() string { return "tracks" }

func (tracks *ObjectTracking_Tracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracks *ObjectTracking_Tracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracks *ObjectTracking_Tracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracks *ObjectTracking_Tracks) SetParent(parent types.Entity) { tracks.parent = parent }

func (tracks *ObjectTracking_Tracks) GetParent() types.Entity { return tracks.parent }

func (tracks *ObjectTracking_Tracks) GetParentYangName() string { return "object-tracking" }

// ObjectTracking_Tracks_Track
// Track name - maximum 32 characters
type ObjectTracking_Tracks_Track struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Track name. The type is string with length: 1..32.
    TrackName interface{}

    // track info. The type is slice of ObjectTracking_Tracks_Track_TrackInfo.
    TrackInfo []ObjectTracking_Tracks_Track_TrackInfo
}

func (track *ObjectTracking_Tracks_Track) GetFilter() yfilter.YFilter { return track.YFilter }

func (track *ObjectTracking_Tracks_Track) SetFilter(yf yfilter.YFilter) { track.YFilter = yf }

func (track *ObjectTracking_Tracks_Track) GetGoName(yname string) string {
    if yname == "track-name" { return "TrackName" }
    if yname == "track-info" { return "TrackInfo" }
    return ""
}

func (track *ObjectTracking_Tracks_Track) GetSegmentPath() string {
    return "track" + "[track-name='" + fmt.Sprintf("%v", track.TrackName) + "']"
}

func (track *ObjectTracking_Tracks_Track) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-info" {
        for _, c := range track.TrackInfo {
            if track.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_Tracks_Track_TrackInfo{}
        track.TrackInfo = append(track.TrackInfo, child)
        return &track.TrackInfo[len(track.TrackInfo)-1]
    }
    return nil
}

func (track *ObjectTracking_Tracks_Track) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range track.TrackInfo {
        children[track.TrackInfo[i].GetSegmentPath()] = &track.TrackInfo[i]
    }
    return children
}

func (track *ObjectTracking_Tracks_Track) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-name"] = track.TrackName
    return leafs
}

func (track *ObjectTracking_Tracks_Track) GetBundleName() string { return "cisco_ios_xr" }

func (track *ObjectTracking_Tracks_Track) GetYangName() string { return "track" }

func (track *ObjectTracking_Tracks_Track) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (track *ObjectTracking_Tracks_Track) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (track *ObjectTracking_Tracks_Track) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (track *ObjectTracking_Tracks_Track) SetParent(parent types.Entity) { track.parent = parent }

func (track *ObjectTracking_Tracks_Track) GetParent() types.Entity { return track.parent }

func (track *ObjectTracking_Tracks_Track) GetParentYangName() string { return "tracks" }

// ObjectTracking_Tracks_Track_TrackInfo
// track info
type ObjectTracking_Tracks_Track_TrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // State Change Counter. The type is interface{} with range: 0..4294967295.
    StateChangeCounter interface{}

    // Seconds Last Change. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecondsLastChange interface{}

    // User specified threshold upper limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdUp interface{}

    // User specified threshold lower limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdDown interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo

    // boolean objects.
    BoolTracks ObjectTracking_Tracks_Track_TrackInfo_BoolTracks

    // Threshold objects.
    ThresholdTracks ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks

    // Tracking Interfaces.
    TrackingInteraces ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces

    // Is the state change delay counter in progress.
    Delayed ObjectTracking_Tracks_Track_TrackInfo_Delayed
}

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetFilter() yfilter.YFilter { return trackInfo.YFilter }

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) SetFilter(yf yfilter.YFilter) { trackInfo.YFilter = yf }

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetGoName(yname string) string {
    if yname == "tracke-name" { return "TrackeName" }
    if yname == "type" { return "Type" }
    if yname == "track-state" { return "TrackState" }
    if yname == "state-change-counter" { return "StateChangeCounter" }
    if yname == "seconds-last-change" { return "SecondsLastChange" }
    if yname == "threshold-up" { return "ThresholdUp" }
    if yname == "threshold-down" { return "ThresholdDown" }
    if yname == "track-type-info" { return "TrackTypeInfo" }
    if yname == "bool-tracks" { return "BoolTracks" }
    if yname == "threshold-tracks" { return "ThresholdTracks" }
    if yname == "tracking-interaces" { return "TrackingInteraces" }
    if yname == "delayed" { return "Delayed" }
    return ""
}

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetSegmentPath() string {
    return "track-info"
}

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-type-info" {
        return &trackInfo.TrackTypeInfo
    }
    if childYangName == "bool-tracks" {
        return &trackInfo.BoolTracks
    }
    if childYangName == "threshold-tracks" {
        return &trackInfo.ThresholdTracks
    }
    if childYangName == "tracking-interaces" {
        return &trackInfo.TrackingInteraces
    }
    if childYangName == "delayed" {
        return &trackInfo.Delayed
    }
    return nil
}

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-type-info"] = &trackInfo.TrackTypeInfo
    children["bool-tracks"] = &trackInfo.BoolTracks
    children["threshold-tracks"] = &trackInfo.ThresholdTracks
    children["tracking-interaces"] = &trackInfo.TrackingInteraces
    children["delayed"] = &trackInfo.Delayed
    return children
}

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tracke-name"] = trackInfo.TrackeName
    leafs["type"] = trackInfo.Type
    leafs["track-state"] = trackInfo.TrackState
    leafs["state-change-counter"] = trackInfo.StateChangeCounter
    leafs["seconds-last-change"] = trackInfo.SecondsLastChange
    leafs["threshold-up"] = trackInfo.ThresholdUp
    leafs["threshold-down"] = trackInfo.ThresholdDown
    return leafs
}

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetYangName() string { return "track-info" }

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) SetParent(parent types.Entity) { trackInfo.parent = parent }

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetParent() types.Entity { return trackInfo.parent }

func (trackInfo *ObjectTracking_Tracks_Track_TrackInfo) GetParentYangName() string { return "track" }

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo
// Track type information
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetFilter() yfilter.YFilter { return trackTypeInfo.YFilter }

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) SetFilter(yf yfilter.YFilter) { trackTypeInfo.YFilter = yf }

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetGoName(yname string) string {
    if yname == "discriminant" { return "Discriminant" }
    if yname == "interface-tracks" { return "InterfaceTracks" }
    if yname == "route-tracks" { return "RouteTracks" }
    if yname == "ipsla-tracks" { return "IpslaTracks" }
    if yname == "bfd-tracks" { return "BfdTracks" }
    return ""
}

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetSegmentPath() string {
    return "track-type-info"
}

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracks" {
        return &trackTypeInfo.InterfaceTracks
    }
    if childYangName == "route-tracks" {
        return &trackTypeInfo.RouteTracks
    }
    if childYangName == "ipsla-tracks" {
        return &trackTypeInfo.IpslaTracks
    }
    if childYangName == "bfd-tracks" {
        return &trackTypeInfo.BfdTracks
    }
    return nil
}

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-tracks"] = &trackTypeInfo.InterfaceTracks
    children["route-tracks"] = &trackTypeInfo.RouteTracks
    children["ipsla-tracks"] = &trackTypeInfo.IpslaTracks
    children["bfd-tracks"] = &trackTypeInfo.BfdTracks
    return children
}

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminant"] = trackTypeInfo.Discriminant
    return leafs
}

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetYangName() string { return "track-type-info" }

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) SetParent(parent types.Entity) { trackTypeInfo.parent = parent }

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetParent() types.Entity { return trackTypeInfo.parent }

func (trackTypeInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo) GetParentYangName() string { return "track-info" }

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetFilter() yfilter.YFilter { return interfaceTracks.YFilter }

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) SetFilter(yf yfilter.YFilter) { interfaceTracks.YFilter = yf }

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetSegmentPath() string {
    return "interface-tracks"
}

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTracks.InterfaceName
    return leafs
}

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetYangName() string { return "interface-tracks" }

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) SetParent(parent types.Entity) { interfaceTracks.parent = parent }

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetParent() types.Entity { return interfaceTracks.parent }

func (interfaceTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_InterfaceTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetFilter() yfilter.YFilter { return routeTracks.YFilter }

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) SetFilter(yf yfilter.YFilter) { routeTracks.YFilter = yf }

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "vrf" { return "Vrf" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetSegmentPath() string {
    return "route-tracks"
}

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = routeTracks.Prefix
    leafs["prefix-length"] = routeTracks.PrefixLength
    leafs["vrf"] = routeTracks.Vrf
    leafs["next-hop"] = routeTracks.NextHop
    return leafs
}

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetBundleName() string { return "cisco_ios_xr" }

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetYangName() string { return "route-tracks" }

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) SetParent(parent types.Entity) { routeTracks.parent = parent }

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetParent() types.Entity { return routeTracks.parent }

func (routeTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_RouteTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetFilter() yfilter.YFilter { return ipslaTracks.YFilter }

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) SetFilter(yf yfilter.YFilter) { ipslaTracks.YFilter = yf }

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetGoName(yname string) string {
    if yname == "ipsla-op-id" { return "IpslaOpId" }
    if yname == "rtt" { return "Rtt" }
    if yname == "return-code" { return "ReturnCode" }
    return ""
}

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetSegmentPath() string {
    return "ipsla-tracks"
}

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipsla-op-id"] = ipslaTracks.IpslaOpId
    leafs["rtt"] = ipslaTracks.Rtt
    leafs["return-code"] = ipslaTracks.ReturnCode
    return leafs
}

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetBundleName() string { return "cisco_ios_xr" }

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetYangName() string { return "ipsla-tracks" }

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) SetParent(parent types.Entity) { ipslaTracks.parent = parent }

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetParent() types.Entity { return ipslaTracks.parent }

func (ipslaTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_IpslaTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetFilter() yfilter.YFilter { return bfdTracks.YFilter }

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) SetFilter(yf yfilter.YFilter) { bfdTracks.YFilter = yf }

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "rate" { return "Rate" }
    if yname == "debounce-count" { return "DebounceCount" }
    return ""
}

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetSegmentPath() string {
    return "bfd-tracks"
}

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bfdTracks.InterfaceName
    leafs["destination-address"] = bfdTracks.DestinationAddress
    leafs["rate"] = bfdTracks.Rate
    leafs["debounce-count"] = bfdTracks.DebounceCount
    return leafs
}

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetYangName() string { return "bfd-tracks" }

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) SetParent(parent types.Entity) { bfdTracks.parent = parent }

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetParent() types.Entity { return bfdTracks.parent }

func (bfdTracks *ObjectTracking_Tracks_Track_TrackInfo_TrackTypeInfo_BfdTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_Tracks_Track_TrackInfo_BoolTracks
// boolean objects
type ObjectTracking_Tracks_Track_TrackInfo_BoolTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bool track info. The type is slice of
    // ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo.
    BoolTrackInfo []ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo
}

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetFilter() yfilter.YFilter { return boolTracks.YFilter }

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) SetFilter(yf yfilter.YFilter) { boolTracks.YFilter = yf }

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetGoName(yname string) string {
    if yname == "bool-track-info" { return "BoolTrackInfo" }
    return ""
}

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetSegmentPath() string {
    return "bool-tracks"
}

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bool-track-info" {
        for _, c := range boolTracks.BoolTrackInfo {
            if boolTracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo{}
        boolTracks.BoolTrackInfo = append(boolTracks.BoolTrackInfo, child)
        return &boolTracks.BoolTrackInfo[len(boolTracks.BoolTrackInfo)-1]
    }
    return nil
}

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range boolTracks.BoolTrackInfo {
        children[boolTracks.BoolTrackInfo[i].GetSegmentPath()] = &boolTracks.BoolTrackInfo[i]
    }
    return children
}

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetBundleName() string { return "cisco_ios_xr" }

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetYangName() string { return "bool-tracks" }

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) SetParent(parent types.Entity) { boolTracks.parent = parent }

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetParent() types.Entity { return boolTracks.parent }

func (boolTracks *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks) GetParentYangName() string { return "track-info" }

// ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo
// bool track info
type ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object Name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track object with Not. The type is bool.
    WithNot interface{}
}

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetFilter() yfilter.YFilter { return boolTrackInfo.YFilter }

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) SetFilter(yf yfilter.YFilter) { boolTrackInfo.YFilter = yf }

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "track-state" { return "TrackState" }
    if yname == "with-not" { return "WithNot" }
    return ""
}

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetSegmentPath() string {
    return "bool-track-info"
}

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = boolTrackInfo.ObjectName
    leafs["track-state"] = boolTrackInfo.TrackState
    leafs["with-not"] = boolTrackInfo.WithNot
    return leafs
}

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetYangName() string { return "bool-track-info" }

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) SetParent(parent types.Entity) { boolTrackInfo.parent = parent }

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetParent() types.Entity { return boolTrackInfo.parent }

func (boolTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_BoolTracks_BoolTrackInfo) GetParentYangName() string { return "bool-tracks" }

// ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks
// Threshold objects
type ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // threshold track info. The type is slice of
    // ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo.
    ThresholdTrackInfo []ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo
}

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetFilter() yfilter.YFilter { return thresholdTracks.YFilter }

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) SetFilter(yf yfilter.YFilter) { thresholdTracks.YFilter = yf }

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetGoName(yname string) string {
    if yname == "threshold-track-info" { return "ThresholdTrackInfo" }
    return ""
}

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetSegmentPath() string {
    return "threshold-tracks"
}

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-track-info" {
        for _, c := range thresholdTracks.ThresholdTrackInfo {
            if thresholdTracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo{}
        thresholdTracks.ThresholdTrackInfo = append(thresholdTracks.ThresholdTrackInfo, child)
        return &thresholdTracks.ThresholdTrackInfo[len(thresholdTracks.ThresholdTrackInfo)-1]
    }
    return nil
}

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholdTracks.ThresholdTrackInfo {
        children[thresholdTracks.ThresholdTrackInfo[i].GetSegmentPath()] = &thresholdTracks.ThresholdTrackInfo[i]
    }
    return children
}

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetYangName() string { return "threshold-tracks" }

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) SetParent(parent types.Entity) { thresholdTracks.parent = parent }

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetParent() types.Entity { return thresholdTracks.parent }

func (thresholdTracks *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks) GetParentYangName() string { return "track-info" }

// ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo
// threshold track info
type ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. True means track is up; False means track is down. The type is
    // bool.
    TrackState interface{}

    // Weight is the number assigned to a track object . In case of a type
    // threshold weight( i.e. weighted sum list), weight is asigned by User at the
    // time of configuration. In case of a type threshold percentage (i.e.
    // percentage based list), weight is internally computed by (1/N)x100, where N
    // is the number of objects in the list. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Weight interface{}
}

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetFilter() yfilter.YFilter { return thresholdTrackInfo.YFilter }

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) SetFilter(yf yfilter.YFilter) { thresholdTrackInfo.YFilter = yf }

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "track-state" { return "TrackState" }
    if yname == "weight" { return "Weight" }
    return ""
}

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetSegmentPath() string {
    return "threshold-track-info"
}

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = thresholdTrackInfo.ObjectName
    leafs["track-state"] = thresholdTrackInfo.TrackState
    leafs["weight"] = thresholdTrackInfo.Weight
    return leafs
}

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetYangName() string { return "threshold-track-info" }

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) SetParent(parent types.Entity) { thresholdTrackInfo.parent = parent }

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetParent() types.Entity { return thresholdTrackInfo.parent }

func (thresholdTrackInfo *ObjectTracking_Tracks_Track_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetParentYangName() string { return "threshold-tracks" }

// ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces
// Tracking Interfaces
type ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface tracking info. The type is slice of
    // ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo.
    InterfaceTrackingInfo []ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
}

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetFilter() yfilter.YFilter { return trackingInteraces.YFilter }

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) SetFilter(yf yfilter.YFilter) { trackingInteraces.YFilter = yf }

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetGoName(yname string) string {
    if yname == "interface-tracking-info" { return "InterfaceTrackingInfo" }
    return ""
}

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetSegmentPath() string {
    return "tracking-interaces"
}

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracking-info" {
        for _, c := range trackingInteraces.InterfaceTrackingInfo {
            if trackingInteraces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo{}
        trackingInteraces.InterfaceTrackingInfo = append(trackingInteraces.InterfaceTrackingInfo, child)
        return &trackingInteraces.InterfaceTrackingInfo[len(trackingInteraces.InterfaceTrackingInfo)-1]
    }
    return nil
}

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackingInteraces.InterfaceTrackingInfo {
        children[trackingInteraces.InterfaceTrackingInfo[i].GetSegmentPath()] = &trackingInteraces.InterfaceTrackingInfo[i]
    }
    return children
}

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetBundleName() string { return "cisco_ios_xr" }

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetYangName() string { return "tracking-interaces" }

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) SetParent(parent types.Entity) { trackingInteraces.parent = parent }

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetParent() types.Entity { return trackingInteraces.parent }

func (trackingInteraces *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces) GetParentYangName() string { return "track-info" }

// ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
// interface tracking info
type ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetFilter() yfilter.YFilter { return interfaceTrackingInfo.YFilter }

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) SetFilter(yf yfilter.YFilter) { interfaceTrackingInfo.YFilter = yf }

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetSegmentPath() string {
    return "interface-tracking-info"
}

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTrackingInfo.InterfaceName
    return leafs
}

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetYangName() string { return "interface-tracking-info" }

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) SetParent(parent types.Entity) { interfaceTrackingInfo.parent = parent }

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetParent() types.Entity { return interfaceTrackingInfo.parent }

func (interfaceTrackingInfo *ObjectTracking_Tracks_Track_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetParentYangName() string { return "tracking-interaces" }

// ObjectTracking_Tracks_Track_TrackInfo_Delayed
// Is the state change delay counter in progress
type ObjectTracking_Tracks_Track_TrackInfo_Delayed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The time remaining in seconds for the counter to trigger state change. The
    // type is interface{} with range: 0..4294967295. Units are second.
    TimeRemaining interface{}

    // State the track will transition to. Track state. True means track is up;
    // False means track is down. The type is bool.
    TrackState interface{}
}

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetFilter() yfilter.YFilter { return delayed.YFilter }

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) SetFilter(yf yfilter.YFilter) { delayed.YFilter = yf }

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetGoName(yname string) string {
    if yname == "time-remaining" { return "TimeRemaining" }
    if yname == "track-state" { return "TrackState" }
    return ""
}

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetSegmentPath() string {
    return "delayed"
}

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-remaining"] = delayed.TimeRemaining
    leafs["track-state"] = delayed.TrackState
    return leafs
}

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetBundleName() string { return "cisco_ios_xr" }

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetYangName() string { return "delayed" }

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) SetParent(parent types.Entity) { delayed.parent = parent }

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetParent() types.Entity { return delayed.parent }

func (delayed *ObjectTracking_Tracks_Track_TrackInfo_Delayed) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeIpv4RouteBrief
// Object Tracking Type Ipv4 Route brief info
type ObjectTracking_TrackTypeIpv4RouteBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // track info brief. The type is slice of
    // ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief.
    TrackInfoBrief []ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief
}

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetFilter() yfilter.YFilter { return trackTypeIpv4RouteBrief.YFilter }

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) SetFilter(yf yfilter.YFilter) { trackTypeIpv4RouteBrief.YFilter = yf }

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetGoName(yname string) string {
    if yname == "track-info-brief" { return "TrackInfoBrief" }
    return ""
}

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetSegmentPath() string {
    return "track-type-ipv4-route-brief"
}

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-info-brief" {
        for _, c := range trackTypeIpv4RouteBrief.TrackInfoBrief {
            if trackTypeIpv4RouteBrief.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief{}
        trackTypeIpv4RouteBrief.TrackInfoBrief = append(trackTypeIpv4RouteBrief.TrackInfoBrief, child)
        return &trackTypeIpv4RouteBrief.TrackInfoBrief[len(trackTypeIpv4RouteBrief.TrackInfoBrief)-1]
    }
    return nil
}

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackTypeIpv4RouteBrief.TrackInfoBrief {
        children[trackTypeIpv4RouteBrief.TrackInfoBrief[i].GetSegmentPath()] = &trackTypeIpv4RouteBrief.TrackInfoBrief[i]
    }
    return children
}

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetYangName() string { return "track-type-ipv4-route-brief" }

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) SetParent(parent types.Entity) { trackTypeIpv4RouteBrief.parent = parent }

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetParent() types.Entity { return trackTypeIpv4RouteBrief.parent }

func (trackTypeIpv4RouteBrief *ObjectTracking_TrackTypeIpv4RouteBrief) GetParentYangName() string { return "object-tracking" }

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief
// track info brief
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo
}

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetFilter() yfilter.YFilter { return trackInfoBrief.YFilter }

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) SetFilter(yf yfilter.YFilter) { trackInfoBrief.YFilter = yf }

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetGoName(yname string) string {
    if yname == "tracke-name" { return "TrackeName" }
    if yname == "type" { return "Type" }
    if yname == "track-state" { return "TrackState" }
    if yname == "track-type-info" { return "TrackTypeInfo" }
    return ""
}

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetSegmentPath() string {
    return "track-info-brief"
}

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-type-info" {
        return &trackInfoBrief.TrackTypeInfo
    }
    return nil
}

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-type-info"] = &trackInfoBrief.TrackTypeInfo
    return children
}

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tracke-name"] = trackInfoBrief.TrackeName
    leafs["type"] = trackInfoBrief.Type
    leafs["track-state"] = trackInfoBrief.TrackState
    return leafs
}

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetBundleName() string { return "cisco_ios_xr" }

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetYangName() string { return "track-info-brief" }

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) SetParent(parent types.Entity) { trackInfoBrief.parent = parent }

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetParent() types.Entity { return trackInfoBrief.parent }

func (trackInfoBrief *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief) GetParentYangName() string { return "track-type-ipv4-route-brief" }

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetFilter() yfilter.YFilter { return trackTypeInfo.YFilter }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) SetFilter(yf yfilter.YFilter) { trackTypeInfo.YFilter = yf }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetGoName(yname string) string {
    if yname == "discriminant" { return "Discriminant" }
    if yname == "interface-tracks" { return "InterfaceTracks" }
    if yname == "route-tracks" { return "RouteTracks" }
    if yname == "ipsla-tracks" { return "IpslaTracks" }
    if yname == "bfd-tracks" { return "BfdTracks" }
    return ""
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetSegmentPath() string {
    return "track-type-info"
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracks" {
        return &trackTypeInfo.InterfaceTracks
    }
    if childYangName == "route-tracks" {
        return &trackTypeInfo.RouteTracks
    }
    if childYangName == "ipsla-tracks" {
        return &trackTypeInfo.IpslaTracks
    }
    if childYangName == "bfd-tracks" {
        return &trackTypeInfo.BfdTracks
    }
    return nil
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-tracks"] = &trackTypeInfo.InterfaceTracks
    children["route-tracks"] = &trackTypeInfo.RouteTracks
    children["ipsla-tracks"] = &trackTypeInfo.IpslaTracks
    children["bfd-tracks"] = &trackTypeInfo.BfdTracks
    return children
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminant"] = trackTypeInfo.Discriminant
    return leafs
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetYangName() string { return "track-type-info" }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) SetParent(parent types.Entity) { trackTypeInfo.parent = parent }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetParent() types.Entity { return trackTypeInfo.parent }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo) GetParentYangName() string { return "track-info-brief" }

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetFilter() yfilter.YFilter { return interfaceTracks.YFilter }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) SetFilter(yf yfilter.YFilter) { interfaceTracks.YFilter = yf }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetSegmentPath() string {
    return "interface-tracks"
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTracks.InterfaceName
    return leafs
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetYangName() string { return "interface-tracks" }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) SetParent(parent types.Entity) { interfaceTracks.parent = parent }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetParent() types.Entity { return interfaceTracks.parent }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetFilter() yfilter.YFilter { return routeTracks.YFilter }

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) SetFilter(yf yfilter.YFilter) { routeTracks.YFilter = yf }

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "vrf" { return "Vrf" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetSegmentPath() string {
    return "route-tracks"
}

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = routeTracks.Prefix
    leafs["prefix-length"] = routeTracks.PrefixLength
    leafs["vrf"] = routeTracks.Vrf
    leafs["next-hop"] = routeTracks.NextHop
    return leafs
}

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetBundleName() string { return "cisco_ios_xr" }

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetYangName() string { return "route-tracks" }

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) SetParent(parent types.Entity) { routeTracks.parent = parent }

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetParent() types.Entity { return routeTracks.parent }

func (routeTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetFilter() yfilter.YFilter { return ipslaTracks.YFilter }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) SetFilter(yf yfilter.YFilter) { ipslaTracks.YFilter = yf }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetGoName(yname string) string {
    if yname == "ipsla-op-id" { return "IpslaOpId" }
    if yname == "rtt" { return "Rtt" }
    if yname == "return-code" { return "ReturnCode" }
    return ""
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetSegmentPath() string {
    return "ipsla-tracks"
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipsla-op-id"] = ipslaTracks.IpslaOpId
    leafs["rtt"] = ipslaTracks.Rtt
    leafs["return-code"] = ipslaTracks.ReturnCode
    return leafs
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetBundleName() string { return "cisco_ios_xr" }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetYangName() string { return "ipsla-tracks" }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) SetParent(parent types.Entity) { ipslaTracks.parent = parent }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetParent() types.Entity { return ipslaTracks.parent }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetFilter() yfilter.YFilter { return bfdTracks.YFilter }

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) SetFilter(yf yfilter.YFilter) { bfdTracks.YFilter = yf }

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "rate" { return "Rate" }
    if yname == "debounce-count" { return "DebounceCount" }
    return ""
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetSegmentPath() string {
    return "bfd-tracks"
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bfdTracks.InterfaceName
    leafs["destination-address"] = bfdTracks.DestinationAddress
    leafs["rate"] = bfdTracks.Rate
    leafs["debounce-count"] = bfdTracks.DebounceCount
    return leafs
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetYangName() string { return "bfd-tracks" }

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) SetParent(parent types.Entity) { bfdTracks.parent = parent }

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetParent() types.Entity { return bfdTracks.parent }

func (bfdTracks *ObjectTracking_TrackTypeIpv4RouteBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeIpv4Route
// Object Tracking Type IPV4 route info
type ObjectTracking_TrackTypeIpv4Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // track info. The type is slice of
    // ObjectTracking_TrackTypeIpv4Route_TrackInfo.
    TrackInfo []ObjectTracking_TrackTypeIpv4Route_TrackInfo
}

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetFilter() yfilter.YFilter { return trackTypeIpv4Route.YFilter }

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) SetFilter(yf yfilter.YFilter) { trackTypeIpv4Route.YFilter = yf }

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetGoName(yname string) string {
    if yname == "track-info" { return "TrackInfo" }
    return ""
}

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetSegmentPath() string {
    return "track-type-ipv4-route"
}

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-info" {
        for _, c := range trackTypeIpv4Route.TrackInfo {
            if trackTypeIpv4Route.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeIpv4Route_TrackInfo{}
        trackTypeIpv4Route.TrackInfo = append(trackTypeIpv4Route.TrackInfo, child)
        return &trackTypeIpv4Route.TrackInfo[len(trackTypeIpv4Route.TrackInfo)-1]
    }
    return nil
}

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackTypeIpv4Route.TrackInfo {
        children[trackTypeIpv4Route.TrackInfo[i].GetSegmentPath()] = &trackTypeIpv4Route.TrackInfo[i]
    }
    return children
}

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetYangName() string { return "track-type-ipv4-route" }

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) SetParent(parent types.Entity) { trackTypeIpv4Route.parent = parent }

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetParent() types.Entity { return trackTypeIpv4Route.parent }

func (trackTypeIpv4Route *ObjectTracking_TrackTypeIpv4Route) GetParentYangName() string { return "object-tracking" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo
// track info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // State Change Counter. The type is interface{} with range: 0..4294967295.
    StateChangeCounter interface{}

    // Seconds Last Change. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SecondsLastChange interface{}

    // User specified threshold upper limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdUp interface{}

    // User specified threshold lower limit. The type is interface{} with range:
    // 0..4294967295.
    ThresholdDown interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo

    // boolean objects.
    BoolTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks

    // Threshold objects.
    ThresholdTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks

    // Tracking Interfaces.
    TrackingInteraces ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces

    // Is the state change delay counter in progress.
    Delayed ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed
}

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetFilter() yfilter.YFilter { return trackInfo.YFilter }

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) SetFilter(yf yfilter.YFilter) { trackInfo.YFilter = yf }

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetGoName(yname string) string {
    if yname == "tracke-name" { return "TrackeName" }
    if yname == "type" { return "Type" }
    if yname == "track-state" { return "TrackState" }
    if yname == "state-change-counter" { return "StateChangeCounter" }
    if yname == "seconds-last-change" { return "SecondsLastChange" }
    if yname == "threshold-up" { return "ThresholdUp" }
    if yname == "threshold-down" { return "ThresholdDown" }
    if yname == "track-type-info" { return "TrackTypeInfo" }
    if yname == "bool-tracks" { return "BoolTracks" }
    if yname == "threshold-tracks" { return "ThresholdTracks" }
    if yname == "tracking-interaces" { return "TrackingInteraces" }
    if yname == "delayed" { return "Delayed" }
    return ""
}

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetSegmentPath() string {
    return "track-info"
}

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-type-info" {
        return &trackInfo.TrackTypeInfo
    }
    if childYangName == "bool-tracks" {
        return &trackInfo.BoolTracks
    }
    if childYangName == "threshold-tracks" {
        return &trackInfo.ThresholdTracks
    }
    if childYangName == "tracking-interaces" {
        return &trackInfo.TrackingInteraces
    }
    if childYangName == "delayed" {
        return &trackInfo.Delayed
    }
    return nil
}

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-type-info"] = &trackInfo.TrackTypeInfo
    children["bool-tracks"] = &trackInfo.BoolTracks
    children["threshold-tracks"] = &trackInfo.ThresholdTracks
    children["tracking-interaces"] = &trackInfo.TrackingInteraces
    children["delayed"] = &trackInfo.Delayed
    return children
}

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tracke-name"] = trackInfo.TrackeName
    leafs["type"] = trackInfo.Type
    leafs["track-state"] = trackInfo.TrackState
    leafs["state-change-counter"] = trackInfo.StateChangeCounter
    leafs["seconds-last-change"] = trackInfo.SecondsLastChange
    leafs["threshold-up"] = trackInfo.ThresholdUp
    leafs["threshold-down"] = trackInfo.ThresholdDown
    return leafs
}

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetYangName() string { return "track-info" }

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) SetParent(parent types.Entity) { trackInfo.parent = parent }

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetParent() types.Entity { return trackInfo.parent }

func (trackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo) GetParentYangName() string { return "track-type-ipv4-route" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetFilter() yfilter.YFilter { return trackTypeInfo.YFilter }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) SetFilter(yf yfilter.YFilter) { trackTypeInfo.YFilter = yf }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetGoName(yname string) string {
    if yname == "discriminant" { return "Discriminant" }
    if yname == "interface-tracks" { return "InterfaceTracks" }
    if yname == "route-tracks" { return "RouteTracks" }
    if yname == "ipsla-tracks" { return "IpslaTracks" }
    if yname == "bfd-tracks" { return "BfdTracks" }
    return ""
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetSegmentPath() string {
    return "track-type-info"
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracks" {
        return &trackTypeInfo.InterfaceTracks
    }
    if childYangName == "route-tracks" {
        return &trackTypeInfo.RouteTracks
    }
    if childYangName == "ipsla-tracks" {
        return &trackTypeInfo.IpslaTracks
    }
    if childYangName == "bfd-tracks" {
        return &trackTypeInfo.BfdTracks
    }
    return nil
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-tracks"] = &trackTypeInfo.InterfaceTracks
    children["route-tracks"] = &trackTypeInfo.RouteTracks
    children["ipsla-tracks"] = &trackTypeInfo.IpslaTracks
    children["bfd-tracks"] = &trackTypeInfo.BfdTracks
    return children
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminant"] = trackTypeInfo.Discriminant
    return leafs
}

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetYangName() string { return "track-type-info" }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) SetParent(parent types.Entity) { trackTypeInfo.parent = parent }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetParent() types.Entity { return trackTypeInfo.parent }

func (trackTypeInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetFilter() yfilter.YFilter { return interfaceTracks.YFilter }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) SetFilter(yf yfilter.YFilter) { interfaceTracks.YFilter = yf }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetSegmentPath() string {
    return "interface-tracks"
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTracks.InterfaceName
    return leafs
}

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetYangName() string { return "interface-tracks" }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) SetParent(parent types.Entity) { interfaceTracks.parent = parent }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetParent() types.Entity { return interfaceTracks.parent }

func (interfaceTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_InterfaceTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetFilter() yfilter.YFilter { return routeTracks.YFilter }

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) SetFilter(yf yfilter.YFilter) { routeTracks.YFilter = yf }

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "vrf" { return "Vrf" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetSegmentPath() string {
    return "route-tracks"
}

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = routeTracks.Prefix
    leafs["prefix-length"] = routeTracks.PrefixLength
    leafs["vrf"] = routeTracks.Vrf
    leafs["next-hop"] = routeTracks.NextHop
    return leafs
}

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetBundleName() string { return "cisco_ios_xr" }

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetYangName() string { return "route-tracks" }

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) SetParent(parent types.Entity) { routeTracks.parent = parent }

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetParent() types.Entity { return routeTracks.parent }

func (routeTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_RouteTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetFilter() yfilter.YFilter { return ipslaTracks.YFilter }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) SetFilter(yf yfilter.YFilter) { ipslaTracks.YFilter = yf }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetGoName(yname string) string {
    if yname == "ipsla-op-id" { return "IpslaOpId" }
    if yname == "rtt" { return "Rtt" }
    if yname == "return-code" { return "ReturnCode" }
    return ""
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetSegmentPath() string {
    return "ipsla-tracks"
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipsla-op-id"] = ipslaTracks.IpslaOpId
    leafs["rtt"] = ipslaTracks.Rtt
    leafs["return-code"] = ipslaTracks.ReturnCode
    return leafs
}

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetBundleName() string { return "cisco_ios_xr" }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetYangName() string { return "ipsla-tracks" }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) SetParent(parent types.Entity) { ipslaTracks.parent = parent }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetParent() types.Entity { return ipslaTracks.parent }

func (ipslaTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_IpslaTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetFilter() yfilter.YFilter { return bfdTracks.YFilter }

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) SetFilter(yf yfilter.YFilter) { bfdTracks.YFilter = yf }

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "rate" { return "Rate" }
    if yname == "debounce-count" { return "DebounceCount" }
    return ""
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetSegmentPath() string {
    return "bfd-tracks"
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bfdTracks.InterfaceName
    leafs["destination-address"] = bfdTracks.DestinationAddress
    leafs["rate"] = bfdTracks.Rate
    leafs["debounce-count"] = bfdTracks.DebounceCount
    return leafs
}

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetYangName() string { return "bfd-tracks" }

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) SetParent(parent types.Entity) { bfdTracks.parent = parent }

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetParent() types.Entity { return bfdTracks.parent }

func (bfdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackTypeInfo_BfdTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks
// boolean objects
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // bool track info. The type is slice of
    // ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo.
    BoolTrackInfo []ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo
}

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetFilter() yfilter.YFilter { return boolTracks.YFilter }

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) SetFilter(yf yfilter.YFilter) { boolTracks.YFilter = yf }

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetGoName(yname string) string {
    if yname == "bool-track-info" { return "BoolTrackInfo" }
    return ""
}

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetSegmentPath() string {
    return "bool-tracks"
}

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bool-track-info" {
        for _, c := range boolTracks.BoolTrackInfo {
            if boolTracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo{}
        boolTracks.BoolTrackInfo = append(boolTracks.BoolTrackInfo, child)
        return &boolTracks.BoolTrackInfo[len(boolTracks.BoolTrackInfo)-1]
    }
    return nil
}

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range boolTracks.BoolTrackInfo {
        children[boolTracks.BoolTrackInfo[i].GetSegmentPath()] = &boolTracks.BoolTrackInfo[i]
    }
    return children
}

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetBundleName() string { return "cisco_ios_xr" }

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetYangName() string { return "bool-tracks" }

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) SetParent(parent types.Entity) { boolTracks.parent = parent }

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetParent() types.Entity { return boolTracks.parent }

func (boolTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo
// bool track info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object Name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track object with Not. The type is bool.
    WithNot interface{}
}

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetFilter() yfilter.YFilter { return boolTrackInfo.YFilter }

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) SetFilter(yf yfilter.YFilter) { boolTrackInfo.YFilter = yf }

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "track-state" { return "TrackState" }
    if yname == "with-not" { return "WithNot" }
    return ""
}

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetSegmentPath() string {
    return "bool-track-info"
}

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = boolTrackInfo.ObjectName
    leafs["track-state"] = boolTrackInfo.TrackState
    leafs["with-not"] = boolTrackInfo.WithNot
    return leafs
}

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetYangName() string { return "bool-track-info" }

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) SetParent(parent types.Entity) { boolTrackInfo.parent = parent }

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetParent() types.Entity { return boolTrackInfo.parent }

func (boolTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_BoolTracks_BoolTrackInfo) GetParentYangName() string { return "bool-tracks" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks
// Threshold objects
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // threshold track info. The type is slice of
    // ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo.
    ThresholdTrackInfo []ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo
}

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetFilter() yfilter.YFilter { return thresholdTracks.YFilter }

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) SetFilter(yf yfilter.YFilter) { thresholdTracks.YFilter = yf }

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetGoName(yname string) string {
    if yname == "threshold-track-info" { return "ThresholdTrackInfo" }
    return ""
}

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetSegmentPath() string {
    return "threshold-tracks"
}

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-track-info" {
        for _, c := range thresholdTracks.ThresholdTrackInfo {
            if thresholdTracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo{}
        thresholdTracks.ThresholdTrackInfo = append(thresholdTracks.ThresholdTrackInfo, child)
        return &thresholdTracks.ThresholdTrackInfo[len(thresholdTracks.ThresholdTrackInfo)-1]
    }
    return nil
}

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholdTracks.ThresholdTrackInfo {
        children[thresholdTracks.ThresholdTrackInfo[i].GetSegmentPath()] = &thresholdTracks.ThresholdTrackInfo[i]
    }
    return children
}

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetYangName() string { return "threshold-tracks" }

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) SetParent(parent types.Entity) { thresholdTracks.parent = parent }

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetParent() types.Entity { return thresholdTracks.parent }

func (thresholdTracks *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo
// threshold track info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object name. The type is string with length: 0..33.
    ObjectName interface{}

    // Track state. True means track is up; False means track is down. The type is
    // bool.
    TrackState interface{}

    // Weight is the number assigned to a track object . In case of a type
    // threshold weight( i.e. weighted sum list), weight is asigned by User at the
    // time of configuration. In case of a type threshold percentage (i.e.
    // percentage based list), weight is internally computed by (1/N)x100, where N
    // is the number of objects in the list. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    Weight interface{}
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetFilter() yfilter.YFilter { return thresholdTrackInfo.YFilter }

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) SetFilter(yf yfilter.YFilter) { thresholdTrackInfo.YFilter = yf }

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "track-state" { return "TrackState" }
    if yname == "weight" { return "Weight" }
    return ""
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetSegmentPath() string {
    return "threshold-track-info"
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = thresholdTrackInfo.ObjectName
    leafs["track-state"] = thresholdTrackInfo.TrackState
    leafs["weight"] = thresholdTrackInfo.Weight
    return leafs
}

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetYangName() string { return "threshold-track-info" }

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) SetParent(parent types.Entity) { thresholdTrackInfo.parent = parent }

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetParent() types.Entity { return thresholdTrackInfo.parent }

func (thresholdTrackInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_ThresholdTracks_ThresholdTrackInfo) GetParentYangName() string { return "threshold-tracks" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces
// Tracking Interfaces
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface tracking info. The type is slice of
    // ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo.
    InterfaceTrackingInfo []ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
}

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetFilter() yfilter.YFilter { return trackingInteraces.YFilter }

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) SetFilter(yf yfilter.YFilter) { trackingInteraces.YFilter = yf }

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetGoName(yname string) string {
    if yname == "interface-tracking-info" { return "InterfaceTrackingInfo" }
    return ""
}

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetSegmentPath() string {
    return "tracking-interaces"
}

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracking-info" {
        for _, c := range trackingInteraces.InterfaceTrackingInfo {
            if trackingInteraces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo{}
        trackingInteraces.InterfaceTrackingInfo = append(trackingInteraces.InterfaceTrackingInfo, child)
        return &trackingInteraces.InterfaceTrackingInfo[len(trackingInteraces.InterfaceTrackingInfo)-1]
    }
    return nil
}

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackingInteraces.InterfaceTrackingInfo {
        children[trackingInteraces.InterfaceTrackingInfo[i].GetSegmentPath()] = &trackingInteraces.InterfaceTrackingInfo[i]
    }
    return children
}

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetBundleName() string { return "cisco_ios_xr" }

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetYangName() string { return "tracking-interaces" }

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) SetParent(parent types.Entity) { trackingInteraces.parent = parent }

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetParent() types.Entity { return trackingInteraces.parent }

func (trackingInteraces *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo
// interface tracking info
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetFilter() yfilter.YFilter { return interfaceTrackingInfo.YFilter }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) SetFilter(yf yfilter.YFilter) { interfaceTrackingInfo.YFilter = yf }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetSegmentPath() string {
    return "interface-tracking-info"
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTrackingInfo.InterfaceName
    return leafs
}

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetYangName() string { return "interface-tracking-info" }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) SetParent(parent types.Entity) { interfaceTrackingInfo.parent = parent }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetParent() types.Entity { return interfaceTrackingInfo.parent }

func (interfaceTrackingInfo *ObjectTracking_TrackTypeIpv4Route_TrackInfo_TrackingInteraces_InterfaceTrackingInfo) GetParentYangName() string { return "tracking-interaces" }

// ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed
// Is the state change delay counter in progress
type ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The time remaining in seconds for the counter to trigger state change. The
    // type is interface{} with range: 0..4294967295. Units are second.
    TimeRemaining interface{}

    // State the track will transition to. Track state. True means track is up;
    // False means track is down. The type is bool.
    TrackState interface{}
}

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetFilter() yfilter.YFilter { return delayed.YFilter }

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) SetFilter(yf yfilter.YFilter) { delayed.YFilter = yf }

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetGoName(yname string) string {
    if yname == "time-remaining" { return "TimeRemaining" }
    if yname == "track-state" { return "TrackState" }
    return ""
}

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetSegmentPath() string {
    return "delayed"
}

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-remaining"] = delayed.TimeRemaining
    leafs["track-state"] = delayed.TrackState
    return leafs
}

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetBundleName() string { return "cisco_ios_xr" }

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetYangName() string { return "delayed" }

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) SetParent(parent types.Entity) { delayed.parent = parent }

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetParent() types.Entity { return delayed.parent }

func (delayed *ObjectTracking_TrackTypeIpv4Route_TrackInfo_Delayed) GetParentYangName() string { return "track-info" }

// ObjectTracking_TrackTypeInterfaceBrief
// Object Tracking Type Interface brief info
type ObjectTracking_TrackTypeInterfaceBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // track info brief. The type is slice of
    // ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief.
    TrackInfoBrief []ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief
}

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetFilter() yfilter.YFilter { return trackTypeInterfaceBrief.YFilter }

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) SetFilter(yf yfilter.YFilter) { trackTypeInterfaceBrief.YFilter = yf }

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetGoName(yname string) string {
    if yname == "track-info-brief" { return "TrackInfoBrief" }
    return ""
}

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetSegmentPath() string {
    return "track-type-interface-brief"
}

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-info-brief" {
        for _, c := range trackTypeInterfaceBrief.TrackInfoBrief {
            if trackTypeInterfaceBrief.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief{}
        trackTypeInterfaceBrief.TrackInfoBrief = append(trackTypeInterfaceBrief.TrackInfoBrief, child)
        return &trackTypeInterfaceBrief.TrackInfoBrief[len(trackTypeInterfaceBrief.TrackInfoBrief)-1]
    }
    return nil
}

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackTypeInterfaceBrief.TrackInfoBrief {
        children[trackTypeInterfaceBrief.TrackInfoBrief[i].GetSegmentPath()] = &trackTypeInterfaceBrief.TrackInfoBrief[i]
    }
    return children
}

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetYangName() string { return "track-type-interface-brief" }

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) SetParent(parent types.Entity) { trackTypeInterfaceBrief.parent = parent }

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetParent() types.Entity { return trackTypeInterfaceBrief.parent }

func (trackTypeInterfaceBrief *ObjectTracking_TrackTypeInterfaceBrief) GetParentYangName() string { return "object-tracking" }

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief
// track info brief
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track Name. The type is string with length: 0..33.
    TrackeName interface{}

    // Track type. The type is Track.
    Type interface{}

    // Track state. The type is bool.
    TrackState interface{}

    // Track type information.
    TrackTypeInfo ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo
}

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetFilter() yfilter.YFilter { return trackInfoBrief.YFilter }

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) SetFilter(yf yfilter.YFilter) { trackInfoBrief.YFilter = yf }

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetGoName(yname string) string {
    if yname == "tracke-name" { return "TrackeName" }
    if yname == "type" { return "Type" }
    if yname == "track-state" { return "TrackState" }
    if yname == "track-type-info" { return "TrackTypeInfo" }
    return ""
}

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetSegmentPath() string {
    return "track-info-brief"
}

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track-type-info" {
        return &trackInfoBrief.TrackTypeInfo
    }
    return nil
}

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["track-type-info"] = &trackInfoBrief.TrackTypeInfo
    return children
}

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tracke-name"] = trackInfoBrief.TrackeName
    leafs["type"] = trackInfoBrief.Type
    leafs["track-state"] = trackInfoBrief.TrackState
    return leafs
}

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetBundleName() string { return "cisco_ios_xr" }

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetYangName() string { return "track-info-brief" }

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) SetParent(parent types.Entity) { trackInfoBrief.parent = parent }

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetParent() types.Entity { return trackInfoBrief.parent }

func (trackInfoBrief *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief) GetParentYangName() string { return "track-type-interface-brief" }

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo
// Track type information
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // discriminant. The type is Track.
    Discriminant interface{}

    // track type interface info.
    InterfaceTracks ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks

    // track type route info.
    RouteTracks ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks

    // track type rtr info.
    IpslaTracks ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks

    // track type bfdrtr info.
    BfdTracks ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetFilter() yfilter.YFilter { return trackTypeInfo.YFilter }

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) SetFilter(yf yfilter.YFilter) { trackTypeInfo.YFilter = yf }

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetGoName(yname string) string {
    if yname == "discriminant" { return "Discriminant" }
    if yname == "interface-tracks" { return "InterfaceTracks" }
    if yname == "route-tracks" { return "RouteTracks" }
    if yname == "ipsla-tracks" { return "IpslaTracks" }
    if yname == "bfd-tracks" { return "BfdTracks" }
    return ""
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetSegmentPath() string {
    return "track-type-info"
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-tracks" {
        return &trackTypeInfo.InterfaceTracks
    }
    if childYangName == "route-tracks" {
        return &trackTypeInfo.RouteTracks
    }
    if childYangName == "ipsla-tracks" {
        return &trackTypeInfo.IpslaTracks
    }
    if childYangName == "bfd-tracks" {
        return &trackTypeInfo.BfdTracks
    }
    return nil
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-tracks"] = &trackTypeInfo.InterfaceTracks
    children["route-tracks"] = &trackTypeInfo.RouteTracks
    children["ipsla-tracks"] = &trackTypeInfo.IpslaTracks
    children["bfd-tracks"] = &trackTypeInfo.BfdTracks
    return children
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminant"] = trackTypeInfo.Discriminant
    return leafs
}

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetYangName() string { return "track-type-info" }

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) SetParent(parent types.Entity) { trackTypeInfo.parent = parent }

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetParent() types.Entity { return trackTypeInfo.parent }

func (trackTypeInfo *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo) GetParentYangName() string { return "track-info-brief" }

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks
// track type interface info
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}
}

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetFilter() yfilter.YFilter { return interfaceTracks.YFilter }

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) SetFilter(yf yfilter.YFilter) { interfaceTracks.YFilter = yf }

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetSegmentPath() string {
    return "interface-tracks"
}

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceTracks.InterfaceName
    return leafs
}

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetYangName() string { return "interface-tracks" }

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) SetParent(parent types.Entity) { interfaceTracks.parent = parent }

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetParent() types.Entity { return interfaceTracks.parent }

func (interfaceTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_InterfaceTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks
// track type route info
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // Prefix Length. The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // VRF Name. The type is string with length: 0..120.
    Vrf interface{}

    // Next Hop. The type is string with length: 0..120.
    NextHop interface{}
}

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetFilter() yfilter.YFilter { return routeTracks.YFilter }

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) SetFilter(yf yfilter.YFilter) { routeTracks.YFilter = yf }

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "vrf" { return "Vrf" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetSegmentPath() string {
    return "route-tracks"
}

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = routeTracks.Prefix
    leafs["prefix-length"] = routeTracks.PrefixLength
    leafs["vrf"] = routeTracks.Vrf
    leafs["next-hop"] = routeTracks.NextHop
    return leafs
}

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetBundleName() string { return "cisco_ios_xr" }

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetYangName() string { return "route-tracks" }

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) SetParent(parent types.Entity) { routeTracks.parent = parent }

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetParent() types.Entity { return routeTracks.parent }

func (routeTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_RouteTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks
// track type rtr info
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Op Id. The type is interface{} with range: 0..4294967295.
    IpslaOpId interface{}

    // Latest RTT. The type is interface{} with range: 0..4294967295.
    Rtt interface{}

    // Latest Return Code. The type is interface{} with range: 0..4294967295.
    ReturnCode interface{}
}

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetFilter() yfilter.YFilter { return ipslaTracks.YFilter }

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) SetFilter(yf yfilter.YFilter) { ipslaTracks.YFilter = yf }

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetGoName(yname string) string {
    if yname == "ipsla-op-id" { return "IpslaOpId" }
    if yname == "rtt" { return "Rtt" }
    if yname == "return-code" { return "ReturnCode" }
    return ""
}

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetSegmentPath() string {
    return "ipsla-tracks"
}

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipsla-op-id"] = ipslaTracks.IpslaOpId
    leafs["rtt"] = ipslaTracks.Rtt
    leafs["return-code"] = ipslaTracks.ReturnCode
    return leafs
}

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetBundleName() string { return "cisco_ios_xr" }

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetYangName() string { return "ipsla-tracks" }

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) SetParent(parent types.Entity) { ipslaTracks.parent = parent }

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetParent() types.Entity { return ipslaTracks.parent }

func (ipslaTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_IpslaTracks) GetParentYangName() string { return "track-type-info" }

// ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks
// track type bfdrtr info
type ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..120.
    InterfaceName interface{}

    // Destination Address. The type is interface{} with range: 0..4294967295.
    DestinationAddress interface{}

    // Rate. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 0..4294967295.
    DebounceCount interface{}
}

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetFilter() yfilter.YFilter { return bfdTracks.YFilter }

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) SetFilter(yf yfilter.YFilter) { bfdTracks.YFilter = yf }

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "rate" { return "Rate" }
    if yname == "debounce-count" { return "DebounceCount" }
    return ""
}

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetSegmentPath() string {
    return "bfd-tracks"
}

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bfdTracks.InterfaceName
    leafs["destination-address"] = bfdTracks.DestinationAddress
    leafs["rate"] = bfdTracks.Rate
    leafs["debounce-count"] = bfdTracks.DebounceCount
    return leafs
}

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetBundleName() string { return "cisco_ios_xr" }

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetYangName() string { return "bfd-tracks" }

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) SetParent(parent types.Entity) { bfdTracks.parent = parent }

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetParent() types.Entity { return bfdTracks.parent }

func (bfdTracks *ObjectTracking_TrackTypeInterfaceBrief_TrackInfoBrief_TrackTypeInfo_BfdTracks) GetParentYangName() string { return "track-type-info" }

