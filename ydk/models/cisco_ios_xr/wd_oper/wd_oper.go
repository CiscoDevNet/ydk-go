// This module contains a collection of YANG definitions
// for Cisco IOS-XR wd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   watchdog: Watchdog information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package wd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package wd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-wd-oper watchdog}", reflect.TypeOf(Watchdog{}))
    ydk.RegisterEntity("Cisco-IOS-XR-wd-oper:watchdog", reflect.TypeOf(Watchdog{}))
}

// MemoryState represents Memory state options
type MemoryState string

const (
    // Memory state unknown
    MemoryState_unknown MemoryState = "unknown"

    // Memory state normal
    MemoryState_normal MemoryState = "normal"

    // Memory state minor
    MemoryState_minor MemoryState = "minor"

    // Memory state severe
    MemoryState_severe MemoryState = "severe"

    // Memory state critical
    MemoryState_critical MemoryState = "critical"
)

// OverloadCtrlNotif represents Overload control notification
type OverloadCtrlNotif string

const (
    // Diabled
    OverloadCtrlNotif_disabled OverloadCtrlNotif = "disabled"

    // Enabled
    OverloadCtrlNotif_enabled OverloadCtrlNotif = "enabled"
)

// Watchdog
// Watchdog information
type Watchdog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes.
    Nodes Watchdog_Nodes
}

func (watchdog *Watchdog) GetEntityData() *types.CommonEntityData {
    watchdog.EntityData.YFilter = watchdog.YFilter
    watchdog.EntityData.YangName = "watchdog"
    watchdog.EntityData.BundleName = "cisco_ios_xr"
    watchdog.EntityData.ParentYangName = "Cisco-IOS-XR-wd-oper"
    watchdog.EntityData.SegmentPath = "Cisco-IOS-XR-wd-oper:watchdog"
    watchdog.EntityData.AbsolutePath = watchdog.EntityData.SegmentPath
    watchdog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    watchdog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    watchdog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    watchdog.EntityData.Children = types.NewOrderedMap()
    watchdog.EntityData.Children.Append("nodes", types.YChild{"Nodes", &watchdog.Nodes})
    watchdog.EntityData.Leafs = types.NewOrderedMap()

    watchdog.EntityData.YListKeys = []string {}

    return &(watchdog.EntityData)
}

// Watchdog_Nodes
// List of nodes
type Watchdog_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node ID. The type is slice of Watchdog_Nodes_Node.
    Node []*Watchdog_Nodes_Node
}

func (nodes *Watchdog_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "watchdog"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Watchdog_Nodes_Node
// Node ID
type Watchdog_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Threshold memory.
    ThresholdMemory Watchdog_Nodes_Node_ThresholdMemory

    // Memory state.
    MemoryState Watchdog_Nodes_Node_MemoryState

    // Display overload control state.
    OverloadState Watchdog_Nodes_Node_OverloadState
}

func (node *Watchdog_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("threshold-memory", types.YChild{"ThresholdMemory", &node.ThresholdMemory})
    node.EntityData.Children.Append("memory-state", types.YChild{"MemoryState", &node.MemoryState})
    node.EntityData.Children.Append("overload-state", types.YChild{"OverloadState", &node.OverloadState})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Watchdog_Nodes_Node_ThresholdMemory
// Threshold memory
type Watchdog_Nodes_Node_ThresholdMemory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // System default memory.
    Default Watchdog_Nodes_Node_ThresholdMemory_Default

    // Memory configured by user.
    Configured Watchdog_Nodes_Node_ThresholdMemory_Configured
}

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetEntityData() *types.CommonEntityData {
    thresholdMemory.EntityData.YFilter = thresholdMemory.YFilter
    thresholdMemory.EntityData.YangName = "threshold-memory"
    thresholdMemory.EntityData.BundleName = "cisco_ios_xr"
    thresholdMemory.EntityData.ParentYangName = "node"
    thresholdMemory.EntityData.SegmentPath = "threshold-memory"
    thresholdMemory.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/node/" + thresholdMemory.EntityData.SegmentPath
    thresholdMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdMemory.EntityData.Children = types.NewOrderedMap()
    thresholdMemory.EntityData.Children.Append("default", types.YChild{"Default", &thresholdMemory.Default})
    thresholdMemory.EntityData.Children.Append("configured", types.YChild{"Configured", &thresholdMemory.Configured})
    thresholdMemory.EntityData.Leafs = types.NewOrderedMap()

    thresholdMemory.EntityData.YListKeys = []string {}

    return &(thresholdMemory.EntityData)
}

