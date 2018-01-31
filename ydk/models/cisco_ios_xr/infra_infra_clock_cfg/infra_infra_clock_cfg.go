// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-infra-clock package configuration.
// 
// This module contains definitions
// for the following management objects:
//   clock: Configure time-of-day clock
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_infra_clock_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_infra_clock_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-infra-clock-cfg clock}", reflect.TypeOf(Clock{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-infra-clock-cfg:clock", reflect.TypeOf(Clock{}))
}

// ClockMonth represents Clock month
type ClockMonth string

const (
    // January
    ClockMonth_january ClockMonth = "january"

    // February
    ClockMonth_february ClockMonth = "february"

    // March
    ClockMonth_march ClockMonth = "march"

    // April
    ClockMonth_april ClockMonth = "april"

    // May
    ClockMonth_may ClockMonth = "may"

    // June
    ClockMonth_june ClockMonth = "june"

    // July
    ClockMonth_july ClockMonth = "july"

    // August
    ClockMonth_august ClockMonth = "august"

    // September
    ClockMonth_september ClockMonth = "september"

    // October
    ClockMonth_october ClockMonth = "october"

    // November
    ClockMonth_november ClockMonth = "november"

    // December
    ClockMonth_december ClockMonth = "december"
)

// ClockSummerTimeMode represents Clock summer time mode
type ClockSummerTimeMode string

const (
    // Recurring summer time
    ClockSummerTimeMode_recurring ClockSummerTimeMode = "recurring"

    // Absolute summer time
    ClockSummerTimeMode_date ClockSummerTimeMode = "date"
)

// Clock
// Configure time-of-day clock
type Clock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure summer (daylight savings) time.
    SummerTime Clock_SummerTime

    // Configure time zone.
    TimeZone Clock_TimeZone
}

func (clock *Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *Clock) GetGoName(yname string) string {
    if yname == "summer-time" { return "SummerTime" }
    if yname == "time-zone" { return "TimeZone" }
    return ""
}

func (clock *Clock) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-infra-clock-cfg:clock"
}

func (clock *Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summer-time" {
        return &clock.SummerTime
    }
    if childYangName == "time-zone" {
        return &clock.TimeZone
    }
    return nil
}

func (clock *Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summer-time"] = &clock.SummerTime
    children["time-zone"] = &clock.TimeZone
    return children
}

func (clock *Clock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clock *Clock) GetBundleName() string { return "cisco_ios_xr" }

func (clock *Clock) GetYangName() string { return "clock" }

func (clock *Clock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clock *Clock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clock *Clock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clock *Clock) SetParent(parent types.Entity) { clock.parent = parent }

func (clock *Clock) GetParent() types.Entity { return clock.parent }

func (clock *Clock) GetParentYangName() string { return "Cisco-IOS-XR-infra-infra-clock-cfg" }

// Clock_SummerTime
// Configure summer (daylight savings) time
// This type is a presence type.
type Clock_SummerTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of time zone in summer. The type is string. This attribute is
    // mandatory.
    TimeZoneName interface{}

    // Summer time mode. The type is ClockSummerTimeMode. This attribute is
    // mandatory.
    Mode interface{}

    // If Mode is set to 'Recurring' specify Week number of the Month to start
    // (first and last strings are not allowed as they are in CLI) , if Mode is
    // set to 'Date' specify Date to start . The type is interface{} with range:
    // 1..31.
    StartWeekNumberOrStartDate interface{}

    // If Mode is set to 'Recurring' specify Weekday to start , if Mode is set to
    // 'Date' specify Year to start . The type is interface{} with range: 0..2035.
    StartWeekdayOrStartYear interface{}

    // Month to start . The type is ClockMonth.
    StartMonth interface{}

    // Hour to start . The type is interface{} with range: 0..23.
    StartHour interface{}

    // Minute to start . The type is interface{} with range: 0..59.
    StartMinute interface{}

    // If Mode is set to 'Recurring' specify Week number of the Month to end
    // (first and last strings are not allowed as they are in CLI), if Mode is set
    // to 'Date' specify Date to End. The type is interface{} with range: 1..31.
    EndWeekNumberOrEndDate interface{}

    // If Mode is set to 'Recurring' specify Weekday to end , if Mode is set to
    // 'Date' specify Year to end. The type is interface{} with range: 0..2035.
    EndWeekdayOrEndYear interface{}

    // Month to end . The type is ClockMonth.
    EndMonth interface{}

    // Hour to end . The type is interface{} with range: 0..23.
    EndHour interface{}

    // Minute to end . The type is interface{} with range: 0..59.
    EndMinute interface{}

    // Offset to add in minutes . The type is interface{} with range: 1..1440.
    // Units are minute.
    Offset interface{}
}

func (summerTime *Clock_SummerTime) GetFilter() yfilter.YFilter { return summerTime.YFilter }

