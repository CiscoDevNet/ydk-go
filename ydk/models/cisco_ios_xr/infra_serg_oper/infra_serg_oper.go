// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-serg package operational data.
// 
// This module contains definitions
// for the following management objects:
//   session-redundancy-manager: Session Redundancy Manager
//     information
//   session-redundancy-agent: session redundancy agent
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_serg_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_serg_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-serg-oper session-redundancy-manager}", reflect.TypeOf(SessionRedundancyManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-serg-oper:session-redundancy-manager", reflect.TypeOf(SessionRedundancyManager{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-serg-oper session-redundancy-agent}", reflect.TypeOf(SessionRedundancyAgent{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-serg-oper:session-redundancy-agent", reflect.TypeOf(SessionRedundancyAgent{}))
}

// SergShowSessionError represents SERG Session Error Operation
type SergShowSessionError string

const (
    // Invalid Error
    SergShowSessionError_none SergShowSessionError = "none"

    // Hard Error
    SergShowSessionError_hard SergShowSessionError = "hard"

    // Soft Error
    SergShowSessionError_soft SergShowSessionError = "soft"
)

// SergShowSlaveMode represents SERG Slave Mode
type SergShowSlaveMode string

const (
    // Not Configured
    SergShowSlaveMode_none SergShowSlaveMode = "none"

    // Warm Modem
    SergShowSlaveMode_warm SergShowSlaveMode = "warm"

    // Hot Mode
    SergShowSlaveMode_hot SergShowSlaveMode = "hot"
)

// SergShowSoReason represents Session Redundancy Switchover Reason
type SergShowSoReason string

const (
    // SwitchOver for Internal Reason
    SergShowSoReason_internal SergShowSoReason = "internal"

    // SwitchOver for Admin
    SergShowSoReason_admin SergShowSoReason = "admin"

    // SwitchOver for Peer UP
    SergShowSoReason_peer_up SergShowSoReason = "peer-up"

    // SwitchOver for Peer Down
    SergShowSoReason_peer_down SergShowSoReason = "peer-down"

    // SwitchOver for Object Tracking status change
    SergShowSoReason_object_tracking_status_change SergShowSoReason = "object-tracking-status-change"

    // Unknown Switchover Reason
    SergShowSoReason_serg_show_so_reason_max SergShowSoReason = "serg-show-so-reason-max"
)

// SergShowMem represents SERG Memory Manager type
type SergShowMem string

const (
    // Standard type
    SergShowMem_standard SergShowMem = "standard"

    // Chunk type
    SergShowMem_chunk SergShowMem = "chunk"

    // EDM type
    SergShowMem_edm SergShowMem = "edm"

    // String type
    SergShowMem_string SergShowMem = "string"

    // Static type
    SergShowMem_static SergShowMem = "static"

    // Unknown type
    SergShowMem_unknown SergShowMem = "unknown"
)

// SergPeerStatus represents SERG Peer Status
type SergPeerStatus string

const (
    // Peer not configured
    SergPeerStatus_not_configured SergPeerStatus = "not-configured"

    // Peer initialization
    SergPeerStatus_initialize SergPeerStatus = "initialize"

    // Peer retry pending
    SergPeerStatus_retry SergPeerStatus = "retry"

    // Connection in Progress
    SergPeerStatus_connect SergPeerStatus = "connect"

    // Listening in Progress
    SergPeerStatus_listen SergPeerStatus = "listen"

    // Awaiting Registration from Peer
    SergPeerStatus_registration SergPeerStatus = "registration"

    // Peer cleanup in progress
    SergPeerStatus_cleanup SergPeerStatus = "cleanup"

    // Peer Connected
    SergPeerStatus_connected SergPeerStatus = "connected"

    // Peer Established
    SergPeerStatus_established SergPeerStatus = "established"
)

// SergShowImRole represents SERG Interface Management Role
type SergShowImRole string

const (
    // Not Determined
    SergShowImRole_none SergShowImRole = "none"

    // Master Role
    SergShowImRole_master SergShowImRole = "master"

    // Slave Role
    SergShowImRole_slave SergShowImRole = "slave"
)

// SergShowComp represents SERG Components
type SergShowComp string

const (
    // SERG Agent
    SergShowComp_serga SergShowComp = "serga"

    // IPv6ND
    SergShowComp_ipv6nd SergShowComp = "ipv6nd"

    // DHCPv6
    SergShowComp_dhcpv6 SergShowComp = "dhcpv6"
)

// SergShowSessionOperation represents SERG Session Operation
type SergShowSessionOperation string

const (
    // No Operation
    SergShowSessionOperation_none SergShowSessionOperation = "none"

    // SERG Session Update Operation
    SergShowSessionOperation_update SergShowSessionOperation = "update"

    // SERG Session Delete Operation
    SergShowSessionOperation_delete SergShowSessionOperation = "delete"

    // SERG Session In Sync
    SergShowSessionOperation_in_sync SergShowSessionOperation = "in-sync"
)

// SergShowRole represents SERG Role
type SergShowRole string

const (
    // Not Configured
    SergShowRole_none SergShowRole = "none"

    // Master Role
    SergShowRole_master SergShowRole = "master"

    // Slave Role
    SergShowRole_slave SergShowRole = "slave"
)

// SessionRedundancyManager
// Session Redundancy Manager information
type SessionRedundancyManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber Redundancy Manager interface table.
    Interfaces SessionRedundancyManager_Interfaces

    // Session Redundancy Manager group table.
    Groups SessionRedundancyManager_Groups

    // Session redundancy manager summary.
    Summary SessionRedundancyManager_Summary
}

func (sessionRedundancyManager *SessionRedundancyManager) GetFilter() yfilter.YFilter { return sessionRedundancyManager.YFilter }

func (sessionRedundancyManager *SessionRedundancyManager) SetFilter(yf yfilter.YFilter) { sessionRedundancyManager.YFilter = yf }

func (sessionRedundancyManager *SessionRedundancyManager) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    if yname == "groups" { return "Groups" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (sessionRedundancyManager *SessionRedundancyManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-serg-oper:session-redundancy-manager"
}

func (sessionRedundancyManager *SessionRedundancyManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &sessionRedundancyManager.Interfaces
    }
    if childYangName == "groups" {
        return &sessionRedundancyManager.Groups
    }
    if childYangName == "summary" {
        return &sessionRedundancyManager.Summary
    }
    return nil
}

func (sessionRedundancyManager *SessionRedundancyManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &sessionRedundancyManager.Interfaces
    children["groups"] = &sessionRedundancyManager.Groups
    children["summary"] = &sessionRedundancyManager.Summary
    return children
}

func (sessionRedundancyManager *SessionRedundancyManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessionRedundancyManager *SessionRedundancyManager) GetBundleName() string { return "cisco_ios_xr" }

func (sessionRedundancyManager *SessionRedundancyManager) GetYangName() string { return "session-redundancy-manager" }

func (sessionRedundancyManager *SessionRedundancyManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionRedundancyManager *SessionRedundancyManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionRedundancyManager *SessionRedundancyManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionRedundancyManager *SessionRedundancyManager) SetParent(parent types.Entity) { sessionRedundancyManager.parent = parent }

func (sessionRedundancyManager *SessionRedundancyManager) GetParent() types.Entity { return sessionRedundancyManager.parent }

func (sessionRedundancyManager *SessionRedundancyManager) GetParentYangName() string { return "Cisco-IOS-XR-infra-serg-oper" }

// SessionRedundancyManager_Interfaces
// Subscriber Redundancy Manager interface table
type SessionRedundancyManager_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface redundancy manager interface. The type is slice of
    // SessionRedundancyManager_Interfaces_Interface.
    Interface []SessionRedundancyManager_Interfaces_Interface
}

func (interfaces *SessionRedundancyManager_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *SessionRedundancyManager_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *SessionRedundancyManager_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *SessionRedundancyManager_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *SessionRedundancyManager_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyManager_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *SessionRedundancyManager_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *SessionRedundancyManager_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *SessionRedundancyManager_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *SessionRedundancyManager_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *SessionRedundancyManager_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *SessionRedundancyManager_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *SessionRedundancyManager_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *SessionRedundancyManager_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *SessionRedundancyManager_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *SessionRedundancyManager_Interfaces) GetParentYangName() string { return "session-redundancy-manager" }

// SessionRedundancyManager_Interfaces_Interface
// interface redundancy manager interface
type SessionRedundancyManager_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Interface interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Interface Mapping ID. The type is interface{} with range: 0..4294967295.
    InterfaceMappingId interface{}

    // Forward Referenced. The type is bool.
    ForwardReferenced interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // SERG Role. The type is SergShowImRole.
    Role interface{}
}

func (self *SessionRedundancyManager_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *SessionRedundancyManager_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *SessionRedundancyManager_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-mapping-id" { return "InterfaceMappingId" }
    if yname == "forward-referenced" { return "ForwardReferenced" }
    if yname == "group-id" { return "GroupId" }
    if yname == "role" { return "Role" }
    return ""
}

func (self *SessionRedundancyManager_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *SessionRedundancyManager_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *SessionRedundancyManager_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *SessionRedundancyManager_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = self.Interface
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-mapping-id"] = self.InterfaceMappingId
    leafs["forward-referenced"] = self.ForwardReferenced
    leafs["group-id"] = self.GroupId
    leafs["role"] = self.Role
    return leafs
}

func (self *SessionRedundancyManager_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *SessionRedundancyManager_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *SessionRedundancyManager_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *SessionRedundancyManager_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *SessionRedundancyManager_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *SessionRedundancyManager_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *SessionRedundancyManager_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *SessionRedundancyManager_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// SessionRedundancyManager_Groups
// Session Redundancy Manager group table
type SessionRedundancyManager_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session redundancy manager group. The type is slice of
    // SessionRedundancyManager_Groups_Group.
    Group []SessionRedundancyManager_Groups_Group
}

func (groups *SessionRedundancyManager_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *SessionRedundancyManager_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *SessionRedundancyManager_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *SessionRedundancyManager_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *SessionRedundancyManager_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyManager_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *SessionRedundancyManager_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *SessionRedundancyManager_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *SessionRedundancyManager_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *SessionRedundancyManager_Groups) GetYangName() string { return "groups" }

func (groups *SessionRedundancyManager_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *SessionRedundancyManager_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *SessionRedundancyManager_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *SessionRedundancyManager_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *SessionRedundancyManager_Groups) GetParent() types.Entity { return groups.parent }

func (groups *SessionRedundancyManager_Groups) GetParentYangName() string { return "session-redundancy-manager" }

// SessionRedundancyManager_Groups_Group
// Session redundancy manager group
type SessionRedundancyManager_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Group interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // Group Description. The type is string.
    Description interface{}

    // Disabled by Config. The type is bool.
    Disabled interface{}

    // SERG Role. The type is SergShowImRole.
    Role interface{}

    // Peer IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerIpv4Address interface{}

    // Peer IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerIpv6Address interface{}

    // Interface Count. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Preferred Role. The type is SergShowRole.
    PreferredRole interface{}

    // Slave Mode. The type is SergShowSlaveMode.
    SlaveMode interface{}

    // Object Tracking Status (Enabled/Disabled). The type is bool.
    ObjectTrackingStatus interface{}

    // Node Information. The type is string.
    NodeName interface{}
}

func (group *SessionRedundancyManager_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *SessionRedundancyManager_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *SessionRedundancyManager_Groups_Group) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    if yname == "group-id" { return "GroupId" }
    if yname == "description" { return "Description" }
    if yname == "disabled" { return "Disabled" }
    if yname == "role" { return "Role" }
    if yname == "peer-ipv4-address" { return "PeerIpv4Address" }
    if yname == "peer-ipv6-address" { return "PeerIpv6Address" }
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "preferred-role" { return "PreferredRole" }
    if yname == "slave-mode" { return "SlaveMode" }
    if yname == "object-tracking-status" { return "ObjectTrackingStatus" }
    if yname == "node-name" { return "NodeName" }
    return ""
}

func (group *SessionRedundancyManager_Groups_Group) GetSegmentPath() string {
    return "group" + "[group='" + fmt.Sprintf("%v", group.Group) + "']"
}

