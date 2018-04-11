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

// SubscriberRedundancyGroupRole represents Subscriber redundancy group role
type SubscriberRedundancyGroupRole string

const (
    // Master Role
    SubscriberRedundancyGroupRole_master SubscriberRedundancyGroupRole = "master"

    // Slave Role
    SubscriberRedundancyGroupRole_slave SubscriberRedundancyGroupRole = "slave"
)

// SrgAddrFamily represents Srg addr family
type SrgAddrFamily string

const (
    // IPv4
    SrgAddrFamily_ipv4 SrgAddrFamily = "ipv4"

    // IPv6
    SrgAddrFamily_ipv6 SrgAddrFamily = "ipv6"
)

// SubscriberRedundancy
// Subscriber Redundancy configuration
type SubscriberRedundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Subscriber Redundancy configuration. Deletion of this object also
    // causes deletion of all associated objects under SubscriberRedundancy. The
    // type is interface{}.
    Enable interface{}

    // Virtual MAC Prefix for Subscriber Redundancy. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    VirtualMacPrefix interface{}

    // Set preferred role. The type is SubscriberRedundancyGroupRole.
    PreferredRole interface{}

    // Source Interface for Redundancy Peer Communication. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
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

func (subscriberRedundancy *SubscriberRedundancy) GetEntityData() *types.CommonEntityData {
    subscriberRedundancy.EntityData.YFilter = subscriberRedundancy.YFilter
    subscriberRedundancy.EntityData.YangName = "subscriber-redundancy"
    subscriberRedundancy.EntityData.BundleName = "cisco_ios_xr"
    subscriberRedundancy.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-srg-cfg"
    subscriberRedundancy.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy"
    subscriberRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberRedundancy.EntityData.Children = make(map[string]types.YChild)
    subscriberRedundancy.EntityData.Children["groups"] = types.YChild{"Groups", &subscriberRedundancy.Groups}
    subscriberRedundancy.EntityData.Children["revertive-timer"] = types.YChild{"RevertiveTimer", &subscriberRedundancy.RevertiveTimer}
    subscriberRedundancy.EntityData.Leafs = make(map[string]types.YLeaf)
    subscriberRedundancy.EntityData.Leafs["enable"] = types.YLeaf{"Enable", subscriberRedundancy.Enable}
    subscriberRedundancy.EntityData.Leafs["virtual-mac-prefix"] = types.YLeaf{"VirtualMacPrefix", subscriberRedundancy.VirtualMacPrefix}
    subscriberRedundancy.EntityData.Leafs["preferred-role"] = types.YLeaf{"PreferredRole", subscriberRedundancy.PreferredRole}
    subscriberRedundancy.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", subscriberRedundancy.SourceInterface}
    subscriberRedundancy.EntityData.Leafs["slave-mode"] = types.YLeaf{"SlaveMode", subscriberRedundancy.SlaveMode}
    subscriberRedundancy.EntityData.Leafs["hold-timer"] = types.YLeaf{"HoldTimer", subscriberRedundancy.HoldTimer}
    subscriberRedundancy.EntityData.Leafs["redundancy-disable"] = types.YLeaf{"RedundancyDisable", subscriberRedundancy.RedundancyDisable}
    return &(subscriberRedundancy.EntityData)
}

// SubscriberRedundancy_Groups
// Table of Group
type SubscriberRedundancy_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redundancy Group configuration. The type is slice of
    // SubscriberRedundancy_Groups_Group.
    Group []SubscriberRedundancy_Groups_Group
}

func (groups *SubscriberRedundancy_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "subscriber-redundancy"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groups.EntityData)
}

