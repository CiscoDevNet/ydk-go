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
    parent types.Entity
    YFilter yfilter.YFilter

    // Data nodes for Total CPU Utilization Statistics.
    CpuUtilization CpuUsage_CpuUtilization
}

func (cpuUsage *CpuUsage) GetFilter() yfilter.YFilter { return cpuUsage.YFilter }

func (cpuUsage *CpuUsage) SetFilter(yf yfilter.YFilter) { cpuUsage.YFilter = yf }

func (cpuUsage *CpuUsage) GetGoName(yname string) string {
    if yname == "cpu-utilization" { return "CpuUtilization" }
    return ""
}

func (cpuUsage *CpuUsage) GetSegmentPath() string {
    return "Cisco-IOS-XE-process-cpu-oper:cpu-usage"
}

func (cpuUsage *CpuUsage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpu-utilization" {
        return &cpuUsage.CpuUtilization
    }
    return nil
}

func (cpuUsage *CpuUsage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpu-utilization"] = &cpuUsage.CpuUtilization
    return children
}

func (cpuUsage *CpuUsage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpuUsage *CpuUsage) GetBundleName() string { return "cisco_ios_xe" }

func (cpuUsage *CpuUsage) GetYangName() string { return "cpu-usage" }

func (cpuUsage *CpuUsage) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpuUsage *CpuUsage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpuUsage *CpuUsage) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpuUsage *CpuUsage) SetParent(parent types.Entity) { cpuUsage.parent = parent }

func (cpuUsage *CpuUsage) GetParent() types.Entity { return cpuUsage.parent }

func (cpuUsage *CpuUsage) GetParentYangName() string { return "Cisco-IOS-XE-process-cpu-oper" }

// CpuUsage_CpuUtilization
// Data nodes for Total CPU Utilization Statistics.
type CpuUsage_CpuUtilization struct {
    parent types.Entity
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

func (cpuUtilization *CpuUsage_CpuUtilization) GetFilter() yfilter.YFilter { return cpuUtilization.YFilter }

func (cpuUtilization *CpuUsage_CpuUtilization) SetFilter(yf yfilter.YFilter) { cpuUtilization.YFilter = yf }

func (cpuUtilization *CpuUsage_CpuUtilization) GetGoName(yname string) string {
    if yname == "five-seconds" { return "FiveSeconds" }
    if yname == "five-seconds-intr" { return "FiveSecondsIntr" }
    if yname == "one-minute" { return "OneMinute" }
    if yname == "five-minutes" { return "FiveMinutes" }
    if yname == "cpu-usage-processes" { return "CpuUsageProcesses" }
    return ""
}

func (cpuUtilization *CpuUsage_CpuUtilization) GetSegmentPath() string {
    return "cpu-utilization"
}

func (cpuUtilization *CpuUsage_CpuUtilization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpu-usage-processes" {
        return &cpuUtilization.CpuUsageProcesses
    }
    return nil
}

func (cpuUtilization *CpuUsage_CpuUtilization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpu-usage-processes"] = &cpuUtilization.CpuUsageProcesses
    return children
}

func (cpuUtilization *CpuUsage_CpuUtilization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["five-seconds"] = cpuUtilization.FiveSeconds
    leafs["five-seconds-intr"] = cpuUtilization.FiveSecondsIntr
    leafs["one-minute"] = cpuUtilization.OneMinute
    leafs["five-minutes"] = cpuUtilization.FiveMinutes
    return leafs
}

func (cpuUtilization *CpuUsage_CpuUtilization) GetBundleName() string { return "cisco_ios_xe" }

func (cpuUtilization *CpuUsage_CpuUtilization) GetYangName() string { return "cpu-utilization" }

func (cpuUtilization *CpuUsage_CpuUtilization) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpuUtilization *CpuUsage_CpuUtilization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpuUtilization *CpuUsage_CpuUtilization) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpuUtilization *CpuUsage_CpuUtilization) SetParent(parent types.Entity) { cpuUtilization.parent = parent }

func (cpuUtilization *CpuUsage_CpuUtilization) GetParent() types.Entity { return cpuUtilization.parent }

