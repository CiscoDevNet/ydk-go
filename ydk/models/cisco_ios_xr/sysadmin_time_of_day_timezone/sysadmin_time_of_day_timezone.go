// This module contains a collection of YANG definitions
// for Cisco IOS-XR syadmin TOD configuration and cli.
// 
// This module contains definitions
// for the following management objects:
// Time of the Day(TOD) Cli and configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_time_of_day_timezone

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_time_of_day_timezone"))
    ydk.RegisterEntity("{http://cisco.com/calvados/Cisco-IOS-XR-sysadmin-time-of-day-timezone clock}", reflect.TypeOf(Clock{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-time-of-day-timezone:clock", reflect.TypeOf(Clock{}))
    ydk.RegisterEntity("{http://cisco.com/calvados/Cisco-IOS-XR-sysadmin-time-of-day-timezone trace}", reflect.TypeOf(Trace{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-time-of-day-timezone:trace", reflect.TypeOf(Trace{}))
}

// Clock
type Clock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Timezone Clock_Timezone
}

func (clock *Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-time-of-day-timezone"
    clock.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-time-of-day-timezone:clock"
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = make(map[string]types.YChild)
    clock.EntityData.Children["timezone"] = types.YChild{"Timezone", &clock.Timezone}
    clock.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clock.EntityData)
}

// Clock_Timezone
type Clock_Timezone struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Tzname interface{}

    // The type is string.
    Area interface{}
}

func (timezone *Clock_Timezone) GetEntityData() *types.CommonEntityData {
    timezone.EntityData.YFilter = timezone.YFilter
    timezone.EntityData.YangName = "timezone"
    timezone.EntityData.BundleName = "cisco_ios_xr"
    timezone.EntityData.ParentYangName = "clock"
    timezone.EntityData.SegmentPath = "timezone"
    timezone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timezone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timezone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timezone.EntityData.Children = make(map[string]types.YChild)
    timezone.EntityData.Leafs = make(map[string]types.YLeaf)
    timezone.EntityData.Leafs["tzname"] = types.YLeaf{"Tzname", timezone.Tzname}
    timezone.EntityData.Leafs["area"] = types.YLeaf{"Area", timezone.Area}
    return &(timezone.EntityData)
}

// Trace
type Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    TimezoneConfig Trace_TimezoneConfig

    
    TimezoneNotify Trace_TimezoneNotify
}

func (trace *Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-time-of-day-timezone"
    trace.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-time-of-day-timezone:trace"
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = make(map[string]types.YChild)
    trace.EntityData.Children["timezone_config"] = types.YChild{"TimezoneConfig", &trace.TimezoneConfig}
    trace.EntityData.Children["timezone_notify"] = types.YChild{"TimezoneNotify", &trace.TimezoneNotify}
    trace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trace.EntityData)
}

// Trace_TimezoneConfig
type Trace_TimezoneConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Trace_TimezoneConfig_Trace.
    Trace []Trace_TimezoneConfig_Trace_
}

func (timezoneConfig *Trace_TimezoneConfig) GetEntityData() *types.CommonEntityData {
    timezoneConfig.EntityData.YFilter = timezoneConfig.YFilter
    timezoneConfig.EntityData.YangName = "timezone_config"
    timezoneConfig.EntityData.BundleName = "cisco_ios_xr"
    timezoneConfig.EntityData.ParentYangName = "trace"
    timezoneConfig.EntityData.SegmentPath = "timezone_config"
    timezoneConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timezoneConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timezoneConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timezoneConfig.EntityData.Children = make(map[string]types.YChild)
    timezoneConfig.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range timezoneConfig.Trace {
        timezoneConfig.EntityData.Children[types.GetSegmentPath(&timezoneConfig.Trace[i])] = types.YChild{"Trace", &timezoneConfig.Trace[i]}
    }
    timezoneConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(timezoneConfig.EntityData)
}

// Trace_TimezoneConfig_Trace_
// show traceable processes
type Trace_TimezoneConfig_Trace_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Trace_TimezoneConfig_Trace__Location.
    Location []Trace_TimezoneConfig_Trace__Location
}

