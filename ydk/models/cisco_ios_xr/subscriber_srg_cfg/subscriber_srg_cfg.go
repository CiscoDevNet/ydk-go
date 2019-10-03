// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-srg package configuration.
// 
// This module contains definitions
// for the following management objects:
//   subscriber-redundancy: Subscriber Redundancy configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
// This type is a presence type.
type SubscriberRedundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable Subscriber Redundancy configuration. Deletion of this object also
    // causes deletion of all associated objects under SubscriberRedundancy. The
    // type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Virtual MAC Prefix for Subscriber Redundancy. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacPrefix interface{}

    // Set preferred role. The type is SubscriberRedundancyGroupRole.
    PreferredRole interface{}

    // Source Interface for Redundancy Peer Communication. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}

    // Set slave. The type is SubscriberRedundancyGroupSlaveMode.
    SlaveMode interface{}

    // Set hold time (in Minutes). The type is interface{} with range: 1..65535.
    // Units are minute.
    HoldTimer interface{}

    // Set sync time (in Minutes). The type is interface{} with range: 1..255.
    // Units are minute.
    SyncTimer interface{}

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
    subscriberRedundancy.EntityData.AbsolutePath = subscriberRedundancy.EntityData.SegmentPath
    subscriberRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberRedundancy.EntityData.Children = types.NewOrderedMap()
    subscriberRedundancy.EntityData.Children.Append("groups", types.YChild{"Groups", &subscriberRedundancy.Groups})
    subscriberRedundancy.EntityData.Children.Append("revertive-timer", types.YChild{"RevertiveTimer", &subscriberRedundancy.RevertiveTimer})
    subscriberRedundancy.EntityData.Leafs = types.NewOrderedMap()
    subscriberRedundancy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", subscriberRedundancy.Enable})
    subscriberRedundancy.EntityData.Leafs.Append("virtual-mac-prefix", types.YLeaf{"VirtualMacPrefix", subscriberRedundancy.VirtualMacPrefix})
    subscriberRedundancy.EntityData.Leafs.Append("preferred-role", types.YLeaf{"PreferredRole", subscriberRedundancy.PreferredRole})
    subscriberRedundancy.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", subscriberRedundancy.SourceInterface})
    subscriberRedundancy.EntityData.Leafs.Append("slave-mode", types.YLeaf{"SlaveMode", subscriberRedundancy.SlaveMode})
    subscriberRedundancy.EntityData.Leafs.Append("hold-timer", types.YLeaf{"HoldTimer", subscriberRedundancy.HoldTimer})
    subscriberRedundancy.EntityData.Leafs.Append("sync-timer", types.YLeaf{"SyncTimer", subscriberRedundancy.SyncTimer})
    subscriberRedundancy.EntityData.Leafs.Append("redundancy-disable", types.YLeaf{"RedundancyDisable", subscriberRedundancy.RedundancyDisable})

    subscriberRedundancy.EntityData.YListKeys = []string {}

    return &(subscriberRedundancy.EntityData)
}

// SubscriberRedundancy_Groups
// Table of Group
type SubscriberRedundancy_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redundancy Group configuration. The type is slice of
    // SubscriberRedundancy_Groups_Group.
    Group []*SubscriberRedundancy_Groups_Group
}

