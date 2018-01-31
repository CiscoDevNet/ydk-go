// This module contains a collection of YANG definitions
// for Cisco IOS-XR fretta-bcm-dpa-hw-resources package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dpa: Stats Data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package fretta_bcm_dpa_hw_resources_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fretta_bcm_dpa_hw_resources_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-fretta-bcm-dpa-hw-resources-oper dpa}", reflect.TypeOf(Dpa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-fretta-bcm-dpa-hw-resources-oper:dpa", reflect.TypeOf(Dpa{}))
}

// Resource represents Resource
type Resource string

const (
    // lem
    Resource_lem Resource = "lem"

    // lpm
    Resource_lpm Resource = "lpm"

    // encap
    Resource_encap Resource = "encap"

    // ext tcam ipv4
    Resource_ext_tcam_ipv4 Resource = "ext-tcam-ipv4"

    // ext tcam ipv6 short
    Resource_ext_tcam_ipv6_short Resource = "ext-tcam-ipv6-short"

    // ext tcam ipv6 long
    Resource_ext_tcam_ipv6_long Resource = "ext-tcam-ipv6-long"

    // fec
    Resource_fec Resource = "fec"

    // ecmp fec
    Resource_ecmpfec Resource = "ecmpfec"
)

// Dpa
// Stats Data
type Dpa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Voq or Trap Data.
    Stats Dpa_Stats
}

func (dpa *Dpa) GetFilter() yfilter.YFilter { return dpa.YFilter }

func (dpa *Dpa) SetFilter(yf yfilter.YFilter) { dpa.YFilter = yf }

func (dpa *Dpa) GetGoName(yname string) string {
    if yname == "stats" { return "Stats" }
    return ""
}

func (dpa *Dpa) GetSegmentPath() string {
    return "Cisco-IOS-XR-fretta-bcm-dpa-hw-resources-oper:dpa"
}

func (dpa *Dpa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &dpa.Stats
    }
    return nil
}

func (dpa *Dpa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &dpa.Stats
    return children
}

func (dpa *Dpa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dpa *Dpa) GetBundleName() string { return "cisco_ios_xr" }

func (dpa *Dpa) GetYangName() string { return "dpa" }

func (dpa *Dpa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dpa *Dpa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dpa *Dpa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dpa *Dpa) SetParent(parent types.Entity) { dpa.parent = parent }

func (dpa *Dpa) GetParent() types.Entity { return dpa.parent }

func (dpa *Dpa) GetParentYangName() string { return "Cisco-IOS-XR-fretta-bcm-dpa-hw-resources-oper" }

// Dpa_Stats
// Voq or Trap Data
type Dpa_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DPA data for available nodes.
    Nodes Dpa_Stats_Nodes
}

func (stats *Dpa_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *Dpa_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *Dpa_Stats) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (stats *Dpa_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *Dpa_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &stats.Nodes
    }
    return nil
}

func (stats *Dpa_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &stats.Nodes
    return children
}

func (stats *Dpa_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stats *Dpa_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *Dpa_Stats) GetYangName() string { return "stats" }

func (stats *Dpa_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *Dpa_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *Dpa_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *Dpa_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *Dpa_Stats) GetParent() types.Entity { return stats.parent }

func (stats *Dpa_Stats) GetParentYangName() string { return "dpa" }

// Dpa_Stats_Nodes
// DPA data for available nodes
type Dpa_Stats_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DPA operational data for a particular node. The type is slice of
    // Dpa_Stats_Nodes_Node.
    Node []Dpa_Stats_Nodes_Node
}

func (nodes *Dpa_Stats_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Dpa_Stats_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Dpa_Stats_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Dpa_Stats_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Dpa_Stats_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Dpa_Stats_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Dpa_Stats_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Dpa_Stats_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Dpa_Stats_Nodes) GetYangName() string { return "nodes" }

func (nodes *Dpa_Stats_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Dpa_Stats_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Dpa_Stats_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Dpa_Stats_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Dpa_Stats_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Dpa_Stats_Nodes) GetParentYangName() string { return "stats" }

// Dpa_Stats_Nodes_Node
// DPA operational data for a particular node
type Dpa_Stats_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // DPA hw resources stats .
    HwResourcesDatas Dpa_Stats_Nodes_Node_HwResourcesDatas

    // ASIC statistics table.
    AsicStatistics Dpa_Stats_Nodes_Node_AsicStatistics

    // Ingress Stats.
    NpuNumbers Dpa_Stats_Nodes_Node_NpuNumbers

    // DPA stats hw resources info .
    NpuIds Dpa_Stats_Nodes_Node_NpuIds
}

func (node *Dpa_Stats_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Dpa_Stats_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Dpa_Stats_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "hw-resources-datas" { return "HwResourcesDatas" }
    if yname == "asic-statistics" { return "AsicStatistics" }
    if yname == "npu-numbers" { return "NpuNumbers" }
    if yname == "npu-ids" { return "NpuIds" }
    return ""
}

func (node *Dpa_Stats_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Dpa_Stats_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-resources-datas" {
        return &node.HwResourcesDatas
    }
    if childYangName == "asic-statistics" {
        return &node.AsicStatistics
    }
    if childYangName == "npu-numbers" {
        return &node.NpuNumbers
    }
    if childYangName == "npu-ids" {
        return &node.NpuIds
    }
    return nil
}

func (node *Dpa_Stats_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-resources-datas"] = &node.HwResourcesDatas
    children["asic-statistics"] = &node.AsicStatistics
    children["npu-numbers"] = &node.NpuNumbers
    children["npu-ids"] = &node.NpuIds
    return children
}

func (node *Dpa_Stats_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Dpa_Stats_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Dpa_Stats_Nodes_Node) GetYangName() string { return "node" }

func (node *Dpa_Stats_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Dpa_Stats_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Dpa_Stats_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Dpa_Stats_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Dpa_Stats_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Dpa_Stats_Nodes_Node) GetParentYangName() string { return "nodes" }

// Dpa_Stats_Nodes_Node_HwResourcesDatas
// DPA hw resources stats 
type Dpa_Stats_Nodes_Node_HwResourcesDatas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hardware resources table. The type is slice of
    // Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData.
    HwResourcesData []Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData
}

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetFilter() yfilter.YFilter { return hwResourcesDatas.YFilter }

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) SetFilter(yf yfilter.YFilter) { hwResourcesDatas.YFilter = yf }

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetGoName(yname string) string {
    if yname == "hw-resources-data" { return "HwResourcesData" }
    return ""
}

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetSegmentPath() string {
    return "hw-resources-datas"
}

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-resources-data" {
        for _, c := range hwResourcesDatas.HwResourcesData {
            if hwResourcesDatas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData{}
        hwResourcesDatas.HwResourcesData = append(hwResourcesDatas.HwResourcesData, child)
        return &hwResourcesDatas.HwResourcesData[len(hwResourcesDatas.HwResourcesData)-1]
    }
    return nil
}

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwResourcesDatas.HwResourcesData {
        children[hwResourcesDatas.HwResourcesData[i].GetSegmentPath()] = &hwResourcesDatas.HwResourcesData[i]
    }
    return children
}

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetBundleName() string { return "cisco_ios_xr" }

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetYangName() string { return "hw-resources-datas" }

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) SetParent(parent types.Entity) { hwResourcesDatas.parent = parent }

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetParent() types.Entity { return hwResourcesDatas.parent }

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetParentYangName() string { return "node" }

// Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData
// Hardware resources table
type Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Resource type. The type is Resource.
    Resource interface{}

    // resource id. The type is interface{} with range: 0..4294967295.
    ResourceId interface{}

    // name. The type is string.
    Name interface{}

    // num npus. The type is interface{} with range: 0..4294967295.
    NumNpus interface{}

    // npu hwr. The type is slice of
    // Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr.
    NpuHwr []Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr
}

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetFilter() yfilter.YFilter { return hwResourcesData.YFilter }

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) SetFilter(yf yfilter.YFilter) { hwResourcesData.YFilter = yf }

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetGoName(yname string) string {
    if yname == "resource" { return "Resource" }
    if yname == "resource-id" { return "ResourceId" }
    if yname == "name" { return "Name" }
    if yname == "num-npus" { return "NumNpus" }
    if yname == "npu-hwr" { return "NpuHwr" }
    return ""
}

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetSegmentPath() string {
    return "hw-resources-data" + "[resource='" + fmt.Sprintf("%v", hwResourcesData.Resource) + "']"
}

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "npu-hwr" {
        for _, c := range hwResourcesData.NpuHwr {
            if hwResourcesData.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr{}
        hwResourcesData.NpuHwr = append(hwResourcesData.NpuHwr, child)
        return &hwResourcesData.NpuHwr[len(hwResourcesData.NpuHwr)-1]
    }
    return nil
}

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hwResourcesData.NpuHwr {
        children[hwResourcesData.NpuHwr[i].GetSegmentPath()] = &hwResourcesData.NpuHwr[i]
    }
    return children
}

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["resource"] = hwResourcesData.Resource
    leafs["resource-id"] = hwResourcesData.ResourceId
    leafs["name"] = hwResourcesData.Name
    leafs["num-npus"] = hwResourcesData.NumNpus
    return leafs
}

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetBundleName() string { return "cisco_ios_xr" }

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetYangName() string { return "hw-resources-data" }

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) SetParent(parent types.Entity) { hwResourcesData.parent = parent }

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetParent() types.Entity { return hwResourcesData.parent }

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetParentYangName() string { return "hw-resources-datas" }

// Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr
// npu hwr
type Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // max allowed. The type is interface{} with range: 0..4294967295.
    MaxAllowed interface{}

    // npu id. The type is interface{} with range: 0..4294967295.
    NpuId interface{}

    // max entries. The type is interface{} with range: 0..4294967295.
    MaxEntries interface{}

    // red oor threshold. The type is interface{} with range: 0..4294967295.
    RedOorThreshold interface{}

    // red oor threshold percent. The type is interface{} with range:
    // 0..4294967295.
    RedOorThresholdPercent interface{}

    // yellow oor threshold. The type is interface{} with range: 0..4294967295.
    YellowOorThreshold interface{}

    // yellow oor threshold percent. The type is interface{} with range:
    // 0..4294967295.
    YellowOorThresholdPercent interface{}

    // inuse objects. The type is interface{} with range: 0..4294967295.
    InuseObjects interface{}

    // num lt. The type is interface{} with range: 0..4294967295.
    NumLt interface{}

    // oor change count. The type is interface{} with range: 0..4294967295.
    OorChangeCount interface{}

    // oor state change time1. The type is string.
    OorStateChangeTime1 interface{}

    // oor state change time2. The type is string.
    OorStateChangeTime2 interface{}

    // oor state. The type is string.
    OorState interface{}

    // lt hwr. The type is slice of
    // Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr.
    LtHwr []Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr
}

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetFilter() yfilter.YFilter { return npuHwr.YFilter }

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) SetFilter(yf yfilter.YFilter) { npuHwr.YFilter = yf }

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetGoName(yname string) string {
    if yname == "max-allowed" { return "MaxAllowed" }
    if yname == "npu-id" { return "NpuId" }
    if yname == "max-entries" { return "MaxEntries" }
    if yname == "red-oor-threshold" { return "RedOorThreshold" }
    if yname == "red-oor-threshold-percent" { return "RedOorThresholdPercent" }
    if yname == "yellow-oor-threshold" { return "YellowOorThreshold" }
    if yname == "yellow-oor-threshold-percent" { return "YellowOorThresholdPercent" }
    if yname == "inuse-objects" { return "InuseObjects" }
    if yname == "num-lt" { return "NumLt" }
    if yname == "oor-change-count" { return "OorChangeCount" }
    if yname == "oor-state-change-time1" { return "OorStateChangeTime1" }
    if yname == "oor-state-change-time2" { return "OorStateChangeTime2" }
    if yname == "oor-state" { return "OorState" }
    if yname == "lt-hwr" { return "LtHwr" }
    return ""
}

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetSegmentPath() string {
    return "npu-hwr"
}

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lt-hwr" {
        for _, c := range npuHwr.LtHwr {
            if npuHwr.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr{}
        npuHwr.LtHwr = append(npuHwr.LtHwr, child)
        return &npuHwr.LtHwr[len(npuHwr.LtHwr)-1]
    }
    return nil
}

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range npuHwr.LtHwr {
        children[npuHwr.LtHwr[i].GetSegmentPath()] = &npuHwr.LtHwr[i]
    }
    return children
}

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-allowed"] = npuHwr.MaxAllowed
    leafs["npu-id"] = npuHwr.NpuId
    leafs["max-entries"] = npuHwr.MaxEntries
    leafs["red-oor-threshold"] = npuHwr.RedOorThreshold
    leafs["red-oor-threshold-percent"] = npuHwr.RedOorThresholdPercent
    leafs["yellow-oor-threshold"] = npuHwr.YellowOorThreshold
    leafs["yellow-oor-threshold-percent"] = npuHwr.YellowOorThresholdPercent
    leafs["inuse-objects"] = npuHwr.InuseObjects
    leafs["num-lt"] = npuHwr.NumLt
    leafs["oor-change-count"] = npuHwr.OorChangeCount
    leafs["oor-state-change-time1"] = npuHwr.OorStateChangeTime1
    leafs["oor-state-change-time2"] = npuHwr.OorStateChangeTime2
    leafs["oor-state"] = npuHwr.OorState
    return leafs
}

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetBundleName() string { return "cisco_ios_xr" }

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetYangName() string { return "npu-hwr" }

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) SetParent(parent types.Entity) { npuHwr.parent = parent }

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetParent() types.Entity { return npuHwr.parent }

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetParentYangName() string { return "hw-resources-data" }

// Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr
// lt hwr
type Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lt id. The type is interface{} with range: 0..4294967295.
    LtId interface{}

    // name. The type is string.
    Name interface{}

    // hw entries. The type is interface{} with range: 0..4294967295.
    HwEntries interface{}

    // sw entries. The type is interface{} with range: 0..4294967295.
    SwEntries interface{}
}

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetFilter() yfilter.YFilter { return ltHwr.YFilter }

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) SetFilter(yf yfilter.YFilter) { ltHwr.YFilter = yf }

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetGoName(yname string) string {
    if yname == "lt-id" { return "LtId" }
    if yname == "name" { return "Name" }
    if yname == "hw-entries" { return "HwEntries" }
    if yname == "sw-entries" { return "SwEntries" }
    return ""
}

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetSegmentPath() string {
    return "lt-hwr"
}

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lt-id"] = ltHwr.LtId
    leafs["name"] = ltHwr.Name
    leafs["hw-entries"] = ltHwr.HwEntries
    leafs["sw-entries"] = ltHwr.SwEntries
    return leafs
}

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetBundleName() string { return "cisco_ios_xr" }

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetYangName() string { return "lt-hwr" }

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) SetParent(parent types.Entity) { ltHwr.parent = parent }

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetParent() types.Entity { return ltHwr.parent }

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetParentYangName() string { return "npu-hwr" }

