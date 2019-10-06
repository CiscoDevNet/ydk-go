// This module contains a collection of YANG definitions
// for Cisco IOS-XR fretta-bcm-dpa-drop-stats package operational data.
// 
// This module contains definitions
// for the following management objects:
//   drop: Drop stats data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Drop data per node.
    Nodes Drop_Nodes
}

func (drop *Drop) GetEntityData() *types.CommonEntityData {
    drop.EntityData.YFilter = drop.YFilter
    drop.EntityData.YangName = "drop"
    drop.EntityData.BundleName = "cisco_ios_xr"
    drop.EntityData.ParentYangName = "Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper"
    drop.EntityData.SegmentPath = "Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper:drop"
    drop.EntityData.AbsolutePath = drop.EntityData.SegmentPath
    drop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    drop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    drop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    drop.EntityData.Children = types.NewOrderedMap()
    drop.EntityData.Children.Append("nodes", types.YChild{"Nodes", &drop.Nodes})
    drop.EntityData.Leafs = types.NewOrderedMap()

    drop.EntityData.YListKeys = []string {}

    return &(drop.EntityData)
}

// Drop_Nodes
// Drop data per node
type Drop_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Drop stats data for a particular node. The type is slice of
    // Drop_Nodes_Node.
    Node []*Drop_Nodes_Node
}

func (nodes *Drop_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "drop"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper:drop/" + nodes.EntityData.SegmentPath
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

// Drop_Nodes_Node
// Drop stats data for a particular node
type Drop_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // NPU drop stats.
    NpuNumberForDropStats Drop_Nodes_Node_NpuNumberForDropStats
}

func (node *Drop_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper:drop/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("npu-number-for-drop-stats", types.YChild{"NpuNumberForDropStats", &node.NpuNumberForDropStats})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Drop_Nodes_Node_NpuNumberForDropStats
// NPU drop stats
type Drop_Nodes_Node_NpuNumberForDropStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All drop stats for a particular NPU. The type is slice of
    // Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat.
    NpuNumberForDropStat []*Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat
}

func (npuNumberForDropStats *Drop_Nodes_Node_NpuNumberForDropStats) GetEntityData() *types.CommonEntityData {
    npuNumberForDropStats.EntityData.YFilter = npuNumberForDropStats.YFilter
    npuNumberForDropStats.EntityData.YangName = "npu-number-for-drop-stats"
    npuNumberForDropStats.EntityData.BundleName = "cisco_ios_xr"
    npuNumberForDropStats.EntityData.ParentYangName = "node"
    npuNumberForDropStats.EntityData.SegmentPath = "npu-number-for-drop-stats"
    npuNumberForDropStats.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper:drop/nodes/node/" + npuNumberForDropStats.EntityData.SegmentPath
    npuNumberForDropStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuNumberForDropStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuNumberForDropStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuNumberForDropStats.EntityData.Children = types.NewOrderedMap()
    npuNumberForDropStats.EntityData.Children.Append("npu-number-for-drop-stat", types.YChild{"NpuNumberForDropStat", nil})
    for i := range npuNumberForDropStats.NpuNumberForDropStat {
        npuNumberForDropStats.EntityData.Children.Append(types.GetSegmentPath(npuNumberForDropStats.NpuNumberForDropStat[i]), types.YChild{"NpuNumberForDropStat", npuNumberForDropStats.NpuNumberForDropStat[i]})
    }
    npuNumberForDropStats.EntityData.Leafs = types.NewOrderedMap()

    npuNumberForDropStats.EntityData.YListKeys = []string {}

    return &(npuNumberForDropStats.EntityData)
}

// Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat
// All drop stats for a particular NPU
type Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NPU number. The type is interface{} with range:
    // 0..4294967295.
    NpuId interface{}

    // Second argument to the module. The type is slice of
    // Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData.
    DropSpecificStatsData []*Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData
}