func (groups *SubscriberRedundancy_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "subscriber-redundancy"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/" + groups.EntityData.SegmentPath
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// SubscriberRedundancy_Groups_Group
// Redundancy Group configuration
type SubscriberRedundancy_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..4000.
    GroupId interface{}

    // Disable Tracking Object for this Group. The type is interface{}.
    DisableTrackingObject interface{}

    // Core Tracking Object for this Group. The type is string.
    CoreTrackingObject interface{}

    // Enable Redundancy Group configuration. Deletion of this object also causes
    // deletion of all associated objects under Group. The type is interface{}.
    // This attribute is mandatory.
    Enable interface{}

    // Set preferred role. The type is SubscriberRedundancyGroupRole.
    PreferredRole interface{}

    // Description for this Group. The type is string.
    Description interface{}

    // Enter an IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    L2tpSourceIpAddress interface{}

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
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupId, "group-id")
    group.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("interface-list", types.YChild{"InterfaceList", &group.InterfaceList})
    group.EntityData.Children.Append("peer", types.YChild{"Peer", &group.Peer})
    group.EntityData.Children.Append("revertive-timer", types.YChild{"RevertiveTimer", &group.RevertiveTimer})
    group.EntityData.Children.Append("virtual-mac", types.YChild{"VirtualMac", &group.VirtualMac})
    group.EntityData.Children.Append("state-control-route", types.YChild{"StateControlRoute", &group.StateControlRoute})
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", group.GroupId})
    group.EntityData.Leafs.Append("disable-tracking-object", types.YLeaf{"DisableTrackingObject", group.DisableTrackingObject})
    group.EntityData.Leafs.Append("core-tracking-object", types.YLeaf{"CoreTrackingObject", group.CoreTrackingObject})
    group.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", group.Enable})
    group.EntityData.Leafs.Append("preferred-role", types.YLeaf{"PreferredRole", group.PreferredRole})
    group.EntityData.Leafs.Append("description", types.YLeaf{"Description", group.Description})
    group.EntityData.Leafs.Append("l2tp-source-ip-address", types.YLeaf{"L2tpSourceIpAddress", group.L2tpSourceIpAddress})
    group.EntityData.Leafs.Append("slave-mode", types.YLeaf{"SlaveMode", group.SlaveMode})
    group.EntityData.Leafs.Append("hold-timer", types.YLeaf{"HoldTimer", group.HoldTimer})
    group.EntityData.Leafs.Append("access-tracking-object", types.YLeaf{"AccessTrackingObject", group.AccessTrackingObject})
    group.EntityData.Leafs.Append("enable-fast-switchover", types.YLeaf{"EnableFastSwitchover", group.EnableFastSwitchover})
    group.EntityData.Leafs.Append("redundancy-disable", types.YLeaf{"RedundancyDisable", group.RedundancyDisable})

    group.EntityData.YListKeys = []string {"GroupId"}

    return &(group.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList
// List of Interfaces for this Group
// This type is a presence type.
type SubscriberRedundancy_Groups_Group_InterfaceList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable List of Interfaces for this Group. Deletion of this object also
    // causes deletion of all associated objects under InterfaceList . The type is
    // interface{}. This attribute is mandatory.
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
    interfaceList.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/" + interfaceList.EntityData.SegmentPath
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = types.NewOrderedMap()
    interfaceList.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaceList.Interfaces})
    interfaceList.EntityData.Children.Append("interface-ranges", types.YChild{"InterfaceRanges", &interfaceList.InterfaceRanges})
    interfaceList.EntityData.Leafs = types.NewOrderedMap()
    interfaceList.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceList.Enable})

    interfaceList.EntityData.YListKeys = []string {}

    return &(interfaceList.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces
// Table of Interface
type SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface.
    Interface []*SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
}

func (interfaces *SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-list"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/interface-list/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
// Interface for this Group
type SubscriberRedundancy_Groups_Group_InterfaceList_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/interface-list/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", self.InterfaceId})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges
// Table of InterfaceRange
type SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange.
    InterfaceRange []*SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
}

