// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Alarm management YANG model
// 
// Copyright(c) 2011-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_alarm_mgr

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_alarm_mgr"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-alarm-mgr alarm_mgr}", reflect.TypeOf(AlarmMgr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr", reflect.TypeOf(AlarmMgr{}))
}

// SeverityTd represents The severity enumeration type of an alarm 
type SeverityTd string

const (
    SeverityTd_unknown SeverityTd = "unknown"

    SeverityTd_not_reported SeverityTd = "not_reported"

    SeverityTd_not_alarmed SeverityTd = "not_alarmed"

    SeverityTd_minor SeverityTd = "minor"

    SeverityTd_major SeverityTd = "major"

    SeverityTd_critical SeverityTd = "critical"
)

// StatusTd represents The status enumeration type of an alarm 
type StatusTd string

const (
    StatusTd_unknown StatusTd = "unknown"

    StatusTd_set StatusTd = "set"

    StatusTd_clear StatusTd = "clear"

    StatusTd_suppress StatusTd = "suppress"
)

// GroupTd represents The group enumeration type of an alarm 
type GroupTd string

const (
    GroupTd_unknown GroupTd = "unknown"

    GroupTd_environ GroupTd = "environ"

    GroupTd_ethernet GroupTd = "ethernet"

    GroupTd_fabric GroupTd = "fabric"

    GroupTd_power GroupTd = "power"

    GroupTd_software GroupTd = "software"

    GroupTd_slice GroupTd = "slice"

    GroupTd_cpu GroupTd = "cpu"

    GroupTd_controller GroupTd = "controller"

    GroupTd_sonet GroupTd = "sonet"

    GroupTd_otn GroupTd = "otn"

    GroupTd_sdh_controller GroupTd = "sdh_controller"

    GroupTd_asic GroupTd = "asic"

    GroupTd_fpd_infra GroupTd = "fpd_infra"

    GroupTd_shelf GroupTd = "shelf"

    GroupTd_mpa GroupTd = "mpa"

    GroupTd_ots GroupTd = "ots"

    GroupTd_timing GroupTd = "timing"

    GroupTd_last GroupTd = "last"
)

// AgentStateTd
type AgentStateTd string

const (
    AgentStateTd_start AgentStateTd = "start"

    AgentStateTd_init AgentStateTd = "init"

    AgentStateTd_connecting AgentStateTd = "connecting"

    AgentStateTd_connected AgentStateTd = "connected"

    AgentStateTd_registered AgentStateTd = "registered"

    AgentStateTd_disconnected AgentStateTd = "disconnected"
)

// AgentTypeTd
type AgentTypeTd string

const (
    AgentTypeTd_unknown AgentTypeTd = "unknown"

    AgentTypeTd_producer AgentTypeTd = "producer"

    AgentTypeTd_consumer AgentTypeTd = "consumer"

    AgentTypeTd_subscriber AgentTypeTd = "subscriber"
)

// AlarmMgr
// Calvados alarms operational data model
type AlarmMgr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of AlarmMgr_Trace.
    Trace []*AlarmMgr_Trace

    // A set of brief alarm commands.
    Brief AlarmMgr_Brief

    // A set of detail alarm commands.
    Detail AlarmMgr_Detail
}

func (alarmMgr *AlarmMgr) GetEntityData() *types.CommonEntityData {
    alarmMgr.EntityData.YFilter = alarmMgr.YFilter
    alarmMgr.EntityData.YangName = "alarm_mgr"
    alarmMgr.EntityData.BundleName = "cisco_ios_xr"
    alarmMgr.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-alarm-mgr"
    alarmMgr.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr"
    alarmMgr.EntityData.AbsolutePath = alarmMgr.EntityData.SegmentPath
    alarmMgr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmMgr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmMgr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmMgr.EntityData.Children = types.NewOrderedMap()
    alarmMgr.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range alarmMgr.Trace {
        alarmMgr.EntityData.Children.Append(types.GetSegmentPath(alarmMgr.Trace[i]), types.YChild{"Trace", alarmMgr.Trace[i]})
    }
    alarmMgr.EntityData.Children.Append("brief", types.YChild{"Brief", &alarmMgr.Brief})
    alarmMgr.EntityData.Children.Append("detail", types.YChild{"Detail", &alarmMgr.Detail})
    alarmMgr.EntityData.Leafs = types.NewOrderedMap()

    alarmMgr.EntityData.YListKeys = []string {}

    return &(alarmMgr.EntityData)
}

// AlarmMgr_Trace
// show traceable processes
type AlarmMgr_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of AlarmMgr_Trace_Location.
    Location []*AlarmMgr_Trace_Location
}

func (trace *AlarmMgr_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "alarm_mgr"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/" + trace.EntityData.SegmentPath
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

// AlarmMgr_Trace_Location
type AlarmMgr_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of AlarmMgr_Trace_Location_AllOptions.
    AllOptions []*AlarmMgr_Trace_Location_AllOptions
}

func (location *AlarmMgr_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/trace/" + location.EntityData.SegmentPath
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

// AlarmMgr_Trace_Location_AllOptions
type AlarmMgr_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of AlarmMgr_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*AlarmMgr_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *AlarmMgr_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/trace/location/" + allOptions.EntityData.SegmentPath
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

// AlarmMgr_Trace_Location_AllOptions_TraceBlocks
type AlarmMgr_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *AlarmMgr_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

// AlarmMgr_Brief
// A set of brief alarm commands
type AlarmMgr_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarms reported at the local card as  specified by the location parameter.
    Card AlarmMgr_Brief_Card

    // Alarms reported at the rack scope    specified by the rack-id.
    Rack AlarmMgr_Brief_Rack

    // Alarms reported at the system scope.
    System AlarmMgr_Brief_System
}

func (brief *AlarmMgr_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "alarm_mgr"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/" + brief.EntityData.SegmentPath
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Children.Append("card", types.YChild{"Card", &brief.Card})
    brief.EntityData.Children.Append("rack", types.YChild{"Rack", &brief.Rack})
    brief.EntityData.Children.Append("system", types.YChild{"System", &brief.System})
    brief.EntityData.Leafs = types.NewOrderedMap()

    brief.EntityData.YListKeys = []string {}

    return &(brief.EntityData)
}

// AlarmMgr_Brief_Card
// Alarms reported at the local card as 
// specified by the location parameter
type AlarmMgr_Brief_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AlarmMgr_Brief_Card_Location.
    Location []*AlarmMgr_Brief_Card_Location
}

func (card *AlarmMgr_Brief_Card) GetEntityData() *types.CommonEntityData {
    card.EntityData.YFilter = card.YFilter
    card.EntityData.YangName = "card"
    card.EntityData.BundleName = "cisco_ios_xr"
    card.EntityData.ParentYangName = "brief"
    card.EntityData.SegmentPath = "card"
    card.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/" + card.EntityData.SegmentPath
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

// AlarmMgr_Brief_Card_Location
type AlarmMgr_Brief_Card_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'((([bB][0-9])/(([a-zA-Z]){2}\\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    Locations interface{}

    // The type is slice of AlarmMgr_Brief_Card_Location_Active.
    Active []*AlarmMgr_Brief_Card_Location_Active

    // The type is slice of AlarmMgr_Brief_Card_Location_History.
    History []*AlarmMgr_Brief_Card_Location_History

    // The type is slice of AlarmMgr_Brief_Card_Location_Suppressed.
    Suppressed []*AlarmMgr_Brief_Card_Location_Suppressed
}

func (location *AlarmMgr_Brief_Card_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "card"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Locations, "locations")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/card/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("active", types.YChild{"Active", nil})
    for i := range location.Active {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Active[i]), types.YChild{"Active", location.Active[i]})
    }
    location.EntityData.Children.Append("history", types.YChild{"History", nil})
    for i := range location.History {
        location.EntityData.Children.Append(types.GetSegmentPath(location.History[i]), types.YChild{"History", location.History[i]})
    }
    location.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", nil})
    for i := range location.Suppressed {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Suppressed[i]), types.YChild{"Suppressed", location.Suppressed[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("locations", types.YLeaf{"Locations", location.Locations})

    location.EntityData.YListKeys = []string {"Locations"}

    return &(location.EntityData)
}

// AlarmMgr_Brief_Card_Location_Active
type AlarmMgr_Brief_Card_Location_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}
}

func (active *AlarmMgr_Brief_Card_Location_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "location"
    active.EntityData.SegmentPath = "active" + types.AddKeyToken(active.Aid, "aid") + types.AddKeyToken(active.Eid, "eid")
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/card/location/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", active.Aid})
    active.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", active.Eid})
    active.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", active.Tag})
    active.EntityData.Leafs.Append("module", types.YLeaf{"Module", active.Module})
    active.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", active.GenLocation})
    active.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", active.Severity})
    active.EntityData.Leafs.Append("group", types.YLeaf{"Group", active.Group})
    active.EntityData.Leafs.Append("description", types.YLeaf{"Description", active.Description})
    active.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", active.SetTime})

    active.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(active.EntityData)
}

// AlarmMgr_Brief_Card_Location_History
type AlarmMgr_Brief_Card_Location_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // Alarm clear time stamp. The type is string.
    ClearTime interface{}
}

