// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-infra-clock-linux package configuration.
// 
// This module contains definitions
// for the following management objects:
//   clock: Configure time-of-day clock
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_infra_clock_linux_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_infra_clock_linux_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-infra-clock-linux-cfg clock}", reflect.TypeOf(Clock{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-infra-clock-linux-cfg:clock", reflect.TypeOf(Clock{}))
}

// Clock
// Configure time-of-day clock
type Clock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure time zone.
    TimeZone Clock_TimeZone
}

func (clock *Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *Clock) GetGoName(yname string) string {
    if yname == "time-zone" { return "TimeZone" }
    return ""
}

func (clock *Clock) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-infra-clock-linux-cfg:clock"
}

func (clock *Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "time-zone" {
        return &clock.TimeZone
    }
    return nil
}

func (clock *Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
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

func (clock *Clock) GetParentYangName() string { return "Cisco-IOS-XR-infra-infra-clock-linux-cfg" }

// Clock_TimeZone
// Configure time zone
// This type is a presence type.
type Clock_TimeZone struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of time zone. The type is string. This attribute is mandatory.
    TimeZoneName interface{}

    // Area File in zoneinfo package. The type is string. This attribute is
    // mandatory.
    AreaName interface{}
}

func (timeZone *Clock_TimeZone) GetFilter() yfilter.YFilter { return timeZone.YFilter }

func (timeZone *Clock_TimeZone) SetFilter(yf yfilter.YFilter) { timeZone.YFilter = yf }

func (timeZone *Clock_TimeZone) GetGoName(yname string) string {
    if yname == "time-zone-name" { return "TimeZoneName" }
    if yname == "area-name" { return "AreaName" }
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
    leafs["area-name"] = timeZone.AreaName
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

