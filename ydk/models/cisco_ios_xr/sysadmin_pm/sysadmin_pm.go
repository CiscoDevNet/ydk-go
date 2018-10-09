// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// The Process Manager (PM).
// 
// Copyright(c) 2011-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_pm

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_pm"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-pm processes}", reflect.TypeOf(Processes{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-pm:processes", reflect.TypeOf(Processes{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-pm process-manager}", reflect.TypeOf(ProcessManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-pm:process-manager", reflect.TypeOf(ProcessManager{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-pm pm}", reflect.TypeOf(Pm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-pm:pm", reflect.TypeOf(Pm{}))
}

// StartupMode
type StartupMode string

const (
    StartupMode_ON_BOOTUP StartupMode = "ON-BOOTUP"

    StartupMode_ON_SELECTION StartupMode = "ON-SELECTION"

    StartupMode_ON_DEMAND StartupMode = "ON-DEMAND"
)

// ServiceRole
type ServiceRole string

const (
    ServiceRole_NONE ServiceRole = "NONE"

    ServiceRole_ACTIVE ServiceRole = "ACTIVE"

    ServiceRole_STANDBY ServiceRole = "STANDBY"
)

// ServiceState
type ServiceState string

const (
    ServiceState_SS_IDLE ServiceState = "SS_IDLE"

    ServiceState_SS_RUNNING ServiceState = "SS_RUNNING"

    ServiceState_SS_ACK_PENDING ServiceState = "SS_ACK_PENDING"
)

// ProcessState
type ProcessState string

const (
    ProcessState_IDLE ProcessState = "IDLE"

    ProcessState_RUNNING ProcessState = "RUNNING"

    ProcessState_STOPPING ProcessState = "STOPPING"

    ProcessState_STOPPED ProcessState = "STOPPED"

    ProcessState_DESELECTING ProcessState = "DESELECTING"

    ProcessState_DESELECTED ProcessState = "DESELECTED"
)

// ServiceScope
type ServiceScope string

const (
    ServiceScope_SYSTEM ServiceScope = "SYSTEM"

    ServiceScope_RACK ServiceScope = "RACK"
)

// Processes
// Process Info
type Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Processes_AllLocations.
    AllLocations []*Processes_AllLocations
}

func (processes *Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xr"
    processes.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-pm"
    processes.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-pm:processes"
    processes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processes.EntityData.Children = types.NewOrderedMap()
    processes.EntityData.Children.Append("all-locations", types.YChild{"AllLocations", nil})
    for i := range processes.AllLocations {
        processes.EntityData.Children.Append(types.GetSegmentPath(processes.AllLocations[i]), types.YChild{"AllLocations", processes.AllLocations[i]})
    }
    processes.EntityData.Leafs = types.NewOrderedMap()

    processes.EntityData.YListKeys = []string {}

    return &(processes.EntityData)
}

// Processes_AllLocations
type Processes_AllLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}

    // IP address of the location. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddr interface{}

    // Total number of process control blocks. The type is interface{} with range:
    // 0..4294967295.
    Pcbs interface{}

    // The type is slice of Processes_AllLocations_Name.
    Name []*Processes_AllLocations_Name
}

func (allLocations *Processes_AllLocations) GetEntityData() *types.CommonEntityData {
    allLocations.EntityData.YFilter = allLocations.YFilter
    allLocations.EntityData.YangName = "all-locations"
    allLocations.EntityData.BundleName = "cisco_ios_xr"
    allLocations.EntityData.ParentYangName = "processes"
    allLocations.EntityData.SegmentPath = "all-locations" + types.AddKeyToken(allLocations.Location, "location")
    allLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocations.EntityData.Children = types.NewOrderedMap()
    allLocations.EntityData.Children.Append("name", types.YChild{"Name", nil})
    for i := range allLocations.Name {
        allLocations.EntityData.Children.Append(types.GetSegmentPath(allLocations.Name[i]), types.YChild{"Name", allLocations.Name[i]})
    }
    allLocations.EntityData.Leafs = types.NewOrderedMap()
    allLocations.EntityData.Leafs.Append("location", types.YLeaf{"Location", allLocations.Location})
    allLocations.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", allLocations.IpAddr})
    allLocations.EntityData.Leafs.Append("pcbs", types.YLeaf{"Pcbs", allLocations.Pcbs})

    allLocations.EntityData.YListKeys = []string {"Location"}

    return &(allLocations.EntityData)
}

