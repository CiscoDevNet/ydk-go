// This module contains a collection of YANG definitions
// for Cisco IOS-XR sysmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   system-process: Process information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysmgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-sysmgr-oper system-process}", reflect.TypeOf(SystemProcess{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysmgr-oper:system-process", reflect.TypeOf(SystemProcess{}))
}

// PlacementState represents Process placement state
type PlacementState string

const (
    // Process runs everywhere (ubiquitous)
    PlacementState_place_null PlacementState = "place-null"

    // Process runs on node chosen by PlaceD
    PlacementState_place_placeable PlacementState = "place-placeable"

    // Process runs on dSDRSC only
    PlacementState_place_dlrsc_tracker PlacementState = "place-dlrsc-tracker"

    // Process runs on RP of each rack
    PlacementState_place_rack_centric PlacementState = "place-rack-centric"

    // Process runs on DSC only
    PlacementState_place_dsc_tracker PlacementState = "place-dsc-tracker"
)

// ProcessState represents Process state
type ProcessState string

const (
    // NONE
    ProcessState_none ProcessState = "none"

    // RUN
    ProcessState_run ProcessState = "run"

    // EXITED
    ProcessState_exited ProcessState = "exited"

    // HOLD
    ProcessState_hold ProcessState = "hold"

    // WAIT
    ProcessState_wait ProcessState = "wait"

    // RESTART
    ProcessState_restart ProcessState = "restart"

    // INITIALIZING
    ProcessState_initializing ProcessState = "initializing"

    // KILLED
    ProcessState_killed ProcessState = "killed"

    // QUEUED
    ProcessState_queued ProcessState = "queued"

    // ERROR
    ProcessState_error_ ProcessState = "error"

    // TUPLESET
    ProcessState_tuple_set ProcessState = "tuple-set"

    // UNKNOWN
    ProcessState_unknown ProcessState = "unknown"
)

// SystemProcess
// Process information
type SystemProcess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes.
    NodeTable SystemProcess_NodeTable
}

func (systemProcess *SystemProcess) GetEntityData() *types.CommonEntityData {
    systemProcess.EntityData.YFilter = systemProcess.YFilter
    systemProcess.EntityData.YangName = "system-process"
    systemProcess.EntityData.BundleName = "cisco_ios_xr"
    systemProcess.EntityData.ParentYangName = "Cisco-IOS-XR-sysmgr-oper"
    systemProcess.EntityData.SegmentPath = "Cisco-IOS-XR-sysmgr-oper:system-process"
    systemProcess.EntityData.AbsolutePath = systemProcess.EntityData.SegmentPath
    systemProcess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemProcess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemProcess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemProcess.EntityData.Children = types.NewOrderedMap()
    systemProcess.EntityData.Children.Append("node-table", types.YChild{"NodeTable", &systemProcess.NodeTable})
    systemProcess.EntityData.Leafs = types.NewOrderedMap()

    systemProcess.EntityData.YListKeys = []string {}

    return &(systemProcess.EntityData)
}

// SystemProcess_NodeTable
// List of nodes
type SystemProcess_NodeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process information per node. The type is slice of
    // SystemProcess_NodeTable_Node.
    Node []*SystemProcess_NodeTable_Node
}

func (nodeTable *SystemProcess_NodeTable) GetEntityData() *types.CommonEntityData {
    nodeTable.EntityData.YFilter = nodeTable.YFilter
    nodeTable.EntityData.YangName = "node-table"
    nodeTable.EntityData.BundleName = "cisco_ios_xr"
    nodeTable.EntityData.ParentYangName = "system-process"
    nodeTable.EntityData.SegmentPath = "node-table"
    nodeTable.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/" + nodeTable.EntityData.SegmentPath
    nodeTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeTable.EntityData.Children = types.NewOrderedMap()
    nodeTable.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodeTable.Node {
        nodeTable.EntityData.Children.Append(types.GetSegmentPath(nodeTable.Node[i]), types.YChild{"Node", nodeTable.Node[i]})
    }
    nodeTable.EntityData.Leafs = types.NewOrderedMap()

    nodeTable.EntityData.YListKeys = []string {}

    return &(nodeTable.EntityData)
}

// SystemProcess_NodeTable_Node
// Process information per node
type SystemProcess_NodeTable_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Process <WORD> information.
    Name SystemProcess_NodeTable_Node_Name

    // Process job id information.
    Jids SystemProcess_NodeTable_Node_Jids

    // Process Dynamic information.
    Dynamic SystemProcess_NodeTable_Node_Dynamic

    // Process Boot Stalled information.
    BootStalled SystemProcess_NodeTable_Node_BootStalled

    // Process all information.
    Processes SystemProcess_NodeTable_Node_Processes

    // Process Startup information.
    Startup SystemProcess_NodeTable_Node_Startup

    // Mandatory Process information.
    Mandatory SystemProcess_NodeTable_Node_Mandatory

    // Process Abort information.
    Abort SystemProcess_NodeTable_Node_Abort

    // Process Failover information.
    Failover SystemProcess_NodeTable_Node_Failover

    // Process Boot information.
    Boot SystemProcess_NodeTable_Node_Boot

    // Process Log information.
    Logs SystemProcess_NodeTable_Node_Logs

    // Process Searchpath information.
    Searchpath SystemProcess_NodeTable_Node_Searchpath
}

func (node *SystemProcess_NodeTable_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "node-table"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("name", types.YChild{"Name", &node.Name})
    node.EntityData.Children.Append("jids", types.YChild{"Jids", &node.Jids})
    node.EntityData.Children.Append("dynamic", types.YChild{"Dynamic", &node.Dynamic})
    node.EntityData.Children.Append("boot-stalled", types.YChild{"BootStalled", &node.BootStalled})
    node.EntityData.Children.Append("processes", types.YChild{"Processes", &node.Processes})
    node.EntityData.Children.Append("startup", types.YChild{"Startup", &node.Startup})
    node.EntityData.Children.Append("mandatory", types.YChild{"Mandatory", &node.Mandatory})
    node.EntityData.Children.Append("abort", types.YChild{"Abort", &node.Abort})
    node.EntityData.Children.Append("failover", types.YChild{"Failover", &node.Failover})
    node.EntityData.Children.Append("boot", types.YChild{"Boot", &node.Boot})
    node.EntityData.Children.Append("logs", types.YChild{"Logs", &node.Logs})
    node.EntityData.Children.Append("searchpath", types.YChild{"Searchpath", &node.Searchpath})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// SystemProcess_NodeTable_Node_Name
// Process <WORD> information
type SystemProcess_NodeTable_Node_Name struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process <WORD> information.
    ProcessNameRunInfos SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos

    // Process <WORD> information.
    ProcessNameInfos SystemProcess_NodeTable_Node_Name_ProcessNameInfos

    // Process <WORD> information.
    ProcessNameRunDetails SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails

    // Process <WORD> information.
    ProcessNameRunverboses SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses

    // Process <WORD> information.
    ProcessNameDetails SystemProcess_NodeTable_Node_Name_ProcessNameDetails

    // Process <WORD> information.
    ProcessNameVerboses SystemProcess_NodeTable_Node_Name_ProcessNameVerboses
}

