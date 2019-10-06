// This module contains a collection of YANG definitions
// for Cisco IOS-XR sysmgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   process-mandatory: Process mandatory configuration
//   process-single-crash: process single crash
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysmgr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysmgr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-sysmgr-cfg process-mandatory}", reflect.TypeOf(ProcessMandatory{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysmgr-cfg:process-mandatory", reflect.TypeOf(ProcessMandatory{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-sysmgr-cfg process-single-crash}", reflect.TypeOf(ProcessSingleCrash{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysmgr-cfg:process-single-crash", reflect.TypeOf(ProcessSingleCrash{}))
}

// ProcessMandatory
// Process mandatory configuration
type ProcessMandatory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of mandatory nodes.
    Nodes ProcessMandatory_Nodes

    // Mandatory for all nodes.
    All ProcessMandatory_All
}

func (processMandatory *ProcessMandatory) GetEntityData() *types.CommonEntityData {
    processMandatory.EntityData.YFilter = processMandatory.YFilter
    processMandatory.EntityData.YangName = "process-mandatory"
    processMandatory.EntityData.BundleName = "cisco_ios_xr"
    processMandatory.EntityData.ParentYangName = "Cisco-IOS-XR-sysmgr-cfg"
    processMandatory.EntityData.SegmentPath = "Cisco-IOS-XR-sysmgr-cfg:process-mandatory"
    processMandatory.EntityData.AbsolutePath = processMandatory.EntityData.SegmentPath
    processMandatory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processMandatory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processMandatory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processMandatory.EntityData.Children = types.NewOrderedMap()
    processMandatory.EntityData.Children.Append("nodes", types.YChild{"Nodes", &processMandatory.Nodes})
    processMandatory.EntityData.Children.Append("all", types.YChild{"All", &processMandatory.All})
    processMandatory.EntityData.Leafs = types.NewOrderedMap()

    processMandatory.EntityData.YListKeys = []string {}

    return &(processMandatory.EntityData)
}

// ProcessMandatory_Nodes
// Table of mandatory nodes
type ProcessMandatory_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mandatory node. The type is slice of ProcessMandatory_Nodes_Node.
    Node []*ProcessMandatory_Nodes_Node
}

func (nodes *ProcessMandatory_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "process-mandatory"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-cfg:process-mandatory/" + nodes.EntityData.SegmentPath
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

// ProcessMandatory_Nodes_Node
// Mandatory node
type ProcessMandatory_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Table of processes.
    Processes ProcessMandatory_Nodes_Node_Processes
}

func (node *ProcessMandatory_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-cfg:process-mandatory/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("processes", types.YChild{"Processes", &node.Processes})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// ProcessMandatory_Nodes_Node_Processes
// Table of processes
type ProcessMandatory_Nodes_Node_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the executable process. The type is slice of
    // ProcessMandatory_Nodes_Node_Processes_Process.
    Process []*ProcessMandatory_Nodes_Node_Processes_Process
}

func (processes *ProcessMandatory_Nodes_Node_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xr"
    processes.EntityData.ParentYangName = "node"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-cfg:process-mandatory/nodes/node/" + processes.EntityData.SegmentPath
    processes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processes.EntityData.Children = types.NewOrderedMap()
    processes.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range processes.Process {
        processes.EntityData.Children.Append(types.GetSegmentPath(processes.Process[i]), types.YChild{"Process", processes.Process[i]})
    }
    processes.EntityData.Leafs = types.NewOrderedMap()

    processes.EntityData.YListKeys = []string {}

    return &(processes.EntityData)
}

// ProcessMandatory_Nodes_Node_Processes_Process
// Name of the executable process
type ProcessMandatory_Nodes_Node_Processes_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProcessName interface{}
}

func (process *ProcessMandatory_Nodes_Node_Processes_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "processes"
    process.EntityData.SegmentPath = "process" + types.AddKeyToken(process.ProcessName, "process-name")
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-cfg:process-mandatory/nodes/node/processes/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", process.ProcessName})

    process.EntityData.YListKeys = []string {"ProcessName"}

    return &(process.EntityData)
}

// ProcessMandatory_All
// Mandatory for all nodes
type ProcessMandatory_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of processes.
    Processes ProcessMandatory_All_Processes
}

func (all *ProcessMandatory_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "process-mandatory"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-cfg:process-mandatory/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("processes", types.YChild{"Processes", &all.Processes})
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// ProcessMandatory_All_Processes
// Table of processes
type ProcessMandatory_All_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the executable process. The type is slice of
    // ProcessMandatory_All_Processes_Process.
    Process []*ProcessMandatory_All_Processes_Process
}

func (processes *ProcessMandatory_All_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xr"
    processes.EntityData.ParentYangName = "all"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-cfg:process-mandatory/all/" + processes.EntityData.SegmentPath
    processes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processes.EntityData.Children = types.NewOrderedMap()
    processes.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range processes.Process {
        processes.EntityData.Children.Append(types.GetSegmentPath(processes.Process[i]), types.YChild{"Process", processes.Process[i]})
    }
    processes.EntityData.Leafs = types.NewOrderedMap()

    processes.EntityData.YListKeys = []string {}

    return &(processes.EntityData)
}

// ProcessMandatory_All_Processes_Process
// Name of the executable process
type ProcessMandatory_All_Processes_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProcessName interface{}
}

func (process *ProcessMandatory_All_Processes_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "processes"
    process.EntityData.SegmentPath = "process" + types.AddKeyToken(process.ProcessName, "process-name")
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-cfg:process-mandatory/all/processes/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", process.ProcessName})

    process.EntityData.YListKeys = []string {"ProcessName"}

    return &(process.EntityData)
}

// ProcessSingleCrash
// process single crash
// This type is a presence type.
type ProcessSingleCrash struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Number of crashes for a process to trigger reboot. The type is interface{}
    // with range: 1..500. This attribute is mandatory.
    Crashes interface{}

    // Minimum process up time (in seconds) to reset crash count. The type is
    // interface{} with range: 0..4294967295. Units are second. The default value
    // is 0.
    MinimumUpTime interface{}
}

func (processSingleCrash *ProcessSingleCrash) GetEntityData() *types.CommonEntityData {
    processSingleCrash.EntityData.YFilter = processSingleCrash.YFilter
    processSingleCrash.EntityData.YangName = "process-single-crash"
    processSingleCrash.EntityData.BundleName = "cisco_ios_xr"
    processSingleCrash.EntityData.ParentYangName = "Cisco-IOS-XR-sysmgr-cfg"
    processSingleCrash.EntityData.SegmentPath = "Cisco-IOS-XR-sysmgr-cfg:process-single-crash"
    processSingleCrash.EntityData.AbsolutePath = processSingleCrash.EntityData.SegmentPath
    processSingleCrash.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processSingleCrash.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processSingleCrash.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processSingleCrash.EntityData.Children = types.NewOrderedMap()
    processSingleCrash.EntityData.Leafs = types.NewOrderedMap()
    processSingleCrash.EntityData.Leafs.Append("crashes", types.YLeaf{"Crashes", processSingleCrash.Crashes})
    processSingleCrash.EntityData.Leafs.Append("minimum-up-time", types.YLeaf{"MinimumUpTime", processSingleCrash.MinimumUpTime})

    processSingleCrash.EntityData.YListKeys = []string {}

    return &(processSingleCrash.EntityData)
}

