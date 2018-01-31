// This module contains a collection of YANG definitions for
// monitoring memory usage of processes in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package process_memory_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package process_memory_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-process-memory-oper memory-usage-processes}", reflect.TypeOf(MemoryUsageProcesses{}))
    ydk.RegisterEntity("Cisco-IOS-XE-process-memory-oper:memory-usage-processes", reflect.TypeOf(MemoryUsageProcesses{}))
}

// MemoryUsageProcesses
// Data nodes for System wide Process Memory Statistics.
type MemoryUsageProcesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of software processes on the device. The type is slice of
    // MemoryUsageProcesses_MemoryUsageProcess.
    MemoryUsageProcess []MemoryUsageProcesses_MemoryUsageProcess
}

func (memoryUsageProcesses *MemoryUsageProcesses) GetFilter() yfilter.YFilter { return memoryUsageProcesses.YFilter }

func (memoryUsageProcesses *MemoryUsageProcesses) SetFilter(yf yfilter.YFilter) { memoryUsageProcesses.YFilter = yf }

func (memoryUsageProcesses *MemoryUsageProcesses) GetGoName(yname string) string {
    if yname == "memory-usage-process" { return "MemoryUsageProcess" }
    return ""
}

func (memoryUsageProcesses *MemoryUsageProcesses) GetSegmentPath() string {
    return "Cisco-IOS-XE-process-memory-oper:memory-usage-processes"
}

func (memoryUsageProcesses *MemoryUsageProcesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "memory-usage-process" {
        for _, c := range memoryUsageProcesses.MemoryUsageProcess {
            if memoryUsageProcesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MemoryUsageProcesses_MemoryUsageProcess{}
        memoryUsageProcesses.MemoryUsageProcess = append(memoryUsageProcesses.MemoryUsageProcess, child)
        return &memoryUsageProcesses.MemoryUsageProcess[len(memoryUsageProcesses.MemoryUsageProcess)-1]
    }
    return nil
}

func (memoryUsageProcesses *MemoryUsageProcesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range memoryUsageProcesses.MemoryUsageProcess {
        children[memoryUsageProcesses.MemoryUsageProcess[i].GetSegmentPath()] = &memoryUsageProcesses.MemoryUsageProcess[i]
    }
    return children
}

func (memoryUsageProcesses *MemoryUsageProcesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memoryUsageProcesses *MemoryUsageProcesses) GetBundleName() string { return "cisco_ios_xe" }

func (memoryUsageProcesses *MemoryUsageProcesses) GetYangName() string { return "memory-usage-processes" }

func (memoryUsageProcesses *MemoryUsageProcesses) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (memoryUsageProcesses *MemoryUsageProcesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (memoryUsageProcesses *MemoryUsageProcesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (memoryUsageProcesses *MemoryUsageProcesses) SetParent(parent types.Entity) { memoryUsageProcesses.parent = parent }

func (memoryUsageProcesses *MemoryUsageProcesses) GetParent() types.Entity { return memoryUsageProcesses.parent }

func (memoryUsageProcesses *MemoryUsageProcesses) GetParentYangName() string { return "Cisco-IOS-XE-process-memory-oper" }

// MemoryUsageProcesses_MemoryUsageProcess
// The list of software processes on the device.
type MemoryUsageProcesses_MemoryUsageProcess struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Process-ID of the process. The type is interface{}
    // with range: 0..4294967295.
    Pid interface{}

    // This attribute is a key. The name of the process. The type is string.
    Name interface{}

    // TTY bound to by the process. The type is interface{} with range: 0..65535.
    Tty interface{}

    // Total memory allocated to this process (bytes). The type is interface{}
    // with range: 0..18446744073709551615. Units are bytes.
    AllocatedMemory interface{}

    // Total memory freed by this process (bytes). The type is interface{} with
    // range: 0..18446744073709551615. Units are bytes.
    FreedMemory interface{}

    // Total memory currently held by this process (bytes). The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    HoldingMemory interface{}

    // Get Buffers of this process (bytes). The type is interface{} with range:
    // 0..4294967295.
    GetBuffers interface{}

    // Return Buffers of this process (bytes). The type is interface{} with range:
    // 0..4294967295.
    RetBuffers interface{}
}

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetFilter() yfilter.YFilter { return memoryUsageProcess.YFilter }

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) SetFilter(yf yfilter.YFilter) { memoryUsageProcess.YFilter = yf }

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetGoName(yname string) string {
    if yname == "pid" { return "Pid" }
    if yname == "name" { return "Name" }
    if yname == "tty" { return "Tty" }
    if yname == "allocated-memory" { return "AllocatedMemory" }
    if yname == "freed-memory" { return "FreedMemory" }
    if yname == "holding-memory" { return "HoldingMemory" }
    if yname == "get-buffers" { return "GetBuffers" }
    if yname == "ret-buffers" { return "RetBuffers" }
    return ""
}

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetSegmentPath() string {
    return "memory-usage-process" + "[pid='" + fmt.Sprintf("%v", memoryUsageProcess.Pid) + "']" + "[name='" + fmt.Sprintf("%v", memoryUsageProcess.Name) + "']"
}

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pid"] = memoryUsageProcess.Pid
    leafs["name"] = memoryUsageProcess.Name
    leafs["tty"] = memoryUsageProcess.Tty
    leafs["allocated-memory"] = memoryUsageProcess.AllocatedMemory
    leafs["freed-memory"] = memoryUsageProcess.FreedMemory
    leafs["holding-memory"] = memoryUsageProcess.HoldingMemory
    leafs["get-buffers"] = memoryUsageProcess.GetBuffers
    leafs["ret-buffers"] = memoryUsageProcess.RetBuffers
    return leafs
}

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetBundleName() string { return "cisco_ios_xe" }

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetYangName() string { return "memory-usage-process" }

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) SetParent(parent types.Entity) { memoryUsageProcess.parent = parent }

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetParent() types.Entity { return memoryUsageProcess.parent }

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetParentYangName() string { return "memory-usage-processes" }

