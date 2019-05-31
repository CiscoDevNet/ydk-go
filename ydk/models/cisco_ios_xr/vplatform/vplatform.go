// This module contains definitions
// for the Calvados model objects.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package vplatform

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package vplatform"))
    ydk.RegisterEntity("{http://cisco.com/panini/calvados/vplatform virtual-platform}", reflect.TypeOf(VirtualPlatform{}))
    ydk.RegisterEntity("vplatform:virtual-platform", reflect.TypeOf(VirtualPlatform{}))
}

// VirtualPlatform
type VirtualPlatform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Vplatform VirtualPlatform_Vplatform

    
    Udi VirtualPlatform_Udi

    
    Cpu VirtualPlatform_Cpu

    
    Processor VirtualPlatform_Processor

    
    Memory VirtualPlatform_Memory

    
    Disk VirtualPlatform_Disk

    
    Action VirtualPlatform_Action
}

func (virtualPlatform *VirtualPlatform) GetEntityData() *types.CommonEntityData {
    virtualPlatform.EntityData.YFilter = virtualPlatform.YFilter
    virtualPlatform.EntityData.YangName = "virtual-platform"
    virtualPlatform.EntityData.BundleName = "cisco_ios_xr"
    virtualPlatform.EntityData.ParentYangName = "vplatform"
    virtualPlatform.EntityData.SegmentPath = "vplatform:virtual-platform"
    virtualPlatform.EntityData.AbsolutePath = virtualPlatform.EntityData.SegmentPath
    virtualPlatform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualPlatform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualPlatform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualPlatform.EntityData.Children = types.NewOrderedMap()
    virtualPlatform.EntityData.Children.Append("vplatform", types.YChild{"Vplatform", &virtualPlatform.Vplatform})
    virtualPlatform.EntityData.Children.Append("udi", types.YChild{"Udi", &virtualPlatform.Udi})
    virtualPlatform.EntityData.Children.Append("cpu", types.YChild{"Cpu", &virtualPlatform.Cpu})
    virtualPlatform.EntityData.Children.Append("processor", types.YChild{"Processor", &virtualPlatform.Processor})
    virtualPlatform.EntityData.Children.Append("memory", types.YChild{"Memory", &virtualPlatform.Memory})
    virtualPlatform.EntityData.Children.Append("disk", types.YChild{"Disk", &virtualPlatform.Disk})
    virtualPlatform.EntityData.Children.Append("Action", types.YChild{"Action", &virtualPlatform.Action})
    virtualPlatform.EntityData.Leafs = types.NewOrderedMap()

    virtualPlatform.EntityData.YListKeys = []string {}

    return &(virtualPlatform.EntityData)
}

// VirtualPlatform_Vplatform
type VirtualPlatform_Vplatform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of VirtualPlatform_Vplatform_Summary.
    Summary []*VirtualPlatform_Vplatform_Summary
}

func (vplatform *VirtualPlatform_Vplatform) GetEntityData() *types.CommonEntityData {
    vplatform.EntityData.YFilter = vplatform.YFilter
    vplatform.EntityData.YangName = "vplatform"
    vplatform.EntityData.BundleName = "cisco_ios_xr"
    vplatform.EntityData.ParentYangName = "virtual-platform"
    vplatform.EntityData.SegmentPath = "vplatform"
    vplatform.EntityData.AbsolutePath = "vplatform:virtual-platform/" + vplatform.EntityData.SegmentPath
    vplatform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vplatform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vplatform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vplatform.EntityData.Children = types.NewOrderedMap()
    vplatform.EntityData.Children.Append("summary", types.YChild{"Summary", nil})
    for i := range vplatform.Summary {
        types.SetYListKey(vplatform.Summary[i], i)
        vplatform.EntityData.Children.Append(types.GetSegmentPath(vplatform.Summary[i]), types.YChild{"Summary", vplatform.Summary[i]})
    }
    vplatform.EntityData.Leafs = types.NewOrderedMap()

    vplatform.EntityData.YListKeys = []string {}

    return &(vplatform.EntityData)
}

