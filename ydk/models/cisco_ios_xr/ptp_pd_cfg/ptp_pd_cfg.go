// This module contains a collection of YANG definitions
// for Cisco IOS-XR ptp-pd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   log-servo-root: Servo Log for Platform
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ptp_pd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ptp_pd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ptp-pd-cfg log-servo-root}", reflect.TypeOf(LogServoRoot{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ptp-pd-cfg:log-servo-root", reflect.TypeOf(LogServoRoot{}))
}

// LogServoRoot
// Servo Log for Platform
type LogServoRoot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Servo change log events. The type is bool.
    ServoEventEnable interface{}
}

func (logServoRoot *LogServoRoot) GetEntityData() *types.CommonEntityData {
    logServoRoot.EntityData.YFilter = logServoRoot.YFilter
    logServoRoot.EntityData.YangName = "log-servo-root"
    logServoRoot.EntityData.BundleName = "cisco_ios_xr"
    logServoRoot.EntityData.ParentYangName = "Cisco-IOS-XR-ptp-pd-cfg"
    logServoRoot.EntityData.SegmentPath = "Cisco-IOS-XR-ptp-pd-cfg:log-servo-root"
    logServoRoot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logServoRoot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logServoRoot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logServoRoot.EntityData.Children = types.NewOrderedMap()
    logServoRoot.EntityData.Leafs = types.NewOrderedMap()
    logServoRoot.EntityData.Leafs.Append("servo-event-enable", types.YLeaf{"ServoEventEnable", logServoRoot.ServoEventEnable})

    logServoRoot.EntityData.YListKeys = []string {}

    return &(logServoRoot.EntityData)
}