func (history *AlarmMgr_Brief_Card_Location_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "location"
    history.EntityData.SegmentPath = "history" + types.AddKeyToken(history.Aid, "aid") + types.AddKeyToken(history.Eid, "eid")
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/card/location/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", history.Aid})
    history.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", history.Eid})
    history.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", history.Tag})
    history.EntityData.Leafs.Append("module", types.YLeaf{"Module", history.Module})
    history.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", history.GenLocation})
    history.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", history.Severity})
    history.EntityData.Leafs.Append("group", types.YLeaf{"Group", history.Group})
    history.EntityData.Leafs.Append("description", types.YLeaf{"Description", history.Description})
    history.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", history.SetTime})
    history.EntityData.Leafs.Append("clear_time", types.YLeaf{"ClearTime", history.ClearTime})

    history.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(history.EntityData)
}

// AlarmMgr_Brief_Card_Location_Suppressed
type AlarmMgr_Brief_Card_Location_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // Alarm suppressed time stamp. The type is string.
    SuppressedTime interface{}
}

func (suppressed *AlarmMgr_Brief_Card_Location_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "location"
    suppressed.EntityData.SegmentPath = "suppressed" + types.AddKeyToken(suppressed.Aid, "aid") + types.AddKeyToken(suppressed.Eid, "eid")
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/card/location/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Leafs = types.NewOrderedMap()
    suppressed.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", suppressed.Aid})
    suppressed.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", suppressed.Eid})
    suppressed.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", suppressed.Tag})
    suppressed.EntityData.Leafs.Append("module", types.YLeaf{"Module", suppressed.Module})
    suppressed.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", suppressed.GenLocation})
    suppressed.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressed.Severity})
    suppressed.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressed.Group})
    suppressed.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressed.Description})
    suppressed.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", suppressed.SetTime})
    suppressed.EntityData.Leafs.Append("suppressed_time", types.YLeaf{"SuppressedTime", suppressed.SuppressedTime})

    suppressed.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(suppressed.EntityData)
}

// AlarmMgr_Brief_Rack
// Alarms reported at the rack scope   
// specified by the rack-id
type AlarmMgr_Brief_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AlarmMgr_Brief_Rack_RackLocations.
    RackLocations []*AlarmMgr_Brief_Rack_RackLocations
}

func (rack *AlarmMgr_Brief_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "brief"
    rack.EntityData.SegmentPath = "rack"
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("rack_locations", types.YChild{"RackLocations", nil})
    for i := range rack.RackLocations {
        rack.EntityData.Children.Append(types.GetSegmentPath(rack.RackLocations[i]), types.YChild{"RackLocations", rack.RackLocations[i]})
    }
    rack.EntityData.Leafs = types.NewOrderedMap()

    rack.EntityData.YListKeys = []string {}

    return &(rack.EntityData)
}

// AlarmMgr_Brief_Rack_RackLocations
type AlarmMgr_Brief_Rack_RackLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Rackid interface{}

    // The type is slice of AlarmMgr_Brief_Rack_RackLocations_Active.
    Active []*AlarmMgr_Brief_Rack_RackLocations_Active

    // The type is slice of AlarmMgr_Brief_Rack_RackLocations_History.
    History []*AlarmMgr_Brief_Rack_RackLocations_History

    // The type is slice of AlarmMgr_Brief_Rack_RackLocations_Suppressed.
    Suppressed []*AlarmMgr_Brief_Rack_RackLocations_Suppressed
}

func (rackLocations *AlarmMgr_Brief_Rack_RackLocations) GetEntityData() *types.CommonEntityData {
    rackLocations.EntityData.YFilter = rackLocations.YFilter
    rackLocations.EntityData.YangName = "rack_locations"
    rackLocations.EntityData.BundleName = "cisco_ios_xr"
    rackLocations.EntityData.ParentYangName = "rack"
    rackLocations.EntityData.SegmentPath = "rack_locations" + types.AddKeyToken(rackLocations.Rackid, "rackid")
    rackLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/rack/" + rackLocations.EntityData.SegmentPath
    rackLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLocations.EntityData.Children = types.NewOrderedMap()
    rackLocations.EntityData.Children.Append("active", types.YChild{"Active", nil})
    for i := range rackLocations.Active {
        rackLocations.EntityData.Children.Append(types.GetSegmentPath(rackLocations.Active[i]), types.YChild{"Active", rackLocations.Active[i]})
    }
    rackLocations.EntityData.Children.Append("history", types.YChild{"History", nil})
    for i := range rackLocations.History {
        rackLocations.EntityData.Children.Append(types.GetSegmentPath(rackLocations.History[i]), types.YChild{"History", rackLocations.History[i]})
    }
    rackLocations.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", nil})
    for i := range rackLocations.Suppressed {
        rackLocations.EntityData.Children.Append(types.GetSegmentPath(rackLocations.Suppressed[i]), types.YChild{"Suppressed", rackLocations.Suppressed[i]})
    }
    rackLocations.EntityData.Leafs = types.NewOrderedMap()
    rackLocations.EntityData.Leafs.Append("rackid", types.YLeaf{"Rackid", rackLocations.Rackid})

    rackLocations.EntityData.YListKeys = []string {"Rackid"}

    return &(rackLocations.EntityData)
}

// AlarmMgr_Brief_Rack_RackLocations_Active
type AlarmMgr_Brief_Rack_RackLocations_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}
}

func (active *AlarmMgr_Brief_Rack_RackLocations_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "rack_locations"
    active.EntityData.SegmentPath = "active" + types.AddKeyToken(active.Aid, "aid") + types.AddKeyToken(active.Eid, "eid")
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/rack/rack_locations/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", active.Aid})
    active.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", active.Eid})
    active.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", active.Tag})
    active.EntityData.Leafs.Append("module", types.YLeaf{"Module", active.Module})
    active.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", active.GenLocation})
    active.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", active.Severity})
    active.EntityData.Leafs.Append("group", types.YLeaf{"Group", active.Group})
    active.EntityData.Leafs.Append("description", types.YLeaf{"Description", active.Description})
    active.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", active.SetTime})

    active.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(active.EntityData)
}

// AlarmMgr_Brief_Rack_RackLocations_History
type AlarmMgr_Brief_Rack_RackLocations_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // Alarm clear time stamp. The type is string.
    ClearTime interface{}
}

func (history *AlarmMgr_Brief_Rack_RackLocations_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "rack_locations"
    history.EntityData.SegmentPath = "history" + types.AddKeyToken(history.Aid, "aid") + types.AddKeyToken(history.Eid, "eid")
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/rack/rack_locations/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", history.Aid})
    history.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", history.Eid})
    history.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", history.Tag})
    history.EntityData.Leafs.Append("module", types.YLeaf{"Module", history.Module})
    history.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", history.GenLocation})
    history.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", history.Severity})
    history.EntityData.Leafs.Append("group", types.YLeaf{"Group", history.Group})
    history.EntityData.Leafs.Append("description", types.YLeaf{"Description", history.Description})
    history.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", history.SetTime})
    history.EntityData.Leafs.Append("clear_time", types.YLeaf{"ClearTime", history.ClearTime})

    history.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(history.EntityData)
}

// AlarmMgr_Brief_Rack_RackLocations_Suppressed
type AlarmMgr_Brief_Rack_RackLocations_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // Alarm suppressed time stamp. The type is string.
    SuppressedTime interface{}
}

func (suppressed *AlarmMgr_Brief_Rack_RackLocations_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "rack_locations"
    suppressed.EntityData.SegmentPath = "suppressed" + types.AddKeyToken(suppressed.Aid, "aid") + types.AddKeyToken(suppressed.Eid, "eid")
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/rack/rack_locations/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Leafs = types.NewOrderedMap()
    suppressed.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", suppressed.Aid})
    suppressed.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", suppressed.Eid})
    suppressed.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", suppressed.Tag})
    suppressed.EntityData.Leafs.Append("module", types.YLeaf{"Module", suppressed.Module})
    suppressed.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", suppressed.GenLocation})
    suppressed.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressed.Severity})
    suppressed.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressed.Group})
    suppressed.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressed.Description})
    suppressed.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", suppressed.SetTime})
    suppressed.EntityData.Leafs.Append("suppressed_time", types.YLeaf{"SuppressedTime", suppressed.SuppressedTime})

    suppressed.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(suppressed.EntityData)
}

// AlarmMgr_Brief_System
// Alarms reported at the system scope
type AlarmMgr_Brief_System struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AlarmMgr_Brief_System_Active.
    Active []*AlarmMgr_Brief_System_Active

    // The type is slice of AlarmMgr_Brief_System_History.
    History []*AlarmMgr_Brief_System_History

    // The type is slice of AlarmMgr_Brief_System_Suppressed.
    Suppressed []*AlarmMgr_Brief_System_Suppressed
}

