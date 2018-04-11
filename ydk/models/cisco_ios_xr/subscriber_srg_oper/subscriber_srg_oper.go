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

// SubscriberRedundancyManager
// Subscriber Redundancy Manager information
type SubscriberRedundancyManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber Redundancy Manager group table.
    Groups SubscriberRedundancyManager_Groups

    // Subscriber redundancy manager summary.
    Summary SubscriberRedundancyManager_Summary

    // Subscriber Redundancy Manager interface table.
    Interfaces SubscriberRedundancyManager_Interfaces
}

func (subscriberRedundancyManager *SubscriberRedundancyManager) GetEntityData() *types.CommonEntityData {
    subscriberRedundancyManager.EntityData.YFilter = subscriberRedundancyManager.YFilter
    subscriberRedundancyManager.EntityData.YangName = "subscriber-redundancy-manager"
    subscriberRedundancyManager.EntityData.BundleName = "cisco_ios_xr"
    subscriberRedundancyManager.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-srg-oper"
    subscriberRedundancyManager.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-srg-oper:subscriber-redundancy-manager"
    subscriberRedundancyManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberRedundancyManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberRedundancyManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberRedundancyManager.EntityData.Children = make(map[string]types.YChild)
    subscriberRedundancyManager.EntityData.Children["groups"] = types.YChild{"Groups", &subscriberRedundancyManager.Groups}
    subscriberRedundancyManager.EntityData.Children["summary"] = types.YChild{"Summary", &subscriberRedundancyManager.Summary}
    subscriberRedundancyManager.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &subscriberRedundancyManager.Interfaces}
    subscriberRedundancyManager.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(subscriberRedundancyManager.EntityData)
}

// SubscriberRedundancyManager_Groups
// Subscriber Redundancy Manager group table
type SubscriberRedundancyManager_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber redundancy manager group. The type is slice of
    // SubscriberRedundancyManager_Groups_Group.
    Group []SubscriberRedundancyManager_Groups_Group
}

func (groups *SubscriberRedundancyManager_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "subscriber-redundancy-manager"
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

// SubscriberRedundancyManager_Groups_Group
// Subscriber redundancy manager group
type SubscriberRedundancyManager_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerIpv4Address interface{}

    // Peer IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (group *SubscriberRedundancyManager_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[group='" + fmt.Sprintf("%v", group.Group) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["group"] = types.YLeaf{"Group", group.Group}
    group.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", group.GroupId}
    group.EntityData.Leafs["description"] = types.YLeaf{"Description", group.Description}
    group.EntityData.Leafs["disabled"] = types.YLeaf{"Disabled", group.Disabled}
    group.EntityData.Leafs["role"] = types.YLeaf{"Role", group.Role}
    group.EntityData.Leafs["peer-ipv4-address"] = types.YLeaf{"PeerIpv4Address", group.PeerIpv4Address}
    group.EntityData.Leafs["peer-ipv6-address"] = types.YLeaf{"PeerIpv6Address", group.PeerIpv6Address}
    group.EntityData.Leafs["interface-count"] = types.YLeaf{"InterfaceCount", group.InterfaceCount}
    group.EntityData.Leafs["preferred-role"] = types.YLeaf{"PreferredRole", group.PreferredRole}
    group.EntityData.Leafs["slave-mode"] = types.YLeaf{"SlaveMode", group.SlaveMode}
    group.EntityData.Leafs["object-tracking-status"] = types.YLeaf{"ObjectTrackingStatus", group.ObjectTrackingStatus}
    group.EntityData.Leafs["virtual-mac-address"] = types.YLeaf{"VirtualMacAddress", group.VirtualMacAddress}
    group.EntityData.Leafs["virtual-mac-address-disable"] = types.YLeaf{"VirtualMacAddressDisable", group.VirtualMacAddressDisable}
    group.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", group.NodeName}
    return &(group.EntityData)
}

