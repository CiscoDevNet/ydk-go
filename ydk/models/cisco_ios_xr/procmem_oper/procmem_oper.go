// This module contains a collection of YANG definitions
// for Cisco IOS-XR procmem package operational data.
// 
// This module contains definitions
// for the following management objects:
//   processes-memory: Process statistics
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    processesMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processesMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processesMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processesMemory.EntityData.Children = make(map[string]types.YChild)
    processesMemory.EntityData.Children["nodes"] = types.YChild{"Nodes", &processesMemory.Nodes}
    processesMemory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processesMemory.EntityData)
}

// ProcessesMemory_Nodes
// List of nodes
type ProcessesMemory_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node ID. The type is slice of ProcessesMemory_Nodes_Node.
    Node []ProcessesMemory_Nodes_Node
}

func (nodes *ProcessesMemory_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "processes-memory"
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

// ProcessesMemory_Nodes_Node
// Node ID
type ProcessesMemory_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["process-ids"] = types.YChild{"ProcessIds", &node.ProcessIds}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// ProcessesMemory_Nodes_Node_ProcessIds
// List of jobs
type ProcessesMemory_Nodes_Node_ProcessIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process Id. The type is slice of
    // ProcessesMemory_Nodes_Node_ProcessIds_ProcessId.
    ProcessId []ProcessesMemory_Nodes_Node_ProcessIds_ProcessId
}

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetEntityData() *types.CommonEntityData {
    processIds.EntityData.YFilter = processIds.YFilter
    processIds.EntityData.YangName = "process-ids"
    processIds.EntityData.BundleName = "cisco_ios_xr"
    processIds.EntityData.ParentYangName = "node"
    processIds.EntityData.SegmentPath = "process-ids"
    processIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processIds.EntityData.Children = make(map[string]types.YChild)
    processIds.EntityData.Children["process-id"] = types.YChild{"ProcessId", nil}
    for i := range processIds.ProcessId {
        processIds.EntityData.Children[types.GetSegmentPath(&processIds.ProcessId[i])] = types.YChild{"ProcessId", &processIds.ProcessId[i]}
    }
    processIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processIds.EntityData)
}

// ProcessesMemory_Nodes_Node_ProcessIds_ProcessId
// Process Id
type ProcessesMemory_Nodes_Node_ProcessIds_ProcessId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Process Id. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcessId interface{}

    // Process name. The type is string.
    Name interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Text Segment Size. The type is interface{} with range: 0..4294967295.
    TextSegSize interface{}

    // Data Segment Size. The type is interface{} with range: 0..4294967295.
    DataSegSize interface{}

    // Stack Segment Size. The type is interface{} with range: 0..4294967295.
    StackSegSize interface{}

    // Malloced Memory Size. The type is interface{} with range: 0..4294967295.
    MallocSize interface{}

    // Dynamic memory limit. The type is interface{} with range: 0..4294967295.
    DynLimit interface{}

    // Shared memory size. The type is interface{} with range: 0..4294967295.
    SharedMem interface{}

    // Physical memory size. The type is interface{} with range: 0..4294967295.
    PhysicalMem interface{}
}

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetEntityData() *types.CommonEntityData {
    processId.EntityData.YFilter = processId.YFilter
    processId.EntityData.YangName = "process-id"
    processId.EntityData.BundleName = "cisco_ios_xr"
    processId.EntityData.ParentYangName = "process-ids"
    processId.EntityData.SegmentPath = "process-id" + "[process-id='" + fmt.Sprintf("%v", processId.ProcessId) + "']"
    processId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processId.EntityData.Children = make(map[string]types.YChild)
    processId.EntityData.Leafs = make(map[string]types.YLeaf)
    processId.EntityData.Leafs["process-id"] = types.YLeaf{"ProcessId", processId.ProcessId}
    processId.EntityData.Leafs["name"] = types.YLeaf{"Name", processId.Name}
    processId.EntityData.Leafs["jid"] = types.YLeaf{"Jid", processId.Jid}
    processId.EntityData.Leafs["pid"] = types.YLeaf{"Pid", processId.Pid}
    processId.EntityData.Leafs["text-seg-size"] = types.YLeaf{"TextSegSize", processId.TextSegSize}
    processId.EntityData.Leafs["data-seg-size"] = types.YLeaf{"DataSegSize", processId.DataSegSize}
    processId.EntityData.Leafs["stack-seg-size"] = types.YLeaf{"StackSegSize", processId.StackSegSize}
    processId.EntityData.Leafs["malloc-size"] = types.YLeaf{"MallocSize", processId.MallocSize}
    processId.EntityData.Leafs["dyn-limit"] = types.YLeaf{"DynLimit", processId.DynLimit}
    processId.EntityData.Leafs["shared-mem"] = types.YLeaf{"SharedMem", processId.SharedMem}
    processId.EntityData.Leafs["physical-mem"] = types.YLeaf{"PhysicalMem", processId.PhysicalMem}
    return &(processId.EntityData)
}