func (system *AlarmMgr_Brief_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xr"
    system.EntityData.ParentYangName = "brief"
    system.EntityData.SegmentPath = "system"
    system.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/" + system.EntityData.SegmentPath
    system.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    system.EntityData.Children = types.NewOrderedMap()
    system.EntityData.Children.Append("active", types.YChild{"Active", nil})
    for i := range system.Active {
        system.EntityData.Children.Append(types.GetSegmentPath(system.Active[i]), types.YChild{"Active", system.Active[i]})
    }
    system.EntityData.Children.Append("history", types.YChild{"History", nil})
    for i := range system.History {
        system.EntityData.Children.Append(types.GetSegmentPath(system.History[i]), types.YChild{"History", system.History[i]})
    }
    system.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", nil})
    for i := range system.Suppressed {
        system.EntityData.Children.Append(types.GetSegmentPath(system.Suppressed[i]), types.YChild{"Suppressed", system.Suppressed[i]})
    }
    system.EntityData.Leafs = types.NewOrderedMap()

    system.EntityData.YListKeys = []string {}

    return &(system.EntityData)
}

// AlarmMgr_Brief_System_Active
type AlarmMgr_Brief_System_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}
}

func (active *AlarmMgr_Brief_System_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "system"
    active.EntityData.SegmentPath = "active" + types.AddKeyToken(active.Aid, "aid") + types.AddKeyToken(active.Eid, "eid")
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/system/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", active.Aid})
    active.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", active.Eid})
    active.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", active.Tag})
    active.EntityData.Leafs.Append("module", types.YLeaf{"Module", active.Module})
    active.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", active.GenLocation})
    active.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", active.Severity})
    active.EntityData.Leafs.Append("group", types.YLeaf{"Group", active.Group})
    active.EntityData.Leafs.Append("description", types.YLeaf{"Description", active.Description})
    active.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", active.SetTime})

    active.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(active.EntityData)
}

// AlarmMgr_Brief_System_History
type AlarmMgr_Brief_System_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // Alarm clear time stamp. The type is string.
    ClearTime interface{}
}

func (history *AlarmMgr_Brief_System_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "system"
    history.EntityData.SegmentPath = "history" + types.AddKeyToken(history.Aid, "aid") + types.AddKeyToken(history.Eid, "eid")
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/system/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", history.Aid})
    history.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", history.Eid})
    history.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", history.Tag})
    history.EntityData.Leafs.Append("module", types.YLeaf{"Module", history.Module})
    history.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", history.GenLocation})
    history.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", history.Severity})
    history.EntityData.Leafs.Append("group", types.YLeaf{"Group", history.Group})
    history.EntityData.Leafs.Append("description", types.YLeaf{"Description", history.Description})
    history.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", history.SetTime})
    history.EntityData.Leafs.Append("clear_time", types.YLeaf{"ClearTime", history.ClearTime})

    history.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(history.EntityData)
}

// AlarmMgr_Brief_System_Suppressed
type AlarmMgr_Brief_System_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // Alarm suppressed time stamp. The type is string.
    SuppressedTime interface{}
}

func (suppressed *AlarmMgr_Brief_System_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "system"
    suppressed.EntityData.SegmentPath = "suppressed" + types.AddKeyToken(suppressed.Aid, "aid") + types.AddKeyToken(suppressed.Eid, "eid")
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/brief/system/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Leafs = types.NewOrderedMap()
    suppressed.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", suppressed.Aid})
    suppressed.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", suppressed.Eid})
    suppressed.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", suppressed.Tag})
    suppressed.EntityData.Leafs.Append("module", types.YLeaf{"Module", suppressed.Module})
    suppressed.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", suppressed.GenLocation})
    suppressed.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressed.Severity})
    suppressed.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressed.Group})
    suppressed.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressed.Description})
    suppressed.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", suppressed.SetTime})
    suppressed.EntityData.Leafs.Append("suppressed_time", types.YLeaf{"SuppressedTime", suppressed.SuppressedTime})

    suppressed.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(suppressed.EntityData)
}

// AlarmMgr_Detail
// A set of detail alarm commands
type AlarmMgr_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarms reported at the local card as  specified by the location parameter.
    Card AlarmMgr_Detail_Card

    // Alarms reported at the rack as  specified by the location parameter.
    Rack AlarmMgr_Detail_Rack

    // Alarms reported at the system as  specified by the location parameter.
    System AlarmMgr_Detail_System
}

func (detail *AlarmMgr_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "alarm_mgr"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("card", types.YChild{"Card", &detail.Card})
    detail.EntityData.Children.Append("rack", types.YChild{"Rack", &detail.Rack})
    detail.EntityData.Children.Append("system", types.YChild{"System", &detail.System})
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// AlarmMgr_Detail_Card
// Alarms reported at the local card as 
// specified by the location parameter
type AlarmMgr_Detail_Card struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AlarmMgr_Detail_Card_Location.
    Location []*AlarmMgr_Detail_Card_Location
}

func (card *AlarmMgr_Detail_Card) GetEntityData() *types.CommonEntityData {
    card.EntityData.YFilter = card.YFilter
    card.EntityData.YangName = "card"
    card.EntityData.BundleName = "cisco_ios_xr"
    card.EntityData.ParentYangName = "detail"
    card.EntityData.SegmentPath = "card"
    card.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/" + card.EntityData.SegmentPath
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

// AlarmMgr_Detail_Card_Location
type AlarmMgr_Detail_Card_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'((([bB][0-9])/(([a-zA-Z]){2}\\d{1,2}))|(([fF][0-7])/(([a-zA-Z]){2}\\d{1,2}))|((0?[0-9]|1[0-5])/((([a-zA-Z]){2,3})?\\d{1,2})))(/[cC][pP][uU]0)?'.
    Locations interface{}

    // The type is slice of AlarmMgr_Detail_Card_Location_Active.
    Active []*AlarmMgr_Detail_Card_Location_Active

    // The type is slice of AlarmMgr_Detail_Card_Location_History.
    History []*AlarmMgr_Detail_Card_Location_History

    // The type is slice of AlarmMgr_Detail_Card_Location_Stats.
    Stats []*AlarmMgr_Detail_Card_Location_Stats

    // The type is slice of AlarmMgr_Detail_Card_Location_Clients.
    Clients []*AlarmMgr_Detail_Card_Location_Clients

    // The type is slice of AlarmMgr_Detail_Card_Location_Suppressed.
    Suppressed []*AlarmMgr_Detail_Card_Location_Suppressed
}

func (location *AlarmMgr_Detail_Card_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "card"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Locations, "locations")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/card/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("active", types.YChild{"Active", nil})
    for i := range location.Active {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Active[i]), types.YChild{"Active", location.Active[i]})
    }
    location.EntityData.Children.Append("history", types.YChild{"History", nil})
    for i := range location.History {
        location.EntityData.Children.Append(types.GetSegmentPath(location.History[i]), types.YChild{"History", location.History[i]})
    }
    location.EntityData.Children.Append("stats", types.YChild{"Stats", nil})
    for i := range location.Stats {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Stats[i]), types.YChild{"Stats", location.Stats[i]})
    }
    location.EntityData.Children.Append("clients", types.YChild{"Clients", nil})
    for i := range location.Clients {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Clients[i]), types.YChild{"Clients", location.Clients[i]})
    }
    location.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", nil})
    for i := range location.Suppressed {
        location.EntityData.Children.Append(types.GetSegmentPath(location.Suppressed[i]), types.YChild{"Suppressed", location.Suppressed[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("locations", types.YLeaf{"Locations", location.Locations})

    location.EntityData.YListKeys = []string {"Locations"}

    return &(location.EntityData)
}

// AlarmMgr_Detail_Card_Location_Active
type AlarmMgr_Detail_Card_Location_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // The current state of the bi-state alarm. The type is StatusTd.
    State interface{}

    // The reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // The condition bit-flags of the alarm. The type is interface{} with range:
    // 0..4294967295.
    Resynced interface{}

    // Alarm detailed description. The type is string.
    DetailDesc interface{}

    // Alarm clear time stamp. The type is string.
    ClearTime interface{}
}

func (active *AlarmMgr_Detail_Card_Location_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "location"
    active.EntityData.SegmentPath = "active" + types.AddKeyToken(active.Aid, "aid") + types.AddKeyToken(active.Eid, "eid")
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/card/location/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", active.Aid})
    active.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", active.Eid})
    active.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", active.Tag})
    active.EntityData.Leafs.Append("module", types.YLeaf{"Module", active.Module})
    active.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", active.GenLocation})
    active.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", active.Severity})
    active.EntityData.Leafs.Append("group", types.YLeaf{"Group", active.Group})
    active.EntityData.Leafs.Append("description", types.YLeaf{"Description", active.Description})
    active.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", active.SetTime})
    active.EntityData.Leafs.Append("state", types.YLeaf{"State", active.State})
    active.EntityData.Leafs.Append("reporting_agent_id", types.YLeaf{"ReportingAgentId", active.ReportingAgentId})
    active.EntityData.Leafs.Append("resynced", types.YLeaf{"Resynced", active.Resynced})
    active.EntityData.Leafs.Append("detail_desc", types.YLeaf{"DetailDesc", active.DetailDesc})
    active.EntityData.Leafs.Append("clear_time", types.YLeaf{"ClearTime", active.ClearTime})

    active.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(active.EntityData)
}

// AlarmMgr_Detail_Card_Location_History
type AlarmMgr_Detail_Card_Location_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // The current state of the bi-state alarm. The type is StatusTd.
    State interface{}

    // The reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // The condition bit-flags of the alarm. The type is interface{} with range:
    // 0..4294967295.
    Resynced interface{}

    // Alarm detailed description. The type is string.
    DetailDesc interface{}

    // Alarm clear time stamp. The type is string.
    ClearTime interface{}
}

func (history *AlarmMgr_Detail_Card_Location_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "location"
    history.EntityData.SegmentPath = "history" + types.AddKeyToken(history.Aid, "aid") + types.AddKeyToken(history.Eid, "eid")
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/card/location/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", history.Aid})
    history.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", history.Eid})
    history.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", history.Tag})
    history.EntityData.Leafs.Append("module", types.YLeaf{"Module", history.Module})
    history.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", history.GenLocation})
    history.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", history.Severity})
    history.EntityData.Leafs.Append("group", types.YLeaf{"Group", history.Group})
    history.EntityData.Leafs.Append("description", types.YLeaf{"Description", history.Description})
    history.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", history.SetTime})
    history.EntityData.Leafs.Append("state", types.YLeaf{"State", history.State})
    history.EntityData.Leafs.Append("reporting_agent_id", types.YLeaf{"ReportingAgentId", history.ReportingAgentId})
    history.EntityData.Leafs.Append("resynced", types.YLeaf{"Resynced", history.Resynced})
    history.EntityData.Leafs.Append("detail_desc", types.YLeaf{"DetailDesc", history.DetailDesc})
    history.EntityData.Leafs.Append("clear_time", types.YLeaf{"ClearTime", history.ClearTime})

    history.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(history.EntityData)
}

// AlarmMgr_Detail_Card_Location_Stats
type AlarmMgr_Detail_Card_Location_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Alarms statistics at specified time. The type is
    // string.
    Attime interface{}

    // Total alarms reported into this service. The type is interface{} with
    // range: 0..18446744073709551615.
    Reported interface{}

    // Alarms dropped due to some error or other. The type is interface{} with
    // range: 0..18446744073709551615.
    Dropped interface{}

    // Total active alarms current in this service. The type is interface{} with
    // range: 0..18446744073709551615.
    BiSet interface{}

    // Alarms that are currently in the cleared state. The type is interface{}
    // with range: 0..18446744073709551615.
    BiClear interface{}

    // Alarms that are currently in the suppressed state. The type is interface{}
    // with range: 0..18446744073709551615.
    Suppressed interface{}

    // Alarms dropped due to invalid aid in the alarm report. The type is
    // interface{} with range: 0..18446744073709551615.
    DropInvAid interface{}

    // Alarms dropped due to low memory condition. The type is interface{} with
    // range: 0..18446744073709551615.
    DropNoMem interface{}

    // Alarms dropped due to database internal error. The type is interface{} with
    // range: 0..18446744073709551615.
    DropDbError interface{}

    // Alarms dropped due to clear without a set. The type is interface{} with
    // range: 0..18446744073709551615.
    DropClearNoSet interface{}

    // Alarms dropped due to duplicate alarm report. The type is interface{} with
    // range: 0..18446744073709551615.
    DropDup interface{}

    // Alarms present in the cached for show. The type is interface{} with range:
    // 0..18446744073709551615.
    CacheHit interface{}

    // Alarms not present in the cached for show. The type is interface{} with
    // range: 0..18446744073709551615.
    CacheMiss interface{}
}

