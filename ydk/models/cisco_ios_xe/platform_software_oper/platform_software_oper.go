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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about control processes.
    ControlProcesses CiscoPlatformSoftware_ControlProcesses
}

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetFilter() yfilter.YFilter { return ciscoPlatformSoftware.YFilter }

func (ciscoPlatformSoftware *CiscoPlatformSoftware) SetFilter(yf yfilter.YFilter) { ciscoPlatformSoftware.YFilter = yf }

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetGoName(yname string) string {
    if yname == "control-processes" { return "ControlProcesses" }
    return ""
}

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetSegmentPath() string {
    return "Cisco-IOS-XE-platform-software-oper:cisco-platform-software"
}

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "control-processes" {
        return &ciscoPlatformSoftware.ControlProcesses
    }
    return nil
}

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["control-processes"] = &ciscoPlatformSoftware.ControlProcesses
    return children
}

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetYangName() string { return "cisco-platform-software" }

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoPlatformSoftware *CiscoPlatformSoftware) SetParent(parent types.Entity) { ciscoPlatformSoftware.parent = parent }

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetParent() types.Entity { return ciscoPlatformSoftware.parent }

func (ciscoPlatformSoftware *CiscoPlatformSoftware) GetParentYangName() string { return "Cisco-IOS-XE-platform-software-oper" }

// CiscoPlatformSoftware_ControlProcesses
// Information about control processes
type CiscoPlatformSoftware_ControlProcesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of control processes. The type is slice of
    // CiscoPlatformSoftware_ControlProcesses_ControlProcess.
    ControlProcess []CiscoPlatformSoftware_ControlProcesses_ControlProcess
}

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetFilter() yfilter.YFilter { return controlProcesses.YFilter }

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) SetFilter(yf yfilter.YFilter) { controlProcesses.YFilter = yf }

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetGoName(yname string) string {
    if yname == "control-process" { return "ControlProcess" }
    return ""
}

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetSegmentPath() string {
    return "control-processes"
}

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "control-process" {
        for _, c := range controlProcesses.ControlProcess {
            if controlProcesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CiscoPlatformSoftware_ControlProcesses_ControlProcess{}
        controlProcesses.ControlProcess = append(controlProcesses.ControlProcess, child)
        return &controlProcesses.ControlProcess[len(controlProcesses.ControlProcess)-1]
    }
    return nil
}

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range controlProcesses.ControlProcess {
        children[controlProcesses.ControlProcess[i].GetSegmentPath()] = &controlProcesses.ControlProcess[i]
    }
    return children
}

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetBundleName() string { return "cisco_ios_xe" }

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetYangName() string { return "control-processes" }

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) SetParent(parent types.Entity) { controlProcesses.parent = parent }

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetParent() types.Entity { return controlProcesses.parent }

func (controlProcesses *CiscoPlatformSoftware_ControlProcesses) GetParentYangName() string { return "cisco-platform-software" }

