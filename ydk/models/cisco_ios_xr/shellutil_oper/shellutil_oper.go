// This module contains a collection of YANG definitions
// for Cisco IOS-XR shellutil package operational data.
// 
// This module contains definitions
// for the following management objects:
//   system-time: System time information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package shellutil_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package shellutil_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-shellutil-oper system-time}", reflect.TypeOf(SystemTime{}))
    ydk.RegisterEntity("Cisco-IOS-XR-shellutil-oper:system-time", reflect.TypeOf(SystemTime{}))
}

// TimeSource represents Time source
type TimeSource string

const (
    // Error
    TimeSource_error_ TimeSource = "error"

    // Unsynchronized time
    TimeSource_none TimeSource = "none"

    // Network time protocol
    TimeSource_ntp TimeSource = "ntp"

    // User configured
    TimeSource_manual TimeSource = "manual"

    // Hardware calendar
    TimeSource_calendar TimeSource = "calendar"
)

// SystemTime
// System time information
type SystemTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // System clock information.
    Clock SystemTime_Clock

    // System uptime information.
    Uptime SystemTime_Uptime
}

func (systemTime *SystemTime) GetEntityData() *types.CommonEntityData {
    systemTime.EntityData.YFilter = systemTime.YFilter
    systemTime.EntityData.YangName = "system-time"
    systemTime.EntityData.BundleName = "cisco_ios_xr"
    systemTime.EntityData.ParentYangName = "Cisco-IOS-XR-shellutil-oper"
    systemTime.EntityData.SegmentPath = "Cisco-IOS-XR-shellutil-oper:system-time"
    systemTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemTime.EntityData.Children = types.NewOrderedMap()
    systemTime.EntityData.Children.Append("clock", types.YChild{"Clock", &systemTime.Clock})
    systemTime.EntityData.Children.Append("uptime", types.YChild{"Uptime", &systemTime.Uptime})
    systemTime.EntityData.Leafs = types.NewOrderedMap()

    systemTime.EntityData.YListKeys = []string {}

    return &(systemTime.EntityData)
}

// SystemTime_Clock
// System clock information
type SystemTime_Clock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Year [0..65535]. The type is interface{} with range: 0..65535.
    Year interface{}

    // Month [1..12]. The type is interface{} with range: 0..255.
    Month interface{}

    // Day [1..31]. The type is interface{} with range: 0..255.
    Day interface{}

    // Hour [0..23]. The type is interface{} with range: 0..255.
    Hour interface{}

    // Minute [0..59]. The type is interface{} with range: 0..255.
    Minute interface{}

    // Second [0..60], use 60 for leap-second. The type is interface{} with range:
    // 0..255.
    Second interface{}

    // Millisecond [0..999]. The type is interface{} with range: 0..65535.
    Millisecond interface{}

    // Week Day [0..6]. The type is interface{} with range: 0..65535.
    Wday interface{}

    // Time zone. The type is string.
    TimeZone interface{}

    // Time source. The type is TimeSource.
    TimeSource interface{}
}

func (clock *SystemTime_Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "system-time"
    clock.EntityData.SegmentPath = "clock"
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Leafs = types.NewOrderedMap()
    clock.EntityData.Leafs.Append("year", types.YLeaf{"Year", clock.Year})
    clock.EntityData.Leafs.Append("month", types.YLeaf{"Month", clock.Month})
    clock.EntityData.Leafs.Append("day", types.YLeaf{"Day", clock.Day})
    clock.EntityData.Leafs.Append("hour", types.YLeaf{"Hour", clock.Hour})
    clock.EntityData.Leafs.Append("minute", types.YLeaf{"Minute", clock.Minute})
    clock.EntityData.Leafs.Append("second", types.YLeaf{"Second", clock.Second})
    clock.EntityData.Leafs.Append("millisecond", types.YLeaf{"Millisecond", clock.Millisecond})
    clock.EntityData.Leafs.Append("wday", types.YLeaf{"Wday", clock.Wday})
    clock.EntityData.Leafs.Append("time-zone", types.YLeaf{"TimeZone", clock.TimeZone})
    clock.EntityData.Leafs.Append("time-source", types.YLeaf{"TimeSource", clock.TimeSource})

    clock.EntityData.YListKeys = []string {}

    return &(clock.EntityData)
}

// SystemTime_Uptime
// System uptime information
type SystemTime_Uptime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Host name. The type is string.
    HostName interface{}

    // Amount of time in seconds since this system     was last initialized. The
    // type is interface{} with range: 0..4294967295. Units are second.
    Uptime interface{}
}

func (uptime *SystemTime_Uptime) GetEntityData() *types.CommonEntityData {
    uptime.EntityData.YFilter = uptime.YFilter
    uptime.EntityData.YangName = "uptime"
    uptime.EntityData.BundleName = "cisco_ios_xr"
    uptime.EntityData.ParentYangName = "system-time"
    uptime.EntityData.SegmentPath = "uptime"
    uptime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    uptime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    uptime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    uptime.EntityData.Children = types.NewOrderedMap()
    uptime.EntityData.Leafs = types.NewOrderedMap()
    uptime.EntityData.Leafs.Append("host-name", types.YLeaf{"HostName", uptime.HostName})
    uptime.EntityData.Leafs.Append("uptime", types.YLeaf{"Uptime", uptime.Uptime})

    uptime.EntityData.YListKeys = []string {}

    return &(uptime.EntityData)
}

