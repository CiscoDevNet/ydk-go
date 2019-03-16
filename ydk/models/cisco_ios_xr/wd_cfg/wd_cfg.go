// This module contains a collection of YANG definitions
// for Cisco IOS-XR wd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   watchdog: watchdog
// 
// This YANG module augments the
//   Cisco-IOS-XR-config-mda-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package wd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package wd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-wd-cfg watchdog}", reflect.TypeOf(Watchdog{}))
    ydk.RegisterEntity("Cisco-IOS-XR-wd-cfg:watchdog", reflect.TypeOf(Watchdog{}))
}

// Watchdog
// watchdog
type Watchdog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // switchover the RP at configured memory state. The type is interface{} with
    // range: 2..4.
    ThresholdMemorySwitchover interface{}

    // Disable watchdog restart deadlock. The type is interface{}.
    RestartDeadlockDisable interface{}

    // Watchdog monitor transport qnet timeout. The type is interface{} with
    // range: 10..3600. Units are second.
    MonitorQnetTimeout interface{}

    // Watchdog monitor cpu-hog persistent timeout configuration. The type is
    // interface{} with range: 10..3600. Units are second.
    MonitorCpuhogTimeout interface{}

    // Watchdog monitor procnto timeout configuration. The type is interface{}
    // with range: 60..3600. Units are second.
    MonitorProcntoTimeout interface{}

    // Disable critical event notification. The type is interface{}.
    OverloadNotification interface{}

    // Disable watchdog restart cpu-hog. The type is interface{}.
    RestartCpuhogDisable interface{}

    // Disable watchdog restart memory-hog. The type is interface{}.
    RestartMemoryhogDisable interface{}

    // Watchdog overload throttle timeout configuration. The type is interface{}
    // with range: 5..120. Units are second.
    OverloadThrottleTimeout interface{}

    // Memory thresholds.
    ThresholdMemory Watchdog_ThresholdMemory
}

func (watchdog *Watchdog) GetEntityData() *types.CommonEntityData {
    watchdog.EntityData.YFilter = watchdog.YFilter
    watchdog.EntityData.YangName = "watchdog"
    watchdog.EntityData.BundleName = "cisco_ios_xr"
    watchdog.EntityData.ParentYangName = "Cisco-IOS-XR-wd-cfg"
    watchdog.EntityData.SegmentPath = "Cisco-IOS-XR-wd-cfg:watchdog"
    watchdog.EntityData.AbsolutePath = watchdog.EntityData.SegmentPath
    watchdog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    watchdog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    watchdog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    watchdog.EntityData.Children = types.NewOrderedMap()
    watchdog.EntityData.Children.Append("threshold-memory", types.YChild{"ThresholdMemory", &watchdog.ThresholdMemory})
    watchdog.EntityData.Leafs = types.NewOrderedMap()
    watchdog.EntityData.Leafs.Append("threshold-memory-switchover", types.YLeaf{"ThresholdMemorySwitchover", watchdog.ThresholdMemorySwitchover})
    watchdog.EntityData.Leafs.Append("restart-deadlock-disable", types.YLeaf{"RestartDeadlockDisable", watchdog.RestartDeadlockDisable})
    watchdog.EntityData.Leafs.Append("monitor-qnet-timeout", types.YLeaf{"MonitorQnetTimeout", watchdog.MonitorQnetTimeout})
    watchdog.EntityData.Leafs.Append("monitor-cpuhog-timeout", types.YLeaf{"MonitorCpuhogTimeout", watchdog.MonitorCpuhogTimeout})
    watchdog.EntityData.Leafs.Append("monitor-procnto-timeout", types.YLeaf{"MonitorProcntoTimeout", watchdog.MonitorProcntoTimeout})
    watchdog.EntityData.Leafs.Append("overload-notification", types.YLeaf{"OverloadNotification", watchdog.OverloadNotification})
    watchdog.EntityData.Leafs.Append("restart-cpuhog-disable", types.YLeaf{"RestartCpuhogDisable", watchdog.RestartCpuhogDisable})
    watchdog.EntityData.Leafs.Append("restart-memoryhog-disable", types.YLeaf{"RestartMemoryhogDisable", watchdog.RestartMemoryhogDisable})
    watchdog.EntityData.Leafs.Append("overload-throttle-timeout", types.YLeaf{"OverloadThrottleTimeout", watchdog.OverloadThrottleTimeout})

    watchdog.EntityData.YListKeys = []string {}

    return &(watchdog.EntityData)
}

// Watchdog_ThresholdMemory
// Memory thresholds
type Watchdog_ThresholdMemory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (thresholdMemory *Watchdog_ThresholdMemory) GetEntityData() *types.CommonEntityData {
    thresholdMemory.EntityData.YFilter = thresholdMemory.YFilter
    thresholdMemory.EntityData.YangName = "threshold-memory"
    thresholdMemory.EntityData.BundleName = "cisco_ios_xr"
    thresholdMemory.EntityData.ParentYangName = "watchdog"
    thresholdMemory.EntityData.SegmentPath = "threshold-memory"
    thresholdMemory.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-cfg:watchdog/" + thresholdMemory.EntityData.SegmentPath
    thresholdMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdMemory.EntityData.Children = types.NewOrderedMap()
    thresholdMemory.EntityData.Leafs = types.NewOrderedMap()
    thresholdMemory.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", thresholdMemory.Minor})
    thresholdMemory.EntityData.Leafs.Append("severe", types.YLeaf{"Severe", thresholdMemory.Severe})
    thresholdMemory.EntityData.Leafs.Append("critical", types.YLeaf{"Critical", thresholdMemory.Critical})

    thresholdMemory.EntityData.YListKeys = []string {}

    return &(thresholdMemory.EntityData)
}

