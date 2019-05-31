// This module contains a collection of YANG definitions
// for Cisco IOS-XR procfind package operational data.
// 
// This module contains definitions
// for the following management objects:
//   proc-distribution: Process distribution information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package procfind_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package procfind_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-procfind-oper proc-distribution}", reflect.TypeOf(ProcDistribution{}))
    ydk.RegisterEntity("Cisco-IOS-XR-procfind-oper:proc-distribution", reflect.TypeOf(ProcDistribution{}))
}

// ProcDistribution
// Process distribution information
type ProcDistribution struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes.
    Nodes ProcDistribution_Nodes
}

func (procDistribution *ProcDistribution) GetEntityData() *types.CommonEntityData {
    procDistribution.EntityData.YFilter = procDistribution.YFilter
    procDistribution.EntityData.YangName = "proc-distribution"
    procDistribution.EntityData.BundleName = "cisco_ios_xr"
    procDistribution.EntityData.ParentYangName = "Cisco-IOS-XR-procfind-oper"
    procDistribution.EntityData.SegmentPath = "Cisco-IOS-XR-procfind-oper:proc-distribution"
    procDistribution.EntityData.AbsolutePath = procDistribution.EntityData.SegmentPath
    procDistribution.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procDistribution.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procDistribution.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procDistribution.EntityData.Children = types.NewOrderedMap()
    procDistribution.EntityData.Children.Append("nodes", types.YChild{"Nodes", &procDistribution.Nodes})
    procDistribution.EntityData.Leafs = types.NewOrderedMap()

    procDistribution.EntityData.YListKeys = []string {}

    return &(procDistribution.EntityData)
}

// ProcDistribution_Nodes
// List of nodes
type ProcDistribution_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process distribution information per node. The type is slice of
    // ProcDistribution_Nodes_Node.
    Node []*ProcDistribution_Nodes_Node
}

func (nodes *ProcDistribution_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "proc-distribution"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-procfind-oper:proc-distribution/" + nodes.EntityData.SegmentPath
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

// ProcDistribution_Nodes_Node
// Process distribution information per node
type ProcDistribution_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Process distribution information. The type is slice of
    // ProcDistribution_Nodes_Node_Process.
    Process []*ProcDistribution_Nodes_Node_Process
}

func (node *ProcDistribution_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-procfind-oper:proc-distribution/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range node.Process {
        node.EntityData.Children.Append(types.GetSegmentPath(node.Process[i]), types.YChild{"Process", node.Process[i]})
    }
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// ProcDistribution_Nodes_Node_Process
// Process distribution information
type ProcDistribution_Nodes_Node_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process Name. The type is string.
    ProcName interface{}

    // Process distribution information. The type is slice of
    // ProcDistribution_Nodes_Node_Process_FilterType.
    FilterType []*ProcDistribution_Nodes_Node_Process_FilterType
}

func (process *ProcDistribution_Nodes_Node_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "node"
    process.EntityData.SegmentPath = "process" + types.AddKeyToken(process.ProcName, "proc-name")
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-procfind-oper:proc-distribution/nodes/node/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("filter-type", types.YChild{"FilterType", nil})
    for i := range process.FilterType {
        process.EntityData.Children.Append(types.GetSegmentPath(process.FilterType[i]), types.YChild{"FilterType", process.FilterType[i]})
    }
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", process.ProcName})

    process.EntityData.YListKeys = []string {"ProcName"}

    return &(process.EntityData)
}

// ProcDistribution_Nodes_Node_Process_FilterType
// Process distribution information
type ProcDistribution_Nodes_Node_Process_FilterType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Filter Type. The type is string.
    FilterType interface{}

    // Node ID. The type is interface{} with range: 0..4294967295.
    Nodeid interface{}

    // Node type. The type is interface{} with range: 0..4294967295.
    Nodetype interface{}

    // Process ID. The type is interface{} with range: -2147483648..2147483647.
    Pid interface{}

    // Job ID. The type is interface{} with range: -2147483648..2147483647.
    Jid interface{}

    // Number of threads. The type is interface{} with range:
    // -2147483648..2147483647.
    NumThreads interface{}

    // Process name. The type is string.
    Name interface{}
}

func (filterType *ProcDistribution_Nodes_Node_Process_FilterType) GetEntityData() *types.CommonEntityData {
    filterType.EntityData.YFilter = filterType.YFilter
    filterType.EntityData.YangName = "filter-type"
    filterType.EntityData.BundleName = "cisco_ios_xr"
    filterType.EntityData.ParentYangName = "process"
    filterType.EntityData.SegmentPath = "filter-type" + types.AddKeyToken(filterType.FilterType, "filter-type")
    filterType.EntityData.AbsolutePath = "Cisco-IOS-XR-procfind-oper:proc-distribution/nodes/node/process/" + filterType.EntityData.SegmentPath
    filterType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filterType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filterType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filterType.EntityData.Children = types.NewOrderedMap()
    filterType.EntityData.Leafs = types.NewOrderedMap()
    filterType.EntityData.Leafs.Append("filter-type", types.YLeaf{"FilterType", filterType.FilterType})
    filterType.EntityData.Leafs.Append("nodeid", types.YLeaf{"Nodeid", filterType.Nodeid})
    filterType.EntityData.Leafs.Append("nodetype", types.YLeaf{"Nodetype", filterType.Nodetype})
    filterType.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", filterType.Pid})
    filterType.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", filterType.Jid})
    filterType.EntityData.Leafs.Append("num-threads", types.YLeaf{"NumThreads", filterType.NumThreads})
    filterType.EntityData.Leafs.Append("name", types.YLeaf{"Name", filterType.Name})

    filterType.EntityData.YListKeys = []string {"FilterType"}

    return &(filterType.EntityData)
}

