// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module holds Shelf Management configuration data.
// 
// Copyright(c) 2011-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_sm

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_sm"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-sm oper}", reflect.TypeOf(Oper{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-sm:oper", reflect.TypeOf(Oper{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-sm config}", reflect.TypeOf(Config{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-sm:config", reflect.TypeOf(Config{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-sm actions}", reflect.TypeOf(Actions{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-sm:actions", reflect.TypeOf(Actions{}))
}

// Oper
type Oper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    ShelfMgr Oper_ShelfMgr

    
    Platform Oper_Platform

    
    Chassis Oper_Chassis

    
    Reload Oper_Reload

    
    RebootHistory Oper_RebootHistory

    // The type is slice of Oper_Interface_.
    Interface_ []Oper_Interface

    
    ReloadVm Oper_ReloadVm

    
    Macpool Oper_Macpool
}

func (oper *Oper) GetEntityData() *types.CommonEntityData {
    oper.EntityData.YFilter = oper.YFilter
    oper.EntityData.YangName = "oper"
    oper.EntityData.BundleName = "cisco_ios_xr"
    oper.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-sm"
    oper.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-sm:oper"
    oper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oper.EntityData.Children = make(map[string]types.YChild)
    oper.EntityData.Children["shelf_mgr"] = types.YChild{"ShelfMgr", &oper.ShelfMgr}
    oper.EntityData.Children["platform"] = types.YChild{"Platform", &oper.Platform}
    oper.EntityData.Children["chassis"] = types.YChild{"Chassis", &oper.Chassis}
    oper.EntityData.Children["reload"] = types.YChild{"Reload", &oper.Reload}
    oper.EntityData.Children["reboot-history"] = types.YChild{"RebootHistory", &oper.RebootHistory}
    oper.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range oper.Interface_ {
        oper.EntityData.Children[types.GetSegmentPath(&oper.Interface_[i])] = types.YChild{"Interface_", &oper.Interface_[i]}
    }
    oper.EntityData.Children["reload_vm"] = types.YChild{"ReloadVm", &oper.ReloadVm}
    oper.EntityData.Children["macpool"] = types.YChild{"Macpool", &oper.Macpool}
    oper.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oper.EntityData)
}

// Oper_ShelfMgr
type Oper_ShelfMgr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Oper_ShelfMgr_Trace.
    Trace []Oper_ShelfMgr_Trace
}

func (shelfMgr *Oper_ShelfMgr) GetEntityData() *types.CommonEntityData {
    shelfMgr.EntityData.YFilter = shelfMgr.YFilter
    shelfMgr.EntityData.YangName = "shelf_mgr"
    shelfMgr.EntityData.BundleName = "cisco_ios_xr"
    shelfMgr.EntityData.ParentYangName = "oper"
    shelfMgr.EntityData.SegmentPath = "shelf_mgr"
    shelfMgr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shelfMgr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shelfMgr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shelfMgr.EntityData.Children = make(map[string]types.YChild)
    shelfMgr.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range shelfMgr.Trace {
        shelfMgr.EntityData.Children[types.GetSegmentPath(&shelfMgr.Trace[i])] = types.YChild{"Trace", &shelfMgr.Trace[i]}
    }
    shelfMgr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shelfMgr.EntityData)
}

// Oper_ShelfMgr_Trace
// show traceable processes
type Oper_ShelfMgr_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Oper_ShelfMgr_Trace_Location.
    Location []Oper_ShelfMgr_Trace_Location
}

func (trace *Oper_ShelfMgr_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "shelf_mgr"
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

// Oper_ShelfMgr_Trace_Location
type Oper_ShelfMgr_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Oper_ShelfMgr_Trace_Location_AllOptions.
    AllOptions []Oper_ShelfMgr_Trace_Location_AllOptions
}

func (location *Oper_ShelfMgr_Trace_Location) GetEntityData() *types.CommonEntityData {
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

// Oper_ShelfMgr_Trace_Location_AllOptions
type Oper_ShelfMgr_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Oper_ShelfMgr_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Oper_ShelfMgr_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Oper_ShelfMgr_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
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

// Oper_ShelfMgr_Trace_Location_AllOptions_TraceBlocks
type Oper_ShelfMgr_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Oper_ShelfMgr_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
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

// Oper_Platform
type Oper_Platform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Summary Oper_Platform_Summary

    
    Detail Oper_Platform_Detail

    
    Slices Oper_Platform_Slices
}

func (platform *Oper_Platform) GetEntityData() *types.CommonEntityData {
    platform.EntityData.YFilter = platform.YFilter
    platform.EntityData.YangName = "platform"
    platform.EntityData.BundleName = "cisco_ios_xr"
    platform.EntityData.ParentYangName = "oper"
    platform.EntityData.SegmentPath = "platform"
    platform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platform.EntityData.Children = make(map[string]types.YChild)
    platform.EntityData.Children["summary"] = types.YChild{"Summary", &platform.Summary}
    platform.EntityData.Children["detail"] = types.YChild{"Detail", &platform.Detail}
    platform.EntityData.Children["slices"] = types.YChild{"Slices", &platform.Slices}
    platform.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platform.EntityData)
}

// Oper_Platform_Summary
type Oper_Platform_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Platform_Summary_Location.
    Location []Oper_Platform_Summary_Location
}

func (summary *Oper_Platform_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "platform"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range summary.Location {
        summary.EntityData.Children[types.GetSegmentPath(&summary.Location[i])] = types.YChild{"Location", &summary.Location[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summary.EntityData)
}

// Oper_Platform_Summary_Location
type Oper_Platform_Summary_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    Location interface{}

    
    SummaryData Oper_Platform_Summary_Location_SummaryData
}

func (location *Oper_Platform_Summary_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "summary"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["summary-data"] = types.YChild{"SummaryData", &location.SummaryData}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Oper_Platform_Summary_Location_SummaryData
type Oper_Platform_Summary_Location_SummaryData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    CardType interface{}

    // The type is string.
    HwState interface{}

    // The type is string.
    SwState interface{}

    // The type is string.
    ConfigState interface{}
}

func (summaryData *Oper_Platform_Summary_Location_SummaryData) GetEntityData() *types.CommonEntityData {
    summaryData.EntityData.YFilter = summaryData.YFilter
    summaryData.EntityData.YangName = "summary-data"
    summaryData.EntityData.BundleName = "cisco_ios_xr"
    summaryData.EntityData.ParentYangName = "location"
    summaryData.EntityData.SegmentPath = "summary-data"
    summaryData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryData.EntityData.Children = make(map[string]types.YChild)
    summaryData.EntityData.Leafs = make(map[string]types.YLeaf)
    summaryData.EntityData.Leafs["card_type"] = types.YLeaf{"CardType", summaryData.CardType}
    summaryData.EntityData.Leafs["hw_state"] = types.YLeaf{"HwState", summaryData.HwState}
    summaryData.EntityData.Leafs["sw_state"] = types.YLeaf{"SwState", summaryData.SwState}
    summaryData.EntityData.Leafs["config_state"] = types.YLeaf{"ConfigState", summaryData.ConfigState}
    return &(summaryData.EntityData)
}

// Oper_Platform_Detail
type Oper_Platform_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Platform_Detail_Location.
    Location []Oper_Platform_Detail_Location
}