func (group *SessionRedundancyManager_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (group *SessionRedundancyManager_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (group *SessionRedundancyManager_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group"] = group.Group
    leafs["group-id"] = group.GroupId
    leafs["description"] = group.Description
    leafs["disabled"] = group.Disabled
    leafs["role"] = group.Role
    leafs["peer-ipv4-address"] = group.PeerIpv4Address
    leafs["peer-ipv6-address"] = group.PeerIpv6Address
    leafs["interface-count"] = group.InterfaceCount
    leafs["preferred-role"] = group.PreferredRole
    leafs["slave-mode"] = group.SlaveMode
    leafs["object-tracking-status"] = group.ObjectTrackingStatus
    leafs["node-name"] = group.NodeName
    return leafs
}

func (group *SessionRedundancyManager_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *SessionRedundancyManager_Groups_Group) GetYangName() string { return "group" }

func (group *SessionRedundancyManager_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *SessionRedundancyManager_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *SessionRedundancyManager_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *SessionRedundancyManager_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *SessionRedundancyManager_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *SessionRedundancyManager_Groups_Group) GetParentYangName() string { return "groups" }

// SessionRedundancyManager_Summary
// Session redundancy manager summary
type SessionRedundancyManager_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disabled by Config. The type is bool.
    Disabled interface{}

    // Process Active State. The type is bool.
    ActiveState interface{}

    // Preferred Role. The type is SergShowRole.
    PreferredRole interface{}

    // Slave Mode. The type is SergShowSlaveMode.
    SlaveMode interface{}

    // Switch Over Hold Time. The type is interface{} with range: 0..4294967295.
    HoldTimer interface{}

    // Source Interface Name. The type is string.
    SourceInterfaceName interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Source Interface IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceInterfaceIpv4Address interface{}

    // Source Interface IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceInterfaceIpv6Address interface{}

    // No. of Configured Groups. The type is interface{} with range:
    // 0..4294967295.
    GroupCount interface{}

    // No. of Disabled Groups by Config. The type is interface{} with range:
    // 0..4294967295.
    DisabledGroupCount interface{}

    // No. of Master/Active Groups. The type is interface{} with range:
    // 0..4294967295.
    MasterGroupCount interface{}

    // No. of Slave Groups. The type is interface{} with range: 0..4294967295.
    SlaveGroupCount interface{}

    // No. of Configured Interfaces. The type is interface{} with range:
    // 0..4294967295.
    InterfaceCount interface{}

    // No. of Master/Active Interfaces. The type is interface{} with range:
    // 0..4294967295.
    MasterInterfaceCount interface{}

    // No. of Slave Interfaces. The type is interface{} with range: 0..4294967295.
    SlaveInterfaceCount interface{}
}

func (summary *SessionRedundancyManager_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *SessionRedundancyManager_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *SessionRedundancyManager_Summary) GetGoName(yname string) string {
    if yname == "disabled" { return "Disabled" }
    if yname == "active-state" { return "ActiveState" }
    if yname == "preferred-role" { return "PreferredRole" }
    if yname == "slave-mode" { return "SlaveMode" }
    if yname == "hold-timer" { return "HoldTimer" }
    if yname == "source-interface-name" { return "SourceInterfaceName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source-interface-ipv4-address" { return "SourceInterfaceIpv4Address" }
    if yname == "source-interface-ipv6-address" { return "SourceInterfaceIpv6Address" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "disabled-group-count" { return "DisabledGroupCount" }
    if yname == "master-group-count" { return "MasterGroupCount" }
    if yname == "slave-group-count" { return "SlaveGroupCount" }
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "master-interface-count" { return "MasterInterfaceCount" }
    if yname == "slave-interface-count" { return "SlaveInterfaceCount" }
    return ""
}

func (summary *SessionRedundancyManager_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *SessionRedundancyManager_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *SessionRedundancyManager_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *SessionRedundancyManager_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disabled"] = summary.Disabled
    leafs["active-state"] = summary.ActiveState
    leafs["preferred-role"] = summary.PreferredRole
    leafs["slave-mode"] = summary.SlaveMode
    leafs["hold-timer"] = summary.HoldTimer
    leafs["source-interface-name"] = summary.SourceInterfaceName
    leafs["vrf-name"] = summary.VrfName
    leafs["source-interface-ipv4-address"] = summary.SourceInterfaceIpv4Address
    leafs["source-interface-ipv6-address"] = summary.SourceInterfaceIpv6Address
    leafs["group-count"] = summary.GroupCount
    leafs["disabled-group-count"] = summary.DisabledGroupCount
    leafs["master-group-count"] = summary.MasterGroupCount
    leafs["slave-group-count"] = summary.SlaveGroupCount
    leafs["interface-count"] = summary.InterfaceCount
    leafs["master-interface-count"] = summary.MasterInterfaceCount
    leafs["slave-interface-count"] = summary.SlaveInterfaceCount
    return leafs
}

func (summary *SessionRedundancyManager_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *SessionRedundancyManager_Summary) GetYangName() string { return "summary" }

func (summary *SessionRedundancyManager_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *SessionRedundancyManager_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *SessionRedundancyManager_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *SessionRedundancyManager_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *SessionRedundancyManager_Summary) GetParent() types.Entity { return summary.parent }

func (summary *SessionRedundancyManager_Summary) GetParentYangName() string { return "session-redundancy-manager" }

// SessionRedundancyAgent
// session redundancy agent
type SessionRedundancyAgent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes for which session data is collected.
    Nodes SessionRedundancyAgent_Nodes
}

func (sessionRedundancyAgent *SessionRedundancyAgent) GetFilter() yfilter.YFilter { return sessionRedundancyAgent.YFilter }

func (sessionRedundancyAgent *SessionRedundancyAgent) SetFilter(yf yfilter.YFilter) { sessionRedundancyAgent.YFilter = yf }

func (sessionRedundancyAgent *SessionRedundancyAgent) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (sessionRedundancyAgent *SessionRedundancyAgent) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-serg-oper:session-redundancy-agent"
}

func (sessionRedundancyAgent *SessionRedundancyAgent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &sessionRedundancyAgent.Nodes
    }
    return nil
}

func (sessionRedundancyAgent *SessionRedundancyAgent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &sessionRedundancyAgent.Nodes
    return children
}

func (sessionRedundancyAgent *SessionRedundancyAgent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessionRedundancyAgent *SessionRedundancyAgent) GetBundleName() string { return "cisco_ios_xr" }

func (sessionRedundancyAgent *SessionRedundancyAgent) GetYangName() string { return "session-redundancy-agent" }

func (sessionRedundancyAgent *SessionRedundancyAgent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionRedundancyAgent *SessionRedundancyAgent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionRedundancyAgent *SessionRedundancyAgent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionRedundancyAgent *SessionRedundancyAgent) SetParent(parent types.Entity) { sessionRedundancyAgent.parent = parent }

func (sessionRedundancyAgent *SessionRedundancyAgent) GetParent() types.Entity { return sessionRedundancyAgent.parent }

func (sessionRedundancyAgent *SessionRedundancyAgent) GetParentYangName() string { return "Cisco-IOS-XR-infra-serg-oper" }

// SessionRedundancyAgent_Nodes
// List of nodes for which session data is
// collected
type SessionRedundancyAgent_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session data for a particular node. The type is slice of
    // SessionRedundancyAgent_Nodes_Node.
    Node []SessionRedundancyAgent_Nodes_Node
}

func (nodes *SessionRedundancyAgent_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *SessionRedundancyAgent_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *SessionRedundancyAgent_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *SessionRedundancyAgent_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *SessionRedundancyAgent_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *SessionRedundancyAgent_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *SessionRedundancyAgent_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *SessionRedundancyAgent_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *SessionRedundancyAgent_Nodes) GetYangName() string { return "nodes" }

func (nodes *SessionRedundancyAgent_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *SessionRedundancyAgent_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *SessionRedundancyAgent_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *SessionRedundancyAgent_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *SessionRedundancyAgent_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *SessionRedundancyAgent_Nodes) GetParentYangName() string { return "session-redundancy-agent" }

// SessionRedundancyAgent_Nodes_Node
// Session data for a particular node
type SessionRedundancyAgent_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Data for particular subscriber group session.
    GroupIdXr SessionRedundancyAgent_Nodes_Node_GroupIdXr

    // Stats Client.
    ClientIds SessionRedundancyAgent_Nodes_Node_ClientIds

    // Show Memory.
    Memory SessionRedundancyAgent_Nodes_Node_Memory

    // Data for particular subscriber group .
    GroupIds SessionRedundancyAgent_Nodes_Node_GroupIds

    // List of interfaces.
    Interfaces SessionRedundancyAgent_Nodes_Node_Interfaces

    // Stats Global.
    StatsGlobal SessionRedundancyAgent_Nodes_Node_StatsGlobal

    // Session data for a particular node.
    GroupSummaries SessionRedundancyAgent_Nodes_Node_GroupSummaries
}

func (node *SessionRedundancyAgent_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *SessionRedundancyAgent_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *SessionRedundancyAgent_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "group-id-xr" { return "GroupIdXr" }
    if yname == "client-ids" { return "ClientIds" }
    if yname == "memory" { return "Memory" }
    if yname == "group-ids" { return "GroupIds" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "stats-global" { return "StatsGlobal" }
    if yname == "group-summaries" { return "GroupSummaries" }
    return ""
}

func (node *SessionRedundancyAgent_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *SessionRedundancyAgent_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-id-xr" {
        return &node.GroupIdXr
    }
    if childYangName == "client-ids" {
        return &node.ClientIds
    }
    if childYangName == "memory" {
        return &node.Memory
    }
    if childYangName == "group-ids" {
        return &node.GroupIds
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "stats-global" {
        return &node.StatsGlobal
    }
    if childYangName == "group-summaries" {
        return &node.GroupSummaries
    }
    return nil
}

func (node *SessionRedundancyAgent_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["group-id-xr"] = &node.GroupIdXr
    children["client-ids"] = &node.ClientIds
    children["memory"] = &node.Memory
    children["group-ids"] = &node.GroupIds
    children["interfaces"] = &node.Interfaces
    children["stats-global"] = &node.StatsGlobal
    children["group-summaries"] = &node.GroupSummaries
    return children
}

func (node *SessionRedundancyAgent_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *SessionRedundancyAgent_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *SessionRedundancyAgent_Nodes_Node) GetYangName() string { return "node" }

func (node *SessionRedundancyAgent_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *SessionRedundancyAgent_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *SessionRedundancyAgent_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *SessionRedundancyAgent_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *SessionRedundancyAgent_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *SessionRedundancyAgent_Nodes_Node) GetParentYangName() string { return "nodes" }

// SessionRedundancyAgent_Nodes_Node_GroupIdXr
// Data for particular subscriber group session
type SessionRedundancyAgent_Nodes_Node_GroupIdXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group id for subscriber group session. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId.
    GroupId []SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId
}

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetFilter() yfilter.YFilter { return groupIdXr.YFilter }

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) SetFilter(yf yfilter.YFilter) { groupIdXr.YFilter = yf }

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    return ""
}

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetSegmentPath() string {
    return "group-id-xr"
}

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-id" {
        for _, c := range groupIdXr.GroupId {
            if groupIdXr.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId{}
        groupIdXr.GroupId = append(groupIdXr.GroupId, child)
        return &groupIdXr.GroupId[len(groupIdXr.GroupId)-1]
    }
    return nil
}

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupIdXr.GroupId {
        children[groupIdXr.GroupId[i].GetSegmentPath()] = &groupIdXr.GroupId[i]
    }
    return children
}

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetBundleName() string { return "cisco_ios_xr" }

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetYangName() string { return "group-id-xr" }

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) SetParent(parent types.Entity) { groupIdXr.parent = parent }

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetParent() types.Entity { return groupIdXr.parent }

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetParentYangName() string { return "node" }

// SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId
// Group id for subscriber group session
type SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. GroupId. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    GroupId interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupIdXr interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Key index. The type is string.
    KeyIndex interface{}

    // Master Role is Set. The type is bool.
    RoleMaster interface{}

    // Negative Acknowledgement Update Flag is Set. The type is bool.
    NegativeAcknowledgementUpdateAll interface{}

    // More Session Information. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation.
    SessionDetailedInformation []SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation

    // Session Synchroniation Error Information. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation.
    SessionSyncErrorInformation []SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetFilter() yfilter.YFilter { return groupId.YFilter }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) SetFilter(yf yfilter.YFilter) { groupId.YFilter = yf }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "group-id-xr" { return "GroupIdXr" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "key-index" { return "KeyIndex" }
    if yname == "role-master" { return "RoleMaster" }
    if yname == "negative-acknowledgement-update-all" { return "NegativeAcknowledgementUpdateAll" }
    if yname == "session-detailed-information" { return "SessionDetailedInformation" }
    if yname == "session-sync-error-information" { return "SessionSyncErrorInformation" }
    return ""
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetSegmentPath() string {
    return "group-id" + "[group-id='" + fmt.Sprintf("%v", groupId.GroupId) + "']"
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-detailed-information" {
        for _, c := range groupId.SessionDetailedInformation {
            if groupId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation{}
        groupId.SessionDetailedInformation = append(groupId.SessionDetailedInformation, child)
        return &groupId.SessionDetailedInformation[len(groupId.SessionDetailedInformation)-1]
    }
    if childYangName == "session-sync-error-information" {
        for _, c := range groupId.SessionSyncErrorInformation {
            if groupId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation{}
        groupId.SessionSyncErrorInformation = append(groupId.SessionSyncErrorInformation, child)
        return &groupId.SessionSyncErrorInformation[len(groupId.SessionSyncErrorInformation)-1]
    }
    return nil
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupId.SessionDetailedInformation {
        children[groupId.SessionDetailedInformation[i].GetSegmentPath()] = &groupId.SessionDetailedInformation[i]
    }
    for i := range groupId.SessionSyncErrorInformation {
        children[groupId.SessionSyncErrorInformation[i].GetSegmentPath()] = &groupId.SessionSyncErrorInformation[i]
    }
    return children
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = groupId.GroupId
    leafs["group-id-xr"] = groupId.GroupIdXr
    leafs["interface-name"] = groupId.InterfaceName
    leafs["key-index"] = groupId.KeyIndex
    leafs["role-master"] = groupId.RoleMaster
    leafs["negative-acknowledgement-update-all"] = groupId.NegativeAcknowledgementUpdateAll
    return leafs
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetBundleName() string { return "cisco_ios_xr" }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetYangName() string { return "group-id" }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) SetParent(parent types.Entity) { groupId.parent = parent }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetParent() types.Entity { return groupId.parent }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetParentYangName() string { return "group-id-xr" }

// SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation
// More Session Information
type SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Component. The type is SergShowComp.
    Component interface{}

    // Operation Code. The type is SergShowSessionOperation.
    Operation interface{}

    // Tx List Queue Failed. The type is bool.
    TxListQueueFail interface{}

    // Marked For Sweeping. The type is bool.
    MarkedForSweeping interface{}

    // Marked For Cleanup. The type is bool.
    MarkedForCleanup interface{}
}

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetFilter() yfilter.YFilter { return sessionDetailedInformation.YFilter }

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) SetFilter(yf yfilter.YFilter) { sessionDetailedInformation.YFilter = yf }

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetGoName(yname string) string {
    if yname == "component" { return "Component" }
    if yname == "operation" { return "Operation" }
    if yname == "tx-list-queue-fail" { return "TxListQueueFail" }
    if yname == "marked-for-sweeping" { return "MarkedForSweeping" }
    if yname == "marked-for-cleanup" { return "MarkedForCleanup" }
    return ""
}

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetSegmentPath() string {
    return "session-detailed-information"
}

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["component"] = sessionDetailedInformation.Component
    leafs["operation"] = sessionDetailedInformation.Operation
    leafs["tx-list-queue-fail"] = sessionDetailedInformation.TxListQueueFail
    leafs["marked-for-sweeping"] = sessionDetailedInformation.MarkedForSweeping
    leafs["marked-for-cleanup"] = sessionDetailedInformation.MarkedForCleanup
    return leafs
}

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetBundleName() string { return "cisco_ios_xr" }

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetYangName() string { return "session-detailed-information" }

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) SetParent(parent types.Entity) { sessionDetailedInformation.parent = parent }

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetParent() types.Entity { return sessionDetailedInformation.parent }

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetParentYangName() string { return "group-id" }

// SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation
// Session Synchroniation Error Information
type SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // No. of Errors occured during Synchronization. The type is interface{} with
    // range: 0..65535.
    SyncErrorCount interface{}

    // Last Error Code. The type is interface{} with range: 0..4294967295.
    LastErrorCode interface{}

    // Last Error Type. The type is SergShowSessionError.
    LastErrorType interface{}
}

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetFilter() yfilter.YFilter { return sessionSyncErrorInformation.YFilter }

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) SetFilter(yf yfilter.YFilter) { sessionSyncErrorInformation.YFilter = yf }

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetGoName(yname string) string {
    if yname == "sync-error-count" { return "SyncErrorCount" }
    if yname == "last-error-code" { return "LastErrorCode" }
    if yname == "last-error-type" { return "LastErrorType" }
    return ""
}

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetSegmentPath() string {
    return "session-sync-error-information"
}

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sync-error-count"] = sessionSyncErrorInformation.SyncErrorCount
    leafs["last-error-code"] = sessionSyncErrorInformation.LastErrorCode
    leafs["last-error-type"] = sessionSyncErrorInformation.LastErrorType
    return leafs
}

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetBundleName() string { return "cisco_ios_xr" }

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetYangName() string { return "session-sync-error-information" }

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) SetParent(parent types.Entity) { sessionSyncErrorInformation.parent = parent }

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetParent() types.Entity { return sessionSyncErrorInformation.parent }

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetParentYangName() string { return "group-id" }

// SessionRedundancyAgent_Nodes_Node_ClientIds
// Stats Client
type SessionRedundancyAgent_Nodes_Node_ClientIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify stats client. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId.
    ClientId []SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId
}

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetFilter() yfilter.YFilter { return clientIds.YFilter }

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) SetFilter(yf yfilter.YFilter) { clientIds.YFilter = yf }

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    return ""
}

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetSegmentPath() string {
    return "client-ids"
}

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-id" {
        for _, c := range clientIds.ClientId {
            if clientIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId{}
        clientIds.ClientId = append(clientIds.ClientId, child)
        return &clientIds.ClientId[len(clientIds.ClientId)-1]
    }
    return nil
}

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clientIds.ClientId {
        children[clientIds.ClientId[i].GetSegmentPath()] = &clientIds.ClientId[i]
    }
    return children
}

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetBundleName() string { return "cisco_ios_xr" }

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetYangName() string { return "client-ids" }

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) SetParent(parent types.Entity) { clientIds.parent = parent }

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetParent() types.Entity { return clientIds.parent }

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetParentYangName() string { return "node" }

// SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId
// Specify stats client
type SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client Id. The type is interface{} with range:
    // -2147483648..2147483647.
    StatsClientId interface{}

    // TxListStartOfDownloadAddOk. The type is interface{} with range:
    // 0..4294967295.
    TxListStartOfDownloadAddOk interface{}

    // TxListStartOfDownloadAddNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListStartOfDownloadAddNotOk interface{}

    // TxListEndOfDownloadAddOk. The type is interface{} with range:
    // 0..4294967295.
    TxListEndOfDownloadAddOk interface{}

    // TxListEndOfDownloadAddNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListEndOfDownloadAddNotOk interface{}

    // TxListEndOfMessageAddOk. The type is interface{} with range: 0..4294967295.
    TxListEndOfMessageAddOk interface{}

    // TxListEndOfMessageAddNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListEndOfMessageAddNotOk interface{}

    // TxListClearAllAddOk. The type is interface{} with range: 0..4294967295.
    TxListClearAllAddOk interface{}

    // TxListClearAllAddNotOk. The type is interface{} with range: 0..4294967295.
    TxListClearAllAddNotOk interface{}

    // TxListClearSelectedAddOk. The type is interface{} with range:
    // 0..4294967295.
    TxListClearSelectedAddOk interface{}

    // TxListClearSelectedAddNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListClearSelectedAddNotOk interface{}

    // TxListReplayAllAddOk. The type is interface{} with range: 0..4294967295.
    TxListReplayAllAddOk interface{}

    // TxListReplayAllAddNotOk. The type is interface{} with range: 0..4294967295.
    TxListReplayAllAddNotOk interface{}

    // TxListReplaySelectedAddOk. The type is interface{} with range:
    // 0..4294967295.
    TxListReplaySelectedAddOk interface{}

    // TxListReplaySelectedAddNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListReplaySelectedAddNotOk interface{}

    // TxListSessionSessionAddOk. The type is interface{} with range:
    // 0..4294967295.
    TxListSessionSessionAddOk interface{}

    // TxListSessionSessionAddNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListSessionSessionAddNotOk interface{}

    // TxListSessionSessionUpdateOk. The type is interface{} with range:
    // 0..4294967295.
    TxListSessionSessionUpdateOk interface{}

    // TxListSessionSessionUpdateNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListSessionSessionUpdateNotOk interface{}

    // TxListSessionSessionDelete. The type is interface{} with range:
    // 0..4294967295.
    TxListSessionSessionDelete interface{}

    // CleanCallBack. The type is interface{} with range: 0..4294967295.
    CleanCallBack interface{}

    // TxListEncodeSessionSessionOk. The type is interface{} with range:
    // 0..4294967295.
    TxListEncodeSessionSessionOk interface{}

    // TxListEncodeSessionSessionPartialWrite. The type is interface{} with range:
    // 0..4294967295.
    TxListEncodeSessionSessionPartialWrite interface{}

    // LastReplayAllCount. The type is interface{} with range: 0..4294967295.
    LastReplayAllCount interface{}

    // TxListReceiveCommandReplayAll. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveCommandReplayAll interface{}

    // TxListReceiveCommandReplaySelected. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveCommandReplaySelected interface{}

    // TxListReceiveSessionSessionDeleteValid. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionDeleteValid interface{}

    // TxListReceiveSessionSessionDeleteInValid. The type is interface{} with
    // range: 0..4294967295.
    TxListReceiveSessionSessionDeleteInvalid interface{}

    // TxListReceiveSessionSessionUpdateValid. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionUpdateValid interface{}

    // TxListReceiveSessionSessionUpdateInValid. The type is interface{} with
    // range: 0..4294967295.
    TxListReceiveSessionSessionUpdateInvalid interface{}

    // TxListReceiveSessionSessionSODAll. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionSodAll interface{}

    // TxListReceiveSessionSessionSODSelected. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionSodSelected interface{}

    // TxListReceiveSessionSessionEODAll. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionEodAll interface{}

    // TxListReceiveSessionSessionEODSelected. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionEodSelected interface{}

    // TxListReceiveSessionSessionEOMS. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionEoms interface{}

    // TxListReceiveSessionSessionClearAll. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionClearAll interface{}

    // TxListReceiveSessionSessionClearSelected. The type is interface{} with
    // range: 0..4294967295.
    TxListReceiveSessionSessionClearSelected interface{}

    // TxListReceiveSessionSessionNegAck. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionNegAck interface{}

    // TxListReceiveSessionSessionNegAckNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListReceiveSessionSessionNegAckNotOk interface{}

    // TxListClientRegistrationOk. The type is interface{} with range:
    // 0..4294967295.
    TxListClientRegistrationOk interface{}

    // TxListClientRegistrationNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListClientRegistrationNotOk interface{}

    // TxListClientDeRegistration. The type is interface{} with range:
    // 0..4294967295.
    TxListClientDeRegistration interface{}

    // TxListClientConnectionDown. The type is interface{} with range:
    // 0..4294967295.
    TxListClientConnectionDown interface{}

    // TxListClientCleanup. The type is interface{} with range: 0..4294967295.
    TxListClientCleanup interface{}

    // TxListActiveOk. The type is interface{} with range: 0..4294967295.
    TxListActiveOk interface{}

    // TxListActiveNotOk. The type is interface{} with range: 0..4294967295.
    TxListActiveNotOk interface{}

    // TxListDeActiveOk. The type is interface{} with range: 0..4294967295.
    TxListDeActiveOk interface{}

    // TxListDeActiveNotOk. The type is interface{} with range: 0..4294967295.
    TxListDeActiveNotOk interface{}
}

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetFilter() yfilter.YFilter { return clientId.YFilter }

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) SetFilter(yf yfilter.YFilter) { clientId.YFilter = yf }

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetGoName(yname string) string {
    if yname == "stats-client-id" { return "StatsClientId" }
    if yname == "tx-list-start-of-download-add-ok" { return "TxListStartOfDownloadAddOk" }
    if yname == "tx-list-start-of-download-add-not-ok" { return "TxListStartOfDownloadAddNotOk" }
    if yname == "tx-list-end-of-download-add-ok" { return "TxListEndOfDownloadAddOk" }
    if yname == "tx-list-end-of-download-add-not-ok" { return "TxListEndOfDownloadAddNotOk" }
    if yname == "tx-list-end-of-message-add-ok" { return "TxListEndOfMessageAddOk" }
    if yname == "tx-list-end-of-message-add-not-ok" { return "TxListEndOfMessageAddNotOk" }
    if yname == "tx-list-clear-all-add-ok" { return "TxListClearAllAddOk" }
    if yname == "tx-list-clear-all-add-not-ok" { return "TxListClearAllAddNotOk" }
    if yname == "tx-list-clear-selected-add-ok" { return "TxListClearSelectedAddOk" }
    if yname == "tx-list-clear-selected-add-not-ok" { return "TxListClearSelectedAddNotOk" }
    if yname == "tx-list-replay-all-add-ok" { return "TxListReplayAllAddOk" }
    if yname == "tx-list-replay-all-add-not-ok" { return "TxListReplayAllAddNotOk" }
    if yname == "tx-list-replay-selected-add-ok" { return "TxListReplaySelectedAddOk" }
    if yname == "tx-list-replay-selected-add-not-ok" { return "TxListReplaySelectedAddNotOk" }
    if yname == "tx-list-session-session-add-ok" { return "TxListSessionSessionAddOk" }
    if yname == "tx-list-session-session-add-not-ok" { return "TxListSessionSessionAddNotOk" }
    if yname == "tx-list-session-session-update-ok" { return "TxListSessionSessionUpdateOk" }
    if yname == "tx-list-session-session-update-not-ok" { return "TxListSessionSessionUpdateNotOk" }
    if yname == "tx-list-session-session-delete" { return "TxListSessionSessionDelete" }
    if yname == "clean-call-back" { return "CleanCallBack" }
    if yname == "tx-list-encode-session-session-ok" { return "TxListEncodeSessionSessionOk" }
    if yname == "tx-list-encode-session-session-partial-write" { return "TxListEncodeSessionSessionPartialWrite" }
    if yname == "last-replay-all-count" { return "LastReplayAllCount" }
    if yname == "tx-list-receive-command-replay-all" { return "TxListReceiveCommandReplayAll" }
    if yname == "tx-list-receive-command-replay-selected" { return "TxListReceiveCommandReplaySelected" }
    if yname == "tx-list-receive-session-session-delete-valid" { return "TxListReceiveSessionSessionDeleteValid" }
    if yname == "tx-list-receive-session-session-delete-invalid" { return "TxListReceiveSessionSessionDeleteInvalid" }
    if yname == "tx-list-receive-session-session-update-valid" { return "TxListReceiveSessionSessionUpdateValid" }
    if yname == "tx-list-receive-session-session-update-invalid" { return "TxListReceiveSessionSessionUpdateInvalid" }
    if yname == "tx-list-receive-session-session-sod-all" { return "TxListReceiveSessionSessionSodAll" }
    if yname == "tx-list-receive-session-session-sod-selected" { return "TxListReceiveSessionSessionSodSelected" }
    if yname == "tx-list-receive-session-session-eod-all" { return "TxListReceiveSessionSessionEodAll" }
    if yname == "tx-list-receive-session-session-eod-selected" { return "TxListReceiveSessionSessionEodSelected" }
    if yname == "tx-list-receive-session-session-eoms" { return "TxListReceiveSessionSessionEoms" }
    if yname == "tx-list-receive-session-session-clear-all" { return "TxListReceiveSessionSessionClearAll" }
    if yname == "tx-list-receive-session-session-clear-selected" { return "TxListReceiveSessionSessionClearSelected" }
    if yname == "tx-list-receive-session-session-neg-ack" { return "TxListReceiveSessionSessionNegAck" }
    if yname == "tx-list-receive-session-session-neg-ack-not-ok" { return "TxListReceiveSessionSessionNegAckNotOk" }
    if yname == "tx-list-client-registration-ok" { return "TxListClientRegistrationOk" }
    if yname == "tx-list-client-registration-not-ok" { return "TxListClientRegistrationNotOk" }
    if yname == "tx-list-client-de-registration" { return "TxListClientDeRegistration" }
    if yname == "tx-list-client-connection-down" { return "TxListClientConnectionDown" }
    if yname == "tx-list-client-cleanup" { return "TxListClientCleanup" }
    if yname == "tx-list-active-ok" { return "TxListActiveOk" }
    if yname == "tx-list-active-not-ok" { return "TxListActiveNotOk" }
    if yname == "tx-list-de-active-ok" { return "TxListDeActiveOk" }
    if yname == "tx-list-de-active-not-ok" { return "TxListDeActiveNotOk" }
    return ""
}

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetSegmentPath() string {
    return "client-id" + "[stats-client-id='" + fmt.Sprintf("%v", clientId.StatsClientId) + "']"
}

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stats-client-id"] = clientId.StatsClientId
    leafs["tx-list-start-of-download-add-ok"] = clientId.TxListStartOfDownloadAddOk
    leafs["tx-list-start-of-download-add-not-ok"] = clientId.TxListStartOfDownloadAddNotOk
    leafs["tx-list-end-of-download-add-ok"] = clientId.TxListEndOfDownloadAddOk
    leafs["tx-list-end-of-download-add-not-ok"] = clientId.TxListEndOfDownloadAddNotOk
    leafs["tx-list-end-of-message-add-ok"] = clientId.TxListEndOfMessageAddOk
    leafs["tx-list-end-of-message-add-not-ok"] = clientId.TxListEndOfMessageAddNotOk
    leafs["tx-list-clear-all-add-ok"] = clientId.TxListClearAllAddOk
    leafs["tx-list-clear-all-add-not-ok"] = clientId.TxListClearAllAddNotOk
    leafs["tx-list-clear-selected-add-ok"] = clientId.TxListClearSelectedAddOk
    leafs["tx-list-clear-selected-add-not-ok"] = clientId.TxListClearSelectedAddNotOk
    leafs["tx-list-replay-all-add-ok"] = clientId.TxListReplayAllAddOk
    leafs["tx-list-replay-all-add-not-ok"] = clientId.TxListReplayAllAddNotOk
    leafs["tx-list-replay-selected-add-ok"] = clientId.TxListReplaySelectedAddOk
    leafs["tx-list-replay-selected-add-not-ok"] = clientId.TxListReplaySelectedAddNotOk
    leafs["tx-list-session-session-add-ok"] = clientId.TxListSessionSessionAddOk
    leafs["tx-list-session-session-add-not-ok"] = clientId.TxListSessionSessionAddNotOk
    leafs["tx-list-session-session-update-ok"] = clientId.TxListSessionSessionUpdateOk
    leafs["tx-list-session-session-update-not-ok"] = clientId.TxListSessionSessionUpdateNotOk
    leafs["tx-list-session-session-delete"] = clientId.TxListSessionSessionDelete
    leafs["clean-call-back"] = clientId.CleanCallBack
    leafs["tx-list-encode-session-session-ok"] = clientId.TxListEncodeSessionSessionOk
    leafs["tx-list-encode-session-session-partial-write"] = clientId.TxListEncodeSessionSessionPartialWrite
    leafs["last-replay-all-count"] = clientId.LastReplayAllCount
    leafs["tx-list-receive-command-replay-all"] = clientId.TxListReceiveCommandReplayAll
    leafs["tx-list-receive-command-replay-selected"] = clientId.TxListReceiveCommandReplaySelected
    leafs["tx-list-receive-session-session-delete-valid"] = clientId.TxListReceiveSessionSessionDeleteValid
    leafs["tx-list-receive-session-session-delete-invalid"] = clientId.TxListReceiveSessionSessionDeleteInvalid
    leafs["tx-list-receive-session-session-update-valid"] = clientId.TxListReceiveSessionSessionUpdateValid
    leafs["tx-list-receive-session-session-update-invalid"] = clientId.TxListReceiveSessionSessionUpdateInvalid
    leafs["tx-list-receive-session-session-sod-all"] = clientId.TxListReceiveSessionSessionSodAll
    leafs["tx-list-receive-session-session-sod-selected"] = clientId.TxListReceiveSessionSessionSodSelected
    leafs["tx-list-receive-session-session-eod-all"] = clientId.TxListReceiveSessionSessionEodAll
    leafs["tx-list-receive-session-session-eod-selected"] = clientId.TxListReceiveSessionSessionEodSelected
    leafs["tx-list-receive-session-session-eoms"] = clientId.TxListReceiveSessionSessionEoms
    leafs["tx-list-receive-session-session-clear-all"] = clientId.TxListReceiveSessionSessionClearAll
    leafs["tx-list-receive-session-session-clear-selected"] = clientId.TxListReceiveSessionSessionClearSelected
    leafs["tx-list-receive-session-session-neg-ack"] = clientId.TxListReceiveSessionSessionNegAck
    leafs["tx-list-receive-session-session-neg-ack-not-ok"] = clientId.TxListReceiveSessionSessionNegAckNotOk
    leafs["tx-list-client-registration-ok"] = clientId.TxListClientRegistrationOk
    leafs["tx-list-client-registration-not-ok"] = clientId.TxListClientRegistrationNotOk
    leafs["tx-list-client-de-registration"] = clientId.TxListClientDeRegistration
    leafs["tx-list-client-connection-down"] = clientId.TxListClientConnectionDown
    leafs["tx-list-client-cleanup"] = clientId.TxListClientCleanup
    leafs["tx-list-active-ok"] = clientId.TxListActiveOk
    leafs["tx-list-active-not-ok"] = clientId.TxListActiveNotOk
    leafs["tx-list-de-active-ok"] = clientId.TxListDeActiveOk
    leafs["tx-list-de-active-not-ok"] = clientId.TxListDeActiveNotOk
    return leafs
}

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetBundleName() string { return "cisco_ios_xr" }

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetYangName() string { return "client-id" }

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) SetParent(parent types.Entity) { clientId.parent = parent }

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetParent() types.Entity { return clientId.parent }

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetParentYangName() string { return "client-ids" }