func (npuNumberForDropStat *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat) GetEntityData() *types.CommonEntityData {
    npuNumberForDropStat.EntityData.YFilter = npuNumberForDropStat.YFilter
    npuNumberForDropStat.EntityData.YangName = "npu-number-for-drop-stat"
    npuNumberForDropStat.EntityData.BundleName = "cisco_ios_xr"
    npuNumberForDropStat.EntityData.ParentYangName = "npu-number-for-drop-stats"
    npuNumberForDropStat.EntityData.SegmentPath = "npu-number-for-drop-stat" + types.AddKeyToken(npuNumberForDropStat.NpuId, "npu-id")
    npuNumberForDropStat.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper:drop/nodes/node/npu-number-for-drop-stats/" + npuNumberForDropStat.EntityData.SegmentPath
    npuNumberForDropStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuNumberForDropStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuNumberForDropStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuNumberForDropStat.EntityData.Children = types.NewOrderedMap()
    npuNumberForDropStat.EntityData.Children.Append("drop-specific-stats-data", types.YChild{"DropSpecificStatsData", nil})
    for i := range npuNumberForDropStat.DropSpecificStatsData {
        npuNumberForDropStat.EntityData.Children.Append(types.GetSegmentPath(npuNumberForDropStat.DropSpecificStatsData[i]), types.YChild{"DropSpecificStatsData", npuNumberForDropStat.DropSpecificStatsData[i]})
    }
    npuNumberForDropStat.EntityData.Leafs = types.NewOrderedMap()
    npuNumberForDropStat.EntityData.Leafs.Append("npu-id", types.YLeaf{"NpuId", npuNumberForDropStat.NpuId})

    npuNumberForDropStat.EntityData.YListKeys = []string {"NpuId"}

    return &(npuNumberForDropStat.EntityData)
}

// Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData
// Second argument to the module
type Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Drop ID. The type is interface{} with range:
    // 0..4294967295.
    DropData interface{}

    // id. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // name. The type is string.
    Name interface{}

    // count. The type is interface{} with range: 0..18446744073709551615.
    Count interface{}
}

func (dropSpecificStatsData *Drop_Nodes_Node_NpuNumberForDropStats_NpuNumberForDropStat_DropSpecificStatsData) GetEntityData() *types.CommonEntityData {
    dropSpecificStatsData.EntityData.YFilter = dropSpecificStatsData.YFilter
    dropSpecificStatsData.EntityData.YangName = "drop-specific-stats-data"
    dropSpecificStatsData.EntityData.BundleName = "cisco_ios_xr"
    dropSpecificStatsData.EntityData.ParentYangName = "npu-number-for-drop-stat"
    dropSpecificStatsData.EntityData.SegmentPath = "drop-specific-stats-data" + types.AddKeyToken(dropSpecificStatsData.DropData, "drop-data")
    dropSpecificStatsData.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-bcm-dpa-drop-stats-oper:drop/nodes/node/npu-number-for-drop-stats/npu-number-for-drop-stat/" + dropSpecificStatsData.EntityData.SegmentPath
    dropSpecificStatsData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropSpecificStatsData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropSpecificStatsData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropSpecificStatsData.EntityData.Children = types.NewOrderedMap()
    dropSpecificStatsData.EntityData.Leafs = types.NewOrderedMap()
    dropSpecificStatsData.EntityData.Leafs.Append("drop-data", types.YLeaf{"DropData", dropSpecificStatsData.DropData})
    dropSpecificStatsData.EntityData.Leafs.Append("id", types.YLeaf{"Id", dropSpecificStatsData.Id})
    dropSpecificStatsData.EntityData.Leafs.Append("name", types.YLeaf{"Name", dropSpecificStatsData.Name})
    dropSpecificStatsData.EntityData.Leafs.Append("count", types.YLeaf{"Count", dropSpecificStatsData.Count})

    dropSpecificStatsData.EntityData.YListKeys = []string {"DropData"}

    return &(dropSpecificStatsData.EntityData)
}