// Processes_AllLocations_Name
type Processes_AllLocations_Name struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the process. The type is string.
    ProcName interface{}

    // This attribute is a key. Instance identifier. The type is interface{} with
    // range: 0..4294967295.
    InstanceId interface{}

    // Process path. The type is string.
    Path interface{}

    // Process startup file. The type is string.
    StartupFile interface{}

    // When is a process started. The type is StartupMode.
    StartupMode interface{}

    // Heart beat timeout in sec. The type is interface{} with range:
    // 0..4294967295.
    HeartBeatTimeout interface{}

    // How long ago last heart beat was detected. The type is string.
    LastHeartBeatTime interface{}

    // Maximum num of restarts. The type is interface{} with range: 0..4294967295.
    MaxRestarts interface{}

    // Respawn reset timer in min. The type is interface{} with range:
    // 0..4294967295.
    RespawnResetTimer interface{}

    // Mandatory process. The type is bool.
    Mandatory interface{}

    // Should run during maintenance mode. The type is bool.
    MaintMode interface{}

    // Process arguments. The type is string.
    Args interface{}

    // State of the process. The type is ProcessState.
    ProcState interface{}

    // Process ID. The type is interface{} with range: -2147483648..2147483647.
    Pid interface{}

    // Whether the processes ever aborted. The type is bool.
    ProcAborted interface{}

    // Last exit status/info of the process. The type is string.
    ExitStatus interface{}

    // Total number of respawns of the process. The type is interface{} with
    // range: -2147483648..2147483647.
    Respawns interface{}

    // Last start date and time. The type is string.
    StartTime interface{}

    // Time for ready from start-time. The type is string.
    ReadyTime interface{}

    // Last exit date and time. The type is string.
    LastExitTime interface{}

    // The type is slice of Processes_AllLocations_Name_Services.
    Services []*Processes_AllLocations_Name_Services
}

func (name *Processes_AllLocations_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "all-locations"
    name.EntityData.SegmentPath = "name" + types.AddKeyToken(name.ProcName, "proc-name") + types.AddKeyToken(name.InstanceId, "instance-id")
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = types.NewOrderedMap()
    name.EntityData.Children.Append("services", types.YChild{"Services", nil})
    for i := range name.Services {
        name.EntityData.Children.Append(types.GetSegmentPath(name.Services[i]), types.YChild{"Services", name.Services[i]})
    }
    name.EntityData.Leafs = types.NewOrderedMap()
    name.EntityData.Leafs.Append("proc-name", types.YLeaf{"ProcName", name.ProcName})
    name.EntityData.Leafs.Append("instance-id", types.YLeaf{"InstanceId", name.InstanceId})
    name.EntityData.Leafs.Append("path", types.YLeaf{"Path", name.Path})
    name.EntityData.Leafs.Append("startup-file", types.YLeaf{"StartupFile", name.StartupFile})
    name.EntityData.Leafs.Append("startup-mode", types.YLeaf{"StartupMode", name.StartupMode})
    name.EntityData.Leafs.Append("heart-beat-timeout", types.YLeaf{"HeartBeatTimeout", name.HeartBeatTimeout})
    name.EntityData.Leafs.Append("last-heart-beat-time", types.YLeaf{"LastHeartBeatTime", name.LastHeartBeatTime})
    name.EntityData.Leafs.Append("max-restarts", types.YLeaf{"MaxRestarts", name.MaxRestarts})
    name.EntityData.Leafs.Append("respawn-reset-timer", types.YLeaf{"RespawnResetTimer", name.RespawnResetTimer})
    name.EntityData.Leafs.Append("mandatory", types.YLeaf{"Mandatory", name.Mandatory})
    name.EntityData.Leafs.Append("maint-mode", types.YLeaf{"MaintMode", name.MaintMode})
    name.EntityData.Leafs.Append("args", types.YLeaf{"Args", name.Args})
    name.EntityData.Leafs.Append("proc-state", types.YLeaf{"ProcState", name.ProcState})
    name.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", name.Pid})
    name.EntityData.Leafs.Append("proc-aborted", types.YLeaf{"ProcAborted", name.ProcAborted})
    name.EntityData.Leafs.Append("exit-status", types.YLeaf{"ExitStatus", name.ExitStatus})
    name.EntityData.Leafs.Append("respawns", types.YLeaf{"Respawns", name.Respawns})
    name.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", name.StartTime})
    name.EntityData.Leafs.Append("ready-time", types.YLeaf{"ReadyTime", name.ReadyTime})
    name.EntityData.Leafs.Append("last-exit-time", types.YLeaf{"LastExitTime", name.LastExitTime})

    name.EntityData.YListKeys = []string {"ProcName", "InstanceId"}

    return &(name.EntityData)
}

