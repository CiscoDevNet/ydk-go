// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-srg package configuration.
// 
// This module contains definitions
// for the following management objects:
//   subscriber-redundancy: Subscriber Redundancy configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_srg_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_srg_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-srg-cfg subscriber-redundancy}", reflect.TypeOf(SubscriberRedundancy{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy", reflect.TypeOf(SubscriberRedundancy{}))
}

// SubscriberRedundancyGroupSlaveMode represents Subscriber redundancy group slave mode
type SubscriberRedundancyGroupSlaveMode string

const (
    // Warm Mode
    SubscriberRedundancyGroupSlaveMode_warm SubscriberRedundancyGroupSlaveMode = "warm"

    // Hot Mode
    SubscriberRedundancyGroupSlaveMode_hot SubscriberRedundancyGroupSlaveMode = "hot"
)

// SrgAddrFamily represents Srg addr family
type SrgAddrFamily string

const (
    // IPv4
    SrgAddrFamily_ipv4 SrgAddrFamily = "ipv4"

    // IPv6
    SrgAddrFamily_ipv6 SrgAddrFamily = "ipv6"
)

// SubscriberRedundancyGroupRole represents Subscriber redundancy group role
type SubscriberRedundancyGroupRole string

const (
    // Master Role
    SubscriberRedundancyGroupRole_master SubscriberRedundancyGroupRole = "master"

    // Slave Role
    SubscriberRedundancyGroupRole_slave SubscriberRedundancyGroupRole = "slave"
)

// SubscriberRedundancy
// Subscriber Redundancy configuration
type SubscriberRedundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Subscriber Redundancy configuration. Deletion of this object also
    // causes deletion of all associated objects under SubscriberRedundancy. The
    // type is interface{}.
    Enable interface{}

    // Virtual MAC Prefix for Subscriber Redundancy. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacPrefix interface{}

    // Set preferred role. The type is SubscriberRedundancyGroupRole.
    PreferredRole interface{}

    // Source Interface for Redundancy Peer Communication. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Set slave. The type is SubscriberRedundancyGroupSlaveMode.
    SlaveMode interface{}

    // Set hold time (in Minutes). The type is interface{} with range: 1..65535.
    // Units are minute.
    HoldTimer interface{}

    // Disable. The type is interface{}.
    RedundancyDisable interface{}

    // Table of Group.
    Groups SubscriberRedundancy_Groups

    // None.
    RevertiveTimer SubscriberRedundancy_RevertiveTimer
}

func (subscriberRedundancy *SubscriberRedundancy) GetFilter() yfilter.YFilter { return subscriberRedundancy.YFilter }

func (subscriberRedundancy *SubscriberRedundancy) SetFilter(yf yfilter.YFilter) { subscriberRedundancy.YFilter = yf }

func (subscriberRedundancy *SubscriberRedundancy) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "virtual-mac-prefix" { return "VirtualMacPrefix" }
    if yname == "preferred-role" { return "PreferredRole" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "slave-mode" { return "SlaveMode" }
    if yname == "hold-timer" { return "HoldTimer" }
    if yname == "redundancy-disable" { return "RedundancyDisable" }
    if yname == "groups" { return "Groups" }
    if yname == "revertive-timer" { return "RevertiveTimer" }
    return ""
}

func (subscriberRedundancy *SubscriberRedundancy) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy"
}

func (subscriberRedundancy *SubscriberRedundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &subscriberRedundancy.Groups
    }
    if childYangName == "revertive-timer" {
        return &subscriberRedundancy.RevertiveTimer
    }
    return nil
}

func (subscriberRedundancy *SubscriberRedundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &subscriberRedundancy.Groups
    children["revertive-timer"] = &subscriberRedundancy.RevertiveTimer
    return children
}

func (subscriberRedundancy *SubscriberRedundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = subscriberRedundancy.Enable
    leafs["virtual-mac-prefix"] = subscriberRedundancy.VirtualMacPrefix
    leafs["preferred-role"] = subscriberRedundancy.PreferredRole
    leafs["source-interface"] = subscriberRedundancy.SourceInterface
    leafs["slave-mode"] = subscriberRedundancy.SlaveMode
    leafs["hold-timer"] = subscriberRedundancy.HoldTimer
    leafs["redundancy-disable"] = subscriberRedundancy.RedundancyDisable
    return leafs
}

