// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_show_trace_vmm

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_show_trace_vmm"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-show-trace-vmm vmm}", reflect.TypeOf(Vmm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-show-trace-vmm:vmm", reflect.TypeOf(Vmm{}))
}

// Vmm
type Vmm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    VmManager Vmm_VmManager
}

func (vmm *Vmm) GetEntityData() *types.CommonEntityData {
    vmm.EntityData.YFilter = vmm.YFilter
    vmm.EntityData.YangName = "vmm"
    vmm.EntityData.BundleName = "cisco_ios_xr"
    vmm.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-show-trace-vmm"
    vmm.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-show-trace-vmm:vmm"
    vmm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vmm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vmm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vmm.EntityData.Children = types.NewOrderedMap()
    vmm.EntityData.Children.Append("vm_manager", types.YChild{"VmManager", &vmm.VmManager})
    vmm.EntityData.Leafs = types.NewOrderedMap()

    vmm.EntityData.YListKeys = []string {}

    return &(vmm.EntityData)
}

// Vmm_VmManager
type Vmm_VmManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Vmm_VmManager_Trace.
    Trace []*Vmm_VmManager_Trace
}

func (vmManager *Vmm_VmManager) GetEntityData() *types.CommonEntityData {
    vmManager.EntityData.YFilter = vmManager.YFilter
    vmManager.EntityData.YangName = "vm_manager"
    vmManager.EntityData.BundleName = "cisco_ios_xr"
    vmManager.EntityData.ParentYangName = "vmm"
    vmManager.EntityData.SegmentPath = "vm_manager"
    vmManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vmManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vmManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vmManager.EntityData.Children = types.NewOrderedMap()
    vmManager.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range vmManager.Trace {
        vmManager.EntityData.Children.Append(types.GetSegmentPath(vmManager.Trace[i]), types.YChild{"Trace", vmManager.Trace[i]})
    }
    vmManager.EntityData.Leafs = types.NewOrderedMap()

    vmManager.EntityData.YListKeys = []string {}

    return &(vmManager.EntityData)
}

// Vmm_VmManager_Trace
// show traceable processes
type Vmm_VmManager_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Vmm_VmManager_Trace_Location.
    Location []*Vmm_VmManager_Trace_Location
}

func (trace *Vmm_VmManager_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "vm_manager"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range trace.Location {
        trace.EntityData.Children.Append(types.GetSegmentPath(trace.Location[i]), types.YChild{"Location", trace.Location[i]})
    }
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("buffer", types.YLeaf{"Buffer", trace.Buffer})

    trace.EntityData.YListKeys = []string {"Buffer"}

    return &(trace.EntityData)
}

// Vmm_VmManager_Trace_Location
type Vmm_VmManager_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Vmm_VmManager_Trace_Location_AllOptions.
    AllOptions []*Vmm_VmManager_Trace_Location_AllOptions
}

func (location *Vmm_VmManager_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("all-options", types.YChild{"AllOptions", nil})
    for i := range location.AllOptions {
        location.EntityData.Children.Append(types.GetSegmentPath(location.AllOptions[i]), types.YChild{"AllOptions", location.AllOptions[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location_name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// Vmm_VmManager_Trace_Location_AllOptions
type Vmm_VmManager_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Vmm_VmManager_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Vmm_VmManager_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Vmm_VmManager_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// Vmm_VmManager_Trace_Location_AllOptions_TraceBlocks
type Vmm_VmManager_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Vmm_VmManager_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks"
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

