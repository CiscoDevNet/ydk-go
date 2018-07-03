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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes.
    Nodes ProcessesMemory_Nodes
}

func (processesMemory *ProcessesMemory) GetEntityData() *types.CommonEntityData {
    processesMemory.EntityData.YFilter = processesMemory.YFilter
    processesMemory.EntityData.YangName = "processes-memory"
    processesMemory.EntityData.BundleName = "cisco_ios_xr"
    processesMemory.EntityData.ParentYangName = "Cisco-IOS-XR-nto-misc-shprocmem-oper"
    processesMemory.EntityData.SegmentPath = "Cisco-IOS-XR-nto-misc-shprocmem-oper:processes-memory"
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

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // List of jobs.
    JobIds ProcessesMemory_Nodes_Node_JobIds
}

func (node *ProcessesMemory_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("job-ids", types.YChild{"JobIds", &node.JobIds})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// ProcessesMemory_Nodes_Node_JobIds
// List of jobs
type ProcessesMemory_Nodes_Node_JobIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Job Id. The type is slice of ProcessesMemory_Nodes_Node_JobIds_JobId.
    JobId []*ProcessesMemory_Nodes_Node_JobIds_JobId
}

func (jobIds *ProcessesMemory_Nodes_Node_JobIds) GetEntityData() *types.CommonEntityData {
    jobIds.EntityData.YFilter = jobIds.YFilter
    jobIds.EntityData.YangName = "job-ids"
    jobIds.EntityData.BundleName = "cisco_ios_xr"
    jobIds.EntityData.ParentYangName = "node"
    jobIds.EntityData.SegmentPath = "job-ids"
    jobIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    jobIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    jobIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    jobIds.EntityData.Children = types.NewOrderedMap()
    jobIds.EntityData.Children.Append("job-id", types.YChild{"JobId", nil})
    for i := range jobIds.JobId {
        jobIds.EntityData.Children.Append(types.GetSegmentPath(jobIds.JobId[i]), types.YChild{"JobId", jobIds.JobId[i]})
    }
    jobIds.EntityData.Leafs = types.NewOrderedMap()

    jobIds.EntityData.YListKeys = []string {}

    return &(jobIds.EntityData)
}

// ProcessesMemory_Nodes_Node_JobIds_JobId
// Job Id
type ProcessesMemory_Nodes_Node_JobIds_JobId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Job Id. The type is interface{} with range:
    // 0..4294967295.
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

func (jobId *ProcessesMemory_Nodes_Node_JobIds_JobId) GetEntityData() *types.CommonEntityData {
    jobId.EntityData.YFilter = jobId.YFilter
    jobId.EntityData.YangName = "job-id"
    jobId.EntityData.BundleName = "cisco_ios_xr"
    jobId.EntityData.ParentYangName = "job-ids"
    jobId.EntityData.SegmentPath = "job-id" + types.AddKeyToken(jobId.JobId, "job-id")
    jobId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    jobId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    jobId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    jobId.EntityData.Children = types.NewOrderedMap()
    jobId.EntityData.Leafs = types.NewOrderedMap()
    jobId.EntityData.Leafs.Append("job-id", types.YLeaf{"JobId", jobId.JobId})
    jobId.EntityData.Leafs.Append("name", types.YLeaf{"Name", jobId.Name})
    jobId.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", jobId.Jid})
    jobId.EntityData.Leafs.Append("text-seg-size", types.YLeaf{"TextSegSize", jobId.TextSegSize})
    jobId.EntityData.Leafs.Append("data-seg-size", types.YLeaf{"DataSegSize", jobId.DataSegSize})
    jobId.EntityData.Leafs.Append("stack-seg-size", types.YLeaf{"StackSegSize", jobId.StackSegSize})
    jobId.EntityData.Leafs.Append("malloc-size", types.YLeaf{"MallocSize", jobId.MallocSize})

    jobId.EntityData.YListKeys = []string {"JobId"}

    return &(jobId.EntityData)
}