func (detail *Oper_Platform_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "platform"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range detail.Location {
        detail.EntityData.Children[types.GetSegmentPath(&detail.Location[i])] = types.YChild{"Location", &detail.Location[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detail.EntityData)
}

// Oper_Platform_Detail_Location
type Oper_Platform_Detail_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    Location interface{}

    
    DetailData Oper_Platform_Detail_Location_DetailData
}

func (location *Oper_Platform_Detail_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "detail"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["detail-data"] = types.YChild{"DetailData", &location.DetailData}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Oper_Platform_Detail_Location_DetailData
type Oper_Platform_Detail_Location_DetailData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Pid interface{}

    // The type is string.
    Description interface{}

    // The type is string.
    VidSn interface{}

    // The type is string.
    HwState interface{}

    // The type is string.
    SwState interface{}

    // The type is string.
    ConfigWordy interface{}

    // The type is string.
    HwVer interface{}

    // The type is string.
    LastEvent interface{}

    // The type is string.
    LastEvReasonStr interface{}
}

func (detailData *Oper_Platform_Detail_Location_DetailData) GetEntityData() *types.CommonEntityData {
    detailData.EntityData.YFilter = detailData.YFilter
    detailData.EntityData.YangName = "detail-data"
    detailData.EntityData.BundleName = "cisco_ios_xr"
    detailData.EntityData.ParentYangName = "location"
    detailData.EntityData.SegmentPath = "detail-data"
    detailData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailData.EntityData.Children = make(map[string]types.YChild)
    detailData.EntityData.Leafs = make(map[string]types.YLeaf)
    detailData.EntityData.Leafs["pid"] = types.YLeaf{"Pid", detailData.Pid}
    detailData.EntityData.Leafs["description"] = types.YLeaf{"Description", detailData.Description}
    detailData.EntityData.Leafs["vid_sn"] = types.YLeaf{"VidSn", detailData.VidSn}
    detailData.EntityData.Leafs["hw_state"] = types.YLeaf{"HwState", detailData.HwState}
    detailData.EntityData.Leafs["sw_state"] = types.YLeaf{"SwState", detailData.SwState}
    detailData.EntityData.Leafs["config_wordy"] = types.YLeaf{"ConfigWordy", detailData.ConfigWordy}
    detailData.EntityData.Leafs["hw_ver"] = types.YLeaf{"HwVer", detailData.HwVer}
    detailData.EntityData.Leafs["last_event"] = types.YLeaf{"LastEvent", detailData.LastEvent}
    detailData.EntityData.Leafs["last_ev_reason_str"] = types.YLeaf{"LastEvReasonStr", detailData.LastEvReasonStr}
    return &(detailData.EntityData)
}

// Oper_Platform_Slices
type Oper_Platform_Slices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Platform_Slices_Location.
    Location []Oper_Platform_Slices_Location
}

func (slices *Oper_Platform_Slices) GetEntityData() *types.CommonEntityData {
    slices.EntityData.YFilter = slices.YFilter
    slices.EntityData.YangName = "slices"
    slices.EntityData.BundleName = "cisco_ios_xr"
    slices.EntityData.ParentYangName = "platform"
    slices.EntityData.SegmentPath = "slices"
    slices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slices.EntityData.Children = make(map[string]types.YChild)
    slices.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range slices.Location {
        slices.EntityData.Children[types.GetSegmentPath(&slices.Location[i])] = types.YChild{"Location", &slices.Location[i]}
    }
    slices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slices.EntityData)
}

// Oper_Platform_Slices_Location
type Oper_Platform_Slices_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    Location interface{}

    // The type is slice of Oper_Platform_Slices_Location_SliceValues.
    SliceValues []Oper_Platform_Slices_Location_SliceValues
}

func (location *Oper_Platform_Slices_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "slices"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["slice_values"] = types.YChild{"SliceValues", nil}
    for i := range location.SliceValues {
        location.EntityData.Children[types.GetSegmentPath(&location.SliceValues[i])] = types.YChild{"SliceValues", &location.SliceValues[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Oper_Platform_Slices_Location_SliceValues
type Oper_Platform_Slices_Location_SliceValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    SliceIdx interface{}

    
    Slice Oper_Platform_Slices_Location_SliceValues_Slice
}

func (sliceValues *Oper_Platform_Slices_Location_SliceValues) GetEntityData() *types.CommonEntityData {
    sliceValues.EntityData.YFilter = sliceValues.YFilter
    sliceValues.EntityData.YangName = "slice_values"
    sliceValues.EntityData.BundleName = "cisco_ios_xr"
    sliceValues.EntityData.ParentYangName = "location"
    sliceValues.EntityData.SegmentPath = "slice_values" + "[slice_idx='" + fmt.Sprintf("%v", sliceValues.SliceIdx) + "']"
    sliceValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceValues.EntityData.Children = make(map[string]types.YChild)
    sliceValues.EntityData.Children["slice"] = types.YChild{"Slice", &sliceValues.Slice}
    sliceValues.EntityData.Leafs = make(map[string]types.YLeaf)
    sliceValues.EntityData.Leafs["slice_idx"] = types.YLeaf{"SliceIdx", sliceValues.SliceIdx}
    return &(sliceValues.EntityData)
}

// Oper_Platform_Slices_Location_SliceValues_Slice
type Oper_Platform_Slices_Location_SliceValues_Slice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    SliceNum interface{}

    // The type is string.
    AdminState interface{}

    // The type is string.
    OperState interface{}
}

func (slice *Oper_Platform_Slices_Location_SliceValues_Slice) GetEntityData() *types.CommonEntityData {
    slice.EntityData.YFilter = slice.YFilter
    slice.EntityData.YangName = "slice"
    slice.EntityData.BundleName = "cisco_ios_xr"
    slice.EntityData.ParentYangName = "slice_values"
    slice.EntityData.SegmentPath = "slice"
    slice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slice.EntityData.Children = make(map[string]types.YChild)
    slice.EntityData.Leafs = make(map[string]types.YLeaf)
    slice.EntityData.Leafs["slice_num"] = types.YLeaf{"SliceNum", slice.SliceNum}
    slice.EntityData.Leafs["admin_state"] = types.YLeaf{"AdminState", slice.AdminState}
    slice.EntityData.Leafs["oper_state"] = types.YLeaf{"OperState", slice.OperState}
    return &(slice.EntityData)
}

// Oper_Chassis
type Oper_Chassis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Brief Oper_Chassis_Brief
}

func (chassis *Oper_Chassis) GetEntityData() *types.CommonEntityData {
    chassis.EntityData.YFilter = chassis.YFilter
    chassis.EntityData.YangName = "chassis"
    chassis.EntityData.BundleName = "cisco_ios_xr"
    chassis.EntityData.ParentYangName = "oper"
    chassis.EntityData.SegmentPath = "chassis"
    chassis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassis.EntityData.Children = make(map[string]types.YChild)
    chassis.EntityData.Children["brief"] = types.YChild{"Brief", &chassis.Brief}
    chassis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(chassis.EntityData)
}

// Oper_Chassis_Brief
type Oper_Chassis_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Chassis_Brief_ChassisSerial.
    ChassisSerial []Oper_Chassis_Brief_ChassisSerial
}

func (brief *Oper_Chassis_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "chassis"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Children["chassis_serial"] = types.YChild{"ChassisSerial", nil}
    for i := range brief.ChassisSerial {
        brief.EntityData.Children[types.GetSegmentPath(&brief.ChassisSerial[i])] = types.YChild{"ChassisSerial", &brief.ChassisSerial[i]}
    }
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(brief.EntityData)
}