func (interfaceRanges *SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetEntityData() *types.CommonEntityData {
    interfaceRanges.EntityData.YFilter = interfaceRanges.YFilter
    interfaceRanges.EntityData.YangName = "interface-ranges"
    interfaceRanges.EntityData.BundleName = "cisco_ios_xr"
    interfaceRanges.EntityData.ParentYangName = "interface-list"
    interfaceRanges.EntityData.SegmentPath = "interface-ranges"
    interfaceRanges.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/interface-list/" + interfaceRanges.EntityData.SegmentPath
    interfaceRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceRanges.EntityData.Children = types.NewOrderedMap()
    interfaceRanges.EntityData.Children.Append("interface-range", types.YChild{"InterfaceRange", nil})
    for i := range interfaceRanges.InterfaceRange {
        interfaceRanges.EntityData.Children.Append(types.GetSegmentPath(interfaceRanges.InterfaceRange[i]), types.YChild{"InterfaceRange", interfaceRanges.InterfaceRange[i]})
    }
    interfaceRanges.EntityData.Leafs = types.NewOrderedMap()

    interfaceRanges.EntityData.YListKeys = []string {}

    return &(interfaceRanges.EntityData)
}

// SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
// Interface for this Group
type SubscriberRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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
    interfaceRange.EntityData.SegmentPath = "interface-range" + types.AddKeyToken(interfaceRange.InterfaceName, "interface-name") + types.AddKeyToken(interfaceRange.SubInterfaceRangeStart, "sub-interface-range-start") + types.AddKeyToken(interfaceRange.SubInterfaceRangeEnd, "sub-interface-range-end")
    interfaceRange.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/interface-list/interface-ranges/" + interfaceRange.EntityData.SegmentPath
    interfaceRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceRange.EntityData.Children = types.NewOrderedMap()
    interfaceRange.EntityData.Leafs = types.NewOrderedMap()
    interfaceRange.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceRange.InterfaceName})
    interfaceRange.EntityData.Leafs.Append("sub-interface-range-start", types.YLeaf{"SubInterfaceRangeStart", interfaceRange.SubInterfaceRangeStart})
    interfaceRange.EntityData.Leafs.Append("sub-interface-range-end", types.YLeaf{"SubInterfaceRangeEnd", interfaceRange.SubInterfaceRangeEnd})
    interfaceRange.EntityData.Leafs.Append("interface-id-range-start", types.YLeaf{"InterfaceIdRangeStart", interfaceRange.InterfaceIdRangeStart})
    interfaceRange.EntityData.Leafs.Append("interface-id-range-end", types.YLeaf{"InterfaceIdRangeEnd", interfaceRange.InterfaceIdRangeEnd})

    interfaceRange.EntityData.YListKeys = []string {"InterfaceName", "SubInterfaceRangeStart", "SubInterfaceRangeEnd"}

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
    peer.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Children.Append("ipaddress", types.YChild{"Ipaddress", &peer.Ipaddress})
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("route-add-disable", types.YLeaf{"RouteAddDisable", peer.RouteAddDisable})

    peer.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixString interface{}
}

func (ipaddress *SubscriberRedundancy_Groups_Group_Peer_Ipaddress) GetEntityData() *types.CommonEntityData {
    ipaddress.EntityData.YFilter = ipaddress.YFilter
    ipaddress.EntityData.YangName = "ipaddress"
    ipaddress.EntityData.BundleName = "cisco_ios_xr"
    ipaddress.EntityData.ParentYangName = "peer"
    ipaddress.EntityData.SegmentPath = "ipaddress"
    ipaddress.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/peer/" + ipaddress.EntityData.SegmentPath
    ipaddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipaddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipaddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipaddress.EntityData.Children = types.NewOrderedMap()
    ipaddress.EntityData.Leafs = types.NewOrderedMap()
    ipaddress.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", ipaddress.AddressFamily})
    ipaddress.EntityData.Leafs.Append("prefix-string", types.YLeaf{"PrefixString", ipaddress.PrefixString})

    ipaddress.EntityData.YListKeys = []string {}

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
    revertiveTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/" + revertiveTimer.EntityData.SegmentPath
    revertiveTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    revertiveTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    revertiveTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    revertiveTimer.EntityData.Children = types.NewOrderedMap()
    revertiveTimer.EntityData.Leafs = types.NewOrderedMap()
    revertiveTimer.EntityData.Leafs.Append("max-value", types.YLeaf{"MaxValue", revertiveTimer.MaxValue})
    revertiveTimer.EntityData.Leafs.Append("value", types.YLeaf{"Value", revertiveTimer.Value})

    revertiveTimer.EntityData.YListKeys = []string {}

    return &(revertiveTimer.EntityData)
}

// SubscriberRedundancy_Groups_Group_VirtualMac
// Virtual MAC Address for this Group
type SubscriberRedundancy_Groups_Group_VirtualMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Virtual MAC Address for this Group. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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
    virtualMac.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/" + virtualMac.EntityData.SegmentPath
    virtualMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualMac.EntityData.Children = types.NewOrderedMap()
    virtualMac.EntityData.Leafs = types.NewOrderedMap()
    virtualMac.EntityData.Leafs.Append("address", types.YLeaf{"Address", virtualMac.Address})
    virtualMac.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", virtualMac.Disable})

    virtualMac.EntityData.YListKeys = []string {}

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
    stateControlRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/" + stateControlRoute.EntityData.SegmentPath
    stateControlRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateControlRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateControlRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateControlRoute.EntityData.Children = types.NewOrderedMap()
    stateControlRoute.EntityData.Children.Append("ipv4-routes", types.YChild{"Ipv4Routes", &stateControlRoute.Ipv4Routes})
    stateControlRoute.EntityData.Children.Append("ipv6-route", types.YChild{"Ipv6Route", &stateControlRoute.Ipv6Route})
    stateControlRoute.EntityData.Leafs = types.NewOrderedMap()

    stateControlRoute.EntityData.YListKeys = []string {}

    return &(stateControlRoute.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes
// Table of IPv4Route
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route.
    Ipv4Route []*SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route
}

