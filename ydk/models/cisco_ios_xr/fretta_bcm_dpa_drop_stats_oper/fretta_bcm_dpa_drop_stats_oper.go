// This module contains a collection of YANG definitions
// for Cisco IOS-XR fretta-bcm-dpa-drop-stats package operational data.
// 
// This module contains definitions
// for the following management objects:
//   drop: Drop stats data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package fretta_bcm_dpa_drop_stats_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fretta_bcm_dpa_drop_stats_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper drop}", reflect.TypeOf(Drop{}))
    ydk.RegisterEntity("Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper:drop", reflect.TypeOf(Drop{}))
}

// Drop
// Drop stats data
type Drop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Drop data per node.
    Nodes Drop_Nodes
}

func (drop *Drop) GetFilter() yfilter.YFilter { return drop.YFilter }

func (drop *Drop) SetFilter(yf yfilter.YFilter) { drop.YFilter = yf }

func (drop *Drop) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (drop *Drop) GetSegmentPath() string {
    return "Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper:drop"
}

func (drop *Drop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &drop.Nodes
    }
    return nil
}

func (drop *Drop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &drop.Nodes
    return children
}

func (drop *Drop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (drop *Drop) GetBundleName() string { return "cisco_ios_xr" }

func (drop *Drop) GetYangName() string { return "drop" }

func (drop *Drop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (drop *Drop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (drop *Drop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (drop *Drop) SetParent(parent types.Entity) { drop.parent = parent }

func (drop *Drop) GetParent() types.Entity { return drop.parent }

func (drop *Drop) GetParentYangName() string { return "Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper" }

// Drop_Nodes
// Drop data per node
type Drop_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Drop stats data for a particular node. The type is slice of
    // Drop_Nodes_Node.
    Node []Drop_Nodes_Node
}

func (nodes *Drop_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Drop_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Drop_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Drop_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Drop_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Drop_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Drop_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Drop_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Drop_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Drop_Nodes) GetYangName() string { return "nodes" }

func (nodes *Drop_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Drop_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Drop_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Drop_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Drop_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Drop_Nodes) GetParentYangName() string { return "drop" }

// Drop_Nodes_Node
// Drop stats data for a particular node
type Drop_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // NPU drop stats.
    NpuNumberForDropStats Drop_Nodes_Node_NpuNumberForDropStats
}

func (node *Drop_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Drop_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Drop_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "npu-number-for-drop-stats" { return "NpuNumberForDropStats" }
    return ""
}

func (node *Drop_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Drop_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "npu-number-for-drop-stats" {
        return &node.NpuNumberForDropStats
    }
    return nil
}

func (node *Drop_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["npu-number-for-drop-stats"] = &node.NpuNumberForDropStats
    return children
}

func (node *Drop_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Drop_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Drop_Nodes_Node) GetYangName() string { return "node" }

func (node *Drop_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Drop_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Drop_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Drop_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Drop_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Drop_Nodes_Node) GetParentYangName() string { return "nodes" }

// Drop_Nodes_Node_NpuNumberForDropStats
// NPU drop stats
type Drop_Nodes_Node_NpuNumberForDropStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All drop stats for a particular NPU. The type is slice of
    // Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat.
    NpuNumberForDropStat []Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat
}

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetFilter() yfilter.YFilter { return npuNumberForDropStats.YFilter }

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) SetFilter(yf yfilter.YFilter) { npuNumberForDropStats.YFilter = yf }

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetGoName(yname string) string {
    if yname == "npu-number-for-drop-stat" { return "NpuNumberForDropStat" }
    return ""
}

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetSegmentPath() string {
    return "npu-number-for-drop-stats"
}

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "npu-number-for-drop-stat" {
        for _, c := range npuNumberForDropStats.NpuNumberForDropStat {
            if npuNumberForDropStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat{}
        npuNumberForDropStats.NpuNumberForDropStat = append(npuNumberForDropStats.NpuNumberForDropStat, child)
        return &npuNumberForDropStats.NpuNumberForDropStat[len(npuNumberForDropStats.NpuNumberForDropStat)-1]
    }
    return nil
}

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range npuNumberForDropStats.NpuNumberForDropStat {
        children[npuNumberForDropStats.NpuNumberForDropStat[i].GetSegmentPath()] = &npuNumberForDropStats.NpuNumberForDropStat[i]
    }
    return children
}

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetBundleName() string { return "cisco_ios_xr" }

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetYangName() string { return "npu-number-for-drop-stats" }

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) SetParent(parent types.Entity) { npuNumberForDropStats.parent = parent }

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetParent() types.Entity { return npuNumberForDropStats.parent }

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetParentYangName() string { return "node" }

// Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat
// All drop stats for a particular NPU
type Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NPU number. The type is interface{} with range:
    // -2147483648..2147483647.
    NpuId interface{}

    // Second argument to the module. The type is slice of
    // Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData.
    DropSpecificStatsData []Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData
}

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetFilter() yfilter.YFilter { return npuNumberForDropStat.YFilter }

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) SetFilter(yf yfilter.YFilter) { npuNumberForDropStat.YFilter = yf }

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "drop-specific-stats-data" { return "DropSpecificStatsData" }
    return ""
}

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetSegmentPath() string {
    return "npu-number-for-drop-stat" + "[npu-id='" + fmt.Sprintf("%v", npuNumberForDropStat.NpuId) + "']"
}

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "drop-specific-stats-data" {
        for _, c := range npuNumberForDropStat.DropSpecificStatsData {
            if npuNumberForDropStat.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData{}
        npuNumberForDropStat.DropSpecificStatsData = append(npuNumberForDropStat.DropSpecificStatsData, child)
        return &npuNumberForDropStat.DropSpecificStatsData[len(npuNumberForDropStat.DropSpecificStatsData)-1]
    }
    return nil
}

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range npuNumberForDropStat.DropSpecificStatsData {
        children[npuNumberForDropStat.DropSpecificStatsData[i].GetSegmentPath()] = &npuNumberForDropStat.DropSpecificStatsData[i]
    }
    return children
}

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = npuNumberForDropStat.NpuId
    return leafs
}

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetBundleName() string { return "cisco_ios_xr" }

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetYangName() string { return "npu-number-for-drop-stat" }

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) SetParent(parent types.Entity) { npuNumberForDropStat.parent = parent }

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetParent() types.Entity { return npuNumberForDropStat.parent }

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetParentYangName() string { return "npu-number-for-drop-stats" }

// Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData
// Second argument to the module
type Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Drop ID. The type is interface{} with range:
    // -2147483648..2147483647.
    DropData interface{}

    // id. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // name. The type is string.
    Name interface{}

    // count. The type is interface{} with range: 0..18446744073709551615.
    Count interface{}
}

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetFilter() yfilter.YFilter { return dropSpecificStatsData.YFilter }

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) SetFilter(yf yfilter.YFilter) { dropSpecificStatsData.YFilter = yf }

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetGoName(yname string) string {
    if yname == "drop-data" { return "DropData" }
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "count" { return "Count" }
    return ""
}

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetSegmentPath() string {
    return "drop-specific-stats-data" + "[drop-data='" + fmt.Sprintf("%v", dropSpecificStatsData.DropData) + "']"
}

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["drop-data"] = dropSpecificStatsData.DropData
    leafs["id"] = dropSpecificStatsData.Id
    leafs["name"] = dropSpecificStatsData.Name
    leafs["count"] = dropSpecificStatsData.Count
    return leafs
}

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetBundleName() string { return "cisco_ios_xr" }

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetYangName() string { return "drop-specific-stats-data" }

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) SetParent(parent types.Entity) { dropSpecificStatsData.parent = parent }

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetParent() types.Entity { return dropSpecificStatsData.parent }

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetParentYangName() string { return "npu-number-for-drop-stat" }