func (stats *AlarmMgr_Detail_Card_Location_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "location"
    stats.EntityData.SegmentPath = "stats" + types.AddKeyToken(stats.Attime, "attime")
    stats.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/card/location/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("attime", types.YLeaf{"Attime", stats.Attime})
    stats.EntityData.Leafs.Append("reported", types.YLeaf{"Reported", stats.Reported})
    stats.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", stats.Dropped})
    stats.EntityData.Leafs.Append("bi_set", types.YLeaf{"BiSet", stats.BiSet})
    stats.EntityData.Leafs.Append("bi_clear", types.YLeaf{"BiClear", stats.BiClear})
    stats.EntityData.Leafs.Append("suppressed", types.YLeaf{"Suppressed", stats.Suppressed})
    stats.EntityData.Leafs.Append("drop_inv_aid", types.YLeaf{"DropInvAid", stats.DropInvAid})
    stats.EntityData.Leafs.Append("drop_no_mem", types.YLeaf{"DropNoMem", stats.DropNoMem})
    stats.EntityData.Leafs.Append("drop_db_error", types.YLeaf{"DropDbError", stats.DropDbError})
    stats.EntityData.Leafs.Append("drop_clear_no_set", types.YLeaf{"DropClearNoSet", stats.DropClearNoSet})
    stats.EntityData.Leafs.Append("drop_dup", types.YLeaf{"DropDup", stats.DropDup})
    stats.EntityData.Leafs.Append("cache_hit", types.YLeaf{"CacheHit", stats.CacheHit})
    stats.EntityData.Leafs.Append("cache_miss", types.YLeaf{"CacheMiss", stats.CacheMiss})

    stats.EntityData.YListKeys = []string {"Attime"}

    return &(stats.EntityData)
}

// AlarmMgr_Detail_Card_Location_Clients
type AlarmMgr_Detail_Card_Location_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The client handle through which interface. The
    // type is string.
    AgentHandle interface{}

    // Alarms client. The type is string.
    AgentName interface{}

    // Alarms agent id of the client. The type is interface{} with range:
    // 0..4294967295.
    AgentId interface{}

    // The location of this client. The type is string.
    AgentLocation interface{}

    // The current state of the client. The type is AgentStateTd.
    AgentState interface{}

    // The type of  the client. The type is AgentTypeTd.
    AgentType interface{}

    // The current subscription status of the client. The type is bool.
    AgentFilterDisp interface{}

    // The subscriber id of the client. The type is interface{} with range:
    // 0..4294967295.
    AgentSubsId interface{}

    // The filter used for alarm bi-state state. The type is StatusTd.
    AgentFilterState interface{}

    // The filter used for alarm severity. The type is SeverityTd.
    AgentFilterSeverity interface{}

    // The filter used for alarm group. The type is GroupTd.
    AgentFilterGroup interface{}

    // agent alarm_sm register with sdr_id in SOMT mode. The type is interface{}
    // with range: 0..4294967295.
    AgentSdrId interface{}

    // Number of times the agent connected to the alarm mgr. The type is
    // interface{} with range: 0..18446744073709551615.
    AgentConnectCount interface{}

    // Agent connect timestamp. The type is string.
    AgentConnectTime interface{}

    // Number of times the agent queried for alarms. The type is interface{} with
    // range: 0..18446744073709551615.
    AgentGetCount interface{}

    // Number of times the agent subscribed for alarms. The type is interface{}
    // with range: 0..18446744073709551615.
    AgentSubscribeCount interface{}

    // Number of times the agent reported alarms. The type is interface{} with
    // range: 0..18446744073709551615.
    AgentReportCount interface{}
}

func (clients *AlarmMgr_Detail_Card_Location_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "location"
    clients.EntityData.SegmentPath = "clients" + types.AddKeyToken(clients.AgentHandle, "agent_handle")
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/card/location/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Leafs = types.NewOrderedMap()
    clients.EntityData.Leafs.Append("agent_handle", types.YLeaf{"AgentHandle", clients.AgentHandle})
    clients.EntityData.Leafs.Append("agent_name", types.YLeaf{"AgentName", clients.AgentName})
    clients.EntityData.Leafs.Append("agent_id", types.YLeaf{"AgentId", clients.AgentId})
    clients.EntityData.Leafs.Append("agent_location", types.YLeaf{"AgentLocation", clients.AgentLocation})
    clients.EntityData.Leafs.Append("agent_state", types.YLeaf{"AgentState", clients.AgentState})
    clients.EntityData.Leafs.Append("agent_type", types.YLeaf{"AgentType", clients.AgentType})
    clients.EntityData.Leafs.Append("agent_filter_disp", types.YLeaf{"AgentFilterDisp", clients.AgentFilterDisp})
    clients.EntityData.Leafs.Append("agent_subs_id", types.YLeaf{"AgentSubsId", clients.AgentSubsId})
    clients.EntityData.Leafs.Append("agent_filter_state", types.YLeaf{"AgentFilterState", clients.AgentFilterState})
    clients.EntityData.Leafs.Append("agent_filter_severity", types.YLeaf{"AgentFilterSeverity", clients.AgentFilterSeverity})
    clients.EntityData.Leafs.Append("agent_filter_group", types.YLeaf{"AgentFilterGroup", clients.AgentFilterGroup})
    clients.EntityData.Leafs.Append("agent_sdr_id", types.YLeaf{"AgentSdrId", clients.AgentSdrId})
    clients.EntityData.Leafs.Append("agent_connect_count", types.YLeaf{"AgentConnectCount", clients.AgentConnectCount})
    clients.EntityData.Leafs.Append("agent_connect_time", types.YLeaf{"AgentConnectTime", clients.AgentConnectTime})
    clients.EntityData.Leafs.Append("agent_get_count", types.YLeaf{"AgentGetCount", clients.AgentGetCount})
    clients.EntityData.Leafs.Append("agent_subscribe_count", types.YLeaf{"AgentSubscribeCount", clients.AgentSubscribeCount})
    clients.EntityData.Leafs.Append("agent_report_count", types.YLeaf{"AgentReportCount", clients.AgentReportCount})

    clients.EntityData.YListKeys = []string {"AgentHandle"}

    return &(clients.EntityData)
}

