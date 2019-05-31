// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-serg package configuration.
// 
// This module contains definitions
// for the following management objects:
//   session-redundancy: Session Redundancy configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// SessionRedundancyGroupRole represents Session redundancy group role
type SessionRedundancyGroupRole string

const (
    // Master Role
    SessionRedundancyGroupRole_master SessionRedundancyGroupRole = "master"

    // Slave Role
    SessionRedundancyGroupRole_slave SessionRedundancyGroupRole = "slave"
)

// SergAddrFamily represents Serg addr family
type SergAddrFamily string

const (
    // IPv4
    SergAddrFamily_ipv4 SergAddrFamily = "ipv4"

    // IPv6
    SergAddrFamily_ipv6 SergAddrFamily = "ipv6"
)

// SessionRedundancy
// Session Redundancy configuration
type SessionRedundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable. The type is interface{}.
    RedundancyDisable interface{}

    // Enable Session Redundancy configuration. Deletion of this object also
    // causes deletion of all associated objects under SessionRedundancy. The type
    // is interface{}.
    Enable interface{}

    // Source Interface for Redundancy Peer Communication. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
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

func (sessionRedundancy *SessionRedundancy) GetEntityData() *types.CommonEntityData {
    sessionRedundancy.EntityData.YFilter = sessionRedundancy.YFilter
    sessionRedundancy.EntityData.YangName = "session-redundancy"
    sessionRedundancy.EntityData.BundleName = "cisco_ios_xr"
    sessionRedundancy.EntityData.ParentYangName = "Cisco-IOS-XR-infra-serg-cfg"
    sessionRedundancy.EntityData.SegmentPath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy"
    sessionRedundancy.EntityData.AbsolutePath = sessionRedundancy.EntityData.SegmentPath
    sessionRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionRedundancy.EntityData.Children = types.NewOrderedMap()
    sessionRedundancy.EntityData.Children.Append("groups", types.YChild{"Groups", &sessionRedundancy.Groups})
    sessionRedundancy.EntityData.Children.Append("revertive-timer", types.YChild{"RevertiveTimer", &sessionRedundancy.RevertiveTimer})
    sessionRedundancy.EntityData.Leafs = types.NewOrderedMap()
    sessionRedundancy.EntityData.Leafs.Append("redundancy-disable", types.YLeaf{"RedundancyDisable", sessionRedundancy.RedundancyDisable})
    sessionRedundancy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", sessionRedundancy.Enable})
    sessionRedundancy.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", sessionRedundancy.SourceInterface})
    sessionRedundancy.EntityData.Leafs.Append("preferred-role", types.YLeaf{"PreferredRole", sessionRedundancy.PreferredRole})
    sessionRedundancy.EntityData.Leafs.Append("hold-timer", types.YLeaf{"HoldTimer", sessionRedundancy.HoldTimer})

    sessionRedundancy.EntityData.YListKeys = []string {}

    return &(sessionRedundancy.EntityData)
}

// SessionRedundancy_Groups
// Table of Group
type SessionRedundancy_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redundancy Group configuration. The type is slice of
    // SessionRedundancy_Groups_Group.
    Group []*SessionRedundancy_Groups_Group
}

func (groups *SessionRedundancy_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "session-redundancy"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/" + groups.EntityData.SegmentPath
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

// SessionRedundancy_Groups_Group
// Redundancy Group configuration
type SessionRedundancy_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Set operation mode. The type is interface{}.
    ModeActive interface{}

    // None.
    Peer SessionRedundancy_Groups_Group_Peer

    // None.
    RevertiveTimer SessionRedundancy_Groups_Group_RevertiveTimer

    // List of Interfaces for this Group.
    InterfaceList SessionRedundancy_Groups_Group_InterfaceList

    // List of Pool-names for this Group.
    PoolList SessionRedundancy_Groups_Group_PoolList
}

