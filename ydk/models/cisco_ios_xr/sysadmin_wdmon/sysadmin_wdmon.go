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
    Trace []Wdmon_Trace
}

func (wdmon *Wdmon) GetEntityData() *types.CommonEntityData {
    wdmon.EntityData.YFilter = wdmon.YFilter
    wdmon.EntityData.YangName = "wdmon"
    wdmon.EntityData.BundleName = "cisco_ios_xr"
    wdmon.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-wdmon"
    wdmon.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-wdmon:wdmon"
    wdmon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wdmon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wdmon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wdmon.EntityData.Children = make(map[string]types.YChild)
    wdmon.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range wdmon.Trace {
        wdmon.EntityData.Children[types.GetSegmentPath(&wdmon.Trace[i])] = types.YChild{"Trace", &wdmon.Trace[i]}
    }
    wdmon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(wdmon.EntityData)
}

// Wdmon_Trace
// show traceable processes
type Wdmon_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Wdmon_Trace_Location.
    Location []Wdmon_Trace_Location
}

func (trace *Wdmon_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "wdmon"
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

// Wdmon_Trace_Location
type Wdmon_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Wdmon_Trace_Location_AllOptions.
    AllOptions []Wdmon_Trace_Location_AllOptions
}

func (location *Wdmon_Trace_Location) GetEntityData() *types.CommonEntityData {
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

// Wdmon_Trace_Location_AllOptions
type Wdmon_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of Wdmon_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Wdmon_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Wdmon_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
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

// Wdmon_Trace_Location_AllOptions_TraceBlocks
type Wdmon_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Wdmon_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
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

// WdmonInfo
type WdmonInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of WdmonInfo_AllLocations.
    AllLocations []WdmonInfo_AllLocations
}

func (wdmonInfo *WdmonInfo) GetEntityData() *types.CommonEntityData {
    wdmonInfo.EntityData.YFilter = wdmonInfo.YFilter
    wdmonInfo.EntityData.YangName = "wdmon-info"
    wdmonInfo.EntityData.BundleName = "cisco_ios_xr"
    wdmonInfo.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-wdmon"
    wdmonInfo.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-wdmon:wdmon-info"
    wdmonInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wdmonInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wdmonInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wdmonInfo.EntityData.Children = make(map[string]types.YChild)
    wdmonInfo.EntityData.Children["all-locations"] = types.YChild{"AllLocations", nil}
    for i := range wdmonInfo.AllLocations {
        wdmonInfo.EntityData.Children[types.GetSegmentPath(&wdmonInfo.AllLocations[i])] = types.YChild{"AllLocations", &wdmonInfo.AllLocations[i]}
    }
    wdmonInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(wdmonInfo.EntityData)
}

// WdmonInfo_AllLocations
type WdmonInfo_AllLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    allLocations.EntityData.SegmentPath = "all-locations" + "[location='" + fmt.Sprintf("%v", allLocations.Location) + "']"
    allLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocations.EntityData.Children = make(map[string]types.YChild)
    allLocations.EntityData.Leafs = make(map[string]types.YLeaf)
    allLocations.EntityData.Leafs["location"] = types.YLeaf{"Location", allLocations.Location}
    allLocations.EntityData.Leafs["start-timestamp"] = types.YLeaf{"StartTimestamp", allLocations.StartTimestamp}
    allLocations.EntityData.Leafs["hushd-timeout"] = types.YLeaf{"HushdTimeout", allLocations.HushdTimeout}
    allLocations.EntityData.Leafs["calv-restart-timeout"] = types.YLeaf{"CalvRestartTimeout", allLocations.CalvRestartTimeout}
    allLocations.EntityData.Leafs["hushd-wd-action-disable"] = types.YLeaf{"HushdWdActionDisable", allLocations.HushdWdActionDisable}
    allLocations.EntityData.Leafs["hushd-wd-action-timestamp"] = types.YLeaf{"HushdWdActionTimestamp", allLocations.HushdWdActionTimestamp}
    allLocations.EntityData.Leafs["hushd-capi-up"] = types.YLeaf{"HushdCapiUp", allLocations.HushdCapiUp}
    allLocations.EntityData.Leafs["hushd-pending-resp"] = types.YLeaf{"HushdPendingResp", allLocations.HushdPendingResp}
    allLocations.EntityData.Leafs["hushd-stop-punching"] = types.YLeaf{"HushdStopPunching", allLocations.HushdStopPunching}
    allLocations.EntityData.Leafs["hushd-capi-up-timestamp"] = types.YLeaf{"HushdCapiUpTimestamp", allLocations.HushdCapiUpTimestamp}
    allLocations.EntityData.Leafs["hushd-last-hb-resp"] = types.YLeaf{"HushdLastHbResp", allLocations.HushdLastHbResp}
    allLocations.EntityData.Leafs["hushd-num-capi-connects"] = types.YLeaf{"HushdNumCapiConnects", allLocations.HushdNumCapiConnects}
    allLocations.EntityData.Leafs["wds-action-disable"] = types.YLeaf{"WdsActionDisable", allLocations.WdsActionDisable}
    allLocations.EntityData.Leafs["wds-action-timestamp"] = types.YLeaf{"WdsActionTimestamp", allLocations.WdsActionTimestamp}
    allLocations.EntityData.Leafs["wds-restart-timeout"] = types.YLeaf{"WdsRestartTimeout", allLocations.WdsRestartTimeout}
    allLocations.EntityData.Leafs["wds-liveness-timeout"] = types.YLeaf{"WdsLivenessTimeout", allLocations.WdsLivenessTimeout}
    allLocations.EntityData.Leafs["wds-client-up"] = types.YLeaf{"WdsClientUp", allLocations.WdsClientUp}
    allLocations.EntityData.Leafs["wds-client-pid"] = types.YLeaf{"WdsClientPid", allLocations.WdsClientPid}
    allLocations.EntityData.Leafs["wds-client-up-timestamp"] = types.YLeaf{"WdsClientUpTimestamp", allLocations.WdsClientUpTimestamp}
    allLocations.EntityData.Leafs["wds-client-last-hb"] = types.YLeaf{"WdsClientLastHb", allLocations.WdsClientLastHb}
    allLocations.EntityData.Leafs["wds-client-num-connects"] = types.YLeaf{"WdsClientNumConnects", allLocations.WdsClientNumConnects}
    allLocations.EntityData.Leafs["wds-num-liveness-timeout"] = types.YLeaf{"WdsNumLivenessTimeout", allLocations.WdsNumLivenessTimeout}
    allLocations.EntityData.Leafs["wds-num-restart-timeout"] = types.YLeaf{"WdsNumRestartTimeout", allLocations.WdsNumRestartTimeout}
    allLocations.EntityData.Leafs["wds-client-reported-status"] = types.YLeaf{"WdsClientReportedStatus", allLocations.WdsClientReportedStatus}
    return &(allLocations.EntityData)
}

