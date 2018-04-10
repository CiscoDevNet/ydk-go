// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2017 by Cisco Systems, Inc.
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
    Trace []Cm_Trace

    // System Admin Manager lspdb of a location.
    Lsp Cm_Lsp
}

func (cm *Cm) GetEntityData() *types.CommonEntityData {
    cm.EntityData.YFilter = cm.YFilter
    cm.EntityData.YangName = "cm"
    cm.EntityData.BundleName = "cisco_ios_xr"
    cm.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-show-trace-cm"
    cm.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-show-trace-cm:cm"
    cm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cm.EntityData.Children = make(map[string]types.YChild)
    cm.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range cm.Trace {
        cm.EntityData.Children[types.GetSegmentPath(&cm.Trace[i])] = types.YChild{"Trace", &cm.Trace[i]}
    }
    cm.EntityData.Children["Cisco-IOS-XR-sysadmin-cm:lsp"] = types.YChild{"Lsp", &cm.Lsp}
    cm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cm.EntityData)
}

// Cm_Trace
// show traceable processes
type Cm_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Cm_Trace_Location.
    Location []Cm_Trace_Location
}

func (trace *Cm_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "cm"
    trace.EntityData.SegmentPath = "trace" + "[buffer='" + fmt.Sprintf("%v", trace.Buffer) + "']"
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = make(map[string]types.YChild)
    trace.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range trace.Location {
        trace.EntityData.Children[types.GetSegmentPath(&trace.Location[i])] = types.YChild{"Location", &trace.Location[i]}
    }
    trace.EntityData.Leafs = make(map[string]types.YLeaf)
    trace.EntityData.Leafs["buffer"] = types.YLeaf{"Buffer", trace.Buffer}
    return &(trace.EntityData)
}

// Cm_Trace_Location
type Cm_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Cm_Trace_Location_AllOptions.
    AllOptions []Cm_Trace_Location_AllOptions
}

func (location *Cm_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + "[location_name='" + fmt.Sprintf("%v", location.LocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["all-options"] = types.YChild{"AllOptions", nil}
    for i := range location.AllOptions {
        location.EntityData.Children[types.GetSegmentPath(&location.AllOptions[i])] = types.YChild{"AllOptions", &location.AllOptions[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location_name"] = types.YLeaf{"LocationName", location.LocationName}
    return &(location.EntityData)
}

// Cm_Trace_Location_AllOptions
type Cm_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Cm_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Cm_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Cm_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + "[option='" + fmt.Sprintf("%v", allOptions.Option) + "']"
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = make(map[string]types.YChild)
    allOptions.EntityData.Children["trace-blocks"] = types.YChild{"TraceBlocks", nil}
    for i := range allOptions.TraceBlocks {
        allOptions.EntityData.Children[types.GetSegmentPath(&allOptions.TraceBlocks[i])] = types.YChild{"TraceBlocks", &allOptions.TraceBlocks[i]}
    }
    allOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    allOptions.EntityData.Leafs["option"] = types.YLeaf{"Option", allOptions.Option}
    return &(allOptions.EntityData)
}

// Cm_Trace_Location_AllOptions_TraceBlocks
type Cm_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Cm_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks"
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = make(map[string]types.YChild)
    traceBlocks.EntityData.Leafs = make(map[string]types.YLeaf)
    traceBlocks.EntityData.Leafs["data"] = types.YLeaf{"Data", traceBlocks.Data}
    return &(traceBlocks.EntityData)
}

// Cm_Lsp
// System Admin Manager lspdb of a location
type Cm_Lsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Cm_Lsp_LspdbLocations.
    LspdbLocations []Cm_Lsp_LspdbLocations
}

func (lsp *Cm_Lsp) GetEntityData() *types.CommonEntityData {
    lsp.EntityData.YFilter = lsp.YFilter
    lsp.EntityData.YangName = "lsp"
    lsp.EntityData.BundleName = "cisco_ios_xr"
    lsp.EntityData.ParentYangName = "cm"
    lsp.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-cm:lsp"
    lsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsp.EntityData.Children = make(map[string]types.YChild)
    lsp.EntityData.Children["lspdb_locations"] = types.YChild{"LspdbLocations", nil}
    for i := range lsp.LspdbLocations {
        lsp.EntityData.Children[types.GetSegmentPath(&lsp.LspdbLocations[i])] = types.YChild{"LspdbLocations", &lsp.LspdbLocations[i]}
    }
    lsp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lsp.EntityData)
}

// Cm_Lsp_LspdbLocations
type Cm_Lsp_LspdbLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    LspdbLocation interface{}

    // The type is slice of Cm_Lsp_LspdbLocations_Lspdbi.
    Lspdbi []Cm_Lsp_LspdbLocations_Lspdbi
}

func (lspdbLocations *Cm_Lsp_LspdbLocations) GetEntityData() *types.CommonEntityData {
    lspdbLocations.EntityData.YFilter = lspdbLocations.YFilter
    lspdbLocations.EntityData.YangName = "lspdb_locations"
    lspdbLocations.EntityData.BundleName = "cisco_ios_xr"
    lspdbLocations.EntityData.ParentYangName = "lsp"
    lspdbLocations.EntityData.SegmentPath = "lspdb_locations" + "[lspdb_location='" + fmt.Sprintf("%v", lspdbLocations.LspdbLocation) + "']"
    lspdbLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspdbLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspdbLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspdbLocations.EntityData.Children = make(map[string]types.YChild)
    lspdbLocations.EntityData.Children["lspdbi"] = types.YChild{"Lspdbi", nil}
    for i := range lspdbLocations.Lspdbi {
        lspdbLocations.EntityData.Children[types.GetSegmentPath(&lspdbLocations.Lspdbi[i])] = types.YChild{"Lspdbi", &lspdbLocations.Lspdbi[i]}
    }
    lspdbLocations.EntityData.Leafs = make(map[string]types.YLeaf)
    lspdbLocations.EntityData.Leafs["lspdb_location"] = types.YLeaf{"LspdbLocation", lspdbLocations.LspdbLocation}
    return &(lspdbLocations.EntityData)
}

// Cm_Lsp_LspdbLocations_Lspdbi
type Cm_Lsp_LspdbLocations_Lspdbi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    lspdbi.EntityData.SegmentPath = "lspdbi" + "[lsp_id='" + fmt.Sprintf("%v", lspdbi.LspId) + "']" + "[lsp_area_type='" + fmt.Sprintf("%v", lspdbi.LspAreaType) + "']"
    lspdbi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspdbi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspdbi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspdbi.EntityData.Children = make(map[string]types.YChild)
    lspdbi.EntityData.Leafs = make(map[string]types.YLeaf)
    lspdbi.EntityData.Leafs["lsp_id"] = types.YLeaf{"LspId", lspdbi.LspId}
    lspdbi.EntityData.Leafs["lsp_area_type"] = types.YLeaf{"LspAreaType", lspdbi.LspAreaType}
    lspdbi.EntityData.Leafs["lsp_sequence"] = types.YLeaf{"LspSequence", lspdbi.LspSequence}
    lspdbi.EntityData.Leafs["lsp_core"] = types.YLeaf{"LspCore", lspdbi.LspCore}
    lspdbi.EntityData.Leafs["lsp_tlvs"] = types.YLeaf{"LspTlvs", lspdbi.LspTlvs}
    return &(lspdbi.EntityData)
}