// SessionRedundancyAgent_Nodes_Node_Memory
// Show Memory
type SessionRedundancyAgent_Nodes_Node_Memory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory Info. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo.
    MemoryInfo []SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo

    // EDM Memory Info. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo.
    EdmMemoryInfo []SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo

    // String Memory Info. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo.
    StringMemoryInfo []SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo
}

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetFilter() yfilter.YFilter { return memory.YFilter }

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) SetFilter(yf yfilter.YFilter) { memory.YFilter = yf }

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetGoName(yname string) string {
    if yname == "memory-info" { return "MemoryInfo" }
    if yname == "edm-memory-info" { return "EdmMemoryInfo" }
    if yname == "string-memory-info" { return "StringMemoryInfo" }
    return ""
}

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetSegmentPath() string {
    return "memory"
}

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "memory-info" {
        for _, c := range memory.MemoryInfo {
            if memory.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo{}
        memory.MemoryInfo = append(memory.MemoryInfo, child)
        return &memory.MemoryInfo[len(memory.MemoryInfo)-1]
    }
    if childYangName == "edm-memory-info" {
        for _, c := range memory.EdmMemoryInfo {
            if memory.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo{}
        memory.EdmMemoryInfo = append(memory.EdmMemoryInfo, child)
        return &memory.EdmMemoryInfo[len(memory.EdmMemoryInfo)-1]
    }
    if childYangName == "string-memory-info" {
        for _, c := range memory.StringMemoryInfo {
            if memory.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo{}
        memory.StringMemoryInfo = append(memory.StringMemoryInfo, child)
        return &memory.StringMemoryInfo[len(memory.StringMemoryInfo)-1]
    }
    return nil
}

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range memory.MemoryInfo {
        children[memory.MemoryInfo[i].GetSegmentPath()] = &memory.MemoryInfo[i]
    }
    for i := range memory.EdmMemoryInfo {
        children[memory.EdmMemoryInfo[i].GetSegmentPath()] = &memory.EdmMemoryInfo[i]
    }
    for i := range memory.StringMemoryInfo {
        children[memory.StringMemoryInfo[i].GetSegmentPath()] = &memory.StringMemoryInfo[i]
    }
    return children
}

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetBundleName() string { return "cisco_ios_xr" }

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetYangName() string { return "memory" }

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) SetParent(parent types.Entity) { memory.parent = parent }

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetParent() types.Entity { return memory.parent }

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetParentYangName() string { return "node" }

// SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo
// Memory Info
type SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Structure Name. The type is string.
    StructureName interface{}

    // Size of the datastructure. The type is interface{} with range:
    // 0..4294967295.
    Size interface{}

    // Current Count. The type is interface{} with range: 0..4294967295.
    CurrentCount interface{}

    // Allocation Fails. The type is interface{} with range: 0..4294967295.
    AllocFails interface{}

    // Allocated count. The type is interface{} with range: 0..4294967295.
    AllocCount interface{}

    // Freed Count. The type is interface{} with range: 0..4294967295.
    FreedCount interface{}

    // Memory Type. The type is SergShowMem.
    MemoryType interface{}
}

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetFilter() yfilter.YFilter { return memoryInfo.YFilter }

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) SetFilter(yf yfilter.YFilter) { memoryInfo.YFilter = yf }

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetGoName(yname string) string {
    if yname == "structure-name" { return "StructureName" }
    if yname == "size" { return "Size" }
    if yname == "current-count" { return "CurrentCount" }
    if yname == "alloc-fails" { return "AllocFails" }
    if yname == "alloc-count" { return "AllocCount" }
    if yname == "freed-count" { return "FreedCount" }
    if yname == "memory-type" { return "MemoryType" }
    return ""
}

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetSegmentPath() string {
    return "memory-info"
}

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["structure-name"] = memoryInfo.StructureName
    leafs["size"] = memoryInfo.Size
    leafs["current-count"] = memoryInfo.CurrentCount
    leafs["alloc-fails"] = memoryInfo.AllocFails
    leafs["alloc-count"] = memoryInfo.AllocCount
    leafs["freed-count"] = memoryInfo.FreedCount
    leafs["memory-type"] = memoryInfo.MemoryType
    return leafs
}

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetBundleName() string { return "cisco_ios_xr" }

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetYangName() string { return "memory-info" }

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) SetParent(parent types.Entity) { memoryInfo.parent = parent }

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetParent() types.Entity { return memoryInfo.parent }

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetParentYangName() string { return "memory" }

// SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo
// EDM Memory Info
type SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size of the block. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Total request. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // Cache-hit success. The type is interface{} with range: 0..4294967295.
    Success interface{}

    // Cache-hit failure. The type is interface{} with range: 0..4294967295.
    Failure interface{}
}

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetFilter() yfilter.YFilter { return edmMemoryInfo.YFilter }

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) SetFilter(yf yfilter.YFilter) { edmMemoryInfo.YFilter = yf }

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "total" { return "Total" }
    if yname == "success" { return "Success" }
    if yname == "failure" { return "Failure" }
    return ""
}

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetSegmentPath() string {
    return "edm-memory-info"
}

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = edmMemoryInfo.Size
    leafs["total"] = edmMemoryInfo.Total
    leafs["success"] = edmMemoryInfo.Success
    leafs["failure"] = edmMemoryInfo.Failure
    return leafs
}

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetBundleName() string { return "cisco_ios_xr" }

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetYangName() string { return "edm-memory-info" }

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) SetParent(parent types.Entity) { edmMemoryInfo.parent = parent }

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetParent() types.Entity { return edmMemoryInfo.parent }

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetParentYangName() string { return "memory" }

// SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo
// String Memory Info
type SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size of the block. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // Total request. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // Cache-hit success. The type is interface{} with range: 0..4294967295.
    Success interface{}

    // Cache-hit failure. The type is interface{} with range: 0..4294967295.
    Failure interface{}
}

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetFilter() yfilter.YFilter { return stringMemoryInfo.YFilter }

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) SetFilter(yf yfilter.YFilter) { stringMemoryInfo.YFilter = yf }

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "total" { return "Total" }
    if yname == "success" { return "Success" }
    if yname == "failure" { return "Failure" }
    return ""
}

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetSegmentPath() string {
    return "string-memory-info"
}

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = stringMemoryInfo.Size
    leafs["total"] = stringMemoryInfo.Total
    leafs["success"] = stringMemoryInfo.Success
    leafs["failure"] = stringMemoryInfo.Failure
    return leafs
}

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetBundleName() string { return "cisco_ios_xr" }

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetYangName() string { return "string-memory-info" }

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) SetParent(parent types.Entity) { stringMemoryInfo.parent = parent }

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetParent() types.Entity { return stringMemoryInfo.parent }

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetParentYangName() string { return "memory" }

// SessionRedundancyAgent_Nodes_Node_GroupIds
// Data for particular subscriber group 
type SessionRedundancyAgent_Nodes_Node_GroupIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group id for subscriber group. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId.
    GroupId []SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId
}

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetFilter() yfilter.YFilter { return groupIds.YFilter }

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) SetFilter(yf yfilter.YFilter) { groupIds.YFilter = yf }

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    return ""
}

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetSegmentPath() string {
    return "group-ids"
}

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-id" {
        for _, c := range groupIds.GroupId {
            if groupIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId{}
        groupIds.GroupId = append(groupIds.GroupId, child)
        return &groupIds.GroupId[len(groupIds.GroupId)-1]
    }
    return nil
}

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupIds.GroupId {
        children[groupIds.GroupId[i].GetSegmentPath()] = &groupIds.GroupId[i]
    }
    return children
}

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetBundleName() string { return "cisco_ios_xr" }

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetYangName() string { return "group-ids" }

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) SetParent(parent types.Entity) { groupIds.parent = parent }

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetParent() types.Entity { return groupIds.parent }

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetParentYangName() string { return "node" }

// SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId
// Group id for subscriber group
type SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group Id. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    GroupId interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupIdXr interface{}

    // Group Description. The type is string.
    Description interface{}

    // Disabled by Config. The type is bool.
    Disabled interface{}

    // Preferred Init Role. The type is SergShowRole.
    InitRole interface{}

    // Negotiating Role. The type is SergShowRole.
    NegotiatingRole interface{}

    // Current Role. The type is SergShowRole.
    CurrentRole interface{}

    // Slave Mode. The type is SergShowSlaveMode.
    SlaveMode interface{}

    // Switch Over Hold Time. The type is interface{} with range: 0..4294967295.
    HoldTimer interface{}

    // Core Object Tracking Name. The type is string.
    CoreTrackingObjectName interface{}

    // Core Object Tracking Status. The type is bool.
    CoreTrackingObjectStatus interface{}

    // Access Object Tracking Name. The type is string.
    AccessTrackingObjectName interface{}

    // Access Object Tracking Status. The type is bool.
    AccessTrackingObjectStatus interface{}

    // Object Tracking Status (Enabled/Disabled). The type is bool.
    ObjectTrackingStatus interface{}

    // Peer IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerIpv4Address interface{}

    // Peer IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerIpv6Address interface{}

    // Peer Status. The type is SergPeerStatus.
    PeerStatus interface{}

    // Last Negotiation time of Peer. The type is string.
    PeerLastNegotiationTime interface{}

    // Last UP time of Peer. The type is string.
    PeerLastUpTime interface{}

    // Last Down time of Peer. The type is string.
    PeerLastDownTime interface{}

    // Peer Preferred Init Role. The type is SergShowRole.
    PeerInitRole interface{}

    // Peer Negotiating Role. The type is SergShowRole.
    PeerNegotiatingRole interface{}

    // Peer Current Role. The type is SergShowRole.
    PeerCurrentRole interface{}

    // Peer Object Tracking Status. The type is bool.
    PeerObjectTrackingStatus interface{}

    // Last Switchover time. The type is string.
    LastSwitchoverTime interface{}

    // Switchover Count. The type is interface{} with range: 0..4294967295.
    SwitchoverCount interface{}

    // Last Switchover Reason. The type is SergShowSoReason.
    LastSwitchoverReason interface{}

    // Switchover Hold Time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SwitchoverHoldTime interface{}

    // Session Count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}

    // Slave Session update fail count. The type is interface{} with range:
    // 0..4294967295.
    SlaveUpdateFailureCount interface{}

    // Pending Session Update Count. The type is interface{} with range:
    // 0..4294967295.
    PendingSessionUpdateCount interface{}

    // Pending Session Delete Count. The type is interface{} with range:
    // 0..4294967295.
    PendingSessionDeleteCount interface{}

    // No. of Configured Interfaces. The type is interface{} with range:
    // 0..4294967295.
    InterfaceCount interface{}

    // Revertive timer for SWO back. The type is interface{} with range:
    // 0..4294967295.
    RevertiveTimer interface{}

    // Switchover Revert Time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SwitchoverRevertTime interface{}

    // Client Session Count. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount.
    ClientSessionCount []SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount

    // Interface List. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface.
    Interface []SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetFilter() yfilter.YFilter { return groupId.YFilter }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) SetFilter(yf yfilter.YFilter) { groupId.YFilter = yf }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "group-id-xr" { return "GroupIdXr" }
    if yname == "description" { return "Description" }
    if yname == "disabled" { return "Disabled" }
    if yname == "init-role" { return "InitRole" }
    if yname == "negotiating-role" { return "NegotiatingRole" }
    if yname == "current-role" { return "CurrentRole" }
    if yname == "slave-mode" { return "SlaveMode" }
    if yname == "hold-timer" { return "HoldTimer" }
    if yname == "core-tracking-object-name" { return "CoreTrackingObjectName" }
    if yname == "core-tracking-object-status" { return "CoreTrackingObjectStatus" }
    if yname == "access-tracking-object-name" { return "AccessTrackingObjectName" }
    if yname == "access-tracking-object-status" { return "AccessTrackingObjectStatus" }
    if yname == "object-tracking-status" { return "ObjectTrackingStatus" }
    if yname == "peer-ipv4-address" { return "PeerIpv4Address" }
    if yname == "peer-ipv6-address" { return "PeerIpv6Address" }
    if yname == "peer-status" { return "PeerStatus" }
    if yname == "peer-last-negotiation-time" { return "PeerLastNegotiationTime" }
    if yname == "peer-last-up-time" { return "PeerLastUpTime" }
    if yname == "peer-last-down-time" { return "PeerLastDownTime" }
    if yname == "peer-init-role" { return "PeerInitRole" }
    if yname == "peer-negotiating-role" { return "PeerNegotiatingRole" }
    if yname == "peer-current-role" { return "PeerCurrentRole" }
    if yname == "peer-object-tracking-status" { return "PeerObjectTrackingStatus" }
    if yname == "last-switchover-time" { return "LastSwitchoverTime" }
    if yname == "switchover-count" { return "SwitchoverCount" }
    if yname == "last-switchover-reason" { return "LastSwitchoverReason" }
    if yname == "switchover-hold-time" { return "SwitchoverHoldTime" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "slave-update-failure-count" { return "SlaveUpdateFailureCount" }
    if yname == "pending-session-update-count" { return "PendingSessionUpdateCount" }
    if yname == "pending-session-delete-count" { return "PendingSessionDeleteCount" }
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "revertive-timer" { return "RevertiveTimer" }
    if yname == "switchover-revert-time" { return "SwitchoverRevertTime" }
    if yname == "client-session-count" { return "ClientSessionCount" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetSegmentPath() string {
    return "group-id" + "[group-id='" + fmt.Sprintf("%v", groupId.GroupId) + "']"
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-session-count" {
        for _, c := range groupId.ClientSessionCount {
            if groupId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount{}
        groupId.ClientSessionCount = append(groupId.ClientSessionCount, child)
        return &groupId.ClientSessionCount[len(groupId.ClientSessionCount)-1]
    }
    if childYangName == "interface" {
        for _, c := range groupId.Interface {
            if groupId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface{}
        groupId.Interface = append(groupId.Interface, child)
        return &groupId.Interface[len(groupId.Interface)-1]
    }
    return nil
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupId.ClientSessionCount {
        children[groupId.ClientSessionCount[i].GetSegmentPath()] = &groupId.ClientSessionCount[i]
    }
    for i := range groupId.Interface {
        children[groupId.Interface[i].GetSegmentPath()] = &groupId.Interface[i]
    }
    return children
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = groupId.GroupId
    leafs["group-id-xr"] = groupId.GroupIdXr
    leafs["description"] = groupId.Description
    leafs["disabled"] = groupId.Disabled
    leafs["init-role"] = groupId.InitRole
    leafs["negotiating-role"] = groupId.NegotiatingRole
    leafs["current-role"] = groupId.CurrentRole
    leafs["slave-mode"] = groupId.SlaveMode
    leafs["hold-timer"] = groupId.HoldTimer
    leafs["core-tracking-object-name"] = groupId.CoreTrackingObjectName
    leafs["core-tracking-object-status"] = groupId.CoreTrackingObjectStatus
    leafs["access-tracking-object-name"] = groupId.AccessTrackingObjectName
    leafs["access-tracking-object-status"] = groupId.AccessTrackingObjectStatus
    leafs["object-tracking-status"] = groupId.ObjectTrackingStatus
    leafs["peer-ipv4-address"] = groupId.PeerIpv4Address
    leafs["peer-ipv6-address"] = groupId.PeerIpv6Address
    leafs["peer-status"] = groupId.PeerStatus
    leafs["peer-last-negotiation-time"] = groupId.PeerLastNegotiationTime
    leafs["peer-last-up-time"] = groupId.PeerLastUpTime
    leafs["peer-last-down-time"] = groupId.PeerLastDownTime
    leafs["peer-init-role"] = groupId.PeerInitRole
    leafs["peer-negotiating-role"] = groupId.PeerNegotiatingRole
    leafs["peer-current-role"] = groupId.PeerCurrentRole
    leafs["peer-object-tracking-status"] = groupId.PeerObjectTrackingStatus
    leafs["last-switchover-time"] = groupId.LastSwitchoverTime
    leafs["switchover-count"] = groupId.SwitchoverCount
    leafs["last-switchover-reason"] = groupId.LastSwitchoverReason
    leafs["switchover-hold-time"] = groupId.SwitchoverHoldTime
    leafs["session-count"] = groupId.SessionCount
    leafs["slave-update-failure-count"] = groupId.SlaveUpdateFailureCount
    leafs["pending-session-update-count"] = groupId.PendingSessionUpdateCount
    leafs["pending-session-delete-count"] = groupId.PendingSessionDeleteCount
    leafs["interface-count"] = groupId.InterfaceCount
    leafs["revertive-timer"] = groupId.RevertiveTimer
    leafs["switchover-revert-time"] = groupId.SwitchoverRevertTime
    return leafs
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetBundleName() string { return "cisco_ios_xr" }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetYangName() string { return "group-id" }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) SetParent(parent types.Entity) { groupId.parent = parent }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetParent() types.Entity { return groupId.parent }

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetParentYangName() string { return "group-ids" }

// SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount
// Client Session Count
type SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Component. The type is SergShowComp.
    Component interface{}

    // Session count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}
}

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetFilter() yfilter.YFilter { return clientSessionCount.YFilter }

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) SetFilter(yf yfilter.YFilter) { clientSessionCount.YFilter = yf }

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetGoName(yname string) string {
    if yname == "component" { return "Component" }
    if yname == "session-count" { return "SessionCount" }
    return ""
}

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetSegmentPath() string {
    return "client-session-count"
}

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["component"] = clientSessionCount.Component
    leafs["session-count"] = clientSessionCount.SessionCount
    return leafs
}

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetBundleName() string { return "cisco_ios_xr" }

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetYangName() string { return "client-session-count" }

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) SetParent(parent types.Entity) { clientSessionCount.parent = parent }

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetParent() types.Entity { return clientSessionCount.parent }

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetParentYangName() string { return "group-id" }

// SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface
// Interface List
type SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Interface Synchronization ID. The type is interface{} with range:
    // 0..4294967295.
    InterfaceSynchronizationId interface{}

    // Forward Referenced. The type is bool.
    ForwardReferenced interface{}

    // Session Count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}
}

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-synchronization-id" { return "InterfaceSynchronizationId" }
    if yname == "forward-referenced" { return "ForwardReferenced" }
    if yname == "session-count" { return "SessionCount" }
    return ""
}

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-synchronization-id"] = self.InterfaceSynchronizationId
    leafs["forward-referenced"] = self.ForwardReferenced
    leafs["session-count"] = self.SessionCount
    return leafs
}

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetYangName() string { return "interface" }

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetParent() types.Entity { return self.parent }

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetParentYangName() string { return "group-id" }

// SessionRedundancyAgent_Nodes_Node_Interfaces
// List of interfaces
type SessionRedundancyAgent_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify interface name. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_Interfaces_Interface.
    Interface []SessionRedundancyAgent_Nodes_Node_Interfaces_Interface
}

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// SessionRedundancyAgent_Nodes_Node_Interfaces_Interface
// Specify interface name
type SessionRedundancyAgent_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Interface Sync ID. The type is interface{} with range: 0..4294967295.
    InterfaceSynchronizationId interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // SERG Role. The type is SergShowImRole.
    Role interface{}

    // Forward Referenced. The type is bool.
    ForwardReferenced interface{}

    // Session Count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}

    // Interface Enable Error Count. The type is interface{} with range:
    // 0..4294967295.
    InterfaceEnableErrorCount interface{}

    // Interface Disable Error Count. The type is interface{} with range:
    // 0..4294967295.
    InterfaceDisableErrorCount interface{}

    // Interface Caps Add Error Count. The type is interface{} with range:
    // 0..4294967295.
    InterfaceCapsAddErrorCount interface{}

    // Interface Caps Remove Error Count. The type is interface{} with range:
    // 0..4294967295.
    InterfaceCapsRemoveErrorCount interface{}

    // Interface Attribute Update Error Count. The type is interface{} with range:
    // 0..4294967295.
    InterfaceAttributeUpdateErrorCount interface{}

    // Interface Batch Operation.
    InterfaceOper SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper

    // Interface Status.
    InterfaceStatus SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus

    // Interface status for each client. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus.
    ClientStatus []SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus
}

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-synchronization-id" { return "InterfaceSynchronizationId" }
    if yname == "group-id" { return "GroupId" }
    if yname == "role" { return "Role" }
    if yname == "forward-referenced" { return "ForwardReferenced" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "interface-enable-error-count" { return "InterfaceEnableErrorCount" }
    if yname == "interface-disable-error-count" { return "InterfaceDisableErrorCount" }
    if yname == "interface-caps-add-error-count" { return "InterfaceCapsAddErrorCount" }
    if yname == "interface-caps-remove-error-count" { return "InterfaceCapsRemoveErrorCount" }
    if yname == "interface-attribute-update-error-count" { return "InterfaceAttributeUpdateErrorCount" }
    if yname == "interface-oper" { return "InterfaceOper" }
    if yname == "interface-status" { return "InterfaceStatus" }
    if yname == "client-status" { return "ClientStatus" }
    return ""
}

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-oper" {
        return &self.InterfaceOper
    }
    if childYangName == "interface-status" {
        return &self.InterfaceStatus
    }
    if childYangName == "client-status" {
        for _, c := range self.ClientStatus {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus{}
        self.ClientStatus = append(self.ClientStatus, child)
        return &self.ClientStatus[len(self.ClientStatus)-1]
    }
    return nil
}

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-oper"] = &self.InterfaceOper
    children["interface-status"] = &self.InterfaceStatus
    for i := range self.ClientStatus {
        children[self.ClientStatus[i].GetSegmentPath()] = &self.ClientStatus[i]
    }
    return children
}

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = self.Interface
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-synchronization-id"] = self.InterfaceSynchronizationId
    leafs["group-id"] = self.GroupId
    leafs["role"] = self.Role
    leafs["forward-referenced"] = self.ForwardReferenced
    leafs["session-count"] = self.SessionCount
    leafs["interface-enable-error-count"] = self.InterfaceEnableErrorCount
    leafs["interface-disable-error-count"] = self.InterfaceDisableErrorCount
    leafs["interface-caps-add-error-count"] = self.InterfaceCapsAddErrorCount
    leafs["interface-caps-remove-error-count"] = self.InterfaceCapsRemoveErrorCount
    leafs["interface-attribute-update-error-count"] = self.InterfaceAttributeUpdateErrorCount
    return leafs
}

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper
// Interface Batch Operation
type SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational Registration Enabled. The type is bool.
    IdbOperRegEnable interface{}

    // Operational Registration Disabled. The type is bool.
    IdbOperRegDisable interface{}

    // Operational Caps Add. The type is bool.
    IdbOperCapsAdd interface{}

    // Operational Caps Remove. The type is bool.
    IdbOperCapsRemove interface{}

    // Operational Attribute Update. The type is bool.
    IdbOperAttrUpdate interface{}
}

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetFilter() yfilter.YFilter { return interfaceOper.YFilter }

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) SetFilter(yf yfilter.YFilter) { interfaceOper.YFilter = yf }

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetGoName(yname string) string {
    if yname == "idb-oper-reg-enable" { return "IdbOperRegEnable" }
    if yname == "idb-oper-reg-disable" { return "IdbOperRegDisable" }
    if yname == "idb-oper-caps-add" { return "IdbOperCapsAdd" }
    if yname == "idb-oper-caps-remove" { return "IdbOperCapsRemove" }
    if yname == "idb-oper-attr-update" { return "IdbOperAttrUpdate" }
    return ""
}

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetSegmentPath() string {
    return "interface-oper"
}

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["idb-oper-reg-enable"] = interfaceOper.IdbOperRegEnable
    leafs["idb-oper-reg-disable"] = interfaceOper.IdbOperRegDisable
    leafs["idb-oper-caps-add"] = interfaceOper.IdbOperCapsAdd
    leafs["idb-oper-caps-remove"] = interfaceOper.IdbOperCapsRemove
    leafs["idb-oper-attr-update"] = interfaceOper.IdbOperAttrUpdate
    return leafs
}

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetYangName() string { return "interface-oper" }

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) SetParent(parent types.Entity) { interfaceOper.parent = parent }

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetParent() types.Entity { return interfaceOper.parent }

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetParentYangName() string { return "interface" }

// SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus
// Interface Status
type SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Forward Referenced. The type is bool.
    IdbStateFwdRef interface{}

    // Interface State Stale. The type is bool.
    IdbStateStale interface{}

    // Interface State Registered. The type is bool.
    IdbStateRegistered interface{}

    // Interface State Caps Added. The type is bool.
    IdbStateCapsAdded interface{}

    // Interface State Owned Resource. The type is bool.
    IdbStateOwnedReSource interface{}

    // Interface Client EOMS Pending. The type is bool.
    IdbClientEomsPending interface{}

    // Interface Caps Remove Pending. The type is bool.
    IdbStatePEndCapsRem interface{}

    // Interface Registration Disable Pending. The type is bool.
    IdbStatePEndRegDisable interface{}
}

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetFilter() yfilter.YFilter { return interfaceStatus.YFilter }

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) SetFilter(yf yfilter.YFilter) { interfaceStatus.YFilter = yf }

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetGoName(yname string) string {
    if yname == "idb-state-fwd-ref" { return "IdbStateFwdRef" }
    if yname == "idb-state-stale" { return "IdbStateStale" }
    if yname == "idb-state-registered" { return "IdbStateRegistered" }
    if yname == "idb-state-caps-added" { return "IdbStateCapsAdded" }
    if yname == "idb-state-owned-re-source" { return "IdbStateOwnedReSource" }
    if yname == "idb-client-eoms-pending" { return "IdbClientEomsPending" }
    if yname == "idb-state-p-end-caps-rem" { return "IdbStatePEndCapsRem" }
    if yname == "idb-state-p-end-reg-disable" { return "IdbStatePEndRegDisable" }
    return ""
}

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetSegmentPath() string {
    return "interface-status"
}

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["idb-state-fwd-ref"] = interfaceStatus.IdbStateFwdRef
    leafs["idb-state-stale"] = interfaceStatus.IdbStateStale
    leafs["idb-state-registered"] = interfaceStatus.IdbStateRegistered
    leafs["idb-state-caps-added"] = interfaceStatus.IdbStateCapsAdded
    leafs["idb-state-owned-re-source"] = interfaceStatus.IdbStateOwnedReSource
    leafs["idb-client-eoms-pending"] = interfaceStatus.IdbClientEomsPending
    leafs["idb-state-p-end-caps-rem"] = interfaceStatus.IdbStatePEndCapsRem
    leafs["idb-state-p-end-reg-disable"] = interfaceStatus.IdbStatePEndRegDisable
    return leafs
}

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetYangName() string { return "interface-status" }

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) SetParent(parent types.Entity) { interfaceStatus.parent = parent }

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetParent() types.Entity { return interfaceStatus.parent }

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetParentYangName() string { return "interface" }

// SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus
// Interface status for each client
type SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Component. The type is SergShowComp.
    Component interface{}

    // SERG SHOW IDB CLIENT EOMS PENDING. The type is bool.
    SergShowIdbClientEomsPending interface{}

    // SERG SHOW IDB CLIENT SYNC EOD PENDING. The type is bool.
    SergShowIdbClientSyncEodPending interface{}

    // session count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetFilter() yfilter.YFilter { return clientStatus.YFilter }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) SetFilter(yf yfilter.YFilter) { clientStatus.YFilter = yf }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetGoName(yname string) string {
    if yname == "component" { return "Component" }
    if yname == "serg-show-idb-client-eoms-pending" { return "SergShowIdbClientEomsPending" }
    if yname == "serg-show-idb-client-sync-eod-pending" { return "SergShowIdbClientSyncEodPending" }
    if yname == "session-count" { return "SessionCount" }
    return ""
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetSegmentPath() string {
    return "client-status"
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["component"] = clientStatus.Component
    leafs["serg-show-idb-client-eoms-pending"] = clientStatus.SergShowIdbClientEomsPending
    leafs["serg-show-idb-client-sync-eod-pending"] = clientStatus.SergShowIdbClientSyncEodPending
    leafs["session-count"] = clientStatus.SessionCount
    return leafs
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetBundleName() string { return "cisco_ios_xr" }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetYangName() string { return "client-status" }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) SetParent(parent types.Entity) { clientStatus.parent = parent }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetParent() types.Entity { return clientStatus.parent }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetParentYangName() string { return "interface" }

