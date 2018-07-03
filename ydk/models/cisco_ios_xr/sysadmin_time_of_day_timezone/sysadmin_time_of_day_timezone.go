// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG definitions
// for Cisco IOS-XR syadmin TOD configuration and cli.
// 
// This module contains definitions
// for the following management objects:
// Time of the Day(TOD) Cli and configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2017 by Cisco Systems, Inc.
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

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Children.Append("timezone", types.YChild{"Timezone", &clock.Timezone})
    clock.EntityData.Leafs = types.NewOrderedMap()

    clock.EntityData.YListKeys = []string {}

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

    timezone.EntityData.Children = types.NewOrderedMap()
    timezone.EntityData.Leafs = types.NewOrderedMap()
    timezone.EntityData.Leafs.Append("tzname", types.YLeaf{"Tzname", timezone.Tzname})
    timezone.EntityData.Leafs.Append("area", types.YLeaf{"Area", timezone.Area})

    timezone.EntityData.YListKeys = []string {}

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

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Children.Append("timezone_config", types.YChild{"TimezoneConfig", &trace.TimezoneConfig})
    trace.EntityData.Children.Append("timezone_notify", types.YChild{"TimezoneNotify", &trace.TimezoneNotify})
    trace.EntityData.Leafs = types.NewOrderedMap()

    trace.EntityData.YListKeys = []string {}

    return &(trace.EntityData)
}

// Trace_TimezoneConfig
type Trace_TimezoneConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Trace_TimezoneConfig_Trace.
    Trace []*Trace_TimezoneConfig_Trace
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

    timezoneConfig.EntityData.Children = types.NewOrderedMap()
    timezoneConfig.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range timezoneConfig.Trace {
        timezoneConfig.EntityData.Children.Append(types.GetSegmentPath(timezoneConfig.Trace[i]), types.YChild{"Trace", timezoneConfig.Trace[i]})
    }
    timezoneConfig.EntityData.Leafs = types.NewOrderedMap()

    timezoneConfig.EntityData.YListKeys = []string {}

    return &(timezoneConfig.EntityData)
}

// Trace_TimezoneConfig_Trace
// show traceable processes
type Trace_TimezoneConfig_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Trace_TimezoneConfig_Trace_Location.
    Location []*Trace_TimezoneConfig_Trace_Location
}

func (trace *Trace_TimezoneConfig_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "timezone_config"
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

// Trace_TimezoneConfig_Trace_Location
type Trace_TimezoneConfig_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Trace_TimezoneConfig_Trace_Location_AllOptions.
    AllOptions []*Trace_TimezoneConfig_Trace_Location_AllOptions
}

func (location *Trace_TimezoneConfig_Trace_Location) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneConfig_Trace_Location_AllOptions
type Trace_TimezoneConfig_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // Trace_TimezoneConfig_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Trace_TimezoneConfig_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Trace_TimezoneConfig_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneConfig_Trace_Location_AllOptions_TraceBlocks
type Trace_TimezoneConfig_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Trace_TimezoneConfig_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneNotify
type Trace_TimezoneNotify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of Trace_TimezoneNotify_Trace.
    Trace []*Trace_TimezoneNotify_Trace
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

    timezoneNotify.EntityData.Children = types.NewOrderedMap()
    timezoneNotify.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range timezoneNotify.Trace {
        timezoneNotify.EntityData.Children.Append(types.GetSegmentPath(timezoneNotify.Trace[i]), types.YChild{"Trace", timezoneNotify.Trace[i]})
    }
    timezoneNotify.EntityData.Leafs = types.NewOrderedMap()

    timezoneNotify.EntityData.YListKeys = []string {}

    return &(timezoneNotify.EntityData)
}

// Trace_TimezoneNotify_Trace
// show traceable processes
type Trace_TimezoneNotify_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Trace_TimezoneNotify_Trace_Location.
    Location []*Trace_TimezoneNotify_Trace_Location
}

func (trace *Trace_TimezoneNotify_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "timezone_notify"
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

// Trace_TimezoneNotify_Trace_Location
type Trace_TimezoneNotify_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of Trace_TimezoneNotify_Trace_Location_AllOptions.
    AllOptions []*Trace_TimezoneNotify_Trace_Location_AllOptions
}

func (location *Trace_TimezoneNotify_Trace_Location) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneNotify_Trace_Location_AllOptions
type Trace_TimezoneNotify_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // Trace_TimezoneNotify_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*Trace_TimezoneNotify_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Trace_TimezoneNotify_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
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

// Trace_TimezoneNotify_Trace_Location_AllOptions_TraceBlocks
type Trace_TimezoneNotify_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Trace_TimezoneNotify_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
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

