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
    QFilesystem []CiscoPlatformSoftware_QFilesystem
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

    ciscoPlatformSoftware.EntityData.Children = make(map[string]types.YChild)
    ciscoPlatformSoftware.EntityData.Children["control-processes"] = types.YChild{"ControlProcesses", &ciscoPlatformSoftware.ControlProcesses}
    ciscoPlatformSoftware.EntityData.Children["q-filesystem"] = types.YChild{"QFilesystem", nil}
    for i := range ciscoPlatformSoftware.QFilesystem {
        ciscoPlatformSoftware.EntityData.Children[types.GetSegmentPath(&ciscoPlatformSoftware.QFilesystem[i])] = types.YChild{"QFilesystem", &ciscoPlatformSoftware.QFilesystem[i]}
    }
    ciscoPlatformSoftware.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoPlatformSoftware.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses
// Information about control processes
type CiscoPlatformSoftware_ControlProcesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of control processes. The type is slice of
    // CiscoPlatformSoftware_ControlProcesses_ControlProcess.
    ControlProcess []CiscoPlatformSoftware_ControlProcesses_ControlProcess
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

    controlProcesses.EntityData.Children = make(map[string]types.YChild)
    controlProcesses.EntityData.Children["control-process"] = types.YChild{"ControlProcess", nil}
    for i := range controlProcesses.ControlProcess {
        controlProcesses.EntityData.Children[types.GetSegmentPath(&controlProcesses.ControlProcess[i])] = types.YChild{"ControlProcess", &controlProcesses.ControlProcess[i]}
    }
    controlProcesses.EntityData.Leafs = make(map[string]types.YLeaf)
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
    controlProcess.EntityData.SegmentPath = "control-process" + "[fru='" + fmt.Sprintf("%v", controlProcess.Fru) + "']" + "[slotnum='" + fmt.Sprintf("%v", controlProcess.Slotnum) + "']" + "[baynum='" + fmt.Sprintf("%v", controlProcess.Baynum) + "']" + "[chassisnum='" + fmt.Sprintf("%v", controlProcess.Chassisnum) + "']"
    controlProcess.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    controlProcess.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    controlProcess.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    controlProcess.EntityData.Children = make(map[string]types.YChild)
    controlProcess.EntityData.Children["load-average-stats"] = types.YChild{"LoadAverageStats", &controlProcess.LoadAverageStats}
    controlProcess.EntityData.Children["load-avg-minutes"] = types.YChild{"LoadAvgMinutes", &controlProcess.LoadAvgMinutes}
    controlProcess.EntityData.Children["memory-stats"] = types.YChild{"MemoryStats", &controlProcess.MemoryStats}
    controlProcess.EntityData.Children["per-core-stats"] = types.YChild{"PerCoreStats", &controlProcess.PerCoreStats}
    controlProcess.EntityData.Leafs = make(map[string]types.YLeaf)
    controlProcess.EntityData.Leafs["fru"] = types.YLeaf{"Fru", controlProcess.Fru}
    controlProcess.EntityData.Leafs["slotnum"] = types.YLeaf{"Slotnum", controlProcess.Slotnum}
    controlProcess.EntityData.Leafs["baynum"] = types.YLeaf{"Baynum", controlProcess.Baynum}
    controlProcess.EntityData.Leafs["chassisnum"] = types.YLeaf{"Chassisnum", controlProcess.Chassisnum}
    controlProcess.EntityData.Leafs["control-process-status"] = types.YLeaf{"ControlProcessStatus", controlProcess.ControlProcessStatus}
    controlProcess.EntityData.Leafs["updated"] = types.YLeaf{"Updated", controlProcess.Updated}
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

    loadAverageStats.EntityData.Children = make(map[string]types.YChild)
    loadAverageStats.EntityData.Leafs = make(map[string]types.YLeaf)
    loadAverageStats.EntityData.Leafs["load-average-status"] = types.YLeaf{"LoadAverageStatus", loadAverageStats.LoadAverageStatus}
    return &(loadAverageStats.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes
// Load average statistics calculated over a period of time
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Load averages based on a time frame. The type is slice of
    // CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute.
    LoadAvgMinute []CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute
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

    loadAvgMinutes.EntityData.Children = make(map[string]types.YChild)
    loadAvgMinutes.EntityData.Children["load-avg-minute"] = types.YChild{"LoadAvgMinute", nil}
    for i := range loadAvgMinutes.LoadAvgMinute {
        loadAvgMinutes.EntityData.Children[types.GetSegmentPath(&loadAvgMinutes.LoadAvgMinute[i])] = types.YChild{"LoadAvgMinute", &loadAvgMinutes.LoadAvgMinute[i]}
    }
    loadAvgMinutes.EntityData.Leafs = make(map[string]types.YLeaf)
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
    loadAvgMinute.EntityData.SegmentPath = "load-avg-minute" + "[number='" + fmt.Sprintf("%v", loadAvgMinute.Number) + "']"
    loadAvgMinute.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    loadAvgMinute.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    loadAvgMinute.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    loadAvgMinute.EntityData.Children = make(map[string]types.YChild)
    loadAvgMinute.EntityData.Children["status"] = types.YChild{"Status", &loadAvgMinute.Status}
    loadAvgMinute.EntityData.Leafs = make(map[string]types.YLeaf)
    loadAvgMinute.EntityData.Leafs["number"] = types.YLeaf{"Number", loadAvgMinute.Number}
    loadAvgMinute.EntityData.Leafs["average"] = types.YLeaf{"Average", loadAvgMinute.Average}
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

    status.EntityData.Children = make(map[string]types.YChild)
    status.EntityData.Leafs = make(map[string]types.YLeaf)
    status.EntityData.Leafs["condition"] = types.YLeaf{"Condition", status.Condition}
    status.EntityData.Leafs["threshold-status"] = types.YLeaf{"ThresholdStatus", status.ThresholdStatus}
    status.EntityData.Leafs["threshold-value"] = types.YLeaf{"ThresholdValue", status.ThresholdValue}
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

    memoryStats.EntityData.Children = make(map[string]types.YChild)
    memoryStats.EntityData.Children["status"] = types.YChild{"Status", &memoryStats.Status}
    memoryStats.EntityData.Leafs = make(map[string]types.YLeaf)
    memoryStats.EntityData.Leafs["memory-status"] = types.YLeaf{"MemoryStatus", memoryStats.MemoryStatus}
    memoryStats.EntityData.Leafs["total"] = types.YLeaf{"Total", memoryStats.Total}
    memoryStats.EntityData.Leafs["used-number"] = types.YLeaf{"UsedNumber", memoryStats.UsedNumber}
    memoryStats.EntityData.Leafs["used-percent"] = types.YLeaf{"UsedPercent", memoryStats.UsedPercent}
    memoryStats.EntityData.Leafs["free-number"] = types.YLeaf{"FreeNumber", memoryStats.FreeNumber}
    memoryStats.EntityData.Leafs["free-percent"] = types.YLeaf{"FreePercent", memoryStats.FreePercent}
    memoryStats.EntityData.Leafs["available-number"] = types.YLeaf{"AvailableNumber", memoryStats.AvailableNumber}
    memoryStats.EntityData.Leafs["available-percent"] = types.YLeaf{"AvailablePercent", memoryStats.AvailablePercent}
    memoryStats.EntityData.Leafs["committed-number"] = types.YLeaf{"CommittedNumber", memoryStats.CommittedNumber}
    memoryStats.EntityData.Leafs["committed-percent"] = types.YLeaf{"CommittedPercent", memoryStats.CommittedPercent}
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

    status.EntityData.Children = make(map[string]types.YChild)
    status.EntityData.Leafs = make(map[string]types.YLeaf)
    status.EntityData.Leafs["warning-threshold-percent"] = types.YLeaf{"WarningThresholdPercent", status.WarningThresholdPercent}
    status.EntityData.Leafs["critical-threshold-percent"] = types.YLeaf{"CriticalThresholdPercent", status.CriticalThresholdPercent}
    return &(status.EntityData)
}

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats
// Processor core statistics
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of processor cores. The type is slice of
    // CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat.
    PerCoreStat []CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat
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

    perCoreStats.EntityData.Children = make(map[string]types.YChild)
    perCoreStats.EntityData.Children["per-core-stat"] = types.YChild{"PerCoreStat", nil}
    for i := range perCoreStats.PerCoreStat {
        perCoreStats.EntityData.Children[types.GetSegmentPath(&perCoreStats.PerCoreStat[i])] = types.YChild{"PerCoreStat", &perCoreStats.PerCoreStat[i]}
    }
    perCoreStats.EntityData.Leafs = make(map[string]types.YLeaf)
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
    perCoreStat.EntityData.SegmentPath = "per-core-stat" + "[name='" + fmt.Sprintf("%v", perCoreStat.Name) + "']"
    perCoreStat.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    perCoreStat.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    perCoreStat.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    perCoreStat.EntityData.Children = make(map[string]types.YChild)
    perCoreStat.EntityData.Leafs = make(map[string]types.YLeaf)
    perCoreStat.EntityData.Leafs["name"] = types.YLeaf{"Name", perCoreStat.Name}
    perCoreStat.EntityData.Leafs["user"] = types.YLeaf{"User", perCoreStat.User}
    perCoreStat.EntityData.Leafs["system"] = types.YLeaf{"System", perCoreStat.System}
    perCoreStat.EntityData.Leafs["nice"] = types.YLeaf{"Nice", perCoreStat.Nice}
    perCoreStat.EntityData.Leafs["idle"] = types.YLeaf{"Idle", perCoreStat.Idle}
    perCoreStat.EntityData.Leafs["irq"] = types.YLeaf{"Irq", perCoreStat.Irq}
    perCoreStat.EntityData.Leafs["sirq"] = types.YLeaf{"Sirq", perCoreStat.Sirq}
    perCoreStat.EntityData.Leafs["io-wait"] = types.YLeaf{"IoWait", perCoreStat.IoWait}
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
    Partitions []CiscoPlatformSoftware_QFilesystem_Partitions

    // Information about core files. The type is slice of
    // CiscoPlatformSoftware_QFilesystem_CoreFiles.
    CoreFiles []CiscoPlatformSoftware_QFilesystem_CoreFiles
}