// SessionRedundancyAgent_Nodes_Node_StatsGlobal
// Stats Global
type SessionRedundancyAgent_Nodes_Node_StatsGlobal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source Interface Name. The type is string.
    SourceInterfaceName interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Source Interface IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceInterfaceIpv4Address interface{}

    // Source Interface IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceInterfaceIpv6Address interface{}

    // Redundancy Role. The type is string.
    RedundancyRole interface{}

    // Restart Client Sync In Progress Flag. The type is bool.
    RestartClientSyncInProgress interface{}

    // Synchronization TimeStamp. The type is string.
    ClientInitSyncTimeStamp interface{}

    // Restart Peer Sync In Progress Flag. The type is bool.
    RestartPeerSyncInProgress interface{}

    // Synchronization TimeStamp. The type is string.
    PeerInitSyncTimeStamp interface{}

    // Sync In Progress Flag. The type is bool.
    SyncInProgress interface{}

    // Value in Seconds to trigger the peer actions. The type is interface{} with
    // range: 0..4294967295. Units are second.
    PeerActionTimer interface{}

    // Value in Seconds to trigger the retry. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RetryTimerRemaining interface{}

    // TxListClientRegistrationInvalid. The type is interface{} with range:
    // 0..4294967295.
    TxListClientRegistrationInvalid interface{}

    // TxListClientDeRegistrationInvalid. The type is interface{} with range:
    // 0..4294967295.
    TxListClientDeRegistrationInvalid interface{}

    // TxListClientConnectionUp. The type is interface{} with range:
    // 0..4294967295.
    TxListClientConnectionUp interface{}

    // TxListClientConnectionDown. The type is interface{} with range:
    // 0..4294967295.
    TxListClientConnectionDown interface{}

    // TxListClientPeerDone. The type is interface{} with range: 0..4294967295.
    TxListClientPeerDone interface{}

    // TxListClientMessageCallBack. The type is interface{} with range:
    // 0..4294967295.
    TxListClientMessageCallBack interface{}

    // TxListClientReceiveValid. The type is interface{} with range:
    // 0..4294967295.
    TxListClientReceiveValid interface{}

    // TxListClientReceiveInValid. The type is interface{} with range:
    // 0..4294967295.
    TxListClientReceiveInvalid interface{}

    // TxListClientReceiveCommandValid. The type is interface{} with range:
    // 0..4294967295.
    TxListClientReceiveCommandValid interface{}

    // TxListClientReceiveCommandInValid. The type is interface{} with range:
    // 0..4294967295.
    TxListClientReceiveCommandInvalid interface{}

    // TxListClientReceiveSessionSessionValid. The type is interface{} with range:
    // 0..4294967295.
    TxListClientReceiveSessionSessionvalid interface{}

    // TxListClientReceiveSessionSessionInValid. The type is interface{} with
    // range: 0..4294967295.
    TxListClientReceiveSessionSessionInvalid interface{}

    // TxListPeerTimerHandler. The type is interface{} with range: 0..4294967295.
    TxListPeerTimerHandler interface{}

    // TxListPeerRegistrationInValid. The type is interface{} with range:
    // 0..4294967295.
    TxListPeerRegistrationInvalid interface{}

    // TxListPeerDeRegistrationInValid. The type is interface{} with range:
    // 0..4294967295.
    TxListPeerDeRegistrationInvalid interface{}

    // TxListPeerMessageCallBackValid. The type is interface{} with range:
    // 0..4294967295.
    TxListPeerMessageCallBackValid interface{}

    // TxListPeerMessageCallBackInValid. The type is interface{} with range:
    // 0..4294967295.
    TxListPeerMessageCallBackInvalid interface{}

    // TxListPeerDone. The type is interface{} with range: 0..4294967295.
    TxListPeerDone interface{}

    // TxListPeerCmdConnectionUpNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListPeerCmdConnectionUpNotOk interface{}

    // TxListPeerCmdConnectionDownNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListPeerCmdConnectionDownNotOk interface{}

    // TxListPeerSessionConnectionUpNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListPeerSessionConnectionUpNotOk interface{}

    // TxListPeerSessionConnectionDownNotOk. The type is interface{} with range:
    // 0..4294967295.
    TxListPeerSessionConnectionDownNotOk interface{}

    // intf status statistics.
    IntfStatusStatistics SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics

    // tx list statistics.
    TxListStatistics SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics

    // Client Status. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus.
    ClientStatus []SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus

    // Opaque memory Stats. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus.
    OpaqueMemoryStatus []SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus

    // TCP TxList Statistics. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus.
    TxListOverTcpStatus []SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus
}

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetFilter() yfilter.YFilter { return statsGlobal.YFilter }

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) SetFilter(yf yfilter.YFilter) { statsGlobal.YFilter = yf }

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetGoName(yname string) string {
    if yname == "source-interface-name" { return "SourceInterfaceName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source-interface-ipv4-address" { return "SourceInterfaceIpv4Address" }
    if yname == "source-interface-ipv6-address" { return "SourceInterfaceIpv6Address" }
    if yname == "redundancy-role" { return "RedundancyRole" }
    if yname == "restart-client-sync-in-progress" { return "RestartClientSyncInProgress" }
    if yname == "client-init-sync-time-stamp" { return "ClientInitSyncTimeStamp" }
    if yname == "restart-peer-sync-in-progress" { return "RestartPeerSyncInProgress" }
    if yname == "peer-init-sync-time-stamp" { return "PeerInitSyncTimeStamp" }
    if yname == "sync-in-progress" { return "SyncInProgress" }
    if yname == "peer-action-timer" { return "PeerActionTimer" }
    if yname == "retry-timer-remaining" { return "RetryTimerRemaining" }
    if yname == "tx-list-client-registration-invalid" { return "TxListClientRegistrationInvalid" }
    if yname == "tx-list-client-de-registration-invalid" { return "TxListClientDeRegistrationInvalid" }
    if yname == "tx-list-client-connection-up" { return "TxListClientConnectionUp" }
    if yname == "tx-list-client-connection-down" { return "TxListClientConnectionDown" }
    if yname == "tx-list-client-peer-done" { return "TxListClientPeerDone" }
    if yname == "tx-list-client-message-call-back" { return "TxListClientMessageCallBack" }
    if yname == "tx-list-client-receive-valid" { return "TxListClientReceiveValid" }
    if yname == "tx-list-client-receive-invalid" { return "TxListClientReceiveInvalid" }
    if yname == "tx-list-client-receive-command-valid" { return "TxListClientReceiveCommandValid" }
    if yname == "tx-list-client-receive-command-invalid" { return "TxListClientReceiveCommandInvalid" }
    if yname == "tx-list-client-receive-session-sessionvalid" { return "TxListClientReceiveSessionSessionvalid" }
    if yname == "tx-list-client-receive-session-session-invalid" { return "TxListClientReceiveSessionSessionInvalid" }
    if yname == "tx-list-peer-timer-handler" { return "TxListPeerTimerHandler" }
    if yname == "tx-list-peer-registration-invalid" { return "TxListPeerRegistrationInvalid" }
    if yname == "tx-list-peer-de-registration-invalid" { return "TxListPeerDeRegistrationInvalid" }
    if yname == "tx-list-peer-message-call-back-valid" { return "TxListPeerMessageCallBackValid" }
    if yname == "tx-list-peer-message-call-back-invalid" { return "TxListPeerMessageCallBackInvalid" }
    if yname == "tx-list-peer-done" { return "TxListPeerDone" }
    if yname == "tx-list-peer-cmd-connection-up-not-ok" { return "TxListPeerCmdConnectionUpNotOk" }
    if yname == "tx-list-peer-cmd-connection-down-not-ok" { return "TxListPeerCmdConnectionDownNotOk" }
    if yname == "tx-list-peer-session-connection-up-not-ok" { return "TxListPeerSessionConnectionUpNotOk" }
    if yname == "tx-list-peer-session-connection-down-not-ok" { return "TxListPeerSessionConnectionDownNotOk" }
    if yname == "intf-status-statistics" { return "IntfStatusStatistics" }
    if yname == "tx-list-statistics" { return "TxListStatistics" }
    if yname == "client-status" { return "ClientStatus" }
    if yname == "opaque-memory-status" { return "OpaqueMemoryStatus" }
    if yname == "tx-list-over-tcp-status" { return "TxListOverTcpStatus" }
    return ""
}

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetSegmentPath() string {
    return "stats-global"
}

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intf-status-statistics" {
        return &statsGlobal.IntfStatusStatistics
    }
    if childYangName == "tx-list-statistics" {
        return &statsGlobal.TxListStatistics
    }
    if childYangName == "client-status" {
        for _, c := range statsGlobal.ClientStatus {
            if statsGlobal.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus{}
        statsGlobal.ClientStatus = append(statsGlobal.ClientStatus, child)
        return &statsGlobal.ClientStatus[len(statsGlobal.ClientStatus)-1]
    }
    if childYangName == "opaque-memory-status" {
        for _, c := range statsGlobal.OpaqueMemoryStatus {
            if statsGlobal.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus{}
        statsGlobal.OpaqueMemoryStatus = append(statsGlobal.OpaqueMemoryStatus, child)
        return &statsGlobal.OpaqueMemoryStatus[len(statsGlobal.OpaqueMemoryStatus)-1]
    }
    if childYangName == "tx-list-over-tcp-status" {
        for _, c := range statsGlobal.TxListOverTcpStatus {
            if statsGlobal.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus{}
        statsGlobal.TxListOverTcpStatus = append(statsGlobal.TxListOverTcpStatus, child)
        return &statsGlobal.TxListOverTcpStatus[len(statsGlobal.TxListOverTcpStatus)-1]
    }
    return nil
}

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["intf-status-statistics"] = &statsGlobal.IntfStatusStatistics
    children["tx-list-statistics"] = &statsGlobal.TxListStatistics
    for i := range statsGlobal.ClientStatus {
        children[statsGlobal.ClientStatus[i].GetSegmentPath()] = &statsGlobal.ClientStatus[i]
    }
    for i := range statsGlobal.OpaqueMemoryStatus {
        children[statsGlobal.OpaqueMemoryStatus[i].GetSegmentPath()] = &statsGlobal.OpaqueMemoryStatus[i]
    }
    for i := range statsGlobal.TxListOverTcpStatus {
        children[statsGlobal.TxListOverTcpStatus[i].GetSegmentPath()] = &statsGlobal.TxListOverTcpStatus[i]
    }
    return children
}

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-interface-name"] = statsGlobal.SourceInterfaceName
    leafs["vrf-name"] = statsGlobal.VrfName
    leafs["source-interface-ipv4-address"] = statsGlobal.SourceInterfaceIpv4Address
    leafs["source-interface-ipv6-address"] = statsGlobal.SourceInterfaceIpv6Address
    leafs["redundancy-role"] = statsGlobal.RedundancyRole
    leafs["restart-client-sync-in-progress"] = statsGlobal.RestartClientSyncInProgress
    leafs["client-init-sync-time-stamp"] = statsGlobal.ClientInitSyncTimeStamp
    leafs["restart-peer-sync-in-progress"] = statsGlobal.RestartPeerSyncInProgress
    leafs["peer-init-sync-time-stamp"] = statsGlobal.PeerInitSyncTimeStamp
    leafs["sync-in-progress"] = statsGlobal.SyncInProgress
    leafs["peer-action-timer"] = statsGlobal.PeerActionTimer
    leafs["retry-timer-remaining"] = statsGlobal.RetryTimerRemaining
    leafs["tx-list-client-registration-invalid"] = statsGlobal.TxListClientRegistrationInvalid
    leafs["tx-list-client-de-registration-invalid"] = statsGlobal.TxListClientDeRegistrationInvalid
    leafs["tx-list-client-connection-up"] = statsGlobal.TxListClientConnectionUp
    leafs["tx-list-client-connection-down"] = statsGlobal.TxListClientConnectionDown
    leafs["tx-list-client-peer-done"] = statsGlobal.TxListClientPeerDone
    leafs["tx-list-client-message-call-back"] = statsGlobal.TxListClientMessageCallBack
    leafs["tx-list-client-receive-valid"] = statsGlobal.TxListClientReceiveValid
    leafs["tx-list-client-receive-invalid"] = statsGlobal.TxListClientReceiveInvalid
    leafs["tx-list-client-receive-command-valid"] = statsGlobal.TxListClientReceiveCommandValid
    leafs["tx-list-client-receive-command-invalid"] = statsGlobal.TxListClientReceiveCommandInvalid
    leafs["tx-list-client-receive-session-sessionvalid"] = statsGlobal.TxListClientReceiveSessionSessionvalid
    leafs["tx-list-client-receive-session-session-invalid"] = statsGlobal.TxListClientReceiveSessionSessionInvalid
    leafs["tx-list-peer-timer-handler"] = statsGlobal.TxListPeerTimerHandler
    leafs["tx-list-peer-registration-invalid"] = statsGlobal.TxListPeerRegistrationInvalid
    leafs["tx-list-peer-de-registration-invalid"] = statsGlobal.TxListPeerDeRegistrationInvalid
    leafs["tx-list-peer-message-call-back-valid"] = statsGlobal.TxListPeerMessageCallBackValid
    leafs["tx-list-peer-message-call-back-invalid"] = statsGlobal.TxListPeerMessageCallBackInvalid
    leafs["tx-list-peer-done"] = statsGlobal.TxListPeerDone
    leafs["tx-list-peer-cmd-connection-up-not-ok"] = statsGlobal.TxListPeerCmdConnectionUpNotOk
    leafs["tx-list-peer-cmd-connection-down-not-ok"] = statsGlobal.TxListPeerCmdConnectionDownNotOk
    leafs["tx-list-peer-session-connection-up-not-ok"] = statsGlobal.TxListPeerSessionConnectionUpNotOk
    leafs["tx-list-peer-session-connection-down-not-ok"] = statsGlobal.TxListPeerSessionConnectionDownNotOk
    return leafs
}

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetBundleName() string { return "cisco_ios_xr" }

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetYangName() string { return "stats-global" }

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) SetParent(parent types.Entity) { statsGlobal.parent = parent }

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetParent() types.Entity { return statsGlobal.parent }

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetParentYangName() string { return "node" }

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics
// intf status statistics
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // No. of interfaces pending caps remove. The type is interface{} with range:
    // 0..4294967295.
    PendCapsRemCnt interface{}

    // No. of interfaces pending red disable. The type is interface{} with range:
    // 0..4294967295.
    PendRegDisableCnt interface{}

    // No. of interfaces pending for other (except unreg/caps rem) batch oper. The
    // type is interface{} with range: 0..4294967295.
    PendOtherBatchOperCnt interface{}

    // No. of non stale interfaces. The type is interface{} with range:
    // 0..4294967295.
    NonStaleCnt interface{}

    // No. of interface bound to a group. The type is interface{} with range:
    // 0..4294967295.
    GrpBoundCnt interface{}
}

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetFilter() yfilter.YFilter { return intfStatusStatistics.YFilter }

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) SetFilter(yf yfilter.YFilter) { intfStatusStatistics.YFilter = yf }

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetGoName(yname string) string {
    if yname == "pend-caps-rem-cnt" { return "PendCapsRemCnt" }
    if yname == "pend-reg-disable-cnt" { return "PendRegDisableCnt" }
    if yname == "pend-other-batch-oper-cnt" { return "PendOtherBatchOperCnt" }
    if yname == "non-stale-cnt" { return "NonStaleCnt" }
    if yname == "grp-bound-cnt" { return "GrpBoundCnt" }
    return ""
}

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetSegmentPath() string {
    return "intf-status-statistics"
}

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pend-caps-rem-cnt"] = intfStatusStatistics.PendCapsRemCnt
    leafs["pend-reg-disable-cnt"] = intfStatusStatistics.PendRegDisableCnt
    leafs["pend-other-batch-oper-cnt"] = intfStatusStatistics.PendOtherBatchOperCnt
    leafs["non-stale-cnt"] = intfStatusStatistics.NonStaleCnt
    leafs["grp-bound-cnt"] = intfStatusStatistics.GrpBoundCnt
    return leafs
}

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetYangName() string { return "intf-status-statistics" }

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) SetParent(parent types.Entity) { intfStatusStatistics.parent = parent }

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetParent() types.Entity { return intfStatusStatistics.parent }

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetParentYangName() string { return "stats-global" }

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics
// tx list statistics
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TxListEncodeMarkerOk. The type is interface{} with range: 0..4294967295.
    TxListEncodeMarkerOk interface{}

    // TxListEncodeMarkerPartialWrite. The type is interface{} with range:
    // 0..4294967295.
    TxListEncodeMarkerPartialWrite interface{}

    // TxListCleanMarker. The type is interface{} with range: 0..4294967295.
    TxListCleanMarker interface{}

    // TxListEncodeCommandOk. The type is interface{} with range: 0..4294967295.
    TxListEncodeCommandOk interface{}

    // TxListEncodeCommandPartialWrite. The type is interface{} with range:
    // 0..4294967295.
    TxListEncodeCommandPartialWrite interface{}

    // TxListCleanCommand. The type is interface{} with range: 0..4294967295.
    TxListCleanCommand interface{}

    // TxListEncodeNegotiationOk. The type is interface{} with range:
    // 0..4294967295.
    TxListEncodeNegotiationOk interface{}

    // TxListEncodeNegotiationPartialWrite. The type is interface{} with range:
    // 0..4294967295.
    TxListEncodeNegotiationPartialWrite interface{}

    // TxListCleanNegotiation. The type is interface{} with range: 0..4294967295.
    TxListCleanNegotiation interface{}
}

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetFilter() yfilter.YFilter { return txListStatistics.YFilter }

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) SetFilter(yf yfilter.YFilter) { txListStatistics.YFilter = yf }

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetGoName(yname string) string {
    if yname == "tx-list-encode-marker-ok" { return "TxListEncodeMarkerOk" }
    if yname == "tx-list-encode-marker-partial-write" { return "TxListEncodeMarkerPartialWrite" }
    if yname == "tx-list-clean-marker" { return "TxListCleanMarker" }
    if yname == "tx-list-encode-command-ok" { return "TxListEncodeCommandOk" }
    if yname == "tx-list-encode-command-partial-write" { return "TxListEncodeCommandPartialWrite" }
    if yname == "tx-list-clean-command" { return "TxListCleanCommand" }
    if yname == "tx-list-encode-negotiation-ok" { return "TxListEncodeNegotiationOk" }
    if yname == "tx-list-encode-negotiation-partial-write" { return "TxListEncodeNegotiationPartialWrite" }
    if yname == "tx-list-clean-negotiation" { return "TxListCleanNegotiation" }
    return ""
}

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetSegmentPath() string {
    return "tx-list-statistics"
}

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-list-encode-marker-ok"] = txListStatistics.TxListEncodeMarkerOk
    leafs["tx-list-encode-marker-partial-write"] = txListStatistics.TxListEncodeMarkerPartialWrite
    leafs["tx-list-clean-marker"] = txListStatistics.TxListCleanMarker
    leafs["tx-list-encode-command-ok"] = txListStatistics.TxListEncodeCommandOk
    leafs["tx-list-encode-command-partial-write"] = txListStatistics.TxListEncodeCommandPartialWrite
    leafs["tx-list-clean-command"] = txListStatistics.TxListCleanCommand
    leafs["tx-list-encode-negotiation-ok"] = txListStatistics.TxListEncodeNegotiationOk
    leafs["tx-list-encode-negotiation-partial-write"] = txListStatistics.TxListEncodeNegotiationPartialWrite
    leafs["tx-list-clean-negotiation"] = txListStatistics.TxListCleanNegotiation
    return leafs
}

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetYangName() string { return "tx-list-statistics" }

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) SetParent(parent types.Entity) { txListStatistics.parent = parent }

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetParent() types.Entity { return txListStatistics.parent }

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetParentYangName() string { return "stats-global" }

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus
// Client Status
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Component. The type is SergShowComp.
    Component interface{}

    // ClientConnectionStatus. The type is bool.
    ClientConnectionStatus interface{}

    // ClientInitializationSynchronizationPending. The type is bool.
    ClientInitializationSynchronizationPending interface{}

    // ClientSynchronizationEndOfDownloadPending. The type is bool.
    ClientSynchronizationEndOfDownloadPending interface{}

    // UpTimeStamp. The type is string.
    UpTimeStamp interface{}

    // DownTimeStamp. The type is string.
    DownTimeStamp interface{}

    // Value in Seconds to trigger the client cleanup. The type is interface{}
    // with range: 0..4294967295. Units are second.
    CleanUpTimerRemaining interface{}
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetFilter() yfilter.YFilter { return clientStatus.YFilter }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) SetFilter(yf yfilter.YFilter) { clientStatus.YFilter = yf }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetGoName(yname string) string {
    if yname == "component" { return "Component" }
    if yname == "client-connection-status" { return "ClientConnectionStatus" }
    if yname == "client-initialization-synchronization-pending" { return "ClientInitializationSynchronizationPending" }
    if yname == "client-synchronization-end-of-download-pending" { return "ClientSynchronizationEndOfDownloadPending" }
    if yname == "up-time-stamp" { return "UpTimeStamp" }
    if yname == "down-time-stamp" { return "DownTimeStamp" }
    if yname == "clean-up-timer-remaining" { return "CleanUpTimerRemaining" }
    return ""
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetSegmentPath() string {
    return "client-status"
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["component"] = clientStatus.Component
    leafs["client-connection-status"] = clientStatus.ClientConnectionStatus
    leafs["client-initialization-synchronization-pending"] = clientStatus.ClientInitializationSynchronizationPending
    leafs["client-synchronization-end-of-download-pending"] = clientStatus.ClientSynchronizationEndOfDownloadPending
    leafs["up-time-stamp"] = clientStatus.UpTimeStamp
    leafs["down-time-stamp"] = clientStatus.DownTimeStamp
    leafs["clean-up-timer-remaining"] = clientStatus.CleanUpTimerRemaining
    return leafs
}

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetBundleName() string { return "cisco_ios_xr" }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetYangName() string { return "client-status" }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) SetParent(parent types.Entity) { clientStatus.parent = parent }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetParent() types.Entity { return clientStatus.parent }

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetParentYangName() string { return "stats-global" }

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus
// Opaque memory Stats
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Component. The type is SergShowComp.
    Component interface{}

    // Session count for each component. The type is interface{} with range:
    // 0..4294967295.
    SessionCount interface{}

    // Current Opaque Memory Size for each component. The type is interface{} with
    // range: 0..4294967295.
    OpaqueSize interface{}

    // High Watermark of Opaque Data Size for each component. The type is
    // interface{} with range: 0..4294967295.
    OpaqueHighSize interface{}

    // Current Opaque Data Size for each component. The type is interface{} with
    // range: 0..4294967295.
    OpaqueDataSize interface{}
}

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetFilter() yfilter.YFilter { return opaqueMemoryStatus.YFilter }

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) SetFilter(yf yfilter.YFilter) { opaqueMemoryStatus.YFilter = yf }

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetGoName(yname string) string {
    if yname == "component" { return "Component" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "opaque-size" { return "OpaqueSize" }
    if yname == "opaque-high-size" { return "OpaqueHighSize" }
    if yname == "opaque-data-size" { return "OpaqueDataSize" }
    return ""
}

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetSegmentPath() string {
    return "opaque-memory-status"
}

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["component"] = opaqueMemoryStatus.Component
    leafs["session-count"] = opaqueMemoryStatus.SessionCount
    leafs["opaque-size"] = opaqueMemoryStatus.OpaqueSize
    leafs["opaque-high-size"] = opaqueMemoryStatus.OpaqueHighSize
    leafs["opaque-data-size"] = opaqueMemoryStatus.OpaqueDataSize
    return leafs
}

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetBundleName() string { return "cisco_ios_xr" }

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetYangName() string { return "opaque-memory-status" }

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) SetParent(parent types.Entity) { opaqueMemoryStatus.parent = parent }

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetParent() types.Entity { return opaqueMemoryStatus.parent }

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetParentYangName() string { return "stats-global" }

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus
// TCP TxList Statistics
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Message Sent Count. The type is interface{} with range: 0..4294967295.
    MessagesSent interface{}

    // Bytes Sent Count. The type is interface{} with range: 0..4294967295. Units
    // are byte.
    BytesSent interface{}

    // Message Received Count. The type is interface{} with range: 0..4294967295.
    MessagesReceived interface{}

    // Bytes Received Count. The type is interface{} with range: 0..4294967295.
    // Units are byte.
    BytesReceived interface{}

    // Socket Connect Count. The type is interface{} with range: 0..4294967295.
    ConnectCount interface{}

    // Blocked Socket Connect Count. The type is interface{} with range:
    // 0..4294967295.
    BlockedConnectCount interface{}

    // Socket Connect Retry Count. The type is interface{} with range:
    // 0..4294967295.
    ConnectRetryCount interface{}

    // Remote Peer DisConnect Count. The type is interface{} with range:
    // 0..4294967295.
    RemoteNodeDownCount interface{}

    // Socket Accept Count. The type is interface{} with range: 0..4294967295.
    AcceptCount interface{}

    // Maximum Size of Sent Message. The type is interface{} with range:
    // 0..4294967295.
    MaximumSentMessageSize interface{}

    // Maximum Size of Received Message. The type is interface{} with range:
    // 0..4294967295.
    MaximumReceivedMessageSize interface{}

    // Peer Pause Count. The type is interface{} with range: 0..4294967295.
    PeerPauseCount interface{}

    // Buffer Full on Write Occurence Count. The type is interface{} with range:
    // 0..4294967295.
    BufferFullOccurenceCount interface{}

    // Partial Message Memory Move Occurence Count. The type is interface{} with
    // range: 0..4294967295.
    MemMoveMessageCount interface{}

    // Partial Message Memory Move Byte Count. The type is interface{} with range:
    // 0..4294967295.
    MemMoveBytesCount interface{}

    // Socket Read Count. The type is interface{} with range: 0..4294967295.
    SocketReadCount interface{}

    // Socket Write Count. The type is interface{} with range: 0..4294967295.
    SocketWriteCount interface{}

    // Active Socket Count. The type is interface{} with range: 0..65535.
    ActiveSocketCount interface{}

    // Maximum Size of Requested Buffer. The type is interface{} with range:
    // 0..4294967295.
    MaximumRequestedBufferSize interface{}

    // Buffer Free Count. The type is interface{} with range: 0..65535.
    BufferFreedCount interface{}

    // Buffer Cache Hit Count. The type is interface{} with range: 0..65535.
    BufferCacheHit interface{}

    // Buffer Cache Miss Count. The type is interface{} with range: 0..65535.
    BufferCacheMiss interface{}
}

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetFilter() yfilter.YFilter { return txListOverTcpStatus.YFilter }

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) SetFilter(yf yfilter.YFilter) { txListOverTcpStatus.YFilter = yf }

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetGoName(yname string) string {
    if yname == "messages-sent" { return "MessagesSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "messages-received" { return "MessagesReceived" }
    if yname == "bytes-received" { return "BytesReceived" }
    if yname == "connect-count" { return "ConnectCount" }
    if yname == "blocked-connect-count" { return "BlockedConnectCount" }
    if yname == "connect-retry-count" { return "ConnectRetryCount" }
    if yname == "remote-node-down-count" { return "RemoteNodeDownCount" }
    if yname == "accept-count" { return "AcceptCount" }
    if yname == "maximum-sent-message-size" { return "MaximumSentMessageSize" }
    if yname == "maximum-received-message-size" { return "MaximumReceivedMessageSize" }
    if yname == "peer-pause-count" { return "PeerPauseCount" }
    if yname == "buffer-full-occurence-count" { return "BufferFullOccurenceCount" }
    if yname == "mem-move-message-count" { return "MemMoveMessageCount" }
    if yname == "mem-move-bytes-count" { return "MemMoveBytesCount" }
    if yname == "socket-read-count" { return "SocketReadCount" }
    if yname == "socket-write-count" { return "SocketWriteCount" }
    if yname == "active-socket-count" { return "ActiveSocketCount" }
    if yname == "maximum-requested-buffer-size" { return "MaximumRequestedBufferSize" }
    if yname == "buffer-freed-count" { return "BufferFreedCount" }
    if yname == "buffer-cache-hit" { return "BufferCacheHit" }
    if yname == "buffer-cache-miss" { return "BufferCacheMiss" }
    return ""
}

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetSegmentPath() string {
    return "tx-list-over-tcp-status"
}

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["messages-sent"] = txListOverTcpStatus.MessagesSent
    leafs["bytes-sent"] = txListOverTcpStatus.BytesSent
    leafs["messages-received"] = txListOverTcpStatus.MessagesReceived
    leafs["bytes-received"] = txListOverTcpStatus.BytesReceived
    leafs["connect-count"] = txListOverTcpStatus.ConnectCount
    leafs["blocked-connect-count"] = txListOverTcpStatus.BlockedConnectCount
    leafs["connect-retry-count"] = txListOverTcpStatus.ConnectRetryCount
    leafs["remote-node-down-count"] = txListOverTcpStatus.RemoteNodeDownCount
    leafs["accept-count"] = txListOverTcpStatus.AcceptCount
    leafs["maximum-sent-message-size"] = txListOverTcpStatus.MaximumSentMessageSize
    leafs["maximum-received-message-size"] = txListOverTcpStatus.MaximumReceivedMessageSize
    leafs["peer-pause-count"] = txListOverTcpStatus.PeerPauseCount
    leafs["buffer-full-occurence-count"] = txListOverTcpStatus.BufferFullOccurenceCount
    leafs["mem-move-message-count"] = txListOverTcpStatus.MemMoveMessageCount
    leafs["mem-move-bytes-count"] = txListOverTcpStatus.MemMoveBytesCount
    leafs["socket-read-count"] = txListOverTcpStatus.SocketReadCount
    leafs["socket-write-count"] = txListOverTcpStatus.SocketWriteCount
    leafs["active-socket-count"] = txListOverTcpStatus.ActiveSocketCount
    leafs["maximum-requested-buffer-size"] = txListOverTcpStatus.MaximumRequestedBufferSize
    leafs["buffer-freed-count"] = txListOverTcpStatus.BufferFreedCount
    leafs["buffer-cache-hit"] = txListOverTcpStatus.BufferCacheHit
    leafs["buffer-cache-miss"] = txListOverTcpStatus.BufferCacheMiss
    return leafs
}

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetBundleName() string { return "cisco_ios_xr" }

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetYangName() string { return "tx-list-over-tcp-status" }

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) SetParent(parent types.Entity) { txListOverTcpStatus.parent = parent }

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetParent() types.Entity { return txListOverTcpStatus.parent }

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetParentYangName() string { return "stats-global" }