// CiscoPlatformSoftware_ControlProcesses_ControlProcess
// The list of control processes
type CiscoPlatformSoftware_ControlProcesses_ControlProcess struct {
    parent types.Entity
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

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetFilter() yfilter.YFilter { return controlProcess.YFilter }

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) SetFilter(yf yfilter.YFilter) { controlProcess.YFilter = yf }

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetGoName(yname string) string {
    if yname == "fru" { return "Fru" }
    if yname == "slotnum" { return "Slotnum" }
    if yname == "baynum" { return "Baynum" }
    if yname == "chassisnum" { return "Chassisnum" }
    if yname == "control-process-status" { return "ControlProcessStatus" }
    if yname == "updated" { return "Updated" }
    if yname == "load-average-stats" { return "LoadAverageStats" }
    if yname == "load-avg-minutes" { return "LoadAvgMinutes" }
    if yname == "memory-stats" { return "MemoryStats" }
    if yname == "per-core-stats" { return "PerCoreStats" }
    return ""
}

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetSegmentPath() string {
    return "control-process" + "[fru='" + fmt.Sprintf("%v", controlProcess.Fru) + "']" + "[slotnum='" + fmt.Sprintf("%v", controlProcess.Slotnum) + "']" + "[baynum='" + fmt.Sprintf("%v", controlProcess.Baynum) + "']" + "[chassisnum='" + fmt.Sprintf("%v", controlProcess.Chassisnum) + "']"
}

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-average-stats" {
        return &controlProcess.LoadAverageStats
    }
    if childYangName == "load-avg-minutes" {
        return &controlProcess.LoadAvgMinutes
    }
    if childYangName == "memory-stats" {
        return &controlProcess.MemoryStats
    }
    if childYangName == "per-core-stats" {
        return &controlProcess.PerCoreStats
    }
    return nil
}

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["load-average-stats"] = &controlProcess.LoadAverageStats
    children["load-avg-minutes"] = &controlProcess.LoadAvgMinutes
    children["memory-stats"] = &controlProcess.MemoryStats
    children["per-core-stats"] = &controlProcess.PerCoreStats
    return children
}

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fru"] = controlProcess.Fru
    leafs["slotnum"] = controlProcess.Slotnum
    leafs["baynum"] = controlProcess.Baynum
    leafs["chassisnum"] = controlProcess.Chassisnum
    leafs["control-process-status"] = controlProcess.ControlProcessStatus
    leafs["updated"] = controlProcess.Updated
    return leafs
}

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetBundleName() string { return "cisco_ios_xe" }

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetYangName() string { return "control-process" }

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) SetParent(parent types.Entity) { controlProcess.parent = parent }

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetParent() types.Entity { return controlProcess.parent }

func (controlProcess *CiscoPlatformSoftware_ControlProcesses_ControlProcess) GetParentYangName() string { return "control-processes" }

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats
// Load average statictics
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Load average status. The type is string.
    LoadAverageStatus interface{}
}

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetFilter() yfilter.YFilter { return loadAverageStats.YFilter }

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) SetFilter(yf yfilter.YFilter) { loadAverageStats.YFilter = yf }

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetGoName(yname string) string {
    if yname == "load-average-status" { return "LoadAverageStatus" }
    return ""
}

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetSegmentPath() string {
    return "load-average-stats"
}

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["load-average-status"] = loadAverageStats.LoadAverageStatus
    return leafs
}

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetBundleName() string { return "cisco_ios_xe" }

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetYangName() string { return "load-average-stats" }

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) SetParent(parent types.Entity) { loadAverageStats.parent = parent }

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetParent() types.Entity { return loadAverageStats.parent }

func (loadAverageStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAverageStats) GetParentYangName() string { return "control-process" }

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes
// Load average statistics calculated over a period of time
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of Load averages based on a time frame. The type is slice of
    // CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute.
    LoadAvgMinute []CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute
}

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetFilter() yfilter.YFilter { return loadAvgMinutes.YFilter }

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) SetFilter(yf yfilter.YFilter) { loadAvgMinutes.YFilter = yf }

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetGoName(yname string) string {
    if yname == "load-avg-minute" { return "LoadAvgMinute" }
    return ""
}

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetSegmentPath() string {
    return "load-avg-minutes"
}

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-avg-minute" {
        for _, c := range loadAvgMinutes.LoadAvgMinute {
            if loadAvgMinutes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute{}
        loadAvgMinutes.LoadAvgMinute = append(loadAvgMinutes.LoadAvgMinute, child)
        return &loadAvgMinutes.LoadAvgMinute[len(loadAvgMinutes.LoadAvgMinute)-1]
    }
    return nil
}

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range loadAvgMinutes.LoadAvgMinute {
        children[loadAvgMinutes.LoadAvgMinute[i].GetSegmentPath()] = &loadAvgMinutes.LoadAvgMinute[i]
    }
    return children
}

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetBundleName() string { return "cisco_ios_xe" }

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetYangName() string { return "load-avg-minutes" }

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) SetParent(parent types.Entity) { loadAvgMinutes.parent = parent }

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetParent() types.Entity { return loadAvgMinutes.parent }

