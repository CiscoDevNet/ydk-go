// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-rcmd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   router-convergence: Configure Router Convergence Monitoring
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_rcmd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_rcmd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-rcmd-cfg router-convergence}", reflect.TypeOf(RouterConvergence{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-rcmd-cfg:router-convergence", reflect.TypeOf(RouterConvergence{}))
}

// RcmdPriority represents Rcmd priority
type RcmdPriority string

const (
    // Critical routes
    RcmdPriority_critical RcmdPriority = "critical"

    // High priority routes
    RcmdPriority_high RcmdPriority = "high"

    // Medium priority routes
    RcmdPriority_medium RcmdPriority = "medium"

    // Low priority routes
    RcmdPriority_low RcmdPriority = "low"
)

// ProtocolName represents Protocol name
type ProtocolName string

const (
    // Configure parameters related to OSPF
    ProtocolName_ospf ProtocolName = "ospf"

    // Configure parameters related to ISIS
    ProtocolName_isis ProtocolName = "isis"
)

// RouterConvergence
// Configure Router Convergence Monitoring
type RouterConvergence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event buffer size for storing event traces (as number of events). The type
    // is interface{} with range: 100..500.
    EventBufferSize interface{}

    // Limits Individual Prefix Monitoring. The type is interface{} with range:
    // 0..100.
    PrefixMonitorLimit interface{}

    // Disable the monitoring of route convergence on the entire router. The type
    // is interface{}.
    Disable interface{}

    // Enable Configure Router Convergence Monitoring. Deletion of this object
    // also causes deletion of all associated objects under RouterConvergence. The
    // type is interface{}.
    Enable interface{}

    // Maximum number of events stored in the server. The type is interface{} with
    // range: 10..500.
    MaxEventsStored interface{}

    // Interval in which to collect logs (in mins). The type is interface{} with
    // range: 5..120. Units are minute.
    MonitoringInterval interface{}

    // Table of Protocol.
    Protocols RouterConvergence_Protocols

    // Absolute directory path for saving the archive files. Example /disk0:/rcmd/
    // or <tftp-location>/rcmd/.
    StorageLocation RouterConvergence_StorageLocation

    // RCMD related configuration for MPLS-LDP.
    MplsLdp RouterConvergence_MplsLdp

    // Table of CollectDiagnostics.
    CollectDiagnostics RouterConvergence_CollectDiagnostics

    // Table of Node.
    Nodes RouterConvergence_Nodes
}

func (routerConvergence *RouterConvergence) GetEntityData() *types.CommonEntityData {
    routerConvergence.EntityData.YFilter = routerConvergence.YFilter
    routerConvergence.EntityData.YangName = "router-convergence"
    routerConvergence.EntityData.BundleName = "cisco_ios_xr"
    routerConvergence.EntityData.ParentYangName = "Cisco-IOS-XR-infra-rcmd-cfg"
    routerConvergence.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence"
    routerConvergence.EntityData.AbsolutePath = routerConvergence.EntityData.SegmentPath
    routerConvergence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routerConvergence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routerConvergence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routerConvergence.EntityData.Children = types.NewOrderedMap()
    routerConvergence.EntityData.Children.Append("protocols", types.YChild{"Protocols", &routerConvergence.Protocols})
    routerConvergence.EntityData.Children.Append("storage-location", types.YChild{"StorageLocation", &routerConvergence.StorageLocation})
    routerConvergence.EntityData.Children.Append("mpls-ldp", types.YChild{"MplsLdp", &routerConvergence.MplsLdp})
    routerConvergence.EntityData.Children.Append("collect-diagnostics", types.YChild{"CollectDiagnostics", &routerConvergence.CollectDiagnostics})
    routerConvergence.EntityData.Children.Append("nodes", types.YChild{"Nodes", &routerConvergence.Nodes})
    routerConvergence.EntityData.Leafs = types.NewOrderedMap()
    routerConvergence.EntityData.Leafs.Append("event-buffer-size", types.YLeaf{"EventBufferSize", routerConvergence.EventBufferSize})
    routerConvergence.EntityData.Leafs.Append("prefix-monitor-limit", types.YLeaf{"PrefixMonitorLimit", routerConvergence.PrefixMonitorLimit})
    routerConvergence.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", routerConvergence.Disable})
    routerConvergence.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", routerConvergence.Enable})
    routerConvergence.EntityData.Leafs.Append("max-events-stored", types.YLeaf{"MaxEventsStored", routerConvergence.MaxEventsStored})
    routerConvergence.EntityData.Leafs.Append("monitoring-interval", types.YLeaf{"MonitoringInterval", routerConvergence.MonitoringInterval})

    routerConvergence.EntityData.YListKeys = []string {}

    return &(routerConvergence.EntityData)
}