// Dpa_Stats_Nodes_Node_AsicStatistics
// ASIC statistics table
type Dpa_Stats_Nodes_Node_AsicStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ASIC statistics.
    AsicStatisticsForNpuIds Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds

    // Detailed ASIC statistics.
    AsicStatisticsDetailForNpuIds Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds
}

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetFilter() yfilter.YFilter { return asicStatistics.YFilter }

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) SetFilter(yf yfilter.YFilter) { asicStatistics.YFilter = yf }

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetGoName(yname string) string {
    if yname == "asic-statistics-for-npu-ids" { return "AsicStatisticsForNpuIds" }
    if yname == "asic-statistics-detail-for-npu-ids" { return "AsicStatisticsDetailForNpuIds" }
    return ""
}

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetSegmentPath() string {
    return "asic-statistics"
}

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-statistics-for-npu-ids" {
        return &asicStatistics.AsicStatisticsForNpuIds
    }
    if childYangName == "asic-statistics-detail-for-npu-ids" {
        return &asicStatistics.AsicStatisticsDetailForNpuIds
    }
    return nil
}

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["asic-statistics-for-npu-ids"] = &asicStatistics.AsicStatisticsForNpuIds
    children["asic-statistics-detail-for-npu-ids"] = &asicStatistics.AsicStatisticsDetailForNpuIds
    return children
}

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetYangName() string { return "asic-statistics" }

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) SetParent(parent types.Entity) { asicStatistics.parent = parent }

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetParent() types.Entity { return asicStatistics.parent }

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetParentYangName() string { return "node" }

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds
// ASIC statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ASIC statistics for a particular NPU. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId.
    AsicStatisticsForNpuId []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId
}

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetFilter() yfilter.YFilter { return asicStatisticsForNpuIds.YFilter }

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) SetFilter(yf yfilter.YFilter) { asicStatisticsForNpuIds.YFilter = yf }

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetGoName(yname string) string {
    if yname == "asic-statistics-for-npu-id" { return "AsicStatisticsForNpuId" }
    return ""
}

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetSegmentPath() string {
    return "asic-statistics-for-npu-ids"
}

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-statistics-for-npu-id" {
        for _, c := range asicStatisticsForNpuIds.AsicStatisticsForNpuId {
            if asicStatisticsForNpuIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId{}
        asicStatisticsForNpuIds.AsicStatisticsForNpuId = append(asicStatisticsForNpuIds.AsicStatisticsForNpuId, child)
        return &asicStatisticsForNpuIds.AsicStatisticsForNpuId[len(asicStatisticsForNpuIds.AsicStatisticsForNpuId)-1]
    }
    return nil
}

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicStatisticsForNpuIds.AsicStatisticsForNpuId {
        children[asicStatisticsForNpuIds.AsicStatisticsForNpuId[i].GetSegmentPath()] = &asicStatisticsForNpuIds.AsicStatisticsForNpuId[i]
    }
    return children
}

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetBundleName() string { return "cisco_ios_xr" }

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetYangName() string { return "asic-statistics-for-npu-ids" }

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) SetParent(parent types.Entity) { asicStatisticsForNpuIds.parent = parent }

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetParent() types.Entity { return asicStatisticsForNpuIds.parent }

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetParentYangName() string { return "asic-statistics" }

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId
// ASIC statistics for a particular NPU
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NPU number. The type is interface{} with range:
    // -2147483648..2147483647.
    NpuId interface{}

    // Flag to indicate if data is valid. The type is bool.
    Valid interface{}

    // Rack number. The type is interface{} with range: 0..4294967295.
    RackNumber interface{}

    // Slot number. The type is interface{} with range: 0..4294967295.
    SlotNumber interface{}

    // ASIC instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}

    // Chip version. The type is interface{} with range: 0..65535.
    ChipVersion interface{}

    // Statistics.
    Statistics Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics
}

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetFilter() yfilter.YFilter { return asicStatisticsForNpuId.YFilter }

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) SetFilter(yf yfilter.YFilter) { asicStatisticsForNpuId.YFilter = yf }

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "valid" { return "Valid" }
    if yname == "rack-number" { return "RackNumber" }
    if yname == "slot-number" { return "SlotNumber" }
    if yname == "asic-instance" { return "AsicInstance" }
    if yname == "chip-version" { return "ChipVersion" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetSegmentPath() string {
    return "asic-statistics-for-npu-id" + "[npu-id='" + fmt.Sprintf("%v", asicStatisticsForNpuId.NpuId) + "']"
}

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &asicStatisticsForNpuId.Statistics
    }
    return nil
}

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &asicStatisticsForNpuId.Statistics
    return children
}

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = asicStatisticsForNpuId.NpuId
    leafs["valid"] = asicStatisticsForNpuId.Valid
    leafs["rack-number"] = asicStatisticsForNpuId.RackNumber
    leafs["slot-number"] = asicStatisticsForNpuId.SlotNumber
    leafs["asic-instance"] = asicStatisticsForNpuId.AsicInstance
    leafs["chip-version"] = asicStatisticsForNpuId.ChipVersion
    return leafs
}

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetBundleName() string { return "cisco_ios_xr" }

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetYangName() string { return "asic-statistics-for-npu-id" }

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) SetParent(parent types.Entity) { asicStatisticsForNpuId.parent = parent }

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetParent() types.Entity { return asicStatisticsForNpuId.parent }

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetParentYangName() string { return "asic-statistics-for-npu-ids" }

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics
// Statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total bytes sent from NIF to IRE. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    NbiRxTotalByteCnt interface{}

    // Total packets sent from NIF to IRE. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxTotalPktCnt interface{}

    // CPU ingress received packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IreCpuPktCnt interface{}

    // NIF received packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IreNifPktCnt interface{}

    // OAMP ingress received packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IreOampPktCnt interface{}

    // OLP ingress received packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IreOlpPktCnt interface{}

    // Recycling ingress received packet counter. The type is interface{} with
    // range: 0..18446744073709551615.
    IreRcyPktCnt interface{}

    // Performance counter of the FDT interface. The type is interface{} with
    // range: 0..18446744073709551615.
    IreFdtIfCnt interface{}

    // Performance counter of the MMU interface. The type is interface{} with
    // range: 0..18446744073709551615.
    IdrMmuIfCnt interface{}

    // Performance counter of the OCB interface. The type is interface{} with
    // range: 0..18446744073709551615.
    IdrOcbIfCnt interface{}

    // Counts enqueued packets. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmEnqueuePktCnt interface{}

    // Counts dequeued packets. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmDequeuePktCnt interface{}

    // Counts matched packets discarded in the DEQ process. The type is
    // interface{} with range: 0..18446744073709551615.
    IqmDeletedPktCnt interface{}

    // Counts all packets discarded at the ENQ pipe. The type is interface{} with
    // range: 0..18446744073709551615.
    IqmEnqDiscardedPktCnt interface{}

    // EGQ packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IptEgqPktCnt interface{}

    // ENQ packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IptEnqPktCnt interface{}

    // FDT packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IptFdtPktCnt interface{}

    // Configurable event counter. The type is interface{} with range:
    // 0..18446744073709551615.
    IptCfgEventCnt interface{}

    // Configurable bytes counter. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    IptCfgByteCnt interface{}

    // Descriptor cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtIptDescCellCnt interface{}

    // IRE internal descriptor cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtIreDescCellCnt interface{}

    // Counts all transmitted data cells. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtTransmittedDataCellsCnt interface{}

    // FDR total incoming cell counter at pipe 1. The type is interface{} with
    // range: 0..18446744073709551615.
    FdrP1CellInCnt interface{}

    // FDR total incoming cell counter at pipe 2. The type is interface{} with
    // range: 0..18446744073709551615.
    FdrP2CellInCnt interface{}

    // FDR total incoming cell counter at pipe 3. The type is interface{} with
    // range: 0..18446744073709551615.
    FdrP3CellInCnt interface{}

    // FDR total incoming cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrCellInCntTotal interface{}

    // FDA input cell counter P1. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInCntP1 interface{}

    // FDA input cell counter P2. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInCntP2 interface{}

    // FDA input cell counter P3. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInCntP3 interface{}

    // FDA input cell counter TDM. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInTdmCnt interface{}

    // FDA input cell counter MESHMC. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInMeshmcCnt interface{}

    // FDA input cell counter IPT. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsInIptCnt interface{}

    // FDA output cell counter P1. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutCntP1 interface{}

    // FDA output cell counter P2. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutCntP2 interface{}

    // FDA output cell counter P3. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutCntP3 interface{}

    // FDA output cell counter TDM. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutTdmCnt interface{}

    // FDA output cell counter MESHMC. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutMeshmcCnt interface{}

    // FDA output cell counter IPT. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaCellsOutIptCnt interface{}

    // FDA EGQ drop counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaEgqDropCnt interface{}

    // FDA EGQ MESHMC drop counter. The type is interface{} with range:
    // 0..18446744073709551615.
    FdaEgqMeshmcDropCnt interface{}

    // FQP2EPE packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqFqpPktCnt interface{}

    // PQP2FQP unicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpUcPktCnt interface{}

    // PQP discarded unicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpDiscardUcPktCnt interface{}

    // PQP2FQP unicast bytes counter. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    EgqPqpUcBytesCnt interface{}

    // PQP2FQP multicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpMcPktCnt interface{}

    // PQP discarded multicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpDiscardMcPktCnt interface{}

    // PQP2FQP multicast bytes counter. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    EgqPqpMcBytesCnt interface{}

    // EHP2PQP unicast packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEhpUcPktCnt interface{}

    // EHP2PQP multicast high packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEhpMcHighPktCnt interface{}

    // EHP2PQP multicast low packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEhpMcLowPktCnt interface{}

    // EHP2PQP discarded packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqDeletedPktCnt interface{}

    // Number of multicast high packets discarded because multicast FIFO is full.
    // The type is interface{} with range: 0..18446744073709551615.
    EgqEhpMcHighDiscardCnt interface{}

    // Number of multicast low packets discarded because multicast FIFO is full.
    // The type is interface{} with range: 0..18446744073709551615.
    EgqEhpMcLowDiscardCnt interface{}

    // Number of packet descriptors discarded due to LAG multicast pruning. The
    // type is interface{} with range: 0..18446744073709551615.
    EgqErppLagPruningDiscardCnt interface{}

    // Number of packet descriptors discarded due to ERPP PMF. The type is
    // interface{} with range: 0..18446744073709551615.
    EgqErppPmfDiscardCnt interface{}

    // Number of packet descriptors discarded because of egress VLAN membership.
    // The type is interface{} with range: 0..18446744073709551615.
    EgqErppVlanMbrDiscardCnt interface{}

    // EPE2PNI bytes counter. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    EpniEpeByteCnt interface{}

    // EPE2PNI packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniEpePktCnt interface{}

    // EPE discarded packet counter. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniEpeDiscardCnt interface{}

    // Total bytes sent from EGQ to NIF. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    NbiTxTotalByteCnt interface{}

    // Total packets sent from EGQ to NIF. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiTxTotalPktCnt interface{}
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetGoName(yname string) string {
    if yname == "nbi-rx-total-byte-cnt" { return "NbiRxTotalByteCnt" }
    if yname == "nbi-rx-total-pkt-cnt" { return "NbiRxTotalPktCnt" }
    if yname == "ire-cpu-pkt-cnt" { return "IreCpuPktCnt" }
    if yname == "ire-nif-pkt-cnt" { return "IreNifPktCnt" }
    if yname == "ire-oamp-pkt-cnt" { return "IreOampPktCnt" }
    if yname == "ire-olp-pkt-cnt" { return "IreOlpPktCnt" }
    if yname == "ire-rcy-pkt-cnt" { return "IreRcyPktCnt" }
    if yname == "ire-fdt-if-cnt" { return "IreFdtIfCnt" }
    if yname == "idr-mmu-if-cnt" { return "IdrMmuIfCnt" }
    if yname == "idr-ocb-if-cnt" { return "IdrOcbIfCnt" }
    if yname == "iqm-enqueue-pkt-cnt" { return "IqmEnqueuePktCnt" }
    if yname == "iqm-dequeue-pkt-cnt" { return "IqmDequeuePktCnt" }
    if yname == "iqm-deleted-pkt-cnt" { return "IqmDeletedPktCnt" }
    if yname == "iqm-enq-discarded-pkt-cnt" { return "IqmEnqDiscardedPktCnt" }
    if yname == "ipt-egq-pkt-cnt" { return "IptEgqPktCnt" }
    if yname == "ipt-enq-pkt-cnt" { return "IptEnqPktCnt" }
    if yname == "ipt-fdt-pkt-cnt" { return "IptFdtPktCnt" }
    if yname == "ipt-cfg-event-cnt" { return "IptCfgEventCnt" }
    if yname == "ipt-cfg-byte-cnt" { return "IptCfgByteCnt" }
    if yname == "fdt-ipt-desc-cell-cnt" { return "FdtIptDescCellCnt" }
    if yname == "fdt-ire-desc-cell-cnt" { return "FdtIreDescCellCnt" }
    if yname == "fdt-transmitted-data-cells-cnt" { return "FdtTransmittedDataCellsCnt" }
    if yname == "fdr-p1-cell-in-cnt" { return "FdrP1CellInCnt" }
    if yname == "fdr-p2-cell-in-cnt" { return "FdrP2CellInCnt" }
    if yname == "fdr-p3-cell-in-cnt" { return "FdrP3CellInCnt" }
    if yname == "fdr-cell-in-cnt-total" { return "FdrCellInCntTotal" }
    if yname == "fda-cells-in-cnt-p1" { return "FdaCellsInCntP1" }
    if yname == "fda-cells-in-cnt-p2" { return "FdaCellsInCntP2" }
    if yname == "fda-cells-in-cnt-p3" { return "FdaCellsInCntP3" }
    if yname == "fda-cells-in-tdm-cnt" { return "FdaCellsInTdmCnt" }
    if yname == "fda-cells-in-meshmc-cnt" { return "FdaCellsInMeshmcCnt" }
    if yname == "fda-cells-in-ipt-cnt" { return "FdaCellsInIptCnt" }
    if yname == "fda-cells-out-cnt-p1" { return "FdaCellsOutCntP1" }
    if yname == "fda-cells-out-cnt-p2" { return "FdaCellsOutCntP2" }
    if yname == "fda-cells-out-cnt-p3" { return "FdaCellsOutCntP3" }
    if yname == "fda-cells-out-tdm-cnt" { return "FdaCellsOutTdmCnt" }
    if yname == "fda-cells-out-meshmc-cnt" { return "FdaCellsOutMeshmcCnt" }
    if yname == "fda-cells-out-ipt-cnt" { return "FdaCellsOutIptCnt" }
    if yname == "fda-egq-drop-cnt" { return "FdaEgqDropCnt" }
    if yname == "fda-egq-meshmc-drop-cnt" { return "FdaEgqMeshmcDropCnt" }
    if yname == "egq-fqp-pkt-cnt" { return "EgqFqpPktCnt" }
    if yname == "egq-pqp-uc-pkt-cnt" { return "EgqPqpUcPktCnt" }
    if yname == "egq-pqp-discard-uc-pkt-cnt" { return "EgqPqpDiscardUcPktCnt" }
    if yname == "egq-pqp-uc-bytes-cnt" { return "EgqPqpUcBytesCnt" }
    if yname == "egq-pqp-mc-pkt-cnt" { return "EgqPqpMcPktCnt" }
    if yname == "egq-pqp-discard-mc-pkt-cnt" { return "EgqPqpDiscardMcPktCnt" }
    if yname == "egq-pqp-mc-bytes-cnt" { return "EgqPqpMcBytesCnt" }
    if yname == "egq-ehp-uc-pkt-cnt" { return "EgqEhpUcPktCnt" }
    if yname == "egq-ehp-mc-high-pkt-cnt" { return "EgqEhpMcHighPktCnt" }
    if yname == "egq-ehp-mc-low-pkt-cnt" { return "EgqEhpMcLowPktCnt" }
    if yname == "egq-deleted-pkt-cnt" { return "EgqDeletedPktCnt" }
    if yname == "egq-ehp-mc-high-discard-cnt" { return "EgqEhpMcHighDiscardCnt" }
    if yname == "egq-ehp-mc-low-discard-cnt" { return "EgqEhpMcLowDiscardCnt" }
    if yname == "egq-erpp-lag-pruning-discard-cnt" { return "EgqErppLagPruningDiscardCnt" }
    if yname == "egq-erpp-pmf-discard-cnt" { return "EgqErppPmfDiscardCnt" }
    if yname == "egq-erpp-vlan-mbr-discard-cnt" { return "EgqErppVlanMbrDiscardCnt" }
    if yname == "epni-epe-byte-cnt" { return "EpniEpeByteCnt" }
    if yname == "epni-epe-pkt-cnt" { return "EpniEpePktCnt" }
    if yname == "epni-epe-discard-cnt" { return "EpniEpeDiscardCnt" }
    if yname == "nbi-tx-total-byte-cnt" { return "NbiTxTotalByteCnt" }
    if yname == "nbi-tx-total-pkt-cnt" { return "NbiTxTotalPktCnt" }
    return ""
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nbi-rx-total-byte-cnt"] = statistics.NbiRxTotalByteCnt
    leafs["nbi-rx-total-pkt-cnt"] = statistics.NbiRxTotalPktCnt
    leafs["ire-cpu-pkt-cnt"] = statistics.IreCpuPktCnt
    leafs["ire-nif-pkt-cnt"] = statistics.IreNifPktCnt
    leafs["ire-oamp-pkt-cnt"] = statistics.IreOampPktCnt
    leafs["ire-olp-pkt-cnt"] = statistics.IreOlpPktCnt
    leafs["ire-rcy-pkt-cnt"] = statistics.IreRcyPktCnt
    leafs["ire-fdt-if-cnt"] = statistics.IreFdtIfCnt
    leafs["idr-mmu-if-cnt"] = statistics.IdrMmuIfCnt
    leafs["idr-ocb-if-cnt"] = statistics.IdrOcbIfCnt
    leafs["iqm-enqueue-pkt-cnt"] = statistics.IqmEnqueuePktCnt
    leafs["iqm-dequeue-pkt-cnt"] = statistics.IqmDequeuePktCnt
    leafs["iqm-deleted-pkt-cnt"] = statistics.IqmDeletedPktCnt
    leafs["iqm-enq-discarded-pkt-cnt"] = statistics.IqmEnqDiscardedPktCnt
    leafs["ipt-egq-pkt-cnt"] = statistics.IptEgqPktCnt
    leafs["ipt-enq-pkt-cnt"] = statistics.IptEnqPktCnt
    leafs["ipt-fdt-pkt-cnt"] = statistics.IptFdtPktCnt
    leafs["ipt-cfg-event-cnt"] = statistics.IptCfgEventCnt
    leafs["ipt-cfg-byte-cnt"] = statistics.IptCfgByteCnt
    leafs["fdt-ipt-desc-cell-cnt"] = statistics.FdtIptDescCellCnt
    leafs["fdt-ire-desc-cell-cnt"] = statistics.FdtIreDescCellCnt
    leafs["fdt-transmitted-data-cells-cnt"] = statistics.FdtTransmittedDataCellsCnt
    leafs["fdr-p1-cell-in-cnt"] = statistics.FdrP1CellInCnt
    leafs["fdr-p2-cell-in-cnt"] = statistics.FdrP2CellInCnt
    leafs["fdr-p3-cell-in-cnt"] = statistics.FdrP3CellInCnt
    leafs["fdr-cell-in-cnt-total"] = statistics.FdrCellInCntTotal
    leafs["fda-cells-in-cnt-p1"] = statistics.FdaCellsInCntP1
    leafs["fda-cells-in-cnt-p2"] = statistics.FdaCellsInCntP2
    leafs["fda-cells-in-cnt-p3"] = statistics.FdaCellsInCntP3
    leafs["fda-cells-in-tdm-cnt"] = statistics.FdaCellsInTdmCnt
    leafs["fda-cells-in-meshmc-cnt"] = statistics.FdaCellsInMeshmcCnt
    leafs["fda-cells-in-ipt-cnt"] = statistics.FdaCellsInIptCnt
    leafs["fda-cells-out-cnt-p1"] = statistics.FdaCellsOutCntP1
    leafs["fda-cells-out-cnt-p2"] = statistics.FdaCellsOutCntP2
    leafs["fda-cells-out-cnt-p3"] = statistics.FdaCellsOutCntP3
    leafs["fda-cells-out-tdm-cnt"] = statistics.FdaCellsOutTdmCnt
    leafs["fda-cells-out-meshmc-cnt"] = statistics.FdaCellsOutMeshmcCnt
    leafs["fda-cells-out-ipt-cnt"] = statistics.FdaCellsOutIptCnt
    leafs["fda-egq-drop-cnt"] = statistics.FdaEgqDropCnt
    leafs["fda-egq-meshmc-drop-cnt"] = statistics.FdaEgqMeshmcDropCnt
    leafs["egq-fqp-pkt-cnt"] = statistics.EgqFqpPktCnt
    leafs["egq-pqp-uc-pkt-cnt"] = statistics.EgqPqpUcPktCnt
    leafs["egq-pqp-discard-uc-pkt-cnt"] = statistics.EgqPqpDiscardUcPktCnt
    leafs["egq-pqp-uc-bytes-cnt"] = statistics.EgqPqpUcBytesCnt
    leafs["egq-pqp-mc-pkt-cnt"] = statistics.EgqPqpMcPktCnt
    leafs["egq-pqp-discard-mc-pkt-cnt"] = statistics.EgqPqpDiscardMcPktCnt
    leafs["egq-pqp-mc-bytes-cnt"] = statistics.EgqPqpMcBytesCnt
    leafs["egq-ehp-uc-pkt-cnt"] = statistics.EgqEhpUcPktCnt
    leafs["egq-ehp-mc-high-pkt-cnt"] = statistics.EgqEhpMcHighPktCnt
    leafs["egq-ehp-mc-low-pkt-cnt"] = statistics.EgqEhpMcLowPktCnt
    leafs["egq-deleted-pkt-cnt"] = statistics.EgqDeletedPktCnt
    leafs["egq-ehp-mc-high-discard-cnt"] = statistics.EgqEhpMcHighDiscardCnt
    leafs["egq-ehp-mc-low-discard-cnt"] = statistics.EgqEhpMcLowDiscardCnt
    leafs["egq-erpp-lag-pruning-discard-cnt"] = statistics.EgqErppLagPruningDiscardCnt
    leafs["egq-erpp-pmf-discard-cnt"] = statistics.EgqErppPmfDiscardCnt
    leafs["egq-erpp-vlan-mbr-discard-cnt"] = statistics.EgqErppVlanMbrDiscardCnt
    leafs["epni-epe-byte-cnt"] = statistics.EpniEpeByteCnt
    leafs["epni-epe-pkt-cnt"] = statistics.EpniEpePktCnt
    leafs["epni-epe-discard-cnt"] = statistics.EpniEpeDiscardCnt
    leafs["nbi-tx-total-byte-cnt"] = statistics.NbiTxTotalByteCnt
    leafs["nbi-tx-total-pkt-cnt"] = statistics.NbiTxTotalPktCnt
    return leafs
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetParentYangName() string { return "asic-statistics-for-npu-id" }

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds
// Detailed ASIC statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed ASIC statistics for a particular NPU. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId.
    AsicStatisticsDetailForNpuId []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId
}

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetFilter() yfilter.YFilter { return asicStatisticsDetailForNpuIds.YFilter }

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) SetFilter(yf yfilter.YFilter) { asicStatisticsDetailForNpuIds.YFilter = yf }

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetGoName(yname string) string {
    if yname == "asic-statistics-detail-for-npu-id" { return "AsicStatisticsDetailForNpuId" }
    return ""
}

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetSegmentPath() string {
    return "asic-statistics-detail-for-npu-ids"
}

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-statistics-detail-for-npu-id" {
        for _, c := range asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId {
            if asicStatisticsDetailForNpuIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId{}
        asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId = append(asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId, child)
        return &asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId[len(asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId)-1]
    }
    return nil
}

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId {
        children[asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId[i].GetSegmentPath()] = &asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId[i]
    }
    return children
}

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetBundleName() string { return "cisco_ios_xr" }

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetYangName() string { return "asic-statistics-detail-for-npu-ids" }

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) SetParent(parent types.Entity) { asicStatisticsDetailForNpuIds.parent = parent }

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetParent() types.Entity { return asicStatisticsDetailForNpuIds.parent }

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetParentYangName() string { return "asic-statistics" }

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId
// Detailed ASIC statistics for a particular
// NPU
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NPU number. The type is interface{} with range:
    // -2147483648..2147483647.
    NpuId interface{}

    // Flag to indicate if data is valid. The type is bool.
    Valid interface{}

    // Rack number. The type is interface{} with range: 0..4294967295.
    RackNumber interface{}

    // Slot number. The type is interface{} with range: 0..4294967295.
    SlotNumber interface{}

    // ASIC instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}

    // Chip version. The type is interface{} with range: 0..65535.
    ChipVersion interface{}

    // Statistics.
    Statistics Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics
}

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetFilter() yfilter.YFilter { return asicStatisticsDetailForNpuId.YFilter }

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) SetFilter(yf yfilter.YFilter) { asicStatisticsDetailForNpuId.YFilter = yf }

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "valid" { return "Valid" }
    if yname == "rack-number" { return "RackNumber" }
    if yname == "slot-number" { return "SlotNumber" }
    if yname == "asic-instance" { return "AsicInstance" }
    if yname == "chip-version" { return "ChipVersion" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetSegmentPath() string {
    return "asic-statistics-detail-for-npu-id" + "[npu-id='" + fmt.Sprintf("%v", asicStatisticsDetailForNpuId.NpuId) + "']"
}

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &asicStatisticsDetailForNpuId.Statistics
    }
    return nil
}

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &asicStatisticsDetailForNpuId.Statistics
    return children
}

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = asicStatisticsDetailForNpuId.NpuId
    leafs["valid"] = asicStatisticsDetailForNpuId.Valid
    leafs["rack-number"] = asicStatisticsDetailForNpuId.RackNumber
    leafs["slot-number"] = asicStatisticsDetailForNpuId.SlotNumber
    leafs["asic-instance"] = asicStatisticsDetailForNpuId.AsicInstance
    leafs["chip-version"] = asicStatisticsDetailForNpuId.ChipVersion
    return leafs
}

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetBundleName() string { return "cisco_ios_xr" }

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetYangName() string { return "asic-statistics-detail-for-npu-id" }

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) SetParent(parent types.Entity) { asicStatisticsDetailForNpuId.parent = parent }

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetParent() types.Entity { return asicStatisticsDetailForNpuId.parent }

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetParentYangName() string { return "asic-statistics-detail-for-npu-ids" }

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics
// Statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of blocks. The type is interface{} with range: 0..255.
    NumBlocks interface{}

    // Block information. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo.
    BlockInfo []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetGoName(yname string) string {
    if yname == "num-blocks" { return "NumBlocks" }
    if yname == "block-info" { return "BlockInfo" }
    return ""
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "block-info" {
        for _, c := range statistics.BlockInfo {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo{}
        statistics.BlockInfo = append(statistics.BlockInfo, child)
        return &statistics.BlockInfo[len(statistics.BlockInfo)-1]
    }
    return nil
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.BlockInfo {
        children[statistics.BlockInfo[i].GetSegmentPath()] = &statistics.BlockInfo[i]
    }
    return children
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-blocks"] = statistics.NumBlocks
    return leafs
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetYangName() string { return "statistics" }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetParentYangName() string { return "asic-statistics-detail-for-npu-id" }

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo
// Block information
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Block name. The type is string with length: 0..10.
    BlockName interface{}

    // Number of fields. The type is interface{} with range: 0..255.
    NumFields interface{}

    // Field information. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo.
    FieldInfo []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo
}

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetFilter() yfilter.YFilter { return blockInfo.YFilter }

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) SetFilter(yf yfilter.YFilter) { blockInfo.YFilter = yf }

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetGoName(yname string) string {
    if yname == "block-name" { return "BlockName" }
    if yname == "num-fields" { return "NumFields" }
    if yname == "field-info" { return "FieldInfo" }
    return ""
}

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetSegmentPath() string {
    return "block-info"
}

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "field-info" {
        for _, c := range blockInfo.FieldInfo {
            if blockInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo{}
        blockInfo.FieldInfo = append(blockInfo.FieldInfo, child)
        return &blockInfo.FieldInfo[len(blockInfo.FieldInfo)-1]
    }
    return nil
}

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range blockInfo.FieldInfo {
        children[blockInfo.FieldInfo[i].GetSegmentPath()] = &blockInfo.FieldInfo[i]
    }
    return children
}

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["block-name"] = blockInfo.BlockName
    leafs["num-fields"] = blockInfo.NumFields
    return leafs
}

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetBundleName() string { return "cisco_ios_xr" }

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetYangName() string { return "block-info" }

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) SetParent(parent types.Entity) { blockInfo.parent = parent }

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetParent() types.Entity { return blockInfo.parent }

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetParentYangName() string { return "statistics" }

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo
// Field information
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field name. The type is string with length: 0..80.
    FieldName interface{}

    // Field value. The type is interface{} with range: 0..18446744073709551615.
    FieldValue interface{}

    // Flag to indicate overflow. The type is bool.
    IsOverflow interface{}
}

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetFilter() yfilter.YFilter { return fieldInfo.YFilter }

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) SetFilter(yf yfilter.YFilter) { fieldInfo.YFilter = yf }

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetGoName(yname string) string {
    if yname == "field-name" { return "FieldName" }
    if yname == "field-value" { return "FieldValue" }
    if yname == "is-overflow" { return "IsOverflow" }
    return ""
}

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetSegmentPath() string {
    return "field-info"
}

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["field-name"] = fieldInfo.FieldName
    leafs["field-value"] = fieldInfo.FieldValue
    leafs["is-overflow"] = fieldInfo.IsOverflow
    return leafs
}

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetYangName() string { return "field-info" }

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) SetParent(parent types.Entity) { fieldInfo.parent = parent }

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetParent() types.Entity { return fieldInfo.parent }

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetParentYangName() string { return "block-info" }

