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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Processes CPU utilization information. The type is slice of
    // SystemMonitoring_CpuUtilization.
    CpuUtilization []SystemMonitoring_CpuUtilization
}

func (systemMonitoring *SystemMonitoring) GetEntityData() *types.CommonEntityData {
    systemMonitoring.EntityData.YFilter = systemMonitoring.YFilter
    systemMonitoring.EntityData.YangName = "system-monitoring"
    systemMonitoring.EntityData.BundleName = "cisco_ios_xr"
    systemMonitoring.EntityData.ParentYangName = "Cisco-IOS-XR-wdsysmon-fd-oper"
    systemMonitoring.EntityData.SegmentPath = "Cisco-IOS-XR-wdsysmon-fd-oper:system-monitoring"
    systemMonitoring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemMonitoring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemMonitoring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemMonitoring.EntityData.Children = make(map[string]types.YChild)
    systemMonitoring.EntityData.Children["cpu-utilization"] = types.YChild{"CpuUtilization", nil}
    for i := range systemMonitoring.CpuUtilization {
        systemMonitoring.EntityData.Children[types.GetSegmentPath(&systemMonitoring.CpuUtilization[i])] = types.YChild{"CpuUtilization", &systemMonitoring.CpuUtilization[i]}
    }
    systemMonitoring.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(systemMonitoring.EntityData)
}

// SystemMonitoring_CpuUtilization
// Processes CPU utilization information
type SystemMonitoring_CpuUtilization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (cpuUtilization *SystemMonitoring_CpuUtilization) GetEntityData() *types.CommonEntityData {
    cpuUtilization.EntityData.YFilter = cpuUtilization.YFilter
    cpuUtilization.EntityData.YangName = "cpu-utilization"
    cpuUtilization.EntityData.BundleName = "cisco_ios_xr"
    cpuUtilization.EntityData.ParentYangName = "system-monitoring"
    cpuUtilization.EntityData.SegmentPath = "cpu-utilization" + "[node-name='" + fmt.Sprintf("%v", cpuUtilization.NodeName) + "']"
    cpuUtilization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpuUtilization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpuUtilization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpuUtilization.EntityData.Children = make(map[string]types.YChild)
    cpuUtilization.EntityData.Children["process-cpu"] = types.YChild{"ProcessCpu", nil}
    for i := range cpuUtilization.ProcessCpu {
        cpuUtilization.EntityData.Children[types.GetSegmentPath(&cpuUtilization.ProcessCpu[i])] = types.YChild{"ProcessCpu", &cpuUtilization.ProcessCpu[i]}
    }
    cpuUtilization.EntityData.Leafs = make(map[string]types.YLeaf)
    cpuUtilization.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", cpuUtilization.NodeName}
    cpuUtilization.EntityData.Leafs["total-cpu-one-minute"] = types.YLeaf{"TotalCpuOneMinute", cpuUtilization.TotalCpuOneMinute}
    cpuUtilization.EntityData.Leafs["total-cpu-five-minute"] = types.YLeaf{"TotalCpuFiveMinute", cpuUtilization.TotalCpuFiveMinute}
    cpuUtilization.EntityData.Leafs["total-cpu-fifteen-minute"] = types.YLeaf{"TotalCpuFifteenMinute", cpuUtilization.TotalCpuFifteenMinute}
    return &(cpuUtilization.EntityData)
}

// SystemMonitoring_CpuUtilization_ProcessCpu
// Per process CPU utilization
type SystemMonitoring_CpuUtilization_ProcessCpu struct {
    EntityData types.CommonEntityData
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

func (processCpu *SystemMonitoring_CpuUtilization_ProcessCpu) GetEntityData() *types.CommonEntityData {
    processCpu.EntityData.YFilter = processCpu.YFilter
    processCpu.EntityData.YangName = "process-cpu"
    processCpu.EntityData.BundleName = "cisco_ios_xr"
    processCpu.EntityData.ParentYangName = "cpu-utilization"
    processCpu.EntityData.SegmentPath = "process-cpu"
    processCpu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processCpu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processCpu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processCpu.EntityData.Children = make(map[string]types.YChild)
    processCpu.EntityData.Leafs = make(map[string]types.YLeaf)
    processCpu.EntityData.Leafs["process-name"] = types.YLeaf{"ProcessName", processCpu.ProcessName}
    processCpu.EntityData.Leafs["process-id"] = types.YLeaf{"ProcessId", processCpu.ProcessId}
    processCpu.EntityData.Leafs["process-cpu-one-minute"] = types.YLeaf{"ProcessCpuOneMinute", processCpu.ProcessCpuOneMinute}
    processCpu.EntityData.Leafs["process-cpu-five-minute"] = types.YLeaf{"ProcessCpuFiveMinute", processCpu.ProcessCpuFiveMinute}
    processCpu.EntityData.Leafs["process-cpu-fifteen-minute"] = types.YLeaf{"ProcessCpuFifteenMinute", processCpu.ProcessCpuFifteenMinute}
    return &(processCpu.EntityData)
}

