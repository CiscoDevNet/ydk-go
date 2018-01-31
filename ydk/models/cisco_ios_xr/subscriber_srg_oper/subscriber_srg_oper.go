// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-srg package operational data.
// 
// This module contains definitions
// for the following management objects:
//   subscriber-redundancy-manager: Subscriber Redundancy Manager
//     information
//   subscriber-redundancy-agent: subscriber redundancy agent
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_srg_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_srg_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-srg-oper subscriber-redundancy-manager}", reflect.TypeOf(SubscriberRedundancyManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-srg-oper:subscriber-redundancy-manager", reflect.TypeOf(SubscriberRedundancyManager{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-srg-oper subscriber-redundancy-agent}", reflect.TypeOf(SubscriberRedundancyAgent{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-srg-oper:subscriber-redundancy-agent", reflect.TypeOf(SubscriberRedundancyAgent{}))
}

// SrgShowImRole represents SRG Interface Management Role
type SrgShowImRole string

const (
    // Not Determined
    SrgShowImRole_none SrgShowImRole = "none"

    // Master Role
    SrgShowImRole_master SrgShowImRole = "master"

    // Slave Role
    SrgShowImRole_slave SrgShowImRole = "slave"
)

// SrgShowSessionError represents SRG Session Error Operation
type SrgShowSessionError string

const (
    // Invalid Error
    SrgShowSessionError_none SrgShowSessionError = "none"

    // Hard Error
    SrgShowSessionError_hard SrgShowSessionError = "hard"

    // Soft Error
    SrgShowSessionError_soft SrgShowSessionError = "soft"
)

// SrgShowSoReason represents Subscriber Redundancy Switchover Reason
type SrgShowSoReason string

const (
    // SwitchOver for Internal Reason
    SrgShowSoReason_internal SrgShowSoReason = "internal"

    // SwitchOver for Admin
    SrgShowSoReason_admin SrgShowSoReason = "admin"

    // SwitchOver for Peer UP
    SrgShowSoReason_peer_up SrgShowSoReason = "peer-up"

    // SwitchOver for Peer Down
    SrgShowSoReason_peer_down SrgShowSoReason = "peer-down"

    // SwitchOver for Object Tracking status change
    SrgShowSoReason_object_tracking_status_change SrgShowSoReason = "object-tracking-status-change"

    // Unknown Switchover Reason
    SrgShowSoReason_srg_show_so_reason_max SrgShowSoReason = "srg-show-so-reason-max"
)

// SrgShowSessionOperation represents SRG Session Operation
type SrgShowSessionOperation string

const (
    // No Operation
    SrgShowSessionOperation_none SrgShowSessionOperation = "none"

    // SRG Session Update Operation
    SrgShowSessionOperation_update SrgShowSessionOperation = "update"

    // SRG Session Delete Operation
    SrgShowSessionOperation_delete SrgShowSessionOperation = "delete"

    // SRG Session In Sync
    SrgShowSessionOperation_in_sync SrgShowSessionOperation = "in-sync"
)

// SrgPeerStatus represents SRG Peer Status
type SrgPeerStatus string

const (
    // Peer not configured
    SrgPeerStatus_not_configured SrgPeerStatus = "not-configured"

    // Peer initialization
    SrgPeerStatus_initialize SrgPeerStatus = "initialize"

    // Peer retry pending
    SrgPeerStatus_retry SrgPeerStatus = "retry"

    // Connection in Progress
    SrgPeerStatus_connect SrgPeerStatus = "connect"

    // Listening in Progress
    SrgPeerStatus_listen SrgPeerStatus = "listen"

    // Awaiting Registration from Peer
    SrgPeerStatus_registration SrgPeerStatus = "registration"

    // Peer cleanup in progress
    SrgPeerStatus_cleanup SrgPeerStatus = "cleanup"

    // Peer Connected
    SrgPeerStatus_connected SrgPeerStatus = "connected"

    // Peer Established
    SrgPeerStatus_established SrgPeerStatus = "established"
)

// SrgShowRole represents SRG Role
type SrgShowRole string

const (
    // Not Configured
    SrgShowRole_none SrgShowRole = "none"

    // Master Role
    SrgShowRole_master SrgShowRole = "master"

    // Slave Role
    SrgShowRole_slave SrgShowRole = "slave"
)

// SrgShowSlaveMode represents SRG Slave Mode
type SrgShowSlaveMode string

const (
    // Not Configured
    SrgShowSlaveMode_none SrgShowSlaveMode = "none"

    // Warm Modem
    SrgShowSlaveMode_warm SrgShowSlaveMode = "warm"

    // Hot Mode
    SrgShowSlaveMode_hot SrgShowSlaveMode = "hot"
)

// SrgShowComp represents SRG Components
type SrgShowComp string