func (subscriberRedundancy *SubscriberRedundancy) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberRedundancy *SubscriberRedundancy) GetYangName() string { return "subscriber-redundancy" }

func (subscriberRedundancy *SubscriberRedundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberRedundancy *SubscriberRedundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberRedundancy *SubscriberRedundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberRedundancy *SubscriberRedundancy) SetParent(parent types.Entity) { subscriberRedundancy.parent = parent }

func (subscriberRedundancy *SubscriberRedundancy) GetParent() types.Entity { return subscriberRedundancy.parent }

func (subscriberRedundancy *SubscriberRedundancy) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-srg-cfg" }

// SubscriberRedundancy_Groups
// Table of Group
type SubscriberRedundancy_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redundancy Group configuration. The type is slice of
    // SubscriberRedundancy_Groups_Group.
    Group []SubscriberRedundancy_Groups_Group
}

func (groups *SubscriberRedundancy_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *SubscriberRedundancy_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *SubscriberRedundancy_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *SubscriberRedundancy_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *SubscriberRedundancy_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancy_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *SubscriberRedundancy_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *SubscriberRedundancy_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *SubscriberRedundancy_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *SubscriberRedundancy_Groups) GetYangName() string { return "groups" }

func (groups *SubscriberRedundancy_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *SubscriberRedundancy_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *SubscriberRedundancy_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *SubscriberRedundancy_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *SubscriberRedundancy_Groups) GetParent() types.Entity { return groups.parent }

func (groups *SubscriberRedundancy_Groups) GetParentYangName() string { return "subscriber-redundancy" }

// SubscriberRedundancy_Groups_Group
// Redundancy Group configuration
type SubscriberRedundancy_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..500.
    GroupId interface{}

    // Disable Tracking Object for this Group. The type is interface{}.
    DisableTrackingObject interface{}

    // Core Tracking Object for this Group. The type is string.
    CoreTrackingObject interface{}

    // Enable Redundancy Group configuration. Deletion of this object also causes
    // deletion of all associated objects under Group. The type is interface{}.
    Enable interface{}

    // Set preferred role. The type is SubscriberRedundancyGroupRole.
    PreferredRole interface{}

    // Description for this Group. The type is string.
    Description interface{}

    // Enter an IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    L2TpSourceIpAddress interface{}

    // Set Slave Mode. The type is SubscriberRedundancyGroupSlaveMode.
    SlaveMode interface{}

    // Set hold time (in Minutes). The type is interface{} with range: 1..65535.
    // Units are minute.
    HoldTimer interface{}

    // Access Tracking Object for this Group. The type is string.
    AccessTrackingObject interface{}

    // Enable fast switchover for this Group. The type is interface{}.
    EnableFastSwitchover interface{}

    // Disable. The type is interface{}.
    RedundancyDisable interface{}

    // List of Interfaces for this Group.
    InterfaceList SubscriberRedundancy_Groups_Group_InterfaceList

    // None.
    Peer SubscriberRedundancy_Groups_Group_Peer

    // None.
    RevertiveTimer SubscriberRedundancy_Groups_Group_RevertiveTimer

    // Virtual MAC Address for this Group.
    VirtualMac SubscriberRedundancy_Groups_Group_VirtualMac

    // None.
    StateControlRoute SubscriberRedundancy_Groups_Group_StateControlRoute
}

func (group *SubscriberRedundancy_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *SubscriberRedundancy_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *SubscriberRedundancy_Groups_Group) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "disable-tracking-object" { return "DisableTrackingObject" }
    if yname == "core-tracking-object" { return "CoreTrackingObject" }
    if yname == "enable" { return "Enable" }
    if yname == "preferred-role" { return "PreferredRole" }
    if yname == "description" { return "Description" }
    if yname == "l2tp-source-ip-address" { return "L2TpSourceIpAddress" }
    if yname == "slave-mode" { return "SlaveMode" }
    if yname == "hold-timer" { return "HoldTimer" }
    if yname == "access-tracking-object" { return "AccessTrackingObject" }
    if yname == "enable-fast-switchover" { return "EnableFastSwitchover" }
    if yname == "redundancy-disable" { return "RedundancyDisable" }
    if yname == "interface-list" { return "InterfaceList" }
    if yname == "peer" { return "Peer" }
    if yname == "revertive-timer" { return "RevertiveTimer" }
    if yname == "virtual-mac" { return "VirtualMac" }
    if yname == "state-control-route" { return "StateControlRoute" }
    return ""
}

