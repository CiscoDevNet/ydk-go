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
    SergShowMem_string_ SergShowMem = "string"

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
    SergShowSessionOperation_delete_ SergShowSessionOperation = "delete"

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber Redundancy Manager interface table.
    Interfaces SessionRedundancyManager_Interfaces

    // Session Redundancy Manager group table.
    Groups SessionRedundancyManager_Groups

    // Session redundancy manager summary.
    Summary SessionRedundancyManager_Summary
}

func (sessionRedundancyManager *SessionRedundancyManager) GetEntityData() *types.CommonEntityData {
    sessionRedundancyManager.EntityData.YFilter = sessionRedundancyManager.YFilter
    sessionRedundancyManager.EntityData.YangName = "session-redundancy-manager"
    sessionRedundancyManager.EntityData.BundleName = "cisco_ios_xr"
    sessionRedundancyManager.EntityData.ParentYangName = "Cisco-IOS-XR-infra-serg-oper"
    sessionRedundancyManager.EntityData.SegmentPath = "Cisco-IOS-XR-infra-serg-oper:session-redundancy-manager"
    sessionRedundancyManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionRedundancyManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionRedundancyManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionRedundancyManager.EntityData.Children = types.NewOrderedMap()
    sessionRedundancyManager.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &sessionRedundancyManager.Interfaces})
    sessionRedundancyManager.EntityData.Children.Append("groups", types.YChild{"Groups", &sessionRedundancyManager.Groups})
    sessionRedundancyManager.EntityData.Children.Append("summary", types.YChild{"Summary", &sessionRedundancyManager.Summary})
    sessionRedundancyManager.EntityData.Leafs = types.NewOrderedMap()

    sessionRedundancyManager.EntityData.YListKeys = []string {}

    return &(sessionRedundancyManager.EntityData)
}

// SessionRedundancyManager_Interfaces
// Subscriber Redundancy Manager interface table
type SessionRedundancyManager_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface redundancy manager interface. The type is slice of
    // SessionRedundancyManager_Interfaces_Interface.
    Interface []*SessionRedundancyManager_Interfaces_Interface
}

func (interfaces *SessionRedundancyManager_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "session-redundancy-manager"
    interfaces.EntityData.SegmentPath = "interfaces"
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

// SessionRedundancyManager_Interfaces_Interface
// interface redundancy manager interface
type SessionRedundancyManager_Interfaces_Interface struct {
    EntityData types.CommonEntityData
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

func (self *SessionRedundancyManager_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Interface, "interface")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-mapping-id", types.YLeaf{"InterfaceMappingId", self.InterfaceMappingId})
    self.EntityData.Leafs.Append("forward-referenced", types.YLeaf{"ForwardReferenced", self.ForwardReferenced})
    self.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", self.GroupId})
    self.EntityData.Leafs.Append("role", types.YLeaf{"Role", self.Role})

    self.EntityData.YListKeys = []string {"Interface"}

    return &(self.EntityData)
}

// SessionRedundancyManager_Groups
// Session Redundancy Manager group table
type SessionRedundancyManager_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session redundancy manager group. The type is slice of
    // SessionRedundancyManager_Groups_Group.
    Group []*SessionRedundancyManager_Groups_Group
}

func (groups *SessionRedundancyManager_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "session-redundancy-manager"
    groups.EntityData.SegmentPath = "groups"
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

// SessionRedundancyManager_Groups_Group
// Session redundancy manager group
type SessionRedundancyManager_Groups_Group struct {
    EntityData types.CommonEntityData
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

func (group *SessionRedundancyManager_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.Group, "group")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group", types.YLeaf{"Group", group.Group})
    group.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", group.GroupId})
    group.EntityData.Leafs.Append("description", types.YLeaf{"Description", group.Description})
    group.EntityData.Leafs.Append("disabled", types.YLeaf{"Disabled", group.Disabled})
    group.EntityData.Leafs.Append("role", types.YLeaf{"Role", group.Role})
    group.EntityData.Leafs.Append("peer-ipv4-address", types.YLeaf{"PeerIpv4Address", group.PeerIpv4Address})
    group.EntityData.Leafs.Append("peer-ipv6-address", types.YLeaf{"PeerIpv6Address", group.PeerIpv6Address})
    group.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", group.InterfaceCount})
    group.EntityData.Leafs.Append("preferred-role", types.YLeaf{"PreferredRole", group.PreferredRole})
    group.EntityData.Leafs.Append("slave-mode", types.YLeaf{"SlaveMode", group.SlaveMode})
    group.EntityData.Leafs.Append("object-tracking-status", types.YLeaf{"ObjectTrackingStatus", group.ObjectTrackingStatus})
    group.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", group.NodeName})

    group.EntityData.YListKeys = []string {"Group"}

    return &(group.EntityData)
}

// SessionRedundancyManager_Summary
// Session redundancy manager summary
type SessionRedundancyManager_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *SessionRedundancyManager_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "session-redundancy-manager"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("disabled", types.YLeaf{"Disabled", summary.Disabled})
    summary.EntityData.Leafs.Append("active-state", types.YLeaf{"ActiveState", summary.ActiveState})
    summary.EntityData.Leafs.Append("preferred-role", types.YLeaf{"PreferredRole", summary.PreferredRole})
    summary.EntityData.Leafs.Append("slave-mode", types.YLeaf{"SlaveMode", summary.SlaveMode})
    summary.EntityData.Leafs.Append("hold-timer", types.YLeaf{"HoldTimer", summary.HoldTimer})
    summary.EntityData.Leafs.Append("source-interface-name", types.YLeaf{"SourceInterfaceName", summary.SourceInterfaceName})
    summary.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", summary.VrfName})
    summary.EntityData.Leafs.Append("source-interface-ipv4-address", types.YLeaf{"SourceInterfaceIpv4Address", summary.SourceInterfaceIpv4Address})
    summary.EntityData.Leafs.Append("source-interface-ipv6-address", types.YLeaf{"SourceInterfaceIpv6Address", summary.SourceInterfaceIpv6Address})
    summary.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", summary.GroupCount})
    summary.EntityData.Leafs.Append("disabled-group-count", types.YLeaf{"DisabledGroupCount", summary.DisabledGroupCount})
    summary.EntityData.Leafs.Append("master-group-count", types.YLeaf{"MasterGroupCount", summary.MasterGroupCount})
    summary.EntityData.Leafs.Append("slave-group-count", types.YLeaf{"SlaveGroupCount", summary.SlaveGroupCount})
    summary.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", summary.InterfaceCount})
    summary.EntityData.Leafs.Append("master-interface-count", types.YLeaf{"MasterInterfaceCount", summary.MasterInterfaceCount})
    summary.EntityData.Leafs.Append("slave-interface-count", types.YLeaf{"SlaveInterfaceCount", summary.SlaveInterfaceCount})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// SessionRedundancyAgent
// session redundancy agent
type SessionRedundancyAgent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes for which session data is collected.
    Nodes SessionRedundancyAgent_Nodes
}

func (sessionRedundancyAgent *SessionRedundancyAgent) GetEntityData() *types.CommonEntityData {
    sessionRedundancyAgent.EntityData.YFilter = sessionRedundancyAgent.YFilter
    sessionRedundancyAgent.EntityData.YangName = "session-redundancy-agent"
    sessionRedundancyAgent.EntityData.BundleName = "cisco_ios_xr"
    sessionRedundancyAgent.EntityData.ParentYangName = "Cisco-IOS-XR-infra-serg-oper"
    sessionRedundancyAgent.EntityData.SegmentPath = "Cisco-IOS-XR-infra-serg-oper:session-redundancy-agent"
    sessionRedundancyAgent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionRedundancyAgent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionRedundancyAgent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionRedundancyAgent.EntityData.Children = types.NewOrderedMap()
    sessionRedundancyAgent.EntityData.Children.Append("nodes", types.YChild{"Nodes", &sessionRedundancyAgent.Nodes})
    sessionRedundancyAgent.EntityData.Leafs = types.NewOrderedMap()

    sessionRedundancyAgent.EntityData.YListKeys = []string {}

    return &(sessionRedundancyAgent.EntityData)
}

// SessionRedundancyAgent_Nodes
// List of nodes for which session data is
// collected
type SessionRedundancyAgent_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session data for a particular node. The type is slice of
    // SessionRedundancyAgent_Nodes_Node.
    Node []*SessionRedundancyAgent_Nodes_Node
}

func (nodes *SessionRedundancyAgent_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "session-redundancy-agent"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// SessionRedundancyAgent_Nodes_Node
// Session data for a particular node
type SessionRedundancyAgent_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *SessionRedundancyAgent_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("group-id-xr", types.YChild{"GroupIdXr", &node.GroupIdXr})
    node.EntityData.Children.Append("client-ids", types.YChild{"ClientIds", &node.ClientIds})
    node.EntityData.Children.Append("memory", types.YChild{"Memory", &node.Memory})
    node.EntityData.Children.Append("group-ids", types.YChild{"GroupIds", &node.GroupIds})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Children.Append("stats-global", types.YChild{"StatsGlobal", &node.StatsGlobal})
    node.EntityData.Children.Append("group-summaries", types.YChild{"GroupSummaries", &node.GroupSummaries})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupIdXr
// Data for particular subscriber group session
type SessionRedundancyAgent_Nodes_Node_GroupIdXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group id for subscriber group session. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId.
    GroupId []*SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId
}

