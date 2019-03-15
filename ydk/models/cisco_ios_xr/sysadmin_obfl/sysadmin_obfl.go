// This module contains definitions
// for the Calvados model objects.
// 
// This module holds OBFL data.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_obfl

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_obfl"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-obfl obfl}", reflect.TypeOf(Obfl{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-obfl:obfl", reflect.TypeOf(Obfl{}))
}

// Obfl
type Obfl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    ObflMgr Obfl_ObflMgr

    
    ObflShow Obfl_ObflShow
}

func (obfl *Obfl) GetEntityData() *types.CommonEntityData {
    obfl.EntityData.YFilter = obfl.YFilter
    obfl.EntityData.YangName = "obfl"
    obfl.EntityData.BundleName = "cisco_ios_xr"
    obfl.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-obfl"
    obfl.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-obfl:obfl"
    obfl.EntityData.AbsolutePath = obfl.EntityData.SegmentPath
    obfl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    obfl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    obfl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    obfl.EntityData.Children = types.NewOrderedMap()
    obfl.EntityData.Children.Append("obfl_mgr", types.YChild{"ObflMgr", &obfl.ObflMgr})
    obfl.EntityData.Children.Append("obfl_show", types.YChild{"ObflShow", &obfl.ObflShow})
    obfl.EntityData.Leafs = types.NewOrderedMap()

    obfl.EntityData.YListKeys = []string {}

    return &(obfl.EntityData)
}

// Obfl_ObflMgr
type Obfl_ObflMgr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Obfl_ObflMgr_Trace.
    Trace []*Obfl_ObflMgr_Trace
}

func (obflMgr *Obfl_ObflMgr) GetEntityData() *types.CommonEntityData {
    obflMgr.EntityData.YFilter = obflMgr.YFilter
    obflMgr.EntityData.YangName = "obfl_mgr"
    obflMgr.EntityData.BundleName = "cisco_ios_xr"
    obflMgr.EntityData.ParentYangName = "obfl"
    obflMgr.EntityData.SegmentPath = "obfl_mgr"
    obflMgr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/" + obflMgr.EntityData.SegmentPath
    obflMgr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    obflMgr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    obflMgr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    obflMgr.EntityData.Children = types.NewOrderedMap()
    obflMgr.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range obflMgr.Trace {
        obflMgr.EntityData.Children.Append(types.GetSegmentPath(obflMgr.Trace[i]), types.YChild{"Trace", obflMgr.Trace[i]})
    }
    obflMgr.EntityData.Leafs = types.NewOrderedMap()

    obflMgr.EntityData.YListKeys = []string {}

    return &(obflMgr.EntityData)
}

// Obfl_ObflMgr_Trace
// show traceable processes
type Obfl_ObflMgr_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Obfl_ObflMgr_Trace_Location.
    Location []*Obfl_ObflMgr_Trace_Location
}

func (trace *Obfl_ObflMgr_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "obfl_mgr"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/obfl_mgr/" + trace.EntityData.SegmentPath
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

// Obfl_ObflMgr_Trace_Location
type Obfl_ObflMgr_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Obfl_ObflMgr_Trace_Location_AllOptions.
    AllOptions []*Obfl_ObflMgr_Trace_Location_AllOptions
}

func (location *Obfl_ObflMgr_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/obfl_mgr/trace/" + location.EntityData.SegmentPath
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

// Obfl_ObflMgr_Trace_Location_AllOptions
type Obfl_ObflMgr_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Obfl_ObflMgr_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/obfl_mgr/trace/location/" + allOptions.EntityData.SegmentPath
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        types.SetYListKey(allOptions.TraceBlocks[i], i)
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks
type Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Obfl_ObflMgr_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/obfl_mgr/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

// Obfl_ObflShow
type Obfl_ObflShow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Obfl_ObflShow_Trace.
    Trace []*Obfl_ObflShow_Trace
}

func (obflShow *Obfl_ObflShow) GetEntityData() *types.CommonEntityData {
    obflShow.EntityData.YFilter = obflShow.YFilter
    obflShow.EntityData.YangName = "obfl_show"
    obflShow.EntityData.BundleName = "cisco_ios_xr"
    obflShow.EntityData.ParentYangName = "obfl"
    obflShow.EntityData.SegmentPath = "obfl_show"
    obflShow.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/" + obflShow.EntityData.SegmentPath
    obflShow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    obflShow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    obflShow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    obflShow.EntityData.Children = types.NewOrderedMap()
    obflShow.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range obflShow.Trace {
        obflShow.EntityData.Children.Append(types.GetSegmentPath(obflShow.Trace[i]), types.YChild{"Trace", obflShow.Trace[i]})
    }
    obflShow.EntityData.Leafs = types.NewOrderedMap()

    obflShow.EntityData.YListKeys = []string {}

    return &(obflShow.EntityData)
}

// Obfl_ObflShow_Trace
// show traceable processes
type Obfl_ObflShow_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Obfl_ObflShow_Trace_Location.
    Location []*Obfl_ObflShow_Trace_Location
}

func (trace *Obfl_ObflShow_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "obfl_show"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/obfl_show/" + trace.EntityData.SegmentPath
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

// Obfl_ObflShow_Trace_Location
type Obfl_ObflShow_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Obfl_ObflShow_Trace_Location_AllOptions.
    AllOptions []*Obfl_ObflShow_Trace_Location_AllOptions
}

func (location *Obfl_ObflShow_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/obfl_show/trace/" + location.EntityData.SegmentPath
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

// Obfl_ObflShow_Trace_Location_AllOptions
type Obfl_ObflShow_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Obfl_ObflShow_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/obfl_show/trace/location/" + allOptions.EntityData.SegmentPath
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        types.SetYListKey(allOptions.TraceBlocks[i], i)
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks
type Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Obfl_ObflShow_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-obfl:obfl/obfl_show/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