// VirtualPlatform_Vplatform_Summary
type VirtualPlatform_Vplatform_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    Lines interface{}
}

func (summary *VirtualPlatform_Vplatform_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "vplatform"
    summary.EntityData.SegmentPath = "summary" + types.AddNoKeyToken(summary)
    summary.EntityData.AbsolutePath = "vplatform:virtual-platform/vplatform/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("lines", types.YLeaf{"Lines", summary.Lines})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// VirtualPlatform_Udi
type VirtualPlatform_Udi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    History VirtualPlatform_Udi_History
}

func (udi *VirtualPlatform_Udi) GetEntityData() *types.CommonEntityData {
    udi.EntityData.YFilter = udi.YFilter
    udi.EntityData.YangName = "udi"
    udi.EntityData.BundleName = "cisco_ios_xr"
    udi.EntityData.ParentYangName = "virtual-platform"
    udi.EntityData.SegmentPath = "udi"
    udi.EntityData.AbsolutePath = "vplatform:virtual-platform/" + udi.EntityData.SegmentPath
    udi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udi.EntityData.Children = types.NewOrderedMap()
    udi.EntityData.Children.Append("history", types.YChild{"History", &udi.History})
    udi.EntityData.Leafs = types.NewOrderedMap()

    udi.EntityData.YListKeys = []string {}

    return &(udi.EntityData)
}

// VirtualPlatform_Udi_History
type VirtualPlatform_Udi_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of VirtualPlatform_Udi_History_UdiHis.
    UdiHis []*VirtualPlatform_Udi_History_UdiHis
}

