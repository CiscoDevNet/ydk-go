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
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
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

    // The type is slice of Oper_Interface.
    Interface []*Oper_Interface

    
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

    oper.EntityData.Children = types.NewOrderedMap()
    oper.EntityData.Children.Append("shelf_mgr", types.YChild{"ShelfMgr", &oper.ShelfMgr})
    oper.EntityData.Children.Append("platform", types.YChild{"Platform", &oper.Platform})
    oper.EntityData.Children.Append("chassis", types.YChild{"Chassis", &oper.Chassis})
    oper.EntityData.Children.Append("reload", types.YChild{"Reload", &oper.Reload})
    oper.EntityData.Children.Append("reboot-history", types.YChild{"RebootHistory", &oper.RebootHistory})
    oper.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range oper.Interface {
        oper.EntityData.Children.Append(types.GetSegmentPath(oper.Interface[i]), types.YChild{"Interface", oper.Interface[i]})
    }
    oper.EntityData.Children.Append("reload_vm", types.YChild{"ReloadVm", &oper.ReloadVm})
    oper.EntityData.Children.Append("macpool", types.YChild{"Macpool", &oper.Macpool})
    oper.EntityData.Leafs = types.NewOrderedMap()

    oper.EntityData.YListKeys = []string {}

    return &(oper.EntityData)
}

// Oper_ShelfMgr
type Oper_ShelfMgr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Oper_ShelfMgr_Trace.
    Trace []*Oper_ShelfMgr_Trace
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

    shelfMgr.EntityData.Children = types.NewOrderedMap()
    shelfMgr.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range shelfMgr.Trace {
        shelfMgr.EntityData.Children.Append(types.GetSegmentPath(shelfMgr.Trace[i]), types.YChild{"Trace", shelfMgr.Trace[i]})
    }
    shelfMgr.EntityData.Leafs = types.NewOrderedMap()

    shelfMgr.EntityData.YListKeys = []string {}

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
    Location []*Oper_ShelfMgr_Trace_Location
}

func (trace *Oper_ShelfMgr_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "shelf_mgr"
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

// Oper_ShelfMgr_Trace_Location
type Oper_ShelfMgr_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Oper_ShelfMgr_Trace_Location_AllOptions.
    AllOptions []*Oper_ShelfMgr_Trace_Location_AllOptions
}

func (location *Oper_ShelfMgr_Trace_Location) GetEntityData() *types.CommonEntityData {
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

// Oper_ShelfMgr_Trace_Location_AllOptions
type Oper_ShelfMgr_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Oper_ShelfMgr_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Oper_ShelfMgr_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Oper_ShelfMgr_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
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

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

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

    platform.EntityData.Children = types.NewOrderedMap()
    platform.EntityData.Children.Append("summary", types.YChild{"Summary", &platform.Summary})
    platform.EntityData.Children.Append("detail", types.YChild{"Detail", &platform.Detail})
    platform.EntityData.Children.Append("slices", types.YChild{"Slices", &platform.Slices})
    platform.EntityData.Leafs = types.NewOrderedMap()

    platform.EntityData.YListKeys = []string {}

    return &(platform.EntityData)
}