func (group *SubscriberRedundancy_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-id='" + fmt.Sprintf("%v", group.GroupId) + "']"
}

func (group *SubscriberRedundancy_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-list" {
        return &group.InterfaceList
    }
    if childYangName == "peer" {
        return &group.Peer
    }
    if childYangName == "revertive-timer" {
        return &group.RevertiveTimer
    }
    if childYangName == "virtual-mac" {
        return &group.VirtualMac
    }
    if childYangName == "state-control-route" {
        return &group.StateControlRoute
    }
    return nil
}

func (group *SubscriberRedundancy_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-list"] = &group.InterfaceList
    children["peer"] = &group.Peer
    children["revertive-timer"] = &group.RevertiveTimer
    children["virtual-mac"] = &group.VirtualMac
    children["state-control-route"] = &group.StateControlRoute
    return children
}

func (group *SubscriberRedundancy_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = group.GroupId
    leafs["disable-tracking-object"] = group.DisableTrackingObject
    leafs["core-tracking-object"] = group.CoreTrackingObject
    leafs["enable"] = group.Enable
    leafs["preferred-role"] = group.PreferredRole
    leafs["description"] = group.Description
    leafs["l2tp-source-ip-address"] = group.L2TpSourceIpAddress
    leafs["slave-mode"] = group.SlaveMode
    leafs["hold-timer"] = group.HoldTimer
    leafs["access-tracking-object"] = group.AccessTrackingObject
    leafs["enable-fast-switchover"] = group.EnableFastSwitchover
    leafs["redundancy-disable"] = group.RedundancyDisable
    return leafs
}

func (group *SubscriberRedundancy_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *SubscriberRedundancy_Groups_Group) GetYangName() string { return "group" }

func (group *SubscriberRedundancy_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *SubscriberRedundancy_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *SubscriberRedundancy_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *SubscriberRedundancy_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *SubscriberRedundancy_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *SubscriberRedundancy_Groups_Group) GetParentYangName() string { return "groups" }

// SubscriberRedundancy_Groups_Group_InterfaceList
// List of Interfaces for this Group
type SubscriberRedundancy_Groups_Group_InterfaceList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable List of Interfaces for this Group. Deletion of this object also
    // causes deletion of all associated objects under InterfaceList . The type is
    // interface{}.
    Enable interface{}

    // Table of Interface.
    Interfaces SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces

    // Table of InterfaceRange.
    InterfaceRanges SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges
}

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetFilter() yfilter.YFilter { return interfaceList.YFilter }

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) SetFilter(yf yfilter.YFilter) { interfaceList.YFilter = yf }

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "interface-ranges" { return "InterfaceRanges" }
    return ""
}

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetSegmentPath() string {
    return "interface-list"
}

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &interfaceList.Interfaces
    }
    if childYangName == "interface-ranges" {
        return &interfaceList.InterfaceRanges
    }
    return nil
}

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &interfaceList.Interfaces
    children["interface-ranges"] = &interfaceList.InterfaceRanges
    return children
}

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = interfaceList.Enable
    return leafs
}

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetYangName() string { return "interface-list" }

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) SetParent(parent types.Entity) { interfaceList.parent = parent }

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetParent() types.Entity { return interfaceList.parent }

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetParentYangName() string { return "group" }

// SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces
// Table of Interface
type SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface.
    Interface []SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
}

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetParentYangName() string { return "interface-list" }

// SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
// Interface for this Group
type SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface Id for the interface. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    InterfaceId interface{}
}

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-id" { return "InterfaceId" }
    return ""
}

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-id"] = self.InterfaceId
    return leafs
}

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges
// Table of InterfaceRange
type SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange.
    InterfaceRange []SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
}

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetFilter() yfilter.YFilter { return interfaceRanges.YFilter }

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) SetFilter(yf yfilter.YFilter) { interfaceRanges.YFilter = yf }

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetGoName(yname string) string {
    if yname == "interface-range" { return "InterfaceRange" }
    return ""
}

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetSegmentPath() string {
    return "interface-ranges"
}

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-range" {
        for _, c := range interfaceRanges.InterfaceRange {
            if interfaceRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange{}
        interfaceRanges.InterfaceRange = append(interfaceRanges.InterfaceRange, child)
        return &interfaceRanges.InterfaceRange[len(interfaceRanges.InterfaceRange)-1]
    }
    return nil
}

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceRanges.InterfaceRange {
        children[interfaceRanges.InterfaceRange[i].GetSegmentPath()] = &interfaceRanges.InterfaceRange[i]
    }
    return children
}

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetYangName() string { return "interface-ranges" }

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) SetParent(parent types.Entity) { interfaceRanges.parent = parent }

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetParent() types.Entity { return interfaceRanges.parent }

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetParentYangName() string { return "interface-list" }

// SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
// Interface for this Group
type SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange struct {
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

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetFilter() yfilter.YFilter { return interfaceRange.YFilter }

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) SetFilter(yf yfilter.YFilter) { interfaceRange.YFilter = yf }

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "sub-interface-range-start" { return "SubInterfaceRangeStart" }
    if yname == "sub-interface-range-end" { return "SubInterfaceRangeEnd" }
    if yname == "interface-id-range-start" { return "InterfaceIdRangeStart" }
    if yname == "interface-id-range-end" { return "InterfaceIdRangeEnd" }
    return ""
}

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetSegmentPath() string {
    return "interface-range" + "[interface-name='" + fmt.Sprintf("%v", interfaceRange.InterfaceName) + "']" + "[sub-interface-range-start='" + fmt.Sprintf("%v", interfaceRange.SubInterfaceRangeStart) + "']" + "[sub-interface-range-end='" + fmt.Sprintf("%v", interfaceRange.SubInterfaceRangeEnd) + "']"
}

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceRange.InterfaceName
    leafs["sub-interface-range-start"] = interfaceRange.SubInterfaceRangeStart
    leafs["sub-interface-range-end"] = interfaceRange.SubInterfaceRangeEnd
    leafs["interface-id-range-start"] = interfaceRange.InterfaceIdRangeStart
    leafs["interface-id-range-end"] = interfaceRange.InterfaceIdRangeEnd
    return leafs
}

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetYangName() string { return "interface-range" }

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) SetParent(parent types.Entity) { interfaceRange.parent = parent }

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetParent() types.Entity { return interfaceRange.parent }

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetParentYangName() string { return "interface-ranges" }

// SubscriberRedundancy_Groups_Group_Peer
// None
type SubscriberRedundancy_Groups_Group_Peer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set Route add disable. The type is interface{}.
    RouteAddDisable interface{}

    // IPv4 or IPv6 Address of SRG Peer.
    Ipaddress SubscriberRedundancy_Groups_Group_Peer_Ipaddress
}

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *SubscriberRedundancy_Groups_Group_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetGoName(yname string) string {
    if yname == "route-add-disable" { return "RouteAddDisable" }
    if yname == "ipaddress" { return "Ipaddress" }
    return ""
}

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetSegmentPath() string {
    return "peer"
}

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipaddress" {
        return &peer.Ipaddress
    }
    return nil
}

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipaddress"] = &peer.Ipaddress
    return children
}

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-add-disable"] = peer.RouteAddDisable
    return leafs
}

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetYangName() string { return "peer" }

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *SubscriberRedundancy_Groups_Group_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetParent() types.Entity { return peer.parent }

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetParentYangName() string { return "group" }

// SubscriberRedundancy_Groups_Group_Peer_Ipaddress
// IPv4 or IPv6 Address of SRG Peer
type SubscriberRedundancy_Groups_Group_Peer_Ipaddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of IPv4/IPv6 address. The type is SrgAddrFamily.
    AddressFamily interface{}

    // IPv4/IPv6 address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixString interface{}
}

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetFilter() yfilter.YFilter { return ipaddress.YFilter }

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) SetFilter(yf yfilter.YFilter) { ipaddress.YFilter = yf }

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "prefix-string" { return "PrefixString" }
    return ""
}

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetSegmentPath() string {
    return "ipaddress"
}

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = ipaddress.AddressFamily
    leafs["prefix-string"] = ipaddress.PrefixString
    return leafs
}

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetBundleName() string { return "cisco_ios_xr" }

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetYangName() string { return "ipaddress" }

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) SetParent(parent types.Entity) { ipaddress.parent = parent }

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetParent() types.Entity { return ipaddress.parent }

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetParentYangName() string { return "peer" }