func (loadAvgMinutes *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes) GetParentYangName() string { return "control-process" }

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute
// List of Load averages based on a time frame
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute struct {
    parent types.Entity
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

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetFilter() yfilter.YFilter { return loadAvgMinute.YFilter }

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) SetFilter(yf yfilter.YFilter) { loadAvgMinute.YFilter = yf }

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "average" { return "Average" }
    if yname == "status" { return "Status" }
    return ""
}

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetSegmentPath() string {
    return "load-avg-minute" + "[number='" + fmt.Sprintf("%v", loadAvgMinute.Number) + "']"
}

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "status" {
        return &loadAvgMinute.Status
    }
    return nil
}

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["status"] = &loadAvgMinute.Status
    return children
}

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = loadAvgMinute.Number
    leafs["average"] = loadAvgMinute.Average
    return leafs
}

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetBundleName() string { return "cisco_ios_xe" }

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetYangName() string { return "load-avg-minute" }

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) SetParent(parent types.Entity) { loadAvgMinute.parent = parent }

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetParent() types.Entity { return loadAvgMinute.parent }

func (loadAvgMinute *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute) GetParentYangName() string { return "load-avg-minutes" }

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status
// Load average statistics minute status
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Load average condition. The type is string.
    Condition interface{}

    // Load average status. The type is string.
    ThresholdStatus interface{}

    // Load average threshold. The type is string with range:
    // -92233720368547758.08..92233720368547758.07.
    ThresholdValue interface{}
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetFilter() yfilter.YFilter { return status.YFilter }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) SetFilter(yf yfilter.YFilter) { status.YFilter = yf }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetGoName(yname string) string {
    if yname == "condition" { return "Condition" }
    if yname == "threshold-status" { return "ThresholdStatus" }
    if yname == "threshold-value" { return "ThresholdValue" }
    return ""
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetSegmentPath() string {
    return "status"
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["condition"] = status.Condition
    leafs["threshold-status"] = status.ThresholdStatus
    leafs["threshold-value"] = status.ThresholdValue
    return leafs
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetBundleName() string { return "cisco_ios_xe" }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetYangName() string { return "status" }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) SetParent(parent types.Entity) { status.parent = parent }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetParent() types.Entity { return status.parent }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_LoadAvgMinutes_LoadAvgMinute_Status) GetParentYangName() string { return "load-avg-minute" }

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats
// Memory statistics
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats struct {
    parent types.Entity
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

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetFilter() yfilter.YFilter { return memoryStats.YFilter }

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) SetFilter(yf yfilter.YFilter) { memoryStats.YFilter = yf }

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetGoName(yname string) string {
    if yname == "memory-status" { return "MemoryStatus" }
    if yname == "total" { return "Total" }
    if yname == "used-number" { return "UsedNumber" }
    if yname == "used-percent" { return "UsedPercent" }
    if yname == "free-number" { return "FreeNumber" }
    if yname == "free-percent" { return "FreePercent" }
    if yname == "available-number" { return "AvailableNumber" }
    if yname == "available-percent" { return "AvailablePercent" }
    if yname == "committed-number" { return "CommittedNumber" }
    if yname == "committed-percent" { return "CommittedPercent" }
    if yname == "status" { return "Status" }
    return ""
}

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetSegmentPath() string {
    return "memory-stats"
}

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "status" {
        return &memoryStats.Status
    }
    return nil
}

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["status"] = &memoryStats.Status
    return children
}

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["memory-status"] = memoryStats.MemoryStatus
    leafs["total"] = memoryStats.Total
    leafs["used-number"] = memoryStats.UsedNumber
    leafs["used-percent"] = memoryStats.UsedPercent
    leafs["free-number"] = memoryStats.FreeNumber
    leafs["free-percent"] = memoryStats.FreePercent
    leafs["available-number"] = memoryStats.AvailableNumber
    leafs["available-percent"] = memoryStats.AvailablePercent
    leafs["committed-number"] = memoryStats.CommittedNumber
    leafs["committed-percent"] = memoryStats.CommittedPercent
    return leafs
}

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetBundleName() string { return "cisco_ios_xe" }

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetYangName() string { return "memory-stats" }

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) SetParent(parent types.Entity) { memoryStats.parent = parent }

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetParent() types.Entity { return memoryStats.parent }

