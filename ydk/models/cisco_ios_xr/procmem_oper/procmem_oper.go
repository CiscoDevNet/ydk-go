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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes.
    Nodes ProcessesMemory_Nodes
}

func (processesMemory *ProcessesMemory) GetFilter() yfilter.YFilter { return processesMemory.YFilter }

func (processesMemory *ProcessesMemory) SetFilter(yf yfilter.YFilter) { processesMemory.YFilter = yf }

func (processesMemory *ProcessesMemory) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (processesMemory *ProcessesMemory) GetSegmentPath() string {
    return "Cisco-IOS-XR-procmem-oper:processes-memory"
}

func (processesMemory *ProcessesMemory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &processesMemory.Nodes
    }
    return nil
}

func (processesMemory *ProcessesMemory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &processesMemory.Nodes
    return children
}

func (processesMemory *ProcessesMemory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processesMemory *ProcessesMemory) GetBundleName() string { return "cisco_ios_xr" }

func (processesMemory *ProcessesMemory) GetYangName() string { return "processes-memory" }

func (processesMemory *ProcessesMemory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processesMemory *ProcessesMemory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processesMemory *ProcessesMemory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processesMemory *ProcessesMemory) SetParent(parent types.Entity) { processesMemory.parent = parent }

func (processesMemory *ProcessesMemory) GetParent() types.Entity { return processesMemory.parent }

func (processesMemory *ProcessesMemory) GetParentYangName() string { return "Cisco-IOS-XR-procmem-oper" }

// ProcessesMemory_Nodes
// List of nodes
type ProcessesMemory_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node ID. The type is slice of ProcessesMemory_Nodes_Node.
    Node []ProcessesMemory_Nodes_Node
}

func (nodes *ProcessesMemory_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *ProcessesMemory_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *ProcessesMemory_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *ProcessesMemory_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *ProcessesMemory_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ProcessesMemory_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *ProcessesMemory_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *ProcessesMemory_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *ProcessesMemory_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *ProcessesMemory_Nodes) GetYangName() string { return "nodes" }

func (nodes *ProcessesMemory_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *ProcessesMemory_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *ProcessesMemory_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *ProcessesMemory_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *ProcessesMemory_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *ProcessesMemory_Nodes) GetParentYangName() string { return "processes-memory" }

// ProcessesMemory_Nodes_Node
// Node ID
type ProcessesMemory_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // List of jobs.
    ProcessIds ProcessesMemory_Nodes_Node_ProcessIds
}

func (node *ProcessesMemory_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *ProcessesMemory_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *ProcessesMemory_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "process-ids" { return "ProcessIds" }
    return ""
}

func (node *ProcessesMemory_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *ProcessesMemory_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-ids" {
        return &node.ProcessIds
    }
    return nil
}

func (node *ProcessesMemory_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["process-ids"] = &node.ProcessIds
    return children
}

func (node *ProcessesMemory_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *ProcessesMemory_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *ProcessesMemory_Nodes_Node) GetYangName() string { return "node" }

func (node *ProcessesMemory_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *ProcessesMemory_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *ProcessesMemory_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *ProcessesMemory_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *ProcessesMemory_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *ProcessesMemory_Nodes_Node) GetParentYangName() string { return "nodes" }

// ProcessesMemory_Nodes_Node_ProcessIds
// List of jobs
type ProcessesMemory_Nodes_Node_ProcessIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process Id. The type is slice of
    // ProcessesMemory_Nodes_Node_ProcessIds_ProcessId.
    ProcessId []ProcessesMemory_Nodes_Node_ProcessIds_ProcessId
}

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetFilter() yfilter.YFilter { return processIds.YFilter }

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) SetFilter(yf yfilter.YFilter) { processIds.YFilter = yf }

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetGoName(yname string) string {
    if yname == "process-id" { return "ProcessId" }
    return ""
}

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetSegmentPath() string {
    return "process-ids"
}

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-id" {
        for _, c := range processIds.ProcessId {
            if processIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ProcessesMemory_Nodes_Node_ProcessIds_ProcessId{}
        processIds.ProcessId = append(processIds.ProcessId, child)
        return &processIds.ProcessId[len(processIds.ProcessId)-1]
    }
    return nil
}

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range processIds.ProcessId {
        children[processIds.ProcessId[i].GetSegmentPath()] = &processIds.ProcessId[i]
    }
    return children
}

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetBundleName() string { return "cisco_ios_xr" }

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetYangName() string { return "process-ids" }

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) SetParent(parent types.Entity) { processIds.parent = parent }

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetParent() types.Entity { return processIds.parent }

func (processIds *ProcessesMemory_Nodes_Node_ProcessIds) GetParentYangName() string { return "node" }

// ProcessesMemory_Nodes_Node_ProcessIds_ProcessId
// Process Id
type ProcessesMemory_Nodes_Node_ProcessIds_ProcessId struct {
    parent types.Entity
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

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetFilter() yfilter.YFilter { return processId.YFilter }

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) SetFilter(yf yfilter.YFilter) { processId.YFilter = yf }

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetGoName(yname string) string {
    if yname == "process-id" { return "ProcessId" }
    if yname == "name" { return "Name" }
    if yname == "jid" { return "Jid" }
    if yname == "pid" { return "Pid" }
    if yname == "text-seg-size" { return "TextSegSize" }
    if yname == "data-seg-size" { return "DataSegSize" }
    if yname == "stack-seg-size" { return "StackSegSize" }
    if yname == "malloc-size" { return "MallocSize" }
    if yname == "dyn-limit" { return "DynLimit" }
    if yname == "shared-mem" { return "SharedMem" }
    if yname == "physical-mem" { return "PhysicalMem" }
    return ""
}

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetSegmentPath() string {
    return "process-id" + "[process-id='" + fmt.Sprintf("%v", processId.ProcessId) + "']"
}

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-id"] = processId.ProcessId
    leafs["name"] = processId.Name
    leafs["jid"] = processId.Jid
    leafs["pid"] = processId.Pid
    leafs["text-seg-size"] = processId.TextSegSize
    leafs["data-seg-size"] = processId.DataSegSize
    leafs["stack-seg-size"] = processId.StackSegSize
    leafs["malloc-size"] = processId.MallocSize
    leafs["dyn-limit"] = processId.DynLimit
    leafs["shared-mem"] = processId.SharedMem
    leafs["physical-mem"] = processId.PhysicalMem
    return leafs
}

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetBundleName() string { return "cisco_ios_xr" }

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetYangName() string { return "process-id" }

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) SetParent(parent types.Entity) { processId.parent = parent }

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetParent() types.Entity { return processId.parent }

func (processId *ProcessesMemory_Nodes_Node_ProcessIds_ProcessId) GetParentYangName() string { return "process-ids" }