const (
    // SRG Agent
    SrgShowComp_srga SrgShowComp = "srga"

    // DHCPv4
    SrgShowComp_dhcpv4 SrgShowComp = "dhcpv4"

    // DHCPv6
    SrgShowComp_dhcpv6 SrgShowComp = "dhcpv6"

    // PPPoE
    SrgShowComp_pppoe SrgShowComp = "pppoe"

    // PPP
    SrgShowComp_ppp SrgShowComp = "ppp"

    // L2TP
    SrgShowComp_l2tp SrgShowComp = "l2tp"

    // iEdge
    SrgShowComp_iedge SrgShowComp = "iedge"
)

// SubscriberRedundancyManager
// Subscriber Redundancy Manager information
type SubscriberRedundancyManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber Redundancy Manager group table.
    Groups SubscriberRedundancyManager_Groups

    // Subscriber redundancy manager summary.
    Summary SubscriberRedundancyManager_Summary

    // Subscriber Redundancy Manager interface table.
    Interfaces SubscriberRedundancyManager_Interfaces
}

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetFilter() yfilter.YFilter { return subscriberRedundancyManager.YFilter }

func (subscriberRedundancyManager *SubscriberRedundancyManager) SetFilter(yf yfilter.YFilter) { subscriberRedundancyManager.YFilter = yf }

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetGoName(yname string) string {
    if yname == "groups" { return "Groups" }
    if yname == "summary" { return "Summary" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-srg-oper:subscriber-redundancy-manager"
}

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &subscriberRedundancyManager.Groups
    }
    if childYangName == "summary" {
        return &subscriberRedundancyManager.Summary
    }
    if childYangName == "interfaces" {
        return &subscriberRedundancyManager.Interfaces
    }
    return nil
}

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &subscriberRedundancyManager.Groups
    children["summary"] = &subscriberRedundancyManager.Summary
    children["interfaces"] = &subscriberRedundancyManager.Interfaces
    return children
}

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetYangName() string { return "subscriber-redundancy-manager" }

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberRedundancyManager *SubscriberRedundancyManager) SetParent(parent types.Entity) { subscriberRedundancyManager.parent = parent }

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetParent() types.Entity { return subscriberRedundancyManager.parent }

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-srg-oper" }

// SubscriberRedundancyManager_Groups
// Subscriber Redundancy Manager group table
type SubscriberRedundancyManager_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber redundancy manager group. The type is slice of
    // SubscriberRedundancyManager_Groups_Group.
    Group []SubscriberRedundancyManager_Groups_Group
}

func (groups *SubscriberRedundancyManager_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *SubscriberRedundancyManager_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *SubscriberRedundancyManager_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *SubscriberRedundancyManager_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *SubscriberRedundancyManager_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyManager_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *SubscriberRedundancyManager_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *SubscriberRedundancyManager_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *SubscriberRedundancyManager_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *SubscriberRedundancyManager_Groups) GetYangName() string { return "groups" }

func (groups *SubscriberRedundancyManager_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *SubscriberRedundancyManager_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *SubscriberRedundancyManager_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *SubscriberRedundancyManager_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *SubscriberRedundancyManager_Groups) GetParent() types.Entity { return groups.parent }

func (groups *SubscriberRedundancyManager_Groups) GetParentYangName() string { return "subscriber-redundancy-manager" }

// SubscriberRedundancyManager_Groups_Group
// Subscriber redundancy manager group
type SubscriberRedundancyManager_Groups_Group struct {
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

    // SRG Role. The type is SrgShowImRole.
    Role interface{}

    // Peer IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerIpv4Address interface{}

    // Peer IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerIpv6Address interface{}

    // Interface Count. The type is interface{} with range: 0..4294967295.
    InterfaceCount interface{}

    // Preferred Role. The type is SrgShowRole.
    PreferredRole interface{}

    // Slave Mode. The type is SrgShowSlaveMode.
    SlaveMode interface{}

    // Object Tracking Status (Enabled/Disabled). The type is bool.
    ObjectTrackingStatus interface{}

    // Virtual MAC Address. The type is string.
    VirtualMacAddress interface{}

    // Virtual MAC Address Disable. The type is bool.
    VirtualMacAddressDisable interface{}

    // Node Information. The type is string.
    NodeName interface{}
}

func (group *SubscriberRedundancyManager_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *SubscriberRedundancyManager_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *SubscriberRedundancyManager_Groups_Group) GetGoName(yname string) string {
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
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "virtual-mac-address-disable" { return "VirtualMacAddressDisable" }
    if yname == "node-name" { return "NodeName" }
    return ""
}

func (group *SubscriberRedundancyManager_Groups_Group) GetSegmentPath() string {
    return "group" + "[group='" + fmt.Sprintf("%v", group.Group) + "']"
}