// SubscriberRedundancy_Groups_Group_RevertiveTimer
// None
type SubscriberRedundancy_Groups_Group_RevertiveTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Value of MAX Revertive Timer. The type is interface{} with range: 1..65535.
    MaxValue interface{}

    // Value of revertive time in minutes. The type is interface{} with range:
    // 1..65535. Units are minute.
    Value interface{}
}

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetFilter() yfilter.YFilter { return revertiveTimer.YFilter }

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) SetFilter(yf yfilter.YFilter) { revertiveTimer.YFilter = yf }

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetGoName(yname string) string {
    if yname == "max-value" { return "MaxValue" }
    if yname == "value" { return "Value" }
    return ""
}

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetSegmentPath() string {
    return "revertive-timer"
}

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-value"] = revertiveTimer.MaxValue
    leafs["value"] = revertiveTimer.Value
    return leafs
}

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetBundleName() string { return "cisco_ios_xr" }

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetYangName() string { return "revertive-timer" }

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) SetParent(parent types.Entity) { revertiveTimer.parent = parent }

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetParent() types.Entity { return revertiveTimer.parent }

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetParentYangName() string { return "group" }

// SubscriberRedundancy_Groups_Group_VirtualMac
// Virtual MAC Address for this Group
type SubscriberRedundancy_Groups_Group_VirtualMac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Virtual MAC Address for this Group. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Address interface{}

    // Disable Virtual MAC Address for this Group. The type is interface{}.
    Disable interface{}
}

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetFilter() yfilter.YFilter { return virtualMac.YFilter }

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) SetFilter(yf yfilter.YFilter) { virtualMac.YFilter = yf }

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "disable" { return "Disable" }
    return ""
}

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetSegmentPath() string {
    return "virtual-mac"
}

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = virtualMac.Address
    leafs["disable"] = virtualMac.Disable
    return leafs
}

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetBundleName() string { return "cisco_ios_xr" }

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetYangName() string { return "virtual-mac" }

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) SetParent(parent types.Entity) { virtualMac.parent = parent }

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetParent() types.Entity { return virtualMac.parent }

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetParentYangName() string { return "group" }

