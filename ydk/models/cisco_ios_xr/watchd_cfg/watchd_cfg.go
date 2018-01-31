// This module contains a collection of YANG definitions
// for Cisco IOS-XR watchd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   watchdog: Watchdog configuration commands
//   watchd: watchd
// 
// This YANG module augments the
//   Cisco-IOS-XR-config-mda-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package watchd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package watchd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-watchd-cfg watchdog}", reflect.TypeOf(Watchdog{}))
    ydk.RegisterEntity("Cisco-IOS-XR-watchd-cfg:watchdog", reflect.TypeOf(Watchdog{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-watchd-cfg watchd}", reflect.TypeOf(Watchd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-watchd-cfg:watchd", reflect.TypeOf(Watchd{}))
}

// Watchdog
// Watchdog configuration commands
type Watchdog struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable critical event notification. The type is interface{}.
    OverloadNotification interface{}

    // Disable watchdog restart deadlock. The type is interface{}.
    RestartDeadlockDisable interface{}

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
    if yname == "overload-notification" { return "OverloadNotification" }
    if yname == "restart-deadlock-disable" { return "RestartDeadlockDisable" }
    if yname == "restart-memoryhog-disable" { return "RestartMemoryhogDisable" }
    if yname == "overload-throttle-timeout" { return "OverloadThrottleTimeout" }
    if yname == "threshold-memory" { return "ThresholdMemory" }
    return ""
}

func (watchdog *Watchdog) GetSegmentPath() string {
    return "Cisco-IOS-XR-watchd-cfg:watchdog"
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
    leafs["overload-notification"] = watchdog.OverloadNotification
    leafs["restart-deadlock-disable"] = watchdog.RestartDeadlockDisable
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

func (watchdog *Watchdog) GetParentYangName() string { return "Cisco-IOS-XR-watchd-cfg" }

// Watchdog_ThresholdMemory
// Memory thresholds
type Watchdog_ThresholdMemory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold, Range (5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range (4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range (3, severe). The type is interface{} with range: 3..40.
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

// Watchd
// watchd
type Watchd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Length of timeout in seconds. The type is interface{} with range: 1..10.
    // Units are second.
    Timeout interface{}
}

func (watchd *Watchd) GetFilter() yfilter.YFilter { return watchd.YFilter }

func (watchd *Watchd) SetFilter(yf yfilter.YFilter) { watchd.YFilter = yf }

func (watchd *Watchd) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (watchd *Watchd) GetSegmentPath() string {
    return "Cisco-IOS-XR-watchd-cfg:watchd"
}

func (watchd *Watchd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (watchd *Watchd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (watchd *Watchd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = watchd.Timeout
    return leafs
}

func (watchd *Watchd) GetBundleName() string { return "cisco_ios_xr" }

func (watchd *Watchd) GetYangName() string { return "watchd" }

func (watchd *Watchd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (watchd *Watchd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (watchd *Watchd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (watchd *Watchd) SetParent(parent types.Entity) { watchd.parent = parent }

func (watchd *Watchd) GetParent() types.Entity { return watchd.parent }

func (watchd *Watchd) GetParentYangName() string { return "Cisco-IOS-XR-watchd-cfg" }

