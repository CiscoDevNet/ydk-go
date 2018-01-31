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
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of racks.
    Racks AsicErrorStats_Racks
}

func (asicErrorStats *AsicErrorStats) GetFilter() yfilter.YFilter { return asicErrorStats.YFilter }

func (asicErrorStats *AsicErrorStats) SetFilter(yf yfilter.YFilter) { asicErrorStats.YFilter = yf }

func (asicErrorStats *AsicErrorStats) GetGoName(yname string) string {
    if yname == "racks" { return "Racks" }
    return ""
}

func (asicErrorStats *AsicErrorStats) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-asic-errors-oper:asic-error-stats"
}

func (asicErrorStats *AsicErrorStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "racks" {
        return &asicErrorStats.Racks
    }
    return nil
}

func (asicErrorStats *AsicErrorStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["racks"] = &asicErrorStats.Racks
    return children
}

func (asicErrorStats *AsicErrorStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicErrorStats *AsicErrorStats) GetBundleName() string { return "cisco_ios_xr" }

func (asicErrorStats *AsicErrorStats) GetYangName() string { return "asic-error-stats" }

func (asicErrorStats *AsicErrorStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicErrorStats *AsicErrorStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicErrorStats *AsicErrorStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicErrorStats *AsicErrorStats) SetParent(parent types.Entity) { asicErrorStats.parent = parent }

func (asicErrorStats *AsicErrorStats) GetParent() types.Entity { return asicErrorStats.parent }

func (asicErrorStats *AsicErrorStats) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-asic-errors-oper" }

// AsicErrorStats_Racks
// Table of racks
type AsicErrorStats_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of AsicErrorStats_Racks_Rack.
    Rack []AsicErrorStats_Racks_Rack
}

func (racks *AsicErrorStats_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *AsicErrorStats_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *AsicErrorStats_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *AsicErrorStats_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *AsicErrorStats_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrorStats_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *AsicErrorStats_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *AsicErrorStats_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *AsicErrorStats_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *AsicErrorStats_Racks) GetYangName() string { return "racks" }

func (racks *AsicErrorStats_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *AsicErrorStats_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *AsicErrorStats_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *AsicErrorStats_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *AsicErrorStats_Racks) GetParent() types.Entity { return racks.parent }

func (racks *AsicErrorStats_Racks) GetParentYangName() string { return "asic-error-stats" }

// AsicErrorStats_Racks_Rack
// Number
type AsicErrorStats_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of Nodes.
    Nodes AsicErrorStats_Racks_Rack_Nodes
}

func (rack *AsicErrorStats_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *AsicErrorStats_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *AsicErrorStats_Racks_Rack) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (rack *AsicErrorStats_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
}

func (rack *AsicErrorStats_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &rack.Nodes
    }
    return nil
}

func (rack *AsicErrorStats_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &rack.Nodes
    return children
}

func (rack *AsicErrorStats_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack"] = rack.Rack
    return leafs
}

func (rack *AsicErrorStats_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *AsicErrorStats_Racks_Rack) GetYangName() string { return "rack" }

func (rack *AsicErrorStats_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *AsicErrorStats_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *AsicErrorStats_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *AsicErrorStats_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *AsicErrorStats_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *AsicErrorStats_Racks_Rack) GetParentYangName() string { return "racks" }

// AsicErrorStats_Racks_Rack_Nodes
// Table of Nodes
type AsicErrorStats_Racks_Rack_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of
    // AsicErrorStats_Racks_Rack_Nodes_Node.
    Node []AsicErrorStats_Racks_Rack_Nodes_Node
}

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *AsicErrorStats_Racks_Rack_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrorStats_Racks_Rack_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetYangName() string { return "nodes" }

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *AsicErrorStats_Racks_Rack_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *AsicErrorStats_Racks_Rack_Nodes) GetParentYangName() string { return "rack" }

// AsicErrorStats_Racks_Rack_Nodes_Node
// Information about a particular node
type AsicErrorStats_Racks_Rack_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Table of all Asic Types information on a node.
    Counts AsicErrorStats_Racks_Rack_Nodes_Node_Counts
}

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "counts" { return "Counts" }
    return ""
}

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counts" {
        return &node.Counts
    }
    return nil
}

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counts"] = &node.Counts
    return children
}

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetYangName() string { return "node" }

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *AsicErrorStats_Racks_Rack_Nodes_Node) GetParentYangName() string { return "nodes" }