// Oper_Chassis_Brief_ChassisSerial
type Oper_Chassis_Brief_ChassisSerial struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    SerialNumber interface{}

    
    BriefData Oper_Chassis_Brief_ChassisSerial_BriefData
}

func (chassisSerial *Oper_Chassis_Brief_ChassisSerial) GetEntityData() *types.CommonEntityData {
    chassisSerial.EntityData.YFilter = chassisSerial.YFilter
    chassisSerial.EntityData.YangName = "chassis_serial"
    chassisSerial.EntityData.BundleName = "cisco_ios_xr"
    chassisSerial.EntityData.ParentYangName = "brief"
    chassisSerial.EntityData.SegmentPath = "chassis_serial" + "[serial_number='" + fmt.Sprintf("%v", chassisSerial.SerialNumber) + "']"
    chassisSerial.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisSerial.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisSerial.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisSerial.EntityData.Children = make(map[string]types.YChild)
    chassisSerial.EntityData.Children["brief-data"] = types.YChild{"BriefData", &chassisSerial.BriefData}
    chassisSerial.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisSerial.EntityData.Leafs["serial_number"] = types.YLeaf{"SerialNumber", chassisSerial.SerialNumber}
    return &(chassisSerial.EntityData)
}

// Oper_Chassis_Brief_ChassisSerial_BriefData
type Oper_Chassis_Brief_ChassisSerial_BriefData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Racknum interface{}

    // The type is string.
    RackType interface{}

    // The type is string.
    RackState interface{}

    // The type is string.
    DataPlane interface{}

    // The type is string.
    CtrlPlane interface{}
}

func (briefData *Oper_Chassis_Brief_ChassisSerial_BriefData) GetEntityData() *types.CommonEntityData {
    briefData.EntityData.YFilter = briefData.YFilter
    briefData.EntityData.YangName = "brief-data"
    briefData.EntityData.BundleName = "cisco_ios_xr"
    briefData.EntityData.ParentYangName = "chassis_serial"
    briefData.EntityData.SegmentPath = "brief-data"
    briefData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefData.EntityData.Children = make(map[string]types.YChild)
    briefData.EntityData.Leafs = make(map[string]types.YLeaf)
    briefData.EntityData.Leafs["racknum"] = types.YLeaf{"Racknum", briefData.Racknum}
    briefData.EntityData.Leafs["rack_type"] = types.YLeaf{"RackType", briefData.RackType}
    briefData.EntityData.Leafs["rack_state"] = types.YLeaf{"RackState", briefData.RackState}
    briefData.EntityData.Leafs["data_plane"] = types.YLeaf{"DataPlane", briefData.DataPlane}
    briefData.EntityData.Leafs["ctrl_plane"] = types.YLeaf{"CtrlPlane", briefData.CtrlPlane}
    return &(briefData.EntityData)
}

// Oper_Reload
type Oper_Reload struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rack Oper_Reload_Rack
}

func (reload *Oper_Reload) GetEntityData() *types.CommonEntityData {
    reload.EntityData.YFilter = reload.YFilter
    reload.EntityData.YangName = "reload"
    reload.EntityData.BundleName = "cisco_ios_xr"
    reload.EntityData.ParentYangName = "oper"
    reload.EntityData.SegmentPath = "reload"
    reload.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reload.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reload.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reload.EntityData.Children = make(map[string]types.YChild)
    reload.EntityData.Children["rack"] = types.YChild{"Rack", &reload.Rack}
    reload.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reload.EntityData)
}

// Oper_Reload_Rack
type Oper_Reload_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Reload_Rack_Racks.
    Racks []Oper_Reload_Rack_Racks
}

func (rack *Oper_Reload_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "reload"
    rack.EntityData.SegmentPath = "rack"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["racks"] = types.YChild{"Racks", nil}
    for i := range rack.Racks {
        rack.EntityData.Children[types.GetSegmentPath(&rack.Racks[i])] = types.YChild{"Racks", &rack.Racks[i]}
    }
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rack.EntityData)
}

// Oper_Reload_Rack_Racks
type Oper_Reload_Rack_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Rack interface{}
}

func (racks *Oper_Reload_Rack_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "rack"
    racks.EntityData.SegmentPath = "racks" + "[rack='" + fmt.Sprintf("%v", racks.Rack) + "']"
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = make(map[string]types.YChild)
    racks.EntityData.Leafs = make(map[string]types.YLeaf)
    racks.EntityData.Leafs["rack"] = types.YLeaf{"Rack", racks.Rack}
    return &(racks.EntityData)
}

// Oper_RebootHistory
type Oper_RebootHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Card Oper_RebootHistory_Card

    
    AdminVm Oper_RebootHistory_AdminVm

    
    Reverse Oper_RebootHistory_Reverse
}

func (rebootHistory *Oper_RebootHistory) GetEntityData() *types.CommonEntityData {
    rebootHistory.EntityData.YFilter = rebootHistory.YFilter
    rebootHistory.EntityData.YangName = "reboot-history"
    rebootHistory.EntityData.BundleName = "cisco_ios_xr"
    rebootHistory.EntityData.ParentYangName = "oper"
    rebootHistory.EntityData.SegmentPath = "reboot-history"
    rebootHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebootHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebootHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebootHistory.EntityData.Children = make(map[string]types.YChild)
    rebootHistory.EntityData.Children["card"] = types.YChild{"Card", &rebootHistory.Card}
    rebootHistory.EntityData.Children["admin-vm"] = types.YChild{"AdminVm", &rebootHistory.AdminVm}
    rebootHistory.EntityData.Children["reverse"] = types.YChild{"Reverse", &rebootHistory.Reverse}
    rebootHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rebootHistory.EntityData)
}

// Oper_RebootHistory_Card
type Oper_RebootHistory_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_RebootHistory_Card_Location.
    Location []Oper_RebootHistory_Card_Location
}

func (card *Oper_RebootHistory_Card) GetEntityData() *types.CommonEntityData {
    card.EntityData.YFilter = card.YFilter
    card.EntityData.YangName = "card"
    card.EntityData.BundleName = "cisco_ios_xr"
    card.EntityData.ParentYangName = "reboot-history"
    card.EntityData.SegmentPath = "card"
    card.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    card.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    card.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    card.EntityData.Children = make(map[string]types.YChild)
    card.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range card.Location {
        card.EntityData.Children[types.GetSegmentPath(&card.Location[i])] = types.YChild{"Location", &card.Location[i]}
    }
    card.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(card.EntityData)
}

// Oper_RebootHistory_Card_Location
type Oper_RebootHistory_Card_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    Location interface{}

    // The type is slice of Oper_RebootHistory_Card_Location_Events.
    Events []Oper_RebootHistory_Card_Location_Events
}

