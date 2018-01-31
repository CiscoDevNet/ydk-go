// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-fsi package operational data.
// 
// This module contains definitions
// for the following management objects:
//   fabric-stats: Fabric stats operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_fsi_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_fsi_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-fsi-oper fabric-stats}", reflect.TypeOf(FabricStats{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-fsi-oper:fabric-stats", reflect.TypeOf(FabricStats{}))
}

// FabricStats
// Fabric stats operational data
type FabricStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Nodes.
    Nodes FabricStats_Nodes
}

func (fabricStats *FabricStats) GetFilter() yfilter.YFilter { return fabricStats.YFilter }

func (fabricStats *FabricStats) SetFilter(yf yfilter.YFilter) { fabricStats.YFilter = yf }

func (fabricStats *FabricStats) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (fabricStats *FabricStats) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-fsi-oper:fabric-stats"
}

func (fabricStats *FabricStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &fabricStats.Nodes
    }
    return nil
}

func (fabricStats *FabricStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &fabricStats.Nodes
    return children
}

func (fabricStats *FabricStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fabricStats *FabricStats) GetBundleName() string { return "cisco_ios_xr" }

func (fabricStats *FabricStats) GetYangName() string { return "fabric-stats" }

func (fabricStats *FabricStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fabricStats *FabricStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fabricStats *FabricStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fabricStats *FabricStats) SetParent(parent types.Entity) { fabricStats.parent = parent }

func (fabricStats *FabricStats) GetParent() types.Entity { return fabricStats.parent }

func (fabricStats *FabricStats) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-fsi-oper" }

// FabricStats_Nodes
// Table of Nodes
type FabricStats_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of
    // FabricStats_Nodes_Node.
    Node []FabricStats_Nodes_Node
}

func (nodes *FabricStats_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *FabricStats_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *FabricStats_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *FabricStats_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *FabricStats_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FabricStats_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *FabricStats_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *FabricStats_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *FabricStats_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *FabricStats_Nodes) GetYangName() string { return "nodes" }

func (nodes *FabricStats_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *FabricStats_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *FabricStats_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *FabricStats_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *FabricStats_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *FabricStats_Nodes) GetParentYangName() string { return "fabric-stats" }

// FabricStats_Nodes_Node
// Information about a particular node
type FabricStats_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Table of stats information.
    Statses FabricStats_Nodes_Node_Statses
}

func (node *FabricStats_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *FabricStats_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *FabricStats_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "statses" { return "Statses" }
    return ""
}

func (node *FabricStats_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *FabricStats_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statses" {
        return &node.Statses
    }
    return nil
}

func (node *FabricStats_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statses"] = &node.Statses
    return children
}

func (node *FabricStats_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *FabricStats_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *FabricStats_Nodes_Node) GetYangName() string { return "node" }

func (node *FabricStats_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *FabricStats_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *FabricStats_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *FabricStats_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *FabricStats_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *FabricStats_Nodes_Node) GetParentYangName() string { return "nodes" }

// FabricStats_Nodes_Node_Statses
// Table of stats information
type FabricStats_Nodes_Node_Statses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats information for a particular type. The type is slice of
    // FabricStats_Nodes_Node_Statses_Stats.
    Stats []FabricStats_Nodes_Node_Statses_Stats
}

func (statses *FabricStats_Nodes_Node_Statses) GetFilter() yfilter.YFilter { return statses.YFilter }

func (statses *FabricStats_Nodes_Node_Statses) SetFilter(yf yfilter.YFilter) { statses.YFilter = yf }

func (statses *FabricStats_Nodes_Node_Statses) GetGoName(yname string) string {
    if yname == "stats" { return "Stats" }
    return ""
}

func (statses *FabricStats_Nodes_Node_Statses) GetSegmentPath() string {
    return "statses"
}

func (statses *FabricStats_Nodes_Node_Statses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        for _, c := range statses.Stats {
            if statses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FabricStats_Nodes_Node_Statses_Stats{}
        statses.Stats = append(statses.Stats, child)
        return &statses.Stats[len(statses.Stats)-1]
    }
    return nil
}

func (statses *FabricStats_Nodes_Node_Statses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statses.Stats {
        children[statses.Stats[i].GetSegmentPath()] = &statses.Stats[i]
    }
    return children
}

func (statses *FabricStats_Nodes_Node_Statses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statses *FabricStats_Nodes_Node_Statses) GetBundleName() string { return "cisco_ios_xr" }

func (statses *FabricStats_Nodes_Node_Statses) GetYangName() string { return "statses" }