func (group *SessionRedundancy_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupId, "group-id")
    group.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("peer", types.YChild{"Peer", &group.Peer})
    group.EntityData.Children.Append("revertive-timer", types.YChild{"RevertiveTimer", &group.RevertiveTimer})
    group.EntityData.Children.Append("interface-list", types.YChild{"InterfaceList", &group.InterfaceList})
    group.EntityData.Children.Append("pool-list", types.YChild{"PoolList", &group.PoolList})
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", group.GroupId})
    group.EntityData.Leafs.Append("core-tracking-object", types.YLeaf{"CoreTrackingObject", group.CoreTrackingObject})
    group.EntityData.Leafs.Append("disable-tracking-object", types.YLeaf{"DisableTrackingObject", group.DisableTrackingObject})
    group.EntityData.Leafs.Append("redundancy-disable", types.YLeaf{"RedundancyDisable", group.RedundancyDisable})
    group.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", group.Enable})
    group.EntityData.Leafs.Append("description", types.YLeaf{"Description", group.Description})
    group.EntityData.Leafs.Append("access-tracking-object", types.YLeaf{"AccessTrackingObject", group.AccessTrackingObject})
    group.EntityData.Leafs.Append("preferred-role", types.YLeaf{"PreferredRole", group.PreferredRole})
    group.EntityData.Leafs.Append("hold-timer", types.YLeaf{"HoldTimer", group.HoldTimer})
    group.EntityData.Leafs.Append("mode-active", types.YLeaf{"ModeActive", group.ModeActive})

    group.EntityData.YListKeys = []string {"GroupId"}

    return &(group.EntityData)
}

// SessionRedundancy_Groups_Group_Peer
// None
type SessionRedundancy_Groups_Group_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 or IPv6 Address of SERG Peer.
    Ipaddress SessionRedundancy_Groups_Group_Peer_Ipaddress
}

func (peer *SessionRedundancy_Groups_Group_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "group"
    peer.EntityData.SegmentPath = "peer"
    peer.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Children.Append("ipaddress", types.YChild{"Ipaddress", &peer.Ipaddress})
    peer.EntityData.Leafs = types.NewOrderedMap()

    peer.EntityData.YListKeys = []string {}

    return &(peer.EntityData)
}

// SessionRedundancy_Groups_Group_Peer_Ipaddress
// IPv4 or IPv6 Address of SERG Peer
type SessionRedundancy_Groups_Group_Peer_Ipaddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of IPv4/IPv6 address. The type is SergAddrFamily.
    AddressFamily interface{}

    // IPv4/IPv6 address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PrefixString interface{}
}

func (ipaddress *SessionRedundancy_Groups_Group_Peer_Ipaddress) GetEntityData() *types.CommonEntityData {
    ipaddress.EntityData.YFilter = ipaddress.YFilter
    ipaddress.EntityData.YangName = "ipaddress"
    ipaddress.EntityData.BundleName = "cisco_ios_xr"
    ipaddress.EntityData.ParentYangName = "peer"
    ipaddress.EntityData.SegmentPath = "ipaddress"
    ipaddress.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/peer/" + ipaddress.EntityData.SegmentPath
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

// SessionRedundancy_Groups_Group_RevertiveTimer
// None
type SessionRedundancy_Groups_Group_RevertiveTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Value of MAX Revertive Timer. The type is interface{} with range: 1..65535.
    MaxValue interface{}

    // Value of revertive time in minutes. The type is interface{} with range:
    // 1..65535. Units are minute.
    Value interface{}
}

