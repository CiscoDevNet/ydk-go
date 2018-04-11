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

// Startupmode
type Startupmode string

const (
    Startupmode_ON_BOOTUP Startupmode = "ON-BOOTUP"

    Startupmode_ON_SELECTION Startupmode = "ON-SELECTION"

    Startupmode_ON_DEMAND Startupmode = "ON-DEMAND"
)

// Processstate
type Processstate string

const (
    Processstate_IDLE Processstate = "IDLE"

    Processstate_RUNNING Processstate = "RUNNING"

    Processstate_STOPPING Processstate = "STOPPING"

    Processstate_STOPPED Processstate = "STOPPED"

    Processstate_DESELECTING Processstate = "DESELECTING"

    Processstate_DESELECTED Processstate = "DESELECTED"
)

// Servicescope
type Servicescope string

const (
    Servicescope_SYSTEM Servicescope = "SYSTEM"

    Servicescope_RACK Servicescope = "RACK"
)

// Servicestate
type Servicestate string

const (
    Servicestate_SS_IDLE Servicestate = "SS_IDLE"

    Servicestate_SS_RUNNING Servicestate = "SS_RUNNING"

    Servicestate_SS_ACK_PENDING Servicestate = "SS_ACK_PENDING"
)

// Servicerole
type Servicerole string

const (
    Servicerole_NONE Servicerole = "NONE"

    Servicerole_ACTIVE Servicerole = "ACTIVE"

    Servicerole_STANDBY Servicerole = "STANDBY"
)

// Processes
// Process Info
type Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Processes_AllLocations.
    AllLocations []Processes_AllLocations
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

    processes.EntityData.Children = make(map[string]types.YChild)
    processes.EntityData.Children["all-locations"] = types.YChild{"AllLocations", nil}
    for i := range processes.AllLocations {
        processes.EntityData.Children[types.GetSegmentPath(&processes.AllLocations[i])] = types.YChild{"AllLocations", &processes.AllLocations[i]}
    }
    processes.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddr interface{}

    // Total number of process control blocks. The type is interface{} with range:
    // 0..4294967295.
    Pcbs interface{}

    // The type is slice of Processes_AllLocations_Name.
    Name []Processes_AllLocations_Name
}

func (allLocations *Processes_AllLocations) GetEntityData() *types.CommonEntityData {
    allLocations.EntityData.YFilter = allLocations.YFilter
    allLocations.EntityData.YangName = "all-locations"
    allLocations.EntityData.BundleName = "cisco_ios_xr"
    allLocations.EntityData.ParentYangName = "processes"
    allLocations.EntityData.SegmentPath = "all-locations" + "[location='" + fmt.Sprintf("%v", allLocations.Location) + "']"
    allLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocations.EntityData.Children = make(map[string]types.YChild)
    allLocations.EntityData.Children["name"] = types.YChild{"Name", nil}
    for i := range allLocations.Name {
        allLocations.EntityData.Children[types.GetSegmentPath(&allLocations.Name[i])] = types.YChild{"Name", &allLocations.Name[i]}
    }
    allLocations.EntityData.Leafs = make(map[string]types.YLeaf)
    allLocations.EntityData.Leafs["location"] = types.YLeaf{"Location", allLocations.Location}
    allLocations.EntityData.Leafs["ip-addr"] = types.YLeaf{"IpAddr", allLocations.IpAddr}
    allLocations.EntityData.Leafs["pcbs"] = types.YLeaf{"Pcbs", allLocations.Pcbs}
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

    // When is a process started. The type is Startupmode.
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

    // State of the process. The type is Processstate.
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
    Services []Processes_AllLocations_Name_Services
}

