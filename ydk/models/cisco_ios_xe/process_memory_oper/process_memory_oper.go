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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of software processes on the device. The type is slice of
    // MemoryUsageProcesses_MemoryUsageProcess.
    MemoryUsageProcess []MemoryUsageProcesses_MemoryUsageProcess
}

func (memoryUsageProcesses *MemoryUsageProcesses) GetEntityData() *types.CommonEntityData {
    memoryUsageProcesses.EntityData.YFilter = memoryUsageProcesses.YFilter
    memoryUsageProcesses.EntityData.YangName = "memory-usage-processes"
    memoryUsageProcesses.EntityData.BundleName = "cisco_ios_xe"
    memoryUsageProcesses.EntityData.ParentYangName = "Cisco-IOS-XE-process-memory-oper"
    memoryUsageProcesses.EntityData.SegmentPath = "Cisco-IOS-XE-process-memory-oper:memory-usage-processes"
    memoryUsageProcesses.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    memoryUsageProcesses.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    memoryUsageProcesses.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    memoryUsageProcesses.EntityData.Children = make(map[string]types.YChild)
    memoryUsageProcesses.EntityData.Children["memory-usage-process"] = types.YChild{"MemoryUsageProcess", nil}
    for i := range memoryUsageProcesses.MemoryUsageProcess {
        memoryUsageProcesses.EntityData.Children[types.GetSegmentPath(&memoryUsageProcesses.MemoryUsageProcess[i])] = types.YChild{"MemoryUsageProcess", &memoryUsageProcesses.MemoryUsageProcess[i]}
    }
    memoryUsageProcesses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memoryUsageProcesses.EntityData)
}

// MemoryUsageProcesses_MemoryUsageProcess
// The list of software processes on the device.
type MemoryUsageProcesses_MemoryUsageProcess struct {
    EntityData types.CommonEntityData
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

func (memoryUsageProcess *MemoryUsageProcesses_MemoryUsageProcess) GetEntityData() *types.CommonEntityData {
    memoryUsageProcess.EntityData.YFilter = memoryUsageProcess.YFilter
    memoryUsageProcess.EntityData.YangName = "memory-usage-process"
    memoryUsageProcess.EntityData.BundleName = "cisco_ios_xe"
    memoryUsageProcess.EntityData.ParentYangName = "memory-usage-processes"
    memoryUsageProcess.EntityData.SegmentPath = "memory-usage-process" + "[pid='" + fmt.Sprintf("%v", memoryUsageProcess.Pid) + "']" + "[name='" + fmt.Sprintf("%v", memoryUsageProcess.Name) + "']"
    memoryUsageProcess.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    memoryUsageProcess.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    memoryUsageProcess.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    memoryUsageProcess.EntityData.Children = make(map[string]types.YChild)
    memoryUsageProcess.EntityData.Leafs = make(map[string]types.YLeaf)
    memoryUsageProcess.EntityData.Leafs["pid"] = types.YLeaf{"Pid", memoryUsageProcess.Pid}
    memoryUsageProcess.EntityData.Leafs["name"] = types.YLeaf{"Name", memoryUsageProcess.Name}
    memoryUsageProcess.EntityData.Leafs["tty"] = types.YLeaf{"Tty", memoryUsageProcess.Tty}
    memoryUsageProcess.EntityData.Leafs["allocated-memory"] = types.YLeaf{"AllocatedMemory", memoryUsageProcess.AllocatedMemory}
    memoryUsageProcess.EntityData.Leafs["freed-memory"] = types.YLeaf{"FreedMemory", memoryUsageProcess.FreedMemory}
    memoryUsageProcess.EntityData.Leafs["holding-memory"] = types.YLeaf{"HoldingMemory", memoryUsageProcess.HoldingMemory}
    memoryUsageProcess.EntityData.Leafs["get-buffers"] = types.YLeaf{"GetBuffers", memoryUsageProcess.GetBuffers}
    memoryUsageProcess.EntityData.Leafs["ret-buffers"] = types.YLeaf{"RetBuffers", memoryUsageProcess.RetBuffers}
    return &(memoryUsageProcess.EntityData)
}

