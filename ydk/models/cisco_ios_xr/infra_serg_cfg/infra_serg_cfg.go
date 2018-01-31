// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-serg package configuration.
// 
// This module contains definitions
// for the following management objects:
//   session-redundancy: Session Redundancy configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_serg_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_serg_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-serg-cfg session-redundancy}", reflect.TypeOf(SessionRedundancy{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-serg-cfg:session-redundancy", reflect.TypeOf(SessionRedundancy{}))
}

// SergAddrFamily represents Serg addr family
type SergAddrFamily string

const (
    // IPv4
    SergAddrFamily_ipv4 SergAddrFamily = "ipv4"

    // IPv6
    SergAddrFamily_ipv6 SergAddrFamily = "ipv6"
)

// SessionRedundancyGroupRole represents Session redundancy group role
type SessionRedundancyGroupRole string

const (
    // Master Role
    SessionRedundancyGroupRole_master SessionRedundancyGroupRole = "master"

    // Slave Role
    SessionRedundancyGroupRole_slave SessionRedundancyGroupRole = "slave"
)

// SessionRedundancy
// Session Redundancy configuration
type SessionRedundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable. The type is interface{}.
    RedundancyDisable interface{}

    // Enable Session Redundancy configuration. Deletion of this object also
    // causes deletion of all associated objects under SessionRedundancy. The type
    // is interface{}.
    Enable interface{}

    // Source Interface for Redundancy Peer Communication. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Set preferred role. The type is SessionRedundancyGroupRole.
    PreferredRole interface{}

    // Set hold time (in Minutes). The type is interface{} with range: 1..65535.
    // Units are minute.
    HoldTimer interface{}

    // Table of Group.
    Groups SessionRedundancy_Groups

    // None.
    RevertiveTimer SessionRedundancy_RevertiveTimer
}

func (sessionRedundancy *SessionRedundancy) GetFilter() yfilter.YFilter { return sessionRedundancy.YFilter }

func (sessionRedundancy *SessionRedundancy) SetFilter(yf yfilter.YFilter) { sessionRedundancy.YFilter = yf }

func (sessionRedundancy *SessionRedundancy) GetGoName(yname string) string {
    if yname == "redundancy-disable" { return "RedundancyDisable" }
    if yname == "enable" { return "Enable" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "preferred-role" { return "PreferredRole" }
    if yname == "hold-timer" { return "HoldTimer" }
    if yname == "groups" { return "Groups" }
    if yname == "revertive-timer" { return "RevertiveTimer" }
    return ""
}

func (sessionRedundancy *SessionRedundancy) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-serg-cfg:session-redundancy"
}

func (sessionRedundancy *SessionRedundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &sessionRedundancy.Groups
    }
    if childYangName == "revertive-timer" {
        return &sessionRedundancy.RevertiveTimer
    }
    return nil
}

func (sessionRedundancy *SessionRedundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &sessionRedundancy.Groups
    children["revertive-timer"] = &sessionRedundancy.RevertiveTimer
    return children
}

func (sessionRedundancy *SessionRedundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["redundancy-disable"] = sessionRedundancy.RedundancyDisable
    leafs["enable"] = sessionRedundancy.Enable
    leafs["source-interface"] = sessionRedundancy.SourceInterface
    leafs["preferred-role"] = sessionRedundancy.PreferredRole
    leafs["hold-timer"] = sessionRedundancy.HoldTimer
    return leafs
}

func (sessionRedundancy *SessionRedundancy) GetBundleName() string { return "cisco_ios_xr" }

func (sessionRedundancy *SessionRedundancy) GetYangName() string { return "session-redundancy" }

func (sessionRedundancy *SessionRedundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionRedundancy *SessionRedundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionRedundancy *SessionRedundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionRedundancy *SessionRedundancy) SetParent(parent types.Entity) { sessionRedundancy.parent = parent }

func (sessionRedundancy *SessionRedundancy) GetParent() types.Entity { return sessionRedundancy.parent }

func (sessionRedundancy *SessionRedundancy) GetParentYangName() string { return "Cisco-IOS-XR-infra-serg-cfg" }

// SessionRedundancy_Groups
// Table of Group
type SessionRedundancy_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redundancy Group configuration. The type is slice of
    // SessionRedundancy_Groups_Group.
    Group []SessionRedundancy_Groups_Group
}

func (groups *SessionRedundancy_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *SessionRedundancy_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *SessionRedundancy_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *SessionRedundancy_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *SessionRedundancy_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancy_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *SessionRedundancy_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *SessionRedundancy_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *SessionRedundancy_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *SessionRedundancy_Groups) GetYangName() string { return "groups" }