func (location *Oper_RebootHistory_Card_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "card"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["events"] = types.YChild{"Events", nil}
    for i := range location.Events {
        location.EntityData.Children[types.GetSegmentPath(&location.Events[i])] = types.YChild{"Events", &location.Events[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Oper_RebootHistory_Card_Location_Events
type Oper_RebootHistory_Card_Location_Events struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    EventIdx interface{}

    
    Event Oper_RebootHistory_Card_Location_Events_Event
}

func (events *Oper_RebootHistory_Card_Location_Events) GetEntityData() *types.CommonEntityData {
    events.EntityData.YFilter = events.YFilter
    events.EntityData.YangName = "events"
    events.EntityData.BundleName = "cisco_ios_xr"
    events.EntityData.ParentYangName = "location"
    events.EntityData.SegmentPath = "events" + "[event_idx='" + fmt.Sprintf("%v", events.EventIdx) + "']"
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = make(map[string]types.YChild)
    events.EntityData.Children["event"] = types.YChild{"Event", &events.Event}
    events.EntityData.Leafs = make(map[string]types.YLeaf)
    events.EntityData.Leafs["event_idx"] = types.YLeaf{"EventIdx", events.EventIdx}
    return &(events.EntityData)
}

// Oper_RebootHistory_Card_Location_Events_Event
type Oper_RebootHistory_Card_Location_Events_Event struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Timestamp interface{}

    // The type is interface{} with range: 0..4294967295.
    ReasonCode interface{}

    // The type is string.
    Reason interface{}

    // The type is string.
    SrcLoc interface{}

    // The type is string.
    SrcName interface{}
}

func (event *Oper_RebootHistory_Card_Location_Events_Event) GetEntityData() *types.CommonEntityData {
    event.EntityData.YFilter = event.YFilter
    event.EntityData.YangName = "event"
    event.EntityData.BundleName = "cisco_ios_xr"
    event.EntityData.ParentYangName = "events"
    event.EntityData.SegmentPath = "event"
    event.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    event.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    event.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    event.EntityData.Children = make(map[string]types.YChild)
    event.EntityData.Leafs = make(map[string]types.YLeaf)
    event.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", event.Timestamp}
    event.EntityData.Leafs["reason_code"] = types.YLeaf{"ReasonCode", event.ReasonCode}
    event.EntityData.Leafs["reason"] = types.YLeaf{"Reason", event.Reason}
    event.EntityData.Leafs["src_loc"] = types.YLeaf{"SrcLoc", event.SrcLoc}
    event.EntityData.Leafs["src_name"] = types.YLeaf{"SrcName", event.SrcName}
    return &(event.EntityData)
}

// Oper_RebootHistory_AdminVm
type Oper_RebootHistory_AdminVm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_RebootHistory_AdminVm_Location.
    Location []Oper_RebootHistory_AdminVm_Location
}

func (adminVm *Oper_RebootHistory_AdminVm) GetEntityData() *types.CommonEntityData {
    adminVm.EntityData.YFilter = adminVm.YFilter
    adminVm.EntityData.YangName = "admin-vm"
    adminVm.EntityData.BundleName = "cisco_ios_xr"
    adminVm.EntityData.ParentYangName = "reboot-history"
    adminVm.EntityData.SegmentPath = "admin-vm"
    adminVm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminVm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminVm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminVm.EntityData.Children = make(map[string]types.YChild)
    adminVm.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range adminVm.Location {
        adminVm.EntityData.Children[types.GetSegmentPath(&adminVm.Location[i])] = types.YChild{"Location", &adminVm.Location[i]}
    }
    adminVm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adminVm.EntityData)
}

// Oper_RebootHistory_AdminVm_Location
type Oper_RebootHistory_AdminVm_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    Location interface{}

    // The type is slice of Oper_RebootHistory_AdminVm_Location_Events.
    Events []Oper_RebootHistory_AdminVm_Location_Events
}

func (location *Oper_RebootHistory_AdminVm_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "admin-vm"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["events"] = types.YChild{"Events", nil}
    for i := range location.Events {
        location.EntityData.Children[types.GetSegmentPath(&location.Events[i])] = types.YChild{"Events", &location.Events[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Oper_RebootHistory_AdminVm_Location_Events
type Oper_RebootHistory_AdminVm_Location_Events struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    EventIdx interface{}

    
    Event Oper_RebootHistory_AdminVm_Location_Events_Event
}

func (events *Oper_RebootHistory_AdminVm_Location_Events) GetEntityData() *types.CommonEntityData {
    events.EntityData.YFilter = events.YFilter
    events.EntityData.YangName = "events"
    events.EntityData.BundleName = "cisco_ios_xr"
    events.EntityData.ParentYangName = "location"
    events.EntityData.SegmentPath = "events" + "[event_idx='" + fmt.Sprintf("%v", events.EventIdx) + "']"
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = make(map[string]types.YChild)
    events.EntityData.Children["event"] = types.YChild{"Event", &events.Event}
    events.EntityData.Leafs = make(map[string]types.YLeaf)
    events.EntityData.Leafs["event_idx"] = types.YLeaf{"EventIdx", events.EventIdx}
    return &(events.EntityData)
}

// Oper_RebootHistory_AdminVm_Location_Events_Event
type Oper_RebootHistory_AdminVm_Location_Events_Event struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Timestamp interface{}

    // The type is interface{} with range: 0..4294967295.
    ReasonCode interface{}

    // The type is string.
    Reason interface{}

    // The type is string.
    SrcLoc interface{}

    // The type is string.
    SrcName interface{}
}

func (event *Oper_RebootHistory_AdminVm_Location_Events_Event) GetEntityData() *types.CommonEntityData {
    event.EntityData.YFilter = event.YFilter
    event.EntityData.YangName = "event"
    event.EntityData.BundleName = "cisco_ios_xr"
    event.EntityData.ParentYangName = "events"
    event.EntityData.SegmentPath = "event"
    event.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    event.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    event.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    event.EntityData.Children = make(map[string]types.YChild)
    event.EntityData.Leafs = make(map[string]types.YLeaf)
    event.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", event.Timestamp}
    event.EntityData.Leafs["reason_code"] = types.YLeaf{"ReasonCode", event.ReasonCode}
    event.EntityData.Leafs["reason"] = types.YLeaf{"Reason", event.Reason}
    event.EntityData.Leafs["src_loc"] = types.YLeaf{"SrcLoc", event.SrcLoc}
    event.EntityData.Leafs["src_name"] = types.YLeaf{"SrcName", event.SrcName}
    return &(event.EntityData)
}

// Oper_RebootHistory_Reverse
type Oper_RebootHistory_Reverse struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Card Oper_RebootHistory_Reverse_Card

    
    AdminVm Oper_RebootHistory_Reverse_AdminVm
}

func (reverse *Oper_RebootHistory_Reverse) GetEntityData() *types.CommonEntityData {
    reverse.EntityData.YFilter = reverse.YFilter
    reverse.EntityData.YangName = "reverse"
    reverse.EntityData.BundleName = "cisco_ios_xr"
    reverse.EntityData.ParentYangName = "reboot-history"
    reverse.EntityData.SegmentPath = "reverse"
    reverse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reverse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reverse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reverse.EntityData.Children = make(map[string]types.YChild)
    reverse.EntityData.Children["card"] = types.YChild{"Card", &reverse.Card}
    reverse.EntityData.Children["admin-vm"] = types.YChild{"AdminVm", &reverse.AdminVm}
    reverse.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reverse.EntityData)
}

// Oper_RebootHistory_Reverse_Card
type Oper_RebootHistory_Reverse_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_RebootHistory_Reverse_Card_Location.
    Location []Oper_RebootHistory_Reverse_Card_Location
}

func (card *Oper_RebootHistory_Reverse_Card) GetEntityData() *types.CommonEntityData {
    card.EntityData.YFilter = card.YFilter
    card.EntityData.YangName = "card"
    card.EntityData.BundleName = "cisco_ios_xr"
    card.EntityData.ParentYangName = "reverse"
    card.EntityData.SegmentPath = "card"
    card.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    card.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    card.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    card.EntityData.Children = make(map[string]types.YChild)
    card.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range card.Location {
        card.EntityData.Children[types.GetSegmentPath(&card.Location[i])] = types.YChild{"Location", &card.Location[i]}
    }
    card.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(card.EntityData)
}

