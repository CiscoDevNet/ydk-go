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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
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

    // Disk thresholds.
    DiskLimit Watchdog_DiskLimit
}

func (watchdog *Watchdog) GetEntityData() *types.CommonEntityData {
    watchdog.EntityData.YFilter = watchdog.YFilter
    watchdog.EntityData.YangName = "watchdog"
    watchdog.EntityData.BundleName = "cisco_ios_xr"
    watchdog.EntityData.ParentYangName = "Cisco-IOS-XR-watchd-cfg"
    watchdog.EntityData.SegmentPath = "Cisco-IOS-XR-watchd-cfg:watchdog"
    watchdog.EntityData.AbsolutePath = watchdog.EntityData.SegmentPath
    watchdog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    watchdog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    watchdog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    watchdog.EntityData.Children = types.NewOrderedMap()
    watchdog.EntityData.Children.Append("threshold-memory", types.YChild{"ThresholdMemory", &watchdog.ThresholdMemory})
    watchdog.EntityData.Children.Append("disk-limit", types.YChild{"DiskLimit", &watchdog.DiskLimit})
    watchdog.EntityData.Leafs = types.NewOrderedMap()
    watchdog.EntityData.Leafs.Append("overload-notification", types.YLeaf{"OverloadNotification", watchdog.OverloadNotification})
    watchdog.EntityData.Leafs.Append("restart-deadlock-disable", types.YLeaf{"RestartDeadlockDisable", watchdog.RestartDeadlockDisable})
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

    // Threshold, Range (5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range (4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range (3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (thresholdMemory *Watchdog_ThresholdMemory) GetEntityData() *types.CommonEntityData {
    thresholdMemory.EntityData.YFilter = thresholdMemory.YFilter
    thresholdMemory.EntityData.YangName = "threshold-memory"
    thresholdMemory.EntityData.BundleName = "cisco_ios_xr"
    thresholdMemory.EntityData.ParentYangName = "watchdog"
    thresholdMemory.EntityData.SegmentPath = "threshold-memory"
    thresholdMemory.EntityData.AbsolutePath = "Cisco-IOS-XR-watchd-cfg:watchdog/" + thresholdMemory.EntityData.SegmentPath
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

// Watchdog_DiskLimit
// Disk thresholds
type Watchdog_DiskLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold, Range (5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range (4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range (3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (diskLimit *Watchdog_DiskLimit) GetEntityData() *types.CommonEntityData {
    diskLimit.EntityData.YFilter = diskLimit.YFilter
    diskLimit.EntityData.YangName = "disk-limit"
    diskLimit.EntityData.BundleName = "cisco_ios_xr"
    diskLimit.EntityData.ParentYangName = "watchdog"
    diskLimit.EntityData.SegmentPath = "disk-limit"
    diskLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-watchd-cfg:watchdog/" + diskLimit.EntityData.SegmentPath
    diskLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diskLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diskLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diskLimit.EntityData.Children = types.NewOrderedMap()
    diskLimit.EntityData.Leafs = types.NewOrderedMap()
    diskLimit.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", diskLimit.Minor})
    diskLimit.EntityData.Leafs.Append("severe", types.YLeaf{"Severe", diskLimit.Severe})
    diskLimit.EntityData.Leafs.Append("critical", types.YLeaf{"Critical", diskLimit.Critical})

    diskLimit.EntityData.YListKeys = []string {}

    return &(diskLimit.EntityData)
}

// Watchd
// watchd
type Watchd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Length of timeout in seconds. The type is interface{} with range: 1..10.
    // Units are second.
    Timeout interface{}
}

func (watchd *Watchd) GetEntityData() *types.CommonEntityData {
    watchd.EntityData.YFilter = watchd.YFilter
    watchd.EntityData.YangName = "watchd"
    watchd.EntityData.BundleName = "cisco_ios_xr"
    watchd.EntityData.ParentYangName = "Cisco-IOS-XR-watchd-cfg"
    watchd.EntityData.SegmentPath = "Cisco-IOS-XR-watchd-cfg:watchd"
    watchd.EntityData.AbsolutePath = watchd.EntityData.SegmentPath
    watchd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    watchd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    watchd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    watchd.EntityData.Children = types.NewOrderedMap()
    watchd.EntityData.Leafs = types.NewOrderedMap()
    watchd.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", watchd.Timeout})

    watchd.EntityData.YListKeys = []string {}

    return &(watchd.EntityData)
}