func (ipv4Routes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes) GetEntityData() *types.CommonEntityData {
    ipv4Routes.EntityData.YFilter = ipv4Routes.YFilter
    ipv4Routes.EntityData.YangName = "ipv4-routes"
    ipv4Routes.EntityData.BundleName = "cisco_ios_xr"
    ipv4Routes.EntityData.ParentYangName = "state-control-route"
    ipv4Routes.EntityData.SegmentPath = "ipv4-routes"
    ipv4Routes.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/state-control-route/" + ipv4Routes.EntityData.SegmentPath
    ipv4Routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Routes.EntityData.Children = types.NewOrderedMap()
    ipv4Routes.EntityData.Children.Append("ipv4-route", types.YChild{"Ipv4Route", nil})
    for i := range ipv4Routes.Ipv4Route {
        ipv4Routes.EntityData.Children.Append(types.GetSegmentPath(ipv4Routes.Ipv4Route[i]), types.YChild{"Ipv4Route", ipv4Routes.Ipv4Route[i]})
    }
    ipv4Routes.EntityData.Leafs = types.NewOrderedMap()

    ipv4Routes.EntityData.YListKeys = []string {}

    return &(ipv4Routes.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Vrfname interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: 0..4294967295.
    PrefixLength interface{}

    // This attribute is a key. IPv4 address with prefix-length. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixString interface{}

    // Tag value. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory.
    Tagvalue interface{}
}

func (ipv4Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv4Routes_Ipv4Route) GetEntityData() *types.CommonEntityData {
    ipv4Route.EntityData.YFilter = ipv4Route.YFilter
    ipv4Route.EntityData.YangName = "ipv4-route"
    ipv4Route.EntityData.BundleName = "cisco_ios_xr"
    ipv4Route.EntityData.ParentYangName = "ipv4-routes"
    ipv4Route.EntityData.SegmentPath = "ipv4-route" + types.AddKeyToken(ipv4Route.Vrfname, "vrfname") + types.AddKeyToken(ipv4Route.PrefixLength, "prefix-length") + types.AddKeyToken(ipv4Route.PrefixString, "prefix-string")
    ipv4Route.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/state-control-route/ipv4-routes/" + ipv4Route.EntityData.SegmentPath
    ipv4Route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Route.EntityData.Children = types.NewOrderedMap()
    ipv4Route.EntityData.Leafs = types.NewOrderedMap()
    ipv4Route.EntityData.Leafs.Append("vrfname", types.YLeaf{"Vrfname", ipv4Route.Vrfname})
    ipv4Route.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", ipv4Route.PrefixLength})
    ipv4Route.EntityData.Leafs.Append("prefix-string", types.YLeaf{"PrefixString", ipv4Route.PrefixString})
    ipv4Route.EntityData.Leafs.Append("tagvalue", types.YLeaf{"Tagvalue", ipv4Route.Tagvalue})

    ipv4Route.EntityData.YListKeys = []string {"Vrfname", "PrefixLength", "PrefixString"}

    return &(ipv4Route.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv6NARoute.
    Ipv6naRoutes SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6naRoutes

    // Table of IPv6PDRoute.
    Ipv6pdRoutes SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6pdRoutes
}

func (ipv6Route *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route) GetEntityData() *types.CommonEntityData {
    ipv6Route.EntityData.YFilter = ipv6Route.YFilter
    ipv6Route.EntityData.YangName = "ipv6-route"
    ipv6Route.EntityData.BundleName = "cisco_ios_xr"
    ipv6Route.EntityData.ParentYangName = "state-control-route"
    ipv6Route.EntityData.SegmentPath = "ipv6-route"
    ipv6Route.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/state-control-route/" + ipv6Route.EntityData.SegmentPath
    ipv6Route.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Route.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Route.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Route.EntityData.Children = types.NewOrderedMap()
    ipv6Route.EntityData.Children.Append("ipv6na-routes", types.YChild{"Ipv6naRoutes", &ipv6Route.Ipv6naRoutes})
    ipv6Route.EntityData.Children.Append("ipv6pd-routes", types.YChild{"Ipv6pdRoutes", &ipv6Route.Ipv6pdRoutes})
    ipv6Route.EntityData.Leafs = types.NewOrderedMap()

    ipv6Route.EntityData.YListKeys = []string {}

    return &(ipv6Route.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6naRoutes
// Table of IPv6NARoute
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6naRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6naRoutes_Ipv6naRoute.
    Ipv6naRoute []*SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6naRoutes_Ipv6naRoute
}

func (ipv6naRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6naRoutes) GetEntityData() *types.CommonEntityData {
    ipv6naRoutes.EntityData.YFilter = ipv6naRoutes.YFilter
    ipv6naRoutes.EntityData.YangName = "ipv6na-routes"
    ipv6naRoutes.EntityData.BundleName = "cisco_ios_xr"
    ipv6naRoutes.EntityData.ParentYangName = "ipv6-route"
    ipv6naRoutes.EntityData.SegmentPath = "ipv6na-routes"
    ipv6naRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/state-control-route/ipv6-route/" + ipv6naRoutes.EntityData.SegmentPath
    ipv6naRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6naRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6naRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6naRoutes.EntityData.Children = types.NewOrderedMap()
    ipv6naRoutes.EntityData.Children.Append("ipv6na-route", types.YChild{"Ipv6naRoute", nil})
    for i := range ipv6naRoutes.Ipv6naRoute {
        ipv6naRoutes.EntityData.Children.Append(types.GetSegmentPath(ipv6naRoutes.Ipv6naRoute[i]), types.YChild{"Ipv6naRoute", ipv6naRoutes.Ipv6naRoute[i]})
    }
    ipv6naRoutes.EntityData.Leafs = types.NewOrderedMap()

    ipv6naRoutes.EntityData.YListKeys = []string {}

    return &(ipv6naRoutes.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6naRoutes_Ipv6naRoute
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6naRoutes_Ipv6naRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Vrfname interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: 0..4294967295.
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

func (ipv6naRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6naRoutes_Ipv6naRoute) GetEntityData() *types.CommonEntityData {
    ipv6naRoute.EntityData.YFilter = ipv6naRoute.YFilter
    ipv6naRoute.EntityData.YangName = "ipv6na-route"
    ipv6naRoute.EntityData.BundleName = "cisco_ios_xr"
    ipv6naRoute.EntityData.ParentYangName = "ipv6na-routes"
    ipv6naRoute.EntityData.SegmentPath = "ipv6na-route" + types.AddKeyToken(ipv6naRoute.Vrfname, "vrfname") + types.AddKeyToken(ipv6naRoute.PrefixLength, "prefix-length") + types.AddKeyToken(ipv6naRoute.PrefixString, "prefix-string")
    ipv6naRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/state-control-route/ipv6-route/ipv6na-routes/" + ipv6naRoute.EntityData.SegmentPath
    ipv6naRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6naRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6naRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6naRoute.EntityData.Children = types.NewOrderedMap()
    ipv6naRoute.EntityData.Leafs = types.NewOrderedMap()
    ipv6naRoute.EntityData.Leafs.Append("vrfname", types.YLeaf{"Vrfname", ipv6naRoute.Vrfname})
    ipv6naRoute.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", ipv6naRoute.PrefixLength})
    ipv6naRoute.EntityData.Leafs.Append("prefix-string", types.YLeaf{"PrefixString", ipv6naRoute.PrefixString})
    ipv6naRoute.EntityData.Leafs.Append("tagvalue", types.YLeaf{"Tagvalue", ipv6naRoute.Tagvalue})

    ipv6naRoute.EntityData.YListKeys = []string {"Vrfname", "PrefixLength", "PrefixString"}

    return &(ipv6naRoute.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6pdRoutes
// Table of IPv6PDRoute
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6pdRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None. The type is slice of
    // SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6pdRoutes_Ipv6pdRoute.
    Ipv6pdRoute []*SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6pdRoutes_Ipv6pdRoute
}

func (ipv6pdRoutes *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6pdRoutes) GetEntityData() *types.CommonEntityData {
    ipv6pdRoutes.EntityData.YFilter = ipv6pdRoutes.YFilter
    ipv6pdRoutes.EntityData.YangName = "ipv6pd-routes"
    ipv6pdRoutes.EntityData.BundleName = "cisco_ios_xr"
    ipv6pdRoutes.EntityData.ParentYangName = "ipv6-route"
    ipv6pdRoutes.EntityData.SegmentPath = "ipv6pd-routes"
    ipv6pdRoutes.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/state-control-route/ipv6-route/" + ipv6pdRoutes.EntityData.SegmentPath
    ipv6pdRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6pdRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6pdRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6pdRoutes.EntityData.Children = types.NewOrderedMap()
    ipv6pdRoutes.EntityData.Children.Append("ipv6pd-route", types.YChild{"Ipv6pdRoute", nil})
    for i := range ipv6pdRoutes.Ipv6pdRoute {
        ipv6pdRoutes.EntityData.Children.Append(types.GetSegmentPath(ipv6pdRoutes.Ipv6pdRoute[i]), types.YChild{"Ipv6pdRoute", ipv6pdRoutes.Ipv6pdRoute[i]})
    }
    ipv6pdRoutes.EntityData.Leafs = types.NewOrderedMap()

    ipv6pdRoutes.EntityData.YListKeys = []string {}

    return &(ipv6pdRoutes.EntityData)
}

// SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6pdRoutes_Ipv6pdRoute
// None
type SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6pdRoutes_Ipv6pdRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Vrfname interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: 0..4294967295.
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

func (ipv6pdRoute *SubscriberRedundancy_Groups_Group_StateControlRoute_Ipv6Route_Ipv6pdRoutes_Ipv6pdRoute) GetEntityData() *types.CommonEntityData {
    ipv6pdRoute.EntityData.YFilter = ipv6pdRoute.YFilter
    ipv6pdRoute.EntityData.YangName = "ipv6pd-route"
    ipv6pdRoute.EntityData.BundleName = "cisco_ios_xr"
    ipv6pdRoute.EntityData.ParentYangName = "ipv6pd-routes"
    ipv6pdRoute.EntityData.SegmentPath = "ipv6pd-route" + types.AddKeyToken(ipv6pdRoute.Vrfname, "vrfname") + types.AddKeyToken(ipv6pdRoute.PrefixLength, "prefix-length") + types.AddKeyToken(ipv6pdRoute.PrefixString, "prefix-string")
    ipv6pdRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/groups/group/state-control-route/ipv6-route/ipv6pd-routes/" + ipv6pdRoute.EntityData.SegmentPath
    ipv6pdRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6pdRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6pdRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6pdRoute.EntityData.Children = types.NewOrderedMap()
    ipv6pdRoute.EntityData.Leafs = types.NewOrderedMap()
    ipv6pdRoute.EntityData.Leafs.Append("vrfname", types.YLeaf{"Vrfname", ipv6pdRoute.Vrfname})
    ipv6pdRoute.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", ipv6pdRoute.PrefixLength})
    ipv6pdRoute.EntityData.Leafs.Append("prefix-string", types.YLeaf{"PrefixString", ipv6pdRoute.PrefixString})
    ipv6pdRoute.EntityData.Leafs.Append("tagvalue", types.YLeaf{"Tagvalue", ipv6pdRoute.Tagvalue})

    ipv6pdRoute.EntityData.YListKeys = []string {"Vrfname", "PrefixLength", "PrefixString"}

    return &(ipv6pdRoute.EntityData)
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
    revertiveTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-srg-cfg:subscriber-redundancy/" + revertiveTimer.EntityData.SegmentPath
    revertiveTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    revertiveTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    revertiveTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    revertiveTimer.EntityData.Children = types.NewOrderedMap()
    revertiveTimer.EntityData.Leafs = types.NewOrderedMap()
    revertiveTimer.EntityData.Leafs.Append("max-value", types.YLeaf{"MaxValue", revertiveTimer.MaxValue})
    revertiveTimer.EntityData.Leafs.Append("value", types.YLeaf{"Value", revertiveTimer.Value})

    revertiveTimer.EntityData.YListKeys = []string {}

    return &(revertiveTimer.EntityData)
}