// Oper_RebootHistory_Reverse_Card_Location
type Oper_RebootHistory_Reverse_Card_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    Location interface{}

    // The type is slice of Oper_RebootHistory_Reverse_Card_Location_Events.
    Events []Oper_RebootHistory_Reverse_Card_Location_Events
}

func (location *Oper_RebootHistory_Reverse_Card_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "card"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["events"] = types.YChild{"Events", nil}
    for i := range location.Events {
        location.EntityData.Children[types.GetSegmentPath(&location.Events[i])] = types.YChild{"Events", &location.Events[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Oper_RebootHistory_Reverse_Card_Location_Events
type Oper_RebootHistory_Reverse_Card_Location_Events struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    EventIdx interface{}

    
    Event Oper_RebootHistory_Reverse_Card_Location_Events_Event
}

func (events *Oper_RebootHistory_Reverse_Card_Location_Events) GetEntityData() *types.CommonEntityData {
    events.EntityData.YFilter = events.YFilter
    events.EntityData.YangName = "events"
    events.EntityData.BundleName = "cisco_ios_xr"
    events.EntityData.ParentYangName = "location"
    events.EntityData.SegmentPath = "events" + "[event_idx='" + fmt.Sprintf("%v", events.EventIdx) + "']"
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = make(map[string]types.YChild)
    events.EntityData.Children["event"] = types.YChild{"Event", &events.Event}
    events.EntityData.Leafs = make(map[string]types.YLeaf)
    events.EntityData.Leafs["event_idx"] = types.YLeaf{"EventIdx", events.EventIdx}
    return &(events.EntityData)
}

// Oper_RebootHistory_Reverse_Card_Location_Events_Event
type Oper_RebootHistory_Reverse_Card_Location_Events_Event struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Timestamp interface{}

    // The type is interface{} with range: 0..4294967295.
    ReasonCode interface{}

    // The type is string.
    Reason interface{}

    // The type is string.
    SrcLoc interface{}

    // The type is string.
    SrcName interface{}
}

func (event *Oper_RebootHistory_Reverse_Card_Location_Events_Event) GetEntityData() *types.CommonEntityData {
    event.EntityData.YFilter = event.YFilter
    event.EntityData.YangName = "event"
    event.EntityData.BundleName = "cisco_ios_xr"
    event.EntityData.ParentYangName = "events"
    event.EntityData.SegmentPath = "event"
    event.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    event.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    event.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    event.EntityData.Children = make(map[string]types.YChild)
    event.EntityData.Leafs = make(map[string]types.YLeaf)
    event.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", event.Timestamp}
    event.EntityData.Leafs["reason_code"] = types.YLeaf{"ReasonCode", event.ReasonCode}
    event.EntityData.Leafs["reason"] = types.YLeaf{"Reason", event.Reason}
    event.EntityData.Leafs["src_loc"] = types.YLeaf{"SrcLoc", event.SrcLoc}
    event.EntityData.Leafs["src_name"] = types.YLeaf{"SrcName", event.SrcName}
    return &(event.EntityData)
}

// Oper_RebootHistory_Reverse_AdminVm
type Oper_RebootHistory_Reverse_AdminVm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_RebootHistory_Reverse_AdminVm_Location.
    Location []Oper_RebootHistory_Reverse_AdminVm_Location
}

func (adminVm *Oper_RebootHistory_Reverse_AdminVm) GetEntityData() *types.CommonEntityData {
    adminVm.EntityData.YFilter = adminVm.YFilter
    adminVm.EntityData.YangName = "admin-vm"
    adminVm.EntityData.BundleName = "cisco_ios_xr"
    adminVm.EntityData.ParentYangName = "reverse"
    adminVm.EntityData.SegmentPath = "admin-vm"
    adminVm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminVm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminVm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminVm.EntityData.Children = make(map[string]types.YChild)
    adminVm.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range adminVm.Location {
        adminVm.EntityData.Children[types.GetSegmentPath(&adminVm.Location[i])] = types.YChild{"Location", &adminVm.Location[i]}
    }
    adminVm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adminVm.EntityData)
}

// Oper_RebootHistory_Reverse_AdminVm_Location
type Oper_RebootHistory_Reverse_AdminVm_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'((([fF][0-3])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[1-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    Location interface{}

    // The type is slice of Oper_RebootHistory_Reverse_AdminVm_Location_Events.
    Events []Oper_RebootHistory_Reverse_AdminVm_Location_Events
}

func (location *Oper_RebootHistory_Reverse_AdminVm_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "admin-vm"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["events"] = types.YChild{"Events", nil}
    for i := range location.Events {
        location.EntityData.Children[types.GetSegmentPath(&location.Events[i])] = types.YChild{"Events", &location.Events[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Oper_RebootHistory_Reverse_AdminVm_Location_Events
type Oper_RebootHistory_Reverse_AdminVm_Location_Events struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    EventIdx interface{}

    
    Event Oper_RebootHistory_Reverse_AdminVm_Location_Events_Event
}

func (events *Oper_RebootHistory_Reverse_AdminVm_Location_Events) GetEntityData() *types.CommonEntityData {
    events.EntityData.YFilter = events.YFilter
    events.EntityData.YangName = "events"
    events.EntityData.BundleName = "cisco_ios_xr"
    events.EntityData.ParentYangName = "location"
    events.EntityData.SegmentPath = "events" + "[event_idx='" + fmt.Sprintf("%v", events.EventIdx) + "']"
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = make(map[string]types.YChild)
    events.EntityData.Children["event"] = types.YChild{"Event", &events.Event}
    events.EntityData.Leafs = make(map[string]types.YLeaf)
    events.EntityData.Leafs["event_idx"] = types.YLeaf{"EventIdx", events.EventIdx}
    return &(events.EntityData)
}

// Oper_RebootHistory_Reverse_AdminVm_Location_Events_Event
type Oper_RebootHistory_Reverse_AdminVm_Location_Events_Event struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Timestamp interface{}

    // The type is interface{} with range: 0..4294967295.
    ReasonCode interface{}

    // The type is string.
    Reason interface{}

    // The type is string.
    SrcLoc interface{}

    // The type is string.
    SrcName interface{}
}

func (event *Oper_RebootHistory_Reverse_AdminVm_Location_Events_Event) GetEntityData() *types.CommonEntityData {
    event.EntityData.YFilter = event.YFilter
    event.EntityData.YangName = "event"
    event.EntityData.BundleName = "cisco_ios_xr"
    event.EntityData.ParentYangName = "events"
    event.EntityData.SegmentPath = "event"
    event.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    event.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    event.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    event.EntityData.Children = make(map[string]types.YChild)
    event.EntityData.Leafs = make(map[string]types.YLeaf)
    event.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", event.Timestamp}
    event.EntityData.Leafs["reason_code"] = types.YLeaf{"ReasonCode", event.ReasonCode}
    event.EntityData.Leafs["reason"] = types.YLeaf{"Reason", event.Reason}
    event.EntityData.Leafs["src_loc"] = types.YLeaf{"SrcLoc", event.SrcLoc}
    event.EntityData.Leafs["src_name"] = types.YLeaf{"SrcName", event.SrcName}
    return &(event.EntityData)
}

// Oper_Interface
type Oper_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Ifname interface{}

    
    InterfaceData Oper_Interface_InterfaceData
}

func (self *Oper_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "oper"
    self.EntityData.SegmentPath = "interface" + "[ifname='" + fmt.Sprintf("%v", self.Ifname) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["interface-data"] = types.YChild{"InterfaceData", &self.InterfaceData}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["ifname"] = types.YLeaf{"Ifname", self.Ifname}
    return &(self.EntityData)
}

// Oper_Interface_InterfaceData
type Oper_Interface_InterfaceData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Mac interface{}

    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // The type is string.
    Flagstr interface{}

    // The type is string.
    PortStatus interface{}

    // The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // The type is interface{} with range: 0..4294967295.
    RxPak interface{}

    // The type is interface{} with range: 0..4294967295.
    RxErrors interface{}

    // The type is interface{} with range: 0..4294967295.
    RxDropped interface{}

    // The type is interface{} with range: 0..4294967295.
    RxOverruns interface{}

    // The type is interface{} with range: 0..4294967295.
    RxFrame interface{}

    // The type is interface{} with range: 0..4294967295.
    TxPak interface{}

    // The type is interface{} with range: 0..4294967295.
    TxErrors interface{}

    // The type is interface{} with range: 0..4294967295.
    TxDropped interface{}

    // The type is interface{} with range: 0..4294967295.
    TxOverruns interface{}

    // The type is interface{} with range: 0..4294967295.
    TxCarrier interface{}

    // The type is interface{} with range: 0..4294967295.
    Collisions interface{}

    // The type is interface{} with range: 0..4294967295.
    TxQueuelen interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    RxBytes interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    TxBytes interface{}

    // The type is interface{} with range: 0..4294967295.
    IntfNum interface{}
}