// Processes_AllLocations_Name_Services
type Processes_AllLocations_Name_Services struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the service. The type is string.
    ServiceName interface{}

    // Scope of the service. The type is ServiceScope.
    Scope interface{}

    // Service redundancy support. The type is bool.
    Redundancy interface{}

    // Standby ready for HA. The type is bool.
    HaReady interface{}

    // State of the service. The type is ServiceState.
    ServiceState interface{}

    // Service role. The type is ServiceRole.
    HaRole interface{}

    // New service role, different if PM in process of assigning. The type is
    // ServiceRole.
    NewHaRole interface{}

    // Service seleted to run on the node. The type is bool.
    Selected interface{}

    // First IP address in the selection. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip1 interface{}

    // Second IP address in the selection. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip2 interface{}

    // Last start date and time. The type is string.
    SvcStartTime interface{}

    // Time it took to get ready since start time. The type is string.
    SvcReadyTime interface{}

    // Time it took to get HA-ready since start time. The type is string.
    SvcHareadyTime interface{}
}

func (services *Processes_AllLocations_Name_Services) GetEntityData() *types.CommonEntityData {
    services.EntityData.YFilter = services.YFilter
    services.EntityData.YangName = "services"
    services.EntityData.BundleName = "cisco_ios_xr"
    services.EntityData.ParentYangName = "name"
    services.EntityData.SegmentPath = "services" + types.AddKeyToken(services.ServiceName, "service-name")
    services.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    services.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    services.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    services.EntityData.Children = types.NewOrderedMap()
    services.EntityData.Leafs = types.NewOrderedMap()
    services.EntityData.Leafs.Append("service-name", types.YLeaf{"ServiceName", services.ServiceName})
    services.EntityData.Leafs.Append("scope", types.YLeaf{"Scope", services.Scope})
    services.EntityData.Leafs.Append("redundancy", types.YLeaf{"Redundancy", services.Redundancy})
    services.EntityData.Leafs.Append("ha-ready", types.YLeaf{"HaReady", services.HaReady})
    services.EntityData.Leafs.Append("service-state", types.YLeaf{"ServiceState", services.ServiceState})
    services.EntityData.Leafs.Append("ha-role", types.YLeaf{"HaRole", services.HaRole})
    services.EntityData.Leafs.Append("new-ha-role", types.YLeaf{"NewHaRole", services.NewHaRole})
    services.EntityData.Leafs.Append("selected", types.YLeaf{"Selected", services.Selected})
    services.EntityData.Leafs.Append("ip1", types.YLeaf{"Ip1", services.Ip1})
    services.EntityData.Leafs.Append("ip2", types.YLeaf{"Ip2", services.Ip2})
    services.EntityData.Leafs.Append("svc-start-time", types.YLeaf{"SvcStartTime", services.SvcStartTime})
    services.EntityData.Leafs.Append("svc-ready-time", types.YLeaf{"SvcReadyTime", services.SvcReadyTime})
    services.EntityData.Leafs.Append("svc-haready-time", types.YLeaf{"SvcHareadyTime", services.SvcHareadyTime})

    services.EntityData.YListKeys = []string {"ServiceName"}

    return &(services.EntityData)
}