// RouterConvergence_Protocols
// Table of Protocol
type RouterConvergence_Protocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol for which to configure RCMD parameters. The type is slice of
    // RouterConvergence_Protocols_Protocol.
    Protocol []*RouterConvergence_Protocols_Protocol
}

func (protocols *RouterConvergence_Protocols) GetEntityData() *types.CommonEntityData {
    protocols.EntityData.YFilter = protocols.YFilter
    protocols.EntityData.YangName = "protocols"
    protocols.EntityData.BundleName = "cisco_ios_xr"
    protocols.EntityData.ParentYangName = "router-convergence"
    protocols.EntityData.SegmentPath = "protocols"
    protocols.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/" + protocols.EntityData.SegmentPath
    protocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocols.EntityData.Children = types.NewOrderedMap()
    protocols.EntityData.Children.Append("protocol", types.YChild{"Protocol", nil})
    for i := range protocols.Protocol {
        protocols.EntityData.Children.Append(types.GetSegmentPath(protocols.Protocol[i]), types.YChild{"Protocol", protocols.Protocol[i]})
    }
    protocols.EntityData.Leafs = types.NewOrderedMap()

    protocols.EntityData.YListKeys = []string {}

    return &(protocols.EntityData)
}

// RouterConvergence_Protocols_Protocol
// Protocol for which to configure RCMD parameters
type RouterConvergence_Protocols_Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Specify the protocol. The type is ProtocolName.
    ProtocolName interface{}

    // Enable Protocol for which to configure RCMD parameters. Deletion of this
    // object also causes deletion of all associated objects under Protocol. The
    // type is interface{}.
    Enable interface{}

    // Table of Priority.
    Priorities RouterConvergence_Protocols_Protocol_Priorities
}

func (protocol *RouterConvergence_Protocols_Protocol) GetEntityData() *types.CommonEntityData {
    protocol.EntityData.YFilter = protocol.YFilter
    protocol.EntityData.YangName = "protocol"
    protocol.EntityData.BundleName = "cisco_ios_xr"
    protocol.EntityData.ParentYangName = "protocols"
    protocol.EntityData.SegmentPath = "protocol" + types.AddKeyToken(protocol.ProtocolName, "protocol-name")
    protocol.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/protocols/" + protocol.EntityData.SegmentPath
    protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocol.EntityData.Children = types.NewOrderedMap()
    protocol.EntityData.Children.Append("priorities", types.YChild{"Priorities", &protocol.Priorities})
    protocol.EntityData.Leafs = types.NewOrderedMap()
    protocol.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", protocol.ProtocolName})
    protocol.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", protocol.Enable})

    protocol.EntityData.YListKeys = []string {"ProtocolName"}

    return &(protocol.EntityData)
}

// RouterConvergence_Protocols_Protocol_Priorities
// Table of Priority
type RouterConvergence_Protocols_Protocol_Priorities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority. The type is slice of
    // RouterConvergence_Protocols_Protocol_Priorities_Priority.
    Priority []*RouterConvergence_Protocols_Protocol_Priorities_Priority
}