// SubscriberRedundancyManager_Summary
// Subscriber redundancy manager summary
type SubscriberRedundancyManager_Summary struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceInterfaceIpv4Address interface{}

    // Source Interface IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (summary *SubscriberRedundancyManager_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "subscriber-redundancy-manager"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    summary.EntityData.Leafs["disabled"] = types.YLeaf{"Disabled", summary.Disabled}
    summary.EntityData.Leafs["active-state"] = types.YLeaf{"ActiveState", summary.ActiveState}
    summary.EntityData.Leafs["preferred-role"] = types.YLeaf{"PreferredRole", summary.PreferredRole}
    summary.EntityData.Leafs["slave-mode"] = types.YLeaf{"SlaveMode", summary.SlaveMode}
    summary.EntityData.Leafs["hold-timer"] = types.YLeaf{"HoldTimer", summary.HoldTimer}
    summary.EntityData.Leafs["source-interface-name"] = types.YLeaf{"SourceInterfaceName", summary.SourceInterfaceName}
    summary.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", summary.VrfName}
    summary.EntityData.Leafs["source-interface-ipv4-address"] = types.YLeaf{"SourceInterfaceIpv4Address", summary.SourceInterfaceIpv4Address}
    summary.EntityData.Leafs["source-interface-ipv6-address"] = types.YLeaf{"SourceInterfaceIpv6Address", summary.SourceInterfaceIpv6Address}
    summary.EntityData.Leafs["group-count"] = types.YLeaf{"GroupCount", summary.GroupCount}
    summary.EntityData.Leafs["disabled-group-count"] = types.YLeaf{"DisabledGroupCount", summary.DisabledGroupCount}
    summary.EntityData.Leafs["master-group-count"] = types.YLeaf{"MasterGroupCount", summary.MasterGroupCount}
    summary.EntityData.Leafs["slave-group-count"] = types.YLeaf{"SlaveGroupCount", summary.SlaveGroupCount}
    summary.EntityData.Leafs["interface-count"] = types.YLeaf{"InterfaceCount", summary.InterfaceCount}
    summary.EntityData.Leafs["master-interface-count"] = types.YLeaf{"MasterInterfaceCount", summary.MasterInterfaceCount}
    summary.EntityData.Leafs["slave-interface-count"] = types.YLeaf{"SlaveInterfaceCount", summary.SlaveInterfaceCount}
    return &(summary.EntityData)
}

// SubscriberRedundancyManager_Interfaces
// Subscriber Redundancy Manager interface table
type SubscriberRedundancyManager_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber redundancy manager interface. The type is slice of
    // SubscriberRedundancyManager_Interfaces_Interface_.
    Interface_ []SubscriberRedundancyManager_Interfaces_Interface
}

func (interfaces *SubscriberRedundancyManager_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "subscriber-redundancy-manager"
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

// SubscriberRedundancyManager_Interfaces_Interface
// Subscriber redundancy manager interface
type SubscriberRedundancyManager_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Interface_ interface{}

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

func (self *SubscriberRedundancyManager_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface_) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-mapping-id"] = types.YLeaf{"InterfaceMappingId", self.InterfaceMappingId}
    self.EntityData.Leafs["forward-referenced"] = types.YLeaf{"ForwardReferenced", self.ForwardReferenced}
    self.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", self.GroupId}
    self.EntityData.Leafs["role"] = types.YLeaf{"Role", self.Role}
    return &(self.EntityData)
}

// SubscriberRedundancyAgent
// subscriber redundancy agent
type SubscriberRedundancyAgent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes for which subscriber data is collected.
    Nodes SubscriberRedundancyAgent_Nodes
}

func (subscriberRedundancyAgent *SubscriberRedundancyAgent) GetEntityData() *types.CommonEntityData {
    subscriberRedundancyAgent.EntityData.YFilter = subscriberRedundancyAgent.YFilter
    subscriberRedundancyAgent.EntityData.YangName = "subscriber-redundancy-agent"
    subscriberRedundancyAgent.EntityData.BundleName = "cisco_ios_xr"
    subscriberRedundancyAgent.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-srg-oper"
    subscriberRedundancyAgent.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-srg-oper:subscriber-redundancy-agent"
    subscriberRedundancyAgent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberRedundancyAgent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberRedundancyAgent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberRedundancyAgent.EntityData.Children = make(map[string]types.YChild)
    subscriberRedundancyAgent.EntityData.Children["nodes"] = types.YChild{"Nodes", &subscriberRedundancyAgent.Nodes}
    subscriberRedundancyAgent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(subscriberRedundancyAgent.EntityData)
}

// SubscriberRedundancyAgent_Nodes
// List of nodes for which subscriber data is
// collected
type SubscriberRedundancyAgent_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber data for a particular node. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node.
    Node []SubscriberRedundancyAgent_Nodes_Node
}

