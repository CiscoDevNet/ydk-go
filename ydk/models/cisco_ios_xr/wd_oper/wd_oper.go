// This module contains a collection of YANG definitions
// for Cisco IOS-XR wd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   watchdog: Watchdog information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes.
    Nodes Watchdog_Nodes
}

func (watchdog *Watchdog) GetFilter() yfilter.YFilter { return watchdog.YFilter }

func (watchdog *Watchdog) SetFilter(yf yfilter.YFilter) { watchdog.YFilter = yf }

func (watchdog *Watchdog) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (watchdog *Watchdog) GetSegmentPath() string {
    return "Cisco-IOS-XR-wd-oper:watchdog"
}

func (watchdog *Watchdog) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &watchdog.Nodes
    }
    return nil
}

func (watchdog *Watchdog) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &watchdog.Nodes
    return children
}

func (watchdog *Watchdog) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
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

func (watchdog *Watchdog) GetParentYangName() string { return "Cisco-IOS-XR-wd-oper" }

// Watchdog_Nodes
// List of nodes
type Watchdog_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node ID. The type is slice of Watchdog_Nodes_Node.
    Node []Watchdog_Nodes_Node
}

func (nodes *Watchdog_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Watchdog_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Watchdog_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Watchdog_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Watchdog_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Watchdog_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Watchdog_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Watchdog_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Watchdog_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Watchdog_Nodes) GetYangName() string { return "nodes" }

func (nodes *Watchdog_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Watchdog_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Watchdog_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Watchdog_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Watchdog_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Watchdog_Nodes) GetParentYangName() string { return "watchdog" }

// Watchdog_Nodes_Node
// Node ID
type Watchdog_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (node *Watchdog_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Watchdog_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Watchdog_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "threshold-memory" { return "ThresholdMemory" }
    if yname == "memory-state" { return "MemoryState" }
    if yname == "overload-state" { return "OverloadState" }
    return ""
}

func (node *Watchdog_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Watchdog_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-memory" {
        return &node.ThresholdMemory
    }
    if childYangName == "memory-state" {
        return &node.MemoryState
    }
    if childYangName == "overload-state" {
        return &node.OverloadState
    }
    return nil
}

func (node *Watchdog_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold-memory"] = &node.ThresholdMemory
    children["memory-state"] = &node.MemoryState
    children["overload-state"] = &node.OverloadState
    return children
}

func (node *Watchdog_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Watchdog_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Watchdog_Nodes_Node) GetYangName() string { return "node" }

func (node *Watchdog_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Watchdog_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Watchdog_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Watchdog_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Watchdog_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Watchdog_Nodes_Node) GetParentYangName() string { return "nodes" }

// Watchdog_Nodes_Node_ThresholdMemory
// Threshold memory
type Watchdog_Nodes_Node_ThresholdMemory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // System default memory.
    Default Watchdog_Nodes_Node_ThresholdMemory_Default

    // Memory configured by user.
    Configured Watchdog_Nodes_Node_ThresholdMemory_Configured
}

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetFilter() yfilter.YFilter { return thresholdMemory.YFilter }

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) SetFilter(yf yfilter.YFilter) { thresholdMemory.YFilter = yf }

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetGoName(yname string) string {
    if yname == "default" { return "Default" }
    if yname == "configured" { return "Configured" }
    return ""
}

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetSegmentPath() string {
    return "threshold-memory"
}

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default" {
        return &thresholdMemory.Default
    }
    if childYangName == "configured" {
        return &thresholdMemory.Configured
    }
    return nil
}

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default"] = &thresholdMemory.Default
    children["configured"] = &thresholdMemory.Configured
    return children
}

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetYangName() string { return "threshold-memory" }

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) SetParent(parent types.Entity) { thresholdMemory.parent = parent }

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetParent() types.Entity { return thresholdMemory.parent }

func (thresholdMemory *Watchdog_Nodes_Node_ThresholdMemory) GetParentYangName() string { return "node" }

// Watchdog_Nodes_Node_ThresholdMemory_Default
// System default memory
type Watchdog_Nodes_Node_ThresholdMemory_Default struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured memory.
    ConfiguredMemory Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory

    // Memory Information.
    Memory Watchdog_Nodes_Node_ThresholdMemory_Default_Memory
}

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetGoName(yname string) string {
    if yname == "configured-memory" { return "ConfiguredMemory" }
    if yname == "memory" { return "Memory" }
    return ""
}

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetSegmentPath() string {
    return "default"
}

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-memory" {
        return &self.ConfiguredMemory
    }
    if childYangName == "memory" {
        return &self.Memory
    }
    return nil
}

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configured-memory"] = &self.ConfiguredMemory
    children["memory"] = &self.Memory
    return children
}

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetBundleName() string { return "cisco_ios_xr" }

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetYangName() string { return "default" }

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) SetParent(parent types.Entity) { self.parent = parent }

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetParent() types.Entity { return self.parent }