// SessionRedundancyAgent_Nodes_Node_GroupSummaries
// Session data for a particular node
type SessionRedundancyAgent_Nodes_Node_GroupSummaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session redundancy agent group summary. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary.
    GroupSummary []SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary
}

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetFilter() yfilter.YFilter { return groupSummaries.YFilter }

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) SetFilter(yf yfilter.YFilter) { groupSummaries.YFilter = yf }

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetGoName(yname string) string {
    if yname == "group-summary" { return "GroupSummary" }
    return ""
}

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetSegmentPath() string {
    return "group-summaries"
}

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-summary" {
        for _, c := range groupSummaries.GroupSummary {
            if groupSummaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary{}
        groupSummaries.GroupSummary = append(groupSummaries.GroupSummary, child)
        return &groupSummaries.GroupSummary[len(groupSummaries.GroupSummary)-1]
    }
    return nil
}

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupSummaries.GroupSummary {
        children[groupSummaries.GroupSummary[i].GetSegmentPath()] = &groupSummaries.GroupSummary[i]
    }
    return children
}

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetBundleName() string { return "cisco_ios_xr" }

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetYangName() string { return "group-summaries" }

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) SetParent(parent types.Entity) { groupSummaries.parent = parent }

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetParent() types.Entity { return groupSummaries.parent }

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetParentYangName() string { return "node" }

// SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary
// Session redundancy agent group summary
type SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. GroupId. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    GroupId interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupIdXr interface{}

    // SERG Role. The type is SergShowImRole.
    Role interface{}

    // Disabled by Config. The type is bool.
    Disabled interface{}

    // Peer IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerIpv4Address interface{}

    // Peer IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerIpv6Address interface{}

    // Peer Status. The type is SergPeerStatus.
    PeerStatus interface{}

    // Preferred Role. The type is SergShowRole.
    PreferredRole interface{}

    // Slave Mode. The type is SergShowSlaveMode.
    SlaveMode interface{}

    // Object Tracking Status (Enabled/Disabled). The type is bool.
    ObjectTrackingStatus interface{}

    // Interface Count. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Session Count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}

    // Pending Session Count for Synchornization. The type is interface{} with
    // range: 0..4294967295.
    PendingAddSessionCount interface{}
}

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetFilter() yfilter.YFilter { return groupSummary.YFilter }

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) SetFilter(yf yfilter.YFilter) { groupSummary.YFilter = yf }

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "group-id-xr" { return "GroupIdXr" }
    if yname == "role" { return "Role" }
    if yname == "disabled" { return "Disabled" }
    if yname == "peer-ipv4-address" { return "PeerIpv4Address" }
    if yname == "peer-ipv6-address" { return "PeerIpv6Address" }
    if yname == "peer-status" { return "PeerStatus" }
    if yname == "preferred-role" { return "PreferredRole" }
    if yname == "slave-mode" { return "SlaveMode" }
    if yname == "object-tracking-status" { return "ObjectTrackingStatus" }
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "pending-add-session-count" { return "PendingAddSessionCount" }
    return ""
}

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetSegmentPath() string {
    return "group-summary" + "[group-id='" + fmt.Sprintf("%v", groupSummary.GroupId) + "']"
}

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = groupSummary.GroupId
    leafs["group-id-xr"] = groupSummary.GroupIdXr
    leafs["role"] = groupSummary.Role
    leafs["disabled"] = groupSummary.Disabled
    leafs["peer-ipv4-address"] = groupSummary.PeerIpv4Address
    leafs["peer-ipv6-address"] = groupSummary.PeerIpv6Address
    leafs["peer-status"] = groupSummary.PeerStatus
    leafs["preferred-role"] = groupSummary.PreferredRole
    leafs["slave-mode"] = groupSummary.SlaveMode
    leafs["object-tracking-status"] = groupSummary.ObjectTrackingStatus
    leafs["interface-count"] = groupSummary.InterfaceCount
    leafs["session-count"] = groupSummary.SessionCount
    leafs["pending-add-session-count"] = groupSummary.PendingAddSessionCount
    return leafs
}

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetBundleName() string { return "cisco_ios_xr" }

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetYangName() string { return "group-summary" }

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) SetParent(parent types.Entity) { groupSummary.parent = parent }

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetParent() types.Entity { return groupSummary.parent }

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetParentYangName() string { return "group-summaries" }