// SubscriberRedundancy_Groups_Group_StateControlRoute
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of IPv4Route.
    Ipv4Routes SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes

    // None.
    Ipv6Route SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route
}

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetFilter() yfilter.YFilter { return stateControlRoute.YFilter }

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) SetFilter(yf yfilter.YFilter) { stateControlRoute.YFilter = yf }

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetGoName(yname string) string {
    if yname == "ipv4-routes" { return "Ipv4Routes" }
    if yname == "ipv6-route" { return "Ipv6Route" }
    return ""
}

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetSegmentPath() string {
    return "state-control-route"
}

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-routes" {
        return &stateControlRoute.Ipv4Routes
    }
    if childYangName == "ipv6-route" {
        return &stateControlRoute.Ipv6Route
    }
    return nil
}

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-routes"] = &stateControlRoute.Ipv4Routes
    children["ipv6-route"] = &stateControlRoute.Ipv6Route
    return children
}

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetBundleName() string { return "cisco_ios_xr" }

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetYangName() string { return "state-control-route" }

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) SetParent(parent types.Entity) { stateControlRoute.parent = parent }

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetParent() types.Entity { return stateControlRoute.parent }

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetParentYangName() string { return "group" }

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes
// Table of IPv4Route
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route.
    Ipv4Route []SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route
}

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetFilter() yfilter.YFilter { return ipv4Routes.YFilter }

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) SetFilter(yf yfilter.YFilter) { ipv4Routes.YFilter = yf }

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetGoName(yname string) string {
    if yname == "ipv4-route" { return "Ipv4Route" }
    return ""
}

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetSegmentPath() string {
    return "ipv4-routes"
}

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-route" {
        for _, c := range ipv4Routes.Ipv4Route {
            if ipv4Routes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route{}
        ipv4Routes.Ipv4Route = append(ipv4Routes.Ipv4Route, child)
        return &ipv4Routes.Ipv4Route[len(ipv4Routes.Ipv4Route)-1]
    }
    return nil
}

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4Routes.Ipv4Route {
        children[ipv4Routes.Ipv4Route[i].GetSegmentPath()] = &ipv4Routes.Ipv4Route[i]
    }
    return children
}

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetYangName() string { return "ipv4-routes" }

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) SetParent(parent types.Entity) { ipv4Routes.parent = parent }

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetParent() types.Entity { return ipv4Routes.parent }

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetParentYangName() string { return "state-control-route" }

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address with prefix-length. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixString interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: -2147483648..2147483647.
    PrefixLength interface{}

    // Data container.
    Ipv4RouteData SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData

    // keys: vrfname. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname.
    Vrfname []SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname
}

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetFilter() yfilter.YFilter { return ipv4Route.YFilter }

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) SetFilter(yf yfilter.YFilter) { ipv4Route.YFilter = yf }

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetGoName(yname string) string {
    if yname == "prefix-string" { return "PrefixString" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "ipv4-route-data" { return "Ipv4RouteData" }
    if yname == "vrfname" { return "Vrfname" }
    return ""
}

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetSegmentPath() string {
    return "ipv4-route" + "[prefix-string='" + fmt.Sprintf("%v", ipv4Route.PrefixString) + "']" + "[prefix-length='" + fmt.Sprintf("%v", ipv4Route.PrefixLength) + "']"
}

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-route-data" {
        return &ipv4Route.Ipv4RouteData
    }
    if childYangName == "vrfname" {
        for _, c := range ipv4Route.Vrfname {
            if ipv4Route.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname{}
        ipv4Route.Vrfname = append(ipv4Route.Vrfname, child)
        return &ipv4Route.Vrfname[len(ipv4Route.Vrfname)-1]
    }
    return nil
}

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-route-data"] = &ipv4Route.Ipv4RouteData
    for i := range ipv4Route.Vrfname {
        children[ipv4Route.Vrfname[i].GetSegmentPath()] = &ipv4Route.Vrfname[i]
    }
    return children
}

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-string"] = ipv4Route.PrefixString
    leafs["prefix-length"] = ipv4Route.PrefixLength
    return leafs
}

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetYangName() string { return "ipv4-route" }

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) SetParent(parent types.Entity) { ipv4Route.parent = parent }

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetParent() types.Entity { return ipv4Route.parent }

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetParentYangName() string { return "ipv4-routes" }

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData
// Data container.
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tag value. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    Tagvalue interface{}
}

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetFilter() yfilter.YFilter { return ipv4RouteData.YFilter }

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) SetFilter(yf yfilter.YFilter) { ipv4RouteData.YFilter = yf }

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetGoName(yname string) string {
    if yname == "tagvalue" { return "Tagvalue" }
    return ""
}

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetSegmentPath() string {
    return "ipv4-route-data"
}

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tagvalue"] = ipv4RouteData.Tagvalue
    return leafs
}

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetYangName() string { return "ipv4-route-data" }

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) SetParent(parent types.Entity) { ipv4RouteData.parent = parent }

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetParent() types.Entity { return ipv4RouteData.parent }

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetParentYangName() string { return "ipv4-route" }

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname
// keys: vrfname
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Vrfname interface{}

    // Tag value. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    Tagvalue interface{}
}

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetFilter() yfilter.YFilter { return vrfname.YFilter }

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) SetFilter(yf yfilter.YFilter) { vrfname.YFilter = yf }

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetGoName(yname string) string {
    if yname == "vrfname" { return "Vrfname" }
    if yname == "tagvalue" { return "Tagvalue" }
    return ""
}

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetSegmentPath() string {
    return "vrfname" + "[vrfname='" + fmt.Sprintf("%v", vrfname.Vrfname) + "']"
}

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrfname"] = vrfname.Vrfname
    leafs["tagvalue"] = vrfname.Tagvalue
    return leafs
}

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetBundleName() string { return "cisco_ios_xr" }

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetYangName() string { return "vrfname" }

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) SetParent(parent types.Entity) { vrfname.parent = parent }

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetParent() types.Entity { return vrfname.parent }

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetParentYangName() string { return "ipv4-route" }

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of IPv6NARoute.
    Ipv6NaRoutes SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes

    // Table of IPv6PDRoute.
    Ipv6PdRoutes SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes
}

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetFilter() yfilter.YFilter { return ipv6Route.YFilter }

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) SetFilter(yf yfilter.YFilter) { ipv6Route.YFilter = yf }

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetGoName(yname string) string {
    if yname == "ipv6na-routes" { return "Ipv6NaRoutes" }
    if yname == "ipv6pd-routes" { return "Ipv6PdRoutes" }
    return ""
}

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetSegmentPath() string {
    return "ipv6-route"
}

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6na-routes" {
        return &ipv6Route.Ipv6NaRoutes
    }
    if childYangName == "ipv6pd-routes" {
        return &ipv6Route.Ipv6PdRoutes
    }
    return nil
}

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6na-routes"] = &ipv6Route.Ipv6NaRoutes
    children["ipv6pd-routes"] = &ipv6Route.Ipv6PdRoutes
    return children
}

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetYangName() string { return "ipv6-route" }

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) SetParent(parent types.Entity) { ipv6Route.parent = parent }

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetParent() types.Entity { return ipv6Route.parent }

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetParentYangName() string { return "state-control-route" }

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes
// Table of IPv6NARoute
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute.
    Ipv6NaRoute []SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute
}

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetFilter() yfilter.YFilter { return ipv6NaRoutes.YFilter }

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) SetFilter(yf yfilter.YFilter) { ipv6NaRoutes.YFilter = yf }

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetGoName(yname string) string {
    if yname == "ipv6na-route" { return "Ipv6NaRoute" }
    return ""
}

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetSegmentPath() string {
    return "ipv6na-routes"
}

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6na-route" {
        for _, c := range ipv6NaRoutes.Ipv6NaRoute {
            if ipv6NaRoutes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute{}
        ipv6NaRoutes.Ipv6NaRoute = append(ipv6NaRoutes.Ipv6NaRoute, child)
        return &ipv6NaRoutes.Ipv6NaRoute[len(ipv6NaRoutes.Ipv6NaRoute)-1]
    }
    return nil
}

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6NaRoutes.Ipv6NaRoute {
        children[ipv6NaRoutes.Ipv6NaRoute[i].GetSegmentPath()] = &ipv6NaRoutes.Ipv6NaRoute[i]
    }
    return children
}

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetYangName() string { return "ipv6na-routes" }

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) SetParent(parent types.Entity) { ipv6NaRoutes.parent = parent }

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetParent() types.Entity { return ipv6NaRoutes.parent }

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetParentYangName() string { return "ipv6-route" }

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Vrfname interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: -2147483648..2147483647.
    PrefixLength interface{}

    // This attribute is a key. IPv6 address with prefix-length. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixString interface{}

    // Tag value. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    Tagvalue interface{}
}

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetFilter() yfilter.YFilter { return ipv6NaRoute.YFilter }

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) SetFilter(yf yfilter.YFilter) { ipv6NaRoute.YFilter = yf }

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetGoName(yname string) string {
    if yname == "vrfname" { return "Vrfname" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix-string" { return "PrefixString" }
    if yname == "tagvalue" { return "Tagvalue" }
    return ""
}

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetSegmentPath() string {
    return "ipv6na-route" + "[vrfname='" + fmt.Sprintf("%v", ipv6NaRoute.Vrfname) + "']" + "[prefix-length='" + fmt.Sprintf("%v", ipv6NaRoute.PrefixLength) + "']" + "[prefix-string='" + fmt.Sprintf("%v", ipv6NaRoute.PrefixString) + "']"
}

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrfname"] = ipv6NaRoute.Vrfname
    leafs["prefix-length"] = ipv6NaRoute.PrefixLength
    leafs["prefix-string"] = ipv6NaRoute.PrefixString
    leafs["tagvalue"] = ipv6NaRoute.Tagvalue
    return leafs
}

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetYangName() string { return "ipv6na-route" }

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) SetParent(parent types.Entity) { ipv6NaRoute.parent = parent }

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetParent() types.Entity { return ipv6NaRoute.parent }

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetParentYangName() string { return "ipv6na-routes" }

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes
// Table of IPv6PDRoute
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute.
    Ipv6PdRoute []SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute
}

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetFilter() yfilter.YFilter { return ipv6PdRoutes.YFilter }

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) SetFilter(yf yfilter.YFilter) { ipv6PdRoutes.YFilter = yf }

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetGoName(yname string) string {
    if yname == "ipv6pd-route" { return "Ipv6PdRoute" }
    return ""
}

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetSegmentPath() string {
    return "ipv6pd-routes"
}

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6pd-route" {
        for _, c := range ipv6PdRoutes.Ipv6PdRoute {
            if ipv6PdRoutes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute{}
        ipv6PdRoutes.Ipv6PdRoute = append(ipv6PdRoutes.Ipv6PdRoute, child)
        return &ipv6PdRoutes.Ipv6PdRoute[len(ipv6PdRoutes.Ipv6PdRoute)-1]
    }
    return nil
}

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6PdRoutes.Ipv6PdRoute {
        children[ipv6PdRoutes.Ipv6PdRoute[i].GetSegmentPath()] = &ipv6PdRoutes.Ipv6PdRoute[i]
    }
    return children
}

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetYangName() string { return "ipv6pd-routes" }

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) SetParent(parent types.Entity) { ipv6PdRoutes.parent = parent }

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetParent() types.Entity { return ipv6PdRoutes.parent }

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetParentYangName() string { return "ipv6-route" }

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Vrfname interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: -2147483648..2147483647.
    PrefixLength interface{}

    // This attribute is a key. IPv6 address with prefix-length. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixString interface{}

    // Tag value. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    Tagvalue interface{}
}

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetFilter() yfilter.YFilter { return ipv6PdRoute.YFilter }

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) SetFilter(yf yfilter.YFilter) { ipv6PdRoute.YFilter = yf }

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetGoName(yname string) string {
    if yname == "vrfname" { return "Vrfname" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix-string" { return "PrefixString" }
    if yname == "tagvalue" { return "Tagvalue" }
    return ""
}

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetSegmentPath() string {
    return "ipv6pd-route" + "[vrfname='" + fmt.Sprintf("%v", ipv6PdRoute.Vrfname) + "']" + "[prefix-length='" + fmt.Sprintf("%v", ipv6PdRoute.PrefixLength) + "']" + "[prefix-string='" + fmt.Sprintf("%v", ipv6PdRoute.PrefixString) + "']"
}

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrfname"] = ipv6PdRoute.Vrfname
    leafs["prefix-length"] = ipv6PdRoute.PrefixLength
    leafs["prefix-string"] = ipv6PdRoute.PrefixString
    leafs["tagvalue"] = ipv6PdRoute.Tagvalue
    return leafs
}

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetYangName() string { return "ipv6pd-route" }

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) SetParent(parent types.Entity) { ipv6PdRoute.parent = parent }

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetParent() types.Entity { return ipv6PdRoute.parent }

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetParentYangName() string { return "ipv6pd-routes" }

// SubscriberRedundancy_RevertiveTimer
// None
type SubscriberRedundancy_RevertiveTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Value of MAX Revertive Timer. The type is interface{} with range: 1..65535.
    MaxValue interface{}

    // Value of revertive time in minutes. The type is interface{} with range:
    // 1..65535. Units are minute.
    Value interface{}
}

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetFilter() yfilter.YFilter { return revertiveTimer.YFilter }

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) SetFilter(yf yfilter.YFilter) { revertiveTimer.YFilter = yf }

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetGoName(yname string) string {
    if yname == "max-value" { return "MaxValue" }
    if yname == "value" { return "Value" }
    return ""
}

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetSegmentPath() string {
    return "revertive-timer"
}

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-value"] = revertiveTimer.MaxValue
    leafs["value"] = revertiveTimer.Value
    return leafs
}

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetBundleName() string { return "cisco_ios_xr" }

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetYangName() string { return "revertive-timer" }

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) SetParent(parent types.Entity) { revertiveTimer.parent = parent }

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetParent() types.Entity { return revertiveTimer.parent }

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetParentYangName() string { return "subscriber-redundancy" }

