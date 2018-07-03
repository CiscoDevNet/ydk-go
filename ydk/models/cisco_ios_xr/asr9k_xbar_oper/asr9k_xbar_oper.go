// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-xbar package operational data.
// 
// This module contains definitions
// for the following management objects:
//   cross-bar-stats: Crossbar stats operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_xbar_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_xbar_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-xbar-oper cross-bar-stats}", reflect.TypeOf(CrossBarStats{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-xbar-oper:cross-bar-stats", reflect.TypeOf(CrossBarStats{}))
}

// CrossBarStats
// Crossbar stats operational data
type CrossBarStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Nodes.
    Nodes CrossBarStats_Nodes
}

func (crossBarStats *CrossBarStats) GetEntityData() *types.CommonEntityData {
    crossBarStats.EntityData.YFilter = crossBarStats.YFilter
    crossBarStats.EntityData.YangName = "cross-bar-stats"
    crossBarStats.EntityData.BundleName = "cisco_ios_xr"
    crossBarStats.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-xbar-oper"
    crossBarStats.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-xbar-oper:cross-bar-stats"
    crossBarStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crossBarStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crossBarStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crossBarStats.EntityData.Children = types.NewOrderedMap()
    crossBarStats.EntityData.Children.Append("nodes", types.YChild{"Nodes", &crossBarStats.Nodes})
    crossBarStats.EntityData.Leafs = types.NewOrderedMap()

    crossBarStats.EntityData.YListKeys = []string {}

    return &(crossBarStats.EntityData)
}

// CrossBarStats_Nodes
// Table of Nodes
type CrossBarStats_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of
    // CrossBarStats_Nodes_Node.
    Node []*CrossBarStats_Nodes_Node
}

func (nodes *CrossBarStats_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "cross-bar-stats"
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

// CrossBarStats_Nodes_Node
// Information about a particular node
type CrossBarStats_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Table of stats information.
    CrossBarTable CrossBarStats_Nodes_Node_CrossBarTable
}

func (node *CrossBarStats_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("cross-bar-table", types.YChild{"CrossBarTable", &node.CrossBarTable})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable
// Table of stats information
type CrossBarStats_Nodes_Node_CrossBarTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of packet stats.
    PktStats CrossBarStats_Nodes_Node_CrossBarTable_PktStats

    // Table of packet stats for SM15.
    Sm15Stats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats
}