// Oper_Platform_Summary
type Oper_Platform_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Platform_Summary_Location.
    Location []*Oper_Platform_Summary_Location
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

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range summary.Location {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.Location[i]), types.YChild{"Location", summary.Location[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Oper_Platform_Summary_Location
type Oper_Platform_Summary_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    Location interface{}

    
    SummaryData Oper_Platform_Summary_Location_SummaryData
}

func (location *Oper_Platform_Summary_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "summary"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("summary-data", types.YChild{"SummaryData", &location.SummaryData})
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

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

    summaryData.EntityData.Children = types.NewOrderedMap()
    summaryData.EntityData.Leafs = types.NewOrderedMap()
    summaryData.EntityData.Leafs.Append("card_type", types.YLeaf{"CardType", summaryData.CardType})
    summaryData.EntityData.Leafs.Append("hw_state", types.YLeaf{"HwState", summaryData.HwState})
    summaryData.EntityData.Leafs.Append("sw_state", types.YLeaf{"SwState", summaryData.SwState})
    summaryData.EntityData.Leafs.Append("config_state", types.YLeaf{"ConfigState", summaryData.ConfigState})

    summaryData.EntityData.YListKeys = []string {}

    return &(summaryData.EntityData)
}

// Oper_Platform_Detail
type Oper_Platform_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Platform_Detail_Location.
    Location []*Oper_Platform_Detail_Location
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

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range detail.Location {
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.Location[i]), types.YChild{"Location", detail.Location[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Oper_Platform_Detail_Location
type Oper_Platform_Detail_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    Location interface{}

    
    DetailData Oper_Platform_Detail_Location_DetailData
}

func (location *Oper_Platform_Detail_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "detail"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("detail-data", types.YChild{"DetailData", &location.DetailData})
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

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

    detailData.EntityData.Children = types.NewOrderedMap()
    detailData.EntityData.Leafs = types.NewOrderedMap()
    detailData.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", detailData.Pid})
    detailData.EntityData.Leafs.Append("description", types.YLeaf{"Description", detailData.Description})
    detailData.EntityData.Leafs.Append("vid_sn", types.YLeaf{"VidSn", detailData.VidSn})
    detailData.EntityData.Leafs.Append("hw_state", types.YLeaf{"HwState", detailData.HwState})
    detailData.EntityData.Leafs.Append("sw_state", types.YLeaf{"SwState", detailData.SwState})
    detailData.EntityData.Leafs.Append("config_wordy", types.YLeaf{"ConfigWordy", detailData.ConfigWordy})
    detailData.EntityData.Leafs.Append("hw_ver", types.YLeaf{"HwVer", detailData.HwVer})
    detailData.EntityData.Leafs.Append("last_event", types.YLeaf{"LastEvent", detailData.LastEvent})
    detailData.EntityData.Leafs.Append("last_ev_reason_str", types.YLeaf{"LastEvReasonStr", detailData.LastEvReasonStr})

    detailData.EntityData.YListKeys = []string {}

    return &(detailData.EntityData)
}

// Oper_Platform_Slices
type Oper_Platform_Slices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Platform_Slices_Location.
    Location []*Oper_Platform_Slices_Location
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

    slices.EntityData.Children = types.NewOrderedMap()
    slices.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range slices.Location {
        slices.EntityData.Children.Append(types.GetSegmentPath(slices.Location[i]), types.YChild{"Location", slices.Location[i]})
    }
    slices.EntityData.Leafs = types.NewOrderedMap()

    slices.EntityData.YListKeys = []string {}

    return &(slices.EntityData)
}

// Oper_Platform_Slices_Location
type Oper_Platform_Slices_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    Location interface{}

    // The type is slice of Oper_Platform_Slices_Location_SliceValues.
    SliceValues []*Oper_Platform_Slices_Location_SliceValues
}