func (history *VirtualPlatform_Udi_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "udi"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "vplatform:virtual-platform/udi/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("udi_his", types.YChild{"UdiHis", nil})
    for i := range history.UdiHis {
        types.SetYListKey(history.UdiHis[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.UdiHis[i]), types.YChild{"UdiHis", history.UdiHis[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// VirtualPlatform_Udi_History_UdiHis
type VirtualPlatform_Udi_History_UdiHis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    HistoryRecords interface{}
}

func (udiHis *VirtualPlatform_Udi_History_UdiHis) GetEntityData() *types.CommonEntityData {
    udiHis.EntityData.YFilter = udiHis.YFilter
    udiHis.EntityData.YangName = "udi_his"
    udiHis.EntityData.BundleName = "cisco_ios_xr"
    udiHis.EntityData.ParentYangName = "history"
    udiHis.EntityData.SegmentPath = "udi_his" + types.AddNoKeyToken(udiHis)
    udiHis.EntityData.AbsolutePath = "vplatform:virtual-platform/udi/history/" + udiHis.EntityData.SegmentPath
    udiHis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udiHis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udiHis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udiHis.EntityData.Children = types.NewOrderedMap()
    udiHis.EntityData.Leafs = types.NewOrderedMap()
    udiHis.EntityData.Leafs.Append("history_records", types.YLeaf{"HistoryRecords", udiHis.HistoryRecords})

    udiHis.EntityData.YListKeys = []string {}

    return &(udiHis.EntityData)
}

// VirtualPlatform_Cpu
type VirtualPlatform_Cpu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    History VirtualPlatform_Cpu_History
}

func (cpu *VirtualPlatform_Cpu) GetEntityData() *types.CommonEntityData {
    cpu.EntityData.YFilter = cpu.YFilter
    cpu.EntityData.YangName = "cpu"
    cpu.EntityData.BundleName = "cisco_ios_xr"
    cpu.EntityData.ParentYangName = "virtual-platform"
    cpu.EntityData.SegmentPath = "cpu"
    cpu.EntityData.AbsolutePath = "vplatform:virtual-platform/" + cpu.EntityData.SegmentPath
    cpu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpu.EntityData.Children = types.NewOrderedMap()
    cpu.EntityData.Children.Append("history", types.YChild{"History", &cpu.History})
    cpu.EntityData.Leafs = types.NewOrderedMap()

    cpu.EntityData.YListKeys = []string {}

    return &(cpu.EntityData)
}

// VirtualPlatform_Cpu_History
type VirtualPlatform_Cpu_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of VirtualPlatform_Cpu_History_CpuUtl.
    CpuUtl []*VirtualPlatform_Cpu_History_CpuUtl
}

func (history *VirtualPlatform_Cpu_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "cpu"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "vplatform:virtual-platform/cpu/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("cpu_utl", types.YChild{"CpuUtl", nil})
    for i := range history.CpuUtl {
        types.SetYListKey(history.CpuUtl[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.CpuUtl[i]), types.YChild{"CpuUtl", history.CpuUtl[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// VirtualPlatform_Cpu_History_CpuUtl
type VirtualPlatform_Cpu_History_CpuUtl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    HistoryRecords interface{}
}

func (cpuUtl *VirtualPlatform_Cpu_History_CpuUtl) GetEntityData() *types.CommonEntityData {
    cpuUtl.EntityData.YFilter = cpuUtl.YFilter
    cpuUtl.EntityData.YangName = "cpu_utl"
    cpuUtl.EntityData.BundleName = "cisco_ios_xr"
    cpuUtl.EntityData.ParentYangName = "history"
    cpuUtl.EntityData.SegmentPath = "cpu_utl" + types.AddNoKeyToken(cpuUtl)
    cpuUtl.EntityData.AbsolutePath = "vplatform:virtual-platform/cpu/history/" + cpuUtl.EntityData.SegmentPath
    cpuUtl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpuUtl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpuUtl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpuUtl.EntityData.Children = types.NewOrderedMap()
    cpuUtl.EntityData.Leafs = types.NewOrderedMap()
    cpuUtl.EntityData.Leafs.Append("history_records", types.YLeaf{"HistoryRecords", cpuUtl.HistoryRecords})

    cpuUtl.EntityData.YListKeys = []string {}

    return &(cpuUtl.EntityData)
}

// VirtualPlatform_Processor
type VirtualPlatform_Processor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    History VirtualPlatform_Processor_History
}

func (processor *VirtualPlatform_Processor) GetEntityData() *types.CommonEntityData {
    processor.EntityData.YFilter = processor.YFilter
    processor.EntityData.YangName = "processor"
    processor.EntityData.BundleName = "cisco_ios_xr"
    processor.EntityData.ParentYangName = "virtual-platform"
    processor.EntityData.SegmentPath = "processor"
    processor.EntityData.AbsolutePath = "vplatform:virtual-platform/" + processor.EntityData.SegmentPath
    processor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processor.EntityData.Children = types.NewOrderedMap()
    processor.EntityData.Children.Append("history", types.YChild{"History", &processor.History})
    processor.EntityData.Leafs = types.NewOrderedMap()

    processor.EntityData.YListKeys = []string {}

    return &(processor.EntityData)
}

// VirtualPlatform_Processor_History
type VirtualPlatform_Processor_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of VirtualPlatform_Processor_History_ProcInfo.
    ProcInfo []*VirtualPlatform_Processor_History_ProcInfo
}

func (history *VirtualPlatform_Processor_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "processor"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "vplatform:virtual-platform/processor/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("proc_info", types.YChild{"ProcInfo", nil})
    for i := range history.ProcInfo {
        types.SetYListKey(history.ProcInfo[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.ProcInfo[i]), types.YChild{"ProcInfo", history.ProcInfo[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// VirtualPlatform_Processor_History_ProcInfo
type VirtualPlatform_Processor_History_ProcInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    HistoryRecords interface{}
}

func (procInfo *VirtualPlatform_Processor_History_ProcInfo) GetEntityData() *types.CommonEntityData {
    procInfo.EntityData.YFilter = procInfo.YFilter
    procInfo.EntityData.YangName = "proc_info"
    procInfo.EntityData.BundleName = "cisco_ios_xr"
    procInfo.EntityData.ParentYangName = "history"
    procInfo.EntityData.SegmentPath = "proc_info" + types.AddNoKeyToken(procInfo)
    procInfo.EntityData.AbsolutePath = "vplatform:virtual-platform/processor/history/" + procInfo.EntityData.SegmentPath
    procInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procInfo.EntityData.Children = types.NewOrderedMap()
    procInfo.EntityData.Leafs = types.NewOrderedMap()
    procInfo.EntityData.Leafs.Append("history_records", types.YLeaf{"HistoryRecords", procInfo.HistoryRecords})

    procInfo.EntityData.YListKeys = []string {}

    return &(procInfo.EntityData)
}

// VirtualPlatform_Memory
type VirtualPlatform_Memory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    History VirtualPlatform_Memory_History
}

func (memory *VirtualPlatform_Memory) GetEntityData() *types.CommonEntityData {
    memory.EntityData.YFilter = memory.YFilter
    memory.EntityData.YangName = "memory"
    memory.EntityData.BundleName = "cisco_ios_xr"
    memory.EntityData.ParentYangName = "virtual-platform"
    memory.EntityData.SegmentPath = "memory"
    memory.EntityData.AbsolutePath = "vplatform:virtual-platform/" + memory.EntityData.SegmentPath
    memory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memory.EntityData.Children = types.NewOrderedMap()
    memory.EntityData.Children.Append("history", types.YChild{"History", &memory.History})
    memory.EntityData.Leafs = types.NewOrderedMap()

    memory.EntityData.YListKeys = []string {}

    return &(memory.EntityData)
}

// VirtualPlatform_Memory_History
type VirtualPlatform_Memory_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of VirtualPlatform_Memory_History_MemUsg.
    MemUsg []*VirtualPlatform_Memory_History_MemUsg
}

func (history *VirtualPlatform_Memory_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "memory"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "vplatform:virtual-platform/memory/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("mem_usg", types.YChild{"MemUsg", nil})
    for i := range history.MemUsg {
        types.SetYListKey(history.MemUsg[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.MemUsg[i]), types.YChild{"MemUsg", history.MemUsg[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// VirtualPlatform_Memory_History_MemUsg
type VirtualPlatform_Memory_History_MemUsg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    HistoryRecords interface{}
}

func (memUsg *VirtualPlatform_Memory_History_MemUsg) GetEntityData() *types.CommonEntityData {
    memUsg.EntityData.YFilter = memUsg.YFilter
    memUsg.EntityData.YangName = "mem_usg"
    memUsg.EntityData.BundleName = "cisco_ios_xr"
    memUsg.EntityData.ParentYangName = "history"
    memUsg.EntityData.SegmentPath = "mem_usg" + types.AddNoKeyToken(memUsg)
    memUsg.EntityData.AbsolutePath = "vplatform:virtual-platform/memory/history/" + memUsg.EntityData.SegmentPath
    memUsg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memUsg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memUsg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memUsg.EntityData.Children = types.NewOrderedMap()
    memUsg.EntityData.Leafs = types.NewOrderedMap()
    memUsg.EntityData.Leafs.Append("history_records", types.YLeaf{"HistoryRecords", memUsg.HistoryRecords})

    memUsg.EntityData.YListKeys = []string {}

    return &(memUsg.EntityData)
}

// VirtualPlatform_Disk
type VirtualPlatform_Disk struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    History VirtualPlatform_Disk_History
}

func (disk *VirtualPlatform_Disk) GetEntityData() *types.CommonEntityData {
    disk.EntityData.YFilter = disk.YFilter
    disk.EntityData.YangName = "disk"
    disk.EntityData.BundleName = "cisco_ios_xr"
    disk.EntityData.ParentYangName = "virtual-platform"
    disk.EntityData.SegmentPath = "disk"
    disk.EntityData.AbsolutePath = "vplatform:virtual-platform/" + disk.EntityData.SegmentPath
    disk.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disk.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disk.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disk.EntityData.Children = types.NewOrderedMap()
    disk.EntityData.Children.Append("history", types.YChild{"History", &disk.History})
    disk.EntityData.Leafs = types.NewOrderedMap()

    disk.EntityData.YListKeys = []string {}

    return &(disk.EntityData)
}

// VirtualPlatform_Disk_History
type VirtualPlatform_Disk_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of VirtualPlatform_Disk_History_DiskUtl.
    DiskUtl []*VirtualPlatform_Disk_History_DiskUtl
}

func (history *VirtualPlatform_Disk_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "disk"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "vplatform:virtual-platform/disk/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("disk_utl", types.YChild{"DiskUtl", nil})
    for i := range history.DiskUtl {
        types.SetYListKey(history.DiskUtl[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.DiskUtl[i]), types.YChild{"DiskUtl", history.DiskUtl[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// VirtualPlatform_Disk_History_DiskUtl
type VirtualPlatform_Disk_History_DiskUtl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string.
    HistoryRecords interface{}
}

func (diskUtl *VirtualPlatform_Disk_History_DiskUtl) GetEntityData() *types.CommonEntityData {
    diskUtl.EntityData.YFilter = diskUtl.YFilter
    diskUtl.EntityData.YangName = "disk_utl"
    diskUtl.EntityData.BundleName = "cisco_ios_xr"
    diskUtl.EntityData.ParentYangName = "history"
    diskUtl.EntityData.SegmentPath = "disk_utl" + types.AddNoKeyToken(diskUtl)
    diskUtl.EntityData.AbsolutePath = "vplatform:virtual-platform/disk/history/" + diskUtl.EntityData.SegmentPath
    diskUtl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diskUtl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diskUtl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diskUtl.EntityData.Children = types.NewOrderedMap()
    diskUtl.EntityData.Leafs = types.NewOrderedMap()
    diskUtl.EntityData.Leafs.Append("history_records", types.YLeaf{"HistoryRecords", diskUtl.HistoryRecords})

    diskUtl.EntityData.YListKeys = []string {}

    return &(diskUtl.EntityData)
}

// VirtualPlatform_Action
type VirtualPlatform_Action struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UDI Operations. The type is slice of VirtualPlatform_Action_Udi.
    Udi []*VirtualPlatform_Action_Udi
}

func (action *VirtualPlatform_Action) GetEntityData() *types.CommonEntityData {
    action.EntityData.YFilter = action.YFilter
    action.EntityData.YangName = "Action"
    action.EntityData.BundleName = "cisco_ios_xr"
    action.EntityData.ParentYangName = "virtual-platform"
    action.EntityData.SegmentPath = "Action"
    action.EntityData.AbsolutePath = "vplatform:virtual-platform/" + action.EntityData.SegmentPath
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = types.NewOrderedMap()
    action.EntityData.Children.Append("udi", types.YChild{"Udi", nil})
    for i := range action.Udi {
        types.SetYListKey(action.Udi[i], i)
        action.EntityData.Children.Append(types.GetSegmentPath(action.Udi[i]), types.YChild{"Udi", action.Udi[i]})
    }
    action.EntityData.Leafs = types.NewOrderedMap()

    action.EntityData.YListKeys = []string {}

    return &(action.EntityData)
}

// VirtualPlatform_Action_Udi
// UDI Operations
type VirtualPlatform_Action_Udi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string
}

func (udi *VirtualPlatform_Action_Udi) GetEntityData() *types.CommonEntityData {
    udi.EntityData.YFilter = udi.YFilter
    udi.EntityData.YangName = "udi"
    udi.EntityData.BundleName = "cisco_ios_xr"
    udi.EntityData.ParentYangName = "Action"
    udi.EntityData.SegmentPath = "udi" + types.AddNoKeyToken(udi)
    udi.EntityData.AbsolutePath = "vplatform:virtual-platform/Action/" + udi.EntityData.SegmentPath
    udi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udi.EntityData.Children = types.NewOrderedMap()
    udi.EntityData.Leafs = types.NewOrderedMap()

    udi.EntityData.YListKeys = []string {}

    return &(udi.EntityData)
}

