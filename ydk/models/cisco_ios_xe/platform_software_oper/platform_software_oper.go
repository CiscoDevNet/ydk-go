// This module contains a collection of YANG definitions
// for monitoring platform software in a Network Element
package platform_software_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package platform_software_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-platform-software-oper cisco-platform-software}", reflect.TypeOf(CiscoPlatformSoftware{}))
    ydk.RegisterEntity("Cisco-IOS-XE-platform-software-oper:cisco-platform-software", reflect.TypeOf(CiscoPlatformSoftware{}))
}

// BFru represents FRU type
type BFru string

const (
    BFru_platform_fru_rp BFru = "platform-fru-rp"

    BFru_platform_fru_fp BFru = "platform-fru-fp"

    BFru_platform_fru_cc BFru = "platform-fru-cc"

    BFru_platform_fru_max BFru = "platform-fru-max"
)

// CiscoPlatformSoftware
// Cisco platform software information
type CiscoPlatformSoftware struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about control processes.
    ControlProcesses CiscoPlatformSoftware_ControlProcesses

    // Information about the filesystem. The type is slice of
    // CiscoPlatformSoftware_QFilesystem.
    QFilesystem []*CiscoPlatformSoftware_QFilesystem
}

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetEntityData() *types.CommonEntityData {
    ciscoPlatformSoftware.EntityData.YFilter = ciscoPlatformSoftware.YFilter
    ciscoPlatformSoftware.EntityData.YangName = "cisco-platform-software"
    ciscoPlatformSoftware.EntityData.BundleName = "cisco_ios_xe"
    ciscoPlatformSoftware.EntityData.ParentYangName = "Cisco-IOS-XE-platform-software-oper"
    ciscoPlatformSoftware.EntityData.SegmentPath = "Cisco-IOS-XE-platform-software-oper:cisco-platform-software"
    ciscoPlatformSoftware.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoPlatformSoftware.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoPlatformSoftware.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoPlatformSoftware.EntityData.Children = types.NewOrderedMap()
    ciscoPlatformSoftware.EntityData.Children.Append("control-processes", types.YChild{"ControlProcesses", &ciscoPlatformSoftware.ControlProcesses})
    ciscoPlatformSoftware.EntityData.Children.Append("q-filesystem", types.YChild{"QFilesystem", nil})
    for i := range ciscoPlatformSoftware.QFilesystem {
        ciscoPlatformSoftware.EntityData.Children.Append(types.GetSegmentPath(ciscoPlatformSoftware.QFilesystem[i]), types.YChild{"QFilesystem", ciscoPlatformSoftware.QFilesystem[i]})
    }
    ciscoPlatformSoftware.EntityData.Leafs = types.NewOrderedMap()

    ciscoPlatformSoftware.EntityData.YListKeys = []string {}

    return &(ciscoPlatformSoftware.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses
// Information about control processes
type CiscoPlatformSoftware_ControlProcesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of control processes. The type is slice of
    // CiscoPlatformSoftware_ControlProcesses_ControlProcess.
    ControlProcess []*CiscoPlatformSoftware_ControlProcesses_ControlProcess
}

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetEntityData() *types.CommonEntityData {
    controlProcesses.EntityData.YFilter = controlProcesses.YFilter
    controlProcesses.EntityData.YangName = "control-processes"
    controlProcesses.EntityData.BundleName = "cisco_ios_xe"
    controlProcesses.EntityData.ParentYangName = "cisco-platform-software"
    controlProcesses.EntityData.SegmentPath = "control-processes"
    controlProcesses.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    controlProcesses.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    controlProcesses.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    controlProcesses.EntityData.Children = types.NewOrderedMap()
    controlProcesses.EntityData.Children.Append("control-process", types.YChild{"ControlProcess", nil})
    for i := range controlProcesses.ControlProcess {
        controlProcesses.EntityData.Children.Append(types.GetSegmentPath(controlProcesses.ControlProcess[i]), types.YChild{"ControlProcess", controlProcesses.ControlProcess[i]})
    }
    controlProcesses.EntityData.Leafs = types.NewOrderedMap()

    controlProcesses.EntityData.YListKeys = []string {}

    return &(controlProcesses.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess
// The list of control processes
type CiscoPlatformSoftware_ControlProcesses_ControlProcess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Field replaceable unit. The type is BFru.
    Fru interface{}

    // This attribute is a key. Slot number. The type is interface{} with range:
    // -32768..32767.
    Slotnum interface{}

    // This attribute is a key. Bay number. The type is interface{} with range:
    // -32768..32767.
    Baynum interface{}

    // This attribute is a key. Chassis number. The type is interface{} with
    // range: -32768..32767.
    Chassisnum interface{}

    // Status of the control process. The type is string.
    ControlProcessStatus interface{}

    // Number of seconds since the data has been updated. The type is interface{}
    // with range: 0..18446744073709551615.
    Updated interface{}

    // Load average statictics.
    LoadAverageStats CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats

    // Load average statistics calculated over a period of time.
    LoadAvgMinutes CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes

    // Memory statistics.
    MemoryStats CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats

    // Processor core statistics.
    PerCoreStats CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats
}

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetEntityData() *types.CommonEntityData {
    controlProcess.EntityData.YFilter = controlProcess.YFilter
    controlProcess.EntityData.YangName = "control-process"
    controlProcess.EntityData.BundleName = "cisco_ios_xe"
    controlProcess.EntityData.ParentYangName = "control-processes"
    controlProcess.EntityData.SegmentPath = "control-process" + types.AddKeyToken(controlProcess.Fru, "fru") + types.AddKeyToken(controlProcess.Slotnum, "slotnum") + types.AddKeyToken(controlProcess.Baynum, "baynum") + types.AddKeyToken(controlProcess.Chassisnum, "chassisnum")
    controlProcess.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    controlProcess.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    controlProcess.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    controlProcess.EntityData.Children = types.NewOrderedMap()
    controlProcess.EntityData.Children.Append("load-average-stats", types.YChild{"LoadAverageStats", &controlProcess.LoadAverageStats})
    controlProcess.EntityData.Children.Append("load-avg-minutes", types.YChild{"LoadAvgMinutes", &controlProcess.LoadAvgMinutes})
    controlProcess.EntityData.Children.Append("memory-stats", types.YChild{"MemoryStats", &controlProcess.MemoryStats})
    controlProcess.EntityData.Children.Append("per-core-stats", types.YChild{"PerCoreStats", &controlProcess.PerCoreStats})
    controlProcess.EntityData.Leafs = types.NewOrderedMap()
    controlProcess.EntityData.Leafs.Append("fru", types.YLeaf{"Fru", controlProcess.Fru})
    controlProcess.EntityData.Leafs.Append("slotnum", types.YLeaf{"Slotnum", controlProcess.Slotnum})
    controlProcess.EntityData.Leafs.Append("baynum", types.YLeaf{"Baynum", controlProcess.Baynum})
    controlProcess.EntityData.Leafs.Append("chassisnum", types.YLeaf{"Chassisnum", controlProcess.Chassisnum})
    controlProcess.EntityData.Leafs.Append("control-process-status", types.YLeaf{"ControlProcessStatus", controlProcess.ControlProcessStatus})
    controlProcess.EntityData.Leafs.Append("updated", types.YLeaf{"Updated", controlProcess.Updated})

    controlProcess.EntityData.YListKeys = []string {"Fru", "Slotnum", "Baynum", "Chassisnum"}

    return &(controlProcess.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats
// Load average statictics
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Load average status. The type is string.
    LoadAverageStatus interface{}
}

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetEntityData() *types.CommonEntityData {
    loadAverageStats.EntityData.YFilter = loadAverageStats.YFilter
    loadAverageStats.EntityData.YangName = "load-average-stats"
    loadAverageStats.EntityData.BundleName = "cisco_ios_xe"
    loadAverageStats.EntityData.ParentYangName = "control-process"
    loadAverageStats.EntityData.SegmentPath = "load-average-stats"
    loadAverageStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    loadAverageStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    loadAverageStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    loadAverageStats.EntityData.Children = types.NewOrderedMap()
    loadAverageStats.EntityData.Leafs = types.NewOrderedMap()
    loadAverageStats.EntityData.Leafs.Append("load-average-status", types.YLeaf{"LoadAverageStatus", loadAverageStats.LoadAverageStatus})

    loadAverageStats.EntityData.YListKeys = []string {}

    return &(loadAverageStats.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes
// Load average statistics calculated over a period of time
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Load averages based on a time frame. The type is slice of
    // CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute.
    LoadAvgMinute []*CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute
}

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetEntityData() *types.CommonEntityData {
    loadAvgMinutes.EntityData.YFilter = loadAvgMinutes.YFilter
    loadAvgMinutes.EntityData.YangName = "load-avg-minutes"
    loadAvgMinutes.EntityData.BundleName = "cisco_ios_xe"
    loadAvgMinutes.EntityData.ParentYangName = "control-process"
    loadAvgMinutes.EntityData.SegmentPath = "load-avg-minutes"
    loadAvgMinutes.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    loadAvgMinutes.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    loadAvgMinutes.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    loadAvgMinutes.EntityData.Children = types.NewOrderedMap()
    loadAvgMinutes.EntityData.Children.Append("load-avg-minute", types.YChild{"LoadAvgMinute", nil})
    for i := range loadAvgMinutes.LoadAvgMinute {
        loadAvgMinutes.EntityData.Children.Append(types.GetSegmentPath(loadAvgMinutes.LoadAvgMinute[i]), types.YChild{"LoadAvgMinute", loadAvgMinutes.LoadAvgMinute[i]})
    }
    loadAvgMinutes.EntityData.Leafs = types.NewOrderedMap()

    loadAvgMinutes.EntityData.YListKeys = []string {}

    return &(loadAvgMinutes.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute
// List of Load averages based on a time frame
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The number of minutes the average was calculated
    // on. The type is interface{} with range: 0..18446744073709551615.
    Number interface{}

    // Calculated average. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Average interface{}

    // Load average statistics minute status.
    Status CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status
}

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetEntityData() *types.CommonEntityData {
    loadAvgMinute.EntityData.YFilter = loadAvgMinute.YFilter
    loadAvgMinute.EntityData.YangName = "load-avg-minute"
    loadAvgMinute.EntityData.BundleName = "cisco_ios_xe"
    loadAvgMinute.EntityData.ParentYangName = "load-avg-minutes"
    loadAvgMinute.EntityData.SegmentPath = "load-avg-minute" + types.AddKeyToken(loadAvgMinute.Number, "number")
    loadAvgMinute.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    loadAvgMinute.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    loadAvgMinute.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    loadAvgMinute.EntityData.Children = types.NewOrderedMap()
    loadAvgMinute.EntityData.Children.Append("status", types.YChild{"Status", &loadAvgMinute.Status})
    loadAvgMinute.EntityData.Leafs = types.NewOrderedMap()
    loadAvgMinute.EntityData.Leafs.Append("number", types.YLeaf{"Number", loadAvgMinute.Number})
    loadAvgMinute.EntityData.Leafs.Append("average", types.YLeaf{"Average", loadAvgMinute.Average})

    loadAvgMinute.EntityData.YListKeys = []string {"Number"}

    return &(loadAvgMinute.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status
// Load average statistics minute status
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Load average condition. The type is string.
    Condition interface{}

    // Load average status. The type is string.
    ThresholdStatus interface{}

    // Load average threshold. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    ThresholdValue interface{}
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xe"
    status.EntityData.ParentYangName = "load-avg-minute"
    status.EntityData.SegmentPath = "status"
    status.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    status.EntityData.Children = types.NewOrderedMap()
    status.EntityData.Leafs = types.NewOrderedMap()
    status.EntityData.Leafs.Append("condition", types.YLeaf{"Condition", status.Condition})
    status.EntityData.Leafs.Append("threshold-status", types.YLeaf{"ThresholdStatus", status.ThresholdStatus})
    status.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", status.ThresholdValue})

    status.EntityData.YListKeys = []string {}

    return &(status.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats
// Memory statistics
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The status of the memory. The type is string.
    MemoryStatus interface{}

    // The total amount of memory in kb. The type is interface{} with range:
    // 0..18446744073709551615.
    Total interface{}

    // The amount of memory being used in kb. The type is interface{} with range:
    // 0..18446744073709551615.
    UsedNumber interface{}

    // The percentage of memory being used. The type is interface{} with range:
    // 0..18446744073709551615.
    UsedPercent interface{}

    // The amount of free memory in kb. The type is interface{} with range:
    // 0..18446744073709551615.
    FreeNumber interface{}

    // The percentage of free memory. The type is interface{} with range:
    // 0..18446744073709551615.
    FreePercent interface{}

    // The amount of available memory in kb. The type is interface{} with range:
    // 0..18446744073709551615.
    AvailableNumber interface{}

    // The percentage of available memory. The type is interface{} with range:
    // 0..18446744073709551615.
    AvailablePercent interface{}

    // The amount of committed memory in kb. The type is interface{} with range:
    // 0..18446744073709551615.
    CommittedNumber interface{}

    // The percentage of committed memory. The type is interface{} with range:
    // 0..255.
    CommittedPercent interface{}

    // Memory status.
    Status CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status
}

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetEntityData() *types.CommonEntityData {
    memoryStats.EntityData.YFilter = memoryStats.YFilter
    memoryStats.EntityData.YangName = "memory-stats"
    memoryStats.EntityData.BundleName = "cisco_ios_xe"
    memoryStats.EntityData.ParentYangName = "control-process"
    memoryStats.EntityData.SegmentPath = "memory-stats"
    memoryStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    memoryStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    memoryStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    memoryStats.EntityData.Children = types.NewOrderedMap()
    memoryStats.EntityData.Children.Append("status", types.YChild{"Status", &memoryStats.Status})
    memoryStats.EntityData.Leafs = types.NewOrderedMap()
    memoryStats.EntityData.Leafs.Append("memory-status", types.YLeaf{"MemoryStatus", memoryStats.MemoryStatus})
    memoryStats.EntityData.Leafs.Append("total", types.YLeaf{"Total", memoryStats.Total})
    memoryStats.EntityData.Leafs.Append("used-number", types.YLeaf{"UsedNumber", memoryStats.UsedNumber})
    memoryStats.EntityData.Leafs.Append("used-percent", types.YLeaf{"UsedPercent", memoryStats.UsedPercent})
    memoryStats.EntityData.Leafs.Append("free-number", types.YLeaf{"FreeNumber", memoryStats.FreeNumber})
    memoryStats.EntityData.Leafs.Append("free-percent", types.YLeaf{"FreePercent", memoryStats.FreePercent})
    memoryStats.EntityData.Leafs.Append("available-number", types.YLeaf{"AvailableNumber", memoryStats.AvailableNumber})
    memoryStats.EntityData.Leafs.Append("available-percent", types.YLeaf{"AvailablePercent", memoryStats.AvailablePercent})
    memoryStats.EntityData.Leafs.Append("committed-number", types.YLeaf{"CommittedNumber", memoryStats.CommittedNumber})
    memoryStats.EntityData.Leafs.Append("committed-percent", types.YLeaf{"CommittedPercent", memoryStats.CommittedPercent})

    memoryStats.EntityData.YListKeys = []string {}

    return &(memoryStats.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status
// Memory status
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Memory warning threshold value percent. The type is interface{} with range:
    // 0..4294967295.
    WarningThresholdPercent interface{}

    // Memory critical threshold value percent. The type is interface{} with
    // range: 0..4294967295.
    CriticalThresholdPercent interface{}
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xe"
    status.EntityData.ParentYangName = "memory-stats"
    status.EntityData.SegmentPath = "status"
    status.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    status.EntityData.Children = types.NewOrderedMap()
    status.EntityData.Leafs = types.NewOrderedMap()
    status.EntityData.Leafs.Append("warning-threshold-percent", types.YLeaf{"WarningThresholdPercent", status.WarningThresholdPercent})
    status.EntityData.Leafs.Append("critical-threshold-percent", types.YLeaf{"CriticalThresholdPercent", status.CriticalThresholdPercent})

    status.EntityData.YListKeys = []string {}

    return &(status.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats
// Processor core statistics
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of processor cores. The type is slice of
    // CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat.
    PerCoreStat []*CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat
}

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetEntityData() *types.CommonEntityData {
    perCoreStats.EntityData.YFilter = perCoreStats.YFilter
    perCoreStats.EntityData.YangName = "per-core-stats"
    perCoreStats.EntityData.BundleName = "cisco_ios_xe"
    perCoreStats.EntityData.ParentYangName = "control-process"
    perCoreStats.EntityData.SegmentPath = "per-core-stats"
    perCoreStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    perCoreStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    perCoreStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    perCoreStats.EntityData.Children = types.NewOrderedMap()
    perCoreStats.EntityData.Children.Append("per-core-stat", types.YChild{"PerCoreStat", nil})
    for i := range perCoreStats.PerCoreStat {
        perCoreStats.EntityData.Children.Append(types.GetSegmentPath(perCoreStats.PerCoreStat[i]), types.YChild{"PerCoreStat", perCoreStats.PerCoreStat[i]})
    }
    perCoreStats.EntityData.Leafs = types.NewOrderedMap()

    perCoreStats.EntityData.YListKeys = []string {}

    return &(perCoreStats.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat
// List of processor cores
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier of the core. The type is
    // interface{} with range: 0..4294967295.
    Name interface{}

    // CPU utilization in user mode. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    User interface{}

    // CPU utilization in system mode. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    System interface{}

    // Nice level. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Nice interface{}

    // Idle percentage. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Idle interface{}

    // The percentage of utilization by irq. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Irq interface{}

    // The percentage of utilization by sirq. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    Sirq interface{}

    // IOWait percentage. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    IoWait interface{}
}

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetEntityData() *types.CommonEntityData {
    perCoreStat.EntityData.YFilter = perCoreStat.YFilter
    perCoreStat.EntityData.YangName = "per-core-stat"
    perCoreStat.EntityData.BundleName = "cisco_ios_xe"
    perCoreStat.EntityData.ParentYangName = "per-core-stats"
    perCoreStat.EntityData.SegmentPath = "per-core-stat" + types.AddKeyToken(perCoreStat.Name, "name")
    perCoreStat.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    perCoreStat.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    perCoreStat.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    perCoreStat.EntityData.Children = types.NewOrderedMap()
    perCoreStat.EntityData.Leafs = types.NewOrderedMap()
    perCoreStat.EntityData.Leafs.Append("name", types.YLeaf{"Name", perCoreStat.Name})
    perCoreStat.EntityData.Leafs.Append("user", types.YLeaf{"User", perCoreStat.User})
    perCoreStat.EntityData.Leafs.Append("system", types.YLeaf{"System", perCoreStat.System})
    perCoreStat.EntityData.Leafs.Append("nice", types.YLeaf{"Nice", perCoreStat.Nice})
    perCoreStat.EntityData.Leafs.Append("idle", types.YLeaf{"Idle", perCoreStat.Idle})
    perCoreStat.EntityData.Leafs.Append("irq", types.YLeaf{"Irq", perCoreStat.Irq})
    perCoreStat.EntityData.Leafs.Append("sirq", types.YLeaf{"Sirq", perCoreStat.Sirq})
    perCoreStat.EntityData.Leafs.Append("io-wait", types.YLeaf{"IoWait", perCoreStat.IoWait})

    perCoreStat.EntityData.YListKeys = []string {"Name"}

    return &(perCoreStat.EntityData)
}

// CiscoPlatformSoftware_QFilesystem
// Information about the filesystem
type CiscoPlatformSoftware_QFilesystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Field replaceable unit. The type is BFru.
    Fru interface{}

    // This attribute is a key. Slot number. The type is interface{} with range:
    // -32768..32767.
    Slotnum interface{}

    // This attribute is a key. Bay number. The type is interface{} with range:
    // -32768..32767.
    Baynum interface{}

    // This attribute is a key. Chassis number. The type is interface{} with
    // range: -32768..32767.
    Chassisnum interface{}

    // Information about partitions. The type is slice of
    // CiscoPlatformSoftware_QFilesystem_Partitions.
    Partitions []*CiscoPlatformSoftware_QFilesystem_Partitions

    // Information about core files. The type is slice of
    // CiscoPlatformSoftware_QFilesystem_CoreFiles.
    CoreFiles []*CiscoPlatformSoftware_QFilesystem_CoreFiles
}

func (qFilesystem *CiscoPlatformSoftware_QFilesystem) GetEntityData() *types.CommonEntityData {
    qFilesystem.EntityData.YFilter = qFilesystem.YFilter
    qFilesystem.EntityData.YangName = "q-filesystem"
    qFilesystem.EntityData.BundleName = "cisco_ios_xe"
    qFilesystem.EntityData.ParentYangName = "cisco-platform-software"
    qFilesystem.EntityData.SegmentPath = "q-filesystem" + types.AddKeyToken(qFilesystem.Fru, "fru") + types.AddKeyToken(qFilesystem.Slotnum, "slotnum") + types.AddKeyToken(qFilesystem.Baynum, "baynum") + types.AddKeyToken(qFilesystem.Chassisnum, "chassisnum")
    qFilesystem.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qFilesystem.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qFilesystem.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qFilesystem.EntityData.Children = types.NewOrderedMap()
    qFilesystem.EntityData.Children.Append("partitions", types.YChild{"Partitions", nil})
    for i := range qFilesystem.Partitions {
        qFilesystem.EntityData.Children.Append(types.GetSegmentPath(qFilesystem.Partitions[i]), types.YChild{"Partitions", qFilesystem.Partitions[i]})
    }
    qFilesystem.EntityData.Children.Append("core-files", types.YChild{"CoreFiles", nil})
    for i := range qFilesystem.CoreFiles {
        qFilesystem.EntityData.Children.Append(types.GetSegmentPath(qFilesystem.CoreFiles[i]), types.YChild{"CoreFiles", qFilesystem.CoreFiles[i]})
    }
    qFilesystem.EntityData.Leafs = types.NewOrderedMap()
    qFilesystem.EntityData.Leafs.Append("fru", types.YLeaf{"Fru", qFilesystem.Fru})
    qFilesystem.EntityData.Leafs.Append("slotnum", types.YLeaf{"Slotnum", qFilesystem.Slotnum})
    qFilesystem.EntityData.Leafs.Append("baynum", types.YLeaf{"Baynum", qFilesystem.Baynum})
    qFilesystem.EntityData.Leafs.Append("chassisnum", types.YLeaf{"Chassisnum", qFilesystem.Chassisnum})

    qFilesystem.EntityData.YListKeys = []string {"Fru", "Slotnum", "Baynum", "Chassisnum"}

    return &(qFilesystem.EntityData)
}

// CiscoPlatformSoftware_QFilesystem_Partitions
// Information about partitions
type CiscoPlatformSoftware_QFilesystem_Partitions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the partition. The type is string.
    Name interface{}

    // Total size of the partition in Kilobytes. The type is interface{} with
    // range: 0..18446744073709551615.
    TotalSize interface{}

    // Size used in Kilobytes. The type is interface{} with range:
    // 0..18446744073709551615.
    UsedSize interface{}
}

func (partitions *CiscoPlatformSoftware_QFilesystem_Partitions) GetEntityData() *types.CommonEntityData {
    partitions.EntityData.YFilter = partitions.YFilter
    partitions.EntityData.YangName = "partitions"
    partitions.EntityData.BundleName = "cisco_ios_xe"
    partitions.EntityData.ParentYangName = "q-filesystem"
    partitions.EntityData.SegmentPath = "partitions" + types.AddKeyToken(partitions.Name, "name")
    partitions.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    partitions.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    partitions.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    partitions.EntityData.Children = types.NewOrderedMap()
    partitions.EntityData.Leafs = types.NewOrderedMap()
    partitions.EntityData.Leafs.Append("name", types.YLeaf{"Name", partitions.Name})
    partitions.EntityData.Leafs.Append("total-size", types.YLeaf{"TotalSize", partitions.TotalSize})
    partitions.EntityData.Leafs.Append("used-size", types.YLeaf{"UsedSize", partitions.UsedSize})

    partitions.EntityData.YListKeys = []string {"Name"}

    return &(partitions.EntityData)
}

// CiscoPlatformSoftware_QFilesystem_CoreFiles
// Information about core files
type CiscoPlatformSoftware_QFilesystem_CoreFiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The core filename. The type is string.
    Filename interface{}

    // The date of generation. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Time interface{}
}

func (coreFiles *CiscoPlatformSoftware_QFilesystem_CoreFiles) GetEntityData() *types.CommonEntityData {
    coreFiles.EntityData.YFilter = coreFiles.YFilter
    coreFiles.EntityData.YangName = "core-files"
    coreFiles.EntityData.BundleName = "cisco_ios_xe"
    coreFiles.EntityData.ParentYangName = "q-filesystem"
    coreFiles.EntityData.SegmentPath = "core-files" + types.AddKeyToken(coreFiles.Filename, "filename")
    coreFiles.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    coreFiles.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    coreFiles.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    coreFiles.EntityData.Children = types.NewOrderedMap()
    coreFiles.EntityData.Leafs = types.NewOrderedMap()
    coreFiles.EntityData.Leafs.Append("filename", types.YLeaf{"Filename", coreFiles.Filename})
    coreFiles.EntityData.Leafs.Append("time", types.YLeaf{"Time", coreFiles.Time})

    coreFiles.EntityData.YListKeys = []string {"Filename"}

    return &(coreFiles.EntityData)
}

