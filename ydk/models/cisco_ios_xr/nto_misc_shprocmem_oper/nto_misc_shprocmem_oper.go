// This module contains a collection of YANG definitions
// for Cisco IOS-XR nto-misc-shprocmem package operational data.
// 
// This module contains definitions
// for the following management objects:
//   processes-memory: Process statistics
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package nto_misc_shprocmem_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package nto_misc_shprocmem_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-nto-misc-shprocmem-oper processes-memory}", reflect.TypeOf(ProcessesMemory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-nto-misc-shprocmem-oper:processes-memory", reflect.TypeOf(ProcessesMemory{}))
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
    return "Cisco-IOS-XR-nto-misc-shprocmem-oper:processes-memory"
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

func (processesMemory *ProcessesMemory) GetParentYangName() string { return "Cisco-IOS-XR-nto-misc-shprocmem-oper" }

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
    JobIds ProcessesMemory_Nodes_Node_JobIds
}

func (node *ProcessesMemory_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *ProcessesMemory_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *ProcessesMemory_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "job-ids" { return "JobIds" }
    return ""
}

func (node *ProcessesMemory_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *ProcessesMemory_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "job-ids" {
        return &node.JobIds
    }
    return nil
}

func (node *ProcessesMemory_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["job-ids"] = &node.JobIds
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

// ProcessesMemory_Nodes_Node_JobIds
// List of jobs
type ProcessesMemory_Nodes_Node_JobIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Job Id. The type is slice of ProcessesMemory_Nodes_Node_JobIds_JobId.
    JobId []ProcessesMemory_Nodes_Node_JobIds_JobId
}

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetFilter() yfilter.YFilter { return jobIds.YFilter }

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) SetFilter(yf yfilter.YFilter) { jobIds.YFilter = yf }

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetGoName(yname string) string {
    if yname == "job-id" { return "JobId" }
    return ""
}

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetSegmentPath() string {
    return "job-ids"
}

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "job-id" {
        for _, c := range jobIds.JobId {
            if jobIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ProcessesMemory_Nodes_Node_JobIds_JobId{}
        jobIds.JobId = append(jobIds.JobId, child)
        return &jobIds.JobId[len(jobIds.JobId)-1]
    }
    return nil
}

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range jobIds.JobId {
        children[jobIds.JobId[i].GetSegmentPath()] = &jobIds.JobId[i]
    }
    return children
}

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetBundleName() string { return "cisco_ios_xr" }

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetYangName() string { return "job-ids" }

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) SetParent(parent types.Entity) { jobIds.parent = parent }

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetParent() types.Entity { return jobIds.parent }

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetParentYangName() string { return "node" }

// ProcessesMemory_Nodes_Node_JobIds_JobId
// Job Id
type ProcessesMemory_Nodes_Node_JobIds_JobId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Job Id. The type is interface{} with range:
    // -2147483648..2147483647.
    JobId interface{}

    // Process name. The type is string.
    Name interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Text Segment Size. The type is interface{} with range: 0..4294967295.
    TextSegSize interface{}

    // Data Segment Size. The type is interface{} with range: 0..4294967295.
    DataSegSize interface{}

    // Stack Segment Size. The type is interface{} with range: 0..4294967295.
    StackSegSize interface{}

    // Malloced Memory Size. The type is interface{} with range: 0..4294967295.
    MallocSize interface{}
}

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetFilter() yfilter.YFilter { return jobId.YFilter }

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) SetFilter(yf yfilter.YFilter) { jobId.YFilter = yf }

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetGoName(yname string) string {
    if yname == "job-id" { return "JobId" }
    if yname == "name" { return "Name" }
    if yname == "jid" { return "Jid" }
    if yname == "text-seg-size" { return "TextSegSize" }
    if yname == "data-seg-size" { return "DataSegSize" }
    if yname == "stack-seg-size" { return "StackSegSize" }
    if yname == "malloc-size" { return "MallocSize" }
    return ""
}

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetSegmentPath() string {
    return "job-id" + "[job-id='" + fmt.Sprintf("%v", jobId.JobId) + "']"
}

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["job-id"] = jobId.JobId
    leafs["name"] = jobId.Name
    leafs["jid"] = jobId.Jid
    leafs["text-seg-size"] = jobId.TextSegSize
    leafs["data-seg-size"] = jobId.DataSegSize
    leafs["stack-seg-size"] = jobId.StackSegSize
    leafs["malloc-size"] = jobId.MallocSize
    return leafs
}

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetBundleName() string { return "cisco_ios_xr" }

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetYangName() string { return "job-id" }

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) SetParent(parent types.Entity) { jobId.parent = parent }

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetParent() types.Entity { return jobId.parent }

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetParentYangName() string { return "job-ids" }