// AlarmMgr_Detail_Card_Location_Suppressed
type AlarmMgr_Detail_Card_Location_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // The current state of the bi-state alarm. The type is StatusTd.
    State interface{}

    // The reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // The condition bit-flags of the alarm. The type is interface{} with range:
    // 0..4294967295.
    Resynced interface{}

    // Alarm detailed description. The type is string.
    DetailDesc interface{}

    // Alarm suppressed time stamp. The type is string.
    SuppressedTime interface{}
}

func (suppressed *AlarmMgr_Detail_Card_Location_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "location"
    suppressed.EntityData.SegmentPath = "suppressed" + types.AddKeyToken(suppressed.Aid, "aid") + types.AddKeyToken(suppressed.Eid, "eid")
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/card/location/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Leafs = types.NewOrderedMap()
    suppressed.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", suppressed.Aid})
    suppressed.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", suppressed.Eid})
    suppressed.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", suppressed.Tag})
    suppressed.EntityData.Leafs.Append("module", types.YLeaf{"Module", suppressed.Module})
    suppressed.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", suppressed.GenLocation})
    suppressed.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressed.Severity})
    suppressed.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressed.Group})
    suppressed.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressed.Description})
    suppressed.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", suppressed.SetTime})
    suppressed.EntityData.Leafs.Append("state", types.YLeaf{"State", suppressed.State})
    suppressed.EntityData.Leafs.Append("reporting_agent_id", types.YLeaf{"ReportingAgentId", suppressed.ReportingAgentId})
    suppressed.EntityData.Leafs.Append("resynced", types.YLeaf{"Resynced", suppressed.Resynced})
    suppressed.EntityData.Leafs.Append("detail_desc", types.YLeaf{"DetailDesc", suppressed.DetailDesc})
    suppressed.EntityData.Leafs.Append("suppressed_time", types.YLeaf{"SuppressedTime", suppressed.SuppressedTime})

    suppressed.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(suppressed.EntityData)
}

// AlarmMgr_Detail_Rack
// Alarms reported at the rack as 
// specified by the location parameter
type AlarmMgr_Detail_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AlarmMgr_Detail_Rack_RackLocations.
    RackLocations []*AlarmMgr_Detail_Rack_RackLocations
}

func (rack *AlarmMgr_Detail_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "detail"
    rack.EntityData.SegmentPath = "rack"
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("rack_locations", types.YChild{"RackLocations", nil})
    for i := range rack.RackLocations {
        rack.EntityData.Children.Append(types.GetSegmentPath(rack.RackLocations[i]), types.YChild{"RackLocations", rack.RackLocations[i]})
    }
    rack.EntityData.Leafs = types.NewOrderedMap()

    rack.EntityData.YListKeys = []string {}

    return &(rack.EntityData)
}

// AlarmMgr_Detail_Rack_RackLocations
type AlarmMgr_Detail_Rack_RackLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Rackid interface{}

    // The type is slice of AlarmMgr_Detail_Rack_RackLocations_Active.
    Active []*AlarmMgr_Detail_Rack_RackLocations_Active

    // The type is slice of AlarmMgr_Detail_Rack_RackLocations_History.
    History []*AlarmMgr_Detail_Rack_RackLocations_History

    // The type is slice of AlarmMgr_Detail_Rack_RackLocations_Stats.
    Stats []*AlarmMgr_Detail_Rack_RackLocations_Stats

    // The type is slice of AlarmMgr_Detail_Rack_RackLocations_Clients.
    Clients []*AlarmMgr_Detail_Rack_RackLocations_Clients

    // The type is slice of AlarmMgr_Detail_Rack_RackLocations_Suppressed.
    Suppressed []*AlarmMgr_Detail_Rack_RackLocations_Suppressed
}

func (rackLocations *AlarmMgr_Detail_Rack_RackLocations) GetEntityData() *types.CommonEntityData {
    rackLocations.EntityData.YFilter = rackLocations.YFilter
    rackLocations.EntityData.YangName = "rack_locations"
    rackLocations.EntityData.BundleName = "cisco_ios_xr"
    rackLocations.EntityData.ParentYangName = "rack"
    rackLocations.EntityData.SegmentPath = "rack_locations" + types.AddKeyToken(rackLocations.Rackid, "rackid")
    rackLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/rack/" + rackLocations.EntityData.SegmentPath
    rackLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackLocations.EntityData.Children = types.NewOrderedMap()
    rackLocations.EntityData.Children.Append("active", types.YChild{"Active", nil})
    for i := range rackLocations.Active {
        rackLocations.EntityData.Children.Append(types.GetSegmentPath(rackLocations.Active[i]), types.YChild{"Active", rackLocations.Active[i]})
    }
    rackLocations.EntityData.Children.Append("history", types.YChild{"History", nil})
    for i := range rackLocations.History {
        rackLocations.EntityData.Children.Append(types.GetSegmentPath(rackLocations.History[i]), types.YChild{"History", rackLocations.History[i]})
    }
    rackLocations.EntityData.Children.Append("stats", types.YChild{"Stats", nil})
    for i := range rackLocations.Stats {
        rackLocations.EntityData.Children.Append(types.GetSegmentPath(rackLocations.Stats[i]), types.YChild{"Stats", rackLocations.Stats[i]})
    }
    rackLocations.EntityData.Children.Append("clients", types.YChild{"Clients", nil})
    for i := range rackLocations.Clients {
        rackLocations.EntityData.Children.Append(types.GetSegmentPath(rackLocations.Clients[i]), types.YChild{"Clients", rackLocations.Clients[i]})
    }
    rackLocations.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", nil})
    for i := range rackLocations.Suppressed {
        rackLocations.EntityData.Children.Append(types.GetSegmentPath(rackLocations.Suppressed[i]), types.YChild{"Suppressed", rackLocations.Suppressed[i]})
    }
    rackLocations.EntityData.Leafs = types.NewOrderedMap()
    rackLocations.EntityData.Leafs.Append("rackid", types.YLeaf{"Rackid", rackLocations.Rackid})

    rackLocations.EntityData.YListKeys = []string {"Rackid"}

    return &(rackLocations.EntityData)
}

// AlarmMgr_Detail_Rack_RackLocations_Active
type AlarmMgr_Detail_Rack_RackLocations_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // The current state of the bi-state alarm. The type is StatusTd.
    State interface{}

    // The reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // The condition bit-flags of the alarm. The type is interface{} with range:
    // 0..4294967295.
    Resynced interface{}

    // Alarm detailed description. The type is string.
    DetailDesc interface{}

    // Alarm clear time stamp. The type is string.
    ClearTime interface{}
}

func (active *AlarmMgr_Detail_Rack_RackLocations_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "rack_locations"
    active.EntityData.SegmentPath = "active" + types.AddKeyToken(active.Aid, "aid") + types.AddKeyToken(active.Eid, "eid")
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/rack/rack_locations/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", active.Aid})
    active.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", active.Eid})
    active.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", active.Tag})
    active.EntityData.Leafs.Append("module", types.YLeaf{"Module", active.Module})
    active.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", active.GenLocation})
    active.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", active.Severity})
    active.EntityData.Leafs.Append("group", types.YLeaf{"Group", active.Group})
    active.EntityData.Leafs.Append("description", types.YLeaf{"Description", active.Description})
    active.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", active.SetTime})
    active.EntityData.Leafs.Append("state", types.YLeaf{"State", active.State})
    active.EntityData.Leafs.Append("reporting_agent_id", types.YLeaf{"ReportingAgentId", active.ReportingAgentId})
    active.EntityData.Leafs.Append("resynced", types.YLeaf{"Resynced", active.Resynced})
    active.EntityData.Leafs.Append("detail_desc", types.YLeaf{"DetailDesc", active.DetailDesc})
    active.EntityData.Leafs.Append("clear_time", types.YLeaf{"ClearTime", active.ClearTime})

    active.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(active.EntityData)
}

// AlarmMgr_Detail_Rack_RackLocations_History
type AlarmMgr_Detail_Rack_RackLocations_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // The current state of the bi-state alarm. The type is StatusTd.
    State interface{}

    // The reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // The condition bit-flags of the alarm. The type is interface{} with range:
    // 0..4294967295.
    Resynced interface{}

    // Alarm detailed description. The type is string.
    DetailDesc interface{}

    // Alarm clear time stamp. The type is string.
    ClearTime interface{}
}

func (history *AlarmMgr_Detail_Rack_RackLocations_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "rack_locations"
    history.EntityData.SegmentPath = "history" + types.AddKeyToken(history.Aid, "aid") + types.AddKeyToken(history.Eid, "eid")
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/rack/rack_locations/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", history.Aid})
    history.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", history.Eid})
    history.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", history.Tag})
    history.EntityData.Leafs.Append("module", types.YLeaf{"Module", history.Module})
    history.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", history.GenLocation})
    history.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", history.Severity})
    history.EntityData.Leafs.Append("group", types.YLeaf{"Group", history.Group})
    history.EntityData.Leafs.Append("description", types.YLeaf{"Description", history.Description})
    history.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", history.SetTime})
    history.EntityData.Leafs.Append("state", types.YLeaf{"State", history.State})
    history.EntityData.Leafs.Append("reporting_agent_id", types.YLeaf{"ReportingAgentId", history.ReportingAgentId})
    history.EntityData.Leafs.Append("resynced", types.YLeaf{"Resynced", history.Resynced})
    history.EntityData.Leafs.Append("detail_desc", types.YLeaf{"DetailDesc", history.DetailDesc})
    history.EntityData.Leafs.Append("clear_time", types.YLeaf{"ClearTime", history.ClearTime})

    history.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(history.EntityData)
}