func (priorities *RouterConvergence_Protocols_Protocol_Priorities) GetEntityData() *types.CommonEntityData {
    priorities.EntityData.YFilter = priorities.YFilter
    priorities.EntityData.YangName = "priorities"
    priorities.EntityData.BundleName = "cisco_ios_xr"
    priorities.EntityData.ParentYangName = "protocol"
    priorities.EntityData.SegmentPath = "priorities"
    priorities.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/protocols/protocol/" + priorities.EntityData.SegmentPath
    priorities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priorities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priorities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priorities.EntityData.Children = types.NewOrderedMap()
    priorities.EntityData.Children.Append("priority", types.YChild{"Priority", nil})
    for i := range priorities.Priority {
        priorities.EntityData.Children.Append(types.GetSegmentPath(priorities.Priority[i]), types.YChild{"Priority", priorities.Priority[i]})
    }
    priorities.EntityData.Leafs = types.NewOrderedMap()

    priorities.EntityData.YListKeys = []string {}

    return &(priorities.EntityData)
}

// RouterConvergence_Protocols_Protocol_Priorities_Priority
// Priority
type RouterConvergence_Protocols_Protocol_Priorities_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Specify the priority. The type is RcmdPriority.
    RcmdPriority interface{}

    // Threshold value for convergence (in msec). The type is interface{} with
    // range: 0..4294967295.
    Threshold interface{}

    // Specify the maximum number of leaf networks monitored. The type is
    // interface{} with range: 10..100.
    LeafNetworks interface{}

    // Disables the monitoring of route convergence for specified priority. The
    // type is interface{}.
    Disable interface{}

    // Enable Priority. Deletion of this object also causes deletion of all
    // associated objects under Priority. The type is interface{}.
    Enable interface{}

    // Threshold value for Fast ReRoute Coverage (in percentage). The type is
    // interface{} with range: 1..100. Units are percentage.
    FrrThreshold interface{}
}

func (priority *RouterConvergence_Protocols_Protocol_Priorities_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "priorities"
    priority.EntityData.SegmentPath = "priority" + types.AddKeyToken(priority.RcmdPriority, "rcmd-priority")
    priority.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/protocols/protocol/priorities/" + priority.EntityData.SegmentPath
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("rcmd-priority", types.YLeaf{"RcmdPriority", priority.RcmdPriority})
    priority.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", priority.Threshold})
    priority.EntityData.Leafs.Append("leaf-networks", types.YLeaf{"LeafNetworks", priority.LeafNetworks})
    priority.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", priority.Disable})
    priority.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", priority.Enable})
    priority.EntityData.Leafs.Append("frr-threshold", types.YLeaf{"FrrThreshold", priority.FrrThreshold})

    priority.EntityData.YListKeys = []string {"RcmdPriority"}

    return &(priority.EntityData)
}

// RouterConvergence_StorageLocation
// Absolute directory path for saving the archive
// files. Example /disk0:/rcmd/ or
// <tftp-location>/rcmd/
// This type is a presence type.
type RouterConvergence_StorageLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Absolute directory path for storing diagnostic reports. Example
    // /disk0:/rcmd/ or <tftp-location>/rcmd/. The type is string.
    Diagnostics interface{}

    // Maximum size of diagnostics dir (5% - 80%) for local storage. The type is
    // interface{} with range: 5..80.
    DiagnosticsSize interface{}

    // Maximum size of reports dir (5% - 80%) for local storage. The type is
    // interface{} with range: 5..80.
    ReportsSize interface{}

    // Absolute directory path for storing reports. Example /disk0:/rcmd/ or
    // <tftp-location>/rcmd/. The type is string.
    Reports interface{}
}