func (groups *SessionRedundancy_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *SessionRedundancy_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *SessionRedundancy_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *SessionRedundancy_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *SessionRedundancy_Groups) GetParent() types.Entity { return groups.parent }

func (groups *SessionRedundancy_Groups) GetParentYangName() string { return "session-redundancy" }

// SessionRedundancy_Groups_Group
// Redundancy Group configuration
type SessionRedundancy_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..500.
    GroupId interface{}

    // Core Tracking Object for this Group. The type is string.
    CoreTrackingObject interface{}

    // Disable Tracking Object for this Group. The type is interface{}.
    DisableTrackingObject interface{}

    // Disable. The type is interface{}.
    RedundancyDisable interface{}

    // Enable Redundancy Group configuration. Deletion of this object also causes
    // deletion of all associated objects under Group. The type is interface{}.
    Enable interface{}

    // Description for this Group. The type is string.
    Description interface{}

    // Access Tracking Object for this Group. The type is string.
    AccessTrackingObject interface{}

    // Set preferred role. The type is SessionRedundancyGroupRole.
    PreferredRole interface{}

    // Set hold time (in Minutes). The type is interface{} with range: 1..65535.
    // Units are minute.
    HoldTimer interface{}

    // None.
    Peer SessionRedundancy_Groups_Group_Peer

    // None.
    RevertiveTimer SessionRedundancy_Groups_Group_RevertiveTimer

    // List of Interfaces for this Group.
    InterfaceList SessionRedundancy_Groups_Group_InterfaceList
}

func (group *SessionRedundancy_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *SessionRedundancy_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *SessionRedundancy_Groups_Group) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "core-tracking-object" { return "CoreTrackingObject" }
    if yname == "disable-tracking-object" { return "DisableTrackingObject" }
    if yname == "redundancy-disable" { return "RedundancyDisable" }
    if yname == "enable" { return "Enable" }
    if yname == "description" { return "Description" }
    if yname == "access-tracking-object" { return "AccessTrackingObject" }
    if yname == "preferred-role" { return "PreferredRole" }
    if yname == "hold-timer" { return "HoldTimer" }
    if yname == "peer" { return "Peer" }
    if yname == "revertive-timer" { return "RevertiveTimer" }
    if yname == "interface-list" { return "InterfaceList" }
    return ""
}

func (group *SessionRedundancy_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-id='" + fmt.Sprintf("%v", group.GroupId) + "']"
}

func (group *SessionRedundancy_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer" {
        return &group.Peer
    }
    if childYangName == "revertive-timer" {
        return &group.RevertiveTimer
    }
    if childYangName == "interface-list" {
        return &group.InterfaceList
    }
    return nil
}

func (group *SessionRedundancy_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer"] = &group.Peer
    children["revertive-timer"] = &group.RevertiveTimer
    children["interface-list"] = &group.InterfaceList
    return children
}

func (group *SessionRedundancy_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = group.GroupId
    leafs["core-tracking-object"] = group.CoreTrackingObject
    leafs["disable-tracking-object"] = group.DisableTrackingObject
    leafs["redundancy-disable"] = group.RedundancyDisable
    leafs["enable"] = group.Enable
    leafs["description"] = group.Description
    leafs["access-tracking-object"] = group.AccessTrackingObject
    leafs["preferred-role"] = group.PreferredRole
    leafs["hold-timer"] = group.HoldTimer
    return leafs
}

func (group *SessionRedundancy_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *SessionRedundancy_Groups_Group) GetYangName() string { return "group" }

func (group *SessionRedundancy_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *SessionRedundancy_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *SessionRedundancy_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *SessionRedundancy_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *SessionRedundancy_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *SessionRedundancy_Groups_Group) GetParentYangName() string { return "groups" }

// SessionRedundancy_Groups_Group_Peer
// None
type SessionRedundancy_Groups_Group_Peer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 or IPv6 Address of SERG Peer.
    Ipaddress SessionRedundancy_Groups_Group_Peer_Ipaddress
}

func (peer *SessionRedundancy_Groups_Group_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *SessionRedundancy_Groups_Group_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *SessionRedundancy_Groups_Group_Peer) GetGoName(yname string) string {
    if yname == "ipaddress" { return "Ipaddress" }
    return ""
}

func (peer *SessionRedundancy_Groups_Group_Peer) GetSegmentPath() string {
    return "peer"
}

func (peer *SessionRedundancy_Groups_Group_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipaddress" {
        return &peer.Ipaddress
    }
    return nil
}

func (peer *SessionRedundancy_Groups_Group_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipaddress"] = &peer.Ipaddress
    return children
}

