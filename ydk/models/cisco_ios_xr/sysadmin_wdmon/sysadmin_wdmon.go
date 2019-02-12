package sysadmin_wdmon

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_wdmon"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-wdmon wdmon}", reflect.TypeOf(Wdmon{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-wdmon:wdmon", reflect.TypeOf(Wdmon{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-wdmon wdmon-info}", reflect.TypeOf(WdmonInfo{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-wdmon:wdmon-info", reflect.TypeOf(WdmonInfo{}))
}

// Wdmon
type Wdmon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Wdmon_Trace.
    Trace []*Wdmon_Trace
}

func (wdmon *Wdmon) GetEntityData() *types.CommonEntityData {
    wdmon.EntityData.YFilter = wdmon.YFilter
    wdmon.EntityData.YangName = "wdmon"
    wdmon.EntityData.BundleName = "cisco_ios_xr"
    wdmon.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-wdmon"
    wdmon.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-wdmon:wdmon"
    wdmon.EntityData.AbsolutePath = wdmon.EntityData.SegmentPath
    wdmon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wdmon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wdmon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wdmon.EntityData.Children = types.NewOrderedMap()
    wdmon.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range wdmon.Trace {
        wdmon.EntityData.Children.Append(types.GetSegmentPath(wdmon.Trace[i]), types.YChild{"Trace", wdmon.Trace[i]})
    }
    wdmon.EntityData.Leafs = types.NewOrderedMap()

    wdmon.EntityData.YListKeys = []string {}

    return &(wdmon.EntityData)
}

// Wdmon_Trace
// show traceable processes
type Wdmon_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Wdmon_Trace_Location.
    Location []*Wdmon_Trace_Location
}

func (trace *Wdmon_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "wdmon"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-wdmon:wdmon/" + trace.EntityData.SegmentPath
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

// Wdmon_Trace_Location
type Wdmon_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Wdmon_Trace_Location_AllOptions.
    AllOptions []*Wdmon_Trace_Location_AllOptions
}

func (location *Wdmon_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-wdmon:wdmon/trace/" + location.EntityData.SegmentPath
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

// Wdmon_Trace_Location_AllOptions
type Wdmon_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Wdmon_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Wdmon_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Wdmon_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-wdmon:wdmon/trace/location/" + allOptions.EntityData.SegmentPath
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

// Wdmon_Trace_Location_AllOptions_TraceBlocks
type Wdmon_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Wdmon_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-wdmon:wdmon/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

// WdmonInfo
type WdmonInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of WdmonInfo_AllLocations.
    AllLocations []*WdmonInfo_AllLocations
}

func (wdmonInfo *WdmonInfo) GetEntityData() *types.CommonEntityData {
    wdmonInfo.EntityData.YFilter = wdmonInfo.YFilter
    wdmonInfo.EntityData.YangName = "wdmon-info"
    wdmonInfo.EntityData.BundleName = "cisco_ios_xr"
    wdmonInfo.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-wdmon"
    wdmonInfo.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-wdmon:wdmon-info"
    wdmonInfo.EntityData.AbsolutePath = wdmonInfo.EntityData.SegmentPath
    wdmonInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wdmonInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wdmonInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wdmonInfo.EntityData.Children = types.NewOrderedMap()
    wdmonInfo.EntityData.Children.Append("all-locations", types.YChild{"AllLocations", nil})
    for i := range wdmonInfo.AllLocations {
        wdmonInfo.EntityData.Children.Append(types.GetSegmentPath(wdmonInfo.AllLocations[i]), types.YChild{"AllLocations", wdmonInfo.AllLocations[i]})
    }
    wdmonInfo.EntityData.Leafs = types.NewOrderedMap()

    wdmonInfo.EntityData.YListKeys = []string {}

    return &(wdmonInfo.EntityData)
}

// WdmonInfo_AllLocations
type WdmonInfo_AllLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // Last start date and time for wdmon. The type is string.
    StartTimestamp interface{}

    // wdmon/Calv/Hushd liveness timeout. The type is string.
    HushdTimeout interface{}

    // Calvados restart timeout. The type is interface{} with range:
    // 0..4294967295.
    CalvRestartTimeout interface{}

    // Hushd WD expire action disable or not. The type is bool.
    HushdWdActionDisable interface{}

    // Last Hushd WD expire action date and time. The type is string.
    HushdWdActionTimestamp interface{}

    // Hushd CAPI up or not. The type is bool.
    HushdCapiUp interface{}

    // Any pending response from Hushd. The type is bool.
    HushdPendingResp interface{}

    // Hushd HB punching stopped. The type is bool.
    HushdStopPunching interface{}

    // Last Hushd CAPI up date and time. The type is string.
    HushdCapiUpTimestamp interface{}

    // How long ago was last HB response from Hushd. The type is string.
    HushdLastHbResp interface{}

    // Num of Hushd CAPI connects. The type is interface{} with range:
    // 0..4294967295.
    HushdNumCapiConnects interface{}

    // wdmon client timeout action disabled or not. The type is bool.
    WdsActionDisable interface{}

    // Last wdmon client timeout action date and time. The type is string.
    WdsActionTimestamp interface{}

    // wdmon client restart timeout. The type is interface{} with range:
    // 0..4294967295.
    WdsRestartTimeout interface{}

    // wdmon client liveness timeout. The type is interface{} with range:
    // 0..4294967295.
    WdsLivenessTimeout interface{}

    // wdmon client is up or not. The type is bool.
    WdsClientUp interface{}

    // Process ID of the wdmon client. The type is interface{} with range:
    // 0..4294967295.
    WdsClientPid interface{}

    // Last wdmon client connect date and time. The type is string.
    WdsClientUpTimestamp interface{}

    // How long ago was last HB from wdmon client. The type is string.
    WdsClientLastHb interface{}

    // Total number of client connects. The type is interface{} with range:
    // 0..4294967295.
    WdsClientNumConnects interface{}

    // Total number of client liveness timeouts. The type is interface{} with
    // range: 0..4294967295.
    WdsNumLivenessTimeout interface{}

    // Total number of client restart timeouts. The type is interface{} with
    // range: 0..4294967295.
    WdsNumRestartTimeout interface{}

    // Status reported by wdmon client. The type is string.
    WdsClientReportedStatus interface{}
}

