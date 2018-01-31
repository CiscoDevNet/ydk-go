// This module contains a collection of YANG definitions
// for Cisco IOS-XR prm-server package operational data.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: PRM data
//   prm: prm
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package prm_server_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package prm_server_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-prm-server-oper hardware-module}", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-prm-server-oper:hardware-module", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-prm-server-oper prm}", reflect.TypeOf(Prm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-prm-server-oper:prm", reflect.TypeOf(Prm{}))
}

// HardwareModule
// PRM data
type HardwareModule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of PRM Nodes.
    Nodes HardwareModule_Nodes
}

func (hardwareModule *HardwareModule) GetFilter() yfilter.YFilter { return hardwareModule.YFilter }

func (hardwareModule *HardwareModule) SetFilter(yf yfilter.YFilter) { hardwareModule.YFilter = yf }

func (hardwareModule *HardwareModule) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (hardwareModule *HardwareModule) GetSegmentPath() string {
    return "Cisco-IOS-XR-prm-server-oper:hardware-module"
}

func (hardwareModule *HardwareModule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &hardwareModule.Nodes
    }
    return nil
}

func (hardwareModule *HardwareModule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &hardwareModule.Nodes
    return children
}

func (hardwareModule *HardwareModule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModule *HardwareModule) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModule *HardwareModule) GetYangName() string { return "hardware-module" }

func (hardwareModule *HardwareModule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModule *HardwareModule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModule *HardwareModule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModule *HardwareModule) SetParent(parent types.Entity) { hardwareModule.parent = parent }

func (hardwareModule *HardwareModule) GetParent() types.Entity { return hardwareModule.parent }

func (hardwareModule *HardwareModule) GetParentYangName() string { return "Cisco-IOS-XR-prm-server-oper" }

// HardwareModule_Nodes
// List of PRM Nodes
type HardwareModule_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Information. The type is slice of HardwareModule_Nodes_Node.
    Node []HardwareModule_Nodes_Node
}

func (nodes *HardwareModule_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *HardwareModule_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *HardwareModule_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *HardwareModule_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *HardwareModule_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModule_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *HardwareModule_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *HardwareModule_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *HardwareModule_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *HardwareModule_Nodes) GetYangName() string { return "nodes" }

func (nodes *HardwareModule_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *HardwareModule_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *HardwareModule_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *HardwareModule_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *HardwareModule_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *HardwareModule_Nodes) GetParentYangName() string { return "hardware-module" }

// HardwareModule_Nodes_Node
// Node Information
type HardwareModule_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Server specific.
    Np HardwareModule_Nodes_Node_Np
}

func (node *HardwareModule_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModule_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModule_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "np" { return "Np" }
    return ""
}

func (node *HardwareModule_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *HardwareModule_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np" {
        return &node.Np
    }
    return nil
}

func (node *HardwareModule_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["np"] = &node.Np
    return children
}

func (node *HardwareModule_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *HardwareModule_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModule_Nodes_Node) GetYangName() string { return "node" }

func (node *HardwareModule_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModule_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModule_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModule_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModule_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModule_Nodes_Node) GetParentYangName() string { return "nodes" }

// HardwareModule_Nodes_Node_Np
// Server specific
type HardwareModule_Nodes_Node_Np struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource specific.
    Cpu HardwareModule_Nodes_Node_Np_Cpu

    // Platform drops.
    PlatformDrop HardwareModule_Nodes_Node_Np_PlatformDrop
}

func (np *HardwareModule_Nodes_Node_Np) GetFilter() yfilter.YFilter { return np.YFilter }

func (np *HardwareModule_Nodes_Node_Np) SetFilter(yf yfilter.YFilter) { np.YFilter = yf }

func (np *HardwareModule_Nodes_Node_Np) GetGoName(yname string) string {
    if yname == "cpu" { return "Cpu" }
    if yname == "platform-drop" { return "PlatformDrop" }
    return ""
}

func (np *HardwareModule_Nodes_Node_Np) GetSegmentPath() string {
    return "np"
}

func (np *HardwareModule_Nodes_Node_Np) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpu" {
        return &np.Cpu
    }
    if childYangName == "platform-drop" {
        return &np.PlatformDrop
    }
    return nil
}

func (np *HardwareModule_Nodes_Node_Np) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpu"] = &np.Cpu
    children["platform-drop"] = &np.PlatformDrop
    return children
}

