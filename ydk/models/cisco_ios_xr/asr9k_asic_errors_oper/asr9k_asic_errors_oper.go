// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-asic-errors package operational data.
// 
// This module contains definitions
// for the following management objects:
//   asic-error-stats: Asic error stats operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_asic_errors_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_asic_errors_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-asic-errors-oper asic-error-stats}", reflect.TypeOf(AsicErrorStats{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-asic-errors-oper:asic-error-stats", reflect.TypeOf(AsicErrorStats{}))
}

// AsicErrorStats
// Asic error stats operational data
type AsicErrorStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of racks.
    Racks AsicErrorStats_Racks
}

func (asicErrorStats *AsicErrorStats) GetEntityData() *types.CommonEntityData {
    asicErrorStats.EntityData.YFilter = asicErrorStats.YFilter
    asicErrorStats.EntityData.YangName = "asic-error-stats"
    asicErrorStats.EntityData.BundleName = "cisco_ios_xr"
    asicErrorStats.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-asic-errors-oper"
    asicErrorStats.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-asic-errors-oper:asic-error-stats"
    asicErrorStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicErrorStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicErrorStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicErrorStats.EntityData.Children = types.NewOrderedMap()
    asicErrorStats.EntityData.Children.Append("racks", types.YChild{"Racks", &asicErrorStats.Racks})
    asicErrorStats.EntityData.Leafs = types.NewOrderedMap()

    asicErrorStats.EntityData.YListKeys = []string {}

    return &(asicErrorStats.EntityData)
}

// AsicErrorStats_Racks
// Table of racks
type AsicErrorStats_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of AsicErrorStats_Racks_Rack.
    Rack []*AsicErrorStats_Racks_Rack
}

func (racks *AsicErrorStats_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "asic-error-stats"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = types.NewOrderedMap()
    racks.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range racks.Rack {
        racks.EntityData.Children.Append(types.GetSegmentPath(racks.Rack[i]), types.YChild{"Rack", racks.Rack[i]})
    }
    racks.EntityData.Leafs = types.NewOrderedMap()

    racks.EntityData.YListKeys = []string {}

    return &(racks.EntityData)
}

// AsicErrorStats_Racks_Rack
// Number
type AsicErrorStats_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // 0..4294967295.
    Rack interface{}

    // Table of Nodes.
    Nodes AsicErrorStats_Racks_Rack_Nodes
}

func (rack *AsicErrorStats_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.Rack, "rack")
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("nodes", types.YChild{"Nodes", &rack.Nodes})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("rack", types.YLeaf{"Rack", rack.Rack})

    rack.EntityData.YListKeys = []string {"Rack"}

    return &(rack.EntityData)
}

// AsicErrorStats_Racks_Rack_Nodes
// Table of Nodes
type AsicErrorStats_Racks_Rack_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of
    // AsicErrorStats_Racks_Rack_Nodes_Node.
    Node []*AsicErrorStats_Racks_Rack_Nodes_Node
}

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "rack"
    nodes.EntityData.SegmentPath = "nodes"
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

// AsicErrorStats_Racks_Rack_Nodes_Node
// Information about a particular node
type AsicErrorStats_Racks_Rack_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Table of all Asic Types information on a node.
    Counts AsicErrorStats_Racks_Rack_Nodes_Node_Counts
}

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("counts", types.YChild{"Counts", &node.Counts})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// AsicErrorStats_Racks_Rack_Nodes_Node_Counts
// Table of all Asic Types information on a
// node
type AsicErrorStats_Racks_Rack_Nodes_Node_Counts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary Asic error counts for a Asic Type. The type is slice of
    // AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count.
    Count []*AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count
}

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetEntityData() *types.CommonEntityData {
    counts.EntityData.YFilter = counts.YFilter
    counts.EntityData.YangName = "counts"
    counts.EntityData.BundleName = "cisco_ios_xr"
    counts.EntityData.ParentYangName = "node"
    counts.EntityData.SegmentPath = "counts"
    counts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counts.EntityData.Children = types.NewOrderedMap()
    counts.EntityData.Children.Append("count", types.YChild{"Count", nil})
    for i := range counts.Count {
        counts.EntityData.Children.Append(types.GetSegmentPath(counts.Count[i]), types.YChild{"Count", counts.Count[i]})
    }
    counts.EntityData.Leafs = types.NewOrderedMap()

    counts.EntityData.YListKeys = []string {}

    return &(counts.EntityData)
}

// AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count
// Summary Asic error counts for a Asic Type
type AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Asic Type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // sum data. The type is slice of
    // AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData.
    SumData []*AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData
}

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetEntityData() *types.CommonEntityData {
    count.EntityData.YFilter = count.YFilter
    count.EntityData.YangName = "count"
    count.EntityData.BundleName = "cisco_ios_xr"
    count.EntityData.ParentYangName = "counts"
    count.EntityData.SegmentPath = "count" + types.AddKeyToken(count.Type, "type")
    count.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    count.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    count.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    count.EntityData.Children = types.NewOrderedMap()
    count.EntityData.Children.Append("sum-data", types.YChild{"SumData", nil})
    for i := range count.SumData {
        count.EntityData.Children.Append(types.GetSegmentPath(count.SumData[i]), types.YChild{"SumData", count.SumData[i]})
    }
    count.EntityData.Leafs = types.NewOrderedMap()
    count.EntityData.Leafs.Append("type", types.YLeaf{"Type", count.Type})

    count.EntityData.YListKeys = []string {"Type"}

    return &(count.EntityData)
}

// AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData
// sum data
type AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // instance. The type is interface{} with range: 0..4294967295.
    Instance interface{}

    // num nodes. The type is interface{} with range: 0..4294967295.
    NumNodes interface{}

    // crc err count. The type is interface{} with range: 0..4294967295.
    CrcErrCount interface{}

    // sbe err count. The type is interface{} with range: 0..4294967295.
    SbeErrCount interface{}

    // mbe err count. The type is interface{} with range: 0..4294967295.
    MbeErrCount interface{}

    // par err count. The type is interface{} with range: 0..4294967295.
    ParErrCount interface{}

    // gen err count. The type is interface{} with range: 0..4294967295.
    GenErrCount interface{}

    // reset err count. The type is interface{} with range: 0..4294967295.
    ResetErrCount interface{}

    // node key. The type is slice of interface{} with range: 0..4294967295.
    NodeKey []interface{}
}

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetEntityData() *types.CommonEntityData {
    sumData.EntityData.YFilter = sumData.YFilter
    sumData.EntityData.YangName = "sum-data"
    sumData.EntityData.BundleName = "cisco_ios_xr"
    sumData.EntityData.ParentYangName = "count"
    sumData.EntityData.SegmentPath = "sum-data"
    sumData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumData.EntityData.Children = types.NewOrderedMap()
    sumData.EntityData.Leafs = types.NewOrderedMap()
    sumData.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", sumData.Instance})
    sumData.EntityData.Leafs.Append("num-nodes", types.YLeaf{"NumNodes", sumData.NumNodes})
    sumData.EntityData.Leafs.Append("crc-err-count", types.YLeaf{"CrcErrCount", sumData.CrcErrCount})
    sumData.EntityData.Leafs.Append("sbe-err-count", types.YLeaf{"SbeErrCount", sumData.SbeErrCount})
    sumData.EntityData.Leafs.Append("mbe-err-count", types.YLeaf{"MbeErrCount", sumData.MbeErrCount})
    sumData.EntityData.Leafs.Append("par-err-count", types.YLeaf{"ParErrCount", sumData.ParErrCount})
    sumData.EntityData.Leafs.Append("gen-err-count", types.YLeaf{"GenErrCount", sumData.GenErrCount})
    sumData.EntityData.Leafs.Append("reset-err-count", types.YLeaf{"ResetErrCount", sumData.ResetErrCount})
    sumData.EntityData.Leafs.Append("node-key", types.YLeaf{"NodeKey", sumData.NodeKey})

    sumData.EntityData.YListKeys = []string {}

    return &(sumData.EntityData)
}