// SubscriberRedundancy_Groups_Group
// Redundancy Group configuration
type SubscriberRedundancy_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..4000.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (group *SubscriberRedundancy_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[group-id='" + fmt.Sprintf("%v", group.GroupId) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Children["interface-list"] = types.YChild{"InterfaceList", &group.InterfaceList}
    group.EntityData.Children["peer"] = types.YChild{"Peer", &group.Peer}
    group.EntityData.Children["revertive-timer"] = types.YChild{"RevertiveTimer", &group.RevertiveTimer}
    group.EntityData.Children["virtual-mac"] = types.YChild{"VirtualMac", &group.VirtualMac}
    group.EntityData.Children["state-control-route"] = types.YChild{"StateControlRoute", &group.StateControlRoute}
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", group.GroupId}
    group.EntityData.Leafs["disable-tracking-object"] = types.YLeaf{"DisableTrackingObject", group.DisableTrackingObject}
    group.EntityData.Leafs["core-tracking-object"] = types.YLeaf{"CoreTrackingObject", group.CoreTrackingObject}
    group.EntityData.Leafs["enable"] = types.YLeaf{"Enable", group.Enable}
    group.EntityData.Leafs["preferred-role"] = types.YLeaf{"PreferredRole", group.PreferredRole}
    group.EntityData.Leafs["description"] = types.YLeaf{"Description", group.Description}
    group.EntityData.Leafs["l2tp-source-ip-address"] = types.YLeaf{"L2TpSourceIpAddress", group.L2TpSourceIpAddress}
    group.EntityData.Leafs["slave-mode"] = types.YLeaf{"SlaveMode", group.SlaveMode}
    group.EntityData.Leafs["hold-timer"] = types.YLeaf{"HoldTimer", group.HoldTimer}
    group.EntityData.Leafs["access-tracking-object"] = types.YLeaf{"AccessTrackingObject", group.AccessTrackingObject}
    group.EntityData.Leafs["enable-fast-switchover"] = types.YLeaf{"EnableFastSwitchover", group.EnableFastSwitchover}
    group.EntityData.Leafs["redundancy-disable"] = types.YLeaf{"RedundancyDisable", group.RedundancyDisable}
    return &(group.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList
// List of Interfaces for this Group
type SubscriberRedundancy_Groups_Group_InterfaceList struct {
    EntityData types.CommonEntityData
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

func (interfaceList *SubscriberRedundancy_Groups_Group_InterfaceList) GetEntityData() *types.CommonEntityData {
    interfaceList.EntityData.YFilter = interfaceList.YFilter
    interfaceList.EntityData.YangName = "interface-list"
    interfaceList.EntityData.BundleName = "cisco_ios_xr"
    interfaceList.EntityData.ParentYangName = "group"
    interfaceList.EntityData.SegmentPath = "interface-list"
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = make(map[string]types.YChild)
    interfaceList.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &interfaceList.Interfaces}
    interfaceList.EntityData.Children["interface-ranges"] = types.YChild{"InterfaceRanges", &interfaceList.InterfaceRanges}
    interfaceList.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceList.EntityData.Leafs["enable"] = types.YLeaf{"Enable", interfaceList.Enable}
    return &(interfaceList.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces
// Table of Interface
type SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface_.
    Interface_ []SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
}

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-list"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
// Interface for this Group
type SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface Id for the interface. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    InterfaceId interface{}
}

func (self *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-id"] = types.YLeaf{"InterfaceId", self.InterfaceId}
    return &(self.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges
// Table of InterfaceRange
type SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange.
    InterfaceRange []SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
}

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetEntityData() *types.CommonEntityData {
    interfaceRanges.EntityData.YFilter = interfaceRanges.YFilter
    interfaceRanges.EntityData.YangName = "interface-ranges"
    interfaceRanges.EntityData.BundleName = "cisco_ios_xr"
    interfaceRanges.EntityData.ParentYangName = "interface-list"
    interfaceRanges.EntityData.SegmentPath = "interface-ranges"
    interfaceRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceRanges.EntityData.Children = make(map[string]types.YChild)
    interfaceRanges.EntityData.Children["interface-range"] = types.YChild{"InterfaceRange", nil}
    for i := range interfaceRanges.InterfaceRange {
        interfaceRanges.EntityData.Children[types.GetSegmentPath(&interfaceRanges.InterfaceRange[i])] = types.YChild{"InterfaceRange", &interfaceRanges.InterfaceRange[i]}
    }
    interfaceRanges.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceRanges.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
// Interface for this Group
type SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (interfaceRange *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetEntityData() *types.CommonEntityData {
    interfaceRange.EntityData.YFilter = interfaceRange.YFilter
    interfaceRange.EntityData.YangName = "interface-range"
    interfaceRange.EntityData.BundleName = "cisco_ios_xr"
    interfaceRange.EntityData.ParentYangName = "interface-ranges"
    interfaceRange.EntityData.SegmentPath = "interface-range" + "[interface-name='" + fmt.Sprintf("%v", interfaceRange.InterfaceName) + "']" + "[sub-interface-range-start='" + fmt.Sprintf("%v", interfaceRange.SubInterfaceRangeStart) + "']" + "[sub-interface-range-end='" + fmt.Sprintf("%v", interfaceRange.SubInterfaceRangeEnd) + "']"
    interfaceRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceRange.EntityData.Children = make(map[string]types.YChild)
    interfaceRange.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceRange.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceRange.InterfaceName}
    interfaceRange.EntityData.Leafs["sub-interface-range-start"] = types.YLeaf{"SubInterfaceRangeStart", interfaceRange.SubInterfaceRangeStart}
    interfaceRange.EntityData.Leafs["sub-interface-range-end"] = types.YLeaf{"SubInterfaceRangeEnd", interfaceRange.SubInterfaceRangeEnd}
    interfaceRange.EntityData.Leafs["interface-id-range-start"] = types.YLeaf{"InterfaceIdRangeStart", interfaceRange.InterfaceIdRangeStart}
    interfaceRange.EntityData.Leafs["interface-id-range-end"] = types.YLeaf{"InterfaceIdRangeEnd", interfaceRange.InterfaceIdRangeEnd}
    return &(interfaceRange.EntityData)
}

// SubscriberRedundancy_Groups_Group_Peer
// None
type SubscriberRedundancy_Groups_Group_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set Route add disable. The type is interface{}.
    RouteAddDisable interface{}

    // IPv4 or IPv6 Address of SRG Peer.
    Ipaddress SubscriberRedundancy_Groups_Group_Peer_Ipaddress
}

func (peer *SubscriberRedundancy_Groups_Group_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "group"
    peer.EntityData.SegmentPath = "peer"
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = make(map[string]types.YChild)
    peer.EntityData.Children["ipaddress"] = types.YChild{"Ipaddress", &peer.Ipaddress}
    peer.EntityData.Leafs = make(map[string]types.YLeaf)
    peer.EntityData.Leafs["route-add-disable"] = types.YLeaf{"RouteAddDisable", peer.RouteAddDisable}
    return &(peer.EntityData)
}

// SubscriberRedundancy_Groups_Group_Peer_Ipaddress
// IPv4 or IPv6 Address of SRG Peer
type SubscriberRedundancy_Groups_Group_Peer_Ipaddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of IPv4/IPv6 address. The type is SrgAddrFamily.
    AddressFamily interface{}

    // IPv4/IPv6 address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PrefixString interface{}
}

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetEntityData() *types.CommonEntityData {
    ipaddress.EntityData.YFilter = ipaddress.YFilter
    ipaddress.EntityData.YangName = "ipaddress"
    ipaddress.EntityData.BundleName = "cisco_ios_xr"
    ipaddress.EntityData.ParentYangName = "peer"
    ipaddress.EntityData.SegmentPath = "ipaddress"
    ipaddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipaddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipaddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipaddress.EntityData.Children = make(map[string]types.YChild)
    ipaddress.EntityData.Leafs = make(map[string]types.YLeaf)
    ipaddress.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", ipaddress.AddressFamily}
    ipaddress.EntityData.Leafs["prefix-string"] = types.YLeaf{"PrefixString", ipaddress.PrefixString}
    return &(ipaddress.EntityData)
}

// SubscriberRedundancy_Groups_Group_RevertiveTimer
// None
type SubscriberRedundancy_Groups_Group_RevertiveTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Value of MAX Revertive Timer. The type is interface{} with range: 1..65535.
    MaxValue interface{}

    // Value of revertive time in minutes. The type is interface{} with range:
    // 1..65535. Units are minute.
    Value interface{}
}

func (revertiveTimer *SubscriberRedundancy_Groups_Group_RevertiveTimer) GetEntityData() *types.CommonEntityData {
    revertiveTimer.EntityData.YFilter = revertiveTimer.YFilter
    revertiveTimer.EntityData.YangName = "revertive-timer"
    revertiveTimer.EntityData.BundleName = "cisco_ios_xr"
    revertiveTimer.EntityData.ParentYangName = "group"
    revertiveTimer.EntityData.SegmentPath = "revertive-timer"
    revertiveTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    revertiveTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    revertiveTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    revertiveTimer.EntityData.Children = make(map[string]types.YChild)
    revertiveTimer.EntityData.Leafs = make(map[string]types.YLeaf)
    revertiveTimer.EntityData.Leafs["max-value"] = types.YLeaf{"MaxValue", revertiveTimer.MaxValue}
    revertiveTimer.EntityData.Leafs["value"] = types.YLeaf{"Value", revertiveTimer.Value}
    return &(revertiveTimer.EntityData)
}

// SubscriberRedundancy_Groups_Group_VirtualMac
// Virtual MAC Address for this Group
type SubscriberRedundancy_Groups_Group_VirtualMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual MAC Address for this Group. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Address interface{}

    // Disable Virtual MAC Address for this Group. The type is interface{}.
    Disable interface{}
}

func (virtualMac *SubscriberRedundancy_Groups_Group_VirtualMac) GetEntityData() *types.CommonEntityData {
    virtualMac.EntityData.YFilter = virtualMac.YFilter
    virtualMac.EntityData.YangName = "virtual-mac"
    virtualMac.EntityData.BundleName = "cisco_ios_xr"
    virtualMac.EntityData.ParentYangName = "group"
    virtualMac.EntityData.SegmentPath = "virtual-mac"
    virtualMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualMac.EntityData.Children = make(map[string]types.YChild)
    virtualMac.EntityData.Leafs = make(map[string]types.YLeaf)
    virtualMac.EntityData.Leafs["address"] = types.YLeaf{"Address", virtualMac.Address}
    virtualMac.EntityData.Leafs["disable"] = types.YLeaf{"Disable", virtualMac.Disable}
    return &(virtualMac.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv4Route.
    Ipv4Routes SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes

    // None.
    Ipv6Route SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route
}

func (stateControlRoute *SubscriberRedundancy_Groups_Group_StateControlRoute) GetEntityData() *types.CommonEntityData {
    stateControlRoute.EntityData.YFilter = stateControlRoute.YFilter
    stateControlRoute.EntityData.YangName = "state-control-route"
    stateControlRoute.EntityData.BundleName = "cisco_ios_xr"
    stateControlRoute.EntityData.ParentYangName = "group"
    stateControlRoute.EntityData.SegmentPath = "state-control-route"
    stateControlRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateControlRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateControlRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateControlRoute.EntityData.Children = make(map[string]types.YChild)
    stateControlRoute.EntityData.Children["ipv4-routes"] = types.YChild{"Ipv4Routes", &stateControlRoute.Ipv4Routes}
    stateControlRoute.EntityData.Children["ipv6-route"] = types.YChild{"Ipv6Route", &stateControlRoute.Ipv6Route}
    stateControlRoute.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateControlRoute.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes
// Table of IPv4Route
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route.
    Ipv4Route []SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route
}

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetEntityData() *types.CommonEntityData {
    ipv4Routes.EntityData.YFilter = ipv4Routes.YFilter
    ipv4Routes.EntityData.YangName = "ipv4-routes"
    ipv4Routes.EntityData.BundleName = "cisco_ios_xr"
    ipv4Routes.EntityData.ParentYangName = "state-control-route"
    ipv4Routes.EntityData.SegmentPath = "ipv4-routes"
    ipv4Routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Routes.EntityData.Children = make(map[string]types.YChild)
    ipv4Routes.EntityData.Children["ipv4-route"] = types.YChild{"Ipv4Route", nil}
    for i := range ipv4Routes.Ipv4Route {
        ipv4Routes.EntityData.Children[types.GetSegmentPath(&ipv4Routes.Ipv4Route[i])] = types.YChild{"Ipv4Route", &ipv4Routes.Ipv4Route[i]}
    }
    ipv4Routes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4Routes.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address with prefix-length. The type is one
    // of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetEntityData() *types.CommonEntityData {
    ipv4Route.EntityData.YFilter = ipv4Route.YFilter
    ipv4Route.EntityData.YangName = "ipv4-route"
    ipv4Route.EntityData.BundleName = "cisco_ios_xr"
    ipv4Route.EntityData.ParentYangName = "ipv4-routes"
    ipv4Route.EntityData.SegmentPath = "ipv4-route" + "[prefix-string='" + fmt.Sprintf("%v", ipv4Route.PrefixString) + "']" + "[prefix-length='" + fmt.Sprintf("%v", ipv4Route.PrefixLength) + "']"
    ipv4Route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Route.EntityData.Children = make(map[string]types.YChild)
    ipv4Route.EntityData.Children["ipv4-route-data"] = types.YChild{"Ipv4RouteData", &ipv4Route.Ipv4RouteData}
    ipv4Route.EntityData.Children["vrfname"] = types.YChild{"Vrfname", nil}
    for i := range ipv4Route.Vrfname {
        ipv4Route.EntityData.Children[types.GetSegmentPath(&ipv4Route.Vrfname[i])] = types.YChild{"Vrfname", &ipv4Route.Vrfname[i]}
    }
    ipv4Route.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Route.EntityData.Leafs["prefix-string"] = types.YLeaf{"PrefixString", ipv4Route.PrefixString}
    ipv4Route.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", ipv4Route.PrefixLength}
    return &(ipv4Route.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData
// Data container.
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tag value. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    Tagvalue interface{}
}

func (ipv4RouteData *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Ipv4RouteData) GetEntityData() *types.CommonEntityData {
    ipv4RouteData.EntityData.YFilter = ipv4RouteData.YFilter
    ipv4RouteData.EntityData.YangName = "ipv4-route-data"
    ipv4RouteData.EntityData.BundleName = "cisco_ios_xr"
    ipv4RouteData.EntityData.ParentYangName = "ipv4-route"
    ipv4RouteData.EntityData.SegmentPath = "ipv4-route-data"
    ipv4RouteData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4RouteData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4RouteData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4RouteData.EntityData.Children = make(map[string]types.YChild)
    ipv4RouteData.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4RouteData.EntityData.Leafs["tagvalue"] = types.YLeaf{"Tagvalue", ipv4RouteData.Tagvalue}
    return &(ipv4RouteData.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname
// keys: vrfname
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Vrfname interface{}

    // Tag value. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    Tagvalue interface{}
}

func (vrfname *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route_Vrfname) GetEntityData() *types.CommonEntityData {
    vrfname.EntityData.YFilter = vrfname.YFilter
    vrfname.EntityData.YangName = "vrfname"
    vrfname.EntityData.BundleName = "cisco_ios_xr"
    vrfname.EntityData.ParentYangName = "ipv4-route"
    vrfname.EntityData.SegmentPath = "vrfname" + "[vrfname='" + fmt.Sprintf("%v", vrfname.Vrfname) + "']"
    vrfname.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfname.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfname.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfname.EntityData.Children = make(map[string]types.YChild)
    vrfname.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfname.EntityData.Leafs["vrfname"] = types.YLeaf{"Vrfname", vrfname.Vrfname}
    vrfname.EntityData.Leafs["tagvalue"] = types.YLeaf{"Tagvalue", vrfname.Tagvalue}
    return &(vrfname.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv6NARoute.
    Ipv6NaRoutes SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes

    // Table of IPv6PDRoute.
    Ipv6PdRoutes SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes
}

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetEntityData() *types.CommonEntityData {
    ipv6Route.EntityData.YFilter = ipv6Route.YFilter
    ipv6Route.EntityData.YangName = "ipv6-route"
    ipv6Route.EntityData.BundleName = "cisco_ios_xr"
    ipv6Route.EntityData.ParentYangName = "state-control-route"
    ipv6Route.EntityData.SegmentPath = "ipv6-route"
    ipv6Route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Route.EntityData.Children = make(map[string]types.YChild)
    ipv6Route.EntityData.Children["ipv6na-routes"] = types.YChild{"Ipv6NaRoutes", &ipv6Route.Ipv6NaRoutes}
    ipv6Route.EntityData.Children["ipv6pd-routes"] = types.YChild{"Ipv6PdRoutes", &ipv6Route.Ipv6PdRoutes}
    ipv6Route.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6Route.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes
// Table of IPv6NARoute
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute.
    Ipv6NaRoute []SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute
}

func (ipv6NaRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes) GetEntityData() *types.CommonEntityData {
    ipv6NaRoutes.EntityData.YFilter = ipv6NaRoutes.YFilter
    ipv6NaRoutes.EntityData.YangName = "ipv6na-routes"
    ipv6NaRoutes.EntityData.BundleName = "cisco_ios_xr"
    ipv6NaRoutes.EntityData.ParentYangName = "ipv6-route"
    ipv6NaRoutes.EntityData.SegmentPath = "ipv6na-routes"
    ipv6NaRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6NaRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6NaRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6NaRoutes.EntityData.Children = make(map[string]types.YChild)
    ipv6NaRoutes.EntityData.Children["ipv6na-route"] = types.YChild{"Ipv6NaRoute", nil}
    for i := range ipv6NaRoutes.Ipv6NaRoute {
        ipv6NaRoutes.EntityData.Children[types.GetSegmentPath(&ipv6NaRoutes.Ipv6NaRoute[i])] = types.YChild{"Ipv6NaRoute", &ipv6NaRoutes.Ipv6NaRoute[i]}
    }
    ipv6NaRoutes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6NaRoutes.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Vrfname interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: -2147483648..2147483647.
    PrefixLength interface{}

    // This attribute is a key. IPv6 address with prefix-length. The type is one
    // of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PrefixString interface{}

    // Tag value. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    Tagvalue interface{}
}

func (ipv6NaRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6NaRoutes_Ipv6NaRoute) GetEntityData() *types.CommonEntityData {
    ipv6NaRoute.EntityData.YFilter = ipv6NaRoute.YFilter
    ipv6NaRoute.EntityData.YangName = "ipv6na-route"
    ipv6NaRoute.EntityData.BundleName = "cisco_ios_xr"
    ipv6NaRoute.EntityData.ParentYangName = "ipv6na-routes"
    ipv6NaRoute.EntityData.SegmentPath = "ipv6na-route" + "[vrfname='" + fmt.Sprintf("%v", ipv6NaRoute.Vrfname) + "']" + "[prefix-length='" + fmt.Sprintf("%v", ipv6NaRoute.PrefixLength) + "']" + "[prefix-string='" + fmt.Sprintf("%v", ipv6NaRoute.PrefixString) + "']"
    ipv6NaRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6NaRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6NaRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6NaRoute.EntityData.Children = make(map[string]types.YChild)
    ipv6NaRoute.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6NaRoute.EntityData.Leafs["vrfname"] = types.YLeaf{"Vrfname", ipv6NaRoute.Vrfname}
    ipv6NaRoute.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", ipv6NaRoute.PrefixLength}
    ipv6NaRoute.EntityData.Leafs["prefix-string"] = types.YLeaf{"PrefixString", ipv6NaRoute.PrefixString}
    ipv6NaRoute.EntityData.Leafs["tagvalue"] = types.YLeaf{"Tagvalue", ipv6NaRoute.Tagvalue}
    return &(ipv6NaRoute.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes
// Table of IPv6PDRoute
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute.
    Ipv6PdRoute []SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute
}

func (ipv6PdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes) GetEntityData() *types.CommonEntityData {
    ipv6PdRoutes.EntityData.YFilter = ipv6PdRoutes.YFilter
    ipv6PdRoutes.EntityData.YangName = "ipv6pd-routes"
    ipv6PdRoutes.EntityData.BundleName = "cisco_ios_xr"
    ipv6PdRoutes.EntityData.ParentYangName = "ipv6-route"
    ipv6PdRoutes.EntityData.SegmentPath = "ipv6pd-routes"
    ipv6PdRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6PdRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6PdRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6PdRoutes.EntityData.Children = make(map[string]types.YChild)
    ipv6PdRoutes.EntityData.Children["ipv6pd-route"] = types.YChild{"Ipv6PdRoute", nil}
    for i := range ipv6PdRoutes.Ipv6PdRoute {
        ipv6PdRoutes.EntityData.Children[types.GetSegmentPath(&ipv6PdRoutes.Ipv6PdRoute[i])] = types.YChild{"Ipv6PdRoute", &ipv6PdRoutes.Ipv6PdRoute[i]}
    }
    ipv6PdRoutes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6PdRoutes.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Vrfname interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: -2147483648..2147483647.
    PrefixLength interface{}

    // This attribute is a key. IPv6 address with prefix-length. The type is one
    // of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PrefixString interface{}

    // Tag value. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    Tagvalue interface{}
}

func (ipv6PdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6PdRoutes_Ipv6PdRoute) GetEntityData() *types.CommonEntityData {
    ipv6PdRoute.EntityData.YFilter = ipv6PdRoute.YFilter
    ipv6PdRoute.EntityData.YangName = "ipv6pd-route"
    ipv6PdRoute.EntityData.BundleName = "cisco_ios_xr"
    ipv6PdRoute.EntityData.ParentYangName = "ipv6pd-routes"
    ipv6PdRoute.EntityData.SegmentPath = "ipv6pd-route" + "[vrfname='" + fmt.Sprintf("%v", ipv6PdRoute.Vrfname) + "']" + "[prefix-length='" + fmt.Sprintf("%v", ipv6PdRoute.PrefixLength) + "']" + "[prefix-string='" + fmt.Sprintf("%v", ipv6PdRoute.PrefixString) + "']"
    ipv6PdRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6PdRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6PdRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6PdRoute.EntityData.Children = make(map[string]types.YChild)
    ipv6PdRoute.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6PdRoute.EntityData.Leafs["vrfname"] = types.YLeaf{"Vrfname", ipv6PdRoute.Vrfname}
    ipv6PdRoute.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", ipv6PdRoute.PrefixLength}
    ipv6PdRoute.EntityData.Leafs["prefix-string"] = types.YLeaf{"PrefixString", ipv6PdRoute.PrefixString}
    ipv6PdRoute.EntityData.Leafs["tagvalue"] = types.YLeaf{"Tagvalue", ipv6PdRoute.Tagvalue}
    return &(ipv6PdRoute.EntityData)
}

// SubscriberRedundancy_RevertiveTimer
// None
type SubscriberRedundancy_RevertiveTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Value of MAX Revertive Timer. The type is interface{} with range: 1..65535.
    MaxValue interface{}

    // Value of revertive time in minutes. The type is interface{} with range:
    // 1..65535. Units are minute.
    Value interface{}
}

func (revertiveTimer *SubscriberRedundancy_RevertiveTimer) GetEntityData() *types.CommonEntityData {
    revertiveTimer.EntityData.YFilter = revertiveTimer.YFilter
    revertiveTimer.EntityData.YangName = "revertive-timer"
    revertiveTimer.EntityData.BundleName = "cisco_ios_xr"
    revertiveTimer.EntityData.ParentYangName = "subscriber-redundancy"
    revertiveTimer.EntityData.SegmentPath = "revertive-timer"
    revertiveTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    revertiveTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    revertiveTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    revertiveTimer.EntityData.Children = make(map[string]types.YChild)
    revertiveTimer.EntityData.Leafs = make(map[string]types.YLeaf)
    revertiveTimer.EntityData.Leafs["max-value"] = types.YLeaf{"MaxValue", revertiveTimer.MaxValue}
    revertiveTimer.EntityData.Leafs["value"] = types.YLeaf{"Value", revertiveTimer.Value}
    return &(revertiveTimer.EntityData)
}

