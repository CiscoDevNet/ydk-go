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
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    watchdog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    watchdog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    watchdog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    watchdog.EntityData.Children = make(map[string]types.YChild)
    watchdog.EntityData.Children["threshold-memory"] = types.YChild{"ThresholdMemory", &watchdog.ThresholdMemory}
    watchdog.EntityData.Leafs = make(map[string]types.YLeaf)
    watchdog.EntityData.Leafs["threshold-memory-switchover"] = types.YLeaf{"ThresholdMemorySwitchover", watchdog.ThresholdMemorySwitchover}
    watchdog.EntityData.Leafs["restart-deadlock-disable"] = types.YLeaf{"RestartDeadlockDisable", watchdog.RestartDeadlockDisable}
    watchdog.EntityData.Leafs["monitor-qnet-timeout"] = types.YLeaf{"MonitorQnetTimeout", watchdog.MonitorQnetTimeout}
    watchdog.EntityData.Leafs["monitor-cpuhog-timeout"] = types.YLeaf{"MonitorCpuhogTimeout", watchdog.MonitorCpuhogTimeout}
    watchdog.EntityData.Leafs["monitor-procnto-timeout"] = types.YLeaf{"MonitorProcntoTimeout", watchdog.MonitorProcntoTimeout}
    watchdog.EntityData.Leafs["overload-notification"] = types.YLeaf{"OverloadNotification", watchdog.OverloadNotification}
    watchdog.EntityData.Leafs["restart-cpuhog-disable"] = types.YLeaf{"RestartCpuhogDisable", watchdog.RestartCpuhogDisable}
    watchdog.EntityData.Leafs["restart-memoryhog-disable"] = types.YLeaf{"RestartMemoryhogDisable", watchdog.RestartMemoryhogDisable}
    watchdog.EntityData.Leafs["overload-throttle-timeout"] = types.YLeaf{"OverloadThrottleTimeout", watchdog.OverloadThrottleTimeout}
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
    thresholdMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdMemory.EntityData.Children = make(map[string]types.YChild)
    thresholdMemory.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdMemory.EntityData.Leafs["minor"] = types.YLeaf{"Minor", thresholdMemory.Minor}
    thresholdMemory.EntityData.Leafs["severe"] = types.YLeaf{"Severe", thresholdMemory.Severe}
    thresholdMemory.EntityData.Leafs["critical"] = types.YLeaf{"Critical", thresholdMemory.Critical}
    return &(thresholdMemory.EntityData)
}

