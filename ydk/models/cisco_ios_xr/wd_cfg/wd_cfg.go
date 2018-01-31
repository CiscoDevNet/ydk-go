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
    parent types.Entity
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

func (watchdog *Watchdog) GetFilter() yfilter.YFilter { return watchdog.YFilter }

func (watchdog *Watchdog) SetFilter(yf yfilter.YFilter) { watchdog.YFilter = yf }

func (watchdog *Watchdog) GetGoName(yname string) string {
    if yname == "threshold-memory-switchover" { return "ThresholdMemorySwitchover" }
    if yname == "restart-deadlock-disable" { return "RestartDeadlockDisable" }
    if yname == "monitor-qnet-timeout" { return "MonitorQnetTimeout" }
    if yname == "monitor-cpuhog-timeout" { return "MonitorCpuhogTimeout" }
    if yname == "monitor-procnto-timeout" { return "MonitorProcntoTimeout" }
    if yname == "overload-notification" { return "OverloadNotification" }
    if yname == "restart-cpuhog-disable" { return "RestartCpuhogDisable" }
    if yname == "restart-memoryhog-disable" { return "RestartMemoryhogDisable" }
    if yname == "overload-throttle-timeout" { return "OverloadThrottleTimeout" }
    if yname == "threshold-memory" { return "ThresholdMemory" }
    return ""
}

func (watchdog *Watchdog) GetSegmentPath() string {
    return "Cisco-IOS-XR-wd-cfg:watchdog"
}

func (watchdog *Watchdog) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-memory" {
        return &watchdog.ThresholdMemory
    }
    return nil
}

func (watchdog *Watchdog) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold-memory"] = &watchdog.ThresholdMemory
    return children
}

func (watchdog *Watchdog) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-memory-switchover"] = watchdog.ThresholdMemorySwitchover
    leafs["restart-deadlock-disable"] = watchdog.RestartDeadlockDisable
    leafs["monitor-qnet-timeout"] = watchdog.MonitorQnetTimeout
    leafs["monitor-cpuhog-timeout"] = watchdog.MonitorCpuhogTimeout
    leafs["monitor-procnto-timeout"] = watchdog.MonitorProcntoTimeout
    leafs["overload-notification"] = watchdog.OverloadNotification
    leafs["restart-cpuhog-disable"] = watchdog.RestartCpuhogDisable
    leafs["restart-memoryhog-disable"] = watchdog.RestartMemoryhogDisable
    leafs["overload-throttle-timeout"] = watchdog.OverloadThrottleTimeout
    return leafs
}

func (watchdog *Watchdog) GetBundleName() string { return "cisco_ios_xr" }

func (watchdog *Watchdog) GetYangName() string { return "watchdog" }

func (watchdog *Watchdog) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (watchdog *Watchdog) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (watchdog *Watchdog) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (watchdog *Watchdog) SetParent(parent types.Entity) { watchdog.parent = parent }

func (watchdog *Watchdog) GetParent() types.Entity { return watchdog.parent }

func (watchdog *Watchdog) GetParentYangName() string { return "Cisco-IOS-XR-wd-cfg" }

// Watchdog_ThresholdMemory
// Memory thresholds
type Watchdog_ThresholdMemory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (thresholdMemory *Watchdog_ThresholdMemory) GetFilter() yfilter.YFilter { return thresholdMemory.YFilter }

func (thresholdMemory *Watchdog_ThresholdMemory) SetFilter(yf yfilter.YFilter) { thresholdMemory.YFilter = yf }

func (thresholdMemory *Watchdog_ThresholdMemory) GetGoName(yname string) string {
    if yname == "minor" { return "Minor" }
    if yname == "severe" { return "Severe" }
    if yname == "critical" { return "Critical" }
    return ""
}

func (thresholdMemory *Watchdog_ThresholdMemory) GetSegmentPath() string {
    return "threshold-memory"
}

func (thresholdMemory *Watchdog_ThresholdMemory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholdMemory *Watchdog_ThresholdMemory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholdMemory *Watchdog_ThresholdMemory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minor"] = thresholdMemory.Minor
    leafs["severe"] = thresholdMemory.Severe
    leafs["critical"] = thresholdMemory.Critical
    return leafs
}

func (thresholdMemory *Watchdog_ThresholdMemory) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdMemory *Watchdog_ThresholdMemory) GetYangName() string { return "threshold-memory" }

func (thresholdMemory *Watchdog_ThresholdMemory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdMemory *Watchdog_ThresholdMemory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdMemory *Watchdog_ThresholdMemory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdMemory *Watchdog_ThresholdMemory) SetParent(parent types.Entity) { thresholdMemory.parent = parent }

func (thresholdMemory *Watchdog_ThresholdMemory) GetParent() types.Entity { return thresholdMemory.parent }

func (thresholdMemory *Watchdog_ThresholdMemory) GetParentYangName() string { return "watchdog" }

