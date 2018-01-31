// This module contains a collection of YANG definitions
// for Cisco IOS-XR nto-misc-shmem package operational data.
// 
// This module contains definitions
// for the following management objects:
//   memory-summary: Memory summary information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package nto_misc_shmem_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package nto_misc_shmem_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-nto-misc-shmem-oper memory-summary}", reflect.TypeOf(MemorySummary{}))
    ydk.RegisterEntity("Cisco-IOS-XR-nto-misc-shmem-oper:memory-summary", reflect.TypeOf(MemorySummary{}))
}

// MemorySummary
// Memory summary information
type MemorySummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes.
    Nodes MemorySummary_Nodes
}

func (memorySummary *MemorySummary) GetFilter() yfilter.YFilter { return memorySummary.YFilter }

func (memorySummary *MemorySummary) SetFilter(yf yfilter.YFilter) { memorySummary.YFilter = yf }

func (memorySummary *MemorySummary) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (memorySummary *MemorySummary) GetSegmentPath() string {
    return "Cisco-IOS-XR-nto-misc-shmem-oper:memory-summary"
}

func (memorySummary *MemorySummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &memorySummary.Nodes
    }
    return nil
}

func (memorySummary *MemorySummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &memorySummary.Nodes
    return children
}

func (memorySummary *MemorySummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memorySummary *MemorySummary) GetBundleName() string { return "cisco_ios_xr" }

func (memorySummary *MemorySummary) GetYangName() string { return "memory-summary" }

func (memorySummary *MemorySummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memorySummary *MemorySummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memorySummary *MemorySummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memorySummary *MemorySummary) SetParent(parent types.Entity) { memorySummary.parent = parent }

func (memorySummary *MemorySummary) GetParent() types.Entity { return memorySummary.parent }

func (memorySummary *MemorySummary) GetParentYangName() string { return "Cisco-IOS-XR-nto-misc-shmem-oper" }

// MemorySummary_Nodes
// List of nodes
type MemorySummary_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of nodes. The type is slice of MemorySummary_Nodes_Node.
    Node []MemorySummary_Nodes_Node
}

func (nodes *MemorySummary_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *MemorySummary_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *MemorySummary_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *MemorySummary_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *MemorySummary_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MemorySummary_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *MemorySummary_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *MemorySummary_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *MemorySummary_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *MemorySummary_Nodes) GetYangName() string { return "nodes" }

func (nodes *MemorySummary_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *MemorySummary_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *MemorySummary_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *MemorySummary_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *MemorySummary_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *MemorySummary_Nodes) GetParentYangName() string { return "memory-summary" }

// MemorySummary_Nodes_Node
// Name of nodes
type MemorySummary_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Memory summary information for a specific node.
    Summary MemorySummary_Nodes_Node_Summary

    // Detail Memory summary information for a specific node.
    Detail MemorySummary_Nodes_Node_Detail
}

func (node *MemorySummary_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *MemorySummary_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *MemorySummary_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "summary" { return "Summary" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (node *MemorySummary_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *MemorySummary_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &node.Summary
    }
    if childYangName == "detail" {
        return &node.Detail
    }
    return nil
}

func (node *MemorySummary_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &node.Summary
    children["detail"] = &node.Detail
    return children
}

func (node *MemorySummary_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *MemorySummary_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *MemorySummary_Nodes_Node) GetYangName() string { return "node" }

func (node *MemorySummary_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *MemorySummary_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *MemorySummary_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *MemorySummary_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *MemorySummary_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *MemorySummary_Nodes_Node) GetParentYangName() string { return "nodes" }

// MemorySummary_Nodes_Node_Summary
// Memory summary information for a specific node
type MemorySummary_Nodes_Node_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Page size in bytes. The type is interface{} with range: 0..4294967295.
    // Units are byte.
    PageSize interface{}

    // Physical memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    RamMemory interface{}

    // Physical memory available in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FreePhysicalMemory interface{}

    // Application memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    SystemRamMemory interface{}

    // Application memory available in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FreeApplicationMemory interface{}

    // Image memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ImageMemory interface{}

    // Boot RAM size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BootRamSize interface{}

    // Reserved memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ReservedMemory interface{}

    // IO memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    IoMemory interface{}

    // Flash System size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FlashSystem interface{}
}