func (storageLocation *RouterConvergence_StorageLocation) GetEntityData() *types.CommonEntityData {
    storageLocation.EntityData.YFilter = storageLocation.YFilter
    storageLocation.EntityData.YangName = "storage-location"
    storageLocation.EntityData.BundleName = "cisco_ios_xr"
    storageLocation.EntityData.ParentYangName = "router-convergence"
    storageLocation.EntityData.SegmentPath = "storage-location"
    storageLocation.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/" + storageLocation.EntityData.SegmentPath
    storageLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    storageLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    storageLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    storageLocation.EntityData.Children = types.NewOrderedMap()
    storageLocation.EntityData.Leafs = types.NewOrderedMap()
    storageLocation.EntityData.Leafs.Append("diagnostics", types.YLeaf{"Diagnostics", storageLocation.Diagnostics})
    storageLocation.EntityData.Leafs.Append("diagnostics-size", types.YLeaf{"DiagnosticsSize", storageLocation.DiagnosticsSize})
    storageLocation.EntityData.Leafs.Append("reports-size", types.YLeaf{"ReportsSize", storageLocation.ReportsSize})
    storageLocation.EntityData.Leafs.Append("reports", types.YLeaf{"Reports", storageLocation.Reports})

    storageLocation.EntityData.YListKeys = []string {}

    return &(storageLocation.EntityData)
}

// RouterConvergence_MplsLdp
// RCMD related configuration for MPLS-LDP
// This type is a presence type.
type RouterConvergence_MplsLdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Monitoring configuration for Remote LFA.
    RemoteLfa RouterConvergence_MplsLdp_RemoteLfa
}

func (mplsLdp *RouterConvergence_MplsLdp) GetEntityData() *types.CommonEntityData {
    mplsLdp.EntityData.YFilter = mplsLdp.YFilter
    mplsLdp.EntityData.YangName = "mpls-ldp"
    mplsLdp.EntityData.BundleName = "cisco_ios_xr"
    mplsLdp.EntityData.ParentYangName = "router-convergence"
    mplsLdp.EntityData.SegmentPath = "mpls-ldp"
    mplsLdp.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/" + mplsLdp.EntityData.SegmentPath
    mplsLdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLdp.EntityData.Children = types.NewOrderedMap()
    mplsLdp.EntityData.Children.Append("remote-lfa", types.YChild{"RemoteLfa", &mplsLdp.RemoteLfa})
    mplsLdp.EntityData.Leafs = types.NewOrderedMap()

    mplsLdp.EntityData.YListKeys = []string {}

    return &(mplsLdp.EntityData)
}

// RouterConvergence_MplsLdp_RemoteLfa
// Monitoring configuration for Remote LFA
// This type is a presence type.
type RouterConvergence_MplsLdp_RemoteLfa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Threshold value for label coverage (in percentage). The type is interface{}
    // with range: 1..100. Units are percentage.
    Threshold interface{}
}

func (remoteLfa *RouterConvergence_MplsLdp_RemoteLfa) GetEntityData() *types.CommonEntityData {
    remoteLfa.EntityData.YFilter = remoteLfa.YFilter
    remoteLfa.EntityData.YangName = "remote-lfa"
    remoteLfa.EntityData.BundleName = "cisco_ios_xr"
    remoteLfa.EntityData.ParentYangName = "mpls-ldp"
    remoteLfa.EntityData.SegmentPath = "remote-lfa"
    remoteLfa.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/mpls-ldp/" + remoteLfa.EntityData.SegmentPath
    remoteLfa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteLfa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteLfa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteLfa.EntityData.Children = types.NewOrderedMap()
    remoteLfa.EntityData.Leafs = types.NewOrderedMap()
    remoteLfa.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", remoteLfa.Threshold})

    remoteLfa.EntityData.YListKeys = []string {}

    return &(remoteLfa.EntityData)
}

// RouterConvergence_CollectDiagnostics
// Table of CollectDiagnostics
type RouterConvergence_CollectDiagnostics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collect diagnostics on specified node. The type is slice of
    // RouterConvergence_CollectDiagnostics_CollectDiagnostic.
    CollectDiagnostic []*RouterConvergence_CollectDiagnostics_CollectDiagnostic
}