// ProcessManager
// Process Manager Info
type ProcessManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ProcessManager_AllLocationsInfo.
    AllLocationsInfo []*ProcessManager_AllLocationsInfo
}

func (processManager *ProcessManager) GetEntityData() *types.CommonEntityData {
    processManager.EntityData.YFilter = processManager.YFilter
    processManager.EntityData.YangName = "process-manager"
    processManager.EntityData.BundleName = "cisco_ios_xr"
    processManager.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-pm"
    processManager.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-pm:process-manager"
    processManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processManager.EntityData.Children = types.NewOrderedMap()
    processManager.EntityData.Children.Append("all-locations-info", types.YChild{"AllLocationsInfo", nil})
    for i := range processManager.AllLocationsInfo {
        processManager.EntityData.Children.Append(types.GetSegmentPath(processManager.AllLocationsInfo[i]), types.YChild{"AllLocationsInfo", processManager.AllLocationsInfo[i]})
    }
    processManager.EntityData.Leafs = types.NewOrderedMap()

    processManager.EntityData.YListKeys = []string {}

    return &(processManager.EntityData)
}

// ProcessManager_AllLocationsInfo
type ProcessManager_AllLocationsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationInfo interface{}

    // IP address of the location. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddrInfo interface{}

    // Last start date and time for PM. The type is string.
    PmStartTime interface{}

    // PM in mandatory process down state. The type is bool.
    MandProcDown interface{}

    // Status of CAPI with vm-manager. The type is bool.
    VmmCapiUp interface{}

    // Status of CAPI with wdmon. The type is bool.
    WdmonCapiUp interface{}

    // Date and time of last wdmon CAPI connection establish. The type is string.
    WdmonCapiTimestamp interface{}

    // Number of times wdmon CAPI connection extablished. The type is interface{}
    // with range: 0..4294967295.
    WdmonNumCapiConnects interface{}
}

func (allLocationsInfo *ProcessManager_AllLocationsInfo) GetEntityData() *types.CommonEntityData {
    allLocationsInfo.EntityData.YFilter = allLocationsInfo.YFilter
    allLocationsInfo.EntityData.YangName = "all-locations-info"
    allLocationsInfo.EntityData.BundleName = "cisco_ios_xr"
    allLocationsInfo.EntityData.ParentYangName = "process-manager"
    allLocationsInfo.EntityData.SegmentPath = "all-locations-info" + types.AddKeyToken(allLocationsInfo.LocationInfo, "location-info")
    allLocationsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocationsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocationsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocationsInfo.EntityData.Children = types.NewOrderedMap()
    allLocationsInfo.EntityData.Leafs = types.NewOrderedMap()
    allLocationsInfo.EntityData.Leafs.Append("location-info", types.YLeaf{"LocationInfo", allLocationsInfo.LocationInfo})
    allLocationsInfo.EntityData.Leafs.Append("ip-addr-info", types.YLeaf{"IpAddrInfo", allLocationsInfo.IpAddrInfo})
    allLocationsInfo.EntityData.Leafs.Append("pm-start-time", types.YLeaf{"PmStartTime", allLocationsInfo.PmStartTime})
    allLocationsInfo.EntityData.Leafs.Append("mand-proc-down", types.YLeaf{"MandProcDown", allLocationsInfo.MandProcDown})
    allLocationsInfo.EntityData.Leafs.Append("vmm-capi-up", types.YLeaf{"VmmCapiUp", allLocationsInfo.VmmCapiUp})
    allLocationsInfo.EntityData.Leafs.Append("wdmon-capi-up", types.YLeaf{"WdmonCapiUp", allLocationsInfo.WdmonCapiUp})
    allLocationsInfo.EntityData.Leafs.Append("wdmon-capi-timestamp", types.YLeaf{"WdmonCapiTimestamp", allLocationsInfo.WdmonCapiTimestamp})
    allLocationsInfo.EntityData.Leafs.Append("wdmon-num-capi-connects", types.YLeaf{"WdmonNumCapiConnects", allLocationsInfo.WdmonNumCapiConnects})

    allLocationsInfo.EntityData.YListKeys = []string {"LocationInfo"}

    return &(allLocationsInfo.EntityData)
}