func (summary *MemorySummary_Nodes_Node_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *MemorySummary_Nodes_Node_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *MemorySummary_Nodes_Node_Summary) GetGoName(yname string) string {
    if yname == "page-size" { return "PageSize" }
    if yname == "ram-memory" { return "RamMemory" }
    if yname == "free-physical-memory" { return "FreePhysicalMemory" }
    if yname == "system-ram-memory" { return "SystemRamMemory" }
    if yname == "free-application-memory" { return "FreeApplicationMemory" }
    if yname == "image-memory" { return "ImageMemory" }
    if yname == "boot-ram-size" { return "BootRamSize" }
    if yname == "reserved-memory" { return "ReservedMemory" }
    if yname == "io-memory" { return "IoMemory" }
    if yname == "flash-system" { return "FlashSystem" }
    return ""
}

func (summary *MemorySummary_Nodes_Node_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *MemorySummary_Nodes_Node_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *MemorySummary_Nodes_Node_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *MemorySummary_Nodes_Node_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["page-size"] = summary.PageSize
    leafs["ram-memory"] = summary.RamMemory
    leafs["free-physical-memory"] = summary.FreePhysicalMemory
    leafs["system-ram-memory"] = summary.SystemRamMemory
    leafs["free-application-memory"] = summary.FreeApplicationMemory
    leafs["image-memory"] = summary.ImageMemory
    leafs["boot-ram-size"] = summary.BootRamSize
    leafs["reserved-memory"] = summary.ReservedMemory
    leafs["io-memory"] = summary.IoMemory
    leafs["flash-system"] = summary.FlashSystem
    return leafs
}

func (summary *MemorySummary_Nodes_Node_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *MemorySummary_Nodes_Node_Summary) GetYangName() string { return "summary" }

func (summary *MemorySummary_Nodes_Node_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *MemorySummary_Nodes_Node_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *MemorySummary_Nodes_Node_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *MemorySummary_Nodes_Node_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *MemorySummary_Nodes_Node_Summary) GetParent() types.Entity { return summary.parent }

func (summary *MemorySummary_Nodes_Node_Summary) GetParentYangName() string { return "node" }

// MemorySummary_Nodes_Node_Detail
// Detail Memory summary information for a
// specific node
type MemorySummary_Nodes_Node_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Page size in bytes. The type is interface{} with range: 0..4294967295.
    // Units are byte.
    PageSize interface{}

    // Physical memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    RamMemory interface{}

    // Physical memory available in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FreePhysicalMemory interface{}

    // Private Physical memory in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    PrivatePhysicalMemory interface{}

    // Application memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    SystemRamMemory interface{}

    // Application memory available in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FreeApplicationMemory interface{}

    // Image memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ImageMemory interface{}

    // Boot RAM size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    BootRamSize interface{}

    // Reserved memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ReservedMemory interface{}

    // IO memory size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    IoMemory interface{}

    // Flash System size in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FlashSystem interface{}

    // Total Shared window. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalSharedWindow interface{}

    // Allocated Memory Size. The type is interface{} with range:
    // 0..18446744073709551615.
    AllocatedMemory interface{}

    // Program Text Size. The type is interface{} with range:
    // 0..18446744073709551615.
    ProgramText interface{}

    // Program Data Size. The type is interface{} with range:
    // 0..18446744073709551615.
    ProgramData interface{}

    // Program Stack Size. The type is interface{} with range:
    // 0..18446744073709551615.
    ProgramStack interface{}

    // Available Shared windows. The type is slice of
    // MemorySummary_Nodes_Node_Detail_SharedWindow.
    SharedWindow []MemorySummary_Nodes_Node_Detail_SharedWindow
}