// AlarmMgr_Detail_Rack_RackLocations_Stats
type AlarmMgr_Detail_Rack_RackLocations_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Alarms statistics at specified time. The type is
    // string.
    Attime interface{}

    // Total alarms reported into this service. The type is interface{} with
    // range: 0..18446744073709551615.
    Reported interface{}

    // Alarms dropped due to some error or other. The type is interface{} with
    // range: 0..18446744073709551615.
    Dropped interface{}

    // Total active alarms current in this service. The type is interface{} with
    // range: 0..18446744073709551615.
    BiSet interface{}

    // Alarms that are currently in the cleared state. The type is interface{}
    // with range: 0..18446744073709551615.
    BiClear interface{}

    // Alarms that are currently in the suppressed state. The type is interface{}
    // with range: 0..18446744073709551615.
    Suppressed interface{}

    // Alarms dropped due to invalid aid in the alarm report. The type is
    // interface{} with range: 0..18446744073709551615.
    DropInvAid interface{}

    // Alarms dropped due to low memory condition. The type is interface{} with
    // range: 0..18446744073709551615.
    DropNoMem interface{}

    // Alarms dropped due to database internal error. The type is interface{} with
    // range: 0..18446744073709551615.
    DropDbError interface{}

    // Alarms dropped due to clear without a set. The type is interface{} with
    // range: 0..18446744073709551615.
    DropClearNoSet interface{}

    // Alarms dropped due to duplicate alarm report. The type is interface{} with
    // range: 0..18446744073709551615.
    DropDup interface{}

    // Alarms present in the cached for show. The type is interface{} with range:
    // 0..18446744073709551615.
    CacheHit interface{}

    // Alarms not present in the cached for show. The type is interface{} with
    // range: 0..18446744073709551615.
    CacheMiss interface{}
}

func (stats *AlarmMgr_Detail_Rack_RackLocations_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "rack_locations"
    stats.EntityData.SegmentPath = "stats" + types.AddKeyToken(stats.Attime, "attime")
    stats.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/rack/rack_locations/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("attime", types.YLeaf{"Attime", stats.Attime})
    stats.EntityData.Leafs.Append("reported", types.YLeaf{"Reported", stats.Reported})
    stats.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", stats.Dropped})
    stats.EntityData.Leafs.Append("bi_set", types.YLeaf{"BiSet", stats.BiSet})
    stats.EntityData.Leafs.Append("bi_clear", types.YLeaf{"BiClear", stats.BiClear})
    stats.EntityData.Leafs.Append("suppressed", types.YLeaf{"Suppressed", stats.Suppressed})
    stats.EntityData.Leafs.Append("drop_inv_aid", types.YLeaf{"DropInvAid", stats.DropInvAid})
    stats.EntityData.Leafs.Append("drop_no_mem", types.YLeaf{"DropNoMem", stats.DropNoMem})
    stats.EntityData.Leafs.Append("drop_db_error", types.YLeaf{"DropDbError", stats.DropDbError})
    stats.EntityData.Leafs.Append("drop_clear_no_set", types.YLeaf{"DropClearNoSet", stats.DropClearNoSet})
    stats.EntityData.Leafs.Append("drop_dup", types.YLeaf{"DropDup", stats.DropDup})
    stats.EntityData.Leafs.Append("cache_hit", types.YLeaf{"CacheHit", stats.CacheHit})
    stats.EntityData.Leafs.Append("cache_miss", types.YLeaf{"CacheMiss", stats.CacheMiss})

    stats.EntityData.YListKeys = []string {"Attime"}

    return &(stats.EntityData)
}

// AlarmMgr_Detail_Rack_RackLocations_Clients
type AlarmMgr_Detail_Rack_RackLocations_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The client handle through which interface. The
    // type is string.
    AgentHandle interface{}

    // Alarms client. The type is string.
    AgentName interface{}

    // Alarms agent id of the client. The type is interface{} with range:
    // 0..4294967295.
    AgentId interface{}

    // The location of this client. The type is string.
    AgentLocation interface{}

    // The current state of the client. The type is AgentStateTd.
    AgentState interface{}

    // The type of  the client. The type is AgentTypeTd.
    AgentType interface{}

    // The current subscription status of the client. The type is bool.
    AgentFilterDisp interface{}

    // The subscriber id of the client. The type is interface{} with range:
    // 0..4294967295.
    AgentSubsId interface{}

    // The filter used for alarm bi-state state. The type is StatusTd.
    AgentFilterState interface{}

    // The filter used for alarm severity. The type is SeverityTd.
    AgentFilterSeverity interface{}

    // The filter used for alarm group. The type is GroupTd.
    AgentFilterGroup interface{}

    // agent alarm_sm register with sdr_id in SOMT mode. The type is interface{}
    // with range: 0..4294967295.
    AgentSdrId interface{}

    // Number of times the agent connected to the alarm mgr. The type is
    // interface{} with range: 0..18446744073709551615.
    AgentConnectCount interface{}

    // Agent connect timestamp. The type is string.
    AgentConnectTime interface{}

    // Number of times the agent queried for alarms. The type is interface{} with
    // range: 0..18446744073709551615.
    AgentGetCount interface{}

    // Number of times the agent subscribed for alarms. The type is interface{}
    // with range: 0..18446744073709551615.
    AgentSubscribeCount interface{}

    // Number of times the agent reported alarms. The type is interface{} with
    // range: 0..18446744073709551615.
    AgentReportCount interface{}
}

func (clients *AlarmMgr_Detail_Rack_RackLocations_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "rack_locations"
    clients.EntityData.SegmentPath = "clients" + types.AddKeyToken(clients.AgentHandle, "agent_handle")
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/rack/rack_locations/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Leafs = types.NewOrderedMap()
    clients.EntityData.Leafs.Append("agent_handle", types.YLeaf{"AgentHandle", clients.AgentHandle})
    clients.EntityData.Leafs.Append("agent_name", types.YLeaf{"AgentName", clients.AgentName})
    clients.EntityData.Leafs.Append("agent_id", types.YLeaf{"AgentId", clients.AgentId})
    clients.EntityData.Leafs.Append("agent_location", types.YLeaf{"AgentLocation", clients.AgentLocation})
    clients.EntityData.Leafs.Append("agent_state", types.YLeaf{"AgentState", clients.AgentState})
    clients.EntityData.Leafs.Append("agent_type", types.YLeaf{"AgentType", clients.AgentType})
    clients.EntityData.Leafs.Append("agent_filter_disp", types.YLeaf{"AgentFilterDisp", clients.AgentFilterDisp})
    clients.EntityData.Leafs.Append("agent_subs_id", types.YLeaf{"AgentSubsId", clients.AgentSubsId})
    clients.EntityData.Leafs.Append("agent_filter_state", types.YLeaf{"AgentFilterState", clients.AgentFilterState})
    clients.EntityData.Leafs.Append("agent_filter_severity", types.YLeaf{"AgentFilterSeverity", clients.AgentFilterSeverity})
    clients.EntityData.Leafs.Append("agent_filter_group", types.YLeaf{"AgentFilterGroup", clients.AgentFilterGroup})
    clients.EntityData.Leafs.Append("agent_sdr_id", types.YLeaf{"AgentSdrId", clients.AgentSdrId})
    clients.EntityData.Leafs.Append("agent_connect_count", types.YLeaf{"AgentConnectCount", clients.AgentConnectCount})
    clients.EntityData.Leafs.Append("agent_connect_time", types.YLeaf{"AgentConnectTime", clients.AgentConnectTime})
    clients.EntityData.Leafs.Append("agent_get_count", types.YLeaf{"AgentGetCount", clients.AgentGetCount})
    clients.EntityData.Leafs.Append("agent_subscribe_count", types.YLeaf{"AgentSubscribeCount", clients.AgentSubscribeCount})
    clients.EntityData.Leafs.Append("agent_report_count", types.YLeaf{"AgentReportCount", clients.AgentReportCount})

    clients.EntityData.YListKeys = []string {"AgentHandle"}

    return &(clients.EntityData)
}

// AlarmMgr_Detail_Rack_RackLocations_Suppressed
type AlarmMgr_Detail_Rack_RackLocations_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // The current state of the bi-state alarm. The type is StatusTd.
    State interface{}

    // The reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // The condition bit-flags of the alarm. The type is interface{} with range:
    // 0..4294967295.
    Resynced interface{}

    // Alarm detailed description. The type is string.
    DetailDesc interface{}

    // Alarm suppressed time stamp. The type is string.
    SuppressedTime interface{}
}