func (trace_ *Trace_TimezoneConfig_Trace_) GetEntityData() *types.CommonEntityData {
    trace_.EntityData.YFilter = trace_.YFilter
    trace_.EntityData.YangName = "trace"
    trace_.EntityData.BundleName = "cisco_ios_xr"
    trace_.EntityData.ParentYangName = "timezone_config"
    trace_.EntityData.SegmentPath = "trace" + "[buffer='" + fmt.Sprintf("%v", trace_.Buffer) + "']"
    trace_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace_.EntityData.Children = make(map[string]types.YChild)
    trace_.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range trace_.Location {
        trace_.EntityData.Children[types.GetSegmentPath(&trace_.Location[i])] = types.YChild{"Location", &trace_.Location[i]}
    }
    trace_.EntityData.Leafs = make(map[string]types.YLeaf)
    trace_.EntityData.Leafs["buffer"] = types.YLeaf{"Buffer", trace_.Buffer}
    return &(trace_.EntityData)
}

// Trace_TimezoneConfig_Trace__Location
type Trace_TimezoneConfig_Trace__Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Trace_TimezoneConfig_Trace__Location_AllOptions.
    AllOptions []Trace_TimezoneConfig_Trace__Location_AllOptions
}

func (location *Trace_TimezoneConfig_Trace__Location) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneConfig_Trace__Location_AllOptions
type Trace_TimezoneConfig_Trace__Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // Trace_TimezoneConfig_Trace__Location_AllOptions_TraceBlocks.
    TraceBlocks []Trace_TimezoneConfig_Trace__Location_AllOptions_TraceBlocks
}

func (allOptions *Trace_TimezoneConfig_Trace__Location_AllOptions) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneConfig_Trace__Location_AllOptions_TraceBlocks
type Trace_TimezoneConfig_Trace__Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Trace_TimezoneConfig_Trace__Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneNotify
type Trace_TimezoneNotify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Trace_TimezoneNotify_Trace.
    Trace []Trace_TimezoneNotify_Trace_
}

func (timezoneNotify *Trace_TimezoneNotify) GetEntityData() *types.CommonEntityData {
    timezoneNotify.EntityData.YFilter = timezoneNotify.YFilter
    timezoneNotify.EntityData.YangName = "timezone_notify"
    timezoneNotify.EntityData.BundleName = "cisco_ios_xr"
    timezoneNotify.EntityData.ParentYangName = "trace"
    timezoneNotify.EntityData.SegmentPath = "timezone_notify"
    timezoneNotify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timezoneNotify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timezoneNotify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timezoneNotify.EntityData.Children = make(map[string]types.YChild)
    timezoneNotify.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range timezoneNotify.Trace {
        timezoneNotify.EntityData.Children[types.GetSegmentPath(&timezoneNotify.Trace[i])] = types.YChild{"Trace", &timezoneNotify.Trace[i]}
    }
    timezoneNotify.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(timezoneNotify.EntityData)
}

// Trace_TimezoneNotify_Trace_
// show traceable processes
type Trace_TimezoneNotify_Trace_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Trace_TimezoneNotify_Trace__Location.
    Location []Trace_TimezoneNotify_Trace__Location
}

func (trace_ *Trace_TimezoneNotify_Trace_) GetEntityData() *types.CommonEntityData {
    trace_.EntityData.YFilter = trace_.YFilter
    trace_.EntityData.YangName = "trace"
    trace_.EntityData.BundleName = "cisco_ios_xr"
    trace_.EntityData.ParentYangName = "timezone_notify"
    trace_.EntityData.SegmentPath = "trace" + "[buffer='" + fmt.Sprintf("%v", trace_.Buffer) + "']"
    trace_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace_.EntityData.Children = make(map[string]types.YChild)
    trace_.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range trace_.Location {
        trace_.EntityData.Children[types.GetSegmentPath(&trace_.Location[i])] = types.YChild{"Location", &trace_.Location[i]}
    }
    trace_.EntityData.Leafs = make(map[string]types.YLeaf)
    trace_.EntityData.Leafs["buffer"] = types.YLeaf{"Buffer", trace_.Buffer}
    return &(trace_.EntityData)
}

// Trace_TimezoneNotify_Trace__Location
type Trace_TimezoneNotify_Trace__Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Trace_TimezoneNotify_Trace__Location_AllOptions.
    AllOptions []Trace_TimezoneNotify_Trace__Location_AllOptions
}

func (location *Trace_TimezoneNotify_Trace__Location) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneNotify_Trace__Location_AllOptions
type Trace_TimezoneNotify_Trace__Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // Trace_TimezoneNotify_Trace__Location_AllOptions_TraceBlocks.
    TraceBlocks []Trace_TimezoneNotify_Trace__Location_AllOptions_TraceBlocks
}

func (allOptions *Trace_TimezoneNotify_Trace__Location_AllOptions) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneNotify_Trace__Location_AllOptions_TraceBlocks
type Trace_TimezoneNotify_Trace__Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Trace_TimezoneNotify_Trace__Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
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