func (name *Processes_AllLocations_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "all-locations"
    name.EntityData.SegmentPath = "name" + "[proc-name='" + fmt.Sprintf("%v", name.ProcName) + "']" + "[instance-id='" + fmt.Sprintf("%v", name.InstanceId) + "']"
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = make(map[string]types.YChild)
    name.EntityData.Children["services"] = types.YChild{"Services", nil}
    for i := range name.Services {
        name.EntityData.Children[types.GetSegmentPath(&name.Services[i])] = types.YChild{"Services", &name.Services[i]}
    }
    name.EntityData.Leafs = make(map[string]types.YLeaf)
    name.EntityData.Leafs["proc-name"] = types.YLeaf{"ProcName", name.ProcName}
    name.EntityData.Leafs["instance-id"] = types.YLeaf{"InstanceId", name.InstanceId}
    name.EntityData.Leafs["path"] = types.YLeaf{"Path", name.Path}
    name.EntityData.Leafs["startup-file"] = types.YLeaf{"StartupFile", name.StartupFile}
    name.EntityData.Leafs["startup-mode"] = types.YLeaf{"StartupMode", name.StartupMode}
    name.EntityData.Leafs["heart-beat-timeout"] = types.YLeaf{"HeartBeatTimeout", name.HeartBeatTimeout}
    name.EntityData.Leafs["last-heart-beat-time"] = types.YLeaf{"LastHeartBeatTime", name.LastHeartBeatTime}
    name.EntityData.Leafs["max-restarts"] = types.YLeaf{"MaxRestarts", name.MaxRestarts}
    name.EntityData.Leafs["respawn-reset-timer"] = types.YLeaf{"RespawnResetTimer", name.RespawnResetTimer}
    name.EntityData.Leafs["mandatory"] = types.YLeaf{"Mandatory", name.Mandatory}
    name.EntityData.Leafs["maint-mode"] = types.YLeaf{"MaintMode", name.MaintMode}
    name.EntityData.Leafs["args"] = types.YLeaf{"Args", name.Args}
    name.EntityData.Leafs["proc-state"] = types.YLeaf{"ProcState", name.ProcState}
    name.EntityData.Leafs["pid"] = types.YLeaf{"Pid", name.Pid}
    name.EntityData.Leafs["proc-aborted"] = types.YLeaf{"ProcAborted", name.ProcAborted}
    name.EntityData.Leafs["exit-status"] = types.YLeaf{"ExitStatus", name.ExitStatus}
    name.EntityData.Leafs["respawns"] = types.YLeaf{"Respawns", name.Respawns}
    name.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", name.StartTime}
    name.EntityData.Leafs["ready-time"] = types.YLeaf{"ReadyTime", name.ReadyTime}
    name.EntityData.Leafs["last-exit-time"] = types.YLeaf{"LastExitTime", name.LastExitTime}
    return &(name.EntityData)
}

// Processes_AllLocations_Name_Services
type Processes_AllLocations_Name_Services struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the service. The type is string.
    ServiceName interface{}

    // Scope of the service. The type is Servicescope.
    Scope interface{}

    // Service redundancy support. The type is bool.
    Redundancy interface{}

    // Standby ready for HA. The type is bool.
    HaReady interface{}

    // State of the service. The type is Servicestate.
    ServiceState interface{}

    // Service role. The type is Servicerole.
    HaRole interface{}

    // New service role, different if PM in process of assigning. The type is
    // Servicerole.
    NewHaRole interface{}

    // Service seleted to run on the node. The type is bool.
    Selected interface{}

    // First IP address in the selection. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip1 interface{}

    // Second IP address in the selection. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    services.EntityData.SegmentPath = "services" + "[service-name='" + fmt.Sprintf("%v", services.ServiceName) + "']"
    services.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    services.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    services.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    services.EntityData.Children = make(map[string]types.YChild)
    services.EntityData.Leafs = make(map[string]types.YLeaf)
    services.EntityData.Leafs["service-name"] = types.YLeaf{"ServiceName", services.ServiceName}
    services.EntityData.Leafs["scope"] = types.YLeaf{"Scope", services.Scope}
    services.EntityData.Leafs["redundancy"] = types.YLeaf{"Redundancy", services.Redundancy}
    services.EntityData.Leafs["ha-ready"] = types.YLeaf{"HaReady", services.HaReady}
    services.EntityData.Leafs["service-state"] = types.YLeaf{"ServiceState", services.ServiceState}
    services.EntityData.Leafs["ha-role"] = types.YLeaf{"HaRole", services.HaRole}
    services.EntityData.Leafs["new-ha-role"] = types.YLeaf{"NewHaRole", services.NewHaRole}
    services.EntityData.Leafs["selected"] = types.YLeaf{"Selected", services.Selected}
    services.EntityData.Leafs["ip1"] = types.YLeaf{"Ip1", services.Ip1}
    services.EntityData.Leafs["ip2"] = types.YLeaf{"Ip2", services.Ip2}
    services.EntityData.Leafs["svc-start-time"] = types.YLeaf{"SvcStartTime", services.SvcStartTime}
    services.EntityData.Leafs["svc-ready-time"] = types.YLeaf{"SvcReadyTime", services.SvcReadyTime}
    services.EntityData.Leafs["svc-haready-time"] = types.YLeaf{"SvcHareadyTime", services.SvcHareadyTime}
    return &(services.EntityData)
}

