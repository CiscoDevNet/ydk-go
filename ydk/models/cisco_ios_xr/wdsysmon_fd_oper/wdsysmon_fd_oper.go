// This module contains a collection of YANG definitions
// for Cisco IOS-XR wdsysmon-fd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   system-monitoring: Processes operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package wdsysmon_fd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package wdsysmon_fd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-wdsysmon-fd-oper system-monitoring}", reflect.TypeOf(SystemMonitoring{}))
    ydk.RegisterEntity("Cisco-IOS-XR-wdsysmon-fd-oper:system-monitoring", reflect.TypeOf(SystemMonitoring{}))
}

// SystemMonitoring
// Processes operational data
type SystemMonitoring struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Processes CPU utilization information. The type is slice of
    // SystemMonitoring_CpuUtilization.
    CpuUtilization []SystemMonitoring_CpuUtilization
}

func (systemMonitoring *SystemMonitoring) GetFilter() yfilter.YFilter { return systemMonitoring.YFilter }

func (systemMonitoring *SystemMonitoring) SetFilter(yf yfilter.YFilter) { systemMonitoring.YFilter = yf }

func (systemMonitoring *SystemMonitoring) GetGoName(yname string) string {
    if yname == "cpu-utilization" { return "CpuUtilization" }
    return ""
}

func (systemMonitoring *SystemMonitoring) GetSegmentPath() string {
    return "Cisco-IOS-XR-wdsysmon-fd-oper:system-monitoring"
}

func (systemMonitoring *SystemMonitoring) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpu-utilization" {
        for _, c := range systemMonitoring.CpuUtilization {
            if systemMonitoring.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SystemMonitoring_CpuUtilization{}
        systemMonitoring.CpuUtilization = append(systemMonitoring.CpuUtilization, child)
        return &systemMonitoring.CpuUtilization[len(systemMonitoring.CpuUtilization)-1]
    }
    return nil
}

func (systemMonitoring *SystemMonitoring) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range systemMonitoring.CpuUtilization {
        children[systemMonitoring.CpuUtilization[i].GetSegmentPath()] = &systemMonitoring.CpuUtilization[i]
    }
    return children
}

func (systemMonitoring *SystemMonitoring) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (systemMonitoring *SystemMonitoring) GetBundleName() string { return "cisco_ios_xr" }

func (systemMonitoring *SystemMonitoring) GetYangName() string { return "system-monitoring" }

func (systemMonitoring *SystemMonitoring) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemMonitoring *SystemMonitoring) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemMonitoring *SystemMonitoring) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemMonitoring *SystemMonitoring) SetParent(parent types.Entity) { systemMonitoring.parent = parent }

func (systemMonitoring *SystemMonitoring) GetParent() types.Entity { return systemMonitoring.parent }

func (systemMonitoring *SystemMonitoring) GetParentYangName() string { return "Cisco-IOS-XR-wdsysmon-fd-oper" }

// SystemMonitoring_CpuUtilization
// Processes CPU utilization information
type SystemMonitoring_CpuUtilization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Total CPU utilization in past 1 minute. The type is interface{} with range:
    // 0..4294967295.
    TotalCpuOneMinute interface{}

    // Total CPU utilization in past 5 minute. The type is interface{} with range:
    // 0..4294967295.
    TotalCpuFiveMinute interface{}

    // Total CPU utilization in past 15 minute. The type is interface{} with
    // range: 0..4294967295.
    TotalCpuFifteenMinute interface{}

    // Per process CPU utilization. The type is slice of
    // SystemMonitoring_CpuUtilization_ProcessCpu.
    ProcessCpu []SystemMonitoring_CpuUtilization_ProcessCpu
}

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetFilter() yfilter.YFilter { return cpuUtilization.YFilter }