func (np *HardwareModule_Nodes_Node_Np) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (np *HardwareModule_Nodes_Node_Np) GetBundleName() string { return "cisco_ios_xr" }

func (np *HardwareModule_Nodes_Node_Np) GetYangName() string { return "np" }

func (np *HardwareModule_Nodes_Node_Np) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (np *HardwareModule_Nodes_Node_Np) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (np *HardwareModule_Nodes_Node_Np) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (np *HardwareModule_Nodes_Node_Np) SetParent(parent types.Entity) { np.parent = parent }

func (np *HardwareModule_Nodes_Node_Np) GetParent() types.Entity { return np.parent }

func (np *HardwareModule_Nodes_Node_Np) GetParentYangName() string { return "node" }

// HardwareModule_Nodes_Node_Np_Cpu
// Resource specific
type HardwareModule_Nodes_Node_Np_Cpu struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data for software resource.
    Indexes HardwareModule_Nodes_Node_Np_Cpu_Indexes
}

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetFilter() yfilter.YFilter { return cpu.YFilter }

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) SetFilter(yf yfilter.YFilter) { cpu.YFilter = yf }

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetGoName(yname string) string {
    if yname == "indexes" { return "Indexes" }
    return ""
}

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetSegmentPath() string {
    return "cpu"
}

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "indexes" {
        return &cpu.Indexes
    }
    return nil
}

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["indexes"] = &cpu.Indexes
    return children
}

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetBundleName() string { return "cisco_ios_xr" }

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetYangName() string { return "cpu" }

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) SetParent(parent types.Entity) { cpu.parent = parent }

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetParent() types.Entity { return cpu.parent }

func (cpu *HardwareModule_Nodes_Node_Np_Cpu) GetParentYangName() string { return "np" }

// HardwareModule_Nodes_Node_Np_Cpu_Indexes
// Data for software resource
type HardwareModule_Nodes_Node_Np_Cpu_Indexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Queue Stats. The type is slice of
    // HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index.
    Index []HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index
}

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetFilter() yfilter.YFilter { return indexes.YFilter }

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) SetFilter(yf yfilter.YFilter) { indexes.YFilter = yf }

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    return ""
}

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetSegmentPath() string {
    return "indexes"
}

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "index" {
        for _, c := range indexes.Index {
            if indexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index{}
        indexes.Index = append(indexes.Index, child)
        return &indexes.Index[len(indexes.Index)-1]
    }
    return nil
}

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range indexes.Index {
        children[indexes.Index[i].GetSegmentPath()] = &indexes.Index[i]
    }
    return children
}

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetBundleName() string { return "cisco_ios_xr" }

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetYangName() string { return "indexes" }

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) SetParent(parent types.Entity) { indexes.parent = parent }

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetParent() types.Entity { return indexes.parent }

func (indexes *HardwareModule_Nodes_Node_Np_Cpu_Indexes) GetParentYangName() string { return "cpu" }

// HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index
// Queue Stats
type HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index value. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // CosQ Name. The type is string with length: 0..1024.
    CosQName interface{}

    // CosQ No. The type is interface{} with range: 0..255.
    CosQ interface{}

    // Rx DMA Channel. The type is interface{} with range: 0..4294967295.
    RxChannel interface{}

    // Flow Rate. The type is interface{} with range: 0..4294967295.
    FlowRate interface{}

    // Burst. The type is interface{} with range: 0..4294967295.
    Burst interface{}

    // Accepted. The type is interface{} with range: 0..18446744073709551615.
    Accepted interface{}

    // Dropped. The type is interface{} with range: 0..18446744073709551615.
    Dropped interface{}
}

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetFilter() yfilter.YFilter { return index.YFilter }

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) SetFilter(yf yfilter.YFilter) { index.YFilter = yf }

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "cos-q-name" { return "CosQName" }
    if yname == "cos-q" { return "CosQ" }
    if yname == "rx-channel" { return "RxChannel" }
    if yname == "flow-rate" { return "FlowRate" }
    if yname == "burst" { return "Burst" }
    if yname == "accepted" { return "Accepted" }
    if yname == "dropped" { return "Dropped" }
    return ""
}

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetSegmentPath() string {
    return "index" + "[index='" + fmt.Sprintf("%v", index.Index) + "']"
}

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = index.Index
    leafs["cos-q-name"] = index.CosQName
    leafs["cos-q"] = index.CosQ
    leafs["rx-channel"] = index.RxChannel
    leafs["flow-rate"] = index.FlowRate
    leafs["burst"] = index.Burst
    leafs["accepted"] = index.Accepted
    leafs["dropped"] = index.Dropped
    return leafs
}

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetBundleName() string { return "cisco_ios_xr" }

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetYangName() string { return "index" }

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) SetParent(parent types.Entity) { index.parent = parent }

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetParent() types.Entity { return index.parent }