func (location *Oper_Platform_Slices_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "slices"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("slice_values", types.YChild{"SliceValues", nil})
    for i := range location.SliceValues {
        location.EntityData.Children.Append(types.GetSegmentPath(location.SliceValues[i]), types.YChild{"SliceValues", location.SliceValues[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

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
    sliceValues.EntityData.SegmentPath = "slice_values" + types.AddKeyToken(sliceValues.SliceIdx, "slice_idx")
    sliceValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceValues.EntityData.Children = types.NewOrderedMap()
    sliceValues.EntityData.Children.Append("slice", types.YChild{"Slice", &sliceValues.Slice})
    sliceValues.EntityData.Leafs = types.NewOrderedMap()
    sliceValues.EntityData.Leafs.Append("slice_idx", types.YLeaf{"SliceIdx", sliceValues.SliceIdx})

    sliceValues.EntityData.YListKeys = []string {"SliceIdx"}

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

    slice.EntityData.Children = types.NewOrderedMap()
    slice.EntityData.Leafs = types.NewOrderedMap()
    slice.EntityData.Leafs.Append("slice_num", types.YLeaf{"SliceNum", slice.SliceNum})
    slice.EntityData.Leafs.Append("admin_state", types.YLeaf{"AdminState", slice.AdminState})
    slice.EntityData.Leafs.Append("oper_state", types.YLeaf{"OperState", slice.OperState})

    slice.EntityData.YListKeys = []string {}

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

    chassis.EntityData.Children = types.NewOrderedMap()
    chassis.EntityData.Children.Append("brief", types.YChild{"Brief", &chassis.Brief})
    chassis.EntityData.Leafs = types.NewOrderedMap()

    chassis.EntityData.YListKeys = []string {}

    return &(chassis.EntityData)
}

// Oper_Chassis_Brief
type Oper_Chassis_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Chassis_Brief_ChassisSerial.
    ChassisSerial []*Oper_Chassis_Brief_ChassisSerial
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

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Children.Append("chassis_serial", types.YChild{"ChassisSerial", nil})
    for i := range brief.ChassisSerial {
        brief.EntityData.Children.Append(types.GetSegmentPath(brief.ChassisSerial[i]), types.YChild{"ChassisSerial", brief.ChassisSerial[i]})
    }
    brief.EntityData.Leafs = types.NewOrderedMap()

    brief.EntityData.YListKeys = []string {}

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
    chassisSerial.EntityData.SegmentPath = "chassis_serial" + types.AddKeyToken(chassisSerial.SerialNumber, "serial_number")
    chassisSerial.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisSerial.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisSerial.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisSerial.EntityData.Children = types.NewOrderedMap()
    chassisSerial.EntityData.Children.Append("brief-data", types.YChild{"BriefData", &chassisSerial.BriefData})
    chassisSerial.EntityData.Leafs = types.NewOrderedMap()
    chassisSerial.EntityData.Leafs.Append("serial_number", types.YLeaf{"SerialNumber", chassisSerial.SerialNumber})

    chassisSerial.EntityData.YListKeys = []string {"SerialNumber"}

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

    briefData.EntityData.Children = types.NewOrderedMap()
    briefData.EntityData.Leafs = types.NewOrderedMap()
    briefData.EntityData.Leafs.Append("racknum", types.YLeaf{"Racknum", briefData.Racknum})
    briefData.EntityData.Leafs.Append("rack_type", types.YLeaf{"RackType", briefData.RackType})
    briefData.EntityData.Leafs.Append("rack_state", types.YLeaf{"RackState", briefData.RackState})
    briefData.EntityData.Leafs.Append("data_plane", types.YLeaf{"DataPlane", briefData.DataPlane})
    briefData.EntityData.Leafs.Append("ctrl_plane", types.YLeaf{"CtrlPlane", briefData.CtrlPlane})

    briefData.EntityData.YListKeys = []string {}

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

    reload.EntityData.Children = types.NewOrderedMap()
    reload.EntityData.Children.Append("rack", types.YChild{"Rack", &reload.Rack})
    reload.EntityData.Leafs = types.NewOrderedMap()

    reload.EntityData.YListKeys = []string {}

    return &(reload.EntityData)
}

// Oper_Reload_Rack
type Oper_Reload_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Reload_Rack_Racks.
    Racks []*Oper_Reload_Rack_Racks
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

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("racks", types.YChild{"Racks", nil})
    for i := range rack.Racks {
        rack.EntityData.Children.Append(types.GetSegmentPath(rack.Racks[i]), types.YChild{"Racks", rack.Racks[i]})
    }
    rack.EntityData.Leafs = types.NewOrderedMap()

    rack.EntityData.YListKeys = []string {}

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
    racks.EntityData.SegmentPath = "racks" + types.AddKeyToken(racks.Rack, "rack")
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = types.NewOrderedMap()
    racks.EntityData.Leafs = types.NewOrderedMap()
    racks.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", racks.Rack})

    racks.EntityData.YListKeys = []string {"Rack"}

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

    rebootHistory.EntityData.Children = types.NewOrderedMap()
    rebootHistory.EntityData.Children.Append("card", types.YChild{"Card", &rebootHistory.Card})
    rebootHistory.EntityData.Children.Append("admin-vm", types.YChild{"AdminVm", &rebootHistory.AdminVm})
    rebootHistory.EntityData.Children.Append("reverse", types.YChild{"Reverse", &rebootHistory.Reverse})
    rebootHistory.EntityData.Leafs = types.NewOrderedMap()

    rebootHistory.EntityData.YListKeys = []string {}

    return &(rebootHistory.EntityData)
}

// Oper_RebootHistory_Card
type Oper_RebootHistory_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_RebootHistory_Card_Location.
    Location []*Oper_RebootHistory_Card_Location
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

    card.EntityData.Children = types.NewOrderedMap()
    card.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range card.Location {
        card.EntityData.Children.Append(types.GetSegmentPath(card.Location[i]), types.YChild{"Location", card.Location[i]})
    }
    card.EntityData.Leafs = types.NewOrderedMap()

    card.EntityData.YListKeys = []string {}

    return &(card.EntityData)
}

// Oper_RebootHistory_Card_Location
type Oper_RebootHistory_Card_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    Location interface{}

    // The type is slice of Oper_RebootHistory_Card_Location_Events.
    Events []*Oper_RebootHistory_Card_Location_Events
}

