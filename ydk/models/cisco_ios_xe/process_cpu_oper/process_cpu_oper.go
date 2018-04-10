// This module contains a collection of YANG definitions for
// monitoring CPU usage of processes in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package process_cpu_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package process_cpu_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-process-cpu-oper cpu-usage}", reflect.TypeOf(CpuUsage{}))
    ydk.RegisterEntity("Cisco-IOS-XE-process-cpu-oper:cpu-usage", reflect.TypeOf(CpuUsage{}))
}

// CpuUsage
// CPU Utilization data
type CpuUsage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data nodes for Total CPU Utilization Statistics.
    CpuUtilization CpuUsage_CpuUtilization
}

func (cpuUsage *CpuUsage) GetEntityData() *types.CommonEntityData {
    cpuUsage.EntityData.YFilter = cpuUsage.YFilter
    cpuUsage.EntityData.YangName = "cpu-usage"
    cpuUsage.EntityData.BundleName = "cisco_ios_xe"
    cpuUsage.EntityData.ParentYangName = "Cisco-IOS-XE-process-cpu-oper"
    cpuUsage.EntityData.SegmentPath = "Cisco-IOS-XE-process-cpu-oper:cpu-usage"
    cpuUsage.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpuUsage.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpuUsage.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpuUsage.EntityData.Children = make(map[string]types.YChild)
    cpuUsage.EntityData.Children["cpu-utilization"] = types.YChild{"CpuUtilization", &cpuUsage.CpuUtilization}
    cpuUsage.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpuUsage.EntityData)
}

// CpuUsage_CpuUtilization
// Data nodes for Total CPU Utilization Statistics.
type CpuUsage_CpuUtilization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Busy percentage in last 5-seconds. The type is interface{} with range:
    // 0..255. Units are percent.
    FiveSeconds interface{}

    // Interrupt busy percentage in last 5-seconds. The type is interface{} with
    // range: 0..255. Units are percent.
    FiveSecondsIntr interface{}

    // Busy percentage in last one minute. The type is interface{} with range:
    // 0..255. Units are percent.
    OneMinute interface{}

    // Busy percentage in last five minutes. The type is interface{} with range:
    // 0..255. Units are percent.
    FiveMinutes interface{}

    // Data nodes for System wide Process CPU usage Statistics.
    CpuUsageProcesses CpuUsage_CpuUtilization_CpuUsageProcesses
}

func (cpuUtilization *CpuUsage_CpuUtilization) GetEntityData() *types.CommonEntityData {
    cpuUtilization.EntityData.YFilter = cpuUtilization.YFilter
    cpuUtilization.EntityData.YangName = "cpu-utilization"
    cpuUtilization.EntityData.BundleName = "cisco_ios_xe"
    cpuUtilization.EntityData.ParentYangName = "cpu-usage"
    cpuUtilization.EntityData.SegmentPath = "cpu-utilization"
    cpuUtilization.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpuUtilization.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpuUtilization.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpuUtilization.EntityData.Children = make(map[string]types.YChild)
    cpuUtilization.EntityData.Children["cpu-usage-processes"] = types.YChild{"CpuUsageProcesses", &cpuUtilization.CpuUsageProcesses}
    cpuUtilization.EntityData.Leafs = make(map[string]types.YLeaf)
    cpuUtilization.EntityData.Leafs["five-seconds"] = types.YLeaf{"FiveSeconds", cpuUtilization.FiveSeconds}
    cpuUtilization.EntityData.Leafs["five-seconds-intr"] = types.YLeaf{"FiveSecondsIntr", cpuUtilization.FiveSecondsIntr}
    cpuUtilization.EntityData.Leafs["one-minute"] = types.YLeaf{"OneMinute", cpuUtilization.OneMinute}
    cpuUtilization.EntityData.Leafs["five-minutes"] = types.YLeaf{"FiveMinutes", cpuUtilization.FiveMinutes}
    return &(cpuUtilization.EntityData)
}

