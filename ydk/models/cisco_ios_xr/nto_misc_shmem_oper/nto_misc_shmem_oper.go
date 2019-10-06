// This module contains a collection of YANG definitions
// for Cisco IOS-XR nto-misc-shmem package operational data.
// 
// This module contains definitions
// for the following management objects:
//   memory-summary: Memory summary information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes.
    Nodes MemorySummary_Nodes
}

func (memorySummary *MemorySummary) GetEntityData() *types.CommonEntityData {
    memorySummary.EntityData.YFilter = memorySummary.YFilter
    memorySummary.EntityData.YangName = "memory-summary"
    memorySummary.EntityData.BundleName = "cisco_ios_xr"
    memorySummary.EntityData.ParentYangName = "Cisco-IOS-XR-nto-misc-shmem-oper"
    memorySummary.EntityData.SegmentPath = "Cisco-IOS-XR-nto-misc-shmem-oper:memory-summary"
    memorySummary.EntityData.AbsolutePath = memorySummary.EntityData.SegmentPath
    memorySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memorySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memorySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memorySummary.EntityData.Children = types.NewOrderedMap()
    memorySummary.EntityData.Children.Append("nodes", types.YChild{"Nodes", &memorySummary.Nodes})
    memorySummary.EntityData.Leafs = types.NewOrderedMap()

    memorySummary.EntityData.YListKeys = []string {}

    return &(memorySummary.EntityData)
}

// MemorySummary_Nodes
// List of nodes
type MemorySummary_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of nodes. The type is slice of MemorySummary_Nodes_Node.
    Node []*MemorySummary_Nodes_Node
}

func (nodes *MemorySummary_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "memory-summary"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-nto-misc-shmem-oper:memory-summary/" + nodes.EntityData.SegmentPath
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

// MemorySummary_Nodes_Node
// Name of nodes
type MemorySummary_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Memory summary information for a specific node.
    Summary MemorySummary_Nodes_Node_Summary

    // Detail Memory summary information for a specific node.
    Detail MemorySummary_Nodes_Node_Detail
}