// Pm
type Pm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Pm Pm_Pm
}

func (pm *Pm) GetEntityData() *types.CommonEntityData {
    pm.EntityData.YFilter = pm.YFilter
    pm.EntityData.YangName = "pm"
    pm.EntityData.BundleName = "cisco_ios_xr"
    pm.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-pm"
    pm.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-pm:pm"
    pm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pm.EntityData.Children = types.NewOrderedMap()
    pm.EntityData.Children.Append("pm", types.YChild{"Pm", &pm.Pm})
    pm.EntityData.Leafs = types.NewOrderedMap()

    pm.EntityData.YListKeys = []string {}

    return &(pm.EntityData)
}

// Pm_Pm
type Pm_Pm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Pm_Pm_Trace.
    Trace []*Pm_Pm_Trace
}

func (pm *Pm_Pm) GetEntityData() *types.CommonEntityData {
    pm.EntityData.YFilter = pm.YFilter
    pm.EntityData.YangName = "pm"
    pm.EntityData.BundleName = "cisco_ios_xr"
    pm.EntityData.ParentYangName = "pm"
    pm.EntityData.SegmentPath = "pm"
    pm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pm.EntityData.Children = types.NewOrderedMap()
    pm.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range pm.Trace {
        pm.EntityData.Children.Append(types.GetSegmentPath(pm.Trace[i]), types.YChild{"Trace", pm.Trace[i]})
    }
    pm.EntityData.Leafs = types.NewOrderedMap()

    pm.EntityData.YListKeys = []string {}

    return &(pm.EntityData)
}

// Pm_Pm_Trace
// show traceable processes
type Pm_Pm_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Pm_Pm_Trace_Location.
    Location []*Pm_Pm_Trace_Location
}

func (trace *Pm_Pm_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "pm"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range trace.Location {
        trace.EntityData.Children.Append(types.GetSegmentPath(trace.Location[i]), types.YChild{"Location", trace.Location[i]})
    }
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("buffer", types.YLeaf{"Buffer", trace.Buffer})

    trace.EntityData.YListKeys = []string {"Buffer"}

    return &(trace.EntityData)
}

// Pm_Pm_Trace_Location
type Pm_Pm_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Pm_Pm_Trace_Location_AllOptions.
    AllOptions []*Pm_Pm_Trace_Location_AllOptions
}

func (location *Pm_Pm_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("all-options", types.YChild{"AllOptions", nil})
    for i := range location.AllOptions {
        location.EntityData.Children.Append(types.GetSegmentPath(location.AllOptions[i]), types.YChild{"AllOptions", location.AllOptions[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location_name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Pm_Pm_Trace_Location_AllOptions
type Pm_Pm_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Pm_Pm_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Pm_Pm_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Pm_Pm_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// Pm_Pm_Trace_Location_AllOptions_TraceBlocks
type Pm_Pm_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Pm_Pm_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks"
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