func (interfaceData *Oper_Interface_InterfaceData) GetEntityData() *types.CommonEntityData {
    interfaceData.EntityData.YFilter = interfaceData.YFilter
    interfaceData.EntityData.YangName = "interface-data"
    interfaceData.EntityData.BundleName = "cisco_ios_xr"
    interfaceData.EntityData.ParentYangName = "interface"
    interfaceData.EntityData.SegmentPath = "interface-data"
    interfaceData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceData.EntityData.Children = make(map[string]types.YChild)
    interfaceData.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceData.EntityData.Leafs["mac"] = types.YLeaf{"Mac", interfaceData.Mac}
    interfaceData.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", interfaceData.Ipv4}
    interfaceData.EntityData.Leafs["flagstr"] = types.YLeaf{"Flagstr", interfaceData.Flagstr}
    interfaceData.EntityData.Leafs["port_status"] = types.YLeaf{"PortStatus", interfaceData.PortStatus}
    interfaceData.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", interfaceData.Mtu}
    interfaceData.EntityData.Leafs["metric"] = types.YLeaf{"Metric", interfaceData.Metric}
    interfaceData.EntityData.Leafs["rx_pak"] = types.YLeaf{"RxPak", interfaceData.RxPak}
    interfaceData.EntityData.Leafs["rx_errors"] = types.YLeaf{"RxErrors", interfaceData.RxErrors}
    interfaceData.EntityData.Leafs["rx_dropped"] = types.YLeaf{"RxDropped", interfaceData.RxDropped}
    interfaceData.EntityData.Leafs["rx_overruns"] = types.YLeaf{"RxOverruns", interfaceData.RxOverruns}
    interfaceData.EntityData.Leafs["rx_frame"] = types.YLeaf{"RxFrame", interfaceData.RxFrame}
    interfaceData.EntityData.Leafs["tx_pak"] = types.YLeaf{"TxPak", interfaceData.TxPak}
    interfaceData.EntityData.Leafs["tx_errors"] = types.YLeaf{"TxErrors", interfaceData.TxErrors}
    interfaceData.EntityData.Leafs["tx_dropped"] = types.YLeaf{"TxDropped", interfaceData.TxDropped}
    interfaceData.EntityData.Leafs["tx_overruns"] = types.YLeaf{"TxOverruns", interfaceData.TxOverruns}
    interfaceData.EntityData.Leafs["tx_carrier"] = types.YLeaf{"TxCarrier", interfaceData.TxCarrier}
    interfaceData.EntityData.Leafs["collisions"] = types.YLeaf{"Collisions", interfaceData.Collisions}
    interfaceData.EntityData.Leafs["tx_queuelen"] = types.YLeaf{"TxQueuelen", interfaceData.TxQueuelen}
    interfaceData.EntityData.Leafs["rx_bytes"] = types.YLeaf{"RxBytes", interfaceData.RxBytes}
    interfaceData.EntityData.Leafs["tx_bytes"] = types.YLeaf{"TxBytes", interfaceData.TxBytes}
    interfaceData.EntityData.Leafs["intf_num"] = types.YLeaf{"IntfNum", interfaceData.IntfNum}
    return &(interfaceData.EntityData)
}

// Oper_ReloadVm
type Oper_ReloadVm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_ReloadVm_Location.
    Location []Oper_ReloadVm_Location
}

func (reloadVm *Oper_ReloadVm) GetEntityData() *types.CommonEntityData {
    reloadVm.EntityData.YFilter = reloadVm.YFilter
    reloadVm.EntityData.YangName = "reload_vm"
    reloadVm.EntityData.BundleName = "cisco_ios_xr"
    reloadVm.EntityData.ParentYangName = "oper"
    reloadVm.EntityData.SegmentPath = "reload_vm"
    reloadVm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reloadVm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reloadVm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reloadVm.EntityData.Children = make(map[string]types.YChild)
    reloadVm.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range reloadVm.Location {
        reloadVm.EntityData.Children[types.GetSegmentPath(&reloadVm.Location[i])] = types.YChild{"Location", &reloadVm.Location[i]}
    }
    reloadVm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reloadVm.EntityData)
}

// Oper_ReloadVm_Location
type Oper_ReloadVm_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Location interface{}
}

func (location *Oper_ReloadVm_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "reload_vm"
    location.EntityData.SegmentPath = "location" + "[location='" + fmt.Sprintf("%v", location.Location) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location"] = types.YLeaf{"Location", location.Location}
    return &(location.EntityData)
}

// Oper_Macpool
type Oper_Macpool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Brief Oper_Macpool_Brief
}

func (macpool *Oper_Macpool) GetEntityData() *types.CommonEntityData {
    macpool.EntityData.YFilter = macpool.YFilter
    macpool.EntityData.YangName = "macpool"
    macpool.EntityData.BundleName = "cisco_ios_xr"
    macpool.EntityData.ParentYangName = "oper"
    macpool.EntityData.SegmentPath = "macpool"
    macpool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macpool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macpool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macpool.EntityData.Children = make(map[string]types.YChild)
    macpool.EntityData.Children["brief"] = types.YChild{"Brief", &macpool.Brief}
    macpool.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macpool.EntityData)
}

// Oper_Macpool_Brief
type Oper_Macpool_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Macpool_Brief_Rack.
    Rack []Oper_Macpool_Brief_Rack
}

func (brief *Oper_Macpool_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "macpool"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Children["rack"] = types.YChild{"Rack", nil}
    for i := range brief.Rack {
        brief.EntityData.Children[types.GetSegmentPath(&brief.Rack[i])] = types.YChild{"Rack", &brief.Rack[i]}
    }
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(brief.EntityData)
}

// Oper_Macpool_Brief_Rack
type Oper_Macpool_Brief_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    SerialNumber interface{}

    
    BriefData Oper_Macpool_Brief_Rack_BriefData
}