func (location *Oper_RebootHistory_Card_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "card"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("events", types.YChild{"Events", nil})
    for i := range location.Events {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Events[i]), types.YChild{"Events", location.Events[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

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
    events.EntityData.SegmentPath = "events" + types.AddKeyToken(events.EventIdx, "event_idx")
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = types.NewOrderedMap()
    events.EntityData.Children.Append("event", types.YChild{"Event", &events.Event})
    events.EntityData.Leafs = types.NewOrderedMap()
    events.EntityData.Leafs.Append("event_idx", types.YLeaf{"EventIdx", events.EventIdx})

    events.EntityData.YListKeys = []string {"EventIdx"}

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

    event.EntityData.Children = types.NewOrderedMap()
    event.EntityData.Leafs = types.NewOrderedMap()
    event.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", event.Timestamp})
    event.EntityData.Leafs.Append("reason_code", types.YLeaf{"ReasonCode", event.ReasonCode})
    event.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", event.Reason})
    event.EntityData.Leafs.Append("src_loc", types.YLeaf{"SrcLoc", event.SrcLoc})
    event.EntityData.Leafs.Append("src_name", types.YLeaf{"SrcName", event.SrcName})

    event.EntityData.YListKeys = []string {}

    return &(event.EntityData)
}

// Oper_RebootHistory_AdminVm
type Oper_RebootHistory_AdminVm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_RebootHistory_AdminVm_Location.
    Location []*Oper_RebootHistory_AdminVm_Location
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

    adminVm.EntityData.Children = types.NewOrderedMap()
    adminVm.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range adminVm.Location {
        adminVm.EntityData.Children.Append(types.GetSegmentPath(adminVm.Location[i]), types.YChild{"Location", adminVm.Location[i]})
    }
    adminVm.EntityData.Leafs = types.NewOrderedMap()

    adminVm.EntityData.YListKeys = []string {}

    return &(adminVm.EntityData)
}

// Oper_RebootHistory_AdminVm_Location
type Oper_RebootHistory_AdminVm_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    Location interface{}

    // The type is slice of Oper_RebootHistory_AdminVm_Location_Events.
    Events []*Oper_RebootHistory_AdminVm_Location_Events
}

func (location *Oper_RebootHistory_AdminVm_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "admin-vm"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("events", types.YChild{"Events", nil})
    for i := range location.Events {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Events[i]), types.YChild{"Events", location.Events[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

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
    events.EntityData.SegmentPath = "events" + types.AddKeyToken(events.EventIdx, "event_idx")
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = types.NewOrderedMap()
    events.EntityData.Children.Append("event", types.YChild{"Event", &events.Event})
    events.EntityData.Leafs = types.NewOrderedMap()
    events.EntityData.Leafs.Append("event_idx", types.YLeaf{"EventIdx", events.EventIdx})

    events.EntityData.YListKeys = []string {"EventIdx"}

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

    event.EntityData.Children = types.NewOrderedMap()
    event.EntityData.Leafs = types.NewOrderedMap()
    event.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", event.Timestamp})
    event.EntityData.Leafs.Append("reason_code", types.YLeaf{"ReasonCode", event.ReasonCode})
    event.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", event.Reason})
    event.EntityData.Leafs.Append("src_loc", types.YLeaf{"SrcLoc", event.SrcLoc})
    event.EntityData.Leafs.Append("src_name", types.YLeaf{"SrcName", event.SrcName})

    event.EntityData.YListKeys = []string {}

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

    reverse.EntityData.Children = types.NewOrderedMap()
    reverse.EntityData.Children.Append("card", types.YChild{"Card", &reverse.Card})
    reverse.EntityData.Children.Append("admin-vm", types.YChild{"AdminVm", &reverse.AdminVm})
    reverse.EntityData.Leafs = types.NewOrderedMap()

    reverse.EntityData.YListKeys = []string {}

    return &(reverse.EntityData)
}

// Oper_RebootHistory_Reverse_Card
type Oper_RebootHistory_Reverse_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_RebootHistory_Reverse_Card_Location.
    Location []*Oper_RebootHistory_Reverse_Card_Location
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

    card.EntityData.Children = types.NewOrderedMap()
    card.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range card.Location {
        card.EntityData.Children.Append(types.GetSegmentPath(card.Location[i]), types.YChild{"Location", card.Location[i]})
    }
    card.EntityData.Leafs = types.NewOrderedMap()

    card.EntityData.YListKeys = []string {}

    return &(card.EntityData)
}

// Oper_RebootHistory_Reverse_Card_Location
type Oper_RebootHistory_Reverse_Card_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    Location interface{}

    // The type is slice of Oper_RebootHistory_Reverse_Card_Location_Events.
    Events []*Oper_RebootHistory_Reverse_Card_Location_Events
}

