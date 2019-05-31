// This module contains a collection of YANG definitions
// for Cisco IOS-XR procmem package operational data.
// 
// This module contains definitions
// for the following management objects:
//   processes-memory: Process statistics
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package procmem_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package procmem_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-procmem-oper processes-memory}", reflect.TypeOf(ProcessesMemory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-procmem-oper:processes-memory", reflect.TypeOf(ProcessesMemory{}))
}

// ProcessesMemory
// Process statistics
type ProcessesMemory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes.
    Nodes ProcessesMemory_Nodes
}

func (processesMemory *ProcessesMemory) GetEntityData() *types.CommonEntityData {
    processesMemory.EntityData.YFilter = processesMemory.YFilter
    processesMemory.EntityData.YangName = "processes-memory"
    processesMemory.EntityData.BundleName = "cisco_ios_xr"
    processesMemory.EntityData.ParentYangName = "Cisco-IOS-XR-procmem-oper"
    processesMemory.EntityData.SegmentPath = "Cisco-IOS-XR-procmem-oper:processes-memory"
    processesMemory.EntityData.AbsolutePath = processesMemory.EntityData.SegmentPath
    processesMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processesMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processesMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processesMemory.EntityData.Children = types.NewOrderedMap()
    processesMemory.EntityData.Children.Append("nodes", types.YChild{"Nodes", &processesMemory.Nodes})
    processesMemory.EntityData.Leafs = types.NewOrderedMap()

    processesMemory.EntityData.YListKeys = []string {}

    return &(processesMemory.EntityData)
}

// ProcessesMemory_Nodes
// List of nodes
type ProcessesMemory_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node ID. The type is slice of ProcessesMemory_Nodes_Node.
    Node []*ProcessesMemory_Nodes_Node
}

func (nodes *ProcessesMemory_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "processes-memory"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-procmem-oper:processes-memory/" + nodes.EntityData.SegmentPath
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

// ProcessesMemory_Nodes_Node
// Node ID
type ProcessesMemory_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // List of jobs.
    ProcessIds ProcessesMemory_Nodes_Node_ProcessIds
}

func (node *ProcessesMemory_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-procmem-oper:processes-memory/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("process-ids", types.YChild{"ProcessIds", &node.ProcessIds})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// ProcessesMemory_Nodes_Node_ProcessIds
// List of jobs
type ProcessesMemory_Nodes_Node_ProcessIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process Id. The type is slice of
    // ProcessesMemory_Nodes_Node_ProcessIds_ProcessId.
    ProcessId []*ProcessesMemory_Nodes_Node_ProcessIds_ProcessId
}

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetEntityData() *types.CommonEntityData {
    processIds.EntityData.YFilter = processIds.YFilter
    processIds.EntityData.YangName = "process-ids"
    processIds.EntityData.BundleName = "cisco_ios_xr"
    processIds.EntityData.ParentYangName = "node"
    processIds.EntityData.SegmentPath = "process-ids"
    processIds.EntityData.AbsolutePath = "Cisco-IOS-XR-procmem-oper:processes-memory/nodes/node/" + processIds.EntityData.SegmentPath
    processIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processIds.EntityData.Children = types.NewOrderedMap()
    processIds.EntityData.Children.Append("process-id", types.YChild{"ProcessId", nil})
    for i := range processIds.ProcessId {
        processIds.EntityData.Children.Append(types.GetSegmentPath(processIds.ProcessId[i]), types.YChild{"ProcessId", processIds.ProcessId[i]})
    }
    processIds.EntityData.Leafs = types.NewOrderedMap()

    processIds.EntityData.YListKeys = []string {}

    return &(processIds.EntityData)
}

// ProcessesMemory_Nodes_Node_ProcessIds_ProcessId
// Process Id
type ProcessesMemory_Nodes_Node_ProcessIds_ProcessId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process Id. The type is interface{} with range:
    // 0..4294967295.
    ProcessId interface{}

    // Process name. The type is string.
    Name interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Text Segment Size in KB. The type is interface{} with range: 0..4294967295.
    TextSegSize interface{}

    // Data Segment Size in KB. The type is interface{} with range: 0..4294967295.
    DataSegSize interface{}

    // Stack Segment Size in KB. The type is interface{} with range:
    // 0..4294967295.
    StackSegSize interface{}

    // Malloced Memory Size in KB. The type is interface{} with range:
    // 0..4294967295.
    MallocSize interface{}

    // Dynamic memory limit in KB (4294967295 for RLIM_INFINITY). The type is
    // interface{} with range: 0..4294967295.
    DynLimit interface{}

    // Shared memory size in KB. The type is interface{} with range:
    // 0..4294967295.
    SharedMem interface{}

    // Physical memory size in KB. The type is interface{} with range:
    // 0..4294967295.
    PhysicalMem interface{}
}

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetEntityData() *types.CommonEntityData {
    processId.EntityData.YFilter = processId.YFilter
    processId.EntityData.YangName = "process-id"
    processId.EntityData.BundleName = "cisco_ios_xr"
    processId.EntityData.ParentYangName = "process-ids"
    processId.EntityData.SegmentPath = "process-id" + types.AddKeyToken(processId.ProcessId, "process-id")
    processId.EntityData.AbsolutePath = "Cisco-IOS-XR-procmem-oper:processes-memory/nodes/node/process-ids/" + processId.EntityData.SegmentPath
    processId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processId.EntityData.Children = types.NewOrderedMap()
    processId.EntityData.Leafs = types.NewOrderedMap()
    processId.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", processId.ProcessId})
    processId.EntityData.Leafs.Append("name", types.YLeaf{"Name", processId.Name})
    processId.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", processId.Jid})
    processId.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", processId.Pid})
    processId.EntityData.Leafs.Append("text-seg-size", types.YLeaf{"TextSegSize", processId.TextSegSize})
    processId.EntityData.Leafs.Append("data-seg-size", types.YLeaf{"DataSegSize", processId.DataSegSize})
    processId.EntityData.Leafs.Append("stack-seg-size", types.YLeaf{"StackSegSize", processId.StackSegSize})
    processId.EntityData.Leafs.Append("malloc-size", types.YLeaf{"MallocSize", processId.MallocSize})
    processId.EntityData.Leafs.Append("dyn-limit", types.YLeaf{"DynLimit", processId.DynLimit})
    processId.EntityData.Leafs.Append("shared-mem", types.YLeaf{"SharedMem", processId.SharedMem})
    processId.EntityData.Leafs.Append("physical-mem", types.YLeaf{"PhysicalMem", processId.PhysicalMem})

    processId.EntityData.YListKeys = []string {"ProcessId"}

    return &(processId.EntityData)
}