// Dpa_Stats_Nodes_Node_NpuNumbers
// Ingress Stats
type Dpa_Stats_Nodes_Node_NpuNumbers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats for a particular npu. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber.
    NpuNumber []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber
}

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetFilter() yfilter.YFilter { return npuNumbers.YFilter }

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) SetFilter(yf yfilter.YFilter) { npuNumbers.YFilter = yf }

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetGoName(yname string) string {
    if yname == "npu-number" { return "NpuNumber" }
    return ""
}

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetSegmentPath() string {
    return "npu-numbers"
}

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "npu-number" {
        for _, c := range npuNumbers.NpuNumber {
            if npuNumbers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber{}
        npuNumbers.NpuNumber = append(npuNumbers.NpuNumber, child)
        return &npuNumbers.NpuNumber[len(npuNumbers.NpuNumber)-1]
    }
    return nil
}

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range npuNumbers.NpuNumber {
        children[npuNumbers.NpuNumber[i].GetSegmentPath()] = &npuNumbers.NpuNumber[i]
    }
    return children
}

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetBundleName() string { return "cisco_ios_xr" }

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetYangName() string { return "npu-numbers" }

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) SetParent(parent types.Entity) { npuNumbers.parent = parent }

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetParent() types.Entity { return npuNumbers.parent }

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetParentYangName() string { return "node" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber
// Stats for a particular npu
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Npu number. The type is interface{} with range:
    // -2147483648..2147483647.
    NpuId interface{}

    // show npu specific voq or trap stats.
    Display Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display
}

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetFilter() yfilter.YFilter { return npuNumber.YFilter }

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) SetFilter(yf yfilter.YFilter) { npuNumber.YFilter = yf }

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "display" { return "Display" }
    return ""
}

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetSegmentPath() string {
    return "npu-number" + "[npu-id='" + fmt.Sprintf("%v", npuNumber.NpuId) + "']"
}

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "display" {
        return &npuNumber.Display
    }
    return nil
}

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["display"] = &npuNumber.Display
    return children
}

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = npuNumber.NpuId
    return leafs
}

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetBundleName() string { return "cisco_ios_xr" }

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetYangName() string { return "npu-number" }

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) SetParent(parent types.Entity) { npuNumber.parent = parent }

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetParent() types.Entity { return npuNumber.parent }

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetParentYangName() string { return "npu-numbers" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display
// show npu specific voq or trap stats
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Voq stats grouped by voq base numbers.
    BaseNumbers Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers

    // Trap stats for a particular npu.
    TrapIds Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds

    // Voq stats grouped by interface handle.
    InterfaceHandles Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles
}

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetFilter() yfilter.YFilter { return display.YFilter }

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) SetFilter(yf yfilter.YFilter) { display.YFilter = yf }

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetGoName(yname string) string {
    if yname == "base-numbers" { return "BaseNumbers" }
    if yname == "trap-ids" { return "TrapIds" }
    if yname == "interface-handles" { return "InterfaceHandles" }
    return ""
}

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetSegmentPath() string {
    return "display"
}

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "base-numbers" {
        return &display.BaseNumbers
    }
    if childYangName == "trap-ids" {
        return &display.TrapIds
    }
    if childYangName == "interface-handles" {
        return &display.InterfaceHandles
    }
    return nil
}

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["base-numbers"] = &display.BaseNumbers
    children["trap-ids"] = &display.TrapIds
    children["interface-handles"] = &display.InterfaceHandles
    return children
}

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetBundleName() string { return "cisco_ios_xr" }

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetYangName() string { return "display" }

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) SetParent(parent types.Entity) { display.parent = parent }

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetParent() types.Entity { return display.parent }

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetParentYangName() string { return "npu-number" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers
// Voq stats grouped by voq base numbers
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Voq Base Number for a particular voq. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber.
    BaseNumber []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber
}

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetFilter() yfilter.YFilter { return baseNumbers.YFilter }

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) SetFilter(yf yfilter.YFilter) { baseNumbers.YFilter = yf }

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetGoName(yname string) string {
    if yname == "base-number" { return "BaseNumber" }
    return ""
}

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetSegmentPath() string {
    return "base-numbers"
}

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "base-number" {
        for _, c := range baseNumbers.BaseNumber {
            if baseNumbers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber{}
        baseNumbers.BaseNumber = append(baseNumbers.BaseNumber, child)
        return &baseNumbers.BaseNumber[len(baseNumbers.BaseNumber)-1]
    }
    return nil
}

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range baseNumbers.BaseNumber {
        children[baseNumbers.BaseNumber[i].GetSegmentPath()] = &baseNumbers.BaseNumber[i]
    }
    return children
}

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetBundleName() string { return "cisco_ios_xr" }

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetYangName() string { return "base-numbers" }

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) SetParent(parent types.Entity) { baseNumbers.parent = parent }

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetParent() types.Entity { return baseNumbers.parent }

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetParentYangName() string { return "display" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber
// Voq Base Number for a particular voq
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface handle. The type is interface{} with
    // range: 0..4294967295.
    BaseNumber interface{}

    // Flag to indicate if port is in use. The type is bool.
    InUse interface{}

    // Rack of port. The type is interface{} with range: 0..255.
    RackNum interface{}

    // Slot of port. The type is interface{} with range: 0..255.
    SlotNum interface{}

    // NPU of port. The type is interface{} with range: 0..255.
    NpuNum interface{}

    // NPU core of port. The type is interface{} with range: 0..255.
    NpuCore interface{}

    // Port Number of port. The type is interface{} with range: 0..255.
    PortNum interface{}

    // IfHandle of port. The type is interface{} with range: 0..4294967295.
    IfHandle interface{}

    // System port of port. The type is interface{} with range: 0..4294967295.
    SysPort interface{}

    // PP Port number of port. The type is interface{} with range: 0..4294967295.
    PpPort interface{}

    // Port speed of port. The type is interface{} with range: 0..4294967295.
    PortSpeed interface{}

    // Voq Base number of port. The type is interface{} with range: 0..4294967295.
    VoqBase interface{}

    // Connector id of port. The type is interface{} with range: 0..4294967295.
    ConnectorId interface{}

    // Flag to indicate if port is local to the node. The type is bool.
    IsLocalPort interface{}

    // Keeps a record of the received and dropped packets and bytes on the port.
    // The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat.
    VoqStat []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat
}

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetFilter() yfilter.YFilter { return baseNumber.YFilter }

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) SetFilter(yf yfilter.YFilter) { baseNumber.YFilter = yf }

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetGoName(yname string) string {
    if yname == "base-number" { return "BaseNumber" }
    if yname == "in-use" { return "InUse" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "slot-num" { return "SlotNum" }
    if yname == "npu-num" { return "NpuNum" }
    if yname == "npu-core" { return "NpuCore" }
    if yname == "port-num" { return "PortNum" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "sys-port" { return "SysPort" }
    if yname == "pp-port" { return "PpPort" }
    if yname == "port-speed" { return "PortSpeed" }
    if yname == "voq-base" { return "VoqBase" }
    if yname == "connector-id" { return "ConnectorId" }
    if yname == "is-local-port" { return "IsLocalPort" }
    if yname == "voq-stat" { return "VoqStat" }
    return ""
}

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetSegmentPath() string {
    return "base-number" + "[base-number='" + fmt.Sprintf("%v", baseNumber.BaseNumber) + "']"
}

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "voq-stat" {
        for _, c := range baseNumber.VoqStat {
            if baseNumber.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat{}
        baseNumber.VoqStat = append(baseNumber.VoqStat, child)
        return &baseNumber.VoqStat[len(baseNumber.VoqStat)-1]
    }
    return nil
}

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range baseNumber.VoqStat {
        children[baseNumber.VoqStat[i].GetSegmentPath()] = &baseNumber.VoqStat[i]
    }
    return children
}

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["base-number"] = baseNumber.BaseNumber
    leafs["in-use"] = baseNumber.InUse
    leafs["rack-num"] = baseNumber.RackNum
    leafs["slot-num"] = baseNumber.SlotNum
    leafs["npu-num"] = baseNumber.NpuNum
    leafs["npu-core"] = baseNumber.NpuCore
    leafs["port-num"] = baseNumber.PortNum
    leafs["if-handle"] = baseNumber.IfHandle
    leafs["sys-port"] = baseNumber.SysPort
    leafs["pp-port"] = baseNumber.PpPort
    leafs["port-speed"] = baseNumber.PortSpeed
    leafs["voq-base"] = baseNumber.VoqBase
    leafs["connector-id"] = baseNumber.ConnectorId
    leafs["is-local-port"] = baseNumber.IsLocalPort
    return leafs
}

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetBundleName() string { return "cisco_ios_xr" }

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetYangName() string { return "base-number" }

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) SetParent(parent types.Entity) { baseNumber.parent = parent }

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetParent() types.Entity { return baseNumber.parent }

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetParentYangName() string { return "base-numbers" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat
// Keeps a record of the received and dropped
// packets and bytes on the port
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bytes Received on the port. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ReceivedBytes interface{}

    // Packets Received on the port. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Bytes Dropped on the port. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    DroppedBytes interface{}

    // Packets Dropeed on the port. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetFilter() yfilter.YFilter { return voqStat.YFilter }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) SetFilter(yf yfilter.YFilter) { voqStat.YFilter = yf }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetGoName(yname string) string {
    if yname == "received-bytes" { return "ReceivedBytes" }
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "dropped-bytes" { return "DroppedBytes" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetSegmentPath() string {
    return "voq-stat"
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-bytes"] = voqStat.ReceivedBytes
    leafs["received-packets"] = voqStat.ReceivedPackets
    leafs["dropped-bytes"] = voqStat.DroppedBytes
    leafs["dropped-packets"] = voqStat.DroppedPackets
    return leafs
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetBundleName() string { return "cisco_ios_xr" }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetYangName() string { return "voq-stat" }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) SetParent(parent types.Entity) { voqStat.parent = parent }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetParent() types.Entity { return voqStat.parent }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetParentYangName() string { return "base-number" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds
// Trap stats for a particular npu
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Filter by specific trap id. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId.
    TrapId []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId
}

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetFilter() yfilter.YFilter { return trapIds.YFilter }

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) SetFilter(yf yfilter.YFilter) { trapIds.YFilter = yf }

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetGoName(yname string) string {
    if yname == "trap-id" { return "TrapId" }
    return ""
}

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetSegmentPath() string {
    return "trap-ids"
}

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-id" {
        for _, c := range trapIds.TrapId {
            if trapIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId{}
        trapIds.TrapId = append(trapIds.TrapId, child)
        return &trapIds.TrapId[len(trapIds.TrapId)-1]
    }
    return nil
}

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapIds.TrapId {
        children[trapIds.TrapId[i].GetSegmentPath()] = &trapIds.TrapId[i]
    }
    return children
}

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetBundleName() string { return "cisco_ios_xr" }

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetYangName() string { return "trap-ids" }

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) SetParent(parent types.Entity) { trapIds.parent = parent }

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetParent() types.Entity { return trapIds.parent }

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetParentYangName() string { return "display" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId
// Filter by specific trap id
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Trap ID. The type is interface{} with range:
    // 0..4294967295.
    TrapId interface{}

    // Trap Strength of the trap. The type is interface{} with range:
    // 0..4294967295.
    TrapStrength interface{}

    // Priority of the trap. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Id of the trap. The type is interface{} with range: 0..4294967295.
    TrapIdXr interface{}

    // Gport of the trap. The type is interface{} with range: 0..4294967295.
    Gport interface{}

    // Fec id of the trap. The type is interface{} with range: 0..4294967295.
    FecId interface{}

    // Id of the policer on the trap. The type is interface{} with range:
    // 0..4294967295.
    PolicerId interface{}

    // Stats Id of the trap. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // Encap Id of the trap. The type is interface{} with range: 0..4294967295.
    EncapId interface{}

    // McGroup of the trap. The type is interface{} with range: 0..4294967295.
    McGroup interface{}

    // Name String of the trap. The type is string.
    TrapString interface{}

    // Id for internal use. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // Offset for internal use. The type is interface{} with range:
    // 0..18446744073709551615.
    Offset interface{}

    // NpuId on which trap is enabled. The type is interface{} with range:
    // 0..18446744073709551615.
    NpuId interface{}

    // Number of packets dropped after hitting the trap. The type is interface{}
    // with range: 0..18446744073709551615.
    PacketDropped interface{}

    // Number of packets accepted after hitting the trap. The type is interface{}
    // with range: 0..18446744073709551615.
    PacketAccepted interface{}
}

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetFilter() yfilter.YFilter { return trapId.YFilter }

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) SetFilter(yf yfilter.YFilter) { trapId.YFilter = yf }

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetGoName(yname string) string {
    if yname == "trap-id" { return "TrapId" }
    if yname == "trap-strength" { return "TrapStrength" }
    if yname == "priority" { return "Priority" }
    if yname == "trap-id-xr" { return "TrapIdXr" }
    if yname == "gport" { return "Gport" }
    if yname == "fec-id" { return "FecId" }
    if yname == "policer-id" { return "PolicerId" }
    if yname == "stats-id" { return "StatsId" }
    if yname == "encap-id" { return "EncapId" }
    if yname == "mc-group" { return "McGroup" }
    if yname == "trap-string" { return "TrapString" }
    if yname == "id" { return "Id" }
    if yname == "offset" { return "Offset" }
    if yname == "npu-id" { return "NpuId" }
    if yname == "packet-dropped" { return "PacketDropped" }
    if yname == "packet-accepted" { return "PacketAccepted" }
    return ""
}

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetSegmentPath() string {
    return "trap-id" + "[trap-id='" + fmt.Sprintf("%v", trapId.TrapId) + "']"
}

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trap-id"] = trapId.TrapId
    leafs["trap-strength"] = trapId.TrapStrength
    leafs["priority"] = trapId.Priority
    leafs["trap-id-xr"] = trapId.TrapIdXr
    leafs["gport"] = trapId.Gport
    leafs["fec-id"] = trapId.FecId
    leafs["policer-id"] = trapId.PolicerId
    leafs["stats-id"] = trapId.StatsId
    leafs["encap-id"] = trapId.EncapId
    leafs["mc-group"] = trapId.McGroup
    leafs["trap-string"] = trapId.TrapString
    leafs["id"] = trapId.Id
    leafs["offset"] = trapId.Offset
    leafs["npu-id"] = trapId.NpuId
    leafs["packet-dropped"] = trapId.PacketDropped
    leafs["packet-accepted"] = trapId.PacketAccepted
    return leafs
}

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetBundleName() string { return "cisco_ios_xr" }

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetYangName() string { return "trap-id" }

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) SetParent(parent types.Entity) { trapId.parent = parent }

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetParent() types.Entity { return trapId.parent }

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetParentYangName() string { return "trap-ids" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles
// Voq stats grouped by interface handle
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Voq stats for a particular interface handle. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle.
    InterfaceHandle []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle
}

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetFilter() yfilter.YFilter { return interfaceHandles.YFilter }

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) SetFilter(yf yfilter.YFilter) { interfaceHandles.YFilter = yf }

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetGoName(yname string) string {
    if yname == "interface-handle" { return "InterfaceHandle" }
    return ""
}

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetSegmentPath() string {
    return "interface-handles"
}

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-handle" {
        for _, c := range interfaceHandles.InterfaceHandle {
            if interfaceHandles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle{}
        interfaceHandles.InterfaceHandle = append(interfaceHandles.InterfaceHandle, child)
        return &interfaceHandles.InterfaceHandle[len(interfaceHandles.InterfaceHandle)-1]
    }
    return nil
}

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceHandles.InterfaceHandle {
        children[interfaceHandles.InterfaceHandle[i].GetSegmentPath()] = &interfaceHandles.InterfaceHandle[i]
    }
    return children
}

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetYangName() string { return "interface-handles" }

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) SetParent(parent types.Entity) { interfaceHandles.parent = parent }

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetParent() types.Entity { return interfaceHandles.parent }

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetParentYangName() string { return "display" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle
// Voq stats for a particular interface
// handle
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Handle. The type is interface{} with
    // range: 0..4294967295.
    InterfaceHandle interface{}

    // Flag to indicate if port is in use. The type is bool.
    InUse interface{}

    // Rack of port. The type is interface{} with range: 0..255.
    RackNum interface{}

    // Slot of port. The type is interface{} with range: 0..255.
    SlotNum interface{}

    // NPU of port. The type is interface{} with range: 0..255.
    NpuNum interface{}

    // NPU core of port. The type is interface{} with range: 0..255.
    NpuCore interface{}

    // Port Number of port. The type is interface{} with range: 0..255.
    PortNum interface{}

    // IfHandle of port. The type is interface{} with range: 0..4294967295.
    IfHandle interface{}

    // System port of port. The type is interface{} with range: 0..4294967295.
    SysPort interface{}

    // PP Port number of port. The type is interface{} with range: 0..4294967295.
    PpPort interface{}

    // Port speed of port. The type is interface{} with range: 0..4294967295.
    PortSpeed interface{}

    // Voq Base number of port. The type is interface{} with range: 0..4294967295.
    VoqBase interface{}

    // Connector id of port. The type is interface{} with range: 0..4294967295.
    ConnectorId interface{}

    // Flag to indicate if port is local to the node. The type is bool.
    IsLocalPort interface{}

    // Keeps a record of the received and dropped packets and bytes on the port.
    // The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat.
    VoqStat []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat
}

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetFilter() yfilter.YFilter { return interfaceHandle.YFilter }

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) SetFilter(yf yfilter.YFilter) { interfaceHandle.YFilter = yf }

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetGoName(yname string) string {
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "in-use" { return "InUse" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "slot-num" { return "SlotNum" }
    if yname == "npu-num" { return "NpuNum" }
    if yname == "npu-core" { return "NpuCore" }
    if yname == "port-num" { return "PortNum" }
    if yname == "if-handle" { return "IfHandle" }
    if yname == "sys-port" { return "SysPort" }
    if yname == "pp-port" { return "PpPort" }
    if yname == "port-speed" { return "PortSpeed" }
    if yname == "voq-base" { return "VoqBase" }
    if yname == "connector-id" { return "ConnectorId" }
    if yname == "is-local-port" { return "IsLocalPort" }
    if yname == "voq-stat" { return "VoqStat" }
    return ""
}

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetSegmentPath() string {
    return "interface-handle" + "[interface-handle='" + fmt.Sprintf("%v", interfaceHandle.InterfaceHandle) + "']"
}

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "voq-stat" {
        for _, c := range interfaceHandle.VoqStat {
            if interfaceHandle.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat{}
        interfaceHandle.VoqStat = append(interfaceHandle.VoqStat, child)
        return &interfaceHandle.VoqStat[len(interfaceHandle.VoqStat)-1]
    }
    return nil
}

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceHandle.VoqStat {
        children[interfaceHandle.VoqStat[i].GetSegmentPath()] = &interfaceHandle.VoqStat[i]
    }
    return children
}

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-handle"] = interfaceHandle.InterfaceHandle
    leafs["in-use"] = interfaceHandle.InUse
    leafs["rack-num"] = interfaceHandle.RackNum
    leafs["slot-num"] = interfaceHandle.SlotNum
    leafs["npu-num"] = interfaceHandle.NpuNum
    leafs["npu-core"] = interfaceHandle.NpuCore
    leafs["port-num"] = interfaceHandle.PortNum
    leafs["if-handle"] = interfaceHandle.IfHandle
    leafs["sys-port"] = interfaceHandle.SysPort
    leafs["pp-port"] = interfaceHandle.PpPort
    leafs["port-speed"] = interfaceHandle.PortSpeed
    leafs["voq-base"] = interfaceHandle.VoqBase
    leafs["connector-id"] = interfaceHandle.ConnectorId
    leafs["is-local-port"] = interfaceHandle.IsLocalPort
    return leafs
}

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetYangName() string { return "interface-handle" }

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) SetParent(parent types.Entity) { interfaceHandle.parent = parent }

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetParent() types.Entity { return interfaceHandle.parent }

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetParentYangName() string { return "interface-handles" }

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat
// Keeps a record of the received and dropped
// packets and bytes on the port
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bytes Received on the port. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ReceivedBytes interface{}

    // Packets Received on the port. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPackets interface{}

    // Bytes Dropped on the port. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    DroppedBytes interface{}

    // Packets Dropeed on the port. The type is interface{} with range:
    // 0..18446744073709551615.
    DroppedPackets interface{}
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetFilter() yfilter.YFilter { return voqStat.YFilter }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) SetFilter(yf yfilter.YFilter) { voqStat.YFilter = yf }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetGoName(yname string) string {
    if yname == "received-bytes" { return "ReceivedBytes" }
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "dropped-bytes" { return "DroppedBytes" }
    if yname == "dropped-packets" { return "DroppedPackets" }
    return ""
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetSegmentPath() string {
    return "voq-stat"
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-bytes"] = voqStat.ReceivedBytes
    leafs["received-packets"] = voqStat.ReceivedPackets
    leafs["dropped-bytes"] = voqStat.DroppedBytes
    leafs["dropped-packets"] = voqStat.DroppedPackets
    return leafs
}

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetBundleName() string { return "cisco_ios_xr" }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetYangName() string { return "voq-stat" }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) SetParent(parent types.Entity) { voqStat.parent = parent }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetParent() types.Entity { return voqStat.parent }

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetParentYangName() string { return "interface-handle" }