func (location *Oper_RebootHistory_Reverse_Card_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "card"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("events", types.YChild{"Events", nil})
    for i := range location.Events {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Events[i]), types.YChild{"Events", location.Events[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

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
    events.EntityData.SegmentPath = "events" + types.AddKeyToken(events.EventIdx, "event_idx")
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = types.NewOrderedMap()
    events.EntityData.Children.Append("event", types.YChild{"Event", &events.Event})
    events.EntityData.Leafs = types.NewOrderedMap()
    events.EntityData.Leafs.Append("event_idx", types.YLeaf{"EventIdx", events.EventIdx})

    events.EntityData.YListKeys = []string {"EventIdx"}

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

    event.EntityData.Children = types.NewOrderedMap()
    event.EntityData.Leafs = types.NewOrderedMap()
    event.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", event.Timestamp})
    event.EntityData.Leafs.Append("reason_code", types.YLeaf{"ReasonCode", event.ReasonCode})
    event.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", event.Reason})
    event.EntityData.Leafs.Append("src_loc", types.YLeaf{"SrcLoc", event.SrcLoc})
    event.EntityData.Leafs.Append("src_name", types.YLeaf{"SrcName", event.SrcName})

    event.EntityData.YListKeys = []string {}

    return &(event.EntityData)
}

// Oper_RebootHistory_Reverse_AdminVm
type Oper_RebootHistory_Reverse_AdminVm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_RebootHistory_Reverse_AdminVm_Location.
    Location []*Oper_RebootHistory_Reverse_AdminVm_Location
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

    adminVm.EntityData.Children = types.NewOrderedMap()
    adminVm.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range adminVm.Location {
        adminVm.EntityData.Children.Append(types.GetSegmentPath(adminVm.Location[i]), types.YChild{"Location", adminVm.Location[i]})
    }
    adminVm.EntityData.Leafs = types.NewOrderedMap()

    adminVm.EntityData.YListKeys = []string {}

    return &(adminVm.EntityData)
}

// Oper_RebootHistory_Reverse_AdminVm_Location
type Oper_RebootHistory_Reverse_AdminVm_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // ((([bB][0-9])/(([a-zA-Z]){2}\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\d{1,2})))(/[cC][pP][uU]0)?.
    Location interface{}

    // The type is slice of Oper_RebootHistory_Reverse_AdminVm_Location_Events.
    Events []*Oper_RebootHistory_Reverse_AdminVm_Location_Events
}

func (location *Oper_RebootHistory_Reverse_AdminVm_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "admin-vm"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("events", types.YChild{"Events", nil})
    for i := range location.Events {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Events[i]), types.YChild{"Events", location.Events[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

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
    events.EntityData.SegmentPath = "events" + types.AddKeyToken(events.EventIdx, "event_idx")
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = types.NewOrderedMap()
    events.EntityData.Children.Append("event", types.YChild{"Event", &events.Event})
    events.EntityData.Leafs = types.NewOrderedMap()
    events.EntityData.Leafs.Append("event_idx", types.YLeaf{"EventIdx", events.EventIdx})

    events.EntityData.YListKeys = []string {"EventIdx"}

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

    event.EntityData.Children = types.NewOrderedMap()
    event.EntityData.Leafs = types.NewOrderedMap()
    event.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", event.Timestamp})
    event.EntityData.Leafs.Append("reason_code", types.YLeaf{"ReasonCode", event.ReasonCode})
    event.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", event.Reason})
    event.EntityData.Leafs.Append("src_loc", types.YLeaf{"SrcLoc", event.SrcLoc})
    event.EntityData.Leafs.Append("src_name", types.YLeaf{"SrcName", event.SrcName})

    event.EntityData.YListKeys = []string {}

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
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Ifname, "ifname")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("interface-data", types.YChild{"InterfaceData", &self.InterfaceData})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("ifname", types.YLeaf{"Ifname", self.Ifname})

    self.EntityData.YListKeys = []string {"Ifname"}

    return &(self.EntityData)
}

// Oper_Interface_InterfaceData
type Oper_Interface_InterfaceData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Mac interface{}

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    interfaceData.EntityData.Children = types.NewOrderedMap()
    interfaceData.EntityData.Leafs = types.NewOrderedMap()
    interfaceData.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", interfaceData.Mac})
    interfaceData.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", interfaceData.Ipv4})
    interfaceData.EntityData.Leafs.Append("flagstr", types.YLeaf{"Flagstr", interfaceData.Flagstr})
    interfaceData.EntityData.Leafs.Append("port_status", types.YLeaf{"PortStatus", interfaceData.PortStatus})
    interfaceData.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", interfaceData.Mtu})
    interfaceData.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", interfaceData.Metric})
    interfaceData.EntityData.Leafs.Append("rx_pak", types.YLeaf{"RxPak", interfaceData.RxPak})
    interfaceData.EntityData.Leafs.Append("rx_errors", types.YLeaf{"RxErrors", interfaceData.RxErrors})
    interfaceData.EntityData.Leafs.Append("rx_dropped", types.YLeaf{"RxDropped", interfaceData.RxDropped})
    interfaceData.EntityData.Leafs.Append("rx_overruns", types.YLeaf{"RxOverruns", interfaceData.RxOverruns})
    interfaceData.EntityData.Leafs.Append("rx_frame", types.YLeaf{"RxFrame", interfaceData.RxFrame})
    interfaceData.EntityData.Leafs.Append("tx_pak", types.YLeaf{"TxPak", interfaceData.TxPak})
    interfaceData.EntityData.Leafs.Append("tx_errors", types.YLeaf{"TxErrors", interfaceData.TxErrors})
    interfaceData.EntityData.Leafs.Append("tx_dropped", types.YLeaf{"TxDropped", interfaceData.TxDropped})
    interfaceData.EntityData.Leafs.Append("tx_overruns", types.YLeaf{"TxOverruns", interfaceData.TxOverruns})
    interfaceData.EntityData.Leafs.Append("tx_carrier", types.YLeaf{"TxCarrier", interfaceData.TxCarrier})
    interfaceData.EntityData.Leafs.Append("collisions", types.YLeaf{"Collisions", interfaceData.Collisions})
    interfaceData.EntityData.Leafs.Append("tx_queuelen", types.YLeaf{"TxQueuelen", interfaceData.TxQueuelen})
    interfaceData.EntityData.Leafs.Append("rx_bytes", types.YLeaf{"RxBytes", interfaceData.RxBytes})
    interfaceData.EntityData.Leafs.Append("tx_bytes", types.YLeaf{"TxBytes", interfaceData.TxBytes})
    interfaceData.EntityData.Leafs.Append("intf_num", types.YLeaf{"IntfNum", interfaceData.IntfNum})

    interfaceData.EntityData.YListKeys = []string {}

    return &(interfaceData.EntityData)
}