func (rack *Oper_Macpool_Brief_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "brief"
    rack.EntityData.SegmentPath = "rack" + "[serial_number='" + fmt.Sprintf("%v", rack.SerialNumber) + "']"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["brief-data"] = types.YChild{"BriefData", &rack.BriefData}
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    rack.EntityData.Leafs["serial_number"] = types.YLeaf{"SerialNumber", rack.SerialNumber}
    return &(rack.EntityData)
}

// Oper_Macpool_Brief_Rack_BriefData
type Oper_Macpool_Brief_Rack_BriefData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Racknum interface{}

    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacBase interface{}

    // The type is interface{} with range: 0..4294967295.
    MacCount interface{}

    // The type is bool.
    IsSelected interface{}

    // The type is interface{} with range: 0..4294967295.
    AllocatedCount interface{}
}

func (briefData *Oper_Macpool_Brief_Rack_BriefData) GetEntityData() *types.CommonEntityData {
    briefData.EntityData.YFilter = briefData.YFilter
    briefData.EntityData.YangName = "brief-data"
    briefData.EntityData.BundleName = "cisco_ios_xr"
    briefData.EntityData.ParentYangName = "rack"
    briefData.EntityData.SegmentPath = "brief-data"
    briefData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefData.EntityData.Children = make(map[string]types.YChild)
    briefData.EntityData.Leafs = make(map[string]types.YLeaf)
    briefData.EntityData.Leafs["racknum"] = types.YLeaf{"Racknum", briefData.Racknum}
    briefData.EntityData.Leafs["mac_base"] = types.YLeaf{"MacBase", briefData.MacBase}
    briefData.EntityData.Leafs["mac_count"] = types.YLeaf{"MacCount", briefData.MacCount}
    briefData.EntityData.Leafs["is_selected"] = types.YLeaf{"IsSelected", briefData.IsSelected}
    briefData.EntityData.Leafs["allocated_count"] = types.YLeaf{"AllocatedCount", briefData.AllocatedCount}
    return &(briefData.EntityData)
}

// Config
type Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Chassis Config_Chassis

    
    Interface_ Config_Interface

    
    Domain Config_Domain

    
    VirtualMacaddrRange Config_VirtualMacaddrRange
}

func (config *Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-sm"
    config.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-sm:config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["chassis"] = types.YChild{"Chassis", &config.Chassis}
    config.EntityData.Children["interface"] = types.YChild{"Interface_", &config.Interface_}
    config.EntityData.Children["domain"] = types.YChild{"Domain", &config.Domain}
    config.EntityData.Children["virtual-macaddr-range"] = types.YChild{"VirtualMacaddrRange", &config.VirtualMacaddrRange}
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(config.EntityData)
}

// Config_Chassis
type Config_Chassis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Chassis_Serial.
    Serial []Config_Chassis_Serial
}

func (chassis *Config_Chassis) GetEntityData() *types.CommonEntityData {
    chassis.EntityData.YFilter = chassis.YFilter
    chassis.EntityData.YangName = "chassis"
    chassis.EntityData.BundleName = "cisco_ios_xr"
    chassis.EntityData.ParentYangName = "config"
    chassis.EntityData.SegmentPath = "chassis"
    chassis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassis.EntityData.Children = make(map[string]types.YChild)
    chassis.EntityData.Children["serial"] = types.YChild{"Serial", nil}
    for i := range chassis.Serial {
        chassis.EntityData.Children[types.GetSegmentPath(&chassis.Serial[i])] = types.YChild{"Serial", &chassis.Serial[i]}
    }
    chassis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(chassis.EntityData)
}

// Config_Chassis_Serial
type Config_Chassis_Serial struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. This attribute is mandatory.
    Serial interface{}

    // The type is string with pattern: b'[fF][0-7]|[0-9]|[1][0-5]|[2][4][0-7]'.
    // This attribute is mandatory.
    Rack interface{}
}

func (serial *Config_Chassis_Serial) GetEntityData() *types.CommonEntityData {
    serial.EntityData.YFilter = serial.YFilter
    serial.EntityData.YangName = "serial"
    serial.EntityData.BundleName = "cisco_ios_xr"
    serial.EntityData.ParentYangName = "chassis"
    serial.EntityData.SegmentPath = "serial" + "[serial='" + fmt.Sprintf("%v", serial.Serial) + "']"
    serial.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serial.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serial.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serial.EntityData.Children = make(map[string]types.YChild)
    serial.EntityData.Leafs = make(map[string]types.YLeaf)
    serial.EntityData.Leafs["serial"] = types.YLeaf{"Serial", serial.Serial}
    serial.EntityData.Leafs["rack"] = types.YLeaf{"Rack", serial.Rack}
    return &(serial.EntityData)
}

// Config_Interface
type Config_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Mgmteth Config_Interface_MgmtEth
}

func (self *Config_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "config"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["MgmtEth"] = types.YChild{"Mgmteth", &self.Mgmteth}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Config_Interface_MgmtEth
type Config_Interface_MgmtEth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Interface_MgmtEth_Locport.
    Locport []Config_Interface_MgmtEth_Locport
}

func (mgmtEth *Config_Interface_MgmtEth) GetEntityData() *types.CommonEntityData {
    mgmtEth.EntityData.YFilter = mgmtEth.YFilter
    mgmtEth.EntityData.YangName = "MgmtEth"
    mgmtEth.EntityData.BundleName = "cisco_ios_xr"
    mgmtEth.EntityData.ParentYangName = "interface"
    mgmtEth.EntityData.SegmentPath = "MgmtEth"
    mgmtEth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mgmtEth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mgmtEth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mgmtEth.EntityData.Children = make(map[string]types.YChild)
    mgmtEth.EntityData.Children["locport"] = types.YChild{"Locport", nil}
    for i := range mgmtEth.Locport {
        mgmtEth.EntityData.Children[types.GetSegmentPath(&mgmtEth.Locport[i])] = types.YChild{"Locport", &mgmtEth.Locport[i]}
    }
    mgmtEth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mgmtEth.EntityData)
}

// Config_Interface_MgmtEth_Locport
type Config_Interface_MgmtEth_Locport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'[fF][0-7]|[0-9]|[1][0-5]'.
    Rack interface{}

    // This attribute is a key. The type is string with pattern:
    // b'[Rr][Pp][0-1]|[Rr][Ss][Pp][0-1]|[Ss][Cc][0-1]'.
    Slot interface{}

    // This attribute is a key. The type is interface{} with range: 0..None.
    Intf interface{}

    // This attribute is a key. The type is interface{} with range: 0..None.
    Port interface{}

    // The type is interface{}.
    Shutdown interface{}

    // The type is interface{} with range: 48..9000.
    Mtu interface{}

    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DefaultGw interface{}

    
    Ipv4 Config_Interface_MgmtEth_Locport_Ipv4

    
    Arp Config_Interface_MgmtEth_Locport_Arp
}