// Watchdog_Nodes_Node_ThresholdMemory_Default
// System default memory
type Watchdog_Nodes_Node_ThresholdMemory_Default struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured memory.
    ConfiguredMemory Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory

    // Memory Information.
    Memory Watchdog_Nodes_Node_ThresholdMemory_Default_Memory
}

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "default"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "threshold-memory"
    self.EntityData.SegmentPath = "default"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/node/threshold-memory/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("configured-memory", types.YChild{"ConfiguredMemory", &self.ConfiguredMemory})
    self.EntityData.Children.Append("memory", types.YChild{"Memory", &self.Memory})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory
// Configured memory
type Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minor memory threshold in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    Minor interface{}

    // Severe memory threshold in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    Severe interface{}

    // Critical memory in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    Critical interface{}
}

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetEntityData() *types.CommonEntityData {
    configuredMemory.EntityData.YFilter = configuredMemory.YFilter
    configuredMemory.EntityData.YangName = "configured-memory"
    configuredMemory.EntityData.BundleName = "cisco_ios_xr"
    configuredMemory.EntityData.ParentYangName = "default"
    configuredMemory.EntityData.SegmentPath = "configured-memory"
    configuredMemory.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/node/threshold-memory/default/" + configuredMemory.EntityData.SegmentPath
    configuredMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredMemory.EntityData.Children = types.NewOrderedMap()
    configuredMemory.EntityData.Leafs = types.NewOrderedMap()
    configuredMemory.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", configuredMemory.Minor})
    configuredMemory.EntityData.Leafs.Append("severe", types.YLeaf{"Severe", configuredMemory.Severe})
    configuredMemory.EntityData.Leafs.Append("critical", types.YLeaf{"Critical", configuredMemory.Critical})

    configuredMemory.EntityData.YListKeys = []string {}

    return &(configuredMemory.EntityData)
}

// Watchdog_Nodes_Node_ThresholdMemory_Default_Memory
// Memory Information
type Watchdog_Nodes_Node_ThresholdMemory_Default_Memory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Physical memory in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    PhysicalMemory interface{}

    // Free memory in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FreeMemory interface{}

    // State of memory. The type is MemoryState.
    MemoryState interface{}
}

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetEntityData() *types.CommonEntityData {
    memory.EntityData.YFilter = memory.YFilter
    memory.EntityData.YangName = "memory"
    memory.EntityData.BundleName = "cisco_ios_xr"
    memory.EntityData.ParentYangName = "default"
    memory.EntityData.SegmentPath = "memory"
    memory.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/node/threshold-memory/default/" + memory.EntityData.SegmentPath
    memory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memory.EntityData.Children = types.NewOrderedMap()
    memory.EntityData.Leafs = types.NewOrderedMap()
    memory.EntityData.Leafs.Append("physical-memory", types.YLeaf{"PhysicalMemory", memory.PhysicalMemory})
    memory.EntityData.Leafs.Append("free-memory", types.YLeaf{"FreeMemory", memory.FreeMemory})
    memory.EntityData.Leafs.Append("memory-state", types.YLeaf{"MemoryState", memory.MemoryState})

    memory.EntityData.YListKeys = []string {}

    return &(memory.EntityData)
}