func (qFilesystem *CiscoPlatformSoftware_QFilesystem) GetEntityData() *types.CommonEntityData {
    qFilesystem.EntityData.YFilter = qFilesystem.YFilter
    qFilesystem.EntityData.YangName = "q-filesystem"
    qFilesystem.EntityData.BundleName = "cisco_ios_xe"
    qFilesystem.EntityData.ParentYangName = "cisco-platform-software"
    qFilesystem.EntityData.SegmentPath = "q-filesystem" + "[fru='" + fmt.Sprintf("%v", qFilesystem.Fru) + "']" + "[slotnum='" + fmt.Sprintf("%v", qFilesystem.Slotnum) + "']" + "[baynum='" + fmt.Sprintf("%v", qFilesystem.Baynum) + "']" + "[chassisnum='" + fmt.Sprintf("%v", qFilesystem.Chassisnum) + "']"
    qFilesystem.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qFilesystem.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qFilesystem.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qFilesystem.EntityData.Children = make(map[string]types.YChild)
    qFilesystem.EntityData.Children["partitions"] = types.YChild{"Partitions", nil}
    for i := range qFilesystem.Partitions {
        qFilesystem.EntityData.Children[types.GetSegmentPath(&qFilesystem.Partitions[i])] = types.YChild{"Partitions", &qFilesystem.Partitions[i]}
    }
    qFilesystem.EntityData.Children["core-files"] = types.YChild{"CoreFiles", nil}
    for i := range qFilesystem.CoreFiles {
        qFilesystem.EntityData.Children[types.GetSegmentPath(&qFilesystem.CoreFiles[i])] = types.YChild{"CoreFiles", &qFilesystem.CoreFiles[i]}
    }
    qFilesystem.EntityData.Leafs = make(map[string]types.YLeaf)
    qFilesystem.EntityData.Leafs["fru"] = types.YLeaf{"Fru", qFilesystem.Fru}
    qFilesystem.EntityData.Leafs["slotnum"] = types.YLeaf{"Slotnum", qFilesystem.Slotnum}
    qFilesystem.EntityData.Leafs["baynum"] = types.YLeaf{"Baynum", qFilesystem.Baynum}
    qFilesystem.EntityData.Leafs["chassisnum"] = types.YLeaf{"Chassisnum", qFilesystem.Chassisnum}
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
    partitions.EntityData.SegmentPath = "partitions" + "[name='" + fmt.Sprintf("%v", partitions.Name) + "']"
    partitions.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    partitions.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    partitions.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    partitions.EntityData.Children = make(map[string]types.YChild)
    partitions.EntityData.Leafs = make(map[string]types.YLeaf)
    partitions.EntityData.Leafs["name"] = types.YLeaf{"Name", partitions.Name}
    partitions.EntityData.Leafs["total-size"] = types.YLeaf{"TotalSize", partitions.TotalSize}
    partitions.EntityData.Leafs["used-size"] = types.YLeaf{"UsedSize", partitions.UsedSize}
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
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    Time interface{}
}

func (coreFiles *CiscoPlatformSoftware_QFilesystem_CoreFiles) GetEntityData() *types.CommonEntityData {
    coreFiles.EntityData.YFilter = coreFiles.YFilter
    coreFiles.EntityData.YangName = "core-files"
    coreFiles.EntityData.BundleName = "cisco_ios_xe"
    coreFiles.EntityData.ParentYangName = "q-filesystem"
    coreFiles.EntityData.SegmentPath = "core-files" + "[filename='" + fmt.Sprintf("%v", coreFiles.Filename) + "']"
    coreFiles.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    coreFiles.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    coreFiles.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    coreFiles.EntityData.Children = make(map[string]types.YChild)
    coreFiles.EntityData.Leafs = make(map[string]types.YLeaf)
    coreFiles.EntityData.Leafs["filename"] = types.YLeaf{"Filename", coreFiles.Filename}
    coreFiles.EntityData.Leafs["time"] = types.YLeaf{"Time", coreFiles.Time}
    return &(coreFiles.EntityData)
}