func (self *Watchdog_Nodes_Node_ThresholdMemory_Default) GetParentYangName() string { return "threshold-memory" }

// Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory
// Configured memory
type Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory struct {
    parent types.Entity
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

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetFilter() yfilter.YFilter { return configuredMemory.YFilter }

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) SetFilter(yf yfilter.YFilter) { configuredMemory.YFilter = yf }

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetGoName(yname string) string {
    if yname == "minor" { return "Minor" }
    if yname == "severe" { return "Severe" }
    if yname == "critical" { return "Critical" }
    return ""
}

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetSegmentPath() string {
    return "configured-memory"
}

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minor"] = configuredMemory.Minor
    leafs["severe"] = configuredMemory.Severe
    leafs["critical"] = configuredMemory.Critical
    return leafs
}

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetBundleName() string { return "cisco_ios_xr" }

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetYangName() string { return "configured-memory" }

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) SetParent(parent types.Entity) { configuredMemory.parent = parent }

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetParent() types.Entity { return configuredMemory.parent }

func (configuredMemory *Watchdog_Nodes_Node_ThresholdMemory_Default_ConfiguredMemory) GetParentYangName() string { return "default" }

// Watchdog_Nodes_Node_ThresholdMemory_Default_Memory
// Memory Information
type Watchdog_Nodes_Node_ThresholdMemory_Default_Memory struct {
    parent types.Entity
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

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetFilter() yfilter.YFilter { return memory.YFilter }

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) SetFilter(yf yfilter.YFilter) { memory.YFilter = yf }

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetGoName(yname string) string {
    if yname == "physical-memory" { return "PhysicalMemory" }
    if yname == "free-memory" { return "FreeMemory" }
    if yname == "memory-state" { return "MemoryState" }
    return ""
}

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetSegmentPath() string {
    return "memory"
}

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["physical-memory"] = memory.PhysicalMemory
    leafs["free-memory"] = memory.FreeMemory
    leafs["memory-state"] = memory.MemoryState
    return leafs
}

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetBundleName() string { return "cisco_ios_xr" }

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetYangName() string { return "memory" }

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) SetParent(parent types.Entity) { memory.parent = parent }

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetParent() types.Entity { return memory.parent }

func (memory *Watchdog_Nodes_Node_ThresholdMemory_Default_Memory) GetParentYangName() string { return "default" }

// Watchdog_Nodes_Node_ThresholdMemory_Configured
// Memory configured by user
type Watchdog_Nodes_Node_ThresholdMemory_Configured struct {
    parent types.Entity
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

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetFilter() yfilter.YFilter { return configured.YFilter }

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) SetFilter(yf yfilter.YFilter) { configured.YFilter = yf }

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetGoName(yname string) string {
    if yname == "minor" { return "Minor" }
    if yname == "severe" { return "Severe" }
    if yname == "critical" { return "Critical" }
    return ""
}

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetSegmentPath() string {
    return "configured"
}

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minor"] = configured.Minor
    leafs["severe"] = configured.Severe
    leafs["critical"] = configured.Critical
    return leafs
}

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetBundleName() string { return "cisco_ios_xr" }

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetYangName() string { return "configured" }

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) SetParent(parent types.Entity) { configured.parent = parent }

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetParent() types.Entity { return configured.parent }

func (configured *Watchdog_Nodes_Node_ThresholdMemory_Configured) GetParentYangName() string { return "threshold-memory" }

// Watchdog_Nodes_Node_MemoryState
// Memory state
type Watchdog_Nodes_Node_MemoryState struct {
    parent types.Entity
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

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetFilter() yfilter.YFilter { return memoryState.YFilter }

func (memoryState *Watchdog_Nodes_Node_MemoryState) SetFilter(yf yfilter.YFilter) { memoryState.YFilter = yf }

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetGoName(yname string) string {
    if yname == "physical-memory" { return "PhysicalMemory" }
    if yname == "free-memory" { return "FreeMemory" }
    if yname == "memory-state" { return "MemoryState" }
    return ""
}

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetSegmentPath() string {
    return "memory-state"
}

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["physical-memory"] = memoryState.PhysicalMemory
    leafs["free-memory"] = memoryState.FreeMemory
    leafs["memory-state"] = memoryState.MemoryState
    return leafs
}

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetBundleName() string { return "cisco_ios_xr" }

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetYangName() string { return "memory-state" }

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryState *Watchdog_Nodes_Node_MemoryState) SetParent(parent types.Entity) { memoryState.parent = parent }

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetParent() types.Entity { return memoryState.parent }

func (memoryState *Watchdog_Nodes_Node_MemoryState) GetParentYangName() string { return "node" }