// Watchdog_Nodes_Node_ThresholdMemory_Configured
// Memory configured by user
type Watchdog_Nodes_Node_ThresholdMemory_Configured struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minor memory threshold in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    Minor interface{}

    // Severe memory threshold in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    Severe interface{}

    // Critical memory in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    Critical interface{}
}

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetEntityData() *types.CommonEntityData {
    configured.EntityData.YFilter = configured.YFilter
    configured.EntityData.YangName = "configured"
    configured.EntityData.BundleName = "cisco_ios_xr"
    configured.EntityData.ParentYangName = "threshold-memory"
    configured.EntityData.SegmentPath = "configured"
    configured.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/node/threshold-memory/" + configured.EntityData.SegmentPath
    configured.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configured.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configured.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configured.EntityData.Children = types.NewOrderedMap()
    configured.EntityData.Leafs = types.NewOrderedMap()
    configured.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", configured.Minor})
    configured.EntityData.Leafs.Append("severe", types.YLeaf{"Severe", configured.Severe})
    configured.EntityData.Leafs.Append("critical", types.YLeaf{"Critical", configured.Critical})

    configured.EntityData.YListKeys = []string {}

    return &(configured.EntityData)
}

// Watchdog_Nodes_Node_MemoryState
// Memory state
type Watchdog_Nodes_Node_MemoryState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Physical memory in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    PhysicalMemory interface{}

    // Free memory in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FreeMemory interface{}

    // State of memory. The type is MemoryState.
    MemoryState interface{}
}

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetEntityData() *types.CommonEntityData {
    memoryState.EntityData.YFilter = memoryState.YFilter
    memoryState.EntityData.YangName = "memory-state"
    memoryState.EntityData.BundleName = "cisco_ios_xr"
    memoryState.EntityData.ParentYangName = "node"
    memoryState.EntityData.SegmentPath = "memory-state"
    memoryState.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/node/" + memoryState.EntityData.SegmentPath
    memoryState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryState.EntityData.Children = types.NewOrderedMap()
    memoryState.EntityData.Leafs = types.NewOrderedMap()
    memoryState.EntityData.Leafs.Append("physical-memory", types.YLeaf{"PhysicalMemory", memoryState.PhysicalMemory})
    memoryState.EntityData.Leafs.Append("free-memory", types.YLeaf{"FreeMemory", memoryState.FreeMemory})
    memoryState.EntityData.Leafs.Append("memory-state", types.YLeaf{"MemoryState", memoryState.MemoryState})

    memoryState.EntityData.YListKeys = []string {}

    return &(memoryState.EntityData)
}

// Watchdog_Nodes_Node_OverloadState
// Display overload control state
type Watchdog_Nodes_Node_OverloadState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State of overload control notification. The type is OverloadCtrlNotif.
    OverloadControlNotification interface{}

    // Default resmon throttle. The type is interface{} with range: 0..4294967295.
    DefaultWdsysmonThrottle interface{}

    // Configured resmon throttle. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredWdsysmonThrottle interface{}

    // Current throttle information.
    CurrentThrottle Watchdog_Nodes_Node_OverloadState_CurrentThrottle

    // Last throttle information. The type is slice of
    // Watchdog_Nodes_Node_OverloadState_LastThrottle.
    LastThrottle []*Watchdog_Nodes_Node_OverloadState_LastThrottle
}

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetEntityData() *types.CommonEntityData {
    overloadState.EntityData.YFilter = overloadState.YFilter
    overloadState.EntityData.YangName = "overload-state"
    overloadState.EntityData.BundleName = "cisco_ios_xr"
    overloadState.EntityData.ParentYangName = "node"
    overloadState.EntityData.SegmentPath = "overload-state"
    overloadState.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/node/" + overloadState.EntityData.SegmentPath
    overloadState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overloadState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overloadState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overloadState.EntityData.Children = types.NewOrderedMap()
    overloadState.EntityData.Children.Append("current-throttle", types.YChild{"CurrentThrottle", &overloadState.CurrentThrottle})
    overloadState.EntityData.Children.Append("last-throttle", types.YChild{"LastThrottle", nil})
    for i := range overloadState.LastThrottle {
        types.SetYListKey(overloadState.LastThrottle[i], i)
        overloadState.EntityData.Children.Append(types.GetSegmentPath(overloadState.LastThrottle[i]), types.YChild{"LastThrottle", overloadState.LastThrottle[i]})
    }
    overloadState.EntityData.Leafs = types.NewOrderedMap()
    overloadState.EntityData.Leafs.Append("overload-control-notification", types.YLeaf{"OverloadControlNotification", overloadState.OverloadControlNotification})
    overloadState.EntityData.Leafs.Append("default-wdsysmon-throttle", types.YLeaf{"DefaultWdsysmonThrottle", overloadState.DefaultWdsysmonThrottle})
    overloadState.EntityData.Leafs.Append("configured-wdsysmon-throttle", types.YLeaf{"ConfiguredWdsysmonThrottle", overloadState.ConfiguredWdsysmonThrottle})

    overloadState.EntityData.YListKeys = []string {}

    return &(overloadState.EntityData)
}

