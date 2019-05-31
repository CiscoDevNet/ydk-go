// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_show_trace_cm

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_show_trace_cm"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-show-trace-cm cm}", reflect.TypeOf(Cm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-show-trace-cm:cm", reflect.TypeOf(Cm{}))
}

// Cm
type Cm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Cm_Trace.
    Trace []*Cm_Trace

    // System Admin Manager lspdb of a location.
    Lsp Cm_Lsp
}

func (cm *Cm) GetEntityData() *types.CommonEntityData {
    cm.EntityData.YFilter = cm.YFilter
    cm.EntityData.YangName = "cm"
    cm.EntityData.BundleName = "cisco_ios_xr"
    cm.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-show-trace-cm"
    cm.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-show-trace-cm:cm"
    cm.EntityData.AbsolutePath = cm.EntityData.SegmentPath
    cm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cm.EntityData.Children = types.NewOrderedMap()
    cm.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range cm.Trace {
        cm.EntityData.Children.Append(types.GetSegmentPath(cm.Trace[i]), types.YChild{"Trace", cm.Trace[i]})
    }
    cm.EntityData.Children.Append("Cisco-IOS-XR-sysadmin-cm:lsp", types.YChild{"Lsp", &cm.Lsp})
    cm.EntityData.Leafs = types.NewOrderedMap()

    cm.EntityData.YListKeys = []string {}

    return &(cm.EntityData)
}

// Cm_Trace
// show traceable processes
type Cm_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Cm_Trace_Location.
    Location []*Cm_Trace_Location
}

func (trace *Cm_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "cm"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-trace-cm:cm/" + trace.EntityData.SegmentPath
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

// Cm_Trace_Location
type Cm_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Cm_Trace_Location_AllOptions.
    AllOptions []*Cm_Trace_Location_AllOptions
}

func (location *Cm_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-trace-cm:cm/trace/" + location.EntityData.SegmentPath
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

// Cm_Trace_Location_AllOptions
type Cm_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Cm_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Cm_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Cm_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-trace-cm:cm/trace/location/" + allOptions.EntityData.SegmentPath
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

// Cm_Trace_Location_AllOptions_TraceBlocks
type Cm_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Cm_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-trace-cm:cm/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

// Cm_Lsp
// System Admin Manager lspdb of a location
type Cm_Lsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Cm_Lsp_LspdbLocations.
    LspdbLocations []*Cm_Lsp_LspdbLocations
}

func (lsp *Cm_Lsp) GetEntityData() *types.CommonEntityData {
    lsp.EntityData.YFilter = lsp.YFilter
    lsp.EntityData.YangName = "lsp"
    lsp.EntityData.BundleName = "cisco_ios_xr"
    lsp.EntityData.ParentYangName = "cm"
    lsp.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:lsp"
    lsp.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-trace-cm:cm/" + lsp.EntityData.SegmentPath
    lsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsp.EntityData.Children = types.NewOrderedMap()
    lsp.EntityData.Children.Append("lspdb_locations", types.YChild{"LspdbLocations", nil})
    for i := range lsp.LspdbLocations {
        lsp.EntityData.Children.Append(types.GetSegmentPath(lsp.LspdbLocations[i]), types.YChild{"LspdbLocations", lsp.LspdbLocations[i]})
    }
    lsp.EntityData.Leafs = types.NewOrderedMap()

    lsp.EntityData.YListKeys = []string {}

    return &(lsp.EntityData)
}

// Cm_Lsp_LspdbLocations
type Cm_Lsp_LspdbLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'((([bB][0-9])/(([a-zA-Z]){2}\\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    LspdbLocation interface{}

    // The type is slice of Cm_Lsp_LspdbLocations_Lspdbi.
    Lspdbi []*Cm_Lsp_LspdbLocations_Lspdbi
}