func (cpuUtilization *CpuUsage_CpuUtilization) GetParentYangName() string { return "cpu-usage" }

// CpuUsage_CpuUtilization_CpuUsageProcesses
// Data nodes for System wide Process CPU usage Statistics.
type CpuUsage_CpuUtilization_CpuUsageProcesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of software processes on the device. The type is slice of
    // CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess.
    CpuUsageProcess []CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess
}

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetFilter() yfilter.YFilter { return cpuUsageProcesses.YFilter }

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) SetFilter(yf yfilter.YFilter) { cpuUsageProcesses.YFilter = yf }

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetGoName(yname string) string {
    if yname == "cpu-usage-process" { return "CpuUsageProcess" }
    return ""
}

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetSegmentPath() string {
    return "cpu-usage-processes"
}

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpu-usage-process" {
        for _, c := range cpuUsageProcesses.CpuUsageProcess {
            if cpuUsageProcesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess{}
        cpuUsageProcesses.CpuUsageProcess = append(cpuUsageProcesses.CpuUsageProcess, child)
        return &cpuUsageProcesses.CpuUsageProcess[len(cpuUsageProcesses.CpuUsageProcess)-1]
    }
    return nil
}

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpuUsageProcesses.CpuUsageProcess {
        children[cpuUsageProcesses.CpuUsageProcess[i].GetSegmentPath()] = &cpuUsageProcesses.CpuUsageProcess[i]
    }
    return children
}

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetBundleName() string { return "cisco_ios_xe" }

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetYangName() string { return "cpu-usage-processes" }

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) SetParent(parent types.Entity) { cpuUsageProcesses.parent = parent }

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetParent() types.Entity { return cpuUsageProcesses.parent }

func (cpuUsageProcesses *CpuUsage_CpuUtilization_CpuUsageProcesses) GetParentYangName() string { return "cpu-utilization" }

// CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess
// The list of software processes on the device.
type CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess struct {
    parent types.Entity
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

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetFilter() yfilter.YFilter { return cpuUsageProcess.YFilter }

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) SetFilter(yf yfilter.YFilter) { cpuUsageProcess.YFilter = yf }

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetGoName(yname string) string {
    if yname == "pid" { return "Pid" }
    if yname == "name" { return "Name" }
    if yname == "tty" { return "Tty" }
    if yname == "total-run-time" { return "TotalRunTime" }
    if yname == "invocation-count" { return "InvocationCount" }
    if yname == "avg-run-time" { return "AvgRunTime" }
    if yname == "five-seconds" { return "FiveSeconds" }
    if yname == "one-minute" { return "OneMinute" }
    if yname == "five-minutes" { return "FiveMinutes" }
    return ""
}

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetSegmentPath() string {
    return "cpu-usage-process" + "[pid='" + fmt.Sprintf("%v", cpuUsageProcess.Pid) + "']" + "[name='" + fmt.Sprintf("%v", cpuUsageProcess.Name) + "']"
}

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pid"] = cpuUsageProcess.Pid
    leafs["name"] = cpuUsageProcess.Name
    leafs["tty"] = cpuUsageProcess.Tty
    leafs["total-run-time"] = cpuUsageProcess.TotalRunTime
    leafs["invocation-count"] = cpuUsageProcess.InvocationCount
    leafs["avg-run-time"] = cpuUsageProcess.AvgRunTime
    leafs["five-seconds"] = cpuUsageProcess.FiveSeconds
    leafs["one-minute"] = cpuUsageProcess.OneMinute
    leafs["five-minutes"] = cpuUsageProcess.FiveMinutes
    return leafs
}

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetBundleName() string { return "cisco_ios_xe" }

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetYangName() string { return "cpu-usage-process" }

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) SetParent(parent types.Entity) { cpuUsageProcess.parent = parent }

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetParent() types.Entity { return cpuUsageProcess.parent }

func (cpuUsageProcess *CpuUsage_CpuUtilization_CpuUsageProcesses_CpuUsageProcess) GetParentYangName() string { return "cpu-usage-processes" }