// Oper_ReloadVm
type Oper_ReloadVm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_ReloadVm_Location.
    Location []*Oper_ReloadVm_Location
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

    reloadVm.EntityData.Children = types.NewOrderedMap()
    reloadVm.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range reloadVm.Location {
        reloadVm.EntityData.Children.Append(types.GetSegmentPath(reloadVm.Location[i]), types.YChild{"Location", reloadVm.Location[i]})
    }
    reloadVm.EntityData.Leafs = types.NewOrderedMap()

    reloadVm.EntityData.YListKeys = []string {}

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
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location, "location")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location", types.YLeaf{"Location", location.Location})

    location.EntityData.YListKeys = []string {"Location"}

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

    macpool.EntityData.Children = types.NewOrderedMap()
    macpool.EntityData.Children.Append("brief", types.YChild{"Brief", &macpool.Brief})
    macpool.EntityData.Leafs = types.NewOrderedMap()

    macpool.EntityData.YListKeys = []string {}

    return &(macpool.EntityData)
}

// Oper_Macpool_Brief
type Oper_Macpool_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_Macpool_Brief_Rack.
    Rack []*Oper_Macpool_Brief_Rack
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

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range brief.Rack {
        brief.EntityData.Children.Append(types.GetSegmentPath(brief.Rack[i]), types.YChild{"Rack", brief.Rack[i]})
    }
    brief.EntityData.Leafs = types.NewOrderedMap()

    brief.EntityData.YListKeys = []string {}

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
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.SerialNumber, "serial_number")
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("brief-data", types.YChild{"BriefData", &rack.BriefData})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("serial_number", types.YLeaf{"SerialNumber", rack.SerialNumber})

    rack.EntityData.YListKeys = []string {"SerialNumber"}

    return &(rack.EntityData)
}

// Oper_Macpool_Brief_Rack_BriefData
type Oper_Macpool_Brief_Rack_BriefData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Racknum interface{}

    // The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

    briefData.EntityData.Children = types.NewOrderedMap()
    briefData.EntityData.Leafs = types.NewOrderedMap()
    briefData.EntityData.Leafs.Append("racknum", types.YLeaf{"Racknum", briefData.Racknum})
    briefData.EntityData.Leafs.Append("mac_base", types.YLeaf{"MacBase", briefData.MacBase})
    briefData.EntityData.Leafs.Append("mac_count", types.YLeaf{"MacCount", briefData.MacCount})
    briefData.EntityData.Leafs.Append("is_selected", types.YLeaf{"IsSelected", briefData.IsSelected})
    briefData.EntityData.Leafs.Append("allocated_count", types.YLeaf{"AllocatedCount", briefData.AllocatedCount})

    briefData.EntityData.YListKeys = []string {}

    return &(briefData.EntityData)
}