func (locport *Config_Interface_MgmtEth_Locport) GetEntityData() *types.CommonEntityData {
    locport.EntityData.YFilter = locport.YFilter
    locport.EntityData.YangName = "locport"
    locport.EntityData.BundleName = "cisco_ios_xr"
    locport.EntityData.ParentYangName = "MgmtEth"
    locport.EntityData.SegmentPath = "locport" + "[rack='" + fmt.Sprintf("%v", locport.Rack) + "']" + "[slot='" + fmt.Sprintf("%v", locport.Slot) + "']" + "[intf='" + fmt.Sprintf("%v", locport.Intf) + "']" + "[port='" + fmt.Sprintf("%v", locport.Port) + "']"
    locport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locport.EntityData.Children = make(map[string]types.YChild)
    locport.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &locport.Ipv4}
    locport.EntityData.Children["arp"] = types.YChild{"Arp", &locport.Arp}
    locport.EntityData.Leafs = make(map[string]types.YLeaf)
    locport.EntityData.Leafs["rack"] = types.YLeaf{"Rack", locport.Rack}
    locport.EntityData.Leafs["slot"] = types.YLeaf{"Slot", locport.Slot}
    locport.EntityData.Leafs["intf"] = types.YLeaf{"Intf", locport.Intf}
    locport.EntityData.Leafs["port"] = types.YLeaf{"Port", locport.Port}
    locport.EntityData.Leafs["shutdown"] = types.YLeaf{"Shutdown", locport.Shutdown}
    locport.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", locport.Mtu}
    locport.EntityData.Leafs["default-gw"] = types.YLeaf{"DefaultGw", locport.DefaultGw}
    return &(locport.EntityData)
}

// Config_Interface_MgmtEth_Locport_Ipv4
type Config_Interface_MgmtEth_Locport_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])((
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))|(/([0-9]+)))'.
    Address interface{}
}

func (ipv4 *Config_Interface_MgmtEth_Locport_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "locport"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4.Address}
    return &(ipv4.EntityData)
}

// Config_Interface_MgmtEth_Locport_Arp
type Config_Interface_MgmtEth_Locport_Arp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Interface_MgmtEth_Locport_Arp_Ip.
    Ip []Config_Interface_MgmtEth_Locport_Arp_Ip
}

func (arp *Config_Interface_MgmtEth_Locport_Arp) GetEntityData() *types.CommonEntityData {
    arp.EntityData.YFilter = arp.YFilter
    arp.EntityData.YangName = "arp"
    arp.EntityData.BundleName = "cisco_ios_xr"
    arp.EntityData.ParentYangName = "locport"
    arp.EntityData.SegmentPath = "arp"
    arp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    arp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    arp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    arp.EntityData.Children = make(map[string]types.YChild)
    arp.EntityData.Children["ip"] = types.YChild{"Ip", nil}
    for i := range arp.Ip {
        arp.EntityData.Children[types.GetSegmentPath(&arp.Ip[i])] = types.YChild{"Ip", &arp.Ip[i]}
    }
    arp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(arp.EntityData)
}

// Config_Interface_MgmtEth_Locport_Arp_Ip
type Config_Interface_MgmtEth_Locport_Arp_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Ip interface{}

    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    // This attribute is mandatory.
    Mac interface{}
}

func (ip *Config_Interface_MgmtEth_Locport_Arp_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "arp"
    ip.EntityData.SegmentPath = "ip" + "[ip='" + fmt.Sprintf("%v", ip.Ip) + "']"
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = make(map[string]types.YChild)
    ip.EntityData.Leafs = make(map[string]types.YLeaf)
    ip.EntityData.Leafs["ip"] = types.YLeaf{"Ip", ip.Ip}
    ip.EntityData.Leafs["mac"] = types.YLeaf{"Mac", ip.Mac}
    return &(ip.EntityData)
}

// Config_Domain
type Config_Domain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Domain_Name.
    Name []Config_Domain_Name

    // The type is slice of Config_Domain_NameServer.
    NameServer []Config_Domain_NameServer
}

func (domain *Config_Domain) GetEntityData() *types.CommonEntityData {
    domain.EntityData.YFilter = domain.YFilter
    domain.EntityData.YangName = "domain"
    domain.EntityData.BundleName = "cisco_ios_xr"
    domain.EntityData.ParentYangName = "config"
    domain.EntityData.SegmentPath = "domain"
    domain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domain.EntityData.Children = make(map[string]types.YChild)
    domain.EntityData.Children["name"] = types.YChild{"Name", nil}
    for i := range domain.Name {
        domain.EntityData.Children[types.GetSegmentPath(&domain.Name[i])] = types.YChild{"Name", &domain.Name[i]}
    }
    domain.EntityData.Children["name-server"] = types.YChild{"NameServer", nil}
    for i := range domain.NameServer {
        domain.EntityData.Children[types.GetSegmentPath(&domain.NameServer[i])] = types.YChild{"NameServer", &domain.NameServer[i]}
    }
    domain.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(domain.EntityData)
}

// Config_Domain_Name
type Config_Domain_Name struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Name interface{}
}

func (name *Config_Domain_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "domain"
    name.EntityData.SegmentPath = "name" + "[name='" + fmt.Sprintf("%v", name.Name) + "']"
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = make(map[string]types.YChild)
    name.EntityData.Leafs = make(map[string]types.YLeaf)
    name.EntityData.Leafs["name"] = types.YLeaf{"Name", name.Name}
    return &(name.EntityData)
}

// Config_Domain_NameServer
type Config_Domain_NameServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NameServer interface{}
}

func (nameServer *Config_Domain_NameServer) GetEntityData() *types.CommonEntityData {
    nameServer.EntityData.YFilter = nameServer.YFilter
    nameServer.EntityData.YangName = "name-server"
    nameServer.EntityData.BundleName = "cisco_ios_xr"
    nameServer.EntityData.ParentYangName = "domain"
    nameServer.EntityData.SegmentPath = "name-server" + "[name-server='" + fmt.Sprintf("%v", nameServer.NameServer) + "']"
    nameServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nameServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nameServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nameServer.EntityData.Children = make(map[string]types.YChild)
    nameServer.EntityData.Leafs = make(map[string]types.YLeaf)
    nameServer.EntityData.Leafs["name-server"] = types.YLeaf{"NameServer", nameServer.NameServer}
    return &(nameServer.EntityData)
}

// Config_VirtualMacaddrRange
type Config_VirtualMacaddrRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Base interface{}

    // The type is interface{} with range: 1..65535.
    Count interface{}
}

func (virtualMacaddrRange *Config_VirtualMacaddrRange) GetEntityData() *types.CommonEntityData {
    virtualMacaddrRange.EntityData.YFilter = virtualMacaddrRange.YFilter
    virtualMacaddrRange.EntityData.YangName = "virtual-macaddr-range"
    virtualMacaddrRange.EntityData.BundleName = "cisco_ios_xr"
    virtualMacaddrRange.EntityData.ParentYangName = "config"
    virtualMacaddrRange.EntityData.SegmentPath = "virtual-macaddr-range"
    virtualMacaddrRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualMacaddrRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualMacaddrRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualMacaddrRange.EntityData.Children = make(map[string]types.YChild)
    virtualMacaddrRange.EntityData.Leafs = make(map[string]types.YLeaf)
    virtualMacaddrRange.EntityData.Leafs["base"] = types.YLeaf{"Base", virtualMacaddrRange.Base}
    virtualMacaddrRange.EntityData.Leafs["count"] = types.YLeaf{"Count", virtualMacaddrRange.Count}
    return &(virtualMacaddrRange.EntityData)
}

// Actions
type Actions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (actions *Actions) GetEntityData() *types.CommonEntityData {
    actions.EntityData.YFilter = actions.YFilter
    actions.EntityData.YangName = "actions"
    actions.EntityData.BundleName = "cisco_ios_xr"
    actions.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-sm"
    actions.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-sm:actions"
    actions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actions.EntityData.Children = make(map[string]types.YChild)
    actions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(actions.EntityData)
}