// Dpa_Stats_Nodes_Node_NpuIds
// DPA stats hw resources info 
type Dpa_Stats_Nodes_Node_NpuIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats Hardware Info for particular NPU. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuIds_NpuId.
    NpuId []Dpa_Stats_Nodes_Node_NpuIds_NpuId
}

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetFilter() yfilter.YFilter { return npuIds.YFilter }

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) SetFilter(yf yfilter.YFilter) { npuIds.YFilter = yf }

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    return ""
}

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetSegmentPath() string {
    return "npu-ids"
}

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "npu-id" {
        for _, c := range npuIds.NpuId {
            if npuIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_NpuIds_NpuId{}
        npuIds.NpuId = append(npuIds.NpuId, child)
        return &npuIds.NpuId[len(npuIds.NpuId)-1]
    }
    return nil
}

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range npuIds.NpuId {
        children[npuIds.NpuId[i].GetSegmentPath()] = &npuIds.NpuId[i]
    }
    return children
}

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetBundleName() string { return "cisco_ios_xr" }

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetYangName() string { return "npu-ids" }

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) SetParent(parent types.Entity) { npuIds.parent = parent }

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetParent() types.Entity { return npuIds.parent }

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetParentYangName() string { return "node" }

// Dpa_Stats_Nodes_Node_NpuIds_NpuId
// Stats Hardware Info for particular NPU
type Dpa_Stats_Nodes_Node_NpuIds_NpuId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. NPU number. The type is interface{} with range:
    // -2147483648..2147483647.
    NpuId interface{}

    // sys cp cnfg prof. The type is interface{} with range: 0..4294967295.
    SysCpCnfgProf interface{}

    // next avail cp id. The type is interface{} with range: 0..4294967295.
    NextAvailCpId interface{}

    // num cntr engines. The type is interface{} with range: 0..4294967295.
    NumCntrEngines interface{}

    // cntr engine. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine.
    CntrEngine []Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine
}

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetFilter() yfilter.YFilter { return npuId.YFilter }

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) SetFilter(yf yfilter.YFilter) { npuId.YFilter = yf }

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "sys-cp-cnfg-prof" { return "SysCpCnfgProf" }
    if yname == "next-avail-cp-id" { return "NextAvailCpId" }
    if yname == "num-cntr-engines" { return "NumCntrEngines" }
    if yname == "cntr-engine" { return "CntrEngine" }
    return ""
}

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetSegmentPath() string {
    return "npu-id" + "[npu-id='" + fmt.Sprintf("%v", npuId.NpuId) + "']"
}

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cntr-engine" {
        for _, c := range npuId.CntrEngine {
            if npuId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine{}
        npuId.CntrEngine = append(npuId.CntrEngine, child)
        return &npuId.CntrEngine[len(npuId.CntrEngine)-1]
    }
    return nil
}

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range npuId.CntrEngine {
        children[npuId.CntrEngine[i].GetSegmentPath()] = &npuId.CntrEngine[i]
    }
    return children
}

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = npuId.NpuId
    leafs["sys-cp-cnfg-prof"] = npuId.SysCpCnfgProf
    leafs["next-avail-cp-id"] = npuId.NextAvailCpId
    leafs["num-cntr-engines"] = npuId.NumCntrEngines
    return leafs
}

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetBundleName() string { return "cisco_ios_xr" }

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetYangName() string { return "npu-id" }

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) SetParent(parent types.Entity) { npuId.parent = parent }

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetParent() types.Entity { return npuId.parent }

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetParentYangName() string { return "npu-ids" }

// Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine
// cntr engine
type Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // state. The type is string.
    State interface{}

    // core id. The type is interface{} with range: 0..4294967295.
    CoreId interface{}

    // apps info. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo.
    AppsInfo []Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo
}

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetFilter() yfilter.YFilter { return cntrEngine.YFilter }

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) SetFilter(yf yfilter.YFilter) { cntrEngine.YFilter = yf }

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    if yname == "core-id" { return "CoreId" }
    if yname == "apps-info" { return "AppsInfo" }
    return ""
}

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetSegmentPath() string {
    return "cntr-engine"
}

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "apps-info" {
        for _, c := range cntrEngine.AppsInfo {
            if cntrEngine.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo{}
        cntrEngine.AppsInfo = append(cntrEngine.AppsInfo, child)
        return &cntrEngine.AppsInfo[len(cntrEngine.AppsInfo)-1]
    }
    return nil
}

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cntrEngine.AppsInfo {
        children[cntrEngine.AppsInfo[i].GetSegmentPath()] = &cntrEngine.AppsInfo[i]
    }
    return children
}

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state"] = cntrEngine.State
    leafs["core-id"] = cntrEngine.CoreId
    return leafs
}

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetBundleName() string { return "cisco_ios_xr" }

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetYangName() string { return "cntr-engine" }

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) SetParent(parent types.Entity) { cntrEngine.parent = parent }

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetParent() types.Entity { return cntrEngine.parent }

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetParentYangName() string { return "npu-id" }

// Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo
// apps info
type Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // app type. The type is string.
    AppType interface{}

    // num cntrs for app. The type is interface{} with range: 0..4294967295.
    NumCntrsForApp interface{}

    // num cntrs used. The type is interface{} with range: 0..4294967295.
    NumCntrsUsed interface{}
}

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetFilter() yfilter.YFilter { return appsInfo.YFilter }

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) SetFilter(yf yfilter.YFilter) { appsInfo.YFilter = yf }

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetGoName(yname string) string {
    if yname == "app-type" { return "AppType" }
    if yname == "num-cntrs-for-app" { return "NumCntrsForApp" }
    if yname == "num-cntrs-used" { return "NumCntrsUsed" }
    return ""
}

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetSegmentPath() string {
    return "apps-info"
}

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["app-type"] = appsInfo.AppType
    leafs["num-cntrs-for-app"] = appsInfo.NumCntrsForApp
    leafs["num-cntrs-used"] = appsInfo.NumCntrsUsed
    return leafs
}

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetYangName() string { return "apps-info" }

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) SetParent(parent types.Entity) { appsInfo.parent = parent }

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetParent() types.Entity { return appsInfo.parent }

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetParentYangName() string { return "cntr-engine" }