func (cpuUtilization *SystemMonitoring_CpuUtilization) SetFilter(yf yfilter.YFilter) { cpuUtilization.YFilter = yf }

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "total-cpu-one-minute" { return "TotalCpuOneMinute" }
    if yname == "total-cpu-five-minute" { return "TotalCpuFiveMinute" }
    if yname == "total-cpu-fifteen-minute" { return "TotalCpuFifteenMinute" }
    if yname == "process-cpu" { return "ProcessCpu" }
    return ""
}

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetSegmentPath() string {
    return "cpu-utilization" + "[node-name='" + fmt.Sprintf("%v", cpuUtilization.NodeName) + "']"
}

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-cpu" {
        for _, c := range cpuUtilization.ProcessCpu {
            if cpuUtilization.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SystemMonitoring_CpuUtilization_ProcessCpu{}
        cpuUtilization.ProcessCpu = append(cpuUtilization.ProcessCpu, child)
        return &cpuUtilization.ProcessCpu[len(cpuUtilization.ProcessCpu)-1]
    }
    return nil
}

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpuUtilization.ProcessCpu {
        children[cpuUtilization.ProcessCpu[i].GetSegmentPath()] = &cpuUtilization.ProcessCpu[i]
    }
    return children
}

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = cpuUtilization.NodeName
    leafs["total-cpu-one-minute"] = cpuUtilization.TotalCpuOneMinute
    leafs["total-cpu-five-minute"] = cpuUtilization.TotalCpuFiveMinute
    leafs["total-cpu-fifteen-minute"] = cpuUtilization.TotalCpuFifteenMinute
    return leafs
}

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetBundleName() string { return "cisco_ios_xr" }

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetYangName() string { return "cpu-utilization" }

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cpuUtilization *SystemMonitoring_CpuUtilization) SetParent(parent types.Entity) { cpuUtilization.parent = parent }

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetParent() types.Entity { return cpuUtilization.parent }

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetParentYangName() string { return "system-monitoring" }

// SystemMonitoring_CpuUtilization_ProcessCpu
// Per process CPU utilization
type SystemMonitoring_CpuUtilization_ProcessCpu struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process name. The type is string.
    ProcessName interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    ProcessId interface{}

    // Process CPU utilization in percent for past 1 minute. The type is
    // interface{} with range: 0..4294967295. Units are percentage.
    ProcessCpuOneMinute interface{}

    // Process CPU utilization in percent for past 5 minute. The type is
    // interface{} with range: 0..4294967295. Units are percentage.
    ProcessCpuFiveMinute interface{}

    // Process CPU utilization in percent for past 15 minute. The type is
    // interface{} with range: 0..4294967295. Units are percentage.
    ProcessCpuFifteenMinute interface{}
}

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetFilter() yfilter.YFilter { return processCpu.YFilter }

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) SetFilter(yf yfilter.YFilter) { processCpu.YFilter = yf }

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetGoName(yname string) string {
    if yname == "process-name" { return "ProcessName" }
    if yname == "process-id" { return "ProcessId" }
    if yname == "process-cpu-one-minute" { return "ProcessCpuOneMinute" }
    if yname == "process-cpu-five-minute" { return "ProcessCpuFiveMinute" }
    if yname == "process-cpu-fifteen-minute" { return "ProcessCpuFifteenMinute" }
    return ""
}

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetSegmentPath() string {
    return "process-cpu"
}

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-name"] = processCpu.ProcessName
    leafs["process-id"] = processCpu.ProcessId
    leafs["process-cpu-one-minute"] = processCpu.ProcessCpuOneMinute
    leafs["process-cpu-five-minute"] = processCpu.ProcessCpuFiveMinute
    leafs["process-cpu-fifteen-minute"] = processCpu.ProcessCpuFifteenMinute
    return leafs
}

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetBundleName() string { return "cisco_ios_xr" }

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetYangName() string { return "process-cpu" }

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) SetParent(parent types.Entity) { processCpu.parent = parent }

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetParent() types.Entity { return processCpu.parent }

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetParentYangName() string { return "cpu-utilization" }