func (collectDiagnostics *RouterConvergence_CollectDiagnostics) GetEntityData() *types.CommonEntityData {
    collectDiagnostics.EntityData.YFilter = collectDiagnostics.YFilter
    collectDiagnostics.EntityData.YangName = "collect-diagnostics"
    collectDiagnostics.EntityData.BundleName = "cisco_ios_xr"
    collectDiagnostics.EntityData.ParentYangName = "router-convergence"
    collectDiagnostics.EntityData.SegmentPath = "collect-diagnostics"
    collectDiagnostics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/" + collectDiagnostics.EntityData.SegmentPath
    collectDiagnostics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collectDiagnostics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collectDiagnostics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collectDiagnostics.EntityData.Children = types.NewOrderedMap()
    collectDiagnostics.EntityData.Children.Append("collect-diagnostic", types.YChild{"CollectDiagnostic", nil})
    for i := range collectDiagnostics.CollectDiagnostic {
        collectDiagnostics.EntityData.Children.Append(types.GetSegmentPath(collectDiagnostics.CollectDiagnostic[i]), types.YChild{"CollectDiagnostic", collectDiagnostics.CollectDiagnostic[i]})
    }
    collectDiagnostics.EntityData.Leafs = types.NewOrderedMap()

    collectDiagnostics.EntityData.YListKeys = []string {}

    return &(collectDiagnostics.EntityData)
}

// RouterConvergence_CollectDiagnostics_CollectDiagnostic
// Collect diagnostics on specified node
type RouterConvergence_CollectDiagnostics_CollectDiagnostic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Specified location. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Enables collection of diagnostics on the specified location. The type is
    // interface{}.
    Enable interface{}
}

func (collectDiagnostic *RouterConvergence_CollectDiagnostics_CollectDiagnostic) GetEntityData() *types.CommonEntityData {
    collectDiagnostic.EntityData.YFilter = collectDiagnostic.YFilter
    collectDiagnostic.EntityData.YangName = "collect-diagnostic"
    collectDiagnostic.EntityData.BundleName = "cisco_ios_xr"
    collectDiagnostic.EntityData.ParentYangName = "collect-diagnostics"
    collectDiagnostic.EntityData.SegmentPath = "collect-diagnostic" + types.AddKeyToken(collectDiagnostic.NodeName, "node-name")
    collectDiagnostic.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/collect-diagnostics/" + collectDiagnostic.EntityData.SegmentPath
    collectDiagnostic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collectDiagnostic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collectDiagnostic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collectDiagnostic.EntityData.Children = types.NewOrderedMap()
    collectDiagnostic.EntityData.Leafs = types.NewOrderedMap()
    collectDiagnostic.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", collectDiagnostic.NodeName})
    collectDiagnostic.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", collectDiagnostic.Enable})

    collectDiagnostic.EntityData.YListKeys = []string {"NodeName"}

    return &(collectDiagnostic.EntityData)
}

// RouterConvergence_Nodes
// Table of Node
type RouterConvergence_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure parameters for the specified node (Partially qualified location
    // allowed). The type is slice of RouterConvergence_Nodes_Node.
    Node []*RouterConvergence_Nodes_Node
}

func (nodes *RouterConvergence_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "router-convergence"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/" + nodes.EntityData.SegmentPath
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

// RouterConvergence_Nodes_Node
// Configure parameters for the specified node
// (Partially qualified location allowed)
type RouterConvergence_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Wildcard expression(eg. */*/*, R/*/*, R/S/*,
    // R/S/I). The type is string with pattern:
    // ((([a-zA-Z0-9_]*\d+)|(\*))/){2}(([a-zA-Z0-9_]*\d+)|(\*)).
    NodeName interface{}

    // Disables the monitoring of route convergence on specified location. The
    // type is interface{}.
    Disable interface{}

    // Enable Configure parameters for the specified node (Partially qualified
    // location allowed). Deletion of this object also causes deletion of all
    // associated objects under Node. The type is interface{}.
    Enable interface{}
}

func (node *RouterConvergence_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-rcmd-cfg:router-convergence/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", node.Disable})
    node.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", node.Enable})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

