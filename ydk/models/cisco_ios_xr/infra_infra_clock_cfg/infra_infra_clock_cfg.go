// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-infra-clock package configuration.
// 
// This module contains definitions
// for the following management objects:
//   clock: Configure time-of-day clock
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure summer (daylight savings) time.
    SummerTime Clock_SummerTime

    // Configure time zone.
    TimeZone Clock_TimeZone
}

func (clock *Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "Cisco-IOS-XR-infra-infra-clock-cfg"
    clock.EntityData.SegmentPath = "Cisco-IOS-XR-infra-infra-clock-cfg:clock"
    clock.EntityData.AbsolutePath = clock.EntityData.SegmentPath
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Children.Append("summer-time", types.YChild{"SummerTime", &clock.SummerTime})
    clock.EntityData.Children.Append("time-zone", types.YChild{"TimeZone", &clock.TimeZone})
    clock.EntityData.Leafs = types.NewOrderedMap()

    clock.EntityData.YListKeys = []string {}

    return &(clock.EntityData)
}

// Clock_SummerTime
// Configure summer (daylight savings) time
// This type is a presence type.
type Clock_SummerTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (summerTime *Clock_SummerTime) GetEntityData() *types.CommonEntityData {
    summerTime.EntityData.YFilter = summerTime.YFilter
    summerTime.EntityData.YangName = "summer-time"
    summerTime.EntityData.BundleName = "cisco_ios_xr"
    summerTime.EntityData.ParentYangName = "clock"
    summerTime.EntityData.SegmentPath = "summer-time"
    summerTime.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-infra-clock-cfg:clock/" + summerTime.EntityData.SegmentPath
    summerTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summerTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summerTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summerTime.EntityData.Children = types.NewOrderedMap()
    summerTime.EntityData.Leafs = types.NewOrderedMap()
    summerTime.EntityData.Leafs.Append("time-zone-name", types.YLeaf{"TimeZoneName", summerTime.TimeZoneName})
    summerTime.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", summerTime.Mode})
    summerTime.EntityData.Leafs.Append("start-week-number-or-start-date", types.YLeaf{"StartWeekNumberOrStartDate", summerTime.StartWeekNumberOrStartDate})
    summerTime.EntityData.Leafs.Append("start-weekday-or-start-year", types.YLeaf{"StartWeekdayOrStartYear", summerTime.StartWeekdayOrStartYear})
    summerTime.EntityData.Leafs.Append("start-month", types.YLeaf{"StartMonth", summerTime.StartMonth})
    summerTime.EntityData.Leafs.Append("start-hour", types.YLeaf{"StartHour", summerTime.StartHour})
    summerTime.EntityData.Leafs.Append("start-minute", types.YLeaf{"StartMinute", summerTime.StartMinute})
    summerTime.EntityData.Leafs.Append("end-week-number-or-end-date", types.YLeaf{"EndWeekNumberOrEndDate", summerTime.EndWeekNumberOrEndDate})
    summerTime.EntityData.Leafs.Append("end-weekday-or-end-year", types.YLeaf{"EndWeekdayOrEndYear", summerTime.EndWeekdayOrEndYear})
    summerTime.EntityData.Leafs.Append("end-month", types.YLeaf{"EndMonth", summerTime.EndMonth})
    summerTime.EntityData.Leafs.Append("end-hour", types.YLeaf{"EndHour", summerTime.EndHour})
    summerTime.EntityData.Leafs.Append("end-minute", types.YLeaf{"EndMinute", summerTime.EndMinute})
    summerTime.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", summerTime.Offset})

    summerTime.EntityData.YListKeys = []string {}

    return &(summerTime.EntityData)
}

// Clock_TimeZone
// Configure time zone
// This type is a presence type.
type Clock_TimeZone struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Name of time zone. The type is string. This attribute is mandatory.
    TimeZoneName interface{}

    // Hours offset from UTC. The type is interface{} with range: -23..23. This
    // attribute is mandatory. Units are hour.
    HourOffset interface{}

    // Minutes offset from UTC. The type is interface{} with range: 0..59. Units
    // are minute. The default value is 0.
    MinuteOffset interface{}
}

func (timeZone *Clock_TimeZone) GetEntityData() *types.CommonEntityData {
    timeZone.EntityData.YFilter = timeZone.YFilter
    timeZone.EntityData.YangName = "time-zone"
    timeZone.EntityData.BundleName = "cisco_ios_xr"
    timeZone.EntityData.ParentYangName = "clock"
    timeZone.EntityData.SegmentPath = "time-zone"
    timeZone.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-infra-clock-cfg:clock/" + timeZone.EntityData.SegmentPath
    timeZone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeZone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeZone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeZone.EntityData.Children = types.NewOrderedMap()
    timeZone.EntityData.Leafs = types.NewOrderedMap()
    timeZone.EntityData.Leafs.Append("time-zone-name", types.YLeaf{"TimeZoneName", timeZone.TimeZoneName})
    timeZone.EntityData.Leafs.Append("hour-offset", types.YLeaf{"HourOffset", timeZone.HourOffset})
    timeZone.EntityData.Leafs.Append("minute-offset", types.YLeaf{"MinuteOffset", timeZone.MinuteOffset})

    timeZone.EntityData.YListKeys = []string {}

    return &(timeZone.EntityData)
}