func (revertiveTimer *SessionRedundancy_Groups_Group_RevertiveTimer) GetEntityData() *types.CommonEntityData {
    revertiveTimer.EntityData.YFilter = revertiveTimer.YFilter
    revertiveTimer.EntityData.YangName = "revertive-timer"
    revertiveTimer.EntityData.BundleName = "cisco_ios_xr"
    revertiveTimer.EntityData.ParentYangName = "group"
    revertiveTimer.EntityData.SegmentPath = "revertive-timer"
    revertiveTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/" + revertiveTimer.EntityData.SegmentPath
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

// SessionRedundancy_Groups_Group_InterfaceList
// List of Interfaces for this Group
type SessionRedundancy_Groups_Group_InterfaceList struct {
    EntityData types.CommonEntityData
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

func (interfaceList *SessionRedundancy_Groups_Group_InterfaceList) GetEntityData() *types.CommonEntityData {
    interfaceList.EntityData.YFilter = interfaceList.YFilter
    interfaceList.EntityData.YangName = "interface-list"
    interfaceList.EntityData.BundleName = "cisco_ios_xr"
    interfaceList.EntityData.ParentYangName = "group"
    interfaceList.EntityData.SegmentPath = "interface-list"
    interfaceList.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/" + interfaceList.EntityData.SegmentPath
    interfaceList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceList.EntityData.Children = types.NewOrderedMap()
    interfaceList.EntityData.Children.Append("interface-ranges", types.YChild{"InterfaceRanges", &interfaceList.InterfaceRanges})
    interfaceList.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaceList.Interfaces})
    interfaceList.EntityData.Leafs = types.NewOrderedMap()
    interfaceList.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", interfaceList.Enable})

    interfaceList.EntityData.YListKeys = []string {}

    return &(interfaceList.EntityData)
}

// SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges
// Table of InterfaceRange
type SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange.
    InterfaceRange []*SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
}

func (interfaceRanges *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges) GetEntityData() *types.CommonEntityData {
    interfaceRanges.EntityData.YFilter = interfaceRanges.YFilter
    interfaceRanges.EntityData.YangName = "interface-ranges"
    interfaceRanges.EntityData.BundleName = "cisco_ios_xr"
    interfaceRanges.EntityData.ParentYangName = "interface-list"
    interfaceRanges.EntityData.SegmentPath = "interface-ranges"
    interfaceRanges.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/interface-list/" + interfaceRanges.EntityData.SegmentPath
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

// SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange
// Interface for this Group
type SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
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

func (interfaceRange *SessionRedundancy_Groups_Group_InterfaceList_InterfaceRanges_InterfaceRange) GetEntityData() *types.CommonEntityData {
    interfaceRange.EntityData.YFilter = interfaceRange.YFilter
    interfaceRange.EntityData.YangName = "interface-range"
    interfaceRange.EntityData.BundleName = "cisco_ios_xr"
    interfaceRange.EntityData.ParentYangName = "interface-ranges"
    interfaceRange.EntityData.SegmentPath = "interface-range" + types.AddKeyToken(interfaceRange.InterfaceName, "interface-name") + types.AddKeyToken(interfaceRange.SubInterfaceRangeStart, "sub-interface-range-start") + types.AddKeyToken(interfaceRange.SubInterfaceRangeEnd, "sub-interface-range-end")
    interfaceRange.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/interface-list/interface-ranges/" + interfaceRange.EntityData.SegmentPath
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

// SessionRedundancy_Groups_Group_InterfaceList_Interfaces
// Table of Interface
type SessionRedundancy_Groups_Group_InterfaceList_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface for this Group. The type is slice of
    // SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface.
    Interface []*SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
}

func (interfaces *SessionRedundancy_Groups_Group_InterfaceList_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-list"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/interface-list/" + interfaces.EntityData.SegmentPath
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

// SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface
// Interface for this Group
type SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Interface Id for the interface. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    InterfaceId interface{}
}

func (self *SessionRedundancy_Groups_Group_InterfaceList_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/interface-list/interfaces/" + self.EntityData.SegmentPath
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

// SessionRedundancy_Groups_Group_PoolList
// List of Pool-names for this Group
type SessionRedundancy_Groups_Group_PoolList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable List of Pool-names for this Group. Deletion of this object also
    // causes deletion of all associated objects under PoolList. The type is
    // interface{}.
    Enable interface{}

    // Table of PoolName.
    PoolNames SessionRedundancy_Groups_Group_PoolList_PoolNames
}

func (poolList *SessionRedundancy_Groups_Group_PoolList) GetEntityData() *types.CommonEntityData {
    poolList.EntityData.YFilter = poolList.YFilter
    poolList.EntityData.YangName = "pool-list"
    poolList.EntityData.BundleName = "cisco_ios_xr"
    poolList.EntityData.ParentYangName = "group"
    poolList.EntityData.SegmentPath = "pool-list"
    poolList.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/" + poolList.EntityData.SegmentPath
    poolList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    poolList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    poolList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    poolList.EntityData.Children = types.NewOrderedMap()
    poolList.EntityData.Children.Append("pool-names", types.YChild{"PoolNames", &poolList.PoolNames})
    poolList.EntityData.Leafs = types.NewOrderedMap()
    poolList.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", poolList.Enable})

    poolList.EntityData.YListKeys = []string {}

    return &(poolList.EntityData)
}

// SessionRedundancy_Groups_Group_PoolList_PoolNames
// Table of PoolName
type SessionRedundancy_Groups_Group_PoolList_PoolNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Pool Name. The type is slice of
    // SessionRedundancy_Groups_Group_PoolList_PoolNames_PoolName.
    PoolName []*SessionRedundancy_Groups_Group_PoolList_PoolNames_PoolName
}