func (index *HardwareModule_Nodes_Node_Np_Cpu_Indexes_Index) GetParentYangName() string { return "indexes" }

// HardwareModule_Nodes_Node_Np_PlatformDrop
// Platform drops
type HardwareModule_Nodes_Node_Np_PlatformDrop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Captured Packets.
    Indxes HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes

    // Stats for Drop packets.
    Idxes HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes
}

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetFilter() yfilter.YFilter { return platformDrop.YFilter }

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) SetFilter(yf yfilter.YFilter) { platformDrop.YFilter = yf }

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetGoName(yname string) string {
    if yname == "indxes" { return "Indxes" }
    if yname == "idxes" { return "Idxes" }
    return ""
}

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetSegmentPath() string {
    return "platform-drop"
}

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "indxes" {
        return &platformDrop.Indxes
    }
    if childYangName == "idxes" {
        return &platformDrop.Idxes
    }
    return nil
}

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["indxes"] = &platformDrop.Indxes
    children["idxes"] = &platformDrop.Idxes
    return children
}

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetBundleName() string { return "cisco_ios_xr" }

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetYangName() string { return "platform-drop" }

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) SetParent(parent types.Entity) { platformDrop.parent = parent }

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetParent() types.Entity { return platformDrop.parent }

func (platformDrop *HardwareModule_Nodes_Node_Np_PlatformDrop) GetParentYangName() string { return "np" }

// HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes
// Captured Packets
type HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Captured packets. The type is slice of
    // HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx.
    Indx []HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx
}

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetFilter() yfilter.YFilter { return indxes.YFilter }

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) SetFilter(yf yfilter.YFilter) { indxes.YFilter = yf }

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetGoName(yname string) string {
    if yname == "indx" { return "Indx" }
    return ""
}

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetSegmentPath() string {
    return "indxes"
}

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "indx" {
        for _, c := range indxes.Indx {
            if indxes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx{}
        indxes.Indx = append(indxes.Indx, child)
        return &indxes.Indx[len(indxes.Indx)-1]
    }
    return nil
}

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range indxes.Indx {
        children[indxes.Indx[i].GetSegmentPath()] = &indxes.Indx[i]
    }
    return children
}

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetBundleName() string { return "cisco_ios_xr" }

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetYangName() string { return "indxes" }

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) SetParent(parent types.Entity) { indxes.parent = parent }

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetParent() types.Entity { return indxes.parent }

func (indxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes) GetParentYangName() string { return "platform-drop" }

// HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx
// Captured packets
type HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index value. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // Total packets Captured. The type is interface{} with range: 0..4294967295.
    TotalCaptured interface{}

    // Captured Packet. The type is string with length: 0..1024.
    CapturedPak interface{}

    // Packet Index. The type is interface{} with range: 0..255.
    PktIndex interface{}

    // If Handle. The type is interface{} with range: 0..4294967295.
    Ifhandle interface{}

    // Buffer Length. The type is interface{} with range: 0..4294967295.
    BufferLen interface{}

    // Reason Hi. The type is interface{} with range: 0..4294967295.
    ReasonHi interface{}

    // Reason. The type is interface{} with range: 0..4294967295.
    Reason interface{}

    // Year. The type is interface{} with range: 0..18446744073709551615.
    Years interface{}

    // Hours. The type is interface{} with range: 0..18446744073709551615. Units
    // are hour.
    Hours interface{}

    // Days. The type is interface{} with range: 0..18446744073709551615. Units
    // are day.
    Days interface{}

    // Minutes. The type is interface{} with range: 0..18446744073709551615. Units
    // are minute.
    Mins interface{}

    // Seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Secs interface{}
}

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetFilter() yfilter.YFilter { return indx.YFilter }

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) SetFilter(yf yfilter.YFilter) { indx.YFilter = yf }

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "total-captured" { return "TotalCaptured" }
    if yname == "captured-pak" { return "CapturedPak" }
    if yname == "pkt-index" { return "PktIndex" }
    if yname == "ifhandle" { return "Ifhandle" }
    if yname == "buffer-len" { return "BufferLen" }
    if yname == "reason-hi" { return "ReasonHi" }
    if yname == "reason" { return "Reason" }
    if yname == "years" { return "Years" }
    if yname == "hours" { return "Hours" }
    if yname == "days" { return "Days" }
    if yname == "mins" { return "Mins" }
    if yname == "secs" { return "Secs" }
    return ""
}

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetSegmentPath() string {
    return "indx" + "[index='" + fmt.Sprintf("%v", indx.Index) + "']"
}

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = indx.Index
    leafs["total-captured"] = indx.TotalCaptured
    leafs["captured-pak"] = indx.CapturedPak
    leafs["pkt-index"] = indx.PktIndex
    leafs["ifhandle"] = indx.Ifhandle
    leafs["buffer-len"] = indx.BufferLen
    leafs["reason-hi"] = indx.ReasonHi
    leafs["reason"] = indx.Reason
    leafs["years"] = indx.Years
    leafs["hours"] = indx.Hours
    leafs["days"] = indx.Days
    leafs["mins"] = indx.Mins
    leafs["secs"] = indx.Secs
    return leafs
}

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetBundleName() string { return "cisco_ios_xr" }

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetYangName() string { return "indx" }

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) SetParent(parent types.Entity) { indx.parent = parent }

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetParent() types.Entity { return indx.parent }

func (indx *HardwareModule_Nodes_Node_Np_PlatformDrop_Indxes_Indx) GetParentYangName() string { return "indxes" }

// HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes
// Stats for Drop packets
type HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Drop Stats. The type is slice of
    // HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx.
    Idx []HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx
}

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetFilter() yfilter.YFilter { return idxes.YFilter }

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) SetFilter(yf yfilter.YFilter) { idxes.YFilter = yf }

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetGoName(yname string) string {
    if yname == "idx" { return "Idx" }
    return ""
}

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetSegmentPath() string {
    return "idxes"
}

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "idx" {
        for _, c := range idxes.Idx {
            if idxes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx{}
        idxes.Idx = append(idxes.Idx, child)
        return &idxes.Idx[len(idxes.Idx)-1]
    }
    return nil
}

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range idxes.Idx {
        children[idxes.Idx[i].GetSegmentPath()] = &idxes.Idx[i]
    }
    return children
}

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetBundleName() string { return "cisco_ios_xr" }

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetYangName() string { return "idxes" }

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) SetParent(parent types.Entity) { idxes.parent = parent }

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetParent() types.Entity { return idxes.parent }

func (idxes *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes) GetParentYangName() string { return "platform-drop" }

// HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx
// Drop Stats
type HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index value. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // Drop Reason. The type is string with length: 0..1024.
    DropReason interface{}

    // Counter. The type is interface{} with range: 0..4294967295.
    Counters interface{}
}

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetFilter() yfilter.YFilter { return idx.YFilter }

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) SetFilter(yf yfilter.YFilter) { idx.YFilter = yf }

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "drop-reason" { return "DropReason" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetSegmentPath() string {
    return "idx" + "[index='" + fmt.Sprintf("%v", idx.Index) + "']"
}

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = idx.Index
    leafs["drop-reason"] = idx.DropReason
    leafs["counters"] = idx.Counters
    return leafs
}

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetBundleName() string { return "cisco_ios_xr" }

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetYangName() string { return "idx" }

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) SetParent(parent types.Entity) { idx.parent = parent }

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetParent() types.Entity { return idx.parent }

func (idx *HardwareModule_Nodes_Node_Np_PlatformDrop_Idxes_Idx) GetParentYangName() string { return "idxes" }

// Prm
// prm
type Prm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of PRM Nodes.
    Nodes Prm_Nodes
}

func (prm *Prm) GetFilter() yfilter.YFilter { return prm.YFilter }

func (prm *Prm) SetFilter(yf yfilter.YFilter) { prm.YFilter = yf }

func (prm *Prm) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (prm *Prm) GetSegmentPath() string {
    return "Cisco-IOS-XR-prm-server-oper:prm"
}