func (suppressed *AlarmMgr_Detail_Rack_RackLocations_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "rack_locations"
    suppressed.EntityData.SegmentPath = "suppressed" + types.AddKeyToken(suppressed.Aid, "aid") + types.AddKeyToken(suppressed.Eid, "eid")
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/rack/rack_locations/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Leafs = types.NewOrderedMap()
    suppressed.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", suppressed.Aid})
    suppressed.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", suppressed.Eid})
    suppressed.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", suppressed.Tag})
    suppressed.EntityData.Leafs.Append("module", types.YLeaf{"Module", suppressed.Module})
    suppressed.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", suppressed.GenLocation})
    suppressed.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressed.Severity})
    suppressed.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressed.Group})
    suppressed.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressed.Description})
    suppressed.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", suppressed.SetTime})
    suppressed.EntityData.Leafs.Append("state", types.YLeaf{"State", suppressed.State})
    suppressed.EntityData.Leafs.Append("reporting_agent_id", types.YLeaf{"ReportingAgentId", suppressed.ReportingAgentId})
    suppressed.EntityData.Leafs.Append("resynced", types.YLeaf{"Resynced", suppressed.Resynced})
    suppressed.EntityData.Leafs.Append("detail_desc", types.YLeaf{"DetailDesc", suppressed.DetailDesc})
    suppressed.EntityData.Leafs.Append("suppressed_time", types.YLeaf{"SuppressedTime", suppressed.SuppressedTime})

    suppressed.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(suppressed.EntityData)
}

// AlarmMgr_Detail_System
// Alarms reported at the system as 
// specified by the location parameter
type AlarmMgr_Detail_System struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of AlarmMgr_Detail_System_Active.
    Active []*AlarmMgr_Detail_System_Active

    // The type is slice of AlarmMgr_Detail_System_History.
    History []*AlarmMgr_Detail_System_History

    // The type is slice of AlarmMgr_Detail_System_Stats.
    Stats []*AlarmMgr_Detail_System_Stats

    // The type is slice of AlarmMgr_Detail_System_Clients.
    Clients []*AlarmMgr_Detail_System_Clients

    // The type is slice of AlarmMgr_Detail_System_Suppressed.
    Suppressed []*AlarmMgr_Detail_System_Suppressed
}

func (system *AlarmMgr_Detail_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xr"
    system.EntityData.ParentYangName = "detail"
    system.EntityData.SegmentPath = "system"
    system.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/" + system.EntityData.SegmentPath
    system.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    system.EntityData.Children = types.NewOrderedMap()
    system.EntityData.Children.Append("active", types.YChild{"Active", nil})
    for i := range system.Active {
        system.EntityData.Children.Append(types.GetSegmentPath(system.Active[i]), types.YChild{"Active", system.Active[i]})
    }
    system.EntityData.Children.Append("history", types.YChild{"History", nil})
    for i := range system.History {
        system.EntityData.Children.Append(types.GetSegmentPath(system.History[i]), types.YChild{"History", system.History[i]})
    }
    system.EntityData.Children.Append("stats", types.YChild{"Stats", nil})
    for i := range system.Stats {
        system.EntityData.Children.Append(types.GetSegmentPath(system.Stats[i]), types.YChild{"Stats", system.Stats[i]})
    }
    system.EntityData.Children.Append("clients", types.YChild{"Clients", nil})
    for i := range system.Clients {
        system.EntityData.Children.Append(types.GetSegmentPath(system.Clients[i]), types.YChild{"Clients", system.Clients[i]})
    }
    system.EntityData.Children.Append("suppressed", types.YChild{"Suppressed", nil})
    for i := range system.Suppressed {
        system.EntityData.Children.Append(types.GetSegmentPath(system.Suppressed[i]), types.YChild{"Suppressed", system.Suppressed[i]})
    }
    system.EntityData.Leafs = types.NewOrderedMap()

    system.EntityData.YListKeys = []string {}

    return &(system.EntityData)
}

// AlarmMgr_Detail_System_Active
type AlarmMgr_Detail_System_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // The current state of the bi-state alarm. The type is StatusTd.
    State interface{}

    // The reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // The condition bit-flags of the alarm. The type is interface{} with range:
    // 0..4294967295.
    Resynced interface{}

    // Alarm detailed description. The type is string.
    DetailDesc interface{}

    // Alarm clear time stamp. The type is string.
    ClearTime interface{}
}

func (active *AlarmMgr_Detail_System_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "system"
    active.EntityData.SegmentPath = "active" + types.AddKeyToken(active.Aid, "aid") + types.AddKeyToken(active.Eid, "eid")
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/system/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", active.Aid})
    active.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", active.Eid})
    active.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", active.Tag})
    active.EntityData.Leafs.Append("module", types.YLeaf{"Module", active.Module})
    active.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", active.GenLocation})
    active.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", active.Severity})
    active.EntityData.Leafs.Append("group", types.YLeaf{"Group", active.Group})
    active.EntityData.Leafs.Append("description", types.YLeaf{"Description", active.Description})
    active.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", active.SetTime})
    active.EntityData.Leafs.Append("state", types.YLeaf{"State", active.State})
    active.EntityData.Leafs.Append("reporting_agent_id", types.YLeaf{"ReportingAgentId", active.ReportingAgentId})
    active.EntityData.Leafs.Append("resynced", types.YLeaf{"Resynced", active.Resynced})
    active.EntityData.Leafs.Append("detail_desc", types.YLeaf{"DetailDesc", active.DetailDesc})
    active.EntityData.Leafs.Append("clear_time", types.YLeaf{"ClearTime", active.ClearTime})

    active.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(active.EntityData)
}

// AlarmMgr_Detail_System_History
type AlarmMgr_Detail_System_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // The current state of the bi-state alarm. The type is StatusTd.
    State interface{}

    // The reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // The condition bit-flags of the alarm. The type is interface{} with range:
    // 0..4294967295.
    Resynced interface{}

    // Alarm detailed description. The type is string.
    DetailDesc interface{}

    // Alarm clear time stamp. The type is string.
    ClearTime interface{}
}

func (history *AlarmMgr_Detail_System_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "system"
    history.EntityData.SegmentPath = "history" + types.AddKeyToken(history.Aid, "aid") + types.AddKeyToken(history.Eid, "eid")
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/system/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", history.Aid})
    history.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", history.Eid})
    history.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", history.Tag})
    history.EntityData.Leafs.Append("module", types.YLeaf{"Module", history.Module})
    history.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", history.GenLocation})
    history.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", history.Severity})
    history.EntityData.Leafs.Append("group", types.YLeaf{"Group", history.Group})
    history.EntityData.Leafs.Append("description", types.YLeaf{"Description", history.Description})
    history.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", history.SetTime})
    history.EntityData.Leafs.Append("state", types.YLeaf{"State", history.State})
    history.EntityData.Leafs.Append("reporting_agent_id", types.YLeaf{"ReportingAgentId", history.ReportingAgentId})
    history.EntityData.Leafs.Append("resynced", types.YLeaf{"Resynced", history.Resynced})
    history.EntityData.Leafs.Append("detail_desc", types.YLeaf{"DetailDesc", history.DetailDesc})
    history.EntityData.Leafs.Append("clear_time", types.YLeaf{"ClearTime", history.ClearTime})

    history.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(history.EntityData)
}

// AlarmMgr_Detail_System_Stats
type AlarmMgr_Detail_System_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Alarms statistics at specified time. The type is
    // string.
    Attime interface{}

    // Total alarms reported into this service. The type is interface{} with
    // range: 0..18446744073709551615.
    Reported interface{}

    // Alarms dropped due to some error or other. The type is interface{} with
    // range: 0..18446744073709551615.
    Dropped interface{}

    // Total active alarms current in this service. The type is interface{} with
    // range: 0..18446744073709551615.
    BiSet interface{}

    // Alarms that are currently in the cleared state. The type is interface{}
    // with range: 0..18446744073709551615.
    BiClear interface{}

    // Alarms that are currently in the suppressed state. The type is interface{}
    // with range: 0..18446744073709551615.
    Suppressed interface{}

    // Alarms dropped due to invalid aid in the alarm report. The type is
    // interface{} with range: 0..18446744073709551615.
    DropInvAid interface{}

    // Alarms dropped due to low memory condition. The type is interface{} with
    // range: 0..18446744073709551615.
    DropNoMem interface{}

    // Alarms dropped due to database internal error. The type is interface{} with
    // range: 0..18446744073709551615.
    DropDbError interface{}

    // Alarms dropped due to clear without a set. The type is interface{} with
    // range: 0..18446744073709551615.
    DropClearNoSet interface{}

    // Alarms dropped due to duplicate alarm report. The type is interface{} with
    // range: 0..18446744073709551615.
    DropDup interface{}

    // Alarms present in the cached for show. The type is interface{} with range:
    // 0..18446744073709551615.
    CacheHit interface{}

    // Alarms not present in the cached for show. The type is interface{} with
    // range: 0..18446744073709551615.
    CacheMiss interface{}
}