func (groupIdXr *SessionRedundancyAgent_Nodes_Node_GroupIdXr) GetEntityData() *types.CommonEntityData {
    groupIdXr.EntityData.YFilter = groupIdXr.YFilter
    groupIdXr.EntityData.YangName = "group-id-xr"
    groupIdXr.EntityData.BundleName = "cisco_ios_xr"
    groupIdXr.EntityData.ParentYangName = "node"
    groupIdXr.EntityData.SegmentPath = "group-id-xr"
    groupIdXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupIdXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupIdXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupIdXr.EntityData.Children = types.NewOrderedMap()
    groupIdXr.EntityData.Children.Append("group-id", types.YChild{"GroupId", nil})
    for i := range groupIdXr.GroupId {
        groupIdXr.EntityData.Children.Append(types.GetSegmentPath(groupIdXr.GroupId[i]), types.YChild{"GroupId", groupIdXr.GroupId[i]})
    }
    groupIdXr.EntityData.Leafs = types.NewOrderedMap()

    groupIdXr.EntityData.YListKeys = []string {}

    return &(groupIdXr.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId
// Group id for subscriber group session
type SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId struct {
    EntityData types.CommonEntityData
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
    SessionDetailedInformation []*SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation

    // Session Synchroniation Error Information. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation.
    SessionSyncErrorInformation []*SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId) GetEntityData() *types.CommonEntityData {
    groupId.EntityData.YFilter = groupId.YFilter
    groupId.EntityData.YangName = "group-id"
    groupId.EntityData.BundleName = "cisco_ios_xr"
    groupId.EntityData.ParentYangName = "group-id-xr"
    groupId.EntityData.SegmentPath = "group-id" + types.AddKeyToken(groupId.GroupId, "group-id")
    groupId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupId.EntityData.Children = types.NewOrderedMap()
    groupId.EntityData.Children.Append("session-detailed-information", types.YChild{"SessionDetailedInformation", nil})
    for i := range groupId.SessionDetailedInformation {
        groupId.EntityData.Children.Append(types.GetSegmentPath(groupId.SessionDetailedInformation[i]), types.YChild{"SessionDetailedInformation", groupId.SessionDetailedInformation[i]})
    }
    groupId.EntityData.Children.Append("session-sync-error-information", types.YChild{"SessionSyncErrorInformation", nil})
    for i := range groupId.SessionSyncErrorInformation {
        groupId.EntityData.Children.Append(types.GetSegmentPath(groupId.SessionSyncErrorInformation[i]), types.YChild{"SessionSyncErrorInformation", groupId.SessionSyncErrorInformation[i]})
    }
    groupId.EntityData.Leafs = types.NewOrderedMap()
    groupId.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", groupId.GroupId})
    groupId.EntityData.Leafs.Append("group-id-xr", types.YLeaf{"GroupIdXr", groupId.GroupIdXr})
    groupId.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", groupId.InterfaceName})
    groupId.EntityData.Leafs.Append("key-index", types.YLeaf{"KeyIndex", groupId.KeyIndex})
    groupId.EntityData.Leafs.Append("role-master", types.YLeaf{"RoleMaster", groupId.RoleMaster})
    groupId.EntityData.Leafs.Append("negative-acknowledgement-update-all", types.YLeaf{"NegativeAcknowledgementUpdateAll", groupId.NegativeAcknowledgementUpdateAll})

    groupId.EntityData.YListKeys = []string {"GroupId"}

    return &(groupId.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation
// More Session Information
type SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation struct {
    EntityData types.CommonEntityData
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

func (sessionDetailedInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionDetailedInformation) GetEntityData() *types.CommonEntityData {
    sessionDetailedInformation.EntityData.YFilter = sessionDetailedInformation.YFilter
    sessionDetailedInformation.EntityData.YangName = "session-detailed-information"
    sessionDetailedInformation.EntityData.BundleName = "cisco_ios_xr"
    sessionDetailedInformation.EntityData.ParentYangName = "group-id"
    sessionDetailedInformation.EntityData.SegmentPath = "session-detailed-information"
    sessionDetailedInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDetailedInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDetailedInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDetailedInformation.EntityData.Children = types.NewOrderedMap()
    sessionDetailedInformation.EntityData.Leafs = types.NewOrderedMap()
    sessionDetailedInformation.EntityData.Leafs.Append("component", types.YLeaf{"Component", sessionDetailedInformation.Component})
    sessionDetailedInformation.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", sessionDetailedInformation.Operation})
    sessionDetailedInformation.EntityData.Leafs.Append("tx-list-queue-fail", types.YLeaf{"TxListQueueFail", sessionDetailedInformation.TxListQueueFail})
    sessionDetailedInformation.EntityData.Leafs.Append("marked-for-sweeping", types.YLeaf{"MarkedForSweeping", sessionDetailedInformation.MarkedForSweeping})
    sessionDetailedInformation.EntityData.Leafs.Append("marked-for-cleanup", types.YLeaf{"MarkedForCleanup", sessionDetailedInformation.MarkedForCleanup})

    sessionDetailedInformation.EntityData.YListKeys = []string {}

    return &(sessionDetailedInformation.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation
// Session Synchroniation Error Information
type SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // No. of Errors occured during Synchronization. The type is interface{} with
    // range: 0..65535.
    SyncErrorCount interface{}

    // Last Error Code. The type is interface{} with range: 0..4294967295.
    LastErrorCode interface{}

    // Last Error Type. The type is SergShowSessionError.
    LastErrorType interface{}
}

func (sessionSyncErrorInformation *SessionRedundancyAgent_Nodes_Node_GroupIdXr_GroupId_SessionSyncErrorInformation) GetEntityData() *types.CommonEntityData {
    sessionSyncErrorInformation.EntityData.YFilter = sessionSyncErrorInformation.YFilter
    sessionSyncErrorInformation.EntityData.YangName = "session-sync-error-information"
    sessionSyncErrorInformation.EntityData.BundleName = "cisco_ios_xr"
    sessionSyncErrorInformation.EntityData.ParentYangName = "group-id"
    sessionSyncErrorInformation.EntityData.SegmentPath = "session-sync-error-information"
    sessionSyncErrorInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionSyncErrorInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionSyncErrorInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionSyncErrorInformation.EntityData.Children = types.NewOrderedMap()
    sessionSyncErrorInformation.EntityData.Leafs = types.NewOrderedMap()
    sessionSyncErrorInformation.EntityData.Leafs.Append("sync-error-count", types.YLeaf{"SyncErrorCount", sessionSyncErrorInformation.SyncErrorCount})
    sessionSyncErrorInformation.EntityData.Leafs.Append("last-error-code", types.YLeaf{"LastErrorCode", sessionSyncErrorInformation.LastErrorCode})
    sessionSyncErrorInformation.EntityData.Leafs.Append("last-error-type", types.YLeaf{"LastErrorType", sessionSyncErrorInformation.LastErrorType})

    sessionSyncErrorInformation.EntityData.YListKeys = []string {}

    return &(sessionSyncErrorInformation.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_ClientIds
// Stats Client
type SessionRedundancyAgent_Nodes_Node_ClientIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify stats client. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId.
    ClientId []*SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId
}

func (clientIds *SessionRedundancyAgent_Nodes_Node_ClientIds) GetEntityData() *types.CommonEntityData {
    clientIds.EntityData.YFilter = clientIds.YFilter
    clientIds.EntityData.YangName = "client-ids"
    clientIds.EntityData.BundleName = "cisco_ios_xr"
    clientIds.EntityData.ParentYangName = "node"
    clientIds.EntityData.SegmentPath = "client-ids"
    clientIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientIds.EntityData.Children = types.NewOrderedMap()
    clientIds.EntityData.Children.Append("client-id", types.YChild{"ClientId", nil})
    for i := range clientIds.ClientId {
        clientIds.EntityData.Children.Append(types.GetSegmentPath(clientIds.ClientId[i]), types.YChild{"ClientId", clientIds.ClientId[i]})
    }
    clientIds.EntityData.Leafs = types.NewOrderedMap()

    clientIds.EntityData.YListKeys = []string {}

    return &(clientIds.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId
// Specify stats client
type SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client Id. The type is interface{} with range:
    // 0..4294967295.
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

func (clientId *SessionRedundancyAgent_Nodes_Node_ClientIds_ClientId) GetEntityData() *types.CommonEntityData {
    clientId.EntityData.YFilter = clientId.YFilter
    clientId.EntityData.YangName = "client-id"
    clientId.EntityData.BundleName = "cisco_ios_xr"
    clientId.EntityData.ParentYangName = "client-ids"
    clientId.EntityData.SegmentPath = "client-id" + types.AddKeyToken(clientId.StatsClientId, "stats-client-id")
    clientId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientId.EntityData.Children = types.NewOrderedMap()
    clientId.EntityData.Leafs = types.NewOrderedMap()
    clientId.EntityData.Leafs.Append("stats-client-id", types.YLeaf{"StatsClientId", clientId.StatsClientId})
    clientId.EntityData.Leafs.Append("tx-list-start-of-download-add-ok", types.YLeaf{"TxListStartOfDownloadAddOk", clientId.TxListStartOfDownloadAddOk})
    clientId.EntityData.Leafs.Append("tx-list-start-of-download-add-not-ok", types.YLeaf{"TxListStartOfDownloadAddNotOk", clientId.TxListStartOfDownloadAddNotOk})
    clientId.EntityData.Leafs.Append("tx-list-end-of-download-add-ok", types.YLeaf{"TxListEndOfDownloadAddOk", clientId.TxListEndOfDownloadAddOk})
    clientId.EntityData.Leafs.Append("tx-list-end-of-download-add-not-ok", types.YLeaf{"TxListEndOfDownloadAddNotOk", clientId.TxListEndOfDownloadAddNotOk})
    clientId.EntityData.Leafs.Append("tx-list-end-of-message-add-ok", types.YLeaf{"TxListEndOfMessageAddOk", clientId.TxListEndOfMessageAddOk})
    clientId.EntityData.Leafs.Append("tx-list-end-of-message-add-not-ok", types.YLeaf{"TxListEndOfMessageAddNotOk", clientId.TxListEndOfMessageAddNotOk})
    clientId.EntityData.Leafs.Append("tx-list-clear-all-add-ok", types.YLeaf{"TxListClearAllAddOk", clientId.TxListClearAllAddOk})
    clientId.EntityData.Leafs.Append("tx-list-clear-all-add-not-ok", types.YLeaf{"TxListClearAllAddNotOk", clientId.TxListClearAllAddNotOk})
    clientId.EntityData.Leafs.Append("tx-list-clear-selected-add-ok", types.YLeaf{"TxListClearSelectedAddOk", clientId.TxListClearSelectedAddOk})
    clientId.EntityData.Leafs.Append("tx-list-clear-selected-add-not-ok", types.YLeaf{"TxListClearSelectedAddNotOk", clientId.TxListClearSelectedAddNotOk})
    clientId.EntityData.Leafs.Append("tx-list-replay-all-add-ok", types.YLeaf{"TxListReplayAllAddOk", clientId.TxListReplayAllAddOk})
    clientId.EntityData.Leafs.Append("tx-list-replay-all-add-not-ok", types.YLeaf{"TxListReplayAllAddNotOk", clientId.TxListReplayAllAddNotOk})
    clientId.EntityData.Leafs.Append("tx-list-replay-selected-add-ok", types.YLeaf{"TxListReplaySelectedAddOk", clientId.TxListReplaySelectedAddOk})
    clientId.EntityData.Leafs.Append("tx-list-replay-selected-add-not-ok", types.YLeaf{"TxListReplaySelectedAddNotOk", clientId.TxListReplaySelectedAddNotOk})
    clientId.EntityData.Leafs.Append("tx-list-session-session-add-ok", types.YLeaf{"TxListSessionSessionAddOk", clientId.TxListSessionSessionAddOk})
    clientId.EntityData.Leafs.Append("tx-list-session-session-add-not-ok", types.YLeaf{"TxListSessionSessionAddNotOk", clientId.TxListSessionSessionAddNotOk})
    clientId.EntityData.Leafs.Append("tx-list-session-session-update-ok", types.YLeaf{"TxListSessionSessionUpdateOk", clientId.TxListSessionSessionUpdateOk})
    clientId.EntityData.Leafs.Append("tx-list-session-session-update-not-ok", types.YLeaf{"TxListSessionSessionUpdateNotOk", clientId.TxListSessionSessionUpdateNotOk})
    clientId.EntityData.Leafs.Append("tx-list-session-session-delete", types.YLeaf{"TxListSessionSessionDelete", clientId.TxListSessionSessionDelete})
    clientId.EntityData.Leafs.Append("clean-call-back", types.YLeaf{"CleanCallBack", clientId.CleanCallBack})
    clientId.EntityData.Leafs.Append("tx-list-encode-session-session-ok", types.YLeaf{"TxListEncodeSessionSessionOk", clientId.TxListEncodeSessionSessionOk})
    clientId.EntityData.Leafs.Append("tx-list-encode-session-session-partial-write", types.YLeaf{"TxListEncodeSessionSessionPartialWrite", clientId.TxListEncodeSessionSessionPartialWrite})
    clientId.EntityData.Leafs.Append("last-replay-all-count", types.YLeaf{"LastReplayAllCount", clientId.LastReplayAllCount})
    clientId.EntityData.Leafs.Append("tx-list-receive-command-replay-all", types.YLeaf{"TxListReceiveCommandReplayAll", clientId.TxListReceiveCommandReplayAll})
    clientId.EntityData.Leafs.Append("tx-list-receive-command-replay-selected", types.YLeaf{"TxListReceiveCommandReplaySelected", clientId.TxListReceiveCommandReplaySelected})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-delete-valid", types.YLeaf{"TxListReceiveSessionSessionDeleteValid", clientId.TxListReceiveSessionSessionDeleteValid})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-delete-invalid", types.YLeaf{"TxListReceiveSessionSessionDeleteInvalid", clientId.TxListReceiveSessionSessionDeleteInvalid})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-update-valid", types.YLeaf{"TxListReceiveSessionSessionUpdateValid", clientId.TxListReceiveSessionSessionUpdateValid})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-update-invalid", types.YLeaf{"TxListReceiveSessionSessionUpdateInvalid", clientId.TxListReceiveSessionSessionUpdateInvalid})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-sod-all", types.YLeaf{"TxListReceiveSessionSessionSodAll", clientId.TxListReceiveSessionSessionSodAll})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-sod-selected", types.YLeaf{"TxListReceiveSessionSessionSodSelected", clientId.TxListReceiveSessionSessionSodSelected})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-eod-all", types.YLeaf{"TxListReceiveSessionSessionEodAll", clientId.TxListReceiveSessionSessionEodAll})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-eod-selected", types.YLeaf{"TxListReceiveSessionSessionEodSelected", clientId.TxListReceiveSessionSessionEodSelected})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-eoms", types.YLeaf{"TxListReceiveSessionSessionEoms", clientId.TxListReceiveSessionSessionEoms})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-clear-all", types.YLeaf{"TxListReceiveSessionSessionClearAll", clientId.TxListReceiveSessionSessionClearAll})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-clear-selected", types.YLeaf{"TxListReceiveSessionSessionClearSelected", clientId.TxListReceiveSessionSessionClearSelected})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-neg-ack", types.YLeaf{"TxListReceiveSessionSessionNegAck", clientId.TxListReceiveSessionSessionNegAck})
    clientId.EntityData.Leafs.Append("tx-list-receive-session-session-neg-ack-not-ok", types.YLeaf{"TxListReceiveSessionSessionNegAckNotOk", clientId.TxListReceiveSessionSessionNegAckNotOk})
    clientId.EntityData.Leafs.Append("tx-list-client-registration-ok", types.YLeaf{"TxListClientRegistrationOk", clientId.TxListClientRegistrationOk})
    clientId.EntityData.Leafs.Append("tx-list-client-registration-not-ok", types.YLeaf{"TxListClientRegistrationNotOk", clientId.TxListClientRegistrationNotOk})
    clientId.EntityData.Leafs.Append("tx-list-client-de-registration", types.YLeaf{"TxListClientDeRegistration", clientId.TxListClientDeRegistration})
    clientId.EntityData.Leafs.Append("tx-list-client-connection-down", types.YLeaf{"TxListClientConnectionDown", clientId.TxListClientConnectionDown})
    clientId.EntityData.Leafs.Append("tx-list-client-cleanup", types.YLeaf{"TxListClientCleanup", clientId.TxListClientCleanup})
    clientId.EntityData.Leafs.Append("tx-list-active-ok", types.YLeaf{"TxListActiveOk", clientId.TxListActiveOk})
    clientId.EntityData.Leafs.Append("tx-list-active-not-ok", types.YLeaf{"TxListActiveNotOk", clientId.TxListActiveNotOk})
    clientId.EntityData.Leafs.Append("tx-list-de-active-ok", types.YLeaf{"TxListDeActiveOk", clientId.TxListDeActiveOk})
    clientId.EntityData.Leafs.Append("tx-list-de-active-not-ok", types.YLeaf{"TxListDeActiveNotOk", clientId.TxListDeActiveNotOk})

    clientId.EntityData.YListKeys = []string {"StatsClientId"}

    return &(clientId.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_Memory
// Show Memory
type SessionRedundancyAgent_Nodes_Node_Memory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Memory Info. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo.
    MemoryInfo []*SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo

    // EDM Memory Info. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo.
    EdmMemoryInfo []*SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo

    // String Memory Info. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo.
    StringMemoryInfo []*SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo
}

func (memory *SessionRedundancyAgent_Nodes_Node_Memory) GetEntityData() *types.CommonEntityData {
    memory.EntityData.YFilter = memory.YFilter
    memory.EntityData.YangName = "memory"
    memory.EntityData.BundleName = "cisco_ios_xr"
    memory.EntityData.ParentYangName = "node"
    memory.EntityData.SegmentPath = "memory"
    memory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memory.EntityData.Children = types.NewOrderedMap()
    memory.EntityData.Children.Append("memory-info", types.YChild{"MemoryInfo", nil})
    for i := range memory.MemoryInfo {
        memory.EntityData.Children.Append(types.GetSegmentPath(memory.MemoryInfo[i]), types.YChild{"MemoryInfo", memory.MemoryInfo[i]})
    }
    memory.EntityData.Children.Append("edm-memory-info", types.YChild{"EdmMemoryInfo", nil})
    for i := range memory.EdmMemoryInfo {
        memory.EntityData.Children.Append(types.GetSegmentPath(memory.EdmMemoryInfo[i]), types.YChild{"EdmMemoryInfo", memory.EdmMemoryInfo[i]})
    }
    memory.EntityData.Children.Append("string-memory-info", types.YChild{"StringMemoryInfo", nil})
    for i := range memory.StringMemoryInfo {
        memory.EntityData.Children.Append(types.GetSegmentPath(memory.StringMemoryInfo[i]), types.YChild{"StringMemoryInfo", memory.StringMemoryInfo[i]})
    }
    memory.EntityData.Leafs = types.NewOrderedMap()

    memory.EntityData.YListKeys = []string {}

    return &(memory.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo
// Memory Info
type SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo struct {
    EntityData types.CommonEntityData
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

func (memoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_MemoryInfo) GetEntityData() *types.CommonEntityData {
    memoryInfo.EntityData.YFilter = memoryInfo.YFilter
    memoryInfo.EntityData.YangName = "memory-info"
    memoryInfo.EntityData.BundleName = "cisco_ios_xr"
    memoryInfo.EntityData.ParentYangName = "memory"
    memoryInfo.EntityData.SegmentPath = "memory-info"
    memoryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryInfo.EntityData.Children = types.NewOrderedMap()
    memoryInfo.EntityData.Leafs = types.NewOrderedMap()
    memoryInfo.EntityData.Leafs.Append("structure-name", types.YLeaf{"StructureName", memoryInfo.StructureName})
    memoryInfo.EntityData.Leafs.Append("size", types.YLeaf{"Size", memoryInfo.Size})
    memoryInfo.EntityData.Leafs.Append("current-count", types.YLeaf{"CurrentCount", memoryInfo.CurrentCount})
    memoryInfo.EntityData.Leafs.Append("alloc-fails", types.YLeaf{"AllocFails", memoryInfo.AllocFails})
    memoryInfo.EntityData.Leafs.Append("alloc-count", types.YLeaf{"AllocCount", memoryInfo.AllocCount})
    memoryInfo.EntityData.Leafs.Append("freed-count", types.YLeaf{"FreedCount", memoryInfo.FreedCount})
    memoryInfo.EntityData.Leafs.Append("memory-type", types.YLeaf{"MemoryType", memoryInfo.MemoryType})

    memoryInfo.EntityData.YListKeys = []string {}

    return &(memoryInfo.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo
// EDM Memory Info
type SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo struct {
    EntityData types.CommonEntityData
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

func (edmMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_EdmMemoryInfo) GetEntityData() *types.CommonEntityData {
    edmMemoryInfo.EntityData.YFilter = edmMemoryInfo.YFilter
    edmMemoryInfo.EntityData.YangName = "edm-memory-info"
    edmMemoryInfo.EntityData.BundleName = "cisco_ios_xr"
    edmMemoryInfo.EntityData.ParentYangName = "memory"
    edmMemoryInfo.EntityData.SegmentPath = "edm-memory-info"
    edmMemoryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    edmMemoryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    edmMemoryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    edmMemoryInfo.EntityData.Children = types.NewOrderedMap()
    edmMemoryInfo.EntityData.Leafs = types.NewOrderedMap()
    edmMemoryInfo.EntityData.Leafs.Append("size", types.YLeaf{"Size", edmMemoryInfo.Size})
    edmMemoryInfo.EntityData.Leafs.Append("total", types.YLeaf{"Total", edmMemoryInfo.Total})
    edmMemoryInfo.EntityData.Leafs.Append("success", types.YLeaf{"Success", edmMemoryInfo.Success})
    edmMemoryInfo.EntityData.Leafs.Append("failure", types.YLeaf{"Failure", edmMemoryInfo.Failure})

    edmMemoryInfo.EntityData.YListKeys = []string {}

    return &(edmMemoryInfo.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo
// String Memory Info
type SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo struct {
    EntityData types.CommonEntityData
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

func (stringMemoryInfo *SessionRedundancyAgent_Nodes_Node_Memory_StringMemoryInfo) GetEntityData() *types.CommonEntityData {
    stringMemoryInfo.EntityData.YFilter = stringMemoryInfo.YFilter
    stringMemoryInfo.EntityData.YangName = "string-memory-info"
    stringMemoryInfo.EntityData.BundleName = "cisco_ios_xr"
    stringMemoryInfo.EntityData.ParentYangName = "memory"
    stringMemoryInfo.EntityData.SegmentPath = "string-memory-info"
    stringMemoryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stringMemoryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stringMemoryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stringMemoryInfo.EntityData.Children = types.NewOrderedMap()
    stringMemoryInfo.EntityData.Leafs = types.NewOrderedMap()
    stringMemoryInfo.EntityData.Leafs.Append("size", types.YLeaf{"Size", stringMemoryInfo.Size})
    stringMemoryInfo.EntityData.Leafs.Append("total", types.YLeaf{"Total", stringMemoryInfo.Total})
    stringMemoryInfo.EntityData.Leafs.Append("success", types.YLeaf{"Success", stringMemoryInfo.Success})
    stringMemoryInfo.EntityData.Leafs.Append("failure", types.YLeaf{"Failure", stringMemoryInfo.Failure})

    stringMemoryInfo.EntityData.YListKeys = []string {}

    return &(stringMemoryInfo.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupIds
// Data for particular subscriber group 
type SessionRedundancyAgent_Nodes_Node_GroupIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group id for subscriber group. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId.
    GroupId []*SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId
}

func (groupIds *SessionRedundancyAgent_Nodes_Node_GroupIds) GetEntityData() *types.CommonEntityData {
    groupIds.EntityData.YFilter = groupIds.YFilter
    groupIds.EntityData.YangName = "group-ids"
    groupIds.EntityData.BundleName = "cisco_ios_xr"
    groupIds.EntityData.ParentYangName = "node"
    groupIds.EntityData.SegmentPath = "group-ids"
    groupIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupIds.EntityData.Children = types.NewOrderedMap()
    groupIds.EntityData.Children.Append("group-id", types.YChild{"GroupId", nil})
    for i := range groupIds.GroupId {
        groupIds.EntityData.Children.Append(types.GetSegmentPath(groupIds.GroupId[i]), types.YChild{"GroupId", groupIds.GroupId[i]})
    }
    groupIds.EntityData.Leafs = types.NewOrderedMap()

    groupIds.EntityData.YListKeys = []string {}

    return &(groupIds.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId
// Group id for subscriber group
type SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId struct {
    EntityData types.CommonEntityData
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
    ClientSessionCount []*SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount

    // Interface List. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface.
    Interface []*SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface
}

func (groupId *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId) GetEntityData() *types.CommonEntityData {
    groupId.EntityData.YFilter = groupId.YFilter
    groupId.EntityData.YangName = "group-id"
    groupId.EntityData.BundleName = "cisco_ios_xr"
    groupId.EntityData.ParentYangName = "group-ids"
    groupId.EntityData.SegmentPath = "group-id" + types.AddKeyToken(groupId.GroupId, "group-id")
    groupId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupId.EntityData.Children = types.NewOrderedMap()
    groupId.EntityData.Children.Append("client-session-count", types.YChild{"ClientSessionCount", nil})
    for i := range groupId.ClientSessionCount {
        groupId.EntityData.Children.Append(types.GetSegmentPath(groupId.ClientSessionCount[i]), types.YChild{"ClientSessionCount", groupId.ClientSessionCount[i]})
    }
    groupId.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range groupId.Interface {
        groupId.EntityData.Children.Append(types.GetSegmentPath(groupId.Interface[i]), types.YChild{"Interface", groupId.Interface[i]})
    }
    groupId.EntityData.Leafs = types.NewOrderedMap()
    groupId.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", groupId.GroupId})
    groupId.EntityData.Leafs.Append("group-id-xr", types.YLeaf{"GroupIdXr", groupId.GroupIdXr})
    groupId.EntityData.Leafs.Append("description", types.YLeaf{"Description", groupId.Description})
    groupId.EntityData.Leafs.Append("disabled", types.YLeaf{"Disabled", groupId.Disabled})
    groupId.EntityData.Leafs.Append("init-role", types.YLeaf{"InitRole", groupId.InitRole})
    groupId.EntityData.Leafs.Append("negotiating-role", types.YLeaf{"NegotiatingRole", groupId.NegotiatingRole})
    groupId.EntityData.Leafs.Append("current-role", types.YLeaf{"CurrentRole", groupId.CurrentRole})
    groupId.EntityData.Leafs.Append("slave-mode", types.YLeaf{"SlaveMode", groupId.SlaveMode})
    groupId.EntityData.Leafs.Append("hold-timer", types.YLeaf{"HoldTimer", groupId.HoldTimer})
    groupId.EntityData.Leafs.Append("core-tracking-object-name", types.YLeaf{"CoreTrackingObjectName", groupId.CoreTrackingObjectName})
    groupId.EntityData.Leafs.Append("core-tracking-object-status", types.YLeaf{"CoreTrackingObjectStatus", groupId.CoreTrackingObjectStatus})
    groupId.EntityData.Leafs.Append("access-tracking-object-name", types.YLeaf{"AccessTrackingObjectName", groupId.AccessTrackingObjectName})
    groupId.EntityData.Leafs.Append("access-tracking-object-status", types.YLeaf{"AccessTrackingObjectStatus", groupId.AccessTrackingObjectStatus})
    groupId.EntityData.Leafs.Append("object-tracking-status", types.YLeaf{"ObjectTrackingStatus", groupId.ObjectTrackingStatus})
    groupId.EntityData.Leafs.Append("peer-ipv4-address", types.YLeaf{"PeerIpv4Address", groupId.PeerIpv4Address})
    groupId.EntityData.Leafs.Append("peer-ipv6-address", types.YLeaf{"PeerIpv6Address", groupId.PeerIpv6Address})
    groupId.EntityData.Leafs.Append("peer-status", types.YLeaf{"PeerStatus", groupId.PeerStatus})
    groupId.EntityData.Leafs.Append("peer-last-negotiation-time", types.YLeaf{"PeerLastNegotiationTime", groupId.PeerLastNegotiationTime})
    groupId.EntityData.Leafs.Append("peer-last-up-time", types.YLeaf{"PeerLastUpTime", groupId.PeerLastUpTime})
    groupId.EntityData.Leafs.Append("peer-last-down-time", types.YLeaf{"PeerLastDownTime", groupId.PeerLastDownTime})
    groupId.EntityData.Leafs.Append("peer-init-role", types.YLeaf{"PeerInitRole", groupId.PeerInitRole})
    groupId.EntityData.Leafs.Append("peer-negotiating-role", types.YLeaf{"PeerNegotiatingRole", groupId.PeerNegotiatingRole})
    groupId.EntityData.Leafs.Append("peer-current-role", types.YLeaf{"PeerCurrentRole", groupId.PeerCurrentRole})
    groupId.EntityData.Leafs.Append("peer-object-tracking-status", types.YLeaf{"PeerObjectTrackingStatus", groupId.PeerObjectTrackingStatus})
    groupId.EntityData.Leafs.Append("last-switchover-time", types.YLeaf{"LastSwitchoverTime", groupId.LastSwitchoverTime})
    groupId.EntityData.Leafs.Append("switchover-count", types.YLeaf{"SwitchoverCount", groupId.SwitchoverCount})
    groupId.EntityData.Leafs.Append("last-switchover-reason", types.YLeaf{"LastSwitchoverReason", groupId.LastSwitchoverReason})
    groupId.EntityData.Leafs.Append("switchover-hold-time", types.YLeaf{"SwitchoverHoldTime", groupId.SwitchoverHoldTime})
    groupId.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", groupId.SessionCount})
    groupId.EntityData.Leafs.Append("slave-update-failure-count", types.YLeaf{"SlaveUpdateFailureCount", groupId.SlaveUpdateFailureCount})
    groupId.EntityData.Leafs.Append("pending-session-update-count", types.YLeaf{"PendingSessionUpdateCount", groupId.PendingSessionUpdateCount})
    groupId.EntityData.Leafs.Append("pending-session-delete-count", types.YLeaf{"PendingSessionDeleteCount", groupId.PendingSessionDeleteCount})
    groupId.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", groupId.InterfaceCount})
    groupId.EntityData.Leafs.Append("revertive-timer", types.YLeaf{"RevertiveTimer", groupId.RevertiveTimer})
    groupId.EntityData.Leafs.Append("switchover-revert-time", types.YLeaf{"SwitchoverRevertTime", groupId.SwitchoverRevertTime})

    groupId.EntityData.YListKeys = []string {"GroupId"}

    return &(groupId.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount
// Client Session Count
type SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Component. The type is SergShowComp.
    Component interface{}

    // Session count. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}
}

func (clientSessionCount *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_ClientSessionCount) GetEntityData() *types.CommonEntityData {
    clientSessionCount.EntityData.YFilter = clientSessionCount.YFilter
    clientSessionCount.EntityData.YangName = "client-session-count"
    clientSessionCount.EntityData.BundleName = "cisco_ios_xr"
    clientSessionCount.EntityData.ParentYangName = "group-id"
    clientSessionCount.EntityData.SegmentPath = "client-session-count"
    clientSessionCount.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientSessionCount.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientSessionCount.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientSessionCount.EntityData.Children = types.NewOrderedMap()
    clientSessionCount.EntityData.Leafs = types.NewOrderedMap()
    clientSessionCount.EntityData.Leafs.Append("component", types.YLeaf{"Component", clientSessionCount.Component})
    clientSessionCount.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", clientSessionCount.SessionCount})

    clientSessionCount.EntityData.YListKeys = []string {}

    return &(clientSessionCount.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface
// Interface List
type SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface struct {
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

func (self *SessionRedundancyAgent_Nodes_Node_GroupIds_GroupId_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "group-id"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-synchronization-id", types.YLeaf{"InterfaceSynchronizationId", self.InterfaceSynchronizationId})
    self.EntityData.Leafs.Append("forward-referenced", types.YLeaf{"ForwardReferenced", self.ForwardReferenced})
    self.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", self.SessionCount})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_Interfaces
// List of interfaces
type SessionRedundancyAgent_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify interface name. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_Interfaces_Interface.
    Interface []*SessionRedundancyAgent_Nodes_Node_Interfaces_Interface
}

func (interfaces *SessionRedundancyAgent_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
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

// SessionRedundancyAgent_Nodes_Node_Interfaces_Interface
// Specify interface name
type SessionRedundancyAgent_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
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
    ClientStatus []*SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus
}

func (self *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Interface, "interface")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("interface-oper", types.YChild{"InterfaceOper", &self.InterfaceOper})
    self.EntityData.Children.Append("interface-status", types.YChild{"InterfaceStatus", &self.InterfaceStatus})
    self.EntityData.Children.Append("client-status", types.YChild{"ClientStatus", nil})
    for i := range self.ClientStatus {
        self.EntityData.Children.Append(types.GetSegmentPath(self.ClientStatus[i]), types.YChild{"ClientStatus", self.ClientStatus[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-synchronization-id", types.YLeaf{"InterfaceSynchronizationId", self.InterfaceSynchronizationId})
    self.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", self.GroupId})
    self.EntityData.Leafs.Append("role", types.YLeaf{"Role", self.Role})
    self.EntityData.Leafs.Append("forward-referenced", types.YLeaf{"ForwardReferenced", self.ForwardReferenced})
    self.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", self.SessionCount})
    self.EntityData.Leafs.Append("interface-enable-error-count", types.YLeaf{"InterfaceEnableErrorCount", self.InterfaceEnableErrorCount})
    self.EntityData.Leafs.Append("interface-disable-error-count", types.YLeaf{"InterfaceDisableErrorCount", self.InterfaceDisableErrorCount})
    self.EntityData.Leafs.Append("interface-caps-add-error-count", types.YLeaf{"InterfaceCapsAddErrorCount", self.InterfaceCapsAddErrorCount})
    self.EntityData.Leafs.Append("interface-caps-remove-error-count", types.YLeaf{"InterfaceCapsRemoveErrorCount", self.InterfaceCapsRemoveErrorCount})
    self.EntityData.Leafs.Append("interface-attribute-update-error-count", types.YLeaf{"InterfaceAttributeUpdateErrorCount", self.InterfaceAttributeUpdateErrorCount})

    self.EntityData.YListKeys = []string {"Interface"}

    return &(self.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper
// Interface Batch Operation
type SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper struct {
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

func (interfaceOper *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceOper) GetEntityData() *types.CommonEntityData {
    interfaceOper.EntityData.YFilter = interfaceOper.YFilter
    interfaceOper.EntityData.YangName = "interface-oper"
    interfaceOper.EntityData.BundleName = "cisco_ios_xr"
    interfaceOper.EntityData.ParentYangName = "interface"
    interfaceOper.EntityData.SegmentPath = "interface-oper"
    interfaceOper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceOper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceOper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceOper.EntityData.Children = types.NewOrderedMap()
    interfaceOper.EntityData.Leafs = types.NewOrderedMap()
    interfaceOper.EntityData.Leafs.Append("idb-oper-reg-enable", types.YLeaf{"IdbOperRegEnable", interfaceOper.IdbOperRegEnable})
    interfaceOper.EntityData.Leafs.Append("idb-oper-reg-disable", types.YLeaf{"IdbOperRegDisable", interfaceOper.IdbOperRegDisable})
    interfaceOper.EntityData.Leafs.Append("idb-oper-caps-add", types.YLeaf{"IdbOperCapsAdd", interfaceOper.IdbOperCapsAdd})
    interfaceOper.EntityData.Leafs.Append("idb-oper-caps-remove", types.YLeaf{"IdbOperCapsRemove", interfaceOper.IdbOperCapsRemove})
    interfaceOper.EntityData.Leafs.Append("idb-oper-attr-update", types.YLeaf{"IdbOperAttrUpdate", interfaceOper.IdbOperAttrUpdate})

    interfaceOper.EntityData.YListKeys = []string {}

    return &(interfaceOper.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus
// Interface Status
type SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus struct {
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

func (interfaceStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_InterfaceStatus) GetEntityData() *types.CommonEntityData {
    interfaceStatus.EntityData.YFilter = interfaceStatus.YFilter
    interfaceStatus.EntityData.YangName = "interface-status"
    interfaceStatus.EntityData.BundleName = "cisco_ios_xr"
    interfaceStatus.EntityData.ParentYangName = "interface"
    interfaceStatus.EntityData.SegmentPath = "interface-status"
    interfaceStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStatus.EntityData.Children = types.NewOrderedMap()
    interfaceStatus.EntityData.Leafs = types.NewOrderedMap()
    interfaceStatus.EntityData.Leafs.Append("idb-state-fwd-ref", types.YLeaf{"IdbStateFwdRef", interfaceStatus.IdbStateFwdRef})
    interfaceStatus.EntityData.Leafs.Append("idb-state-stale", types.YLeaf{"IdbStateStale", interfaceStatus.IdbStateStale})
    interfaceStatus.EntityData.Leafs.Append("idb-state-registered", types.YLeaf{"IdbStateRegistered", interfaceStatus.IdbStateRegistered})
    interfaceStatus.EntityData.Leafs.Append("idb-state-caps-added", types.YLeaf{"IdbStateCapsAdded", interfaceStatus.IdbStateCapsAdded})
    interfaceStatus.EntityData.Leafs.Append("idb-state-owned-re-source", types.YLeaf{"IdbStateOwnedReSource", interfaceStatus.IdbStateOwnedReSource})
    interfaceStatus.EntityData.Leafs.Append("idb-client-eoms-pending", types.YLeaf{"IdbClientEomsPending", interfaceStatus.IdbClientEomsPending})
    interfaceStatus.EntityData.Leafs.Append("idb-state-p-end-caps-rem", types.YLeaf{"IdbStatePEndCapsRem", interfaceStatus.IdbStatePEndCapsRem})
    interfaceStatus.EntityData.Leafs.Append("idb-state-p-end-reg-disable", types.YLeaf{"IdbStatePEndRegDisable", interfaceStatus.IdbStatePEndRegDisable})

    interfaceStatus.EntityData.YListKeys = []string {}

    return &(interfaceStatus.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus
// Interface status for each client
type SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus struct {
    EntityData types.CommonEntityData
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

func (clientStatus *SessionRedundancyAgent_Nodes_Node_Interfaces_Interface_ClientStatus) GetEntityData() *types.CommonEntityData {
    clientStatus.EntityData.YFilter = clientStatus.YFilter
    clientStatus.EntityData.YangName = "client-status"
    clientStatus.EntityData.BundleName = "cisco_ios_xr"
    clientStatus.EntityData.ParentYangName = "interface"
    clientStatus.EntityData.SegmentPath = "client-status"
    clientStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientStatus.EntityData.Children = types.NewOrderedMap()
    clientStatus.EntityData.Leafs = types.NewOrderedMap()
    clientStatus.EntityData.Leafs.Append("component", types.YLeaf{"Component", clientStatus.Component})
    clientStatus.EntityData.Leafs.Append("serg-show-idb-client-eoms-pending", types.YLeaf{"SergShowIdbClientEomsPending", clientStatus.SergShowIdbClientEomsPending})
    clientStatus.EntityData.Leafs.Append("serg-show-idb-client-sync-eod-pending", types.YLeaf{"SergShowIdbClientSyncEodPending", clientStatus.SergShowIdbClientSyncEodPending})
    clientStatus.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", clientStatus.SessionCount})

    clientStatus.EntityData.YListKeys = []string {}

    return &(clientStatus.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_StatsGlobal
// Stats Global
type SessionRedundancyAgent_Nodes_Node_StatsGlobal struct {
    EntityData types.CommonEntityData
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
    ClientStatus []*SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus

    // Opaque memory Stats. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus.
    OpaqueMemoryStatus []*SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus

    // TCP TxList Statistics. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus.
    TxListOverTcpStatus []*SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus
}

func (statsGlobal *SessionRedundancyAgent_Nodes_Node_StatsGlobal) GetEntityData() *types.CommonEntityData {
    statsGlobal.EntityData.YFilter = statsGlobal.YFilter
    statsGlobal.EntityData.YangName = "stats-global"
    statsGlobal.EntityData.BundleName = "cisco_ios_xr"
    statsGlobal.EntityData.ParentYangName = "node"
    statsGlobal.EntityData.SegmentPath = "stats-global"
    statsGlobal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsGlobal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsGlobal.EntityData.Children = types.NewOrderedMap()
    statsGlobal.EntityData.Children.Append("intf-status-statistics", types.YChild{"IntfStatusStatistics", &statsGlobal.IntfStatusStatistics})
    statsGlobal.EntityData.Children.Append("tx-list-statistics", types.YChild{"TxListStatistics", &statsGlobal.TxListStatistics})
    statsGlobal.EntityData.Children.Append("client-status", types.YChild{"ClientStatus", nil})
    for i := range statsGlobal.ClientStatus {
        statsGlobal.EntityData.Children.Append(types.GetSegmentPath(statsGlobal.ClientStatus[i]), types.YChild{"ClientStatus", statsGlobal.ClientStatus[i]})
    }
    statsGlobal.EntityData.Children.Append("opaque-memory-status", types.YChild{"OpaqueMemoryStatus", nil})
    for i := range statsGlobal.OpaqueMemoryStatus {
        statsGlobal.EntityData.Children.Append(types.GetSegmentPath(statsGlobal.OpaqueMemoryStatus[i]), types.YChild{"OpaqueMemoryStatus", statsGlobal.OpaqueMemoryStatus[i]})
    }
    statsGlobal.EntityData.Children.Append("tx-list-over-tcp-status", types.YChild{"TxListOverTcpStatus", nil})
    for i := range statsGlobal.TxListOverTcpStatus {
        statsGlobal.EntityData.Children.Append(types.GetSegmentPath(statsGlobal.TxListOverTcpStatus[i]), types.YChild{"TxListOverTcpStatus", statsGlobal.TxListOverTcpStatus[i]})
    }
    statsGlobal.EntityData.Leafs = types.NewOrderedMap()
    statsGlobal.EntityData.Leafs.Append("source-interface-name", types.YLeaf{"SourceInterfaceName", statsGlobal.SourceInterfaceName})
    statsGlobal.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", statsGlobal.VrfName})
    statsGlobal.EntityData.Leafs.Append("source-interface-ipv4-address", types.YLeaf{"SourceInterfaceIpv4Address", statsGlobal.SourceInterfaceIpv4Address})
    statsGlobal.EntityData.Leafs.Append("source-interface-ipv6-address", types.YLeaf{"SourceInterfaceIpv6Address", statsGlobal.SourceInterfaceIpv6Address})
    statsGlobal.EntityData.Leafs.Append("redundancy-role", types.YLeaf{"RedundancyRole", statsGlobal.RedundancyRole})
    statsGlobal.EntityData.Leafs.Append("restart-client-sync-in-progress", types.YLeaf{"RestartClientSyncInProgress", statsGlobal.RestartClientSyncInProgress})
    statsGlobal.EntityData.Leafs.Append("client-init-sync-time-stamp", types.YLeaf{"ClientInitSyncTimeStamp", statsGlobal.ClientInitSyncTimeStamp})
    statsGlobal.EntityData.Leafs.Append("restart-peer-sync-in-progress", types.YLeaf{"RestartPeerSyncInProgress", statsGlobal.RestartPeerSyncInProgress})
    statsGlobal.EntityData.Leafs.Append("peer-init-sync-time-stamp", types.YLeaf{"PeerInitSyncTimeStamp", statsGlobal.PeerInitSyncTimeStamp})
    statsGlobal.EntityData.Leafs.Append("sync-in-progress", types.YLeaf{"SyncInProgress", statsGlobal.SyncInProgress})
    statsGlobal.EntityData.Leafs.Append("peer-action-timer", types.YLeaf{"PeerActionTimer", statsGlobal.PeerActionTimer})
    statsGlobal.EntityData.Leafs.Append("retry-timer-remaining", types.YLeaf{"RetryTimerRemaining", statsGlobal.RetryTimerRemaining})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-registration-invalid", types.YLeaf{"TxListClientRegistrationInvalid", statsGlobal.TxListClientRegistrationInvalid})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-de-registration-invalid", types.YLeaf{"TxListClientDeRegistrationInvalid", statsGlobal.TxListClientDeRegistrationInvalid})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-connection-up", types.YLeaf{"TxListClientConnectionUp", statsGlobal.TxListClientConnectionUp})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-connection-down", types.YLeaf{"TxListClientConnectionDown", statsGlobal.TxListClientConnectionDown})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-peer-done", types.YLeaf{"TxListClientPeerDone", statsGlobal.TxListClientPeerDone})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-message-call-back", types.YLeaf{"TxListClientMessageCallBack", statsGlobal.TxListClientMessageCallBack})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-receive-valid", types.YLeaf{"TxListClientReceiveValid", statsGlobal.TxListClientReceiveValid})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-receive-invalid", types.YLeaf{"TxListClientReceiveInvalid", statsGlobal.TxListClientReceiveInvalid})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-receive-command-valid", types.YLeaf{"TxListClientReceiveCommandValid", statsGlobal.TxListClientReceiveCommandValid})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-receive-command-invalid", types.YLeaf{"TxListClientReceiveCommandInvalid", statsGlobal.TxListClientReceiveCommandInvalid})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-receive-session-sessionvalid", types.YLeaf{"TxListClientReceiveSessionSessionvalid", statsGlobal.TxListClientReceiveSessionSessionvalid})
    statsGlobal.EntityData.Leafs.Append("tx-list-client-receive-session-session-invalid", types.YLeaf{"TxListClientReceiveSessionSessionInvalid", statsGlobal.TxListClientReceiveSessionSessionInvalid})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-timer-handler", types.YLeaf{"TxListPeerTimerHandler", statsGlobal.TxListPeerTimerHandler})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-registration-invalid", types.YLeaf{"TxListPeerRegistrationInvalid", statsGlobal.TxListPeerRegistrationInvalid})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-de-registration-invalid", types.YLeaf{"TxListPeerDeRegistrationInvalid", statsGlobal.TxListPeerDeRegistrationInvalid})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-message-call-back-valid", types.YLeaf{"TxListPeerMessageCallBackValid", statsGlobal.TxListPeerMessageCallBackValid})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-message-call-back-invalid", types.YLeaf{"TxListPeerMessageCallBackInvalid", statsGlobal.TxListPeerMessageCallBackInvalid})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-done", types.YLeaf{"TxListPeerDone", statsGlobal.TxListPeerDone})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-cmd-connection-up-not-ok", types.YLeaf{"TxListPeerCmdConnectionUpNotOk", statsGlobal.TxListPeerCmdConnectionUpNotOk})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-cmd-connection-down-not-ok", types.YLeaf{"TxListPeerCmdConnectionDownNotOk", statsGlobal.TxListPeerCmdConnectionDownNotOk})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-session-connection-up-not-ok", types.YLeaf{"TxListPeerSessionConnectionUpNotOk", statsGlobal.TxListPeerSessionConnectionUpNotOk})
    statsGlobal.EntityData.Leafs.Append("tx-list-peer-session-connection-down-not-ok", types.YLeaf{"TxListPeerSessionConnectionDownNotOk", statsGlobal.TxListPeerSessionConnectionDownNotOk})

    statsGlobal.EntityData.YListKeys = []string {}

    return &(statsGlobal.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics
// intf status statistics
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics struct {
    EntityData types.CommonEntityData
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

func (intfStatusStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_IntfStatusStatistics) GetEntityData() *types.CommonEntityData {
    intfStatusStatistics.EntityData.YFilter = intfStatusStatistics.YFilter
    intfStatusStatistics.EntityData.YangName = "intf-status-statistics"
    intfStatusStatistics.EntityData.BundleName = "cisco_ios_xr"
    intfStatusStatistics.EntityData.ParentYangName = "stats-global"
    intfStatusStatistics.EntityData.SegmentPath = "intf-status-statistics"
    intfStatusStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intfStatusStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intfStatusStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intfStatusStatistics.EntityData.Children = types.NewOrderedMap()
    intfStatusStatistics.EntityData.Leafs = types.NewOrderedMap()
    intfStatusStatistics.EntityData.Leafs.Append("pend-caps-rem-cnt", types.YLeaf{"PendCapsRemCnt", intfStatusStatistics.PendCapsRemCnt})
    intfStatusStatistics.EntityData.Leafs.Append("pend-reg-disable-cnt", types.YLeaf{"PendRegDisableCnt", intfStatusStatistics.PendRegDisableCnt})
    intfStatusStatistics.EntityData.Leafs.Append("pend-other-batch-oper-cnt", types.YLeaf{"PendOtherBatchOperCnt", intfStatusStatistics.PendOtherBatchOperCnt})
    intfStatusStatistics.EntityData.Leafs.Append("non-stale-cnt", types.YLeaf{"NonStaleCnt", intfStatusStatistics.NonStaleCnt})
    intfStatusStatistics.EntityData.Leafs.Append("grp-bound-cnt", types.YLeaf{"GrpBoundCnt", intfStatusStatistics.GrpBoundCnt})

    intfStatusStatistics.EntityData.YListKeys = []string {}

    return &(intfStatusStatistics.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics
// tx list statistics
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics struct {
    EntityData types.CommonEntityData
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

func (txListStatistics *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListStatistics) GetEntityData() *types.CommonEntityData {
    txListStatistics.EntityData.YFilter = txListStatistics.YFilter
    txListStatistics.EntityData.YangName = "tx-list-statistics"
    txListStatistics.EntityData.BundleName = "cisco_ios_xr"
    txListStatistics.EntityData.ParentYangName = "stats-global"
    txListStatistics.EntityData.SegmentPath = "tx-list-statistics"
    txListStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txListStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txListStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txListStatistics.EntityData.Children = types.NewOrderedMap()
    txListStatistics.EntityData.Leafs = types.NewOrderedMap()
    txListStatistics.EntityData.Leafs.Append("tx-list-encode-marker-ok", types.YLeaf{"TxListEncodeMarkerOk", txListStatistics.TxListEncodeMarkerOk})
    txListStatistics.EntityData.Leafs.Append("tx-list-encode-marker-partial-write", types.YLeaf{"TxListEncodeMarkerPartialWrite", txListStatistics.TxListEncodeMarkerPartialWrite})
    txListStatistics.EntityData.Leafs.Append("tx-list-clean-marker", types.YLeaf{"TxListCleanMarker", txListStatistics.TxListCleanMarker})
    txListStatistics.EntityData.Leafs.Append("tx-list-encode-command-ok", types.YLeaf{"TxListEncodeCommandOk", txListStatistics.TxListEncodeCommandOk})
    txListStatistics.EntityData.Leafs.Append("tx-list-encode-command-partial-write", types.YLeaf{"TxListEncodeCommandPartialWrite", txListStatistics.TxListEncodeCommandPartialWrite})
    txListStatistics.EntityData.Leafs.Append("tx-list-clean-command", types.YLeaf{"TxListCleanCommand", txListStatistics.TxListCleanCommand})
    txListStatistics.EntityData.Leafs.Append("tx-list-encode-negotiation-ok", types.YLeaf{"TxListEncodeNegotiationOk", txListStatistics.TxListEncodeNegotiationOk})
    txListStatistics.EntityData.Leafs.Append("tx-list-encode-negotiation-partial-write", types.YLeaf{"TxListEncodeNegotiationPartialWrite", txListStatistics.TxListEncodeNegotiationPartialWrite})
    txListStatistics.EntityData.Leafs.Append("tx-list-clean-negotiation", types.YLeaf{"TxListCleanNegotiation", txListStatistics.TxListCleanNegotiation})

    txListStatistics.EntityData.YListKeys = []string {}

    return &(txListStatistics.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus
// Client Status
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus struct {
    EntityData types.CommonEntityData
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

func (clientStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_ClientStatus) GetEntityData() *types.CommonEntityData {
    clientStatus.EntityData.YFilter = clientStatus.YFilter
    clientStatus.EntityData.YangName = "client-status"
    clientStatus.EntityData.BundleName = "cisco_ios_xr"
    clientStatus.EntityData.ParentYangName = "stats-global"
    clientStatus.EntityData.SegmentPath = "client-status"
    clientStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientStatus.EntityData.Children = types.NewOrderedMap()
    clientStatus.EntityData.Leafs = types.NewOrderedMap()
    clientStatus.EntityData.Leafs.Append("component", types.YLeaf{"Component", clientStatus.Component})
    clientStatus.EntityData.Leafs.Append("client-connection-status", types.YLeaf{"ClientConnectionStatus", clientStatus.ClientConnectionStatus})
    clientStatus.EntityData.Leafs.Append("client-initialization-synchronization-pending", types.YLeaf{"ClientInitializationSynchronizationPending", clientStatus.ClientInitializationSynchronizationPending})
    clientStatus.EntityData.Leafs.Append("client-synchronization-end-of-download-pending", types.YLeaf{"ClientSynchronizationEndOfDownloadPending", clientStatus.ClientSynchronizationEndOfDownloadPending})
    clientStatus.EntityData.Leafs.Append("up-time-stamp", types.YLeaf{"UpTimeStamp", clientStatus.UpTimeStamp})
    clientStatus.EntityData.Leafs.Append("down-time-stamp", types.YLeaf{"DownTimeStamp", clientStatus.DownTimeStamp})
    clientStatus.EntityData.Leafs.Append("clean-up-timer-remaining", types.YLeaf{"CleanUpTimerRemaining", clientStatus.CleanUpTimerRemaining})

    clientStatus.EntityData.YListKeys = []string {}

    return &(clientStatus.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus
// Opaque memory Stats
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus struct {
    EntityData types.CommonEntityData
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

func (opaqueMemoryStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_OpaqueMemoryStatus) GetEntityData() *types.CommonEntityData {
    opaqueMemoryStatus.EntityData.YFilter = opaqueMemoryStatus.YFilter
    opaqueMemoryStatus.EntityData.YangName = "opaque-memory-status"
    opaqueMemoryStatus.EntityData.BundleName = "cisco_ios_xr"
    opaqueMemoryStatus.EntityData.ParentYangName = "stats-global"
    opaqueMemoryStatus.EntityData.SegmentPath = "opaque-memory-status"
    opaqueMemoryStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opaqueMemoryStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opaqueMemoryStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opaqueMemoryStatus.EntityData.Children = types.NewOrderedMap()
    opaqueMemoryStatus.EntityData.Leafs = types.NewOrderedMap()
    opaqueMemoryStatus.EntityData.Leafs.Append("component", types.YLeaf{"Component", opaqueMemoryStatus.Component})
    opaqueMemoryStatus.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", opaqueMemoryStatus.SessionCount})
    opaqueMemoryStatus.EntityData.Leafs.Append("opaque-size", types.YLeaf{"OpaqueSize", opaqueMemoryStatus.OpaqueSize})
    opaqueMemoryStatus.EntityData.Leafs.Append("opaque-high-size", types.YLeaf{"OpaqueHighSize", opaqueMemoryStatus.OpaqueHighSize})
    opaqueMemoryStatus.EntityData.Leafs.Append("opaque-data-size", types.YLeaf{"OpaqueDataSize", opaqueMemoryStatus.OpaqueDataSize})

    opaqueMemoryStatus.EntityData.YListKeys = []string {}

    return &(opaqueMemoryStatus.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus
// TCP TxList Statistics
type SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus struct {
    EntityData types.CommonEntityData
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

func (txListOverTcpStatus *SessionRedundancyAgent_Nodes_Node_StatsGlobal_TxListOverTcpStatus) GetEntityData() *types.CommonEntityData {
    txListOverTcpStatus.EntityData.YFilter = txListOverTcpStatus.YFilter
    txListOverTcpStatus.EntityData.YangName = "tx-list-over-tcp-status"
    txListOverTcpStatus.EntityData.BundleName = "cisco_ios_xr"
    txListOverTcpStatus.EntityData.ParentYangName = "stats-global"
    txListOverTcpStatus.EntityData.SegmentPath = "tx-list-over-tcp-status"
    txListOverTcpStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txListOverTcpStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txListOverTcpStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txListOverTcpStatus.EntityData.Children = types.NewOrderedMap()
    txListOverTcpStatus.EntityData.Leafs = types.NewOrderedMap()
    txListOverTcpStatus.EntityData.Leafs.Append("messages-sent", types.YLeaf{"MessagesSent", txListOverTcpStatus.MessagesSent})
    txListOverTcpStatus.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", txListOverTcpStatus.BytesSent})
    txListOverTcpStatus.EntityData.Leafs.Append("messages-received", types.YLeaf{"MessagesReceived", txListOverTcpStatus.MessagesReceived})
    txListOverTcpStatus.EntityData.Leafs.Append("bytes-received", types.YLeaf{"BytesReceived", txListOverTcpStatus.BytesReceived})
    txListOverTcpStatus.EntityData.Leafs.Append("connect-count", types.YLeaf{"ConnectCount", txListOverTcpStatus.ConnectCount})
    txListOverTcpStatus.EntityData.Leafs.Append("blocked-connect-count", types.YLeaf{"BlockedConnectCount", txListOverTcpStatus.BlockedConnectCount})
    txListOverTcpStatus.EntityData.Leafs.Append("connect-retry-count", types.YLeaf{"ConnectRetryCount", txListOverTcpStatus.ConnectRetryCount})
    txListOverTcpStatus.EntityData.Leafs.Append("remote-node-down-count", types.YLeaf{"RemoteNodeDownCount", txListOverTcpStatus.RemoteNodeDownCount})
    txListOverTcpStatus.EntityData.Leafs.Append("accept-count", types.YLeaf{"AcceptCount", txListOverTcpStatus.AcceptCount})
    txListOverTcpStatus.EntityData.Leafs.Append("maximum-sent-message-size", types.YLeaf{"MaximumSentMessageSize", txListOverTcpStatus.MaximumSentMessageSize})
    txListOverTcpStatus.EntityData.Leafs.Append("maximum-received-message-size", types.YLeaf{"MaximumReceivedMessageSize", txListOverTcpStatus.MaximumReceivedMessageSize})
    txListOverTcpStatus.EntityData.Leafs.Append("peer-pause-count", types.YLeaf{"PeerPauseCount", txListOverTcpStatus.PeerPauseCount})
    txListOverTcpStatus.EntityData.Leafs.Append("buffer-full-occurence-count", types.YLeaf{"BufferFullOccurenceCount", txListOverTcpStatus.BufferFullOccurenceCount})
    txListOverTcpStatus.EntityData.Leafs.Append("mem-move-message-count", types.YLeaf{"MemMoveMessageCount", txListOverTcpStatus.MemMoveMessageCount})
    txListOverTcpStatus.EntityData.Leafs.Append("mem-move-bytes-count", types.YLeaf{"MemMoveBytesCount", txListOverTcpStatus.MemMoveBytesCount})
    txListOverTcpStatus.EntityData.Leafs.Append("socket-read-count", types.YLeaf{"SocketReadCount", txListOverTcpStatus.SocketReadCount})
    txListOverTcpStatus.EntityData.Leafs.Append("socket-write-count", types.YLeaf{"SocketWriteCount", txListOverTcpStatus.SocketWriteCount})
    txListOverTcpStatus.EntityData.Leafs.Append("active-socket-count", types.YLeaf{"ActiveSocketCount", txListOverTcpStatus.ActiveSocketCount})
    txListOverTcpStatus.EntityData.Leafs.Append("maximum-requested-buffer-size", types.YLeaf{"MaximumRequestedBufferSize", txListOverTcpStatus.MaximumRequestedBufferSize})
    txListOverTcpStatus.EntityData.Leafs.Append("buffer-freed-count", types.YLeaf{"BufferFreedCount", txListOverTcpStatus.BufferFreedCount})
    txListOverTcpStatus.EntityData.Leafs.Append("buffer-cache-hit", types.YLeaf{"BufferCacheHit", txListOverTcpStatus.BufferCacheHit})
    txListOverTcpStatus.EntityData.Leafs.Append("buffer-cache-miss", types.YLeaf{"BufferCacheMiss", txListOverTcpStatus.BufferCacheMiss})

    txListOverTcpStatus.EntityData.YListKeys = []string {}

    return &(txListOverTcpStatus.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupSummaries
// Session data for a particular node
type SessionRedundancyAgent_Nodes_Node_GroupSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session redundancy agent group summary. The type is slice of
    // SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary.
    GroupSummary []*SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary
}

func (groupSummaries *SessionRedundancyAgent_Nodes_Node_GroupSummaries) GetEntityData() *types.CommonEntityData {
    groupSummaries.EntityData.YFilter = groupSummaries.YFilter
    groupSummaries.EntityData.YangName = "group-summaries"
    groupSummaries.EntityData.BundleName = "cisco_ios_xr"
    groupSummaries.EntityData.ParentYangName = "node"
    groupSummaries.EntityData.SegmentPath = "group-summaries"
    groupSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupSummaries.EntityData.Children = types.NewOrderedMap()
    groupSummaries.EntityData.Children.Append("group-summary", types.YChild{"GroupSummary", nil})
    for i := range groupSummaries.GroupSummary {
        groupSummaries.EntityData.Children.Append(types.GetSegmentPath(groupSummaries.GroupSummary[i]), types.YChild{"GroupSummary", groupSummaries.GroupSummary[i]})
    }
    groupSummaries.EntityData.Leafs = types.NewOrderedMap()

    groupSummaries.EntityData.YListKeys = []string {}

    return &(groupSummaries.EntityData)
}

// SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary
// Session redundancy agent group summary
type SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary struct {
    EntityData types.CommonEntityData
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

func (groupSummary *SessionRedundancyAgent_Nodes_Node_GroupSummaries_GroupSummary) GetEntityData() *types.CommonEntityData {
    groupSummary.EntityData.YFilter = groupSummary.YFilter
    groupSummary.EntityData.YangName = "group-summary"
    groupSummary.EntityData.BundleName = "cisco_ios_xr"
    groupSummary.EntityData.ParentYangName = "group-summaries"
    groupSummary.EntityData.SegmentPath = "group-summary" + types.AddKeyToken(groupSummary.GroupId, "group-id")
    groupSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupSummary.EntityData.Children = types.NewOrderedMap()
    groupSummary.EntityData.Leafs = types.NewOrderedMap()
    groupSummary.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", groupSummary.GroupId})
    groupSummary.EntityData.Leafs.Append("group-id-xr", types.YLeaf{"GroupIdXr", groupSummary.GroupIdXr})
    groupSummary.EntityData.Leafs.Append("role", types.YLeaf{"Role", groupSummary.Role})
    groupSummary.EntityData.Leafs.Append("disabled", types.YLeaf{"Disabled", groupSummary.Disabled})
    groupSummary.EntityData.Leafs.Append("peer-ipv4-address", types.YLeaf{"PeerIpv4Address", groupSummary.PeerIpv4Address})
    groupSummary.EntityData.Leafs.Append("peer-ipv6-address", types.YLeaf{"PeerIpv6Address", groupSummary.PeerIpv6Address})
    groupSummary.EntityData.Leafs.Append("peer-status", types.YLeaf{"PeerStatus", groupSummary.PeerStatus})
    groupSummary.EntityData.Leafs.Append("preferred-role", types.YLeaf{"PreferredRole", groupSummary.PreferredRole})
    groupSummary.EntityData.Leafs.Append("slave-mode", types.YLeaf{"SlaveMode", groupSummary.SlaveMode})
    groupSummary.EntityData.Leafs.Append("object-tracking-status", types.YLeaf{"ObjectTrackingStatus", groupSummary.ObjectTrackingStatus})
    groupSummary.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", groupSummary.InterfaceCount})
    groupSummary.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", groupSummary.SessionCount})
    groupSummary.EntityData.Leafs.Append("pending-add-session-count", types.YLeaf{"PendingAddSessionCount", groupSummary.PendingAddSessionCount})

    groupSummary.EntityData.YListKeys = []string {"GroupId"}

    return &(groupSummary.EntityData)
}