// AsicErrorStats_Racks_Rack_Nodes_Node_Counts
// Table of all Asic Types information on a
// node
type AsicErrorStats_Racks_Rack_Nodes_Node_Counts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary Asic error counts for a Asic Type. The type is slice of
    // AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count.
    Count []AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count
}

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetFilter() yfilter.YFilter { return counts.YFilter }

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) SetFilter(yf yfilter.YFilter) { counts.YFilter = yf }

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetGoName(yname string) string {
    if yname == "count" { return "Count" }
    return ""
}

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetSegmentPath() string {
    return "counts"
}

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "count" {
        for _, c := range counts.Count {
            if counts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count{}
        counts.Count = append(counts.Count, child)
        return &counts.Count[len(counts.Count)-1]
    }
    return nil
}

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range counts.Count {
        children[counts.Count[i].GetSegmentPath()] = &counts.Count[i]
    }
    return children
}

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetBundleName() string { return "cisco_ios_xr" }

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetYangName() string { return "counts" }

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) SetParent(parent types.Entity) { counts.parent = parent }

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetParent() types.Entity { return counts.parent }

func (counts *AsicErrorStats_Racks_Rack_Nodes_Node_Counts) GetParentYangName() string { return "node" }

// AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count
// Summary Asic error counts for a Asic Type
type AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Asic Type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // sum data. The type is slice of
    // AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData.
    SumData []AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData
}

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetFilter() yfilter.YFilter { return count.YFilter }

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) SetFilter(yf yfilter.YFilter) { count.YFilter = yf }

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "sum-data" { return "SumData" }
    return ""
}

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetSegmentPath() string {
    return "count" + "[type='" + fmt.Sprintf("%v", count.Type) + "']"
}

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sum-data" {
        for _, c := range count.SumData {
            if count.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData{}
        count.SumData = append(count.SumData, child)
        return &count.SumData[len(count.SumData)-1]
    }
    return nil
}

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range count.SumData {
        children[count.SumData[i].GetSegmentPath()] = &count.SumData[i]
    }
    return children
}

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = count.Type
    return leafs
}

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetBundleName() string { return "cisco_ios_xr" }

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetYangName() string { return "count" }

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) SetParent(parent types.Entity) { count.parent = parent }

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetParent() types.Entity { return count.parent }

func (count *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count) GetParentYangName() string { return "counts" }

// AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData
// sum data
type AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData struct {
    parent types.Entity
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

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetFilter() yfilter.YFilter { return sumData.YFilter }

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) SetFilter(yf yfilter.YFilter) { sumData.YFilter = yf }

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    if yname == "num-nodes" { return "NumNodes" }
    if yname == "crc-err-count" { return "CrcErrCount" }
    if yname == "sbe-err-count" { return "SbeErrCount" }
    if yname == "mbe-err-count" { return "MbeErrCount" }
    if yname == "par-err-count" { return "ParErrCount" }
    if yname == "gen-err-count" { return "GenErrCount" }
    if yname == "reset-err-count" { return "ResetErrCount" }
    if yname == "node-key" { return "NodeKey" }
    return ""
}

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetSegmentPath() string {
    return "sum-data"
}

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance"] = sumData.Instance
    leafs["num-nodes"] = sumData.NumNodes
    leafs["crc-err-count"] = sumData.CrcErrCount
    leafs["sbe-err-count"] = sumData.SbeErrCount
    leafs["mbe-err-count"] = sumData.MbeErrCount
    leafs["par-err-count"] = sumData.ParErrCount
    leafs["gen-err-count"] = sumData.GenErrCount
    leafs["reset-err-count"] = sumData.ResetErrCount
    leafs["node-key"] = sumData.NodeKey
    return leafs
}

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetBundleName() string { return "cisco_ios_xr" }

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetYangName() string { return "sum-data" }

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) SetParent(parent types.Entity) { sumData.parent = parent }

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetParent() types.Entity { return sumData.parent }

func (sumData *AsicErrorStats_Racks_Rack_Nodes_Node_Counts_Count_SumData) GetParentYangName() string { return "count" }

