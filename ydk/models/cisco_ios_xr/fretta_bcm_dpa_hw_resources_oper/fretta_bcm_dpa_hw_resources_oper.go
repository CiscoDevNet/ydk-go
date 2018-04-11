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

    // ext tcam ipv6
    Resource_ext_tcam_ipv6 Resource = "ext-tcam-ipv6"
)

// Dpa
// Stats Data
type Dpa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Voq or Trap Data.
    Stats Dpa_Stats
}

func (dpa *Dpa) GetEntityData() *types.CommonEntityData {
    dpa.EntityData.YFilter = dpa.YFilter
    dpa.EntityData.YangName = "dpa"
    dpa.EntityData.BundleName = "cisco_ios_xr"
    dpa.EntityData.ParentYangName = "Cisco-IOS-XR-fretta-bcm-dpa-hw-resources-oper"
    dpa.EntityData.SegmentPath = "Cisco-IOS-XR-fretta-bcm-dpa-hw-resources-oper:dpa"
    dpa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dpa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dpa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dpa.EntityData.Children = make(map[string]types.YChild)
    dpa.EntityData.Children["stats"] = types.YChild{"Stats", &dpa.Stats}
    dpa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dpa.EntityData)
}

// Dpa_Stats
// Voq or Trap Data
type Dpa_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA data for available nodes.
    Nodes Dpa_Stats_Nodes
}

func (stats *Dpa_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "dpa"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Children["nodes"] = types.YChild{"Nodes", &stats.Nodes}
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stats.EntityData)
}

// Dpa_Stats_Nodes
// DPA data for available nodes
type Dpa_Stats_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DPA operational data for a particular node. The type is slice of
    // Dpa_Stats_Nodes_Node.
    Node []Dpa_Stats_Nodes_Node
}

func (nodes *Dpa_Stats_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "stats"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Dpa_Stats_Nodes_Node
// DPA operational data for a particular node
type Dpa_Stats_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (node *Dpa_Stats_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["hw-resources-datas"] = types.YChild{"HwResourcesDatas", &node.HwResourcesDatas}
    node.EntityData.Children["asic-statistics"] = types.YChild{"AsicStatistics", &node.AsicStatistics}
    node.EntityData.Children["npu-numbers"] = types.YChild{"NpuNumbers", &node.NpuNumbers}
    node.EntityData.Children["npu-ids"] = types.YChild{"NpuIds", &node.NpuIds}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Dpa_Stats_Nodes_Node_HwResourcesDatas
// DPA hw resources stats 
type Dpa_Stats_Nodes_Node_HwResourcesDatas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hardware resources table. The type is slice of
    // Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData.
    HwResourcesData []Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData
}

func (hwResourcesDatas *Dpa_Stats_Nodes_Node_HwResourcesDatas) GetEntityData() *types.CommonEntityData {
    hwResourcesDatas.EntityData.YFilter = hwResourcesDatas.YFilter
    hwResourcesDatas.EntityData.YangName = "hw-resources-datas"
    hwResourcesDatas.EntityData.BundleName = "cisco_ios_xr"
    hwResourcesDatas.EntityData.ParentYangName = "node"
    hwResourcesDatas.EntityData.SegmentPath = "hw-resources-datas"
    hwResourcesDatas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwResourcesDatas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwResourcesDatas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwResourcesDatas.EntityData.Children = make(map[string]types.YChild)
    hwResourcesDatas.EntityData.Children["hw-resources-data"] = types.YChild{"HwResourcesData", nil}
    for i := range hwResourcesDatas.HwResourcesData {
        hwResourcesDatas.EntityData.Children[types.GetSegmentPath(&hwResourcesDatas.HwResourcesData[i])] = types.YChild{"HwResourcesData", &hwResourcesDatas.HwResourcesData[i]}
    }
    hwResourcesDatas.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hwResourcesDatas.EntityData)
}

// Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData
// Hardware resources table
type Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Resource type. The type is Resource.
    Resource interface{}

    // resource id. The type is interface{} with range: 0..4294967295.
    ResourceId interface{}

    // name. The type is string.
    Name interface{}

    // num npus. The type is interface{} with range: 0..4294967295.
    NumNpus interface{}

    // cmd invalid. The type is bool.
    CmdInvalid interface{}

    // npu hwr. The type is slice of
    // Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr.
    NpuHwr []Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr
}

func (hwResourcesData *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData) GetEntityData() *types.CommonEntityData {
    hwResourcesData.EntityData.YFilter = hwResourcesData.YFilter
    hwResourcesData.EntityData.YangName = "hw-resources-data"
    hwResourcesData.EntityData.BundleName = "cisco_ios_xr"
    hwResourcesData.EntityData.ParentYangName = "hw-resources-datas"
    hwResourcesData.EntityData.SegmentPath = "hw-resources-data" + "[resource='" + fmt.Sprintf("%v", hwResourcesData.Resource) + "']"
    hwResourcesData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwResourcesData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwResourcesData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwResourcesData.EntityData.Children = make(map[string]types.YChild)
    hwResourcesData.EntityData.Children["npu-hwr"] = types.YChild{"NpuHwr", nil}
    for i := range hwResourcesData.NpuHwr {
        hwResourcesData.EntityData.Children[types.GetSegmentPath(&hwResourcesData.NpuHwr[i])] = types.YChild{"NpuHwr", &hwResourcesData.NpuHwr[i]}
    }
    hwResourcesData.EntityData.Leafs = make(map[string]types.YLeaf)
    hwResourcesData.EntityData.Leafs["resource"] = types.YLeaf{"Resource", hwResourcesData.Resource}
    hwResourcesData.EntityData.Leafs["resource-id"] = types.YLeaf{"ResourceId", hwResourcesData.ResourceId}
    hwResourcesData.EntityData.Leafs["name"] = types.YLeaf{"Name", hwResourcesData.Name}
    hwResourcesData.EntityData.Leafs["num-npus"] = types.YLeaf{"NumNpus", hwResourcesData.NumNpus}
    hwResourcesData.EntityData.Leafs["cmd-invalid"] = types.YLeaf{"CmdInvalid", hwResourcesData.CmdInvalid}
    return &(hwResourcesData.EntityData)
}

// Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr
// npu hwr
type Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr struct {
    EntityData types.CommonEntityData
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

func (npuHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr) GetEntityData() *types.CommonEntityData {
    npuHwr.EntityData.YFilter = npuHwr.YFilter
    npuHwr.EntityData.YangName = "npu-hwr"
    npuHwr.EntityData.BundleName = "cisco_ios_xr"
    npuHwr.EntityData.ParentYangName = "hw-resources-data"
    npuHwr.EntityData.SegmentPath = "npu-hwr"
    npuHwr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuHwr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuHwr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuHwr.EntityData.Children = make(map[string]types.YChild)
    npuHwr.EntityData.Children["lt-hwr"] = types.YChild{"LtHwr", nil}
    for i := range npuHwr.LtHwr {
        npuHwr.EntityData.Children[types.GetSegmentPath(&npuHwr.LtHwr[i])] = types.YChild{"LtHwr", &npuHwr.LtHwr[i]}
    }
    npuHwr.EntityData.Leafs = make(map[string]types.YLeaf)
    npuHwr.EntityData.Leafs["max-allowed"] = types.YLeaf{"MaxAllowed", npuHwr.MaxAllowed}
    npuHwr.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", npuHwr.NpuId}
    npuHwr.EntityData.Leafs["max-entries"] = types.YLeaf{"MaxEntries", npuHwr.MaxEntries}
    npuHwr.EntityData.Leafs["red-oor-threshold"] = types.YLeaf{"RedOorThreshold", npuHwr.RedOorThreshold}
    npuHwr.EntityData.Leafs["red-oor-threshold-percent"] = types.YLeaf{"RedOorThresholdPercent", npuHwr.RedOorThresholdPercent}
    npuHwr.EntityData.Leafs["yellow-oor-threshold"] = types.YLeaf{"YellowOorThreshold", npuHwr.YellowOorThreshold}
    npuHwr.EntityData.Leafs["yellow-oor-threshold-percent"] = types.YLeaf{"YellowOorThresholdPercent", npuHwr.YellowOorThresholdPercent}
    npuHwr.EntityData.Leafs["inuse-objects"] = types.YLeaf{"InuseObjects", npuHwr.InuseObjects}
    npuHwr.EntityData.Leafs["num-lt"] = types.YLeaf{"NumLt", npuHwr.NumLt}
    npuHwr.EntityData.Leafs["oor-change-count"] = types.YLeaf{"OorChangeCount", npuHwr.OorChangeCount}
    npuHwr.EntityData.Leafs["oor-state-change-time1"] = types.YLeaf{"OorStateChangeTime1", npuHwr.OorStateChangeTime1}
    npuHwr.EntityData.Leafs["oor-state-change-time2"] = types.YLeaf{"OorStateChangeTime2", npuHwr.OorStateChangeTime2}
    npuHwr.EntityData.Leafs["oor-state"] = types.YLeaf{"OorState", npuHwr.OorState}
    return &(npuHwr.EntityData)
}

// Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr
// lt hwr
type Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr struct {
    EntityData types.CommonEntityData
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

func (ltHwr *Dpa_Stats_Nodes_Node_HwResourcesDatas_HwResourcesData_NpuHwr_LtHwr) GetEntityData() *types.CommonEntityData {
    ltHwr.EntityData.YFilter = ltHwr.YFilter
    ltHwr.EntityData.YangName = "lt-hwr"
    ltHwr.EntityData.BundleName = "cisco_ios_xr"
    ltHwr.EntityData.ParentYangName = "npu-hwr"
    ltHwr.EntityData.SegmentPath = "lt-hwr"
    ltHwr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ltHwr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ltHwr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ltHwr.EntityData.Children = make(map[string]types.YChild)
    ltHwr.EntityData.Leafs = make(map[string]types.YLeaf)
    ltHwr.EntityData.Leafs["lt-id"] = types.YLeaf{"LtId", ltHwr.LtId}
    ltHwr.EntityData.Leafs["name"] = types.YLeaf{"Name", ltHwr.Name}
    ltHwr.EntityData.Leafs["hw-entries"] = types.YLeaf{"HwEntries", ltHwr.HwEntries}
    ltHwr.EntityData.Leafs["sw-entries"] = types.YLeaf{"SwEntries", ltHwr.SwEntries}
    return &(ltHwr.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics
// ASIC statistics table
type Dpa_Stats_Nodes_Node_AsicStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ASIC statistics.
    AsicStatisticsForNpuIds Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds

    // Detailed ASIC statistics.
    AsicStatisticsDetailForNpuIds Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds
}

func (asicStatistics *Dpa_Stats_Nodes_Node_AsicStatistics) GetEntityData() *types.CommonEntityData {
    asicStatistics.EntityData.YFilter = asicStatistics.YFilter
    asicStatistics.EntityData.YangName = "asic-statistics"
    asicStatistics.EntityData.BundleName = "cisco_ios_xr"
    asicStatistics.EntityData.ParentYangName = "node"
    asicStatistics.EntityData.SegmentPath = "asic-statistics"
    asicStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatistics.EntityData.Children = make(map[string]types.YChild)
    asicStatistics.EntityData.Children["asic-statistics-for-npu-ids"] = types.YChild{"AsicStatisticsForNpuIds", &asicStatistics.AsicStatisticsForNpuIds}
    asicStatistics.EntityData.Children["asic-statistics-detail-for-npu-ids"] = types.YChild{"AsicStatisticsDetailForNpuIds", &asicStatistics.AsicStatisticsDetailForNpuIds}
    asicStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicStatistics.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds
// ASIC statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ASIC statistics for a particular NPU. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId.
    AsicStatisticsForNpuId []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId
}

func (asicStatisticsForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds) GetEntityData() *types.CommonEntityData {
    asicStatisticsForNpuIds.EntityData.YFilter = asicStatisticsForNpuIds.YFilter
    asicStatisticsForNpuIds.EntityData.YangName = "asic-statistics-for-npu-ids"
    asicStatisticsForNpuIds.EntityData.BundleName = "cisco_ios_xr"
    asicStatisticsForNpuIds.EntityData.ParentYangName = "asic-statistics"
    asicStatisticsForNpuIds.EntityData.SegmentPath = "asic-statistics-for-npu-ids"
    asicStatisticsForNpuIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatisticsForNpuIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatisticsForNpuIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatisticsForNpuIds.EntityData.Children = make(map[string]types.YChild)
    asicStatisticsForNpuIds.EntityData.Children["asic-statistics-for-npu-id"] = types.YChild{"AsicStatisticsForNpuId", nil}
    for i := range asicStatisticsForNpuIds.AsicStatisticsForNpuId {
        asicStatisticsForNpuIds.EntityData.Children[types.GetSegmentPath(&asicStatisticsForNpuIds.AsicStatisticsForNpuId[i])] = types.YChild{"AsicStatisticsForNpuId", &asicStatisticsForNpuIds.AsicStatisticsForNpuId[i]}
    }
    asicStatisticsForNpuIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicStatisticsForNpuIds.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId
// ASIC statistics for a particular NPU
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId struct {
    EntityData types.CommonEntityData
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

func (asicStatisticsForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId) GetEntityData() *types.CommonEntityData {
    asicStatisticsForNpuId.EntityData.YFilter = asicStatisticsForNpuId.YFilter
    asicStatisticsForNpuId.EntityData.YangName = "asic-statistics-for-npu-id"
    asicStatisticsForNpuId.EntityData.BundleName = "cisco_ios_xr"
    asicStatisticsForNpuId.EntityData.ParentYangName = "asic-statistics-for-npu-ids"
    asicStatisticsForNpuId.EntityData.SegmentPath = "asic-statistics-for-npu-id" + "[npu-id='" + fmt.Sprintf("%v", asicStatisticsForNpuId.NpuId) + "']"
    asicStatisticsForNpuId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatisticsForNpuId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatisticsForNpuId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatisticsForNpuId.EntityData.Children = make(map[string]types.YChild)
    asicStatisticsForNpuId.EntityData.Children["statistics"] = types.YChild{"Statistics", &asicStatisticsForNpuId.Statistics}
    asicStatisticsForNpuId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicStatisticsForNpuId.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", asicStatisticsForNpuId.NpuId}
    asicStatisticsForNpuId.EntityData.Leafs["valid"] = types.YLeaf{"Valid", asicStatisticsForNpuId.Valid}
    asicStatisticsForNpuId.EntityData.Leafs["rack-number"] = types.YLeaf{"RackNumber", asicStatisticsForNpuId.RackNumber}
    asicStatisticsForNpuId.EntityData.Leafs["slot-number"] = types.YLeaf{"SlotNumber", asicStatisticsForNpuId.SlotNumber}
    asicStatisticsForNpuId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicStatisticsForNpuId.AsicInstance}
    asicStatisticsForNpuId.EntityData.Leafs["chip-version"] = types.YLeaf{"ChipVersion", asicStatisticsForNpuId.ChipVersion}
    return &(asicStatisticsForNpuId.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics
// Statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsForNpuIds_AsicStatisticsForNpuId_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "asic-statistics-for-npu-id"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["nbi-rx-total-byte-cnt"] = types.YLeaf{"NbiRxTotalByteCnt", statistics.NbiRxTotalByteCnt}
    statistics.EntityData.Leafs["nbi-rx-total-pkt-cnt"] = types.YLeaf{"NbiRxTotalPktCnt", statistics.NbiRxTotalPktCnt}
    statistics.EntityData.Leafs["ire-cpu-pkt-cnt"] = types.YLeaf{"IreCpuPktCnt", statistics.IreCpuPktCnt}
    statistics.EntityData.Leafs["ire-nif-pkt-cnt"] = types.YLeaf{"IreNifPktCnt", statistics.IreNifPktCnt}
    statistics.EntityData.Leafs["ire-oamp-pkt-cnt"] = types.YLeaf{"IreOampPktCnt", statistics.IreOampPktCnt}
    statistics.EntityData.Leafs["ire-olp-pkt-cnt"] = types.YLeaf{"IreOlpPktCnt", statistics.IreOlpPktCnt}
    statistics.EntityData.Leafs["ire-rcy-pkt-cnt"] = types.YLeaf{"IreRcyPktCnt", statistics.IreRcyPktCnt}
    statistics.EntityData.Leafs["ire-fdt-if-cnt"] = types.YLeaf{"IreFdtIfCnt", statistics.IreFdtIfCnt}
    statistics.EntityData.Leafs["idr-mmu-if-cnt"] = types.YLeaf{"IdrMmuIfCnt", statistics.IdrMmuIfCnt}
    statistics.EntityData.Leafs["idr-ocb-if-cnt"] = types.YLeaf{"IdrOcbIfCnt", statistics.IdrOcbIfCnt}
    statistics.EntityData.Leafs["iqm-enqueue-pkt-cnt"] = types.YLeaf{"IqmEnqueuePktCnt", statistics.IqmEnqueuePktCnt}
    statistics.EntityData.Leafs["iqm-dequeue-pkt-cnt"] = types.YLeaf{"IqmDequeuePktCnt", statistics.IqmDequeuePktCnt}
    statistics.EntityData.Leafs["iqm-deleted-pkt-cnt"] = types.YLeaf{"IqmDeletedPktCnt", statistics.IqmDeletedPktCnt}
    statistics.EntityData.Leafs["iqm-enq-discarded-pkt-cnt"] = types.YLeaf{"IqmEnqDiscardedPktCnt", statistics.IqmEnqDiscardedPktCnt}
    statistics.EntityData.Leafs["ipt-egq-pkt-cnt"] = types.YLeaf{"IptEgqPktCnt", statistics.IptEgqPktCnt}
    statistics.EntityData.Leafs["ipt-enq-pkt-cnt"] = types.YLeaf{"IptEnqPktCnt", statistics.IptEnqPktCnt}
    statistics.EntityData.Leafs["ipt-fdt-pkt-cnt"] = types.YLeaf{"IptFdtPktCnt", statistics.IptFdtPktCnt}
    statistics.EntityData.Leafs["ipt-cfg-event-cnt"] = types.YLeaf{"IptCfgEventCnt", statistics.IptCfgEventCnt}
    statistics.EntityData.Leafs["ipt-cfg-byte-cnt"] = types.YLeaf{"IptCfgByteCnt", statistics.IptCfgByteCnt}
    statistics.EntityData.Leafs["fdt-ipt-desc-cell-cnt"] = types.YLeaf{"FdtIptDescCellCnt", statistics.FdtIptDescCellCnt}
    statistics.EntityData.Leafs["fdt-ire-desc-cell-cnt"] = types.YLeaf{"FdtIreDescCellCnt", statistics.FdtIreDescCellCnt}
    statistics.EntityData.Leafs["fdt-transmitted-data-cells-cnt"] = types.YLeaf{"FdtTransmittedDataCellsCnt", statistics.FdtTransmittedDataCellsCnt}
    statistics.EntityData.Leafs["fdr-p1-cell-in-cnt"] = types.YLeaf{"FdrP1CellInCnt", statistics.FdrP1CellInCnt}
    statistics.EntityData.Leafs["fdr-p2-cell-in-cnt"] = types.YLeaf{"FdrP2CellInCnt", statistics.FdrP2CellInCnt}
    statistics.EntityData.Leafs["fdr-p3-cell-in-cnt"] = types.YLeaf{"FdrP3CellInCnt", statistics.FdrP3CellInCnt}
    statistics.EntityData.Leafs["fdr-cell-in-cnt-total"] = types.YLeaf{"FdrCellInCntTotal", statistics.FdrCellInCntTotal}
    statistics.EntityData.Leafs["fda-cells-in-cnt-p1"] = types.YLeaf{"FdaCellsInCntP1", statistics.FdaCellsInCntP1}
    statistics.EntityData.Leafs["fda-cells-in-cnt-p2"] = types.YLeaf{"FdaCellsInCntP2", statistics.FdaCellsInCntP2}
    statistics.EntityData.Leafs["fda-cells-in-cnt-p3"] = types.YLeaf{"FdaCellsInCntP3", statistics.FdaCellsInCntP3}
    statistics.EntityData.Leafs["fda-cells-in-tdm-cnt"] = types.YLeaf{"FdaCellsInTdmCnt", statistics.FdaCellsInTdmCnt}
    statistics.EntityData.Leafs["fda-cells-in-meshmc-cnt"] = types.YLeaf{"FdaCellsInMeshmcCnt", statistics.FdaCellsInMeshmcCnt}
    statistics.EntityData.Leafs["fda-cells-in-ipt-cnt"] = types.YLeaf{"FdaCellsInIptCnt", statistics.FdaCellsInIptCnt}
    statistics.EntityData.Leafs["fda-cells-out-cnt-p1"] = types.YLeaf{"FdaCellsOutCntP1", statistics.FdaCellsOutCntP1}
    statistics.EntityData.Leafs["fda-cells-out-cnt-p2"] = types.YLeaf{"FdaCellsOutCntP2", statistics.FdaCellsOutCntP2}
    statistics.EntityData.Leafs["fda-cells-out-cnt-p3"] = types.YLeaf{"FdaCellsOutCntP3", statistics.FdaCellsOutCntP3}
    statistics.EntityData.Leafs["fda-cells-out-tdm-cnt"] = types.YLeaf{"FdaCellsOutTdmCnt", statistics.FdaCellsOutTdmCnt}
    statistics.EntityData.Leafs["fda-cells-out-meshmc-cnt"] = types.YLeaf{"FdaCellsOutMeshmcCnt", statistics.FdaCellsOutMeshmcCnt}
    statistics.EntityData.Leafs["fda-cells-out-ipt-cnt"] = types.YLeaf{"FdaCellsOutIptCnt", statistics.FdaCellsOutIptCnt}
    statistics.EntityData.Leafs["fda-egq-drop-cnt"] = types.YLeaf{"FdaEgqDropCnt", statistics.FdaEgqDropCnt}
    statistics.EntityData.Leafs["fda-egq-meshmc-drop-cnt"] = types.YLeaf{"FdaEgqMeshmcDropCnt", statistics.FdaEgqMeshmcDropCnt}
    statistics.EntityData.Leafs["egq-fqp-pkt-cnt"] = types.YLeaf{"EgqFqpPktCnt", statistics.EgqFqpPktCnt}
    statistics.EntityData.Leafs["egq-pqp-uc-pkt-cnt"] = types.YLeaf{"EgqPqpUcPktCnt", statistics.EgqPqpUcPktCnt}
    statistics.EntityData.Leafs["egq-pqp-discard-uc-pkt-cnt"] = types.YLeaf{"EgqPqpDiscardUcPktCnt", statistics.EgqPqpDiscardUcPktCnt}
    statistics.EntityData.Leafs["egq-pqp-uc-bytes-cnt"] = types.YLeaf{"EgqPqpUcBytesCnt", statistics.EgqPqpUcBytesCnt}
    statistics.EntityData.Leafs["egq-pqp-mc-pkt-cnt"] = types.YLeaf{"EgqPqpMcPktCnt", statistics.EgqPqpMcPktCnt}
    statistics.EntityData.Leafs["egq-pqp-discard-mc-pkt-cnt"] = types.YLeaf{"EgqPqpDiscardMcPktCnt", statistics.EgqPqpDiscardMcPktCnt}
    statistics.EntityData.Leafs["egq-pqp-mc-bytes-cnt"] = types.YLeaf{"EgqPqpMcBytesCnt", statistics.EgqPqpMcBytesCnt}
    statistics.EntityData.Leafs["egq-ehp-uc-pkt-cnt"] = types.YLeaf{"EgqEhpUcPktCnt", statistics.EgqEhpUcPktCnt}
    statistics.EntityData.Leafs["egq-ehp-mc-high-pkt-cnt"] = types.YLeaf{"EgqEhpMcHighPktCnt", statistics.EgqEhpMcHighPktCnt}
    statistics.EntityData.Leafs["egq-ehp-mc-low-pkt-cnt"] = types.YLeaf{"EgqEhpMcLowPktCnt", statistics.EgqEhpMcLowPktCnt}
    statistics.EntityData.Leafs["egq-deleted-pkt-cnt"] = types.YLeaf{"EgqDeletedPktCnt", statistics.EgqDeletedPktCnt}
    statistics.EntityData.Leafs["egq-ehp-mc-high-discard-cnt"] = types.YLeaf{"EgqEhpMcHighDiscardCnt", statistics.EgqEhpMcHighDiscardCnt}
    statistics.EntityData.Leafs["egq-ehp-mc-low-discard-cnt"] = types.YLeaf{"EgqEhpMcLowDiscardCnt", statistics.EgqEhpMcLowDiscardCnt}
    statistics.EntityData.Leafs["egq-erpp-lag-pruning-discard-cnt"] = types.YLeaf{"EgqErppLagPruningDiscardCnt", statistics.EgqErppLagPruningDiscardCnt}
    statistics.EntityData.Leafs["egq-erpp-pmf-discard-cnt"] = types.YLeaf{"EgqErppPmfDiscardCnt", statistics.EgqErppPmfDiscardCnt}
    statistics.EntityData.Leafs["egq-erpp-vlan-mbr-discard-cnt"] = types.YLeaf{"EgqErppVlanMbrDiscardCnt", statistics.EgqErppVlanMbrDiscardCnt}
    statistics.EntityData.Leafs["epni-epe-byte-cnt"] = types.YLeaf{"EpniEpeByteCnt", statistics.EpniEpeByteCnt}
    statistics.EntityData.Leafs["epni-epe-pkt-cnt"] = types.YLeaf{"EpniEpePktCnt", statistics.EpniEpePktCnt}
    statistics.EntityData.Leafs["epni-epe-discard-cnt"] = types.YLeaf{"EpniEpeDiscardCnt", statistics.EpniEpeDiscardCnt}
    statistics.EntityData.Leafs["nbi-tx-total-byte-cnt"] = types.YLeaf{"NbiTxTotalByteCnt", statistics.NbiTxTotalByteCnt}
    statistics.EntityData.Leafs["nbi-tx-total-pkt-cnt"] = types.YLeaf{"NbiTxTotalPktCnt", statistics.NbiTxTotalPktCnt}
    return &(statistics.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds
// Detailed ASIC statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed ASIC statistics for a particular NPU. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId.
    AsicStatisticsDetailForNpuId []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId
}

func (asicStatisticsDetailForNpuIds *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds) GetEntityData() *types.CommonEntityData {
    asicStatisticsDetailForNpuIds.EntityData.YFilter = asicStatisticsDetailForNpuIds.YFilter
    asicStatisticsDetailForNpuIds.EntityData.YangName = "asic-statistics-detail-for-npu-ids"
    asicStatisticsDetailForNpuIds.EntityData.BundleName = "cisco_ios_xr"
    asicStatisticsDetailForNpuIds.EntityData.ParentYangName = "asic-statistics"
    asicStatisticsDetailForNpuIds.EntityData.SegmentPath = "asic-statistics-detail-for-npu-ids"
    asicStatisticsDetailForNpuIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatisticsDetailForNpuIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatisticsDetailForNpuIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatisticsDetailForNpuIds.EntityData.Children = make(map[string]types.YChild)
    asicStatisticsDetailForNpuIds.EntityData.Children["asic-statistics-detail-for-npu-id"] = types.YChild{"AsicStatisticsDetailForNpuId", nil}
    for i := range asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId {
        asicStatisticsDetailForNpuIds.EntityData.Children[types.GetSegmentPath(&asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId[i])] = types.YChild{"AsicStatisticsDetailForNpuId", &asicStatisticsDetailForNpuIds.AsicStatisticsDetailForNpuId[i]}
    }
    asicStatisticsDetailForNpuIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(asicStatisticsDetailForNpuIds.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId
// Detailed ASIC statistics for a particular
// NPU
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId struct {
    EntityData types.CommonEntityData
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

func (asicStatisticsDetailForNpuId *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId) GetEntityData() *types.CommonEntityData {
    asicStatisticsDetailForNpuId.EntityData.YFilter = asicStatisticsDetailForNpuId.YFilter
    asicStatisticsDetailForNpuId.EntityData.YangName = "asic-statistics-detail-for-npu-id"
    asicStatisticsDetailForNpuId.EntityData.BundleName = "cisco_ios_xr"
    asicStatisticsDetailForNpuId.EntityData.ParentYangName = "asic-statistics-detail-for-npu-ids"
    asicStatisticsDetailForNpuId.EntityData.SegmentPath = "asic-statistics-detail-for-npu-id" + "[npu-id='" + fmt.Sprintf("%v", asicStatisticsDetailForNpuId.NpuId) + "']"
    asicStatisticsDetailForNpuId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatisticsDetailForNpuId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatisticsDetailForNpuId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatisticsDetailForNpuId.EntityData.Children = make(map[string]types.YChild)
    asicStatisticsDetailForNpuId.EntityData.Children["statistics"] = types.YChild{"Statistics", &asicStatisticsDetailForNpuId.Statistics}
    asicStatisticsDetailForNpuId.EntityData.Leafs = make(map[string]types.YLeaf)
    asicStatisticsDetailForNpuId.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", asicStatisticsDetailForNpuId.NpuId}
    asicStatisticsDetailForNpuId.EntityData.Leafs["valid"] = types.YLeaf{"Valid", asicStatisticsDetailForNpuId.Valid}
    asicStatisticsDetailForNpuId.EntityData.Leafs["rack-number"] = types.YLeaf{"RackNumber", asicStatisticsDetailForNpuId.RackNumber}
    asicStatisticsDetailForNpuId.EntityData.Leafs["slot-number"] = types.YLeaf{"SlotNumber", asicStatisticsDetailForNpuId.SlotNumber}
    asicStatisticsDetailForNpuId.EntityData.Leafs["asic-instance"] = types.YLeaf{"AsicInstance", asicStatisticsDetailForNpuId.AsicInstance}
    asicStatisticsDetailForNpuId.EntityData.Leafs["chip-version"] = types.YLeaf{"ChipVersion", asicStatisticsDetailForNpuId.ChipVersion}
    return &(asicStatisticsDetailForNpuId.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics
// Statistics
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of blocks. The type is interface{} with range: 0..255.
    NumBlocks interface{}

    // Block information. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo.
    BlockInfo []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo
}

func (statistics *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "asic-statistics-detail-for-npu-id"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["block-info"] = types.YChild{"BlockInfo", nil}
    for i := range statistics.BlockInfo {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.BlockInfo[i])] = types.YChild{"BlockInfo", &statistics.BlockInfo[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["num-blocks"] = types.YLeaf{"NumBlocks", statistics.NumBlocks}
    return &(statistics.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo
// Block information
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Block name. The type is string with length: 0..10.
    BlockName interface{}

    // Number of fields. The type is interface{} with range: 0..255.
    NumFields interface{}

    // Field information. The type is slice of
    // Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo.
    FieldInfo []Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo
}

func (blockInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo) GetEntityData() *types.CommonEntityData {
    blockInfo.EntityData.YFilter = blockInfo.YFilter
    blockInfo.EntityData.YangName = "block-info"
    blockInfo.EntityData.BundleName = "cisco_ios_xr"
    blockInfo.EntityData.ParentYangName = "statistics"
    blockInfo.EntityData.SegmentPath = "block-info"
    blockInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    blockInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    blockInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    blockInfo.EntityData.Children = make(map[string]types.YChild)
    blockInfo.EntityData.Children["field-info"] = types.YChild{"FieldInfo", nil}
    for i := range blockInfo.FieldInfo {
        blockInfo.EntityData.Children[types.GetSegmentPath(&blockInfo.FieldInfo[i])] = types.YChild{"FieldInfo", &blockInfo.FieldInfo[i]}
    }
    blockInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    blockInfo.EntityData.Leafs["block-name"] = types.YLeaf{"BlockName", blockInfo.BlockName}
    blockInfo.EntityData.Leafs["num-fields"] = types.YLeaf{"NumFields", blockInfo.NumFields}
    return &(blockInfo.EntityData)
}

// Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo
// Field information
type Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Field name. The type is string with length: 0..80.
    FieldName interface{}

    // Field value. The type is interface{} with range: 0..18446744073709551615.
    FieldValue interface{}

    // Flag to indicate overflow. The type is bool.
    IsOverflow interface{}
}

func (fieldInfo *Dpa_Stats_Nodes_Node_AsicStatistics_AsicStatisticsDetailForNpuIds_AsicStatisticsDetailForNpuId_Statistics_BlockInfo_FieldInfo) GetEntityData() *types.CommonEntityData {
    fieldInfo.EntityData.YFilter = fieldInfo.YFilter
    fieldInfo.EntityData.YangName = "field-info"
    fieldInfo.EntityData.BundleName = "cisco_ios_xr"
    fieldInfo.EntityData.ParentYangName = "block-info"
    fieldInfo.EntityData.SegmentPath = "field-info"
    fieldInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fieldInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fieldInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fieldInfo.EntityData.Children = make(map[string]types.YChild)
    fieldInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    fieldInfo.EntityData.Leafs["field-name"] = types.YLeaf{"FieldName", fieldInfo.FieldName}
    fieldInfo.EntityData.Leafs["field-value"] = types.YLeaf{"FieldValue", fieldInfo.FieldValue}
    fieldInfo.EntityData.Leafs["is-overflow"] = types.YLeaf{"IsOverflow", fieldInfo.IsOverflow}
    return &(fieldInfo.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers
// Ingress Stats
type Dpa_Stats_Nodes_Node_NpuNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats for a particular npu. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber.
    NpuNumber []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber
}

func (npuNumbers *Dpa_Stats_Nodes_Node_NpuNumbers) GetEntityData() *types.CommonEntityData {
    npuNumbers.EntityData.YFilter = npuNumbers.YFilter
    npuNumbers.EntityData.YangName = "npu-numbers"
    npuNumbers.EntityData.BundleName = "cisco_ios_xr"
    npuNumbers.EntityData.ParentYangName = "node"
    npuNumbers.EntityData.SegmentPath = "npu-numbers"
    npuNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuNumbers.EntityData.Children = make(map[string]types.YChild)
    npuNumbers.EntityData.Children["npu-number"] = types.YChild{"NpuNumber", nil}
    for i := range npuNumbers.NpuNumber {
        npuNumbers.EntityData.Children[types.GetSegmentPath(&npuNumbers.NpuNumber[i])] = types.YChild{"NpuNumber", &npuNumbers.NpuNumber[i]}
    }
    npuNumbers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(npuNumbers.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber
// Stats for a particular npu
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Npu number. The type is interface{} with range:
    // -2147483648..2147483647.
    NpuId interface{}

    // show npu specific voq or trap stats.
    Display Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display
}

func (npuNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber) GetEntityData() *types.CommonEntityData {
    npuNumber.EntityData.YFilter = npuNumber.YFilter
    npuNumber.EntityData.YangName = "npu-number"
    npuNumber.EntityData.BundleName = "cisco_ios_xr"
    npuNumber.EntityData.ParentYangName = "npu-numbers"
    npuNumber.EntityData.SegmentPath = "npu-number" + "[npu-id='" + fmt.Sprintf("%v", npuNumber.NpuId) + "']"
    npuNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuNumber.EntityData.Children = make(map[string]types.YChild)
    npuNumber.EntityData.Children["display"] = types.YChild{"Display", &npuNumber.Display}
    npuNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    npuNumber.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", npuNumber.NpuId}
    return &(npuNumber.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display
// show npu specific voq or trap stats
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Voq stats grouped by voq base numbers.
    BaseNumbers Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers

    // Trap stats for a particular npu.
    TrapIds Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds

    // Voq stats grouped by interface handle.
    InterfaceHandles Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles
}

func (display *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display) GetEntityData() *types.CommonEntityData {
    display.EntityData.YFilter = display.YFilter
    display.EntityData.YangName = "display"
    display.EntityData.BundleName = "cisco_ios_xr"
    display.EntityData.ParentYangName = "npu-number"
    display.EntityData.SegmentPath = "display"
    display.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    display.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    display.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    display.EntityData.Children = make(map[string]types.YChild)
    display.EntityData.Children["base-numbers"] = types.YChild{"BaseNumbers", &display.BaseNumbers}
    display.EntityData.Children["trap-ids"] = types.YChild{"TrapIds", &display.TrapIds}
    display.EntityData.Children["interface-handles"] = types.YChild{"InterfaceHandles", &display.InterfaceHandles}
    display.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(display.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers
// Voq stats grouped by voq base numbers
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Voq Base Number for a particular voq. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber.
    BaseNumber []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber
}

func (baseNumbers *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers) GetEntityData() *types.CommonEntityData {
    baseNumbers.EntityData.YFilter = baseNumbers.YFilter
    baseNumbers.EntityData.YangName = "base-numbers"
    baseNumbers.EntityData.BundleName = "cisco_ios_xr"
    baseNumbers.EntityData.ParentYangName = "display"
    baseNumbers.EntityData.SegmentPath = "base-numbers"
    baseNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseNumbers.EntityData.Children = make(map[string]types.YChild)
    baseNumbers.EntityData.Children["base-number"] = types.YChild{"BaseNumber", nil}
    for i := range baseNumbers.BaseNumber {
        baseNumbers.EntityData.Children[types.GetSegmentPath(&baseNumbers.BaseNumber[i])] = types.YChild{"BaseNumber", &baseNumbers.BaseNumber[i]}
    }
    baseNumbers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(baseNumbers.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber
// Voq Base Number for a particular voq
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber struct {
    EntityData types.CommonEntityData
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

func (baseNumber *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber) GetEntityData() *types.CommonEntityData {
    baseNumber.EntityData.YFilter = baseNumber.YFilter
    baseNumber.EntityData.YangName = "base-number"
    baseNumber.EntityData.BundleName = "cisco_ios_xr"
    baseNumber.EntityData.ParentYangName = "base-numbers"
    baseNumber.EntityData.SegmentPath = "base-number" + "[base-number='" + fmt.Sprintf("%v", baseNumber.BaseNumber) + "']"
    baseNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseNumber.EntityData.Children = make(map[string]types.YChild)
    baseNumber.EntityData.Children["voq-stat"] = types.YChild{"VoqStat", nil}
    for i := range baseNumber.VoqStat {
        baseNumber.EntityData.Children[types.GetSegmentPath(&baseNumber.VoqStat[i])] = types.YChild{"VoqStat", &baseNumber.VoqStat[i]}
    }
    baseNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    baseNumber.EntityData.Leafs["base-number"] = types.YLeaf{"BaseNumber", baseNumber.BaseNumber}
    baseNumber.EntityData.Leafs["in-use"] = types.YLeaf{"InUse", baseNumber.InUse}
    baseNumber.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", baseNumber.RackNum}
    baseNumber.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", baseNumber.SlotNum}
    baseNumber.EntityData.Leafs["npu-num"] = types.YLeaf{"NpuNum", baseNumber.NpuNum}
    baseNumber.EntityData.Leafs["npu-core"] = types.YLeaf{"NpuCore", baseNumber.NpuCore}
    baseNumber.EntityData.Leafs["port-num"] = types.YLeaf{"PortNum", baseNumber.PortNum}
    baseNumber.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", baseNumber.IfHandle}
    baseNumber.EntityData.Leafs["sys-port"] = types.YLeaf{"SysPort", baseNumber.SysPort}
    baseNumber.EntityData.Leafs["pp-port"] = types.YLeaf{"PpPort", baseNumber.PpPort}
    baseNumber.EntityData.Leafs["port-speed"] = types.YLeaf{"PortSpeed", baseNumber.PortSpeed}
    baseNumber.EntityData.Leafs["voq-base"] = types.YLeaf{"VoqBase", baseNumber.VoqBase}
    baseNumber.EntityData.Leafs["connector-id"] = types.YLeaf{"ConnectorId", baseNumber.ConnectorId}
    baseNumber.EntityData.Leafs["is-local-port"] = types.YLeaf{"IsLocalPort", baseNumber.IsLocalPort}
    return &(baseNumber.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat
// Keeps a record of the received and dropped
// packets and bytes on the port
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat struct {
    EntityData types.CommonEntityData
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

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_BaseNumbers_BaseNumber_VoqStat) GetEntityData() *types.CommonEntityData {
    voqStat.EntityData.YFilter = voqStat.YFilter
    voqStat.EntityData.YangName = "voq-stat"
    voqStat.EntityData.BundleName = "cisco_ios_xr"
    voqStat.EntityData.ParentYangName = "base-number"
    voqStat.EntityData.SegmentPath = "voq-stat"
    voqStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    voqStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    voqStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    voqStat.EntityData.Children = make(map[string]types.YChild)
    voqStat.EntityData.Leafs = make(map[string]types.YLeaf)
    voqStat.EntityData.Leafs["received-bytes"] = types.YLeaf{"ReceivedBytes", voqStat.ReceivedBytes}
    voqStat.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", voqStat.ReceivedPackets}
    voqStat.EntityData.Leafs["dropped-bytes"] = types.YLeaf{"DroppedBytes", voqStat.DroppedBytes}
    voqStat.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", voqStat.DroppedPackets}
    return &(voqStat.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds
// Trap stats for a particular npu
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter by specific trap id. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId.
    TrapId []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId
}

func (trapIds *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds) GetEntityData() *types.CommonEntityData {
    trapIds.EntityData.YFilter = trapIds.YFilter
    trapIds.EntityData.YangName = "trap-ids"
    trapIds.EntityData.BundleName = "cisco_ios_xr"
    trapIds.EntityData.ParentYangName = "display"
    trapIds.EntityData.SegmentPath = "trap-ids"
    trapIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapIds.EntityData.Children = make(map[string]types.YChild)
    trapIds.EntityData.Children["trap-id"] = types.YChild{"TrapId", nil}
    for i := range trapIds.TrapId {
        trapIds.EntityData.Children[types.GetSegmentPath(&trapIds.TrapId[i])] = types.YChild{"TrapId", &trapIds.TrapId[i]}
    }
    trapIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trapIds.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId
// Filter by specific trap id
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId struct {
    EntityData types.CommonEntityData
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

func (trapId *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_TrapIds_TrapId) GetEntityData() *types.CommonEntityData {
    trapId.EntityData.YFilter = trapId.YFilter
    trapId.EntityData.YangName = "trap-id"
    trapId.EntityData.BundleName = "cisco_ios_xr"
    trapId.EntityData.ParentYangName = "trap-ids"
    trapId.EntityData.SegmentPath = "trap-id" + "[trap-id='" + fmt.Sprintf("%v", trapId.TrapId) + "']"
    trapId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapId.EntityData.Children = make(map[string]types.YChild)
    trapId.EntityData.Leafs = make(map[string]types.YLeaf)
    trapId.EntityData.Leafs["trap-id"] = types.YLeaf{"TrapId", trapId.TrapId}
    trapId.EntityData.Leafs["trap-strength"] = types.YLeaf{"TrapStrength", trapId.TrapStrength}
    trapId.EntityData.Leafs["priority"] = types.YLeaf{"Priority", trapId.Priority}
    trapId.EntityData.Leafs["trap-id-xr"] = types.YLeaf{"TrapIdXr", trapId.TrapIdXr}
    trapId.EntityData.Leafs["gport"] = types.YLeaf{"Gport", trapId.Gport}
    trapId.EntityData.Leafs["fec-id"] = types.YLeaf{"FecId", trapId.FecId}
    trapId.EntityData.Leafs["policer-id"] = types.YLeaf{"PolicerId", trapId.PolicerId}
    trapId.EntityData.Leafs["stats-id"] = types.YLeaf{"StatsId", trapId.StatsId}
    trapId.EntityData.Leafs["encap-id"] = types.YLeaf{"EncapId", trapId.EncapId}
    trapId.EntityData.Leafs["mc-group"] = types.YLeaf{"McGroup", trapId.McGroup}
    trapId.EntityData.Leafs["trap-string"] = types.YLeaf{"TrapString", trapId.TrapString}
    trapId.EntityData.Leafs["id"] = types.YLeaf{"Id", trapId.Id}
    trapId.EntityData.Leafs["offset"] = types.YLeaf{"Offset", trapId.Offset}
    trapId.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", trapId.NpuId}
    trapId.EntityData.Leafs["packet-dropped"] = types.YLeaf{"PacketDropped", trapId.PacketDropped}
    trapId.EntityData.Leafs["packet-accepted"] = types.YLeaf{"PacketAccepted", trapId.PacketAccepted}
    return &(trapId.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles
// Voq stats grouped by interface handle
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Voq stats for a particular interface handle. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle.
    InterfaceHandle []Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle
}

func (interfaceHandles *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles) GetEntityData() *types.CommonEntityData {
    interfaceHandles.EntityData.YFilter = interfaceHandles.YFilter
    interfaceHandles.EntityData.YangName = "interface-handles"
    interfaceHandles.EntityData.BundleName = "cisco_ios_xr"
    interfaceHandles.EntityData.ParentYangName = "display"
    interfaceHandles.EntityData.SegmentPath = "interface-handles"
    interfaceHandles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceHandles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceHandles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceHandles.EntityData.Children = make(map[string]types.YChild)
    interfaceHandles.EntityData.Children["interface-handle"] = types.YChild{"InterfaceHandle", nil}
    for i := range interfaceHandles.InterfaceHandle {
        interfaceHandles.EntityData.Children[types.GetSegmentPath(&interfaceHandles.InterfaceHandle[i])] = types.YChild{"InterfaceHandle", &interfaceHandles.InterfaceHandle[i]}
    }
    interfaceHandles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceHandles.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle
// Voq stats for a particular interface
// handle
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle struct {
    EntityData types.CommonEntityData
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

func (interfaceHandle *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle) GetEntityData() *types.CommonEntityData {
    interfaceHandle.EntityData.YFilter = interfaceHandle.YFilter
    interfaceHandle.EntityData.YangName = "interface-handle"
    interfaceHandle.EntityData.BundleName = "cisco_ios_xr"
    interfaceHandle.EntityData.ParentYangName = "interface-handles"
    interfaceHandle.EntityData.SegmentPath = "interface-handle" + "[interface-handle='" + fmt.Sprintf("%v", interfaceHandle.InterfaceHandle) + "']"
    interfaceHandle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceHandle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceHandle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceHandle.EntityData.Children = make(map[string]types.YChild)
    interfaceHandle.EntityData.Children["voq-stat"] = types.YChild{"VoqStat", nil}
    for i := range interfaceHandle.VoqStat {
        interfaceHandle.EntityData.Children[types.GetSegmentPath(&interfaceHandle.VoqStat[i])] = types.YChild{"VoqStat", &interfaceHandle.VoqStat[i]}
    }
    interfaceHandle.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceHandle.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", interfaceHandle.InterfaceHandle}
    interfaceHandle.EntityData.Leafs["in-use"] = types.YLeaf{"InUse", interfaceHandle.InUse}
    interfaceHandle.EntityData.Leafs["rack-num"] = types.YLeaf{"RackNum", interfaceHandle.RackNum}
    interfaceHandle.EntityData.Leafs["slot-num"] = types.YLeaf{"SlotNum", interfaceHandle.SlotNum}
    interfaceHandle.EntityData.Leafs["npu-num"] = types.YLeaf{"NpuNum", interfaceHandle.NpuNum}
    interfaceHandle.EntityData.Leafs["npu-core"] = types.YLeaf{"NpuCore", interfaceHandle.NpuCore}
    interfaceHandle.EntityData.Leafs["port-num"] = types.YLeaf{"PortNum", interfaceHandle.PortNum}
    interfaceHandle.EntityData.Leafs["if-handle"] = types.YLeaf{"IfHandle", interfaceHandle.IfHandle}
    interfaceHandle.EntityData.Leafs["sys-port"] = types.YLeaf{"SysPort", interfaceHandle.SysPort}
    interfaceHandle.EntityData.Leafs["pp-port"] = types.YLeaf{"PpPort", interfaceHandle.PpPort}
    interfaceHandle.EntityData.Leafs["port-speed"] = types.YLeaf{"PortSpeed", interfaceHandle.PortSpeed}
    interfaceHandle.EntityData.Leafs["voq-base"] = types.YLeaf{"VoqBase", interfaceHandle.VoqBase}
    interfaceHandle.EntityData.Leafs["connector-id"] = types.YLeaf{"ConnectorId", interfaceHandle.ConnectorId}
    interfaceHandle.EntityData.Leafs["is-local-port"] = types.YLeaf{"IsLocalPort", interfaceHandle.IsLocalPort}
    return &(interfaceHandle.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat
// Keeps a record of the received and dropped
// packets and bytes on the port
type Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat struct {
    EntityData types.CommonEntityData
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

func (voqStat *Dpa_Stats_Nodes_Node_NpuNumbers_NpuNumber_Display_InterfaceHandles_InterfaceHandle_VoqStat) GetEntityData() *types.CommonEntityData {
    voqStat.EntityData.YFilter = voqStat.YFilter
    voqStat.EntityData.YangName = "voq-stat"
    voqStat.EntityData.BundleName = "cisco_ios_xr"
    voqStat.EntityData.ParentYangName = "interface-handle"
    voqStat.EntityData.SegmentPath = "voq-stat"
    voqStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    voqStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    voqStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    voqStat.EntityData.Children = make(map[string]types.YChild)
    voqStat.EntityData.Leafs = make(map[string]types.YLeaf)
    voqStat.EntityData.Leafs["received-bytes"] = types.YLeaf{"ReceivedBytes", voqStat.ReceivedBytes}
    voqStat.EntityData.Leafs["received-packets"] = types.YLeaf{"ReceivedPackets", voqStat.ReceivedPackets}
    voqStat.EntityData.Leafs["dropped-bytes"] = types.YLeaf{"DroppedBytes", voqStat.DroppedBytes}
    voqStat.EntityData.Leafs["dropped-packets"] = types.YLeaf{"DroppedPackets", voqStat.DroppedPackets}
    return &(voqStat.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuIds
// DPA stats hw resources info 
type Dpa_Stats_Nodes_Node_NpuIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats Hardware Info for particular NPU. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuIds_NpuId.
    NpuId []Dpa_Stats_Nodes_Node_NpuIds_NpuId
}

func (npuIds *Dpa_Stats_Nodes_Node_NpuIds) GetEntityData() *types.CommonEntityData {
    npuIds.EntityData.YFilter = npuIds.YFilter
    npuIds.EntityData.YangName = "npu-ids"
    npuIds.EntityData.BundleName = "cisco_ios_xr"
    npuIds.EntityData.ParentYangName = "node"
    npuIds.EntityData.SegmentPath = "npu-ids"
    npuIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuIds.EntityData.Children = make(map[string]types.YChild)
    npuIds.EntityData.Children["npu-id"] = types.YChild{"NpuId", nil}
    for i := range npuIds.NpuId {
        npuIds.EntityData.Children[types.GetSegmentPath(&npuIds.NpuId[i])] = types.YChild{"NpuId", &npuIds.NpuId[i]}
    }
    npuIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(npuIds.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuIds_NpuId
// Stats Hardware Info for particular NPU
type Dpa_Stats_Nodes_Node_NpuIds_NpuId struct {
    EntityData types.CommonEntityData
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

func (npuId *Dpa_Stats_Nodes_Node_NpuIds_NpuId) GetEntityData() *types.CommonEntityData {
    npuId.EntityData.YFilter = npuId.YFilter
    npuId.EntityData.YangName = "npu-id"
    npuId.EntityData.BundleName = "cisco_ios_xr"
    npuId.EntityData.ParentYangName = "npu-ids"
    npuId.EntityData.SegmentPath = "npu-id" + "[npu-id='" + fmt.Sprintf("%v", npuId.NpuId) + "']"
    npuId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npuId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npuId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npuId.EntityData.Children = make(map[string]types.YChild)
    npuId.EntityData.Children["cntr-engine"] = types.YChild{"CntrEngine", nil}
    for i := range npuId.CntrEngine {
        npuId.EntityData.Children[types.GetSegmentPath(&npuId.CntrEngine[i])] = types.YChild{"CntrEngine", &npuId.CntrEngine[i]}
    }
    npuId.EntityData.Leafs = make(map[string]types.YLeaf)
    npuId.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", npuId.NpuId}
    npuId.EntityData.Leafs["sys-cp-cnfg-prof"] = types.YLeaf{"SysCpCnfgProf", npuId.SysCpCnfgProf}
    npuId.EntityData.Leafs["next-avail-cp-id"] = types.YLeaf{"NextAvailCpId", npuId.NextAvailCpId}
    npuId.EntityData.Leafs["num-cntr-engines"] = types.YLeaf{"NumCntrEngines", npuId.NumCntrEngines}
    return &(npuId.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine
// cntr engine
type Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // state. The type is string.
    State interface{}

    // core id. The type is interface{} with range: 0..4294967295.
    CoreId interface{}

    // apps info. The type is slice of
    // Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo.
    AppsInfo []Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo
}

func (cntrEngine *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine) GetEntityData() *types.CommonEntityData {
    cntrEngine.EntityData.YFilter = cntrEngine.YFilter
    cntrEngine.EntityData.YangName = "cntr-engine"
    cntrEngine.EntityData.BundleName = "cisco_ios_xr"
    cntrEngine.EntityData.ParentYangName = "npu-id"
    cntrEngine.EntityData.SegmentPath = "cntr-engine"
    cntrEngine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cntrEngine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cntrEngine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cntrEngine.EntityData.Children = make(map[string]types.YChild)
    cntrEngine.EntityData.Children["apps-info"] = types.YChild{"AppsInfo", nil}
    for i := range cntrEngine.AppsInfo {
        cntrEngine.EntityData.Children[types.GetSegmentPath(&cntrEngine.AppsInfo[i])] = types.YChild{"AppsInfo", &cntrEngine.AppsInfo[i]}
    }
    cntrEngine.EntityData.Leafs = make(map[string]types.YLeaf)
    cntrEngine.EntityData.Leafs["state"] = types.YLeaf{"State", cntrEngine.State}
    cntrEngine.EntityData.Leafs["core-id"] = types.YLeaf{"CoreId", cntrEngine.CoreId}
    return &(cntrEngine.EntityData)
}

// Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo
// apps info
type Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // app type. The type is string.
    AppType interface{}

    // num cntrs for app. The type is interface{} with range: 0..4294967295.
    NumCntrsForApp interface{}

    // num cntrs used. The type is interface{} with range: 0..4294967295.
    NumCntrsUsed interface{}
}

func (appsInfo *Dpa_Stats_Nodes_Node_NpuIds_NpuId_CntrEngine_AppsInfo) GetEntityData() *types.CommonEntityData {
    appsInfo.EntityData.YFilter = appsInfo.YFilter
    appsInfo.EntityData.YangName = "apps-info"
    appsInfo.EntityData.BundleName = "cisco_ios_xr"
    appsInfo.EntityData.ParentYangName = "cntr-engine"
    appsInfo.EntityData.SegmentPath = "apps-info"
    appsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appsInfo.EntityData.Children = make(map[string]types.YChild)
    appsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    appsInfo.EntityData.Leafs["app-type"] = types.YLeaf{"AppType", appsInfo.AppType}
    appsInfo.EntityData.Leafs["num-cntrs-for-app"] = types.YLeaf{"NumCntrsForApp", appsInfo.NumCntrsForApp}
    appsInfo.EntityData.Leafs["num-cntrs-used"] = types.YLeaf{"NumCntrsUsed", appsInfo.NumCntrsUsed}
    return &(appsInfo.EntityData)
}