// CpuUsage_CpuUtilization_CpuUsageProcesses
// Data nodes for System wide Process CPU usage Statistics.
type CpuUsage_CpuUtilization_CpuUsageProcesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of software processes on the device. The type is slice of
    // CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess.
    CpuUsageProcess []CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess
}

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetEntityData() *types.CommonEntityData {
    cpuUsageProcesses.EntityData.YFilter = cpuUsageProcesses.YFilter
    cpuUsageProcesses.EntityData.YangName = "cpu-usage-processes"
    cpuUsageProcesses.EntityData.BundleName = "cisco_ios_xe"
    cpuUsageProcesses.EntityData.ParentYangName = "cpu-utilization"
    cpuUsageProcesses.EntityData.SegmentPath = "cpu-usage-processes"
    cpuUsageProcesses.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpuUsageProcesses.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpuUsageProcesses.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpuUsageProcesses.EntityData.Children = make(map[string]types.YChild)
    cpuUsageProcesses.EntityData.Children["cpu-usage-process"] = types.YChild{"CpuUsageProcess", nil}
    for i := range cpuUsageProcesses.CpuUsageProcess {
        cpuUsageProcesses.EntityData.Children[types.GetSegmentPath(&cpuUsageProcesses.CpuUsageProcess[i])] = types.YChild{"CpuUsageProcess", &cpuUsageProcesses.CpuUsageProcess[i]}
    }
    cpuUsageProcesses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpuUsageProcesses.EntityData)
}

// CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess
// The list of software processes on the device.
type CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Process-ID of the process. The type is interface{}
    // with range: 0..4294967295.
    Pid interface{}

    // This attribute is a key. The name of the process. The type is string.
    Name interface{}

    // TTY bound to by the process. The type is interface{} with range: 0..65535.
    Tty interface{}

    // Total Run-time of this process (mSec). The type is interface{} with range:
    // 0..18446744073709551615. Units are milli-seconds.
    TotalRunTime interface{}

    // Total number of invocations. The type is interface{} with range:
    // 0..4294967295.
    InvocationCount interface{}

    // Average Run-time of this process (uSec). The type is interface{} with
    // range: 0..18446744073709551615. Units are micro-seconds.
    AvgRunTime interface{}

    // Busy percentage in last 5-seconds. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are percent.
    FiveSeconds interface{}

    // Busy percentage in last one minute. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are percent.
    OneMinute interface{}

    // Busy percentage in last five minutes. The type is string with range:
    // -92233720368547758.08..92233720368547758.07. Units are percent.
    FiveMinutes interface{}
}

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetEntityData() *types.CommonEntityData {
    cpuUsageProcess.EntityData.YFilter = cpuUsageProcess.YFilter
    cpuUsageProcess.EntityData.YangName = "cpu-usage-process"
    cpuUsageProcess.EntityData.BundleName = "cisco_ios_xe"
    cpuUsageProcess.EntityData.ParentYangName = "cpu-usage-processes"
    cpuUsageProcess.EntityData.SegmentPath = "cpu-usage-process" + "[pid='" + fmt.Sprintf("%v", cpuUsageProcess.Pid) + "']" + "[name='" + fmt.Sprintf("%v", cpuUsageProcess.Name) + "']"
    cpuUsageProcess.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpuUsageProcess.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpuUsageProcess.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpuUsageProcess.EntityData.Children = make(map[string]types.YChild)
    cpuUsageProcess.EntityData.Leafs = make(map[string]types.YLeaf)
    cpuUsageProcess.EntityData.Leafs["pid"] = types.YLeaf{"Pid", cpuUsageProcess.Pid}
    cpuUsageProcess.EntityData.Leafs["name"] = types.YLeaf{"Name", cpuUsageProcess.Name}
    cpuUsageProcess.EntityData.Leafs["tty"] = types.YLeaf{"Tty", cpuUsageProcess.Tty}
    cpuUsageProcess.EntityData.Leafs["total-run-time"] = types.YLeaf{"TotalRunTime", cpuUsageProcess.TotalRunTime}
    cpuUsageProcess.EntityData.Leafs["invocation-count"] = types.YLeaf{"InvocationCount", cpuUsageProcess.InvocationCount}
    cpuUsageProcess.EntityData.Leafs["avg-run-time"] = types.YLeaf{"AvgRunTime", cpuUsageProcess.AvgRunTime}
    cpuUsageProcess.EntityData.Leafs["five-seconds"] = types.YLeaf{"FiveSeconds", cpuUsageProcess.FiveSeconds}
    cpuUsageProcess.EntityData.Leafs["one-minute"] = types.YLeaf{"OneMinute", cpuUsageProcess.OneMinute}
    cpuUsageProcess.EntityData.Leafs["five-minutes"] = types.YLeaf{"FiveMinutes", cpuUsageProcess.FiveMinutes}
    return &(cpuUsageProcess.EntityData)
}