func (prm *Prm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &prm.Nodes
    }
    return nil
}

func (prm *Prm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &prm.Nodes
    return children
}

func (prm *Prm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prm *Prm) GetBundleName() string { return "cisco_ios_xr" }

func (prm *Prm) GetYangName() string { return "prm" }

func (prm *Prm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prm *Prm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prm *Prm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prm *Prm) SetParent(parent types.Entity) { prm.parent = parent }

func (prm *Prm) GetParent() types.Entity { return prm.parent }

func (prm *Prm) GetParentYangName() string { return "Cisco-IOS-XR-prm-server-oper" }

// Prm_Nodes
// List of PRM Nodes
type Prm_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Information. The type is slice of Prm_Nodes_Node.
    Node []Prm_Nodes_Node
}

func (nodes *Prm_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Prm_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Prm_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Prm_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Prm_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Prm_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Prm_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Prm_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Prm_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Prm_Nodes) GetYangName() string { return "nodes" }

func (nodes *Prm_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Prm_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Prm_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Prm_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Prm_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Prm_Nodes) GetParentYangName() string { return "prm" }

// Prm_Nodes_Node
// Node Information
type Prm_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Server specific.
    Server Prm_Nodes_Node_Server
}

func (node *Prm_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Prm_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Prm_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "server" { return "Server" }
    return ""
}

func (node *Prm_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Prm_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server" {
        return &node.Server
    }
    return nil
}

func (node *Prm_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["server"] = &node.Server
    return children
}

func (node *Prm_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Prm_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Prm_Nodes_Node) GetYangName() string { return "node" }

func (node *Prm_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Prm_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Prm_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Prm_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Prm_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Prm_Nodes_Node) GetParentYangName() string { return "nodes" }

// Prm_Nodes_Node_Server
// Server specific
type Prm_Nodes_Node_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource specific.
    Resource Prm_Nodes_Node_Server_Resource
}

func (server *Prm_Nodes_Node_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Prm_Nodes_Node_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Prm_Nodes_Node_Server) GetGoName(yname string) string {
    if yname == "resource" { return "Resource" }
    return ""
}

func (server *Prm_Nodes_Node_Server) GetSegmentPath() string {
    return "server"
}

func (server *Prm_Nodes_Node_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "resource" {
        return &server.Resource
    }
    return nil
}

func (server *Prm_Nodes_Node_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["resource"] = &server.Resource
    return children
}

func (server *Prm_Nodes_Node_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (server *Prm_Nodes_Node_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Prm_Nodes_Node_Server) GetYangName() string { return "server" }

func (server *Prm_Nodes_Node_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Prm_Nodes_Node_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Prm_Nodes_Node_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Prm_Nodes_Node_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Prm_Nodes_Node_Server) GetParent() types.Entity { return server.parent }

func (server *Prm_Nodes_Node_Server) GetParentYangName() string { return "node" }

// Prm_Nodes_Node_Server_Resource
// Resource specific
type Prm_Nodes_Node_Server_Resource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data for software resource.
    Indexes Prm_Nodes_Node_Server_Resource_Indexes
}

func (resource *Prm_Nodes_Node_Server_Resource) GetFilter() yfilter.YFilter { return resource.YFilter }

func (resource *Prm_Nodes_Node_Server_Resource) SetFilter(yf yfilter.YFilter) { resource.YFilter = yf }

func (resource *Prm_Nodes_Node_Server_Resource) GetGoName(yname string) string {
    if yname == "indexes" { return "Indexes" }
    return ""
}

func (resource *Prm_Nodes_Node_Server_Resource) GetSegmentPath() string {
    return "resource"
}

func (resource *Prm_Nodes_Node_Server_Resource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "indexes" {
        return &resource.Indexes
    }
    return nil
}

func (resource *Prm_Nodes_Node_Server_Resource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["indexes"] = &resource.Indexes
    return children
}