func (detail *MemorySummary_Nodes_Node_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *MemorySummary_Nodes_Node_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *MemorySummary_Nodes_Node_Detail) GetGoName(yname string) string {
    if yname == "page-size" { return "PageSize" }
    if yname == "ram-memory" { return "RamMemory" }
    if yname == "free-physical-memory" { return "FreePhysicalMemory" }
    if yname == "private-physical-memory" { return "PrivatePhysicalMemory" }
    if yname == "system-ram-memory" { return "SystemRamMemory" }
    if yname == "free-application-memory" { return "FreeApplicationMemory" }
    if yname == "image-memory" { return "ImageMemory" }
    if yname == "boot-ram-size" { return "BootRamSize" }
    if yname == "reserved-memory" { return "ReservedMemory" }
    if yname == "io-memory" { return "IoMemory" }
    if yname == "flash-system" { return "FlashSystem" }
    if yname == "total-shared-window" { return "TotalSharedWindow" }
    if yname == "allocated-memory" { return "AllocatedMemory" }
    if yname == "program-text" { return "ProgramText" }
    if yname == "program-data" { return "ProgramData" }
    if yname == "program-stack" { return "ProgramStack" }
    if yname == "shared-window" { return "SharedWindow" }
    return ""
}

func (detail *MemorySummary_Nodes_Node_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *MemorySummary_Nodes_Node_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "shared-window" {
        for _, c := range detail.SharedWindow {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MemorySummary_Nodes_Node_Detail_SharedWindow{}
        detail.SharedWindow = append(detail.SharedWindow, child)
        return &detail.SharedWindow[len(detail.SharedWindow)-1]
    }
    return nil
}

func (detail *MemorySummary_Nodes_Node_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range detail.SharedWindow {
        children[detail.SharedWindow[i].GetSegmentPath()] = &detail.SharedWindow[i]
    }
    return children
}

func (detail *MemorySummary_Nodes_Node_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["page-size"] = detail.PageSize
    leafs["ram-memory"] = detail.RamMemory
    leafs["free-physical-memory"] = detail.FreePhysicalMemory
    leafs["private-physical-memory"] = detail.PrivatePhysicalMemory
    leafs["system-ram-memory"] = detail.SystemRamMemory
    leafs["free-application-memory"] = detail.FreeApplicationMemory
    leafs["image-memory"] = detail.ImageMemory
    leafs["boot-ram-size"] = detail.BootRamSize
    leafs["reserved-memory"] = detail.ReservedMemory
    leafs["io-memory"] = detail.IoMemory
    leafs["flash-system"] = detail.FlashSystem
    leafs["total-shared-window"] = detail.TotalSharedWindow
    leafs["allocated-memory"] = detail.AllocatedMemory
    leafs["program-text"] = detail.ProgramText
    leafs["program-data"] = detail.ProgramData
    leafs["program-stack"] = detail.ProgramStack
    return leafs
}

func (detail *MemorySummary_Nodes_Node_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *MemorySummary_Nodes_Node_Detail) GetYangName() string { return "detail" }

func (detail *MemorySummary_Nodes_Node_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *MemorySummary_Nodes_Node_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *MemorySummary_Nodes_Node_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *MemorySummary_Nodes_Node_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *MemorySummary_Nodes_Node_Detail) GetParent() types.Entity { return detail.parent }

func (detail *MemorySummary_Nodes_Node_Detail) GetParentYangName() string { return "node" }

// MemorySummary_Nodes_Node_Detail_SharedWindow
// Available Shared windows
type MemorySummary_Nodes_Node_Detail_SharedWindow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of shared window. The type is string.
    SharedWindow interface{}

    // Size of shared window. The type is interface{} with range:
    // 0..18446744073709551615.
    WindowSize interface{}
}

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetFilter() yfilter.YFilter { return sharedWindow.YFilter }

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) SetFilter(yf yfilter.YFilter) { sharedWindow.YFilter = yf }

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetGoName(yname string) string {
    if yname == "shared-window" { return "SharedWindow" }
    if yname == "window-size" { return "WindowSize" }
    return ""
}

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetSegmentPath() string {
    return "shared-window"
}

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["shared-window"] = sharedWindow.SharedWindow
    leafs["window-size"] = sharedWindow.WindowSize
    return leafs
}

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetBundleName() string { return "cisco_ios_xr" }

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetYangName() string { return "shared-window" }

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) SetParent(parent types.Entity) { sharedWindow.parent = parent }

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetParent() types.Entity { return sharedWindow.parent }

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetParentYangName() string { return "detail" }