func (peer *SessionRedundancy_Groups_Group_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peer *SessionRedundancy_Groups_Group_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *SessionRedundancy_Groups_Group_Peer) GetYangName() string { return "peer" }

func (peer *SessionRedundancy_Groups_Group_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *SessionRedundancy_Groups_Group_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *SessionRedundancy_Groups_Group_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *SessionRedundancy_Groups_Group_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *SessionRedundancy_Groups_Group_Peer) GetParent() types.Entity { return peer.parent }

func (peer *SessionRedundancy_Groups_Group_Peer) GetParentYangName() string { return "group" }

// SessionRedundancy_Groups_Group_Peer_Ipaddress
// IPv4 or IPv6 Address of SERG Peer
type SessionRedundancy_Groups_Group_Peer_Ipaddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of IPv4/IPv6 address. The type is SergAddrFamily.
    AddressFamily interface{}

    // IPv4/IPv6 address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixString interface{}
}

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetFilter() yfilter.YFilter { return ipaddress.YFilter }

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) SetFilter(yf yfilter.YFilter) { ipaddress.YFilter = yf }

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "prefix-string" { return "PrefixString" }
    return ""
}

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetSegmentPath() string {
    return "ipaddress"
}

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = ipaddress.AddressFamily
    leafs["prefix-string"] = ipaddress.PrefixString
    return leafs
}

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetBundleName() string { return "cisco_ios_xr" }

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetYangName() string { return "ipaddress" }

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) SetParent(parent types.Entity) { ipaddress.parent = parent }

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetParent() types.Entity { return ipaddress.parent }

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetParentYangName() string { return "peer" }

// SessionRedundancy_Groups_Group_RevertiveTimer
// None
type SessionRedundancy_Groups_Group_RevertiveTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Value of MAX Revertive Timer. The type is interface{} with range: 1..65535.
    MaxValue interface{}

    // Value of revertive time in minutes. The type is interface{} with range:
    // 1..65535. Units are minute.
    Value interface{}
}

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetFilter() yfilter.YFilter { return revertiveTimer.YFilter }

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) SetFilter(yf yfilter.YFilter) { revertiveTimer.YFilter = yf }

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetGoName(yname string) string {
    if yname == "max-value" { return "MaxValue" }
    if yname == "value" { return "Value" }
    return ""
}

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetSegmentPath() string {
    return "revertive-timer"
}

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-value"] = revertiveTimer.MaxValue
    leafs["value"] = revertiveTimer.Value
    return leafs
}

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetBundleName() string { return "cisco_ios_xr" }

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetYangName() string { return "revertive-timer" }

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) SetParent(parent types.Entity) { revertiveTimer.parent = parent }

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetParent() types.Entity { return revertiveTimer.parent }

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetParentYangName() string { return "group" }

// SessionRedundancy_Groups_Group_InterfaceList
// List of Interfaces for this Group
type SessionRedundancy_Groups_Group_InterfaceList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable List of Interfaces for this Group. Deletion of this object also
    // causes deletion of all associated objects under InterfaceList . The type is
    // interface{}.
    Enable interface{}

    // Table of InterfaceRange.
    InterfaceRanges SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges

    // Table of Interface.
    Interfaces SessionRedundancy_Groups_Group_InterfaceList_Interfaces
}

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetFilter() yfilter.YFilter { return interfaceList.YFilter }

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) SetFilter(yf yfilter.YFilter) { interfaceList.YFilter = yf }

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "interface-ranges" { return "InterfaceRanges" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetSegmentPath() string {
    return "interface-list"
}

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-ranges" {
        return &interfaceList.InterfaceRanges
    }
    if childYangName == "interfaces" {
        return &interfaceList.Interfaces
    }
    return nil
}

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-ranges"] = &interfaceList.InterfaceRanges
    children["interfaces"] = &interfaceList.Interfaces
    return children
}

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = interfaceList.Enable
    return leafs
}

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetYangName() string { return "interface-list" }

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) SetParent(parent types.Entity) { interfaceList.parent = parent }

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetParent() types.Entity { return interfaceList.parent }

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetParentYangName() string { return "group" }

// SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges
// Table of InterfaceRange
type SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange.
    InterfaceRange []SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
}

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetFilter() yfilter.YFilter { return interfaceRanges.YFilter }

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) SetFilter(yf yfilter.YFilter) { interfaceRanges.YFilter = yf }

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetGoName(yname string) string {
    if yname == "interface-range" { return "InterfaceRange" }
    return ""
}

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetSegmentPath() string {
    return "interface-ranges"
}

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-range" {
        for _, c := range interfaceRanges.InterfaceRange {
            if interfaceRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange{}
        interfaceRanges.InterfaceRange = append(interfaceRanges.InterfaceRange, child)
        return &interfaceRanges.InterfaceRange[len(interfaceRanges.InterfaceRange)-1]
    }
    return nil
}

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceRanges.InterfaceRange {
        children[interfaceRanges.InterfaceRange[i].GetSegmentPath()] = &interfaceRanges.InterfaceRange[i]
    }
    return children
}

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetYangName() string { return "interface-ranges" }

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) SetParent(parent types.Entity) { interfaceRanges.parent = parent }

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetParent() types.Entity { return interfaceRanges.parent }

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetParentYangName() string { return "interface-list" }

// SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
// Interface for this Group
type SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. Sub Interface Start Range. The type is interface{}
    // with range: 0..2147483647.
    SubInterfaceRangeStart interface{}

    // This attribute is a key. Sub Interface End Range. The type is interface{}
    // with range: 0..2147483647.
    SubInterfaceRangeEnd interface{}

    // Interface ID Start Range. The type is interface{} with range: 1..65535.
    InterfaceIdRangeStart interface{}

    // Interface ID End Range. The type is interface{} with range: 1..65535.
    InterfaceIdRangeEnd interface{}
}

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetFilter() yfilter.YFilter { return interfaceRange.YFilter }

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) SetFilter(yf yfilter.YFilter) { interfaceRange.YFilter = yf }

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "sub-interface-range-start" { return "SubInterfaceRangeStart" }
    if yname == "sub-interface-range-end" { return "SubInterfaceRangeEnd" }
    if yname == "interface-id-range-start" { return "InterfaceIdRangeStart" }
    if yname == "interface-id-range-end" { return "InterfaceIdRangeEnd" }
    return ""
}

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetSegmentPath() string {
    return "interface-range" + "[interface-name='" + fmt.Sprintf("%v", interfaceRange.InterfaceName) + "']" + "[sub-interface-range-start='" + fmt.Sprintf("%v", interfaceRange.SubInterfaceRangeStart) + "']" + "[sub-interface-range-end='" + fmt.Sprintf("%v", interfaceRange.SubInterfaceRangeEnd) + "']"
}

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceRange.InterfaceName
    leafs["sub-interface-range-start"] = interfaceRange.SubInterfaceRangeStart
    leafs["sub-interface-range-end"] = interfaceRange.SubInterfaceRangeEnd
    leafs["interface-id-range-start"] = interfaceRange.InterfaceIdRangeStart
    leafs["interface-id-range-end"] = interfaceRange.InterfaceIdRangeEnd
    return leafs
}

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetYangName() string { return "interface-range" }

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) SetParent(parent types.Entity) { interfaceRange.parent = parent }

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetParent() types.Entity { return interfaceRange.parent }

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetParentYangName() string { return "interface-ranges" }

// SessionRedundancy_Groups_Group_InterfaceList_Interfaces
// Table of Interface
type SessionRedundancy_Groups_Group_InterfaceList_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface.
    Interface []SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
}

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetParentYangName() string { return "interface-list" }

// SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
// Interface for this Group
type SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface Id for the interface. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    InterfaceId interface{}
}

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-id" { return "InterfaceId" }
    return ""
}

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-id"] = self.InterfaceId
    return leafs
}

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// SessionRedundancy_RevertiveTimer
// None
type SessionRedundancy_RevertiveTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Value of MAX Revertive Timer. The type is interface{} with range: 1..65535.
    MaxValue interface{}

    // Value of revertive time in minutes. The type is interface{} with range:
    // 1..65535. Units are minute.
    Value interface{}
}

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetFilter() yfilter.YFilter { return revertiveTimer.YFilter }

func (revertiveTimer *SessionRedundancy_RevertiveTimer) SetFilter(yf yfilter.YFilter) { revertiveTimer.YFilter = yf }

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetGoName(yname string) string {
    if yname == "max-value" { return "MaxValue" }
    if yname == "value" { return "Value" }
    return ""
}

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetSegmentPath() string {
    return "revertive-timer"
}

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-value"] = revertiveTimer.MaxValue
    leafs["value"] = revertiveTimer.Value
    return leafs
}

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetBundleName() string { return "cisco_ios_xr" }

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetYangName() string { return "revertive-timer" }

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (revertiveTimer *SessionRedundancy_RevertiveTimer) SetParent(parent types.Entity) { revertiveTimer.parent = parent }

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetParent() types.Entity { return revertiveTimer.parent }

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetParentYangName() string { return "session-redundancy" }

