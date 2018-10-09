// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-infra-clock-linux package configuration.
// 
// This module contains definitions
// for the following management objects:
//   clock: Configure time-of-day clock
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure time zone.
    TimeZone Clock_TimeZone
}

func (clock *Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "Cisco-IOS-XR-infra-infra-clock-linux-cfg"
    clock.EntityData.SegmentPath = "Cisco-IOS-XR-infra-infra-clock-linux-cfg:clock"
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Children.Append("time-zone", types.YChild{"TimeZone", &clock.TimeZone})
    clock.EntityData.Leafs = types.NewOrderedMap()

    clock.EntityData.YListKeys = []string {}

    return &(clock.EntityData)
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

    // Area File in zoneinfo package. The type is string. This attribute is
    // mandatory.
    AreaName interface{}
}

func (timeZone *Clock_TimeZone) GetEntityData() *types.CommonEntityData {
    timeZone.EntityData.YFilter = timeZone.YFilter
    timeZone.EntityData.YangName = "time-zone"
    timeZone.EntityData.BundleName = "cisco_ios_xr"
    timeZone.EntityData.ParentYangName = "clock"
    timeZone.EntityData.SegmentPath = "time-zone"
    timeZone.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeZone.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeZone.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeZone.EntityData.Children = types.NewOrderedMap()
    timeZone.EntityData.Leafs = types.NewOrderedMap()
    timeZone.EntityData.Leafs.Append("time-zone-name", types.YLeaf{"TimeZoneName", timeZone.TimeZoneName})
    timeZone.EntityData.Leafs.Append("area-name", types.YLeaf{"AreaName", timeZone.AreaName})

    timeZone.EntityData.YListKeys = []string {}

    return &(timeZone.EntityData)
}