func (stats *AlarmMgr_Detail_System_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "system"
    stats.EntityData.SegmentPath = "stats" + types.AddKeyToken(stats.Attime, "attime")
    stats.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/system/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("attime", types.YLeaf{"Attime", stats.Attime})
    stats.EntityData.Leafs.Append("reported", types.YLeaf{"Reported", stats.Reported})
    stats.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", stats.Dropped})
    stats.EntityData.Leafs.Append("bi_set", types.YLeaf{"BiSet", stats.BiSet})
    stats.EntityData.Leafs.Append("bi_clear", types.YLeaf{"BiClear", stats.BiClear})
    stats.EntityData.Leafs.Append("suppressed", types.YLeaf{"Suppressed", stats.Suppressed})
    stats.EntityData.Leafs.Append("drop_inv_aid", types.YLeaf{"DropInvAid", stats.DropInvAid})
    stats.EntityData.Leafs.Append("drop_no_mem", types.YLeaf{"DropNoMem", stats.DropNoMem})
    stats.EntityData.Leafs.Append("drop_db_error", types.YLeaf{"DropDbError", stats.DropDbError})
    stats.EntityData.Leafs.Append("drop_clear_no_set", types.YLeaf{"DropClearNoSet", stats.DropClearNoSet})
    stats.EntityData.Leafs.Append("drop_dup", types.YLeaf{"DropDup", stats.DropDup})
    stats.EntityData.Leafs.Append("cache_hit", types.YLeaf{"CacheHit", stats.CacheHit})
    stats.EntityData.Leafs.Append("cache_miss", types.YLeaf{"CacheMiss", stats.CacheMiss})

    stats.EntityData.YListKeys = []string {"Attime"}

    return &(stats.EntityData)
}

// AlarmMgr_Detail_System_Clients
type AlarmMgr_Detail_System_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The client handle through which interface. The
    // type is string.
    AgentHandle interface{}

    // Alarms client. The type is string.
    AgentName interface{}

    // Alarms agent id of the client. The type is interface{} with range:
    // 0..4294967295.
    AgentId interface{}

    // The location of this client. The type is string.
    AgentLocation interface{}

    // The current state of the client. The type is AgentStateTd.
    AgentState interface{}

    // The type of  the client. The type is AgentTypeTd.
    AgentType interface{}

    // The current subscription status of the client. The type is bool.
    AgentFilterDisp interface{}

    // The subscriber id of the client. The type is interface{} with range:
    // 0..4294967295.
    AgentSubsId interface{}

    // The filter used for alarm bi-state state. The type is StatusTd.
    AgentFilterState interface{}

    // The filter used for alarm severity. The type is SeverityTd.
    AgentFilterSeverity interface{}

    // The filter used for alarm group. The type is GroupTd.
    AgentFilterGroup interface{}

    // agent alarm_sm register with sdr_id in SOMT mode. The type is interface{}
    // with range: 0..4294967295.
    AgentSdrId interface{}

    // Number of times the agent connected to the alarm mgr. The type is
    // interface{} with range: 0..18446744073709551615.
    AgentConnectCount interface{}

    // Agent connect timestamp. The type is string.
    AgentConnectTime interface{}

    // Number of times the agent queried for alarms. The type is interface{} with
    // range: 0..18446744073709551615.
    AgentGetCount interface{}

    // Number of times the agent subscribed for alarms. The type is interface{}
    // with range: 0..18446744073709551615.
    AgentSubscribeCount interface{}

    // Number of times the agent reported alarms. The type is interface{} with
    // range: 0..18446744073709551615.
    AgentReportCount interface{}
}

func (clients *AlarmMgr_Detail_System_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "system"
    clients.EntityData.SegmentPath = "clients" + types.AddKeyToken(clients.AgentHandle, "agent_handle")
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/system/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Leafs = types.NewOrderedMap()
    clients.EntityData.Leafs.Append("agent_handle", types.YLeaf{"AgentHandle", clients.AgentHandle})
    clients.EntityData.Leafs.Append("agent_name", types.YLeaf{"AgentName", clients.AgentName})
    clients.EntityData.Leafs.Append("agent_id", types.YLeaf{"AgentId", clients.AgentId})
    clients.EntityData.Leafs.Append("agent_location", types.YLeaf{"AgentLocation", clients.AgentLocation})
    clients.EntityData.Leafs.Append("agent_state", types.YLeaf{"AgentState", clients.AgentState})
    clients.EntityData.Leafs.Append("agent_type", types.YLeaf{"AgentType", clients.AgentType})
    clients.EntityData.Leafs.Append("agent_filter_disp", types.YLeaf{"AgentFilterDisp", clients.AgentFilterDisp})
    clients.EntityData.Leafs.Append("agent_subs_id", types.YLeaf{"AgentSubsId", clients.AgentSubsId})
    clients.EntityData.Leafs.Append("agent_filter_state", types.YLeaf{"AgentFilterState", clients.AgentFilterState})
    clients.EntityData.Leafs.Append("agent_filter_severity", types.YLeaf{"AgentFilterSeverity", clients.AgentFilterSeverity})
    clients.EntityData.Leafs.Append("agent_filter_group", types.YLeaf{"AgentFilterGroup", clients.AgentFilterGroup})
    clients.EntityData.Leafs.Append("agent_sdr_id", types.YLeaf{"AgentSdrId", clients.AgentSdrId})
    clients.EntityData.Leafs.Append("agent_connect_count", types.YLeaf{"AgentConnectCount", clients.AgentConnectCount})
    clients.EntityData.Leafs.Append("agent_connect_time", types.YLeaf{"AgentConnectTime", clients.AgentConnectTime})
    clients.EntityData.Leafs.Append("agent_get_count", types.YLeaf{"AgentGetCount", clients.AgentGetCount})
    clients.EntityData.Leafs.Append("agent_subscribe_count", types.YLeaf{"AgentSubscribeCount", clients.AgentSubscribeCount})
    clients.EntityData.Leafs.Append("agent_report_count", types.YLeaf{"AgentReportCount", clients.AgentReportCount})

    clients.EntityData.YListKeys = []string {"AgentHandle"}

    return &(clients.EntityData)
}

// AlarmMgr_Detail_System_Suppressed
type AlarmMgr_Detail_System_Suppressed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The AID for the current alarm. The type is string.
    Aid interface{}

    // This attribute is a key. The Object Id for the current alarm. The type is
    // string.
    Eid interface{}

    // The Fault Tag for the current alarm. The type is string.
    Tag interface{}

    // The Module for the current alarm. The type is string.
    Module interface{}

    // The location for the current alarm. The type is string.
    GenLocation interface{}

    // The alarm severity. The type is SeverityTd.
    Severity interface{}

    // The alarm grouping . The type is GroupTd.
    Group interface{}

    // Alarm description. The type is string.
    Description interface{}

    // Alarm set time stamp. The type is string.
    SetTime interface{}

    // The current state of the bi-state alarm. The type is StatusTd.
    State interface{}

    // The reporting agent id. The type is interface{} with range: 0..4294967295.
    ReportingAgentId interface{}

    // The condition bit-flags of the alarm. The type is interface{} with range:
    // 0..4294967295.
    Resynced interface{}

    // Alarm detailed description. The type is string.
    DetailDesc interface{}

    // Alarm suppressed time stamp. The type is string.
    SuppressedTime interface{}
}

func (suppressed *AlarmMgr_Detail_System_Suppressed) GetEntityData() *types.CommonEntityData {
    suppressed.EntityData.YFilter = suppressed.YFilter
    suppressed.EntityData.YangName = "suppressed"
    suppressed.EntityData.BundleName = "cisco_ios_xr"
    suppressed.EntityData.ParentYangName = "system"
    suppressed.EntityData.SegmentPath = "suppressed" + types.AddKeyToken(suppressed.Aid, "aid") + types.AddKeyToken(suppressed.Eid, "eid")
    suppressed.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-alarm-mgr:alarm_mgr/detail/system/" + suppressed.EntityData.SegmentPath
    suppressed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressed.EntityData.Children = types.NewOrderedMap()
    suppressed.EntityData.Leafs = types.NewOrderedMap()
    suppressed.EntityData.Leafs.Append("aid", types.YLeaf{"Aid", suppressed.Aid})
    suppressed.EntityData.Leafs.Append("eid", types.YLeaf{"Eid", suppressed.Eid})
    suppressed.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", suppressed.Tag})
    suppressed.EntityData.Leafs.Append("module", types.YLeaf{"Module", suppressed.Module})
    suppressed.EntityData.Leafs.Append("gen_location", types.YLeaf{"GenLocation", suppressed.GenLocation})
    suppressed.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", suppressed.Severity})
    suppressed.EntityData.Leafs.Append("group", types.YLeaf{"Group", suppressed.Group})
    suppressed.EntityData.Leafs.Append("description", types.YLeaf{"Description", suppressed.Description})
    suppressed.EntityData.Leafs.Append("set_time", types.YLeaf{"SetTime", suppressed.SetTime})
    suppressed.EntityData.Leafs.Append("state", types.YLeaf{"State", suppressed.State})
    suppressed.EntityData.Leafs.Append("reporting_agent_id", types.YLeaf{"ReportingAgentId", suppressed.ReportingAgentId})
    suppressed.EntityData.Leafs.Append("resynced", types.YLeaf{"Resynced", suppressed.Resynced})
    suppressed.EntityData.Leafs.Append("detail_desc", types.YLeaf{"DetailDesc", suppressed.DetailDesc})
    suppressed.EntityData.Leafs.Append("suppressed_time", types.YLeaf{"SuppressedTime", suppressed.SuppressedTime})

    suppressed.EntityData.YListKeys = []string {"Aid", "Eid"}

    return &(suppressed.EntityData)
}