func (memoryStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats) GetParentYangName() string { return "control-process" }

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status
// Memory status
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory warning threshold value percent. The type is interface{} with range:
    // 0..4294967295.
    WarningThresholdPercent interface{}

    // Memory critical threshold value percent. The type is interface{} with
    // range: 0..4294967295.
    CriticalThresholdPercent interface{}
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetFilter() yfilter.YFilter { return status.YFilter }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) SetFilter(yf yfilter.YFilter) { status.YFilter = yf }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetGoName(yname string) string {
    if yname == "warning-threshold-percent" { return "WarningThresholdPercent" }
    if yname == "critical-threshold-percent" { return "CriticalThresholdPercent" }
    return ""
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetSegmentPath() string {
    return "status"
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["warning-threshold-percent"] = status.WarningThresholdPercent
    leafs["critical-threshold-percent"] = status.CriticalThresholdPercent
    return leafs
}

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetBundleName() string { return "cisco_ios_xe" }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetYangName() string { return "status" }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) SetParent(parent types.Entity) { status.parent = parent }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetParent() types.Entity { return status.parent }

func (status *CiscoPlatformSoftware_ControlProcesses_ControlProcess_MemoryStats_Status) GetParentYangName() string { return "memory-stats" }

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats
// Processor core statistics
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of processor cores. The type is slice of
    // CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat.
    PerCoreStat []CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat
}

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetFilter() yfilter.YFilter { return perCoreStats.YFilter }

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) SetFilter(yf yfilter.YFilter) { perCoreStats.YFilter = yf }

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetGoName(yname string) string {
    if yname == "per-core-stat" { return "PerCoreStat" }
    return ""
}

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetSegmentPath() string {
    return "per-core-stats"
}

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "per-core-stat" {
        for _, c := range perCoreStats.PerCoreStat {
            if perCoreStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat{}
        perCoreStats.PerCoreStat = append(perCoreStats.PerCoreStat, child)
        return &perCoreStats.PerCoreStat[len(perCoreStats.PerCoreStat)-1]
    }
    return nil
}

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range perCoreStats.PerCoreStat {
        children[perCoreStats.PerCoreStat[i].GetSegmentPath()] = &perCoreStats.PerCoreStat[i]
    }
    return children
}

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetBundleName() string { return "cisco_ios_xe" }

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetYangName() string { return "per-core-stats" }

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) SetParent(parent types.Entity) { perCoreStats.parent = parent }

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetParent() types.Entity { return perCoreStats.parent }

func (perCoreStats *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats) GetParentYangName() string { return "control-process" }

// CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat
// List of processor cores
type CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat struct {
    parent types.Entity
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

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetFilter() yfilter.YFilter { return perCoreStat.YFilter }

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) SetFilter(yf yfilter.YFilter) { perCoreStat.YFilter = yf }

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "user" { return "User" }
    if yname == "system" { return "System" }
    if yname == "nice" { return "Nice" }
    if yname == "idle" { return "Idle" }
    if yname == "irq" { return "Irq" }
    if yname == "sirq" { return "Sirq" }
    if yname == "io-wait" { return "IoWait" }
    return ""
}

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetSegmentPath() string {
    return "per-core-stat" + "[name='" + fmt.Sprintf("%v", perCoreStat.Name) + "']"
}

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = perCoreStat.Name
    leafs["user"] = perCoreStat.User
    leafs["system"] = perCoreStat.System
    leafs["nice"] = perCoreStat.Nice
    leafs["idle"] = perCoreStat.Idle
    leafs["irq"] = perCoreStat.Irq
    leafs["sirq"] = perCoreStat.Sirq
    leafs["io-wait"] = perCoreStat.IoWait
    return leafs
}

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetBundleName() string { return "cisco_ios_xe" }

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetYangName() string { return "per-core-stat" }

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) SetParent(parent types.Entity) { perCoreStat.parent = parent }

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetParent() types.Entity { return perCoreStat.parent }

func (perCoreStat *CiscoPlatformSoftware_ControlProcesses_ControlProcess_PerCoreStats_PerCoreStat) GetParentYangName() string { return "per-core-stats" }