// Config
type Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Chassis Config_Chassis

    
    Interface Config_Interface

    
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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("chassis", types.YChild{"Chassis", &config.Chassis})
    config.EntityData.Children.Append("interface", types.YChild{"Interface", &config.Interface})
    config.EntityData.Children.Append("domain", types.YChild{"Domain", &config.Domain})
    config.EntityData.Children.Append("virtual-macaddr-range", types.YChild{"VirtualMacaddrRange", &config.VirtualMacaddrRange})
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Config_Chassis
type Config_Chassis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Chassis_Serial.
    Serial []*Config_Chassis_Serial
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

    chassis.EntityData.Children = types.NewOrderedMap()
    chassis.EntityData.Children.Append("serial", types.YChild{"Serial", nil})
    for i := range chassis.Serial {
        chassis.EntityData.Children.Append(types.GetSegmentPath(chassis.Serial[i]), types.YChild{"Serial", chassis.Serial[i]})
    }
    chassis.EntityData.Leafs = types.NewOrderedMap()

    chassis.EntityData.YListKeys = []string {}

    return &(chassis.EntityData)
}

// Config_Chassis_Serial
type Config_Chassis_Serial struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. This attribute is mandatory.
    Serial interface{}

    // The type is string with pattern:
    // [bB][0-9]|[fF][0-7]|[0-9]|[1][0-5]|[2][4][0-7]. This attribute is
    // mandatory.
    Rack interface{}
}

func (serial *Config_Chassis_Serial) GetEntityData() *types.CommonEntityData {
    serial.EntityData.YFilter = serial.YFilter
    serial.EntityData.YangName = "serial"
    serial.EntityData.BundleName = "cisco_ios_xr"
    serial.EntityData.ParentYangName = "chassis"
    serial.EntityData.SegmentPath = "serial" + types.AddKeyToken(serial.Serial, "serial")
    serial.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serial.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serial.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serial.EntityData.Children = types.NewOrderedMap()
    serial.EntityData.Leafs = types.NewOrderedMap()
    serial.EntityData.Leafs.Append("serial", types.YLeaf{"Serial", serial.Serial})
    serial.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", serial.Rack})

    serial.EntityData.YListKeys = []string {"Serial"}

    return &(serial.EntityData)
}

// Config_Interface
type Config_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    MgmtEth Config_Interface_MgmtEth
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

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("MgmtEth", types.YChild{"MgmtEth", &self.MgmtEth})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Config_Interface_MgmtEth
type Config_Interface_MgmtEth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Interface_MgmtEth_Locport.
    Locport []*Config_Interface_MgmtEth_Locport
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

    mgmtEth.EntityData.Children = types.NewOrderedMap()
    mgmtEth.EntityData.Children.Append("locport", types.YChild{"Locport", nil})
    for i := range mgmtEth.Locport {
        mgmtEth.EntityData.Children.Append(types.GetSegmentPath(mgmtEth.Locport[i]), types.YChild{"Locport", mgmtEth.Locport[i]})
    }
    mgmtEth.EntityData.Leafs = types.NewOrderedMap()

    mgmtEth.EntityData.YListKeys = []string {}

    return &(mgmtEth.EntityData)
}

// Config_Interface_MgmtEth_Locport
type Config_Interface_MgmtEth_Locport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // [fF][0-7]|[0-9]|[1][0-5]|[bB][0-9].
    Rack interface{}

    // This attribute is a key. The type is string with pattern:
    // [Rr][Pp][0-1]|[Rr][Ss][Pp][0-1]|[Ss][Cc][0-1]|[cC][bB][0-9].
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DefaultGw interface{}

    
    Ipv4 Config_Interface_MgmtEth_Locport_Ipv4

    
    Arp Config_Interface_MgmtEth_Locport_Arp
}

func (locport *Config_Interface_MgmtEth_Locport) GetEntityData() *types.CommonEntityData {
    locport.EntityData.YFilter = locport.YFilter
    locport.EntityData.YangName = "locport"
    locport.EntityData.BundleName = "cisco_ios_xr"
    locport.EntityData.ParentYangName = "MgmtEth"
    locport.EntityData.SegmentPath = "locport" + types.AddKeyToken(locport.Rack, "rack") + types.AddKeyToken(locport.Slot, "slot") + types.AddKeyToken(locport.Intf, "intf") + types.AddKeyToken(locport.Port, "port")
    locport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locport.EntityData.Children = types.NewOrderedMap()
    locport.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &locport.Ipv4})
    locport.EntityData.Children.Append("arp", types.YChild{"Arp", &locport.Arp})
    locport.EntityData.Leafs = types.NewOrderedMap()
    locport.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", locport.Rack})
    locport.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", locport.Slot})
    locport.EntityData.Leafs.Append("intf", types.YLeaf{"Intf", locport.Intf})
    locport.EntityData.Leafs.Append("port", types.YLeaf{"Port", locport.Port})
    locport.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", locport.Shutdown})
    locport.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", locport.Mtu})
    locport.EntityData.Leafs.Append("default-gw", types.YLeaf{"DefaultGw", locport.DefaultGw})

    locport.EntityData.YListKeys = []string {"Rack", "Slot", "Intf", "Port"}

    return &(locport.EntityData)
}