func (allLocations *WdmonInfo_AllLocations) GetEntityData() *types.CommonEntityData {
    allLocations.EntityData.YFilter = allLocations.YFilter
    allLocations.EntityData.YangName = "all-locations"
    allLocations.EntityData.BundleName = "cisco_ios_xr"
    allLocations.EntityData.ParentYangName = "wdmon-info"
    allLocations.EntityData.SegmentPath = "all-locations" + types.AddKeyToken(allLocations.Location, "location")
    allLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-wdmon:wdmon-info/" + allLocations.EntityData.SegmentPath
    allLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocations.EntityData.Children = types.NewOrderedMap()
    allLocations.EntityData.Leafs = types.NewOrderedMap()
    allLocations.EntityData.Leafs.Append("location", types.YLeaf{"Location", allLocations.Location})
    allLocations.EntityData.Leafs.Append("start-timestamp", types.YLeaf{"StartTimestamp", allLocations.StartTimestamp})
    allLocations.EntityData.Leafs.Append("hushd-timeout", types.YLeaf{"HushdTimeout", allLocations.HushdTimeout})
    allLocations.EntityData.Leafs.Append("calv-restart-timeout", types.YLeaf{"CalvRestartTimeout", allLocations.CalvRestartTimeout})
    allLocations.EntityData.Leafs.Append("hushd-wd-action-disable", types.YLeaf{"HushdWdActionDisable", allLocations.HushdWdActionDisable})
    allLocations.EntityData.Leafs.Append("hushd-wd-action-timestamp", types.YLeaf{"HushdWdActionTimestamp", allLocations.HushdWdActionTimestamp})
    allLocations.EntityData.Leafs.Append("hushd-capi-up", types.YLeaf{"HushdCapiUp", allLocations.HushdCapiUp})
    allLocations.EntityData.Leafs.Append("hushd-pending-resp", types.YLeaf{"HushdPendingResp", allLocations.HushdPendingResp})
    allLocations.EntityData.Leafs.Append("hushd-stop-punching", types.YLeaf{"HushdStopPunching", allLocations.HushdStopPunching})
    allLocations.EntityData.Leafs.Append("hushd-capi-up-timestamp", types.YLeaf{"HushdCapiUpTimestamp", allLocations.HushdCapiUpTimestamp})
    allLocations.EntityData.Leafs.Append("hushd-last-hb-resp", types.YLeaf{"HushdLastHbResp", allLocations.HushdLastHbResp})
    allLocations.EntityData.Leafs.Append("hushd-num-capi-connects", types.YLeaf{"HushdNumCapiConnects", allLocations.HushdNumCapiConnects})
    allLocations.EntityData.Leafs.Append("wds-action-disable", types.YLeaf{"WdsActionDisable", allLocations.WdsActionDisable})
    allLocations.EntityData.Leafs.Append("wds-action-timestamp", types.YLeaf{"WdsActionTimestamp", allLocations.WdsActionTimestamp})
    allLocations.EntityData.Leafs.Append("wds-restart-timeout", types.YLeaf{"WdsRestartTimeout", allLocations.WdsRestartTimeout})
    allLocations.EntityData.Leafs.Append("wds-liveness-timeout", types.YLeaf{"WdsLivenessTimeout", allLocations.WdsLivenessTimeout})
    allLocations.EntityData.Leafs.Append("wds-client-up", types.YLeaf{"WdsClientUp", allLocations.WdsClientUp})
    allLocations.EntityData.Leafs.Append("wds-client-pid", types.YLeaf{"WdsClientPid", allLocations.WdsClientPid})
    allLocations.EntityData.Leafs.Append("wds-client-up-timestamp", types.YLeaf{"WdsClientUpTimestamp", allLocations.WdsClientUpTimestamp})
    allLocations.EntityData.Leafs.Append("wds-client-last-hb", types.YLeaf{"WdsClientLastHb", allLocations.WdsClientLastHb})
    allLocations.EntityData.Leafs.Append("wds-client-num-connects", types.YLeaf{"WdsClientNumConnects", allLocations.WdsClientNumConnects})
    allLocations.EntityData.Leafs.Append("wds-num-liveness-timeout", types.YLeaf{"WdsNumLivenessTimeout", allLocations.WdsNumLivenessTimeout})
    allLocations.EntityData.Leafs.Append("wds-num-restart-timeout", types.YLeaf{"WdsNumRestartTimeout", allLocations.WdsNumRestartTimeout})
    allLocations.EntityData.Leafs.Append("wds-client-reported-status", types.YLeaf{"WdsClientReportedStatus", allLocations.WdsClientReportedStatus})

    allLocations.EntityData.YListKeys = []string {"Location"}

    return &(allLocations.EntityData)
}