func (poolNames *SessionRedundancy_Groups_Group_PoolList_PoolNames) GetEntityData() *types.CommonEntityData {
    poolNames.EntityData.YFilter = poolNames.YFilter
    poolNames.EntityData.YangName = "pool-names"
    poolNames.EntityData.BundleName = "cisco_ios_xr"
    poolNames.EntityData.ParentYangName = "pool-list"
    poolNames.EntityData.SegmentPath = "pool-names"
    poolNames.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/pool-list/" + poolNames.EntityData.SegmentPath
    poolNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    poolNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    poolNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    poolNames.EntityData.Children = types.NewOrderedMap()
    poolNames.EntityData.Children.Append("pool-name", types.YChild{"PoolName", nil})
    for i := range poolNames.PoolName {
        poolNames.EntityData.Children.Append(types.GetSegmentPath(poolNames.PoolName[i]), types.YChild{"PoolName", poolNames.PoolName[i]})
    }
    poolNames.EntityData.Leafs = types.NewOrderedMap()

    poolNames.EntityData.YListKeys = []string {}

    return &(poolNames.EntityData)
}

// SessionRedundancy_Groups_Group_PoolList_PoolNames_PoolName
// Address Pool Name
type SessionRedundancy_Groups_Group_PoolList_PoolNames_PoolName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Pool name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    PoolName interface{}
}

func (poolName *SessionRedundancy_Groups_Group_PoolList_PoolNames_PoolName) GetEntityData() *types.CommonEntityData {
    poolName.EntityData.YFilter = poolName.YFilter
    poolName.EntityData.YangName = "pool-name"
    poolName.EntityData.BundleName = "cisco_ios_xr"
    poolName.EntityData.ParentYangName = "pool-names"
    poolName.EntityData.SegmentPath = "pool-name" + types.AddKeyToken(poolName.PoolName, "pool-name")
    poolName.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/groups/group/pool-list/pool-names/" + poolName.EntityData.SegmentPath
    poolName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    poolName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    poolName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    poolName.EntityData.Children = types.NewOrderedMap()
    poolName.EntityData.Leafs = types.NewOrderedMap()
    poolName.EntityData.Leafs.Append("pool-name", types.YLeaf{"PoolName", poolName.PoolName})

    poolName.EntityData.YListKeys = []string {"PoolName"}

    return &(poolName.EntityData)
}

// SessionRedundancy_RevertiveTimer
// None
type SessionRedundancy_RevertiveTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Value of MAX Revertive Timer. The type is interface{} with range: 1..65535.
    MaxValue interface{}

    // Value of revertive time in minutes. The type is interface{} with range:
    // 1..65535. Units are minute.
    Value interface{}
}

func (revertiveTimer *SessionRedundancy_RevertiveTimer) GetEntityData() *types.CommonEntityData {
    revertiveTimer.EntityData.YFilter = revertiveTimer.YFilter
    revertiveTimer.EntityData.YangName = "revertive-timer"
    revertiveTimer.EntityData.BundleName = "cisco_ios_xr"
    revertiveTimer.EntityData.ParentYangName = "session-redundancy"
    revertiveTimer.EntityData.SegmentPath = "revertive-timer"
    revertiveTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-serg-cfg:session-redundancy/" + revertiveTimer.EntityData.SegmentPath
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