// Config_Interface_MgmtEth_Locport_Ipv4
type Config_Interface_MgmtEth_Locport_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])((
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))|(/([0-9]+))).
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

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4.Address})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Config_Interface_MgmtEth_Locport_Arp
type Config_Interface_MgmtEth_Locport_Arp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Interface_MgmtEth_Locport_Arp_Ip.
    Ip []*Config_Interface_MgmtEth_Locport_Arp_Ip
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

    arp.EntityData.Children = types.NewOrderedMap()
    arp.EntityData.Children.Append("ip", types.YChild{"Ip", nil})
    for i := range arp.Ip {
        arp.EntityData.Children.Append(types.GetSegmentPath(arp.Ip[i]), types.YChild{"Ip", arp.Ip[i]})
    }
    arp.EntityData.Leafs = types.NewOrderedMap()

    arp.EntityData.YListKeys = []string {}

    return &(arp.EntityData)
}

// Config_Interface_MgmtEth_Locport_Arp_Ip
type Config_Interface_MgmtEth_Locport_Arp_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Ip interface{}

    // The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}. This
    // attribute is mandatory.
    Mac interface{}
}

func (ip *Config_Interface_MgmtEth_Locport_Arp_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "arp"
    ip.EntityData.SegmentPath = "ip" + types.AddKeyToken(ip.Ip, "ip")
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = types.NewOrderedMap()
    ip.EntityData.Leafs = types.NewOrderedMap()
    ip.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", ip.Ip})
    ip.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", ip.Mac})

    ip.EntityData.YListKeys = []string {"Ip"}

    return &(ip.EntityData)
}

// Config_Domain
type Config_Domain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Config_Domain_Name.
    Name []*Config_Domain_Name

    // The type is slice of Config_Domain_NameServer.
    NameServer []*Config_Domain_NameServer
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

    domain.EntityData.Children = types.NewOrderedMap()
    domain.EntityData.Children.Append("name", types.YChild{"Name", nil})
    for i := range domain.Name {
        domain.EntityData.Children.Append(types.GetSegmentPath(domain.Name[i]), types.YChild{"Name", domain.Name[i]})
    }
    domain.EntityData.Children.Append("name-server", types.YChild{"NameServer", nil})
    for i := range domain.NameServer {
        domain.EntityData.Children.Append(types.GetSegmentPath(domain.NameServer[i]), types.YChild{"NameServer", domain.NameServer[i]})
    }
    domain.EntityData.Leafs = types.NewOrderedMap()

    domain.EntityData.YListKeys = []string {}

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
    name.EntityData.SegmentPath = "name" + types.AddKeyToken(name.Name, "name")
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = types.NewOrderedMap()
    name.EntityData.Leafs = types.NewOrderedMap()
    name.EntityData.Leafs.Append("name", types.YLeaf{"Name", name.Name})

    name.EntityData.YListKeys = []string {"Name"}

    return &(name.EntityData)
}

// Config_Domain_NameServer
type Config_Domain_NameServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NameServer interface{}
}

func (nameServer *Config_Domain_NameServer) GetEntityData() *types.CommonEntityData {
    nameServer.EntityData.YFilter = nameServer.YFilter
    nameServer.EntityData.YangName = "name-server"
    nameServer.EntityData.BundleName = "cisco_ios_xr"
    nameServer.EntityData.ParentYangName = "domain"
    nameServer.EntityData.SegmentPath = "name-server" + types.AddKeyToken(nameServer.NameServer, "name-server")
    nameServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nameServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nameServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nameServer.EntityData.Children = types.NewOrderedMap()
    nameServer.EntityData.Leafs = types.NewOrderedMap()
    nameServer.EntityData.Leafs.Append("name-server", types.YLeaf{"NameServer", nameServer.NameServer})

    nameServer.EntityData.YListKeys = []string {"NameServer"}

    return &(nameServer.EntityData)
}

// Config_VirtualMacaddrRange
type Config_VirtualMacaddrRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

    virtualMacaddrRange.EntityData.Children = types.NewOrderedMap()
    virtualMacaddrRange.EntityData.Leafs = types.NewOrderedMap()
    virtualMacaddrRange.EntityData.Leafs.Append("base", types.YLeaf{"Base", virtualMacaddrRange.Base})
    virtualMacaddrRange.EntityData.Leafs.Append("count", types.YLeaf{"Count", virtualMacaddrRange.Count})

    virtualMacaddrRange.EntityData.YListKeys = []string {}

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

    actions.EntityData.Children = types.NewOrderedMap()
    actions.EntityData.Leafs = types.NewOrderedMap()

    actions.EntityData.YListKeys = []string {}

    return &(actions.EntityData)
}