func (statses *FabricStats_Nodes_Node_Statses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statses *FabricStats_Nodes_Node_Statses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statses *FabricStats_Nodes_Node_Statses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statses *FabricStats_Nodes_Node_Statses) SetParent(parent types.Entity) { statses.parent = parent }

func (statses *FabricStats_Nodes_Node_Statses) GetParent() types.Entity { return statses.parent }

func (statses *FabricStats_Nodes_Node_Statses) GetParentYangName() string { return "node" }

// FabricStats_Nodes_Node_Statses_Stats
// Stats information for a particular type
type FabricStats_Nodes_Node_Statses_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fabric asic type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // Last Clear Time. The type is interface{} with range:
    // 0..18446744073709551615.
    LastClearTime interface{}

    // Stat Table Name. The type is string.
    StatTableName interface{}

    // Array of counters . The type is slice of
    // FabricStats_Nodes_Node_Statses_Stats_StatsTable.
    StatsTable []FabricStats_Nodes_Node_Statses_Stats_StatsTable
}

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *FabricStats_Nodes_Node_Statses_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "last-clear-time" { return "LastClearTime" }
    if yname == "stat-table-name" { return "StatTableName" }
    if yname == "stats-table" { return "StatsTable" }
    return ""
}

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetSegmentPath() string {
    return "stats" + "[type='" + fmt.Sprintf("%v", stats.Type) + "']"
}

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats-table" {
        for _, c := range stats.StatsTable {
            if stats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FabricStats_Nodes_Node_Statses_Stats_StatsTable{}
        stats.StatsTable = append(stats.StatsTable, child)
        return &stats.StatsTable[len(stats.StatsTable)-1]
    }
    return nil
}

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stats.StatsTable {
        children[stats.StatsTable[i].GetSegmentPath()] = &stats.StatsTable[i]
    }
    return children
}

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = stats.Type
    leafs["last-clear-time"] = stats.LastClearTime
    leafs["stat-table-name"] = stats.StatTableName
    return leafs
}

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetYangName() string { return "stats" }

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *FabricStats_Nodes_Node_Statses_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetParent() types.Entity { return stats.parent }

func (stats *FabricStats_Nodes_Node_Statses_Stats) GetParentYangName() string { return "statses" }

// FabricStats_Nodes_Node_Statses_Stats_StatsTable
// Array of counters 
type FabricStats_Nodes_Node_Statses_Stats_StatsTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats Table. The type is slice of
    // FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat.
    FsiStat []FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat
}

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetFilter() yfilter.YFilter { return statsTable.YFilter }

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) SetFilter(yf yfilter.YFilter) { statsTable.YFilter = yf }

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetGoName(yname string) string {
    if yname == "fsi-stat" { return "FsiStat" }
    return ""
}

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetSegmentPath() string {
    return "stats-table"
}

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fsi-stat" {
        for _, c := range statsTable.FsiStat {
            if statsTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat{}
        statsTable.FsiStat = append(statsTable.FsiStat, child)
        return &statsTable.FsiStat[len(statsTable.FsiStat)-1]
    }
    return nil
}

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statsTable.FsiStat {
        children[statsTable.FsiStat[i].GetSegmentPath()] = &statsTable.FsiStat[i]
    }
    return children
}

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetBundleName() string { return "cisco_ios_xr" }

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetYangName() string { return "stats-table" }

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) SetParent(parent types.Entity) { statsTable.parent = parent }

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetParent() types.Entity { return statsTable.parent }

func (statsTable *FabricStats_Nodes_Node_Statses_Stats_StatsTable) GetParentYangName() string { return "stats" }

// FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat
// Stats Table
type FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Counter. The type is interface{} with range: 0..18446744073709551615.
    Count interface{}

    // Counter Name. The type is string.
    CounterName interface{}
}

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetFilter() yfilter.YFilter { return fsiStat.YFilter }

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) SetFilter(yf yfilter.YFilter) { fsiStat.YFilter = yf }

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetGoName(yname string) string {
    if yname == "count" { return "Count" }
    if yname == "counter-name" { return "CounterName" }
    return ""
}

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetSegmentPath() string {
    return "fsi-stat"
}

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["count"] = fsiStat.Count
    leafs["counter-name"] = fsiStat.CounterName
    return leafs
}

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetBundleName() string { return "cisco_ios_xr" }

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetYangName() string { return "fsi-stat" }

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) SetParent(parent types.Entity) { fsiStat.parent = parent }

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetParent() types.Entity { return fsiStat.parent }

func (fsiStat *FabricStats_Nodes_Node_Statses_Stats_StatsTable_FsiStat) GetParentYangName() string { return "stats-table" }