func (group *SubscriberRedundancyManager_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (group *SubscriberRedundancyManager_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (group *SubscriberRedundancyManager_Groups_Group) GetLeafs() map[string]interface{} {
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
    leafs["virtual-mac-address"] = group.VirtualMacAddress
    leafs["virtual-mac-address-disable"] = group.VirtualMacAddressDisable
    leafs["node-name"] = group.NodeName
    return leafs
}

func (group *SubscriberRedundancyManager_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *SubscriberRedundancyManager_Groups_Group) GetYangName() string { return "group" }

func (group *SubscriberRedundancyManager_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *SubscriberRedundancyManager_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *SubscriberRedundancyManager_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *SubscriberRedundancyManager_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *SubscriberRedundancyManager_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *SubscriberRedundancyManager_Groups_Group) GetParentYangName() string { return "groups" }

// SubscriberRedundancyManager_Summary
// Subscriber redundancy manager summary
type SubscriberRedundancyManager_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disabled by Config. The type is bool.
    Disabled interface{}

    // Process Active State. The type is bool.
    ActiveState interface{}

    // Preferred Role. The type is SrgShowRole.
    PreferredRole interface{}

    // Slave Mode. The type is SrgShowSlaveMode.
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

func (summary *SubscriberRedundancyManager_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *SubscriberRedundancyManager_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *SubscriberRedundancyManager_Summary) GetGoName(yname string) string {
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

func (summary *SubscriberRedundancyManager_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *SubscriberRedundancyManager_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *SubscriberRedundancyManager_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *SubscriberRedundancyManager_Summary) GetLeafs() map[string]interface{} {
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

func (summary *SubscriberRedundancyManager_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *SubscriberRedundancyManager_Summary) GetYangName() string { return "summary" }

func (summary *SubscriberRedundancyManager_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *SubscriberRedundancyManager_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *SubscriberRedundancyManager_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *SubscriberRedundancyManager_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *SubscriberRedundancyManager_Summary) GetParent() types.Entity { return summary.parent }

func (summary *SubscriberRedundancyManager_Summary) GetParentYangName() string { return "subscriber-redundancy-manager" }

// SubscriberRedundancyManager_Interfaces
// Subscriber Redundancy Manager interface table
type SubscriberRedundancyManager_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber redundancy manager interface. The type is slice of
    // SubscriberRedundancyManager_Interfaces_Interface.
    Interface []SubscriberRedundancyManager_Interfaces_Interface
}

func (interfaces *SubscriberRedundancyManager_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *SubscriberRedundancyManager_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *SubscriberRedundancyManager_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *SubscriberRedundancyManager_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *SubscriberRedundancyManager_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyManager_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *SubscriberRedundancyManager_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *SubscriberRedundancyManager_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *SubscriberRedundancyManager_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *SubscriberRedundancyManager_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *SubscriberRedundancyManager_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *SubscriberRedundancyManager_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *SubscriberRedundancyManager_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *SubscriberRedundancyManager_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *SubscriberRedundancyManager_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *SubscriberRedundancyManager_Interfaces) GetParentYangName() string { return "subscriber-redundancy-manager" }

// SubscriberRedundancyManager_Interfaces_Interface
// Subscriber redundancy manager interface
type SubscriberRedundancyManager_Interfaces_Interface struct {
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

    // SRG Role. The type is SrgShowImRole.
    Role interface{}
}

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *SubscriberRedundancyManager_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-mapping-id" { return "InterfaceMappingId" }
    if yname == "forward-referenced" { return "ForwardReferenced" }
    if yname == "group-id" { return "GroupId" }
    if yname == "role" { return "Role" }
    return ""
}

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = self.Interface
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-mapping-id"] = self.InterfaceMappingId
    leafs["forward-referenced"] = self.ForwardReferenced
    leafs["group-id"] = self.GroupId
    leafs["role"] = self.Role
    return leafs
}

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *SubscriberRedundancyManager_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// SubscriberRedundancyAgent
// subscriber redundancy agent
type SubscriberRedundancyAgent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes for which subscriber data is collected.
    Nodes SubscriberRedundancyAgent_Nodes
}

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetFilter() yfilter.YFilter { return subscriberRedundancyAgent.YFilter }

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) SetFilter(yf yfilter.YFilter) { subscriberRedundancyAgent.YFilter = yf }

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-srg-oper:subscriber-redundancy-agent"
}

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &subscriberRedundancyAgent.Nodes
    }
    return nil
}

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &subscriberRedundancyAgent.Nodes
    return children
}

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetYangName() string { return "subscriber-redundancy-agent" }

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) SetParent(parent types.Entity) { subscriberRedundancyAgent.parent = parent }

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetParent() types.Entity { return subscriberRedundancyAgent.parent }

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-srg-oper" }

// SubscriberRedundancyAgent_Nodes
// List of nodes for which subscriber data is
// collected
type SubscriberRedundancyAgent_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber data for a particular node. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node.
    Node []SubscriberRedundancyAgent_Nodes_Node
}

func (nodes *SubscriberRedundancyAgent_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *SubscriberRedundancyAgent_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *SubscriberRedundancyAgent_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *SubscriberRedundancyAgent_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *SubscriberRedundancyAgent_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyAgent_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *SubscriberRedundancyAgent_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *SubscriberRedundancyAgent_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *SubscriberRedundancyAgent_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *SubscriberRedundancyAgent_Nodes) GetYangName() string { return "nodes" }