func (lspdbLocations *Cm_Lsp_LspdbLocations) GetEntityData() *types.CommonEntityData {
    lspdbLocations.EntityData.YFilter = lspdbLocations.YFilter
    lspdbLocations.EntityData.YangName = "lspdb_locations"
    lspdbLocations.EntityData.BundleName = "cisco_ios_xr"
    lspdbLocations.EntityData.ParentYangName = "lsp"
    lspdbLocations.EntityData.SegmentPath = "lspdb_locations" + types.AddKeyToken(lspdbLocations.LspdbLocation, "lspdb_location")
    lspdbLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-trace-cm:cm/Cisco-IOS-XR-sysadmin-cm:lsp/" + lspdbLocations.EntityData.SegmentPath
    lspdbLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspdbLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspdbLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspdbLocations.EntityData.Children = types.NewOrderedMap()
    lspdbLocations.EntityData.Children.Append("lspdbi", types.YChild{"Lspdbi", nil})
    for i := range lspdbLocations.Lspdbi {
        lspdbLocations.EntityData.Children.Append(types.GetSegmentPath(lspdbLocations.Lspdbi[i]), types.YChild{"Lspdbi", lspdbLocations.Lspdbi[i]})
    }
    lspdbLocations.EntityData.Leafs = types.NewOrderedMap()
    lspdbLocations.EntityData.Leafs.Append("lspdb_location", types.YLeaf{"LspdbLocation", lspdbLocations.LspdbLocation})

    lspdbLocations.EntityData.YListKeys = []string {"LspdbLocation"}

    return &(lspdbLocations.EntityData)
}

// Cm_Lsp_LspdbLocations_Lspdbi
type Cm_Lsp_LspdbLocations_Lspdbi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSP System ID. The type is string.
    LspId interface{}

    // This attribute is a key. LSP Area Type. The type is AreaType.
    LspAreaType interface{}

    // LSP sequence number. The type is interface{} with range: 0..4294967295.
    LspSequence interface{}

    // LSP core data bits. The type is string.
    LspCore interface{}

    // LSP TLV data. The type is slice of string.
    LspTlvs []interface{}
}

func (lspdbi *Cm_Lsp_LspdbLocations_Lspdbi) GetEntityData() *types.CommonEntityData {
    lspdbi.EntityData.YFilter = lspdbi.YFilter
    lspdbi.EntityData.YangName = "lspdbi"
    lspdbi.EntityData.BundleName = "cisco_ios_xr"
    lspdbi.EntityData.ParentYangName = "lspdb_locations"
    lspdbi.EntityData.SegmentPath = "lspdbi" + types.AddKeyToken(lspdbi.LspId, "lsp_id") + types.AddKeyToken(lspdbi.LspAreaType, "lsp_area_type")
    lspdbi.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-show-trace-cm:cm/Cisco-IOS-XR-sysadmin-cm:lsp/lspdb_locations/" + lspdbi.EntityData.SegmentPath
    lspdbi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspdbi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspdbi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspdbi.EntityData.Children = types.NewOrderedMap()
    lspdbi.EntityData.Leafs = types.NewOrderedMap()
    lspdbi.EntityData.Leafs.Append("lsp_id", types.YLeaf{"LspId", lspdbi.LspId})
    lspdbi.EntityData.Leafs.Append("lsp_area_type", types.YLeaf{"LspAreaType", lspdbi.LspAreaType})
    lspdbi.EntityData.Leafs.Append("lsp_sequence", types.YLeaf{"LspSequence", lspdbi.LspSequence})
    lspdbi.EntityData.Leafs.Append("lsp_core", types.YLeaf{"LspCore", lspdbi.LspCore})
    lspdbi.EntityData.Leafs.Append("lsp_tlvs", types.YLeaf{"LspTlvs", lspdbi.LspTlvs})

    lspdbi.EntityData.YListKeys = []string {"LspId", "LspAreaType"}

    return &(lspdbi.EntityData)
}