// Watchdog_Nodes_Node_OverloadState_CurrentThrottle
// Current throttle information
type Watchdog_Nodes_Node_OverloadState_CurrentThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current throttle duration in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    ThrottleDuration interface{}

    // Current throttle start time in format :day-of-week month date-of-month
    // HH:MM:SS year eg: Thu Feb 1 18:32:14 2011. The type is string with length:
    // 0..25.
    StartTime interface{}
}

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetEntityData() *types.CommonEntityData {
    currentThrottle.EntityData.YFilter = currentThrottle.YFilter
    currentThrottle.EntityData.YangName = "current-throttle"
    currentThrottle.EntityData.BundleName = "cisco_ios_xr"
    currentThrottle.EntityData.ParentYangName = "overload-state"
    currentThrottle.EntityData.SegmentPath = "current-throttle"
    currentThrottle.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/node/overload-state/" + currentThrottle.EntityData.SegmentPath
    currentThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentThrottle.EntityData.Children = types.NewOrderedMap()
    currentThrottle.EntityData.Leafs = types.NewOrderedMap()
    currentThrottle.EntityData.Leafs.Append("throttle-duration", types.YLeaf{"ThrottleDuration", currentThrottle.ThrottleDuration})
    currentThrottle.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", currentThrottle.StartTime})

    currentThrottle.EntityData.YListKeys = []string {}

    return &(currentThrottle.EntityData)
}

// Watchdog_Nodes_Node_OverloadState_LastThrottle
// Last throttle information
type Watchdog_Nodes_Node_OverloadState_LastThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Last throttle duration in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    ThrottleDuration interface{}

    // Last throttle start time in format :day-of-week month date-of-month
    // HH:MM:SS year eg: Thu Feb 1 18:32:14 2011. The type is string with length:
    // 0..25.
    StartTime interface{}

    // Last throttle stop time in format :day-of-week month date-of-month HH:MM:SS
    // year eg: Thu Feb 1 18:32:14 2011. The type is string with length: 0..25.
    StopTime interface{}
}

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetEntityData() *types.CommonEntityData {
    lastThrottle.EntityData.YFilter = lastThrottle.YFilter
    lastThrottle.EntityData.YangName = "last-throttle"
    lastThrottle.EntityData.BundleName = "cisco_ios_xr"
    lastThrottle.EntityData.ParentYangName = "overload-state"
    lastThrottle.EntityData.SegmentPath = "last-throttle" + types.AddNoKeyToken(lastThrottle)
    lastThrottle.EntityData.AbsolutePath = "Cisco-IOS-XR-wd-oper:watchdog/nodes/node/overload-state/" + lastThrottle.EntityData.SegmentPath
    lastThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastThrottle.EntityData.Children = types.NewOrderedMap()
    lastThrottle.EntityData.Leafs = types.NewOrderedMap()
    lastThrottle.EntityData.Leafs.Append("throttle-duration", types.YLeaf{"ThrottleDuration", lastThrottle.ThrottleDuration})
    lastThrottle.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", lastThrottle.StartTime})
    lastThrottle.EntityData.Leafs.Append("stop-time", types.YLeaf{"StopTime", lastThrottle.StopTime})

    lastThrottle.EntityData.YListKeys = []string {}

    return &(lastThrottle.EntityData)
}