func (name *SystemProcess_NodeTable_Node_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "node"
    name.EntityData.SegmentPath = "name"
    name.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + name.EntityData.SegmentPath
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = types.NewOrderedMap()
    name.EntityData.Children.Append("process-name-run-infos", types.YChild{"ProcessNameRunInfos", &name.ProcessNameRunInfos})
    name.EntityData.Children.Append("process-name-infos", types.YChild{"ProcessNameInfos", &name.ProcessNameInfos})
    name.EntityData.Children.Append("process-name-run-details", types.YChild{"ProcessNameRunDetails", &name.ProcessNameRunDetails})
    name.EntityData.Children.Append("process-name-runverboses", types.YChild{"ProcessNameRunverboses", &name.ProcessNameRunverboses})
    name.EntityData.Children.Append("process-name-details", types.YChild{"ProcessNameDetails", &name.ProcessNameDetails})
    name.EntityData.Children.Append("process-name-verboses", types.YChild{"ProcessNameVerboses", &name.ProcessNameVerboses})
    name.EntityData.Leafs = types.NewOrderedMap()

    name.EntityData.YListKeys = []string {}

    return &(name.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos
// Process <WORD> information
type SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process <WORD> run information. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo.
    ProcessNameRunInfo []*SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo
}

func (processNameRunInfos *SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos) GetEntityData() *types.CommonEntityData {
    processNameRunInfos.EntityData.YFilter = processNameRunInfos.YFilter
    processNameRunInfos.EntityData.YangName = "process-name-run-infos"
    processNameRunInfos.EntityData.BundleName = "cisco_ios_xr"
    processNameRunInfos.EntityData.ParentYangName = "name"
    processNameRunInfos.EntityData.SegmentPath = "process-name-run-infos"
    processNameRunInfos.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/" + processNameRunInfos.EntityData.SegmentPath
    processNameRunInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameRunInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameRunInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameRunInfos.EntityData.Children = types.NewOrderedMap()
    processNameRunInfos.EntityData.Children.Append("process-name-run-info", types.YChild{"ProcessNameRunInfo", nil})
    for i := range processNameRunInfos.ProcessNameRunInfo {
        processNameRunInfos.EntityData.Children.Append(types.GetSegmentPath(processNameRunInfos.ProcessNameRunInfo[i]), types.YChild{"ProcessNameRunInfo", processNameRunInfos.ProcessNameRunInfo[i]})
    }
    processNameRunInfos.EntityData.Leafs = types.NewOrderedMap()

    processNameRunInfos.EntityData.YListKeys = []string {}

    return &(processNameRunInfos.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo
// Process <WORD> run information
type SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcName interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    JobIdXr interface{}

    // PID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Process name. The type is string.
    ProcessName interface{}

    // Executable name or path. The type is string.
    Executable interface{}

    // Active Path. The type is string.
    ActivePath interface{}

    // Instance ID. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Args. The type is string.
    Args interface{}

    // Version ID. The type is string.
    VersionId interface{}

    // Respawn on/off. The type is string.
    Respawn interface{}

    // Respawn Count. The type is interface{} with range: -2147483648..2147483647.
    RespawnCount interface{}

    // Last Started. The type is string.
    LastStarted interface{}

    // Process State. The type is string.
    ProcessState interface{}

    // Last Exit status. The type is interface{} with range:
    // -2147483648..2147483647.
    LastExitStatus interface{}

    // Last Exit due to. The type is string.
    LastExitReason interface{}

    // Package State. The type is string.
    PackageState interface{}

    // Started on Config. The type is string.
    StartedOnConfig interface{}

    // Feature Name. The type is string.
    FeatureName interface{}

    // Tag. The type is string.
    Tag interface{}

    // Process Group. The type is string.
    Group interface{}

    // Core. The type is string.
    Core interface{}

    // Max core. The type is interface{} with range: -2147483648..2147483647.
    MaxCore interface{}

    // Level. The type is string.
    Level interface{}

    // Is mandatory?. The type is bool.
    Mandatory interface{}

    // Is admin mode process?. The type is bool.
    MaintModeProc interface{}

    // Placement State. The type is string.
    PlacementState interface{}

    // Startup Path. The type is string.
    StartUpPath interface{}

    // Memory Limit. The type is interface{} with range: 0..4294967295.
    MemoryLimit interface{}

    // Elapsed Ready. The type is string.
    Ready interface{}

    // Elapsed Available. The type is string.
    Available interface{}

    // Proces cpu time.
    ProcCpuTime SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo_ProcCpuTime

    // Registered Items. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo_RegisteredItem.
    RegisteredItem []*SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo_RegisteredItem
}

func (processNameRunInfo *SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo) GetEntityData() *types.CommonEntityData {
    processNameRunInfo.EntityData.YFilter = processNameRunInfo.YFilter
    processNameRunInfo.EntityData.YangName = "process-name-run-info"
    processNameRunInfo.EntityData.BundleName = "cisco_ios_xr"
    processNameRunInfo.EntityData.ParentYangName = "process-name-run-infos"
    processNameRunInfo.EntityData.SegmentPath = "process-name-run-info" + types.AddKeyToken(processNameRunInfo.ProcName, "proc-name")
    processNameRunInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-run-infos/" + processNameRunInfo.EntityData.SegmentPath
    processNameRunInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameRunInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameRunInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameRunInfo.EntityData.Children = types.NewOrderedMap()
    processNameRunInfo.EntityData.Children.Append("proc-cpu-time", types.YChild{"ProcCpuTime", &processNameRunInfo.ProcCpuTime})
    processNameRunInfo.EntityData.Children.Append("registered-item", types.YChild{"RegisteredItem", nil})
    for i := range processNameRunInfo.RegisteredItem {
        types.SetYListKey(processNameRunInfo.RegisteredItem[i], i)
        processNameRunInfo.EntityData.Children.Append(types.GetSegmentPath(processNameRunInfo.RegisteredItem[i]), types.YChild{"RegisteredItem", processNameRunInfo.RegisteredItem[i]})
    }
    processNameRunInfo.EntityData.Leafs = types.NewOrderedMap()
    processNameRunInfo.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", processNameRunInfo.ProcName})
    processNameRunInfo.EntityData.Leafs.Append("job-id-xr", types.YLeaf{"JobIdXr", processNameRunInfo.JobIdXr})
    processNameRunInfo.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", processNameRunInfo.ProcessId})
    processNameRunInfo.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", processNameRunInfo.ProcessName})
    processNameRunInfo.EntityData.Leafs.Append("executable", types.YLeaf{"Executable", processNameRunInfo.Executable})
    processNameRunInfo.EntityData.Leafs.Append("active-path", types.YLeaf{"ActivePath", processNameRunInfo.ActivePath})
    processNameRunInfo.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", processNameRunInfo.InstanceId})
    processNameRunInfo.EntityData.Leafs.Append("args", types.YLeaf{"Args", processNameRunInfo.Args})
    processNameRunInfo.EntityData.Leafs.Append("version-id", types.YLeaf{"VersionId", processNameRunInfo.VersionId})
    processNameRunInfo.EntityData.Leafs.Append("respawn", types.YLeaf{"Respawn", processNameRunInfo.Respawn})
    processNameRunInfo.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", processNameRunInfo.RespawnCount})
    processNameRunInfo.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", processNameRunInfo.LastStarted})
    processNameRunInfo.EntityData.Leafs.Append("process-state", types.YLeaf{"ProcessState", processNameRunInfo.ProcessState})
    processNameRunInfo.EntityData.Leafs.Append("last-exit-status", types.YLeaf{"LastExitStatus", processNameRunInfo.LastExitStatus})
    processNameRunInfo.EntityData.Leafs.Append("last-exit-reason", types.YLeaf{"LastExitReason", processNameRunInfo.LastExitReason})
    processNameRunInfo.EntityData.Leafs.Append("package-state", types.YLeaf{"PackageState", processNameRunInfo.PackageState})
    processNameRunInfo.EntityData.Leafs.Append("started-on-config", types.YLeaf{"StartedOnConfig", processNameRunInfo.StartedOnConfig})
    processNameRunInfo.EntityData.Leafs.Append("feature-name", types.YLeaf{"FeatureName", processNameRunInfo.FeatureName})
    processNameRunInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", processNameRunInfo.Tag})
    processNameRunInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", processNameRunInfo.Group})
    processNameRunInfo.EntityData.Leafs.Append("core", types.YLeaf{"Core", processNameRunInfo.Core})
    processNameRunInfo.EntityData.Leafs.Append("max-core", types.YLeaf{"MaxCore", processNameRunInfo.MaxCore})
    processNameRunInfo.EntityData.Leafs.Append("level", types.YLeaf{"Level", processNameRunInfo.Level})
    processNameRunInfo.EntityData.Leafs.Append("mandatory", types.YLeaf{"Mandatory", processNameRunInfo.Mandatory})
    processNameRunInfo.EntityData.Leafs.Append("maint-mode-proc", types.YLeaf{"MaintModeProc", processNameRunInfo.MaintModeProc})
    processNameRunInfo.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", processNameRunInfo.PlacementState})
    processNameRunInfo.EntityData.Leafs.Append("start-up-path", types.YLeaf{"StartUpPath", processNameRunInfo.StartUpPath})
    processNameRunInfo.EntityData.Leafs.Append("memory-limit", types.YLeaf{"MemoryLimit", processNameRunInfo.MemoryLimit})
    processNameRunInfo.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", processNameRunInfo.Ready})
    processNameRunInfo.EntityData.Leafs.Append("available", types.YLeaf{"Available", processNameRunInfo.Available})

    processNameRunInfo.EntityData.YListKeys = []string {"ProcName"}

    return &(processNameRunInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo_ProcCpuTime
// Proces cpu time
type SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo_ProcCpuTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User time. The type is string.
    User interface{}

    // Kernel time. The type is string.
    System interface{}

    // Total time. The type is string.
    Total interface{}
}

func (procCpuTime *SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo_ProcCpuTime) GetEntityData() *types.CommonEntityData {
    procCpuTime.EntityData.YFilter = procCpuTime.YFilter
    procCpuTime.EntityData.YangName = "proc-cpu-time"
    procCpuTime.EntityData.BundleName = "cisco_ios_xr"
    procCpuTime.EntityData.ParentYangName = "process-name-run-info"
    procCpuTime.EntityData.SegmentPath = "proc-cpu-time"
    procCpuTime.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-run-infos/process-name-run-info/" + procCpuTime.EntityData.SegmentPath
    procCpuTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procCpuTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procCpuTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procCpuTime.EntityData.Children = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs.Append("user", types.YLeaf{"User", procCpuTime.User})
    procCpuTime.EntityData.Leafs.Append("system", types.YLeaf{"System", procCpuTime.System})
    procCpuTime.EntityData.Leafs.Append("total", types.YLeaf{"Total", procCpuTime.Total})

    procCpuTime.EntityData.YListKeys = []string {}

    return &(procCpuTime.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo_RegisteredItem
// Registered Items
type SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo_RegisteredItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (registeredItem *SystemProcess_NodeTable_Node_Name_ProcessNameRunInfos_ProcessNameRunInfo_RegisteredItem) GetEntityData() *types.CommonEntityData {
    registeredItem.EntityData.YFilter = registeredItem.YFilter
    registeredItem.EntityData.YangName = "registered-item"
    registeredItem.EntityData.BundleName = "cisco_ios_xr"
    registeredItem.EntityData.ParentYangName = "process-name-run-info"
    registeredItem.EntityData.SegmentPath = "registered-item" + types.AddNoKeyToken(registeredItem)
    registeredItem.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-run-infos/process-name-run-info/" + registeredItem.EntityData.SegmentPath
    registeredItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registeredItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registeredItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registeredItem.EntityData.Children = types.NewOrderedMap()
    registeredItem.EntityData.Leafs = types.NewOrderedMap()
    registeredItem.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", registeredItem.Tuple})

    registeredItem.EntityData.YListKeys = []string {}

    return &(registeredItem.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameInfos
// Process <WORD> information
type SystemProcess_NodeTable_Node_Name_ProcessNameInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process <WORD> information. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo.
    ProcessNameInfo []*SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo
}

func (processNameInfos *SystemProcess_NodeTable_Node_Name_ProcessNameInfos) GetEntityData() *types.CommonEntityData {
    processNameInfos.EntityData.YFilter = processNameInfos.YFilter
    processNameInfos.EntityData.YangName = "process-name-infos"
    processNameInfos.EntityData.BundleName = "cisco_ios_xr"
    processNameInfos.EntityData.ParentYangName = "name"
    processNameInfos.EntityData.SegmentPath = "process-name-infos"
    processNameInfos.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/" + processNameInfos.EntityData.SegmentPath
    processNameInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameInfos.EntityData.Children = types.NewOrderedMap()
    processNameInfos.EntityData.Children.Append("process-name-info", types.YChild{"ProcessNameInfo", nil})
    for i := range processNameInfos.ProcessNameInfo {
        processNameInfos.EntityData.Children.Append(types.GetSegmentPath(processNameInfos.ProcessNameInfo[i]), types.YChild{"ProcessNameInfo", processNameInfos.ProcessNameInfo[i]})
    }
    processNameInfos.EntityData.Leafs = types.NewOrderedMap()

    processNameInfos.EntityData.YListKeys = []string {}

    return &(processNameInfos.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo
// Process <WORD> information
type SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcName interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    JobIdXr interface{}

    // PID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Process name. The type is string.
    ProcessName interface{}

    // Executable name or path. The type is string.
    Executable interface{}

    // Active Path. The type is string.
    ActivePath interface{}

    // Instance ID. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Args. The type is string.
    Args interface{}

    // Version ID. The type is string.
    VersionId interface{}

    // Respawn on/off. The type is string.
    Respawn interface{}

    // Respawn Count. The type is interface{} with range: -2147483648..2147483647.
    RespawnCount interface{}

    // Last Started. The type is string.
    LastStarted interface{}

    // Process State. The type is string.
    ProcessState interface{}

    // Last Exit status. The type is interface{} with range:
    // -2147483648..2147483647.
    LastExitStatus interface{}

    // Last Exit due to. The type is string.
    LastExitReason interface{}

    // Package State. The type is string.
    PackageState interface{}

    // Started on Config. The type is string.
    StartedOnConfig interface{}

    // Feature Name. The type is string.
    FeatureName interface{}

    // Tag. The type is string.
    Tag interface{}

    // Process Group. The type is string.
    Group interface{}

    // Core. The type is string.
    Core interface{}

    // Max core. The type is interface{} with range: -2147483648..2147483647.
    MaxCore interface{}

    // Level. The type is string.
    Level interface{}

    // Is mandatory?. The type is bool.
    Mandatory interface{}

    // Is admin mode process?. The type is bool.
    MaintModeProc interface{}

    // Placement State. The type is string.
    PlacementState interface{}

    // Startup Path. The type is string.
    StartUpPath interface{}

    // Memory Limit. The type is interface{} with range: 0..4294967295.
    MemoryLimit interface{}

    // Elapsed Ready. The type is string.
    Ready interface{}

    // Elapsed Available. The type is string.
    Available interface{}

    // Proces cpu time.
    ProcCpuTime SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo_ProcCpuTime

    // Registered Items. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo_RegisteredItem.
    RegisteredItem []*SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo_RegisteredItem
}

func (processNameInfo *SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo) GetEntityData() *types.CommonEntityData {
    processNameInfo.EntityData.YFilter = processNameInfo.YFilter
    processNameInfo.EntityData.YangName = "process-name-info"
    processNameInfo.EntityData.BundleName = "cisco_ios_xr"
    processNameInfo.EntityData.ParentYangName = "process-name-infos"
    processNameInfo.EntityData.SegmentPath = "process-name-info" + types.AddKeyToken(processNameInfo.ProcName, "proc-name")
    processNameInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-infos/" + processNameInfo.EntityData.SegmentPath
    processNameInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameInfo.EntityData.Children = types.NewOrderedMap()
    processNameInfo.EntityData.Children.Append("proc-cpu-time", types.YChild{"ProcCpuTime", &processNameInfo.ProcCpuTime})
    processNameInfo.EntityData.Children.Append("registered-item", types.YChild{"RegisteredItem", nil})
    for i := range processNameInfo.RegisteredItem {
        types.SetYListKey(processNameInfo.RegisteredItem[i], i)
        processNameInfo.EntityData.Children.Append(types.GetSegmentPath(processNameInfo.RegisteredItem[i]), types.YChild{"RegisteredItem", processNameInfo.RegisteredItem[i]})
    }
    processNameInfo.EntityData.Leafs = types.NewOrderedMap()
    processNameInfo.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", processNameInfo.ProcName})
    processNameInfo.EntityData.Leafs.Append("job-id-xr", types.YLeaf{"JobIdXr", processNameInfo.JobIdXr})
    processNameInfo.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", processNameInfo.ProcessId})
    processNameInfo.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", processNameInfo.ProcessName})
    processNameInfo.EntityData.Leafs.Append("executable", types.YLeaf{"Executable", processNameInfo.Executable})
    processNameInfo.EntityData.Leafs.Append("active-path", types.YLeaf{"ActivePath", processNameInfo.ActivePath})
    processNameInfo.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", processNameInfo.InstanceId})
    processNameInfo.EntityData.Leafs.Append("args", types.YLeaf{"Args", processNameInfo.Args})
    processNameInfo.EntityData.Leafs.Append("version-id", types.YLeaf{"VersionId", processNameInfo.VersionId})
    processNameInfo.EntityData.Leafs.Append("respawn", types.YLeaf{"Respawn", processNameInfo.Respawn})
    processNameInfo.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", processNameInfo.RespawnCount})
    processNameInfo.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", processNameInfo.LastStarted})
    processNameInfo.EntityData.Leafs.Append("process-state", types.YLeaf{"ProcessState", processNameInfo.ProcessState})
    processNameInfo.EntityData.Leafs.Append("last-exit-status", types.YLeaf{"LastExitStatus", processNameInfo.LastExitStatus})
    processNameInfo.EntityData.Leafs.Append("last-exit-reason", types.YLeaf{"LastExitReason", processNameInfo.LastExitReason})
    processNameInfo.EntityData.Leafs.Append("package-state", types.YLeaf{"PackageState", processNameInfo.PackageState})
    processNameInfo.EntityData.Leafs.Append("started-on-config", types.YLeaf{"StartedOnConfig", processNameInfo.StartedOnConfig})
    processNameInfo.EntityData.Leafs.Append("feature-name", types.YLeaf{"FeatureName", processNameInfo.FeatureName})
    processNameInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", processNameInfo.Tag})
    processNameInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", processNameInfo.Group})
    processNameInfo.EntityData.Leafs.Append("core", types.YLeaf{"Core", processNameInfo.Core})
    processNameInfo.EntityData.Leafs.Append("max-core", types.YLeaf{"MaxCore", processNameInfo.MaxCore})
    processNameInfo.EntityData.Leafs.Append("level", types.YLeaf{"Level", processNameInfo.Level})
    processNameInfo.EntityData.Leafs.Append("mandatory", types.YLeaf{"Mandatory", processNameInfo.Mandatory})
    processNameInfo.EntityData.Leafs.Append("maint-mode-proc", types.YLeaf{"MaintModeProc", processNameInfo.MaintModeProc})
    processNameInfo.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", processNameInfo.PlacementState})
    processNameInfo.EntityData.Leafs.Append("start-up-path", types.YLeaf{"StartUpPath", processNameInfo.StartUpPath})
    processNameInfo.EntityData.Leafs.Append("memory-limit", types.YLeaf{"MemoryLimit", processNameInfo.MemoryLimit})
    processNameInfo.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", processNameInfo.Ready})
    processNameInfo.EntityData.Leafs.Append("available", types.YLeaf{"Available", processNameInfo.Available})

    processNameInfo.EntityData.YListKeys = []string {"ProcName"}

    return &(processNameInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo_ProcCpuTime
// Proces cpu time
type SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo_ProcCpuTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User time. The type is string.
    User interface{}

    // Kernel time. The type is string.
    System interface{}

    // Total time. The type is string.
    Total interface{}
}

func (procCpuTime *SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo_ProcCpuTime) GetEntityData() *types.CommonEntityData {
    procCpuTime.EntityData.YFilter = procCpuTime.YFilter
    procCpuTime.EntityData.YangName = "proc-cpu-time"
    procCpuTime.EntityData.BundleName = "cisco_ios_xr"
    procCpuTime.EntityData.ParentYangName = "process-name-info"
    procCpuTime.EntityData.SegmentPath = "proc-cpu-time"
    procCpuTime.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-infos/process-name-info/" + procCpuTime.EntityData.SegmentPath
    procCpuTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procCpuTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procCpuTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procCpuTime.EntityData.Children = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs.Append("user", types.YLeaf{"User", procCpuTime.User})
    procCpuTime.EntityData.Leafs.Append("system", types.YLeaf{"System", procCpuTime.System})
    procCpuTime.EntityData.Leafs.Append("total", types.YLeaf{"Total", procCpuTime.Total})

    procCpuTime.EntityData.YListKeys = []string {}

    return &(procCpuTime.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo_RegisteredItem
// Registered Items
type SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo_RegisteredItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (registeredItem *SystemProcess_NodeTable_Node_Name_ProcessNameInfos_ProcessNameInfo_RegisteredItem) GetEntityData() *types.CommonEntityData {
    registeredItem.EntityData.YFilter = registeredItem.YFilter
    registeredItem.EntityData.YangName = "registered-item"
    registeredItem.EntityData.BundleName = "cisco_ios_xr"
    registeredItem.EntityData.ParentYangName = "process-name-info"
    registeredItem.EntityData.SegmentPath = "registered-item" + types.AddNoKeyToken(registeredItem)
    registeredItem.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-infos/process-name-info/" + registeredItem.EntityData.SegmentPath
    registeredItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registeredItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registeredItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registeredItem.EntityData.Children = types.NewOrderedMap()
    registeredItem.EntityData.Leafs = types.NewOrderedMap()
    registeredItem.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", registeredItem.Tuple})

    registeredItem.EntityData.YListKeys = []string {}

    return &(registeredItem.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails
// Process <WORD> information
type SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process <WORD> run detail information. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail.
    ProcessNameRunDetail []*SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail
}

func (processNameRunDetails *SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails) GetEntityData() *types.CommonEntityData {
    processNameRunDetails.EntityData.YFilter = processNameRunDetails.YFilter
    processNameRunDetails.EntityData.YangName = "process-name-run-details"
    processNameRunDetails.EntityData.BundleName = "cisco_ios_xr"
    processNameRunDetails.EntityData.ParentYangName = "name"
    processNameRunDetails.EntityData.SegmentPath = "process-name-run-details"
    processNameRunDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/" + processNameRunDetails.EntityData.SegmentPath
    processNameRunDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameRunDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameRunDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameRunDetails.EntityData.Children = types.NewOrderedMap()
    processNameRunDetails.EntityData.Children.Append("process-name-run-detail", types.YChild{"ProcessNameRunDetail", nil})
    for i := range processNameRunDetails.ProcessNameRunDetail {
        processNameRunDetails.EntityData.Children.Append(types.GetSegmentPath(processNameRunDetails.ProcessNameRunDetail[i]), types.YChild{"ProcessNameRunDetail", processNameRunDetails.ProcessNameRunDetail[i]})
    }
    processNameRunDetails.EntityData.Leafs = types.NewOrderedMap()

    processNameRunDetails.EntityData.YListKeys = []string {}

    return &(processNameRunDetails.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail
// Process <WORD> run detail information
type SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcName interface{}

    // Process Basic Info.
    BasicInfo SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo

    // Process Detail Info.
    DetailInfo SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_DetailInfo
}

func (processNameRunDetail *SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail) GetEntityData() *types.CommonEntityData {
    processNameRunDetail.EntityData.YFilter = processNameRunDetail.YFilter
    processNameRunDetail.EntityData.YangName = "process-name-run-detail"
    processNameRunDetail.EntityData.BundleName = "cisco_ios_xr"
    processNameRunDetail.EntityData.ParentYangName = "process-name-run-details"
    processNameRunDetail.EntityData.SegmentPath = "process-name-run-detail" + types.AddKeyToken(processNameRunDetail.ProcName, "proc-name")
    processNameRunDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-run-details/" + processNameRunDetail.EntityData.SegmentPath
    processNameRunDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameRunDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameRunDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameRunDetail.EntityData.Children = types.NewOrderedMap()
    processNameRunDetail.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &processNameRunDetail.BasicInfo})
    processNameRunDetail.EntityData.Children.Append("detail-info", types.YChild{"DetailInfo", &processNameRunDetail.DetailInfo})
    processNameRunDetail.EntityData.Leafs = types.NewOrderedMap()
    processNameRunDetail.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", processNameRunDetail.ProcName})

    processNameRunDetail.EntityData.YListKeys = []string {"ProcName"}

    return &(processNameRunDetail.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo
// Process Basic Info
type SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Job ID. The type is interface{} with range: 0..4294967295.
    JobIdXr interface{}

    // PID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Process name. The type is string.
    ProcessName interface{}

    // Executable name or path. The type is string.
    Executable interface{}

    // Active Path. The type is string.
    ActivePath interface{}

    // Instance ID. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Args. The type is string.
    Args interface{}

    // Version ID. The type is string.
    VersionId interface{}

    // Respawn on/off. The type is string.
    Respawn interface{}

    // Respawn Count. The type is interface{} with range: -2147483648..2147483647.
    RespawnCount interface{}

    // Last Started. The type is string.
    LastStarted interface{}

    // Process State. The type is string.
    ProcessState interface{}

    // Last Exit status. The type is interface{} with range:
    // -2147483648..2147483647.
    LastExitStatus interface{}

    // Last Exit due to. The type is string.
    LastExitReason interface{}

    // Package State. The type is string.
    PackageState interface{}

    // Started on Config. The type is string.
    StartedOnConfig interface{}

    // Feature Name. The type is string.
    FeatureName interface{}

    // Tag. The type is string.
    Tag interface{}

    // Process Group. The type is string.
    Group interface{}

    // Core. The type is string.
    Core interface{}

    // Max core. The type is interface{} with range: -2147483648..2147483647.
    MaxCore interface{}

    // Level. The type is string.
    Level interface{}

    // Is mandatory?. The type is bool.
    Mandatory interface{}

    // Is admin mode process?. The type is bool.
    MaintModeProc interface{}

    // Placement State. The type is string.
    PlacementState interface{}

    // Startup Path. The type is string.
    StartUpPath interface{}

    // Memory Limit. The type is interface{} with range: 0..4294967295.
    MemoryLimit interface{}

    // Elapsed Ready. The type is string.
    Ready interface{}

    // Elapsed Available. The type is string.
    Available interface{}

    // Proces cpu time.
    ProcCpuTime SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo_ProcCpuTime

    // Registered Items. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo_RegisteredItem.
    RegisteredItem []*SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo_RegisteredItem
}

func (basicInfo *SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "process-name-run-detail"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-run-details/process-name-run-detail/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Children.Append("proc-cpu-time", types.YChild{"ProcCpuTime", &basicInfo.ProcCpuTime})
    basicInfo.EntityData.Children.Append("registered-item", types.YChild{"RegisteredItem", nil})
    for i := range basicInfo.RegisteredItem {
        types.SetYListKey(basicInfo.RegisteredItem[i], i)
        basicInfo.EntityData.Children.Append(types.GetSegmentPath(basicInfo.RegisteredItem[i]), types.YChild{"RegisteredItem", basicInfo.RegisteredItem[i]})
    }
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("job-id-xr", types.YLeaf{"JobIdXr", basicInfo.JobIdXr})
    basicInfo.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", basicInfo.ProcessId})
    basicInfo.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", basicInfo.ProcessName})
    basicInfo.EntityData.Leafs.Append("executable", types.YLeaf{"Executable", basicInfo.Executable})
    basicInfo.EntityData.Leafs.Append("active-path", types.YLeaf{"ActivePath", basicInfo.ActivePath})
    basicInfo.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", basicInfo.InstanceId})
    basicInfo.EntityData.Leafs.Append("args", types.YLeaf{"Args", basicInfo.Args})
    basicInfo.EntityData.Leafs.Append("version-id", types.YLeaf{"VersionId", basicInfo.VersionId})
    basicInfo.EntityData.Leafs.Append("respawn", types.YLeaf{"Respawn", basicInfo.Respawn})
    basicInfo.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", basicInfo.RespawnCount})
    basicInfo.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", basicInfo.LastStarted})
    basicInfo.EntityData.Leafs.Append("process-state", types.YLeaf{"ProcessState", basicInfo.ProcessState})
    basicInfo.EntityData.Leafs.Append("last-exit-status", types.YLeaf{"LastExitStatus", basicInfo.LastExitStatus})
    basicInfo.EntityData.Leafs.Append("last-exit-reason", types.YLeaf{"LastExitReason", basicInfo.LastExitReason})
    basicInfo.EntityData.Leafs.Append("package-state", types.YLeaf{"PackageState", basicInfo.PackageState})
    basicInfo.EntityData.Leafs.Append("started-on-config", types.YLeaf{"StartedOnConfig", basicInfo.StartedOnConfig})
    basicInfo.EntityData.Leafs.Append("feature-name", types.YLeaf{"FeatureName", basicInfo.FeatureName})
    basicInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", basicInfo.Tag})
    basicInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", basicInfo.Group})
    basicInfo.EntityData.Leafs.Append("core", types.YLeaf{"Core", basicInfo.Core})
    basicInfo.EntityData.Leafs.Append("max-core", types.YLeaf{"MaxCore", basicInfo.MaxCore})
    basicInfo.EntityData.Leafs.Append("level", types.YLeaf{"Level", basicInfo.Level})
    basicInfo.EntityData.Leafs.Append("mandatory", types.YLeaf{"Mandatory", basicInfo.Mandatory})
    basicInfo.EntityData.Leafs.Append("maint-mode-proc", types.YLeaf{"MaintModeProc", basicInfo.MaintModeProc})
    basicInfo.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", basicInfo.PlacementState})
    basicInfo.EntityData.Leafs.Append("start-up-path", types.YLeaf{"StartUpPath", basicInfo.StartUpPath})
    basicInfo.EntityData.Leafs.Append("memory-limit", types.YLeaf{"MemoryLimit", basicInfo.MemoryLimit})
    basicInfo.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", basicInfo.Ready})
    basicInfo.EntityData.Leafs.Append("available", types.YLeaf{"Available", basicInfo.Available})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo_ProcCpuTime
// Proces cpu time
type SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo_ProcCpuTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User time. The type is string.
    User interface{}

    // Kernel time. The type is string.
    System interface{}

    // Total time. The type is string.
    Total interface{}
}

func (procCpuTime *SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo_ProcCpuTime) GetEntityData() *types.CommonEntityData {
    procCpuTime.EntityData.YFilter = procCpuTime.YFilter
    procCpuTime.EntityData.YangName = "proc-cpu-time"
    procCpuTime.EntityData.BundleName = "cisco_ios_xr"
    procCpuTime.EntityData.ParentYangName = "basic-info"
    procCpuTime.EntityData.SegmentPath = "proc-cpu-time"
    procCpuTime.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-run-details/process-name-run-detail/basic-info/" + procCpuTime.EntityData.SegmentPath
    procCpuTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procCpuTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procCpuTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procCpuTime.EntityData.Children = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs.Append("user", types.YLeaf{"User", procCpuTime.User})
    procCpuTime.EntityData.Leafs.Append("system", types.YLeaf{"System", procCpuTime.System})
    procCpuTime.EntityData.Leafs.Append("total", types.YLeaf{"Total", procCpuTime.Total})

    procCpuTime.EntityData.YListKeys = []string {}

    return &(procCpuTime.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo_RegisteredItem
// Registered Items
type SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo_RegisteredItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (registeredItem *SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_BasicInfo_RegisteredItem) GetEntityData() *types.CommonEntityData {
    registeredItem.EntityData.YFilter = registeredItem.YFilter
    registeredItem.EntityData.YangName = "registered-item"
    registeredItem.EntityData.BundleName = "cisco_ios_xr"
    registeredItem.EntityData.ParentYangName = "basic-info"
    registeredItem.EntityData.SegmentPath = "registered-item" + types.AddNoKeyToken(registeredItem)
    registeredItem.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-run-details/process-name-run-detail/basic-info/" + registeredItem.EntityData.SegmentPath
    registeredItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registeredItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registeredItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registeredItem.EntityData.Children = types.NewOrderedMap()
    registeredItem.EntityData.Leafs = types.NewOrderedMap()
    registeredItem.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", registeredItem.Tuple})

    registeredItem.EntityData.YListKeys = []string {}

    return &(registeredItem.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_DetailInfo
// Process Detail Info
type SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_DetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Running path. The type is string.
    RunningPath interface{}

    // Package path. The type is string.
    PackagePath interface{}

    // Job Id Link. The type is interface{} with range: -2147483648..2147483647.
    JobIdLink interface{}

    // Group Jid. The type is string.
    GroupJid interface{}

    // Fail count. The type is interface{} with range: 0..4294967295.
    FailCount interface{}

    // Restart needed. The type is bool.
    RestartNeeded interface{}

    // Init process. The type is bool.
    InitProcess interface{}

    // Last Online. The type is string.
    LastOnline interface{}

    // This PCB. The type is string.
    ThisPcb interface{}

    // Next PCB. The type is string.
    NextPcb interface{}

    // Env variables. The type is string.
    Envs interface{}

    // Wait For /dev/xxx. The type is string.
    WaitFor interface{}

    // Job ID on RP. The type is interface{} with range: -2147483648..2147483647.
    JobIdOnRp interface{}

    // Is standby capable?. The type is bool.
    IsStandbyCapable interface{}

    // Disable kill?. The type is bool.
    DisableKill interface{}

    // Check avail. The type is bool.
    SendAvail interface{}

    // Node Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    NodeEventCliInfo interface{}

    // Node redundancy state. The type is string.
    NodeRedundancyState interface{}

    // Role event cli info. The type is interface{} with range:
    // -2147483648..2147483647.
    RoleEventCliInfo interface{}

    // Proc Role State. The type is string.
    ProcRoleState interface{}

    // Standby Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    StandbyEventCliInfo interface{}

    // Cleanup event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    CleanupEventCliInfo interface{}

    // Band Ready Event CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    BandReadyEventCliInfo interface{}

    // LR Event CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    LrEventCliInfo interface{}

    // Plane Ready Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    PlaneReadyEventCliInfo interface{}

    // MDR is done CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrIsDoneCliInfo interface{}
}

func (detailInfo *SystemProcess_NodeTable_Node_Name_ProcessNameRunDetails_ProcessNameRunDetail_DetailInfo) GetEntityData() *types.CommonEntityData {
    detailInfo.EntityData.YFilter = detailInfo.YFilter
    detailInfo.EntityData.YangName = "detail-info"
    detailInfo.EntityData.BundleName = "cisco_ios_xr"
    detailInfo.EntityData.ParentYangName = "process-name-run-detail"
    detailInfo.EntityData.SegmentPath = "detail-info"
    detailInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-run-details/process-name-run-detail/" + detailInfo.EntityData.SegmentPath
    detailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailInfo.EntityData.Children = types.NewOrderedMap()
    detailInfo.EntityData.Leafs = types.NewOrderedMap()
    detailInfo.EntityData.Leafs.Append("running-path", types.YLeaf{"RunningPath", detailInfo.RunningPath})
    detailInfo.EntityData.Leafs.Append("package-path", types.YLeaf{"PackagePath", detailInfo.PackagePath})
    detailInfo.EntityData.Leafs.Append("job-id-link", types.YLeaf{"JobIdLink", detailInfo.JobIdLink})
    detailInfo.EntityData.Leafs.Append("group-jid", types.YLeaf{"GroupJid", detailInfo.GroupJid})
    detailInfo.EntityData.Leafs.Append("fail-count", types.YLeaf{"FailCount", detailInfo.FailCount})
    detailInfo.EntityData.Leafs.Append("restart-needed", types.YLeaf{"RestartNeeded", detailInfo.RestartNeeded})
    detailInfo.EntityData.Leafs.Append("init-process", types.YLeaf{"InitProcess", detailInfo.InitProcess})
    detailInfo.EntityData.Leafs.Append("last-online", types.YLeaf{"LastOnline", detailInfo.LastOnline})
    detailInfo.EntityData.Leafs.Append("this-pcb", types.YLeaf{"ThisPcb", detailInfo.ThisPcb})
    detailInfo.EntityData.Leafs.Append("next-pcb", types.YLeaf{"NextPcb", detailInfo.NextPcb})
    detailInfo.EntityData.Leafs.Append("envs", types.YLeaf{"Envs", detailInfo.Envs})
    detailInfo.EntityData.Leafs.Append("wait-for", types.YLeaf{"WaitFor", detailInfo.WaitFor})
    detailInfo.EntityData.Leafs.Append("job-id-on-rp", types.YLeaf{"JobIdOnRp", detailInfo.JobIdOnRp})
    detailInfo.EntityData.Leafs.Append("is-standby-capable", types.YLeaf{"IsStandbyCapable", detailInfo.IsStandbyCapable})
    detailInfo.EntityData.Leafs.Append("disable-kill", types.YLeaf{"DisableKill", detailInfo.DisableKill})
    detailInfo.EntityData.Leafs.Append("send-avail", types.YLeaf{"SendAvail", detailInfo.SendAvail})
    detailInfo.EntityData.Leafs.Append("node-event-cli-info", types.YLeaf{"NodeEventCliInfo", detailInfo.NodeEventCliInfo})
    detailInfo.EntityData.Leafs.Append("node-redundancy-state", types.YLeaf{"NodeRedundancyState", detailInfo.NodeRedundancyState})
    detailInfo.EntityData.Leafs.Append("role-event-cli-info", types.YLeaf{"RoleEventCliInfo", detailInfo.RoleEventCliInfo})
    detailInfo.EntityData.Leafs.Append("proc-role-state", types.YLeaf{"ProcRoleState", detailInfo.ProcRoleState})
    detailInfo.EntityData.Leafs.Append("standby-event-cli-info", types.YLeaf{"StandbyEventCliInfo", detailInfo.StandbyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("cleanup-event-cli-info", types.YLeaf{"CleanupEventCliInfo", detailInfo.CleanupEventCliInfo})
    detailInfo.EntityData.Leafs.Append("band-ready-event-cli-info", types.YLeaf{"BandReadyEventCliInfo", detailInfo.BandReadyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("lr-event-cli-info", types.YLeaf{"LrEventCliInfo", detailInfo.LrEventCliInfo})
    detailInfo.EntityData.Leafs.Append("plane-ready-event-cli-info", types.YLeaf{"PlaneReadyEventCliInfo", detailInfo.PlaneReadyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("mdr-is-done-cli-info", types.YLeaf{"MdrIsDoneCliInfo", detailInfo.MdrIsDoneCliInfo})

    detailInfo.EntityData.YListKeys = []string {}

    return &(detailInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses
// Process <WORD> information
type SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process <WORD> run verbose information. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose.
    ProcessNameRunverbose []*SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose
}

func (processNameRunverboses *SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses) GetEntityData() *types.CommonEntityData {
    processNameRunverboses.EntityData.YFilter = processNameRunverboses.YFilter
    processNameRunverboses.EntityData.YangName = "process-name-runverboses"
    processNameRunverboses.EntityData.BundleName = "cisco_ios_xr"
    processNameRunverboses.EntityData.ParentYangName = "name"
    processNameRunverboses.EntityData.SegmentPath = "process-name-runverboses"
    processNameRunverboses.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/" + processNameRunverboses.EntityData.SegmentPath
    processNameRunverboses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameRunverboses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameRunverboses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameRunverboses.EntityData.Children = types.NewOrderedMap()
    processNameRunverboses.EntityData.Children.Append("process-name-runverbose", types.YChild{"ProcessNameRunverbose", nil})
    for i := range processNameRunverboses.ProcessNameRunverbose {
        processNameRunverboses.EntityData.Children.Append(types.GetSegmentPath(processNameRunverboses.ProcessNameRunverbose[i]), types.YChild{"ProcessNameRunverbose", processNameRunverboses.ProcessNameRunverbose[i]})
    }
    processNameRunverboses.EntityData.Leafs = types.NewOrderedMap()

    processNameRunverboses.EntityData.YListKeys = []string {}

    return &(processNameRunverboses.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose
// Process <WORD> run verbose information
type SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcName interface{}

    // Process Basic Info.
    BasicInfo SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo

    // Process Detail Info.
    DetailInfo SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_DetailInfo

    // Process Verbose Info.
    VerboseInfo SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo
}

func (processNameRunverbose *SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose) GetEntityData() *types.CommonEntityData {
    processNameRunverbose.EntityData.YFilter = processNameRunverbose.YFilter
    processNameRunverbose.EntityData.YangName = "process-name-runverbose"
    processNameRunverbose.EntityData.BundleName = "cisco_ios_xr"
    processNameRunverbose.EntityData.ParentYangName = "process-name-runverboses"
    processNameRunverbose.EntityData.SegmentPath = "process-name-runverbose" + types.AddKeyToken(processNameRunverbose.ProcName, "proc-name")
    processNameRunverbose.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-runverboses/" + processNameRunverbose.EntityData.SegmentPath
    processNameRunverbose.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameRunverbose.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameRunverbose.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameRunverbose.EntityData.Children = types.NewOrderedMap()
    processNameRunverbose.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &processNameRunverbose.BasicInfo})
    processNameRunverbose.EntityData.Children.Append("detail-info", types.YChild{"DetailInfo", &processNameRunverbose.DetailInfo})
    processNameRunverbose.EntityData.Children.Append("verbose-info", types.YChild{"VerboseInfo", &processNameRunverbose.VerboseInfo})
    processNameRunverbose.EntityData.Leafs = types.NewOrderedMap()
    processNameRunverbose.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", processNameRunverbose.ProcName})

    processNameRunverbose.EntityData.YListKeys = []string {"ProcName"}

    return &(processNameRunverbose.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo
// Process Basic Info
type SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Job ID. The type is interface{} with range: 0..4294967295.
    JobIdXr interface{}

    // PID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Process name. The type is string.
    ProcessName interface{}

    // Executable name or path. The type is string.
    Executable interface{}

    // Active Path. The type is string.
    ActivePath interface{}

    // Instance ID. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Args. The type is string.
    Args interface{}

    // Version ID. The type is string.
    VersionId interface{}

    // Respawn on/off. The type is string.
    Respawn interface{}

    // Respawn Count. The type is interface{} with range: -2147483648..2147483647.
    RespawnCount interface{}

    // Last Started. The type is string.
    LastStarted interface{}

    // Process State. The type is string.
    ProcessState interface{}

    // Last Exit status. The type is interface{} with range:
    // -2147483648..2147483647.
    LastExitStatus interface{}

    // Last Exit due to. The type is string.
    LastExitReason interface{}

    // Package State. The type is string.
    PackageState interface{}

    // Started on Config. The type is string.
    StartedOnConfig interface{}

    // Feature Name. The type is string.
    FeatureName interface{}

    // Tag. The type is string.
    Tag interface{}

    // Process Group. The type is string.
    Group interface{}

    // Core. The type is string.
    Core interface{}

    // Max core. The type is interface{} with range: -2147483648..2147483647.
    MaxCore interface{}

    // Level. The type is string.
    Level interface{}

    // Is mandatory?. The type is bool.
    Mandatory interface{}

    // Is admin mode process?. The type is bool.
    MaintModeProc interface{}

    // Placement State. The type is string.
    PlacementState interface{}

    // Startup Path. The type is string.
    StartUpPath interface{}

    // Memory Limit. The type is interface{} with range: 0..4294967295.
    MemoryLimit interface{}

    // Elapsed Ready. The type is string.
    Ready interface{}

    // Elapsed Available. The type is string.
    Available interface{}

    // Proces cpu time.
    ProcCpuTime SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo_ProcCpuTime

    // Registered Items. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo_RegisteredItem.
    RegisteredItem []*SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo_RegisteredItem
}

func (basicInfo *SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "process-name-runverbose"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-runverboses/process-name-runverbose/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Children.Append("proc-cpu-time", types.YChild{"ProcCpuTime", &basicInfo.ProcCpuTime})
    basicInfo.EntityData.Children.Append("registered-item", types.YChild{"RegisteredItem", nil})
    for i := range basicInfo.RegisteredItem {
        types.SetYListKey(basicInfo.RegisteredItem[i], i)
        basicInfo.EntityData.Children.Append(types.GetSegmentPath(basicInfo.RegisteredItem[i]), types.YChild{"RegisteredItem", basicInfo.RegisteredItem[i]})
    }
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("job-id-xr", types.YLeaf{"JobIdXr", basicInfo.JobIdXr})
    basicInfo.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", basicInfo.ProcessId})
    basicInfo.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", basicInfo.ProcessName})
    basicInfo.EntityData.Leafs.Append("executable", types.YLeaf{"Executable", basicInfo.Executable})
    basicInfo.EntityData.Leafs.Append("active-path", types.YLeaf{"ActivePath", basicInfo.ActivePath})
    basicInfo.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", basicInfo.InstanceId})
    basicInfo.EntityData.Leafs.Append("args", types.YLeaf{"Args", basicInfo.Args})
    basicInfo.EntityData.Leafs.Append("version-id", types.YLeaf{"VersionId", basicInfo.VersionId})
    basicInfo.EntityData.Leafs.Append("respawn", types.YLeaf{"Respawn", basicInfo.Respawn})
    basicInfo.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", basicInfo.RespawnCount})
    basicInfo.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", basicInfo.LastStarted})
    basicInfo.EntityData.Leafs.Append("process-state", types.YLeaf{"ProcessState", basicInfo.ProcessState})
    basicInfo.EntityData.Leafs.Append("last-exit-status", types.YLeaf{"LastExitStatus", basicInfo.LastExitStatus})
    basicInfo.EntityData.Leafs.Append("last-exit-reason", types.YLeaf{"LastExitReason", basicInfo.LastExitReason})
    basicInfo.EntityData.Leafs.Append("package-state", types.YLeaf{"PackageState", basicInfo.PackageState})
    basicInfo.EntityData.Leafs.Append("started-on-config", types.YLeaf{"StartedOnConfig", basicInfo.StartedOnConfig})
    basicInfo.EntityData.Leafs.Append("feature-name", types.YLeaf{"FeatureName", basicInfo.FeatureName})
    basicInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", basicInfo.Tag})
    basicInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", basicInfo.Group})
    basicInfo.EntityData.Leafs.Append("core", types.YLeaf{"Core", basicInfo.Core})
    basicInfo.EntityData.Leafs.Append("max-core", types.YLeaf{"MaxCore", basicInfo.MaxCore})
    basicInfo.EntityData.Leafs.Append("level", types.YLeaf{"Level", basicInfo.Level})
    basicInfo.EntityData.Leafs.Append("mandatory", types.YLeaf{"Mandatory", basicInfo.Mandatory})
    basicInfo.EntityData.Leafs.Append("maint-mode-proc", types.YLeaf{"MaintModeProc", basicInfo.MaintModeProc})
    basicInfo.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", basicInfo.PlacementState})
    basicInfo.EntityData.Leafs.Append("start-up-path", types.YLeaf{"StartUpPath", basicInfo.StartUpPath})
    basicInfo.EntityData.Leafs.Append("memory-limit", types.YLeaf{"MemoryLimit", basicInfo.MemoryLimit})
    basicInfo.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", basicInfo.Ready})
    basicInfo.EntityData.Leafs.Append("available", types.YLeaf{"Available", basicInfo.Available})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo_ProcCpuTime
// Proces cpu time
type SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo_ProcCpuTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User time. The type is string.
    User interface{}

    // Kernel time. The type is string.
    System interface{}

    // Total time. The type is string.
    Total interface{}
}

func (procCpuTime *SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo_ProcCpuTime) GetEntityData() *types.CommonEntityData {
    procCpuTime.EntityData.YFilter = procCpuTime.YFilter
    procCpuTime.EntityData.YangName = "proc-cpu-time"
    procCpuTime.EntityData.BundleName = "cisco_ios_xr"
    procCpuTime.EntityData.ParentYangName = "basic-info"
    procCpuTime.EntityData.SegmentPath = "proc-cpu-time"
    procCpuTime.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-runverboses/process-name-runverbose/basic-info/" + procCpuTime.EntityData.SegmentPath
    procCpuTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procCpuTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procCpuTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procCpuTime.EntityData.Children = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs.Append("user", types.YLeaf{"User", procCpuTime.User})
    procCpuTime.EntityData.Leafs.Append("system", types.YLeaf{"System", procCpuTime.System})
    procCpuTime.EntityData.Leafs.Append("total", types.YLeaf{"Total", procCpuTime.Total})

    procCpuTime.EntityData.YListKeys = []string {}

    return &(procCpuTime.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo_RegisteredItem
// Registered Items
type SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo_RegisteredItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (registeredItem *SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_BasicInfo_RegisteredItem) GetEntityData() *types.CommonEntityData {
    registeredItem.EntityData.YFilter = registeredItem.YFilter
    registeredItem.EntityData.YangName = "registered-item"
    registeredItem.EntityData.BundleName = "cisco_ios_xr"
    registeredItem.EntityData.ParentYangName = "basic-info"
    registeredItem.EntityData.SegmentPath = "registered-item" + types.AddNoKeyToken(registeredItem)
    registeredItem.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-runverboses/process-name-runverbose/basic-info/" + registeredItem.EntityData.SegmentPath
    registeredItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registeredItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registeredItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registeredItem.EntityData.Children = types.NewOrderedMap()
    registeredItem.EntityData.Leafs = types.NewOrderedMap()
    registeredItem.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", registeredItem.Tuple})

    registeredItem.EntityData.YListKeys = []string {}

    return &(registeredItem.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_DetailInfo
// Process Detail Info
type SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_DetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Running path. The type is string.
    RunningPath interface{}

    // Package path. The type is string.
    PackagePath interface{}

    // Job Id Link. The type is interface{} with range: -2147483648..2147483647.
    JobIdLink interface{}

    // Group Jid. The type is string.
    GroupJid interface{}

    // Fail count. The type is interface{} with range: 0..4294967295.
    FailCount interface{}

    // Restart needed. The type is bool.
    RestartNeeded interface{}

    // Init process. The type is bool.
    InitProcess interface{}

    // Last Online. The type is string.
    LastOnline interface{}

    // This PCB. The type is string.
    ThisPcb interface{}

    // Next PCB. The type is string.
    NextPcb interface{}

    // Env variables. The type is string.
    Envs interface{}

    // Wait For /dev/xxx. The type is string.
    WaitFor interface{}

    // Job ID on RP. The type is interface{} with range: -2147483648..2147483647.
    JobIdOnRp interface{}

    // Is standby capable?. The type is bool.
    IsStandbyCapable interface{}

    // Disable kill?. The type is bool.
    DisableKill interface{}

    // Check avail. The type is bool.
    SendAvail interface{}

    // Node Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    NodeEventCliInfo interface{}

    // Node redundancy state. The type is string.
    NodeRedundancyState interface{}

    // Role event cli info. The type is interface{} with range:
    // -2147483648..2147483647.
    RoleEventCliInfo interface{}

    // Proc Role State. The type is string.
    ProcRoleState interface{}

    // Standby Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    StandbyEventCliInfo interface{}

    // Cleanup event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    CleanupEventCliInfo interface{}

    // Band Ready Event CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    BandReadyEventCliInfo interface{}

    // LR Event CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    LrEventCliInfo interface{}

    // Plane Ready Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    PlaneReadyEventCliInfo interface{}

    // MDR is done CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrIsDoneCliInfo interface{}
}

func (detailInfo *SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_DetailInfo) GetEntityData() *types.CommonEntityData {
    detailInfo.EntityData.YFilter = detailInfo.YFilter
    detailInfo.EntityData.YangName = "detail-info"
    detailInfo.EntityData.BundleName = "cisco_ios_xr"
    detailInfo.EntityData.ParentYangName = "process-name-runverbose"
    detailInfo.EntityData.SegmentPath = "detail-info"
    detailInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-runverboses/process-name-runverbose/" + detailInfo.EntityData.SegmentPath
    detailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailInfo.EntityData.Children = types.NewOrderedMap()
    detailInfo.EntityData.Leafs = types.NewOrderedMap()
    detailInfo.EntityData.Leafs.Append("running-path", types.YLeaf{"RunningPath", detailInfo.RunningPath})
    detailInfo.EntityData.Leafs.Append("package-path", types.YLeaf{"PackagePath", detailInfo.PackagePath})
    detailInfo.EntityData.Leafs.Append("job-id-link", types.YLeaf{"JobIdLink", detailInfo.JobIdLink})
    detailInfo.EntityData.Leafs.Append("group-jid", types.YLeaf{"GroupJid", detailInfo.GroupJid})
    detailInfo.EntityData.Leafs.Append("fail-count", types.YLeaf{"FailCount", detailInfo.FailCount})
    detailInfo.EntityData.Leafs.Append("restart-needed", types.YLeaf{"RestartNeeded", detailInfo.RestartNeeded})
    detailInfo.EntityData.Leafs.Append("init-process", types.YLeaf{"InitProcess", detailInfo.InitProcess})
    detailInfo.EntityData.Leafs.Append("last-online", types.YLeaf{"LastOnline", detailInfo.LastOnline})
    detailInfo.EntityData.Leafs.Append("this-pcb", types.YLeaf{"ThisPcb", detailInfo.ThisPcb})
    detailInfo.EntityData.Leafs.Append("next-pcb", types.YLeaf{"NextPcb", detailInfo.NextPcb})
    detailInfo.EntityData.Leafs.Append("envs", types.YLeaf{"Envs", detailInfo.Envs})
    detailInfo.EntityData.Leafs.Append("wait-for", types.YLeaf{"WaitFor", detailInfo.WaitFor})
    detailInfo.EntityData.Leafs.Append("job-id-on-rp", types.YLeaf{"JobIdOnRp", detailInfo.JobIdOnRp})
    detailInfo.EntityData.Leafs.Append("is-standby-capable", types.YLeaf{"IsStandbyCapable", detailInfo.IsStandbyCapable})
    detailInfo.EntityData.Leafs.Append("disable-kill", types.YLeaf{"DisableKill", detailInfo.DisableKill})
    detailInfo.EntityData.Leafs.Append("send-avail", types.YLeaf{"SendAvail", detailInfo.SendAvail})
    detailInfo.EntityData.Leafs.Append("node-event-cli-info", types.YLeaf{"NodeEventCliInfo", detailInfo.NodeEventCliInfo})
    detailInfo.EntityData.Leafs.Append("node-redundancy-state", types.YLeaf{"NodeRedundancyState", detailInfo.NodeRedundancyState})
    detailInfo.EntityData.Leafs.Append("role-event-cli-info", types.YLeaf{"RoleEventCliInfo", detailInfo.RoleEventCliInfo})
    detailInfo.EntityData.Leafs.Append("proc-role-state", types.YLeaf{"ProcRoleState", detailInfo.ProcRoleState})
    detailInfo.EntityData.Leafs.Append("standby-event-cli-info", types.YLeaf{"StandbyEventCliInfo", detailInfo.StandbyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("cleanup-event-cli-info", types.YLeaf{"CleanupEventCliInfo", detailInfo.CleanupEventCliInfo})
    detailInfo.EntityData.Leafs.Append("band-ready-event-cli-info", types.YLeaf{"BandReadyEventCliInfo", detailInfo.BandReadyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("lr-event-cli-info", types.YLeaf{"LrEventCliInfo", detailInfo.LrEventCliInfo})
    detailInfo.EntityData.Leafs.Append("plane-ready-event-cli-info", types.YLeaf{"PlaneReadyEventCliInfo", detailInfo.PlaneReadyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("mdr-is-done-cli-info", types.YLeaf{"MdrIsDoneCliInfo", detailInfo.MdrIsDoneCliInfo})

    detailInfo.EntityData.YListKeys = []string {}

    return &(detailInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo
// Process Verbose Info
type SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process Group. The type is string.
    ProcessGroup interface{}

    // Is respawn allowed?. The type is interface{} with range:
    // -2147483648..2147483647.
    RespawnAllowed interface{}

    // Wait for exit. The type is interface{} with range: -2147483648..2147483647.
    WaitForExit interface{}

    // Dynamic Tag. The type is interface{} with range: -2147483648..2147483647.
    DynamicTag interface{}

    // Forced stop. The type is interface{} with range: -2147483648..2147483647.
    ForcedStop interface{}

    // Critical Process. The type is interface{} with range:
    // -2147483648..2147483647.
    CriticalProcess interface{}

    // Hold. The type is interface{} with range: -2147483648..2147483647.
    Hold interface{}

    // Transient. The type is interface{} with range: -2147483648..2147483647.
    Transient interface{}

    // Tuple Cfgmgr. The type is interface{} with range: -2147483648..2147483647.
    TupleCfgmgr interface{}

    // Standby capable. The type is interface{} with range:
    // -2147483648..2147483647.
    StandbyCapable interface{}

    // EDM startup. The type is interface{} with range: -2147483648..2147483647.
    EdmStartup interface{}

    // Placement. The type is interface{} with range: -2147483648..2147483647.
    Placement interface{}

    // Skip Kill Notif. The type is interface{} with range:
    // -2147483648..2147483647.
    SkipKillNotif interface{}

    // Init process. The type is interface{} with range: -2147483648..2147483647.
    InitProc interface{}

    // Sysdb Event. The type is interface{} with range: -2147483648..2147483647.
    SysdbEvent interface{}

    // Level Started. The type is interface{} with range: -2147483648..2147483647.
    LevelStarted interface{}

    // Process available. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcAvail interface{}

    // Tuples Scanned. The type is interface{} with range:
    // -2147483648..2147483647.
    TuplesScanned interface{}

    // No checkpoint start. The type is interface{} with range:
    // -2147483648..2147483647.
    NoChkptStart interface{}

    // In Shut Down. The type is interface{} with range: -2147483648..2147483647.
    InShutDown interface{}

    // SM started. The type is interface{} with range: -2147483648..2147483647.
    SmStarted interface{}

    // Ignore on SC. The type is interface{} with range: -2147483648..2147483647.
    IgnoreOnSc interface{}

    // Ignore on EasyBake. The type is interface{} with range:
    // -2147483648..2147483647.
    IgnoreOnEasyBake interface{}

    // Pre init. The type is interface{} with range: -2147483648..2147483647.
    PreInit interface{}

    // EOI received. The type is interface{} with range: -2147483648..2147483647.
    EoiReceived interface{}

    // EOI Timeout. The type is interface{} with range: -2147483648..2147483647.
    EoiTimeout interface{}

    // Avail Timeout. The type is interface{} with range: -2147483648..2147483647.
    AvailTimeout interface{}

    // Reserved Memory. The type is interface{} with range:
    // -2147483648..2147483647.
    ReservedMemory interface{}

    // Allow Warned. The type is interface{} with range: -2147483648..2147483647.
    AllowWarned interface{}

    // Arg Change. The type is interface{} with range: -2147483648..2147483647.
    ArgChange interface{}

    // Restart on tuple. The type is interface{} with range:
    // -2147483648..2147483647.
    RestartOnTuple interface{}

    // Boot Hold. The type is interface{} with range: -2147483648..2147483647.
    BootHold interface{}

    // Reg Id. The type is interface{} with range: -2147483648..2147483647.
    RegId interface{}

    // Memory Limit. The type is interface{} with range: -2147483648..2147483647.
    MemoryLimit interface{}

    // Parent Job ID. The type is interface{} with range: -2147483648..2147483647.
    ParentJobId interface{}

    // Tuple Index. The type is interface{} with range: -2147483648..2147483647.
    TupleIndex interface{}

    // Dump Count. The type is interface{} with range: -2147483648..2147483647.
    DumpCount interface{}

    // Respawn Interval User. The type is interface{} with range:
    // -2147483648..2147483647.
    RespawnIntervalUser interface{}

    // Silent Restart Count. The type is interface{} with range:
    // -2147483648..2147483647.
    SilentRestartCount interface{}

    // Critical Tier. The type is interface{} with range: -2147483648..2147483647.
    CriticalTier interface{}

    // Exit Type. The type is interface{} with range: -2147483648..2147483647.
    ExitType interface{}

    // Init Timeout. The type is interface{} with range: -2147483648..2147483647.
    InitTimeout interface{}

    // Restart by Command. The type is interface{} with range:
    // -2147483648..2147483647.
    RestartByCmd interface{}

    // Boot Pref. The type is interface{} with range: -2147483648..2147483647.
    BootPref interface{}

    // Mdr Mbi proc. The type is interface{} with range: -2147483648..2147483647.
    MdrMbiProc interface{}

    // Mdr Non Mbi Kld. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrNonMbiKld interface{}

    // Mdr Mbi Kld. The type is interface{} with range: -2147483648..2147483647.
    MdrMbiKld interface{}

    // Mdr Shut Delay. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrShutDelay interface{}

    // Mdr Keep Thru. The type is interface{} with range: -2147483648..2147483647.
    MdrKeepThru interface{}

    // Mdr spoofer. The type is interface{} with range: -2147483648..2147483647.
    MdrSpoofer interface{}

    // Mdr spoofed. The type is interface{} with range: -2147483648..2147483647.
    MdrSpoofed interface{}

    // Mdr spoofed last. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrSpoofedLast interface{}

    // Mdr Spoofed Ready. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrSpoofedReady interface{}

    // Mdr PCB Check. The type is interface{} with range: -2147483648..2147483647.
    MdrPcbCheck interface{}

    // Mdr Kill Tier. The type is interface{} with range: -2147483648..2147483647.
    MdrKillTier interface{}

    // Mdr kld. The type is interface{} with range: -2147483648..2147483647.
    MdrKld interface{}

    // Mdr Level. The type is interface{} with range: -2147483648..2147483647.
    MdrLevel interface{}

    // FM restart count. The type is interface{} with range:
    // -2147483648..2147483647.
    FmRestartCnt interface{}

    // Self Managed. The type is interface{} with range: -2147483648..2147483647.
    SelfManaged interface{}

    // Tuple. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_Tuple.
    Tuple []*SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_Tuple

    // Orig Tuple. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_OrigTuple.
    OrigTuple []*SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_OrigTuple
}

func (verboseInfo *SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo) GetEntityData() *types.CommonEntityData {
    verboseInfo.EntityData.YFilter = verboseInfo.YFilter
    verboseInfo.EntityData.YangName = "verbose-info"
    verboseInfo.EntityData.BundleName = "cisco_ios_xr"
    verboseInfo.EntityData.ParentYangName = "process-name-runverbose"
    verboseInfo.EntityData.SegmentPath = "verbose-info"
    verboseInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-runverboses/process-name-runverbose/" + verboseInfo.EntityData.SegmentPath
    verboseInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    verboseInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    verboseInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    verboseInfo.EntityData.Children = types.NewOrderedMap()
    verboseInfo.EntityData.Children.Append("tuple", types.YChild{"Tuple", nil})
    for i := range verboseInfo.Tuple {
        types.SetYListKey(verboseInfo.Tuple[i], i)
        verboseInfo.EntityData.Children.Append(types.GetSegmentPath(verboseInfo.Tuple[i]), types.YChild{"Tuple", verboseInfo.Tuple[i]})
    }
    verboseInfo.EntityData.Children.Append("orig-tuple", types.YChild{"OrigTuple", nil})
    for i := range verboseInfo.OrigTuple {
        types.SetYListKey(verboseInfo.OrigTuple[i], i)
        verboseInfo.EntityData.Children.Append(types.GetSegmentPath(verboseInfo.OrigTuple[i]), types.YChild{"OrigTuple", verboseInfo.OrigTuple[i]})
    }
    verboseInfo.EntityData.Leafs = types.NewOrderedMap()
    verboseInfo.EntityData.Leafs.Append("process-group", types.YLeaf{"ProcessGroup", verboseInfo.ProcessGroup})
    verboseInfo.EntityData.Leafs.Append("respawn-allowed", types.YLeaf{"RespawnAllowed", verboseInfo.RespawnAllowed})
    verboseInfo.EntityData.Leafs.Append("wait-for-exit", types.YLeaf{"WaitForExit", verboseInfo.WaitForExit})
    verboseInfo.EntityData.Leafs.Append("dynamic-tag", types.YLeaf{"DynamicTag", verboseInfo.DynamicTag})
    verboseInfo.EntityData.Leafs.Append("forced-stop", types.YLeaf{"ForcedStop", verboseInfo.ForcedStop})
    verboseInfo.EntityData.Leafs.Append("critical-process", types.YLeaf{"CriticalProcess", verboseInfo.CriticalProcess})
    verboseInfo.EntityData.Leafs.Append("hold", types.YLeaf{"Hold", verboseInfo.Hold})
    verboseInfo.EntityData.Leafs.Append("transient", types.YLeaf{"Transient", verboseInfo.Transient})
    verboseInfo.EntityData.Leafs.Append("tuple-cfgmgr", types.YLeaf{"TupleCfgmgr", verboseInfo.TupleCfgmgr})
    verboseInfo.EntityData.Leafs.Append("standby-capable", types.YLeaf{"StandbyCapable", verboseInfo.StandbyCapable})
    verboseInfo.EntityData.Leafs.Append("edm-startup", types.YLeaf{"EdmStartup", verboseInfo.EdmStartup})
    verboseInfo.EntityData.Leafs.Append("placement", types.YLeaf{"Placement", verboseInfo.Placement})
    verboseInfo.EntityData.Leafs.Append("skip-kill-notif", types.YLeaf{"SkipKillNotif", verboseInfo.SkipKillNotif})
    verboseInfo.EntityData.Leafs.Append("init-proc", types.YLeaf{"InitProc", verboseInfo.InitProc})
    verboseInfo.EntityData.Leafs.Append("sysdb-event", types.YLeaf{"SysdbEvent", verboseInfo.SysdbEvent})
    verboseInfo.EntityData.Leafs.Append("level-started", types.YLeaf{"LevelStarted", verboseInfo.LevelStarted})
    verboseInfo.EntityData.Leafs.Append("proc-avail", types.YLeaf{"ProcAvail", verboseInfo.ProcAvail})
    verboseInfo.EntityData.Leafs.Append("tuples-scanned", types.YLeaf{"TuplesScanned", verboseInfo.TuplesScanned})
    verboseInfo.EntityData.Leafs.Append("no-chkpt-start", types.YLeaf{"NoChkptStart", verboseInfo.NoChkptStart})
    verboseInfo.EntityData.Leafs.Append("in-shut-down", types.YLeaf{"InShutDown", verboseInfo.InShutDown})
    verboseInfo.EntityData.Leafs.Append("sm-started", types.YLeaf{"SmStarted", verboseInfo.SmStarted})
    verboseInfo.EntityData.Leafs.Append("ignore-on-sc", types.YLeaf{"IgnoreOnSc", verboseInfo.IgnoreOnSc})
    verboseInfo.EntityData.Leafs.Append("ignore-on-easy-bake", types.YLeaf{"IgnoreOnEasyBake", verboseInfo.IgnoreOnEasyBake})
    verboseInfo.EntityData.Leafs.Append("pre-init", types.YLeaf{"PreInit", verboseInfo.PreInit})
    verboseInfo.EntityData.Leafs.Append("eoi-received", types.YLeaf{"EoiReceived", verboseInfo.EoiReceived})
    verboseInfo.EntityData.Leafs.Append("eoi-timeout", types.YLeaf{"EoiTimeout", verboseInfo.EoiTimeout})
    verboseInfo.EntityData.Leafs.Append("avail-timeout", types.YLeaf{"AvailTimeout", verboseInfo.AvailTimeout})
    verboseInfo.EntityData.Leafs.Append("reserved-memory", types.YLeaf{"ReservedMemory", verboseInfo.ReservedMemory})
    verboseInfo.EntityData.Leafs.Append("allow-warned", types.YLeaf{"AllowWarned", verboseInfo.AllowWarned})
    verboseInfo.EntityData.Leafs.Append("arg-change", types.YLeaf{"ArgChange", verboseInfo.ArgChange})
    verboseInfo.EntityData.Leafs.Append("restart-on-tuple", types.YLeaf{"RestartOnTuple", verboseInfo.RestartOnTuple})
    verboseInfo.EntityData.Leafs.Append("boot-hold", types.YLeaf{"BootHold", verboseInfo.BootHold})
    verboseInfo.EntityData.Leafs.Append("reg-id", types.YLeaf{"RegId", verboseInfo.RegId})
    verboseInfo.EntityData.Leafs.Append("memory-limit", types.YLeaf{"MemoryLimit", verboseInfo.MemoryLimit})
    verboseInfo.EntityData.Leafs.Append("parent-job-id", types.YLeaf{"ParentJobId", verboseInfo.ParentJobId})
    verboseInfo.EntityData.Leafs.Append("tuple-index", types.YLeaf{"TupleIndex", verboseInfo.TupleIndex})
    verboseInfo.EntityData.Leafs.Append("dump-count", types.YLeaf{"DumpCount", verboseInfo.DumpCount})
    verboseInfo.EntityData.Leafs.Append("respawn-interval-user", types.YLeaf{"RespawnIntervalUser", verboseInfo.RespawnIntervalUser})
    verboseInfo.EntityData.Leafs.Append("silent-restart-count", types.YLeaf{"SilentRestartCount", verboseInfo.SilentRestartCount})
    verboseInfo.EntityData.Leafs.Append("critical-tier", types.YLeaf{"CriticalTier", verboseInfo.CriticalTier})
    verboseInfo.EntityData.Leafs.Append("exit-type", types.YLeaf{"ExitType", verboseInfo.ExitType})
    verboseInfo.EntityData.Leafs.Append("init-timeout", types.YLeaf{"InitTimeout", verboseInfo.InitTimeout})
    verboseInfo.EntityData.Leafs.Append("restart-by-cmd", types.YLeaf{"RestartByCmd", verboseInfo.RestartByCmd})
    verboseInfo.EntityData.Leafs.Append("boot-pref", types.YLeaf{"BootPref", verboseInfo.BootPref})
    verboseInfo.EntityData.Leafs.Append("mdr-mbi-proc", types.YLeaf{"MdrMbiProc", verboseInfo.MdrMbiProc})
    verboseInfo.EntityData.Leafs.Append("mdr-non-mbi-kld", types.YLeaf{"MdrNonMbiKld", verboseInfo.MdrNonMbiKld})
    verboseInfo.EntityData.Leafs.Append("mdr-mbi-kld", types.YLeaf{"MdrMbiKld", verboseInfo.MdrMbiKld})
    verboseInfo.EntityData.Leafs.Append("mdr-shut-delay", types.YLeaf{"MdrShutDelay", verboseInfo.MdrShutDelay})
    verboseInfo.EntityData.Leafs.Append("mdr-keep-thru", types.YLeaf{"MdrKeepThru", verboseInfo.MdrKeepThru})
    verboseInfo.EntityData.Leafs.Append("mdr-spoofer", types.YLeaf{"MdrSpoofer", verboseInfo.MdrSpoofer})
    verboseInfo.EntityData.Leafs.Append("mdr-spoofed", types.YLeaf{"MdrSpoofed", verboseInfo.MdrSpoofed})
    verboseInfo.EntityData.Leafs.Append("mdr-spoofed-last", types.YLeaf{"MdrSpoofedLast", verboseInfo.MdrSpoofedLast})
    verboseInfo.EntityData.Leafs.Append("mdr-spoofed-ready", types.YLeaf{"MdrSpoofedReady", verboseInfo.MdrSpoofedReady})
    verboseInfo.EntityData.Leafs.Append("mdr-pcb-check", types.YLeaf{"MdrPcbCheck", verboseInfo.MdrPcbCheck})
    verboseInfo.EntityData.Leafs.Append("mdr-kill-tier", types.YLeaf{"MdrKillTier", verboseInfo.MdrKillTier})
    verboseInfo.EntityData.Leafs.Append("mdr-kld", types.YLeaf{"MdrKld", verboseInfo.MdrKld})
    verboseInfo.EntityData.Leafs.Append("mdr-level", types.YLeaf{"MdrLevel", verboseInfo.MdrLevel})
    verboseInfo.EntityData.Leafs.Append("fm-restart-cnt", types.YLeaf{"FmRestartCnt", verboseInfo.FmRestartCnt})
    verboseInfo.EntityData.Leafs.Append("self-managed", types.YLeaf{"SelfManaged", verboseInfo.SelfManaged})

    verboseInfo.EntityData.YListKeys = []string {}

    return &(verboseInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_Tuple
// Tuple
type SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_Tuple struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (tuple *SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_Tuple) GetEntityData() *types.CommonEntityData {
    tuple.EntityData.YFilter = tuple.YFilter
    tuple.EntityData.YangName = "tuple"
    tuple.EntityData.BundleName = "cisco_ios_xr"
    tuple.EntityData.ParentYangName = "verbose-info"
    tuple.EntityData.SegmentPath = "tuple" + types.AddNoKeyToken(tuple)
    tuple.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-runverboses/process-name-runverbose/verbose-info/" + tuple.EntityData.SegmentPath
    tuple.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tuple.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tuple.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tuple.EntityData.Children = types.NewOrderedMap()
    tuple.EntityData.Leafs = types.NewOrderedMap()
    tuple.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", tuple.Tuple})

    tuple.EntityData.YListKeys = []string {}

    return &(tuple.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_OrigTuple
// Orig Tuple
type SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_OrigTuple struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (origTuple *SystemProcess_NodeTable_Node_Name_ProcessNameRunverboses_ProcessNameRunverbose_VerboseInfo_OrigTuple) GetEntityData() *types.CommonEntityData {
    origTuple.EntityData.YFilter = origTuple.YFilter
    origTuple.EntityData.YangName = "orig-tuple"
    origTuple.EntityData.BundleName = "cisco_ios_xr"
    origTuple.EntityData.ParentYangName = "verbose-info"
    origTuple.EntityData.SegmentPath = "orig-tuple" + types.AddNoKeyToken(origTuple)
    origTuple.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-runverboses/process-name-runverbose/verbose-info/" + origTuple.EntityData.SegmentPath
    origTuple.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    origTuple.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    origTuple.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    origTuple.EntityData.Children = types.NewOrderedMap()
    origTuple.EntityData.Leafs = types.NewOrderedMap()
    origTuple.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", origTuple.Tuple})

    origTuple.EntityData.YListKeys = []string {}

    return &(origTuple.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameDetails
// Process <WORD> information
type SystemProcess_NodeTable_Node_Name_ProcessNameDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process <WORD> detail information. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail.
    ProcessNameDetail []*SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail
}

func (processNameDetails *SystemProcess_NodeTable_Node_Name_ProcessNameDetails) GetEntityData() *types.CommonEntityData {
    processNameDetails.EntityData.YFilter = processNameDetails.YFilter
    processNameDetails.EntityData.YangName = "process-name-details"
    processNameDetails.EntityData.BundleName = "cisco_ios_xr"
    processNameDetails.EntityData.ParentYangName = "name"
    processNameDetails.EntityData.SegmentPath = "process-name-details"
    processNameDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/" + processNameDetails.EntityData.SegmentPath
    processNameDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameDetails.EntityData.Children = types.NewOrderedMap()
    processNameDetails.EntityData.Children.Append("process-name-detail", types.YChild{"ProcessNameDetail", nil})
    for i := range processNameDetails.ProcessNameDetail {
        processNameDetails.EntityData.Children.Append(types.GetSegmentPath(processNameDetails.ProcessNameDetail[i]), types.YChild{"ProcessNameDetail", processNameDetails.ProcessNameDetail[i]})
    }
    processNameDetails.EntityData.Leafs = types.NewOrderedMap()

    processNameDetails.EntityData.YListKeys = []string {}

    return &(processNameDetails.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail
// Process <WORD> detail information
type SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcName interface{}

    // Process Basic Info.
    BasicInfo SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo

    // Process Detail Info.
    DetailInfo SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_DetailInfo
}

func (processNameDetail *SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail) GetEntityData() *types.CommonEntityData {
    processNameDetail.EntityData.YFilter = processNameDetail.YFilter
    processNameDetail.EntityData.YangName = "process-name-detail"
    processNameDetail.EntityData.BundleName = "cisco_ios_xr"
    processNameDetail.EntityData.ParentYangName = "process-name-details"
    processNameDetail.EntityData.SegmentPath = "process-name-detail" + types.AddKeyToken(processNameDetail.ProcName, "proc-name")
    processNameDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-details/" + processNameDetail.EntityData.SegmentPath
    processNameDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameDetail.EntityData.Children = types.NewOrderedMap()
    processNameDetail.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &processNameDetail.BasicInfo})
    processNameDetail.EntityData.Children.Append("detail-info", types.YChild{"DetailInfo", &processNameDetail.DetailInfo})
    processNameDetail.EntityData.Leafs = types.NewOrderedMap()
    processNameDetail.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", processNameDetail.ProcName})

    processNameDetail.EntityData.YListKeys = []string {"ProcName"}

    return &(processNameDetail.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo
// Process Basic Info
type SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Job ID. The type is interface{} with range: 0..4294967295.
    JobIdXr interface{}

    // PID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Process name. The type is string.
    ProcessName interface{}

    // Executable name or path. The type is string.
    Executable interface{}

    // Active Path. The type is string.
    ActivePath interface{}

    // Instance ID. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Args. The type is string.
    Args interface{}

    // Version ID. The type is string.
    VersionId interface{}

    // Respawn on/off. The type is string.
    Respawn interface{}

    // Respawn Count. The type is interface{} with range: -2147483648..2147483647.
    RespawnCount interface{}

    // Last Started. The type is string.
    LastStarted interface{}

    // Process State. The type is string.
    ProcessState interface{}

    // Last Exit status. The type is interface{} with range:
    // -2147483648..2147483647.
    LastExitStatus interface{}

    // Last Exit due to. The type is string.
    LastExitReason interface{}

    // Package State. The type is string.
    PackageState interface{}

    // Started on Config. The type is string.
    StartedOnConfig interface{}

    // Feature Name. The type is string.
    FeatureName interface{}

    // Tag. The type is string.
    Tag interface{}

    // Process Group. The type is string.
    Group interface{}

    // Core. The type is string.
    Core interface{}

    // Max core. The type is interface{} with range: -2147483648..2147483647.
    MaxCore interface{}

    // Level. The type is string.
    Level interface{}

    // Is mandatory?. The type is bool.
    Mandatory interface{}

    // Is admin mode process?. The type is bool.
    MaintModeProc interface{}

    // Placement State. The type is string.
    PlacementState interface{}

    // Startup Path. The type is string.
    StartUpPath interface{}

    // Memory Limit. The type is interface{} with range: 0..4294967295.
    MemoryLimit interface{}

    // Elapsed Ready. The type is string.
    Ready interface{}

    // Elapsed Available. The type is string.
    Available interface{}

    // Proces cpu time.
    ProcCpuTime SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo_ProcCpuTime

    // Registered Items. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo_RegisteredItem.
    RegisteredItem []*SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo_RegisteredItem
}

func (basicInfo *SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "process-name-detail"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-details/process-name-detail/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Children.Append("proc-cpu-time", types.YChild{"ProcCpuTime", &basicInfo.ProcCpuTime})
    basicInfo.EntityData.Children.Append("registered-item", types.YChild{"RegisteredItem", nil})
    for i := range basicInfo.RegisteredItem {
        types.SetYListKey(basicInfo.RegisteredItem[i], i)
        basicInfo.EntityData.Children.Append(types.GetSegmentPath(basicInfo.RegisteredItem[i]), types.YChild{"RegisteredItem", basicInfo.RegisteredItem[i]})
    }
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("job-id-xr", types.YLeaf{"JobIdXr", basicInfo.JobIdXr})
    basicInfo.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", basicInfo.ProcessId})
    basicInfo.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", basicInfo.ProcessName})
    basicInfo.EntityData.Leafs.Append("executable", types.YLeaf{"Executable", basicInfo.Executable})
    basicInfo.EntityData.Leafs.Append("active-path", types.YLeaf{"ActivePath", basicInfo.ActivePath})
    basicInfo.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", basicInfo.InstanceId})
    basicInfo.EntityData.Leafs.Append("args", types.YLeaf{"Args", basicInfo.Args})
    basicInfo.EntityData.Leafs.Append("version-id", types.YLeaf{"VersionId", basicInfo.VersionId})
    basicInfo.EntityData.Leafs.Append("respawn", types.YLeaf{"Respawn", basicInfo.Respawn})
    basicInfo.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", basicInfo.RespawnCount})
    basicInfo.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", basicInfo.LastStarted})
    basicInfo.EntityData.Leafs.Append("process-state", types.YLeaf{"ProcessState", basicInfo.ProcessState})
    basicInfo.EntityData.Leafs.Append("last-exit-status", types.YLeaf{"LastExitStatus", basicInfo.LastExitStatus})
    basicInfo.EntityData.Leafs.Append("last-exit-reason", types.YLeaf{"LastExitReason", basicInfo.LastExitReason})
    basicInfo.EntityData.Leafs.Append("package-state", types.YLeaf{"PackageState", basicInfo.PackageState})
    basicInfo.EntityData.Leafs.Append("started-on-config", types.YLeaf{"StartedOnConfig", basicInfo.StartedOnConfig})
    basicInfo.EntityData.Leafs.Append("feature-name", types.YLeaf{"FeatureName", basicInfo.FeatureName})
    basicInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", basicInfo.Tag})
    basicInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", basicInfo.Group})
    basicInfo.EntityData.Leafs.Append("core", types.YLeaf{"Core", basicInfo.Core})
    basicInfo.EntityData.Leafs.Append("max-core", types.YLeaf{"MaxCore", basicInfo.MaxCore})
    basicInfo.EntityData.Leafs.Append("level", types.YLeaf{"Level", basicInfo.Level})
    basicInfo.EntityData.Leafs.Append("mandatory", types.YLeaf{"Mandatory", basicInfo.Mandatory})
    basicInfo.EntityData.Leafs.Append("maint-mode-proc", types.YLeaf{"MaintModeProc", basicInfo.MaintModeProc})
    basicInfo.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", basicInfo.PlacementState})
    basicInfo.EntityData.Leafs.Append("start-up-path", types.YLeaf{"StartUpPath", basicInfo.StartUpPath})
    basicInfo.EntityData.Leafs.Append("memory-limit", types.YLeaf{"MemoryLimit", basicInfo.MemoryLimit})
    basicInfo.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", basicInfo.Ready})
    basicInfo.EntityData.Leafs.Append("available", types.YLeaf{"Available", basicInfo.Available})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo_ProcCpuTime
// Proces cpu time
type SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo_ProcCpuTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User time. The type is string.
    User interface{}

    // Kernel time. The type is string.
    System interface{}

    // Total time. The type is string.
    Total interface{}
}

func (procCpuTime *SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo_ProcCpuTime) GetEntityData() *types.CommonEntityData {
    procCpuTime.EntityData.YFilter = procCpuTime.YFilter
    procCpuTime.EntityData.YangName = "proc-cpu-time"
    procCpuTime.EntityData.BundleName = "cisco_ios_xr"
    procCpuTime.EntityData.ParentYangName = "basic-info"
    procCpuTime.EntityData.SegmentPath = "proc-cpu-time"
    procCpuTime.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-details/process-name-detail/basic-info/" + procCpuTime.EntityData.SegmentPath
    procCpuTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procCpuTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procCpuTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procCpuTime.EntityData.Children = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs.Append("user", types.YLeaf{"User", procCpuTime.User})
    procCpuTime.EntityData.Leafs.Append("system", types.YLeaf{"System", procCpuTime.System})
    procCpuTime.EntityData.Leafs.Append("total", types.YLeaf{"Total", procCpuTime.Total})

    procCpuTime.EntityData.YListKeys = []string {}

    return &(procCpuTime.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo_RegisteredItem
// Registered Items
type SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo_RegisteredItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (registeredItem *SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_BasicInfo_RegisteredItem) GetEntityData() *types.CommonEntityData {
    registeredItem.EntityData.YFilter = registeredItem.YFilter
    registeredItem.EntityData.YangName = "registered-item"
    registeredItem.EntityData.BundleName = "cisco_ios_xr"
    registeredItem.EntityData.ParentYangName = "basic-info"
    registeredItem.EntityData.SegmentPath = "registered-item" + types.AddNoKeyToken(registeredItem)
    registeredItem.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-details/process-name-detail/basic-info/" + registeredItem.EntityData.SegmentPath
    registeredItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registeredItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registeredItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registeredItem.EntityData.Children = types.NewOrderedMap()
    registeredItem.EntityData.Leafs = types.NewOrderedMap()
    registeredItem.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", registeredItem.Tuple})

    registeredItem.EntityData.YListKeys = []string {}

    return &(registeredItem.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_DetailInfo
// Process Detail Info
type SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_DetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Running path. The type is string.
    RunningPath interface{}

    // Package path. The type is string.
    PackagePath interface{}

    // Job Id Link. The type is interface{} with range: -2147483648..2147483647.
    JobIdLink interface{}

    // Group Jid. The type is string.
    GroupJid interface{}

    // Fail count. The type is interface{} with range: 0..4294967295.
    FailCount interface{}

    // Restart needed. The type is bool.
    RestartNeeded interface{}

    // Init process. The type is bool.
    InitProcess interface{}

    // Last Online. The type is string.
    LastOnline interface{}

    // This PCB. The type is string.
    ThisPcb interface{}

    // Next PCB. The type is string.
    NextPcb interface{}

    // Env variables. The type is string.
    Envs interface{}

    // Wait For /dev/xxx. The type is string.
    WaitFor interface{}

    // Job ID on RP. The type is interface{} with range: -2147483648..2147483647.
    JobIdOnRp interface{}

    // Is standby capable?. The type is bool.
    IsStandbyCapable interface{}

    // Disable kill?. The type is bool.
    DisableKill interface{}

    // Check avail. The type is bool.
    SendAvail interface{}

    // Node Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    NodeEventCliInfo interface{}

    // Node redundancy state. The type is string.
    NodeRedundancyState interface{}

    // Role event cli info. The type is interface{} with range:
    // -2147483648..2147483647.
    RoleEventCliInfo interface{}

    // Proc Role State. The type is string.
    ProcRoleState interface{}

    // Standby Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    StandbyEventCliInfo interface{}

    // Cleanup event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    CleanupEventCliInfo interface{}

    // Band Ready Event CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    BandReadyEventCliInfo interface{}

    // LR Event CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    LrEventCliInfo interface{}

    // Plane Ready Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    PlaneReadyEventCliInfo interface{}

    // MDR is done CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrIsDoneCliInfo interface{}
}

func (detailInfo *SystemProcess_NodeTable_Node_Name_ProcessNameDetails_ProcessNameDetail_DetailInfo) GetEntityData() *types.CommonEntityData {
    detailInfo.EntityData.YFilter = detailInfo.YFilter
    detailInfo.EntityData.YangName = "detail-info"
    detailInfo.EntityData.BundleName = "cisco_ios_xr"
    detailInfo.EntityData.ParentYangName = "process-name-detail"
    detailInfo.EntityData.SegmentPath = "detail-info"
    detailInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-details/process-name-detail/" + detailInfo.EntityData.SegmentPath
    detailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailInfo.EntityData.Children = types.NewOrderedMap()
    detailInfo.EntityData.Leafs = types.NewOrderedMap()
    detailInfo.EntityData.Leafs.Append("running-path", types.YLeaf{"RunningPath", detailInfo.RunningPath})
    detailInfo.EntityData.Leafs.Append("package-path", types.YLeaf{"PackagePath", detailInfo.PackagePath})
    detailInfo.EntityData.Leafs.Append("job-id-link", types.YLeaf{"JobIdLink", detailInfo.JobIdLink})
    detailInfo.EntityData.Leafs.Append("group-jid", types.YLeaf{"GroupJid", detailInfo.GroupJid})
    detailInfo.EntityData.Leafs.Append("fail-count", types.YLeaf{"FailCount", detailInfo.FailCount})
    detailInfo.EntityData.Leafs.Append("restart-needed", types.YLeaf{"RestartNeeded", detailInfo.RestartNeeded})
    detailInfo.EntityData.Leafs.Append("init-process", types.YLeaf{"InitProcess", detailInfo.InitProcess})
    detailInfo.EntityData.Leafs.Append("last-online", types.YLeaf{"LastOnline", detailInfo.LastOnline})
    detailInfo.EntityData.Leafs.Append("this-pcb", types.YLeaf{"ThisPcb", detailInfo.ThisPcb})
    detailInfo.EntityData.Leafs.Append("next-pcb", types.YLeaf{"NextPcb", detailInfo.NextPcb})
    detailInfo.EntityData.Leafs.Append("envs", types.YLeaf{"Envs", detailInfo.Envs})
    detailInfo.EntityData.Leafs.Append("wait-for", types.YLeaf{"WaitFor", detailInfo.WaitFor})
    detailInfo.EntityData.Leafs.Append("job-id-on-rp", types.YLeaf{"JobIdOnRp", detailInfo.JobIdOnRp})
    detailInfo.EntityData.Leafs.Append("is-standby-capable", types.YLeaf{"IsStandbyCapable", detailInfo.IsStandbyCapable})
    detailInfo.EntityData.Leafs.Append("disable-kill", types.YLeaf{"DisableKill", detailInfo.DisableKill})
    detailInfo.EntityData.Leafs.Append("send-avail", types.YLeaf{"SendAvail", detailInfo.SendAvail})
    detailInfo.EntityData.Leafs.Append("node-event-cli-info", types.YLeaf{"NodeEventCliInfo", detailInfo.NodeEventCliInfo})
    detailInfo.EntityData.Leafs.Append("node-redundancy-state", types.YLeaf{"NodeRedundancyState", detailInfo.NodeRedundancyState})
    detailInfo.EntityData.Leafs.Append("role-event-cli-info", types.YLeaf{"RoleEventCliInfo", detailInfo.RoleEventCliInfo})
    detailInfo.EntityData.Leafs.Append("proc-role-state", types.YLeaf{"ProcRoleState", detailInfo.ProcRoleState})
    detailInfo.EntityData.Leafs.Append("standby-event-cli-info", types.YLeaf{"StandbyEventCliInfo", detailInfo.StandbyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("cleanup-event-cli-info", types.YLeaf{"CleanupEventCliInfo", detailInfo.CleanupEventCliInfo})
    detailInfo.EntityData.Leafs.Append("band-ready-event-cli-info", types.YLeaf{"BandReadyEventCliInfo", detailInfo.BandReadyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("lr-event-cli-info", types.YLeaf{"LrEventCliInfo", detailInfo.LrEventCliInfo})
    detailInfo.EntityData.Leafs.Append("plane-ready-event-cli-info", types.YLeaf{"PlaneReadyEventCliInfo", detailInfo.PlaneReadyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("mdr-is-done-cli-info", types.YLeaf{"MdrIsDoneCliInfo", detailInfo.MdrIsDoneCliInfo})

    detailInfo.EntityData.YListKeys = []string {}

    return &(detailInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameVerboses
// Process <WORD> information
type SystemProcess_NodeTable_Node_Name_ProcessNameVerboses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process <WORD> verbose information. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose.
    ProcessNameVerbose []*SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose
}

func (processNameVerboses *SystemProcess_NodeTable_Node_Name_ProcessNameVerboses) GetEntityData() *types.CommonEntityData {
    processNameVerboses.EntityData.YFilter = processNameVerboses.YFilter
    processNameVerboses.EntityData.YangName = "process-name-verboses"
    processNameVerboses.EntityData.BundleName = "cisco_ios_xr"
    processNameVerboses.EntityData.ParentYangName = "name"
    processNameVerboses.EntityData.SegmentPath = "process-name-verboses"
    processNameVerboses.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/" + processNameVerboses.EntityData.SegmentPath
    processNameVerboses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameVerboses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameVerboses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameVerboses.EntityData.Children = types.NewOrderedMap()
    processNameVerboses.EntityData.Children.Append("process-name-verbose", types.YChild{"ProcessNameVerbose", nil})
    for i := range processNameVerboses.ProcessNameVerbose {
        processNameVerboses.EntityData.Children.Append(types.GetSegmentPath(processNameVerboses.ProcessNameVerbose[i]), types.YChild{"ProcessNameVerbose", processNameVerboses.ProcessNameVerbose[i]})
    }
    processNameVerboses.EntityData.Leafs = types.NewOrderedMap()

    processNameVerboses.EntityData.YListKeys = []string {}

    return &(processNameVerboses.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose
// Process <WORD> verbose information
type SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Process Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcName interface{}

    // Process Basic Info.
    BasicInfo SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo

    // Process Detail Info.
    DetailInfo SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_DetailInfo

    // Process Verbose Info.
    VerboseInfo SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo
}

func (processNameVerbose *SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose) GetEntityData() *types.CommonEntityData {
    processNameVerbose.EntityData.YFilter = processNameVerbose.YFilter
    processNameVerbose.EntityData.YangName = "process-name-verbose"
    processNameVerbose.EntityData.BundleName = "cisco_ios_xr"
    processNameVerbose.EntityData.ParentYangName = "process-name-verboses"
    processNameVerbose.EntityData.SegmentPath = "process-name-verbose" + types.AddKeyToken(processNameVerbose.ProcName, "proc-name")
    processNameVerbose.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-verboses/" + processNameVerbose.EntityData.SegmentPath
    processNameVerbose.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNameVerbose.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNameVerbose.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNameVerbose.EntityData.Children = types.NewOrderedMap()
    processNameVerbose.EntityData.Children.Append("basic-info", types.YChild{"BasicInfo", &processNameVerbose.BasicInfo})
    processNameVerbose.EntityData.Children.Append("detail-info", types.YChild{"DetailInfo", &processNameVerbose.DetailInfo})
    processNameVerbose.EntityData.Children.Append("verbose-info", types.YChild{"VerboseInfo", &processNameVerbose.VerboseInfo})
    processNameVerbose.EntityData.Leafs = types.NewOrderedMap()
    processNameVerbose.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", processNameVerbose.ProcName})

    processNameVerbose.EntityData.YListKeys = []string {"ProcName"}

    return &(processNameVerbose.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo
// Process Basic Info
type SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Job ID. The type is interface{} with range: 0..4294967295.
    JobIdXr interface{}

    // PID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Process name. The type is string.
    ProcessName interface{}

    // Executable name or path. The type is string.
    Executable interface{}

    // Active Path. The type is string.
    ActivePath interface{}

    // Instance ID. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Args. The type is string.
    Args interface{}

    // Version ID. The type is string.
    VersionId interface{}

    // Respawn on/off. The type is string.
    Respawn interface{}

    // Respawn Count. The type is interface{} with range: -2147483648..2147483647.
    RespawnCount interface{}

    // Last Started. The type is string.
    LastStarted interface{}

    // Process State. The type is string.
    ProcessState interface{}

    // Last Exit status. The type is interface{} with range:
    // -2147483648..2147483647.
    LastExitStatus interface{}

    // Last Exit due to. The type is string.
    LastExitReason interface{}

    // Package State. The type is string.
    PackageState interface{}

    // Started on Config. The type is string.
    StartedOnConfig interface{}

    // Feature Name. The type is string.
    FeatureName interface{}

    // Tag. The type is string.
    Tag interface{}

    // Process Group. The type is string.
    Group interface{}

    // Core. The type is string.
    Core interface{}

    // Max core. The type is interface{} with range: -2147483648..2147483647.
    MaxCore interface{}

    // Level. The type is string.
    Level interface{}

    // Is mandatory?. The type is bool.
    Mandatory interface{}

    // Is admin mode process?. The type is bool.
    MaintModeProc interface{}

    // Placement State. The type is string.
    PlacementState interface{}

    // Startup Path. The type is string.
    StartUpPath interface{}

    // Memory Limit. The type is interface{} with range: 0..4294967295.
    MemoryLimit interface{}

    // Elapsed Ready. The type is string.
    Ready interface{}

    // Elapsed Available. The type is string.
    Available interface{}

    // Proces cpu time.
    ProcCpuTime SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo_ProcCpuTime

    // Registered Items. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo_RegisteredItem.
    RegisteredItem []*SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo_RegisteredItem
}

func (basicInfo *SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo) GetEntityData() *types.CommonEntityData {
    basicInfo.EntityData.YFilter = basicInfo.YFilter
    basicInfo.EntityData.YangName = "basic-info"
    basicInfo.EntityData.BundleName = "cisco_ios_xr"
    basicInfo.EntityData.ParentYangName = "process-name-verbose"
    basicInfo.EntityData.SegmentPath = "basic-info"
    basicInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-verboses/process-name-verbose/" + basicInfo.EntityData.SegmentPath
    basicInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicInfo.EntityData.Children = types.NewOrderedMap()
    basicInfo.EntityData.Children.Append("proc-cpu-time", types.YChild{"ProcCpuTime", &basicInfo.ProcCpuTime})
    basicInfo.EntityData.Children.Append("registered-item", types.YChild{"RegisteredItem", nil})
    for i := range basicInfo.RegisteredItem {
        types.SetYListKey(basicInfo.RegisteredItem[i], i)
        basicInfo.EntityData.Children.Append(types.GetSegmentPath(basicInfo.RegisteredItem[i]), types.YChild{"RegisteredItem", basicInfo.RegisteredItem[i]})
    }
    basicInfo.EntityData.Leafs = types.NewOrderedMap()
    basicInfo.EntityData.Leafs.Append("job-id-xr", types.YLeaf{"JobIdXr", basicInfo.JobIdXr})
    basicInfo.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", basicInfo.ProcessId})
    basicInfo.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", basicInfo.ProcessName})
    basicInfo.EntityData.Leafs.Append("executable", types.YLeaf{"Executable", basicInfo.Executable})
    basicInfo.EntityData.Leafs.Append("active-path", types.YLeaf{"ActivePath", basicInfo.ActivePath})
    basicInfo.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", basicInfo.InstanceId})
    basicInfo.EntityData.Leafs.Append("args", types.YLeaf{"Args", basicInfo.Args})
    basicInfo.EntityData.Leafs.Append("version-id", types.YLeaf{"VersionId", basicInfo.VersionId})
    basicInfo.EntityData.Leafs.Append("respawn", types.YLeaf{"Respawn", basicInfo.Respawn})
    basicInfo.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", basicInfo.RespawnCount})
    basicInfo.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", basicInfo.LastStarted})
    basicInfo.EntityData.Leafs.Append("process-state", types.YLeaf{"ProcessState", basicInfo.ProcessState})
    basicInfo.EntityData.Leafs.Append("last-exit-status", types.YLeaf{"LastExitStatus", basicInfo.LastExitStatus})
    basicInfo.EntityData.Leafs.Append("last-exit-reason", types.YLeaf{"LastExitReason", basicInfo.LastExitReason})
    basicInfo.EntityData.Leafs.Append("package-state", types.YLeaf{"PackageState", basicInfo.PackageState})
    basicInfo.EntityData.Leafs.Append("started-on-config", types.YLeaf{"StartedOnConfig", basicInfo.StartedOnConfig})
    basicInfo.EntityData.Leafs.Append("feature-name", types.YLeaf{"FeatureName", basicInfo.FeatureName})
    basicInfo.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", basicInfo.Tag})
    basicInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", basicInfo.Group})
    basicInfo.EntityData.Leafs.Append("core", types.YLeaf{"Core", basicInfo.Core})
    basicInfo.EntityData.Leafs.Append("max-core", types.YLeaf{"MaxCore", basicInfo.MaxCore})
    basicInfo.EntityData.Leafs.Append("level", types.YLeaf{"Level", basicInfo.Level})
    basicInfo.EntityData.Leafs.Append("mandatory", types.YLeaf{"Mandatory", basicInfo.Mandatory})
    basicInfo.EntityData.Leafs.Append("maint-mode-proc", types.YLeaf{"MaintModeProc", basicInfo.MaintModeProc})
    basicInfo.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", basicInfo.PlacementState})
    basicInfo.EntityData.Leafs.Append("start-up-path", types.YLeaf{"StartUpPath", basicInfo.StartUpPath})
    basicInfo.EntityData.Leafs.Append("memory-limit", types.YLeaf{"MemoryLimit", basicInfo.MemoryLimit})
    basicInfo.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", basicInfo.Ready})
    basicInfo.EntityData.Leafs.Append("available", types.YLeaf{"Available", basicInfo.Available})

    basicInfo.EntityData.YListKeys = []string {}

    return &(basicInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo_ProcCpuTime
// Proces cpu time
type SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo_ProcCpuTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User time. The type is string.
    User interface{}

    // Kernel time. The type is string.
    System interface{}

    // Total time. The type is string.
    Total interface{}
}

func (procCpuTime *SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo_ProcCpuTime) GetEntityData() *types.CommonEntityData {
    procCpuTime.EntityData.YFilter = procCpuTime.YFilter
    procCpuTime.EntityData.YangName = "proc-cpu-time"
    procCpuTime.EntityData.BundleName = "cisco_ios_xr"
    procCpuTime.EntityData.ParentYangName = "basic-info"
    procCpuTime.EntityData.SegmentPath = "proc-cpu-time"
    procCpuTime.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-verboses/process-name-verbose/basic-info/" + procCpuTime.EntityData.SegmentPath
    procCpuTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procCpuTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procCpuTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procCpuTime.EntityData.Children = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs.Append("user", types.YLeaf{"User", procCpuTime.User})
    procCpuTime.EntityData.Leafs.Append("system", types.YLeaf{"System", procCpuTime.System})
    procCpuTime.EntityData.Leafs.Append("total", types.YLeaf{"Total", procCpuTime.Total})

    procCpuTime.EntityData.YListKeys = []string {}

    return &(procCpuTime.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo_RegisteredItem
// Registered Items
type SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo_RegisteredItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (registeredItem *SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_BasicInfo_RegisteredItem) GetEntityData() *types.CommonEntityData {
    registeredItem.EntityData.YFilter = registeredItem.YFilter
    registeredItem.EntityData.YangName = "registered-item"
    registeredItem.EntityData.BundleName = "cisco_ios_xr"
    registeredItem.EntityData.ParentYangName = "basic-info"
    registeredItem.EntityData.SegmentPath = "registered-item" + types.AddNoKeyToken(registeredItem)
    registeredItem.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-verboses/process-name-verbose/basic-info/" + registeredItem.EntityData.SegmentPath
    registeredItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registeredItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registeredItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registeredItem.EntityData.Children = types.NewOrderedMap()
    registeredItem.EntityData.Leafs = types.NewOrderedMap()
    registeredItem.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", registeredItem.Tuple})

    registeredItem.EntityData.YListKeys = []string {}

    return &(registeredItem.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_DetailInfo
// Process Detail Info
type SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_DetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Running path. The type is string.
    RunningPath interface{}

    // Package path. The type is string.
    PackagePath interface{}

    // Job Id Link. The type is interface{} with range: -2147483648..2147483647.
    JobIdLink interface{}

    // Group Jid. The type is string.
    GroupJid interface{}

    // Fail count. The type is interface{} with range: 0..4294967295.
    FailCount interface{}

    // Restart needed. The type is bool.
    RestartNeeded interface{}

    // Init process. The type is bool.
    InitProcess interface{}

    // Last Online. The type is string.
    LastOnline interface{}

    // This PCB. The type is string.
    ThisPcb interface{}

    // Next PCB. The type is string.
    NextPcb interface{}

    // Env variables. The type is string.
    Envs interface{}

    // Wait For /dev/xxx. The type is string.
    WaitFor interface{}

    // Job ID on RP. The type is interface{} with range: -2147483648..2147483647.
    JobIdOnRp interface{}

    // Is standby capable?. The type is bool.
    IsStandbyCapable interface{}

    // Disable kill?. The type is bool.
    DisableKill interface{}

    // Check avail. The type is bool.
    SendAvail interface{}

    // Node Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    NodeEventCliInfo interface{}

    // Node redundancy state. The type is string.
    NodeRedundancyState interface{}

    // Role event cli info. The type is interface{} with range:
    // -2147483648..2147483647.
    RoleEventCliInfo interface{}

    // Proc Role State. The type is string.
    ProcRoleState interface{}

    // Standby Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    StandbyEventCliInfo interface{}

    // Cleanup event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    CleanupEventCliInfo interface{}

    // Band Ready Event CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    BandReadyEventCliInfo interface{}

    // LR Event CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    LrEventCliInfo interface{}

    // Plane Ready Event CLI info. The type is interface{} with range:
    // -2147483648..2147483647.
    PlaneReadyEventCliInfo interface{}

    // MDR is done CLI Info. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrIsDoneCliInfo interface{}
}

func (detailInfo *SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_DetailInfo) GetEntityData() *types.CommonEntityData {
    detailInfo.EntityData.YFilter = detailInfo.YFilter
    detailInfo.EntityData.YangName = "detail-info"
    detailInfo.EntityData.BundleName = "cisco_ios_xr"
    detailInfo.EntityData.ParentYangName = "process-name-verbose"
    detailInfo.EntityData.SegmentPath = "detail-info"
    detailInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-verboses/process-name-verbose/" + detailInfo.EntityData.SegmentPath
    detailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailInfo.EntityData.Children = types.NewOrderedMap()
    detailInfo.EntityData.Leafs = types.NewOrderedMap()
    detailInfo.EntityData.Leafs.Append("running-path", types.YLeaf{"RunningPath", detailInfo.RunningPath})
    detailInfo.EntityData.Leafs.Append("package-path", types.YLeaf{"PackagePath", detailInfo.PackagePath})
    detailInfo.EntityData.Leafs.Append("job-id-link", types.YLeaf{"JobIdLink", detailInfo.JobIdLink})
    detailInfo.EntityData.Leafs.Append("group-jid", types.YLeaf{"GroupJid", detailInfo.GroupJid})
    detailInfo.EntityData.Leafs.Append("fail-count", types.YLeaf{"FailCount", detailInfo.FailCount})
    detailInfo.EntityData.Leafs.Append("restart-needed", types.YLeaf{"RestartNeeded", detailInfo.RestartNeeded})
    detailInfo.EntityData.Leafs.Append("init-process", types.YLeaf{"InitProcess", detailInfo.InitProcess})
    detailInfo.EntityData.Leafs.Append("last-online", types.YLeaf{"LastOnline", detailInfo.LastOnline})
    detailInfo.EntityData.Leafs.Append("this-pcb", types.YLeaf{"ThisPcb", detailInfo.ThisPcb})
    detailInfo.EntityData.Leafs.Append("next-pcb", types.YLeaf{"NextPcb", detailInfo.NextPcb})
    detailInfo.EntityData.Leafs.Append("envs", types.YLeaf{"Envs", detailInfo.Envs})
    detailInfo.EntityData.Leafs.Append("wait-for", types.YLeaf{"WaitFor", detailInfo.WaitFor})
    detailInfo.EntityData.Leafs.Append("job-id-on-rp", types.YLeaf{"JobIdOnRp", detailInfo.JobIdOnRp})
    detailInfo.EntityData.Leafs.Append("is-standby-capable", types.YLeaf{"IsStandbyCapable", detailInfo.IsStandbyCapable})
    detailInfo.EntityData.Leafs.Append("disable-kill", types.YLeaf{"DisableKill", detailInfo.DisableKill})
    detailInfo.EntityData.Leafs.Append("send-avail", types.YLeaf{"SendAvail", detailInfo.SendAvail})
    detailInfo.EntityData.Leafs.Append("node-event-cli-info", types.YLeaf{"NodeEventCliInfo", detailInfo.NodeEventCliInfo})
    detailInfo.EntityData.Leafs.Append("node-redundancy-state", types.YLeaf{"NodeRedundancyState", detailInfo.NodeRedundancyState})
    detailInfo.EntityData.Leafs.Append("role-event-cli-info", types.YLeaf{"RoleEventCliInfo", detailInfo.RoleEventCliInfo})
    detailInfo.EntityData.Leafs.Append("proc-role-state", types.YLeaf{"ProcRoleState", detailInfo.ProcRoleState})
    detailInfo.EntityData.Leafs.Append("standby-event-cli-info", types.YLeaf{"StandbyEventCliInfo", detailInfo.StandbyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("cleanup-event-cli-info", types.YLeaf{"CleanupEventCliInfo", detailInfo.CleanupEventCliInfo})
    detailInfo.EntityData.Leafs.Append("band-ready-event-cli-info", types.YLeaf{"BandReadyEventCliInfo", detailInfo.BandReadyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("lr-event-cli-info", types.YLeaf{"LrEventCliInfo", detailInfo.LrEventCliInfo})
    detailInfo.EntityData.Leafs.Append("plane-ready-event-cli-info", types.YLeaf{"PlaneReadyEventCliInfo", detailInfo.PlaneReadyEventCliInfo})
    detailInfo.EntityData.Leafs.Append("mdr-is-done-cli-info", types.YLeaf{"MdrIsDoneCliInfo", detailInfo.MdrIsDoneCliInfo})

    detailInfo.EntityData.YListKeys = []string {}

    return &(detailInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo
// Process Verbose Info
type SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process Group. The type is string.
    ProcessGroup interface{}

    // Is respawn allowed?. The type is interface{} with range:
    // -2147483648..2147483647.
    RespawnAllowed interface{}

    // Wait for exit. The type is interface{} with range: -2147483648..2147483647.
    WaitForExit interface{}

    // Dynamic Tag. The type is interface{} with range: -2147483648..2147483647.
    DynamicTag interface{}

    // Forced stop. The type is interface{} with range: -2147483648..2147483647.
    ForcedStop interface{}

    // Critical Process. The type is interface{} with range:
    // -2147483648..2147483647.
    CriticalProcess interface{}

    // Hold. The type is interface{} with range: -2147483648..2147483647.
    Hold interface{}

    // Transient. The type is interface{} with range: -2147483648..2147483647.
    Transient interface{}

    // Tuple Cfgmgr. The type is interface{} with range: -2147483648..2147483647.
    TupleCfgmgr interface{}

    // Standby capable. The type is interface{} with range:
    // -2147483648..2147483647.
    StandbyCapable interface{}

    // EDM startup. The type is interface{} with range: -2147483648..2147483647.
    EdmStartup interface{}

    // Placement. The type is interface{} with range: -2147483648..2147483647.
    Placement interface{}

    // Skip Kill Notif. The type is interface{} with range:
    // -2147483648..2147483647.
    SkipKillNotif interface{}

    // Init process. The type is interface{} with range: -2147483648..2147483647.
    InitProc interface{}

    // Sysdb Event. The type is interface{} with range: -2147483648..2147483647.
    SysdbEvent interface{}

    // Level Started. The type is interface{} with range: -2147483648..2147483647.
    LevelStarted interface{}

    // Process available. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcAvail interface{}

    // Tuples Scanned. The type is interface{} with range:
    // -2147483648..2147483647.
    TuplesScanned interface{}

    // No checkpoint start. The type is interface{} with range:
    // -2147483648..2147483647.
    NoChkptStart interface{}

    // In Shut Down. The type is interface{} with range: -2147483648..2147483647.
    InShutDown interface{}

    // SM started. The type is interface{} with range: -2147483648..2147483647.
    SmStarted interface{}

    // Ignore on SC. The type is interface{} with range: -2147483648..2147483647.
    IgnoreOnSc interface{}

    // Ignore on EasyBake. The type is interface{} with range:
    // -2147483648..2147483647.
    IgnoreOnEasyBake interface{}

    // Pre init. The type is interface{} with range: -2147483648..2147483647.
    PreInit interface{}

    // EOI received. The type is interface{} with range: -2147483648..2147483647.
    EoiReceived interface{}

    // EOI Timeout. The type is interface{} with range: -2147483648..2147483647.
    EoiTimeout interface{}

    // Avail Timeout. The type is interface{} with range: -2147483648..2147483647.
    AvailTimeout interface{}

    // Reserved Memory. The type is interface{} with range:
    // -2147483648..2147483647.
    ReservedMemory interface{}

    // Allow Warned. The type is interface{} with range: -2147483648..2147483647.
    AllowWarned interface{}

    // Arg Change. The type is interface{} with range: -2147483648..2147483647.
    ArgChange interface{}

    // Restart on tuple. The type is interface{} with range:
    // -2147483648..2147483647.
    RestartOnTuple interface{}

    // Boot Hold. The type is interface{} with range: -2147483648..2147483647.
    BootHold interface{}

    // Reg Id. The type is interface{} with range: -2147483648..2147483647.
    RegId interface{}

    // Memory Limit. The type is interface{} with range: -2147483648..2147483647.
    MemoryLimit interface{}

    // Parent Job ID. The type is interface{} with range: -2147483648..2147483647.
    ParentJobId interface{}

    // Tuple Index. The type is interface{} with range: -2147483648..2147483647.
    TupleIndex interface{}

    // Dump Count. The type is interface{} with range: -2147483648..2147483647.
    DumpCount interface{}

    // Respawn Interval User. The type is interface{} with range:
    // -2147483648..2147483647.
    RespawnIntervalUser interface{}

    // Silent Restart Count. The type is interface{} with range:
    // -2147483648..2147483647.
    SilentRestartCount interface{}

    // Critical Tier. The type is interface{} with range: -2147483648..2147483647.
    CriticalTier interface{}

    // Exit Type. The type is interface{} with range: -2147483648..2147483647.
    ExitType interface{}

    // Init Timeout. The type is interface{} with range: -2147483648..2147483647.
    InitTimeout interface{}

    // Restart by Command. The type is interface{} with range:
    // -2147483648..2147483647.
    RestartByCmd interface{}

    // Boot Pref. The type is interface{} with range: -2147483648..2147483647.
    BootPref interface{}

    // Mdr Mbi proc. The type is interface{} with range: -2147483648..2147483647.
    MdrMbiProc interface{}

    // Mdr Non Mbi Kld. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrNonMbiKld interface{}

    // Mdr Mbi Kld. The type is interface{} with range: -2147483648..2147483647.
    MdrMbiKld interface{}

    // Mdr Shut Delay. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrShutDelay interface{}

    // Mdr Keep Thru. The type is interface{} with range: -2147483648..2147483647.
    MdrKeepThru interface{}

    // Mdr spoofer. The type is interface{} with range: -2147483648..2147483647.
    MdrSpoofer interface{}

    // Mdr spoofed. The type is interface{} with range: -2147483648..2147483647.
    MdrSpoofed interface{}

    // Mdr spoofed last. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrSpoofedLast interface{}

    // Mdr Spoofed Ready. The type is interface{} with range:
    // -2147483648..2147483647.
    MdrSpoofedReady interface{}

    // Mdr PCB Check. The type is interface{} with range: -2147483648..2147483647.
    MdrPcbCheck interface{}

    // Mdr Kill Tier. The type is interface{} with range: -2147483648..2147483647.
    MdrKillTier interface{}

    // Mdr kld. The type is interface{} with range: -2147483648..2147483647.
    MdrKld interface{}

    // Mdr Level. The type is interface{} with range: -2147483648..2147483647.
    MdrLevel interface{}

    // FM restart count. The type is interface{} with range:
    // -2147483648..2147483647.
    FmRestartCnt interface{}

    // Self Managed. The type is interface{} with range: -2147483648..2147483647.
    SelfManaged interface{}

    // Tuple. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_Tuple.
    Tuple []*SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_Tuple

    // Orig Tuple. The type is slice of
    // SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_OrigTuple.
    OrigTuple []*SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_OrigTuple
}

func (verboseInfo *SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo) GetEntityData() *types.CommonEntityData {
    verboseInfo.EntityData.YFilter = verboseInfo.YFilter
    verboseInfo.EntityData.YangName = "verbose-info"
    verboseInfo.EntityData.BundleName = "cisco_ios_xr"
    verboseInfo.EntityData.ParentYangName = "process-name-verbose"
    verboseInfo.EntityData.SegmentPath = "verbose-info"
    verboseInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-verboses/process-name-verbose/" + verboseInfo.EntityData.SegmentPath
    verboseInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    verboseInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    verboseInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    verboseInfo.EntityData.Children = types.NewOrderedMap()
    verboseInfo.EntityData.Children.Append("tuple", types.YChild{"Tuple", nil})
    for i := range verboseInfo.Tuple {
        types.SetYListKey(verboseInfo.Tuple[i], i)
        verboseInfo.EntityData.Children.Append(types.GetSegmentPath(verboseInfo.Tuple[i]), types.YChild{"Tuple", verboseInfo.Tuple[i]})
    }
    verboseInfo.EntityData.Children.Append("orig-tuple", types.YChild{"OrigTuple", nil})
    for i := range verboseInfo.OrigTuple {
        types.SetYListKey(verboseInfo.OrigTuple[i], i)
        verboseInfo.EntityData.Children.Append(types.GetSegmentPath(verboseInfo.OrigTuple[i]), types.YChild{"OrigTuple", verboseInfo.OrigTuple[i]})
    }
    verboseInfo.EntityData.Leafs = types.NewOrderedMap()
    verboseInfo.EntityData.Leafs.Append("process-group", types.YLeaf{"ProcessGroup", verboseInfo.ProcessGroup})
    verboseInfo.EntityData.Leafs.Append("respawn-allowed", types.YLeaf{"RespawnAllowed", verboseInfo.RespawnAllowed})
    verboseInfo.EntityData.Leafs.Append("wait-for-exit", types.YLeaf{"WaitForExit", verboseInfo.WaitForExit})
    verboseInfo.EntityData.Leafs.Append("dynamic-tag", types.YLeaf{"DynamicTag", verboseInfo.DynamicTag})
    verboseInfo.EntityData.Leafs.Append("forced-stop", types.YLeaf{"ForcedStop", verboseInfo.ForcedStop})
    verboseInfo.EntityData.Leafs.Append("critical-process", types.YLeaf{"CriticalProcess", verboseInfo.CriticalProcess})
    verboseInfo.EntityData.Leafs.Append("hold", types.YLeaf{"Hold", verboseInfo.Hold})
    verboseInfo.EntityData.Leafs.Append("transient", types.YLeaf{"Transient", verboseInfo.Transient})
    verboseInfo.EntityData.Leafs.Append("tuple-cfgmgr", types.YLeaf{"TupleCfgmgr", verboseInfo.TupleCfgmgr})
    verboseInfo.EntityData.Leafs.Append("standby-capable", types.YLeaf{"StandbyCapable", verboseInfo.StandbyCapable})
    verboseInfo.EntityData.Leafs.Append("edm-startup", types.YLeaf{"EdmStartup", verboseInfo.EdmStartup})
    verboseInfo.EntityData.Leafs.Append("placement", types.YLeaf{"Placement", verboseInfo.Placement})
    verboseInfo.EntityData.Leafs.Append("skip-kill-notif", types.YLeaf{"SkipKillNotif", verboseInfo.SkipKillNotif})
    verboseInfo.EntityData.Leafs.Append("init-proc", types.YLeaf{"InitProc", verboseInfo.InitProc})
    verboseInfo.EntityData.Leafs.Append("sysdb-event", types.YLeaf{"SysdbEvent", verboseInfo.SysdbEvent})
    verboseInfo.EntityData.Leafs.Append("level-started", types.YLeaf{"LevelStarted", verboseInfo.LevelStarted})
    verboseInfo.EntityData.Leafs.Append("proc-avail", types.YLeaf{"ProcAvail", verboseInfo.ProcAvail})
    verboseInfo.EntityData.Leafs.Append("tuples-scanned", types.YLeaf{"TuplesScanned", verboseInfo.TuplesScanned})
    verboseInfo.EntityData.Leafs.Append("no-chkpt-start", types.YLeaf{"NoChkptStart", verboseInfo.NoChkptStart})
    verboseInfo.EntityData.Leafs.Append("in-shut-down", types.YLeaf{"InShutDown", verboseInfo.InShutDown})
    verboseInfo.EntityData.Leafs.Append("sm-started", types.YLeaf{"SmStarted", verboseInfo.SmStarted})
    verboseInfo.EntityData.Leafs.Append("ignore-on-sc", types.YLeaf{"IgnoreOnSc", verboseInfo.IgnoreOnSc})
    verboseInfo.EntityData.Leafs.Append("ignore-on-easy-bake", types.YLeaf{"IgnoreOnEasyBake", verboseInfo.IgnoreOnEasyBake})
    verboseInfo.EntityData.Leafs.Append("pre-init", types.YLeaf{"PreInit", verboseInfo.PreInit})
    verboseInfo.EntityData.Leafs.Append("eoi-received", types.YLeaf{"EoiReceived", verboseInfo.EoiReceived})
    verboseInfo.EntityData.Leafs.Append("eoi-timeout", types.YLeaf{"EoiTimeout", verboseInfo.EoiTimeout})
    verboseInfo.EntityData.Leafs.Append("avail-timeout", types.YLeaf{"AvailTimeout", verboseInfo.AvailTimeout})
    verboseInfo.EntityData.Leafs.Append("reserved-memory", types.YLeaf{"ReservedMemory", verboseInfo.ReservedMemory})
    verboseInfo.EntityData.Leafs.Append("allow-warned", types.YLeaf{"AllowWarned", verboseInfo.AllowWarned})
    verboseInfo.EntityData.Leafs.Append("arg-change", types.YLeaf{"ArgChange", verboseInfo.ArgChange})
    verboseInfo.EntityData.Leafs.Append("restart-on-tuple", types.YLeaf{"RestartOnTuple", verboseInfo.RestartOnTuple})
    verboseInfo.EntityData.Leafs.Append("boot-hold", types.YLeaf{"BootHold", verboseInfo.BootHold})
    verboseInfo.EntityData.Leafs.Append("reg-id", types.YLeaf{"RegId", verboseInfo.RegId})
    verboseInfo.EntityData.Leafs.Append("memory-limit", types.YLeaf{"MemoryLimit", verboseInfo.MemoryLimit})
    verboseInfo.EntityData.Leafs.Append("parent-job-id", types.YLeaf{"ParentJobId", verboseInfo.ParentJobId})
    verboseInfo.EntityData.Leafs.Append("tuple-index", types.YLeaf{"TupleIndex", verboseInfo.TupleIndex})
    verboseInfo.EntityData.Leafs.Append("dump-count", types.YLeaf{"DumpCount", verboseInfo.DumpCount})
    verboseInfo.EntityData.Leafs.Append("respawn-interval-user", types.YLeaf{"RespawnIntervalUser", verboseInfo.RespawnIntervalUser})
    verboseInfo.EntityData.Leafs.Append("silent-restart-count", types.YLeaf{"SilentRestartCount", verboseInfo.SilentRestartCount})
    verboseInfo.EntityData.Leafs.Append("critical-tier", types.YLeaf{"CriticalTier", verboseInfo.CriticalTier})
    verboseInfo.EntityData.Leafs.Append("exit-type", types.YLeaf{"ExitType", verboseInfo.ExitType})
    verboseInfo.EntityData.Leafs.Append("init-timeout", types.YLeaf{"InitTimeout", verboseInfo.InitTimeout})
    verboseInfo.EntityData.Leafs.Append("restart-by-cmd", types.YLeaf{"RestartByCmd", verboseInfo.RestartByCmd})
    verboseInfo.EntityData.Leafs.Append("boot-pref", types.YLeaf{"BootPref", verboseInfo.BootPref})
    verboseInfo.EntityData.Leafs.Append("mdr-mbi-proc", types.YLeaf{"MdrMbiProc", verboseInfo.MdrMbiProc})
    verboseInfo.EntityData.Leafs.Append("mdr-non-mbi-kld", types.YLeaf{"MdrNonMbiKld", verboseInfo.MdrNonMbiKld})
    verboseInfo.EntityData.Leafs.Append("mdr-mbi-kld", types.YLeaf{"MdrMbiKld", verboseInfo.MdrMbiKld})
    verboseInfo.EntityData.Leafs.Append("mdr-shut-delay", types.YLeaf{"MdrShutDelay", verboseInfo.MdrShutDelay})
    verboseInfo.EntityData.Leafs.Append("mdr-keep-thru", types.YLeaf{"MdrKeepThru", verboseInfo.MdrKeepThru})
    verboseInfo.EntityData.Leafs.Append("mdr-spoofer", types.YLeaf{"MdrSpoofer", verboseInfo.MdrSpoofer})
    verboseInfo.EntityData.Leafs.Append("mdr-spoofed", types.YLeaf{"MdrSpoofed", verboseInfo.MdrSpoofed})
    verboseInfo.EntityData.Leafs.Append("mdr-spoofed-last", types.YLeaf{"MdrSpoofedLast", verboseInfo.MdrSpoofedLast})
    verboseInfo.EntityData.Leafs.Append("mdr-spoofed-ready", types.YLeaf{"MdrSpoofedReady", verboseInfo.MdrSpoofedReady})
    verboseInfo.EntityData.Leafs.Append("mdr-pcb-check", types.YLeaf{"MdrPcbCheck", verboseInfo.MdrPcbCheck})
    verboseInfo.EntityData.Leafs.Append("mdr-kill-tier", types.YLeaf{"MdrKillTier", verboseInfo.MdrKillTier})
    verboseInfo.EntityData.Leafs.Append("mdr-kld", types.YLeaf{"MdrKld", verboseInfo.MdrKld})
    verboseInfo.EntityData.Leafs.Append("mdr-level", types.YLeaf{"MdrLevel", verboseInfo.MdrLevel})
    verboseInfo.EntityData.Leafs.Append("fm-restart-cnt", types.YLeaf{"FmRestartCnt", verboseInfo.FmRestartCnt})
    verboseInfo.EntityData.Leafs.Append("self-managed", types.YLeaf{"SelfManaged", verboseInfo.SelfManaged})

    verboseInfo.EntityData.YListKeys = []string {}

    return &(verboseInfo.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_Tuple
// Tuple
type SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_Tuple struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (tuple *SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_Tuple) GetEntityData() *types.CommonEntityData {
    tuple.EntityData.YFilter = tuple.YFilter
    tuple.EntityData.YangName = "tuple"
    tuple.EntityData.BundleName = "cisco_ios_xr"
    tuple.EntityData.ParentYangName = "verbose-info"
    tuple.EntityData.SegmentPath = "tuple" + types.AddNoKeyToken(tuple)
    tuple.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-verboses/process-name-verbose/verbose-info/" + tuple.EntityData.SegmentPath
    tuple.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tuple.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tuple.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tuple.EntityData.Children = types.NewOrderedMap()
    tuple.EntityData.Leafs = types.NewOrderedMap()
    tuple.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", tuple.Tuple})

    tuple.EntityData.YListKeys = []string {}

    return &(tuple.EntityData)
}

// SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_OrigTuple
// Orig Tuple
type SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_OrigTuple struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (origTuple *SystemProcess_NodeTable_Node_Name_ProcessNameVerboses_ProcessNameVerbose_VerboseInfo_OrigTuple) GetEntityData() *types.CommonEntityData {
    origTuple.EntityData.YFilter = origTuple.YFilter
    origTuple.EntityData.YangName = "orig-tuple"
    origTuple.EntityData.BundleName = "cisco_ios_xr"
    origTuple.EntityData.ParentYangName = "verbose-info"
    origTuple.EntityData.SegmentPath = "orig-tuple" + types.AddNoKeyToken(origTuple)
    origTuple.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/name/process-name-verboses/process-name-verbose/verbose-info/" + origTuple.EntityData.SegmentPath
    origTuple.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    origTuple.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    origTuple.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    origTuple.EntityData.Children = types.NewOrderedMap()
    origTuple.EntityData.Leafs = types.NewOrderedMap()
    origTuple.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", origTuple.Tuple})

    origTuple.EntityData.YListKeys = []string {}

    return &(origTuple.EntityData)
}

// SystemProcess_NodeTable_Node_Jids
// Process job id information
type SystemProcess_NodeTable_Node_Jids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process <jid> information. The type is slice of
    // SystemProcess_NodeTable_Node_Jids_Jid.
    Jid []*SystemProcess_NodeTable_Node_Jids_Jid
}

func (jids *SystemProcess_NodeTable_Node_Jids) GetEntityData() *types.CommonEntityData {
    jids.EntityData.YFilter = jids.YFilter
    jids.EntityData.YangName = "jids"
    jids.EntityData.BundleName = "cisco_ios_xr"
    jids.EntityData.ParentYangName = "node"
    jids.EntityData.SegmentPath = "jids"
    jids.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + jids.EntityData.SegmentPath
    jids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    jids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    jids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    jids.EntityData.Children = types.NewOrderedMap()
    jids.EntityData.Children.Append("jid", types.YChild{"Jid", nil})
    for i := range jids.Jid {
        jids.EntityData.Children.Append(types.GetSegmentPath(jids.Jid[i]), types.YChild{"Jid", jids.Jid[i]})
    }
    jids.EntityData.Leafs = types.NewOrderedMap()

    jids.EntityData.YListKeys = []string {}

    return &(jids.EntityData)
}

// SystemProcess_NodeTable_Node_Jids_Jid
// Process <jid> information
type SystemProcess_NodeTable_Node_Jids_Jid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Job ID. The type is interface{} with range:
    // 0..4294967295.
    JobId interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    JobIdXr interface{}

    // PID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Process name. The type is string.
    ProcessName interface{}

    // Executable name or path. The type is string.
    Executable interface{}

    // Active Path. The type is string.
    ActivePath interface{}

    // Instance ID. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Args. The type is string.
    Args interface{}

    // Version ID. The type is string.
    VersionId interface{}

    // Respawn on/off. The type is string.
    Respawn interface{}

    // Respawn Count. The type is interface{} with range: -2147483648..2147483647.
    RespawnCount interface{}

    // Last Started. The type is string.
    LastStarted interface{}

    // Process State. The type is string.
    ProcessState interface{}

    // Last Exit status. The type is interface{} with range:
    // -2147483648..2147483647.
    LastExitStatus interface{}

    // Last Exit due to. The type is string.
    LastExitReason interface{}

    // Package State. The type is string.
    PackageState interface{}

    // Started on Config. The type is string.
    StartedOnConfig interface{}

    // Feature Name. The type is string.
    FeatureName interface{}

    // Tag. The type is string.
    Tag interface{}

    // Process Group. The type is string.
    Group interface{}

    // Core. The type is string.
    Core interface{}

    // Max core. The type is interface{} with range: -2147483648..2147483647.
    MaxCore interface{}

    // Level. The type is string.
    Level interface{}

    // Is mandatory?. The type is bool.
    Mandatory interface{}

    // Is admin mode process?. The type is bool.
    MaintModeProc interface{}

    // Placement State. The type is string.
    PlacementState interface{}

    // Startup Path. The type is string.
    StartUpPath interface{}

    // Memory Limit. The type is interface{} with range: 0..4294967295.
    MemoryLimit interface{}

    // Elapsed Ready. The type is string.
    Ready interface{}

    // Elapsed Available. The type is string.
    Available interface{}

    // Proces cpu time.
    ProcCpuTime SystemProcess_NodeTable_Node_Jids_Jid_ProcCpuTime

    // Registered Items. The type is slice of
    // SystemProcess_NodeTable_Node_Jids_Jid_RegisteredItem.
    RegisteredItem []*SystemProcess_NodeTable_Node_Jids_Jid_RegisteredItem
}

func (jid *SystemProcess_NodeTable_Node_Jids_Jid) GetEntityData() *types.CommonEntityData {
    jid.EntityData.YFilter = jid.YFilter
    jid.EntityData.YangName = "jid"
    jid.EntityData.BundleName = "cisco_ios_xr"
    jid.EntityData.ParentYangName = "jids"
    jid.EntityData.SegmentPath = "jid" + types.AddKeyToken(jid.JobId, "job-id")
    jid.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/jids/" + jid.EntityData.SegmentPath
    jid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    jid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    jid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    jid.EntityData.Children = types.NewOrderedMap()
    jid.EntityData.Children.Append("proc-cpu-time", types.YChild{"ProcCpuTime", &jid.ProcCpuTime})
    jid.EntityData.Children.Append("registered-item", types.YChild{"RegisteredItem", nil})
    for i := range jid.RegisteredItem {
        types.SetYListKey(jid.RegisteredItem[i], i)
        jid.EntityData.Children.Append(types.GetSegmentPath(jid.RegisteredItem[i]), types.YChild{"RegisteredItem", jid.RegisteredItem[i]})
    }
    jid.EntityData.Leafs = types.NewOrderedMap()
    jid.EntityData.Leafs.Append("job-id", types.YLeaf{"JobId", jid.JobId})
    jid.EntityData.Leafs.Append("job-id-xr", types.YLeaf{"JobIdXr", jid.JobIdXr})
    jid.EntityData.Leafs.Append("process-id", types.YLeaf{"ProcessId", jid.ProcessId})
    jid.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", jid.ProcessName})
    jid.EntityData.Leafs.Append("executable", types.YLeaf{"Executable", jid.Executable})
    jid.EntityData.Leafs.Append("active-path", types.YLeaf{"ActivePath", jid.ActivePath})
    jid.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", jid.InstanceId})
    jid.EntityData.Leafs.Append("args", types.YLeaf{"Args", jid.Args})
    jid.EntityData.Leafs.Append("version-id", types.YLeaf{"VersionId", jid.VersionId})
    jid.EntityData.Leafs.Append("respawn", types.YLeaf{"Respawn", jid.Respawn})
    jid.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", jid.RespawnCount})
    jid.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", jid.LastStarted})
    jid.EntityData.Leafs.Append("process-state", types.YLeaf{"ProcessState", jid.ProcessState})
    jid.EntityData.Leafs.Append("last-exit-status", types.YLeaf{"LastExitStatus", jid.LastExitStatus})
    jid.EntityData.Leafs.Append("last-exit-reason", types.YLeaf{"LastExitReason", jid.LastExitReason})
    jid.EntityData.Leafs.Append("package-state", types.YLeaf{"PackageState", jid.PackageState})
    jid.EntityData.Leafs.Append("started-on-config", types.YLeaf{"StartedOnConfig", jid.StartedOnConfig})
    jid.EntityData.Leafs.Append("feature-name", types.YLeaf{"FeatureName", jid.FeatureName})
    jid.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", jid.Tag})
    jid.EntityData.Leafs.Append("group", types.YLeaf{"Group", jid.Group})
    jid.EntityData.Leafs.Append("core", types.YLeaf{"Core", jid.Core})
    jid.EntityData.Leafs.Append("max-core", types.YLeaf{"MaxCore", jid.MaxCore})
    jid.EntityData.Leafs.Append("level", types.YLeaf{"Level", jid.Level})
    jid.EntityData.Leafs.Append("mandatory", types.YLeaf{"Mandatory", jid.Mandatory})
    jid.EntityData.Leafs.Append("maint-mode-proc", types.YLeaf{"MaintModeProc", jid.MaintModeProc})
    jid.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", jid.PlacementState})
    jid.EntityData.Leafs.Append("start-up-path", types.YLeaf{"StartUpPath", jid.StartUpPath})
    jid.EntityData.Leafs.Append("memory-limit", types.YLeaf{"MemoryLimit", jid.MemoryLimit})
    jid.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", jid.Ready})
    jid.EntityData.Leafs.Append("available", types.YLeaf{"Available", jid.Available})

    jid.EntityData.YListKeys = []string {"JobId"}

    return &(jid.EntityData)
}

// SystemProcess_NodeTable_Node_Jids_Jid_ProcCpuTime
// Proces cpu time
type SystemProcess_NodeTable_Node_Jids_Jid_ProcCpuTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // User time. The type is string.
    User interface{}

    // Kernel time. The type is string.
    System interface{}

    // Total time. The type is string.
    Total interface{}
}

func (procCpuTime *SystemProcess_NodeTable_Node_Jids_Jid_ProcCpuTime) GetEntityData() *types.CommonEntityData {
    procCpuTime.EntityData.YFilter = procCpuTime.YFilter
    procCpuTime.EntityData.YangName = "proc-cpu-time"
    procCpuTime.EntityData.BundleName = "cisco_ios_xr"
    procCpuTime.EntityData.ParentYangName = "jid"
    procCpuTime.EntityData.SegmentPath = "proc-cpu-time"
    procCpuTime.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/jids/jid/" + procCpuTime.EntityData.SegmentPath
    procCpuTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procCpuTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procCpuTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procCpuTime.EntityData.Children = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs = types.NewOrderedMap()
    procCpuTime.EntityData.Leafs.Append("user", types.YLeaf{"User", procCpuTime.User})
    procCpuTime.EntityData.Leafs.Append("system", types.YLeaf{"System", procCpuTime.System})
    procCpuTime.EntityData.Leafs.Append("total", types.YLeaf{"Total", procCpuTime.Total})

    procCpuTime.EntityData.YListKeys = []string {}

    return &(procCpuTime.EntityData)
}

// SystemProcess_NodeTable_Node_Jids_Jid_RegisteredItem
// Registered Items
type SystemProcess_NodeTable_Node_Jids_Jid_RegisteredItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tuple. The type is string.
    Tuple interface{}
}

func (registeredItem *SystemProcess_NodeTable_Node_Jids_Jid_RegisteredItem) GetEntityData() *types.CommonEntityData {
    registeredItem.EntityData.YFilter = registeredItem.YFilter
    registeredItem.EntityData.YangName = "registered-item"
    registeredItem.EntityData.BundleName = "cisco_ios_xr"
    registeredItem.EntityData.ParentYangName = "jid"
    registeredItem.EntityData.SegmentPath = "registered-item" + types.AddNoKeyToken(registeredItem)
    registeredItem.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/jids/jid/" + registeredItem.EntityData.SegmentPath
    registeredItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    registeredItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    registeredItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    registeredItem.EntityData.Children = types.NewOrderedMap()
    registeredItem.EntityData.Leafs = types.NewOrderedMap()
    registeredItem.EntityData.Leafs.Append("tuple", types.YLeaf{"Tuple", registeredItem.Tuple})

    registeredItem.EntityData.YListKeys = []string {}

    return &(registeredItem.EntityData)
}

// SystemProcess_NodeTable_Node_Dynamic
// Process Dynamic information
type SystemProcess_NodeTable_Node_Dynamic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of processes. The type is interface{} with range: 0..4294967295.
    ProcessCount interface{}

    // Array of processes. The type is slice of
    // SystemProcess_NodeTable_Node_Dynamic_Process.
    Process []*SystemProcess_NodeTable_Node_Dynamic_Process
}

func (dynamic *SystemProcess_NodeTable_Node_Dynamic) GetEntityData() *types.CommonEntityData {
    dynamic.EntityData.YFilter = dynamic.YFilter
    dynamic.EntityData.YangName = "dynamic"
    dynamic.EntityData.BundleName = "cisco_ios_xr"
    dynamic.EntityData.ParentYangName = "node"
    dynamic.EntityData.SegmentPath = "dynamic"
    dynamic.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + dynamic.EntityData.SegmentPath
    dynamic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamic.EntityData.Children = types.NewOrderedMap()
    dynamic.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range dynamic.Process {
        types.SetYListKey(dynamic.Process[i], i)
        dynamic.EntityData.Children.Append(types.GetSegmentPath(dynamic.Process[i]), types.YChild{"Process", dynamic.Process[i]})
    }
    dynamic.EntityData.Leafs = types.NewOrderedMap()
    dynamic.EntityData.Leafs.Append("process-count", types.YLeaf{"ProcessCount", dynamic.ProcessCount})

    dynamic.EntityData.YListKeys = []string {}

    return &(dynamic.EntityData)
}

// SystemProcess_NodeTable_Node_Dynamic_Process
// Array of processes
type SystemProcess_NodeTable_Node_Dynamic_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Process name. The type is string.
    Name interface{}

    // Instance ID. The type is interface{} with range: 0..4294967295.
    InstanceId interface{}

    // Arguments. The type is string.
    Args interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Process state. The type is ProcessState.
    State interface{}

    // Date and time of process last started. The type is string.
    LastStarted interface{}

    // Respawn count. The type is interface{} with range: 0..255.
    RespawnCount interface{}

    // Placement state. The type is PlacementState.
    PlacementState interface{}

    // Is process mandatory?. The type is bool.
    IsMandatory interface{}

    // Is maintenance mode?. The type is bool.
    IsMaintenance interface{}
}

func (process *SystemProcess_NodeTable_Node_Dynamic_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "dynamic"
    process.EntityData.SegmentPath = "process" + types.AddNoKeyToken(process)
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/dynamic/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("name", types.YLeaf{"Name", process.Name})
    process.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", process.InstanceId})
    process.EntityData.Leafs.Append("args", types.YLeaf{"Args", process.Args})
    process.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", process.Jid})
    process.EntityData.Leafs.Append("state", types.YLeaf{"State", process.State})
    process.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", process.LastStarted})
    process.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", process.RespawnCount})
    process.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", process.PlacementState})
    process.EntityData.Leafs.Append("is-mandatory", types.YLeaf{"IsMandatory", process.IsMandatory})
    process.EntityData.Leafs.Append("is-maintenance", types.YLeaf{"IsMaintenance", process.IsMaintenance})

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// SystemProcess_NodeTable_Node_BootStalled
// Process Boot Stalled information
type SystemProcess_NodeTable_Node_BootStalled struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Spawn status of the processes. The type is string.
    SpawnStatus interface{}

    // Boot hold information of the processes. The type is slice of
    // SystemProcess_NodeTable_Node_BootStalled_BootHold.
    BootHold []*SystemProcess_NodeTable_Node_BootStalled_BootHold
}

func (bootStalled *SystemProcess_NodeTable_Node_BootStalled) GetEntityData() *types.CommonEntityData {
    bootStalled.EntityData.YFilter = bootStalled.YFilter
    bootStalled.EntityData.YangName = "boot-stalled"
    bootStalled.EntityData.BundleName = "cisco_ios_xr"
    bootStalled.EntityData.ParentYangName = "node"
    bootStalled.EntityData.SegmentPath = "boot-stalled"
    bootStalled.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + bootStalled.EntityData.SegmentPath
    bootStalled.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootStalled.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootStalled.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootStalled.EntityData.Children = types.NewOrderedMap()
    bootStalled.EntityData.Children.Append("boot-hold", types.YChild{"BootHold", nil})
    for i := range bootStalled.BootHold {
        types.SetYListKey(bootStalled.BootHold[i], i)
        bootStalled.EntityData.Children.Append(types.GetSegmentPath(bootStalled.BootHold[i]), types.YChild{"BootHold", bootStalled.BootHold[i]})
    }
    bootStalled.EntityData.Leafs = types.NewOrderedMap()
    bootStalled.EntityData.Leafs.Append("spawn-status", types.YLeaf{"SpawnStatus", bootStalled.SpawnStatus})

    bootStalled.EntityData.YListKeys = []string {}

    return &(bootStalled.EntityData)
}

// SystemProcess_NodeTable_Node_BootStalled_BootHold
// Boot hold information of the processes
type SystemProcess_NodeTable_Node_BootStalled_BootHold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Processs name. The type is string.
    BootHeldByName interface{}

    // Instance Id. The type is interface{} with range: 0..4294967295.
    InstanceId interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    Jid interface{}
}

func (bootHold *SystemProcess_NodeTable_Node_BootStalled_BootHold) GetEntityData() *types.CommonEntityData {
    bootHold.EntityData.YFilter = bootHold.YFilter
    bootHold.EntityData.YangName = "boot-hold"
    bootHold.EntityData.BundleName = "cisco_ios_xr"
    bootHold.EntityData.ParentYangName = "boot-stalled"
    bootHold.EntityData.SegmentPath = "boot-hold" + types.AddNoKeyToken(bootHold)
    bootHold.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/boot-stalled/" + bootHold.EntityData.SegmentPath
    bootHold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootHold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootHold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootHold.EntityData.Children = types.NewOrderedMap()
    bootHold.EntityData.Leafs = types.NewOrderedMap()
    bootHold.EntityData.Leafs.Append("boot-held-by-name", types.YLeaf{"BootHeldByName", bootHold.BootHeldByName})
    bootHold.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", bootHold.InstanceId})
    bootHold.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", bootHold.Jid})

    bootHold.EntityData.YListKeys = []string {}

    return &(bootHold.EntityData)
}

// SystemProcess_NodeTable_Node_Processes
// Process all information
type SystemProcess_NodeTable_Node_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of processes. The type is interface{} with range: 0..4294967295.
    ProcessCount interface{}

    // Array of processes. The type is slice of
    // SystemProcess_NodeTable_Node_Processes_Process.
    Process []*SystemProcess_NodeTable_Node_Processes_Process
}

func (processes *SystemProcess_NodeTable_Node_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xr"
    processes.EntityData.ParentYangName = "node"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + processes.EntityData.SegmentPath
    processes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processes.EntityData.Children = types.NewOrderedMap()
    processes.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range processes.Process {
        types.SetYListKey(processes.Process[i], i)
        processes.EntityData.Children.Append(types.GetSegmentPath(processes.Process[i]), types.YChild{"Process", processes.Process[i]})
    }
    processes.EntityData.Leafs = types.NewOrderedMap()
    processes.EntityData.Leafs.Append("process-count", types.YLeaf{"ProcessCount", processes.ProcessCount})

    processes.EntityData.YListKeys = []string {}

    return &(processes.EntityData)
}

// SystemProcess_NodeTable_Node_Processes_Process
// Array of processes
type SystemProcess_NodeTable_Node_Processes_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Process name. The type is string.
    Name interface{}

    // Instance ID. The type is interface{} with range: 0..4294967295.
    InstanceId interface{}

    // Arguments. The type is string.
    Args interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Process state. The type is ProcessState.
    State interface{}

    // Date and time of process last started. The type is string.
    LastStarted interface{}

    // Respawn count. The type is interface{} with range: 0..255.
    RespawnCount interface{}

    // Placement state. The type is PlacementState.
    PlacementState interface{}

    // Is process mandatory?. The type is bool.
    IsMandatory interface{}

    // Is maintenance mode?. The type is bool.
    IsMaintenance interface{}
}

func (process *SystemProcess_NodeTable_Node_Processes_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "processes"
    process.EntityData.SegmentPath = "process" + types.AddNoKeyToken(process)
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/processes/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("name", types.YLeaf{"Name", process.Name})
    process.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", process.InstanceId})
    process.EntityData.Leafs.Append("args", types.YLeaf{"Args", process.Args})
    process.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", process.Jid})
    process.EntityData.Leafs.Append("state", types.YLeaf{"State", process.State})
    process.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", process.LastStarted})
    process.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", process.RespawnCount})
    process.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", process.PlacementState})
    process.EntityData.Leafs.Append("is-mandatory", types.YLeaf{"IsMandatory", process.IsMandatory})
    process.EntityData.Leafs.Append("is-maintenance", types.YLeaf{"IsMaintenance", process.IsMaintenance})

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// SystemProcess_NodeTable_Node_Startup
// Process Startup information
type SystemProcess_NodeTable_Node_Startup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of processes. The type is interface{} with range: 0..4294967295.
    ProcessCount interface{}

    // Array of processes. The type is slice of
    // SystemProcess_NodeTable_Node_Startup_Process.
    Process []*SystemProcess_NodeTable_Node_Startup_Process
}

func (startup *SystemProcess_NodeTable_Node_Startup) GetEntityData() *types.CommonEntityData {
    startup.EntityData.YFilter = startup.YFilter
    startup.EntityData.YangName = "startup"
    startup.EntityData.BundleName = "cisco_ios_xr"
    startup.EntityData.ParentYangName = "node"
    startup.EntityData.SegmentPath = "startup"
    startup.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + startup.EntityData.SegmentPath
    startup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    startup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    startup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    startup.EntityData.Children = types.NewOrderedMap()
    startup.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range startup.Process {
        types.SetYListKey(startup.Process[i], i)
        startup.EntityData.Children.Append(types.GetSegmentPath(startup.Process[i]), types.YChild{"Process", startup.Process[i]})
    }
    startup.EntityData.Leafs = types.NewOrderedMap()
    startup.EntityData.Leafs.Append("process-count", types.YLeaf{"ProcessCount", startup.ProcessCount})

    startup.EntityData.YListKeys = []string {}

    return &(startup.EntityData)
}

// SystemProcess_NodeTable_Node_Startup_Process
// Array of processes
type SystemProcess_NodeTable_Node_Startup_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Process name. The type is string.
    Name interface{}

    // Instance ID. The type is interface{} with range: 0..4294967295.
    InstanceId interface{}

    // Arguments. The type is string.
    Args interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Process state. The type is ProcessState.
    State interface{}

    // Date and time of process last started. The type is string.
    LastStarted interface{}

    // Respawn count. The type is interface{} with range: 0..255.
    RespawnCount interface{}

    // Placement state. The type is PlacementState.
    PlacementState interface{}

    // Is process mandatory?. The type is bool.
    IsMandatory interface{}

    // Is maintenance mode?. The type is bool.
    IsMaintenance interface{}
}

func (process *SystemProcess_NodeTable_Node_Startup_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "startup"
    process.EntityData.SegmentPath = "process" + types.AddNoKeyToken(process)
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/startup/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("name", types.YLeaf{"Name", process.Name})
    process.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", process.InstanceId})
    process.EntityData.Leafs.Append("args", types.YLeaf{"Args", process.Args})
    process.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", process.Jid})
    process.EntityData.Leafs.Append("state", types.YLeaf{"State", process.State})
    process.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", process.LastStarted})
    process.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", process.RespawnCount})
    process.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", process.PlacementState})
    process.EntityData.Leafs.Append("is-mandatory", types.YLeaf{"IsMandatory", process.IsMandatory})
    process.EntityData.Leafs.Append("is-maintenance", types.YLeaf{"IsMaintenance", process.IsMaintenance})

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// SystemProcess_NodeTable_Node_Mandatory
// Mandatory Process information
type SystemProcess_NodeTable_Node_Mandatory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of processes. The type is interface{} with range: 0..4294967295.
    ProcessCount interface{}

    // Array of processes. The type is slice of
    // SystemProcess_NodeTable_Node_Mandatory_Process.
    Process []*SystemProcess_NodeTable_Node_Mandatory_Process
}

func (mandatory *SystemProcess_NodeTable_Node_Mandatory) GetEntityData() *types.CommonEntityData {
    mandatory.EntityData.YFilter = mandatory.YFilter
    mandatory.EntityData.YangName = "mandatory"
    mandatory.EntityData.BundleName = "cisco_ios_xr"
    mandatory.EntityData.ParentYangName = "node"
    mandatory.EntityData.SegmentPath = "mandatory"
    mandatory.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + mandatory.EntityData.SegmentPath
    mandatory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mandatory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mandatory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mandatory.EntityData.Children = types.NewOrderedMap()
    mandatory.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range mandatory.Process {
        types.SetYListKey(mandatory.Process[i], i)
        mandatory.EntityData.Children.Append(types.GetSegmentPath(mandatory.Process[i]), types.YChild{"Process", mandatory.Process[i]})
    }
    mandatory.EntityData.Leafs = types.NewOrderedMap()
    mandatory.EntityData.Leafs.Append("process-count", types.YLeaf{"ProcessCount", mandatory.ProcessCount})

    mandatory.EntityData.YListKeys = []string {}

    return &(mandatory.EntityData)
}

// SystemProcess_NodeTable_Node_Mandatory_Process
// Array of processes
type SystemProcess_NodeTable_Node_Mandatory_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Process name. The type is string.
    Name interface{}

    // Instance ID. The type is interface{} with range: 0..4294967295.
    InstanceId interface{}

    // Arguments. The type is string.
    Args interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Process state. The type is ProcessState.
    State interface{}

    // Date and time of process last started. The type is string.
    LastStarted interface{}

    // Respawn count. The type is interface{} with range: 0..255.
    RespawnCount interface{}

    // Placement state. The type is PlacementState.
    PlacementState interface{}

    // Is process mandatory?. The type is bool.
    IsMandatory interface{}

    // Is maintenance mode?. The type is bool.
    IsMaintenance interface{}
}

func (process *SystemProcess_NodeTable_Node_Mandatory_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "mandatory"
    process.EntityData.SegmentPath = "process" + types.AddNoKeyToken(process)
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/mandatory/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("name", types.YLeaf{"Name", process.Name})
    process.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", process.InstanceId})
    process.EntityData.Leafs.Append("args", types.YLeaf{"Args", process.Args})
    process.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", process.Jid})
    process.EntityData.Leafs.Append("state", types.YLeaf{"State", process.State})
    process.EntityData.Leafs.Append("last-started", types.YLeaf{"LastStarted", process.LastStarted})
    process.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", process.RespawnCount})
    process.EntityData.Leafs.Append("placement-state", types.YLeaf{"PlacementState", process.PlacementState})
    process.EntityData.Leafs.Append("is-mandatory", types.YLeaf{"IsMandatory", process.IsMandatory})
    process.EntityData.Leafs.Append("is-maintenance", types.YLeaf{"IsMaintenance", process.IsMaintenance})

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// SystemProcess_NodeTable_Node_Abort
// Process Abort information
type SystemProcess_NodeTable_Node_Abort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Aborted Processes. The type is interface{} with range:
    // 0..4294967295.
    ProcessAbortCount interface{}

    // Array of aborted processes. The type is slice of
    // SystemProcess_NodeTable_Node_Abort_Process.
    Process []*SystemProcess_NodeTable_Node_Abort_Process
}

func (abort *SystemProcess_NodeTable_Node_Abort) GetEntityData() *types.CommonEntityData {
    abort.EntityData.YFilter = abort.YFilter
    abort.EntityData.YangName = "abort"
    abort.EntityData.BundleName = "cisco_ios_xr"
    abort.EntityData.ParentYangName = "node"
    abort.EntityData.SegmentPath = "abort"
    abort.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + abort.EntityData.SegmentPath
    abort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    abort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    abort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    abort.EntityData.Children = types.NewOrderedMap()
    abort.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range abort.Process {
        types.SetYListKey(abort.Process[i], i)
        abort.EntityData.Children.Append(types.GetSegmentPath(abort.Process[i]), types.YChild{"Process", abort.Process[i]})
    }
    abort.EntityData.Leafs = types.NewOrderedMap()
    abort.EntityData.Leafs.Append("process-abort-count", types.YLeaf{"ProcessAbortCount", abort.ProcessAbortCount})

    abort.EntityData.YListKeys = []string {}

    return &(abort.EntityData)
}

// SystemProcess_NodeTable_Node_Abort_Process
// Array of aborted processes
type SystemProcess_NodeTable_Node_Abort_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Process name. The type is string.
    Name interface{}

    // Date and time of process abort. The type is string.
    Timebuf interface{}

    // Job ID. The type is interface{} with range: 0..4294967295.
    JobId interface{}

    // Respawn information. The type is string.
    IsRespawned interface{}
}

func (process *SystemProcess_NodeTable_Node_Abort_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "abort"
    process.EntityData.SegmentPath = "process" + types.AddNoKeyToken(process)
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/abort/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("name", types.YLeaf{"Name", process.Name})
    process.EntityData.Leafs.Append("timebuf", types.YLeaf{"Timebuf", process.Timebuf})
    process.EntityData.Leafs.Append("job-id", types.YLeaf{"JobId", process.JobId})
    process.EntityData.Leafs.Append("is-respawned", types.YLeaf{"IsRespawned", process.IsRespawned})

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// SystemProcess_NodeTable_Node_Failover
// Process Failover information
type SystemProcess_NodeTable_Node_Failover struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Failover log message. The type is string.
    FailoverLog interface{}

    // Critical Failover Elapsed Time. The type is string.
    CriticalFailoverElapsedTime interface{}

    // Last process started. The type is string.
    LastProcessStarted interface{}

    // Primary failover elapsed time. The type is string.
    PrimaryFailoverElapsedTime interface{}

    // Last primary process started. The type is string.
    LastPrimaryProcStarted interface{}

    // Standby Band statistics. The type is slice of
    // SystemProcess_NodeTable_Node_Failover_StandbyBandStatistic.
    StandbyBandStatistic []*SystemProcess_NodeTable_Node_Failover_StandbyBandStatistic

    // Active Band statistics. The type is slice of
    // SystemProcess_NodeTable_Node_Failover_ActiveBandStatistic.
    ActiveBandStatistic []*SystemProcess_NodeTable_Node_Failover_ActiveBandStatistic

    // List of booted process as per avail time. The type is slice of
    // SystemProcess_NodeTable_Node_Failover_ActiveTsBootProc.
    ActiveTsBootProc []*SystemProcess_NodeTable_Node_Failover_ActiveTsBootProc

    // List of booted processes per start time. The type is slice of
    // SystemProcess_NodeTable_Node_Failover_StartTsBootProc.
    StartTsBootProc []*SystemProcess_NodeTable_Node_Failover_StartTsBootProc

    // Primary Band statistics. The type is slice of
    // SystemProcess_NodeTable_Node_Failover_PrimaryBandStatistic.
    PrimaryBandStatistic []*SystemProcess_NodeTable_Node_Failover_PrimaryBandStatistic

    // List of booted processes per primary time. The type is slice of
    // SystemProcess_NodeTable_Node_Failover_PrimaryTsBootProc.
    PrimaryTsBootProc []*SystemProcess_NodeTable_Node_Failover_PrimaryTsBootProc

    // List of booted process per primary start time. The type is slice of
    // SystemProcess_NodeTable_Node_Failover_PrimaryStartTsBootProc.
    PrimaryStartTsBootProc []*SystemProcess_NodeTable_Node_Failover_PrimaryStartTsBootProc
}

func (failover *SystemProcess_NodeTable_Node_Failover) GetEntityData() *types.CommonEntityData {
    failover.EntityData.YFilter = failover.YFilter
    failover.EntityData.YangName = "failover"
    failover.EntityData.BundleName = "cisco_ios_xr"
    failover.EntityData.ParentYangName = "node"
    failover.EntityData.SegmentPath = "failover"
    failover.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + failover.EntityData.SegmentPath
    failover.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    failover.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    failover.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    failover.EntityData.Children = types.NewOrderedMap()
    failover.EntityData.Children.Append("standby-band-statistic", types.YChild{"StandbyBandStatistic", nil})
    for i := range failover.StandbyBandStatistic {
        types.SetYListKey(failover.StandbyBandStatistic[i], i)
        failover.EntityData.Children.Append(types.GetSegmentPath(failover.StandbyBandStatistic[i]), types.YChild{"StandbyBandStatistic", failover.StandbyBandStatistic[i]})
    }
    failover.EntityData.Children.Append("active-band-statistic", types.YChild{"ActiveBandStatistic", nil})
    for i := range failover.ActiveBandStatistic {
        types.SetYListKey(failover.ActiveBandStatistic[i], i)
        failover.EntityData.Children.Append(types.GetSegmentPath(failover.ActiveBandStatistic[i]), types.YChild{"ActiveBandStatistic", failover.ActiveBandStatistic[i]})
    }
    failover.EntityData.Children.Append("active-ts-boot-proc", types.YChild{"ActiveTsBootProc", nil})
    for i := range failover.ActiveTsBootProc {
        types.SetYListKey(failover.ActiveTsBootProc[i], i)
        failover.EntityData.Children.Append(types.GetSegmentPath(failover.ActiveTsBootProc[i]), types.YChild{"ActiveTsBootProc", failover.ActiveTsBootProc[i]})
    }
    failover.EntityData.Children.Append("start-ts-boot-proc", types.YChild{"StartTsBootProc", nil})
    for i := range failover.StartTsBootProc {
        types.SetYListKey(failover.StartTsBootProc[i], i)
        failover.EntityData.Children.Append(types.GetSegmentPath(failover.StartTsBootProc[i]), types.YChild{"StartTsBootProc", failover.StartTsBootProc[i]})
    }
    failover.EntityData.Children.Append("primary-band-statistic", types.YChild{"PrimaryBandStatistic", nil})
    for i := range failover.PrimaryBandStatistic {
        types.SetYListKey(failover.PrimaryBandStatistic[i], i)
        failover.EntityData.Children.Append(types.GetSegmentPath(failover.PrimaryBandStatistic[i]), types.YChild{"PrimaryBandStatistic", failover.PrimaryBandStatistic[i]})
    }
    failover.EntityData.Children.Append("primary-ts-boot-proc", types.YChild{"PrimaryTsBootProc", nil})
    for i := range failover.PrimaryTsBootProc {
        types.SetYListKey(failover.PrimaryTsBootProc[i], i)
        failover.EntityData.Children.Append(types.GetSegmentPath(failover.PrimaryTsBootProc[i]), types.YChild{"PrimaryTsBootProc", failover.PrimaryTsBootProc[i]})
    }
    failover.EntityData.Children.Append("primary-start-ts-boot-proc", types.YChild{"PrimaryStartTsBootProc", nil})
    for i := range failover.PrimaryStartTsBootProc {
        types.SetYListKey(failover.PrimaryStartTsBootProc[i], i)
        failover.EntityData.Children.Append(types.GetSegmentPath(failover.PrimaryStartTsBootProc[i]), types.YChild{"PrimaryStartTsBootProc", failover.PrimaryStartTsBootProc[i]})
    }
    failover.EntityData.Leafs = types.NewOrderedMap()
    failover.EntityData.Leafs.Append("failover-log", types.YLeaf{"FailoverLog", failover.FailoverLog})
    failover.EntityData.Leafs.Append("critical-failover-elapsed-time", types.YLeaf{"CriticalFailoverElapsedTime", failover.CriticalFailoverElapsedTime})
    failover.EntityData.Leafs.Append("last-process-started", types.YLeaf{"LastProcessStarted", failover.LastProcessStarted})
    failover.EntityData.Leafs.Append("primary-failover-elapsed-time", types.YLeaf{"PrimaryFailoverElapsedTime", failover.PrimaryFailoverElapsedTime})
    failover.EntityData.Leafs.Append("last-primary-proc-started", types.YLeaf{"LastPrimaryProcStarted", failover.LastPrimaryProcStarted})

    failover.EntityData.YListKeys = []string {}

    return &(failover.EntityData)
}

// SystemProcess_NodeTable_Node_Failover_StandbyBandStatistic
// Standby Band statistics
type SystemProcess_NodeTable_Node_Failover_StandbyBandStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Level. The type is string.
    Level interface{}

    // Band Name. The type is string.
    BandName interface{}

    // Band finish time. The type is string.
    BandFinishTime interface{}

    // Band time. The type is string.
    BandTime interface{}

    // Finish Time. The type is string.
    FinishTime interface{}

    // Idle Percentage. The type is string. Units are percentage.
    IdlePercentage interface{}

    // Jid. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Ready Time. The type is string.
    ReadyTime interface{}

    // Last Process Name. The type is string.
    LastProcess interface{}
}

func (standbyBandStatistic *SystemProcess_NodeTable_Node_Failover_StandbyBandStatistic) GetEntityData() *types.CommonEntityData {
    standbyBandStatistic.EntityData.YFilter = standbyBandStatistic.YFilter
    standbyBandStatistic.EntityData.YangName = "standby-band-statistic"
    standbyBandStatistic.EntityData.BundleName = "cisco_ios_xr"
    standbyBandStatistic.EntityData.ParentYangName = "failover"
    standbyBandStatistic.EntityData.SegmentPath = "standby-band-statistic" + types.AddNoKeyToken(standbyBandStatistic)
    standbyBandStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/failover/" + standbyBandStatistic.EntityData.SegmentPath
    standbyBandStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyBandStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyBandStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyBandStatistic.EntityData.Children = types.NewOrderedMap()
    standbyBandStatistic.EntityData.Leafs = types.NewOrderedMap()
    standbyBandStatistic.EntityData.Leafs.Append("level", types.YLeaf{"Level", standbyBandStatistic.Level})
    standbyBandStatistic.EntityData.Leafs.Append("band-name", types.YLeaf{"BandName", standbyBandStatistic.BandName})
    standbyBandStatistic.EntityData.Leafs.Append("band-finish-time", types.YLeaf{"BandFinishTime", standbyBandStatistic.BandFinishTime})
    standbyBandStatistic.EntityData.Leafs.Append("band-time", types.YLeaf{"BandTime", standbyBandStatistic.BandTime})
    standbyBandStatistic.EntityData.Leafs.Append("finish-time", types.YLeaf{"FinishTime", standbyBandStatistic.FinishTime})
    standbyBandStatistic.EntityData.Leafs.Append("idle-percentage", types.YLeaf{"IdlePercentage", standbyBandStatistic.IdlePercentage})
    standbyBandStatistic.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", standbyBandStatistic.Jid})
    standbyBandStatistic.EntityData.Leafs.Append("ready-time", types.YLeaf{"ReadyTime", standbyBandStatistic.ReadyTime})
    standbyBandStatistic.EntityData.Leafs.Append("last-process", types.YLeaf{"LastProcess", standbyBandStatistic.LastProcess})

    standbyBandStatistic.EntityData.YListKeys = []string {}

    return &(standbyBandStatistic.EntityData)
}

// SystemProcess_NodeTable_Node_Failover_ActiveBandStatistic
// Active Band statistics
type SystemProcess_NodeTable_Node_Failover_ActiveBandStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Level. The type is string.
    Level interface{}

    // Band Name. The type is string.
    BandName interface{}

    // Band finish time. The type is string.
    BandFinishTime interface{}

    // Band time. The type is string.
    BandTime interface{}

    // Finish Time. The type is string.
    FinishTime interface{}

    // Idle Percentage. The type is string. Units are percentage.
    IdlePercentage interface{}

    // Jid. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Ready Time. The type is string.
    ReadyTime interface{}

    // Last Process Name. The type is string.
    LastProcess interface{}
}

func (activeBandStatistic *SystemProcess_NodeTable_Node_Failover_ActiveBandStatistic) GetEntityData() *types.CommonEntityData {
    activeBandStatistic.EntityData.YFilter = activeBandStatistic.YFilter
    activeBandStatistic.EntityData.YangName = "active-band-statistic"
    activeBandStatistic.EntityData.BundleName = "cisco_ios_xr"
    activeBandStatistic.EntityData.ParentYangName = "failover"
    activeBandStatistic.EntityData.SegmentPath = "active-band-statistic" + types.AddNoKeyToken(activeBandStatistic)
    activeBandStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/failover/" + activeBandStatistic.EntityData.SegmentPath
    activeBandStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeBandStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeBandStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeBandStatistic.EntityData.Children = types.NewOrderedMap()
    activeBandStatistic.EntityData.Leafs = types.NewOrderedMap()
    activeBandStatistic.EntityData.Leafs.Append("level", types.YLeaf{"Level", activeBandStatistic.Level})
    activeBandStatistic.EntityData.Leafs.Append("band-name", types.YLeaf{"BandName", activeBandStatistic.BandName})
    activeBandStatistic.EntityData.Leafs.Append("band-finish-time", types.YLeaf{"BandFinishTime", activeBandStatistic.BandFinishTime})
    activeBandStatistic.EntityData.Leafs.Append("band-time", types.YLeaf{"BandTime", activeBandStatistic.BandTime})
    activeBandStatistic.EntityData.Leafs.Append("finish-time", types.YLeaf{"FinishTime", activeBandStatistic.FinishTime})
    activeBandStatistic.EntityData.Leafs.Append("idle-percentage", types.YLeaf{"IdlePercentage", activeBandStatistic.IdlePercentage})
    activeBandStatistic.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", activeBandStatistic.Jid})
    activeBandStatistic.EntityData.Leafs.Append("ready-time", types.YLeaf{"ReadyTime", activeBandStatistic.ReadyTime})
    activeBandStatistic.EntityData.Leafs.Append("last-process", types.YLeaf{"LastProcess", activeBandStatistic.LastProcess})

    activeBandStatistic.EntityData.YListKeys = []string {}

    return &(activeBandStatistic.EntityData)
}

// SystemProcess_NodeTable_Node_Failover_ActiveTsBootProc
// List of booted process as per avail time
type SystemProcess_NodeTable_Node_Failover_ActiveTsBootProc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Active Time Stamp. The type is string.
    ActiveTimeStamp interface{}

    // Go Active time stamp. The type is string.
    GoActive interface{}

    // Level. The type is string.
    Level interface{}

    // Band Name. The type is string.
    BandName interface{}

    // Job Id. The type is interface{} with range: -2147483648..2147483647.
    JobId interface{}

    // Instance Id. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Avail Time Stamp. The type is string.
    AvailTimeStamp interface{}

    // Time since Avail. The type is string.
    Avail interface{}

    // Is Avail timeout. The type is bool.
    IsAvailTimeout interface{}

    // Process Name. The type is string.
    ProcessName interface{}
}

func (activeTsBootProc *SystemProcess_NodeTable_Node_Failover_ActiveTsBootProc) GetEntityData() *types.CommonEntityData {
    activeTsBootProc.EntityData.YFilter = activeTsBootProc.YFilter
    activeTsBootProc.EntityData.YangName = "active-ts-boot-proc"
    activeTsBootProc.EntityData.BundleName = "cisco_ios_xr"
    activeTsBootProc.EntityData.ParentYangName = "failover"
    activeTsBootProc.EntityData.SegmentPath = "active-ts-boot-proc" + types.AddNoKeyToken(activeTsBootProc)
    activeTsBootProc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/failover/" + activeTsBootProc.EntityData.SegmentPath
    activeTsBootProc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeTsBootProc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeTsBootProc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeTsBootProc.EntityData.Children = types.NewOrderedMap()
    activeTsBootProc.EntityData.Leafs = types.NewOrderedMap()
    activeTsBootProc.EntityData.Leafs.Append("active-time-stamp", types.YLeaf{"ActiveTimeStamp", activeTsBootProc.ActiveTimeStamp})
    activeTsBootProc.EntityData.Leafs.Append("go-active", types.YLeaf{"GoActive", activeTsBootProc.GoActive})
    activeTsBootProc.EntityData.Leafs.Append("level", types.YLeaf{"Level", activeTsBootProc.Level})
    activeTsBootProc.EntityData.Leafs.Append("band-name", types.YLeaf{"BandName", activeTsBootProc.BandName})
    activeTsBootProc.EntityData.Leafs.Append("job-id", types.YLeaf{"JobId", activeTsBootProc.JobId})
    activeTsBootProc.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", activeTsBootProc.InstanceId})
    activeTsBootProc.EntityData.Leafs.Append("avail-time-stamp", types.YLeaf{"AvailTimeStamp", activeTsBootProc.AvailTimeStamp})
    activeTsBootProc.EntityData.Leafs.Append("avail", types.YLeaf{"Avail", activeTsBootProc.Avail})
    activeTsBootProc.EntityData.Leafs.Append("is-avail-timeout", types.YLeaf{"IsAvailTimeout", activeTsBootProc.IsAvailTimeout})
    activeTsBootProc.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", activeTsBootProc.ProcessName})

    activeTsBootProc.EntityData.YListKeys = []string {}

    return &(activeTsBootProc.EntityData)
}

// SystemProcess_NodeTable_Node_Failover_StartTsBootProc
// List of booted processes per start time
type SystemProcess_NodeTable_Node_Failover_StartTsBootProc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Start Time Stamp. The type is string.
    StartTimeStamp interface{}

    // Time since started. The type is string.
    Started interface{}

    // Level. The type is string.
    Level interface{}

    // Job Id. The type is interface{} with range: -2147483648..2147483647.
    Jid interface{}

    // Instance Id. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Ready Time Stamp. The type is string.
    ReadyTimeStamp interface{}

    // Time since Ready. The type is string.
    Ready interface{}

    // Is EOI timeout. The type is bool.
    IsEoiTimeout interface{}

    // Process Name. The type is string.
    ProcessName interface{}
}

func (startTsBootProc *SystemProcess_NodeTable_Node_Failover_StartTsBootProc) GetEntityData() *types.CommonEntityData {
    startTsBootProc.EntityData.YFilter = startTsBootProc.YFilter
    startTsBootProc.EntityData.YangName = "start-ts-boot-proc"
    startTsBootProc.EntityData.BundleName = "cisco_ios_xr"
    startTsBootProc.EntityData.ParentYangName = "failover"
    startTsBootProc.EntityData.SegmentPath = "start-ts-boot-proc" + types.AddNoKeyToken(startTsBootProc)
    startTsBootProc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/failover/" + startTsBootProc.EntityData.SegmentPath
    startTsBootProc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    startTsBootProc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    startTsBootProc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    startTsBootProc.EntityData.Children = types.NewOrderedMap()
    startTsBootProc.EntityData.Leafs = types.NewOrderedMap()
    startTsBootProc.EntityData.Leafs.Append("start-time-stamp", types.YLeaf{"StartTimeStamp", startTsBootProc.StartTimeStamp})
    startTsBootProc.EntityData.Leafs.Append("started", types.YLeaf{"Started", startTsBootProc.Started})
    startTsBootProc.EntityData.Leafs.Append("level", types.YLeaf{"Level", startTsBootProc.Level})
    startTsBootProc.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", startTsBootProc.Jid})
    startTsBootProc.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", startTsBootProc.InstanceId})
    startTsBootProc.EntityData.Leafs.Append("ready-time-stamp", types.YLeaf{"ReadyTimeStamp", startTsBootProc.ReadyTimeStamp})
    startTsBootProc.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", startTsBootProc.Ready})
    startTsBootProc.EntityData.Leafs.Append("is-eoi-timeout", types.YLeaf{"IsEoiTimeout", startTsBootProc.IsEoiTimeout})
    startTsBootProc.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", startTsBootProc.ProcessName})

    startTsBootProc.EntityData.YListKeys = []string {}

    return &(startTsBootProc.EntityData)
}

// SystemProcess_NodeTable_Node_Failover_PrimaryBandStatistic
// Primary Band statistics
type SystemProcess_NodeTable_Node_Failover_PrimaryBandStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Level. The type is string.
    Level interface{}

    // Band Name. The type is string.
    BandName interface{}

    // Band finish time. The type is string.
    BandFinishTime interface{}

    // Band time. The type is string.
    BandTime interface{}

    // Finish Time. The type is string.
    FinishTime interface{}

    // Idle Percentage. The type is string. Units are percentage.
    IdlePercentage interface{}

    // Jid. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Ready Time. The type is string.
    ReadyTime interface{}

    // Last Process Name. The type is string.
    LastProcess interface{}
}

func (primaryBandStatistic *SystemProcess_NodeTable_Node_Failover_PrimaryBandStatistic) GetEntityData() *types.CommonEntityData {
    primaryBandStatistic.EntityData.YFilter = primaryBandStatistic.YFilter
    primaryBandStatistic.EntityData.YangName = "primary-band-statistic"
    primaryBandStatistic.EntityData.BundleName = "cisco_ios_xr"
    primaryBandStatistic.EntityData.ParentYangName = "failover"
    primaryBandStatistic.EntityData.SegmentPath = "primary-band-statistic" + types.AddNoKeyToken(primaryBandStatistic)
    primaryBandStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/failover/" + primaryBandStatistic.EntityData.SegmentPath
    primaryBandStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryBandStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryBandStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryBandStatistic.EntityData.Children = types.NewOrderedMap()
    primaryBandStatistic.EntityData.Leafs = types.NewOrderedMap()
    primaryBandStatistic.EntityData.Leafs.Append("level", types.YLeaf{"Level", primaryBandStatistic.Level})
    primaryBandStatistic.EntityData.Leafs.Append("band-name", types.YLeaf{"BandName", primaryBandStatistic.BandName})
    primaryBandStatistic.EntityData.Leafs.Append("band-finish-time", types.YLeaf{"BandFinishTime", primaryBandStatistic.BandFinishTime})
    primaryBandStatistic.EntityData.Leafs.Append("band-time", types.YLeaf{"BandTime", primaryBandStatistic.BandTime})
    primaryBandStatistic.EntityData.Leafs.Append("finish-time", types.YLeaf{"FinishTime", primaryBandStatistic.FinishTime})
    primaryBandStatistic.EntityData.Leafs.Append("idle-percentage", types.YLeaf{"IdlePercentage", primaryBandStatistic.IdlePercentage})
    primaryBandStatistic.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", primaryBandStatistic.Jid})
    primaryBandStatistic.EntityData.Leafs.Append("ready-time", types.YLeaf{"ReadyTime", primaryBandStatistic.ReadyTime})
    primaryBandStatistic.EntityData.Leafs.Append("last-process", types.YLeaf{"LastProcess", primaryBandStatistic.LastProcess})

    primaryBandStatistic.EntityData.YListKeys = []string {}

    return &(primaryBandStatistic.EntityData)
}

// SystemProcess_NodeTable_Node_Failover_PrimaryTsBootProc
// List of booted processes per primary time
type SystemProcess_NodeTable_Node_Failover_PrimaryTsBootProc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Primary Time Stamp. The type is string.
    PrimTimeStamp interface{}

    // Go primary time stamp. The type is string.
    GoPrimary interface{}

    // Level. The type is string.
    Level interface{}

    // Band Name. The type is string.
    BandName interface{}

    // Job Id. The type is interface{} with range: -2147483648..2147483647.
    Jid interface{}

    // Instance Id. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Avail Time Stamp. The type is string.
    AvailTimeStamp interface{}

    // Time since Avail. The type is string.
    Avail interface{}

    // Is EOI timeout. The type is bool.
    IsAvailTimeout interface{}

    // Process Name. The type is string.
    ProcessName interface{}
}

func (primaryTsBootProc *SystemProcess_NodeTable_Node_Failover_PrimaryTsBootProc) GetEntityData() *types.CommonEntityData {
    primaryTsBootProc.EntityData.YFilter = primaryTsBootProc.YFilter
    primaryTsBootProc.EntityData.YangName = "primary-ts-boot-proc"
    primaryTsBootProc.EntityData.BundleName = "cisco_ios_xr"
    primaryTsBootProc.EntityData.ParentYangName = "failover"
    primaryTsBootProc.EntityData.SegmentPath = "primary-ts-boot-proc" + types.AddNoKeyToken(primaryTsBootProc)
    primaryTsBootProc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/failover/" + primaryTsBootProc.EntityData.SegmentPath
    primaryTsBootProc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryTsBootProc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryTsBootProc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryTsBootProc.EntityData.Children = types.NewOrderedMap()
    primaryTsBootProc.EntityData.Leafs = types.NewOrderedMap()
    primaryTsBootProc.EntityData.Leafs.Append("prim-time-stamp", types.YLeaf{"PrimTimeStamp", primaryTsBootProc.PrimTimeStamp})
    primaryTsBootProc.EntityData.Leafs.Append("go-primary", types.YLeaf{"GoPrimary", primaryTsBootProc.GoPrimary})
    primaryTsBootProc.EntityData.Leafs.Append("level", types.YLeaf{"Level", primaryTsBootProc.Level})
    primaryTsBootProc.EntityData.Leafs.Append("band-name", types.YLeaf{"BandName", primaryTsBootProc.BandName})
    primaryTsBootProc.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", primaryTsBootProc.Jid})
    primaryTsBootProc.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", primaryTsBootProc.InstanceId})
    primaryTsBootProc.EntityData.Leafs.Append("avail-time-stamp", types.YLeaf{"AvailTimeStamp", primaryTsBootProc.AvailTimeStamp})
    primaryTsBootProc.EntityData.Leafs.Append("avail", types.YLeaf{"Avail", primaryTsBootProc.Avail})
    primaryTsBootProc.EntityData.Leafs.Append("is-avail-timeout", types.YLeaf{"IsAvailTimeout", primaryTsBootProc.IsAvailTimeout})
    primaryTsBootProc.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", primaryTsBootProc.ProcessName})

    primaryTsBootProc.EntityData.YListKeys = []string {}

    return &(primaryTsBootProc.EntityData)
}

// SystemProcess_NodeTable_Node_Failover_PrimaryStartTsBootProc
// List of booted process per primary start time
type SystemProcess_NodeTable_Node_Failover_PrimaryStartTsBootProc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Start Time Stamp. The type is string.
    StartTimeStamp interface{}

    // Time since started. The type is string.
    Started interface{}

    // Level. The type is string.
    Level interface{}

    // Job Id. The type is interface{} with range: -2147483648..2147483647.
    Jid interface{}

    // Instance Id. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Ready Time Stamp. The type is string.
    ReadyTimeStamp interface{}

    // Time since Ready. The type is string.
    Ready interface{}

    // Is EOI timeout. The type is bool.
    IsEoiTimeout interface{}

    // Process Name. The type is string.
    ProcessName interface{}
}

func (primaryStartTsBootProc *SystemProcess_NodeTable_Node_Failover_PrimaryStartTsBootProc) GetEntityData() *types.CommonEntityData {
    primaryStartTsBootProc.EntityData.YFilter = primaryStartTsBootProc.YFilter
    primaryStartTsBootProc.EntityData.YangName = "primary-start-ts-boot-proc"
    primaryStartTsBootProc.EntityData.BundleName = "cisco_ios_xr"
    primaryStartTsBootProc.EntityData.ParentYangName = "failover"
    primaryStartTsBootProc.EntityData.SegmentPath = "primary-start-ts-boot-proc" + types.AddNoKeyToken(primaryStartTsBootProc)
    primaryStartTsBootProc.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/failover/" + primaryStartTsBootProc.EntityData.SegmentPath
    primaryStartTsBootProc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryStartTsBootProc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryStartTsBootProc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryStartTsBootProc.EntityData.Children = types.NewOrderedMap()
    primaryStartTsBootProc.EntityData.Leafs = types.NewOrderedMap()
    primaryStartTsBootProc.EntityData.Leafs.Append("start-time-stamp", types.YLeaf{"StartTimeStamp", primaryStartTsBootProc.StartTimeStamp})
    primaryStartTsBootProc.EntityData.Leafs.Append("started", types.YLeaf{"Started", primaryStartTsBootProc.Started})
    primaryStartTsBootProc.EntityData.Leafs.Append("level", types.YLeaf{"Level", primaryStartTsBootProc.Level})
    primaryStartTsBootProc.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", primaryStartTsBootProc.Jid})
    primaryStartTsBootProc.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", primaryStartTsBootProc.InstanceId})
    primaryStartTsBootProc.EntityData.Leafs.Append("ready-time-stamp", types.YLeaf{"ReadyTimeStamp", primaryStartTsBootProc.ReadyTimeStamp})
    primaryStartTsBootProc.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", primaryStartTsBootProc.Ready})
    primaryStartTsBootProc.EntityData.Leafs.Append("is-eoi-timeout", types.YLeaf{"IsEoiTimeout", primaryStartTsBootProc.IsEoiTimeout})
    primaryStartTsBootProc.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", primaryStartTsBootProc.ProcessName})

    primaryStartTsBootProc.EntityData.YListKeys = []string {}

    return &(primaryStartTsBootProc.EntityData)
}

// SystemProcess_NodeTable_Node_Boot
// Process Boot information
type SystemProcess_NodeTable_Node_Boot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Last process started. The type is string.
    LastProcessStarted interface{}

    // Standby Band statistics. The type is slice of
    // SystemProcess_NodeTable_Node_Boot_StandbyBandStatistic.
    StandbyBandStatistic []*SystemProcess_NodeTable_Node_Boot_StandbyBandStatistic

    // Active Band statistics. The type is slice of
    // SystemProcess_NodeTable_Node_Boot_ActiveBandStatistic.
    ActiveBandStatistic []*SystemProcess_NodeTable_Node_Boot_ActiveBandStatistic

    // List of booted processes. The type is slice of
    // SystemProcess_NodeTable_Node_Boot_BootedProcess.
    BootedProcess []*SystemProcess_NodeTable_Node_Boot_BootedProcess
}

func (boot *SystemProcess_NodeTable_Node_Boot) GetEntityData() *types.CommonEntityData {
    boot.EntityData.YFilter = boot.YFilter
    boot.EntityData.YangName = "boot"
    boot.EntityData.BundleName = "cisco_ios_xr"
    boot.EntityData.ParentYangName = "node"
    boot.EntityData.SegmentPath = "boot"
    boot.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + boot.EntityData.SegmentPath
    boot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    boot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    boot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    boot.EntityData.Children = types.NewOrderedMap()
    boot.EntityData.Children.Append("standby-band-statistic", types.YChild{"StandbyBandStatistic", nil})
    for i := range boot.StandbyBandStatistic {
        types.SetYListKey(boot.StandbyBandStatistic[i], i)
        boot.EntityData.Children.Append(types.GetSegmentPath(boot.StandbyBandStatistic[i]), types.YChild{"StandbyBandStatistic", boot.StandbyBandStatistic[i]})
    }
    boot.EntityData.Children.Append("active-band-statistic", types.YChild{"ActiveBandStatistic", nil})
    for i := range boot.ActiveBandStatistic {
        types.SetYListKey(boot.ActiveBandStatistic[i], i)
        boot.EntityData.Children.Append(types.GetSegmentPath(boot.ActiveBandStatistic[i]), types.YChild{"ActiveBandStatistic", boot.ActiveBandStatistic[i]})
    }
    boot.EntityData.Children.Append("booted-process", types.YChild{"BootedProcess", nil})
    for i := range boot.BootedProcess {
        types.SetYListKey(boot.BootedProcess[i], i)
        boot.EntityData.Children.Append(types.GetSegmentPath(boot.BootedProcess[i]), types.YChild{"BootedProcess", boot.BootedProcess[i]})
    }
    boot.EntityData.Leafs = types.NewOrderedMap()
    boot.EntityData.Leafs.Append("last-process-started", types.YLeaf{"LastProcessStarted", boot.LastProcessStarted})

    boot.EntityData.YListKeys = []string {}

    return &(boot.EntityData)
}

// SystemProcess_NodeTable_Node_Boot_StandbyBandStatistic
// Standby Band statistics
type SystemProcess_NodeTable_Node_Boot_StandbyBandStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Level. The type is string.
    Level interface{}

    // Band Name. The type is string.
    BandName interface{}

    // Band finish time. The type is string.
    BandFinishTime interface{}

    // Band time. The type is string.
    BandTime interface{}

    // Finish Time. The type is string.
    FinishTime interface{}

    // Idle Percentage. The type is string. Units are percentage.
    IdlePercentage interface{}

    // Jid. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Ready Time. The type is string.
    ReadyTime interface{}

    // Last Process Name. The type is string.
    LastProcess interface{}
}

func (standbyBandStatistic *SystemProcess_NodeTable_Node_Boot_StandbyBandStatistic) GetEntityData() *types.CommonEntityData {
    standbyBandStatistic.EntityData.YFilter = standbyBandStatistic.YFilter
    standbyBandStatistic.EntityData.YangName = "standby-band-statistic"
    standbyBandStatistic.EntityData.BundleName = "cisco_ios_xr"
    standbyBandStatistic.EntityData.ParentYangName = "boot"
    standbyBandStatistic.EntityData.SegmentPath = "standby-band-statistic" + types.AddNoKeyToken(standbyBandStatistic)
    standbyBandStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/boot/" + standbyBandStatistic.EntityData.SegmentPath
    standbyBandStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standbyBandStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standbyBandStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standbyBandStatistic.EntityData.Children = types.NewOrderedMap()
    standbyBandStatistic.EntityData.Leafs = types.NewOrderedMap()
    standbyBandStatistic.EntityData.Leafs.Append("level", types.YLeaf{"Level", standbyBandStatistic.Level})
    standbyBandStatistic.EntityData.Leafs.Append("band-name", types.YLeaf{"BandName", standbyBandStatistic.BandName})
    standbyBandStatistic.EntityData.Leafs.Append("band-finish-time", types.YLeaf{"BandFinishTime", standbyBandStatistic.BandFinishTime})
    standbyBandStatistic.EntityData.Leafs.Append("band-time", types.YLeaf{"BandTime", standbyBandStatistic.BandTime})
    standbyBandStatistic.EntityData.Leafs.Append("finish-time", types.YLeaf{"FinishTime", standbyBandStatistic.FinishTime})
    standbyBandStatistic.EntityData.Leafs.Append("idle-percentage", types.YLeaf{"IdlePercentage", standbyBandStatistic.IdlePercentage})
    standbyBandStatistic.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", standbyBandStatistic.Jid})
    standbyBandStatistic.EntityData.Leafs.Append("ready-time", types.YLeaf{"ReadyTime", standbyBandStatistic.ReadyTime})
    standbyBandStatistic.EntityData.Leafs.Append("last-process", types.YLeaf{"LastProcess", standbyBandStatistic.LastProcess})

    standbyBandStatistic.EntityData.YListKeys = []string {}

    return &(standbyBandStatistic.EntityData)
}

// SystemProcess_NodeTable_Node_Boot_ActiveBandStatistic
// Active Band statistics
type SystemProcess_NodeTable_Node_Boot_ActiveBandStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Level. The type is string.
    Level interface{}

    // Band Name. The type is string.
    BandName interface{}

    // Band finish time. The type is string.
    BandFinishTime interface{}

    // Band time. The type is string.
    BandTime interface{}

    // Finish Time. The type is string.
    FinishTime interface{}

    // Idle Percentage. The type is string. Units are percentage.
    IdlePercentage interface{}

    // Jid. The type is interface{} with range: 0..4294967295.
    Jid interface{}

    // Ready Time. The type is string.
    ReadyTime interface{}

    // Last Process Name. The type is string.
    LastProcess interface{}
}

func (activeBandStatistic *SystemProcess_NodeTable_Node_Boot_ActiveBandStatistic) GetEntityData() *types.CommonEntityData {
    activeBandStatistic.EntityData.YFilter = activeBandStatistic.YFilter
    activeBandStatistic.EntityData.YangName = "active-band-statistic"
    activeBandStatistic.EntityData.BundleName = "cisco_ios_xr"
    activeBandStatistic.EntityData.ParentYangName = "boot"
    activeBandStatistic.EntityData.SegmentPath = "active-band-statistic" + types.AddNoKeyToken(activeBandStatistic)
    activeBandStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/boot/" + activeBandStatistic.EntityData.SegmentPath
    activeBandStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeBandStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeBandStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeBandStatistic.EntityData.Children = types.NewOrderedMap()
    activeBandStatistic.EntityData.Leafs = types.NewOrderedMap()
    activeBandStatistic.EntityData.Leafs.Append("level", types.YLeaf{"Level", activeBandStatistic.Level})
    activeBandStatistic.EntityData.Leafs.Append("band-name", types.YLeaf{"BandName", activeBandStatistic.BandName})
    activeBandStatistic.EntityData.Leafs.Append("band-finish-time", types.YLeaf{"BandFinishTime", activeBandStatistic.BandFinishTime})
    activeBandStatistic.EntityData.Leafs.Append("band-time", types.YLeaf{"BandTime", activeBandStatistic.BandTime})
    activeBandStatistic.EntityData.Leafs.Append("finish-time", types.YLeaf{"FinishTime", activeBandStatistic.FinishTime})
    activeBandStatistic.EntityData.Leafs.Append("idle-percentage", types.YLeaf{"IdlePercentage", activeBandStatistic.IdlePercentage})
    activeBandStatistic.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", activeBandStatistic.Jid})
    activeBandStatistic.EntityData.Leafs.Append("ready-time", types.YLeaf{"ReadyTime", activeBandStatistic.ReadyTime})
    activeBandStatistic.EntityData.Leafs.Append("last-process", types.YLeaf{"LastProcess", activeBandStatistic.LastProcess})

    activeBandStatistic.EntityData.YListKeys = []string {}

    return &(activeBandStatistic.EntityData)
}

// SystemProcess_NodeTable_Node_Boot_BootedProcess
// List of booted processes
type SystemProcess_NodeTable_Node_Boot_BootedProcess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Start Time Stamp. The type is string.
    StartTimeStamp interface{}

    // Time since started. The type is string.
    Started interface{}

    // Level. The type is string.
    Level interface{}

    // Job Id. The type is interface{} with range: -2147483648..2147483647.
    Jid interface{}

    // Instance Id. The type is interface{} with range: -2147483648..2147483647.
    InstanceId interface{}

    // Ready Time Stamp. The type is string.
    ReadyTimeStamp interface{}

    // Time since Ready. The type is string.
    Ready interface{}

    // Is EOI timeout. The type is bool.
    IsEoiTimeout interface{}

    // Process Name. The type is string.
    ProcessName interface{}
}

func (bootedProcess *SystemProcess_NodeTable_Node_Boot_BootedProcess) GetEntityData() *types.CommonEntityData {
    bootedProcess.EntityData.YFilter = bootedProcess.YFilter
    bootedProcess.EntityData.YangName = "booted-process"
    bootedProcess.EntityData.BundleName = "cisco_ios_xr"
    bootedProcess.EntityData.ParentYangName = "boot"
    bootedProcess.EntityData.SegmentPath = "booted-process" + types.AddNoKeyToken(bootedProcess)
    bootedProcess.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/boot/" + bootedProcess.EntityData.SegmentPath
    bootedProcess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bootedProcess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bootedProcess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bootedProcess.EntityData.Children = types.NewOrderedMap()
    bootedProcess.EntityData.Leafs = types.NewOrderedMap()
    bootedProcess.EntityData.Leafs.Append("start-time-stamp", types.YLeaf{"StartTimeStamp", bootedProcess.StartTimeStamp})
    bootedProcess.EntityData.Leafs.Append("started", types.YLeaf{"Started", bootedProcess.Started})
    bootedProcess.EntityData.Leafs.Append("level", types.YLeaf{"Level", bootedProcess.Level})
    bootedProcess.EntityData.Leafs.Append("jid", types.YLeaf{"Jid", bootedProcess.Jid})
    bootedProcess.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", bootedProcess.InstanceId})
    bootedProcess.EntityData.Leafs.Append("ready-time-stamp", types.YLeaf{"ReadyTimeStamp", bootedProcess.ReadyTimeStamp})
    bootedProcess.EntityData.Leafs.Append("ready", types.YLeaf{"Ready", bootedProcess.Ready})
    bootedProcess.EntityData.Leafs.Append("is-eoi-timeout", types.YLeaf{"IsEoiTimeout", bootedProcess.IsEoiTimeout})
    bootedProcess.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", bootedProcess.ProcessName})

    bootedProcess.EntityData.YListKeys = []string {}

    return &(bootedProcess.EntityData)
}

// SystemProcess_NodeTable_Node_Logs
// Process Log information
type SystemProcess_NodeTable_Node_Logs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process log. The type is string.
    Log interface{}
}

func (logs *SystemProcess_NodeTable_Node_Logs) GetEntityData() *types.CommonEntityData {
    logs.EntityData.YFilter = logs.YFilter
    logs.EntityData.YangName = "logs"
    logs.EntityData.BundleName = "cisco_ios_xr"
    logs.EntityData.ParentYangName = "node"
    logs.EntityData.SegmentPath = "logs"
    logs.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + logs.EntityData.SegmentPath
    logs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logs.EntityData.Children = types.NewOrderedMap()
    logs.EntityData.Leafs = types.NewOrderedMap()
    logs.EntityData.Leafs.Append("log", types.YLeaf{"Log", logs.Log})

    logs.EntityData.YListKeys = []string {}

    return &(logs.EntityData)
}

// SystemProcess_NodeTable_Node_Searchpath
// Process Searchpath information
type SystemProcess_NodeTable_Node_Searchpath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // process searchpath. The type is string.
    Path interface{}
}

func (searchpath *SystemProcess_NodeTable_Node_Searchpath) GetEntityData() *types.CommonEntityData {
    searchpath.EntityData.YFilter = searchpath.YFilter
    searchpath.EntityData.YangName = "searchpath"
    searchpath.EntityData.BundleName = "cisco_ios_xr"
    searchpath.EntityData.ParentYangName = "node"
    searchpath.EntityData.SegmentPath = "searchpath"
    searchpath.EntityData.AbsolutePath = "Cisco-IOS-XR-sysmgr-oper:system-process/node-table/node/" + searchpath.EntityData.SegmentPath
    searchpath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    searchpath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    searchpath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    searchpath.EntityData.Children = types.NewOrderedMap()
    searchpath.EntityData.Leafs = types.NewOrderedMap()
    searchpath.EntityData.Leafs.Append("path", types.YLeaf{"Path", searchpath.Path})

    searchpath.EntityData.YListKeys = []string {}

    return &(searchpath.EntityData)
}