func (nodes *SubscriberRedundancyAgent_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "subscriber-redundancy-agent"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node
// Subscriber data for a particular node
type SubscriberRedundancyAgent_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (node *SubscriberRedundancyAgent_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["group-id-xr"] = types.YChild{"GroupIdXr", &node.GroupIdXr}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Children["group-summaries"] = types.YChild{"GroupSummaries", &node.GroupSummaries}
    node.EntityData.Children["group-ids"] = types.YChild{"GroupIds", &node.GroupIds}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_GroupIdXr
// Data for particular subscriber group session
type SubscriberRedundancyAgent_Nodes_Node_GroupIdXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group id for subscriber group session. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId.
    GroupId []SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId
}

func (groupIdXr *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr) GetEntityData() *types.CommonEntityData {
    groupIdXr.EntityData.YFilter = groupIdXr.YFilter
    groupIdXr.EntityData.YangName = "group-id-xr"
    groupIdXr.EntityData.BundleName = "cisco_ios_xr"
    groupIdXr.EntityData.ParentYangName = "node"
    groupIdXr.EntityData.SegmentPath = "group-id-xr"
    groupIdXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupIdXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupIdXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupIdXr.EntityData.Children = make(map[string]types.YChild)
    groupIdXr.EntityData.Children["group-id"] = types.YChild{"GroupId", nil}
    for i := range groupIdXr.GroupId {
        groupIdXr.EntityData.Children[types.GetSegmentPath(&groupIdXr.GroupId[i])] = types.YChild{"GroupId", &groupIdXr.GroupId[i]}
    }
    groupIdXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groupIdXr.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId
// Group id for subscriber group session
type SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. GroupId. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetEntityData() *types.CommonEntityData {
    groupId.EntityData.YFilter = groupId.YFilter
    groupId.EntityData.YangName = "group-id"
    groupId.EntityData.BundleName = "cisco_ios_xr"
    groupId.EntityData.ParentYangName = "group-id-xr"
    groupId.EntityData.SegmentPath = "group-id" + "[group-id='" + fmt.Sprintf("%v", groupId.GroupId) + "']"
    groupId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupId.EntityData.Children = make(map[string]types.YChild)
    groupId.EntityData.Children["session-detailed-information"] = types.YChild{"SessionDetailedInformation", nil}
    for i := range groupId.SessionDetailedInformation {
        groupId.EntityData.Children[types.GetSegmentPath(&groupId.SessionDetailedInformation[i])] = types.YChild{"SessionDetailedInformation", &groupId.SessionDetailedInformation[i]}
    }
    groupId.EntityData.Children["session-sync-error-information"] = types.YChild{"SessionSyncErrorInformation", nil}
    for i := range groupId.SessionSyncErrorInformation {
        groupId.EntityData.Children[types.GetSegmentPath(&groupId.SessionSyncErrorInformation[i])] = types.YChild{"SessionSyncErrorInformation", &groupId.SessionSyncErrorInformation[i]}
    }
    groupId.EntityData.Leafs = make(map[string]types.YLeaf)
    groupId.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", groupId.GroupId}
    groupId.EntityData.Leafs["group-id-xr"] = types.YLeaf{"GroupIdXr", groupId.GroupIdXr}
    groupId.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", groupId.InterfaceName}
    groupId.EntityData.Leafs["outer-vlan"] = types.YLeaf{"OuterVlan", groupId.OuterVlan}
    groupId.EntityData.Leafs["inner-vlan"] = types.YLeaf{"InnerVlan", groupId.InnerVlan}
    groupId.EntityData.Leafs["session-mac-address"] = types.YLeaf{"SessionMacAddress", groupId.SessionMacAddress}
    groupId.EntityData.Leafs["pppoe-session-id"] = types.YLeaf{"PppoeSessionId", groupId.PppoeSessionId}
    groupId.EntityData.Leafs["l2tp-tunnel-id"] = types.YLeaf{"L2TpTunnelId", groupId.L2TpTunnelId}
    groupId.EntityData.Leafs["role-master"] = types.YLeaf{"RoleMaster", groupId.RoleMaster}
    groupId.EntityData.Leafs["valid-mac-address"] = types.YLeaf{"ValidMacAddress", groupId.ValidMacAddress}
    groupId.EntityData.Leafs["negative-acknowledgement-update-all"] = types.YLeaf{"NegativeAcknowledgementUpdateAll", groupId.NegativeAcknowledgementUpdateAll}
    return &(groupId.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation
// More Session Information
type SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation struct {
    EntityData types.CommonEntityData
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

func (sessionDetailedInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetEntityData() *types.CommonEntityData {
    sessionDetailedInformation.EntityData.YFilter = sessionDetailedInformation.YFilter
    sessionDetailedInformation.EntityData.YangName = "session-detailed-information"
    sessionDetailedInformation.EntityData.BundleName = "cisco_ios_xr"
    sessionDetailedInformation.EntityData.ParentYangName = "group-id"
    sessionDetailedInformation.EntityData.SegmentPath = "session-detailed-information"
    sessionDetailedInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDetailedInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDetailedInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDetailedInformation.EntityData.Children = make(map[string]types.YChild)
    sessionDetailedInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionDetailedInformation.EntityData.Leafs["component"] = types.YLeaf{"Component", sessionDetailedInformation.Component}
    sessionDetailedInformation.EntityData.Leafs["operation"] = types.YLeaf{"Operation", sessionDetailedInformation.Operation}
    sessionDetailedInformation.EntityData.Leafs["tx-list-queue-fail"] = types.YLeaf{"TxListQueueFail", sessionDetailedInformation.TxListQueueFail}
    sessionDetailedInformation.EntityData.Leafs["marked-for-sweeping"] = types.YLeaf{"MarkedForSweeping", sessionDetailedInformation.MarkedForSweeping}
    sessionDetailedInformation.EntityData.Leafs["marked-for-cleanup"] = types.YLeaf{"MarkedForCleanup", sessionDetailedInformation.MarkedForCleanup}
    return &(sessionDetailedInformation.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation
// Session Synchroniation Error Information
type SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // No. of Errors occured during Synchronization. The type is interface{} with
    // range: 0..65535.
    SyncErrorCount interface{}

    // Last Error Code. The type is interface{} with range: 0..4294967295.
    LastErrorCode interface{}

    // Last Error Type. The type is SrgShowSessionError.
    LastErrorType interface{}
}

func (sessionSyncErrorInformation *SubscriberRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetEntityData() *types.CommonEntityData {
    sessionSyncErrorInformation.EntityData.YFilter = sessionSyncErrorInformation.YFilter
    sessionSyncErrorInformation.EntityData.YangName = "session-sync-error-information"
    sessionSyncErrorInformation.EntityData.BundleName = "cisco_ios_xr"
    sessionSyncErrorInformation.EntityData.ParentYangName = "group-id"
    sessionSyncErrorInformation.EntityData.SegmentPath = "session-sync-error-information"
    sessionSyncErrorInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionSyncErrorInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionSyncErrorInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionSyncErrorInformation.EntityData.Children = make(map[string]types.YChild)
    sessionSyncErrorInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionSyncErrorInformation.EntityData.Leafs["sync-error-count"] = types.YLeaf{"SyncErrorCount", sessionSyncErrorInformation.SyncErrorCount}
    sessionSyncErrorInformation.EntityData.Leafs["last-error-code"] = types.YLeaf{"LastErrorCode", sessionSyncErrorInformation.LastErrorCode}
    sessionSyncErrorInformation.EntityData.Leafs["last-error-type"] = types.YLeaf{"LastErrorType", sessionSyncErrorInformation.LastErrorType}
    return &(sessionSyncErrorInformation.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_Interfaces
// List of interfaces
type SubscriberRedundancyAgent_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify interface name. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_.
    Interface_ []SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface
}

func (interfaces *SubscriberRedundancyAgent_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
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

// SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface
// Specify interface name
type SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

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

func (self *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface_) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["interface-oper"] = types.YChild{"InterfaceOper", &self.InterfaceOper}
    self.EntityData.Children["interface-status"] = types.YChild{"InterfaceStatus", &self.InterfaceStatus}
    self.EntityData.Children["client-status"] = types.YChild{"ClientStatus", nil}
    for i := range self.ClientStatus {
        self.EntityData.Children[types.GetSegmentPath(&self.ClientStatus[i])] = types.YChild{"ClientStatus", &self.ClientStatus[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-synchronization-id"] = types.YLeaf{"InterfaceSynchronizationId", self.InterfaceSynchronizationId}
    self.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", self.GroupId}
    self.EntityData.Leafs["role"] = types.YLeaf{"Role", self.Role}
    self.EntityData.Leafs["forward-referenced"] = types.YLeaf{"ForwardReferenced", self.ForwardReferenced}
    self.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", self.SessionCount}
    self.EntityData.Leafs["interface-enable-error-count"] = types.YLeaf{"InterfaceEnableErrorCount", self.InterfaceEnableErrorCount}
    self.EntityData.Leafs["interface-disable-error-count"] = types.YLeaf{"InterfaceDisableErrorCount", self.InterfaceDisableErrorCount}
    self.EntityData.Leafs["interface-caps-add-error-count"] = types.YLeaf{"InterfaceCapsAddErrorCount", self.InterfaceCapsAddErrorCount}
    self.EntityData.Leafs["interface-caps-remove-error-count"] = types.YLeaf{"InterfaceCapsRemoveErrorCount", self.InterfaceCapsRemoveErrorCount}
    self.EntityData.Leafs["interface-attribute-update-error-count"] = types.YLeaf{"InterfaceAttributeUpdateErrorCount", self.InterfaceAttributeUpdateErrorCount}
    return &(self.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper
// Interface Batch Operation
type SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper struct {
    EntityData types.CommonEntityData
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

func (interfaceOper *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetEntityData() *types.CommonEntityData {
    interfaceOper.EntityData.YFilter = interfaceOper.YFilter
    interfaceOper.EntityData.YangName = "interface-oper"
    interfaceOper.EntityData.BundleName = "cisco_ios_xr"
    interfaceOper.EntityData.ParentYangName = "interface"
    interfaceOper.EntityData.SegmentPath = "interface-oper"
    interfaceOper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceOper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceOper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceOper.EntityData.Children = make(map[string]types.YChild)
    interfaceOper.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceOper.EntityData.Leafs["idb-oper-reg-enable"] = types.YLeaf{"IdbOperRegEnable", interfaceOper.IdbOperRegEnable}
    interfaceOper.EntityData.Leafs["idb-oper-reg-disable"] = types.YLeaf{"IdbOperRegDisable", interfaceOper.IdbOperRegDisable}
    interfaceOper.EntityData.Leafs["idb-oper-caps-add"] = types.YLeaf{"IdbOperCapsAdd", interfaceOper.IdbOperCapsAdd}
    interfaceOper.EntityData.Leafs["idb-oper-caps-remove"] = types.YLeaf{"IdbOperCapsRemove", interfaceOper.IdbOperCapsRemove}
    interfaceOper.EntityData.Leafs["idb-oper-attr-update"] = types.YLeaf{"IdbOperAttrUpdate", interfaceOper.IdbOperAttrUpdate}
    return &(interfaceOper.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus
// Interface Status
type SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus struct {
    EntityData types.CommonEntityData
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

func (interfaceStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetEntityData() *types.CommonEntityData {
    interfaceStatus.EntityData.YFilter = interfaceStatus.YFilter
    interfaceStatus.EntityData.YangName = "interface-status"
    interfaceStatus.EntityData.BundleName = "cisco_ios_xr"
    interfaceStatus.EntityData.ParentYangName = "interface"
    interfaceStatus.EntityData.SegmentPath = "interface-status"
    interfaceStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStatus.EntityData.Children = make(map[string]types.YChild)
    interfaceStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceStatus.EntityData.Leafs["idb-state-fwd-ref"] = types.YLeaf{"IdbStateFwdRef", interfaceStatus.IdbStateFwdRef}
    interfaceStatus.EntityData.Leafs["idb-state-stale"] = types.YLeaf{"IdbStateStale", interfaceStatus.IdbStateStale}
    interfaceStatus.EntityData.Leafs["idb-state-registered"] = types.YLeaf{"IdbStateRegistered", interfaceStatus.IdbStateRegistered}
    interfaceStatus.EntityData.Leafs["idb-state-caps-added"] = types.YLeaf{"IdbStateCapsAdded", interfaceStatus.IdbStateCapsAdded}
    interfaceStatus.EntityData.Leafs["idb-state-owned-re-source"] = types.YLeaf{"IdbStateOwnedReSource", interfaceStatus.IdbStateOwnedReSource}
    interfaceStatus.EntityData.Leafs["idb-client-eoms-pending"] = types.YLeaf{"IdbClientEomsPending", interfaceStatus.IdbClientEomsPending}
    interfaceStatus.EntityData.Leafs["idb-state-p-end-caps-rem"] = types.YLeaf{"IdbStatePEndCapsRem", interfaceStatus.IdbStatePEndCapsRem}
    interfaceStatus.EntityData.Leafs["idb-state-p-end-reg-disable"] = types.YLeaf{"IdbStatePEndRegDisable", interfaceStatus.IdbStatePEndRegDisable}
    return &(interfaceStatus.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus
// Interface status for each client
type SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus struct {
    EntityData types.CommonEntityData
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

func (clientStatus *SubscriberRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetEntityData() *types.CommonEntityData {
    clientStatus.EntityData.YFilter = clientStatus.YFilter
    clientStatus.EntityData.YangName = "client-status"
    clientStatus.EntityData.BundleName = "cisco_ios_xr"
    clientStatus.EntityData.ParentYangName = "interface"
    clientStatus.EntityData.SegmentPath = "client-status"
    clientStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientStatus.EntityData.Children = make(map[string]types.YChild)
    clientStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    clientStatus.EntityData.Leafs["component"] = types.YLeaf{"Component", clientStatus.Component}
    clientStatus.EntityData.Leafs["srg-show-idb-client-eoms-pending"] = types.YLeaf{"SrgShowIdbClientEomsPending", clientStatus.SrgShowIdbClientEomsPending}
    clientStatus.EntityData.Leafs["srg-show-idb-client-sync-eod-pending"] = types.YLeaf{"SrgShowIdbClientSyncEodPending", clientStatus.SrgShowIdbClientSyncEodPending}
    clientStatus.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", clientStatus.SessionCount}
    return &(clientStatus.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_GroupSummaries
// Subscriber data for a particular node
type SubscriberRedundancyAgent_Nodes_Node_GroupSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber redundancy agent group summary. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary.
    GroupSummary []SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary
}

func (groupSummaries *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries) GetEntityData() *types.CommonEntityData {
    groupSummaries.EntityData.YFilter = groupSummaries.YFilter
    groupSummaries.EntityData.YangName = "group-summaries"
    groupSummaries.EntityData.BundleName = "cisco_ios_xr"
    groupSummaries.EntityData.ParentYangName = "node"
    groupSummaries.EntityData.SegmentPath = "group-summaries"
    groupSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupSummaries.EntityData.Children = make(map[string]types.YChild)
    groupSummaries.EntityData.Children["group-summary"] = types.YChild{"GroupSummary", nil}
    for i := range groupSummaries.GroupSummary {
        groupSummaries.EntityData.Children[types.GetSegmentPath(&groupSummaries.GroupSummary[i])] = types.YChild{"GroupSummary", &groupSummaries.GroupSummary[i]}
    }
    groupSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groupSummaries.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary
// Subscriber redundancy agent group summary
type SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. GroupId. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    GroupId interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupIdXr interface{}

    // SRG Role. The type is SrgShowImRole.
    Role interface{}

    // Disabled by Config. The type is bool.
    Disabled interface{}

    // Peer IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerIpv4Address interface{}

    // Peer IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (groupSummary *SubscriberRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetEntityData() *types.CommonEntityData {
    groupSummary.EntityData.YFilter = groupSummary.YFilter
    groupSummary.EntityData.YangName = "group-summary"
    groupSummary.EntityData.BundleName = "cisco_ios_xr"
    groupSummary.EntityData.ParentYangName = "group-summaries"
    groupSummary.EntityData.SegmentPath = "group-summary" + "[group-id='" + fmt.Sprintf("%v", groupSummary.GroupId) + "']"
    groupSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupSummary.EntityData.Children = make(map[string]types.YChild)
    groupSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    groupSummary.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", groupSummary.GroupId}
    groupSummary.EntityData.Leafs["group-id-xr"] = types.YLeaf{"GroupIdXr", groupSummary.GroupIdXr}
    groupSummary.EntityData.Leafs["role"] = types.YLeaf{"Role", groupSummary.Role}
    groupSummary.EntityData.Leafs["disabled"] = types.YLeaf{"Disabled", groupSummary.Disabled}
    groupSummary.EntityData.Leafs["peer-ipv4-address"] = types.YLeaf{"PeerIpv4Address", groupSummary.PeerIpv4Address}
    groupSummary.EntityData.Leafs["peer-ipv6-address"] = types.YLeaf{"PeerIpv6Address", groupSummary.PeerIpv6Address}
    groupSummary.EntityData.Leafs["peer-status"] = types.YLeaf{"PeerStatus", groupSummary.PeerStatus}
    groupSummary.EntityData.Leafs["preferred-role"] = types.YLeaf{"PreferredRole", groupSummary.PreferredRole}
    groupSummary.EntityData.Leafs["slave-mode"] = types.YLeaf{"SlaveMode", groupSummary.SlaveMode}
    groupSummary.EntityData.Leafs["object-tracking-status"] = types.YLeaf{"ObjectTrackingStatus", groupSummary.ObjectTrackingStatus}
    groupSummary.EntityData.Leafs["interface-count"] = types.YLeaf{"InterfaceCount", groupSummary.InterfaceCount}
    groupSummary.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", groupSummary.SessionCount}
    groupSummary.EntityData.Leafs["pending-add-session-count"] = types.YLeaf{"PendingAddSessionCount", groupSummary.PendingAddSessionCount}
    return &(groupSummary.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_GroupIds
// Data for particular subscriber group 
type SubscriberRedundancyAgent_Nodes_Node_GroupIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group id for subscriber group. The type is slice of
    // SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId.
    GroupId []SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId
}

func (groupIds *SubscriberRedundancyAgent_Nodes_Node_GroupIds) GetEntityData() *types.CommonEntityData {
    groupIds.EntityData.YFilter = groupIds.YFilter
    groupIds.EntityData.YangName = "group-ids"
    groupIds.EntityData.BundleName = "cisco_ios_xr"
    groupIds.EntityData.ParentYangName = "node"
    groupIds.EntityData.SegmentPath = "group-ids"
    groupIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupIds.EntityData.Children = make(map[string]types.YChild)
    groupIds.EntityData.Children["group-id"] = types.YChild{"GroupId", nil}
    for i := range groupIds.GroupId {
        groupIds.EntityData.Children[types.GetSegmentPath(&groupIds.GroupId[i])] = types.YChild{"GroupId", &groupIds.GroupId[i]}
    }
    groupIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groupIds.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId
// Group id for subscriber group
type SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group Id. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerIpv4Address interface{}

    // Peer IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    // SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface_.
    Interface_ []SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface
}

func (groupId *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetEntityData() *types.CommonEntityData {
    groupId.EntityData.YFilter = groupId.YFilter
    groupId.EntityData.YangName = "group-id"
    groupId.EntityData.BundleName = "cisco_ios_xr"
    groupId.EntityData.ParentYangName = "group-ids"
    groupId.EntityData.SegmentPath = "group-id" + "[group-id='" + fmt.Sprintf("%v", groupId.GroupId) + "']"
    groupId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupId.EntityData.Children = make(map[string]types.YChild)
    groupId.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range groupId.Interface_ {
        groupId.EntityData.Children[types.GetSegmentPath(&groupId.Interface_[i])] = types.YChild{"Interface_", &groupId.Interface_[i]}
    }
    groupId.EntityData.Leafs = make(map[string]types.YLeaf)
    groupId.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", groupId.GroupId}
    groupId.EntityData.Leafs["group-id-xr"] = types.YLeaf{"GroupIdXr", groupId.GroupIdXr}
    groupId.EntityData.Leafs["description"] = types.YLeaf{"Description", groupId.Description}
    groupId.EntityData.Leafs["disabled"] = types.YLeaf{"Disabled", groupId.Disabled}
    groupId.EntityData.Leafs["init-role"] = types.YLeaf{"InitRole", groupId.InitRole}
    groupId.EntityData.Leafs["negotiating-role"] = types.YLeaf{"NegotiatingRole", groupId.NegotiatingRole}
    groupId.EntityData.Leafs["current-role"] = types.YLeaf{"CurrentRole", groupId.CurrentRole}
    groupId.EntityData.Leafs["slave-mode"] = types.YLeaf{"SlaveMode", groupId.SlaveMode}
    groupId.EntityData.Leafs["hold-timer"] = types.YLeaf{"HoldTimer", groupId.HoldTimer}
    groupId.EntityData.Leafs["virtual-mac-address"] = types.YLeaf{"VirtualMacAddress", groupId.VirtualMacAddress}
    groupId.EntityData.Leafs["virtual-mac-address-disable"] = types.YLeaf{"VirtualMacAddressDisable", groupId.VirtualMacAddressDisable}
    groupId.EntityData.Leafs["l2tp-source-ip"] = types.YLeaf{"L2TpSourceIp", groupId.L2TpSourceIp}
    groupId.EntityData.Leafs["core-tracking-object-name"] = types.YLeaf{"CoreTrackingObjectName", groupId.CoreTrackingObjectName}
    groupId.EntityData.Leafs["core-tracking-object-status"] = types.YLeaf{"CoreTrackingObjectStatus", groupId.CoreTrackingObjectStatus}
    groupId.EntityData.Leafs["access-tracking-object-name"] = types.YLeaf{"AccessTrackingObjectName", groupId.AccessTrackingObjectName}
    groupId.EntityData.Leafs["access-tracking-object-status"] = types.YLeaf{"AccessTrackingObjectStatus", groupId.AccessTrackingObjectStatus}
    groupId.EntityData.Leafs["object-tracking-status"] = types.YLeaf{"ObjectTrackingStatus", groupId.ObjectTrackingStatus}
    groupId.EntityData.Leafs["peer-ipv4-address"] = types.YLeaf{"PeerIpv4Address", groupId.PeerIpv4Address}
    groupId.EntityData.Leafs["peer-ipv6-address"] = types.YLeaf{"PeerIpv6Address", groupId.PeerIpv6Address}
    groupId.EntityData.Leafs["peer-status"] = types.YLeaf{"PeerStatus", groupId.PeerStatus}
    groupId.EntityData.Leafs["peer-last-negotiation-time"] = types.YLeaf{"PeerLastNegotiationTime", groupId.PeerLastNegotiationTime}
    groupId.EntityData.Leafs["peer-last-up-time"] = types.YLeaf{"PeerLastUpTime", groupId.PeerLastUpTime}
    groupId.EntityData.Leafs["peer-last-down-time"] = types.YLeaf{"PeerLastDownTime", groupId.PeerLastDownTime}
    groupId.EntityData.Leafs["peer-init-role"] = types.YLeaf{"PeerInitRole", groupId.PeerInitRole}
    groupId.EntityData.Leafs["peer-negotiating-role"] = types.YLeaf{"PeerNegotiatingRole", groupId.PeerNegotiatingRole}
    groupId.EntityData.Leafs["peer-current-role"] = types.YLeaf{"PeerCurrentRole", groupId.PeerCurrentRole}
    groupId.EntityData.Leafs["peer-object-tracking-status"] = types.YLeaf{"PeerObjectTrackingStatus", groupId.PeerObjectTrackingStatus}
    groupId.EntityData.Leafs["last-switchover-time"] = types.YLeaf{"LastSwitchoverTime", groupId.LastSwitchoverTime}
    groupId.EntityData.Leafs["switchover-count"] = types.YLeaf{"SwitchoverCount", groupId.SwitchoverCount}
    groupId.EntityData.Leafs["last-switchover-reason"] = types.YLeaf{"LastSwitchoverReason", groupId.LastSwitchoverReason}
    groupId.EntityData.Leafs["switchover-hold-time"] = types.YLeaf{"SwitchoverHoldTime", groupId.SwitchoverHoldTime}
    groupId.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", groupId.SessionCount}
    groupId.EntityData.Leafs["slave-update-failure-count"] = types.YLeaf{"SlaveUpdateFailureCount", groupId.SlaveUpdateFailureCount}
    groupId.EntityData.Leafs["tunnel-count"] = types.YLeaf{"TunnelCount", groupId.TunnelCount}
    groupId.EntityData.Leafs["pending-session-update-count"] = types.YLeaf{"PendingSessionUpdateCount", groupId.PendingSessionUpdateCount}
    groupId.EntityData.Leafs["pending-session-delete-count"] = types.YLeaf{"PendingSessionDeleteCount", groupId.PendingSessionDeleteCount}
    groupId.EntityData.Leafs["interface-count"] = types.YLeaf{"InterfaceCount", groupId.InterfaceCount}
    groupId.EntityData.Leafs["revertive-timer"] = types.YLeaf{"RevertiveTimer", groupId.RevertiveTimer}
    groupId.EntityData.Leafs["switchover-revert-time"] = types.YLeaf{"SwitchoverRevertTime", groupId.SwitchoverRevertTime}
    return &(groupId.EntityData)
}

// SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface
// Interface List
type SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface struct {
    EntityData types.CommonEntityData
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

func (self *SubscriberRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "group-id"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface-synchronization-id"] = types.YLeaf{"InterfaceSynchronizationId", self.InterfaceSynchronizationId}
    self.EntityData.Leafs["forward-referenced"] = types.YLeaf{"ForwardReferenced", self.ForwardReferenced}
    self.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", self.SessionCount}
    return &(self.EntityData)
}