// Watchdog_Nodes_Node_OverloadState
// Display overload control state
type Watchdog_Nodes_Node_OverloadState struct {
    parent types.Entity
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
    LastThrottle []Watchdog_Nodes_Node_OverloadState_LastThrottle
}

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetFilter() yfilter.YFilter { return overloadState.YFilter }

func (overloadState *Watchdog_Nodes_Node_OverloadState) SetFilter(yf yfilter.YFilter) { overloadState.YFilter = yf }

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetGoName(yname string) string {
    if yname == "overload-control-notification" { return "OverloadControlNotification" }
    if yname == "default-wdsysmon-throttle" { return "DefaultWdsysmonThrottle" }
    if yname == "configured-wdsysmon-throttle" { return "ConfiguredWdsysmonThrottle" }
    if yname == "current-throttle" { return "CurrentThrottle" }
    if yname == "last-throttle" { return "LastThrottle" }
    return ""
}

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetSegmentPath() string {
    return "overload-state"
}

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "current-throttle" {
        return &overloadState.CurrentThrottle
    }
    if childYangName == "last-throttle" {
        for _, c := range overloadState.LastThrottle {
            if overloadState.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Watchdog_Nodes_Node_OverloadState_LastThrottle{}
        overloadState.LastThrottle = append(overloadState.LastThrottle, child)
        return &overloadState.LastThrottle[len(overloadState.LastThrottle)-1]
    }
    return nil
}

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["current-throttle"] = &overloadState.CurrentThrottle
    for i := range overloadState.LastThrottle {
        children[overloadState.LastThrottle[i].GetSegmentPath()] = &overloadState.LastThrottle[i]
    }
    return children
}

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["overload-control-notification"] = overloadState.OverloadControlNotification
    leafs["default-wdsysmon-throttle"] = overloadState.DefaultWdsysmonThrottle
    leafs["configured-wdsysmon-throttle"] = overloadState.ConfiguredWdsysmonThrottle
    return leafs
}

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetBundleName() string { return "cisco_ios_xr" }

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetYangName() string { return "overload-state" }

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (overloadState *Watchdog_Nodes_Node_OverloadState) SetParent(parent types.Entity) { overloadState.parent = parent }

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetParent() types.Entity { return overloadState.parent }

func (overloadState *Watchdog_Nodes_Node_OverloadState) GetParentYangName() string { return "node" }

// Watchdog_Nodes_Node_OverloadState_CurrentThrottle
// Current throttle information
type Watchdog_Nodes_Node_OverloadState_CurrentThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current throttle duration in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    ThrottleDuration interface{}

    // Current throttle start time in format  :day-of-week month date-of-month
    // HH:MM:SS year eg: Thu Feb 1 18:32:14 2011. The type is string with length:
    // 0..25.
    StartTime interface{}
}

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetFilter() yfilter.YFilter { return currentThrottle.YFilter }

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) SetFilter(yf yfilter.YFilter) { currentThrottle.YFilter = yf }

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetGoName(yname string) string {
    if yname == "throttle-duration" { return "ThrottleDuration" }
    if yname == "start-time" { return "StartTime" }
    return ""
}

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetSegmentPath() string {
    return "current-throttle"
}

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle-duration"] = currentThrottle.ThrottleDuration
    leafs["start-time"] = currentThrottle.StartTime
    return leafs
}

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetYangName() string { return "current-throttle" }

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) SetParent(parent types.Entity) { currentThrottle.parent = parent }

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetParent() types.Entity { return currentThrottle.parent }

func (currentThrottle *Watchdog_Nodes_Node_OverloadState_CurrentThrottle) GetParentYangName() string { return "overload-state" }

// Watchdog_Nodes_Node_OverloadState_LastThrottle
// Last throttle information
type Watchdog_Nodes_Node_OverloadState_LastThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetFilter() yfilter.YFilter { return lastThrottle.YFilter }

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) SetFilter(yf yfilter.YFilter) { lastThrottle.YFilter = yf }

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetGoName(yname string) string {
    if yname == "throttle-duration" { return "ThrottleDuration" }
    if yname == "start-time" { return "StartTime" }
    if yname == "stop-time" { return "StopTime" }
    return ""
}

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetSegmentPath() string {
    return "last-throttle"
}

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle-duration"] = lastThrottle.ThrottleDuration
    leafs["start-time"] = lastThrottle.StartTime
    leafs["stop-time"] = lastThrottle.StopTime
    return leafs
}

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetYangName() string { return "last-throttle" }

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) SetParent(parent types.Entity) { lastThrottle.parent = parent }

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetParent() types.Entity { return lastThrottle.parent }

func (lastThrottle *Watchdog_Nodes_Node_OverloadState_LastThrottle) GetParentYangName() string { return "overload-state" }