// ProcessManager
// Process Manager Info
type ProcessManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ProcessManager_AllLocationsInfo.
    AllLocationsInfo []ProcessManager_AllLocationsInfo
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

    processManager.EntityData.Children = make(map[string]types.YChild)
    processManager.EntityData.Children["all-locations-info"] = types.YChild{"AllLocationsInfo", nil}
    for i := range processManager.AllLocationsInfo {
        processManager.EntityData.Children[types.GetSegmentPath(&processManager.AllLocationsInfo[i])] = types.YChild{"AllLocationsInfo", &processManager.AllLocationsInfo[i]}
    }
    processManager.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    allLocationsInfo.EntityData.SegmentPath = "all-locations-info" + "[location-info='" + fmt.Sprintf("%v", allLocationsInfo.LocationInfo) + "']"
    allLocationsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocationsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocationsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocationsInfo.EntityData.Children = make(map[string]types.YChild)
    allLocationsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    allLocationsInfo.EntityData.Leafs["location-info"] = types.YLeaf{"LocationInfo", allLocationsInfo.LocationInfo}
    allLocationsInfo.EntityData.Leafs["ip-addr-info"] = types.YLeaf{"IpAddrInfo", allLocationsInfo.IpAddrInfo}
    allLocationsInfo.EntityData.Leafs["pm-start-time"] = types.YLeaf{"PmStartTime", allLocationsInfo.PmStartTime}
    allLocationsInfo.EntityData.Leafs["mand-proc-down"] = types.YLeaf{"MandProcDown", allLocationsInfo.MandProcDown}
    allLocationsInfo.EntityData.Leafs["vmm-capi-up"] = types.YLeaf{"VmmCapiUp", allLocationsInfo.VmmCapiUp}
    allLocationsInfo.EntityData.Leafs["wdmon-capi-up"] = types.YLeaf{"WdmonCapiUp", allLocationsInfo.WdmonCapiUp}
    allLocationsInfo.EntityData.Leafs["wdmon-capi-timestamp"] = types.YLeaf{"WdmonCapiTimestamp", allLocationsInfo.WdmonCapiTimestamp}
    allLocationsInfo.EntityData.Leafs["wdmon-num-capi-connects"] = types.YLeaf{"WdmonNumCapiConnects", allLocationsInfo.WdmonNumCapiConnects}
    return &(allLocationsInfo.EntityData)
}

// Pm
type Pm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Pm Pm_Pm_
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

    pm.EntityData.Children = make(map[string]types.YChild)
    pm.EntityData.Children["pm"] = types.YChild{"Pm", &pm.Pm}
    pm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pm.EntityData)
}

// Pm_Pm_
type Pm_Pm_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Pm_Pm__Trace.
    Trace []Pm_Pm__Trace
}

func (pm_ *Pm_Pm_) GetEntityData() *types.CommonEntityData {
    pm_.EntityData.YFilter = pm_.YFilter
    pm_.EntityData.YangName = "pm"
    pm_.EntityData.BundleName = "cisco_ios_xr"
    pm_.EntityData.ParentYangName = "pm"
    pm_.EntityData.SegmentPath = "pm"
    pm_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pm_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pm_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pm_.EntityData.Children = make(map[string]types.YChild)
    pm_.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range pm_.Trace {
        pm_.EntityData.Children[types.GetSegmentPath(&pm_.Trace[i])] = types.YChild{"Trace", &pm_.Trace[i]}
    }
    pm_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pm_.EntityData)
}

// Pm_Pm__Trace
// show traceable processes
type Pm_Pm__Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Pm_Pm__Trace_Location.
    Location []Pm_Pm__Trace_Location
}

func (trace *Pm_Pm__Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "pm"
    trace.EntityData.SegmentPath = "trace" + "[buffer='" + fmt.Sprintf("%v", trace.Buffer) + "']"
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = make(map[string]types.YChild)
    trace.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range trace.Location {
        trace.EntityData.Children[types.GetSegmentPath(&trace.Location[i])] = types.YChild{"Location", &trace.Location[i]}
    }
    trace.EntityData.Leafs = make(map[string]types.YLeaf)
    trace.EntityData.Leafs["buffer"] = types.YLeaf{"Buffer", trace.Buffer}
    return &(trace.EntityData)
}

// Pm_Pm__Trace_Location
type Pm_Pm__Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Pm_Pm__Trace_Location_AllOptions.
    AllOptions []Pm_Pm__Trace_Location_AllOptions
}

func (location *Pm_Pm__Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + "[location_name='" + fmt.Sprintf("%v", location.LocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["all-options"] = types.YChild{"AllOptions", nil}
    for i := range location.AllOptions {
        location.EntityData.Children[types.GetSegmentPath(&location.AllOptions[i])] = types.YChild{"AllOptions", &location.AllOptions[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location_name"] = types.YLeaf{"LocationName", location.LocationName}
    return &(location.EntityData)
}

// Pm_Pm__Trace_Location_AllOptions
type Pm_Pm__Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Pm_Pm__Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Pm_Pm__Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Pm_Pm__Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + "[option='" + fmt.Sprintf("%v", allOptions.Option) + "']"
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = make(map[string]types.YChild)
    allOptions.EntityData.Children["trace-blocks"] = types.YChild{"TraceBlocks", nil}
    for i := range allOptions.TraceBlocks {
        allOptions.EntityData.Children[types.GetSegmentPath(&allOptions.TraceBlocks[i])] = types.YChild{"TraceBlocks", &allOptions.TraceBlocks[i]}
    }
    allOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    allOptions.EntityData.Leafs["option"] = types.YLeaf{"Option", allOptions.Option}
    return &(allOptions.EntityData)
}

// Pm_Pm__Trace_Location_AllOptions_TraceBlocks
type Pm_Pm__Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Pm_Pm__Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks"
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = make(map[string]types.YChild)
    traceBlocks.EntityData.Leafs = make(map[string]types.YLeaf)
    traceBlocks.EntityData.Leafs["data"] = types.YLeaf{"Data", traceBlocks.Data}
    return &(traceBlocks.EntityData)
}