func (nodes *SubscriberRedundancyAgent_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *SubscriberRedundancyAgent_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *SubscriberRedundancyAgent_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *SubscriberRedundancyAgent_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *SubscriberRedundancyAgent_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *SubscriberRedundancyAgent_Nodes) GetParentYangName() string { return "subscriber-redundancy-agent" }

// SubscriberRedundancyAgent_Nodes_Node
// Subscriber data for a particular node
type SubscriberRedundancyAgent_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Data for particular subscriber group session.
    GroupIdXr SubscriberRedundancyAgent_Nodes_Node_GroupIdXr

    // List of interfaces.
    Interfaces SubscriberRedundancyAgent_Nodes_Node_Interfaces

    // Subscriber data for a particular node.
    GroupSummaries SubscriberRedundancyAgent_Nodes_Node_GroupSummaries

    // Data for particular subscriber group .
    GroupIds SubscriberRedundancyAgent_Nodes_Node_GroupIds
}

func (node *SubscriberRedundancyAgent_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *SubscriberRedundancyAgent_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *SubscriberRedundancyAgent_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "group-id-xr" { return "GroupIdXr" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "group-summaries" { return "GroupSummaries" }
    if yname == "group-ids" { return "GroupIds" }
    return ""
}

func (node *SubscriberRedundancyAgent_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *SubscriberRedundancyAgent_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-id-xr" {
        return &node.GroupIdXr
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "group-summaries" {
        return &node.GroupSummaries
    }
    if childYangName == "group-ids" {
        return &node.GroupIds
    }
    return nil
}

func (node *SubscriberRedundancyAgent_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["group-id-xr"] = &node.GroupIdXr
    children["interfaces"] = &node.Interfaces
    children["group-summaries"] = &node.GroupSummaries
    children["group-ids"] = &node.GroupIds
    return children
}

func (node *SubscriberRedundancyAgent_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *SubscriberRedundancyAgent_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *SubscriberRedundancyAgent_Nodes_Node) GetYangName() string { return "node" }

func (node *SubscriberRedundancyAgent_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *SubscriberRedundancyAgent_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *SubscriberRedundancyAgent_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *SubscriberRedundancyAgent_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *SubscriberRedundancyAgent_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *SubscriberRedundancyAgent_Nodes_Node) GetParentYangName() string { return "nodes" }

// SubscriberRedundancyAgent_Nodes_Node_GroupIdXr
// Data for particular subscriber group session
type SubscriberRedundancyAgent_Nodes_Node_GroupIdXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group id for subscriber group session. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId.
    GroupId []SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId
}

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetFilter() yfilter.YFilter { return groupIdXr.YFilter }

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) SetFilter(yf yfilter.YFilter) { groupIdXr.YFilter = yf }

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    return ""
}

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetSegmentPath() string {
    return "group-id-xr"
}

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-id" {
        for _, c := range groupIdXr.GroupId {
            if groupIdXr.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId{}
        groupIdXr.GroupId = append(groupIdXr.GroupId, child)
        return &groupIdXr.GroupId[len(groupIdXr.GroupId)-1]
    }
    return nil
}

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupIdXr.GroupId {
        children[groupIdXr.GroupId[i].GetSegmentPath()] = &groupIdXr.GroupId[i]
    }
    return children
}

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetBundleName() string { return "cisco_ios_xr" }

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetYangName() string { return "group-id-xr" }

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) SetParent(parent types.Entity) { groupIdXr.parent = parent }

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetParent() types.Entity { return groupIdXr.parent }

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetParentYangName() string { return "node" }

// SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId
// Group id for subscriber group session
type SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. GroupId. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    GroupId interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupIdXr interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Outer VLAN Information. The type is interface{} with range: 0..4294967295.
    OuterVlan interface{}

    // Inner VLAN Information. The type is interface{} with range: 0..4294967295.
    InnerVlan interface{}

    // Session MAC Address. The type is string.
    SessionMacAddress interface{}

    // PPPoE Session ID. The type is interface{} with range: 0..65535.
    PppoeSessionId interface{}

    // L2TP Tunnel local ID. The type is interface{} with range: 0..4294967295.
    L2TpTunnelId interface{}

    // Master Role is Set. The type is bool.
    RoleMaster interface{}

    // Holds a Valid MAC Address. The type is bool.
    ValidMacAddress interface{}

    // Negative Acknowledgement Update Flag is Set. The type is bool.
    NegativeAcknowledgementUpdateAll interface{}

    // More Session Information. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation.
    SessionDetailedInformation []SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation

    // Session Synchroniation Error Information. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation.
    SessionSyncErrorInformation []SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetFilter() yfilter.YFilter { return groupId.YFilter }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) SetFilter(yf yfilter.YFilter) { groupId.YFilter = yf }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "group-id-xr" { return "GroupIdXr" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "outer-vlan" { return "OuterVlan" }
    if yname == "inner-vlan" { return "InnerVlan" }
    if yname == "session-mac-address" { return "SessionMacAddress" }
    if yname == "pppoe-session-id" { return "PppoeSessionId" }
    if yname == "l2tp-tunnel-id" { return "L2TpTunnelId" }
    if yname == "role-master" { return "RoleMaster" }
    if yname == "valid-mac-address" { return "ValidMacAddress" }
    if yname == "negative-acknowledgement-update-all" { return "NegativeAcknowledgementUpdateAll" }
    if yname == "session-detailed-information" { return "SessionDetailedInformation" }
    if yname == "session-sync-error-information" { return "SessionSyncErrorInformation" }
    return ""
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetSegmentPath() string {
    return "group-id" + "[group-id='" + fmt.Sprintf("%v", groupId.GroupId) + "']"
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-detailed-information" {
        for _, c := range groupId.SessionDetailedInformation {
            if groupId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation{}
        groupId.SessionDetailedInformation = append(groupId.SessionDetailedInformation, child)
        return &groupId.SessionDetailedInformation[len(groupId.SessionDetailedInformation)-1]
    }
    if childYangName == "session-sync-error-information" {
        for _, c := range groupId.SessionSyncErrorInformation {
            if groupId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation{}
        groupId.SessionSyncErrorInformation = append(groupId.SessionSyncErrorInformation, child)
        return &groupId.SessionSyncErrorInformation[len(groupId.SessionSyncErrorInformation)-1]
    }
    return nil
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupId.SessionDetailedInformation {
        children[groupId.SessionDetailedInformation[i].GetSegmentPath()] = &groupId.SessionDetailedInformation[i]
    }
    for i := range groupId.SessionSyncErrorInformation {
        children[groupId.SessionSyncErrorInformation[i].GetSegmentPath()] = &groupId.SessionSyncErrorInformation[i]
    }
    return children
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = groupId.GroupId
    leafs["group-id-xr"] = groupId.GroupIdXr
    leafs["interface-name"] = groupId.InterfaceName
    leafs["outer-vlan"] = groupId.OuterVlan
    leafs["inner-vlan"] = groupId.InnerVlan
    leafs["session-mac-address"] = groupId.SessionMacAddress
    leafs["pppoe-session-id"] = groupId.PppoeSessionId
    leafs["l2tp-tunnel-id"] = groupId.L2TpTunnelId
    leafs["role-master"] = groupId.RoleMaster
    leafs["valid-mac-address"] = groupId.ValidMacAddress
    leafs["negative-acknowledgement-update-all"] = groupId.NegativeAcknowledgementUpdateAll
    return leafs
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetBundleName() string { return "cisco_ios_xr" }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetYangName() string { return "group-id" }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) SetParent(parent types.Entity) { groupId.parent = parent }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetParent() types.Entity { return groupId.parent }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetParentYangName() string { return "group-id-xr" }

// SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation
// More Session Information
type SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Component. The type is SrgShowComp.
    Component interface{}

    // Operation Code. The type is SrgShowSessionOperation.
    Operation interface{}

    // Tx List Queue Failed. The type is bool.
    TxListQueueFail interface{}

    // Marked For Sweeping. The type is bool.
    MarkedForSweeping interface{}

    // Marked For Cleanup. The type is bool.
    MarkedForCleanup interface{}
}

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetFilter() yfilter.YFilter { return sessionDetailedInformation.YFilter }

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) SetFilter(yf yfilter.YFilter) { sessionDetailedInformation.YFilter = yf }

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetGoName(yname string) string {
    if yname == "component" { return "Component" }
    if yname == "operation" { return "Operation" }
    if yname == "tx-list-queue-fail" { return "TxListQueueFail" }
    if yname == "marked-for-sweeping" { return "MarkedForSweeping" }
    if yname == "marked-for-cleanup" { return "MarkedForCleanup" }
    return ""
}

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetSegmentPath() string {
    return "session-detailed-information"
}

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["component"] = sessionDetailedInformation.Component
    leafs["operation"] = sessionDetailedInformation.Operation
    leafs["tx-list-queue-fail"] = sessionDetailedInformation.TxListQueueFail
    leafs["marked-for-sweeping"] = sessionDetailedInformation.MarkedForSweeping
    leafs["marked-for-cleanup"] = sessionDetailedInformation.MarkedForCleanup
    return leafs
}

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetBundleName() string { return "cisco_ios_xr" }

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetYangName() string { return "session-detailed-information" }

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) SetParent(parent types.Entity) { sessionDetailedInformation.parent = parent }

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetParent() types.Entity { return sessionDetailedInformation.parent }

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetParentYangName() string { return "group-id" }

// SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation
// Session Synchroniation Error Information
type SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // No. of Errors occured during Synchronization. The type is interface{} with
    // range: 0..65535.
    SyncErrorCount interface{}

    // Last Error Code. The type is interface{} with range: 0..4294967295.
    LastErrorCode interface{}

    // Last Error Type. The type is SrgShowSessionError.
    LastErrorType interface{}
}

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetFilter() yfilter.YFilter { return sessionSyncErrorInformation.YFilter }

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) SetFilter(yf yfilter.YFilter) { sessionSyncErrorInformation.YFilter = yf }

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetGoName(yname string) string {
    if yname == "sync-error-count" { return "SyncErrorCount" }
    if yname == "last-error-code" { return "LastErrorCode" }
    if yname == "last-error-type" { return "LastErrorType" }
    return ""
}

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetSegmentPath() string {
    return "session-sync-error-information"
}

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sync-error-count"] = sessionSyncErrorInformation.SyncErrorCount
    leafs["last-error-code"] = sessionSyncErrorInformation.LastErrorCode
    leafs["last-error-type"] = sessionSyncErrorInformation.LastErrorType
    return leafs
}

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetBundleName() string { return "cisco_ios_xr" }

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetYangName() string { return "session-sync-error-information" }

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) SetParent(parent types.Entity) { sessionSyncErrorInformation.parent = parent }

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetParent() types.Entity { return sessionSyncErrorInformation.parent }

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetParentYangName() string { return "group-id" }

// SubscriberRedundancyAgent_Nodes_Node_Interfaces
// List of interfaces
type SubscriberRedundancyAgent_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify interface name. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface.
    Interface []SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface
}

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface
// Specify interface name
type SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface struct {
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

    // SRG Role. The type is SrgShowImRole.
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
    InterfaceOper SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper

    // Interface Status.
    InterfaceStatus SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus

    // Interface status for each client. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus.
    ClientStatus []SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus
}

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
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

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
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
        child := SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus{}
        self.ClientStatus = append(self.ClientStatus, child)
        return &self.ClientStatus[len(self.ClientStatus)-1]
    }
    return nil
}

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-oper"] = &self.InterfaceOper
    children["interface-status"] = &self.InterfaceStatus
    for i := range self.ClientStatus {
        children[self.ClientStatus[i].GetSegmentPath()] = &self.ClientStatus[i]
    }
    return children
}

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
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

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper
// Interface Batch Operation
type SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper struct {
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

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetFilter() yfilter.YFilter { return interfaceOper.YFilter }

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) SetFilter(yf yfilter.YFilter) { interfaceOper.YFilter = yf }

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetGoName(yname string) string {
    if yname == "idb-oper-reg-enable" { return "IdbOperRegEnable" }
    if yname == "idb-oper-reg-disable" { return "IdbOperRegDisable" }
    if yname == "idb-oper-caps-add" { return "IdbOperCapsAdd" }
    if yname == "idb-oper-caps-remove" { return "IdbOperCapsRemove" }
    if yname == "idb-oper-attr-update" { return "IdbOperAttrUpdate" }
    return ""
}

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetSegmentPath() string {
    return "interface-oper"
}

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["idb-oper-reg-enable"] = interfaceOper.IdbOperRegEnable
    leafs["idb-oper-reg-disable"] = interfaceOper.IdbOperRegDisable
    leafs["idb-oper-caps-add"] = interfaceOper.IdbOperCapsAdd
    leafs["idb-oper-caps-remove"] = interfaceOper.IdbOperCapsRemove
    leafs["idb-oper-attr-update"] = interfaceOper.IdbOperAttrUpdate
    return leafs
}

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetYangName() string { return "interface-oper" }

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) SetParent(parent types.Entity) { interfaceOper.parent = parent }

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetParent() types.Entity { return interfaceOper.parent }

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetParentYangName() string { return "interface" }

// SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus
// Interface Status
type SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus struct {
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

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetFilter() yfilter.YFilter { return interfaceStatus.YFilter }

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) SetFilter(yf yfilter.YFilter) { interfaceStatus.YFilter = yf }

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetGoName(yname string) string {
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

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetSegmentPath() string {
    return "interface-status"
}

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetLeafs() map[string]interface{} {
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

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetYangName() string { return "interface-status" }

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) SetParent(parent types.Entity) { interfaceStatus.parent = parent }

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetParent() types.Entity { return interfaceStatus.parent }

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetParentYangName() string { return "interface" }

// SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus
// Interface status for each client
type SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Component. The type is SrgShowComp.
    Component interface{}

    // SRG SHOW IDB CLIENT EOMS PENDING. The type is bool.
    SrgShowIdbClientEomsPending interface{}

    // SRG SHOW IDB CLIENT SYNC EOD PENDING. The type is bool.
    SrgShowIdbClientSyncEodPending interface{}

    // session count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}
}

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetFilter() yfilter.YFilter { return clientStatus.YFilter }

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) SetFilter(yf yfilter.YFilter) { clientStatus.YFilter = yf }

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetGoName(yname string) string {
    if yname == "component" { return "Component" }
    if yname == "srg-show-idb-client-eoms-pending" { return "SrgShowIdbClientEomsPending" }
    if yname == "srg-show-idb-client-sync-eod-pending" { return "SrgShowIdbClientSyncEodPending" }
    if yname == "session-count" { return "SessionCount" }
    return ""
}

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetSegmentPath() string {
    return "client-status"
}

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["component"] = clientStatus.Component
    leafs["srg-show-idb-client-eoms-pending"] = clientStatus.SrgShowIdbClientEomsPending
    leafs["srg-show-idb-client-sync-eod-pending"] = clientStatus.SrgShowIdbClientSyncEodPending
    leafs["session-count"] = clientStatus.SessionCount
    return leafs
}

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetBundleName() string { return "cisco_ios_xr" }

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetYangName() string { return "client-status" }

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) SetParent(parent types.Entity) { clientStatus.parent = parent }

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetParent() types.Entity { return clientStatus.parent }

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetParentYangName() string { return "interface" }

