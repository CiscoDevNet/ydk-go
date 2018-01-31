// This module contains a collection of YANG definitions
// for Cisco IOS-XR shellutil package operational data.
// 
// This module contains definitions
// for the following management objects:
//   system-time: System time information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    TimeSource_error TimeSource = "error"

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
    parent types.Entity
    YFilter yfilter.YFilter

    // System clock information.
    Clock SystemTime_Clock

    // System uptime information.
    Uptime SystemTime_Uptime
}

func (systemTime *SystemTime) GetFilter() yfilter.YFilter { return systemTime.YFilter }

func (systemTime *SystemTime) SetFilter(yf yfilter.YFilter) { systemTime.YFilter = yf }

func (systemTime *SystemTime) GetGoName(yname string) string {
    if yname == "clock" { return "Clock" }
    if yname == "uptime" { return "Uptime" }
    return ""
}

func (systemTime *SystemTime) GetSegmentPath() string {
    return "Cisco-IOS-XR-shellutil-oper:system-time"
}

func (systemTime *SystemTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock" {
        return &systemTime.Clock
    }
    if childYangName == "uptime" {
        return &systemTime.Uptime
    }
    return nil
}

func (systemTime *SystemTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock"] = &systemTime.Clock
    children["uptime"] = &systemTime.Uptime
    return children
}

func (systemTime *SystemTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (systemTime *SystemTime) GetBundleName() string { return "cisco_ios_xr" }

func (systemTime *SystemTime) GetYangName() string { return "system-time" }

func (systemTime *SystemTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemTime *SystemTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemTime *SystemTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemTime *SystemTime) SetParent(parent types.Entity) { systemTime.parent = parent }

func (systemTime *SystemTime) GetParent() types.Entity { return systemTime.parent }

func (systemTime *SystemTime) GetParentYangName() string { return "Cisco-IOS-XR-shellutil-oper" }

// SystemTime_Clock
// System clock information
type SystemTime_Clock struct {
    parent types.Entity
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

func (clock *SystemTime_Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *SystemTime_Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *SystemTime_Clock) GetGoName(yname string) string {
    if yname == "year" { return "Year" }
    if yname == "month" { return "Month" }
    if yname == "day" { return "Day" }
    if yname == "hour" { return "Hour" }
    if yname == "minute" { return "Minute" }
    if yname == "second" { return "Second" }
    if yname == "millisecond" { return "Millisecond" }
    if yname == "wday" { return "Wday" }
    if yname == "time-zone" { return "TimeZone" }
    if yname == "time-source" { return "TimeSource" }
    return ""
}

func (clock *SystemTime_Clock) GetSegmentPath() string {
    return "clock"
}

func (clock *SystemTime_Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clock *SystemTime_Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clock *SystemTime_Clock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["year"] = clock.Year
    leafs["month"] = clock.Month
    leafs["day"] = clock.Day
    leafs["hour"] = clock.Hour
    leafs["minute"] = clock.Minute
    leafs["second"] = clock.Second
    leafs["millisecond"] = clock.Millisecond
    leafs["wday"] = clock.Wday
    leafs["time-zone"] = clock.TimeZone
    leafs["time-source"] = clock.TimeSource
    return leafs
}

func (clock *SystemTime_Clock) GetBundleName() string { return "cisco_ios_xr" }

func (clock *SystemTime_Clock) GetYangName() string { return "clock" }

func (clock *SystemTime_Clock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clock *SystemTime_Clock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clock *SystemTime_Clock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clock *SystemTime_Clock) SetParent(parent types.Entity) { clock.parent = parent }

func (clock *SystemTime_Clock) GetParent() types.Entity { return clock.parent }

func (clock *SystemTime_Clock) GetParentYangName() string { return "system-time" }

// SystemTime_Uptime
// System uptime information
type SystemTime_Uptime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Host name. The type is string.
    HostName interface{}

    // Amount of time in seconds since this system     was last initialized. The
    // type is interface{} with range: 0..4294967295. Units are second.
    Uptime interface{}
}

func (uptime *SystemTime_Uptime) GetFilter() yfilter.YFilter { return uptime.YFilter }

func (uptime *SystemTime_Uptime) SetFilter(yf yfilter.YFilter) { uptime.YFilter = yf }

func (uptime *SystemTime_Uptime) GetGoName(yname string) string {
    if yname == "host-name" { return "HostName" }
    if yname == "uptime" { return "Uptime" }
    return ""
}

func (uptime *SystemTime_Uptime) GetSegmentPath() string {
    return "uptime"
}

func (uptime *SystemTime_Uptime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (uptime *SystemTime_Uptime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (uptime *SystemTime_Uptime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-name"] = uptime.HostName
    leafs["uptime"] = uptime.Uptime
    return leafs
}

func (uptime *SystemTime_Uptime) GetBundleName() string { return "cisco_ios_xr" }

func (uptime *SystemTime_Uptime) GetYangName() string { return "uptime" }

func (uptime *SystemTime_Uptime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (uptime *SystemTime_Uptime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (uptime *SystemTime_Uptime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (uptime *SystemTime_Uptime) SetParent(parent types.Entity) { uptime.parent = parent }

func (uptime *SystemTime_Uptime) GetParent() types.Entity { return uptime.parent }

func (uptime *SystemTime_Uptime) GetParentYangName() string { return "system-time" }