func (resource *Prm_Nodes_Node_Server_Resource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (resource *Prm_Nodes_Node_Server_Resource) GetBundleName() string { return "cisco_ios_xr" }

func (resource *Prm_Nodes_Node_Server_Resource) GetYangName() string { return "resource" }

func (resource *Prm_Nodes_Node_Server_Resource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resource *Prm_Nodes_Node_Server_Resource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resource *Prm_Nodes_Node_Server_Resource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resource *Prm_Nodes_Node_Server_Resource) SetParent(parent types.Entity) { resource.parent = parent }

func (resource *Prm_Nodes_Node_Server_Resource) GetParent() types.Entity { return resource.parent }

func (resource *Prm_Nodes_Node_Server_Resource) GetParentYangName() string { return "server" }

// Prm_Nodes_Node_Server_Resource_Indexes
// Data for software resource
type Prm_Nodes_Node_Server_Resource_Indexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data for software resource. The type is slice of
    // Prm_Nodes_Node_Server_Resource_Indexes_Index.
    Index []Prm_Nodes_Node_Server_Resource_Indexes_Index
}

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetFilter() yfilter.YFilter { return indexes.YFilter }

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) SetFilter(yf yfilter.YFilter) { indexes.YFilter = yf }

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    return ""
}

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetSegmentPath() string {
    return "indexes"
}

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "index" {
        for _, c := range indexes.Index {
            if indexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Prm_Nodes_Node_Server_Resource_Indexes_Index{}
        indexes.Index = append(indexes.Index, child)
        return &indexes.Index[len(indexes.Index)-1]
    }
    return nil
}

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range indexes.Index {
        children[indexes.Index[i].GetSegmentPath()] = &indexes.Index[i]
    }
    return children
}

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetBundleName() string { return "cisco_ios_xr" }

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetYangName() string { return "indexes" }

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) SetParent(parent types.Entity) { indexes.parent = parent }

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetParent() types.Entity { return indexes.parent }

func (indexes *Prm_Nodes_Node_Server_Resource_Indexes) GetParentYangName() string { return "resource" }

// Prm_Nodes_Node_Server_Resource_Indexes_Index
// Data for software resource
type Prm_Nodes_Node_Server_Resource_Indexes_Index struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Index value. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // Resource Name. The type is string with length: 0..1024.
    ResourceName interface{}

    // Resource Type. The type is interface{} with range: 0..4294967295.
    ResourceType interface{}

    // Total Resource Count. The type is interface{} with range: 0..4294967295.
    TotalNum interface{}

    // Free Resource Count. The type is interface{} with range: 0..4294967295.
    FreeNum interface{}

    // Next Free Index. The type is interface{} with range: 0..4294967295.
    FirstAvailableIndex interface{}

    // Start Index. The type is interface{} with range: 0..4294967295.
    StartIndex interface{}

    // Availability Status. The type is bool.
    AvailabilityStatus interface{}

    // Resource Flags. The type is interface{} with range: 0..255.
    Flags interface{}

    // Inconsistice Flags. The type is bool.
    Inconsistent interface{}
}

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetFilter() yfilter.YFilter { return index.YFilter }

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) SetFilter(yf yfilter.YFilter) { index.YFilter = yf }

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "resource-name" { return "ResourceName" }
    if yname == "resource-type" { return "ResourceType" }
    if yname == "total-num" { return "TotalNum" }
    if yname == "free-num" { return "FreeNum" }
    if yname == "first-available-index" { return "FirstAvailableIndex" }
    if yname == "start-index" { return "StartIndex" }
    if yname == "availability-status" { return "AvailabilityStatus" }
    if yname == "flags" { return "Flags" }
    if yname == "inconsistent" { return "Inconsistent" }
    return ""
}

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetSegmentPath() string {
    return "index" + "[index='" + fmt.Sprintf("%v", index.Index) + "']"
}

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = index.Index
    leafs["resource-name"] = index.ResourceName
    leafs["resource-type"] = index.ResourceType
    leafs["total-num"] = index.TotalNum
    leafs["free-num"] = index.FreeNum
    leafs["first-available-index"] = index.FirstAvailableIndex
    leafs["start-index"] = index.StartIndex
    leafs["availability-status"] = index.AvailabilityStatus
    leafs["flags"] = index.Flags
    leafs["inconsistent"] = index.Inconsistent
    return leafs
}

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetBundleName() string { return "cisco_ios_xr" }

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetYangName() string { return "index" }

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) SetParent(parent types.Entity) { index.parent = parent }

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetParent() types.Entity { return index.parent }

func (index *Prm_Nodes_Node_Server_Resource_Indexes_Index) GetParentYangName() string { return "indexes" }