// SubscriberRedundancyAgent_Nodes_Node_GroupSummaries
// Subscriber data for a particular node
type SubscriberRedundancyAgent_Nodes_Node_GroupSummaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber redundancy agent group summary. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary.
    GroupSummary []SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary
}

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetFilter() yfilter.YFilter { return groupSummaries.YFilter }

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) SetFilter(yf yfilter.YFilter) { groupSummaries.YFilter = yf }

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetGoName(yname string) string {
    if yname == "group-summary" { return "GroupSummary" }
    return ""
}

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetSegmentPath() string {
    return "group-summaries"
}

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-summary" {
        for _, c := range groupSummaries.GroupSummary {
            if groupSummaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary{}
        groupSummaries.GroupSummary = append(groupSummaries.GroupSummary, child)
        return &groupSummaries.GroupSummary[len(groupSummaries.GroupSummary)-1]
    }
    return nil
}

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupSummaries.GroupSummary {
        children[groupSummaries.GroupSummary[i].GetSegmentPath()] = &groupSummaries.GroupSummary[i]
    }
    return children
}

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetBundleName() string { return "cisco_ios_xr" }

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetYangName() string { return "group-summaries" }

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) SetParent(parent types.Entity) { groupSummaries.parent = parent }

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetParent() types.Entity { return groupSummaries.parent }

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetParentYangName() string { return "node" }

// SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary
// Subscriber redundancy agent group summary
type SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. GroupId. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    GroupId interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupIdXr interface{}

    // SRG Role. The type is SrgShowImRole.
    Role interface{}

    // Disabled by Config. The type is bool.
    Disabled interface{}

    // Peer IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerIpv4Address interface{}

    // Peer IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerIpv6Address interface{}

    // Peer Status. The type is SrgPeerStatus.
    PeerStatus interface{}

    // Preferred Role. The type is SrgShowRole.
    PreferredRole interface{}

    // Slave Mode. The type is SrgShowSlaveMode.
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

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetFilter() yfilter.YFilter { return groupSummary.YFilter }

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) SetFilter(yf yfilter.YFilter) { groupSummary.YFilter = yf }

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetGoName(yname string) string {
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

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetSegmentPath() string {
    return "group-summary" + "[group-id='" + fmt.Sprintf("%v", groupSummary.GroupId) + "']"
}

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetLeafs() map[string]interface{} {
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

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetBundleName() string { return "cisco_ios_xr" }

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetYangName() string { return "group-summary" }

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) SetParent(parent types.Entity) { groupSummary.parent = parent }

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetParent() types.Entity { return groupSummary.parent }

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetParentYangName() string { return "group-summaries" }

// SubscriberRedundancyAgent_Nodes_Node_GroupIds
// Data for particular subscriber group 
type SubscriberRedundancyAgent_Nodes_Node_GroupIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group id for subscriber group. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId.
    GroupId []SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId
}

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetFilter() yfilter.YFilter { return groupIds.YFilter }

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) SetFilter(yf yfilter.YFilter) { groupIds.YFilter = yf }

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    return ""
}

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetSegmentPath() string {
    return "group-ids"
}

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-id" {
        for _, c := range groupIds.GroupId {
            if groupIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId{}
        groupIds.GroupId = append(groupIds.GroupId, child)
        return &groupIds.GroupId[len(groupIds.GroupId)-1]
    }
    return nil
}

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupIds.GroupId {
        children[groupIds.GroupId[i].GetSegmentPath()] = &groupIds.GroupId[i]
    }
    return children
}

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetBundleName() string { return "cisco_ios_xr" }

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetYangName() string { return "group-ids" }

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) SetParent(parent types.Entity) { groupIds.parent = parent }

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetParent() types.Entity { return groupIds.parent }

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetParentYangName() string { return "node" }

// SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId
// Group id for subscriber group
type SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId struct {
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

    // Preferred Init Role. The type is SrgShowRole.
    InitRole interface{}

    // Negotiating Role. The type is SrgShowRole.
    NegotiatingRole interface{}

    // Current Role. The type is SrgShowRole.
    CurrentRole interface{}

    // Slave Mode. The type is SrgShowSlaveMode.
    SlaveMode interface{}

    // Switch Over Hold Time. The type is interface{} with range: 0..4294967295.
    HoldTimer interface{}

    // Virtual MAC Address. The type is string.
    VirtualMacAddress interface{}

    // Virtual MAC Address Disable. The type is bool.
    VirtualMacAddressDisable interface{}

    // L2TP Souce IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    L2TpSourceIp interface{}

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

    // Peer Status. The type is SrgPeerStatus.
    PeerStatus interface{}

    // Last Negotiation time of Peer. The type is string.
    PeerLastNegotiationTime interface{}

    // Last UP time of Peer. The type is string.
    PeerLastUpTime interface{}

    // Last Down time of Peer. The type is string.
    PeerLastDownTime interface{}

    // Peer Preferred Init Role. The type is SrgShowRole.
    PeerInitRole interface{}

    // Peer Negotiating Role. The type is SrgShowRole.
    PeerNegotiatingRole interface{}

    // Peer Current Role. The type is SrgShowRole.
    PeerCurrentRole interface{}

    // Peer Object Tracking Status. The type is bool.
    PeerObjectTrackingStatus interface{}

    // Last Switchover time. The type is string.
    LastSwitchoverTime interface{}

    // Switchover Count. The type is interface{} with range: 0..4294967295.
    SwitchoverCount interface{}

    // Last Switchover Reason. The type is SrgShowSoReason.
    LastSwitchoverReason interface{}

    // Switchover Hold Time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SwitchoverHoldTime interface{}

    // Session Count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}

    // Slave Session update fail count. The type is interface{} with range:
    // 0..4294967295.
    SlaveUpdateFailureCount interface{}

    // Tunnel Count. The type is interface{} with range: 0..4294967295.
    TunnelCount interface{}

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

    // Interface List. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface.
    Interface []SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetFilter() yfilter.YFilter { return groupId.YFilter }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) SetFilter(yf yfilter.YFilter) { groupId.YFilter = yf }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "group-id-xr" { return "GroupIdXr" }
    if yname == "description" { return "Description" }
    if yname == "disabled" { return "Disabled" }
    if yname == "init-role" { return "InitRole" }
    if yname == "negotiating-role" { return "NegotiatingRole" }
    if yname == "current-role" { return "CurrentRole" }
    if yname == "slave-mode" { return "SlaveMode" }
    if yname == "hold-timer" { return "HoldTimer" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "virtual-mac-address-disable" { return "VirtualMacAddressDisable" }
    if yname == "l2tp-source-ip" { return "L2TpSourceIp" }
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
    if yname == "tunnel-count" { return "TunnelCount" }
    if yname == "pending-session-update-count" { return "PendingSessionUpdateCount" }
    if yname == "pending-session-delete-count" { return "PendingSessionDeleteCount" }
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "revertive-timer" { return "RevertiveTimer" }
    if yname == "switchover-revert-time" { return "SwitchoverRevertTime" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetSegmentPath() string {
    return "group-id" + "[group-id='" + fmt.Sprintf("%v", groupId.GroupId) + "']"
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range groupId.Interface {
            if groupId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface{}
        groupId.Interface = append(groupId.Interface, child)
        return &groupId.Interface[len(groupId.Interface)-1]
    }
    return nil
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupId.Interface {
        children[groupId.Interface[i].GetSegmentPath()] = &groupId.Interface[i]
    }
    return children
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetLeafs() map[string]interface{} {
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
    leafs["virtual-mac-address"] = groupId.VirtualMacAddress
    leafs["virtual-mac-address-disable"] = groupId.VirtualMacAddressDisable
    leafs["l2tp-source-ip"] = groupId.L2TpSourceIp
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
    leafs["tunnel-count"] = groupId.TunnelCount
    leafs["pending-session-update-count"] = groupId.PendingSessionUpdateCount
    leafs["pending-session-delete-count"] = groupId.PendingSessionDeleteCount
    leafs["interface-count"] = groupId.InterfaceCount
    leafs["revertive-timer"] = groupId.RevertiveTimer
    leafs["switchover-revert-time"] = groupId.SwitchoverRevertTime
    return leafs
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetBundleName() string { return "cisco_ios_xr" }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetYangName() string { return "group-id" }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) SetParent(parent types.Entity) { groupId.parent = parent }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetParent() types.Entity { return groupId.parent }

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetParentYangName() string { return "group-ids" }

// SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface
// Interface List
type SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface struct {
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

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-synchronization-id" { return "InterfaceSynchronizationId" }
    if yname == "forward-referenced" { return "ForwardReferenced" }
    if yname == "session-count" { return "SessionCount" }
    return ""
}

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-synchronization-id"] = self.InterfaceSynchronizationId
    leafs["forward-referenced"] = self.ForwardReferenced
    leafs["session-count"] = self.SessionCount
    return leafs
}

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetYangName() string { return "interface" }

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetParent() types.Entity { return self.parent }

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetParentYangName() string { return "group-id" }