func (node *MemorySummary_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-nto-misc-shmem-oper:memory-summary/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Children.Append("detail", types.YChild{"Detail", &node.Detail})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// MemorySummary_Nodes_Node_Summary
// Memory summary information for a specific node
type MemorySummary_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
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

func (summary *MemorySummary_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-nto-misc-shmem-oper:memory-summary/nodes/node/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("page-size", types.YLeaf{"PageSize", summary.PageSize})
    summary.EntityData.Leafs.Append("ram-memory", types.YLeaf{"RamMemory", summary.RamMemory})
    summary.EntityData.Leafs.Append("free-physical-memory", types.YLeaf{"FreePhysicalMemory", summary.FreePhysicalMemory})
    summary.EntityData.Leafs.Append("system-ram-memory", types.YLeaf{"SystemRamMemory", summary.SystemRamMemory})
    summary.EntityData.Leafs.Append("free-application-memory", types.YLeaf{"FreeApplicationMemory", summary.FreeApplicationMemory})
    summary.EntityData.Leafs.Append("image-memory", types.YLeaf{"ImageMemory", summary.ImageMemory})
    summary.EntityData.Leafs.Append("boot-ram-size", types.YLeaf{"BootRamSize", summary.BootRamSize})
    summary.EntityData.Leafs.Append("reserved-memory", types.YLeaf{"ReservedMemory", summary.ReservedMemory})
    summary.EntityData.Leafs.Append("io-memory", types.YLeaf{"IoMemory", summary.IoMemory})
    summary.EntityData.Leafs.Append("flash-system", types.YLeaf{"FlashSystem", summary.FlashSystem})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// MemorySummary_Nodes_Node_Detail
// Detail Memory summary information for a
// specific node
type MemorySummary_Nodes_Node_Detail struct {
    EntityData types.CommonEntityData
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

    // Total Used. The type is interface{} with range: 0..18446744073709551615.
    TotalUsed interface{}

    // Available Shared windows. The type is slice of
    // MemorySummary_Nodes_Node_Detail_SharedWindow.
    SharedWindow []*MemorySummary_Nodes_Node_Detail_SharedWindow
}

func (detail *MemorySummary_Nodes_Node_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "node"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-nto-misc-shmem-oper:memory-summary/nodes/node/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("shared-window", types.YChild{"SharedWindow", nil})
    for i := range detail.SharedWindow {
        types.SetYListKey(detail.SharedWindow[i], i)
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.SharedWindow[i]), types.YChild{"SharedWindow", detail.SharedWindow[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("page-size", types.YLeaf{"PageSize", detail.PageSize})
    detail.EntityData.Leafs.Append("ram-memory", types.YLeaf{"RamMemory", detail.RamMemory})
    detail.EntityData.Leafs.Append("free-physical-memory", types.YLeaf{"FreePhysicalMemory", detail.FreePhysicalMemory})
    detail.EntityData.Leafs.Append("private-physical-memory", types.YLeaf{"PrivatePhysicalMemory", detail.PrivatePhysicalMemory})
    detail.EntityData.Leafs.Append("system-ram-memory", types.YLeaf{"SystemRamMemory", detail.SystemRamMemory})
    detail.EntityData.Leafs.Append("free-application-memory", types.YLeaf{"FreeApplicationMemory", detail.FreeApplicationMemory})
    detail.EntityData.Leafs.Append("image-memory", types.YLeaf{"ImageMemory", detail.ImageMemory})
    detail.EntityData.Leafs.Append("boot-ram-size", types.YLeaf{"BootRamSize", detail.BootRamSize})
    detail.EntityData.Leafs.Append("reserved-memory", types.YLeaf{"ReservedMemory", detail.ReservedMemory})
    detail.EntityData.Leafs.Append("io-memory", types.YLeaf{"IoMemory", detail.IoMemory})
    detail.EntityData.Leafs.Append("flash-system", types.YLeaf{"FlashSystem", detail.FlashSystem})
    detail.EntityData.Leafs.Append("total-shared-window", types.YLeaf{"TotalSharedWindow", detail.TotalSharedWindow})
    detail.EntityData.Leafs.Append("allocated-memory", types.YLeaf{"AllocatedMemory", detail.AllocatedMemory})
    detail.EntityData.Leafs.Append("program-text", types.YLeaf{"ProgramText", detail.ProgramText})
    detail.EntityData.Leafs.Append("program-data", types.YLeaf{"ProgramData", detail.ProgramData})
    detail.EntityData.Leafs.Append("program-stack", types.YLeaf{"ProgramStack", detail.ProgramStack})
    detail.EntityData.Leafs.Append("total-used", types.YLeaf{"TotalUsed", detail.TotalUsed})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// MemorySummary_Nodes_Node_Detail_SharedWindow
// Available Shared windows
type MemorySummary_Nodes_Node_Detail_SharedWindow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Name of shared window. The type is string.
    SharedWindow interface{}

    // Size of shared window. The type is interface{} with range:
    // 0..18446744073709551615.
    WindowSize interface{}
}

func (sharedWindow *MemorySummary_Nodes_Node_Detail_SharedWindow) GetEntityData() *types.CommonEntityData {
    sharedWindow.EntityData.YFilter = sharedWindow.YFilter
    sharedWindow.EntityData.YangName = "shared-window"
    sharedWindow.EntityData.BundleName = "cisco_ios_xr"
    sharedWindow.EntityData.ParentYangName = "detail"
    sharedWindow.EntityData.SegmentPath = "shared-window" + types.AddNoKeyToken(sharedWindow)
    sharedWindow.EntityData.AbsolutePath = "Cisco-IOS-XR-nto-misc-shmem-oper:memory-summary/nodes/node/detail/" + sharedWindow.EntityData.SegmentPath
    sharedWindow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sharedWindow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sharedWindow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sharedWindow.EntityData.Children = types.NewOrderedMap()
    sharedWindow.EntityData.Leafs = types.NewOrderedMap()
    sharedWindow.EntityData.Leafs.Append("shared-window", types.YLeaf{"SharedWindow", sharedWindow.SharedWindow})
    sharedWindow.EntityData.Leafs.Append("window-size", types.YLeaf{"WindowSize", sharedWindow.WindowSize})

    sharedWindow.EntityData.YListKeys = []string {}

    return &(sharedWindow.EntityData)
}