func (summerTime *Clock_SummerTime) SetFilter(yf yfilter.YFilter) { summerTime.YFilter = yf }

func (summerTime *Clock_SummerTime) GetGoName(yname string) string {
    if yname == "time-zone-name" { return "TimeZoneName" }
    if yname == "mode" { return "Mode" }
    if yname == "start-week-number-or-start-date" { return "StartWeekNumberOrStartDate" }
    if yname == "start-weekday-or-start-year" { return "StartWeekdayOrStartYear" }
    if yname == "start-month" { return "StartMonth" }
    if yname == "start-hour" { return "StartHour" }
    if yname == "start-minute" { return "StartMinute" }
    if yname == "end-week-number-or-end-date" { return "EndWeekNumberOrEndDate" }
    if yname == "end-weekday-or-end-year" { return "EndWeekdayOrEndYear" }
    if yname == "end-month" { return "EndMonth" }
    if yname == "end-hour" { return "EndHour" }
    if yname == "end-minute" { return "EndMinute" }
    if yname == "offset" { return "Offset" }
    return ""
}

func (summerTime *Clock_SummerTime) GetSegmentPath() string {
    return "summer-time"
}

func (summerTime *Clock_SummerTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summerTime *Clock_SummerTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summerTime *Clock_SummerTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-zone-name"] = summerTime.TimeZoneName
    leafs["mode"] = summerTime.Mode
    leafs["start-week-number-or-start-date"] = summerTime.StartWeekNumberOrStartDate
    leafs["start-weekday-or-start-year"] = summerTime.StartWeekdayOrStartYear
    leafs["start-month"] = summerTime.StartMonth
    leafs["start-hour"] = summerTime.StartHour
    leafs["start-minute"] = summerTime.StartMinute
    leafs["end-week-number-or-end-date"] = summerTime.EndWeekNumberOrEndDate
    leafs["end-weekday-or-end-year"] = summerTime.EndWeekdayOrEndYear
    leafs["end-month"] = summerTime.EndMonth
    leafs["end-hour"] = summerTime.EndHour
    leafs["end-minute"] = summerTime.EndMinute
    leafs["offset"] = summerTime.Offset
    return leafs
}

func (summerTime *Clock_SummerTime) GetBundleName() string { return "cisco_ios_xr" }

func (summerTime *Clock_SummerTime) GetYangName() string { return "summer-time" }

func (summerTime *Clock_SummerTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summerTime *Clock_SummerTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summerTime *Clock_SummerTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summerTime *Clock_SummerTime) SetParent(parent types.Entity) { summerTime.parent = parent }

func (summerTime *Clock_SummerTime) GetParent() types.Entity { return summerTime.parent }

func (summerTime *Clock_SummerTime) GetParentYangName() string { return "clock" }

// Clock_TimeZone
// Configure time zone
// This type is a presence type.
type Clock_TimeZone struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of time zone. The type is string. This attribute is mandatory.
    TimeZoneName interface{}

    // Hours offset from UTC. The type is interface{} with range: -23..23. This
    // attribute is mandatory. Units are hour.
    HourOffset interface{}

    // Minutes offset from UTC. The type is interface{} with range: 0..59. Units
    // are minute. The default value is 0.
    MinuteOffset interface{}
}

func (timeZone *Clock_TimeZone) GetFilter() yfilter.YFilter { return timeZone.YFilter }

func (timeZone *Clock_TimeZone) SetFilter(yf yfilter.YFilter) { timeZone.YFilter = yf }

func (timeZone *Clock_TimeZone) GetGoName(yname string) string {
    if yname == "time-zone-name" { return "TimeZoneName" }
    if yname == "hour-offset" { return "HourOffset" }
    if yname == "minute-offset" { return "MinuteOffset" }
    return ""
}

func (timeZone *Clock_TimeZone) GetSegmentPath() string {
    return "time-zone"
}

func (timeZone *Clock_TimeZone) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timeZone *Clock_TimeZone) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timeZone *Clock_TimeZone) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-zone-name"] = timeZone.TimeZoneName
    leafs["hour-offset"] = timeZone.HourOffset
    leafs["minute-offset"] = timeZone.MinuteOffset
    return leafs
}

func (timeZone *Clock_TimeZone) GetBundleName() string { return "cisco_ios_xr" }

func (timeZone *Clock_TimeZone) GetYangName() string { return "time-zone" }

func (timeZone *Clock_TimeZone) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeZone *Clock_TimeZone) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeZone *Clock_TimeZone) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeZone *Clock_TimeZone) SetParent(parent types.Entity) { timeZone.parent = parent }

func (timeZone *Clock_TimeZone) GetParent() types.Entity { return timeZone.parent }

func (timeZone *Clock_TimeZone) GetParentYangName() string { return "clock" }