func (crossBarTable *CrossBarStats_Nodes_Node_CrossBarTable) GetEntityData() *types.CommonEntityData {
    crossBarTable.EntityData.YFilter = crossBarTable.YFilter
    crossBarTable.EntityData.YangName = "cross-bar-table"
    crossBarTable.EntityData.BundleName = "cisco_ios_xr"
    crossBarTable.EntityData.ParentYangName = "node"
    crossBarTable.EntityData.SegmentPath = "cross-bar-table"
    crossBarTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crossBarTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crossBarTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crossBarTable.EntityData.Children = types.NewOrderedMap()
    crossBarTable.EntityData.Children.Append("pkt-stats", types.YChild{"PktStats", &crossBarTable.PktStats})
    crossBarTable.EntityData.Children.Append("sm15-stats", types.YChild{"Sm15Stats", &crossBarTable.Sm15Stats})
    crossBarTable.EntityData.Leafs = types.NewOrderedMap()

    crossBarTable.EntityData.YListKeys = []string {}

    return &(crossBarTable.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_PktStats
// Table of packet stats
type CrossBarStats_Nodes_Node_CrossBarTable_PktStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats information for a particular asic type and port. The type is slice of
    // CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat.
    PktStat []*CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat
}

func (pktStats *CrossBarStats_Nodes_Node_CrossBarTable_PktStats) GetEntityData() *types.CommonEntityData {
    pktStats.EntityData.YFilter = pktStats.YFilter
    pktStats.EntityData.YangName = "pkt-stats"
    pktStats.EntityData.BundleName = "cisco_ios_xr"
    pktStats.EntityData.ParentYangName = "cross-bar-table"
    pktStats.EntityData.SegmentPath = "pkt-stats"
    pktStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pktStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pktStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pktStats.EntityData.Children = types.NewOrderedMap()
    pktStats.EntityData.Children.Append("pkt-stat", types.YChild{"PktStat", nil})
    for i := range pktStats.PktStat {
        pktStats.EntityData.Children.Append(types.GetSegmentPath(pktStats.PktStat[i]), types.YChild{"PktStat", pktStats.PktStat[i]})
    }
    pktStats.EntityData.Leafs = types.NewOrderedMap()

    pktStats.EntityData.YListKeys = []string {}

    return &(pktStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat
// Stats information for a particular asic type
// and port
type CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Asic ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    AsicId interface{}

    // Port. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Port interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    InternalErrorCount interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    InputBufferQueuedPacketCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    IngressPacketCountSinceLastReadHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    IngressChannelUtilizationCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    InputBufferBackPressureCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    XbarTimeoutDropCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    HoldropCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    NullFpoeDropCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DiagnosticPacketCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    InputBufferCorrectableEccErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    InputBufferUncorrectableEccErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    HeaderCrcErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    ShortInputHeaderErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    PacketCrcErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    ShortPacketErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    OutputBufferQueuedPacketCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressPacketCountSinceLastReadHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressChannelUtilizationCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    OutputBufferBackPressureCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    OutputBufferCorrectableEccErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    OutputBufferUncorrectableEccErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    FpoedbCorrectableEccErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    FpoedbUncorrectableEccErrorCountHigh interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    InputBufferQueuedPacketCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    IngressPacketCountSinceLastReadLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    IngressChannelUtilizationCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    InputBufferBackPressureCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    XbarTimeoutDropCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    HoldropCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    NullFpoeDropCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DiagnosticPacketCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    InputBufferCorrectableEccErrorCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    InputBufferUncorrectableEccErrorCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    HeaderCrcErrorCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    ShortInputHeaderErrorCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    PacketCrcErrorCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    ShortPacketErrorCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    OutputBufferQueuedPacketCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressPacketCountSinceLastReadLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressChannelUtilizationCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    OutputBufferBackPressureCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    OutputBufferCorrectableEccErrorCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    OutputBufferUncorrectableEccErrorCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    FpoedbCorrectableEccErrorCountLow interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    FpoedbUncorrectableEccErrorCountLow interface{}
}

func (pktStat *CrossBarStats_Nodes_Node_CrossBarTable_PktStats_PktStat) GetEntityData() *types.CommonEntityData {
    pktStat.EntityData.YFilter = pktStat.YFilter
    pktStat.EntityData.YangName = "pkt-stat"
    pktStat.EntityData.BundleName = "cisco_ios_xr"
    pktStat.EntityData.ParentYangName = "pkt-stats"
    pktStat.EntityData.SegmentPath = "pkt-stat"
    pktStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pktStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pktStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pktStat.EntityData.Children = types.NewOrderedMap()
    pktStat.EntityData.Leafs = types.NewOrderedMap()
    pktStat.EntityData.Leafs.Append("asic-id", types.YLeaf{"AsicId", pktStat.AsicId})
    pktStat.EntityData.Leafs.Append("port", types.YLeaf{"Port", pktStat.Port})
    pktStat.EntityData.Leafs.Append("internal-error-count", types.YLeaf{"InternalErrorCount", pktStat.InternalErrorCount})
    pktStat.EntityData.Leafs.Append("input-buffer-queued-packet-count-high", types.YLeaf{"InputBufferQueuedPacketCountHigh", pktStat.InputBufferQueuedPacketCountHigh})
    pktStat.EntityData.Leafs.Append("ingress-packet-count-since-last-read-high", types.YLeaf{"IngressPacketCountSinceLastReadHigh", pktStat.IngressPacketCountSinceLastReadHigh})
    pktStat.EntityData.Leafs.Append("ingress-channel-utilization-count-high", types.YLeaf{"IngressChannelUtilizationCountHigh", pktStat.IngressChannelUtilizationCountHigh})
    pktStat.EntityData.Leafs.Append("input-buffer-back-pressure-count-high", types.YLeaf{"InputBufferBackPressureCountHigh", pktStat.InputBufferBackPressureCountHigh})
    pktStat.EntityData.Leafs.Append("xbar-timeout-drop-count-high", types.YLeaf{"XbarTimeoutDropCountHigh", pktStat.XbarTimeoutDropCountHigh})
    pktStat.EntityData.Leafs.Append("holdrop-count-high", types.YLeaf{"HoldropCountHigh", pktStat.HoldropCountHigh})
    pktStat.EntityData.Leafs.Append("null-fpoe-drop-count-high", types.YLeaf{"NullFpoeDropCountHigh", pktStat.NullFpoeDropCountHigh})
    pktStat.EntityData.Leafs.Append("diagnostic-packet-count-high", types.YLeaf{"DiagnosticPacketCountHigh", pktStat.DiagnosticPacketCountHigh})
    pktStat.EntityData.Leafs.Append("input-buffer-correctable-ecc-error-count-high", types.YLeaf{"InputBufferCorrectableEccErrorCountHigh", pktStat.InputBufferCorrectableEccErrorCountHigh})
    pktStat.EntityData.Leafs.Append("input-buffer-uncorrectable-ecc-error-count-high", types.YLeaf{"InputBufferUncorrectableEccErrorCountHigh", pktStat.InputBufferUncorrectableEccErrorCountHigh})
    pktStat.EntityData.Leafs.Append("header-crc-error-count-high", types.YLeaf{"HeaderCrcErrorCountHigh", pktStat.HeaderCrcErrorCountHigh})
    pktStat.EntityData.Leafs.Append("short-input-header-error-count-high", types.YLeaf{"ShortInputHeaderErrorCountHigh", pktStat.ShortInputHeaderErrorCountHigh})
    pktStat.EntityData.Leafs.Append("packet-crc-error-count-high", types.YLeaf{"PacketCrcErrorCountHigh", pktStat.PacketCrcErrorCountHigh})
    pktStat.EntityData.Leafs.Append("short-packet-error-count-high", types.YLeaf{"ShortPacketErrorCountHigh", pktStat.ShortPacketErrorCountHigh})
    pktStat.EntityData.Leafs.Append("output-buffer-queued-packet-count-high", types.YLeaf{"OutputBufferQueuedPacketCountHigh", pktStat.OutputBufferQueuedPacketCountHigh})
    pktStat.EntityData.Leafs.Append("egress-packet-count-since-last-read-high", types.YLeaf{"EgressPacketCountSinceLastReadHigh", pktStat.EgressPacketCountSinceLastReadHigh})
    pktStat.EntityData.Leafs.Append("egress-channel-utilization-count-high", types.YLeaf{"EgressChannelUtilizationCountHigh", pktStat.EgressChannelUtilizationCountHigh})
    pktStat.EntityData.Leafs.Append("output-buffer-back-pressure-count-high", types.YLeaf{"OutputBufferBackPressureCountHigh", pktStat.OutputBufferBackPressureCountHigh})
    pktStat.EntityData.Leafs.Append("output-buffer-correctable-ecc-error-count-high", types.YLeaf{"OutputBufferCorrectableEccErrorCountHigh", pktStat.OutputBufferCorrectableEccErrorCountHigh})
    pktStat.EntityData.Leafs.Append("output-buffer-uncorrectable-ecc-error-count-high", types.YLeaf{"OutputBufferUncorrectableEccErrorCountHigh", pktStat.OutputBufferUncorrectableEccErrorCountHigh})
    pktStat.EntityData.Leafs.Append("fpoedb-correctable-ecc-error-count-high", types.YLeaf{"FpoedbCorrectableEccErrorCountHigh", pktStat.FpoedbCorrectableEccErrorCountHigh})
    pktStat.EntityData.Leafs.Append("fpoedb-uncorrectable-ecc-error-count-high", types.YLeaf{"FpoedbUncorrectableEccErrorCountHigh", pktStat.FpoedbUncorrectableEccErrorCountHigh})
    pktStat.EntityData.Leafs.Append("input-buffer-queued-packet-count-low", types.YLeaf{"InputBufferQueuedPacketCountLow", pktStat.InputBufferQueuedPacketCountLow})
    pktStat.EntityData.Leafs.Append("ingress-packet-count-since-last-read-low", types.YLeaf{"IngressPacketCountSinceLastReadLow", pktStat.IngressPacketCountSinceLastReadLow})
    pktStat.EntityData.Leafs.Append("ingress-channel-utilization-count-low", types.YLeaf{"IngressChannelUtilizationCountLow", pktStat.IngressChannelUtilizationCountLow})
    pktStat.EntityData.Leafs.Append("input-buffer-back-pressure-count-low", types.YLeaf{"InputBufferBackPressureCountLow", pktStat.InputBufferBackPressureCountLow})
    pktStat.EntityData.Leafs.Append("xbar-timeout-drop-count-low", types.YLeaf{"XbarTimeoutDropCountLow", pktStat.XbarTimeoutDropCountLow})
    pktStat.EntityData.Leafs.Append("holdrop-count-low", types.YLeaf{"HoldropCountLow", pktStat.HoldropCountLow})
    pktStat.EntityData.Leafs.Append("null-fpoe-drop-count-low", types.YLeaf{"NullFpoeDropCountLow", pktStat.NullFpoeDropCountLow})
    pktStat.EntityData.Leafs.Append("diagnostic-packet-count-low", types.YLeaf{"DiagnosticPacketCountLow", pktStat.DiagnosticPacketCountLow})
    pktStat.EntityData.Leafs.Append("input-buffer-correctable-ecc-error-count-low", types.YLeaf{"InputBufferCorrectableEccErrorCountLow", pktStat.InputBufferCorrectableEccErrorCountLow})
    pktStat.EntityData.Leafs.Append("input-buffer-uncorrectable-ecc-error-count-low", types.YLeaf{"InputBufferUncorrectableEccErrorCountLow", pktStat.InputBufferUncorrectableEccErrorCountLow})
    pktStat.EntityData.Leafs.Append("header-crc-error-count-low", types.YLeaf{"HeaderCrcErrorCountLow", pktStat.HeaderCrcErrorCountLow})
    pktStat.EntityData.Leafs.Append("short-input-header-error-count-low", types.YLeaf{"ShortInputHeaderErrorCountLow", pktStat.ShortInputHeaderErrorCountLow})
    pktStat.EntityData.Leafs.Append("packet-crc-error-count-low", types.YLeaf{"PacketCrcErrorCountLow", pktStat.PacketCrcErrorCountLow})
    pktStat.EntityData.Leafs.Append("short-packet-error-count-low", types.YLeaf{"ShortPacketErrorCountLow", pktStat.ShortPacketErrorCountLow})
    pktStat.EntityData.Leafs.Append("output-buffer-queued-packet-count-low", types.YLeaf{"OutputBufferQueuedPacketCountLow", pktStat.OutputBufferQueuedPacketCountLow})
    pktStat.EntityData.Leafs.Append("egress-packet-count-since-last-read-low", types.YLeaf{"EgressPacketCountSinceLastReadLow", pktStat.EgressPacketCountSinceLastReadLow})
    pktStat.EntityData.Leafs.Append("egress-channel-utilization-count-low", types.YLeaf{"EgressChannelUtilizationCountLow", pktStat.EgressChannelUtilizationCountLow})
    pktStat.EntityData.Leafs.Append("output-buffer-back-pressure-count-low", types.YLeaf{"OutputBufferBackPressureCountLow", pktStat.OutputBufferBackPressureCountLow})
    pktStat.EntityData.Leafs.Append("output-buffer-correctable-ecc-error-count-low", types.YLeaf{"OutputBufferCorrectableEccErrorCountLow", pktStat.OutputBufferCorrectableEccErrorCountLow})
    pktStat.EntityData.Leafs.Append("output-buffer-uncorrectable-ecc-error-count-low", types.YLeaf{"OutputBufferUncorrectableEccErrorCountLow", pktStat.OutputBufferUncorrectableEccErrorCountLow})
    pktStat.EntityData.Leafs.Append("fpoedb-correctable-ecc-error-count-low", types.YLeaf{"FpoedbCorrectableEccErrorCountLow", pktStat.FpoedbCorrectableEccErrorCountLow})
    pktStat.EntityData.Leafs.Append("fpoedb-uncorrectable-ecc-error-count-low", types.YLeaf{"FpoedbUncorrectableEccErrorCountLow", pktStat.FpoedbUncorrectableEccErrorCountLow})

    pktStat.EntityData.YListKeys = []string {}

    return &(pktStat.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats
// Table of packet stats for SM15
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats information for a particular asic type and port. The type is slice of
    // CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat.
    Sm15Stat []*CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat
}

func (sm15Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats) GetEntityData() *types.CommonEntityData {
    sm15Stats.EntityData.YFilter = sm15Stats.YFilter
    sm15Stats.EntityData.YangName = "sm15-stats"
    sm15Stats.EntityData.BundleName = "cisco_ios_xr"
    sm15Stats.EntityData.ParentYangName = "cross-bar-table"
    sm15Stats.EntityData.SegmentPath = "sm15-stats"
    sm15Stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sm15Stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sm15Stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sm15Stats.EntityData.Children = types.NewOrderedMap()
    sm15Stats.EntityData.Children.Append("sm15-stat", types.YChild{"Sm15Stat", nil})
    for i := range sm15Stats.Sm15Stat {
        sm15Stats.EntityData.Children.Append(types.GetSegmentPath(sm15Stats.Sm15Stat[i]), types.YChild{"Sm15Stat", sm15Stats.Sm15Stat[i]})
    }
    sm15Stats.EntityData.Leafs = types.NewOrderedMap()

    sm15Stats.EntityData.YListKeys = []string {}

    return &(sm15Stats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat
// Stats information for a particular asic type
// and port
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Asic ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    AsicId interface{}

    // Port. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Port interface{}

    // internal err cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalErrCnt interface{}

    // ua0 stats.
    Ua0Stats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats

    // ua1 stats.
    Ua1Stats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats

    // ua2 stats.
    Ua2Stats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats

    // ma stats.
    MaStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats

    // ca stats.
    CaStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats

    // pi stats.
    PiStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats

    // pe stats.
    PeStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats

    // pi uc stats.
    PiUcStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats

    // pi mc stats.
    PiMcStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats

    // pi cc stats.
    PiCcStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats

    // pe uc stats.
    PeUcStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats

    // pe mc stats.
    PeMcStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats

    // pe cc stats.
    PeCcStats CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats
}

func (sm15Stat *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat) GetEntityData() *types.CommonEntityData {
    sm15Stat.EntityData.YFilter = sm15Stat.YFilter
    sm15Stat.EntityData.YangName = "sm15-stat"
    sm15Stat.EntityData.BundleName = "cisco_ios_xr"
    sm15Stat.EntityData.ParentYangName = "sm15-stats"
    sm15Stat.EntityData.SegmentPath = "sm15-stat"
    sm15Stat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sm15Stat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sm15Stat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sm15Stat.EntityData.Children = types.NewOrderedMap()
    sm15Stat.EntityData.Children.Append("ua0-stats", types.YChild{"Ua0Stats", &sm15Stat.Ua0Stats})
    sm15Stat.EntityData.Children.Append("ua1-stats", types.YChild{"Ua1Stats", &sm15Stat.Ua1Stats})
    sm15Stat.EntityData.Children.Append("ua2-stats", types.YChild{"Ua2Stats", &sm15Stat.Ua2Stats})
    sm15Stat.EntityData.Children.Append("ma-stats", types.YChild{"MaStats", &sm15Stat.MaStats})
    sm15Stat.EntityData.Children.Append("ca-stats", types.YChild{"CaStats", &sm15Stat.CaStats})
    sm15Stat.EntityData.Children.Append("pi-stats", types.YChild{"PiStats", &sm15Stat.PiStats})
    sm15Stat.EntityData.Children.Append("pe-stats", types.YChild{"PeStats", &sm15Stat.PeStats})
    sm15Stat.EntityData.Children.Append("pi-uc-stats", types.YChild{"PiUcStats", &sm15Stat.PiUcStats})
    sm15Stat.EntityData.Children.Append("pi-mc-stats", types.YChild{"PiMcStats", &sm15Stat.PiMcStats})
    sm15Stat.EntityData.Children.Append("pi-cc-stats", types.YChild{"PiCcStats", &sm15Stat.PiCcStats})
    sm15Stat.EntityData.Children.Append("pe-uc-stats", types.YChild{"PeUcStats", &sm15Stat.PeUcStats})
    sm15Stat.EntityData.Children.Append("pe-mc-stats", types.YChild{"PeMcStats", &sm15Stat.PeMcStats})
    sm15Stat.EntityData.Children.Append("pe-cc-stats", types.YChild{"PeCcStats", &sm15Stat.PeCcStats})
    sm15Stat.EntityData.Leafs = types.NewOrderedMap()
    sm15Stat.EntityData.Leafs.Append("asic-id", types.YLeaf{"AsicId", sm15Stat.AsicId})
    sm15Stat.EntityData.Leafs.Append("port", types.YLeaf{"Port", sm15Stat.Port})
    sm15Stat.EntityData.Leafs.Append("internal-err-cnt", types.YLeaf{"InternalErrCnt", sm15Stat.InternalErrCnt})

    sm15Stat.EntityData.YListKeys = []string {}

    return &(sm15Stat.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats
// ua0 stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DEST DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestDropPktCnt interface{}

    // SRC DEST PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    SrcDestPktCnt interface{}

    // DEST SRC PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestSrcPktCnt interface{}

    // RCV PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    RcvPktCnt interface{}

    // TX PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    TxPktCnt interface{}

    // RX DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDropPktCnt interface{}

    // RX FABRIC TO CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxFabricToCnt interface{}

    // ACK WAIT CNT. The type is interface{} with range: 0..18446744073709551615.
    AckWaitCnt interface{}
}

func (ua0Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua0Stats) GetEntityData() *types.CommonEntityData {
    ua0Stats.EntityData.YFilter = ua0Stats.YFilter
    ua0Stats.EntityData.YangName = "ua0-stats"
    ua0Stats.EntityData.BundleName = "cisco_ios_xr"
    ua0Stats.EntityData.ParentYangName = "sm15-stat"
    ua0Stats.EntityData.SegmentPath = "ua0-stats"
    ua0Stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ua0Stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ua0Stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ua0Stats.EntityData.Children = types.NewOrderedMap()
    ua0Stats.EntityData.Leafs = types.NewOrderedMap()
    ua0Stats.EntityData.Leafs.Append("dest-drop-pkt-cnt", types.YLeaf{"DestDropPktCnt", ua0Stats.DestDropPktCnt})
    ua0Stats.EntityData.Leafs.Append("src-dest-pkt-cnt", types.YLeaf{"SrcDestPktCnt", ua0Stats.SrcDestPktCnt})
    ua0Stats.EntityData.Leafs.Append("dest-src-pkt-cnt", types.YLeaf{"DestSrcPktCnt", ua0Stats.DestSrcPktCnt})
    ua0Stats.EntityData.Leafs.Append("rcv-pkt-cnt", types.YLeaf{"RcvPktCnt", ua0Stats.RcvPktCnt})
    ua0Stats.EntityData.Leafs.Append("tx-pkt-cnt", types.YLeaf{"TxPktCnt", ua0Stats.TxPktCnt})
    ua0Stats.EntityData.Leafs.Append("rx-drop-pkt-cnt", types.YLeaf{"RxDropPktCnt", ua0Stats.RxDropPktCnt})
    ua0Stats.EntityData.Leafs.Append("rx-fabric-to-cnt", types.YLeaf{"RxFabricToCnt", ua0Stats.RxFabricToCnt})
    ua0Stats.EntityData.Leafs.Append("ack-wait-cnt", types.YLeaf{"AckWaitCnt", ua0Stats.AckWaitCnt})

    ua0Stats.EntityData.YListKeys = []string {}

    return &(ua0Stats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats
// ua1 stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DEST DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestDropPktCnt interface{}

    // SRC DEST PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    SrcDestPktCnt interface{}

    // DEST SRC PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestSrcPktCnt interface{}

    // RCV PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    RcvPktCnt interface{}

    // TX PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    TxPktCnt interface{}

    // RX DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDropPktCnt interface{}

    // RX FABRIC TO CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxFabricToCnt interface{}

    // ACK WAIT CNT. The type is interface{} with range: 0..18446744073709551615.
    AckWaitCnt interface{}
}

func (ua1Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua1Stats) GetEntityData() *types.CommonEntityData {
    ua1Stats.EntityData.YFilter = ua1Stats.YFilter
    ua1Stats.EntityData.YangName = "ua1-stats"
    ua1Stats.EntityData.BundleName = "cisco_ios_xr"
    ua1Stats.EntityData.ParentYangName = "sm15-stat"
    ua1Stats.EntityData.SegmentPath = "ua1-stats"
    ua1Stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ua1Stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ua1Stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ua1Stats.EntityData.Children = types.NewOrderedMap()
    ua1Stats.EntityData.Leafs = types.NewOrderedMap()
    ua1Stats.EntityData.Leafs.Append("dest-drop-pkt-cnt", types.YLeaf{"DestDropPktCnt", ua1Stats.DestDropPktCnt})
    ua1Stats.EntityData.Leafs.Append("src-dest-pkt-cnt", types.YLeaf{"SrcDestPktCnt", ua1Stats.SrcDestPktCnt})
    ua1Stats.EntityData.Leafs.Append("dest-src-pkt-cnt", types.YLeaf{"DestSrcPktCnt", ua1Stats.DestSrcPktCnt})
    ua1Stats.EntityData.Leafs.Append("rcv-pkt-cnt", types.YLeaf{"RcvPktCnt", ua1Stats.RcvPktCnt})
    ua1Stats.EntityData.Leafs.Append("tx-pkt-cnt", types.YLeaf{"TxPktCnt", ua1Stats.TxPktCnt})
    ua1Stats.EntityData.Leafs.Append("rx-drop-pkt-cnt", types.YLeaf{"RxDropPktCnt", ua1Stats.RxDropPktCnt})
    ua1Stats.EntityData.Leafs.Append("rx-fabric-to-cnt", types.YLeaf{"RxFabricToCnt", ua1Stats.RxFabricToCnt})
    ua1Stats.EntityData.Leafs.Append("ack-wait-cnt", types.YLeaf{"AckWaitCnt", ua1Stats.AckWaitCnt})

    ua1Stats.EntityData.YListKeys = []string {}

    return &(ua1Stats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats
// ua2 stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DEST DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestDropPktCnt interface{}

    // SRC DEST PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    SrcDestPktCnt interface{}

    // DEST SRC PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestSrcPktCnt interface{}

    // RCV PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    RcvPktCnt interface{}

    // TX PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    TxPktCnt interface{}

    // RX DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDropPktCnt interface{}

    // RX FABRIC TO CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxFabricToCnt interface{}

    // ACK WAIT CNT. The type is interface{} with range: 0..18446744073709551615.
    AckWaitCnt interface{}
}

func (ua2Stats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_Ua2Stats) GetEntityData() *types.CommonEntityData {
    ua2Stats.EntityData.YFilter = ua2Stats.YFilter
    ua2Stats.EntityData.YangName = "ua2-stats"
    ua2Stats.EntityData.BundleName = "cisco_ios_xr"
    ua2Stats.EntityData.ParentYangName = "sm15-stat"
    ua2Stats.EntityData.SegmentPath = "ua2-stats"
    ua2Stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ua2Stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ua2Stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ua2Stats.EntityData.Children = types.NewOrderedMap()
    ua2Stats.EntityData.Leafs = types.NewOrderedMap()
    ua2Stats.EntityData.Leafs.Append("dest-drop-pkt-cnt", types.YLeaf{"DestDropPktCnt", ua2Stats.DestDropPktCnt})
    ua2Stats.EntityData.Leafs.Append("src-dest-pkt-cnt", types.YLeaf{"SrcDestPktCnt", ua2Stats.SrcDestPktCnt})
    ua2Stats.EntityData.Leafs.Append("dest-src-pkt-cnt", types.YLeaf{"DestSrcPktCnt", ua2Stats.DestSrcPktCnt})
    ua2Stats.EntityData.Leafs.Append("rcv-pkt-cnt", types.YLeaf{"RcvPktCnt", ua2Stats.RcvPktCnt})
    ua2Stats.EntityData.Leafs.Append("tx-pkt-cnt", types.YLeaf{"TxPktCnt", ua2Stats.TxPktCnt})
    ua2Stats.EntityData.Leafs.Append("rx-drop-pkt-cnt", types.YLeaf{"RxDropPktCnt", ua2Stats.RxDropPktCnt})
    ua2Stats.EntityData.Leafs.Append("rx-fabric-to-cnt", types.YLeaf{"RxFabricToCnt", ua2Stats.RxFabricToCnt})
    ua2Stats.EntityData.Leafs.Append("ack-wait-cnt", types.YLeaf{"AckWaitCnt", ua2Stats.AckWaitCnt})

    ua2Stats.EntityData.YListKeys = []string {}

    return &(ua2Stats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats
// ma stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DEST DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestDropPktCnt interface{}

    // SRC DEST PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    SrcDestPktCnt interface{}

    // DEST SRC PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestSrcPktCnt interface{}

    // RCV PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    RcvPktCnt interface{}

    // TX PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    TxPktCnt interface{}

    // RX DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDropPktCnt interface{}

    // RX RETRANSMIT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxReTransmitCnt interface{}

    // RX FABRIC TO CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxFabricToCnt interface{}

    // RX HOL TO CNT. The type is interface{} with range: 0..18446744073709551615.
    RxHolToCnt interface{}
}

func (maStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_MaStats) GetEntityData() *types.CommonEntityData {
    maStats.EntityData.YFilter = maStats.YFilter
    maStats.EntityData.YangName = "ma-stats"
    maStats.EntityData.BundleName = "cisco_ios_xr"
    maStats.EntityData.ParentYangName = "sm15-stat"
    maStats.EntityData.SegmentPath = "ma-stats"
    maStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maStats.EntityData.Children = types.NewOrderedMap()
    maStats.EntityData.Leafs = types.NewOrderedMap()
    maStats.EntityData.Leafs.Append("dest-drop-pkt-cnt", types.YLeaf{"DestDropPktCnt", maStats.DestDropPktCnt})
    maStats.EntityData.Leafs.Append("src-dest-pkt-cnt", types.YLeaf{"SrcDestPktCnt", maStats.SrcDestPktCnt})
    maStats.EntityData.Leafs.Append("dest-src-pkt-cnt", types.YLeaf{"DestSrcPktCnt", maStats.DestSrcPktCnt})
    maStats.EntityData.Leafs.Append("rcv-pkt-cnt", types.YLeaf{"RcvPktCnt", maStats.RcvPktCnt})
    maStats.EntityData.Leafs.Append("tx-pkt-cnt", types.YLeaf{"TxPktCnt", maStats.TxPktCnt})
    maStats.EntityData.Leafs.Append("rx-drop-pkt-cnt", types.YLeaf{"RxDropPktCnt", maStats.RxDropPktCnt})
    maStats.EntityData.Leafs.Append("rx-re-transmit-cnt", types.YLeaf{"RxReTransmitCnt", maStats.RxReTransmitCnt})
    maStats.EntityData.Leafs.Append("rx-fabric-to-cnt", types.YLeaf{"RxFabricToCnt", maStats.RxFabricToCnt})
    maStats.EntityData.Leafs.Append("rx-hol-to-cnt", types.YLeaf{"RxHolToCnt", maStats.RxHolToCnt})

    maStats.EntityData.YListKeys = []string {}

    return &(maStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats
// ca stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DEST DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestDropPktCnt interface{}

    // SRC DEST PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    SrcDestPktCnt interface{}

    // DEST SRC PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DestSrcPktCnt interface{}

    // RCV PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    RcvPktCnt interface{}

    // TX PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    TxPktCnt interface{}

    // RX DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDropPktCnt interface{}
}

func (caStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_CaStats) GetEntityData() *types.CommonEntityData {
    caStats.EntityData.YFilter = caStats.YFilter
    caStats.EntityData.YangName = "ca-stats"
    caStats.EntityData.BundleName = "cisco_ios_xr"
    caStats.EntityData.ParentYangName = "sm15-stat"
    caStats.EntityData.SegmentPath = "ca-stats"
    caStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    caStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    caStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    caStats.EntityData.Children = types.NewOrderedMap()
    caStats.EntityData.Leafs = types.NewOrderedMap()
    caStats.EntityData.Leafs.Append("dest-drop-pkt-cnt", types.YLeaf{"DestDropPktCnt", caStats.DestDropPktCnt})
    caStats.EntityData.Leafs.Append("src-dest-pkt-cnt", types.YLeaf{"SrcDestPktCnt", caStats.SrcDestPktCnt})
    caStats.EntityData.Leafs.Append("dest-src-pkt-cnt", types.YLeaf{"DestSrcPktCnt", caStats.DestSrcPktCnt})
    caStats.EntityData.Leafs.Append("rcv-pkt-cnt", types.YLeaf{"RcvPktCnt", caStats.RcvPktCnt})
    caStats.EntityData.Leafs.Append("tx-pkt-cnt", types.YLeaf{"TxPktCnt", caStats.TxPktCnt})
    caStats.EntityData.Leafs.Append("rx-drop-pkt-cnt", types.YLeaf{"RxDropPktCnt", caStats.RxDropPktCnt})

    caStats.EntityData.YListKeys = []string {}

    return &(caStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats
// pi stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TOTAL RATE1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalRate1Cnt interface{}

    // TOTAL RATE2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalRate2Cnt interface{}

    // TOTAL RATE3 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalRate3Cnt interface{}

    // total calc rate. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalCalcRate interface{}
}

func (piStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiStats) GetEntityData() *types.CommonEntityData {
    piStats.EntityData.YFilter = piStats.YFilter
    piStats.EntityData.YangName = "pi-stats"
    piStats.EntityData.BundleName = "cisco_ios_xr"
    piStats.EntityData.ParentYangName = "sm15-stat"
    piStats.EntityData.SegmentPath = "pi-stats"
    piStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    piStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    piStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    piStats.EntityData.Children = types.NewOrderedMap()
    piStats.EntityData.Leafs = types.NewOrderedMap()
    piStats.EntityData.Leafs.Append("total-rate1-cnt", types.YLeaf{"TotalRate1Cnt", piStats.TotalRate1Cnt})
    piStats.EntityData.Leafs.Append("total-rate2-cnt", types.YLeaf{"TotalRate2Cnt", piStats.TotalRate2Cnt})
    piStats.EntityData.Leafs.Append("total-rate3-cnt", types.YLeaf{"TotalRate3Cnt", piStats.TotalRate3Cnt})
    piStats.EntityData.Leafs.Append("total-calc-rate", types.YLeaf{"TotalCalcRate", piStats.TotalCalcRate})

    piStats.EntityData.YListKeys = []string {}

    return &(piStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats
// pe stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TOTAL RATE1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalRate1Cnt interface{}

    // TOTAL RATE2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalRate2Cnt interface{}

    // TOTAL RATE3 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalRate3Cnt interface{}

    // total calc rate. The type is interface{} with range:
    // 0..18446744073709551615.
    TotalCalcRate interface{}

    // MC2UC PREEMPT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Mc2ucPreemptCnt interface{}
}

func (peStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeStats) GetEntityData() *types.CommonEntityData {
    peStats.EntityData.YFilter = peStats.YFilter
    peStats.EntityData.YangName = "pe-stats"
    peStats.EntityData.BundleName = "cisco_ios_xr"
    peStats.EntityData.ParentYangName = "sm15-stat"
    peStats.EntityData.SegmentPath = "pe-stats"
    peStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peStats.EntityData.Children = types.NewOrderedMap()
    peStats.EntityData.Leafs = types.NewOrderedMap()
    peStats.EntityData.Leafs.Append("total-rate1-cnt", types.YLeaf{"TotalRate1Cnt", peStats.TotalRate1Cnt})
    peStats.EntityData.Leafs.Append("total-rate2-cnt", types.YLeaf{"TotalRate2Cnt", peStats.TotalRate2Cnt})
    peStats.EntityData.Leafs.Append("total-rate3-cnt", types.YLeaf{"TotalRate3Cnt", peStats.TotalRate3Cnt})
    peStats.EntityData.Leafs.Append("total-calc-rate", types.YLeaf{"TotalCalcRate", peStats.TotalCalcRate})
    peStats.EntityData.Leafs.Append("mc2uc-preempt-cnt", types.YLeaf{"Mc2ucPreemptCnt", peStats.Mc2ucPreemptCnt})

    peStats.EntityData.YListKeys = []string {}

    return &(peStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats
// pi uc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PKT RCV CNT. The type is interface{} with range: 0..18446744073709551615.
    PktRcvCnt interface{}

    // PKT SEQ ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktSeqErrCnt interface{}

    // INCOMING PKT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    InComingPktErrCnt interface{}

    // MIN PKT LEN ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    MinPktLenErrCnt interface{}

    // MAX PKT LEN ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPktLenErrCnt interface{}

    // LINE ERR DRP PKT. The type is interface{} with range:
    // 0..18446744073709551615.
    LineErrDrpPkt interface{}

    // PKT CRC ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktCrcErrCnt interface{}

    // PKT CFH CRC ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktCfhCrcErrCnt interface{}

    // LINES WRITTEN IN MEM0. The type is interface{} with range:
    // 0..18446744073709551615.
    LineSWrittenInMem0 interface{}

    // LINES WRITTEN IN MEM1. The type is interface{} with range:
    // 0..18446744073709551615.
    LineSWrittenInMem1 interface{}

    // LINES WRITTEN IN MEM2. The type is interface{} with range:
    // 0..18446744073709551615.
    LineSWrittenInMem2 interface{}

    // TAIL DRP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TailDrpPktCnt interface{}

    // UC0 DATA MEM ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc0DataMemEcc1bitErrCnt interface{}

    // UC1 DATA MEM ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc1DataMemEcc1bitErrCnt interface{}

    // UC2 DATA MEM ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc2DataMemEcc1bitErrCnt interface{}

    // UC0 DATA MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc0DataMemEcc2bitErrCnt interface{}

    // UC1 DATA MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc1DataMemEcc2bitErrCnt interface{}

    // UC2 DATA MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Uc2DataMemEcc2bitErrCnt interface{}

    // DIAG PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    DiagPktCnt interface{}

    // PKT SENT TO DISABLED PORT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktSentToDisabledPortCnt interface{}

    // PKT NULL POE SENT UA0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktNullPoeSentUa0Cnt interface{}

    // PKT NULL POE SENT UA1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktNullPoeSentUa1Cnt interface{}

    // PKT NULL POE SENT UA2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktNullPoeSentUa2Cnt interface{}

    // PKT FPOE ADDR RNG HIT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktFpoeAddrRngHitCnt interface{}

    // FPOE MEM ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FpoeMemEcc1bitErrCnt interface{}

    // FPOE MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FpoeMemEcc2bitErrCnt interface{}

    // PKTS SENT TO UX0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsSentToUx0Cnt interface{}

    // PKTS SENT TO UX1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsSentToUx1Cnt interface{}

    // PKTS SENT TO UX2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsSentToUx2Cnt interface{}

    // CPP HEAD DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    CppHeadDropPktCnt interface{}

    // TR HEAD DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TrHeadDropPktCnt interface{}

    // TR PKT SENT TO UX. The type is interface{} with range:
    // 0..18446744073709551615.
    TrPktSentToUx interface{}

    // STOP THRSH HIT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    StopThrshHitCnt interface{}

    // RATE CNT. The type is interface{} with range: 0..18446744073709551615.
    RateCnt interface{}

    // calc rate. The type is interface{} with range: 0..18446744073709551615.
    CalcRate interface{}

    // CRC STOMP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    CrcStompPktCnt interface{}
}

func (piUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiUcStats) GetEntityData() *types.CommonEntityData {
    piUcStats.EntityData.YFilter = piUcStats.YFilter
    piUcStats.EntityData.YangName = "pi-uc-stats"
    piUcStats.EntityData.BundleName = "cisco_ios_xr"
    piUcStats.EntityData.ParentYangName = "sm15-stat"
    piUcStats.EntityData.SegmentPath = "pi-uc-stats"
    piUcStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    piUcStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    piUcStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    piUcStats.EntityData.Children = types.NewOrderedMap()
    piUcStats.EntityData.Leafs = types.NewOrderedMap()
    piUcStats.EntityData.Leafs.Append("pkt-rcv-cnt", types.YLeaf{"PktRcvCnt", piUcStats.PktRcvCnt})
    piUcStats.EntityData.Leafs.Append("pkt-seq-err-cnt", types.YLeaf{"PktSeqErrCnt", piUcStats.PktSeqErrCnt})
    piUcStats.EntityData.Leafs.Append("in-coming-pkt-err-cnt", types.YLeaf{"InComingPktErrCnt", piUcStats.InComingPktErrCnt})
    piUcStats.EntityData.Leafs.Append("min-pkt-len-err-cnt", types.YLeaf{"MinPktLenErrCnt", piUcStats.MinPktLenErrCnt})
    piUcStats.EntityData.Leafs.Append("max-pkt-len-err-cnt", types.YLeaf{"MaxPktLenErrCnt", piUcStats.MaxPktLenErrCnt})
    piUcStats.EntityData.Leafs.Append("line-err-drp-pkt", types.YLeaf{"LineErrDrpPkt", piUcStats.LineErrDrpPkt})
    piUcStats.EntityData.Leafs.Append("pkt-crc-err-cnt", types.YLeaf{"PktCrcErrCnt", piUcStats.PktCrcErrCnt})
    piUcStats.EntityData.Leafs.Append("pkt-cfh-crc-err-cnt", types.YLeaf{"PktCfhCrcErrCnt", piUcStats.PktCfhCrcErrCnt})
    piUcStats.EntityData.Leafs.Append("line-s-written-in-mem0", types.YLeaf{"LineSWrittenInMem0", piUcStats.LineSWrittenInMem0})
    piUcStats.EntityData.Leafs.Append("line-s-written-in-mem1", types.YLeaf{"LineSWrittenInMem1", piUcStats.LineSWrittenInMem1})
    piUcStats.EntityData.Leafs.Append("line-s-written-in-mem2", types.YLeaf{"LineSWrittenInMem2", piUcStats.LineSWrittenInMem2})
    piUcStats.EntityData.Leafs.Append("tail-drp-pkt-cnt", types.YLeaf{"TailDrpPktCnt", piUcStats.TailDrpPktCnt})
    piUcStats.EntityData.Leafs.Append("uc0-data-mem-ecc-1bit-err-cnt", types.YLeaf{"Uc0DataMemEcc1bitErrCnt", piUcStats.Uc0DataMemEcc1bitErrCnt})
    piUcStats.EntityData.Leafs.Append("uc1-data-mem-ecc-1bit-err-cnt", types.YLeaf{"Uc1DataMemEcc1bitErrCnt", piUcStats.Uc1DataMemEcc1bitErrCnt})
    piUcStats.EntityData.Leafs.Append("uc2-data-mem-ecc-1bit-err-cnt", types.YLeaf{"Uc2DataMemEcc1bitErrCnt", piUcStats.Uc2DataMemEcc1bitErrCnt})
    piUcStats.EntityData.Leafs.Append("uc0-data-mem-ecc-2bit-err-cnt", types.YLeaf{"Uc0DataMemEcc2bitErrCnt", piUcStats.Uc0DataMemEcc2bitErrCnt})
    piUcStats.EntityData.Leafs.Append("uc1-data-mem-ecc-2bit-err-cnt", types.YLeaf{"Uc1DataMemEcc2bitErrCnt", piUcStats.Uc1DataMemEcc2bitErrCnt})
    piUcStats.EntityData.Leafs.Append("uc2-data-mem-ecc-2bit-err-cnt", types.YLeaf{"Uc2DataMemEcc2bitErrCnt", piUcStats.Uc2DataMemEcc2bitErrCnt})
    piUcStats.EntityData.Leafs.Append("diag-pkt-cnt", types.YLeaf{"DiagPktCnt", piUcStats.DiagPktCnt})
    piUcStats.EntityData.Leafs.Append("pkt-sent-to-disabled-port-cnt", types.YLeaf{"PktSentToDisabledPortCnt", piUcStats.PktSentToDisabledPortCnt})
    piUcStats.EntityData.Leafs.Append("pkt-null-poe-sent-ua0-cnt", types.YLeaf{"PktNullPoeSentUa0Cnt", piUcStats.PktNullPoeSentUa0Cnt})
    piUcStats.EntityData.Leafs.Append("pkt-null-poe-sent-ua1-cnt", types.YLeaf{"PktNullPoeSentUa1Cnt", piUcStats.PktNullPoeSentUa1Cnt})
    piUcStats.EntityData.Leafs.Append("pkt-null-poe-sent-ua2-cnt", types.YLeaf{"PktNullPoeSentUa2Cnt", piUcStats.PktNullPoeSentUa2Cnt})
    piUcStats.EntityData.Leafs.Append("pkt-fpoe-addr-rng-hit-cnt", types.YLeaf{"PktFpoeAddrRngHitCnt", piUcStats.PktFpoeAddrRngHitCnt})
    piUcStats.EntityData.Leafs.Append("fpoe-mem-ecc-1bit-err-cnt", types.YLeaf{"FpoeMemEcc1bitErrCnt", piUcStats.FpoeMemEcc1bitErrCnt})
    piUcStats.EntityData.Leafs.Append("fpoe-mem-ecc-2bit-err-cnt", types.YLeaf{"FpoeMemEcc2bitErrCnt", piUcStats.FpoeMemEcc2bitErrCnt})
    piUcStats.EntityData.Leafs.Append("pkts-sent-to-ux0-cnt", types.YLeaf{"PktsSentToUx0Cnt", piUcStats.PktsSentToUx0Cnt})
    piUcStats.EntityData.Leafs.Append("pkts-sent-to-ux1-cnt", types.YLeaf{"PktsSentToUx1Cnt", piUcStats.PktsSentToUx1Cnt})
    piUcStats.EntityData.Leafs.Append("pkts-sent-to-ux2-cnt", types.YLeaf{"PktsSentToUx2Cnt", piUcStats.PktsSentToUx2Cnt})
    piUcStats.EntityData.Leafs.Append("cpp-head-drop-pkt-cnt", types.YLeaf{"CppHeadDropPktCnt", piUcStats.CppHeadDropPktCnt})
    piUcStats.EntityData.Leafs.Append("tr-head-drop-pkt-cnt", types.YLeaf{"TrHeadDropPktCnt", piUcStats.TrHeadDropPktCnt})
    piUcStats.EntityData.Leafs.Append("tr-pkt-sent-to-ux", types.YLeaf{"TrPktSentToUx", piUcStats.TrPktSentToUx})
    piUcStats.EntityData.Leafs.Append("stop-thrsh-hit-cnt", types.YLeaf{"StopThrshHitCnt", piUcStats.StopThrshHitCnt})
    piUcStats.EntityData.Leafs.Append("rate-cnt", types.YLeaf{"RateCnt", piUcStats.RateCnt})
    piUcStats.EntityData.Leafs.Append("calc-rate", types.YLeaf{"CalcRate", piUcStats.CalcRate})
    piUcStats.EntityData.Leafs.Append("crc-stomp-pkt-cnt", types.YLeaf{"CrcStompPktCnt", piUcStats.CrcStompPktCnt})

    piUcStats.EntityData.YListKeys = []string {}

    return &(piUcStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats
// pi mc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PKT RCV CNT. The type is interface{} with range: 0..18446744073709551615.
    PktRcvCnt interface{}

    // PKT SEQ ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktSeqErrCnt interface{}

    // INCOMING PKT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    InComingPktErrCnt interface{}

    // MIN PKT LEN ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    MinPktLenErrCnt interface{}

    // MAX PKT LEN ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPktLenErrCnt interface{}

    // LINE ERR DRP PKT. The type is interface{} with range:
    // 0..18446744073709551615.
    LineErrDrpPkt interface{}

    // PKT CRC ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktCrcErrCnt interface{}

    // PKT CFH CRC ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktCfhCrcErrCnt interface{}

    // LINES WRITTEN IN MEM. The type is interface{} with range:
    // 0..18446744073709551615.
    LineSWrittenInMem interface{}

    // TAIL DRP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TailDrpPktCnt interface{}

    // DATA MEM0 ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem0Ecc1bitErrCnt interface{}

    // DATA MEM1 ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem1Ecc1bitErrCnt interface{}

    // DATA MEM2 ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem2Ecc1bitErrCnt interface{}

    // DATA MEM0 ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem0Ecc2bitErrCnt interface{}

    // DATA MEM1 ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem1Ecc2bitErrCnt interface{}

    // DATA MEM2 ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMem2Ecc2bitErrCnt interface{}

    // DIAG PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    DiagPktCnt interface{}

    // PKT SENT TO DISABLED PORT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktSentToDisabledPort interface{}

    // PKT FPOE MATCH HIT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktFpoeMatchHitCnt interface{}

    // PKT NULL POE SENT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktNullPoeSentCnt interface{}

    // PKT FPOE ADDR RNG HIT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktFpoeAddrRngHitCnt interface{}

    // DI HDR LEN ERR PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DiHdrLenErrPktCnt interface{}

    // DI ERR PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DiErrPktCnt interface{}

    // FPOE MEM ECC 1BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FpoeMemEcc1bitErrCnt interface{}

    // FPOE MEM ECC 2BIT ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FpoeMemEcc2bitErrCnt interface{}

    // PKTS SENT TO MX CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsSentToMxCnt interface{}

    // CPP HEAD DROP PKT FROM MA CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    CppHeadDropPktFromMaCnt interface{}

    // TR HEAD DROP PKT FROM MA CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TrHeadDropPktFromMaCnt interface{}

    // TR PKT SENT TO MX. The type is interface{} with range:
    // 0..18446744073709551615.
    TrPktSentToMx interface{}

    // STOP THRSH HIT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    StopThrshHitCnt interface{}

    // RATE CNT. The type is interface{} with range: 0..18446744073709551615.
    RateCnt interface{}

    // calc rate. The type is interface{} with range: 0..18446744073709551615.
    CalcRate interface{}

    // CRC STOMP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    CrcStompPktCnt interface{}
}

func (piMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiMcStats) GetEntityData() *types.CommonEntityData {
    piMcStats.EntityData.YFilter = piMcStats.YFilter
    piMcStats.EntityData.YangName = "pi-mc-stats"
    piMcStats.EntityData.BundleName = "cisco_ios_xr"
    piMcStats.EntityData.ParentYangName = "sm15-stat"
    piMcStats.EntityData.SegmentPath = "pi-mc-stats"
    piMcStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    piMcStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    piMcStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    piMcStats.EntityData.Children = types.NewOrderedMap()
    piMcStats.EntityData.Leafs = types.NewOrderedMap()
    piMcStats.EntityData.Leafs.Append("pkt-rcv-cnt", types.YLeaf{"PktRcvCnt", piMcStats.PktRcvCnt})
    piMcStats.EntityData.Leafs.Append("pkt-seq-err-cnt", types.YLeaf{"PktSeqErrCnt", piMcStats.PktSeqErrCnt})
    piMcStats.EntityData.Leafs.Append("in-coming-pkt-err-cnt", types.YLeaf{"InComingPktErrCnt", piMcStats.InComingPktErrCnt})
    piMcStats.EntityData.Leafs.Append("min-pkt-len-err-cnt", types.YLeaf{"MinPktLenErrCnt", piMcStats.MinPktLenErrCnt})
    piMcStats.EntityData.Leafs.Append("max-pkt-len-err-cnt", types.YLeaf{"MaxPktLenErrCnt", piMcStats.MaxPktLenErrCnt})
    piMcStats.EntityData.Leafs.Append("line-err-drp-pkt", types.YLeaf{"LineErrDrpPkt", piMcStats.LineErrDrpPkt})
    piMcStats.EntityData.Leafs.Append("pkt-crc-err-cnt", types.YLeaf{"PktCrcErrCnt", piMcStats.PktCrcErrCnt})
    piMcStats.EntityData.Leafs.Append("pkt-cfh-crc-err-cnt", types.YLeaf{"PktCfhCrcErrCnt", piMcStats.PktCfhCrcErrCnt})
    piMcStats.EntityData.Leafs.Append("line-s-written-in-mem", types.YLeaf{"LineSWrittenInMem", piMcStats.LineSWrittenInMem})
    piMcStats.EntityData.Leafs.Append("tail-drp-pkt-cnt", types.YLeaf{"TailDrpPktCnt", piMcStats.TailDrpPktCnt})
    piMcStats.EntityData.Leafs.Append("data-mem0-ecc-1bit-err-cnt", types.YLeaf{"DataMem0Ecc1bitErrCnt", piMcStats.DataMem0Ecc1bitErrCnt})
    piMcStats.EntityData.Leafs.Append("data-mem1-ecc-1bit-err-cnt", types.YLeaf{"DataMem1Ecc1bitErrCnt", piMcStats.DataMem1Ecc1bitErrCnt})
    piMcStats.EntityData.Leafs.Append("data-mem2-ecc-1bit-err-cnt", types.YLeaf{"DataMem2Ecc1bitErrCnt", piMcStats.DataMem2Ecc1bitErrCnt})
    piMcStats.EntityData.Leafs.Append("data-mem0-ecc-2bit-err-cnt", types.YLeaf{"DataMem0Ecc2bitErrCnt", piMcStats.DataMem0Ecc2bitErrCnt})
    piMcStats.EntityData.Leafs.Append("data-mem1-ecc-2bit-err-cnt", types.YLeaf{"DataMem1Ecc2bitErrCnt", piMcStats.DataMem1Ecc2bitErrCnt})
    piMcStats.EntityData.Leafs.Append("data-mem2-ecc-2bit-err-cnt", types.YLeaf{"DataMem2Ecc2bitErrCnt", piMcStats.DataMem2Ecc2bitErrCnt})
    piMcStats.EntityData.Leafs.Append("diag-pkt-cnt", types.YLeaf{"DiagPktCnt", piMcStats.DiagPktCnt})
    piMcStats.EntityData.Leafs.Append("pkt-sent-to-disabled-port", types.YLeaf{"PktSentToDisabledPort", piMcStats.PktSentToDisabledPort})
    piMcStats.EntityData.Leafs.Append("pkt-fpoe-match-hit-cnt", types.YLeaf{"PktFpoeMatchHitCnt", piMcStats.PktFpoeMatchHitCnt})
    piMcStats.EntityData.Leafs.Append("pkt-null-poe-sent-cnt", types.YLeaf{"PktNullPoeSentCnt", piMcStats.PktNullPoeSentCnt})
    piMcStats.EntityData.Leafs.Append("pkt-fpoe-addr-rng-hit-cnt", types.YLeaf{"PktFpoeAddrRngHitCnt", piMcStats.PktFpoeAddrRngHitCnt})
    piMcStats.EntityData.Leafs.Append("di-hdr-len-err-pkt-cnt", types.YLeaf{"DiHdrLenErrPktCnt", piMcStats.DiHdrLenErrPktCnt})
    piMcStats.EntityData.Leafs.Append("di-err-pkt-cnt", types.YLeaf{"DiErrPktCnt", piMcStats.DiErrPktCnt})
    piMcStats.EntityData.Leafs.Append("fpoe-mem-ecc-1bit-err-cnt", types.YLeaf{"FpoeMemEcc1bitErrCnt", piMcStats.FpoeMemEcc1bitErrCnt})
    piMcStats.EntityData.Leafs.Append("fpoe-mem-ecc-2bit-err-cnt", types.YLeaf{"FpoeMemEcc2bitErrCnt", piMcStats.FpoeMemEcc2bitErrCnt})
    piMcStats.EntityData.Leafs.Append("pkts-sent-to-mx-cnt", types.YLeaf{"PktsSentToMxCnt", piMcStats.PktsSentToMxCnt})
    piMcStats.EntityData.Leafs.Append("cpp-head-drop-pkt-from-ma-cnt", types.YLeaf{"CppHeadDropPktFromMaCnt", piMcStats.CppHeadDropPktFromMaCnt})
    piMcStats.EntityData.Leafs.Append("tr-head-drop-pkt-from-ma-cnt", types.YLeaf{"TrHeadDropPktFromMaCnt", piMcStats.TrHeadDropPktFromMaCnt})
    piMcStats.EntityData.Leafs.Append("tr-pkt-sent-to-mx", types.YLeaf{"TrPktSentToMx", piMcStats.TrPktSentToMx})
    piMcStats.EntityData.Leafs.Append("stop-thrsh-hit-cnt", types.YLeaf{"StopThrshHitCnt", piMcStats.StopThrshHitCnt})
    piMcStats.EntityData.Leafs.Append("rate-cnt", types.YLeaf{"RateCnt", piMcStats.RateCnt})
    piMcStats.EntityData.Leafs.Append("calc-rate", types.YLeaf{"CalcRate", piMcStats.CalcRate})
    piMcStats.EntityData.Leafs.Append("crc-stomp-pkt-cnt", types.YLeaf{"CrcStompPktCnt", piMcStats.CrcStompPktCnt})

    piMcStats.EntityData.YListKeys = []string {}

    return &(piMcStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats
// pi cc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IN0 ECC SERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    In0EccSerrCnt interface{}

    // IN0 ECC DERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    In0EccDerrCnt interface{}

    // IN1 ECC SERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    In1EccSerrCnt interface{}

    // IN1 ECC DERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    In1EccDerrCnt interface{}

    // DATA MEM ECC SERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMemEccSerrCnt interface{}

    // DATA MEM ECC DERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMemEccDerrCnt interface{}

    // DATA MEM OVF0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMemOvf0Cnt interface{}

    // DATA MEM OVF1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    DataMemOvf1Cnt interface{}

    // FPOE MEM ECC SERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FpoeMemEccSerrCnt interface{}

    // FPOE MEM ECC DERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FpoeMemEccDerrCnt interface{}

    // NULL POE CNT. The type is interface{} with range: 0..18446744073709551615.
    NullPoeCnt interface{}

    // SHUT ACK CNT. The type is interface{} with range: 0..18446744073709551615.
    ShutAckCnt interface{}

    // IN0 FNC ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    In0FncErrCnt interface{}

    // IN1 FNC ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    In1FncErrCnt interface{}

    // IN0 DROP CNT. The type is interface{} with range: 0..18446744073709551615.
    In0DropCnt interface{}

    // IN1 DROP CNT. The type is interface{} with range: 0..18446744073709551615.
    In1DropCnt interface{}

    // IN0 CONG CNT. The type is interface{} with range: 0..18446744073709551615.
    In0CongCnt interface{}

    // IN1 CONG CNT. The type is interface{} with range: 0..18446744073709551615.
    In1CongCnt interface{}

    // IN0 SHUT CNT. The type is interface{} with range: 0..18446744073709551615.
    In0ShutCnt interface{}

    // IN1 SHUT CNT. The type is interface{} with range: 0..18446744073709551615.
    In1ShutCnt interface{}

    // TAIL DROP MSG CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    TailDropMsgCnt interface{}

    // IN0 PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    In0PktCnt interface{}

    // IN1 PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    In1PktCnt interface{}

    // DMEM RD CNT. The type is interface{} with range: 0..18446744073709551615.
    DmemRdCnt interface{}

    // IN DMEM0 CNT. The type is interface{} with range: 0..18446744073709551615.
    InDmem0Cnt interface{}

    // IN DMEM1 CNT. The type is interface{} with range: 0..18446744073709551615.
    InDmem1Cnt interface{}

    // OUT PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    OutPktCnt interface{}

    // STOP THRSH HIT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    StopThrshHitCnt interface{}

    // RATE CNT. The type is interface{} with range: 0..18446744073709551615.
    RateCnt interface{}

    // calc rate. The type is interface{} with range: 0..18446744073709551615.
    CalcRate interface{}
}

func (piCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PiCcStats) GetEntityData() *types.CommonEntityData {
    piCcStats.EntityData.YFilter = piCcStats.YFilter
    piCcStats.EntityData.YangName = "pi-cc-stats"
    piCcStats.EntityData.BundleName = "cisco_ios_xr"
    piCcStats.EntityData.ParentYangName = "sm15-stat"
    piCcStats.EntityData.SegmentPath = "pi-cc-stats"
    piCcStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    piCcStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    piCcStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    piCcStats.EntityData.Children = types.NewOrderedMap()
    piCcStats.EntityData.Leafs = types.NewOrderedMap()
    piCcStats.EntityData.Leafs.Append("in0-ecc-serr-cnt", types.YLeaf{"In0EccSerrCnt", piCcStats.In0EccSerrCnt})
    piCcStats.EntityData.Leafs.Append("in0-ecc-derr-cnt", types.YLeaf{"In0EccDerrCnt", piCcStats.In0EccDerrCnt})
    piCcStats.EntityData.Leafs.Append("in1-ecc-serr-cnt", types.YLeaf{"In1EccSerrCnt", piCcStats.In1EccSerrCnt})
    piCcStats.EntityData.Leafs.Append("in1-ecc-derr-cnt", types.YLeaf{"In1EccDerrCnt", piCcStats.In1EccDerrCnt})
    piCcStats.EntityData.Leafs.Append("data-mem-ecc-serr-cnt", types.YLeaf{"DataMemEccSerrCnt", piCcStats.DataMemEccSerrCnt})
    piCcStats.EntityData.Leafs.Append("data-mem-ecc-derr-cnt", types.YLeaf{"DataMemEccDerrCnt", piCcStats.DataMemEccDerrCnt})
    piCcStats.EntityData.Leafs.Append("data-mem-ovf0-cnt", types.YLeaf{"DataMemOvf0Cnt", piCcStats.DataMemOvf0Cnt})
    piCcStats.EntityData.Leafs.Append("data-mem-ovf1-cnt", types.YLeaf{"DataMemOvf1Cnt", piCcStats.DataMemOvf1Cnt})
    piCcStats.EntityData.Leafs.Append("fpoe-mem-ecc-serr-cnt", types.YLeaf{"FpoeMemEccSerrCnt", piCcStats.FpoeMemEccSerrCnt})
    piCcStats.EntityData.Leafs.Append("fpoe-mem-ecc-derr-cnt", types.YLeaf{"FpoeMemEccDerrCnt", piCcStats.FpoeMemEccDerrCnt})
    piCcStats.EntityData.Leafs.Append("null-poe-cnt", types.YLeaf{"NullPoeCnt", piCcStats.NullPoeCnt})
    piCcStats.EntityData.Leafs.Append("shut-ack-cnt", types.YLeaf{"ShutAckCnt", piCcStats.ShutAckCnt})
    piCcStats.EntityData.Leafs.Append("in0-fnc-err-cnt", types.YLeaf{"In0FncErrCnt", piCcStats.In0FncErrCnt})
    piCcStats.EntityData.Leafs.Append("in1-fnc-err-cnt", types.YLeaf{"In1FncErrCnt", piCcStats.In1FncErrCnt})
    piCcStats.EntityData.Leafs.Append("in0-drop-cnt", types.YLeaf{"In0DropCnt", piCcStats.In0DropCnt})
    piCcStats.EntityData.Leafs.Append("in1-drop-cnt", types.YLeaf{"In1DropCnt", piCcStats.In1DropCnt})
    piCcStats.EntityData.Leafs.Append("in0-cong-cnt", types.YLeaf{"In0CongCnt", piCcStats.In0CongCnt})
    piCcStats.EntityData.Leafs.Append("in1-cong-cnt", types.YLeaf{"In1CongCnt", piCcStats.In1CongCnt})
    piCcStats.EntityData.Leafs.Append("in0-shut-cnt", types.YLeaf{"In0ShutCnt", piCcStats.In0ShutCnt})
    piCcStats.EntityData.Leafs.Append("in1-shut-cnt", types.YLeaf{"In1ShutCnt", piCcStats.In1ShutCnt})
    piCcStats.EntityData.Leafs.Append("tail-drop-msg-cnt", types.YLeaf{"TailDropMsgCnt", piCcStats.TailDropMsgCnt})
    piCcStats.EntityData.Leafs.Append("in0-pkt-cnt", types.YLeaf{"In0PktCnt", piCcStats.In0PktCnt})
    piCcStats.EntityData.Leafs.Append("in1-pkt-cnt", types.YLeaf{"In1PktCnt", piCcStats.In1PktCnt})
    piCcStats.EntityData.Leafs.Append("dmem-rd-cnt", types.YLeaf{"DmemRdCnt", piCcStats.DmemRdCnt})
    piCcStats.EntityData.Leafs.Append("in-dmem0-cnt", types.YLeaf{"InDmem0Cnt", piCcStats.InDmem0Cnt})
    piCcStats.EntityData.Leafs.Append("in-dmem1-cnt", types.YLeaf{"InDmem1Cnt", piCcStats.InDmem1Cnt})
    piCcStats.EntityData.Leafs.Append("out-pkt-cnt", types.YLeaf{"OutPktCnt", piCcStats.OutPktCnt})
    piCcStats.EntityData.Leafs.Append("stop-thrsh-hit-cnt", types.YLeaf{"StopThrshHitCnt", piCcStats.StopThrshHitCnt})
    piCcStats.EntityData.Leafs.Append("rate-cnt", types.YLeaf{"RateCnt", piCcStats.RateCnt})
    piCcStats.EntityData.Leafs.Append("calc-rate", types.YLeaf{"CalcRate", piCcStats.CalcRate})

    piCcStats.EntityData.YListKeys = []string {}

    return &(piCcStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats
// pe uc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IN PKT UC0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUc0Cnt interface{}

    // IN PKT UC1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUc1Cnt interface{}

    // IN PKT UC2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUc2Cnt interface{}

    // IN FULL LINE UC0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    InFullLineUc0Cnt interface{}

    // IN FULL LINE UC1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    InFullLineUc1Cnt interface{}

    // IN FULL LINE UC2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    InFullLineUc2Cnt interface{}

    // PKT TRUNC EOP UC0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktTruncEopUc0Cnt interface{}

    // PKT TRUNC EOP UC1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktTruncEopUc1Cnt interface{}

    // PKT TRUNC EOP UC2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktTruncEopUc2Cnt interface{}

    // PKT SOP DROP UC0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktSopDropUc0Cnt interface{}

    // PKT SOP DROP UC1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktSopDropUc1Cnt interface{}

    // PKT SOP DROP UC2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktSopDropUc2Cnt interface{}

    // PKT ECC ERR DROP UC CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktEccErrDropUcCnt interface{}

    // PKT ECC TRUNC CNT UC CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktEccTruncCntUcCnt interface{}

    // ECC 1BIT ERR UC0 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1bitErrUc00Cnt interface{}

    // ECC 1BIT ERR UC0 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1bitErrUc01Cnt interface{}

    // ECC 1BIT ERR UC1 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1bitErrUc10Cnt interface{}

    // ECC 1BIT ERR UC1 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1bitErrUc11Cnt interface{}

    // ECC 1BIT ERR UC2 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1bitErrUc20Cnt interface{}

    // ECC 1BIT ERR UC2 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1bitErrUc21Cnt interface{}

    // ECC 2BIT ERR UC0 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2bitErrUc00Cnt interface{}

    // ECC 2BIT ERR UC0 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2bitErrUc01Cnt interface{}

    // ECC 2BIT ERR UC1 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2bitErrUc10Cnt interface{}

    // ECC 2BIT ERR UC1 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2bitErrUc11Cnt interface{}

    // ECC 2BIT ERR UC2 0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2bitErrUc20Cnt interface{}

    // ECC 2BIT ERR UC2 1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2bitErrUc21Cnt interface{}

    // OUT PKT UC CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktUcCnt interface{}

    // FE UC SOP EOP PACK CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FeUcSopEopPackCnt interface{}

    // FC UC 0 1 TRANS CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FcUc01TransCnt interface{}

    // RATE CNT. The type is interface{} with range: 0..18446744073709551615.
    RateCnt interface{}

    // calc rate. The type is interface{} with range: 0..18446744073709551615.
    CalcRate interface{}
}

func (peUcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeUcStats) GetEntityData() *types.CommonEntityData {
    peUcStats.EntityData.YFilter = peUcStats.YFilter
    peUcStats.EntityData.YangName = "pe-uc-stats"
    peUcStats.EntityData.BundleName = "cisco_ios_xr"
    peUcStats.EntityData.ParentYangName = "sm15-stat"
    peUcStats.EntityData.SegmentPath = "pe-uc-stats"
    peUcStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peUcStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peUcStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peUcStats.EntityData.Children = types.NewOrderedMap()
    peUcStats.EntityData.Leafs = types.NewOrderedMap()
    peUcStats.EntityData.Leafs.Append("in-pkt-uc0-cnt", types.YLeaf{"InPktUc0Cnt", peUcStats.InPktUc0Cnt})
    peUcStats.EntityData.Leafs.Append("in-pkt-uc1-cnt", types.YLeaf{"InPktUc1Cnt", peUcStats.InPktUc1Cnt})
    peUcStats.EntityData.Leafs.Append("in-pkt-uc2-cnt", types.YLeaf{"InPktUc2Cnt", peUcStats.InPktUc2Cnt})
    peUcStats.EntityData.Leafs.Append("in-full-line-uc0-cnt", types.YLeaf{"InFullLineUc0Cnt", peUcStats.InFullLineUc0Cnt})
    peUcStats.EntityData.Leafs.Append("in-full-line-uc1-cnt", types.YLeaf{"InFullLineUc1Cnt", peUcStats.InFullLineUc1Cnt})
    peUcStats.EntityData.Leafs.Append("in-full-line-uc2-cnt", types.YLeaf{"InFullLineUc2Cnt", peUcStats.InFullLineUc2Cnt})
    peUcStats.EntityData.Leafs.Append("pkt-trunc-eop-uc0-cnt", types.YLeaf{"PktTruncEopUc0Cnt", peUcStats.PktTruncEopUc0Cnt})
    peUcStats.EntityData.Leafs.Append("pkt-trunc-eop-uc1-cnt", types.YLeaf{"PktTruncEopUc1Cnt", peUcStats.PktTruncEopUc1Cnt})
    peUcStats.EntityData.Leafs.Append("pkt-trunc-eop-uc2-cnt", types.YLeaf{"PktTruncEopUc2Cnt", peUcStats.PktTruncEopUc2Cnt})
    peUcStats.EntityData.Leafs.Append("pkt-sop-drop-uc0-cnt", types.YLeaf{"PktSopDropUc0Cnt", peUcStats.PktSopDropUc0Cnt})
    peUcStats.EntityData.Leafs.Append("pkt-sop-drop-uc1-cnt", types.YLeaf{"PktSopDropUc1Cnt", peUcStats.PktSopDropUc1Cnt})
    peUcStats.EntityData.Leafs.Append("pkt-sop-drop-uc2-cnt", types.YLeaf{"PktSopDropUc2Cnt", peUcStats.PktSopDropUc2Cnt})
    peUcStats.EntityData.Leafs.Append("pkt-ecc-err-drop-uc-cnt", types.YLeaf{"PktEccErrDropUcCnt", peUcStats.PktEccErrDropUcCnt})
    peUcStats.EntityData.Leafs.Append("pkt-ecc-trunc-cnt-uc-cnt", types.YLeaf{"PktEccTruncCntUcCnt", peUcStats.PktEccTruncCntUcCnt})
    peUcStats.EntityData.Leafs.Append("ecc-1bit-err-uc0-0-cnt", types.YLeaf{"Ecc1bitErrUc00Cnt", peUcStats.Ecc1bitErrUc00Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-1bit-err-uc0-1-cnt", types.YLeaf{"Ecc1bitErrUc01Cnt", peUcStats.Ecc1bitErrUc01Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-1bit-err-uc1-0-cnt", types.YLeaf{"Ecc1bitErrUc10Cnt", peUcStats.Ecc1bitErrUc10Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-1bit-err-uc1-1-cnt", types.YLeaf{"Ecc1bitErrUc11Cnt", peUcStats.Ecc1bitErrUc11Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-1bit-err-uc2-0-cnt", types.YLeaf{"Ecc1bitErrUc20Cnt", peUcStats.Ecc1bitErrUc20Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-1bit-err-uc2-1-cnt", types.YLeaf{"Ecc1bitErrUc21Cnt", peUcStats.Ecc1bitErrUc21Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-2bit-err-uc0-0-cnt", types.YLeaf{"Ecc2bitErrUc00Cnt", peUcStats.Ecc2bitErrUc00Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-2bit-err-uc0-1-cnt", types.YLeaf{"Ecc2bitErrUc01Cnt", peUcStats.Ecc2bitErrUc01Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-2bit-err-uc1-0-cnt", types.YLeaf{"Ecc2bitErrUc10Cnt", peUcStats.Ecc2bitErrUc10Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-2bit-err-uc1-1-cnt", types.YLeaf{"Ecc2bitErrUc11Cnt", peUcStats.Ecc2bitErrUc11Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-2bit-err-uc2-0-cnt", types.YLeaf{"Ecc2bitErrUc20Cnt", peUcStats.Ecc2bitErrUc20Cnt})
    peUcStats.EntityData.Leafs.Append("ecc-2bit-err-uc2-1-cnt", types.YLeaf{"Ecc2bitErrUc21Cnt", peUcStats.Ecc2bitErrUc21Cnt})
    peUcStats.EntityData.Leafs.Append("out-pkt-uc-cnt", types.YLeaf{"OutPktUcCnt", peUcStats.OutPktUcCnt})
    peUcStats.EntityData.Leafs.Append("fe-uc-sop-eop-pack-cnt", types.YLeaf{"FeUcSopEopPackCnt", peUcStats.FeUcSopEopPackCnt})
    peUcStats.EntityData.Leafs.Append("fc-uc-0-1-trans-cnt", types.YLeaf{"FcUc01TransCnt", peUcStats.FcUc01TransCnt})
    peUcStats.EntityData.Leafs.Append("rate-cnt", types.YLeaf{"RateCnt", peUcStats.RateCnt})
    peUcStats.EntityData.Leafs.Append("calc-rate", types.YLeaf{"CalcRate", peUcStats.CalcRate})

    peUcStats.EntityData.YListKeys = []string {}

    return &(peUcStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats
// pe mc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IN PKT MC CNT. The type is interface{} with range: 0..18446744073709551615.
    InPktMcCnt interface{}

    // IN FULL LINE MC CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    InFullLineMcCnt interface{}

    // PKT TRUNC EOP MC CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktTruncEopMcCnt interface{}

    // PKT SOP DROP MC CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktSopDropMcCnt interface{}

    // PKT ECC ERR DROP MC CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktEccErrDropMcCnt interface{}

    // PKT ECC ERR TRUNC CNT MC CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    PktEccErrTruncCntMcCnt interface{}

    // ECC 1BIT ERR MC0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1bitErrMc0Cnt interface{}

    // ECC 1BIT ERR MC1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1bitErrMc1Cnt interface{}

    // ECC 1BIT ERR MC2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc1bitErrMc2Cnt interface{}

    // ECC 2BIT ERR MC0 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2bitErrMc0Cnt interface{}

    // ECC 2BIT ERR MC1 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2bitErrMc1Cnt interface{}

    // ECC 2BIT ERR MC2 CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Ecc2bitErrMc2Cnt interface{}

    // OUT PKT MC CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktMcCnt interface{}

    // FE MC SOP EOP PACK CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FeMcSopEopPackCnt interface{}

    // FC MC 0 1 TRANS CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FcMc01TransCnt interface{}

    // RATE CNT. The type is interface{} with range: 0..18446744073709551615.
    RateCnt interface{}

    // calc rate. The type is interface{} with range: 0..18446744073709551615.
    CalcRate interface{}
}

func (peMcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeMcStats) GetEntityData() *types.CommonEntityData {
    peMcStats.EntityData.YFilter = peMcStats.YFilter
    peMcStats.EntityData.YangName = "pe-mc-stats"
    peMcStats.EntityData.BundleName = "cisco_ios_xr"
    peMcStats.EntityData.ParentYangName = "sm15-stat"
    peMcStats.EntityData.SegmentPath = "pe-mc-stats"
    peMcStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peMcStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peMcStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peMcStats.EntityData.Children = types.NewOrderedMap()
    peMcStats.EntityData.Leafs = types.NewOrderedMap()
    peMcStats.EntityData.Leafs.Append("in-pkt-mc-cnt", types.YLeaf{"InPktMcCnt", peMcStats.InPktMcCnt})
    peMcStats.EntityData.Leafs.Append("in-full-line-mc-cnt", types.YLeaf{"InFullLineMcCnt", peMcStats.InFullLineMcCnt})
    peMcStats.EntityData.Leafs.Append("pkt-trunc-eop-mc-cnt", types.YLeaf{"PktTruncEopMcCnt", peMcStats.PktTruncEopMcCnt})
    peMcStats.EntityData.Leafs.Append("pkt-sop-drop-mc-cnt", types.YLeaf{"PktSopDropMcCnt", peMcStats.PktSopDropMcCnt})
    peMcStats.EntityData.Leafs.Append("pkt-ecc-err-drop-mc-cnt", types.YLeaf{"PktEccErrDropMcCnt", peMcStats.PktEccErrDropMcCnt})
    peMcStats.EntityData.Leafs.Append("pkt-ecc-err-trunc-cnt-mc-cnt", types.YLeaf{"PktEccErrTruncCntMcCnt", peMcStats.PktEccErrTruncCntMcCnt})
    peMcStats.EntityData.Leafs.Append("ecc-1bit-err-mc0-cnt", types.YLeaf{"Ecc1bitErrMc0Cnt", peMcStats.Ecc1bitErrMc0Cnt})
    peMcStats.EntityData.Leafs.Append("ecc-1bit-err-mc1-cnt", types.YLeaf{"Ecc1bitErrMc1Cnt", peMcStats.Ecc1bitErrMc1Cnt})
    peMcStats.EntityData.Leafs.Append("ecc-1bit-err-mc2-cnt", types.YLeaf{"Ecc1bitErrMc2Cnt", peMcStats.Ecc1bitErrMc2Cnt})
    peMcStats.EntityData.Leafs.Append("ecc-2bit-err-mc0-cnt", types.YLeaf{"Ecc2bitErrMc0Cnt", peMcStats.Ecc2bitErrMc0Cnt})
    peMcStats.EntityData.Leafs.Append("ecc-2bit-err-mc1-cnt", types.YLeaf{"Ecc2bitErrMc1Cnt", peMcStats.Ecc2bitErrMc1Cnt})
    peMcStats.EntityData.Leafs.Append("ecc-2bit-err-mc2-cnt", types.YLeaf{"Ecc2bitErrMc2Cnt", peMcStats.Ecc2bitErrMc2Cnt})
    peMcStats.EntityData.Leafs.Append("out-pkt-mc-cnt", types.YLeaf{"OutPktMcCnt", peMcStats.OutPktMcCnt})
    peMcStats.EntityData.Leafs.Append("fe-mc-sop-eop-pack-cnt", types.YLeaf{"FeMcSopEopPackCnt", peMcStats.FeMcSopEopPackCnt})
    peMcStats.EntityData.Leafs.Append("fc-mc-0-1-trans-cnt", types.YLeaf{"FcMc01TransCnt", peMcStats.FcMc01TransCnt})
    peMcStats.EntityData.Leafs.Append("rate-cnt", types.YLeaf{"RateCnt", peMcStats.RateCnt})
    peMcStats.EntityData.Leafs.Append("calc-rate", types.YLeaf{"CalcRate", peMcStats.CalcRate})

    peMcStats.EntityData.YListKeys = []string {}

    return &(peMcStats.EntityData)
}

// CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats
// pe cc stats
type CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IN PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    InPktCnt interface{}

    // OUT PATH0 PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPath0PktCnt interface{}

    // OUT PATH1 PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPath1PktCnt interface{}

    // XBAR ECC DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    XbarEccDropPktCnt interface{}

    // MEM0 DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Mem0DropPktCnt interface{}

    // MEM1 DROP PKT CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Mem1DropPktCnt interface{}

    // CONGN PKT CNT. The type is interface{} with range: 0..18446744073709551615.
    CongnPktCnt interface{}

    // XBAR ECC SINGLE ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    XbarEccSingleErrCnt interface{}

    // XBAR ECC DOUBLE ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    XbarEccDoubleErrCnt interface{}

    // MEM0 ECC SINGLE ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Mem0EccSingleErrCnt interface{}

    // MEM0 ECC DOUBLE ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Mem0EccDoubleErrCnt interface{}

    // MEM1 ECC SINGLE ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Mem1EccSingleErrCnt interface{}

    // MEM1 ECC DOUBLE ERR CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    Mem1EccDoubleErrCnt interface{}

    // FC CC 0 1 TRANS CNT. The type is interface{} with range:
    // 0..18446744073709551615.
    FcCc01TransCnt interface{}

    // RATE CNT. The type is interface{} with range: 0..18446744073709551615.
    RateCnt interface{}

    // calc rate. The type is interface{} with range: 0..18446744073709551615.
    CalcRate interface{}
}

func (peCcStats *CrossBarStats_Nodes_Node_CrossBarTable_Sm15Stats_Sm15Stat_PeCcStats) GetEntityData() *types.CommonEntityData {
    peCcStats.EntityData.YFilter = peCcStats.YFilter
    peCcStats.EntityData.YangName = "pe-cc-stats"
    peCcStats.EntityData.BundleName = "cisco_ios_xr"
    peCcStats.EntityData.ParentYangName = "sm15-stat"
    peCcStats.EntityData.SegmentPath = "pe-cc-stats"
    peCcStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peCcStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peCcStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peCcStats.EntityData.Children = types.NewOrderedMap()
    peCcStats.EntityData.Leafs = types.NewOrderedMap()
    peCcStats.EntityData.Leafs.Append("in-pkt-cnt", types.YLeaf{"InPktCnt", peCcStats.InPktCnt})
    peCcStats.EntityData.Leafs.Append("out-path0-pkt-cnt", types.YLeaf{"OutPath0PktCnt", peCcStats.OutPath0PktCnt})
    peCcStats.EntityData.Leafs.Append("out-path1-pkt-cnt", types.YLeaf{"OutPath1PktCnt", peCcStats.OutPath1PktCnt})
    peCcStats.EntityData.Leafs.Append("xbar-ecc-drop-pkt-cnt", types.YLeaf{"XbarEccDropPktCnt", peCcStats.XbarEccDropPktCnt})
    peCcStats.EntityData.Leafs.Append("mem0-drop-pkt-cnt", types.YLeaf{"Mem0DropPktCnt", peCcStats.Mem0DropPktCnt})
    peCcStats.EntityData.Leafs.Append("mem1-drop-pkt-cnt", types.YLeaf{"Mem1DropPktCnt", peCcStats.Mem1DropPktCnt})
    peCcStats.EntityData.Leafs.Append("congn-pkt-cnt", types.YLeaf{"CongnPktCnt", peCcStats.CongnPktCnt})
    peCcStats.EntityData.Leafs.Append("xbar-ecc-single-err-cnt", types.YLeaf{"XbarEccSingleErrCnt", peCcStats.XbarEccSingleErrCnt})
    peCcStats.EntityData.Leafs.Append("xbar-ecc-double-err-cnt", types.YLeaf{"XbarEccDoubleErrCnt", peCcStats.XbarEccDoubleErrCnt})
    peCcStats.EntityData.Leafs.Append("mem0-ecc-single-err-cnt", types.YLeaf{"Mem0EccSingleErrCnt", peCcStats.Mem0EccSingleErrCnt})
    peCcStats.EntityData.Leafs.Append("mem0-ecc-double-err-cnt", types.YLeaf{"Mem0EccDoubleErrCnt", peCcStats.Mem0EccDoubleErrCnt})
    peCcStats.EntityData.Leafs.Append("mem1-ecc-single-err-cnt", types.YLeaf{"Mem1EccSingleErrCnt", peCcStats.Mem1EccSingleErrCnt})
    peCcStats.EntityData.Leafs.Append("mem1-ecc-double-err-cnt", types.YLeaf{"Mem1EccDoubleErrCnt", peCcStats.Mem1EccDoubleErrCnt})
    peCcStats.EntityData.Leafs.Append("fc-cc-0-1-trans-cnt", types.YLeaf{"FcCc01TransCnt", peCcStats.FcCc01TransCnt})
    peCcStats.EntityData.Leafs.Append("rate-cnt", types.YLeaf{"RateCnt", peCcStats.RateCnt})
    peCcStats.EntityData.Leafs.Append("calc-rate", types.YLeaf{"CalcRate", peCcStats.CalcRate})

    peCcStats.EntityData.YListKeys = []string {}

    return &(peCcStats.EntityData)
}

